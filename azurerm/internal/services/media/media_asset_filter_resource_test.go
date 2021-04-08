package media_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance/check"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/media/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

type AssetFilterResource struct{}

func TestAccAssetFilter_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_media_asset_filter", "test")
	r := AssetFilterResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("name").HasValue("Filter-1"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAssetFilter_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_media_asset_filter", "test")
	r := AssetFilterResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("name").HasValue("Filter-1"),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport),
	})
}

func TestAccAssetFilter_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_media_asset_filter", "test")
	r := AssetFilterResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("first_quality_bitrate").HasValue("128000"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAssetFilter_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_media_asset_filter", "test")
	r := AssetFilterResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("name").HasValue("Filter-1"),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: resource.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("first_quality_bitrate").HasValue("128000"),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: resource.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("name").HasValue("Filter-1"),
			),
		},
		data.ImportStep(),
	})
}

func (AssetFilterResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	id, err := parse.AssetFilterID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := clients.Media.AssetFiltersClient.Get(ctx, id.ResourceGroup, id.MediaserviceName, id.AssetName, id.Name)
	if err != nil {
		return nil, fmt.Errorf("retrieving Asset Filter %s (Media Services Account %s) (resource group: %s): %v", id.Name, id.MediaserviceName, id.ResourceGroup, err)
	}

	return utils.Bool(resp.FilterProperties != nil), nil
}

func (r AssetFilterResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_media_asset_filter" "test" {
  name                        = "Filter-1"
  asset_id                    = azurerm_media_asset.test.id
}

`, r.template(data))
}

func (r AssetFilterResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_media_asset_filter" "import" {
  name                        = azurerm_media_asset_filter.test.name
  asset_id                    = azurerm_media_asset.test.id
}

`, r.basic(data))
}

func (r AssetFilterResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_media_asset_filter" "test" {
  name                        = "Filter-1"
  asset_id                    = azurerm_media_asset.test.id
  first_quality_bitrate       = 128000

  presentation_time_range {
    start_timescale                  = 0
    end_timescale                    = 170000000
    presentation_window_in_timescale = 9223372036854775000
    live_backoff_in_timescale        = 0
    timescale_increment_in_seconds   = 10000000
    force_end_timescale              = false
  }

  track_selection {
    condition {
      property  = "Type"
      operation = "Equal"
      value     = "Audio"
    }

    condition {
      property  = "Language"
      operation = "NotEqual"
      value     = "en"
    }

    condition {
      property  = "FourCC"
      operation = "NotEqual"
      value     = "EC-3"
    }
  }


  track_selection {
    condition {
      property  = "Type"
      operation = "Equal"
      value     = "Video"
    }

    condition {
      property  = "Bitrate"
      operation = "Equal"
      value     = "3000000-5000000"
    }
  }
}

`, r.template(data))
}

func (AssetFilterResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-media-%d"
  location = "%s"
}

resource "azurerm_storage_account" "test" {
  name                     = "acctestsa1%s"
  resource_group_name      = azurerm_resource_group.test.name
  location                 = azurerm_resource_group.test.location
  account_tier             = "Standard"
  account_replication_type = "GRS"
}

resource "azurerm_media_services_account" "test" {
  name                = "acctestmsa%s"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name

  storage_account {
    id         = azurerm_storage_account.test.id
    is_primary = true
  }
}

resource "azurerm_media_asset" "test" {
  name                        = "test"
  resource_group_name         = azurerm_resource_group.test.name
  media_services_account_name = azurerm_media_services_account.test.name
}
`, data.RandomInteger, data.Locations.Primary, data.RandomString, data.RandomString)
}
