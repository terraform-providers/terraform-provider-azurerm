package azurerm

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/response"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func TestAccAzureRMSubnet_basic(t *testing.T) {
	resourceName := "azurerm_subnet.test"
	ri := acctest.RandInt()
	config := testAccAzureRMSubnet_basic(ri, testLocation())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSubnetExists(resourceName),
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

func TestAccAzureRMSubnet_delegation(t *testing.T) {
	resourceName := "azurerm_subnet.test"
	ri := acctest.RandInt()
	config := testAccAzureRMSubnet_delegation(ri, testLocation())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSubnetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "delegation.#", "1"),
				),
			},
		},
	})
}

func TestAccAzureRMSubnet_routeTableUpdate(t *testing.T) {
	resourceName := "azurerm_subnet.test"
	ri := acctest.RandInt()
	location := testLocation()
	initConfig := testAccAzureRMSubnet_routeTable(ri, location)
	updatedConfig := testAccAzureRMSubnet_updatedRouteTable(ri, location)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: initConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSubnetExists(resourceName),
				),
			},

			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSubnetRouteTableExists(resourceName, fmt.Sprintf("acctest-%d", ri)),
				),
			},
		},
	})
}

func TestAccAzureRMSubnet_routeTableRemove(t *testing.T) {
	resourceName := "azurerm_subnet.test"
	ri := acctest.RandInt()
	location := testLocation()
	initConfig := testAccAzureRMSubnet_routeTable(ri, location)
	updatedConfig := testAccAzureRMSubnet_routeTableUnlinked(ri, location)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: initConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSubnetExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSubnetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "route_table_id", ""),
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

func TestAccAzureRMSubnet_removeNetworkSecurityGroup(t *testing.T) {
	resourceName := "azurerm_subnet.test"
	ri := acctest.RandInt()
	location := testLocation()
	initConfig := testAccAzureRMSubnet_networkSecurityGroup(ri, location)
	updatedConfig := testAccAzureRMSubnet_networkSecurityGroupDetached(ri, location)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: initConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSubnetExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "network_security_group_id"),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSubnetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "network_security_group_id", ""),
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

func TestAccAzureRMSubnet_bug7986(t *testing.T) {
	ri := acctest.RandInt()
	initConfig := testAccAzureRMSubnet_bug7986(ri, testLocation())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: initConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSubnetExists("azurerm_subnet.first"),
					testCheckAzureRMSubnetExists("azurerm_subnet.second"),
				),
			},
		},
	})
}

func TestAccAzureRMSubnet_bug15204(t *testing.T) {
	resourceName := "azurerm_subnet.test"
	ri := acctest.RandInt()
	initConfig := testAccAzureRMSubnet_bug15204(ri, testLocation())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: initConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSubnetExists(resourceName),
				),
			},
		},
	})
}

func TestAccAzureRMSubnet_disappears(t *testing.T) {
	resourceName := "azurerm_subnet.test"
	ri := acctest.RandInt()
	config := testAccAzureRMSubnet_basic(ri, testLocation())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSubnetExists(resourceName),
					testCheckAzureRMSubnetDisappears(resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testCheckAzureRMSubnetExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		log.Printf("[INFO] Checking Subnet addition.")

		name := rs.Primary.Attributes["name"]
		vnetName := rs.Primary.Attributes["virtual_network_name"]
		resourceGroup, hasResourceGroup := rs.Primary.Attributes["resource_group_name"]
		if !hasResourceGroup {
			return fmt.Errorf("Bad: no resource group found in state for subnet: %s", name)
		}

		client := testAccProvider.Meta().(*ArmClient).subnetClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext

		resp, err := client.Get(ctx, resourceGroup, vnetName, name, "")
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Subnet %q (resource group: %q) does not exist", name, resourceGroup)
			}

			return fmt.Errorf("Bad: Get on subnetClient: %+v", err)
		}

		return nil
	}
}

func testCheckAzureRMSubnetRouteTableExists(resourceName string, routeTableId string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		log.Printf("[INFO] Checking Subnet update.")

		subnetName := rs.Primary.Attributes["name"]
		vnetName := rs.Primary.Attributes["virtual_network_name"]
		resourceGroup, hasResourceGroup := rs.Primary.Attributes["resource_group_name"]
		if !hasResourceGroup {
			return fmt.Errorf("Bad: no resource group found in state for subnet: %s", subnetName)
		}

		networksClient := testAccProvider.Meta().(*ArmClient).vnetClient
		subnetsClient := testAccProvider.Meta().(*ArmClient).subnetClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext

		vnetResp, vnetErr := networksClient.Get(ctx, resourceGroup, vnetName, "")
		if vnetErr != nil {
			return fmt.Errorf("Bad: Get on vnetClient: %+v", vnetErr)
		}

		if vnetResp.Subnets == nil {
			return fmt.Errorf("Bad: Vnet %q (resource group: %q) does not have subnets after update", vnetName, resourceGroup)
		}

		resp, err := subnetsClient.Get(ctx, resourceGroup, vnetName, subnetName, "")
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Subnet %q (resource group: %q) does not exist", subnetName, resourceGroup)
			}

			return fmt.Errorf("Bad: Get on subnetClient: %+v", err)
		}

		if resp.RouteTable == nil {
			return fmt.Errorf("Bad: Subnet %q (resource group: %q) does not contain route tables after update", subnetName, resourceGroup)
		}

		if !strings.Contains(*resp.RouteTable.ID, routeTableId) {
			return fmt.Errorf("Bad: Subnet %q (resource group: %q) does not have route table %q", subnetName, resourceGroup, routeTableId)
		}

		return nil
	}
}

func testCheckAzureRMSubnetDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		name := rs.Primary.Attributes["name"]
		vnetName := rs.Primary.Attributes["virtual_network_name"]
		resourceGroup, hasResourceGroup := rs.Primary.Attributes["resource_group_name"]
		if !hasResourceGroup {
			return fmt.Errorf("Bad: no resource group found in state for subnet: %s", name)
		}

		client := testAccProvider.Meta().(*ArmClient).subnetClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext
		future, err := client.Delete(ctx, resourceGroup, vnetName, name)
		if err != nil {
			if !response.WasNotFound(future.Response()) {
				return fmt.Errorf("Error deleting Subnet %q (Network %q / Resource Group %q): %+v", name, vnetName, resourceGroup, err)
			}
		}

		if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
			return fmt.Errorf("Error waiting for completion of Subnet %q (Network %q / Resource Group %q): %+v", name, vnetName, resourceGroup, err)
		}

		return nil
	}
}

func testCheckAzureRMSubnetDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*ArmClient).subnetClient
	ctx := testAccProvider.Meta().(*ArmClient).StopContext

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azurerm_subnet" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		vnetName := rs.Primary.Attributes["virtual_network_name"]
		resourceGroup := rs.Primary.Attributes["resource_group_name"]

		resp, err := client.Get(ctx, resourceGroup, vnetName, name, "")
		if err != nil {
			if !utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Subnet still exists:\n%#v", resp.SubnetPropertiesFormat)
			}
			return nil
		}
	}

	return nil
}

func TestAccAzureRMSubnet_serviceEndpoints(t *testing.T) {

	ri := acctest.RandInt()
	config := testAccAzureRMSubnet_serviceEndpoints(ri, testLocation())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSubnetExists("azurerm_subnet.test"),
				),
			},
		},
	})
}

func testAccAzureRMSubnet_basic(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctestvirtnet%d"
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
`, rInt, location, rInt, rInt)
}

func testAccAzureRMSubnet_delegation(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctestvirtnet%d"
  address_space       = ["10.0.0.0/16"]
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

resource "azurerm_subnet" "test" {
  name                 = "acctestsubnet%d"
  resource_group_name  = "${azurerm_resource_group.test.name}"
  virtual_network_name = "${azurerm_virtual_network.test.name}"
  address_prefix       = "10.0.2.0/24"
  delegation {
     name = "acctestdelegation"
     service_delegation {
       name    = "Microsoft.ContainerInstance/containerGroups"
       actions = ["Microsoft.Network/virtualNetworks/subnets/action"]
     }
  }
}
`, rInt, location, rInt, rInt)
}

func testAccAzureRMSubnet_routeTable(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctestvirtnet%d"
  address_space       = ["10.0.0.0/16"]
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

resource "azurerm_subnet" "test" {
  name                 = "acctestsubnet%d"
  resource_group_name  = "${azurerm_resource_group.test.name}"
  virtual_network_name = "${azurerm_virtual_network.test.name}"
  address_prefix       = "10.0.2.0/24"
  route_table_id       = "${azurerm_route_table.test.id}"
}

resource "azurerm_route_table" "test" {
  name                = "acctest-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  route {
    name                   = "acctest-%d"
    address_prefix         = "10.100.0.0/14"
    next_hop_type          = "VirtualAppliance"
    next_hop_in_ip_address = "10.10.1.1"
  }
}
`, rInt, location, rInt, rInt, rInt, rInt)
}

func testAccAzureRMSubnet_routeTableUnlinked(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctestvirtnet%d"
  address_space       = ["10.0.0.0/16"]
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

resource "azurerm_subnet" "test" {
  name                 = "acctestsubnet%d"
  resource_group_name  = "${azurerm_resource_group.test.name}"
  virtual_network_name = "${azurerm_virtual_network.test.name}"
  address_prefix       = "10.0.2.0/24"
  route_table_id       = ""
}

resource "azurerm_route_table" "test" {
  name                = "acctest-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  route {
    name                   = "acctest-%d"
    address_prefix         = "10.100.0.0/14"
    next_hop_type          = "VirtualAppliance"
    next_hop_in_ip_address = "10.10.1.1"
  }
}
`, rInt, location, rInt, rInt, rInt, rInt)
}

func testAccAzureRMSubnet_updatedRouteTable(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"

  tags {
    environment = "Testing"
  }
}

resource "azurerm_network_security_group" "test_secgroup" {
  name                = "acctest-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  security_rule {
    name                       = "acctest-%d"
    priority                   = 100
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "*"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }

  tags {
    environment = "Testing"
  }
}

resource "azurerm_virtual_network" "test" {
  name                = "acctestvirtnet%d"
  address_space       = ["10.0.0.0/16"]
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  tags {
    environment = "Testing"
  }
}

resource "azurerm_subnet" "test" {
  name                 = "acctestsubnet%d"
  resource_group_name  = "${azurerm_resource_group.test.name}"
  virtual_network_name = "${azurerm_virtual_network.test.name}"
  address_prefix       = "10.0.2.0/24"
  route_table_id       = "${azurerm_route_table.test.id}"
}

resource "azurerm_route_table" "test" {
  name                = "acctest-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  route {
    name                   = "acctest-%d"
    address_prefix         = "10.100.0.0/14"
    next_hop_type          = "VirtualAppliance"
    next_hop_in_ip_address = "10.10.1.1"
  }

  tags {
    environment = "Testing"
  }
}
`, rInt, location, rInt, rInt, rInt, rInt, rInt, rInt)
}

func testAccAzureRMSubnet_networkSecurityGroup(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctest%d-rg"
  location = "%s"
}

resource "azurerm_network_security_group" "test" {
  name                = "acctestnsg%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  security_rule {
    name                       = "test123"
    priority                   = 100
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "*"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }
}

resource "azurerm_virtual_network" "test" {
  name                = "acctest%d-vn"
  address_space       = ["10.0.0.0/16"]
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

resource "azurerm_subnet" "test" {
  name                      = "acctest%d-private"
  resource_group_name       = "${azurerm_resource_group.test.name}"
  virtual_network_name      = "${azurerm_virtual_network.test.name}"
  address_prefix            = "10.0.0.0/24"
  network_security_group_id = "${azurerm_network_security_group.test.id}"
}
`, rInt, location, rInt, rInt, rInt)
}

func testAccAzureRMSubnet_networkSecurityGroupDetached(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctest%d-rg"
  location = "%s"
}

resource "azurerm_network_security_group" "test" {
  name                = "acctestnsg%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  security_rule {
    name                       = "test123"
    priority                   = 100
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "*"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }
}

resource "azurerm_virtual_network" "test" {
  name                = "acctest%d-vn"
  address_space       = ["10.0.0.0/16"]
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

resource "azurerm_subnet" "test" {
  name                 = "acctest%d-private"
  resource_group_name  = "${azurerm_resource_group.test.name}"
  virtual_network_name = "${azurerm_virtual_network.test.name}"
  address_prefix       = "10.0.0.0/24"
}
`, rInt, location, rInt, rInt, rInt)
}

func testAccAzureRMSubnet_bug7986(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctest%d-rg"
  location = "%s"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctest%d-vn"
  address_space       = ["10.0.0.0/16"]
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

resource "azurerm_route_table" "first" {
  name                = "acctest%d-private-1"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  route {
    name           = "acctest%d-private-1"
    address_prefix = "0.0.0.0/0"
    next_hop_type  = "None"
  }
}

resource "azurerm_subnet" "first" {
  name                 = "acctest%d-private-1"
  resource_group_name  = "${azurerm_resource_group.test.name}"
  virtual_network_name = "${azurerm_virtual_network.test.name}"
  address_prefix       = "10.0.0.0/24"
  route_table_id       = "${azurerm_route_table.first.id}"
}

resource "azurerm_route_table" "second" {
  name                = "acctest%d-private-2"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  route {
    name           = "acctest%d-private-2"
    address_prefix = "0.0.0.0/0"
    next_hop_type  = "None"
  }
}

resource "azurerm_subnet" "second" {
  name                 = "acctest%d-private-2"
  resource_group_name  = "${azurerm_resource_group.test.name}"
  virtual_network_name = "${azurerm_virtual_network.test.name}"
  address_prefix       = "10.0.1.0/24"
  route_table_id       = "${azurerm_route_table.second.id}"
}
`, rInt, location, rInt, rInt, rInt, rInt, rInt, rInt, rInt)
}

func testAccAzureRMSubnet_bug15204(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctest-%d"
  location = "%s"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctestvn-%d"
  address_space       = ["10.85.0.0/16"]
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

resource "azurerm_network_security_group" "test" {
  name                = "acctestnsg-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

resource "azurerm_route_table" "test" {
  name                = "acctestrt-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

resource "azurerm_subnet" "test" {
  name                      = "acctestsubnet-%d"
  resource_group_name       = "${azurerm_resource_group.test.name}"
  virtual_network_name      = "${azurerm_virtual_network.test.name}"
  address_prefix            = "10.85.9.0/24"
  route_table_id            = "${azurerm_route_table.test.id}"
  network_security_group_id = "${azurerm_network_security_group.test.id}"
}
`, rInt, location, rInt, rInt, rInt, rInt)
}

func testAccAzureRMSubnet_serviceEndpoints(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctestvirtnet%d"
  address_space       = ["10.0.0.0/16"]
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

resource "azurerm_subnet" "test" {
  name                 = "acctestsubnet%d"
  resource_group_name  = "${azurerm_resource_group.test.name}"
  virtual_network_name = "${azurerm_virtual_network.test.name}"
  address_prefix       = "10.0.2.0/24"
  service_endpoints    = ["Microsoft.Sql", "Microsoft.Storage"]
}
`, rInt, location, rInt, rInt)
}
