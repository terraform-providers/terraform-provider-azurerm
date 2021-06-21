package kusto

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/kusto/mgmt/2020-09-18/kusto"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/location"
	iothubValidate "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/iothub/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/kusto/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/kusto/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceKustoIotHubDataConnection() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Create: resourceKustoIotHubDataConnectionCreate,
		Read:   resourceKustoIotHubDataConnectionRead,
		Delete: resourceKustoIotHubDataConnectionDelete,

		Importer: pluginsdk.ImporterValidatingResourceIdThen(func(id string) error {
			_, err := parse.DataConnectionID(id)
			return err
		}, importDataConnection(kusto.KindIotHub)),

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(60 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.DataConnectionName,
			},

			"resource_group_name": azure.SchemaResourceGroupName(),

			"location": azure.SchemaLocation(),

			"cluster_name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.ClusterName,
			},

			"database_name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.DatabaseName,
			},

			"iothub_id": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: iothubValidate.IotHubID,
			},

			"consumer_group": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: iothubValidate.IoTHubConsumerGroupName,
			},

			"table_name": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validate.EntityName,
			},

			"mapping_rule_name": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validate.EntityName,
			},

			"data_format": {
				Type:     pluginsdk.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(kusto.IotHubDataFormatAPACHEAVRO),
					string(kusto.IotHubDataFormatAVRO),
					string(kusto.IotHubDataFormatCSV),
					string(kusto.IotHubDataFormatJSON),
					string(kusto.IotHubDataFormatMULTIJSON),
					string(kusto.IotHubDataFormatORC),
					string(kusto.IotHubDataFormatPARQUET),
					string(kusto.IotHubDataFormatPSV),
					string(kusto.IotHubDataFormatRAW),
					string(kusto.IotHubDataFormatSCSV),
					string(kusto.IotHubDataFormatSINGLEJSON),
					string(kusto.IotHubDataFormatSOHSV),
					string(kusto.IotHubDataFormatTSV),
					string(kusto.IotHubDataFormatTSVE),
					string(kusto.IotHubDataFormatTXT),
					string(kusto.IotHubDataFormatW3CLOGFILE),
				}, false),
			},

			"shared_access_policy_name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: iothubValidate.IotHubSharedAccessPolicyName,
			},

			"event_system_properties": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"message-id",
						"sequence-number",
						"to",
						"absolute-expiry-time",
						"iothub-enqueuedtime",
						"correlation-id",
						"user-id",
						"iothub-ack",
						"iothub-connection-device-id",
						"iothub-connection-auth-generation-id",
						"iothub-connection-auth-method",
					}, false),
				},
			},
		},
	}
}

func resourceKustoIotHubDataConnectionCreate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Kusto.DataConnectionsClient
	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	log.Printf("[INFO] preparing arguments for Azure Kusto Iot Hub Data Connection creation.")

	id := parse.NewDataConnectionID(subscriptionId, d.Get("resource_group_name").(string), d.Get("cluster_name").(string), d.Get("database_name").(string), d.Get("name").(string))
	resp, err := client.Get(ctx, id.ResourceGroup, id.ClusterName, id.DatabaseName, id.Name)
	if err != nil {
		if !utils.ResponseWasNotFound(resp.Response) {
			return fmt.Errorf("checking for presence of existing %s: %s", id, err)
		}
	}

	if !utils.ResponseWasNotFound(resp.Response) {
		return tf.ImportAsExistsError("azurerm_kusto_iothub_data_connection", id.ID())
	}

	dataConnection := kusto.IotHubDataConnection{
		Location: utils.String(azure.NormalizeLocation(d.Get("location").(string))),
		IotHubConnectionProperties: &kusto.IotHubConnectionProperties{
			IotHubResourceID:       utils.String(d.Get("iothub_id").(string)),
			ConsumerGroup:          utils.String(d.Get("consumer_group").(string)),
			TableName :             utils.String(d.Get("table_name").(string)),
			MappingRuleName :       utils.String(d.Get("mapping_rule_name").(string)),
			DataFormat :            kusto.IotHubDataFormat(d.Get("data_format").(string)),
			SharedAccessPolicyName: utils.String(d.Get("shared_access_policy_name").(string)),
		},
	}

	if eventSystemProperties, ok := d.GetOk("event_system_properties"); ok {
		dataConnection.IotHubConnectionProperties.EventSystemProperties = utils.ExpandStringSlice(eventSystemProperties.(*pluginsdk.Set).List())
	}

	future, err := client.CreateOrUpdate(ctx, id.ResourceGroup, id.ClusterName, id.DatabaseName, id.Name, dataConnection)
	if err != nil {
		return fmt.Errorf("creating or updating %s: %+v", id, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for completion of %s: %+v", id, err)
	}

	d.SetId(id.ID())
	return resourceKustoIotHubDataConnectionRead(d, meta)
}

func resourceKustoIotHubDataConnectionRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Kusto.DataConnectionsClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.DataConnectionID(d.Id())
	if err != nil {
		return err
	}

	connectionModel, err := client.Get(ctx, id.ResourceGroup, id.ClusterName, id.DatabaseName, id.Name)
	if err != nil {
		if utils.ResponseWasNotFound(connectionModel.Response) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("retrieving %s: %+v", id, err)
	}

	d.Set("name", id.Name)
	d.Set("resource_group_name", id.ResourceGroup)
	d.Set("cluster_name", id.ClusterName)
	d.Set("database_name", id.DatabaseName)

	if dataConnection, ok := connectionModel.Value.(kusto.IotHubDataConnection); ok {
		d.Set("location", location.NormalizeNilable(dataConnection.Location))
		if props := dataConnection.IotHubConnectionProperties; props != nil {
			d.Set("iothub_id", props.IotHubResourceID)
			d.Set("consumer_group", props.ConsumerGroup)
			d.Set("table_name", props.TableName)
			d.Set("mapping_rule_name", props.MappingRuleName)
			d.Set("data_format", props.DataFormat)
			d.Set("shared_access_policy_name", props.SharedAccessPolicyName)
			d.Set("event_system_properties", utils.FlattenStringSlice(props.EventSystemProperties))
		}
	}

	return nil
}

func resourceKustoIotHubDataConnectionDelete(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Kusto.DataConnectionsClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.DataConnectionID(d.Id())
	if err != nil {
		return err
	}

	future, err := client.Delete(ctx, id.ResourceGroup, id.ClusterName, id.DatabaseName, id.Name)
	if err != nil {
		return fmt.Errorf("deleting %s: %+v", id, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for deletion of %s: %+v", id, err)
	}

	return nil
}
