package policy

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/policy"
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/policy/parse"
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
					string(policy.TypeBuiltIn),
					string(policy.TypeCustom),
				}, false),
			},

			"management_group_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Computed:      true,
				ConflictsWith: []string{"management_group_name"},
				Deprecated:    "Deprecated in favour of `management_group_name`", // TODO -- remove this in next major version
			},

			"management_group_name": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Computed:      true, // TODO -- remove this when deprecation resolves
				ConflictsWith: []string{"management_group_id"},
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

			"policy_definitions": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: policyDefinitionsDiffSuppressFunc,
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

func policyDefinitionsDiffSuppressFunc(_, old, new string, _ *schema.ResourceData) bool {
	var oldPolicyDefinitions []policy.DefinitionReference
	errOld := json.Unmarshal([]byte(old), &oldPolicyDefinitions)
	if errOld != nil {
		return false
	}

	var newPolicyDefinitions []policy.DefinitionReference
	errNew := json.Unmarshal([]byte(new), &newPolicyDefinitions)
	if errNew != nil {
		return false
	}

	return reflect.DeepEqual(oldPolicyDefinitions, newPolicyDefinitions)
}

func resourceArmPolicySetDefinitionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Policy.SetDefinitionsClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	policyType := d.Get("policy_type").(string)
	displayName := d.Get("display_name").(string)
	description := d.Get("description").(string)
	managementGroupName := ""
	if v, ok := d.GetOk("management_group_name"); ok {
		managementGroupName = v.(string)
	}
	if v, ok := d.GetOk("management_group_id"); ok {
		managementGroupName = v.(string)
	}

	if features.ShouldResourcesBeImported() && d.IsNewResource() {
		existing, err := getPolicySetDefinitionByName(ctx, client, name, managementGroupName)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("Error checking for presence of existing Policy Set Definition %q: %s", name, err)
			}
		}

		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_policy_set_definition", *existing.ID)
		}
	}

	properties := policy.SetDefinitionProperties{
		PolicyType:  policy.Type(policyType),
		DisplayName: utils.String(displayName),
		Description: utils.String(description),
	}

	if metaDataString := d.Get("metadata").(string); metaDataString != "" {
		metaData, err := structure.ExpandJsonFromString(metaDataString)
		if err != nil {
			return fmt.Errorf("unable to expand metadata json: %s", err)
		}
		properties.Metadata = &metaData
	}

	if parametersString := d.Get("parameters").(string); parametersString != "" {
		parameters, err := structure.ExpandJsonFromString(parametersString)
		if err != nil {
			return fmt.Errorf("unable to expand parameters json: %s", err)
		}
		properties.Parameters = &parameters
	}

	if policyDefinitionsString := d.Get("policy_definitions").(string); policyDefinitionsString != "" {
		var policyDefinitions []policy.DefinitionReference
		err := json.Unmarshal([]byte(policyDefinitionsString), &policyDefinitions)
		if err != nil {
			return fmt.Errorf("unable to expand parameters json: %s", err)
		}
		properties.PolicyDefinitions = &policyDefinitions
	}

	definition := policy.SetDefinition{
		Name:                    utils.String(name),
		SetDefinitionProperties: &properties,
	}

	var err error
	if managementGroupName == "" {
		_, err = client.CreateOrUpdate(ctx, name, definition)
	} else {
		_, err = client.CreateOrUpdateAtManagementGroup(ctx, name, definition, managementGroupName)
	}

	if err != nil {
		return fmt.Errorf("Error creating/updating Policy Set Definition %q: %s", name, err)
	}

	// Policy Definitions are eventually consistent; wait for them to stabilize
	log.Printf("[DEBUG] Waiting for Policy Set Definition %q to become available", name)
	stateConf := &resource.StateChangeConf{
		Pending:                   []string{"404"},
		Target:                    []string{"200"},
		Refresh:                   policySetDefinitionRefreshFunc(ctx, client, name, managementGroupName),
		MinTimeout:                10 * time.Second,
		ContinuousTargetOccurence: 10,
	}

	if d.IsNewResource() {
		stateConf.Timeout = d.Timeout(schema.TimeoutCreate)
	} else {
		stateConf.Timeout = d.Timeout(schema.TimeoutUpdate)
	}

	if _, err = stateConf.WaitForState(); err != nil {
		return fmt.Errorf("Error waiting for Policy Set Definition %q to become available: %s", name, err)
	}

	var resp policy.SetDefinition
	resp, err = getPolicySetDefinitionByName(ctx, client, name, managementGroupName)
	if err != nil {
		return fmt.Errorf("Error retrieving Policy Set Definition %q: %s", name, err)
	}

	d.SetId(*resp.ID)

	return resourceArmPolicySetDefinitionRead(d, meta)
}

func resourceArmPolicySetDefinitionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Policy.SetDefinitionsClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.PolicySetDefinitionID(d.Id())
	if err != nil {
		return err
	}

	managementGroupName := ""
	switch scopeId := id.PolicyScopeId.(type) { // nolint gocritic
	case parse.ScopeAtManagementGroup:
		managementGroupName = scopeId.ManagementGroupName
	}

	resp, err := getPolicySetDefinitionByName(ctx, client, id.Name, managementGroupName)

	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[INFO] Error reading Policy Set Definition %q - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error reading Policy Set Definition %+v", err)
	}

	d.Set("name", resp.Name)
	d.Set("management_group_id", managementGroupName)
	d.Set("management_group_name", managementGroupName)

	if props := resp.SetDefinitionProperties; props != nil {
		d.Set("policy_type", string(props.PolicyType))
		d.Set("display_name", props.DisplayName)
		d.Set("description", props.Description)

		if metadata := props.Metadata; metadata != nil {
			metadataVal := metadata.(map[string]interface{})
			metadataStr, err := structure.FlattenJsonToString(metadataVal)
			if err != nil {
				return fmt.Errorf("unable to flatten JSON for `metadata`: %s", err)
			}

			d.Set("metadata", metadataStr)
		}

		if parameters := props.Parameters; parameters != nil {
			paramsVal := parameters.(map[string]interface{})
			parametersStr, err := structure.FlattenJsonToString(paramsVal)
			if err != nil {
				return fmt.Errorf("unable to flatten JSON for `parameters`: %s", err)
			}

			d.Set("parameters", parametersStr)
		}

		if policyDefinitions := props.PolicyDefinitions; policyDefinitions != nil {
			policyDefinitionsRes, err := json.Marshal(policyDefinitions)
			if err != nil {
				return fmt.Errorf("unable to flatten JSON for `policy_defintions`: %s", err)
			}

			policyDefinitionsStr := string(policyDefinitionsRes)
			d.Set("policy_definitions", policyDefinitionsStr)
		}
	}

	return nil
}

func resourceArmPolicySetDefinitionDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Policy.SetDefinitionsClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.PolicySetDefinitionID(d.Id())
	if err != nil {
		return err
	}

	managementGroupName := ""
	switch scopeId := id.PolicyScopeId.(type) { // nolint gocritic
	case parse.ScopeAtManagementGroup:
		managementGroupName = scopeId.ManagementGroupName
	}

	var resp autorest.Response
	if managementGroupName == "" {
		resp, err = client.Delete(ctx, id.Name)
	} else {
		resp, err = client.DeleteAtManagementGroup(ctx, id.Name, managementGroupName)
	}

	if err != nil {
		if utils.ResponseWasNotFound(resp) {
			return nil
		}

		return fmt.Errorf("Error deleting Policy Set Definition %q: %+v", id.Name, err)
	}

	return nil
}

func policySetDefinitionRefreshFunc(ctx context.Context, client *policy.SetDefinitionsClient, name string, managementGroupId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		res, err := getPolicySetDefinitionByName(ctx, client, name, managementGroupId)
		if err != nil {
			return nil, strconv.Itoa(res.StatusCode), fmt.Errorf("Error issuing read request in policySetDefinitionRefreshFunc for Policy Set Definition %q: %s", name, err)
		}

		return res, strconv.Itoa(res.StatusCode), nil
	}
}
