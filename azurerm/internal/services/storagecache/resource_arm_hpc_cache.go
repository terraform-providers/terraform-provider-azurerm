package storagecache

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/storagecache/mgmt/2019-11-01/storagecache"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmHPCCache() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmHPCCacheCreate,
		Read:   resourceArmHPCCacheRead,
		Delete: resourceArmHPCCacheDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"resource_group_name": azure.SchemaResourceGroupName(),

			"location": azure.SchemaLocation(),

			"cache_size": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.IntInSlice([]int{
					3072,
					6144,
					12288,
					24576,
					49152,
				}),
			},

			"subnet_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: azure.ValidateResourceIDOrEmpty,
			},

			"sku_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Standard_2G",
					"Standard_4G",
					"Standard_8G",
				}, false),
			},
		},
	}
}

func resourceArmHPCCacheCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).StorageCache.CachesClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	log.Printf("[INFO] preparing arguments for Azure HPC Cache creation.")
	name := d.Get("name").(string)
	resourceGroup := d.Get("resource_group_name").(string)

	if features.ShouldResourcesBeImported() && d.IsNewResource() {
		existing, err := client.Get(ctx, resourceGroup, name)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("Error checking for presence of existing HPC Cache %q (Resource Group %q): %s", name, resourceGroup, err)
			}
		}

		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_hpc_cache", *existing.ID)
		}
	}

	location := d.Get("location").(string)
	cacheSize := d.Get("cache_size").(int)
	subnet := d.Get("subnet_id").(string)
	skuName := d.Get("sku_name").(string)

	cache := &storagecache.Cache{
		Name:     utils.String(name),
		Location: utils.String(location),
		CacheProperties: &storagecache.CacheProperties{
			CacheSizeGB: utils.Int32(int32(cacheSize)),
			Subnet:      utils.String(subnet),
		},
		Sku: &storagecache.CacheSku{
			Name: utils.String(skuName),
		},
	}

	if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, cache); err != nil {
		return fmt.Errorf("Error creating HPC Cache %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{string(storagecache.Creating)},
		Target:     []string{string(storagecache.Succeeded)},
		MinTimeout: 30 * time.Second,
		Refresh: func() (interface{}, string, error) {
			resp, err := client.Get(ctx, resourceGroup, name)
			if err != nil {
				return resp, "Error", fmt.Errorf("Error retrieving HPC Cache %q (Resource Group %q): %+v", resourceGroup, name, err)
			}

			if properties := resp.CacheProperties; properties != nil {
				return resp, string(properties.ProvisioningState), nil
			}

			return resp, "Unknown", nil
		},
	}

	if features.SupportsCustomTimeouts() {
		if d.IsNewResource() {
			stateConf.Timeout = d.Timeout(schema.TimeoutCreate)
		}
	} else {
		stateConf.Timeout = 30 * time.Minute
	}

	_, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for HPC Cache %q (Resource Group %q) to finish provisioning: %+v", resourceGroup, name, err)
	}

	resp, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		return fmt.Errorf("Error retrieving HPC Cache %q (Resource Group %q): %+v", name, resourceGroup, err)
	}
	if resp.ID == nil {
		return fmt.Errorf("Cannot read ID for HPC Cache %q (Resource Group %q): %+v", name, resourceGroup, err)
	}
	d.SetId(*resp.ID)

	return resourceArmHPCCacheRead(d, meta)
}

func resourceArmHPCCacheRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).StorageCache.CachesClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := azure.ParseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resourceGroup := id.ResourceGroup
	name := id.Path["caches"]

	resp, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[DEBUG] HPC Cache %q was not found in Resource Group %q - removing from state!", name, resourceGroup)
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error retrieving HPC Cache %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	d.Set("name", name)
	d.Set("resource_group_name", resourceGroup)
	d.Set("location", resp.Location)

	if props := resp.CacheProperties; props != nil {
		d.Set("cache_size", props.CacheSizeGB)
		d.Set("subnet_id", props.Subnet)
	}

	if sku := resp.Sku; sku != nil {
		d.Set("sku_name", sku.Name)
	}

	return nil
}

func resourceArmHPCCacheDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).StorageCache.CachesClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := azure.ParseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resourceGroup := id.ResourceGroup
	name := id.Path["caches"]

	future, err := client.Delete(ctx, resourceGroup, name)
	if err != nil {
		return fmt.Errorf("Error deleting HPC Cache %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("Error waiting for deletion of HPC Cache %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	return nil
}
