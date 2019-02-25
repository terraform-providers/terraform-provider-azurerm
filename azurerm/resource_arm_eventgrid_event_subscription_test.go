package azurerm

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func TestAccAzureRMEventGridEventSubscription_basic(t *testing.T) {
	resourceName := "azurerm_eventgrid_event_subscription.test"
	ri := tf.AccRandTimeInt()
	rs := strings.ToLower(acctest.RandString(11))

	location := testLocation()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMEventGridEventSubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMEventGridEventSubscription_basic(ri, rs, location),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventGridEventSubscriptionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "storage_queue_endpoint.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureRMEventGridEventSubscription_eventhub(t *testing.T) {
	resourceName := "azurerm_eventgrid_event_subscription.test"
	ri := tf.AccRandTimeInt()

	location := testLocation()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMEventGridEventSubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMEventGridEventSubscription_eventhub(ri, location),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventGridEventSubscriptionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "eventhub_endpoint.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureRMEventGridEventSubscription_update(t *testing.T) {
	resourceName := "azurerm_eventgrid_event_subscription.test"
	ri := tf.AccRandTimeInt()
	rs := strings.ToLower(acctest.RandString(11))

	location := testLocation()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMEventGridEventSubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMEventGridEventSubscription_eventhub(ri, location),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventGridEventSubscriptionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "eventhub_endpoint.#", "1"),
				),
			},
			{
				Config: testAccAzureRMEventGridEventSubscription_basic(ri, rs, location),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMEventGridEventSubscriptionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "storage_queue_endpoint.#", "1"),
				),
			},
		},
	})
}

func testCheckAzureRMEventGridEventSubscriptionDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*ArmClient).eventGridEventSubscriptionsClient
	ctx := testAccProvider.Meta().(*ArmClient).StopContext

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azurerm_eventgrid_event_subscription" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		scope := rs.Primary.Attributes["scope"]

		resp, err := client.Get(ctx, scope, name)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return err
		}

		if resp.StatusCode != http.StatusNotFound {
			return fmt.Errorf("EventGrid Event Subscription still exists:\n%#v", resp)
		}
	}

	return nil
}

func testCheckAzureRMEventGridEventSubscriptionExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		name := rs.Primary.Attributes["name"]
		scope, hasScope := rs.Primary.Attributes["scope"]
		if !hasScope {
			return fmt.Errorf("Bad: no scope found in state for EventGrid Event Subscription: %s", name)
		}

		client := testAccProvider.Meta().(*ArmClient).eventGridEventSubscriptionsClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext
		resp, err := client.Get(ctx, scope, name)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: EventGrid Event Subscription %q (scope: %s) does not exist", name, scope)
			}

			return fmt.Errorf("Bad: Get on eventGridEventSubscriptionsClient: %s", err)
		}

		return nil
	}
}

func testAccAzureRMEventGridEventSubscription_basic(rInt int, rString string, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_storage_account" "test" {
  name                     = "acctestacc%s"
  resource_group_name      = "${azurerm_resource_group.test.name}"
  location                 = "${azurerm_resource_group.test.location}"
  account_tier             = "Standard"
  account_replication_type = "LRS"
  
  tags {
    environment = "staging"
  }
}
  
resource "azurerm_storage_queue" "test" {
  name                 = "mysamplequeue-%d"
  resource_group_name  = "${azurerm_resource_group.test.name}"
  storage_account_name = "${azurerm_storage_account.test.name}"
}

resource "azurerm_eventgrid_event_subscription" "test" {
  name                = "acctesteg-%d"
  scope = "${azurerm_resource_group.test.id}"
  storage_queue_endpoint {
	  storage_account_id = "${azurerm_storage_account.test.id}"
	  queue_name = "${azurerm_storage_queue.test.name}"
  }
}
`, rInt, location, rString, rInt, rInt)
}

func testAccAzureRMEventGridEventSubscription_eventhub(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}
resource "azurerm_eventhub_namespace" "test" {
  name                = "acctesteventhubnamespace-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  sku                 = "Basic"
}

resource "azurerm_eventhub" "test" {
  name                = "acctesteventhub-%d"
  namespace_name      = "${azurerm_eventhub_namespace.test.name}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  partition_count     = 2
  message_retention   = 1
}

resource "azurerm_eventgrid_event_subscription" "test" {
  name                = "acctesteg-%d"
  scope = "${azurerm_resource_group.test.id}"
  eventhub_endpoint {
	  eventhub_id = "${azurerm_eventhub.test.id}"
  }
}
`, rInt, location, rInt, rInt, rInt)
}
