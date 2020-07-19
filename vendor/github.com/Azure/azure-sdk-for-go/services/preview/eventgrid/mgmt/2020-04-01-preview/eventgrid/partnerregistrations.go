package eventgrid

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

// PartnerRegistrationsClient is the azure EventGrid Management Client
type PartnerRegistrationsClient struct {
	BaseClient
}

// NewPartnerRegistrationsClient creates an instance of the PartnerRegistrationsClient client.
func NewPartnerRegistrationsClient(subscriptionID string) PartnerRegistrationsClient {
	return NewPartnerRegistrationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPartnerRegistrationsClientWithBaseURI creates an instance of the PartnerRegistrationsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewPartnerRegistrationsClientWithBaseURI(baseURI string, subscriptionID string) PartnerRegistrationsClient {
	return PartnerRegistrationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new partner registration with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerRegistrationName - name of the partner registration.
// partnerRegistrationInfo - partnerRegistration information.
func (client PartnerRegistrationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, partnerRegistrationName string, partnerRegistrationInfo PartnerRegistration) (result PartnerRegistration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerRegistrationsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, partnerRegistrationName, partnerRegistrationInfo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PartnerRegistrationsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, partnerRegistrationName string, partnerRegistrationInfo PartnerRegistration) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"partnerRegistrationName": autorest.Encode("path", partnerRegistrationName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}", pathParameters),
		autorest.WithJSON(partnerRegistrationInfo),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerRegistrationsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client PartnerRegistrationsClient) CreateOrUpdateResponder(resp *http.Response) (result PartnerRegistration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a partner registration with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerRegistrationName - name of the partner registration.
func (client PartnerRegistrationsClient) Delete(ctx context.Context, resourceGroupName string, partnerRegistrationName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerRegistrationsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, partnerRegistrationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PartnerRegistrationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, partnerRegistrationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"partnerRegistrationName": autorest.Encode("path", partnerRegistrationName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerRegistrationsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PartnerRegistrationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a partner registration with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerRegistrationName - name of the partner registration.
func (client PartnerRegistrationsClient) Get(ctx context.Context, resourceGroupName string, partnerRegistrationName string) (result PartnerRegistration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerRegistrationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, partnerRegistrationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PartnerRegistrationsClient) GetPreparer(ctx context.Context, resourceGroupName string, partnerRegistrationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"partnerRegistrationName": autorest.Encode("path", partnerRegistrationName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerRegistrationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PartnerRegistrationsClient) GetResponder(resp *http.Response) (result PartnerRegistration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all partners registrations.
func (client PartnerRegistrationsClient) List(ctx context.Context) (result PartnerRegistrationsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerRegistrationsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client PartnerRegistrationsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.EventGrid/partnerRegistrations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerRegistrationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PartnerRegistrationsClient) ListResponder(resp *http.Response) (result PartnerRegistrationsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup list all the partner registrations under a resource group.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// filter - the query used to filter the search results using OData syntax. Filtering is permitted on the
// 'name' property only and with limited number of OData operations. These operations are: the 'contains'
// function as well as the following logical operations: not, and, or, eq (for equal), and ne (for not equal).
// No arithmetic operations are supported. The following is a valid filter example: $filter=contains(namE,
// 'PATTERN') and name ne 'PATTERN-1'. The following is not a valid filter example: $filter=location eq
// 'westus'.
// top - the number of results to return per page for the list operation. Valid range for top parameter is 1 to
// 100. If not specified, the default number of results to be returned is 20 items per page.
func (client PartnerRegistrationsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32) (result PartnerRegistrationsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerRegistrationsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.prlr.Response.Response != nil {
				sc = result.prlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, filter, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.prlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.prlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client PartnerRegistrationsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerRegistrationsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client PartnerRegistrationsClient) ListByResourceGroupResponder(resp *http.Response) (result PartnerRegistrationsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client PartnerRegistrationsClient) listByResourceGroupNextResults(ctx context.Context, lastResults PartnerRegistrationsListResult) (result PartnerRegistrationsListResult, err error) {
	req, err := lastResults.partnerRegistrationsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client PartnerRegistrationsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, top *int32) (result PartnerRegistrationsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerRegistrationsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, filter, top)
	return
}

// ListBySubscription list all the partner registrations under an Azure subscription.
// Parameters:
// filter - the query used to filter the search results using OData syntax. Filtering is permitted on the
// 'name' property only and with limited number of OData operations. These operations are: the 'contains'
// function as well as the following logical operations: not, and, or, eq (for equal), and ne (for not equal).
// No arithmetic operations are supported. The following is a valid filter example: $filter=contains(namE,
// 'PATTERN') and name ne 'PATTERN-1'. The following is not a valid filter example: $filter=location eq
// 'westus'.
// top - the number of results to return per page for the list operation. Valid range for top parameter is 1 to
// 100. If not specified, the default number of results to be returned is 20 items per page.
func (client PartnerRegistrationsClient) ListBySubscription(ctx context.Context, filter string, top *int32) (result PartnerRegistrationsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerRegistrationsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.prlr.Response.Response != nil {
				sc = result.prlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, filter, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.prlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.prlr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client PartnerRegistrationsClient) ListBySubscriptionPreparer(ctx context.Context, filter string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/partnerRegistrations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerRegistrationsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client PartnerRegistrationsClient) ListBySubscriptionResponder(resp *http.Response) (result PartnerRegistrationsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client PartnerRegistrationsClient) listBySubscriptionNextResults(ctx context.Context, lastResults PartnerRegistrationsListResult) (result PartnerRegistrationsListResult, err error) {
	req, err := lastResults.partnerRegistrationsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client PartnerRegistrationsClient) ListBySubscriptionComplete(ctx context.Context, filter string, top *int32) (result PartnerRegistrationsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerRegistrationsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, filter, top)
	return
}

// Update updates a partner registration with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerRegistrationName - name of the partner registration.
// partnerRegistrationUpdateParameters - partner registration update information.
func (client PartnerRegistrationsClient) Update(ctx context.Context, resourceGroupName string, partnerRegistrationName string, partnerRegistrationUpdateParameters PartnerRegistrationUpdateParameters) (result PartnerRegistration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerRegistrationsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, partnerRegistrationName, partnerRegistrationUpdateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerRegistrationsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client PartnerRegistrationsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, partnerRegistrationName string, partnerRegistrationUpdateParameters PartnerRegistrationUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"partnerRegistrationName": autorest.Encode("path", partnerRegistrationName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}", pathParameters),
		autorest.WithJSON(partnerRegistrationUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerRegistrationsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client PartnerRegistrationsClient) UpdateResponder(resp *http.Response) (result PartnerRegistration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
