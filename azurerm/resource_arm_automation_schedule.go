package azurerm

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/arm/automation"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmAutomationSchedule() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmAutomationScheduleCreateUpdate,
		Read:   resourceArmAutomationScheduleRead,
		Update: resourceArmAutomationScheduleCreateUpdate,
		Delete: resourceArmAutomationScheduleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"account_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"resource_group_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"start_time": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateStartTime,
			},

			"expiry_time": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"frequency": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: ignoreCaseDiffSuppressFunc,
				ValidateFunc: validation.StringInSlice([]string{
					string(automation.Day),
					string(automation.Hour),
					string(automation.Month),
					string(automation.OneTime),
					string(automation.Week),
				}, true),
			},
			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func validateStartTime(v interface{}, k string) (ws []string, errors []error) {
	starttime, tperr := time.Parse(time.RFC3339, v.(string))
	if tperr != nil {
		errors = append(errors, fmt.Errorf("Cannot parse %q", k))
	}

	u := time.Until(starttime)
	if u < 5*time.Minute {
		errors = append(errors, fmt.Errorf("%q should be at least 5 minutes in the future", k))
	}

	return
}

func resourceArmAutomationScheduleCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).automationScheduleClient
	log.Printf("[INFO] preparing arguments for AzureRM Automation Schedule creation.")

	name := d.Get("name").(string)
	resGroup := d.Get("resource_group_name").(string)

	accName := d.Get("account_name").(string)
	freqstr := d.Get("frequency").(string)
	freq := automation.ScheduleFrequency(freqstr)

	cst := d.Get("start_time").(string)
	starttime, tperr := time.Parse(time.RFC3339, cst)
	if tperr != nil {
		return fmt.Errorf("Cannot parse start_time: %s", cst)
	}

	ardt := date.Time{Time: starttime}

	description := d.Get("description").(string)
	var timezone string
	if v, ok := d.GetOk("timezone"); ok {
		timezone = v.(string)
	} else {
		timezone = "UTC"
	}
	parameters := automation.ScheduleCreateOrUpdateParameters{
		Name: &name,
		ScheduleCreateOrUpdateProperties: &automation.ScheduleCreateOrUpdateProperties{
			Description: &description,
			Frequency:   freq,
			StartTime:   &ardt,
			TimeZone:    &timezone,
		},
	}

	_, err := client.CreateOrUpdate(resGroup, accName, name, parameters)
	if err != nil {
		return err
	}

	read, err := client.Get(resGroup, accName, name)
	if err != nil {
		return err
	}

	if read.ID == nil {
		return fmt.Errorf("Cannot read Automation Schedule '%s' (resource group %s) ID", name, resGroup)
	}

	d.SetId(*read.ID)

	return resourceArmAutomationScheduleRead(d, meta)
}

func resourceArmAutomationScheduleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).automationScheduleClient
	id, err := parseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resGroup := id.ResourceGroup
	accName := id.Path["automationAccounts"]
	name := id.Path["schedules"]

	resp, err := client.Get(resGroup, accName, name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error making Read request on AzureRM Automation Schedule '%s': %s", name, err)
	}

	d.Set("name", resp.Name)
	d.Set("resource_group_name", resGroup)
	d.Set("account_name", accName)
	d.Set("frequency", resp.Frequency)
	d.Set("description", resp.Description)
	d.Set("start_time", string(resp.StartTime.Format(time.RFC3339)))
	return nil
}

func resourceArmAutomationScheduleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).automationScheduleClient

	id, err := parseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resGroup := id.ResourceGroup
	accName := id.Path["automationAccounts"]
	name := id.Path["schedules"]

	resp, err := client.Delete(resGroup, accName, name)

	if err != nil {
		if utils.ResponseWasNotFound(resp) {
			return nil
		}

		return fmt.Errorf("Error issuing AzureRM delete request for Automation Schedule '%s': %+v", name, err)
	}

	return nil
}
