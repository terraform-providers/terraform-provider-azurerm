package policy

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-09-01/policy"
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/policy/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/policy/validate"
	azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmPolicySetDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmPolicySetDefinitionCreateUpdate,
		Update: resourceArmPolicySetDefinitionCreateUpdate,
		Read:   resourceArmPolicySetDefinitionRead,
		Delete: resourceArmPolicySetDefinitionDelete,

		Importer: azSchema.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.PolicySetDefinitionID(id)
			return err
		}),

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"policy_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(policy.BuiltIn),
					string(policy.Custom),
					string(policy.NotSpecified),
					string(policy.Static),
				}, false),
			},

			"management_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"display_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"metadata": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: policySetDefinitionsMetadataDiffSuppressFunc,
			},

			"parameters": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},

			"policy_definitions": { // TODO -- remove in the next major version
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: policyDefinitionsDiffSuppressFunc,
				ExactlyOneOf:     []string{"policy_definitions", "policy_definition_reference"},
				Deprecated:       "Deprecated in favor of `policy_definition_reference`",
			},

			"policy_definition_reference": { // TODO -- rename this back to `policy_definition` after the deprecation
				Type:         schema.TypeList,
				Optional:     true,                                                          // TODO -- change this to Required after the deprecation
				Computed:     true,                                                          // TODO -- remove Computed after the deprecation
				ExactlyOneOf: []string{"policy_definitions", "policy_definition_reference"}, // TODO -- remove after the deprecation
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_definition_id": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validate.PolicyDefinitionID,
						},

						"parameters": {
							Type:     schema.TypeMap,
							Optional: true,
						},

						"reference_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func policySetDefinitionsMetadataDiffSuppressFunc(_, old, new string, _ *schema.ResourceData) bool {
	var oldPolicySetDefinitionsMetadata map[string]interface{}
	errOld := json.Unmarshal([]byte(old), &oldPolicySetDefinitionsMetadata)
	if errOld != nil {
		return false
	}

	var newPolicySetDefinitionsMetadata map[string]interface{}
	errNew := json.Unmarshal([]byte(new), &newPolicySetDefinitionsMetadata)
	if errNew != nil {
		return false
	}

	// Ignore the following keys if they're found in the metadata JSON
	ignoreKeys := [4]string{"createdBy", "createdOn", "updatedBy", "updatedOn"}
	for _, key := range ignoreKeys {
		delete(oldPolicySetDefinitionsMetadata, key)
		delete(newPolicySetDefinitionsMetadata, key)
	}

	return reflect.DeepEqual(oldPolicySetDefinitionsMetadata, newPolicySetDefinitionsMetadata)
}

// This function only serves the deprecated attribute `policy_definitions` in the old api-version.
// The old api-version only support two attribute - `policy_definition_id` and `parameters` in each element.
// Therefore this function is used for ignoring any other keys and then compare if there is a diff
func policyDefinitionsDiffSuppressFunc(_, old, new string, _ *schema.ResourceData) bool {
	var oldPolicyDefinitions []DefinitionReferenceInOldApiVersion
	errOld := json.Unmarshal([]byte(old), &oldPolicyDefinitions)
	if errOld != nil {
		return false
	}

	var newPolicyDefinitions []DefinitionReferenceInOldApiVersion
	errNew := json.Unmarshal([]byte(new), &newPolicyDefinitions)
	if errNew != nil {
		return false
	}

	return reflect.DeepEqual(oldPolicyDefinitions, newPolicyDefinitions)
}

type DefinitionReferenceInOldApiVersion struct {
	// PolicyDefinitionID - The ID of the policy definition or policy set definition.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty"`
	// Parameters - The parameter values for the referenced policy rule. The keys are the parameter names.
	Parameters map[string]*policy.ParameterValuesValue `json:"parameters"`
}

func resourceArmPolicySetDefinitionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Policy.SetDefinitionsClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	managementGroupID := d.Get("management_group_id").(string)

	if d.IsNewResource() {
		existing, err := getPolicySetDefinitionByName(ctx, client, name, managementGroupID)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("checking for presence of existing Policy Set Definition %q: %+v", name, err)
			}
		}

		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_policy_set_definition", *existing.ID)
		}
	}

	properties := policy.SetDefinitionProperties{
		PolicyType:  policy.Type(d.Get("policy_type").(string)),
		DisplayName: utils.String(d.Get("display_name").(string)),
		Description: utils.String(d.Get("description").(string)),
	}

	if metaDataString := d.Get("metadata").(string); metaDataString != "" {
		metaData, err := structure.ExpandJsonFromString(metaDataString)
		if err != nil {
			return fmt.Errorf("expanding JSON for `metadata`: %+v", err)
		}
		properties.Metadata = &metaData
	}

	if parametersString := d.Get("parameters").(string); parametersString != "" {
		parameters, err := expandParameterDefinitionsValueFromString(parametersString)
		if err != nil {
			return fmt.Errorf("expanding JSON for `parameters`: %+v", err)
		}
		properties.Parameters = parameters
	}

	if v, ok := d.GetOk("policy_definitions"); ok {
		var policyDefinitions []policy.DefinitionReference
		err := json.Unmarshal([]byte(v.(string)), &policyDefinitions)
		if err != nil {
			return fmt.Errorf("expanding JSON for `policy_definitions`: %+v", err)
		}
		properties.PolicyDefinitions = &policyDefinitions
	}
	if v, ok := d.GetOk("policy_definition_reference"); ok {
		properties.PolicyDefinitions = expandAzureRMPolicySetDefinitionPolicyDefinitions(v.([]interface{}))
	}

	definition := policy.SetDefinition{
		Name:                    utils.String(name),
		SetDefinitionProperties: &properties,
	}

	var err error
	if managementGroupID == "" {
		_, err = client.CreateOrUpdate(ctx, name, definition)
	} else {
		_, err = client.CreateOrUpdateAtManagementGroup(ctx, name, definition, managementGroupID)
	}

	if err != nil {
		return fmt.Errorf("creating/updating Policy Set Definition %q: %+v", name, err)
	}

	// Policy Definitions are eventually consistent; wait for them to stabilize
	log.Printf("[DEBUG] Waiting for Policy Set Definition %q to become available", name)
	stateConf := &resource.StateChangeConf{
		Pending:                   []string{"404"},
		Target:                    []string{"200"},
		Refresh:                   policySetDefinitionRefreshFunc(ctx, client, name, managementGroupID),
		MinTimeout:                10 * time.Second,
		ContinuousTargetOccurence: 10,
	}

	if d.IsNewResource() {
		stateConf.Timeout = d.Timeout(schema.TimeoutCreate)
	} else {
		stateConf.Timeout = d.Timeout(schema.TimeoutUpdate)
	}

	if _, err = stateConf.WaitForState(); err != nil {
		return fmt.Errorf("waiting for Policy Set Definition %q to become available: %+v", name, err)
	}

	var resp policy.SetDefinition
	resp, err = getPolicySetDefinitionByName(ctx, client, name, managementGroupID)
	if err != nil {
		return fmt.Errorf("retrieving Policy Set Definition %q: %+v", name, err)
	}

	d.SetId(*resp.ID)

	return resourceArmPolicySetDefinitionRead(d, meta)
}

func resourceArmPolicySetDefinitionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Policy.SetDefinitionsClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name, err := parsePolicySetDefinitionNameFromId(d.Id())
	if err != nil {
		return err
	}

	managementGroupID := parseManagementGroupIdFromPolicySetId(d.Id())

	resp, err := getPolicySetDefinitionByName(ctx, client, name, managementGroupID)

	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[INFO] Error reading Policy Set Definition %q - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf("reading Policy Set Definition %+v", err)
	}

	d.Set("name", resp.Name)
	d.Set("management_group_id", managementGroupID)

	if props := resp.SetDefinitionProperties; props != nil {
		d.Set("policy_type", string(props.PolicyType))
		d.Set("display_name", props.DisplayName)
		d.Set("description", props.Description)

		if metadata := props.Metadata; metadata != nil {
			metadataVal := metadata.(map[string]interface{})
			metadataStr, err := structure.FlattenJsonToString(metadataVal)
			if err != nil {
				return fmt.Errorf("flattening JSON for `metadata`: %+v", err)
			}

			d.Set("metadata", metadataStr)
		}

		if parameters := props.Parameters; parameters != nil {
			parametersStr, err := flattenParameterDefintionsValueToString(parameters)
			if err != nil {
				return fmt.Errorf("flattening JSON for `parameters`: %+v", err)
			}

			d.Set("parameters", parametersStr)
		}

		if policyDefinitions := props.PolicyDefinitions; policyDefinitions != nil {
			policyDefinitionsRes, err := json.Marshal(policyDefinitions)
			if err != nil {
				return fmt.Errorf("flattening JSON for `policy_defintions`: %+v", err)
			}

			d.Set("policy_definitions", string(policyDefinitionsRes))
		}
		if err := d.Set("policy_definition_reference", flattenAzureRMPolicySetDefinitionPolicyDefinitions(props.PolicyDefinitions)); err != nil {
			return fmt.Errorf("setting `policy_definition_reference`: %+v", err)
		}
	}

	return nil
}

func resourceArmPolicySetDefinitionDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Policy.SetDefinitionsClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name, err := parsePolicySetDefinitionNameFromId(d.Id())
	if err != nil {
		return err
	}

	managementGroupID := parseManagementGroupIdFromPolicySetId(d.Id())

	var resp autorest.Response
	if managementGroupID == "" {
		resp, err = client.Delete(ctx, name)
	} else {
		resp, err = client.DeleteAtManagementGroup(ctx, name, managementGroupID)
	}

	if err != nil {
		if utils.ResponseWasNotFound(resp) {
			return nil
		}

		return fmt.Errorf("deleting Policy Set Definition %q: %+v", name, err)
	}

	return nil
}

func parsePolicySetDefinitionNameFromId(id string) (string, error) {
	components := strings.Split(id, "/")

	if len(components) == 0 {
		return "", fmt.Errorf("Azure Policy Set Definition Id is empty or not formatted correctly: %+v", id)
	}

	return components[len(components)-1], nil
}

func parseManagementGroupIdFromPolicySetId(id string) string {
	r := regexp.MustCompile("managementgroups/(.+)/providers/.*$")

	if r.MatchString(id) {
		matches := r.FindAllStringSubmatch(id, -1)[0]
		return matches[1]
	}

	return ""
}

func policySetDefinitionRefreshFunc(ctx context.Context, client *policy.SetDefinitionsClient, name string, managementGroupId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		res, err := getPolicySetDefinitionByName(ctx, client, name, managementGroupId)
		if err != nil {
			return nil, strconv.Itoa(res.StatusCode), fmt.Errorf("issuing read request in policySetDefinitionRefreshFunc for Policy Set Definition %q: %+v", name, err)
		}

		return res, strconv.Itoa(res.StatusCode), nil
	}
}

func expandAzureRMPolicySetDefinitionPolicyDefinitions(input []interface{}) *[]policy.DefinitionReference {
	result := make([]policy.DefinitionReference, 0)

	for _, item := range input {
		v := item.(map[string]interface{})

		parameters := make(map[string]*policy.ParameterValuesValue)
		for k, value := range v["parameters"].(map[string]interface{}) {
			parameters[k] = &policy.ParameterValuesValue{
				Value: value.(string),
			}
		}

		result = append(result, policy.DefinitionReference{
			PolicyDefinitionID:          utils.String(v["policy_definition_id"].(string)),
			Parameters:                  parameters,
			PolicyDefinitionReferenceID: utils.String(v["reference_id"].(string)),
		})
	}

	return &result
}

func flattenAzureRMPolicySetDefinitionPolicyDefinitions(input *[]policy.DefinitionReference) []interface{} {
	result := make([]interface{}, 0)
	if input == nil {
		return result
	}

	for _, definition := range *input {
		policyDefinitionID := ""
		if definition.PolicyDefinitionID != nil {
			policyDefinitionID = *definition.PolicyDefinitionID
		}

		parametersMap := make(map[string]interface{})
		for k, v := range definition.Parameters {
			if v == nil {
				continue
			}
			parametersMap[k] = v.Value
		}

		policyDefinitionReference := ""
		if definition.PolicyDefinitionReferenceID != nil {
			policyDefinitionReference = *definition.PolicyDefinitionReferenceID
		}

		result = append(result, map[string]interface{}{
			"policy_definition_id": policyDefinitionID,
			"parameters":           parametersMap,
			"reference_id":         policyDefinitionReference,
		})
	}
	return result
}
