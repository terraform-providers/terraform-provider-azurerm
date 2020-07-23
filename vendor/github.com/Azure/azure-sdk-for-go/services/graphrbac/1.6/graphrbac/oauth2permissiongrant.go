package graphrbac

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
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// OAuth2PermissionGrantClient is the the Graph RBAC Management Client
type OAuth2PermissionGrantClient struct {
	BaseClient
}

// NewOAuth2PermissionGrantClient creates an instance of the OAuth2PermissionGrantClient client.
func NewOAuth2PermissionGrantClient(tenantID string) OAuth2PermissionGrantClient {
	return NewOAuth2PermissionGrantClientWithBaseURI(DefaultBaseURI, tenantID)
}

// NewOAuth2PermissionGrantClientWithBaseURI creates an instance of the OAuth2PermissionGrantClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewOAuth2PermissionGrantClientWithBaseURI(baseURI string, tenantID string) OAuth2PermissionGrantClient {
	return OAuth2PermissionGrantClient{NewWithBaseURI(baseURI, tenantID)}
}

// Create grants OAuth2 permissions for the relevant resource Ids of an app.
// Parameters:
// body - the relevant app Service Principal Object Id and the Service Principal Object Id you want to grant.
func (client OAuth2PermissionGrantClient) Create(ctx context.Context, body *OAuth2PermissionGrant) (result OAuth2PermissionGrant, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OAuth2PermissionGrantClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2PermissionGrantClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2PermissionGrantClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2PermissionGrantClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client OAuth2PermissionGrantClient) CreatePreparer(ctx context.Context, body *OAuth2PermissionGrant) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/oauth2PermissionGrants", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if body != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(body))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client OAuth2PermissionGrantClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client OAuth2PermissionGrantClient) CreateResponder(resp *http.Response) (result OAuth2PermissionGrant, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a OAuth2 permission grant for the relevant resource Ids of an app.
// Parameters:
// objectID - the object ID of a permission grant.
func (client OAuth2PermissionGrantClient) Delete(ctx context.Context, objectID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OAuth2PermissionGrantClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, objectID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2PermissionGrantClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2PermissionGrantClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2PermissionGrantClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client OAuth2PermissionGrantClient) DeletePreparer(ctx context.Context, objectID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"objectId": autorest.Encode("path", objectID),
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/oauth2PermissionGrants/{objectId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client OAuth2PermissionGrantClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client OAuth2PermissionGrantClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// List queries OAuth2 permissions grants for the relevant SP ObjectId of an app.
// Parameters:
// filter - this is the Service Principal ObjectId associated with the app
func (client OAuth2PermissionGrantClient) List(ctx context.Context, filter string) (result OAuth2PermissionGrantListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OAuth2PermissionGrantClient.List")
		defer func() {
			sc := -1
			if result.oa2pglr.Response.Response != nil {
				sc = result.oa2pglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = func(ctx context.Context, lastResult OAuth2PermissionGrantListResult) (OAuth2PermissionGrantListResult, error) {
		if lastResult.OdataNextLink == nil || len(to.String(lastResult.OdataNextLink)) < 1 {
			return OAuth2PermissionGrantListResult{}, nil
		}
		return client.ListNext(ctx, *lastResult.OdataNextLink)
	}
	req, err := client.ListPreparer(ctx, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2PermissionGrantClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.oa2pglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2PermissionGrantClient", "List", resp, "Failure sending request")
		return
	}

	result.oa2pglr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2PermissionGrantClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client OAuth2PermissionGrantClient) ListPreparer(ctx context.Context, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/oauth2PermissionGrants", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client OAuth2PermissionGrantClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client OAuth2PermissionGrantClient) ListResponder(resp *http.Response) (result OAuth2PermissionGrantListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client OAuth2PermissionGrantClient) ListComplete(ctx context.Context, filter string) (result OAuth2PermissionGrantListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OAuth2PermissionGrantClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, filter)
	return
}

// ListNext gets the next page of OAuth2 permission grants
// Parameters:
// nextLink - next link for the list operation.
func (client OAuth2PermissionGrantClient) ListNext(ctx context.Context, nextLink string) (result OAuth2PermissionGrantListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OAuth2PermissionGrantClient.ListNext")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListNextPreparer(ctx, nextLink)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2PermissionGrantClient", "ListNext", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListNextSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2PermissionGrantClient", "ListNext", resp, "Failure sending request")
		return
	}

	result, err = client.ListNextResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2PermissionGrantClient", "ListNext", resp, "Failure responding to request")
	}

	return
}

// ListNextPreparer prepares the ListNext request.
func (client OAuth2PermissionGrantClient) ListNextPreparer(ctx context.Context, nextLink string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nextLink": nextLink,
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/{nextLink}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListNextSender sends the ListNext request. The method will close the
// http.Response Body if it receives an error.
func (client OAuth2PermissionGrantClient) ListNextSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListNextResponder handles the response to the ListNext request. The method always
// closes the http.Response Body.
func (client OAuth2PermissionGrantClient) ListNextResponder(resp *http.Response) (result OAuth2PermissionGrantListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
