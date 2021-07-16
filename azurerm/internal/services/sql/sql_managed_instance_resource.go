package sql

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2018-06-01-preview/sql"
	"github.com/hashicorp/go-azure-helpers/response"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/mssql/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/sql/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tags"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmSqlMiServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmSqlMiServerCreateUpdate,
		Read:   resourceArmSqlMiServerRead,
		Update: resourceArmSqlMiServerCreateUpdate,
		Delete: resourceArmSqlMiServerDelete,
		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.ManagedInstanceID(id)
			return err
		}),

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(24 * time.Hour),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(24 * time.Hour),
			Delete: schema.DefaultTimeout(24 * time.Hour),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.ValidateMsSqlServerName,
			},

			"location": azure.SchemaLocation(),

			"resource_group_name": azure.SchemaResourceGroupName(),

			"sku_name": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"GP_Gen4",
					"GP_Gen5",
					"BC_Gen4",
					"BC_Gen5",
				}, false),
			},

			"administrator_login": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"administrator_login_password": {
				Type:         schema.TypeString,
				Required:     true,
				Sensitive:    true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"vcores": {
				Type:     schema.TypeInt,
				Required: true,
				ValidateFunc: validation.IntInSlice([]int{
					4,
					8,
					16,
					24,
					32,
					40,
					64,
					80,
				}),
			},

			"storage_size_in_gb": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(32, 8192),
			},

			"license_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"LicenseIncluded",
					"BasePrice",
				}, true),
			},

			"subnet_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: azure.ValidateResourceID,
			},

			"collation": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "SQL_Latin1_General_CP1_CI_AS",
				ValidateFunc: validation.StringIsNotEmpty,
				ForceNew:     true,
			},

			"public_data_endpoint_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"minimum_tls_version": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "1.2",
				ValidateFunc: validation.StringInSlice([]string{
					"1.0",
					"1.1",
					"1.2",
				}, false),
			},

			"proxy_override": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  string(sql.ManagedInstanceProxyOverrideDefault),
				ValidateFunc: validation.StringInSlice([]string{
					string(sql.ManagedInstanceProxyOverrideDefault),
					string(sql.ManagedInstanceProxyOverrideRedirect),
					string(sql.ManagedInstanceProxyOverrideProxy),
				}, false),
			},

			"timezone_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "UTC",
				ValidateFunc: validation.StringIsNotEmpty,
				ForceNew:     true,
			},

			"fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"tags": tags.Schema(),
		},
	}
}

func resourceArmSqlMiServerCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Sql.ManagedInstancesClient
	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	resGroup := d.Get("resource_group_name").(string)
	id := parse.NewManagedInstanceID(subscriptionId, resGroup, name)

	if d.IsNewResource() {
		existing, err := client.Get(ctx, id.ResourceGroup, id.Name)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("checking for presence of existing Managed Instance %q: %s", id.ID(), err)
			}
		}

		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_sql_managed_instance", *existing.ID)
		}
	}

	sku, err := expandManagedInstanceSkuName(d.Get("sku_name").(string))
	if err != nil {
		return fmt.Errorf("error expanding `sku_name` for SQL Managed Instance Server %q: %v", id.ID(), err)
	}

	parameters := sql.ManagedInstance{
		Sku:      sku,
		Location: utils.String(azure.NormalizeLocation(d.Get("location").(string))),
		Tags:     tags.Expand(d.Get("tags").(map[string]interface{})),
		ManagedInstanceProperties: &sql.ManagedInstanceProperties{
			LicenseType:               sql.ManagedInstanceLicenseType(d.Get("license_type").(string)),
			AdministratorLogin:        utils.String(d.Get("administrator_login").(string)),
			SubnetID:                  utils.String(d.Get("subnet_id").(string)),
			StorageSizeInGB:           utils.Int32(int32(d.Get("storage_size_in_gb").(int))),
			VCores:                    utils.Int32(int32(d.Get("vcores").(int))),
			Collation:                 utils.String(d.Get("collation").(string)),
			PublicDataEndpointEnabled: utils.Bool(d.Get("public_data_endpoint_enabled").(bool)),
			MinimalTLSVersion:         utils.String(d.Get("minimum_tls_version").(string)),
			ProxyOverride:             sql.ManagedInstanceProxyOverride(d.Get("proxy_override").(string)),
			TimezoneID:                utils.String(d.Get("timezone_id").(string)),
		},
	}

	if d.HasChange("administrator_login_password") {
		adminPassword := d.Get("administrator_login_password").(string)
		parameters.ManagedInstanceProperties.AdministratorLoginPassword = utils.String(adminPassword)
	}

	future, err := client.CreateOrUpdate(ctx, resGroup, name, parameters)
	if err != nil {
		return err
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		if response.WasConflict(future.Response()) {
			return fmt.Errorf("SQL Server names need to be globally unique and %q is already in use.", name)
		}

		return err
	}

	d.SetId(id.ID())

	return resourceArmSqlMiServerRead(d, meta)
}

func resourceArmSqlMiServerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Sql.ManagedInstancesClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.ManagedInstanceID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.Name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[INFO] Error reading SQL Managed Instance %q - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf("reading SQL Managed Instance %q: %v", id.ID(), err)
	}

	d.Set("name", id.Name)
	d.Set("resource_group_name", id.ResourceGroup)

	if location := resp.Location; location != nil {
		d.Set("location", azure.NormalizeLocation(*location))
	}

	if sku := resp.Sku; sku != nil {
		d.Set("sku_name", sku.Name)
	}

	if props := resp.ManagedInstanceProperties; props != nil {
		d.Set("license_type", props.LicenseType)
		d.Set("administrator_login", props.AdministratorLogin)
		d.Set("subnet_id", props.SubnetID)
		d.Set("storage_size_in_gb", props.StorageSizeInGB)
		d.Set("vcores", props.VCores)
		d.Set("fqdn", props.FullyQualifiedDomainName)
		d.Set("collation", props.Collation)
		d.Set("public_data_endpoint_enabled", props.PublicDataEndpointEnabled)
		d.Set("minimum_tls_version", props.MinimalTLSVersion)
		d.Set("proxy_override", props.ProxyOverride)
		d.Set("timezone_id", props.TimezoneID)
	}

	return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmSqlMiServerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Sql.ManagedInstancesClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.ManagedInstanceID(d.Id())
	if err != nil {
		return err
	}

	future, err := client.Delete(ctx, id.ResourceGroup, id.Name)
	if err != nil {
		return fmt.Errorf("deleting SQL Managed Instance %q: %+v", id.ID(), err)
	}

	return future.WaitForCompletionRef(ctx, client.Client)
}

func expandManagedInstanceSkuName(skuName string) (*sql.Sku, error) {
	parts := strings.Split(skuName, "_")
	if len(parts) != 2 {
		return nil, fmt.Errorf("sku_name (%s) has the wrong number of parts (%d) after splitting on _", skuName, len(parts))
	}

	var tier string
	switch parts[0] {
	case "GP":
		tier = "GeneralPurpose"
	case "BC":
		tier = "BusinessCritical"
	default:
		return nil, fmt.Errorf("sku_name %s has unknown sku tier %s", skuName, parts[0])
	}

	return &sql.Sku{
		Name:   utils.String(skuName),
		Tier:   utils.String(tier),
		Family: utils.String(parts[1]),
	}, nil
}
