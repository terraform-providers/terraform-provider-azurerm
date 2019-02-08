package azurerm

import (
	"fmt"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
)

func TestAccAzureRMLoadBalancerOutboundRule_basic(t *testing.T) {
	var lb network.LoadBalancer
	ri := tf.AccRandTimeInt()
	outboundRuleName := fmt.Sprintf("OutboundRule-%s", ri)

	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
	outboundRuleId := fmt.Sprintf(
		"/subscriptions/%s/resourceGroups/acctestRG-%d/providers/Microsoft.Network/loadBalancers/arm-test-loadbalancer-%d/outboundRules/%s",
		subscriptionID, ri, ri, outboundRuleName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMLoadBalancerOutboundRule_basic(ri, outboundRuleName, testLocation()),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMLoadBalancerExists("azurerm_lb.test", &lb),
					testCheckAzureRMLoadBalancerOutboundRuleExists(outboundRuleName, &lb),
					resource.TestCheckResourceAttr(
						"azurerm_lb_outbound_rule.test", "id", outboundRuleId),
				),
			},
			{
				ResourceName:      "azurerm_lb.test",
				ImportState:       true,
				ImportStateVerify: true,
				// location is deprecated and was never actually used
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func TestAccAzureRMLoadBalancerOutboundRule_requiresImport(t *testing.T) {
	if !requireResourcesToBeImported {
		t.Skip("Skipping since resources aren't required to be imported")
		return
	}

	var lb network.LoadBalancer
	ri := tf.AccRandTimeInt()
	outboundRuleName := fmt.Sprintf("OutboundRule-%s", ri)
	location := testLocation()

	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
	outboundRuleId := fmt.Sprintf(
		"/subscriptions/%s/resourceGroups/acctestRG-%d/providers/Microsoft.Network/loadBalancers/arm-test-loadbalancer-%d/outboundRules/%s",
		subscriptionID, ri, ri, outboundRuleName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMLoadBalancerOutboundRule_basic(ri, outboundRuleName, location),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMLoadBalancerExists("azurerm_lb.test", &lb),
					testCheckAzureRMLoadBalancerOutboundRuleExists(outboundRuleName, &lb),
					resource.TestCheckResourceAttr(
						"azurerm_lb_outbound_rule.test", "id", outboundRuleId),
				),
			},
			{
				Config:      testAccAzureRMLoadBalancerOutboundRule_requiresImport(ri, outboundRuleName, location),
				ExpectError: testRequiresImportError("azurerm_lb_outbound_rule"),
			},
		},
	})
}

func TestAccAzureRMLoadBalancerOutboundRule_removal(t *testing.T) {
	var lb network.LoadBalancer
	ri := tf.AccRandTimeInt()
	outboundRuleName := fmt.Sprintf("OutboundRule-%s", ri)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMLoadBalancerOutboundRule_basic(ri, outboundRuleName, testLocation()),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMLoadBalancerExists("azurerm_lb.test", &lb),
					testCheckAzureRMLoadBalancerOutboundRuleExists(outboundRuleName, &lb),
				),
			},
			{
				Config: testAccAzureRMLoadBalancerOutboundRule_removal(ri, testLocation()),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMLoadBalancerExists("azurerm_lb.test", &lb),
					testCheckAzureRMLoadBalancerOutboundRuleNotExists(outboundRuleName, &lb),
				),
			},
		},
	})
}

func TestAccAzureRMLoadBalancerOutboundRule_update(t *testing.T) {
	var lb network.LoadBalancer
	ri := tf.AccRandTimeInt()
	outboundRuleName := fmt.Sprintf("OutboundRule-%s", ri)
	outboundRule2Name := fmt.Sprintf("OutboundRule-%s", tf.AccRandTimeInt())

	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
	outboundRuleID := fmt.Sprintf(
		"/subscriptions/%s/resourceGroups/acctestRG-%d/providers/Microsoft.Network/loadBalancers/arm-test-loadbalancer-%d/outboundRules/%s",
		subscriptionID, ri, ri, outboundRuleName)

	outboundRule2ID := fmt.Sprintf(
		"/subscriptions/%s/resourceGroups/acctestRG-%d/providers/Microsoft.Network/loadBalancers/arm-test-loadbalancer-%d/outboundRules/%s",
		subscriptionID, ri, ri, outboundRule2Name)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMLoadBalancerOutboundRule_multipleRules(ri, outboundRuleName, outboundRule2Name, testLocation()),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMLoadBalancerExists("azurerm_lb.test", &lb),
					testCheckAzureRMLoadBalancerOutboundRuleExists(outboundRuleName, &lb),
					testCheckAzureRMLoadBalancerOutboundRuleExists(outboundRule2Name, &lb),
					resource.TestCheckResourceAttr("azurerm_lb_outbound_rule.test2", "protocol", "Udp"),
				),
			},
			{
				Config: testAccAzureRMLoadBalancerOutboundRule_multipleRulesUpdate(ri, outboundRuleName, outboundRule2Name, testLocation()),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMLoadBalancerExists("azurerm_lb.test", &lb),
					testCheckAzureRMLoadBalancerOutboundRuleExists(outboundRuleName, &lb),
					testCheckAzureRMLoadBalancerOutboundRuleExists(outboundRule2Name, &lb),
					resource.TestCheckResourceAttr("azurerm_lb_outbound_rule.test2", "protocol", "All"),
				),
			},
		},
	})
}

func TestAccAzureRMLoadBalancerOutboundRule_reapply(t *testing.T) {
	var lb network.LoadBalancer
	ri := tf.AccRandTimeInt()
	outboundRuleName := fmt.Sprintf("OutboundRule-%s", ri)

	deleteOutboundRuleState := func(s *terraform.State) error {
		return s.Remove("azurerm_lb_outbound_rule.test")
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMLoadBalancerOutboundRule_basic(ri, outboundRuleName, testLocation()),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMLoadBalancerExists("azurerm_lb.test", &lb),
					testCheckAzureRMLoadBalancerOutboundRuleExists(outboundRuleName, &lb),
					deleteRuleState,
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccAzureRMLoadBalancerOutboundRule_basic(ri, outboundRuleName, testLocation()),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMLoadBalancerExists("azurerm_lb.test", &lb),
					testCheckAzureRMLoadBalancerOutboundRuleExists(outboundRuleName, &lb),
				),
			},
		},
	})
}

func TestAccAzureRMLoadBalancerOutboundRule_disappears(t *testing.T) {
	var lb network.LoadBalancer
	ri := tf.AccRandTimeInt()
	outboundRuleName := fmt.Sprintf("OutboundRule-%s", ri)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMLoadBalancerOutboundRule_basic(ri, outboundRuleName, testLocation()),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMLoadBalancerExists("azurerm_lb.test", &lb),
					testCheckAzureRMLoadBalancerOutboundRuleExists(outboundRuleName, &lb),
					testCheckAzureRMLoadBalancerOutboundRuleDisappears(outboundRuleName, &lb),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testCheckAzureRMLoadBalancerOutboundRuleExists(outboundRuleName string, lb *network.LoadBalancer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, _, exists := findLoadBalancerOutboundRuleByName(lb, outboundRuleName)
		if !exists {
			return fmt.Errorf("A Load Balancer Outbound Rule with name %q cannot be found.", outboundRuleName)
		}

		return nil
	}
}

func testCheckAzureRMLoadBalancerOutboundRuleNotExists(outboundRuleName string, lb *network.LoadBalancer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, _, exists := findLoadBalancerOutboundRuleByName(lb, outboundRuleName)
		if exists {
			return fmt.Errorf("A Load Balancer Outbound Rule with name %q has been found.", outboundRuleName)
		}

		return nil
	}
}

func testCheckAzureRMLoadBalancerOutboundRuleDisappears(ruleName string, lb *network.LoadBalancer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*ArmClient).loadBalancerClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext

		_, i, exists := findLoadBalancerOutboundRuleByName(lb, ruleName)
		if !exists {
			return fmt.Errorf("A Outbound Rule with name %q cannot be found.", ruleName)
		}

		currentRules := *lb.LoadBalancerPropertiesFormat.OutboundRules
		rules := append(currentRules[:i], currentRules[i+1:]...)
		lb.LoadBalancerPropertiesFormat.OutboundRules = &rules

		id, err := parseAzureResourceID(*lb.ID)
		if err != nil {
			return err
		}

		future, err := client.CreateOrUpdate(ctx, id.ResourceGroup, *lb.Name, *lb)
		if err != nil {
			return fmt.Errorf("Error Creating/Updating Load Balancer %q (Resource Group %q): %+v", *lb.Name, id.ResourceGroup, err)
		}

		if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
			return fmt.Errorf("Error waiting for completion of Load Balancer %q (Resource Group %q): %+v", *lb.Name, id.ResourceGroup, err)
		}

		_, err = client.Get(ctx, id.ResourceGroup, *lb.Name, "")
		return err
	}
}

func testAccAzureRMLoadBalancerOutboundRule_basic(rInt int, outboundRuleName string, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_public_ip" "test" {
  name                = "test-ip-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  allocation_method   = "Static"
}

resource "azurerm_lb" "test" {
  name                = "arm-test-loadbalancer-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  frontend_ip_configuration {
    name                 = "one-%d"
    public_ip_address_id = "${azurerm_public_ip.test.id}"
  }
}

resource "azurerm_lb_outbound_rule" "test" {
  location                       = "${azurerm_resource_group.test.location}"
  resource_group_name            = "${azurerm_resource_group.test.name}"
  loadbalancer_id                = "${azurerm_lb.test.id}"
  name                           = "%s"
  protocol                       = "All"

  frontend_ip_configuration {
    name = "one-%d"
  }
}
`, rInt, location, rInt, rInt, rInt, outboundRuleName, rInt)
}

func testAccAzureRMLoadBalancerOutboundRule_requiresImport(rInt int, name string, location string) string {
	template := testAccAzureRMLoadBalancerOutboundRule_basic(rInt, name, location)
	return fmt.Sprintf(`
%s

resource "azurerm_lb_outbound_rule" "import" {
  name                           = "${azurerm_lb_outbound_rule.test.name}"
  location                       = "${azurerm_lb_outbound_rule.test.location}"
  resource_group_name            = "${azurerm_lb_outbound_rule.test.resource_group_name}"
  loadbalancer_id                = "${azurerm_lb_outbound_rule.test.loadbalancer_id}"
  protocol                       = "All"

  frontend_ip_configuration {
    name = "${azurerm_lb_outbound_rule.test.frontend_ip_configuration_name}"
  }
}
`, template)
}

func testAccAzureRMLoadBalancerOutboundRule_removal(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_public_ip" "test" {
  name                = "test-ip-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  allocation_method   = "Static"
}

resource "azurerm_lb" "test" {
  name                = "arm-test-loadbalancer-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  frontend_ip_configuration {
    name                 = "one-%d"
    public_ip_address_id = "${azurerm_public_ip.test.id}"
  }
}
`, rInt, location, rInt, rInt, rInt)
}

func testAccAzureRMLoadBalancerOutboundRule_multipleRules(rInt int, outboundRuleName, outboundRule2Name string, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_public_ip" "test" {
  name                = "test-ip-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  allocation_method   = "Static"
}

resource "azurerm_lb" "test" {
  name                = "arm-test-loadbalancer-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  frontend_ip_configuration {
    name                 = "one-%d"
    public_ip_address_id = "${azurerm_public_ip.test.id}"
  }
}

resource "azurerm_lb_outbound_rule" "test" {
  location                       = "${azurerm_resource_group.test.location}"
  resource_group_name            = "${azurerm_resource_group.test.name}"
  loadbalancer_id                = "${azurerm_lb.test.id}"
  name                           = "%s"
  protocol                       = "Tcp"
  frontend_ip_configuration {
    name = "one-%d"
  }
}

resource "azurerm_lb_outbound_rule" "test2" {
  location                       = "${azurerm_resource_group.test.location}"
  resource_group_name            = "${azurerm_resource_group.test.name}"
  loadbalancer_id                = "${azurerm_lb.test.id}"
  name                           = "%s"
  protocol                       = "Udp"
  frontend_ip_configuration {
    name = "one-%d"
  }
}
`, rInt, location, rInt, rInt, rInt, outboundRuleName, rInt, outboundRule2Name, rInt)
}

func testAccAzureRMLoadBalancerOutboundRule_multipleRulesUpdate(rInt int, outboundRuleName, outboundRule2Name string, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_public_ip" "test" {
  name                = "test-ip-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  allocation_method   = "Static"
}

resource "azurerm_lb" "test" {
  name                = "arm-test-loadbalancer-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  frontend_ip_configuration {
    name                 = "one-%d"
    public_ip_address_id = "${azurerm_public_ip.test.id}"
  }
}

resource "azurerm_lb_outbound_rule" "test" {
  location                       = "${azurerm_resource_group.test.location}"
  resource_group_name            = "${azurerm_resource_group.test.name}"
  loadbalancer_id                = "${azurerm_lb.test.id}"
  name                           = "%s"
  protocol                       = "All"
  frontend_ip_configuration {
    name = "one-%d"
  }
}

resource "azurerm_lb_outbound_rule" "test2" {
  location                       = "${azurerm_resource_group.test.location}"
  resource_group_name            = "${azurerm_resource_group.test.name}"
  loadbalancer_id                = "${azurerm_lb.test.id}"
  name                           = "%s"
  protocol                       = "All"
  frontend_ip_configuration {
    name = "one-%d"
  }
}
`, rInt, location, rInt, rInt, rInt, outboundRuleName, rInt, outboundRule2Name, rInt)
}
