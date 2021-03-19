package healthbot_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance/check"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/healthbot/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

type HealthbotResource struct{}

func TestAccHealthbot_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_healthbot_service", "test")
	r := HealthbotResource{}
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

func TestAccHealthbot_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_healthbot_service", "test")
	r := HealthbotResource{}
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

func TestAccHealthbot_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_healthbot_service", "test")
	r := HealthbotResource{}
	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccHealthbot_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_healthbot_service", "test")
	r := HealthbotResource{}
	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r HealthbotResource) Exists(ctx context.Context, client *clients.Client, state *terraform.InstanceState) (*bool, error) {
	id, err := parse.HealthbotID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.HealthBot.BotClient.Get(ctx, id.ResourceGroup, id.Name)
	if err != nil {
		return nil, fmt.Errorf("retrieving Healthbot service %q (resource group: %q): %+v", id.Name, id.ResourceGroup, err)
	}

	return utils.Bool(resp.Properties != nil), nil
}

func (r HealthbotResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctest-healthbot-%d"
  location = "%s"
}
`, data.RandomInteger, data.Locations.Primary)
}

func (r HealthbotResource) basic(data acceptance.TestData) string {
	template := r.template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_healthbot_service" "test" {
  name                = "acctest-hb-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku_name            = "F0"
}
`, template, data.RandomInteger)
}

func (r HealthbotResource) requiresImport(data acceptance.TestData) string {
	config := r.basic(data)
	return fmt.Sprintf(`
%s

resource "azurerm_healthbot_service" "import" {
  name                = azurerm_healthbot_service.test.name
  resource_group_name = azurerm_healthbot_service.test.resource_group_name
  location            = azurerm_healthbot_service.test.location
  sku_name            = azurerm_healthbot_service.test.sku_name
}
`, config)
}

func (r HealthbotResource) complete(data acceptance.TestData) string {
	template := r.template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_healthbot_service" "test" {
  name                = "acctest-hb-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku_name            = "S1"

  tags = {
    ENV = "Test"
  }
}
`, template, data.RandomInteger)
}
