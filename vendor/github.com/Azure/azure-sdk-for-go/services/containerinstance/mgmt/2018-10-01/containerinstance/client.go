// Package containerinstance implements the Azure ARM Containerinstance service API version 2018-10-01.
//
//
package containerinstance

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

const (
	// DefaultBaseURI is the default URI used for the service Containerinstance
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Containerinstance.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// ListCachedImages get the list of cached images on specific OS type for a subscription in a region.
// Parameters:
// location - the identifier for the physical azure location.
func (client BaseClient) ListCachedImages(ctx context.Context, location string) (result CachedImagesListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ListCachedImages")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListCachedImagesPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.BaseClient", "ListCachedImages", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListCachedImagesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerinstance.BaseClient", "ListCachedImages", resp, "Failure sending request")
		return
	}

	result, err = client.ListCachedImagesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.BaseClient", "ListCachedImages", resp, "Failure responding to request")
	}

	return
}

// ListCachedImagesPreparer prepares the ListCachedImages request.
func (client BaseClient) ListCachedImagesPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/cachedImages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListCachedImagesSender sends the ListCachedImages request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ListCachedImagesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListCachedImagesResponder handles the response to the ListCachedImages request. The method always
// closes the http.Response Body.
func (client BaseClient) ListCachedImagesResponder(resp *http.Response) (result CachedImagesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListCapabilities get the list of CPU/memory/GPU capabilities of a region.
// Parameters:
// location - the identifier for the physical azure location.
func (client BaseClient) ListCapabilities(ctx context.Context, location string) (result CapabilitiesListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ListCapabilities")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListCapabilitiesPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.BaseClient", "ListCapabilities", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListCapabilitiesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerinstance.BaseClient", "ListCapabilities", resp, "Failure sending request")
		return
	}

	result, err = client.ListCapabilitiesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.BaseClient", "ListCapabilities", resp, "Failure responding to request")
	}

	return
}

// ListCapabilitiesPreparer prepares the ListCapabilities request.
func (client BaseClient) ListCapabilitiesPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/capabilities", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListCapabilitiesSender sends the ListCapabilities request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ListCapabilitiesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListCapabilitiesResponder handles the response to the ListCapabilities request. The method always
// closes the http.Response Body.
func (client BaseClient) ListCapabilitiesResponder(resp *http.Response) (result CapabilitiesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
