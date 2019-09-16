package azurerm

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func TestAccAzureRMDataFactoryIntegrationRuntimeManaged_basic(t *testing.T) {
	ri := tf.AccRandTimeInt()
	config := testAccAzureRMDataFactoryIntegrationRuntimeManaged_basic(ri, testLocation())
	resourceName := "azurerm_data_factory_integration_runtime_managed.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMDataFactoryIntegrationRuntimeManagedDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMDataFactoryIntegrationRuntimeManagedExists(resourceName),
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

func TestAccAzureRMDataFactoryIntegrationRuntimeManaged_vnetIntegration(t *testing.T) {
	ri := tf.AccRandTimeInt()
	config := testAccAzureRMDataFactoryIntegrationRuntimeManaged_vnetIntegration(ri, testLocation())
	resourceName := "azurerm_data_factory_integration_runtime_managed.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMDataFactoryIntegrationRuntimeManagedDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMDataFactoryIntegrationRuntimeManagedExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "vnet_integration.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "vnet_integration.0.vnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "vnet_integration.0.subnet_name"),
				),
			},
		},
	})
}

func testAccAzureRMDataFactoryIntegrationRuntimeManaged_basic(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_data_factory" "test" {
  name                = "acctestdfirm%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

resource "azurerm_data_factory_integration_runtime_managed" "test" {
  name                = "managed-integration-runtime"
  data_factory_name   = azurerm_data_factory.test.name
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location

  node_size                        = "Standard_D8_v3"
  number_of_nodes                  = 2
  max_parallel_executions_per_node = 8
  edition                          = "Standard"
  license_type                     = "LicenseIncluded"
}
`, rInt, location, rInt)
}

func testAccAzureRMDataFactoryIntegrationRuntimeManaged_vnetIntegration(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctestvnet%d"
  address_space       = ["10.0.0.0/16"]
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

resource "azurerm_subnet" "test" {
  name                 = "acctestsubnet%d"
  resource_group_name  = "${azurerm_resource_group.test.name}"
  virtual_network_name = "${azurerm_virtual_network.test.name}"
  address_prefix       = "10.0.2.0/24"
}

resource "azurerm_data_factory" "test" {
  name                = "acctestdfirm%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

resource "azurerm_data_factory_integration_runtime_managed" "test" {
  name                = "managed-integration-runtime"
  data_factory_name   = azurerm_data_factory.test.name
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location

  node_size                        = "Standard_D8_v3"
  number_of_nodes                  = 2
  max_parallel_executions_per_node = 8
  edition                          = "Standard"
  license_type                     = "LicenseIncluded"

  vnet_integration {
    vnet_id     = "${azurerm_virtual_network.test.id}"
    subnet_name = "${azurerm_subnet.test.name}"
  }
}
`, rInt, location, rInt, rInt, rInt)
}

func testCheckAzureRMDataFactoryIntegrationRuntimeManagedExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		name := rs.Primary.Attributes["name"]
		factoryName := rs.Primary.Attributes["data_factory_name"]
		resourceGroup, hasResourceGroup := rs.Primary.Attributes["resource_group_name"]
		if !hasResourceGroup {
			return fmt.Errorf("Bad: no resource group found in state for Data Factory Managed Integration Runtime: %s", name)
		}

		client := testAccProvider.Meta().(*ArmClient).dataFactory.IntegrationRuntimesClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext

		resp, err := client.Get(ctx, resourceGroup, factoryName, name, "")
		if err != nil {
			return fmt.Errorf("Bad: Get on IntegrationRuntimesClient: %+v", err)
		}

		if utils.ResponseWasNotFound(resp.Response) {
			return fmt.Errorf("Bad: Data Factory Managed Integration Runtime %q (Resource Group: %q, Data Factory %q) does not exist", name, factoryName, resourceGroup)
		}

		return nil
	}
}

func testCheckAzureRMDataFactoryIntegrationRuntimeManagedDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*ArmClient).dataFactory.IntegrationRuntimesClient
	ctx := testAccProvider.Meta().(*ArmClient).StopContext

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azurerm_data_factory_integration_managed" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		factoryName := rs.Primary.Attributes["data_factory_name"]
		resourceGroup := rs.Primary.Attributes["resource_group_name"]

		resp, err := client.Get(ctx, resourceGroup, factoryName, name, "")

		if err != nil {
			return nil
		}

		if resp.StatusCode != http.StatusNotFound {
			return fmt.Errorf("Data Factory Managed Integration Runtime still exists:\n%#v", resp.Properties)
		}
	}

	return nil
}
