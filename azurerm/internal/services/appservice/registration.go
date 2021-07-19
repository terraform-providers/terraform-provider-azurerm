package appservice

import (
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/sdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/appservice/serviceplan"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/appservice/sourcecontrol"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/appservice/webapp"
)

var _ sdk.TypedServiceRegistration = Registration{}

type Registration struct{}

func (r Registration) PackagePath() string {
	return "TODO: Not implemented yet"
}

func (r Registration) WebsiteCategories() []string {
	return nil
}

func (r Registration) Name() string {
	return "AppService"
}

func (r Registration) DataSources() []sdk.DataSource {
	if features.ThreePointOh() {
		return []sdk.DataSource{
			sourcecontrol.AppServiceGitHubTokenDataSource{},
		}
	}
	return []sdk.DataSource{}
}

func (r Registration) Resources() []sdk.Resource {
	if features.ThreePointOh() {
		return []sdk.Resource{
			sourcecontrol.AppServiceSourceControlResource{},
			sourcecontrol.AppServiceGitHubTokenResource{},
			webapp.WindowsWebAppResource{},
			webapp.LinuxWebAppResource{},
			serviceplan.AppServicePlanResource{},
		}
	}
	return []sdk.Resource{}
}
