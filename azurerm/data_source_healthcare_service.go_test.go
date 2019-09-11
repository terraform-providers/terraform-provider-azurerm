package azurerm

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"testing"
)

func TestAccAzureRMDataSourceHealthcareService(t *testing.T) {
	dataSourceName := "data.azurerm_healthcare_service.test"
	ri := tf.AccRandTimeInt()
	location := testLocation()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMHealthcareServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMDataSourceHealthcareService_basic(ri, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "resource_group_name"),
					resource.TestCheckResourceAttr(dataSourceName, "tags.%", "2"),
				),
			},
		},
	})
}

func testAccAzureRMDataSourceHealthcareService_basic(rInt int, location string) string {
	resource := testAccAzureRMHealthcareService_basic(rInt)
	s := fmt.Sprintf(`
%s

data "azurerm_healthcare_service" "test" {
  name                = "${azurerm_healthcare_service.test.name}"
  resource_group_name = "${azurerm_healthcare_service.test.resource_group_name}"
  location            = "${azurerm_resource_group.test.location}"

  access_policy_object_ids {
    object_id          = "${data.azurerm_client_config.current.service_principal_object_id}"
  }
}
`, resource)
	return s
}
