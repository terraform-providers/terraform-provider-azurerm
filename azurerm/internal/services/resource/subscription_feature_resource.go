package resource

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2015-12-01/features"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
)

const (
	Pending       = "Pending"
	Registering   = "Registering"
	Unregistering = "Unregistering"
	Registered    = "Registered"
	NotRegistered = "NotRegistered"
	Unregistered  = "Unregistered"
)

func resourceSubscriptionFeature() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Create: resourceSubscriptionFeatureCreate,
		Read:   resourceSubscriptionFeatureRead,
		Delete: resourceSubscriptionFeatureDelete,
		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.FeatureID(id)
			return err
		}),

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
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"provider_namespace": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
		},
	}
}

func resourceSubscriptionFeatureCreate(d *pluginsdk.ResourceData, meta interface{}) error {
	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	client := meta.(*clients.Client).Resource.FeaturesClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id := parse.NewFeatureID(subscriptionId, d.Get("provider_namespace").(string), d.Get("name").(string))

	existing, err := client.Get(ctx, id.ProviderNamespace, id.Name)
	if err != nil {
		return fmt.Errorf("error checking for existing feature %q: %+v", id, err)
	}

	if existing.Properties != nil && existing.Properties.State != nil {
		if strings.EqualFold(*existing.Properties.State, Pending) {
			return fmt.Errorf("%s which requires manual approval should not be managed by terraform", id)
		}
		if !strings.EqualFold(*existing.Properties.State, NotRegistered) && !strings.EqualFold(*existing.Properties.State, Unregistered) {
			return tf.ImportAsExistsError("azurerm_subscription_feature", id.ID())
		}
	}

	resp, err := client.Register(ctx, id.ProviderNamespace, id.Name)
	if err != nil {
		return fmt.Errorf("error registering feature %q: %+v", id, err)
	}

	if resp.Properties != nil && resp.Properties.State != nil {
		if strings.EqualFold(*resp.Properties.State, Pending) {
			return fmt.Errorf("%s which requires manual approval can not be managed by terraform", id)
		}
	}

	deadline, ok := ctx.Deadline()
	if !ok {
		return fmt.Errorf("context had no deadline")
	}
	stateConf := &pluginsdk.StateChangeConf{
		Pending:    []string{Registering},
		Target:     []string{Registered},
		Refresh:    featureRegisteringStateRefreshFunc(ctx, client, id),
		MinTimeout: 3 * time.Minute,
		Timeout:    time.Until(deadline),
	}

	if _, err = stateConf.WaitForStateContext(ctx); err != nil {
		return fmt.Errorf("waiting for %s registering to be completed: %+v", id, err)
	}

	d.SetId(id.ID())

	return resourceSubscriptionFeatureRead(d, meta)
}

func resourceSubscriptionFeatureRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Resource.FeaturesClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.FeatureID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ProviderNamespace, id.Name)
	if err != nil {
		return fmt.Errorf("error retrieving feature %q: %+v", id, err)
	}
	if resp.Properties != nil && resp.Properties.State != nil {
		if strings.EqualFold(*resp.Properties.State, Pending) {
			return fmt.Errorf("%s requires manual registration approval and can not be managed by terraform", id)
		}
		if !strings.EqualFold(*resp.Properties.State, Registered) {
			d.SetId("")
			return nil
		}
	}

	d.Set("name", id.Name)
	d.Set("provider_namespace", id.ProviderNamespace)
	return nil
}

func resourceSubscriptionFeatureDelete(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Resource.FeaturesClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.FeatureID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Unregister(ctx, id.ProviderNamespace, id.Name)
	if err != nil {
		return fmt.Errorf("unregistering feature %q: %+v", id, err)
	}

	if resp.Properties != nil && resp.Properties.State != nil {
		if strings.EqualFold(*resp.Properties.State, Pending) {
			return fmt.Errorf("%s requires manual registration approval and can not be managed by terraform", id)
		}
	}

	deadline, ok := ctx.Deadline()
	if !ok {
		return fmt.Errorf("context had no deadline")
	}
	stateConf := &pluginsdk.StateChangeConf{
		Pending:    []string{Unregistering},
		Target:     []string{NotRegistered, Unregistered},
		Refresh:    featureRegisteringStateRefreshFunc(ctx, client, *id),
		MinTimeout: 3 * time.Minute,
		Timeout:    time.Until(deadline),
	}

	if _, err = stateConf.WaitForStateContext(ctx); err != nil {
		return fmt.Errorf("waiting for %s to be complete unregistering: %+v", id, err)
	}

	return nil
}

func featureRegisteringStateRefreshFunc(ctx context.Context, client *features.Client, id parse.FeatureId) pluginsdk.StateRefreshFunc {
	return func() (interface{}, string, error) {
		res, err := client.Get(ctx, id.ProviderNamespace, id.Name)
		if err != nil {
			return nil, "", fmt.Errorf("retrieving %s: %+v", id, err)
		}
		if res.Properties == nil || res.Properties.State == nil {
			return nil, "", fmt.Errorf("error reading %s registering status: %+v", id, err)
		}

		return res, *res.Properties.State, nil
	}
}
