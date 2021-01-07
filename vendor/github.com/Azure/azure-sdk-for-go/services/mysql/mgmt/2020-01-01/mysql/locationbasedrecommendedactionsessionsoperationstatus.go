package mysql

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

// LocationBasedRecommendedActionSessionsOperationStatusClient is the the Microsoft Azure management API provides
// create, read, update, and delete functionality for Azure MySQL resources including servers, databases, firewall
// rules, VNET rules, log files and configurations with new business model.
type LocationBasedRecommendedActionSessionsOperationStatusClient struct {
	BaseClient
}

// NewLocationBasedRecommendedActionSessionsOperationStatusClient creates an instance of the
// LocationBasedRecommendedActionSessionsOperationStatusClient client.
func NewLocationBasedRecommendedActionSessionsOperationStatusClient(subscriptionID string) LocationBasedRecommendedActionSessionsOperationStatusClient {
	return NewLocationBasedRecommendedActionSessionsOperationStatusClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLocationBasedRecommendedActionSessionsOperationStatusClientWithBaseURI creates an instance of the
// LocationBasedRecommendedActionSessionsOperationStatusClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewLocationBasedRecommendedActionSessionsOperationStatusClientWithBaseURI(baseURI string, subscriptionID string) LocationBasedRecommendedActionSessionsOperationStatusClient {
	return LocationBasedRecommendedActionSessionsOperationStatusClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get recommendation action session operation status.
// Parameters:
// locationName - the name of the location.
// operationID - the operation identifier.
func (client LocationBasedRecommendedActionSessionsOperationStatusClient) Get(ctx context.Context, locationName string, operationID string) (result RecommendedActionSessionsOperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationBasedRecommendedActionSessionsOperationStatusClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("mysql.LocationBasedRecommendedActionSessionsOperationStatusClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, locationName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.LocationBasedRecommendedActionSessionsOperationStatusClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mysql.LocationBasedRecommendedActionSessionsOperationStatusClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.LocationBasedRecommendedActionSessionsOperationStatusClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client LocationBasedRecommendedActionSessionsOperationStatusClient) GetPreparer(ctx context.Context, locationName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"operationId":    autorest.Encode("path", operationID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/locations/{locationName}/recommendedActionSessionsAzureAsyncOperation/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LocationBasedRecommendedActionSessionsOperationStatusClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LocationBasedRecommendedActionSessionsOperationStatusClient) GetResponder(resp *http.Response) (result RecommendedActionSessionsOperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
