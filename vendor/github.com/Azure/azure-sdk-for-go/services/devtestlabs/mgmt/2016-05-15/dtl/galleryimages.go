package dtl

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

// GalleryImagesClient is the the DevTest Labs Client.
type GalleryImagesClient struct {
	BaseClient
}

// NewGalleryImagesClient creates an instance of the GalleryImagesClient client.
func NewGalleryImagesClient(subscriptionID string) GalleryImagesClient {
	return NewGalleryImagesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGalleryImagesClientWithBaseURI creates an instance of the GalleryImagesClient client.
func NewGalleryImagesClientWithBaseURI(baseURI string, subscriptionID string) GalleryImagesClient {
	return GalleryImagesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List list gallery images in a given lab.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// expand - specify the $expand query. Example: 'properties($select=author)'
// filter - the filter to apply to the operation.
// top - the maximum number of resources to return from the operation.
// orderby - the ordering expression for the results, using OData notation.
func (client GalleryImagesClient) List(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationGalleryImagePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryImagesClient.List")
		defer func() {
			sc := -1
			if result.rwcgi.Response.Response != nil {
				sc = result.rwcgi.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, labName, expand, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.GalleryImagesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rwcgi.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.GalleryImagesClient", "List", resp, "Failure sending request")
		return
	}

	result.rwcgi, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.GalleryImagesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client GalleryImagesClient) ListPreparer(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/galleryimages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client GalleryImagesClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client GalleryImagesClient) ListResponder(resp *http.Response) (result ResponseWithContinuationGalleryImage, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client GalleryImagesClient) listNextResults(ctx context.Context, lastResults ResponseWithContinuationGalleryImage) (result ResponseWithContinuationGalleryImage, err error) {
	req, err := lastResults.responseWithContinuationGalleryImagePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dtl.GalleryImagesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dtl.GalleryImagesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.GalleryImagesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client GalleryImagesClient) ListComplete(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationGalleryImageIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryImagesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, labName, expand, filter, top, orderby)
	return
}
