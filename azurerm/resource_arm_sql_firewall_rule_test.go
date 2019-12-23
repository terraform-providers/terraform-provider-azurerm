package azurerm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func TestAccAzureRMSqlFirewallRule_basic(t *testing.T) {
	resourceName := "azurerm_sql_firewall_rule.test"
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMSqlFirewallRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMSqlFirewallRule_basic(ri, acceptance.Location()),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSqlFirewallRuleExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "start_ip_address", "0.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "end_ip_address", "255.255.255.255"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAzureRMSqlFirewallRule_withUpdates(ri, acceptance.Location()),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSqlFirewallRuleExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "start_ip_address", "10.0.17.62"),
					resource.TestCheckResourceAttr(resourceName, "end_ip_address", "10.0.17.62"),
				),
			},
		},
	})
}
func TestAccAzureRMSqlFirewallRule_requiresImport(t *testing.T) {
	if !features.ShouldResourcesBeImported() {
		t.Skip("Skipping since resources aren't required to be imported")
		return
	}
	resourceName := "azurerm_sql_firewall_rule.test"
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMSqlFirewallRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMSqlFirewallRule_basic(ri, acceptance.Location()),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSqlFirewallRuleExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "start_ip_address", "0.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "end_ip_address", "255.255.255.255"),
				),
			},
			{
				Config:      testAccAzureRMSqlFirewallRule_requiresImport(ri, acceptance.Location()),
				ExpectError: acceptance.RequiresImportError("azurerm_sql_firewall_rule"),
			},
		},
	})
}

func TestAccAzureRMSqlFirewallRule_disappears(t *testing.T) {
	resourceName := "azurerm_sql_firewall_rule.test"
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMSqlFirewallRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMSqlFirewallRule_basic(ri, acceptance.Location()),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSqlFirewallRuleExists(resourceName),
					testCheckAzureRMSqlFirewallRuleDisappears(resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testCheckAzureRMSqlFirewallRuleExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		resourceGroup := rs.Primary.Attributes["resource_group_name"]
		serverName := rs.Primary.Attributes["server_name"]
		ruleName := rs.Primary.Attributes["name"]

		client := acceptance.AzureProvider.Meta().(*clients.Client).Sql.FirewallRulesClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		resp, err := client.Get(ctx, resourceGroup, serverName, ruleName)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("SQL Firewall Rule %q (server %q / resource group %q) was not found", ruleName, serverName, resourceGroup)
			}

			return err
		}

		return nil
	}
}

func testCheckAzureRMSqlFirewallRuleDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azurerm_sql_firewall_rule" {
			continue
		}

		resourceGroup := rs.Primary.Attributes["resource_group_name"]
		serverName := rs.Primary.Attributes["server_name"]
		ruleName := rs.Primary.Attributes["name"]

		client := acceptance.AzureProvider.Meta().(*clients.Client).Sql.FirewallRulesClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		resp, err := client.Get(ctx, resourceGroup, serverName, ruleName)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return err
		}

		return fmt.Errorf("SQL Firewall Rule %q (server %q / resource group %q) still exists: %+v", ruleName, serverName, resourceGroup, resp)
	}

	return nil
}

func testCheckAzureRMSqlFirewallRuleDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		resourceGroup := rs.Primary.Attributes["resource_group_name"]
		serverName := rs.Primary.Attributes["server_name"]
		ruleName := rs.Primary.Attributes["name"]

		client := acceptance.AzureProvider.Meta().(*clients.Client).Sql.FirewallRulesClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		resp, err := client.Delete(ctx, resourceGroup, serverName, ruleName)
		if err != nil {
			if utils.ResponseWasNotFound(resp) {
				return nil
			}

			return fmt.Errorf("Bad: Delete on sqlFirewallRulesClient: %+v", err)
		}

		return nil
	}
}

func testAccAzureRMSqlFirewallRule_basic(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_sql_server" "test" {
  name                         = "acctestsqlserver%d"
  resource_group_name          = "${azurerm_resource_group.test.name}"
  location                     = "${azurerm_resource_group.test.location}"
  version                      = "12.0"
  administrator_login          = "mradministrator"
  administrator_login_password = "thisIsDog11"
}

resource "azurerm_sql_firewall_rule" "test" {
  name                = "acctestsqlserver%d"
  resource_group_name = "${azurerm_resource_group.test.name}"
  server_name         = "${azurerm_sql_server.test.name}"
  start_ip_address    = "0.0.0.0"
  end_ip_address      = "255.255.255.255"
}
`, rInt, location, rInt, rInt)
}

func testAccAzureRMSqlFirewallRule_requiresImport(rInt int, location string) string {
	return fmt.Sprintf(`
%s

resource "azurerm_sql_firewall_rule" "import" {
  name                = "${azurerm_sql_firewall_rule.test.name}"
  resource_group_name = "${azurerm_sql_firewall_rule.test.resource_group_name}"
  server_name         = "${azurerm_sql_firewall_rule.test.server_name}"
  start_ip_address    = "${azurerm_sql_firewall_rule.test.start_ip_address}"
  end_ip_address      = "${azurerm_sql_firewall_rule.test.end_ip_address}"
}
`, testAccAzureRMSqlFirewallRule_basic(rInt, location))
}

func testAccAzureRMSqlFirewallRule_withUpdates(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_sql_server" "test" {
  name                         = "acctestsqlserver%d"
  resource_group_name          = "${azurerm_resource_group.test.name}"
  location                     = "${azurerm_resource_group.test.location}"
  version                      = "12.0"
  administrator_login          = "mradministrator"
  administrator_login_password = "thisIsDog11"
}

resource "azurerm_sql_firewall_rule" "test" {
  name                = "acctestsqlserver%d"
  resource_group_name = "${azurerm_resource_group.test.name}"
  server_name         = "${azurerm_sql_server.test.name}"
  start_ip_address    = "10.0.17.62"
  end_ip_address      = "10.0.17.62"
}
`, rInt, location, rInt, rInt)
}
