package hpccache

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/storagecache/mgmt/2021-03-01/storagecache"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/hpccache/parse"
	azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceHPCCache() *schema.Resource {
	return &schema.Resource{
		Create: resourceHPCCacheCreateOrUpdate,
		Update: resourceHPCCacheCreateOrUpdate,
		Read:   resourceHPCCacheRead,
		Delete: resourceHPCCacheDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		Importer: azSchema.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.CacheID(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"resource_group_name": azure.SchemaResourceGroupName(),

			"location": azure.SchemaLocation(),

			"cache_size_in_gb": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.IntInSlice([]int{
					3072,
					6144,
					12288,
					24576,
					49152,
				}),
			},

			"subnet_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: azure.ValidateResourceIDOrEmpty,
			},

			"sku_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Standard_2G",
					"Standard_4G",
					"Standard_8G",
				}, false),
			},

			"mtu": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      1500,
				ValidateFunc: validation.IntBetween(576, 1500),
			},

			"ntp_server": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "time.windows.com",
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"dns": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"servers": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 3,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.IsIPAddress,
							},
						},

						"search_domain": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
					},
				},
			},

			"directory_flat_file": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_file_uri": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"passwd_file_uri": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
					},
				},
				ConflictsWith: []string{"directory_ldap"},
			},

			"directory_ldap": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},

						"base_dn": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},

						"conn_encrypted": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"certificate_validation_uri": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},

						"download_certificate": {
							Type:         schema.TypeBool,
							Optional:     true,
							RequiredWith: []string{"directory_ldap.0.certificate_validation_uri"},
						},

						"bind": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dn": {
										Type:         schema.TypeString,
										Sensitive:    true,
										Required:     true,
										ValidateFunc: validation.StringIsNotEmpty,
									},
									"password": {
										Type:         schema.TypeString,
										Sensitive:    true,
										Required:     true,
										ValidateFunc: validation.StringIsNotEmpty,
									},
								},
							},
						},
					},
				},
				ConflictsWith: []string{"directory_flat_file"},
			},

			// TODO 3.0: remove this property
			"root_squash_enabled": {
				Type:       schema.TypeBool,
				Optional:   true,
				Computed:   true,
				Deprecated: "This property is not functional and will be deprecated in favor of `default_access_policy.0.access_rule.x.root_squash_enabled`, where the scope of access_rule is `default`.",
			},

			"default_access_policy": {
				Type:     schema.TypeList,
				MinItems: 1,
				MaxItems: 1,
				Optional: true,
				// This is computed because there is always a "default" policy in the cache. It is created together with the cache, and users can't remove it.
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_rule": {
							Type:     schema.TypeSet,
							Required: true,
							MinItems: 1,
							MaxItems: 3,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"scope": {
										Type:     schema.TypeString,
										Required: true,
										ValidateFunc: validation.StringInSlice([]string{
											string(storagecache.Default),
											string(storagecache.Network),
											string(storagecache.Host),
										}, false),
									},

									"access": {
										Type:     schema.TypeString,
										Required: true,
										ValidateFunc: validation.StringInSlice([]string{
											string(storagecache.NfsAccessRuleAccessRw),
											string(storagecache.NfsAccessRuleAccessRo),
											string(storagecache.NfsAccessRuleAccessNo),
										}, false),
									},

									"filter": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringIsNotEmpty,
									},

									"suid_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"submount_access_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"root_squash_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"anonymous_uid": {
										Type:         schema.TypeInt,
										Optional:     true,
										ValidateFunc: validation.IntAtLeast(0),
									},

									"anonymous_gid": {
										Type:         schema.TypeInt,
										Optional:     true,
										ValidateFunc: validation.IntAtLeast(0),
									},
								},
							},
						},
					},
				},
			},

			"mount_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceHPCCacheCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).HPCCache.CachesClient
	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	log.Printf("[INFO] preparing arguments for Azure HPC Cache creation.")
	name := d.Get("name").(string)
	resourceGroup := d.Get("resource_group_name").(string)

	id := parse.NewCacheID(subscriptionId, resourceGroup, name)

	if d.IsNewResource() {
		existing, err := client.Get(ctx, id.ResourceGroup, id.Name)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("Error checking for presence of existing HPC Cache %q: %s", id, err)
			}
		}

		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_hpc_cache", id.ID())
		}
	}

	location := d.Get("location").(string)
	cacheSize := d.Get("cache_size_in_gb").(int)
	subnet := d.Get("subnet_id").(string)
	skuName := d.Get("sku_name").(string)

	var accessPolicies []storagecache.NfsAccessPolicy
	if !d.IsNewResource() {
		existing, err := client.Get(ctx, id.ResourceGroup, id.Name)
		if err != nil {
			return fmt.Errorf("retrieving existing HPC Cache %q: %v", id, err)
		}
		if prop := existing.CacheProperties; prop != nil {
			if settings := existing.SecuritySettings; settings != nil {
				if policies := settings.AccessPolicies; policies != nil {
					accessPolicies = *policies
				}
			}
		}
	}
	defaultAccessPolicy := expandStorageCacheDefaultAccessPolicy(d.Get("default_access_policy").([]interface{}))
	if defaultAccessPolicy != nil {
		var err error
		accessPolicies, err = CacheInsertOrUpdateAccessPolicy(accessPolicies, *defaultAccessPolicy)
		if err != nil {
			return err
		}
	}

	directorySetting := expandStorageCacheDirectorySettings(d)

	cache := &storagecache.Cache{
		Name:     utils.String(name),
		Location: utils.String(location),
		CacheProperties: &storagecache.CacheProperties{
			CacheSizeGB:     utils.Int32(int32(cacheSize)),
			Subnet:          utils.String(subnet),
			NetworkSettings: expandStorageCacheNetworkSettings(d),
			SecuritySettings: &storagecache.CacheSecuritySettings{
				AccessPolicies: &accessPolicies,
			},
			DirectoryServicesSettings: directorySetting,
		},
		Sku: &storagecache.CacheSku{
			Name: utils.String(skuName),
		},
	}

	future, err := client.CreateOrUpdate(ctx, resourceGroup, name, cache)
	if err != nil {
		return fmt.Errorf("Error creating HPC Cache %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("Error waiting for HPC Cache %q (Resource Group %q) to finish provisioning: %+v", name, resourceGroup, err)
	}

	// If any directory setting is set, we'll further check the `usernameDownloaded` in response to ensure the configuration is correct, and the cache is functional.
	// There are situations that the LRO succeeded, whilst ends up with a non-functional cache (e.g. providing some invalid flat file setting).
	if directorySetting != nil {
		resp, err := client.Get(ctx, resourceGroup, name)
		if err != nil {
			return fmt.Errorf("Error retrieving HPC Cache %q (Resource Group %q): %+v", name, resourceGroup, err)
		}

		prop := resp.CacheProperties
		if prop == nil {
			return fmt.Errorf("Unepxected nil `cacheProperties` in response")
		}
		ds := prop.DirectoryServicesSettings
		if ds == nil {
			return fmt.Errorf("Unexpected nil `directoryServicesSettings` in response")
		}
		ud := ds.UsernameDownload
		if ud == nil {
			return fmt.Errorf("Unexpected nil `usernameDownload` in response")
		}
		if ud.UsernameDownloaded != storagecache.UsernameDownloadedTypeYes {
			return fmt.Errorf("failed to download directory info, current status: %s", ud.UsernameDownloaded)
		}
	}

	d.SetId(id.ID())

	return resourceHPCCacheRead(d, meta)
}

func resourceHPCCacheRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).HPCCache.CachesClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.CacheID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.Name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[DEBUG] HPC Cache %q was not found in Resource Group %q - removing from state!", id.Name, id.ResourceGroup)
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error retrieving HPC Cache %q: %+v", id, err)
	}

	d.Set("name", id.Name)
	d.Set("resource_group_name", id.ResourceGroup)
	d.Set("location", resp.Location)

	if props := resp.CacheProperties; props != nil {
		d.Set("cache_size_in_gb", props.CacheSizeGB)
		d.Set("subnet_id", props.Subnet)
		d.Set("mount_addresses", utils.FlattenStringSlice(props.MountAddresses))

		mtu, ntpServer, dnsSetting := flattenStorageCacheNetworkSettings(props.NetworkSettings)
		d.Set("mtu", mtu)
		d.Set("ntp_server", ntpServer)
		if err := d.Set("dns", dnsSetting); err != nil {
			return fmt.Errorf("setting `dns`: %v", err)
		}

		flatFile, ldap, err := flattenStorageCacheDirectorySettings(d, props.DirectoryServicesSettings)
		if err != nil {
			return err
		}

		if err := d.Set("directory_flat_file", flatFile); err != nil {
			fmt.Errorf("setting `directory_flat_file`: %v", err)
		}

		if err := d.Set("directory_ldap", ldap); err != nil {
			fmt.Errorf("setting `directory_ldap`: %v", err)
		}

		if securitySettings := props.SecuritySettings; securitySettings != nil {
			if securitySettings.AccessPolicies != nil {
				defaultPolicy := CacheGetAccessPolicyByName(*securitySettings.AccessPolicies, "default")
				if defaultPolicy != nil {
					defaultAccessPolicy, err := flattenStorageCacheNfsDefaultAccessPolicy(*defaultPolicy)
					if err != nil {
						return err
					}
					if err := d.Set("default_access_policy", defaultAccessPolicy); err != nil {
						return fmt.Errorf("setting `default_access_policy`: %v", err)
					}

					// Set the "root_squash_enabled" for whatever is set in the config, to make any existing .tf that has specified this property
					// not encounter plan diff.
					// TODO 3.0 - remove this part.
					d.Set("root_squash_enabled", d.Get("root_squash_enabled"))
				}
			}
		}
	}

	if sku := resp.Sku; sku != nil {
		d.Set("sku_name", sku.Name)
	}

	return nil
}

func resourceHPCCacheDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).HPCCache.CachesClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.CacheID(d.Id())
	if err != nil {
		return err
	}

	future, err := client.Delete(ctx, id.ResourceGroup, id.Name)
	if err != nil {
		return fmt.Errorf("Error deleting HPC Cache %q (Resource Group %q): %+v", id.Name, id.ResourceGroup, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("Error waiting for deletion of HPC Cache %q (Resource Group %q): %+v", id.Name, id.ResourceGroup, err)
	}

	return nil
}

func expandStorageCacheDefaultAccessPolicy(input []interface{}) *storagecache.NfsAccessPolicy {
	if len(input) == 0 {
		return nil
	}

	return &storagecache.NfsAccessPolicy{
		Name:        utils.String("default"),
		AccessRules: expandStorageCacheNfsAccessRules(input[0].(map[string]interface{})["access_rule"].(*schema.Set).List()),
	}
}

func flattenStorageCacheNfsDefaultAccessPolicy(input storagecache.NfsAccessPolicy) ([]interface{}, error) {
	rules, err := flattenStorageCacheNfsAccessRules(input.AccessRules)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		map[string]interface{}{
			"access_rule": rules,
		},
	}, nil
}

func expandStorageCacheNfsAccessRules(input []interface{}) *[]storagecache.NfsAccessRule {
	var out []storagecache.NfsAccessRule
	for _, accessRuleRaw := range input {
		b := accessRuleRaw.(map[string]interface{})
		out = append(out, storagecache.NfsAccessRule{
			Scope:          storagecache.NfsAccessRuleScope(b["scope"].(string)),
			Access:         storagecache.NfsAccessRuleAccess(b["access"].(string)),
			Filter:         utils.String(b["filter"].(string)),
			Suid:           utils.Bool(b["suid_enabled"].(bool)),
			SubmountAccess: utils.Bool(b["submount_access_enabled"].(bool)),
			RootSquash:     utils.Bool(b["root_squash_enabled"].(bool)),
			AnonymousUID:   utils.String(strconv.Itoa(b["anonymous_uid"].(int))),
			AnonymousGID:   utils.String(strconv.Itoa(b["anonymous_gid"].(int))),
		})
	}
	return &out
}

func flattenStorageCacheNfsAccessRules(input *[]storagecache.NfsAccessRule) ([]interface{}, error) {
	if input == nil {
		return nil, nil
	}

	var rules []interface{}
	for _, accessRule := range *input {
		filter := ""
		if accessRule.Filter != nil {
			filter = *accessRule.Filter
		}

		suidEnabled := false
		if accessRule.Suid != nil {
			suidEnabled = *accessRule.Suid
		}

		submountAccessEnabled := false
		if accessRule.SubmountAccess != nil {
			submountAccessEnabled = *accessRule.SubmountAccess
		}

		rootSquashEnabled := false
		if accessRule.RootSquash != nil {
			rootSquashEnabled = *accessRule.RootSquash
		}

		anonymousUID := 0
		if accessRule.AnonymousUID != nil {
			var err error
			anonymousUID, err = strconv.Atoi(*accessRule.AnonymousUID)
			if err != nil {
				return nil, fmt.Errorf("converting `anonymous_uid` from string to int")
			}
		}

		anonymousGID := 0
		if accessRule.AnonymousGID != nil {
			var err error
			anonymousGID, err = strconv.Atoi(*accessRule.AnonymousGID)
			if err != nil {
				return nil, fmt.Errorf("converting `anonymous_gid` from string to int")
			}
		}

		rules = append(rules, map[string]interface{}{
			"scope":                   accessRule.Scope,
			"access":                  accessRule.Access,
			"filter":                  filter,
			"suid_enabled":            suidEnabled,
			"submount_access_enabled": submountAccessEnabled,
			"root_squash_enabled":     rootSquashEnabled,
			"anonymous_uid":           anonymousUID,
			"anonymous_gid":           anonymousGID,
		})
	}

	return rules, nil
}

func expandStorageCacheNetworkSettings(d *schema.ResourceData) *storagecache.CacheNetworkSettings {
	out := &storagecache.CacheNetworkSettings{
		Mtu:       utils.Int32(int32(d.Get("mtu").(int))),
		NtpServer: utils.String(d.Get("ntp_server").(string)),
	}

	if dnsSetting, ok := d.GetOk("dns"); ok {
		dnsSetting := dnsSetting.([]interface{})[0].(map[string]interface{})
		out.DNSServers = utils.ExpandStringSlice(dnsSetting["servers"].([]interface{}))
		searchDomain := dnsSetting["search_domain"].(string)
		if searchDomain != "" {
			out.DNSSearchDomain = &searchDomain
		}
	}
	return out
}

func flattenStorageCacheNetworkSettings(settings *storagecache.CacheNetworkSettings) (mtu int, ntpServer string, dnsSetting []interface{}) {
	if settings == nil {
		return
	}

	if settings.Mtu != nil {
		mtu = int(*settings.Mtu)
	}

	if settings.NtpServer != nil {
		ntpServer = *settings.NtpServer
	}

	if settings.DNSServers != nil {
		dnsServers := utils.FlattenStringSlice(settings.DNSServers)

		searchDomain := ""
		if settings.DNSSearchDomain != nil {
			searchDomain = *settings.DNSSearchDomain
		}

		dnsSetting = []interface{}{
			map[string]interface{}{
				"servers":       dnsServers,
				"search_domain": searchDomain,
			},
		}
	}
	return
}

func expandStorageCacheDirectorySettings(d *schema.ResourceData) *storagecache.CacheDirectorySettings {
	if raw := d.Get("directory_flat_file").([]interface{}); len(raw) != 0 {
		b := raw[0].(map[string]interface{})
		return &storagecache.CacheDirectorySettings{
			UsernameDownload: &storagecache.CacheUsernameDownloadSettings{
				ExtendedGroups: utils.Bool(true),
				UsernameSource: storagecache.UsernameSourceFile,
				GroupFileURI:   utils.String(b["group_file_uri"].(string)),
				UserFileURI:    utils.String(b["passwd_file_uri"].(string)),
			},
		}
	}

	if raw := d.Get("directory_ldap").([]interface{}); len(raw) != 0 {
		b := raw[0].(map[string]interface{})
		var certValidationUriPtr *string
		certValidationUri := b["certificate_validation_uri"].(string)
		if certValidationUri != "" {
			certValidationUriPtr = &certValidationUri
		}
		return &storagecache.CacheDirectorySettings{
			UsernameDownload: &storagecache.CacheUsernameDownloadSettings{
				ExtendedGroups:          utils.Bool(true),
				UsernameSource:          storagecache.UsernameSourceLDAP,
				LdapServer:              utils.String(b["server"].(string)),
				LdapBaseDN:              utils.String(b["base_dn"].(string)),
				EncryptLdapConnection:   utils.Bool(b["conn_encrypted"].(bool)),
				RequireValidCertificate: utils.Bool(certValidationUriPtr != nil),
				AutoDownloadCertificate: utils.Bool(b["download_certificate"].(bool)),
				CaCertificateURI:        certValidationUriPtr,
				Credentials:             expandStorageCacheDirectoryLdapBind(b["bind"].([]interface{})),
			},
		}
	}

	return nil
}

func flattenStorageCacheDirectorySettings(d *schema.ResourceData, input *storagecache.CacheDirectorySettings) (flatFile, ldap []interface{}, err error) {
	if input == nil || input.UsernameDownload == nil {
		return nil, nil, nil
	}

	ud := input.UsernameDownload
	switch ud.UsernameSource {
	case storagecache.UsernameSourceFile:
		var groupFileUri string
		if ud.GroupFileURI != nil {
			groupFileUri = *ud.GroupFileURI
		}

		var passwdFileUri string
		if ud.UserFileURI != nil {
			passwdFileUri = *ud.UserFileURI
		}

		return []interface{}{
			map[string]interface{}{
				"group_file_uri":  groupFileUri,
				"passwd_file_uri": passwdFileUri,
			},
		}, nil, nil
	case storagecache.UsernameSourceLDAP:
		var server string
		if ud.LdapServer != nil {
			server = *ud.LdapServer
		}

		var baseDn string
		if ud.LdapBaseDN != nil {
			baseDn = *ud.LdapBaseDN
		}

		var connEncrypted bool
		if ud.EncryptLdapConnection != nil {
			connEncrypted = *ud.EncryptLdapConnection
		}

		var certValidationUri string
		if ud.CaCertificateURI != nil {
			certValidationUri = *ud.CaCertificateURI
		}

		var downloadCert bool
		if ud.AutoDownloadCertificate != nil {
			downloadCert = *ud.AutoDownloadCertificate
		}

		return nil, []interface{}{
			map[string]interface{}{
				"server":                     server,
				"base_dn":                    baseDn,
				"conn_encrypted":             connEncrypted,
				"certificate_validation_uri": certValidationUri,
				"download_certificate":       downloadCert,
				"bind":                       flattenStorageCacheDirectoryLdapBind(d),
			},
		}, nil
	default:
		return nil, nil, nil
	}
}

func flattenStorageCacheDirectoryLdapBind(d *schema.ResourceData) []interface{} {
	// Since the credentials are never returned from response. We will set whatever specified in the config back to state as the best effort.
	ldap := d.Get("directory_ldap").([]interface{})
	if len(ldap) == 0 {
		return nil
	}
	return ldap[0].(map[string]interface{})["bind"].([]interface{})
}

func expandStorageCacheDirectoryLdapBind(input []interface{}) *storagecache.CacheUsernameDownloadSettingsCredentials {
	if input == nil || len(input) == 0 {
		return nil
	}
	b := input[0].(map[string]interface{})
	return &storagecache.CacheUsernameDownloadSettingsCredentials{
		BindDn:       utils.String(b["dn"].(string)),
		BindPassword: utils.String(b["password"].(string)),
	}
}
