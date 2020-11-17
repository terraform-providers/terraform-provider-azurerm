package containers

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-06-01-preview/containerregistry"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/hashicorp/go-azure-helpers/response"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmContainerRegistryToken() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmContainerRegistryTokenCreate,
		Read:   resourceArmContainerRegistryTokenRead,
		Update: resourceArmContainerRegistryTokenUpdate,
		Delete: resourceArmContainerRegistryTokenDelete,
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
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_group_name": azure.SchemaResourceGroupName(),
			"container_registry_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: ValidateAzureRMContainerRegistryName,
			},
			"scope_map_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: azure.ValidateResourceID,
			},
			"status": {
				Type:     schema.TypeString,
				Required: false,
				Default:  containerregistry.TokenStatusEnabled,
				ValidateFunc: validation.StringInSlice([]string{
					string(containerregistry.TokenStatusDisabled),
					string(containerregistry.TokenStatusEnabled),
				}, false),
			},
			"first_password": {
				Type: schema.TypeMap,
				Required: false,
				Schema: *schema.Resource{
					Schema: map[string]*schema.Schema{
						"expires_at": {
							Type: schema.TypeString,
							Optional: true,
							ValidateFunc: validation.IsRFC3339Time,
						},
						"value": {
							Type: schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"second_password": {
				Type: schema.TypeMap,
				Required: false,
				Schema: *schema.Resource{
					Schema: map[string]*schema.Schema{
						"expires_at": {
							Type: schema.TypeString,
							Optional: true,
							ValidateFunc: validation.IsRFC3339Time,
						},
						"value": {
							Type: schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceArmContainerRegistryTokenCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Containers.TokensClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	log.Printf("[INFO] preparing arguments for AzureRM Container Registry token creation.")
	resourceGroup := d.Get("resource_group_name").(string)
	containerRegistryName := d.Get("container_registry_name").(string)
	name := d.Get("name").(string)

	if d.IsNewResource() {
		existing, err := client.Get(ctx, resourceGroup, containerRegistryName, name)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("Error checking for presence of existing token %q in Container Registry %q (Resource Group %q): %s", name, containerRegistryName, resourceGroup, err)
			}
		}

		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_container_registry_token", *existing.ID)
		}
	}

	scopeMapID := d.Get("scope_map_id").(string)
	status := d.Get("status").(string)

	firstPassword, err := expandTokenPassword(d, "first_password") 

	if err != nil {
		return fmt.Errorf("Error parsing first password for token %q in Container Registry %q (Resource Group %q): %+v", name, containerRegistryName, resourceGroup, err)
	}

	secondPassword, err := expandTokenPassword(d, "second_password")
	
	if err != nil {
		return fmt.Errorf("Error parsing second password for token %q in Container Registry %q (Resource Group %q): %+v", name, containerRegistryName, resourceGroup, err)
	}

	passwords := []containerregistry.TokenPassword{}

	if (firstPassword != nil) {
		passwords = append(passwords, *firstPassword)
	}

	if (secondPassword != nil) {
		passwords = append(passwords, *secondPassword)
	}

	parameters := containerregistry.Token{
		TokenProperties: &containerregistry.TokenProperties{
			ScopeMapID: utils.String(scopeMapID),
			Status:     containerregistry.TokenStatus(status),
			Credentials: &containerregistry.TokenCredentialsProperties{
				Passwords: &passwords,
			},
		},
	}

	future, err := client.Create(ctx, resourceGroup, containerRegistryName, name, parameters)
	if err != nil {
		return fmt.Errorf("Error creating token %q in Container Registry %q (Resource Group %q): %+v", name, containerRegistryName, resourceGroup, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("Error waiting for creation of token %q (Container Registry %q, Resource Group %q): %+v", name, containerRegistryName, resourceGroup, err)
	}

	read, err := client.Get(ctx, resourceGroup, containerRegistryName, name)
	if err != nil {
		return fmt.Errorf("Error retrieving token %q for Container Registry %q (Resource Group %q): %+v", name, containerRegistryName, resourceGroup, err)
	}

	if read.ID == nil {
		return fmt.Errorf("Cannot read token %q for Container Registry %q (resource group %q) ID", name, containerRegistryName, resourceGroup)
	}

	d.SetId(*read.ID)

	return resourceArmContainerRegistryTokenRead(d, meta)
}
func resourceArmContainerRegistryTokenUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Containers.TokensClient
	ctx, cancel := timeouts.ForUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	log.Printf("[INFO] preparing arguments for AzureRM Container Registry token update.")
	resourceGroup := d.Get("resource_group_name").(string)
	containerRegistryName := d.Get("container_registry_name").(string)
	name := d.Get("name").(string)
	scopeMapID := d.Get("scope_map_id").(string)
	status := d.Get("status").(string)

	passwords := []containerregistry.TokenPassword{}

	firstPassword, err := expandTokenPassword(d, "first_password") 

	if err != nil {
		return fmt.Errorf("Error parsing first password for token %q in Container Registry %q (Resource Group %q): %+v", name, containerRegistryName, resourceGroup, err)
	}

	secondPassword, err := expandTokenPassword(d, "second_password")
	
	if err != nil {
		return fmt.Errorf("Error parsing second password for token %q in Container Registry %q (Resource Group %q): %+v", name, containerRegistryName, resourceGroup, err)
	}

	if (firstPassword != nil) {
		passwords = append(passwords, *firstPassword)
	}

	if (secondPassword != nil) {
		passwords = append(passwords, *secondPassword)
	}

	parameters := containerregistry.TokenUpdateParameters{
		TokenUpdateProperties: &containerregistry.TokenUpdateProperties{
			ScopeMapID: utils.String(scopeMapID),
			Status:     containerregistry.TokenStatus(status),
			Credentials: &containerregistry.TokenCredentialsProperties{
				Passwords: &passwords,
			},
		},
	}

	future, err := client.Update(ctx, resourceGroup, containerRegistryName, name, parameters)
	if err != nil {
		return fmt.Errorf("Error updating token %q for Container Registry %q (Resource Group %q): %+v", name, containerRegistryName, resourceGroup, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("Error waiting for update of token %q (Container Registry %q, Resource Group %q): %+v", name, containerRegistryName, resourceGroup, err)
	}

	read, err := client.Get(ctx, resourceGroup, containerRegistryName, name)
	if err != nil {
		return fmt.Errorf("Error retrieving token %q (Container Registry %q, Resource Group %q): %+v", name, containerRegistryName, resourceGroup, err)
	}

	if read.ID == nil {
		return fmt.Errorf("Cannot read token %q (Container Registry %q, resource group %q) ID", name, containerRegistryName, resourceGroup)
	}

	d.SetId(*read.ID)

	return resourceArmContainerRegistryTokenRead(d, meta)
}

func resourceArmContainerRegistryTokenRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Containers.TokensClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := azure.ParseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resourceGroup := id.ResourceGroup
	containerRegistryName := id.Path["container_registry_name"]
	name := id.Path["name"]

	resp, err := client.Get(ctx, resourceGroup, containerRegistryName, name)

	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[DEBUG] Token %q was not found in Container Registry %q in Resource Group %q", name, containerRegistryName, resourceGroup)
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error making Read request on token %q in Azure Container Registry %q (Resource Group %q): %+v", name, containerRegistryName, resourceGroup, err)
	}

	d.Set("name", resp.Name)
	d.Set("status", resp.Status)
	d.Set("resource_group_name", resourceGroup)
	d.Set("scope_map_id", resp.ScopeMapID)

	if resp.Credentials != nil && resp.Credentials.Passwords != nil {
		passwords := *resp.Credentials.Passwords

		if len(passwords) > 0 {
			firstPassword := flattenTokenPassword(passwords[0])
			d.Set("first_password", firstPassword)
		}

		if len(passwords) > 1 {
			secondPassword := flattenTokenPassword(passwords[1])
			d.Set("second_password", secondPassword)
		}
	}

	return nil
}

func resourceArmContainerRegistryTokenDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Containers.TokensClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := azure.ParseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resourceGroup := id.ResourceGroup
	containerRegistryName := id.Path["container_registry_name"]
	name := id.Path["name"]

	future, err := client.Delete(ctx, resourceGroup, containerRegistryName, name)
	if err != nil {
		if response.WasNotFound(future.Response()) {
			return nil
		}
		return fmt.Errorf("Error issuing Azure ARM delete request of Container Registry token '%s': %+v", name, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		if response.WasNotFound(future.Response()) {
			return nil
		}
		return fmt.Errorf("Error issuing Azure ARM delete request of Container Registry token '%s': %+v", name, err)
	}

	return nil
}

func expandTokenPassword(d *schema.ResourceData, key string) (*containerregistry.TokenPassword, error) {
	passwordRaw := d.Get(key).(map[string]interface{})
	
	if (passwordRaw == nil) {
		return nil, nil
	}
	var passwordName containerregistry.TokenPasswordName

	if key == "first_password" {
		passwordName = containerregistry.TokenPasswordNamePassword1
	} else if key == "second_password" {
		passwordName = containerregistry.TokenPasswordNamePassword2
	} else {
		return nil, fmt.Errorf("Invalid password key %q.  Should be 'first_password' or 'second_password'", key)
	}

	password := containerregistry.TokenPassword{
		Name: passwordName,
	}

	if v, ok := passwordRaw["expires_at"]; ok {
		t, err := time.Parse(time.RFC3339, v.(string)) // should be validated by the schema

		if err != nil {
			return nil, err
		}
		password.Expiry = &date.Time{Time: t}
	}
	return &password, nil
}

func flattenTokenPassword(password containerregistry.TokenPassword) (map[string]interface{}) {
	passwordRaw := make(map[string]interface{})

	passwordRaw["value"] = password.Value
	passwordRaw["expires_at"] = password.Expiry.Format(time.RFC3339)

	return passwordRaw
}
