package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
)

func TestAccDataSourceAzureRMSqlServer_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_sql_server", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMSqlServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAzureRMSqlServer_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSqlServerExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "location"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "fqdn"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "version"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "administrator_login"),
					resource.TestCheckResourceAttr(data.ResourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "blob_extended_auditing_policy.0.state", "Disabled"),
				),
			},
		},
	})
}

func TestAccDataSourceAzureRMSqlServer_withBlobAuditing(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_sql_server", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMSqlServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAzureRMSqlServer_withBlobAuditing(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMSqlServerExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "blob_extended_auditing_policy.0.state"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "blob_extended_auditing_policy.0.retention_days"),
				),
			},
		},
	})
}

func testAccDataSourceAzureRMSqlServer_basic(data acceptance.TestData) string {
	template := testAccAzureRMSqlServer_basic(data)
	return fmt.Sprintf(`
%s

data "azurerm_sql_server" "test" {
  name                = "${azurerm_sql_server.test.name}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}
`, template)
}

func testAccDataSourceAzureRMSqlServer_withBlobAuditing(data acceptance.TestData) string {
	template := testAccAzureRMSqlServer_withBlobAuditingPolices(data)
	return fmt.Sprintf(`
%s

data "azurerm_sql_server" "test" {
  name                = "${azurerm_sql_server.test.name}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}
`, template)
}
