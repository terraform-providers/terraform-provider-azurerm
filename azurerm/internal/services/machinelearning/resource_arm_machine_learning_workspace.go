package machinelearning

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2018-09-01/containerregistry"
	"github.com/Azure/azure-sdk-for-go/services/machinelearningservices/mgmt/2019-11-01/machinelearningservices"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/suppress"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/machinelearning/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/machinelearning/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tags"
	azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmMachineLearningWorkspace() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmAmlWorkspaceCreate,
		Read:   resourceArmAmlWorkspaceRead,
		Update: resourceArmAmlWorkspaceUpdate,
		Delete: resourceArmAmlWorkspaceDelete,

		Importer: azSchema.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.WorkspaceID(id)
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
				ValidateFunc: validate.WorkspaceName,
			},

			"location": azure.SchemaLocation(),

			"resource_group_name": azure.SchemaResourceGroupName(),

			"application_insights_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				// TODO -- use the custom validation function of application insights
				ValidateFunc: azure.ValidateResourceID,
				// TODO -- remove when issue https://github.com/Azure/azure-rest-api-specs/issues/8323 is addressed
				DiffSuppressFunc: suppress.CaseDifference,
			},

			"key_vault_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				// TODO -- use the custom validation function of key vault
				ValidateFunc: azure.ValidateResourceID,
				// TODO -- remove when issue https://github.com/Azure/azure-rest-api-specs/issues/8323 is addressed
				DiffSuppressFunc: suppress.CaseDifference,
			},

			"storage_account_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				// TODO -- use the custom validation function of storage account
				ValidateFunc: azure.ValidateResourceID,
				// TODO -- remove when issue https://github.com/Azure/azure-rest-api-specs/issues/8323 is addressed
				DiffSuppressFunc: suppress.CaseDifference,
			},

			"identity": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								string(machinelearningservices.SystemAssigned),
							}, false),
						},
						"principal_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tenant_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"container_registry_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// TODO -- use the custom validation function of container registry
				ValidateFunc: azure.ValidateResourceID,
				// TODO -- remove when issue https://github.com/Azure/azure-rest-api-specs/issues/8323 is addressed
				DiffSuppressFunc: suppress.CaseDifference,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"sku_name": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Basic",
					"Enterprise",
				}, true),
				Default: "Basic",
			},

			"tags": tags.Schema(),
		},
	}
}

func resourceArmAmlWorkspaceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).MachineLearning.WorkspacesClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	resourceGroup := d.Get("resource_group_name").(string)

	if features.ShouldResourcesBeImported() && d.IsNewResource() {
		existing, err := client.Get(ctx, resourceGroup, name)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("Error checking for existing AML Workspace %q (Resource Group %q): %s", name, resourceGroup, err)
			}
		}
		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_machine_learning_workspace", *existing.ID)
		}
	}

	location := azure.NormalizeLocation(d.Get("location").(string))
	storageAccountId := d.Get("storage_account_id").(string)
	keyVaultId := d.Get("key_vault_id").(string)
	applicationInsightsId := d.Get("application_insights_id").(string)
	skuName := d.Get("sku_name").(string)

	identityRaw := d.Get("identity").([]interface{})
	identity := expandMachineLearningWorkspaceIdentity(identityRaw)

	t := d.Get("tags").(map[string]interface{})

	workspace := machinelearningservices.Workspace{
		Name:     &name,
		Location: &location,
		Tags:     tags.Expand(t),
		WorkspaceProperties: &machinelearningservices.WorkspaceProperties{
			StorageAccount:      &storageAccountId,
			ApplicationInsights: &applicationInsightsId,
			KeyVault:            &keyVaultId,
		},
		Identity: identity,
		Sku:      &machinelearningservices.Sku{Name: utils.String(skuName)},
	}

	if v, ok := d.GetOk("description"); ok {
		workspace.Description = utils.String(v.(string))
	}

	if v, ok := d.GetOk("friendly_name"); ok {
		workspace.FriendlyName = utils.String(v.(string))
	}

	if v, ok := d.GetOk("container_registry_id"); ok {
		workspace.ContainerRegistry = utils.String(v.(string))
	}

	accountsClient := meta.(*clients.Client).Storage.AccountsClient
	if err := validateStorageAccount(ctx, accountsClient, storageAccountId); err != nil {
		return fmt.Errorf("Error creating Machine Learning Workspace %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	registriesClient := meta.(*clients.Client).Containers.RegistriesClient
	if err := validateContainerRegistry(ctx, registriesClient, workspace.ContainerRegistry); err != nil {
		return fmt.Errorf("Error creating Machine Learning Workspace %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	future, err := client.CreateOrUpdate(ctx, resourceGroup, name, workspace)
	if err != nil {
		return fmt.Errorf("Error creating Machine Learning Workspace %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("Error waiting for creation of Machine Learning Workspace %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	resp, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		return fmt.Errorf("Error retrieving Machine Learning Workspace %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	if resp.ID == nil {
		return fmt.Errorf("Cannot read Machine Learning Workspace %q (Resource Group %q) ID", name, resourceGroup)
	}

	d.SetId(*resp.ID)

	return resourceArmAmlWorkspaceRead(d, meta)
}

func resourceArmAmlWorkspaceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).MachineLearning.WorkspacesClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.WorkspaceID(d.Id())
	if err != nil {
		return fmt.Errorf("Error parsing Machine Learning Workspace ID `%q`: %+v", d.Id(), err)
	}

	name := id.Name
	resourceGroup := id.ResourceGroup

	resp, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error making Read request on Workspace %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	d.Set("name", resp.Name)
	d.Set("resource_group_name", resourceGroup)

	if location := resp.Location; location != nil {
		d.Set("location", azure.NormalizeLocation(*location))
	}

	if sku := resp.Sku; sku != nil {
		d.Set("sku_name", sku.Name)
	}

	if props := resp.WorkspaceProperties; props != nil {
		d.Set("application_insights_id", props.ApplicationInsights)
		d.Set("storage_account_id", props.StorageAccount)
		d.Set("key_vault_id", props.KeyVault)
		d.Set("container_registry_id", props.ContainerRegistry)
		d.Set("description", props.Description)
		d.Set("friendly_name", props.FriendlyName)
	}

	if err := d.Set("identity", flattenMachineLearningWorkspaceIdentity(resp.Identity)); err != nil {
		return fmt.Errorf("Error flattening identity on Workspace %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmAmlWorkspaceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).MachineLearning.WorkspacesClient
	ctx, cancel := timeouts.ForUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.WorkspaceID(d.Id())
	if err != nil {
		return err
	}

	update := machinelearningservices.WorkspaceUpdateParameters{
		WorkspacePropertiesUpdateParameters: &machinelearningservices.WorkspacePropertiesUpdateParameters{},
	}

	if d.HasChange("sku_name") {
		skuName := d.Get("sku_name").(string)
		update.Sku = &machinelearningservices.Sku{
			Name: &skuName,
		}
	}

	if d.HasChange("description") {
		description := d.Get("description").(string)
		update.WorkspacePropertiesUpdateParameters.Description = &description
	}

	if d.HasChange("friendly_name") {
		friendlyName := d.Get("friendly_name").(string)
		update.WorkspacePropertiesUpdateParameters.FriendlyName = &friendlyName
	}

	if d.HasChange("tags") {
		update.Tags = tags.Expand(d.Get("tags").(map[string]interface{}))
	}

	_, err = client.Update(ctx, id.ResourceGroup, id.Name, update)
	if err != nil {
		return fmt.Errorf("Error updating Machine Learning Workspace %q (Resource Group %q): %+v", id.Name, id.ResourceGroup, err)
	}

	return resourceArmAmlWorkspaceRead(d, meta)
}

func resourceArmAmlWorkspaceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).MachineLearning.WorkspacesClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.WorkspaceID(d.Id())
	if err != nil {
		return fmt.Errorf("Error parsing Machine Learning Workspace ID `%q`: %+v", d.Id(), err)
	}

	name := id.Name
	resourceGroup := id.ResourceGroup

	_, err = client.Delete(ctx, resourceGroup, name)
	if err != nil {
		return fmt.Errorf("Error deleting Machine Learning Workspace %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	return nil
}

func validateStorageAccount(ctx context.Context, client *storage.AccountsClient, accountID string) error {
	if accountID == "" {
		return fmt.Errorf("Error validating Storage Account: Empty ID")
	}

	// TODO -- use parse function "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/storage/parsers".ParseAccountID
	// when issue https://github.com/Azure/azure-rest-api-specs/issues/8323 is addressed
	id, err := parse.AccountIDCaseDiffSuppress(accountID)
	if err != nil {
		return fmt.Errorf("Error validating Storage Account: %+v", err)
	}

	account, err := client.GetProperties(ctx, id.ResourceGroup, id.Name, "")
	if err != nil {
		return fmt.Errorf("Error validating Storage Account %q (Resource Group %q): %+v", id.Name, id.ResourceGroup, err)
	}
	if sku := account.Sku; sku != nil {
		if sku.Tier == storage.Premium {
			return fmt.Errorf("Error validating Storage Account %q (Resource Group %q): The associated Storage Account must not be Premium", id.Name, id.ResourceGroup)
		}
	}

	return nil
}

func validateContainerRegistry(ctx context.Context, client *containerregistry.RegistriesClient, acrID *string) error {
	if acrID == nil {
		return nil
	}

	// TODO: use container registry's custom ID parse function when implemented
	id, err := azure.ParseAzureResourceID(*acrID)
	if err != nil {
		return fmt.Errorf("Error validating Container Registry: %+v", err)
	}

	acrName := id.Path["registries"]
	resourceGroup := id.ResourceGroup
	acr, err := client.Get(ctx, resourceGroup, acrName)
	if err != nil {
		return fmt.Errorf("Error validating Container Registry %q (Resource Group %q): %+v", acrName, resourceGroup, err)
	}
	if acr.AdminUserEnabled == nil || !*acr.AdminUserEnabled {
		return fmt.Errorf("Error validating Container Registry%q (Resource Group %q): The associated Container Registry must set `admin_enabled` to true", acrName, resourceGroup)
	}

	return nil
}

func flattenMachineLearningWorkspaceIdentity(identity *machinelearningservices.Identity) []interface{} {
	if identity == nil {
		return []interface{}{}
	}

	t := string(identity.Type)

	principalID := ""
	if identity.PrincipalID != nil {
		principalID = *identity.PrincipalID
	}

	tenantID := ""
	if identity.TenantID != nil {
		tenantID = *identity.TenantID
	}

	return []interface{}{
		map[string]interface{}{
			"type":         t,
			"principal_id": principalID,
			"tenant_id":    tenantID,
		},
	}
}

func expandMachineLearningWorkspaceIdentity(input []interface{}) *machinelearningservices.Identity {
	if len(input) == 0 {
		return nil
	}

	v := input[0].(map[string]interface{})

	identityType := machinelearningservices.ResourceIdentityType(v["type"].(string))

	identity := machinelearningservices.Identity{
		Type: identityType,
	}

	return &identity
}
