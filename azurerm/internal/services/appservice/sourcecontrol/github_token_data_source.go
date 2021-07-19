package sourcecontrol

import (
	"context"
	"fmt"
	"time"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/sdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/appservice/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

type AppServiceGitHubTokenDataSource struct{}

var _ sdk.DataSource = AppServiceGitHubTokenDataSource{}

func (d AppServiceGitHubTokenDataSource) Arguments() map[string]*pluginsdk.Schema {
	return nil
}

func (d AppServiceGitHubTokenDataSource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"token": {
			Type:      pluginsdk.TypeString,
			Sensitive: true,
			Computed:  true,
		},
	}
}

func (d AppServiceGitHubTokenDataSource) ModelObject() interface{} {
	return AppServiceGitHubTokenModel{}
}

func (d AppServiceGitHubTokenDataSource) ResourceType() string {
	return "azurerm_app_service_github_token"
}

func (d AppServiceGitHubTokenDataSource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.AppService.BaseClient

			resp, err := client.GetSourceControl(ctx, "GitHub")
			if err != nil || resp.SourceControlProperties == nil {
				return fmt.Errorf("reading App Service Source Control GitHub Token")
			}

			state := AppServiceGitHubTokenModel{}

			state.Token = utils.NormalizeNilableString(resp.Token)

			metadata.SetID(parse.AppServiceGitHubTokenId{})

			return metadata.Encode(&state)
		},
	}
}
