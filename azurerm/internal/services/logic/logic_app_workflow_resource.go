package logic

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2019-05-01/logic"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/locks"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/logic/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tags"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

var logicAppResourceName = "azurerm_logic_app"

func resourceLogicAppWorkflow() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Create: resourceLogicAppWorkflowCreate,
		Read:   resourceLogicAppWorkflowRead,
		Update: resourceLogicAppWorkflowUpdate,
		Delete: resourceLogicAppWorkflowDelete,
		// TODO: replace this with an importer which validates the ID during import
		Importer: pluginsdk.DefaultImporter(),

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(30 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(30 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"name": {
				Type:     pluginsdk.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.All(
					validation.StringIsNotEmpty,
					validation.StringMatch(
						regexp.MustCompile("^[-()_.A-Za-z0-9]{1,80}$"),
						"The Logic app name can contain only letters, numbers, periods (.), hyphens (-), brackets (()) and underscores (_), up to 80 characters",
					),
				),
			},

			"location": azure.SchemaLocation(),

			"resource_group_name": azure.SchemaResourceGroupName(),

			"integration_service_environment_id": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validate.IntegrationServiceEnvironmentID,
			},

			"logic_app_integration_account_id": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validate.IntegrationAccountID,
			},

			// TODO: should Parameters be split out into their own object to allow validation on the different sub-types?
			"parameters": {
				Type:     pluginsdk.TypeMap,
				Optional: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"workflow_schema": {
				Type:     pluginsdk.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#",
			},

			"workflow_version": {
				Type:     pluginsdk.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "1.0.0.0",
			},

			"workflow_parameters": {
				Type:     pluginsdk.TypeMap,
				Optional: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"access_endpoint": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"connector_endpoint_ip_addresses": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem:     &pluginsdk.Schema{Type: pluginsdk.TypeString},
			},
			"connector_outbound_ip_addresses": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem:     &pluginsdk.Schema{Type: pluginsdk.TypeString},
			},
			"workflow_endpoint_ip_addresses": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem:     &pluginsdk.Schema{Type: pluginsdk.TypeString},
			},
			"workflow_outbound_ip_addresses": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem:     &pluginsdk.Schema{Type: pluginsdk.TypeString},
			},

			"tags": tags.Schema(),
		},
	}
}

func resourceLogicAppWorkflowCreate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Logic.WorkflowClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	log.Printf("[INFO] preparing arguments for Logic App Workflow creation.")

	name := d.Get("name").(string)
	resourceGroup := d.Get("resource_group_name").(string)

	if d.IsNewResource() {
		existing, err := client.Get(ctx, resourceGroup, name)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("Error checking for presence of existing Logic App Workflow %q (Resource Group %q): %s", name, resourceGroup, err)
			}
		}

		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_logic_app_workflow", *existing.ID)
		}
	}

	location := azure.NormalizeLocation(d.Get("location").(string))

	workflowSchema := d.Get("workflow_schema").(string)
	workflowVersion := d.Get("workflow_version").(string)
	workflowParameters, err := expandLogicAppWorkflowWorkflowParameters(d.Get("workflow_parameters").(map[string]interface{}))
	if err != nil {
		return fmt.Errorf("expanding `workflow_parameters`: %+v", err)
	}

	parameters, err := expandLogicAppWorkflowParameters(d.Get("parameters").(map[string]interface{}), workflowParameters)
	if err != nil {
		return err
	}
	t := d.Get("tags").(map[string]interface{})

	properties := logic.Workflow{
		Location: utils.String(location),
		WorkflowProperties: &logic.WorkflowProperties{
			Definition: &map[string]interface{}{
				"$schema":        workflowSchema,
				"contentVersion": workflowVersion,
				"actions":        make(map[string]interface{}),
				"triggers":       make(map[string]interface{}),
				"parameters":     workflowParameters,
			},
			Parameters: parameters,
		},
		Tags: tags.Expand(t),
	}

	if iseID, ok := d.GetOk("integration_service_environment_id"); ok {
		properties.WorkflowProperties.IntegrationServiceEnvironment = &logic.ResourceReference{
			ID: utils.String(iseID.(string)),
		}
	}

	if v, ok := d.GetOk("logic_app_integration_account_id"); ok {
		properties.WorkflowProperties.IntegrationAccount = &logic.ResourceReference{
			ID: utils.String(v.(string)),
		}
	}

	if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, properties); err != nil {
		return fmt.Errorf("[ERROR] Error creating Logic App Workflow %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	read, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		return fmt.Errorf("[ERROR] Error making Read request on Logic App Workflow %q (Resource Group %q): %+v", name, resourceGroup, err)
	}
	if read.ID == nil {
		return fmt.Errorf("[ERROR] Cannot read Logic App Workflow %q (Resource Group %q) ID", name, resourceGroup)
	}

	d.SetId(*read.ID)

	return resourceLogicAppWorkflowRead(d, meta)
}

func resourceLogicAppWorkflowUpdate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Logic.WorkflowClient
	ctx, cancel := timeouts.ForUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := azure.ParseAzureResourceID(d.Id())
	if err != nil {
		return err
	}

	resourceGroup := id.ResourceGroup
	name := id.Path["workflows"]

	// lock to prevent against Actions, Parameters or Triggers conflicting
	locks.ByName(name, logicAppResourceName)
	defer locks.UnlockByName(name, logicAppResourceName)

	read, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		if utils.ResponseWasNotFound(read.Response) {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("[ERROR] Error making Read request on Logic App Workflow %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	if read.WorkflowProperties == nil {
		return fmt.Errorf("[ERROR] Error parsing Logic App Workflow - `WorkflowProperties` is nil")
	}

	location := azure.NormalizeLocation(d.Get("location").(string))
	workflowParameters, err := expandLogicAppWorkflowWorkflowParameters(d.Get("workflow_parameters").(map[string]interface{}))
	if err != nil {
		return fmt.Errorf("expanding `workflow_parameters`: %+v", err)
	}
	parameters, err := expandLogicAppWorkflowParameters(d.Get("parameters").(map[string]interface{}), workflowParameters)
	if err != nil {
		return err
	}

	t := d.Get("tags").(map[string]interface{})

	definition := read.WorkflowProperties.Definition.(map[string]interface{})
	definition["parameters"] = workflowParameters

	properties := logic.Workflow{
		Location: utils.String(location),
		WorkflowProperties: &logic.WorkflowProperties{
			Definition: definition,
			Parameters: parameters,
		},
		Tags: tags.Expand(t),
	}

	if v, ok := d.GetOk("logic_app_integration_account_id"); ok {
		properties.WorkflowProperties.IntegrationAccount = &logic.ResourceReference{
			ID: utils.String(v.(string)),
		}
	}

	if _, err = client.CreateOrUpdate(ctx, resourceGroup, name, properties); err != nil {
		return fmt.Errorf("Error updating Logic App Workspace %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	return resourceLogicAppWorkflowRead(d, meta)
}

func resourceLogicAppWorkflowRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Logic.WorkflowClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := azure.ParseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resourceGroup := id.ResourceGroup
	name := id.Path["workflows"]

	resp, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[DEBUG] Logic App Workflow %q (Resource Group %q) was not found - removing from state", name, resourceGroup)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Error making Read request on Logic App Workflow %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	d.Set("name", resp.Name)
	d.Set("resource_group_name", resourceGroup)

	if location := resp.Location; location != nil {
		d.Set("location", azure.NormalizeLocation(*location))
	}

	if props := resp.WorkflowProperties; props != nil {
		d.Set("access_endpoint", props.AccessEndpoint)

		if props.EndpointsConfiguration == nil || props.EndpointsConfiguration.Connector == nil {
			d.Set("connector_endpoint_ip_addresses", []interface{}{})
			d.Set("connector_outbound_ip_addresses", []interface{}{})
		} else {
			d.Set("connector_endpoint_ip_addresses", flattenIPAddresses(props.EndpointsConfiguration.Connector.AccessEndpointIPAddresses))
			d.Set("connector_outbound_ip_addresses", flattenIPAddresses(props.EndpointsConfiguration.Connector.OutgoingIPAddresses))
		}

		if props.EndpointsConfiguration == nil || props.EndpointsConfiguration.Workflow == nil {
			d.Set("workflow_endpoint_ip_addresses", []interface{}{})
			d.Set("workflow_outbound_ip_addresses", []interface{}{})
		} else {
			d.Set("workflow_endpoint_ip_addresses", flattenIPAddresses(props.EndpointsConfiguration.Workflow.AccessEndpointIPAddresses))
			d.Set("workflow_outbound_ip_addresses", flattenIPAddresses(props.EndpointsConfiguration.Workflow.OutgoingIPAddresses))
		}
		if definition := props.Definition; definition != nil {
			if v, ok := definition.(map[string]interface{}); ok {
				d.Set("workflow_schema", v["$schema"].(string))
				d.Set("workflow_version", v["contentVersion"].(string))
				if p, ok := v["parameters"]; ok {
					workflowParameters, err := flattenLogicAppWorkflowWorkflowParameters(p.(map[string]interface{}))
					if err != nil {
						return fmt.Errorf("flattening `workflow_parameters`: %+v", err)
					}
					if err := d.Set("workflow_parameters", workflowParameters); err != nil {
						return fmt.Errorf("setting `workflow_parameters`: %+v", err)
					}

					// The props.Parameters (the value of the param) is accompany with the "parameters" (the definition of the param) inside the props.Definition.
					// We will need to make use of the definition of the parameters in order to properly flatten the value of the parameters being set (for kinds of types).
					parameters, err := flattenLogicAppWorkflowParameters(props.Parameters, p.(map[string]interface{}))
					if err != nil {
						return fmt.Errorf("flattening `parameters`: %v", err)
					}
					if err := d.Set("parameters", parameters); err != nil {
						return fmt.Errorf("Error setting `parameters`: %+v", err)
					}
				}
			}
		}

		integrationServiceEnvironmentId := ""
		if props.IntegrationServiceEnvironment != nil && props.IntegrationServiceEnvironment.ID != nil {
			integrationServiceEnvironmentId = *props.IntegrationServiceEnvironment.ID
		}
		d.Set("integration_service_environment_id", integrationServiceEnvironmentId)

		if props.IntegrationAccount != nil && props.IntegrationAccount.ID != nil {
			d.Set("logic_app_integration_account_id", props.IntegrationAccount.ID)
		}

		integrationAccountId := ""
		if props.IntegrationAccount != nil && props.IntegrationAccount.ID != nil {
			integrationAccountId = *props.IntegrationAccount.ID
		}
		d.Set("logic_app_integration_account_id", integrationAccountId)
	}

	return tags.FlattenAndSet(d, resp.Tags)
}

func resourceLogicAppWorkflowDelete(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Logic.WorkflowClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := azure.ParseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resourceGroup := id.ResourceGroup
	name := id.Path["workflows"]

	// lock to prevent against Actions, Parameters or Triggers conflicting
	locks.ByName(name, logicAppResourceName)
	defer locks.UnlockByName(name, logicAppResourceName)

	resp, err := client.Delete(ctx, resourceGroup, name)
	if err != nil {
		if utils.ResponseWasNotFound(resp) {
			return nil
		}

		return fmt.Errorf("Error issuing delete request for Logic App Workflow %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	return nil
}

func expandLogicAppWorkflowParameters(input map[string]interface{}, paramDefs map[string]interface{}) (map[string]*logic.WorkflowParameter, error) {
	output := make(map[string]*logic.WorkflowParameter)

	for k, v := range input {
		defRaw, ok := paramDefs[k]
		if !ok {
			return nil, fmt.Errorf("no parameter definition for %s", k)
		}
		def := defRaw.(map[string]interface{})
		t := logic.ParameterType(def["type"].(string))

		v := v.(string)

		var value interface{}
		switch t {
		case logic.ParameterTypeBool:
			value = v == "true"
		case logic.ParameterTypeFloat:
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, fmt.Errorf("converting %s to float64: %v", v, err)
			}
			value = f
		case logic.ParameterTypeInt:
			i, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("converting %s to int: %v", v, err)
			}
			value = i

		case logic.ParameterTypeArray,
			logic.ParameterTypeObject,
			logic.ParameterTypeSecureObject:
			obj, err := pluginsdk.ExpandJsonFromString(v)
			if err != nil {
				return nil, fmt.Errorf("converting %s to json: %v", v, err)
			}
			value = obj

		case logic.ParameterTypeString,
			logic.ParameterTypeSecureString:
			value = v
		}

		output[k] = &logic.WorkflowParameter{
			Type:  t,
			Value: value,
		}
	}

	return output, nil
}

func flattenLogicAppWorkflowParameters(input map[string]*logic.WorkflowParameter, paramDefs map[string]interface{}) (map[string]interface{}, error) {
	output := make(map[string]interface{})

	for k, v := range input {
		defRaw, ok := paramDefs[k]
		if !ok {
			// This should never happen.
			log.Printf("[WARN] The parameter %s is not defined in the Logic App Workflow", k)
			continue
		}

		if v == nil {
			log.Printf("[WARN] The value of parameter %s is nil", k)
			continue
		}

		def := defRaw.(map[string]interface{})
		t := logic.ParameterType(def["type"].(string))

		var value string
		switch t {
		case logic.ParameterTypeBool:
			tv, ok := v.Value.(bool)
			if !ok {
				log.Printf("[WARN] The value of parameter %s is expected to be bool, but got %T", k, v.Value)
			}
			value = "true"
			if !tv {
				value = "false"
			}
		case logic.ParameterTypeFloat:
			tv, ok := v.Value.(float64)
			if !ok {
				log.Printf("[WARN] The value of parameter %s is expected to be float64, but got %T", k, v.Value)
			}
			value = strconv.FormatFloat(tv, 'f', -1, 64)
		case logic.ParameterTypeInt:
			// Note that the json unmarshalled response doesn't differ between float and int, as json has only type number.
			tv, ok := v.Value.(float64)
			if !ok {
				log.Printf("[WARN] The value of parameter %s is expected to be float64, but got %T", k, v.Value)
			}
			value = strconv.Itoa(int(tv))

		case logic.ParameterTypeArray,
			logic.ParameterTypeObject:
			tv, ok := v.Value.(map[string]interface{})
			if !ok {
				log.Printf("[WARN] The value of parameter %s is expected to be map[string]interface{}, but got %T", k, v.Value)
			}
			obj, err := pluginsdk.FlattenJsonToString(tv)
			if err != nil {
				return nil, fmt.Errorf("converting %+v from json: %v", tv, err)
			}
			value = obj

		case logic.ParameterTypeString:
			tv, ok := v.Value.(string)
			if !ok {
				log.Printf("[WARN] The value of parameter %s is expected to be string, but got %T", k, v.Value)
			}
			value = tv

		case logic.ParameterTypeSecureString,
			logic.ParameterTypeSecureObject:
			// These are not expected to return from API
			continue
		}

		output[k] = value
	}

	return output, nil
}

func expandLogicAppWorkflowWorkflowParameters(input map[string]interface{}) (map[string]interface{}, error) {
	if len(input) == 0 {
		return nil, nil
	}

	output := make(map[string]interface{})
	for k, v := range input {
		obj, err := pluginsdk.ExpandJsonFromString(v.(string))
		if err != nil {
			return nil, err
		}
		output[k] = obj
	}
	return output, nil
}

func flattenLogicAppWorkflowWorkflowParameters(input map[string]interface{}) (map[string]interface{}, error) {
	if input == nil {
		return nil, nil
	}
	output := make(map[string]interface{})
	for k, v := range input {
		objstr, err := pluginsdk.FlattenJsonToString(v.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		output[k] = objstr
	}
	return output, nil
}

func flattenIPAddresses(input *[]logic.IPAddress) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	var addresses []interface{}
	for _, addr := range *input {
		addresses = append(addresses, *addr.Address)
	}
	return addresses
}
