package validate

import (
	"fmt"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/datashare/parse"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func DataShareAccountName() schema.SchemaValidateFunc {
	return validation.StringMatch(
		regexp.MustCompile(`^[^<>%&:\\?/#*$^();,.\|+={}\[\]!~@]{3,90}$`), `Data share account name should have length of 3 - 90, and cannot contain <>%&:\?/#*$^();,.|+={}[]!~@.`,
	)
}

func DatashareName() schema.SchemaValidateFunc {
	return validation.StringMatch(
		regexp.MustCompile(`^\w{2,90}$`), `DataShare name can only contain alphanumeric characters and _, and must be between 2 and 90 characters long.`,
	)
}

func DatashareAccountID(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %q to be string", k))
		return warnings, errors
	}

	if _, err := parse.DataShareAccountID(v); err != nil {
		errors = append(errors, fmt.Errorf("can not parse %q as a Datashare account id: %v", k, err))
	}

	return warnings, errors
}
