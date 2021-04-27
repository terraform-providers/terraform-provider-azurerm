package backup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ProtectedItemsGroupClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ProtectedItemsGroupClient struct {
	BaseClient
}

// NewProtectedItemsGroupClient creates an instance of the ProtectedItemsGroupClient client.
func NewProtectedItemsGroupClient(subscriptionID string) ProtectedItemsGroupClient {
	return NewProtectedItemsGroupClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProtectedItemsGroupClientWithBaseURI creates an instance of the ProtectedItemsGroupClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewProtectedItemsGroupClientWithBaseURI(baseURI string, subscriptionID string) ProtectedItemsGroupClient {
	return ProtectedItemsGroupClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List provides a pageable list of all items that are backed up within a vault.
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// filter - oData filter options.
// skipToken - skipToken Filter.
func (client ProtectedItemsGroupClient) List(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result ProtectedItemResourceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProtectedItemsGroupClient.List")
		defer func() {
			sc := -1
			if result.pirl.Response.Response != nil {
				sc = result.pirl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, vaultName, resourceGroupName, filter, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectedItemsGroupClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pirl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.ProtectedItemsGroupClient", "List", resp, "Failure sending request")
		return
	}

	result.pirl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectedItemsGroupClient", "List", resp, "Failure responding to request")
		return
	}
	if result.pirl.hasNextLink() && result.pirl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ProtectedItemsGroupClient) ListPreparer(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2019-05-13"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupProtectedItems", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectedItemsGroupClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProtectedItemsGroupClient) ListResponder(resp *http.Response) (result ProtectedItemResourceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ProtectedItemsGroupClient) listNextResults(ctx context.Context, lastResults ProtectedItemResourceList) (result ProtectedItemResourceList, err error) {
	req, err := lastResults.protectedItemResourceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "backup.ProtectedItemsGroupClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "backup.ProtectedItemsGroupClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectedItemsGroupClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProtectedItemsGroupClient) ListComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result ProtectedItemResourceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProtectedItemsGroupClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, vaultName, resourceGroupName, filter, skipToken)
	return
}
