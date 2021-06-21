package datafactory

import (
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/datafactory/mgmt/2018-06-01/datafactory"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/datafactory/migration"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/datafactory/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/datafactory/validate"
	keyVaultParse "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/keyvault/parse"
	keyVaultValidate "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/keyvault/validate"
	msiParse "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/msi/parse"
	msiValidate "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/msi/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tags"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceDataFactory() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Create: resourceDataFactoryCreateUpdate,
		Read:   resourceDataFactoryRead,
		Update: resourceDataFactoryCreateUpdate,
		Delete: resourceDataFactoryDelete,

		SchemaVersion: 2,
		StateUpgraders: pluginsdk.StateUpgrades(map[int]pluginsdk.StateUpgrade{
			0: migration.DataFactoryV0ToV1{},
			1: migration.DataFactoryV1ToV2{},
		}),

		// TODO: replace this with an importer which validates the ID during import
		Importer: pluginsdk.DefaultImporter(),

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(30 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(30 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.DataFactoryName(),
			},

			"location": azure.SchemaLocation(),

			// There's a bug in the Azure API where this is returned in lower-case
			// BUG: https://github.com/Azure/azure-rest-api-specs/issues/5788
			"resource_group_name": azure.SchemaResourceGroupNameDiffSuppress(),

			"identity": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"type": {
							Type:     pluginsdk.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								string(datafactory.FactoryIdentityTypeSystemAssigned),
								string(datafactory.FactoryIdentityTypeUserAssigned),
							}, false),
						},

						"identity_ids": {
							Type:     pluginsdk.TypeList,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: msiValidate.UserAssignedIdentityID,
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
				Type:          pluginsdk.TypeList,
				Optional:      true,
				MaxItems:      1,
				ConflictsWith: []string{"vsts_configuration"},
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"account_name": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"branch_name": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"git_url": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"repository_name": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"root_folder": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
					},
				},
			},

			"vsts_configuration": {
				Type:          pluginsdk.TypeList,
				Optional:      true,
				MaxItems:      1,
				ConflictsWith: []string{"github_configuration"},
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"account_name": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"branch_name": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"project_name": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"repository_name": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"root_folder": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"tenant_id": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.IsUUID,
						},
					},
				},
			},

			"global_parameter": {
				Type:     pluginsdk.TypeSet,
				Optional: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"name": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},

						"type": {
							Type:     pluginsdk.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								"Array",
								"Bool",
								"Float",
								"Int",
								"Object",
								"String",
							}, false),
						},

						"value": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
					},
				},
			},

			"public_network_enabled": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Default:  true,
			},

			"customer_managed_key_id": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: keyVaultValidate.NestedItemId,
				RequiredWith: []string{"identity.0.identity_ids"},
			},

			"tags": tags.Schema(),
		},
	}
}

func resourceDataFactoryCreateUpdate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DataFactory.FactoriesClient
	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id := parse.NewDataFactoryID(subscriptionId, d.Get("resource_group_name").(string), d.Get("name").(string))
	if d.IsNewResource() {
		existing, err := client.Get(ctx, id.ResourceGroup, id.FactoryName, "")
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("checking for presence of existing %s: %+v", id, err)
			}
		}

		if !utils.ResponseWasNotFound(existing.Response) {
			return tf.ImportAsExistsError("azurerm_data_factory", id.ID())
		}
	}

	location := azure.NormalizeLocation(d.Get("location").(string))
	dataFactory := datafactory.Factory{
		Location:          &location,
		FactoryProperties: &datafactory.FactoryProperties{},
		Tags:              tags.Expand(d.Get("tags").(map[string]interface{})),
	}

	dataFactory.PublicNetworkAccess = datafactory.PublicNetworkAccessEnabled
	enabled := d.Get("public_network_enabled").(bool)
	if !enabled {
		dataFactory.FactoryProperties.PublicNetworkAccess = datafactory.PublicNetworkAccessDisabled
	}

	if v, ok := d.GetOk("identity.0.type"); ok {
		identityType := v.(string)
		dataFactory.Identity = &datafactory.FactoryIdentity{
			Type: datafactory.FactoryIdentityType(identityType),
		}

		identityIdsRaw := d.Get("identity.0.identity_ids").([]interface{})
		if len(identityIdsRaw) > 0 {
			if identityType != string(datafactory.FactoryIdentityTypeUserAssigned) {
				return fmt.Errorf("`identity_ids` can only be specified when `type` is `%s`", string(datafactory.FactoryIdentityTypeUserAssigned))
			}

			identityIds := make(map[string]interface{})
			for _, v := range identityIdsRaw {
				identityIds[v.(string)] = make(map[string]string)
			}
			dataFactory.Identity.UserAssignedIdentities = identityIds
		}
	}

	if keyVaultKeyID, ok := d.GetOk("customer_managed_key_id"); ok {
		keyVaultKey, err := keyVaultParse.ParseOptionallyVersionedNestedItemID(keyVaultKeyID.(string))
		if err != nil {
			return fmt.Errorf("could not parse Key Vault Key ID: %+v", err)
		}

		identityIdsRaw := d.Get("identity.0.identity_ids").([]interface{})

		dataFactory.FactoryProperties.Encryption = &datafactory.EncryptionConfiguration{
			VaultBaseURL: &keyVaultKey.KeyVaultBaseUrl,
			KeyName:      &keyVaultKey.Name,
			KeyVersion:   &keyVaultKey.Version,
			Identity: &datafactory.CMKIdentityDefinition{
				UserAssignedIdentity: utils.String(identityIdsRaw[0].(string)),
			},
		}
	}

	globalParameters, err := expandDataFactoryGlobalParameters(d.Get("global_parameter").(*pluginsdk.Set).List())
	if err != nil {
		return err
	}
	dataFactory.FactoryProperties.GlobalParameters = globalParameters

	if _, err := client.CreateOrUpdate(ctx, id.ResourceGroup, id.FactoryName, dataFactory, ""); err != nil {
		return fmt.Errorf("creating/updating %s: %+v", id, err)
	}

	if hasRepo, repo := expandDataFactoryRepoConfiguration(d); hasRepo {
		repoUpdate := datafactory.FactoryRepoUpdate{
			FactoryResourceID: utils.String(id.ID()),
			RepoConfiguration: repo,
		}
		if _, err := client.ConfigureFactoryRepo(ctx, location, repoUpdate); err != nil {
			return fmt.Errorf("configuring Repository for %s: %+v", id, err)
		}
	}

	d.SetId(id.ID())

	return resourceDataFactoryRead(d, meta)
}

func resourceDataFactoryRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DataFactory.FactoriesClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.DataFactoryID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.FactoryName, "")
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("retrieving %s: %+v", id, err)
	}

	d.Set("name", id.FactoryName)
	d.Set("resource_group_name", id.ResourceGroup)
	if location := resp.Location; location != nil {
		d.Set("location", azure.NormalizeLocation(*location))
	}

	if factoryProps := resp.FactoryProperties; factoryProps != nil {
		if enc := factoryProps.Encryption; enc != nil {
			if enc.VaultBaseURL != nil && enc.KeyName != nil && enc.KeyVersion != nil {
				versionedKey := fmt.Sprintf("%skeys/%s/%s", *enc.VaultBaseURL, *enc.KeyName, *enc.KeyVersion)
				if err := d.Set("customer_managed_key_id", versionedKey); err != nil {
					return fmt.Errorf("Error setting `customer_managed_key_id`: %+v", err)
				}
			}
		}

		if err := d.Set("global_parameter", flattenDataFactoryGlobalParameters(factoryProps.GlobalParameters)); err != nil {
			return fmt.Errorf("setting `global_parameter`: %+v", err)
		}
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

	// This variable isn't returned from the API if it hasn't been passed in first but we know the default is `true`
	if resp.PublicNetworkAccess != "" {
		d.Set("public_network_enabled", resp.PublicNetworkAccess == datafactory.PublicNetworkAccessEnabled)
	}

	return tags.FlattenAndSet(d, resp.Tags)
}

func resourceDataFactoryDelete(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DataFactory.FactoriesClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.DataFactoryID(d.Id())
	if err != nil {
		return err
	}

	response, err := client.Delete(ctx, id.ResourceGroup, id.FactoryName)
	if err != nil {
		if !utils.ResponseWasNotFound(response) {
			return fmt.Errorf("deleting %s: %+v", id, err)
		}
	}

	return nil
}

func expandDataFactoryRepoConfiguration(d *pluginsdk.ResourceData) (bool, datafactory.BasicFactoryRepoConfiguration) {
	if vstsList, ok := d.GetOk("vsts_configuration"); ok {
		vsts := vstsList.([]interface{})[0].(map[string]interface{})
		accountName := vsts["account_name"].(string)
		branchName := vsts["branch_name"].(string)
		projectName := vsts["project_name"].(string)
		repositoryName := vsts["repository_name"].(string)
		rootFolder := vsts["root_folder"].(string)
		tenantID := vsts["tenant_id"].(string)
		return true, &datafactory.FactoryVSTSConfiguration{
			AccountName:         &accountName,
			CollaborationBranch: &branchName,
			ProjectName:         &projectName,
			RepositoryName:      &repositoryName,
			RootFolder:          &rootFolder,
			TenantID:            &tenantID,
		}
	}

	if githubList, ok := d.GetOk("github_configuration"); ok {
		github := githubList.([]interface{})[0].(map[string]interface{})
		accountName := github["account_name"].(string)
		branchName := github["branch_name"].(string)
		gitURL := github["git_url"].(string)
		repositoryName := github["repository_name"].(string)
		rootFolder := github["root_folder"].(string)
		return true, &datafactory.FactoryGitHubConfiguration{
			AccountName:         &accountName,
			CollaborationBranch: &branchName,
			HostName:            &gitURL,
			RepositoryName:      &repositoryName,
			RootFolder:          &rootFolder,
		}
	}

	return false, nil
}

func expandDataFactoryGlobalParameters(input []interface{}) (map[string]*datafactory.GlobalParameterSpecification, error) {
	if len(input) == 0 {
		return nil, nil
	}
	result := make(map[string]*datafactory.GlobalParameterSpecification)
	for _, item := range input {
		if item == nil {
			continue
		}
		v := item.(map[string]interface{})

		name := v["name"].(string)
		if _, ok := v[name]; ok {
			return nil, fmt.Errorf("duplicate parameter name")
		}

		result[name] = &datafactory.GlobalParameterSpecification{
			Type:  datafactory.GlobalParameterType(v["type"].(string)),
			Value: v["value"].(string),
		}
	}
	return result, nil
}

func flattenDataFactoryRepoConfiguration(factory *datafactory.Factory) (datafactory.TypeBasicFactoryRepoConfiguration, []interface{}) {
	result := make([]interface{}, 0)

	if properties := factory.FactoryProperties; properties != nil {
		repo := properties.RepoConfiguration
		if repo != nil {
			settings := map[string]interface{}{}
			if config, test := repo.AsFactoryGitHubConfiguration(); test {
				if config.AccountName != nil {
					settings["account_name"] = *config.AccountName
				}
				if config.CollaborationBranch != nil {
					settings["branch_name"] = *config.CollaborationBranch
				}
				if config.HostName != nil {
					settings["git_url"] = *config.HostName
				}
				if config.RepositoryName != nil {
					settings["repository_name"] = *config.RepositoryName
				}
				if config.RootFolder != nil {
					settings["root_folder"] = *config.RootFolder
				}
				return datafactory.TypeBasicFactoryRepoConfigurationTypeFactoryGitHubConfiguration, append(result, settings)
			}
			if config, test := repo.AsFactoryVSTSConfiguration(); test {
				if config.AccountName != nil {
					settings["account_name"] = *config.AccountName
				}
				if config.CollaborationBranch != nil {
					settings["branch_name"] = *config.CollaborationBranch
				}
				if config.ProjectName != nil {
					settings["project_name"] = *config.ProjectName
				}
				if config.RepositoryName != nil {
					settings["repository_name"] = *config.RepositoryName
				}
				if config.RootFolder != nil {
					settings["root_folder"] = *config.RootFolder
				}
				if config.TenantID != nil {
					settings["tenant_id"] = *config.TenantID
				}
				return datafactory.TypeBasicFactoryRepoConfigurationTypeFactoryVSTSConfiguration, append(result, settings)
			}
		}
	}
	return datafactory.TypeBasicFactoryRepoConfigurationTypeFactoryRepoConfiguration, result
}

func flattenDataFactoryIdentity(identity *datafactory.FactoryIdentity) (interface{}, error) {
	if identity == nil {
		return []interface{}{}, nil
	}

	principalId := ""
	if identity.PrincipalID != nil {
		principalId = identity.PrincipalID.String()
	}
	tenantId := ""
	if identity.TenantID != nil {
		tenantId = identity.TenantID.String()
	}
	var identityIds []string
	if identity.UserAssignedIdentities != nil {
		for key := range identity.UserAssignedIdentities {
			id, err := msiParse.UserAssignedIdentityID(key)
			if err != nil {
				return nil, err
			}
			identityIds = append(identityIds, id.ID())
		}
	}

	return []interface{}{
		map[string]interface{}{
			"principal_id": principalId,
			"tenant_id":    tenantId,
			"type":         string(identity.Type),
			"identity_ids": identityIds,
		},
	}, nil
}

func flattenDataFactoryGlobalParameters(input map[string]*datafactory.GlobalParameterSpecification) []interface{} {
	if len(input) == 0 {
		return []interface{}{}
	}
	result := make([]interface{}, 0)
	for name, item := range input {
		result = append(result, map[string]interface{}{
			"name":  name,
			"type":  string(item.Type),
			"value": item.Value,
		})
	}
	return result
}
