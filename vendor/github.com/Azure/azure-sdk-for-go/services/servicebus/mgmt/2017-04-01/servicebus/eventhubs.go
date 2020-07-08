package servicebus

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
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

// EventHubsClient is the azure Service Bus client
type EventHubsClient struct {
	BaseClient
}

// NewEventHubsClient creates an instance of the EventHubsClient client.
func NewEventHubsClient(subscriptionID string) EventHubsClient {
	return NewEventHubsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEventHubsClientWithBaseURI creates an instance of the EventHubsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewEventHubsClientWithBaseURI(baseURI string, subscriptionID string) EventHubsClient {
	return EventHubsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByNamespace gets all the Event Hubs in a service bus Namespace.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// namespaceName - the namespace name
func (client EventHubsClient) ListByNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (result EventHubListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventHubsClient.ListByNamespace")
		defer func() {
			sc := -1
			if result.ehlr.Response.Response != nil {
				sc = result.ehlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicebus.EventHubsClient", "ListByNamespace", err.Error())
	}

	result.fn = client.listByNamespaceNextResults
	req, err := client.ListByNamespacePreparer(ctx, resourceGroupName, namespaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.EventHubsClient", "ListByNamespace", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByNamespaceSender(req)
	if err != nil {
		result.ehlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicebus.EventHubsClient", "ListByNamespace", resp, "Failure sending request")
		return
	}

	result.ehlr, err = client.ListByNamespaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.EventHubsClient", "ListByNamespace", resp, "Failure responding to request")
	}

	return
}

// ListByNamespacePreparer prepares the ListByNamespace request.
func (client EventHubsClient) ListByNamespacePreparer(ctx context.Context, resourceGroupName string, namespaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/eventhubs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByNamespaceSender sends the ListByNamespace request. The method will close the
// http.Response Body if it receives an error.
func (client EventHubsClient) ListByNamespaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByNamespaceResponder handles the response to the ListByNamespace request. The method always
// closes the http.Response Body.
func (client EventHubsClient) ListByNamespaceResponder(resp *http.Response) (result EventHubListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByNamespaceNextResults retrieves the next set of results, if any.
func (client EventHubsClient) listByNamespaceNextResults(ctx context.Context, lastResults EventHubListResult) (result EventHubListResult, err error) {
	req, err := lastResults.eventHubListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicebus.EventHubsClient", "listByNamespaceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByNamespaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicebus.EventHubsClient", "listByNamespaceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByNamespaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.EventHubsClient", "listByNamespaceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByNamespaceComplete enumerates all values, automatically crossing page boundaries as required.
func (client EventHubsClient) ListByNamespaceComplete(ctx context.Context, resourceGroupName string, namespaceName string) (result EventHubListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventHubsClient.ListByNamespace")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByNamespace(ctx, resourceGroupName, namespaceName)
	return
}
