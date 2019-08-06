package azurerm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/response"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func TestAccAzureRMProximityPlacementGroup_basic(t *testing.T) {
	resourceName := "azurerm_proximity_placement_group.test"
	ri := tf.AccRandTimeInt()
	config := testAccAzureRMProximityPlacementGroup_basic(ri, testLocation())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMProximityPlacementGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMProximityPlacementGroupExists(resourceName),
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

func TestAccAzureRMProximityPlacementGroup_requiresImport(t *testing.T) {
	if !requireResourcesToBeImported {
		t.Skip("Skipping since resources aren't required to be imported")
		return
	}

	resourceName := "azurerm_proximity_placement_group.test"
	ri := tf.AccRandTimeInt()
	location := testLocation()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMProximityPlacementGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMProximityPlacementGroup_basic(ri, location),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMProximityPlacementGroupExists(resourceName),
				),
			},
			{
				Config:      testAccAzureRMProximityPlacementGroup_requiresImport(ri, location),
				ExpectError: testRequiresImportError("azurerm_proximity_placement_group"),
			},
		},
	})
}

func TestAccAzureRMProximityPlacementGroup_disappears(t *testing.T) {
	resourceName := "azurerm_proximity_placement_group.test"
	ri := tf.AccRandTimeInt()
	config := testAccAzureRMProximityPlacementGroup_basic(ri, testLocation())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMProximityPlacementGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMProximityPlacementGroupExists(resourceName),
					testCheckAzureRMProximityPlacementGroupDisappears(resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAzureRMProximityPlacementGroup_withTags(t *testing.T) {
	resourceName := "azurerm_proximity_placement_group.test"
	ri := tf.AccRandTimeInt()
	location := testLocation()
	preConfig := testAccAzureRMProximityPlacementGroup_withTags(ri, location)
	postConfig := testAccAzureRMProximityPlacementGroup_withUpdatedTags(ri, location)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMProximityPlacementGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: preConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMProximityPlacementGroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.environment", "Production"),
					resource.TestCheckResourceAttr(resourceName, "tags.cost_center", "MSFT"),
				),
			},
			{
				Config: postConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMProximityPlacementGroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.environment", "staging"),
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

func testCheckAzureRMProximityPlacementGroupExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		ppgName := rs.Primary.Attributes["name"]
		resourceGroup, hasResourceGroup := rs.Primary.Attributes["resource_group_name"]
		if !hasResourceGroup {
			return fmt.Errorf("Bad: no resource group found in state for proximity placement group: %s", ppgName)
		}

		client := testAccProvider.Meta().(*ArmClient).ppgClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext
		resp, err := client.Get(ctx, resourceGroup, ppgName)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Availability Set %q (resource group: %q) does not exist", ppgName, resourceGroup)
			}

			return fmt.Errorf("Bad: Get on ppgClient: %+v", err)
		}

		return nil
	}
}

func testCheckAzureRMProximityPlacementGroupDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		ppgName := rs.Primary.Attributes["name"]
		resourceGroup, hasResourceGroup := rs.Primary.Attributes["resource_group_name"]
		if !hasResourceGroup {
			return fmt.Errorf("Bad: no resource group found in state for proximity placement group: %s", ppgName)
		}

		client := testAccProvider.Meta().(*ArmClient).ppgClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext
		resp, err := client.Delete(ctx, resourceGroup, ppgName)
		if err != nil {
			if !response.WasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Delete on ppgClient: %+v", err)
			}
		}

		return nil
	}
}

func testCheckAzureRMProximityPlacementGroupDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azurerm_proximity_placement_group" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		resourceGroup := rs.Primary.Attributes["resource_group_name"]

		client := testAccProvider.Meta().(*ArmClient).ppgClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext
		resp, err := client.Get(ctx, resourceGroup, name)

		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return nil
			}
			return err
		}

		return fmt.Errorf("Proximity placement group still exists:\n%#v", resp.ProximityPlacementGroupProperties)
	}

	return nil
}

func testAccAzureRMProximityPlacementGroup_basic(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_proximity_placement_group" "test" {
  name                = "acctestppg-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}
`, rInt, location, rInt)
}

func testAccAzureRMProximityPlacementGroup_requiresImport(rInt int, location string) string {
	template := testAccAzureRMProximityPlacementGroup_basic(rInt, location)
	return fmt.Sprintf(`
%s

resource "azurerm_proximity_placement_group" "import" {
  name                = "${azurerm_proximity_placement_group.test.name}"
  location            = "${azurerm_proximity_placement_group.test.location}"
  resource_group_name = "${azurerm_proximity_placement_group.test.resource_group_name}"
}
`, template)
}

func testAccAzureRMProximityPlacementGroup_withTags(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_proximity_placement_group" "test" {
  name                = "acctestppg-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  tags = {
    environment = "Production"
    cost_center = "MSFT"
  }
}
`, rInt, location, rInt)
}

func testAccAzureRMProximityPlacementGroup_withUpdatedTags(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_proximity_placement_group" "test" {
  name                = "acctestppg-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  tags = {
    environment = "staging"
  }
}
`, rInt, location, rInt)
}
