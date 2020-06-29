package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccAzureRMKustoClusterCustomerManagedKey_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_kusto_cluster_customer_managed_key", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMKustoClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMKustoClusterCustomerManagedKey_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMKustoClusterCustomerManagedKeyExists(data.ResourceName),
				),
			},
			data.ImportStep(),
			{
				// Delete the encryption settings resource and verify it is gone
				Config: testAccAzureRMKustoClusterCustomerManagedKey_template(data),
				Check: resource.ComposeTestCheckFunc(
					// Then ensure the encryption settings on the storage account
					// have been reverted to their default state
					testCheckAzureRMKustoClusterExistsWithDefaultSettings("azurerm_kusto_cluster.test"),
				),
			},
		},
	})
}

func TestAccAzureRMKustoClusterCustomerManagedKey_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_kusto_cluster_customer_managed_key", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMKustoClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMKustoClusterCustomerManagedKey_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMKustoClusterCustomerManagedKeyExists(data.ResourceName),
				),
			},
			data.RequiresImportErrorStep(testAccAzureRMKustoClusterCustomerManagedKey_requiresImport),
		},
	})
}

func TestAccAzureRMKustoClusterCustomerManagedKey_updateKey(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_kusto_cluster_customer_managed_key", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMKustoClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMKustoClusterCustomerManagedKey_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMKustoClusterCustomerManagedKeyExists(data.ResourceName),
				),
			},
			data.ImportStep(),
			{
				Config: testAccAzureRMKustoClusterCustomerManagedKey_updated(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMKustoClusterCustomerManagedKeyExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func testCheckAzureRMKustoClusterExistsWithDefaultSettings(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		kustoCluster := rs.Primary.Attributes["name"]
		resourceGroup := rs.Primary.Attributes["resource_group_name"]

		// Ensure resource group exists in API
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext
		conn := acceptance.AzureProvider.Meta().(*clients.Client).Kusto.ClustersClient

		resp, err := conn.Get(ctx, resourceGroup, kustoCluster)
		if err != nil {
			return fmt.Errorf("Bad: Get on Kusto ClustersClient: %+v", err)
		}

		if resp.StatusCode == http.StatusNotFound {
			return fmt.Errorf("Bad: Kusto Cluster %q (resource group: %q) does not exist", kustoCluster, resourceGroup)
		}

		if props := resp.ClusterProperties; props != nil {
			return fmt.Errorf("Kusto Cluster encryption properties not found: %s", resourceName)
		}

		return nil
	}
}

func testCheckAzureRMKustoClusterCustomerManagedKeyExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if kustoClusterId := rs.Primary.Attributes["storage_account_id"]; kustoClusterId == "" {
			return fmt.Errorf("Unable to read kustoClusterId: %s", resourceName)
		}

		return nil
	}
}

func testAccAzureRMKustoClusterCustomerManagedKey_basic(data acceptance.TestData) string {
	template := testAccAzureRMKustoClusterCustomerManagedKey_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_kusto_cluster_customer_managed_key" "test" {
  cluster_id   = azurerm_kusto_cluster.test.id
  key_vault_id = azurerm_key_vault.test.id
  key_name     = azurerm_key_vault_key.first.name
  key_version  = azurerm_key_vault_key.first.version
}
`, template)
}

func testAccAzureRMKustoClusterCustomerManagedKey_requiresImport(data acceptance.TestData) string {
	template := testAccAzureRMKustoClusterCustomerManagedKey_basic(data)
	return fmt.Sprintf(`
%s

resource "azurerm_kusto_cluster_customer_managed_key" "import" {
  cluster_id   = azurerm_storage_account_customer_managed_key.test.cluster_id
  key_vault_id = azurerm_storage_account_customer_managed_key.test.key_vault_id
  key_name     = azurerm_storage_account_customer_managed_key.test.key_name
  key_version  = azurerm_storage_account_customer_managed_key.test.key_version
}
`, template)
}

func testAccAzureRMKustoClusterCustomerManagedKey_updated(data acceptance.TestData) string {
	template := testAccAzureRMKustoClusterCustomerManagedKey_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_key_vault_key" "second" {
  name         = "second"
  key_vault_id = azurerm_key_vault.test.id
  key_type     = "RSA"
  key_size     = 2048
  key_opts     = ["decrypt", "encrypt", "sign", "unwrapKey", "verify", "wrapKey"]

  depends_on = [
    azurerm_key_vault_access_policy.client,
    azurerm_key_vault_access_policy.cluster,
  ]
}

resource "azurerm_kusto_cluster_customer_managed_key" "test" {
  cluster_id   = azurerm_kusto_cluster.test.id
  key_vault_id = azurerm_key_vault.test.id
  key_name     = azurerm_key_vault_key.second.name
  key_version  = azurerm_key_vault_key.second.version
}
`, template)
}

func testAccAzureRMKustoClusterCustomerManagedKey_template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {
    key_vault {
      purge_soft_delete_on_destroy = false
    }
  }
}

data "azurerm_client_config" "current" {}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_key_vault" "test" {
  name                     = "acctestkv%s"
  location                 = azurerm_resource_group.test.location
  resource_group_name      = azurerm_resource_group.test.name
  tenant_id                = data.azurerm_client_config.current.tenant_id
  sku_name                 = "standard"
  soft_delete_enabled      = true
  purge_protection_enabled = true
}

resource "azurerm_key_vault_access_policy" "cluster" {
  key_vault_id = azurerm_key_vault.test.id
  tenant_id    = azurerm_kusto_cluster.test.identity.0.tenant_id
  object_id    = azurerm_kusto_cluster.test.identity.0.principal_id

  key_permissions    = ["get", "create", "list", "restore", "recover", "unwrapkey", "wrapkey", "purge", "encrypt", "decrypt", "sign", "verify"]
  secret_permissions = ["get"]
}

resource "azurerm_key_vault_access_policy" "client" {
  key_vault_id = azurerm_key_vault.test.id
  tenant_id    = data.azurerm_client_config.current.tenant_id
  object_id    = data.azurerm_client_config.current.object_id

  key_permissions    = ["get", "create", "delete", "list", "restore", "recover", "unwrapkey", "wrapkey", "purge", "encrypt", "decrypt", "sign", "verify"]
  secret_permissions = ["get"]
}

resource "azurerm_key_vault_key" "test" {
  name         = "test"
  key_vault_id = azurerm_key_vault.test.id
  key_type     = "RSA"
  key_size     = 2048
  key_opts     = ["decrypt", "encrypt", "sign", "unwrapKey", "verify", "wrapKey"]

  depends_on = [
    azurerm_key_vault_access_policy.client,
    azurerm_key_vault_access_policy.cluster,
  ]
}

resource "azurerm_kusto_cluster" "test" {
  name                = "acctestkc%s"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name

  sku {
    name     = "Dev(No SLA)_Standard_D11_v2"
    capacity = 1
  }

  identity {
    type = "SystemAssigned"
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomString, data.RandomString)
}
