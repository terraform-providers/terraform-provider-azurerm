package azurerm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
)

func TestAccDataSourceAzureRMAppServiceCertificateOrder_basic(t *testing.T) {
	dataSourceName := "data.azurerm_app_service_certificate_order.test"
	rInt := tf.AccRandTimeInt()
	location := testLocation()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAppServiceCertificateOrder_basic(rInt, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "csr"),
					resource.TestCheckResourceAttrSet(dataSourceName, "domain_verification_token"),
					resource.TestCheckResourceAttr(dataSourceName, "distinguished_name", "CN=example.com"),
					resource.TestCheckResourceAttr(dataSourceName, "product_type", "standard"),
				),
			},
		},
	})
}

func TestAccDataSourceAzureRMAppServiceCertificateOrder_wildcard(t *testing.T) {
	dataSourceName := "data.azurerm_app_service_certificate_order.test"
	rInt := tf.AccRandTimeInt()
	location := testLocation()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAppServiceCertificateOrder_wildcard(rInt, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "csr"),
					resource.TestCheckResourceAttrSet(dataSourceName, "domain_verification_token"),
					resource.TestCheckResourceAttr(dataSourceName, "distinguished_name", "CN=example.com"),
					resource.TestCheckResourceAttr(dataSourceName, "product_type", "wildcard"),
				),
			},
		},
	})
}

func TestAccDataSourceAzureRMAppServiceCertificateOrder_complete(t *testing.T) {
	dataSourceName := "data.azurerm_app_service_certificate_order.test"
	rInt := tf.AccRandTimeInt()
	location := testLocation()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAppServiceCertificateOrder_complete(rInt, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "csr"),
					resource.TestCheckResourceAttrSet(dataSourceName, "domain_verification_token"),
					resource.TestCheckResourceAttr(dataSourceName, "distinguished_name", "CN=example.com"),
					resource.TestCheckResourceAttr(dataSourceName, "product_type", "standard"),
					resource.TestCheckResourceAttr(dataSourceName, "validatity_in_years", ""),
					resource.TestCheckResourceAttr(dataSourceName, "auto_renew", "false"),
					resource.TestCheckResourceAttr(dataSourceName, "key_size", "4096"),
				),
			},
		},
	})
}

func testAccDataSourceAppServiceCertificateOrder_basic(rInt int, location string) string {
	config := testAccAzureRMAppServiceCertificateOrder_basic(rInt, location)
	return fmt.Sprintf(`
%s

data "azurerm_app_service_certificate_order" "test" {
  name                = "${azurerm_app_service_certificate_order.test.name}"
  resource_group_name = "${azurerm_app_service_certificate_order.test.resource_group_name}"
}
`, config)
}

func testAccDataSourceAppServiceCertificateOrder_wildcard(rInt int, location string) string {
	config := testAccAzureRMAppServiceCertificateOrder_wildcard(rInt, location)
	return fmt.Sprintf(`
%s

data "azurerm_app_service_certificate_order" "test" {
  name                = "${azurerm_app_service_certificate_order.test.name}"
  resource_group_name = "${azurerm_app_service_certificate_order.test.resource_group_name}"
}
`, config)
}

func testAccDataSourceAppServiceCertificateOrder_complete(rInt int, location string) string {
	config := testAccAzureRMAppServiceCertificateOrder_complete(rInt, location)
	return fmt.Sprintf(`
%s

data "azurerm_app_service_certificate_order" "test" {
  name                = "${azurerm_app_service_certificate_order.test.name}"
  resource_group_name = "${azurerm_app_service_certificate_order.test.resource_group_name}"
}
`, config)
}
