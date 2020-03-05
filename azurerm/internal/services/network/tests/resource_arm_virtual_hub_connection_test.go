package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
	networkSvc "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/network"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func TestAccAzureRMVirtualHubConnection_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_hub_connection", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMVirtualHubConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMVirtualHubConnection_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMVirtualHubConnectionExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMVirtualHubConnection_requiresImport(t *testing.T) {
	if !features.ShouldResourcesBeImported() {
		t.Skip("Skipping since resources aren't required to be imported")
		return
	}

	data := acceptance.BuildTestData(t, "azurerm_virtual_hub_connection", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMVirtualHubConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMVirtualHubConnection_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMVirtualHubConnectionExists(data.ResourceName),
				),
			},
			{
				Config:      testAccAzureRMVirtualHubConnection_requiresImport(data),
				ExpectError: acceptance.RequiresImportError("azurerm_virtual_hub_connection"),
			},
		},
	})
}

func TestAccAzureRMVirtualHubConnection_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_virtual_hub_connection", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMVirtualHubConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMVirtualHubConnection_complete(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMVirtualHubConnectionExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func testCheckAzureRMVirtualHubConnectionExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.AzureProvider.Meta().(*clients.Client).Network.VirtualHubClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Virtual Hub Connection not found: %s", resourceName)
		}

		virtualHubId := rs.Primary.Attributes["virtual_hub_id"]
		id, err := networkSvc.ParseVirtualHubID(virtualHubId)
		if err != nil {
			return err
		}

		resourceGroup := id.ResourceGroup
		hubName := id.Name
		name := rs.Primary.Attributes["name"]
		resp, err := client.Get(ctx, resourceGroup, hubName)
		if err != nil {
			return fmt.Errorf("Bad: Get on network.VirtualHubClient: %+v", err)
		}

		if resp.VirtualHubProperties == nil {
			return fmt.Errorf("VirtualHubProperties was nil!")
		}

		props := *resp.VirtualHubProperties
		if props.VirtualNetworkConnections == nil {
			return fmt.Errorf("props.VirtualNetworkConnections was nil")
		}

		conns := *props.VirtualNetworkConnections

		found := false
		for _, conn := range conns {
			if conn.Name != nil && *conn.Name == name {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("Connection %q was not found", name)
		}

		return nil
	}
}

func testCheckAzureRMVirtualHubConnectionDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		client := acceptance.AzureProvider.Meta().(*clients.Client).Network.VirtualHubClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		if rs.Type != "azurerm_virtual_hub_connection" {
			continue
		}

		virtualHubId := rs.Primary.Attributes["virtual_hub_id"]
		id, err := networkSvc.ParseVirtualHubID(virtualHubId)
		if err != nil {
			return err
		}

		resourceGroup := id.ResourceGroup
		hubName := id.Name
		name := rs.Primary.Attributes["name"]

		resp, err := client.Get(ctx, resourceGroup, hubName)
		if err != nil {
			if !utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Get on network.VirtualHubClient: %+v", err)
			}
		}

		// since it's been deleted, that's fine
		if resp.VirtualHubProperties == nil {
			return nil
		}
		props := *resp.VirtualHubProperties
		if props.VirtualNetworkConnections == nil {
			return nil
		}

		conns := *props.VirtualNetworkConnections
		for _, conn := range conns {
			if conn.Name != nil && *conn.Name == name {
				return fmt.Errorf("Connection %q still exists", name)
			}
		}

		return nil
	}

	return nil
}

func testAccAzureRMVirtualHubConnection_basic(data acceptance.TestData) string {
	template := testAccAzureRMVirtualHubConnection_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_virtual_hub_connection" "test" {
  name                      = "acctestvhubconn-%d"
  virtual_hub_id            = azurerm_virtual_hub.test.id
  remote_virtual_network_id = azurerm_virtual_network.test.id
}
`, template, data.RandomInteger)
}

func testAccAzureRMVirtualHubConnection_requiresImport(data acceptance.TestData) string {
	template := testAccAzureRMVirtualHubConnection_basic(data)
	return fmt.Sprintf(`
%s

resource "azurerm_virtual_hub_connection" "import" {
  name                      = azurerm_virtual_hub_connection.test.name
  virtual_hub_id            = azurerm_virtual_hub_connection.test.virtual_hub_id
  remote_virtual_network_id = azurerm_virtual_hub_connection.test.remote_virtual_network_id
}
`, template)
}

func testAccAzureRMVirtualHubConnection_complete(data acceptance.TestData) string {
	template := testAccAzureRMVirtualHubConnection_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_virtual_hub_connection" "test" {
  name                                       = "acctestvhubconn-%d"
  virtual_hub_id                             = azurerm_virtual_hub.test.id
  remote_virtual_network_id                  = azurerm_virtual_network.test.id
  allow_hub_to_remote_vnet_transit           = true
  allow_remote_vnet_to_use_hub_vnet_gateways = false
  enable_internet_security                   = true
}
`, template, data.RandomInteger)
}

func testAccAzureRMVirtualHubConnection_template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestrg-%d"
  location = "%s"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctestvirtnet%d"
  address_space       = ["10.5.0.0/16"]
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_network_security_group" "test" {
  name                = "acctestnsg%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_subnet" "test" {
  name                 = "acctestsubnet%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefix       = "10.5.1.0/24"
}

resource "azurerm_subnet_network_security_group_association" "test" {
  subnet_id                 = azurerm_subnet.test.id
  network_security_group_id = azurerm_network_security_group.test.id
}

resource "azurerm_virtual_wan" "test" {
  name                = "acctestvwan-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}

resource "azurerm_virtual_hub" "test" {
  name                = "acctestvhub-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  virtual_wan_id      = azurerm_virtual_wan.test.id
  address_prefix      = "10.0.2.0/24"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}
