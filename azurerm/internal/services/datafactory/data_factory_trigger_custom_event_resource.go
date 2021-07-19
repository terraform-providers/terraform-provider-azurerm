package datafactory

import (
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/datafactory/mgmt/2018-06-01/datafactory"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/datafactory/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/datafactory/validate"
	eventgridValidate "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/eventgrid/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceDataFactoryTriggerCustomEvent() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Create: resourceDataFactoryTriggerCustomEventCreateUpdate,
		Read:   resourceDataFactoryTriggerCustomEventRead,
		Update: resourceDataFactoryTriggerCustomEventCreateUpdate,
		Delete: resourceDataFactoryTriggerCustomEventDelete,

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.TriggerID(id)
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
				ValidateFunc: validate.DataFactoryPipelineAndTriggerName(),
			},

			"data_factory_id": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.DataFactoryID,
			},

			"eventgrid_topic_id": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: eventgridValidate.TopicID,
			},

			"events": {
				Type:     pluginsdk.TypeSet,
				Required: true,
				MinItems: 1,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"pipeline": {
				Type:     pluginsdk.TypeSet,
				Required: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"name": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validate.DataFactoryPipelineAndTriggerName(),
						},

						"parameters": {
							Type:     pluginsdk.TypeMap,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},
					},
				},
			},

			"additional_properties": {
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
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},

			"description": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"subject_begins_with": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				AtLeastOneOf: []string{"subject_begins_with", "subject_ends_with"},
			},

			"subject_ends_with": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				AtLeastOneOf: []string{"subject_begins_with", "subject_ends_with"},
			},
		},
	}
}

func resourceDataFactoryTriggerCustomEventCreateUpdate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DataFactory.TriggersClient
	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	dataFactoryId, err := parse.DataFactoryID(d.Get("data_factory_id").(string))
	if err != nil {
		return err
	}

	id := parse.NewTriggerID(subscriptionId, dataFactoryId.ResourceGroup, dataFactoryId.FactoryName, d.Get("name").(string))
	if d.IsNewResource() {
		existing, err := client.Get(ctx, id.ResourceGroup, id.FactoryName, id.Name, "")
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("checking for presence of existing %s: %+v", id, err)
			}
		}
		if !utils.ResponseWasNotFound(existing.Response) {
			return tf.ImportAsExistsError("azurerm_data_factory_trigger_custom_event", id.ID())
		}
	}

	events := d.Get("events").(*pluginsdk.Set).List()
	trigger := &datafactory.CustomEventsTrigger{
		CustomEventsTriggerTypeProperties: &datafactory.CustomEventsTriggerTypeProperties{
			Events: &events,
			Scope:  utils.String(d.Get("eventgrid_topic_id").(string)),
		},
		Description: utils.String(d.Get("description").(string)),
		Pipelines:   expandDataFactoryTriggerPipeline(d.Get("pipeline").(*pluginsdk.Set).List()),
		Type:        datafactory.TypeBasicTriggerTypeCustomEventsTrigger,
	}

	if v, ok := d.GetOk("annotations"); ok {
		annotations := v.([]interface{})
		trigger.Annotations = &annotations
	}

	if v, ok := d.GetOk("subject_begins_with"); ok {
		trigger.SubjectBeginsWith = utils.String(v.(string))
	}

	if v, ok := d.GetOk("subject_ends_with"); ok {
		trigger.SubjectEndsWith = utils.String(v.(string))
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		trigger.AdditionalProperties = v.(map[string]interface{})
	}

	resource := datafactory.TriggerResource{
		Properties: trigger,
	}

	if _, err := client.CreateOrUpdate(ctx, id.ResourceGroup, id.FactoryName, id.Name, resource, ""); err != nil {
		return fmt.Errorf("creating %s: %+v", id, err)
	}

	d.SetId(id.ID())

	return resourceDataFactoryTriggerCustomEventRead(d, meta)
}

func resourceDataFactoryTriggerCustomEventRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DataFactory.TriggersClient
	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.TriggerID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.FactoryName, id.Name, "")
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("retrieving %s: %+v", id, err)
	}

	CustomEventsTrigger, ok := resp.Properties.AsCustomEventsTrigger()
	if !ok {
		return fmt.Errorf("classifiying %s: Expected: %q", id, datafactory.TypeBasicTriggerTypeCustomEventsTrigger)
	}

	d.Set("name", id.Name)
	d.Set("data_factory_id", parse.NewDataFactoryID(subscriptionId, id.ResourceGroup, id.FactoryName).ID())

	d.Set("additional_properties", CustomEventsTrigger.AdditionalProperties)
	d.Set("description", CustomEventsTrigger.Description)

	if err := d.Set("annotations", flattenDataFactoryAnnotations(CustomEventsTrigger.Annotations)); err != nil {
		return fmt.Errorf("setting `annotations`: %+v", err)
	}

	if err := d.Set("pipeline", flattenDataFactoryTriggerPipeline(CustomEventsTrigger.Pipelines)); err != nil {
		return fmt.Errorf("setting `pipeline`: %+v", err)
	}

	if props := CustomEventsTrigger.CustomEventsTriggerTypeProperties; props != nil {
		d.Set("eventgrid_topic_id", props.Scope)
		d.Set("events", props.Events)
		d.Set("subject_begins_with", props.SubjectBeginsWith)
		d.Set("subject_ends_with", props.SubjectEndsWith)
	}

	return nil
}

func resourceDataFactoryTriggerCustomEventDelete(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DataFactory.TriggersClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.TriggerID(d.Id())
	if err != nil {
		return err
	}

	if _, err = client.Delete(ctx, id.ResourceGroup, id.FactoryName, id.Name); err != nil {
		return fmt.Errorf("deleting %s: %+v", id, err)
	}

	return nil
}
