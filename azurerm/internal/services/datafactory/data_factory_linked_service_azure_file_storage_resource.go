package datafactory

import (
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/datafactory/mgmt/2018-06-01/datafactory"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/datafactory/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/datafactory/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceDataFactoryLinkedServiceAzureFileStorage() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Create: resourceDataFactoryLinkedServiceAzureFileStorageCreateUpdate,
		Read:   resourceDataFactoryLinkedServiceAzureFileStorageRead,
		Update: resourceDataFactoryLinkedServiceAzureFileStorageCreateUpdate,
		Delete: resourceDataFactoryLinkedServiceAzureFileStorageDelete,

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
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.LinkedServiceDatasetName,
			},

			"data_factory_name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.DataFactoryName(),
			},

			// There's a bug in the Azure API where this is returned in lower-case
			// BUG: https://github.com/Azure/azure-rest-api-specs/issues/5788
			"resource_group_name": azure.SchemaResourceGroupNameDiffSuppress(),

			"connection_string": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				Sensitive:    true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"file_share": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"host": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"user_id": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"password": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Sensitive:    true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"description": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"integration_runtime_name": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"key_vault_password": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"linked_service_name": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},

						"secret_name": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
					},
				},
			},

			"parameters": {
				Type:     pluginsdk.TypeMap,
				Optional: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"annotations": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"additional_properties": {
				Type:     pluginsdk.TypeMap,
				Optional: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},
		},
	}
}

func resourceDataFactoryLinkedServiceAzureFileStorageCreateUpdate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DataFactory.LinkedServiceClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	dataFactoryName := d.Get("data_factory_name").(string)
	resourceGroup := d.Get("resource_group_name").(string)

	if d.IsNewResource() {
		existing, err := client.Get(ctx, resourceGroup, dataFactoryName, name, "")
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("Error checking for presence of existing Data Factory Linked Service Azure File Storage Anonymous %q (Data Factory %q / Resource Group %q): %+v", name, dataFactoryName, resourceGroup, err)
			}
		}

		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_data_factory_linked_service_azure_file_storage", *existing.ID)
		}
	}

	fileStorageProperties := &datafactory.AzureFileStorageLinkedServiceTypeProperties{
		ConnectionString: &datafactory.SecureString{
			Value: utils.String(d.Get("connection_string").(string)),
			Type:  datafactory.TypeSecureString,
		},
		FileShare: d.Get("file_share").(string),
		Host:      d.Get("host").(string),
		UserID:    d.Get("connection_string").(string),
	}

	password := d.Get("password").(string)
	if password != "" {
		fileStorageProperties.Password = &datafactory.SecureString{
			Value: utils.String(d.Get("password").(string)),
			Type:  datafactory.TypeSecureString,
		}
	}

	fileStorageLinkedService := &datafactory.AzureFileStorageLinkedService{
		Description: utils.String(d.Get("description").(string)),
		AzureFileStorageLinkedServiceTypeProperties: fileStorageProperties,
		Type: datafactory.TypeBasicLinkedServiceTypeAzureFileStorage,
	}

	if v, ok := d.GetOk("parameters"); ok {
		fileStorageLinkedService.Parameters = expandDataFactoryParameters(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("integration_runtime_name"); ok {
		fileStorageLinkedService.ConnectVia = expandDataFactoryLinkedServiceIntegrationRuntime(v.(string))
	}

	if v, ok := d.GetOk("key_vault_password"); ok {
		password := v.([]interface{})
		fileStorageProperties.Password = expandAzureKeyVaultSecretReference(password)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		fileStorageLinkedService.AdditionalProperties = v.(map[string]interface{})
	}

	if v, ok := d.GetOk("annotations"); ok {
		annotations := v.([]interface{})
		fileStorageLinkedService.Annotations = &annotations
	}

	linkedService := datafactory.LinkedServiceResource{
		Properties: fileStorageLinkedService,
	}

	if _, err := client.CreateOrUpdate(ctx, resourceGroup, dataFactoryName, name, linkedService, ""); err != nil {
		return fmt.Errorf("Error creating/updating Data Factory Linked Service Azure File Storage %q (Data Factory %q / Resource Group %q): %+v", name, dataFactoryName, resourceGroup, err)
	}

	resp, err := client.Get(ctx, resourceGroup, dataFactoryName, name, "")
	if err != nil {
		return fmt.Errorf("Error retrieving Data Factory Linked Service Azure File Storage %q (Data Factory %q / Resource Group %q): %+v", name, dataFactoryName, resourceGroup, err)
	}

	if resp.ID == nil {
		return fmt.Errorf("Cannot read Data Factory Linked Service Azure File Storage %q (Data Factory %q / Resource Group %q): %+v", name, dataFactoryName, resourceGroup, err)
	}

	d.SetId(*resp.ID)

	return resourceDataFactoryLinkedServiceAzureFileStorageRead(d, meta)
}

func resourceDataFactoryLinkedServiceAzureFileStorageRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DataFactory.LinkedServiceClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.LinkedServiceID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.FactoryName, id.Name, "")
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error retrieving Data Factory Linked Service Azure File Storage %q (Data Factory %q / Resource Group %q): %+v", id.Name, id.FactoryName, id.ResourceGroup, err)
	}

	d.Set("name", resp.Name)
	d.Set("resource_group_name", id.ResourceGroup)
	d.Set("data_factory_name", id.FactoryName)

	fileStorage, ok := resp.Properties.AsAzureFileStorageLinkedService()
	if !ok {
		return fmt.Errorf("Error classifiying Data Factory Linked Service Azure File Storage %q (Data Factory %q / Resource Group %q): Expected: %q Received: %q", id.Name, id.FactoryName, id.ResourceGroup, datafactory.TypeBasicLinkedServiceTypeAzureFileStorage, *resp.Type)
	}

	d.Set("additional_properties", fileStorage.AdditionalProperties)
	d.Set("description", fileStorage.Description)

	if password := fileStorage.Password; password != nil {
		if keyVaultPassword, ok := password.AsAzureKeyVaultSecretReference(); ok {
			if err := d.Set("key_vault_password", flattenAzureKeyVaultSecretReference(keyVaultPassword)); err != nil {
				return fmt.Errorf("setting `key_vault_password`: %+v", err)
			}
		}
	}

	annotations := flattenDataFactoryAnnotations(fileStorage.Annotations)
	if err := d.Set("annotations", annotations); err != nil {
		return fmt.Errorf("Error setting `annotations` for Data Factory Linked Service Azure File Storage %q (Data Factory %q) / Resource Group %q): %+v", id.Name, id.FactoryName, id.ResourceGroup, err)
	}

	parameters := flattenDataFactoryParameters(fileStorage.Parameters)
	if err := d.Set("parameters", parameters); err != nil {
		return fmt.Errorf("Error setting `parameters`: %+v", err)
	}

	if connectVia := fileStorage.ConnectVia; connectVia != nil {
		if connectVia.ReferenceName != nil {
			d.Set("integration_runtime_name", connectVia.ReferenceName)
		}
	}

	if props := fileStorage.AzureFileStorageLinkedServiceTypeProperties; props != nil {
		if props.FileShare != nil {
			d.Set("file_share", props.FileShare)
		}
	}

	return nil
}

func resourceDataFactoryLinkedServiceAzureFileStorageDelete(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DataFactory.LinkedServiceClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.LinkedServiceID(d.Id())
	if err != nil {
		return err
	}

	response, err := client.Delete(ctx, id.ResourceGroup, id.FactoryName, id.Name)
	if err != nil {
		if !utils.ResponseWasNotFound(response) {
			return fmt.Errorf("Error deleting Data Factory Linked Service Azure File Storage %q (Data Factory %q / Resource Group %q): %+v", id.Name, id.FactoryName, id.ResourceGroup, err)
		}
	}

	return nil
}
