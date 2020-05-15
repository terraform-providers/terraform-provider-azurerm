package eventgrid

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2020-04-01-preview/eventgrid"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/hashicorp/go-azure-helpers/response"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/eventgrid/parse"
	azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func getEnpointTypes() []string {
	return []string{"azure_function_endpoint", "eventhub_endpoint", "hybrid_connection_endpoint", "service_bus_queue_endpoint", "service_bus_topic_endpoint", "storage_queue_endpoint", "webhook_endpoint_url"}
}

func resourceArmEventGridEventSubscription() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmEventGridEventSubscriptionCreateUpdate,
		Read:   resourceArmEventGridEventSubscriptionRead,
		Update: resourceArmEventGridEventSubscriptionCreateUpdate,
		Delete: resourceArmEventGridEventSubscriptionDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Importer: azSchema.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.EventGridEventSubscriptionID(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"scope": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"event_delivery_schema": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  string(eventgrid.EventGridSchema),
				ValidateFunc: validation.StringInSlice([]string{
					string(eventgrid.EventGridSchema),
					string(eventgrid.CloudEventSchemaV10),
					string(eventgrid.CustomInputSchema),
				}, false),
			},

			"expiration_time_utc": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"topic_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"azure_function_endpoint": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: utils.RemoveFromStringArray(getEnpointTypes(), "azure_function_endpoint"),
				ValidateFunc:  azure.ValidateResourceID,
			},

			"eventhub_endpoint": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: utils.RemoveFromStringArray(getEnpointTypes(), "eventhub_endpoint"),
				ValidateFunc:  azure.ValidateResourceID,
			},

			"hybrid_connection_endpoint": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: utils.RemoveFromStringArray(getEnpointTypes(), "hybrid_connection_endpoint"),
				ValidateFunc:  azure.ValidateResourceID,
			},

			"service_bus_queue_endpoint": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: utils.RemoveFromStringArray(getEnpointTypes(), "service_bus_queue_endpoint"),
				ValidateFunc:  azure.ValidateResourceID,
			},

			"service_bus_topic_endpoint": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: utils.RemoveFromStringArray(getEnpointTypes(), "service_bus_topic_endpoint"),
				ValidateFunc:  azure.ValidateResourceID,
			},

			"storage_queue_endpoint": {
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				ConflictsWith: utils.RemoveFromStringArray(getEnpointTypes(), "storage_queue_endpoint"),
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"storage_account_id": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: azure.ValidateResourceID,
						},
						"queue_name": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
					},
				},
			},

			"webhook_endpoint_url": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: utils.RemoveFromStringArray(getEnpointTypes(), "webhook_endpoint_url"),
				ValidateFunc:  validation.IsURLWithHTTPS,
			},

			"included_event_types": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},

			"subject_filter": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subject_begins_with": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"subject_ends_with": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"case_sensitive": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"storage_blob_dead_letter_destination": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"storage_account_id": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: azure.ValidateResourceID,
						},
						"storage_blob_container_name": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
					},
				},
			},

			"retry_policy": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_delivery_attempts": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(1, 30),
						},
						"event_time_to_live": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(1, 1440),
						},
					},
				},
			},

			"labels": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceArmEventGridEventSubscriptionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).EventGrid.EventSubscriptionsClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	scope := d.Get("scope").(string)

	if features.ShouldResourcesBeImported() && d.IsNewResource() {
		existing, err := client.Get(ctx, scope, name)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("Error checking for presence of existing EventGrid Event Subscription %q (Scope %q): %s", name, scope, err)
			}
		}

		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_eventgrid_event_subscription", *existing.ID)
		}
	}

	destination := expandEventGridEventSubscriptionDestination(d)
	if destination == nil {
		return fmt.Errorf("One of the following endpoint types must be specificed to create an EventGrid Event Subscription: %q", getEnpointTypes())
	}

	filter := expandEventGridEventSubscriptionFilter(d)

	parsedTime, err := date.ParseTime(time.RFC3339, d.Get("expiration_time_utc").(string))
	if err != nil {
		return fmt.Errorf("Error creating/updating EventGrid Event Subscription %q (Scope %q): %s", name, scope, err)
	}

	expirationTime := date.Time{Time: parsedTime}

	eventSubscriptionProperties := eventgrid.EventSubscriptionProperties{
		Destination:           destination,
		Filter:                filter,
		DeadLetterDestination: expandEventGridEventSubscriptionStorageBlobDeadLetterDestination(d),
		RetryPolicy:           expandEventGridEventSubscriptionRetryPolicy(d),
		Labels:                utils.ExpandStringSlice(d.Get("labels").([]interface{})),
		EventDeliverySchema:   eventgrid.EventDeliverySchema(d.Get("event_delivery_schema").(string)),
		ExpirationTimeUtc:     &expirationTime,
	}

	eventSubscription := eventgrid.EventSubscription{
		EventSubscriptionProperties: &eventSubscriptionProperties,
	}

	log.Printf("[INFO] preparing arguments for AzureRM EventGrid Event Subscription creation with Properties: %+v.", eventSubscription)

	future, err := client.CreateOrUpdate(ctx, scope, name, eventSubscription)
	if err != nil {
		return fmt.Errorf("Error creating/updating EventGrid Event Subscription %q (Scope %q): %s", name, scope, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("Error waiting for EventGrid Event Subscription %q (Scope %q) to become available: %s", name, scope, err)
	}

	read, err := client.Get(ctx, scope, name)
	if err != nil {
		return fmt.Errorf("Error retrieving EventGrid Event Subscription %q (Scope %q): %s", name, scope, err)
	}
	if read.ID == nil {
		return fmt.Errorf("Cannot read EventGrid Event Subscription %s (Scope %s) ID", name, scope)
	}

	d.SetId(*read.ID)

	return resourceArmEventGridEventSubscriptionRead(d, meta)
}

func resourceArmEventGridEventSubscriptionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).EventGrid.EventSubscriptionsClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.EventGridEventSubscriptionID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.Scope, id.Name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[WARN] EventGrid Event Subscription '%s' was not found (resource group '%s')", id.Name, id.Scope)
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error making Read request on EventGrid Event Subscription '%s': %+v", id.Name, err)
	}

	d.Set("name", resp.Name)
	d.Set("scope", id.Scope)

	if props := resp.EventSubscriptionProperties; props != nil {
		d.Set("expiration_time_utc", props.ExpirationTimeUtc.Format(time.RFC3339))
		d.Set("event_delivery_schema", string(props.EventDeliverySchema))

		if props.Topic != nil && *props.Topic != "" {
			d.Set("topic_name", props.Topic)
		}

		if storageQueueEndpoint, ok := props.Destination.AsStorageQueueEventSubscriptionDestination(); ok {
			if err := d.Set("storage_queue_endpoint", flattenEventGridEventSubscriptionStorageQueueEndpoint(storageQueueEndpoint)); err != nil {
				return fmt.Errorf("Error setting `storage_queue_endpoint` for EventGrid Event Subscription %q (Scope %q): %s", id.Name, id.Scope, err)
			}
		}
		if eventHubEndpoint, ok := props.Destination.AsEventHubEventSubscriptionDestination(); ok {
			if err := d.Set("eventhub_endpoint", flattenEventGridEventSubscriptionEventHubEndpoint(eventHubEndpoint)); err != nil {
				return fmt.Errorf("Error setting `eventhub_endpoint` for EventGrid Event Subscription %q (Scope %q): %s", id.Name, id.Scope, err)
			}
		}
		if hybridConnectionEndpoint, ok := props.Destination.AsHybridConnectionEventSubscriptionDestination(); ok {
			if err := d.Set("hybrid_connection_endpoint", flattenEventGridEventSubscriptionHybridConnectionEndpoint(hybridConnectionEndpoint)); err != nil {
				return fmt.Errorf("Error setting `hybrid_connection_endpoint` for EventGrid Event Subscription %q (Scope %q): %s", id.Name, id.Scope, err)
			}
		}
		if serviceBusQueueEndpoint, ok := props.Destination.AsServiceBusQueueEventSubscriptionDestination(); ok {
			if err := d.Set("service_bus_queue_endpoint", flattenEventGridEventSubscriptionServiceBusQueueEndpoint(serviceBusQueueEndpoint)); err != nil {
				return fmt.Errorf("Error setting `service_bus_queue_endpoint` for EventGrid Event Subscription %q (Scope %q): %s", id.Name, id.Scope, err)
			}
		}
		if serviceBusTopicEndpoint, ok := props.Destination.AsServiceBusTopicEventSubscriptionDestination(); ok {
			if err := d.Set("service_bus_topic_endpoint", flattenEventGridEventSubscriptionServiceBusTopicEndpoint(serviceBusTopicEndpoint)); err != nil {
				return fmt.Errorf("Error setting `service_bus_topic_endpoint` for EventGrid Event Subscription %q (Scope %q): %s", id.Name, id.Scope, err)
			}
		}
		if azureFunctionEndpoint, ok := props.Destination.AsAzureFunctionEventSubscriptionDestination(); ok {
			if err := d.Set("azure_function_endpoint", flattenEventGridEventSubscriptionAzureFunctionEndpoint(azureFunctionEndpoint)); err != nil {
				return fmt.Errorf("Error setting `azure_function_endpoint` for EventGrid Event Subscription %q (Scope %q): %s", id.Name, id.Scope, err)
			}
		}
		if _, ok := props.Destination.AsWebHookEventSubscriptionDestination(); ok {
			fullURL, err := client.GetFullURL(ctx, id.Scope, id.Name)
			if err != nil {
				return fmt.Errorf("Error making Read request on EventGrid Event Subscription full URL '%s': %+v", id.Name, err)
			}
			if err := d.Set("webhook_endpoint_url", flattenEventGridEventSubscriptionWebhookEndpoint(&fullURL)); err != nil {
				return fmt.Errorf("Error setting `webhook_endpoint_url` for EventGrid Event Subscription %q (Scope %q): %s", id.Name, id.Scope, err)
			}
		}

		if filter := props.Filter; filter != nil {
			d.Set("included_event_types", filter.IncludedEventTypes)
			if err := d.Set("subject_filter", flattenEventGridEventSubscriptionSubjectFilter(filter)); err != nil {
				return fmt.Errorf("Error setting `subject_filter` for EventGrid Event Subscription %q (Scope %q): %s", id.Name, id.Scope, err)
			}
		}

		if props.DeadLetterDestination != nil {
			if storageBlobDeadLetterDestination, ok := props.DeadLetterDestination.AsStorageBlobDeadLetterDestination(); ok {
				if err := d.Set("storage_blob_dead_letter_destination", flattenEventGridEventSubscriptionStorageBlobDeadLetterDestination(storageBlobDeadLetterDestination)); err != nil {
					return fmt.Errorf("Error setting `storage_blob_dead_letter_destination` for EventGrid Event Subscription %q (Scope %q): %s", id.Name, id.Scope, err)
				}
			}
		}

		if retryPolicy := props.RetryPolicy; retryPolicy != nil {
			if err := d.Set("retry_policy", flattenEventGridEventSubscriptionRetryPolicy(retryPolicy)); err != nil {
				return fmt.Errorf("Error setting `retry_policy` for EventGrid Event Subscription %q (Scope %q): %s", id.Name, id.Scope, err)
			}
		}

		if err := d.Set("labels", props.Labels); err != nil {
			return fmt.Errorf("Error setting `labels` for EventGrid Event Subscription %q (Scope %q): %s", id.Name, id.Scope, err)
		}
	}

	return nil
}

func resourceArmEventGridEventSubscriptionDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).EventGrid.EventSubscriptionsClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.EventGridEventSubscriptionID(d.Id())
	if err != nil {
		return err
	}

	future, err := client.Delete(ctx, id.Scope, id.Name)
	if err != nil {
		if response.WasNotFound(future.Response()) {
			return nil
		}
		return fmt.Errorf("Error deleting Event Grid Event Subscription %q: %+v", id.Name, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		if response.WasNotFound(future.Response()) {
			return nil
		}
		return fmt.Errorf("Error deleting Event Grid Event Subscription %q: %+v", id.Name, err)
	}

	return nil
}

func expandEventGridEventSubscriptionDestination(d *schema.ResourceData) eventgrid.BasicEventSubscriptionDestination {
	if _, ok := d.GetOk("azure_function_endpoint"); ok {
		return expandEventGridEventSubscriptionAzureFunctionEndpoint(d)
	}

	if _, ok := d.GetOk("eventhub_endpoint"); ok {
		return expandEventGridEventSubscriptionEventHubEndpoint(d)
	}

	if _, ok := d.GetOk("hybrid_connection_endpoint"); ok {
		return expandEventGridEventSubscriptionHybridConnectionEndpoint(d)
	}

	if _, ok := d.GetOk("service_bus_queue_endpoint"); ok {
		return expandEventGridEventSubscriptionServiceBusQueueEndpoint(d)
	}

	if _, ok := d.GetOk("service_bus_topic_endpoint"); ok {
		return expandEventGridEventSubscriptionServiceBusTopicEndpoint(d)
	}

	if _, ok := d.GetOk("storage_queue_endpoint"); ok {
		return expandEventGridEventSubscriptionStorageQueueEndpoint(d)
	}

	if _, ok := d.GetOk("webhook_endpoint_url"); ok {
		return expandEventGridEventSubscriptionWebhookEndpoint(d)
	}

	return nil
}

func expandEventGridEventSubscriptionStorageQueueEndpoint(d *schema.ResourceData) eventgrid.BasicEventSubscriptionDestination {
	props := d.Get("storage_queue_endpoint").([]interface{})[0].(map[string]interface{})
	storageAccountID := props["storage_account_id"].(string)
	queueName := props["queue_name"].(string)

	storageQueueEndpoint := eventgrid.StorageQueueEventSubscriptionDestination{
		EndpointType: eventgrid.EndpointTypeStorageQueue,
		StorageQueueEventSubscriptionDestinationProperties: &eventgrid.StorageQueueEventSubscriptionDestinationProperties{
			ResourceID: &storageAccountID,
			QueueName:  &queueName,
		},
	}
	return storageQueueEndpoint
}

func expandEventGridEventSubscriptionEventHubEndpoint(d *schema.ResourceData) eventgrid.BasicEventSubscriptionDestination {
	eventHubID := d.Get("eventhub_endpoint").(string)

	eventHubEndpoint := eventgrid.EventHubEventSubscriptionDestination{
		EndpointType: eventgrid.EndpointTypeEventHub,
		EventHubEventSubscriptionDestinationProperties: &eventgrid.EventHubEventSubscriptionDestinationProperties{
			ResourceID: &eventHubID,
		},
	}
	return eventHubEndpoint
}

func expandEventGridEventSubscriptionHybridConnectionEndpoint(d *schema.ResourceData) eventgrid.BasicEventSubscriptionDestination {
	hybridConnectionID := d.Get("hybrid_connection_endpoint").(string)

	hybridConnectionEndpoint := eventgrid.HybridConnectionEventSubscriptionDestination{
		EndpointType: eventgrid.EndpointTypeHybridConnection,
		HybridConnectionEventSubscriptionDestinationProperties: &eventgrid.HybridConnectionEventSubscriptionDestinationProperties{
			ResourceID: &hybridConnectionID,
		},
	}
	return hybridConnectionEndpoint
}

func expandEventGridEventSubscriptionWebhookEndpoint(d *schema.ResourceData) eventgrid.BasicEventSubscriptionDestination {
	url := d.Get("webhook_endpoint_url").(string)

	webhookEndpoint := eventgrid.WebHookEventSubscriptionDestination{
		EndpointType: eventgrid.EndpointTypeWebHook,
		WebHookEventSubscriptionDestinationProperties: &eventgrid.WebHookEventSubscriptionDestinationProperties{
			EndpointURL: &url,
		},
	}
	return webhookEndpoint
}

func expandEventGridEventSubscriptionServiceBusQueueEndpoint(d *schema.ResourceData) eventgrid.BasicEventSubscriptionDestination {
	serviceBusQueueID := d.Get("service_bus_queue_endpoint").(string)

	serviceBusQueueEndpoint := eventgrid.ServiceBusQueueEventSubscriptionDestination{
		EndpointType: eventgrid.EndpointTypeServiceBusQueue,
		ServiceBusQueueEventSubscriptionDestinationProperties: &eventgrid.ServiceBusQueueEventSubscriptionDestinationProperties{
			ResourceID: &serviceBusQueueID,
		},
	}
	return serviceBusQueueEndpoint
}

func expandEventGridEventSubscriptionServiceBusTopicEndpoint(d *schema.ResourceData) eventgrid.BasicEventSubscriptionDestination {
	serviceBusTopicID := d.Get("service_bus_topic_endpoint").(string)

	serviceBusTopicEndpoint := eventgrid.ServiceBusTopicEventSubscriptionDestination{
		EndpointType: eventgrid.EndpointTypeServiceBusTopic,
		ServiceBusTopicEventSubscriptionDestinationProperties: &eventgrid.ServiceBusTopicEventSubscriptionDestinationProperties{
			ResourceID: &serviceBusTopicID,
		},
	}
	return serviceBusTopicEndpoint
}

func expandEventGridEventSubscriptionAzureFunctionEndpoint(d *schema.ResourceData) eventgrid.BasicEventSubscriptionDestination {
	azureFunctionResourceID := d.Get("azure_function_endpoint").(string)

	azureFunctionEndpoint := eventgrid.AzureFunctionEventSubscriptionDestination{
		EndpointType: eventgrid.EndpointTypeAzureFunction,
		AzureFunctionEventSubscriptionDestinationProperties: &eventgrid.AzureFunctionEventSubscriptionDestinationProperties{
			ResourceID: &azureFunctionResourceID,
		},
	}
	return azureFunctionEndpoint
}

func expandEventGridEventSubscriptionFilter(d *schema.ResourceData) *eventgrid.EventSubscriptionFilter {
	filter := &eventgrid.EventSubscriptionFilter{}

	if includedEvents, ok := d.GetOk("included_event_types"); ok {
		filter.IncludedEventTypes = utils.ExpandStringSlice(includedEvents.([]interface{}))
	}

	if subjectFilter, ok := d.GetOk("subject_filter"); ok {
		config := subjectFilter.([]interface{})[0].(map[string]interface{})
		subjectBeginsWith := config["subject_begins_with"].(string)
		subjectEndsWith := config["subject_ends_with"].(string)
		caseSensitive := config["case_sensitive"].(bool)

		filter.SubjectBeginsWith = &subjectBeginsWith
		filter.SubjectEndsWith = &subjectEndsWith
		filter.IsSubjectCaseSensitive = &caseSensitive
	}

	return filter
}

func expandEventGridEventSubscriptionStorageBlobDeadLetterDestination(d *schema.ResourceData) eventgrid.BasicDeadLetterDestination {
	if v, ok := d.GetOk("storage_blob_dead_letter_destination"); ok {
		dest := v.([]interface{})[0].(map[string]interface{})
		resourceID := dest["storage_account_id"].(string)
		blobName := dest["storage_blob_container_name"].(string)
		return eventgrid.StorageBlobDeadLetterDestination{
			EndpointType: eventgrid.EndpointTypeStorageBlob,
			StorageBlobDeadLetterDestinationProperties: &eventgrid.StorageBlobDeadLetterDestinationProperties{
				ResourceID:        &resourceID,
				BlobContainerName: &blobName,
			},
		}
	}
	return nil
}

func expandEventGridEventSubscriptionRetryPolicy(d *schema.ResourceData) *eventgrid.RetryPolicy {
	if v, ok := d.GetOk("retry_policy"); ok {
		dest := v.([]interface{})[0].(map[string]interface{})
		maxDeliveryAttempts := dest["max_delivery_attempts"].(int)
		eventTimeToLive := dest["event_time_to_live"].(int)
		return &eventgrid.RetryPolicy{
			MaxDeliveryAttempts:      utils.Int32(int32(maxDeliveryAttempts)),
			EventTimeToLiveInMinutes: utils.Int32(int32(eventTimeToLive)),
		}
	}
	return nil
}

func flattenEventGridEventSubscriptionEventHubEndpoint(input *eventgrid.EventHubEventSubscriptionDestination) *string {
	if input == nil || input.ResourceID == nil {
		return nil
	}

	return input.ResourceID
}

func flattenEventGridEventSubscriptionHybridConnectionEndpoint(input *eventgrid.HybridConnectionEventSubscriptionDestination) *string {
	if input == nil || input.ResourceID == nil {
		return nil
	}

	return input.ResourceID
}

func flattenEventGridEventSubscriptionServiceBusQueueEndpoint(input *eventgrid.ServiceBusQueueEventSubscriptionDestination) *string {
	if input == nil || input.ResourceID == nil {
		return nil
	}

	return input.ResourceID
}

func flattenEventGridEventSubscriptionServiceBusTopicEndpoint(input *eventgrid.ServiceBusTopicEventSubscriptionDestination) *string {
	if input == nil || input.ResourceID == nil {
		return nil
	}

	return input.ResourceID
}

func flattenEventGridEventSubscriptionAzureFunctionEndpoint(input *eventgrid.AzureFunctionEventSubscriptionDestination) *string {
	if input == nil || input.ResourceID == nil {
		return nil
	}

	return input.ResourceID
}

func flattenEventGridEventSubscriptionStorageQueueEndpoint(input *eventgrid.StorageQueueEventSubscriptionDestination) []interface{} {
	if input == nil {
		return nil
	}
	result := make(map[string]interface{})

	if input.ResourceID != nil {
		result["storage_account_id"] = *input.ResourceID
	}
	if input.QueueName != nil {
		result["queue_name"] = *input.QueueName
	}

	return []interface{}{result}
}

func flattenEventGridEventSubscriptionWebhookEndpoint(input *eventgrid.EventSubscriptionFullURL) *string {
	if input == nil || input.EndpointURL == nil {
		return nil
	}

	return input.EndpointURL
}

func flattenEventGridEventSubscriptionSubjectFilter(filter *eventgrid.EventSubscriptionFilter) []interface{} {
	if (filter.SubjectBeginsWith != nil && *filter.SubjectBeginsWith == "") && (filter.SubjectEndsWith != nil && *filter.SubjectEndsWith == "") {
		return nil
	}
	result := make(map[string]interface{})

	if filter.SubjectBeginsWith != nil {
		result["subject_begins_with"] = *filter.SubjectBeginsWith
	}

	if filter.SubjectEndsWith != nil {
		result["subject_ends_with"] = *filter.SubjectEndsWith
	}

	if filter.IsSubjectCaseSensitive != nil {
		result["case_sensitive"] = *filter.IsSubjectCaseSensitive
	}

	return []interface{}{result}
}

func flattenEventGridEventSubscriptionStorageBlobDeadLetterDestination(dest *eventgrid.StorageBlobDeadLetterDestination) []interface{} {
	if dest == nil {
		return nil
	}
	result := make(map[string]interface{})

	if dest.ResourceID != nil {
		result["storage_account_id"] = *dest.ResourceID
	}

	if dest.BlobContainerName != nil {
		result["storage_blob_container_name"] = *dest.BlobContainerName
	}

	return []interface{}{result}
}

func flattenEventGridEventSubscriptionRetryPolicy(retryPolicy *eventgrid.RetryPolicy) []interface{} {
	result := make(map[string]interface{})

	if v := retryPolicy.EventTimeToLiveInMinutes; v != nil {
		result["event_time_to_live"] = int(*v)
	}

	if v := retryPolicy.MaxDeliveryAttempts; v != nil {
		result["max_delivery_attempts"] = int(*v)
	}

	return []interface{}{result}
}
