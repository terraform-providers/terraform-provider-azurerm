package costmanagement

import (
	"fmt"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/costmanagement/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"

	"github.com/Azure/azure-sdk-for-go/services/costmanagement/mgmt/2020-06-01/costmanagement"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	billValidate "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/billing/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/costmanagement/parse"
	mgmGrpValidate "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/managementgroup/validate"
	rgValidate "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/validate"
	subValidate "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/subscription/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceCostManagementExport() *schema.Resource {
	return &schema.Resource{
		Create: resourceCostManagementExportCreateUpdate,
		Read:   resourceCostManagementExportRead,
		Update: resourceCostManagementExportCreateUpdate,
		Delete: resourceCostManagementExportDelete,
		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.CostManagementExportID(id)
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
				ValidateFunc: validate.ExportName,
			},

			"scope": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.Any(
					rgValidate.ResourceGroupID,
					subValidate.SubscriptionID,
					mgmGrpValidate.ManagementGroupID,
					billValidate.EnrollmentID,
					billValidate.EnrollmentBillingScopeID,
					billValidate.MicrosoftCustomerAccountBillingScopeID,
					billValidate.CustomerID,
					billValidate.BillingProfileID,
					billValidate.DepartmentID,
					billValidate.BillingAccountID,
				),
			},

			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"recurrence_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(costmanagement.RecurrenceTypeDaily),
					string(costmanagement.RecurrenceTypeWeekly),
					string(costmanagement.RecurrenceTypeMonthly),
					string(costmanagement.RecurrenceTypeAnnually),
				}, false),
			},

			"recurrence_period_start": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsRFC3339Time,
			},

			"recurrence_period_end": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsRFC3339Time,
			},

			"delivery_info": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"storage_account_id": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: azure.ValidateResourceID,
						},
						"container_name": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validate.ExportContainerName,
						},
						"root_folder_path": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
					},
				},
			},

			"query": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},

						"time_frame": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								string(costmanagement.Custom),
								string(costmanagement.BillingMonthToDate),
								string(costmanagement.MonthToDate),
								string(costmanagement.TheLastMonth),
								string(costmanagement.WeekToDate),
								string(costmanagement.MonthToDate),
							}, false),
						},
					},
				},
			},
		},
	}
}

func resourceCostManagementExportCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).CostManagement.ExportClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	scope := d.Get("scope").(string)

	if d.IsNewResource() {
		existing, err := client.Get(ctx, scope, name, "")
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("checking for presence of existing Cost Management Export %q (Scope %q): %s", name, scope, err)
			}
		}

		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_cost_management_export", *existing.ID)
		}
	}

	from, _ := time.Parse(time.RFC3339, d.Get("recurrence_period_start").(string))
	to, _ := time.Parse(time.RFC3339, d.Get("recurrence_period_end").(string))

	status := costmanagement.Active
	if v := d.Get("active"); !v.(bool) {
		status = costmanagement.Inactive
	}

	properties := costmanagement.Export{
		ExportProperties: &costmanagement.ExportProperties{
			Schedule: &costmanagement.ExportSchedule{
				Recurrence: costmanagement.RecurrenceType(d.Get("recurrence_type").(string)),
				RecurrencePeriod: &costmanagement.ExportRecurrencePeriod{
					From: &date.Time{Time: from},
					To:   &date.Time{Time: to},
				},
				Status: status,
			},
			DeliveryInfo: expandExportDeliveryInfo(d.Get("delivery_info").([]interface{})),
			Format:       costmanagement.Csv,
			Definition:   expandExportDefinition(d.Get("query").([]interface{})),
		},
	}

	if _, err := client.CreateOrUpdate(ctx, scope, name, properties); err != nil {
		return fmt.Errorf("creating/updating Cost Management Export %q (Scope %q): %+v", name, scope, err)
	}

	resp, err := client.Get(ctx, scope, name, "")
	if err != nil {
		return fmt.Errorf("retrieving Cost Management Export %q (Scope %q): %+v", name, scope, err)
	}

	if resp.ID == nil {
		return fmt.Errorf("cannot read Cost Management Export %q (Scope %q) ID", name, scope)
	}

	id := *resp.ID
	// The ID is missing the prefix `/` which causes our uri parse to fail
	if !strings.HasPrefix(id, "/") {
		id = fmt.Sprintf("/%s", id)
	}

	d.SetId(id)

	return resourceCostManagementExportRead(d, meta)
}

func resourceCostManagementExportRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).CostManagement.ExportClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.CostManagementExportID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceId, id.Name, "")
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("retrieving Cost Management Export %q (Scope %q): %+v", id.Name, id.ResourceId, err)
	}

	d.Set("name", resp.Name)
	d.Set("scope", id.ResourceId)

	if schedule := resp.Schedule; schedule != nil {
		if recurrencePeriod := schedule.RecurrencePeriod; recurrencePeriod != nil {
			d.Set("recurrence_period_start", recurrencePeriod.From.Format(time.RFC3339))
			d.Set("recurrence_period_end", recurrencePeriod.To.Format(time.RFC3339))
		}
		status := false
		if schedule.Status == costmanagement.Active {
			status = true
		}
		d.Set("active", status)
		d.Set("recurrence_type", schedule.Recurrence)
	}
	if err := d.Set("delivery_info", flattenExportDeliveryInfo(resp.DeliveryInfo)); err != nil {
		return fmt.Errorf("setting `delivery_info`: %+v", err)
	}

	if err := d.Set("query", flattenExportDefinition(resp.Definition)); err != nil {
		return fmt.Errorf("setting `query`: %+v", err)
	}

	return nil
}

func resourceCostManagementExportDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).CostManagement.ExportClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.CostManagementExportID(d.Id())
	if err != nil {
		return err
	}

	response, err := client.Delete(ctx, id.ResourceId, id.Name)
	if err != nil {
		if !utils.ResponseWasNotFound(response) {
			return fmt.Errorf("deleting Cost Management Export %q (Scope %q): %+v", id.Name, id.ResourceId, err)
		}
	}

	return nil
}
