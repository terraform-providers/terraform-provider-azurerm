package azurerm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
)

func TestAccDataSourceAzureRMSqlVirtualMachine_basic(t *testing.T) {
	dataSourceName := "data.azurerm_sql_virtual_machine.test"
	ri := tf.AccRandTimeInt()
	location := testLocation()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSqlVirtualMachine_basic(ri, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "resource_group_name"),
				),
			},
		},
	})
}
func TestAccDataSourceAzureRMSqlVirtualMachine_complete(t *testing.T) {
	dataSourceName := "data.azurerm_sql_virtual_machine.test"
	ri := tf.AccRandTimeInt()
	location := testLocation()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSqlVirtualMachine_complete(ri, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "resource_group_name"),
				),
			},
		},
	})
}

func testAccDataSourceSqlVirtualMachine_basic(rInt int, location string) string {
	config := testAccAzureRMSqlVirtualMachine_basic(rInt, location)
	return fmt.Sprintf(`
%s

data "azurerm_sql_virtual_machine" "test" {
  resource_group           = "${azurerm_sql_virtual_machine.test.resource_group}"
  sql_virtual_machine_name = "${azurerm_sql_virtual_machine.test.sql_virtual_machine_name}"
}
`, config)
}

func testAccDataSourceSqlVirtualMachine_complete(rInt int, location string) string {
	config := testAccAzureRMSqlVirtualMachine_complete(rInt, location)
	return fmt.Sprintf(`
%s

data "azurerm_sql_virtual_machine" "test" {
  resource_group           = "${azurerm_sql_virtual_machine.test.resource_group}"
  sql_virtual_machine_name = "${azurerm_sql_virtual_machine.test.sql_virtual_machine_name}"
}
`, config)
}
