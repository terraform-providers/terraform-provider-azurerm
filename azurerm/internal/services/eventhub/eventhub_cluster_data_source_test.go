package eventhub_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance/check"
)

type EventHubClusterDataSource struct {
}

func TestAccEventHubClusterDataSource_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_eventhub_cluster", "test")
	r := EventHubClusterDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("sku_name").HasValue("Dedicated_1"),
			),
		},
	})
}

func (EventHubClusterDataSource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-eventhub-%[1]d"
  location = "%[2]s"
}

resource "azurerm_eventhub_cluster" "test" {
  name                = "acctesteventhubclusTER-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku_name            = "Dedicated_1"
}

data "azurerm_eventhub_cluster" "test" {
  name                = azurerm_eventhub_cluster.test.name
  resource_group_name = azurerm_resource_group.test.name
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger)
}
