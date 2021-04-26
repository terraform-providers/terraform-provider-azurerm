package privatedns

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

// PrivateZonesClient is the the Private DNS Management Client.
type PrivateZonesClient struct {
	BaseClient
}

// NewPrivateZonesClient creates an instance of the PrivateZonesClient client.
func NewPrivateZonesClient(subscriptionID string) PrivateZonesClient {
	return NewPrivateZonesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateZonesClientWithBaseURI creates an instance of the PrivateZonesClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPrivateZonesClientWithBaseURI(baseURI string, subscriptionID string) PrivateZonesClient {
	return PrivateZonesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a Private DNS zone. Does not modify Links to virtual networks or DNS records
// within the zone.
// Parameters:
// resourceGroupName - the name of the resource group.
// privateZoneName - the name of the Private DNS zone (without a terminating dot).
// parameters - parameters supplied to the CreateOrUpdate operation.
// ifMatch - the ETag of the Private DNS zone. Omit this value to always overwrite the current zone. Specify
// the last-seen ETag value to prevent accidentally overwriting any concurrent changes.
// ifNoneMatch - set to '*' to allow a new Private DNS zone to be created, but to prevent updating an existing
// zone. Other values will be ignored.
func (client PrivateZonesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, parameters PrivateZone, ifMatch string, ifNoneMatch string) (result PrivateZonesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateZonesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, privateZoneName, parameters, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PrivateZonesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, privateZoneName string, parameters PrivateZone, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateZoneName":   autorest.Encode("path", privateZoneName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateZonesClient) CreateOrUpdateSender(req *http.Request) (future PrivateZonesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client PrivateZonesClient) CreateOrUpdateResponder(resp *http.Response) (result PrivateZone, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Private DNS zone. WARNING: All DNS records in the zone will also be deleted. This operation cannot
// be undone. Private DNS zone cannot be deleted unless all virtual network links to it are removed.
// Parameters:
// resourceGroupName - the name of the resource group.
// privateZoneName - the name of the Private DNS zone (without a terminating dot).
// ifMatch - the ETag of the Private DNS zone. Omit this value to always delete the current zone. Specify the
// last-seen ETag value to prevent accidentally deleting any concurrent changes.
func (client PrivateZonesClient) Delete(ctx context.Context, resourceGroupName string, privateZoneName string, ifMatch string) (result PrivateZonesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateZonesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, privateZoneName, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PrivateZonesClient) DeletePreparer(ctx context.Context, resourceGroupName string, privateZoneName string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateZoneName":   autorest.Encode("path", privateZoneName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateZonesClient) DeleteSender(req *http.Request) (future PrivateZonesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PrivateZonesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a Private DNS zone. Retrieves the zone properties, but not the virtual networks links or the record sets
// within the zone.
// Parameters:
// resourceGroupName - the name of the resource group.
// privateZoneName - the name of the Private DNS zone (without a terminating dot).
func (client PrivateZonesClient) Get(ctx context.Context, resourceGroupName string, privateZoneName string) (result PrivateZone, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateZonesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, privateZoneName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrivateZonesClient) GetPreparer(ctx context.Context, resourceGroupName string, privateZoneName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateZoneName":   autorest.Encode("path", privateZoneName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateZonesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrivateZonesClient) GetResponder(resp *http.Response) (result PrivateZone, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the Private DNS zones in all resource groups in a subscription.
// Parameters:
// top - the maximum number of Private DNS zones to return. If not specified, returns up to 100 zones.
func (client PrivateZonesClient) List(ctx context.Context, top *int32) (result PrivateZoneListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateZonesClient.List")
		defer func() {
			sc := -1
			if result.pzlr.Response.Response != nil {
				sc = result.pzlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pzlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "List", resp, "Failure sending request")
		return
	}

	result.pzlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.pzlr.hasNextLink() && result.pzlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client PrivateZonesClient) ListPreparer(ctx context.Context, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/privateDnsZones", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateZonesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PrivateZonesClient) ListResponder(resp *http.Response) (result PrivateZoneListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client PrivateZonesClient) listNextResults(ctx context.Context, lastResults PrivateZoneListResult) (result PrivateZoneListResult, err error) {
	req, err := lastResults.privateZoneListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateZonesClient) ListComplete(ctx context.Context, top *int32) (result PrivateZoneListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateZonesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, top)
	return
}

// ListByResourceGroup lists the Private DNS zones within a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// top - the maximum number of record sets to return. If not specified, returns up to 100 record sets.
func (client PrivateZonesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, top *int32) (result PrivateZoneListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateZonesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.pzlr.Response.Response != nil {
				sc = result.pzlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.pzlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.pzlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.pzlr.hasNextLink() && result.pzlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client PrivateZonesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateZonesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client PrivateZonesClient) ListByResourceGroupResponder(resp *http.Response) (result PrivateZoneListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client PrivateZonesClient) listByResourceGroupNextResults(ctx context.Context, lastResults PrivateZoneListResult) (result PrivateZoneListResult, err error) {
	req, err := lastResults.privateZoneListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateZonesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, top *int32) (result PrivateZoneListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateZonesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, top)
	return
}

// Update updates a Private DNS zone. Does not modify virtual network links or DNS records within the zone.
// Parameters:
// resourceGroupName - the name of the resource group.
// privateZoneName - the name of the Private DNS zone (without a terminating dot).
// parameters - parameters supplied to the Update operation.
// ifMatch - the ETag of the Private DNS zone. Omit this value to always overwrite the current zone. Specify
// the last-seen ETag value to prevent accidentally overwriting any concurrent changes.
func (client PrivateZonesClient) Update(ctx context.Context, resourceGroupName string, privateZoneName string, parameters PrivateZone, ifMatch string) (result PrivateZonesUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateZonesClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, privateZoneName, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatedns.PrivateZonesClient", "Update", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client PrivateZonesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, privateZoneName string, parameters PrivateZone, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateZoneName":   autorest.Encode("path", privateZoneName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateZonesClient) UpdateSender(req *http.Request) (future PrivateZonesUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client PrivateZonesClient) UpdateResponder(resp *http.Response) (result PrivateZone, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
