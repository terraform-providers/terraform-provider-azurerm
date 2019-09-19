package azurerm

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"

	"github.com/Azure/azure-sdk-for-go/services/appinsights/mgmt/2015-05-01/insights"
)

func TestAccAzureRMApplicationInsightsAnalyticsItem_basic(t *testing.T) {
	resourceName := "azurerm_application_insights_analytics_item.test"
	ri := tf.AccRandTimeInt()
	config := testAccAzureRMApplicationInsightsAnalyticsItem_basic(ri, testLocation())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMApplicationInsightAnalyticsItemDestroy(),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMApplicationInsightsAnalyticsItemExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "testquery"),
					resource.TestCheckResourceAttr(resourceName, "scope", "shared"),
					resource.TestCheckResourceAttr(resourceName, "type", "query"),
					resource.TestCheckResourceAttr(resourceName, "content", "requests #test"),
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

func TestAccAzureRMApplicationInsightsAnalyticsItem_update(t *testing.T) {
	resourceName := "azurerm_application_insights_analytics_item.test"
	ri := tf.AccRandTimeInt()
	config1 := testAccAzureRMApplicationInsightsAnalyticsItem_basic(ri, testLocation())
	config2 := testAccAzureRMApplicationInsightsAnalyticsItem_basic2(ri, testLocation())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMApplicationInsightAnalyticsItemDestroy(),
		Steps: []resource.TestStep{
			{
				Config: config1,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMApplicationInsightsAnalyticsItemExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "testquery"),
					resource.TestCheckResourceAttr(resourceName, "scope", "shared"),
					resource.TestCheckResourceAttr(resourceName, "type", "query"),
					resource.TestCheckResourceAttr(resourceName, "content", "requests #test"),
				),
			},
			{
				Config: config2,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMApplicationInsightsAnalyticsItemExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "testquery"),
					resource.TestCheckResourceAttr(resourceName, "scope", "shared"),
					resource.TestCheckResourceAttr(resourceName, "type", "query"),
					resource.TestCheckResourceAttr(resourceName, "content", "requests #updated"),
				),
			},
		},
	})
}

func TestAccAzureRMApplicationInsightsAnalyticsItem_multiple(t *testing.T) {
	resourceName1 := "azurerm_application_insights_analytics_item.test1"
	resourceName2 := "azurerm_application_insights_analytics_item.test2"
	resourceName3 := "azurerm_application_insights_analytics_item.test3"
	ri := tf.AccRandTimeInt()
	config := testAccAzureRMApplicationInsightsAnalyticsItem_multiple(ri, testLocation())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMApplicationInsightAnalyticsItemDestroy(),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMApplicationInsightsAnalyticsItemExists(resourceName1),
					testCheckAzureRMApplicationInsightsAnalyticsItemExists(resourceName2),
					testCheckAzureRMApplicationInsightsAnalyticsItemExists(resourceName3),
					resource.TestCheckResourceAttr(resourceName1, "name", "testquery1"),
					resource.TestCheckResourceAttr(resourceName1, "scope", "shared"),
					resource.TestCheckResourceAttr(resourceName1, "type", "query"),
					resource.TestCheckResourceAttr(resourceName1, "content", "requests #test1"),
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
			resGroup := rs.Primary.Attributes["resource_group_name"]

			exists, err := testCheckAzureRMApplicationInsightsAnalyticsItemExistsInternal(rs)
			if err != nil {
				return fmt.Errorf("Error checking if item has been destroyed: %s", err)
			}
			if exists {
				return fmt.Errorf("Bad: Application Insights AnalyticsItem '%q' (resource group: '%q') still exists", name, resGroup)
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
		resGroup := rs.Primary.Attributes["resource_group_name"]

		exists, err := testCheckAzureRMApplicationInsightsAnalyticsItemExistsInternal(rs)
		if err != nil {
			return fmt.Errorf("Error checking if item exists: %s", err)
		}
		if !exists {
			return fmt.Errorf("Bad: Application Insights AnalyticsItem '%q' (resource group: '%q') does not exist", name, resGroup)
		}

		return nil
	}
}

func testCheckAzureRMApplicationInsightsAnalyticsItemExistsInternal(rs *terraform.ResourceState) (bool, error) {
	resGroup := rs.Primary.Attributes["resource_group_name"]
	scopeName := rs.Primary.Attributes["scope"]
	typeName := rs.Primary.Attributes["type"]
	name := rs.Primary.Attributes["name"]

	itemScope := insights.ItemScope(scopeName)
	itemType := insights.ItemTypeParameter(typeName)

	appInsightsID := rs.Primary.Attributes["application_insights_id"]
	id, err := azure.ParseAzureResourceID(appInsightsID)
	if err != nil {
		return false, fmt.Errorf("Error parsing resource ID: %s", err)
	}

	appInsightsName := id.Path["components"]

	conn := testAccProvider.Meta().(*ArmClient).appInsights.AnalyticsItemsClient
	ctx := testAccProvider.Meta().(*ArmClient).StopContext

	includeContent := false
	resp, err := conn.List(ctx, resGroup, appInsightsName, insights.AnalyticsItems, itemScope, itemType, &includeContent)
	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("Bad: List on appInsightsAnalyticsItemsClient: %+v", err)
	}
	for _, item := range *resp.Value {
		if *item.Name == name && item.Scope == itemScope {
			return true, nil
		}
	}

	return false, nil
}

func testAccAzureRMApplicationInsightsAnalyticsItem_basic(rInt int, location string) string {
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
  resource_group_name     = "${azurerm_resource_group.test.name}"
  application_insights_id = "${azurerm_application_insights.test.id}"
  content                 = "requests #test"
  scope                   = "shared"
  type                    = "query"
}
`, rInt, location, rInt)
}

func testAccAzureRMApplicationInsightsAnalyticsItem_basic2(rInt int, location string) string {
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
  resource_group_name     = "${azurerm_resource_group.test.name}"
  application_insights_id = "${azurerm_application_insights.test.id}"
  content                 = "requests #updated"
  scope                   = "shared"
  type                    = "query"
}
`, rInt, location, rInt)
}

func testAccAzureRMApplicationInsightsAnalyticsItem_multiple(rInt int, location string) string {
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
  resource_group_name     = "${azurerm_resource_group.test.name}"
  application_insights_id = "${azurerm_application_insights.test.id}"
  content                 = "requests #test1"
  scope                   = "shared"
  type                    = "query"
}

resource "azurerm_application_insights_analytics_item" "test2" {
  name                    = "testquery2"
  resource_group_name     = "${azurerm_resource_group.test.name}"
  application_insights_id = "${azurerm_application_insights.test.id}"
  content                 = "requests #test2"
  scope                   = "user"
  type                    = "query"
}

resource "azurerm_application_insights_analytics_item" "test3" {
  name                    = "testfunction1"
  resource_group_name     = "${azurerm_resource_group.test.name}"
  application_insights_id = "${azurerm_application_insights.test.id}"
  content                 = "requests #test3"
  scope                   = "shared"
  type                    = "function"
  function_alias          = "myfunction"
}
`, rInt, location, rInt)
}
