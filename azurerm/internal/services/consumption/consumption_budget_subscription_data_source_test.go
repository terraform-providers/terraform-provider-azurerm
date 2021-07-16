package consumption_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance/check"
)

type BudgetSubscriptionDataSource struct{}

func TestAccBudgetSubscriptionDataSource_current(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_consumption_budget", "current")
	r := BudgetSubscriptionDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.template(data),
		},
		{
			Config: r.basic(),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("subscription_id").HasValue(data.Client().SubscriptionID),
				check.That(data.ResourceName).Key("name").HasValue("acctestconsumptionbudget-sub"),
			),
		},
	})
}

func (d BudgetSubscriptionDataSource) basic() string {
	return `
provider "azurerm" {
  features {}
}

data "azurerm_subscription" "current" {}

data "azurerm_consumption_budget" "current" {
  name            = "acctestconsumptionbudget-sub"
  subscription_id = data.azurerm_subscription.current.subscription_id
}
`
}

func (BudgetSubscriptionDataSource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

data "azurerm_subscription" "current" {}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
  lifecycle {
    ignore_changes = [tags]
  }
}

resource "azurerm_consumption_budget_resource_group" "test" {
  name              = "acctestconsumptionbudget-sub"
  resource_group_id = azurerm_resource_group.test.id

  amount     = 1000
  time_grain = "Monthly"

  time_period {
    start_date = "%s"
  }

  filter {
    tag {
      name = "foo"
      values = [
        "bar",
      ]
    }
  }

  notification {
    enabled   = true
    threshold = 90.0
    operator  = "EqualTo"

    contact_emails = [
      "foo@example.com",
      "bar@example.com",
    ]
  }
}
`, data.RandomInteger, data.Locations.Primary, consumptionBudgetTestStartDate().Format(time.RFC3339))
}
