package validate

import (
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/validation"
)

func NodeDefinitionVMSize() pluginsdk.SchemaValidateFunc {
	return validation.StringInSlice([]string{
		// short of deploying every VM Sku for every node type for every HDInsight Cluster
		// this is the list I've (@tombuildsstuff) found for valid SKU's from an endpoint in the Portal
		// using another SKU causes a bad request from the API - as such this is a best effort UX
		"ExtraSmall",
		"Small",
		"Medium",
		"Large",
		"ExtraLarge",
		"A5",
		"A6",
		"A7",
		"A8",
		"A9",
		"A10",
		"A11",
		"Standard_A1_V2",
		"Standard_A2_V2",
		"Standard_A2m_V2",
		"Standard_A3",
		"Standard_A4_V2",
		"Standard_A4m_V2",
		"Standard_A8_V2",
		"Standard_A8m_V2",
		"Standard_D1",
		"Standard_D2",
		"Standard_D3",
		"Standard_D4",
		"Standard_D11",
		"Standard_D12",
		"Standard_D13",
		"Standard_D14",
		"Standard_D1_V2",
		"Standard_D2_V2",
		"Standard_D3_V2",
		"Standard_D4_V2",
		"Standard_D5_V2",
		"Standard_D11_V2",
		"Standard_D12_V2",
		"Standard_D13_V2",
		"Standard_D14_V2",
		"Standard_DS1_V2",
		"Standard_DS2_V2",
		"Standard_DS3_V2",
		"Standard_DS4_V2",
		"Standard_DS5_V2",
		"Standard_DS11_V2",
		"Standard_DS12_V2",
		"Standard_DS13_V2",
		"Standard_DS14_V2",
		"Standard_D4a_V4",
		"Standard_E2_V3",
		"Standard_E4_V3",
		"Standard_E8_V3",
		"Standard_E16_V3",
		"Standard_E20_V3",
		"Standard_E32_V3",
		"Standard_E64_V3",
		"Standard_E64i_V3",
		"Standard_E2s_V3",
		"Standard_E4s_V3",
		"Standard_E8s_V3",
		"Standard_E16s_V3",
		"Standard_E20s_V3",
		"Standard_E32s_V3",
		"Standard_E64s_V3",
		"Standard_E64is_V3",
		"Standard_G1",
		"Standard_G2",
		"Standard_G3",
		"Standard_G4",
		"Standard_G5",
		"Standard_F2s_V2",
		"Standard_F4s_V2",
		"Standard_F8s_V2",
		"Standard_F16s_V2",
		"Standard_F32s_V2",
		"Standard_F64s_V2",
		"Standard_F72s_V2",
		"Standard_GS1",
		"Standard_GS2",
		"Standard_GS3",
		"Standard_GS4",
		"Standard_GS5",
		"Standard_NC24",
	}, true)
}
