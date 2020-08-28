package compute

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/suppress"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/compute/migration"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/compute/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tags"
	azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmAvailabilitySet() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmAvailabilitySetCreateUpdate,
		Read:   resourceArmAvailabilitySetRead,
		Update: resourceArmAvailabilitySetCreateUpdate,
		Delete: resourceArmAvailabilitySetDelete,
		Importer: azSchema.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.AvailabilitySetID(id)
			return err
		}),

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    migration.AvailabilitySetV0V1Schema().CoreConfigSchema().ImpliedType(),
				Upgrade: migration.AvailabilitySetV0ToV1,
				Version: 0,
			},
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

			"location": azure.SchemaLocation(),

			"platform_update_domain_count": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      5,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(1, 20),
			},

			"platform_fault_domain_count": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      3,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(1, 3),
			},

			"managed": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				ForceNew: true,
			},

			"proximity_placement_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,

				// We have to ignore case due to incorrect capitalisation of resource group name in
				// proximity placement group ID in the response we get from the API request
				//
				// todo can be removed when https://github.com/Azure/azure-sdk-for-go/issues/5699 is fixed
				DiffSuppressFunc: suppress.CaseDifference,
			},

			"tags": tags.Schema(),
		},
	}
}

func resourceArmAvailabilitySetCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	client := meta.(*clients.Client).Compute.AvailabilitySetsClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	log.Printf("[INFO] preparing arguments for AzureRM Availability Set creation.")

	name := d.Get("name").(string)
	resGroup := d.Get("resource_group_name").(string)
	availabilitySetId := parse.NewAvailabilitySetId(resGroup, name)

	if d.IsNewResource() {
		existing, err := client.Get(ctx, resGroup, name)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("Error checking for presence of existing Availability Set %q (Resource Group %q): %s", name, resGroup, err)
			}
		}

		if !utils.ResponseWasNotFound(existing.Response) {
			return tf.ImportAsExistsError("azurerm_availability_set", availabilitySetId.ID(subscriptionId))
		}
	}

	location := azure.NormalizeLocation(d.Get("location").(string))
	updateDomainCount := d.Get("platform_update_domain_count").(int)
	faultDomainCount := d.Get("platform_fault_domain_count").(int)
	managed := d.Get("managed").(bool)
	t := d.Get("tags").(map[string]interface{})

	availSet := compute.AvailabilitySet{
		Name:     &name,
		Location: &location,
		AvailabilitySetProperties: &compute.AvailabilitySetProperties{
			PlatformFaultDomainCount:  utils.Int32(int32(faultDomainCount)),
			PlatformUpdateDomainCount: utils.Int32(int32(updateDomainCount)),
		},
		Tags: tags.Expand(t),
	}

	if v, ok := d.GetOk("proximity_placement_group_id"); ok {
		availSet.AvailabilitySetProperties.ProximityPlacementGroup = &compute.SubResource{
			ID: utils.String(v.(string)),
		}
	}

	if managed {
		n := "Aligned"
		availSet.Sku = &compute.Sku{
			Name: &n,
		}
	}

	if _, err := client.CreateOrUpdate(ctx, resGroup, name, availSet); err != nil {
		return fmt.Errorf("creating/updating Availability Set %q (Resource Group %q): %+v", name, resGroup, err)
	}

	d.SetId(availabilitySetId.ID(subscriptionId))

	return resourceArmAvailabilitySetRead(d, meta)
}

func resourceArmAvailabilitySetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.AvailabilitySetsClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.AvailabilitySetID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.Name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error making Read request on Azure Availability Set %q (Resource Group %q): %+v", id.Name, id.ResourceGroup, err)
	}

	d.Set("name", resp.Name)
	d.Set("resource_group_name", id.ResourceGroup)
	if location := resp.Location; location != nil {
		d.Set("location", azure.NormalizeLocation(*location))
	}
	if resp.Sku != nil && resp.Sku.Name != nil {
		d.Set("managed", strings.EqualFold(*resp.Sku.Name, "Aligned"))
	}

	if props := resp.AvailabilitySetProperties; props != nil {
		d.Set("platform_update_domain_count", props.PlatformUpdateDomainCount)
		d.Set("platform_fault_domain_count", props.PlatformFaultDomainCount)

		if proximityPlacementGroup := props.ProximityPlacementGroup; proximityPlacementGroup != nil {
			d.Set("proximity_placement_group_id", proximityPlacementGroup.ID)
		}
	}

	return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmAvailabilitySetDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.AvailabilitySetsClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.AvailabilitySetID(d.Id())
	if err != nil {
		return err
	}

	if _, err := client.Delete(ctx, id.ResourceGroup, id.Name); err != nil {
		return fmt.Errorf("deleting Availability Set %q (Resource Group %q): %+v", id.Name, id.ResourceGroup, err)
	}

	return nil
}
