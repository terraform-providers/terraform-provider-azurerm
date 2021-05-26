package dtl

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

// ArmTemplatesClient is the the DevTest Labs Client.
type ArmTemplatesClient struct {
	BaseClient
}

// NewArmTemplatesClient creates an instance of the ArmTemplatesClient client.
func NewArmTemplatesClient(subscriptionID string) ArmTemplatesClient {
	return NewArmTemplatesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewArmTemplatesClientWithBaseURI creates an instance of the ArmTemplatesClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewArmTemplatesClientWithBaseURI(baseURI string, subscriptionID string) ArmTemplatesClient {
	return ArmTemplatesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get azure resource manager template.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// artifactSourceName - the name of the artifact source.
// name - the name of the azure Resource Manager template.
// expand - specify the $expand query. Example: 'properties($select=displayName)'
func (client ArmTemplatesClient) Get(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, expand string) (result ArmTemplate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArmTemplatesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, labName, artifactSourceName, name, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ArmTemplatesClient) GetPreparer(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"artifactSourceName": autorest.Encode("path", artifactSourceName),
		"labName":            autorest.Encode("path", labName),
		"name":               autorest.Encode("path", name),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/armtemplates/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ArmTemplatesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ArmTemplatesClient) GetResponder(resp *http.Response) (result ArmTemplate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list azure resource manager templates in a given artifact source.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// artifactSourceName - the name of the artifact source.
// expand - specify the $expand query. Example: 'properties($select=displayName)'
// filter - the filter to apply to the operation.
// top - the maximum number of resources to return from the operation.
// orderby - the ordering expression for the results, using OData notation.
func (client ArmTemplatesClient) List(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationArmTemplatePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArmTemplatesClient.List")
		defer func() {
			sc := -1
			if result.rwcat.Response.Response != nil {
				sc = result.rwcat.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, labName, artifactSourceName, expand, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rwcat.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "List", resp, "Failure sending request")
		return
	}

	result.rwcat, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rwcat.hasNextLink() && result.rwcat.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ArmTemplatesClient) ListPreparer(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, expand string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"artifactSourceName": autorest.Encode("path", artifactSourceName),
		"labName":            autorest.Encode("path", labName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/armtemplates", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ArmTemplatesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ArmTemplatesClient) ListResponder(resp *http.Response) (result ResponseWithContinuationArmTemplate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ArmTemplatesClient) listNextResults(ctx context.Context, lastResults ResponseWithContinuationArmTemplate) (result ResponseWithContinuationArmTemplate, err error) {
	req, err := lastResults.responseWithContinuationArmTemplatePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ArmTemplatesClient) ListComplete(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationArmTemplateIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArmTemplatesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, labName, artifactSourceName, expand, filter, top, orderby)
	return
}
