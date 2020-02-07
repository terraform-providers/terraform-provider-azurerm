package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
)

func TestAccDataSourceAzureRMMachineLearningWorkspace_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_machine_learning_workspace", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMachineLearningWorkspace_basic(data),
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

func testAccDataSourceMachineLearningWorkspace_basic(data acceptance.TestData) string {
	config := testAccAzureRMMachineLearningWorkspace_basic(data)
	return fmt.Sprintf(`
%s

data "azurerm_machine_learning_workspace" "test" {
  resource_group_name = "${azurerm_machine_learning_workspace.test.resource_group_name}"
  name                = "${azurerm_machine_learning_workspace.test.name}"
}
`, config)
}
