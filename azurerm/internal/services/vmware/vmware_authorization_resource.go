package vmware

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/avs/mgmt/2020-03-20/avs"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/vmware/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/vmware/validate"
	azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceVmwareAuthorization() *schema.Resource {
	return &schema.Resource{
		Create: resourceVmwareAuthorizationCreate,
		Read:   resourceVmwareAuthorizationRead,
		Delete: resourceVmwareAuthorizationDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Importer: azSchema.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.AuthorizationID(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"private_cloud_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.PrivateCloudID,
			},

			"express_route_authorization_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"express_route_authorization_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func resourceVmwareAuthorizationCreate(d *schema.ResourceData, meta interface{}) error {
	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	client := meta.(*clients.Client).Vmware.AuthorizationClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	privateCloudId, err := parse.PrivateCloudID(d.Get("private_cloud_id").(string))
	if err != nil {
		return err
	}

	id := parse.NewAuthorizationID(subscriptionId, privateCloudId.ResourceGroup, privateCloudId.Name, name)

	if d.IsNewResource() {
		existing, err := client.Get(ctx, id.ResourceGroup, id.PrivateCloudName, id.Name)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("checking for present of existing %q : %+v", id.ID(), err)
			}
		}
		if !utils.ResponseWasNotFound(existing.Response) {
			return tf.ImportAsExistsError("azurerm_vmware_authorization", id.ID())
		}
	}

	props := avs.ExpressRouteAuthorization{}

	future, err := client.CreateOrUpdate(ctx, privateCloudId.ResourceGroup, privateCloudId.Name, name, props)
	if err != nil {
		return fmt.Errorf("creating %q: %+v", id, err)
	}

	if err := future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for the creation of the %q: %+v", id, err)
	}

	d.SetId(id.ID())
	return resourceVmwareAuthorizationRead(d, meta)
}

func resourceVmwareAuthorizationRead(d *schema.ResourceData, meta interface{}) error {
	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	client := meta.(*clients.Client).Vmware.AuthorizationClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.AuthorizationID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.PrivateCloudName, id.Name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[INFO] avs %q does not exist - removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("retrieving %q: %+v", id, err)
	}

	d.Set("name", id.Name)
	d.Set("private_cloud_id", parse.NewPrivateCloudID(subscriptionId, id.ResourceGroup, id.PrivateCloudName).ID())
	if props := resp.ExpressRouteAuthorizationProperties; props != nil {
		d.Set("express_route_authorization_id", props.ExpressRouteAuthorizationID)
		d.Set("express_route_authorization_key", props.ExpressRouteAuthorizationKey)
	}
	return nil
}

func resourceVmwareAuthorizationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Vmware.AuthorizationClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.AuthorizationID(d.Id())
	if err != nil {
		return err
	}

	future, err := client.Delete(ctx, id.ResourceGroup, id.PrivateCloudName, id.Name)
	if err != nil {
		return fmt.Errorf("deleting %q: %+v", id, err)
	}

	if err := future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for the deletion of the %q: %+v", id, err)
	}
	return nil
}
