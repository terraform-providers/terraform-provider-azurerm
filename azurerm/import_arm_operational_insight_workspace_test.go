package azurerm

import (
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccAzureRMOperationalInsightWorkspace_importrequiredOnly(t *testing.T) {
	resourceName := "azurerm_operational_insight_workspace.test"

	ri := acctest.RandInt()
	config := testAccAzureRMOperationalInsightWorkspace_requiredOnly(ri, testLocation())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMOperationalInsightWorkspaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureRMOperationalInsightWorkspace_importoptional(t *testing.T) {
	resourceName := "azurerm_operational_insight_workspace.test"

	ri := acctest.RandInt()
	config := testAccAzureRMOperationalInsightWorkspace_optional(ri, testLocation())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMOperationalInsightWorkspaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
