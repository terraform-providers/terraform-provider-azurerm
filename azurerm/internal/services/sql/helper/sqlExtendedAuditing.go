package helper

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-03-01-preview/sql"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func BlobExtendedAuditingSchemaFrom(s map[string]*schema.Schema) map[string]*schema.Schema {
	blobAuditing := map[string]*schema.Schema{
		"blob_extended_auditing_policy": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"storage_account_access_key": {
						Type:         schema.TypeString,
						Required:     true,
						Sensitive:    true,
						ValidateFunc: validate.NoEmptyStrings,
					},

					"storage_endpoint": {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validate.URLIsHTTPS,
					},

					"storage_secondary_key_enabled": {
						Type:     schema.TypeBool,
						Optional: true,
					},

					"retention_in_days": {
						Type:         schema.TypeInt,
						Optional:     true,
						ValidateFunc: validation.IntBetween(0, 3285),
					},
				},
			},
		},
	}
	return azure.MergeSchema(s, blobAuditing)
}

func ExpandAzureRmSqlServerBlobAuditingPolicies(input []interface{}) *sql.ExtendedServerBlobAuditingPolicyProperties {
	if len(input) == 0 {
		return &sql.ExtendedServerBlobAuditingPolicyProperties{
			State: sql.BlobAuditingPolicyStateDisabled,
		}
	}
	serverBlobAuditingPolicies := input[0].(map[string]interface{})

	ExtendedServerBlobAuditingPolicyProperties := sql.ExtendedServerBlobAuditingPolicyProperties{
		State:                   sql.BlobAuditingPolicyStateEnabled,
		StorageAccountAccessKey: utils.String(serverBlobAuditingPolicies["storage_account_access_key"].(string)),
		StorageEndpoint:         utils.String(serverBlobAuditingPolicies["storage_endpoint"].(string)),
	}
	if v, ok := serverBlobAuditingPolicies["storage_secondary_key_enabled"]; ok {
		ExtendedServerBlobAuditingPolicyProperties.IsStorageSecondaryKeyInUse = utils.Bool(v.(bool))
	}
	if v, ok := serverBlobAuditingPolicies["retention_in_days"]; ok {
		ExtendedServerBlobAuditingPolicyProperties.RetentionDays = utils.Int32(int32(v.(int)))
	}

	return &ExtendedServerBlobAuditingPolicyProperties
}

func FlattenAzureRmSqlServerBlobAuditingPolicies(extendedServerBlobAuditingPolicy *sql.ExtendedServerBlobAuditingPolicy, d *schema.ResourceData) []interface{} {
	if extendedServerBlobAuditingPolicy == nil || extendedServerBlobAuditingPolicy.State == sql.BlobAuditingPolicyStateDisabled {
		return []interface{}{}
	}
	var storageEndpoint, storageAccessKey string
	// storage_account_access_key will not be returned, so we transfer the schema value
	if v, ok := d.GetOk("blob_extended_auditing_policy.0.storage_account_access_key"); ok {
		storageAccessKey = v.(string)
	}
	if extendedServerBlobAuditingPolicy.StorageEndpoint != nil {
		storageEndpoint = *extendedServerBlobAuditingPolicy.StorageEndpoint
	}

	var secondKeyInUse bool
	if extendedServerBlobAuditingPolicy.IsStorageSecondaryKeyInUse != nil {
		secondKeyInUse = *extendedServerBlobAuditingPolicy.IsStorageSecondaryKeyInUse
	}
	var retentionDays int32
	if extendedServerBlobAuditingPolicy.RetentionDays != nil {
		retentionDays = *extendedServerBlobAuditingPolicy.RetentionDays
	}

	return []interface{}{
		map[string]interface{}{
			"storage_account_access_key":    storageAccessKey,
			"storage_endpoint":              storageEndpoint,
			"storage_secondary_key_enabled": secondKeyInUse,
			"retention_in_days":             retentionDays,
		},
	}
}
