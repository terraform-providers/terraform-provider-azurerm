package datafactory

import (
	"fmt"
	"regexp"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/datafactory/mgmt/2018-06-01/datafactory"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tags"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func dataSourceDataFactory() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Read: dataSourceDataFactoryRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"name": {
				Type:     pluginsdk.TypeString,
				Required: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile(`^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`),
					`Invalid name for Data Factory, see https://docs.microsoft.com/en-us/azure/data-factory/naming-rules`,
				),
			},

			"resource_group_name": azure.SchemaResourceGroupNameForDataSource(),

			"location": azure.SchemaLocationForDataSource(),

			"identity": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"type": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
						"identity_ids": {
							Type:     pluginsdk.TypeList,
							Computed: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},
						"principal_id": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
						"tenant_id": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
					},
				},
			},

			"github_configuration": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"account_name": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
						"branch_name": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
						"git_url": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
						"repository_name": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
						"root_folder": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
					},
				},
			},

			"vsts_configuration": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"account_name": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
						"branch_name": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
						"project_name": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
						"repository_name": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
						"root_folder": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
						"tenant_id": {
							Type:     pluginsdk.TypeString,
							Computed: true,
						},
					},
				},
			},

			"tags": tags.SchemaDataSource(),
		},
	}
}

func dataSourceDataFactoryRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DataFactory.FactoriesClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	resourceGroup := d.Get("resource_group_name").(string)

	resp, err := client.Get(ctx, resourceGroup, name, "")
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return fmt.Errorf("Error Data Factory %q (Resource Group %q) was not found", name, resourceGroup)
		}

		return fmt.Errorf("Error retrieving Data Factory %q (Resource Group %q): %+v", name, resourceGroup, err)
	}
	if resp.ID == nil || *resp.ID == "" {
		return fmt.Errorf("API returns a nil/empty id on Data Factory %q (resource group %q): %+v", name, resourceGroup, err)
	}
	d.SetId(*resp.ID)

	if location := resp.Location; location != nil {
		d.Set("location", azure.NormalizeLocation(*location))
	}

	d.Set("vsts_configuration", []interface{}{})
	d.Set("github_configuration", []interface{}{})
	repoType, repo := flattenDataFactoryRepoConfiguration(&resp)
	if repoType == datafactory.TypeBasicFactoryRepoConfigurationTypeFactoryVSTSConfiguration {
		if err := d.Set("vsts_configuration", repo); err != nil {
			return fmt.Errorf("Error setting `vsts_configuration`: %+v", err)
		}
	}
	if repoType == datafactory.TypeBasicFactoryRepoConfigurationTypeFactoryGitHubConfiguration {
		if err := d.Set("github_configuration", repo); err != nil {
			return fmt.Errorf("Error setting `github_configuration`: %+v", err)
		}
	}
	if repoType == datafactory.TypeBasicFactoryRepoConfigurationTypeFactoryRepoConfiguration {
		d.Set("vsts_configuration", repo)
		d.Set("github_configuration", repo)
	}

	identity, err := flattenDataFactoryIdentity(resp.Identity)
	if err != nil {
		return fmt.Errorf("Error flattening `identity`: %+v", err)
	}
	if err := d.Set("identity", identity); err != nil {
		return fmt.Errorf("Error setting `identity`: %+v", err)
	}

	return tags.FlattenAndSet(d, resp.Tags)
}
