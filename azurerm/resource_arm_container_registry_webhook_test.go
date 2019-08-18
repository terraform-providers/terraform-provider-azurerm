package azurerm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func TestAccAzureRMContainerRegistryWebhook_basic(t *testing.T) {
	resourceName := "azurerm_container_registry_webhook.test"
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMContainerRegistryWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMContainerRegistryWebhook_basic(ri, testLocation()),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMContainerRegistryWebhookExists(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAzureRMContainerRegistryWebhook_basic(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "rg" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_container_registry" "acr" {
  name                     = "acrwebhooktest%d"
  resource_group_name      = azurerm_resource_group.rg.name
  location                 = "%s"
  sku                      = "Standard"
}

resource "azurerm_container_registry_webhook" "test" {
  name                = "testwebhook%d"
  resource_group_name = azurerm_resource_group.rg.name
  registry_name       = azurerm_container_registry.acr.name
  location            = "%s"
  
  service_uri    = "https://mywebhookreceiver.example/mytag"

  actions = [
      "push"
   ]
}
`, rInt, location, rInt, location, rInt, location)
}

func testCheckAzureRMContainerRegistryWebhookDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*ArmClient).containers.WebhooksClient

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azurerm_container_registry_webhook" {
			continue
		}

		resourceGroup := rs.Primary.Attributes["resource_group_name"]
		registryName := rs.Primary.Attributes["registry_name"]
		name := rs.Primary.Attributes["name"]

		ctx := testAccProvider.Meta().(*ArmClient).StopContext
		resp, err := client.Get(ctx, resourceGroup, registryName, name)

		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return nil
			}
			return err
		}

		return nil
	}

	return nil
}

func testCheckAzureRMContainerRegistryWebhookExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		webhookName := rs.Primary.Attributes["name"]
		resourceGroup, hasResourceGroup := rs.Primary.Attributes["resource_group_name"]
		if !hasResourceGroup {
			return fmt.Errorf("Bad: no resource group found in state for Container Registry Webhook: %s", webhookName)
		}

		registryName, hasRegistryName := rs.Primary.Attributes["registry_name"]
		if !hasRegistryName {
			return fmt.Errorf("Bad: no registry name found in state for Container Registry Webhook: %s", webhookName)
		}

		client := testAccProvider.Meta().(*ArmClient).containers.WebhooksClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext
		resp, err := client.Get(ctx, resourceGroup, registryName, webhookName)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Container Registry Webhook %q (resource group: %q) does not exist", webhookName, resourceGroup)
			}

			return fmt.Errorf("Bad: Get on WebhooksClient: %+v", err)
		}

		return nil
	}
}
