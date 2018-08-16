package azurerm

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/response"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmPostgreSQLVirtualNetworkRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmPostgreSQLVirtualNetworkRuleCreateUpdate,
		Read:   resourceArmPostgreSQLVirtualNetworkRuleRead,
		Update: resourceArmPostgreSQLVirtualNetworkRuleCreateUpdate,
		Delete: resourceArmPostgreSQLVirtualNetworkRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validatePostgreSQLVirtualNetworkRuleName,
			},

			"resource_group_name": resourceGroupNameSchema(),

			"server_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NonZeroValues,
			},

			"subnet_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NonZeroValues,
			},

			"ignore_missing_vnet_service_endpoint": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false, //When not provided, Azure defaults to false
			},
		},
	}
}

func resourceArmPostgreSQLVirtualNetworkRuleCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).postgresqlVirtualNetworkRulesClient
	ctx := meta.(*ArmClient).StopContext

	name := d.Get("name").(string)
	serverName := d.Get("server_name").(string)
	resourceGroup := d.Get("resource_group_name").(string)
	virtualNetworkSubnetId := d.Get("subnet_id").(string)
	ignoreMissingVnetServiceEndpoint := d.Get("ignore_missing_vnet_service_endpoint").(bool)

	parameters := postgresql.VirtualNetworkRule{
		VirtualNetworkRuleProperties: &postgresql.VirtualNetworkRuleProperties{
			VirtualNetworkSubnetID:           utils.String(virtualNetworkSubnetId),
			IgnoreMissingVnetServiceEndpoint: utils.Bool(ignoreMissingVnetServiceEndpoint),
		},
	}

	_, err := client.CreateOrUpdate(ctx, resourceGroup, serverName, name, parameters)
	if err != nil {
		return fmt.Errorf("Error creating PostgreSQL Virtual Network Rule %q (PostgreSQL Server: %q, Resource Group: %q): %+v", name, serverName, resourceGroup, err)
	}

	//Wait for the provisioning state to become ready
	log.Printf("[DEBUG] Waiting for PostgreSQL Virtual Network Rule %q (PostgreSQL Server: %q, Resource Group: %q) to become ready: %+v", name, serverName, resourceGroup, err)
	stateConf := &resource.StateChangeConf{
		Pending:                   []string{"Initializing", "InProgress", "Unknown", "ResponseNotFound"},
		Target:                    []string{"Ready"},
		Refresh:                   postgreSQLVirtualNetworkStateStatusCodeRefreshFunc(ctx, client, resourceGroup, serverName, name),
		Timeout:                   10 * time.Minute,
		MinTimeout:                1 * time.Minute,
		ContinuousTargetOccurence: 5,
	}

	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("Error waiting for PostgreSQL Virtual Network Rule %q (PostgreSQL Server: %q, Resource Group: %q) to be created or updated: %+v", name, serverName, resourceGroup, err)
	}

	resp, err := client.Get(ctx, resourceGroup, serverName, name)
	if err != nil {
		return fmt.Errorf("Error retrieving PostgreSQL Virtual Network Rule %q (PostgreSQL Server: %q, Resource Group: %q): %+v", name, serverName, resourceGroup, err)
	}

	d.SetId(*resp.ID)

	return resourceArmPostgreSQLVirtualNetworkRuleRead(d, meta)
}

func resourceArmPostgreSQLVirtualNetworkRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).postgresqlVirtualNetworkRulesClient
	ctx := meta.(*ArmClient).StopContext

	id, err := parseAzureResourceID(d.Id())
	if err != nil {
		return err
	}

	resourceGroup := id.ResourceGroup
	serverName := id.Path["servers"]
	name := id.Path["virtualNetworkRules"]

	resp, err := client.Get(ctx, resourceGroup, serverName, name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[INFO] Error reading PostgreSQL Virtual Network Rule %q - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error reading PostgreSQL Virtual Network Rule: %q (PostgreSQL Server: %q, Resource Group: %q): %+v", name, serverName, resourceGroup, err)
	}

	d.Set("name", resp.Name)
	d.Set("resource_group_name", resourceGroup)
	d.Set("server_name", serverName)

	if props := resp.VirtualNetworkRuleProperties; props != nil {
		d.Set("subnet_id", props.VirtualNetworkSubnetID)
		d.Set("ignore_missing_vnet_service_endpoint", props.IgnoreMissingVnetServiceEndpoint)
	}

	return nil
}

func resourceArmPostgreSQLVirtualNetworkRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).postgresqlVirtualNetworkRulesClient
	ctx := meta.(*ArmClient).StopContext

	id, err := parseAzureResourceID(d.Id())
	if err != nil {
		return err
	}

	resourceGroup := id.ResourceGroup
	serverName := id.Path["servers"]
	name := id.Path["virtualNetworkRules"]

	future, err := client.Delete(ctx, resourceGroup, serverName, name)
	if err != nil {
		if response.WasNotFound(future.Response()) {
			return nil
		}

		return fmt.Errorf("Error deleting PostgreSQL Virtual Network Rule %q (PostgreSQL Server: %q, Resource Group: %q): %+v", name, serverName, resourceGroup, err)
	}

	err = future.WaitForCompletionRef(ctx, client.Client)
	if err != nil {
		if response.WasNotFound(future.Response()) {
			return nil
		}

		return fmt.Errorf("Error deleting PostgreSQL Virtual Network Rule %q (PostgreSQL Server: %q, Resource Group: %q): %+v", name, serverName, resourceGroup, err)
	}

	return nil
}

/*
	This function checks the format of the PostgreSQL Virtual Network Rule Name to make sure that
	it does not contain any potentially invalid values.
*/
func validatePostgreSQLVirtualNetworkRuleName(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	// Cannot be empty
	if len(value) == 0 {
		errors = append(errors, fmt.Errorf(
			"%q cannot be an empty string: %q", k, value))
	}

	// Cannot be more than 128 characters
	if len(value) > 128 {
		errors = append(errors, fmt.Errorf(
			"%q cannot be longer than 128 characters: %q", k, value))
	}

	// Must only contain alphanumeric characters or hyphens
	if !regexp.MustCompile(`^[A-Za-z0-9-]*$`).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q can only contain alphanumeric characters and hyphens: %q",
			k, value))
	}

	// Cannot end in a hyphen
	if regexp.MustCompile(`-$`).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q cannot end with a hyphen: %q", k, value))
	}

	// Cannot start with a number or hyphen
	if regexp.MustCompile(`^[0-9-]`).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q cannot start with a number or hyphen: %q", k, value))
	}

	// There are multiple returns in the case that there is more than one invalid
	// case applied to the name.
	return
}

/*
	This function refreshes and checks the state of the PostgreSQL Virtual Network Rule.

	Response will contain a VirtualNetworkRuleProperties struct with a State property. The state property contain one of the following states (except ResponseNotFound).
	* Deleting
	* Initializing
	* InProgress
	* Unknown
	* Ready
	* ResponseNotFound (Custom state in case of 404)
*/
func postgreSQLVirtualNetworkStateStatusCodeRefreshFunc(ctx context.Context, client postgresql.VirtualNetworkRulesClient, resourceGroup string, serverName string, name string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, err := client.Get(ctx, resourceGroup, serverName, name)

		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				log.Printf("[DEBUG] Retrieving PostgreSQL Virtual Network Rule %q (PostgreSQL Server: %q, Resource Group: %q) returned 404.", resourceGroup, serverName, name)
				return nil, "ResponseNotFound", nil
			}

			return nil, "", fmt.Errorf("Error polling for the state of the PostgreSQL Virtual Network Rule %q (PostgreSQL Server: %q, Resource Group: %q): %+v", name, serverName, resourceGroup, err)
		}

		if props := resp.VirtualNetworkRuleProperties; props != nil {
			log.Printf("[DEBUG] Retrieving PostgreSQL Virtual Network Rule %q (PostgreSQL Server: %q, Resource Group: %q) returned Status %s", resourceGroup, serverName, name, props.State)
			return resp, fmt.Sprintf("%s", props.State), nil
		}

		//Valid response was returned but VirtualNetworkRuleProperties was nil. Basically the rule exists, but with no properties for some reason. Assume Unknown instead of returning error.
		log.Printf("[DEBUG] Retrieving PostgreSQL Virtual Network Rule %q (PostgreSQL Server: %q, Resource Group: %q) returned empty VirtualNetworkRuleProperties", resourceGroup, serverName, name)
		return resp, "Unknown", nil
	}
}
