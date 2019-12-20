package datalake

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
)

func TestAccDataSourceAzureRMDataLakeStore_basic(t *testing.T) {
	dataSourceName := "data.azurerm_data_lake_store.test"
	rInt := tf.AccRandTimeInt()
	rs := acctest.RandString(4)
	location := acceptance.Location()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDataLakeStore_basic(rInt, rs, location),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMDataLakeStoreExists(dataSourceName),
					resource.TestCheckResourceAttr(dataSourceName, "tier", "Consumption"),
				),
			},
		},
	})
}

func TestAccDataSourceAzureRMDataLakeStore_tier(t *testing.T) {
	dataSourceName := "data.azurerm_data_lake_store.test"
	rInt := tf.AccRandTimeInt()
	rs := acctest.RandString(4)
	location := acceptance.Location()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDataLakeStore_tier(rInt, rs, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "tier", "Commitment_1TB"),
					resource.TestCheckResourceAttr(dataSourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "tags.hello", "world"),
				),
			},
		},
	})
}

func testAccDataSourceDataLakeStore_basic(rInt int, rs string, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_data_lake_store" "test" {
  name                = "unlikely23exst2acct%s"
  location            = "%s"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

data "azurerm_data_lake_store" "test" {
  name                = "${azurerm_data_lake_store.test.name}"
  resource_group_name = "${azurerm_data_lake_store.test.resource_group_name}"
}
`, rInt, location, rs, location)
}

func testAccDataSourceDataLakeStore_tier(rInt int, rs string, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_data_lake_store" "test" {
  name                = "unlikely23exst2acct%s"
  location            = "%s"
  tier                = "Commitment_1TB"
  resource_group_name = "${azurerm_resource_group.test.name}"

  tags = {
    hello = "world"
  }
}

data "azurerm_data_lake_store" "test" {
  name                = "${azurerm_data_lake_store.test.name}"
  resource_group_name = "${azurerm_data_lake_store.test.resource_group_name}"
}
`, rInt, location, rs, location)
}
