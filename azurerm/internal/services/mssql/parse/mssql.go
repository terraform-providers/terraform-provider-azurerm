package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type MssqlVmId struct {
	ResourceGroup string
	Name          string
}

type MssqlVmGroupId struct {
	ResourceGroup string
	Name          string
}

func MssqlVmID(input string) (*MssqlVmId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Microsoft Sql VM ID %q: %+v", input, err)
	}

	sqlvm := MssqlVmId{
		ResourceGroup: id.ResourceGroup,
	}

	if sqlvm.Name, err = id.PopSegment("sqlVirtualMachines"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &sqlvm, nil
}

func MssqlVmGroupID(input string) (*MssqlVmGroupId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Microsoft Sql VM Group ID %q: %+v", input, err)
	}

	sqlvmGroup := MssqlVmGroupId{
		ResourceGroup: id.ResourceGroup,
	}

	if sqlvmGroup.Name, err = id.PopSegment("sqlVirtualMachineGroups"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &sqlvmGroup, nil
}
