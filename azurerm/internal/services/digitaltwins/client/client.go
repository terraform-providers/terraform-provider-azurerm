package client

import (
	"github.com/Azure/azure-sdk-for-go/services/digitaltwins/mgmt/2020-10-31/digitaltwins"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/common"
)

type Client struct {
	DigitalTwinsClient *digitaltwins.Client
	EndpointClient     *digitaltwins.EndpointClient
}

func NewClient(o *common.ClientOptions) *Client {
	digitalTwinsClient := digitaltwins.NewClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&digitalTwinsClient.Client, o.ResourceManagerAuthorizer)

	endpointClient := digitaltwins.NewEndpointClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&endpointClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		DigitalTwinsClient: &digitalTwinsClient,
		EndpointClient:     &endpointClient,
	}
}
