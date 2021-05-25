package keyvault

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/keyvault/parse"
	keyVaultValidate "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/keyvault/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tags"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceKeyVaultManagedStorageAccountSasToken() *schema.Resource {
	return &schema.Resource{
		Create: resourceKeyVaultManagedStorageAccountSasTokenCreateUpdate,
		Read:   resourceKeyVaultManagedStorageAccountSasTokenRead,
		Update: resourceKeyVaultManagedStorageAccountSasTokenCreateUpdate,
		Delete: resourceKeyVaultManagedStorageAccountSasTokenDelete,

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.SasDefinitionID(id)
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
				ValidateFunc: keyVaultValidate.NestedItemName,
			},

			"managed_storage_account_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: keyVaultValidate.VersionlessNestedItemId,
			},

			"sas_template_uri": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"sas_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"service",
					"account",
				}, false),
			},

			"secret_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"validity_period": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.ISO8601Duration,
			},

			"tags": tags.ForceNewSchema(),
		},
	}
}

func resourceKeyVaultManagedStorageAccountSasTokenCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	keyVaultsClient := meta.(*clients.Client).KeyVault
	client := meta.(*clients.Client).KeyVault.ManagementClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	resourcesClient := meta.(*clients.Client).Resource

	defer cancel()

	name := d.Get("name").(string)
	storageAccount, err := parse.ParseOptionallyVersionedNestedItemID(d.Get("managed_storage_account_id").(string))
	if err != nil {
		return err
	}

	keyVaultIdRaw, err := keyVaultsClient.KeyVaultIDFromBaseUrl(ctx, resourcesClient, storageAccount.KeyVaultBaseUrl)
	if err != nil {
		return fmt.Errorf("retrieving the Resource ID the Key Vault at URL %q: %s", storageAccount.KeyVaultBaseUrl, err)
	}
	keyVaultId, err := parse.VaultID(*keyVaultIdRaw)
	if err != nil {
		return err
	}

	keyVaultBaseUri, err := keyVaultsClient.BaseUriForKeyVault(ctx, *keyVaultId)
	if err != nil {
		return fmt.Errorf("retrieving base uri for %s: %+v", *keyVaultId, err)
	}

	if d.IsNewResource() {
		existing, err := client.GetSasDefinition(ctx, *keyVaultBaseUri, storageAccount.Name, name)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("checking for presence of existing Managed Storage Account %q (Key Vault %q): %s", name, *keyVaultBaseUri, err)
			}
		}

		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_key_vault_managed_storage_account_sasdefinition", *existing.ID)
		}
	}

	t := d.Get("tags").(map[string]interface{})
	parameters := keyvault.SasDefinitionCreateParameters{
		TemplateURI:    utils.String(d.Get("sas_template_uri").(string)),
		SasType:        keyvault.SasTokenType(d.Get("sas_type").(string)),
		ValidityPeriod: utils.String(d.Get("validity_period").(string)),
		SasDefinitionAttributes: &keyvault.SasDefinitionAttributes{
			Enabled: utils.Bool(true),
		},
		Tags: tags.Expand(t),
	}

	if resp, err := client.SetSasDefinition(ctx, *keyVaultBaseUri, storageAccount.Name, name, parameters); err != nil {
		// In the case that the Storage Account already exists in a Soft Deleted / Recoverable state we check if `recover_soft_deleted_key_vaults` is set
		// and attempt recovery where appropriate
		if meta.(*clients.Client).Features.KeyVault.RecoverSoftDeletedKeyVaults && utils.ResponseWasConflict(resp.Response) {
			recoveredStorageAccount, err := client.RecoverDeletedSasDefinition(ctx, *keyVaultBaseUri, storageAccount.Name, name)
			if err != nil {
				return err
			}
			log.Printf("[DEBUG] Recovering Managed Storage Account Sas Definition %q with ID: %q", name, *recoveredStorageAccount.ID)
			// We need to wait for consistency, recovered Key Vault Child items are not as readily available as newly created
			if secret := recoveredStorageAccount.ID; secret != nil {
				stateConf := &resource.StateChangeConf{
					Pending:                   []string{"pending"},
					Target:                    []string{"available"},
					Refresh:                   keyVaultChildItemRefreshFunc(*secret),
					Delay:                     30 * time.Second,
					PollInterval:              10 * time.Second,
					ContinuousTargetOccurence: 10,
					Timeout:                   d.Timeout(schema.TimeoutCreate),
				}

				if _, err := stateConf.WaitForState(); err != nil {
					return fmt.Errorf("Error waiting for Key Vault Managed Storage Account Sas Definition %q to become available: %s", name, err)
				}
				log.Printf("[DEBUG] Managed Storage Account Sas Definition %q recovered with ID: %q", name, *recoveredStorageAccount.ID)

				_, err := client.SetSasDefinition(ctx, *keyVaultBaseUri, storageAccount.Name, name, parameters)
				if err != nil {
					return err
				}
			}
		} else {
			// If the error response was anything else, or `recover_soft_deleted_key_vaults` is `false` just return the error
			return err
		}
	}

	read, err := client.GetSasDefinition(ctx, *keyVaultBaseUri, storageAccount.Name, name)

	if err != nil {
		return err
	}

	if read.ID == nil {
		return fmt.Errorf("cannot read Managed Storage Account Sas Definition '%s' (in key vault '%s')", name, *keyVaultBaseUri)
	}

	d.SetId(*read.ID)

	return resourceKeyVaultManagedStorageAccountSasTokenRead(d, meta)
}

func resourceKeyVaultManagedStorageAccountSasTokenRead(d *schema.ResourceData, meta interface{}) error {
	keyVaultsClient := meta.(*clients.Client).KeyVault
	client := meta.(*clients.Client).KeyVault.ManagementClient
	resourcesClient := meta.(*clients.Client).Resource
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.SasDefinitionID(d.Id())
	if err != nil {
		return err
	}

	keyVaultIdRaw, err := keyVaultsClient.KeyVaultIDFromBaseUrl(ctx, resourcesClient, id.KeyVaultBaseUrl)
	if err != nil {
		return fmt.Errorf("retrieving the Resource ID the Key Vault at URL %q: %s", id.KeyVaultBaseUrl, err)
	}

	keyVaultId, err := parse.VaultID(*keyVaultIdRaw)
	if err != nil {
		return err
	}

	ok, err := keyVaultsClient.Exists(ctx, *keyVaultId)
	if err != nil {
		return fmt.Errorf("Error checking if key vault %q for Managed Storage Account Sas Definition %q in Vault at url %q exists: %v", *keyVaultId, id.Name, id.KeyVaultBaseUrl, err)
	}
	if !ok {
		log.Printf("[DEBUG] Managed Storage Account Sas Definition %q Key Vault %q was not found in Key Vault at URI %q - removing from state", id.Name, *keyVaultId, id.KeyVaultBaseUrl)
		d.SetId("")
		return nil
	}

	resp, err := client.GetSasDefinition(ctx, id.KeyVaultBaseUrl, id.StorageAccountName, id.Name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[DEBUG] Managed Storage Account Sas Definition %q was not found in Key Vault at URI %q - removing from state", id.Name, id.KeyVaultBaseUrl)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error making Read request on Managed Storage Account Sas Definition %s: %+v", id.Name, err)
	}

	d.Set("name", id.Name)
	d.Set("sas_template_uri", resp.TemplateURI)
	d.Set("sas_type", resp.SasType)
	d.Set("secret_id", resp.SecretID)
	d.Set("validity_period", resp.ValidityPeriod)

	return tags.FlattenAndSet(d, resp.Tags)
}

func resourceKeyVaultManagedStorageAccountSasTokenDelete(d *schema.ResourceData, meta interface{}) error {
	keyVaultsClient := meta.(*clients.Client).KeyVault
	client := meta.(*clients.Client).KeyVault.ManagementClient
	resourcesClient := meta.(*clients.Client).Resource
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.SasDefinitionID(d.Id())
	if err != nil {
		return err
	}

	keyVaultIdRaw, err := keyVaultsClient.KeyVaultIDFromBaseUrl(ctx, resourcesClient, id.KeyVaultBaseUrl)
	if err != nil {
		return fmt.Errorf("Error retrieving the Resource ID the Key Vault at URL %q: %s", id.KeyVaultBaseUrl, err)
	}
	if keyVaultIdRaw == nil {
		return fmt.Errorf("Unable to determine the Resource ID for the Key Vault at URL %q", id.KeyVaultBaseUrl)
	}
	keyVaultId, err := parse.VaultID(*keyVaultIdRaw)
	if err != nil {
		return err
	}

	ok, err := keyVaultsClient.Exists(ctx, *keyVaultId)
	if err != nil {
		return fmt.Errorf("Error checking if key vault %q for Managed Storage Account Sas Definition %q in Vault at url %q exists: %v", *keyVaultId, id.Name, id.KeyVaultBaseUrl, err)
	}
	if !ok {
		log.Printf("[DEBUG] Managed Storage Account Sas Definition %q Key Vault %q was not found in Key Vault at URI %q - removing from state", id.Name, *keyVaultId, id.KeyVaultBaseUrl)
		d.SetId("")
		return nil
	}

	_, err = client.DeleteSasDefinition(ctx, id.KeyVaultBaseUrl, id.StorageAccountName, id.Name)
	return err
}
