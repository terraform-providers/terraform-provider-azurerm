package azurerm

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tags"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func dataSourceSqlServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceArmSqlServerRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"resource_group_name": azure.SchemaResourceGroupNameForDataSource(),

			"location": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"administrator_login": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"identity": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"principal_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tenant_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"blob_extended_auditing_policy": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage_endpoint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"retention_days": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"audit_actions_and_groups": {
							Type:     schema.TypeSet,
							Computed: true,
							Set:      schema.HashString,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"storage_account_subscription_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_storage_secondary_key_in_use": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"predicate_expression": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"tags": tags.SchemaDataSource(),
		},
	}
}

func dataSourceArmSqlServerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).Sql.ServersClient
	ctx, cancel := timeouts.ForRead(meta.(*ArmClient).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	resourceGroup := d.Get("resource_group_name").(string)

	resp, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return fmt.Errorf("Sql Server %q was not found in Resource Group %q", name, resourceGroup)
		}

		return fmt.Errorf("Error retrieving Sql Server %q (Resource Group %q): %s", name, resourceGroup, err)
	}

	if id := resp.ID; id != nil {
		d.SetId(*resp.ID)
	}

	if location := resp.Location; location != nil {
		d.Set("location", azure.NormalizeLocation(*location))
	}

	if props := resp.ServerProperties; props != nil {
		d.Set("fqdn", props.FullyQualifiedDomainName)
		d.Set("version", props.Version)
		d.Set("administrator_login", props.AdministratorLogin)
	}

	if err := d.Set("identity", flattenAzureRmSqlServerIdentity(resp.Identity)); err != nil {
		return fmt.Errorf("Error setting `identity`: %+v", err)
	}

	auditingClient := meta.(*ArmClient).Sql.ExtendedServerBlobAuditingPoliciesClient
	auditingResp, err := auditingClient.Get(ctx, resourceGroup, name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[INFO] Error reading SQL Server %q Blob Auditing Policies - removing from state", d.Id())
		}

		return fmt.Errorf("Error reading SQL Server %s: %v Blob Auditing Policies", name, err)
	}

	d.Set("blob_extended_auditing_policy", flattenAzureRmSqlServerBlobAuditingPolicies(&auditingResp, d))

	return tags.FlattenAndSet(d, resp.Tags)
}
