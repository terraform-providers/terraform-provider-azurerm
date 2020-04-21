package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func TestAccAzureRMRegistrationAssignment_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_registration_assignment", "test")
	// Multiple tenants are needed to test this resource.
	// Second tenant ID needs to be set as a enviornment variable ARM_SECOND_TENANT_ID.
	secondTenantID := os.Getenv("ARM_SECOND_TENANT_ID")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMRegistrationAssignmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMRegistrationAssignment_basic(uuid.New().String(), secondTenantID, data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMRegistrationAssignmentExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "registration_assignment_id"),
				),
			},
		},
	})
}

func TestAccAzureRMRegistrationAssignment_requiresImport(t *testing.T) {
	if !features.ShouldResourcesBeImported() {
		t.Skip("Skipping since resources aren't required to be imported")
		return
	}
	data := acceptance.BuildTestData(t, "azurerm_registration_assignment", "test")
	secondTenantID := os.Getenv("ARM_SECOND_TENANT_ID")
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMRegistrationAssignmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMRegistrationAssignment_basic(id, secondTenantID, data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMRegistrationAssignmentExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "registration_assignment_id"),
				),
			},
			{
				Config:      testAccAzureRMRegistrationAssignment_requiresImport(id, secondTenantID, data),
				ExpectError: acceptance.RequiresImportError("azurerm_registration_assignment"),
			},
		},
	})
}

func TestAccAzureRMRegistrationAssignment_emptyID(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_registration_assignment", "test")
	secondTenantID := os.Getenv("ARM_SECOND_TENANT_ID")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMRegistrationAssignmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMRegistrationAssignment_emptyId(secondTenantID, data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMRegistrationAssignmentExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "registration_assignment_id"),
				),
			},
		},
	})
}

func testCheckAzureRMRegistrationAssignmentExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.AzureProvider.Meta().(*clients.Client).ManagedServices.RegistrationAssignmentsClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %q", resourceName)
		}

		scope := rs.Primary.Attributes["scope"]
		RegistrationAssignmentID := rs.Primary.Attributes["registration_assignment_id"]
		expandRegistrationDefinition := true

		resp, err := client.Get(ctx, scope, RegistrationAssignmentID, &expandRegistrationDefinition)

		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Registration Assignment %q (Scope: %q) does not exist", RegistrationAssignmentID, scope)
			}
			return fmt.Errorf("Bad: Get on RegistrationAssignmentsClient: %+v", err)
		}

		return nil
	}
}

func testCheckAzureRMRegistrationAssignmentDestroy(s *terraform.State) error {
	client := acceptance.AzureProvider.Meta().(*clients.Client).ManagedServices.RegistrationAssignmentsClient
	ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azurerm_registration_assignment" {
			continue
		}

		scope := rs.Primary.Attributes["scope"]
		RegistrationAssignmentID := rs.Primary.Attributes["registration_assignment_id"]
		expandRegistrationDefinition := true

		resp, err := client.Get(ctx, scope, RegistrationAssignmentID, &expandRegistrationDefinition)

		if err != nil {
			if !utils.ResponseWasNotFound(resp.Response) {
				return err
			}
		}

		return nil
	}

	return nil
}

func testAccAzureRMRegistrationAssignment_basic(id string, secondTenantID string, data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

data "azurerm_subscription" "primary" {
}

data "azurerm_client_config" "test" {
}

resource "azurerm_registration_definition" "test" {
  name                       = "acctestrd-%d"
  description				 = "Acceptance Test Registration Definition"
  scope                      = data.azurerm_subscription.primary.id
  managed_by_tenant_id       = "%s"

  authorization {
	principal_id        = data.azurerm_client_config.test.object_id
	role_definition_id  = "b24988ac-6180-42a0-ab88-20f7382dd24c"
  }
}

resource "azurerm_registration_assignment" "test" {
   registration_assignment_id = "%s"
   scope = data.azurerm_subscription.primary.id
   registration_definition_id = azurerm_registration_definition.test.id
}

`, data.RandomInteger, secondTenantID, id)
}

func testAccAzureRMRegistrationAssignment_requiresImport(id string, secondTenantID string, data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_registration_assignment" "import" {
  registration_assignment_id = azurerm_registration_assignment.test.registration_assignment_id
  registration_definition_id = azurerm_registration_assignment.test.registration_definition_id
  scope                      = azurerm_registration_assignment.test.scope
}
`, testAccAzureRMRegistrationAssignment_basic(id, secondTenantID, data))
}

func testAccAzureRMRegistrationAssignment_emptyId(secondTenantID string, data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
	features {}
}
	  
data "azurerm_subscription" "primary" {
}

data "azurerm_client_config" "test" {
}

resource "azurerm_registration_definition" "test" {
  name                       = "acctestrd-%d"
  description				 = "Acceptance Test Registration Definition"
  scope                      = data.azurerm_subscription.primary.id
  managed_by_tenant_id       = "%s"

  authorization {
	principal_id        = data.azurerm_client_config.test.object_id
	role_definition_id  = "b24988ac-6180-42a0-ab88-20f7382dd24c"
  }
}

resource "azurerm_registration_assignment" "test" {
	scope = data.azurerm_subscription.primary.id
	registration_definition_id = azurerm_registration_definition.test.id
 }
`, data.RandomInteger, secondTenantID)
}
