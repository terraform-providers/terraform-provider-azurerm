package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func TestAccAzureRMManagementPartner_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_management_partner", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMManagementPartnerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMManagementPartner_basic(),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMManagementPartnerExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMManagementPartner_requiresImport(t *testing.T) {
	if !features.ShouldResourcesBeImported() {
		t.Skip("Skipping since resources aren't required to be imported")
		return
	}

	data := acceptance.BuildTestData(t, "azurerm_management_partner", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMManagementPartnerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMManagementPartner_basic(),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMManagementPartnerExists(data.ResourceName),
				),
			},
			{
				Config:      testAccAzureRMManagementPartner_requiresImport(),
				ExpectError: acceptance.RequiresImportError("azurerm_management_partner"),
			},
		},
	})
}

func TestAccAzureRMManagementPartner_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_management_partner", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMManagementPartnerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMManagementPartner_basic(),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMManagementPartnerExists(data.ResourceName),
				),
			},
			data.ImportStep(),
			{
				Config: testAccAzureRMManagementPartner_update(),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMManagementPartnerExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func testCheckAzureRMManagementPartnerExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Management Partner not found: %s", resourceName)
		}

		partnerId := rs.Primary.Attributes["partner_id"]

		client := acceptance.AzureProvider.Meta().(*clients.Client).ManagementPartner.PartnerClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		if resp, err := client.Get(ctx, partnerId); err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Management Partner %q does not exist", partnerId)
			}
			return fmt.Errorf("Bad: Get on ManagementPartner.PartnerClient: %+v", err)
		}

		return nil
	}
}

func testCheckAzureRMManagementPartnerDestroy(s *terraform.State) error {
	client := acceptance.AzureProvider.Meta().(*clients.Client).ManagementPartner.PartnerClient
	ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azurerm_management_partner" {
			continue
		}

		partnerId := rs.Primary.Attributes["partner_id"]

		if resp, err := client.Get(ctx, partnerId); err != nil {
			if !utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Get on ManagementPartner.PartnerClient: %+v", err)
			}
		}

		return nil
	}

	return nil
}

func testAccAzureRMManagementPartner_basic() string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_management_partner" "test" {
  partner_id = "6080810"
}
`)
}

func testAccAzureRMManagementPartner_requiresImport() string {
	return fmt.Sprintf(`
%s

resource "azurerm_management_partner" "import" {
  partner_id = "${azurerm_management_partner.test.partner_id}"
}
`, testAccAzureRMManagementPartner_basic())
}

func testAccAzureRMManagementPartner_update() string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_management_partner" "test" {
  partner_id = "6080830"
}
`)
}
