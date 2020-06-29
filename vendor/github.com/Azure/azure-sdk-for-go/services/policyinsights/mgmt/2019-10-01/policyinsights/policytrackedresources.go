package policyinsights

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

// PolicyTrackedResourcesClient is the client for the PolicyTrackedResources methods of the Policyinsights service.
type PolicyTrackedResourcesClient struct {
	BaseClient
}

// NewPolicyTrackedResourcesClient creates an instance of the PolicyTrackedResourcesClient client.
func NewPolicyTrackedResourcesClient() PolicyTrackedResourcesClient {
	return NewPolicyTrackedResourcesClientWithBaseURI(DefaultBaseURI)
}

// NewPolicyTrackedResourcesClientWithBaseURI creates an instance of the PolicyTrackedResourcesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewPolicyTrackedResourcesClientWithBaseURI(baseURI string) PolicyTrackedResourcesClient {
	return PolicyTrackedResourcesClient{NewWithBaseURI(baseURI)}
}

// ListQueryResultsForManagementGroup queries policy tracked resources under the management group.
// Parameters:
// managementGroupName - management group name.
// top - maximum number of records to return.
// filter - oData filter expression.
func (client PolicyTrackedResourcesClient) ListQueryResultsForManagementGroup(ctx context.Context, managementGroupName string, top *int32, filter string) (result PolicyTrackedResourcesQueryResultsPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PolicyTrackedResourcesClient.ListQueryResultsForManagementGroup")
		defer func() {
			sc := -1
			if result.ptrqr.Response.Response != nil {
				sc = result.ptrqr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("policyinsights.PolicyTrackedResourcesClient", "ListQueryResultsForManagementGroup", err.Error())
	}

	result.fn = client.listQueryResultsForManagementGroupNextResults
	req, err := client.ListQueryResultsForManagementGroupPreparer(ctx, managementGroupName, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "ListQueryResultsForManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListQueryResultsForManagementGroupSender(req)
	if err != nil {
		result.ptrqr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "ListQueryResultsForManagementGroup", resp, "Failure sending request")
		return
	}

	result.ptrqr, err = client.ListQueryResultsForManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "ListQueryResultsForManagementGroup", resp, "Failure responding to request")
	}

	return
}

// ListQueryResultsForManagementGroupPreparer prepares the ListQueryResultsForManagementGroup request.
func (client PolicyTrackedResourcesClient) ListQueryResultsForManagementGroupPreparer(ctx context.Context, managementGroupName string, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupName":            autorest.Encode("path", managementGroupName),
		"managementGroupsNamespace":      autorest.Encode("path", "Microsoft.Management"),
		"policyTrackedResourcesResource": autorest.Encode("path", "default"),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/{managementGroupsNamespace}/managementGroups/{managementGroupName}/providers/Microsoft.PolicyInsights/policyTrackedResources/{policyTrackedResourcesResource}/queryResults", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListQueryResultsForManagementGroupSender sends the ListQueryResultsForManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyTrackedResourcesClient) ListQueryResultsForManagementGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListQueryResultsForManagementGroupResponder handles the response to the ListQueryResultsForManagementGroup request. The method always
// closes the http.Response Body.
func (client PolicyTrackedResourcesClient) ListQueryResultsForManagementGroupResponder(resp *http.Response) (result PolicyTrackedResourcesQueryResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listQueryResultsForManagementGroupNextResults retrieves the next set of results, if any.
func (client PolicyTrackedResourcesClient) listQueryResultsForManagementGroupNextResults(ctx context.Context, lastResults PolicyTrackedResourcesQueryResults) (result PolicyTrackedResourcesQueryResults, err error) {
	req, err := lastResults.policyTrackedResourcesQueryResultsPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "listQueryResultsForManagementGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListQueryResultsForManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "listQueryResultsForManagementGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListQueryResultsForManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "listQueryResultsForManagementGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListQueryResultsForManagementGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client PolicyTrackedResourcesClient) ListQueryResultsForManagementGroupComplete(ctx context.Context, managementGroupName string, top *int32, filter string) (result PolicyTrackedResourcesQueryResultsIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PolicyTrackedResourcesClient.ListQueryResultsForManagementGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListQueryResultsForManagementGroup(ctx, managementGroupName, top, filter)
	return
}

// ListQueryResultsForResource queries policy tracked resources under the resource.
// Parameters:
// resourceID - resource ID.
// top - maximum number of records to return.
// filter - oData filter expression.
func (client PolicyTrackedResourcesClient) ListQueryResultsForResource(ctx context.Context, resourceID string, top *int32, filter string) (result PolicyTrackedResourcesQueryResultsPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PolicyTrackedResourcesClient.ListQueryResultsForResource")
		defer func() {
			sc := -1
			if result.ptrqr.Response.Response != nil {
				sc = result.ptrqr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("policyinsights.PolicyTrackedResourcesClient", "ListQueryResultsForResource", err.Error())
	}

	result.fn = client.listQueryResultsForResourceNextResults
	req, err := client.ListQueryResultsForResourcePreparer(ctx, resourceID, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "ListQueryResultsForResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListQueryResultsForResourceSender(req)
	if err != nil {
		result.ptrqr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "ListQueryResultsForResource", resp, "Failure sending request")
		return
	}

	result.ptrqr, err = client.ListQueryResultsForResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "ListQueryResultsForResource", resp, "Failure responding to request")
	}

	return
}

// ListQueryResultsForResourcePreparer prepares the ListQueryResultsForResource request.
func (client PolicyTrackedResourcesClient) ListQueryResultsForResourcePreparer(ctx context.Context, resourceID string, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyTrackedResourcesResource": autorest.Encode("path", "default"),
		"resourceId":                     resourceID,
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceId}/providers/Microsoft.PolicyInsights/policyTrackedResources/{policyTrackedResourcesResource}/queryResults", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListQueryResultsForResourceSender sends the ListQueryResultsForResource request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyTrackedResourcesClient) ListQueryResultsForResourceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListQueryResultsForResourceResponder handles the response to the ListQueryResultsForResource request. The method always
// closes the http.Response Body.
func (client PolicyTrackedResourcesClient) ListQueryResultsForResourceResponder(resp *http.Response) (result PolicyTrackedResourcesQueryResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listQueryResultsForResourceNextResults retrieves the next set of results, if any.
func (client PolicyTrackedResourcesClient) listQueryResultsForResourceNextResults(ctx context.Context, lastResults PolicyTrackedResourcesQueryResults) (result PolicyTrackedResourcesQueryResults, err error) {
	req, err := lastResults.policyTrackedResourcesQueryResultsPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "listQueryResultsForResourceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListQueryResultsForResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "listQueryResultsForResourceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListQueryResultsForResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "listQueryResultsForResourceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListQueryResultsForResourceComplete enumerates all values, automatically crossing page boundaries as required.
func (client PolicyTrackedResourcesClient) ListQueryResultsForResourceComplete(ctx context.Context, resourceID string, top *int32, filter string) (result PolicyTrackedResourcesQueryResultsIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PolicyTrackedResourcesClient.ListQueryResultsForResource")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListQueryResultsForResource(ctx, resourceID, top, filter)
	return
}

// ListQueryResultsForResourceGroup queries policy tracked resources under the resource group.
// Parameters:
// resourceGroupName - resource group name.
// subscriptionID - microsoft Azure subscription ID.
// top - maximum number of records to return.
// filter - oData filter expression.
func (client PolicyTrackedResourcesClient) ListQueryResultsForResourceGroup(ctx context.Context, resourceGroupName string, subscriptionID string, top *int32, filter string) (result PolicyTrackedResourcesQueryResultsPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PolicyTrackedResourcesClient.ListQueryResultsForResourceGroup")
		defer func() {
			sc := -1
			if result.ptrqr.Response.Response != nil {
				sc = result.ptrqr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("policyinsights.PolicyTrackedResourcesClient", "ListQueryResultsForResourceGroup", err.Error())
	}

	result.fn = client.listQueryResultsForResourceGroupNextResults
	req, err := client.ListQueryResultsForResourceGroupPreparer(ctx, resourceGroupName, subscriptionID, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "ListQueryResultsForResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListQueryResultsForResourceGroupSender(req)
	if err != nil {
		result.ptrqr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "ListQueryResultsForResourceGroup", resp, "Failure sending request")
		return
	}

	result.ptrqr, err = client.ListQueryResultsForResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "ListQueryResultsForResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListQueryResultsForResourceGroupPreparer prepares the ListQueryResultsForResourceGroup request.
func (client PolicyTrackedResourcesClient) ListQueryResultsForResourceGroupPreparer(ctx context.Context, resourceGroupName string, subscriptionID string, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyTrackedResourcesResource": autorest.Encode("path", "default"),
		"resourceGroupName":              autorest.Encode("path", resourceGroupName),
		"subscriptionId":                 autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PolicyInsights/policyTrackedResources/{policyTrackedResourcesResource}/queryResults", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListQueryResultsForResourceGroupSender sends the ListQueryResultsForResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyTrackedResourcesClient) ListQueryResultsForResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListQueryResultsForResourceGroupResponder handles the response to the ListQueryResultsForResourceGroup request. The method always
// closes the http.Response Body.
func (client PolicyTrackedResourcesClient) ListQueryResultsForResourceGroupResponder(resp *http.Response) (result PolicyTrackedResourcesQueryResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listQueryResultsForResourceGroupNextResults retrieves the next set of results, if any.
func (client PolicyTrackedResourcesClient) listQueryResultsForResourceGroupNextResults(ctx context.Context, lastResults PolicyTrackedResourcesQueryResults) (result PolicyTrackedResourcesQueryResults, err error) {
	req, err := lastResults.policyTrackedResourcesQueryResultsPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "listQueryResultsForResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListQueryResultsForResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "listQueryResultsForResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListQueryResultsForResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "listQueryResultsForResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListQueryResultsForResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client PolicyTrackedResourcesClient) ListQueryResultsForResourceGroupComplete(ctx context.Context, resourceGroupName string, subscriptionID string, top *int32, filter string) (result PolicyTrackedResourcesQueryResultsIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PolicyTrackedResourcesClient.ListQueryResultsForResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListQueryResultsForResourceGroup(ctx, resourceGroupName, subscriptionID, top, filter)
	return
}

// ListQueryResultsForSubscription queries policy tracked resources under the subscription.
// Parameters:
// subscriptionID - microsoft Azure subscription ID.
// top - maximum number of records to return.
// filter - oData filter expression.
func (client PolicyTrackedResourcesClient) ListQueryResultsForSubscription(ctx context.Context, subscriptionID string, top *int32, filter string) (result PolicyTrackedResourcesQueryResultsPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PolicyTrackedResourcesClient.ListQueryResultsForSubscription")
		defer func() {
			sc := -1
			if result.ptrqr.Response.Response != nil {
				sc = result.ptrqr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("policyinsights.PolicyTrackedResourcesClient", "ListQueryResultsForSubscription", err.Error())
	}

	result.fn = client.listQueryResultsForSubscriptionNextResults
	req, err := client.ListQueryResultsForSubscriptionPreparer(ctx, subscriptionID, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "ListQueryResultsForSubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListQueryResultsForSubscriptionSender(req)
	if err != nil {
		result.ptrqr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "ListQueryResultsForSubscription", resp, "Failure sending request")
		return
	}

	result.ptrqr, err = client.ListQueryResultsForSubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "ListQueryResultsForSubscription", resp, "Failure responding to request")
	}

	return
}

// ListQueryResultsForSubscriptionPreparer prepares the ListQueryResultsForSubscription request.
func (client PolicyTrackedResourcesClient) ListQueryResultsForSubscriptionPreparer(ctx context.Context, subscriptionID string, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyTrackedResourcesResource": autorest.Encode("path", "default"),
		"subscriptionId":                 autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.PolicyInsights/policyTrackedResources/{policyTrackedResourcesResource}/queryResults", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListQueryResultsForSubscriptionSender sends the ListQueryResultsForSubscription request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyTrackedResourcesClient) ListQueryResultsForSubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListQueryResultsForSubscriptionResponder handles the response to the ListQueryResultsForSubscription request. The method always
// closes the http.Response Body.
func (client PolicyTrackedResourcesClient) ListQueryResultsForSubscriptionResponder(resp *http.Response) (result PolicyTrackedResourcesQueryResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listQueryResultsForSubscriptionNextResults retrieves the next set of results, if any.
func (client PolicyTrackedResourcesClient) listQueryResultsForSubscriptionNextResults(ctx context.Context, lastResults PolicyTrackedResourcesQueryResults) (result PolicyTrackedResourcesQueryResults, err error) {
	req, err := lastResults.policyTrackedResourcesQueryResultsPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "listQueryResultsForSubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListQueryResultsForSubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "listQueryResultsForSubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListQueryResultsForSubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyTrackedResourcesClient", "listQueryResultsForSubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListQueryResultsForSubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client PolicyTrackedResourcesClient) ListQueryResultsForSubscriptionComplete(ctx context.Context, subscriptionID string, top *int32, filter string) (result PolicyTrackedResourcesQueryResultsIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PolicyTrackedResourcesClient.ListQueryResultsForSubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListQueryResultsForSubscription(ctx, subscriptionID, top, filter)
	return
}
