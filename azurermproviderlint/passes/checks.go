package passes

import (
	"github.com/terraform-providers/terraform-provider-azurerm/azurermproviderlint/passes/AZRMR001"
	"github.com/terraform-providers/terraform-provider-azurerm/azurermproviderlint/passes/AZRMS001"
	"golang.org/x/tools/go/analysis"
)

var AllChecks = []*analysis.Analyzer{
	AZRMR001.Analyzer,
	AZRMS001.Analyzer,
}
