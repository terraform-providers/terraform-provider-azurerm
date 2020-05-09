package parse

import (
	"fmt"
	"strings"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type EventGridEventSubscriptionId struct {
	Scope string
	Name  string
}

func EventGridEventSubscriptionID(input string) (*EventGridEventSubscriptionId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse EventGrid Event Subscription ID %q: %+v", input, err)
	}

	segments := strings.Split(input, "/providers/Microsoft.EventGrid/eventSubscriptions/")
	if len(segments) != 2 {
		return nil, fmt.Errorf("Expected ID to be in the format `{scope}/providers/Microsoft.EventGrid/eventSubscriptions/{name} - got %d segments", len(segments))
	}

	eventSubscription := EventGridEventSubscriptionId{
		Scope: segments[0],
	}

	if eventSubscription.Name, err = id.PopSegment("eventSubscriptions"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &eventSubscription, nil
}
