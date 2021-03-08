package keyvault_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance/check"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/keyvault/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

type KeyVaultManagedHardwareSecurityModuleResource struct {
}

func TestAccKeyVaultManagedHardwareSecurityModule_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_key_vault_managed_hardware_security_module", "test")
	r := KeyVaultManagedHardwareSecurityModuleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccKeyVaultManagedHardwareSecurityModule_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_key_vault_managed_hardware_security_module", "test")
	r := KeyVaultManagedHardwareSecurityModuleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport),
	})
}

func TestAccKeyVaultManagedHardwareSecurityModule_purgeProtectionAndSoftDelete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_key_vault_managed_hardware_security_module", "test")
	r := KeyVaultManagedHardwareSecurityModuleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.purgeProtectionAndSoftDelete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccKeyVaultManagedHardwareSecurityModule_updateTags(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_key_vault_managed_hardware_security_module", "test")
	r := KeyVaultManagedHardwareSecurityModuleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateTags(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (KeyVaultManagedHardwareSecurityModuleResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	id, err := parse.HardwareSecurityModuleID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := clients.KeyVault.ManagedHsmClient.Get(ctx, id.ResourceGroup, id.ManagedHSMName)
	if err != nil {
		return nil, fmt.Errorf("retrieving %s: %+v", id, err)
	}

	return utils.Bool(resp.Properties != nil), nil
}

func (r KeyVaultManagedHardwareSecurityModuleResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_key_vault_managed_hardware_security_module" "test" {
  name                = "kvHsm%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku_name            = "Standard_B1"
  tenant_id           = data.azurerm_client_config.current.tenant_id
  admin_object_ids    = [data.azurerm_client_config.current.object_id]
}
`, r.template(data), data.RandomInteger)
}

func (r KeyVaultManagedHardwareSecurityModuleResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_key_vault_managed_hardware_security_module" "import" {
  name                = azurerm_key_vault_managed_hardware_security_module.test.name
  resource_group_name = azurerm_key_vault_managed_hardware_security_module.test.resource_group_name
  location            = azurerm_key_vault_managed_hardware_security_module.test.location
  sku_name            = azurerm_key_vault_managed_hardware_security_module.test.sku_name
  tenant_id           = azurerm_key_vault_managed_hardware_security_module.test.tenant_id
  admin_object_ids    = azurerm_key_vault_managed_hardware_security_module.test.admin_object_ids
}
`, r.basic(data))
}

func (r KeyVaultManagedHardwareSecurityModuleResource) purgeProtectionAndSoftDelete(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_key_vault_managed_hardware_security_module" "test" {
  name                       = "kvHsm%d"
  resource_group_name        = azurerm_resource_group.test.name
  location                   = azurerm_resource_group.test.location
  sku_name                   = "Standard_B1"
  soft_delete_retention_days = 7
  purge_protection_enabled   = true
  tenant_id                  = data.azurerm_client_config.current.tenant_id
  admin_object_ids           = [data.azurerm_client_config.current.object_id]
}
`, r.template(data), data.RandomInteger)
}

func (r KeyVaultManagedHardwareSecurityModuleResource) updateTags(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_key_vault_managed_hardware_security_module" "test" {
  name                = "kvHsm%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku_name            = "Standard_B1"
  tenant_id           = data.azurerm_client_config.current.tenant_id
  admin_object_ids    = [data.azurerm_client_config.current.object_id]

  tags = {
    Env = "Test"
  }
}
`, r.template(data), data.RandomInteger)
}

func (KeyVaultManagedHardwareSecurityModuleResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

data "azurerm_client_config" "current" {
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-KV-%d"
  location = "%s"
}
`, data.RandomInteger, data.Locations.Primary)
}
