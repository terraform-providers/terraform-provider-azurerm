package azurerm

import (
	"bytes"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

func passwordCredentialsSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"key_id": {
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validateUUID,
				},

				"start_date": {
					Type:             schema.TypeString,
					Optional:         true,
					Computed:         true,
					DiffSuppressFunc: compareDataAsUTCSuppressFunc,
					ValidateFunc:     validateRFC3339Date,
				},

				"end_date": {
					Type:             schema.TypeString,
					Required:         true,
					DiffSuppressFunc: compareDataAsUTCSuppressFunc,
					ValidateFunc:     validateRFC3339Date,
				},

				"value": {
					Type:      schema.TypeString,
					Required:  true,
					Sensitive: true,
				},
			},
		},
		Set: resourcePasswordCredentialHash,
	}
}

func resourcePasswordCredentialHash(v interface{}) int {
	var buf bytes.Buffer

	if v != nil {
		m := v.(map[string]interface{})
		buf.WriteString(fmt.Sprintf("%s-", m["key_id"].(string)))
	}

	return hashcode.String(buf.String())
}

func flattenAzureRmPasswordCredential(cred *graphrbac.PasswordCredential) interface{} {
	l := make(map[string]interface{})
	l["key_id"] = *cred.KeyID
	l["start_date"] = string((*cred.StartDate).Format(time.RFC3339))
	l["end_date"] = string((*cred.EndDate).Format(time.RFC3339))

	return l
}

func flattenAzureRmPasswordCredentials(creds *[]graphrbac.PasswordCredential) []interface{} {
	result := make([]interface{}, 0, len(*creds))
	for _, cred := range *creds {
		result = append(result, flattenAzureRmPasswordCredential(&cred))
	}
	return result
}

func expandAzureRmPasswordCredential(d *map[string]interface{}) (*graphrbac.PasswordCredential, error) {
	keyId := (*d)["key_id"].(string)
	startDate := (*d)["start_date"].(string)
	endDate := (*d)["end_date"].(string)
	value := (*d)["value"].(string)

	pc := graphrbac.PasswordCredential{
		KeyID: &keyId,
		Value: &value,
	}

	if startDate != "" {
		starttime, sterr := time.Parse(time.RFC3339, startDate)
		if sterr != nil {
			return nil, fmt.Errorf("Cannot parse start_date: %q", startDate)
		}
		stdt := date.Time{Time: starttime}
		pc.StartDate = &stdt
	}

	if endDate != "" {
		endtime, eterr := time.Parse(time.RFC3339, endDate)
		if eterr != nil {
			return nil, fmt.Errorf("Cannot parse end_date: %q", endDate)
		}
		etdt := date.Time{Time: endtime}
		pc.EndDate = &etdt
	}

	return &pc, nil
}

func expandAzureRmPasswordCredentials(d *schema.ResourceData, o *schema.Set) (*[]graphrbac.PasswordCredential, error) {
	creds := d.Get("password_credential").(*schema.Set).List()
	passCreds := make([]graphrbac.PasswordCredential, 0, len(creds))

	for _, v := range creds {
		cfg := v.(map[string]interface{})
		cred, err := expandAzureRmPasswordCredential(&cfg)
		if err != nil {
			return nil, err
		}

		// Azure only allows an in-place update of the Password Credentials list.
		// Existing keys, matched by their KeyID, must be sent back with their
		// Value attribute set to nil. New keys need to provide a Value.
		// By referencing the existing schema (o), we can determine which
		// entries in the list are existing keys.
		if o != nil && o.Contains(v) {
			cred.Value = nil
		}

		passCreds = append(passCreds, *cred)
	}

	return &passCreds, nil
}
