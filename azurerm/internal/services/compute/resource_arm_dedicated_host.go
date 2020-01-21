package compute

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/hashicorp/go-azure-helpers/response"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/compute/parse"
	azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tags"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmDedicatedHost() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmDedicatedHostCreate,
		Read:   resourceArmDedicatedHostRead,
		Update: resourceArmDedicatedHostUpdate,
		Delete: resourceArmDedicatedHostDelete,

		Importer: azSchema.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.DedicatedHostID(id)
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
				ValidateFunc: validateDedicatedHostName(),
			},

			"location": azure.SchemaLocation(),

			"resource_group_name": azure.SchemaResourceGroupName(),

			"host_group_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateDedicatedHostGroupName(),
			},

			// In order to follow convention for both protal and azure cli, `sku_name` here means
			// sku name only.
			"sku_name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"DSv3-Type1",
					"ESv3-Type1",
					"FSv2-Type2",
				}, false),
			},

			"platform_fault_domain": {
				Type:     schema.TypeInt,
				ForceNew: true,
				Required: true,
			},

			"auto_replace_on_failure": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"license_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(compute.DedicatedHostLicenseTypesNone),
					string(compute.DedicatedHostLicenseTypesWindowsServerHybrid),
					string(compute.DedicatedHostLicenseTypesWindowsServerPerpetual),
				}, false),
				Default: string(compute.DedicatedHostLicenseTypesNone),
			},

			"tags": tags.Schema(),
		},
	}
}

func resourceArmDedicatedHostCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.DedicatedHostsClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	hostGroupName := d.Get("host_group_name").(string)

	if features.ShouldResourcesBeImported() && d.IsNewResource() {
		existing, err := client.Get(ctx, resourceGroupName, hostGroupName, name, "")
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("Error checking for present of existing Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", name, hostGroupName, resourceGroupName, err)
			}
		}
		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_dedicated_host", *existing.ID)
		}
	}

	parameters := compute.DedicatedHost{
		Location: utils.String(azure.NormalizeLocation(d.Get("location").(string))),
		DedicatedHostProperties: &compute.DedicatedHostProperties{
			AutoReplaceOnFailure: utils.Bool(d.Get("auto_replace_on_failure").(bool)),
			LicenseType:          compute.DedicatedHostLicenseTypes(d.Get("license_type").(string)),
		},
		Sku: &compute.Sku{
			Name: utils.String(d.Get("sku_name").(string)),
		},
		Tags: tags.Expand(d.Get("tags").(map[string]interface{})),
	}
	if platformFaultDomain, ok := d.GetOk("platform_fault_domain"); ok {
		parameters.DedicatedHostProperties.PlatformFaultDomain = utils.Int32(int32(platformFaultDomain.(int)))
	}

	future, err := client.CreateOrUpdate(ctx, resourceGroupName, hostGroupName, name, parameters)
	if err != nil {
		return fmt.Errorf("Error creating Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", name, hostGroupName, resourceGroupName, err)
	}
	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("Error waiting for creation of Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", name, hostGroupName, resourceGroupName, err)
	}

	resp, err := client.Get(ctx, resourceGroupName, hostGroupName, name, "")
	if err != nil {
		return fmt.Errorf("Error retrieving Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", name, hostGroupName, resourceGroupName, err)
	}
	if resp.ID == nil {
		return fmt.Errorf("Cannot read Dedicated Host %q (Host Group Name %q / Resource Group %q) ID", name, hostGroupName, resourceGroupName)
	}
	d.SetId(*resp.ID)

	return resourceArmDedicatedHostRead(d, meta)
}

func resourceArmDedicatedHostRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.DedicatedHostsClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.DedicatedHostID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.HostGroup, id.Name, "")
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[INFO] Dedicated Host %q does not exist - removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error reading Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", id.Name, id.HostGroup, id.ResourceGroup, err)
	}

	d.Set("name", resp.Name)
	d.Set("resource_group_name", id.ResourceGroup)
	if location := resp.Location; location != nil {
		d.Set("location", azure.NormalizeLocation(*location))
	}
	d.Set("host_group_name", id.HostGroup)
	d.Set("sku_name", resp.Sku.Name)
	if props := resp.DedicatedHostProperties; props != nil {
		d.Set("platform_fault_domain", props.PlatformFaultDomain)
		d.Set("auto_replace_on_failure", props.AutoReplaceOnFailure)
		d.Set("license_type", props.LicenseType)
	}

	return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmDedicatedHostUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.DedicatedHostsClient
	ctx, cancel := timeouts.ForUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	hostGroupName := d.Get("host_group_name").(string)

	parameters := compute.DedicatedHostUpdate{
		DedicatedHostProperties: &compute.DedicatedHostProperties{
			AutoReplaceOnFailure: utils.Bool(d.Get("auto_replace_on_failure").(bool)),
			LicenseType:          compute.DedicatedHostLicenseTypes(d.Get("license_type").(string)),
		},
		Tags: tags.Expand(d.Get("tags").(map[string]interface{})),
	}

	future, err := client.Update(ctx, resourceGroupName, hostGroupName, name, parameters)
	if err != nil {
		return fmt.Errorf("Error updating Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", name, hostGroupName, resourceGroupName, err)
	}
	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("Error waiting for update of Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", name, hostGroupName, resourceGroupName, err)
	}

	return resourceArmDedicatedHostRead(d, meta)
}

func resourceArmDedicatedHostDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.DedicatedHostsClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.DedicatedHostID(d.Id())
	if err != nil {
		return err
	}

	future, err := client.Delete(ctx, id.ResourceGroup, id.HostGroup, id.Name)
	if err != nil {
		return fmt.Errorf("Error deleting Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", id.Name, id.HostGroup, id.ResourceGroup, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		if !response.WasNotFound(future.Response()) {
			return fmt.Errorf("Error waiting for deleting Dedicated Host %q (Host Group Name %q / Resource Group %q): %+v", id.Name, id.HostGroup, id.ResourceGroup, err)
		}
	}

	// API has bug, which appears to be eventually consistent. Tracked by this issue: https://github.com/Azure/azure-rest-api-specs/issues/8137
	log.Printf("[DEBUG] Waiting for Dedicated Host %q (Host Group Name %q / Resource Group %q) to disappear", id.Name, id.HostGroup, id.ResourceGroup)
	stateConf := &resource.StateChangeConf{
		Pending:                   []string{"Exists"},
		Target:                    []string{"NotFound"},
		Refresh:                   dedicatedHostDeletedRefreshFunc(ctx, client, id),
		MinTimeout:                10 * time.Second,
		ContinuousTargetOccurence: 5,
	}

	if features.SupportsCustomTimeouts() {
		stateConf.Timeout = d.Timeout(schema.TimeoutDelete)
	} else {
		stateConf.Timeout = 10 * time.Minute
	}

	if _, err = stateConf.WaitForState(); err != nil {
		return fmt.Errorf("Error waiting for Dedicated Host %q (Host Group Name %q / Resource Group %q) to become available: %+v", id.Name, id.HostGroup, id.ResourceGroup, err)
	}

	return nil
}

func dedicatedHostDeletedRefreshFunc(ctx context.Context, client *compute.DedicatedHostsClient, id *parse.DedicatedHostId) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		res, err := client.Get(ctx, id.ResourceGroup, id.HostGroup, id.Name, "")
		if err != nil {
			if utils.ResponseWasNotFound(res.Response) {
				return "NotFound", "NotFound", nil
			}
			return nil, "", fmt.Errorf("Error issuing read request in  dedicatedHostDeletedRefreshFunc: %+v", err)
		}

		return res, "Exists", nil
	}
}
