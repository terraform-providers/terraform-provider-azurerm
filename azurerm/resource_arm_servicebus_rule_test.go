package azurerm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func TestAccAzureRMServiceBusRule_basicSqlFilter(t *testing.T) {
	ri := acctest.RandInt()
	config := testAccAzureRMServiceBusRule_basicSqlFilter(ri, testLocation())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMServiceBusTopicDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMServiceBusRuleExists(),
				),
			},
		},
	})
}

func TestAccAzureRMServiceBusRule_basicCorrelationFilter(t *testing.T) {
	ri := acctest.RandInt()
	config := testAccAzureRMServiceBusRule_basicCorrelationFilter(ri, testLocation())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMServiceBusTopicDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMServiceBusRuleExists(),
				),
			},
		},
	})
}

func TestAccAzureRMServiceBusRule_sqlFilterWithAction(t *testing.T) {
	ri := acctest.RandInt()
	config := testAccAzureRMServiceBusRule_sqlFilterWithAction(ri, testLocation())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMServiceBusTopicDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMServiceBusRuleExists(),
				),
			},
		},
	})
}

func TestAccAzureRMServiceBusRule_correlationFilterWithAction(t *testing.T) {
	ri := acctest.RandInt()
	config := testAccAzureRMServiceBusRule_correlationFilterWithAction(ri, testLocation())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMServiceBusTopicDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMServiceBusRuleExists(),
				),
			},
		},
	})
}

func TestAccAzureRMServiceBusRule_sqlFilterUpdated(t *testing.T) {
	ri := acctest.RandInt()
	location := testLocation()
	config := testAccAzureRMServiceBusRule_basicSqlFilter(ri, location)
	updatedConfig := testAccAzureRMServiceBusRule_basicSqlFilterUpdated(ri, location)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMServiceBusTopicDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMServiceBusRuleExists(),
					resource.TestCheckResourceAttr(resourceName, "sql_filter", "2=2"),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMServiceBusRuleExists(),
					resource.TestCheckResourceAttr(resourceName, "sql_filter", "3=3"),
				),
			},
		},
	})
}

func TestAccAzureRMServiceBusRule_correlationFilterUpdated(t *testing.T) {
	ri := acctest.RandInt()
	location := testLocation()
	config := testAccAzureRMServiceBusRule_correlationFilter(ri, location)
	updatedConfig := testAccAzureRMServiceBusRule_correlationFilterUpdated(ri, location)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMServiceBusTopicDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMServiceBusRuleExists(),
					resource.TestCheckResourceAttr(resourceName, "correlation_filter.0.message_id", "test_message_id"),
					resource.TestCheckResourceAttr(resourceName, "correlation_filter.0.reply_to", ""),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMServiceBusRuleExists(),
					resource.TestCheckResourceAttr(resourceName, "correlation_filter.0.message_id", "test_message_id_updated"),
					resource.TestCheckResourceAttr(resourceName, "correlation_filter.0.reply_to", "test_reply_to_added"),
				),
			},
		},
	})
}

func TestAccAzureRMServiceBusRule_updateSqlFilterToCorrelationFilter(t *testing.T) {
	ri := acctest.RandInt()
	location := testLocation()
	config := testAccAzureRMServiceBusRule_basicSqlFilter(ri, location)
	updatedConfig := testAccAzureRMServiceBusRule_basicCorrelationFilter(ri, location)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMServiceBusTopicDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMServiceBusRuleExists(),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMServiceBusRuleExists(),
				),
			},
		},
	})
}

func testCheckAzureRMServiceBusRuleExists() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		ruleName := rs.Primary.Attributes["name"]
		subscriptionName := rs.Primary.Attributes["subscription_name"]
		topicName := rs.Primary.Attributes["topic_name"]
		namespaceName := rs.Primary.Attributes["namespace_name"]
		resourceGroup, hasResourceGroup := rs.Primary.Attributes["resource_group_name"]
		if !hasResourceGroup {
			return fmt.Errorf("Bad: no resource group found in state for Rule: %q", ruleName)
		}

		client := testAccProvider.Meta().(*ArmClient).serviceBusRulesClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext

		resp, err := client.Get(ctx, resourceGroup, namespaceName, topicName, subscriptionName, ruleName)
		if err != nil {
			return fmt.Errorf("Bad: Get on serviceBusRulesClient: %+v", err)
		}

		if utils.ResponseWasNotFound(resp.Response) {
			return fmt.Errorf("Bad: Rule %q (resource group: %q) does not exist", ruleName, resourceGroup)
		}

		return nil
	}
}

func testAccAzureRMServiceBusRule_basicSqlFilter(rInt int, location string) string {
	sqlFilter := `
	filter_type = "SqlFilter"
	sql_filter  = "2=2"
	`
	return fmt.Sprintf(testAccAzureRMServiceBusRule_tfTemplate, rInt, location, rInt, rInt, rInt, rInt, sqlFilter)
}

func testAccAzureRMServiceBusRule_basicSqlFilterUpdated(rInt int, location string) string {
	sqlFilter := `
	filter_type = "SqlFilter"
	sql_filter  = "3=3"
	`
	return fmt.Sprintf(testAccAzureRMServiceBusRule_tfTemplate, rInt, location, rInt, rInt, rInt, rInt, sqlFilter)
}

func testAccAzureRMServiceBusRule_sqlFilterWithAction(rInt int, location string) string {
	sqlFilter := `
	filter_type = "SqlFilter"
	sql_filter  = "2=2"
	action      = "SET Test='true'"
	`
	return fmt.Sprintf(testAccAzureRMServiceBusRule_tfTemplate, rInt, location, rInt, rInt, rInt, rInt, sqlFilter)
}

func testAccAzureRMServiceBusRule_basicCorrelationFilter(rInt int, location string) string {
	correlationFilter := `
	filter_type = "CorrelationFilter"
	correlation_filter = {
		correlation_id      = "test_correlation_id"
		message_id          = "test_message_id"
		to                  = "test_to"
		reply_to            = "test_reply_to"
		label               = "test_label"
		session_id          = "test_session_id"
		reply_to_session_id = "test_reply_to_session_id"
		content_type        = "test_content_type"
	  }
	`
	return fmt.Sprintf(testAccAzureRMServiceBusRule_tfTemplate, rInt, location, rInt, rInt, rInt, rInt, correlationFilter)
}

func testAccAzureRMServiceBusRule_correlationFilter(rInt int, location string) string {
	correlationFilter := `
	filter_type = "CorrelationFilter"
	correlation_filter = {
		correlation_id      = "test_correlation_id"
		message_id          = "test_message_id"
	  }
	`
	return fmt.Sprintf(testAccAzureRMServiceBusRule_tfTemplate, rInt, location, rInt, rInt, rInt, rInt, correlationFilter)
}

func testAccAzureRMServiceBusRule_correlationFilterUpdated(rInt int, location string) string {
	correlationFilter := `
	filter_type = "CorrelationFilter"
	correlation_filter = {
		correlation_id      = "test_correlation_id"
		message_id          = "test_message_id_updated"
		reply_to            = "test_reply_to_added"
	  }
	`
	return fmt.Sprintf(testAccAzureRMServiceBusRule_tfTemplate, rInt, location, rInt, rInt, rInt, rInt, correlationFilter)
}

func testAccAzureRMServiceBusRule_correlationFilterWithAction(rInt int, location string) string {
	correlationFilter := `
	filter_type = "CorrelationFilter"
	correlation_filter = {
		correlation_id      = "test_correlation_id"
		message_id          = "test_message_id"
	  }
	  action      = "SET Test='true'"
	`
	return fmt.Sprintf(testAccAzureRMServiceBusRule_tfTemplate, rInt, location, rInt, rInt, rInt, rInt, correlationFilter)
}

const resourceName = "azurerm_servicebus_rule.test"

const testAccAzureRMServiceBusRule_tfTemplate = `
resource "azurerm_resource_group" "test" {
    name     = "acctestRG-%d"
    location = "%s"
}

resource "azurerm_servicebus_namespace" "test" {
    name                = "acctestservicebusnamespace-%d"
    location            = "${azurerm_resource_group.test.location}"
    resource_group_name = "${azurerm_resource_group.test.name}"
    sku                 = "standard"
}

resource "azurerm_servicebus_topic" "test" {
    name                = "acctestservicebustopic-%d"
    namespace_name      = "${azurerm_servicebus_namespace.test.name}"
    resource_group_name = "${azurerm_resource_group.test.name}"
}

resource "azurerm_servicebus_subscription" "test" {
    name                = "acctestservicebussubscription-%d"
    namespace_name      = "${azurerm_servicebus_namespace.test.name}"
    topic_name          = "${azurerm_servicebus_topic.test.name}"
    resource_group_name = "${azurerm_resource_group.test.name}"
    max_delivery_count  = 10
}

resource "azurerm_servicebus_rule" "test" {
	name                = "acctestservicebusrule-%d"
	namespace_name      = "${azurerm_servicebus_namespace.test.name}"
	topic_name          = "${azurerm_servicebus_topic.test.name}"
	subscription_name   = "${azurerm_servicebus_subscription.test.name}"
	resource_group_name = "${azurerm_resource_group.test.name}"	
	%s
  }
`
