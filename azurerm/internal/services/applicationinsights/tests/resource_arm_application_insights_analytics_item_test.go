package tests

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/applicationinsights"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
)

func TestAccAzureRMApplicationInsightsAnalyticsItem_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_application_insights_analytics_item", "test")
	config := testAccAzureRMApplicationInsightsAnalyticsItem_basic(data)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMApplicationInsightAnalyticsItemDestroy(),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMApplicationInsightsAnalyticsItemExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "name", "testquery"),
					resource.TestCheckResourceAttr(data.ResourceName, "scope", "shared"),
					resource.TestCheckResourceAttr(data.ResourceName, "type", "query"),
					resource.TestCheckResourceAttr(data.ResourceName, "content", "requests #test"),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMApplicationInsightsAnalyticsItem_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_application_insights_analytics_item", "test")
	config1 := testAccAzureRMApplicationInsightsAnalyticsItem_basic(data)
	config2 := testAccAzureRMApplicationInsightsAnalyticsItem_basic2(data)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMApplicationInsightAnalyticsItemDestroy(),
		Steps: []resource.TestStep{
			{
				Config: config1,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMApplicationInsightsAnalyticsItemExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "name", "testquery"),
					resource.TestCheckResourceAttr(data.ResourceName, "scope", "shared"),
					resource.TestCheckResourceAttr(data.ResourceName, "type", "query"),
					resource.TestCheckResourceAttr(data.ResourceName, "content", "requests #test"),
				),
			},
			data.ImportStep(),
			{
				Config: config2,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMApplicationInsightsAnalyticsItemExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "name", "testquery"),
					resource.TestCheckResourceAttr(data.ResourceName, "scope", "shared"),
					resource.TestCheckResourceAttr(data.ResourceName, "type", "query"),
					resource.TestCheckResourceAttr(data.ResourceName, "content", "requests #updated"),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMApplicationInsightsAnalyticsItem_multiple(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_application_insights_analytics_item", "test")
	resourceName2 := "azurerm_application_insights_analytics_item.test2"
	resourceName3 := "azurerm_application_insights_analytics_item.test3"
	config := testAccAzureRMApplicationInsightsAnalyticsItem_multiple(data)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMApplicationInsightAnalyticsItemDestroy(),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMApplicationInsightsAnalyticsItemExists(data.ResourceName),
					testCheckAzureRMApplicationInsightsAnalyticsItemExists(resourceName2),
					testCheckAzureRMApplicationInsightsAnalyticsItemExists(resourceName3),
					resource.TestCheckResourceAttr(data.ResourceName, "name", "testquery1"),
					resource.TestCheckResourceAttr(data.ResourceName, "scope", "shared"),
					resource.TestCheckResourceAttr(data.ResourceName, "type", "query"),
					resource.TestCheckResourceAttr(data.ResourceName, "content", "requests #test1"),
					resource.TestCheckResourceAttr(resourceName2, "name", "testquery2"),
					resource.TestCheckResourceAttr(resourceName2, "scope", "user"),
					resource.TestCheckResourceAttr(resourceName2, "type", "query"),
					resource.TestCheckResourceAttr(resourceName2, "content", "requests #test2"),
					resource.TestCheckResourceAttr(resourceName3, "name", "testfunction1"),
					resource.TestCheckResourceAttr(resourceName3, "scope", "shared"),
					resource.TestCheckResourceAttr(resourceName3, "type", "function"),
					resource.TestCheckResourceAttr(resourceName3, "content", "requests #test3"),
					resource.TestCheckResourceAttr(resourceName3, "function_alias", "myfunction"),
				),
			},
			data.ImportStep(),
			{
				ResourceName:      resourceName2,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      resourceName3,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testCheckAzureRMApplicationInsightAnalyticsItemDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "azurerm_application_insights_analytics_item" {
				continue
			}
			name := rs.Primary.Attributes["name"]

			exists, err := testCheckAzureRMApplicationInsightsAnalyticsItemExistsInternal(rs)
			if err != nil {
				return fmt.Errorf("Error checking if item has been destroyed: %s", err)
			}
			if exists {
				return fmt.Errorf("Bad: Application Insights AnalyticsItem '%q' still exists", name)
			}
		}

		return nil
	}
}

func testCheckAzureRMApplicationInsightsAnalyticsItemExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}
		name := rs.Primary.Attributes["name"]

		exists, err := testCheckAzureRMApplicationInsightsAnalyticsItemExistsInternal(rs)
		if err != nil {
			return fmt.Errorf("Error checking if item exists: %s", err)
		}
		if !exists {
			return fmt.Errorf("Bad: Application Insights AnalyticsItem '%q' does not exist", name)
		}

		return nil
	}
}

func testCheckAzureRMApplicationInsightsAnalyticsItemExistsInternal(rs *terraform.ResourceState) (bool, error) {
	id := rs.Primary.Attributes["id"]

	resGroup, appInsightsName, itemScopePath, itemID, err := applicationinsights.ResourcesArmApplicationInsightsAnalyticsItemParseID(id)
	if err != nil {
		return false, fmt.Errorf("Failed to parse ID (id: %s): %+v", id, err)
	}

	conn := acceptance.AzureProvider.Meta().(*clients.Client).AppInsights.AnalyticsItemsClient
	ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

	response, err := conn.Get(ctx, resGroup, appInsightsName, itemScopePath, itemID, "")
	if err != nil {
		if response.Response.IsHTTPStatus(404) {
			return false, nil
		}
		return false, fmt.Errorf("Bad: Get on appInsightsAnalyticsItemsClient (id: %s): %+v", id, err)
	}
	_ = response

	return true, nil
}

func testAccAzureRMApplicationInsightsAnalyticsItem_basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_application_insights" "test" {
  name                = "acctestappinsights-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  application_type    = "web"
}

resource "azurerm_application_insights_analytics_item" "test" {
  name                    = "testquery"
  application_insights_id = "${azurerm_application_insights.test.id}"
  content                 = "requests #test"
  scope                   = "shared"
  type                    = "query"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger)
}

func testAccAzureRMApplicationInsightsAnalyticsItem_basic2(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_application_insights" "test" {
  name                = "acctestappinsights-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  application_type    = "web"
}

resource "azurerm_application_insights_analytics_item" "test" {
  name                    = "testquery"
  application_insights_id = "${azurerm_application_insights.test.id}"
  content                 = "requests #updated"
  scope                   = "shared"
  type                    = "query"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger)
}

func testAccAzureRMApplicationInsightsAnalyticsItem_multiple(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_application_insights" "test" {
  name                = "acctestappinsights-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  application_type    = "web"
}

resource "azurerm_application_insights_analytics_item" "test1" {
  name                    = "testquery1"
  application_insights_id = "${azurerm_application_insights.test.id}"
  content                 = "requests #test1"
  scope                   = "shared"
  type                    = "query"
}

resource "azurerm_application_insights_analytics_item" "test2" {
  name                    = "testquery2"
  application_insights_id = "${azurerm_application_insights.test.id}"
  content                 = "requests #test2"
  scope                   = "user"
  type                    = "query"
}

resource "azurerm_application_insights_analytics_item" "test3" {
  name                    = "testfunction1"
  application_insights_id = "${azurerm_application_insights.test.id}"
  content                 = "requests #test3"
  scope                   = "shared"
  type                    = "function"
  function_alias          = "myfunction"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger)
}
