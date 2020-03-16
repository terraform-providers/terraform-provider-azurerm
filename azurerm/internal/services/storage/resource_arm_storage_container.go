package storage

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/storage/parsers"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
	"github.com/tombuildsstuff/giovanni/storage/2018-11-09/blob/containers"
)

func resourceArmStorageContainer() *schema.Resource {
	return &schema.Resource{
		Create:        resourceArmStorageContainerCreate,
		Read:          resourceArmStorageContainerRead,
		Delete:        resourceArmStorageContainerDelete,
		Update:        resourceArmStorageContainerUpdate,
		MigrateState:  ResourceStorageContainerMigrateState,
		SchemaVersion: 1,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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
				ValidateFunc: validate.StorageContainerName,
			},

			"storage_account_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: ValidateArmStorageAccountName,
			},

			"container_access_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "private",
				ValidateFunc: validation.StringInSlice([]string{
					"blob",
					"container",
					"private",
				}, false),
			},

			"metadata": MetaDataComputedSchema(),

			// TODO: support for ACL's, Legal Holds and Immutability Policies
			"has_immutability_policy": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"has_legal_hold": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"resource_manager_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceArmStorageContainerCreate(d *schema.ResourceData, meta interface{}) error {
	storageClient := meta.(*clients.Client).Storage
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	containerName := d.Get("name").(string)
	accountName := d.Get("storage_account_name").(string)

	accessLevelRaw := d.Get("container_access_type").(string)
	accessLevel, err := expandStorageContainerAccessLevel(accessLevelRaw)
	if err != nil {
		return fmt.Errorf("Invalid container public access Account %q for Container %q: %s", accountName, containerName, err)
	}

	metaDataRaw := d.Get("metadata").(map[string]interface{})
	metaData := expandMetaData(metaDataRaw)

	account, err := storageClient.FindAccount(ctx, accountName)
	if err != nil {
		return fmt.Errorf("Error retrieving Account %q for Container %q: %s", accountName, containerName, err)
	}
	if account == nil {
		return fmt.Errorf("Unable to locate Storage Account %q!", accountName)
	}

	client := storageClient.BlobContainersClient

	id := getResourceID(meta.(*clients.Client).Account.Environment.StorageEndpointSuffix, accountName, containerName)
	if features.ShouldResourcesBeImported() {
		existing, err := client.Get(ctx, account.ResourceGroup, accountName, containerName)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("Error checking for existence of existing Container %q (Account %q / Resource Group %q): %+v", containerName, accountName, account.ResourceGroup, err)
			}
		}

		if !utils.ResponseWasNotFound(existing.Response) {
			return tf.ImportAsExistsError("azurerm_storage_container", id)
		}
	}

	log.Printf("[INFO] Creating Container %q in Storage Account %q", containerName, accountName)
	input := storage.BlobContainer{
		ContainerProperties: &storage.ContainerProperties{
			PublicAccess: accessLevel,
			Metadata:     metaData,
		},
	}
	if _, err := client.Create(ctx, account.ResourceGroup, accountName, containerName, input); err != nil {
		return fmt.Errorf("Error creating Container %q (Account %q / Resource Group %q): %s", containerName, accountName, account.ResourceGroup, err)
	}

	d.SetId(id)
	return resourceArmStorageContainerRead(d, meta)
}

func resourceArmStorageContainerUpdate(d *schema.ResourceData, meta interface{}) error {
	storageClient := meta.(*clients.Client).Storage
	ctx, cancel := timeouts.ForUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parsers.ParseContainerID(d.Id())
	if err != nil {
		return err
	}

	account, err := storageClient.FindAccount(ctx, id.AccountName)
	if err != nil {
		return fmt.Errorf("Error retrieving Account %q for Container %q: %s", id.AccountName, id.ContainerName, err)
	}
	if account == nil {
		return fmt.Errorf("Unable to locate Storage Account %q!", id.AccountName)
	}

	client := storageClient.BlobContainersClient

	log.Printf("[DEBUG] Computing the Access Control for Container %q (Storage Account %q / Resource Group %q)..", id.ContainerName, id.AccountName, account.ResourceGroup)
	accessLevelRaw := d.Get("container_access_type").(string)
	accessLevel, err := expandStorageContainerAccessLevel(accessLevelRaw)
	if err != nil {
		return fmt.Errorf("Invalid public access for Container %q (Storage Account %q / Resource Group %q): %s", id.ContainerName, id.AccountName, account.ResourceGroup, err)
	}

	log.Printf("[DEBUG] Computing the MetaData for Container %q (Storage Account %q / Resource Group %q)..", id.ContainerName, id.AccountName, account.ResourceGroup)
	metaDataRaw := d.Get("metadata").(map[string]interface{})
	metaData := expandMetaData(metaDataRaw)

	if d.HasChange("container_access_type") || d.HasChange("metadata") {
		input := storage.BlobContainer{
			ContainerProperties: &storage.ContainerProperties{
				PublicAccess: accessLevel,
				Metadata:     metaData,
			},
		}
		if _, err := client.Update(ctx, account.ResourceGroup, id.AccountName, id.ContainerName, input); err != nil {
			return fmt.Errorf("Error updating Container %q (Storage Account %q / Resource Group %q): %s", id.ContainerName, id.AccountName, account.ResourceGroup, err)
		}
	}

	return resourceArmStorageContainerRead(d, meta)
}

func resourceArmStorageContainerRead(d *schema.ResourceData, meta interface{}) error {
	storageClient := meta.(*clients.Client).Storage
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parsers.ParseContainerID(d.Id())
	if err != nil {
		return err
	}

	account, err := storageClient.FindAccount(ctx, id.AccountName)
	if err != nil {
		return fmt.Errorf("Error retrieving Account %q for Container %q: %s", id.AccountName, id.ContainerName, err)
	}
	if account == nil {
		log.Printf("[DEBUG] Unable to locate Account %q for Storage Container %q - assuming removed & removing from state", id.AccountName, id.ContainerName)
		d.SetId("")
		return nil
	}

	client := storageClient.BlobContainersClient

	props, err := client.Get(ctx, account.ResourceGroup, id.AccountName, id.ContainerName)
	if err != nil {
		if utils.ResponseWasNotFound(props.Response) {
			log.Printf("[DEBUG] Container %q was not found in Account %q / Resource Group %q - assuming removed & removing from state", id.ContainerName, id.AccountName, account.ResourceGroup)
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error retrieving Container %q (Account %q / Resource Group %q): %s", id.ContainerName, id.AccountName, account.ResourceGroup, err)
	}

	d.Set("name", id.ContainerName)
	d.Set("storage_account_name", id.AccountName)

	accessLevel, err := flattenStorageContainerAccessLevel(props.PublicAccess)
	if err != nil {
		fmt.Errorf("Error retrieving public access Container %q (Account %q / Resource Group %q): %s", id.ContainerName, id.AccountName, account.ResourceGroup, err)
	}

	d.Set("container_access_type", accessLevel)

	if err := d.Set("metadata", flattenMetaData(props.Metadata)); err != nil {
		return fmt.Errorf("Error setting `metadata`: %+v", err)
	}

	d.Set("has_immutability_policy", props.HasImmutabilityPolicy)
	d.Set("has_legal_hold", props.HasLegalHold)

	resourceManagerId := client.GetResourceManagerResourceID(storageClient.SubscriptionId, account.ResourceGroup, id.AccountName, id.ContainerName)
	d.Set("resource_manager_id", resourceManagerId)

	return nil
}

func resourceArmStorageContainerDelete(d *schema.ResourceData, meta interface{}) error {
	storageClient := meta.(*clients.Client).Storage
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := containers.ParseResourceID(d.Id())
	if err != nil {
		return err
	}

	account, err := storageClient.FindAccount(ctx, id.AccountName)
	if err != nil {
		return fmt.Errorf("Error retrieving Account %q for Container %q: %s", id.AccountName, id.ContainerName, err)
	}
	if account == nil {
		return fmt.Errorf("Unable to locate Storage Account %q!", id.AccountName)
	}

	client, err := storageClient.ContainersClient(ctx, *account)
	if err != nil {
		return fmt.Errorf("Error building Containers Client for Storage Account %q (Resource Group %q): %s", id.AccountName, account.ResourceGroup, err)
	}

	if _, err := client.Delete(ctx, id.AccountName, id.ContainerName); err != nil {
		return fmt.Errorf("Error deleting Container %q (Storage Account %q / Resource Group %q): %s", id.ContainerName, id.AccountName, account.ResourceGroup, err)
	}

	return nil
}

func expandStorageContainerAccessLevel(input string) (storage.PublicAccess, error) {
	switch input {
	case "private":
		return storage.PublicAccessNone, nil
	case "container":
		return storage.PublicAccessContainer, nil
	case "blob":
		return storage.PublicAccessBlob, nil
	default:
		return "", fmt.Errorf("%s", input)
	}
}

func flattenStorageContainerAccessLevel(input storage.PublicAccess) (string, error) {
	switch input {
	case storage.PublicAccessNone:
		return "private", nil
	case storage.PublicAccessContainer:
		return "container", nil
	case storage.PublicAccessBlob:
		return "blob", nil
	default:
		return "", fmt.Errorf("%s", input)
	}
}

func getResourceID(baseUri, accountName, containerName string) string {
	domain := parsers.GetBlobEndpoint(baseUri, accountName)
	return fmt.Sprintf("%s/%s", domain, containerName)
}

func expandMetaData(input map[string]interface{}) map[string]*string {
	output := make(map[string]*string)

	for k, v := range input {
		temp := v.(string)
		output[k] = &temp
	}

	return output
}

func flattenMetaData(input map[string]*string) map[string]interface{} {
	output := make(map[string]interface{})

	for k, v := range input {
		output[k] = &v
	}

	return output
}
