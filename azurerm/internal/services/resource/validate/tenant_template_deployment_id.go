package validate

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/parse"
)

func TenantTemplateDeploymentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := parse.TenantTemplateDeploymentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}
