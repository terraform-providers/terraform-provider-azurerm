package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/mssql/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

// TODO: add import tests
func TestAccAzureRMMsSqlDatabase_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_mssql_database", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMMsSqlDatabaseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMMsSqlDatabase_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMMsSqlDatabaseExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMMsSqlDatabase_requiresImport(t *testing.T) {
	if !features.ShouldResourcesBeImported() {
		t.Skip("Skipping since resources aren't required to be imported")
		return
	}

	data := acceptance.BuildTestData(t, "azurerm_mssql_database", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMMsSqlDatabaseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMMsSqlDatabase_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMMsSqlDatabaseExists(data.ResourceName),
				),
			},
			data.RequiresImportErrorStep(testAccAzureRMMsSqlDatabase_requiresImport),
		},
	})
}

func TestAccAzureRMMsSqlDatabase_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_mssql_database", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMMsSqlDatabaseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMMsSqlDatabase_complete(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMMsSqlDatabaseExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "collation", "SQL_AltDiction_CP850_CI_AI"),
					resource.TestCheckResourceAttr(data.ResourceName, "license_type", "BasePrice"),
				),
			},
			data.ImportStep("sample_name"),
			{
				Config: testAccAzureRMMsSqlDatabase_update(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMMsSqlDatabaseExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "license_type", "LicenseIncluded"),
				),
			},
			data.ImportStep("sample_name"),
		},
	})
}

func TestAccAzureRMMsSqlDatabase_elasticPool(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_mssql_database", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMMsSqlDatabaseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMMsSqlDatabase_elasticPool(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMMsSqlDatabaseExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "elastic_pool_id"),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMMsSqlDatabase_GP(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_mssql_database", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMMsSqlDatabaseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMMsSqlDatabase_GP(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMMsSqlDatabaseExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "general_purpose.#", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "general_purpose.0.capacity", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "general_purpose.0.family", "Gen4"),
					resource.TestCheckResourceAttr(data.ResourceName, "general_purpose.0.max_size_gb", "2"),
				),
			},
			data.ImportStep(),
			{
				Config: testAccAzureRMMsSqlDatabase_GPUpdated(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMMsSqlDatabaseExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "general_purpose.#", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "general_purpose.0.capacity", "4"),
					resource.TestCheckResourceAttr(data.ResourceName, "general_purpose.0.family", "Gen5"),
					resource.TestCheckResourceAttr(data.ResourceName, "general_purpose.0.max_size_gb", "5"),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMMsSqlDatabase_BC(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_mssql_database", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMMsSqlDatabaseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMMsSqlDatabase_BC(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMMsSqlDatabaseExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "business_critical.#", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "business_critical.0.capacity", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "business_critical.0.family", "Gen4"),
					resource.TestCheckResourceAttr(data.ResourceName, "business_critical.0.max_size_gb", "1"),
				),
			},
			data.ImportStep(),
			{
				Config: testAccAzureRMMsSqlDatabase_BCUpdated(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMMsSqlDatabaseExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "business_critical.#", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "business_critical.0.capacity", "4"),
					resource.TestCheckResourceAttr(data.ResourceName, "business_critical.0.family", "Gen5"),
					resource.TestCheckResourceAttr(data.ResourceName, "business_critical.0.max_size_gb", "3"),
					resource.TestCheckResourceAttr(data.ResourceName, "business_critical.0.read_scale", "Enabled"),
					resource.TestCheckResourceAttr(data.ResourceName, "business_critical.0.zone_redundant", "true"),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMMsSqlDatabase_HS(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_mssql_database", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMMsSqlDatabaseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMMsSqlDatabase_HS(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMMsSqlDatabaseExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "hyper_scale.#", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "hyper_scale.0.capacity", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "hyper_scale.0.family", "Gen4"),
					resource.TestCheckResourceAttr(data.ResourceName, "hyper_scale.0.read_replica_count", "1"),
				),
			},
			data.ImportStep(),
			{
				Config: testAccAzureRMMsSqlDatabase_HSUpdated(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMMsSqlDatabaseExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "hyper_scale.#", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "hyper_scale.0.capacity", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "hyper_scale.0.family", "Gen5"),
					resource.TestCheckResourceAttr(data.ResourceName, "hyper_scale.0.read_replica_count", "2"),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccAzureRMMsSqlDatabase_createCopyMode(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_mssql_database", "copy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMMsSqlDatabaseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMMsSqlDatabase_createCopyMode(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMMsSqlDatabaseExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "collation", "SQL_AltDiction_CP850_CI_AI"),
					resource.TestCheckResourceAttr(data.ResourceName, "license_type", "BasePrice"),
				),
			},
			data.ImportStep("create_copy_mode.#", "create_copy_mode.0.source_database_id"),
		},
	})
}

func TestAccAzureRMMsSqlDatabase_createPITRMode(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_mssql_database", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMMsSqlDatabaseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMMsSqlDatabase_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMMsSqlDatabaseExists(data.ResourceName),
				),
			},
			data.ImportStep(),

			{
				PreConfig: func() { time.Sleep(7 * time.Minute) },
				Config:    testAccAzureRMMsSqlDatabase_createPITRMode(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMMsSqlDatabaseExists("azurerm_mssql_database.pitr"),
				),
			},

			data.ImportStep("create_pitr_mode.#", "create_pitr_mode.0.source_database_id", "create_pitr_mode.0.restore_point_in_time"),
		},
	})
}

//func TestAccAzureRMMsSqlDatabase_createRecoveryMode(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azurerm_mssql_database", "recovery")
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:     func() { acceptance.PreCheck(t) },
//		Providers:    acceptance.SupportedProviders,
//		CheckDestroy: testCheckAzureRMMsSqlDatabaseDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccAzureRMMsSqlDatabase_createRecoveryMode(data),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckAzureRMMsSqlDatabaseExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "collation", "SQL_AltDiction_CP850_CI_AI"),
//					resource.TestCheckResourceAttr(data.ResourceName, "license_type", "BasePrice"),
//				),
//			},
//			data.ImportStep(),
//		},
//	})
//}
//
//func TestAccAzureRMMsSqlDatabase_createRestoreMode(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azurerm_mssql_database", "test")
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:     func() { acceptance.PreCheck(t) },
//		Providers:    acceptance.SupportedProviders,
//		CheckDestroy: testCheckAzureRMMsSqlDatabaseDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccAzureRMMsSqlDatabase_createRestoreMode(data),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckAzureRMMsSqlDatabaseExists(data.ResourceName),
//					resource.TestCheckResourceAttrSet(data.ResourceName, "create_restore_mode.0.source_database_id"),
//					resource.TestCheckResourceAttr(data.ResourceName, "create_restore_mode.0.source_database_deletion_date", "2020-07-14T06:41:06.613Z"),
//				),
//			},
//			data.ImportStep(),
//		},
//	})
//}

func TestAccAzureRMMsSqlDatabase_createSecondaryMode(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_mssql_database", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMMsSqlDatabaseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMMsSqlDatabase_createSecondaryMode(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMMsSqlDatabaseExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "collation", "SQL_AltDiction_CP850_CI_AI"),
					resource.TestCheckResourceAttr(data.ResourceName, "license_type", "BasePrice"),
				),
			},
			data.ImportStep("create_secondary_mode.#", "create_secondary_mode.0.source_database_id","sample_name"),
		},
	})
}

func testCheckAzureRMMsSqlDatabaseExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.AzureProvider.Meta().(*clients.Client).MSSQL.DatabasesClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		id, err := parse.MsSqlDatabaseID(rs.Primary.ID)
		if err != nil {
			return err
		}

		resp, err := client.Get(ctx, id.ResourceGroup, id.MsSqlServer, id.Name)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: MsSql Database %q (resource group: %q) does not exist", id.Name, id.ResourceGroup)
			}

			return fmt.Errorf("Bad: Get on MsSql Database Client: %+v", err)
		}

		return nil
	}
}

func testCheckAzureRMMsSqlDatabaseDestroy(s *terraform.State) error {
	client := acceptance.AzureProvider.Meta().(*clients.Client).MSSQL.DatabasesClient
	ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azurerm_mssql_database" {
			continue
		}

		id, err := parse.MsSqlDatabaseID(rs.Primary.ID)
		if err != nil {
			return err
		}

		if resp, err := client.Get(ctx, id.ResourceGroup, id.MsSqlServer, id.Name); err != nil {
			if !utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Get on MsSql Database Client: %+v", err)
			}
		}
		return nil
	}

	return nil
}

func testAccAzureRMMsSqlDatabase_template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-mssqldb-%[1]d"
  location = "%[2]s"
}

resource "azurerm_sql_server" "test" {
  name                         = "acctest-sqlserver-%[1]d"
  resource_group_name          = azurerm_resource_group.test.name
  location                     = azurerm_resource_group.test.location
  version                      = "12.0"
  administrator_login          = "mradministrator"
  administrator_login_password = "thisIsDog11"
}
`, data.RandomInteger, data.Locations.Primary)
}

func testAccAzureRMMsSqlDatabase_basic(data acceptance.TestData) string {
	template := testAccAzureRMMsSqlDatabase_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_mssql_database" "test" {
  name            = "acctest-db-%d"
  mssql_server_id = azurerm_sql_server.test.id
}
`, template, data.RandomInteger)
}

func testAccAzureRMMsSqlDatabase_requiresImport(data acceptance.TestData) string {
	template := testAccAzureRMMsSqlDatabase_basic(data)
	return fmt.Sprintf(`
%s

resource "azurerm_mssql_database" "import" {
  name            = azurerm_mssql_database.test.name
  mssql_server_id = azurerm_sql_server.test.id
}
`, template)
}

func testAccAzureRMMsSqlDatabase_complete(data acceptance.TestData) string {
	template := testAccAzureRMMsSqlDatabase_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_mssql_database" "test" {
  name            = "acctest-db-%[2]d"
  mssql_server_id = azurerm_sql_server.test.id
  collation       = "SQL_AltDiction_CP850_CI_AI"
  license_type    = "BasePrice"
  sample_name     = "AdventureWorksLT"
}
`, template, data.RandomInteger)
}

func testAccAzureRMMsSqlDatabase_update(data acceptance.TestData) string {
	template := testAccAzureRMMsSqlDatabase_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_mssql_database" "test" {
  name            = "acctest-db-%[2]d"
  mssql_server_id = azurerm_sql_server.test.id
  collation       = "SQL_AltDiction_CP850_CI_AI"
  license_type    = "LicenseIncluded"
}
`, template, data.RandomInteger)
}

func testAccAzureRMMsSqlDatabase_elasticPool(data acceptance.TestData) string {
	template := testAccAzureRMMsSqlDatabase_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_mssql_elasticpool" "test" {
  name                = "acctest-pool-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  server_name         = azurerm_sql_server.test.name
  max_size_gb         = 4.8828125
  zone_redundant      = false

  sku {
    name     = "BasicPool"
    tier     = "Basic"
    capacity = 50
  }

  per_database_settings {
    min_capacity = 0
    max_capacity = 5
  }
}

resource "azurerm_mssql_database" "test" {
  name            = "acctest-db-%[2]d"
  mssql_server_id = azurerm_sql_server.test.id
  elastic_pool_id = azurerm_mssql_elasticpool.test.id
}
`, template, data.RandomInteger)
}

func testAccAzureRMMsSqlDatabase_GP(data acceptance.TestData) string {
	template := testAccAzureRMMsSqlDatabase_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_mssql_database" "test" {
  name            = "acctest-db-%d"
  mssql_server_id = azurerm_sql_server.test.id
  general_purpose {
    capacity    = 2
    family      = "Gen4"
    max_size_gb = 2
  }
}
`, template, data.RandomInteger)
}

func testAccAzureRMMsSqlDatabase_GPUpdated(data acceptance.TestData) string {
	template := testAccAzureRMMsSqlDatabase_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_mssql_database" "test" {
  name            = "acctest-db-%d"
  mssql_server_id = azurerm_sql_server.test.id
  general_purpose {
    capacity    = 4
    family      = "Gen5"
    max_size_gb = 5
  }
}
`, template, data.RandomInteger)
}

func testAccAzureRMMsSqlDatabase_HS(data acceptance.TestData) string {
	template := testAccAzureRMMsSqlDatabase_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_mssql_database" "test" {
  name            = "acctest-db-%d"
  mssql_server_id = azurerm_sql_server.test.id
  hyper_scale {
    capacity           = 1
    family             = "Gen4"
    read_replica_count = 1
  }
}
`, template, data.RandomInteger)
}

func testAccAzureRMMsSqlDatabase_HSUpdated(data acceptance.TestData) string {
	template := testAccAzureRMMsSqlDatabase_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_mssql_database" "test" {
  name            = "acctest-db-%d"
  mssql_server_id = azurerm_sql_server.test.id
  hyper_scale {
    capacity           = 2
    family             = "Gen5"
    read_replica_count = 2
  }
}
`, template, data.RandomInteger)
}

func testAccAzureRMMsSqlDatabase_BC(data acceptance.TestData) string {
	template := testAccAzureRMMsSqlDatabase_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_mssql_database" "test" {
  name            = "acctest-db-%d"
  mssql_server_id = azurerm_sql_server.test.id
  business_critical {
    capacity    = 2
    family      = "Gen4"
    max_size_gb = 1
  }
}
`, template, data.RandomInteger)
}

func testAccAzureRMMsSqlDatabase_BCUpdated(data acceptance.TestData) string {
	template := testAccAzureRMMsSqlDatabase_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_mssql_database" "test" {
  name            = "acctest-db-%d"
  mssql_server_id = azurerm_sql_server.test.id
  business_critical {
    capacity       = 4
    family         = "Gen5"
    max_size_gb    = 3
    read_scale     = "Enabled"
    zone_redundant = true
  }
}
`, template, data.RandomInteger)
}

func testAccAzureRMMsSqlDatabase_createCopyMode(data acceptance.TestData) string {
	template := testAccAzureRMMsSqlDatabase_complete(data)
	return fmt.Sprintf(`
%s

resource "azurerm_mssql_database" "copy" {
  name            = "acctest-dbc-%d"
  mssql_server_id = azurerm_sql_server.test.id
  create_copy_mode {
    source_database_id = azurerm_mssql_database.test.id
  }

}
`, template, data.RandomInteger)
}

func testAccAzureRMMsSqlDatabase_createPITRMode(data acceptance.TestData) string {
	template := testAccAzureRMMsSqlDatabase_basic(data)
	return fmt.Sprintf(`
%s

resource "azurerm_mssql_database" "pitr" {
  name            = "acctest-dbp-%d"
  mssql_server_id = azurerm_sql_server.test.id
  create_pitr_mode {
    source_database_id    = azurerm_mssql_database.test.id
    restore_point_in_time = "%s"
  }
}
`, template, data.RandomInteger, time.Now().Add(time.Duration(7)*time.Minute).UTC().Format(time.RFC3339))
}

func testAccAzureRMMsSqlDatabase_createSecondaryMode(data acceptance.TestData) string {
	template := testAccAzureRMMsSqlDatabase_complete(data)
	return fmt.Sprintf(`
%s

resource "azurerm_resource_group" "second" {
  name     = "acctestRG-mssqldb2-%[2]d"
  location = "%[3]s"
}

resource "azurerm_sql_server" "second" {
  name                         = "acctest-sqlserver2-%[2]d"
  resource_group_name          = azurerm_resource_group.second.name
  location                     = azurerm_resource_group.second.location
  version                      = "12.0"
  administrator_login          = "mradministrator"
  administrator_login_password = "thisIsDog11"
}

resource "azurerm_mssql_database" "secondary" {
  name            = "acctest-dbs-%[2]d"
  mssql_server_id = azurerm_sql_server.second.id
  create_secondary_mode {
    source_database_id = azurerm_mssql_database.test.id
  }

}
`, template, data.RandomInteger, data.Locations.Secondary)
}
