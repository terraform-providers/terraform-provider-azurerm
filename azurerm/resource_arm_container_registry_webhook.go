package azurerm

import (
	"fmt"
	"log"
	"regexp"

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2018-09-01/containerregistry"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/response"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmContainerRegistryWebhook() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmContainerRegistryWebhookCreate,
		Read:   resourceArmContainerRegistryWebhookRead,
		Update: resourceArmContainerRegistryWebhookUpdate,
		Delete: resourceArmContainerRegistryWebhookDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateAzureRMContainerRegistryWebhookName,
			},

			"resource_group_name": azure.SchemaResourceGroupName(),

			"registry_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateAzureRMContainerRegistryName,
			},

			"service_uri": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAzureRMContainerRegistryWebhookServiceUri,
			},

			"custom_headers": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "enabled",
				ValidateFunc: validateAzureRMContainerRegistryWebhookStatus(),
			},

			"scope": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"actions": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validateAzureRMContainerRegistryWebhookAction(),
				},
			},

			"location": azure.SchemaLocation(),

			"tags": tagsSchema(),
		},
	}
}

func resourceArmContainerRegistryWebhookCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).containers.WebhooksClient
	ctx := meta.(*ArmClient).StopContext
	log.Printf("[INFO] preparing arguments for AzureRM Container Registry Webhook creation.")

	resourceGroup := d.Get("resource_group_name").(string)
	registryName := d.Get("registry_name").(string)
	name := d.Get("name").(string)

	if requireResourcesToBeImported && d.IsNewResource() {
		existing, err := client.Get(ctx, resourceGroup, registryName, name)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("Error checking for presence of existing Container Registry Webhook %q (Resource Group %q, Registry %q): %s", name, resourceGroup, registryName, err)
			}
		}

		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_container_registry_webhook", *existing.ID)
		}
	}

	location := azure.NormalizeLocation(d.Get("location").(string))

	webhookProperties := expandWebhookPropertiesCreateParameters(d)

	tags := d.Get("tags").(map[string]interface{})

	webhook := containerregistry.WebhookCreateParameters{
		Location:                          &location,
		WebhookPropertiesCreateParameters: webhookProperties,
		Tags:                              expandTags(tags),
	}

	future, err := client.Create(ctx, resourceGroup, registryName, name, webhook)
	if err != nil {
		return fmt.Errorf("Error creating Container Registry Webhook %q (Resource Group %q, Registry %q): %+v", name, resourceGroup, registryName, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("Error waiting for creation of Container Registry %q (Resource Group %q, Registry %q): %+v", name, resourceGroup, registryName, err)
	}

	read, err := client.Get(ctx, resourceGroup, registryName, name)
	if err != nil {
		return fmt.Errorf("Error retrieving Container Registry %q (Resource Group %q, Registry %q): %+v", name, resourceGroup, registryName, err)
	}

	if read.ID == nil {
		return fmt.Errorf("Cannot read Container Registry %q (resource group %q, Registry %q) ID", name, resourceGroup, registryName)
	}

	d.SetId(*read.ID)

	return resourceArmContainerRegistryWebhookRead(d, meta)
}

func resourceArmContainerRegistryWebhookUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).containers.WebhooksClient
	ctx := meta.(*ArmClient).StopContext

	log.Printf("[INFO] preparing arguments for AzureRM Container Registry Webhook update.")

	id, err := azure.ParseAzureResourceID(d.Id())
	if err != nil {
		return err
	}

	resourceGroup := id.ResourceGroup
	registryName := id.Path["registries"]
	name := id.Path["webhooks"]

	webhookProperties := expandWebhookPropertiesUpdateParameters(d)

	tags := d.Get("tags").(map[string]interface{})

	webhook := containerregistry.WebhookUpdateParameters{
		WebhookPropertiesUpdateParameters: webhookProperties,
		Tags:                              expandTags(tags),
	}

	future, err := client.Update(ctx, resourceGroup, registryName, name, webhook)
	if err != nil {
		return fmt.Errorf("Error updating Container Registry Webhook %q (Resource Group %q, Registry %q): %+v", name, resourceGroup, registryName, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("Error waiting for completion of Container Registry Webhook %q (Resource Group %q, Registry %q): %+v", name, resourceGroup, registryName, err)
	}

	return resourceArmContainerRegistryWebhookRead(d, meta)
}

func resourceArmContainerRegistryWebhookRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).containers.WebhooksClient
	ctx := meta.(*ArmClient).StopContext

	id, err := azure.ParseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resourceGroup := id.ResourceGroup
	registryName := id.Path["registries"]
	name := id.Path["webhooks"]

	resp, err := client.Get(ctx, resourceGroup, registryName, name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[DEBUG] Container Registry Webhook %q was not found in Resource Group %q for Registry %q", name, resourceGroup, registryName)
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error making Read request on Azure Container Registry Webhook %q (Resource Group %q, Registry %q): %+v", name, resourceGroup, registryName, err)
	}

	callbackConfig, err := client.GetCallbackConfig(ctx, resourceGroup, registryName, name)
	if err != nil {
		return fmt.Errorf("Error making Read request on Azure Container Registry Webhook Callback Config %q (Resource Group %q, Registry %q): %+v", name, resourceGroup, registryName, err)
	}

	d.Set("resource_group_name", resourceGroup)
	d.Set("registry_name", registryName)
	d.Set("name", name)

	if location := resp.Location; location != nil {
		d.Set("location", azure.NormalizeLocation(*location))
	}

	d.Set("service_uri", callbackConfig.ServiceURI)

	if callbackConfig.CustomHeaders != nil {
		customHeaders := make(map[string]string)
		for k, v := range callbackConfig.CustomHeaders {
			customHeaders[k] = *v
		}
		d.Set("custom_headers", customHeaders)
	}

	if webhookProps := resp.WebhookProperties; webhookProps != nil {
		if webhookProps.Status != "" {
			d.Set("status", string(webhookProps.Status))
		}

		if webhookProps.Scope != nil {
			d.Set("scope", webhookProps.Scope)
		}

		webhookActions := make([]string, len(*webhookProps.Actions))
		for i, action := range *webhookProps.Actions {
			webhookActions[i] = string(action)
		}
		d.Set("actions", webhookActions)
	}

	flattenAndSetTags(d, resp.Tags)

	return nil
}

func resourceArmContainerRegistryWebhookDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).containers.WebhooksClient
	ctx := meta.(*ArmClient).StopContext

	id, err := azure.ParseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resourceGroup := id.ResourceGroup
	registryName := id.Path["registries"]
	name := id.Path["webhooks"]

	future, err := client.Delete(ctx, resourceGroup, registryName, name)
	if err != nil {
		if response.WasNotFound(future.Response()) {
			return nil
		}
		return fmt.Errorf("Error issuing Azure ARM delete request of Container Registry Webhook '%s': %+v", name, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		if response.WasNotFound(future.Response()) {
			return nil
		}
		return fmt.Errorf("Error issuing Azure ARM delete request of Container Registry Webhook '%s': %+v", name, err)
	}

	return nil
}

func validateAzureRMContainerRegistryWebhookName(v interface{}, k string) (warnings []string, errors []error) {
	value := v.(string)
	if !regexp.MustCompile(`^[a-zA-Z0-9]{5,50}$`).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"alpha numeric characters only are allowed and between 5 and 50 characters in %q: %q", k, value))
	}

	return warnings, errors
}

func validateAzureRMContainerRegistryWebhookStatus() schema.SchemaValidateFunc {
	statuses := make([]string, len(containerregistry.PossibleWebhookStatusValues()))
	for i, v := range containerregistry.PossibleWebhookStatusValues() {
		statuses[i] = string(v)
	}

	return validation.StringInSlice(statuses, false)
}

func validateAzureRMContainerRegistryWebhookAction() schema.SchemaValidateFunc {
	actions := make([]string, len(containerregistry.PossibleWebhookActionValues()))
	for i, v := range containerregistry.PossibleWebhookActionValues() {
		actions[i] = string(v)
	}

	return validation.StringInSlice(actions, false)
}

func validateAzureRMContainerRegistryWebhookServiceUri(v interface{}, k string) (warnings []string, errors []error) {
	value := v.(string)
	if !regexp.MustCompile(`^https?://[^\s]+$`).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q must start with http:// or https:// and must not contain whitespaces: %q", k, value))
	}

	return warnings, errors
}

func expandWebhookPropertiesCreateParameters(d *schema.ResourceData) *containerregistry.WebhookPropertiesCreateParameters {
	serviceUri := d.Get("service_uri").(string)
	customHeaders := make(map[string]*string)
	for k, v := range d.Get("custom_headers").(map[string]interface{}) {
		customHeaders[k] = utils.String(v.(string))
	}

	actions := expandWebhookActions(d)

	scope := d.Get("scope").(string)

	webhookProperties := containerregistry.WebhookPropertiesCreateParameters{
		ServiceURI:    &serviceUri,
		CustomHeaders: customHeaders,
		Actions:       actions,
		Scope:         &scope,
	}

	if d.Get("status").(string) == string(containerregistry.WebhookStatusEnabled) {
		webhookProperties.Status = containerregistry.WebhookStatusEnabled
	} else {
		webhookProperties.Status = containerregistry.WebhookStatusDisabled
	}

	return &webhookProperties
}

func expandWebhookPropertiesUpdateParameters(d *schema.ResourceData) *containerregistry.WebhookPropertiesUpdateParameters {
	serviceUri := d.Get("service_uri").(string)

	customHeaders := make(map[string]*string)
	for k, v := range d.Get("custom_headers").(map[string]interface{}) {
		customHeaders[k] = utils.String(v.(string))
	}

	actions := expandWebhookActions(d)

	scope := d.Get("scope").(string)

	webhookProperties := containerregistry.WebhookPropertiesUpdateParameters{
		ServiceURI:    &serviceUri,
		CustomHeaders: customHeaders,
		Actions:       actions,
		Scope:         &scope,
	}

	if d.Get("status").(string) == string(containerregistry.WebhookStatusEnabled) {
		webhookProperties.Status = containerregistry.WebhookStatusEnabled
	} else {
		webhookProperties.Status = containerregistry.WebhookStatusDisabled
	}

	return &webhookProperties
}

func expandWebhookActions(d *schema.ResourceData) *[]containerregistry.WebhookAction {
	actions := make([]containerregistry.WebhookAction, 0)
	for _, action := range d.Get("actions").(*schema.Set).List() {
		switch action.(string) {
		case "chart_delete":
			actions = append(actions, containerregistry.ChartDelete)
		case "chart_push":
			actions = append(actions, containerregistry.ChartPush)
		case "delete":
			actions = append(actions, containerregistry.Delete)
		case "push":
			actions = append(actions, containerregistry.Push)
		case "quarantine":
			actions = append(actions, containerregistry.Quarantine)
		}
	}

	return &actions
}
