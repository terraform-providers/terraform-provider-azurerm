package validate

import (
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/databoxedge/mgmt/2020-12-01/databoxedge"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

func DataboxEdgeDeviceSkuName(v interface{}, k string) (warnings []string, errors []error) {
	validSku := false
	validTier := false

	value, ok := v.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %q to be string", k))
		return warnings, errors
	}

	skuParts := strings.Split(value, "-")
	validSkus := getValidSkus()
	validTiers := getValidTiers()

	// Validate the SKU Name section
	for _, str := range validSkus {
		if skuParts[0] == str {
			validSku = true
			break
		}
	}

	if len(skuParts) > 1 {
		// Validate the SKU Tier section
		for _, str := range validTiers {
			if skuParts[1] == str {
				validTier = true
				break
			}
		}
	}

	if !validSku {
		errors = append(errors, fmt.Errorf("expected %q %q segment to be one of [%s], got %q", k, "name", azure.QuotedStringSlice(validSkus), value))
	}
	if !validTier {
		errors = append(errors, fmt.Errorf("expected %q %q segment to be one of [%s], got %q", k, "tier", azure.QuotedStringSlice(validTiers), value))
	}

	return warnings, errors
}

func getValidSkus() []string {
	return []string{
		string(databoxedge.Gateway),
		string(databoxedge.Edge),
		string(databoxedge.TEA1Node),
		string(databoxedge.TEA1NodeUPS),
		string(databoxedge.TEA1NodeHeater),
		string(databoxedge.TEA1NodeUPSHeater),
		string(databoxedge.TEA4NodeHeater),
		string(databoxedge.TEA4NodeUPSHeater),
		string(databoxedge.TMA),
	}
}

func getValidTiers() []string {
	return []string{
		string(databoxedge.Standard),
	}
}
