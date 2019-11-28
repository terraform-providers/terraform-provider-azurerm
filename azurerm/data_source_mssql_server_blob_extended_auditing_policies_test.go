package azurerm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
)

func TestAccDataSourceAzureRMMsSqlServerBlobExtendedAuditingPolicies_basic(t *testing.T) {
	dataSourceName := "data.azurerm_mssql_server_blob_extended_auditing_policies.test"
	ri := tf.AccRandTimeInt()
	location := testLocation()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAzureRMMsSqlServerBlobExtendedAuditingPolicies_basic(ri, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "resource_group_name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "server_name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "state"),
					resource.TestCheckResourceAttrSet(dataSourceName, "retention_days"),
					resource.TestCheckResourceAttrSet(dataSourceName, "storage_account_subscription_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "is_storage_secondary_key_in_use"),
					resource.TestCheckResourceAttrSet(dataSourceName, "is_azure_monitor_target_enabled"),
				),
			},
		},
	})
}

func testAccDataSourceAzureRMMsSqlServerBlobExtendedAuditingPolicies_basic(rInt int, location string) string {
	config := testAccAzureRMMsSqlServerBlobAuditingPolicies_basic(rInt, location)
	return fmt.Sprintf(`
%s
data "azurerm_mssql_server_blob_extended_auditing_policies" "test" {
  resource_group_name = "${azurerm_resource_group.test.name}"
  server_name         = "${azurerm_sql_server.test.name}"
}
`, config)
}
