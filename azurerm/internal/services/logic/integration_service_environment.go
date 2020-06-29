package logic

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2019-05-01/logic"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/set"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/logic/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tags"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmIntegrationServiceEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmIntegrationServiceEnvironmentCreateUpdate,
		Read:   resourceArmIntegrationServiceEnvironmentRead,
		Update: resourceArmIntegrationServiceEnvironmentCreateUpdate,
		Delete: resourceArmIntegrationServiceEnvironmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Hour),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Hour),
			Delete: schema.DefaultTimeout(4 * time.Hour),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.IntegrationServiceEnvironmentName(),
			},

			"location": azure.SchemaLocation(),

			"resource_group_name": azure.SchemaResourceGroupName(),

			"sku_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true, // The SKU cannot be changed once integration service environment has been provisioned.
				ValidateFunc: validation.StringInSlice([]string{
					string(logic.IntegrationServiceEnvironmentSkuNameDeveloper),
					string(logic.IntegrationServiceEnvironmentSkuNamePremium),
				}, false),
			},

			// Maximum scale units that you can add	10 - https://docs.microsoft.com/en-US/azure/logic-apps/logic-apps-limits-and-config#integration-service-environment-ise
			"capacity": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      0,
				ValidateFunc: validation.IntBetween(0, 10),
			},

			"access_endpoint_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true, // The access end point type cannot be changed once the integration service environment is provisioned.
				ValidateFunc: validation.StringInSlice([]string{
					string(logic.IntegrationServiceEnvironmentAccessEndpointTypeInternal),
					string(logic.IntegrationServiceEnvironmentAccessEndpointTypeExternal),
				}, false),
			},

			"virtual_network_subnet_ids": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true, // The network configuration subnets cannot be updated after integration service environment is created.
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validate.ValidateSubnetID,
				},
				Set:      set.HashStringIgnoreCase,
				MinItems: 4,
				MaxItems: 4,
			},

			"connector_endpoint_ip_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"connector_outbound_ip_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"workflow_endpoint_ip_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"workflow_outbound_ip_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"tags": tags.Schema(),
		},
	}
}

func resourceArmIntegrationServiceEnvironmentCreateUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*clients.Client).Logic.IntegrationServiceEnvironmentsClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	log.Printf("[INFO] preparing arguments for Azure ARM Integration Service Environment creation.")

	name := d.Get("name").(string)
	resourceGroup := d.Get("resource_group_name").(string)

	if features.ShouldResourcesBeImported() && d.IsNewResource() {
		existing, err := client.Get(ctx, resourceGroup, name)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("Error checking for presence of existing Integration Service Environment %q (Resource Group %q): %s", name, resourceGroup, err)
			}
		}

		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_integration_service_environment", *existing.ID)
		}
	}

	location := azure.NormalizeLocation(d.Get("location").(string))
	skuName := d.Get("sku_name").(string)
	capacity := d.Get("capacity").(int)
	accessEndpointType := d.Get("access_endpoint_type").(string)
	virtualNetworkSubnetIds := d.Get("virtual_network_subnet_ids").(*schema.Set).List()
	t := d.Get("tags").(map[string]interface{})

	properties := logic.IntegrationServiceEnvironment{
		Name:     &name,
		Location: &location,
		Properties: &logic.IntegrationServiceEnvironmentProperties{
			NetworkConfiguration: &logic.NetworkConfiguration{
				AccessEndpoint: &logic.IntegrationServiceEnvironmentAccessEndpoint{
					Type: logic.IntegrationServiceEnvironmentAccessEndpointType(accessEndpointType),
				},
				Subnets: expandSubnetResourceID(virtualNetworkSubnetIds),
			},
		},
		Sku: &logic.IntegrationServiceEnvironmentSku{
			Name:     logic.IntegrationServiceEnvironmentSkuName(skuName),
			Capacity: utils.Int32(int32(capacity)),
		},
		Tags: tags.Expand(t),
	}

	future, err := client.CreateOrUpdate(ctx, resourceGroup, name, properties)
	if err != nil {
		return fmt.Errorf("creating/updating Integration Service Environment %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for completion of Integration Service Environment %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	read, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		return fmt.Errorf("read request on Integration Service Environment %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	if read.ID == nil {
		return fmt.Errorf("cannot read Integration Service Environment %q (Resource Group %q) ID", name, resourceGroup)
	}

	d.SetId(*read.ID)

	return resourceArmIntegrationServiceEnvironmentRead(d, meta)
}

func resourceArmIntegrationServiceEnvironmentRead(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceArmIntegrationServiceEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func expandSubnetResourceID(input []interface{}) *[]logic.ResourceReference {
	results := make([]logic.ResourceReference, 0)
	for _, item := range input {
		id := item.(string)

		results = append(results, logic.ResourceReference{
			ID: utils.String(id),
		})
	}
	return &results
}
