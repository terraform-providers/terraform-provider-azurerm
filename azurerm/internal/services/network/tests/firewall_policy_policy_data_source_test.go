package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
)

func TestAccDataSourceFirewallPolicyPolicy_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_firewall_policy_policy", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFirewallPolicyPolicy_basic(data),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(data.ResourceName, "name"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "resource_group_name"),
				),
			},
		},
	})
}

func testAccDataSourceFirewallPolicyPolicy_basic(data acceptance.TestData) string {
	config := testAccAzureRMFirewallPolicyPolicy_basic(data)
	return fmt.Sprintf(`
%s

data "azurerm_firewall_policy_policy" "test" {
  name                = azurerm_firewall_policy_policy.test.name
  resource_group_name = azurerm_firewall_policy_policy.test.resource_group_name
}
`, config)
}
