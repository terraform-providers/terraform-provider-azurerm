package azure

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/suppress"
)

func SchemaResourceGroupName() *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: ValidateResourceGroupName,
	}
}

func SchemaResourceGroupNameDeprecated() *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: ValidateResourceGroupName,
		Deprecated:   "This field is no longer used and will be removed in the next major version of the Azure Provider",
	}
}

func SchemaResourceGroupNameDiffSuppress() *schema.Schema {
	return &schema.Schema{
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: suppress.CaseDifference,
		ValidateFunc:     ValidateResourceGroupName,
	}
}

func SchemaResourceGroupNameForDataSource() *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: ValidateResourceGroupName,
	}
}

func SchemaResourceGroupNameOptionalComputed() *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		ForceNew:     true,
		Optional:     true,
		Computed:     true,
		ValidateFunc: ValidateResourceGroupName,
	}
}

func SchemaResourceGroupNameSetOptional() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: ValidateResourceGroupName,
		},
	}
}

func ValidateResourceGroupName(v interface{}, k string) (warnings []string, errors []error) {
	value := v.(string)

	if len(value) > 90 {
		errors = append(errors, fmt.Errorf("%q may not exceed 90 characters in length", k))
	}

	if strings.HasSuffix(value, ".") {
		errors = append(errors, fmt.Errorf("%q may not end with a period", k))
	}

	// regex pulled from https://docs.microsoft.com/en-us/rest/api/resources/resourcegroups/createorupdate
	if matched := regexp.MustCompile(`^[-\w._()]+$`).Match([]byte(value)); !matched {
		errors = append(errors, fmt.Errorf("%q may only contain alphanumeric characters, dash, underscores, parentheses and periods", k))
	}

	return warnings, errors
}
