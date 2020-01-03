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

func TestAccAzureRMManagedApplicationDefinition_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_managed_application_definition", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMManagedApplicationDefinitionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMManagedApplicationDefinition_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMManagedApplicationDefinitionExists(data.ResourceName),
				),
			},
			data.ImportStep("package_file_uri"),
		},
	})
}

func TestAccAzureRMManagedApplicationDefinition_requiresImport(t *testing.T) {
	if !features.ShouldResourcesBeImported() {
		t.Skip("Skipping since resources aren't required to be imported")
		return
	}

	data := acceptance.BuildTestData(t, "azurerm_managed_application_definition", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMManagedApplicationDefinitionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMManagedApplicationDefinition_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMManagedApplicationDefinitionExists(data.ResourceName),
				),
			},
			{
				Config:      testAccAzureRMManagedApplicationDefinition_requiresImport(data),
				ExpectError: acceptance.RequiresImportError("azurerm_managed_application_definition"),
			},
		},
	})
}

func TestAccAzureRMManagedApplicationDefinition_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_managed_application_definition", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMManagedApplicationDefinitionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMManagedApplicationDefinition_complete(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMManagedApplicationDefinitionExists(data.ResourceName),
				),
			},
			data.ImportStep("create_ui_definition", "main_template"),
		},
	})
}

func TestAccAzureRMManagedApplicationDefinition_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_managed_application_definition", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMManagedApplicationDefinitionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMManagedApplicationDefinition_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMManagedApplicationDefinitionExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "display_name", "TestManagedApplicationDefinition"),
					resource.TestCheckResourceAttr(data.ResourceName, "description", "Test Managed Application Definition"),
					resource.TestCheckResourceAttr(data.ResourceName, "tags.%", "0"),
				),
			},
			data.ImportStep("package_file_uri"),
			{
				Config: testAccAzureRMManagedApplicationDefinition_complete(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMManagedApplicationDefinitionExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "display_name", "UpdatedTestManagedApplicationDefinition"),
					resource.TestCheckResourceAttr(data.ResourceName, "description", "Updated Test Managed Application Definition"),
					resource.TestCheckResourceAttr(data.ResourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "tags.ENV", "Prod"),
				),
			},
			data.ImportStep("create_ui_definition", "main_template", "package_file_uri"),
		},
	})
}

func testCheckAzureRMManagedApplicationDefinitionExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Managed Application Definition not found: %s", resourceName)
		}

		name := rs.Primary.Attributes["name"]
		resourceGroup := rs.Primary.Attributes["resource_group_name"]

		client := acceptance.AzureProvider.Meta().(*clients.Client).ManagedApplication.ApplicationDefinitionClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		if resp, err := client.Get(ctx, resourceGroup, name); err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Managed Application Definition %q (Resource Group %q) does not exist", name, resourceGroup)
			}
			return fmt.Errorf("Bad: Get on ManagedApplication.ApplicationDefinitionClient: %+v", err)
		}

		return nil
	}
}

func testCheckAzureRMManagedApplicationDefinitionDestroy(s *terraform.State) error {
	client := acceptance.AzureProvider.Meta().(*clients.Client).ManagedApplication.ApplicationDefinitionClient
	ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azurerm_managed_application_definition" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		resourceGroup := rs.Primary.Attributes["resource_group_name"]

		if resp, err := client.Get(ctx, resourceGroup, name); err != nil {
			if !utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Get on ManagedApplication.ApplicationDefinitionClient: %+v", err)
			}
		}

		return nil
	}

	return nil
}

func testAccAzureRMManagedApplicationDefinition_basic(data acceptance.TestData) string {
	template := testAccAzureRMManagedApplicationDefinition_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_managed_application_definition" "test" {
  name                = "acctestAppDef%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  lock_level           = "ReadOnly"
  package_file_uri     = "https://github.com/Azure/azure-managedapp-samples/raw/master/Managed Application Sample Packages/201-managed-storage-account/managedstorage.zip"
  display_name         = "TestManagedApplicationDefinition"
  description          = "Test Managed Application Definition"

  authorization {
    service_principal_id = "${data.azurerm_client_config.current.object_id}"
    role_definition_id   = "b24988ac-6180-42a0-ab88-20f7382dd24c"
  }
}
`, template, data.RandomInteger)
}

func testAccAzureRMManagedApplicationDefinition_requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_managed_application_definition" "import" {
  name                = "${azurerm_managed_application_definition.test.name}"
  location            = "${azurerm_managed_application_definition.test.location}"
  resource_group_name = "${azurerm_managed_application_definition.test.name}"
}
`, testAccAzureRMManagedApplicationDefinition_basic(data))
}

func testAccAzureRMManagedApplicationDefinition_complete(data acceptance.TestData) string {
	template := testAccAzureRMManagedApplicationDefinition_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_managed_application_definition" "test" {
  name                = "acctestAppDef%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  lock_level           = "ReadOnly"
  display_name         = "UpdatedTestManagedApplicationDefinition"
  description          = "Updated Test Managed Application Definition"
  create_ui_definition = <<CREATE_UI_DEFINITION
    {
      "$schema": "https://schema.management.azure.com/schemas/0.1.2-preview/CreateUIDefinition.MultiVm.json#",
      "handler": "Microsoft.Azure.CreateUIDef",
	  "version": "0.1.2-preview",
	  "parameters": {
         "basics": [
            {}
         ],
         "steps": [
            {
              "name": "storageConfig",
              "label": "Storage settings",
              "subLabel": {
                 "preValidation": "Configure the infrastructure settings",
                 "postValidation": "Done"
              },
              "bladeTitle": "Storage settings",
              "elements": [
                 {
                   "name": "storageAccounts",
                   "type": "Microsoft.Storage.MultiStorageAccountCombo",
                   "label": {
                      "prefix": "Storage account name prefix",
                      "type": "Storage account type"
                   },
                   "defaultValue": {
                      "type": "Standard_LRS"
                   },
                   "constraints": {
                      "allowedTypes": [
                        "Premium_LRS",
                        "Standard_LRS",
                        "Standard_GRS"
                      ]
                   }
                 }
              ]
            }
         ],
         "outputs": {
            "storageAccountNamePrefix": "[steps('storageConfig').storageAccounts.prefix]",
            "storageAccountType": "[steps('storageConfig').storageAccounts.type]",
            "location": "[location()]"
         }
      }
	}
  CREATE_UI_DEFINITION

  main_template = <<MAIN_TEMPLATE
    {
      "$schema": "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
      "contentVersion": "1.0.0.0",
      "parameters": {
         "storageAccountNamePrefix": {
            "type": "string"
         },
         "storageAccountType": {
            "type": "string"
         },
         "location": {
            "type": "string",
            "defaultValue": "[resourceGroup().location]"
         }
      },
      "variables": {
         "storageAccountName": "[concat(parameters('storageAccountNamePrefix'), uniqueString(resourceGroup().id))]"
      },
      "resources": [
         {
           "type": "Microsoft.Storage/storageAccounts",
           "name": "[variables('storageAccountName')]",
           "apiVersion": "2016-01-01",
           "location": "[parameters('location')]",
           "sku": {
              "name": "[parameters('storageAccountType')]"
           },
           "kind": "Storage",
           "properties": {}
         }
      ],
      "outputs": {
         "storageEndpoint": {
           "type": "string",
           "value": "[reference(resourceId('Microsoft.Storage/storageAccounts/', variables('storageAccountName')), '2016-01-01').primaryEndpoints.blob]"
         }
      }
    }
  MAIN_TEMPLATE

  authorization {
    service_principal_id = "${data.azurerm_client_config.current.object_id}"
    role_definition_id   = "8e3af657-a8ff-443c-a75c-2fe8c4bcb635"
  }

  tags = {
    ENV = "Prod"
  }
}
`, template, data.RandomInteger)
}

func testAccAzureRMManagedApplicationDefinition_template(data acceptance.TestData) string {
	return fmt.Sprintf(`
data "azurerm_client_config" "current" {}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-managedapplication-%d"
  location = "%s"
}
`, data.RandomInteger, data.Locations.Primary)
}
