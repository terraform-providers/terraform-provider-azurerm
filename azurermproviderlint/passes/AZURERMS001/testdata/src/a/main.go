package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/suppress"
)

func f() {
	/* Passing case */
	_ = schema.Schema{
		ValidateFunc:     validation.StringInSlice([]string{}, true),
		DiffSuppressFunc: suppress.CaseDifference,
	}

	_ = schema.Schema{
		ValidateFunc: validation.StringInSlice([]string{}, false),
	}

	/* Failing cases */
	_ = schema.Schema{ // want "AZURERMS001: prefer adding `DiffSuppressFunc: suppress.CaseDifference` when ignoring case during validation"
		ValidateFunc: validation.StringInSlice([]string{}, true),
	}
	_ = schema.Schema{ // want "AZURERMS001: prefer adding `DiffSuppressFunc: suppress.CaseDifference` when ignoring case during validation"
		ValidateFunc:     validation.StringInSlice([]string{}, true),
		DiffSuppressFunc: suppress.RFC3339Time,
	}

	/* Comment ignored cases */

	// lintignore:AZURERMS001
	_ = schema.Schema{
		ValidateFunc: validation.StringInSlice([]string{}, true),
	}
	// lintignore:AZURERMS001
	_ = schema.Schema{
		ValidateFunc:     validation.StringInSlice([]string{}, true),
		DiffSuppressFunc: suppress.RFC3339Time,
	}
}
