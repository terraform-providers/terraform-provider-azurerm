package policy

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/guestconfiguration/mgmt/2020-06-25/guestconfiguration"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/location"
	computeParse "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/compute/parse"
	computeValidate "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/compute/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/policy/parse"
	azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceVirtualMachineConfigurationPolicyAssignment() *schema.Resource {
	return &schema.Resource{
		Create: resourceVirtualMachineConfigurationPolicyAssignmentCreateUpdate,
		Read:   resourceVirtualMachineConfigurationPolicyAssignmentRead,
		Update: resourceVirtualMachineConfigurationPolicyAssignmentCreateUpdate,
		Delete: resourceVirtualMachineConfigurationPolicyAssignmentDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Importer: azSchema.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.VirtualMachineConfigurationPolicyAssignmentID(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"location": azure.SchemaLocation(),

			"virtual_machine_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: computeValidate.VirtualMachineID,
			},

			"policy": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},

						"parameter": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringIsNotEmpty,
									},

									"value": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"version": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceVirtualMachineConfigurationPolicyAssignmentCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	client := meta.(*clients.Client).Policy.GuestConfigurationAssignmentsClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	vmId, err := computeParse.VirtualMachineID(d.Get("virtual_machine_id").(string))
	if err != nil {
		return err
	}

	id := parse.NewVirtualMachineConfigurationPolicyAssignmentID(subscriptionId, vmId.ResourceGroup, vmId.Name, d.Get("name").(string))

	if d.IsNewResource() {
		existing, err := client.Get(ctx, id.ResourceGroup, id.GuestConfigurationAssignmentName, id.VirtualMachineName)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("checking for present of existing GuestConfiguration GuestConfigurationAssignment %q (Virtual Machine ID %q): %+v", id.GuestConfigurationAssignmentName, vmId.ID(), err)
			}
		}
		if !utils.ResponseWasNotFound(existing.Response) {
			return tf.ImportAsExistsError("azurerm_virtual_machine_configuration_policy_assignment", id.ID())
		}
	}

	parameter := guestconfiguration.Assignment{
		Name:     utils.String(d.Get("name").(string)),
		Location: utils.String(location.Normalize(d.Get("location").(string))),
		Properties: &guestconfiguration.AssignmentProperties{
			GuestConfiguration: expandGuestConfigurationAssignment(d.Get("policy").([]interface{})),
		},
	}
	future, err := client.CreateOrUpdate(ctx, id.GuestConfigurationAssignmentName, parameter, id.ResourceGroup, id.VirtualMachineName)
	if err != nil {
		return fmt.Errorf("creating/updating GuestConfigurationAssignment %q (Virtual Machine ID %q): %+v", id.GuestConfigurationAssignmentName, vmId.ID(), err)
	}

	if err := future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting on creating/updating future for GuestConfigurationAssignment %q (Virtual Machine ID %q): %+v", id.GuestConfigurationAssignmentName, vmId.ID(), err)
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.GuestConfigurationAssignmentName, id.VirtualMachineName)
	if err != nil {
		return fmt.Errorf("retrieving GuestConfigurationAssignment %q (Virtual Machine ID %q): %+v", id.GuestConfigurationAssignmentName, vmId.ID(), err)
	}

	if resp.ID == nil || *resp.ID == "" {
		return fmt.Errorf("empty or nil ID returned for GuestConfigurationAssignment %q (Virtual Machine ID %q) ID", id.GuestConfigurationAssignmentName, vmId.ID())
	}

	d.SetId(id.ID())

	return resourceVirtualMachineConfigurationPolicyAssignmentRead(d, meta)
}

func resourceVirtualMachineConfigurationPolicyAssignmentRead(d *schema.ResourceData, meta interface{}) error {
	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	client := meta.(*clients.Client).Policy.GuestConfigurationAssignmentsClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.VirtualMachineConfigurationPolicyAssignmentID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.GuestConfigurationAssignmentName, id.VirtualMachineName)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[INFO] guestConfiguration %q does not exist - removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("retrieving GuestConfigurationAssignment %q (Resource Group %q / Virtual Machine %q): %+v", id.GuestConfigurationAssignmentName, id.ResourceGroup, id.VirtualMachineName, err)
	}

	vmId := computeParse.NewVirtualMachineID(subscriptionId, id.ResourceGroup, id.VirtualMachineName)
	d.Set("name", id.GuestConfigurationAssignmentName)
	d.Set("virtual_machine_id", vmId.ID())
	d.Set("location", location.NormalizeNilable(resp.Location))

	if props := resp.Properties; props != nil {
		if err := d.Set("policy", flattenGuestConfigurationAssignment(props.GuestConfiguration)); err != nil {
			return fmt.Errorf("setting `policy`: %+v", err)
		}
	}
	return nil
}

func resourceVirtualMachineConfigurationPolicyAssignmentDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Policy.GuestConfigurationAssignmentsClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.VirtualMachineConfigurationPolicyAssignmentID(d.Id())
	if err != nil {
		return err
	}

	future, err := client.Delete(ctx, id.ResourceGroup, id.GuestConfigurationAssignmentName, id.VirtualMachineName)
	if err != nil {
		return fmt.Errorf("deleting GuestConfiguration GuestConfigurationAssignment %q (Resource Group %q / Virtual Machine %q): %+v", id.GuestConfigurationAssignmentName, id.ResourceGroup, id.VirtualMachineName, err)
	}

	if err := future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting on deleting future for GuestConfiguration GuestConfigurationAssignment %q (Resource Group %q / Virtual Machine %q): %+v", id.GuestConfigurationAssignmentName, id.ResourceGroup, id.VirtualMachineName, err)
	}
	return nil
}

func expandGuestConfigurationAssignment(input []interface{}) *guestconfiguration.Navigation {
	if len(input) == 0 {
		return nil
	}
	v := input[0].(map[string]interface{})
	return &guestconfiguration.Navigation{
		Name:                   utils.String(v["name"].(string)),
		Version:                utils.String(v["version"].(string)),
		ConfigurationParameter: expandGuestConfigurationAssignmentConfigurationParameters(v["parameter"].(*schema.Set).List()),
	}
}

func expandGuestConfigurationAssignmentConfigurationParameters(input []interface{}) *[]guestconfiguration.ConfigurationParameter {
	results := make([]guestconfiguration.ConfigurationParameter, 0)
	for _, item := range input {
		v := item.(map[string]interface{})
		results = append(results, guestconfiguration.ConfigurationParameter{
			Name:  utils.String(v["name"].(string)),
			Value: utils.String(v["value"].(string)),
		})
	}
	return &results
}

func flattenGuestConfigurationAssignment(input *guestconfiguration.Navigation) []interface{} {
	if input == nil {
		return make([]interface{}, 0)
	}

	var name string
	if input.Name != nil {
		name = *input.Name
	}
	var version string
	if input.Version != nil {
		version = *input.Version
	}
	return []interface{}{
		map[string]interface{}{
			"name":      name,
			"parameter": flattenGuestConfigurationAssignmentConfigurationParameters(input.ConfigurationParameter),
			"version":   version,
		},
	}
}

func flattenGuestConfigurationAssignmentConfigurationParameters(input *[]guestconfiguration.ConfigurationParameter) []interface{} {
	results := make([]interface{}, 0)
	if input == nil {
		return results
	}

	for _, item := range *input {
		var name string
		if item.Name != nil {
			name = *item.Name
		}
		var value string
		if item.Value != nil {
			value = *item.Value
		}
		results = append(results, map[string]interface{}{
			"name":  name,
			"value": value,
		})
	}
	return results
}