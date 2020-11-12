package datashare

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ProviderShareSubscriptionsClient is the creates a Microsoft.DataShare management client.
type ProviderShareSubscriptionsClient struct {
	BaseClient
}

// NewProviderShareSubscriptionsClient creates an instance of the ProviderShareSubscriptionsClient client.
func NewProviderShareSubscriptionsClient(subscriptionID string) ProviderShareSubscriptionsClient {
	return NewProviderShareSubscriptionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProviderShareSubscriptionsClientWithBaseURI creates an instance of the ProviderShareSubscriptionsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewProviderShareSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) ProviderShareSubscriptionsClient {
	return ProviderShareSubscriptionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetByShare get share subscription in a provider share
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareName - the name of the share.
// providerShareSubscriptionID - to locate shareSubscription
func (client ProviderShareSubscriptionsClient) GetByShare(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string) (result ProviderShareSubscription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderShareSubscriptionsClient.GetByShare")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetBySharePreparer(ctx, resourceGroupName, accountName, shareName, providerShareSubscriptionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.ProviderShareSubscriptionsClient", "GetByShare", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByShareSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.ProviderShareSubscriptionsClient", "GetByShare", resp, "Failure sending request")
		return
	}

	result, err = client.GetByShareResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.ProviderShareSubscriptionsClient", "GetByShare", resp, "Failure responding to request")
	}

	return
}

// GetBySharePreparer prepares the GetByShare request.
func (client ProviderShareSubscriptionsClient) GetBySharePreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":                 autorest.Encode("path", accountName),
		"providerShareSubscriptionId": autorest.Encode("path", providerShareSubscriptionID),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"shareName":                   autorest.Encode("path", shareName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/providerShareSubscriptions/{providerShareSubscriptionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByShareSender sends the GetByShare request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderShareSubscriptionsClient) GetByShareSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetByShareResponder handles the response to the GetByShare request. The method always
// closes the http.Response Body.
func (client ProviderShareSubscriptionsClient) GetByShareResponder(resp *http.Response) (result ProviderShareSubscription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByShare list share subscriptions in a provider share
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareName - the name of the share.
// skipToken - continuation Token
func (client ProviderShareSubscriptionsClient) ListByShare(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string) (result ProviderShareSubscriptionListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderShareSubscriptionsClient.ListByShare")
		defer func() {
			sc := -1
			if result.pssl.Response.Response != nil {
				sc = result.pssl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByShareNextResults
	req, err := client.ListBySharePreparer(ctx, resourceGroupName, accountName, shareName, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.ProviderShareSubscriptionsClient", "ListByShare", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByShareSender(req)
	if err != nil {
		result.pssl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.ProviderShareSubscriptionsClient", "ListByShare", resp, "Failure sending request")
		return
	}

	result.pssl, err = client.ListByShareResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.ProviderShareSubscriptionsClient", "ListByShare", resp, "Failure responding to request")
	}
	if result.pssl.hasNextLink() && result.pssl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListBySharePreparer prepares the ListByShare request.
func (client ProviderShareSubscriptionsClient) ListBySharePreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"shareName":         autorest.Encode("path", shareName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/providerShareSubscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByShareSender sends the ListByShare request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderShareSubscriptionsClient) ListByShareSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByShareResponder handles the response to the ListByShare request. The method always
// closes the http.Response Body.
func (client ProviderShareSubscriptionsClient) ListByShareResponder(resp *http.Response) (result ProviderShareSubscriptionList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByShareNextResults retrieves the next set of results, if any.
func (client ProviderShareSubscriptionsClient) listByShareNextResults(ctx context.Context, lastResults ProviderShareSubscriptionList) (result ProviderShareSubscriptionList, err error) {
	req, err := lastResults.providerShareSubscriptionListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datashare.ProviderShareSubscriptionsClient", "listByShareNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByShareSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datashare.ProviderShareSubscriptionsClient", "listByShareNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByShareResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.ProviderShareSubscriptionsClient", "listByShareNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByShareComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProviderShareSubscriptionsClient) ListByShareComplete(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string) (result ProviderShareSubscriptionListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderShareSubscriptionsClient.ListByShare")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByShare(ctx, resourceGroupName, accountName, shareName, skipToken)
	return
}

// Reinstate reinstate share subscription in a provider share
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareName - the name of the share.
// providerShareSubscriptionID - to locate shareSubscription
func (client ProviderShareSubscriptionsClient) Reinstate(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string) (result ProviderShareSubscription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderShareSubscriptionsClient.Reinstate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ReinstatePreparer(ctx, resourceGroupName, accountName, shareName, providerShareSubscriptionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.ProviderShareSubscriptionsClient", "Reinstate", nil, "Failure preparing request")
		return
	}

	resp, err := client.ReinstateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.ProviderShareSubscriptionsClient", "Reinstate", resp, "Failure sending request")
		return
	}

	result, err = client.ReinstateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.ProviderShareSubscriptionsClient", "Reinstate", resp, "Failure responding to request")
	}

	return
}

// ReinstatePreparer prepares the Reinstate request.
func (client ProviderShareSubscriptionsClient) ReinstatePreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":                 autorest.Encode("path", accountName),
		"providerShareSubscriptionId": autorest.Encode("path", providerShareSubscriptionID),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"shareName":                   autorest.Encode("path", shareName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/providerShareSubscriptions/{providerShareSubscriptionId}/reinstate", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ReinstateSender sends the Reinstate request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderShareSubscriptionsClient) ReinstateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ReinstateResponder handles the response to the Reinstate request. The method always
// closes the http.Response Body.
func (client ProviderShareSubscriptionsClient) ReinstateResponder(resp *http.Response) (result ProviderShareSubscription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Revoke revoke share subscription in a provider share
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareName - the name of the share.
// providerShareSubscriptionID - to locate shareSubscription
func (client ProviderShareSubscriptionsClient) Revoke(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string) (result ProviderShareSubscriptionsRevokeFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderShareSubscriptionsClient.Revoke")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RevokePreparer(ctx, resourceGroupName, accountName, shareName, providerShareSubscriptionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.ProviderShareSubscriptionsClient", "Revoke", nil, "Failure preparing request")
		return
	}

	result, err = client.RevokeSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.ProviderShareSubscriptionsClient", "Revoke", result.Response(), "Failure sending request")
		return
	}

	return
}

// RevokePreparer prepares the Revoke request.
func (client ProviderShareSubscriptionsClient) RevokePreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, providerShareSubscriptionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":                 autorest.Encode("path", accountName),
		"providerShareSubscriptionId": autorest.Encode("path", providerShareSubscriptionID),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"shareName":                   autorest.Encode("path", shareName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/providerShareSubscriptions/{providerShareSubscriptionId}/revoke", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RevokeSender sends the Revoke request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderShareSubscriptionsClient) RevokeSender(req *http.Request) (future ProviderShareSubscriptionsRevokeFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// RevokeResponder handles the response to the Revoke request. The method always
// closes the http.Response Body.
func (client ProviderShareSubscriptionsClient) RevokeResponder(resp *http.Response) (result ProviderShareSubscription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
