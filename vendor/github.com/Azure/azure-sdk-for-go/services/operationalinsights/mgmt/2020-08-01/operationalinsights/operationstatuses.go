package operationalinsights

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// OperationStatusesClient is the operational Insights Client
type OperationStatusesClient struct {
	BaseClient
}

// NewOperationStatusesClient creates an instance of the OperationStatusesClient client.
func NewOperationStatusesClient(subscriptionID string) OperationStatusesClient {
	return NewOperationStatusesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOperationStatusesClientWithBaseURI creates an instance of the OperationStatusesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewOperationStatusesClientWithBaseURI(baseURI string, subscriptionID string) OperationStatusesClient {
	return OperationStatusesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get the status of a long running azure asynchronous operation.
// Parameters:
// location - the region name of operation.
// asyncOperationID - the operation Id.
func (client OperationStatusesClient) Get(ctx context.Context, location string, asyncOperationID string) (result OperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationStatusesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.OperationStatusesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, location, asyncOperationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.OperationStatusesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationalinsights.OperationStatusesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.OperationStatusesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client OperationStatusesClient) GetPreparer(ctx context.Context, location string, asyncOperationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"asyncOperationId": autorest.Encode("path", asyncOperationID),
		"location":         autorest.Encode("path", location),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.OperationalInsights/locations/{location}/operationStatuses/{asyncOperationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OperationStatusesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OperationStatusesClient) GetResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
