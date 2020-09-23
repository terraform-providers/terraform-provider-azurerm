package parse

import (
	"strings"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type NetworkConnectionMonitorId struct {
	ResourceGroup string
	WatcherId     string
	WatcherName   string
	Name          string
}

func NetworkConnectionMonitorID(input string) (*NetworkConnectionMonitorId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	connectionMonitor := NetworkConnectionMonitorId{
		ResourceGroup: id.ResourceGroup,
		WatcherId:     strings.Split(input, "/connectionMonitors/")[0],
	}

	if connectionMonitor.WatcherName, err = id.PopSegment("networkWatchers"); err != nil {
		return nil, err
	}

	if connectionMonitor.Name, err = id.PopSegment("connectionMonitors"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &connectionMonitor, nil
}
