package managedapplications

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

// ApplicationDefinitionsClient is the ARM applications
type ApplicationDefinitionsClient struct {
	BaseClient
}

// NewApplicationDefinitionsClient creates an instance of the ApplicationDefinitionsClient client.
func NewApplicationDefinitionsClient(subscriptionID string) ApplicationDefinitionsClient {
	return NewApplicationDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationDefinitionsClientWithBaseURI creates an instance of the ApplicationDefinitionsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewApplicationDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationDefinitionsClient {
	return ApplicationDefinitionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new managed application definition.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applicationDefinitionName - the name of the managed application definition.
// parameters - parameters supplied to the create or update an managed application definition.
func (client ApplicationDefinitionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition) (result ApplicationDefinitionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationDefinitionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\p{L}\._\(\)\w]+$`, Chain: nil}}},
		{TargetValue: applicationDefinitionName,
			Constraints: []validation.Constraint{{Target: "applicationDefinitionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationDefinitionName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ApplicationDefinitionProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.ApplicationDefinitionProperties.NotificationPolicy", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.ApplicationDefinitionProperties.NotificationPolicy.NotificationEndpoints", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplicationDefinitionsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, applicationDefinitionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ApplicationDefinitionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationDefinitionName": autorest.Encode("path", applicationDefinitionName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationDefinitionsClient) CreateOrUpdateSender(req *http.Request) (future ApplicationDefinitionsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ApplicationDefinitionsClient) (ad ApplicationDefinition, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("managedapplications.ApplicationDefinitionsCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		ad.Response.Response, err = future.GetResult(sender)
		if ad.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && ad.Response.Response.StatusCode != http.StatusNoContent {
			ad, err = client.CreateOrUpdateResponder(ad.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsCreateOrUpdateFuture", "Result", ad.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ApplicationDefinitionsClient) CreateOrUpdateResponder(resp *http.Response) (result ApplicationDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateByID creates a new managed application definition.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applicationDefinitionName - the name of the managed application definition.
// parameters - parameters supplied to the create or update a managed application definition.
func (client ApplicationDefinitionsClient) CreateOrUpdateByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition) (result ApplicationDefinitionsCreateOrUpdateByIDFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationDefinitionsClient.CreateOrUpdateByID")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: applicationDefinitionName,
			Constraints: []validation.Constraint{{Target: "applicationDefinitionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationDefinitionName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ApplicationDefinitionProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.ApplicationDefinitionProperties.NotificationPolicy", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.ApplicationDefinitionProperties.NotificationPolicy.NotificationEndpoints", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplicationDefinitionsClient", "CreateOrUpdateByID", err.Error())
	}

	req, err := client.CreateOrUpdateByIDPreparer(ctx, resourceGroupName, applicationDefinitionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "CreateOrUpdateByID", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateByIDSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "CreateOrUpdateByID", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateByIDPreparer prepares the CreateOrUpdateByID request.
func (client ApplicationDefinitionsClient) CreateOrUpdateByIDPreparer(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationDefinitionName": autorest.Encode("path", applicationDefinitionName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateByIDSender sends the CreateOrUpdateByID request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationDefinitionsClient) CreateOrUpdateByIDSender(req *http.Request) (future ApplicationDefinitionsCreateOrUpdateByIDFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ApplicationDefinitionsClient) (ad ApplicationDefinition, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsCreateOrUpdateByIDFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("managedapplications.ApplicationDefinitionsCreateOrUpdateByIDFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		ad.Response.Response, err = future.GetResult(sender)
		if ad.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsCreateOrUpdateByIDFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && ad.Response.Response.StatusCode != http.StatusNoContent {
			ad, err = client.CreateOrUpdateByIDResponder(ad.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsCreateOrUpdateByIDFuture", "Result", ad.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateByIDResponder handles the response to the CreateOrUpdateByID request. The method always
// closes the http.Response Body.
func (client ApplicationDefinitionsClient) CreateOrUpdateByIDResponder(resp *http.Response) (result ApplicationDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the managed application definition.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applicationDefinitionName - the name of the managed application definition to delete.
func (client ApplicationDefinitionsClient) Delete(ctx context.Context, resourceGroupName string, applicationDefinitionName string) (result ApplicationDefinitionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationDefinitionsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\p{L}\._\(\)\w]+$`, Chain: nil}}},
		{TargetValue: applicationDefinitionName,
			Constraints: []validation.Constraint{{Target: "applicationDefinitionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationDefinitionName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplicationDefinitionsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, applicationDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationDefinitionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, applicationDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationDefinitionName": autorest.Encode("path", applicationDefinitionName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationDefinitionsClient) DeleteSender(req *http.Request) (future ApplicationDefinitionsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ApplicationDefinitionsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("managedapplications.ApplicationDefinitionsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplicationDefinitionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteByID deletes the managed application definition.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applicationDefinitionName - the name of the managed application definition.
func (client ApplicationDefinitionsClient) DeleteByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string) (result ApplicationDefinitionsDeleteByIDFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationDefinitionsClient.DeleteByID")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: applicationDefinitionName,
			Constraints: []validation.Constraint{{Target: "applicationDefinitionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationDefinitionName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplicationDefinitionsClient", "DeleteByID", err.Error())
	}

	req, err := client.DeleteByIDPreparer(ctx, resourceGroupName, applicationDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "DeleteByID", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteByIDSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "DeleteByID", nil, "Failure sending request")
		return
	}

	return
}

// DeleteByIDPreparer prepares the DeleteByID request.
func (client ApplicationDefinitionsClient) DeleteByIDPreparer(ctx context.Context, resourceGroupName string, applicationDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationDefinitionName": autorest.Encode("path", applicationDefinitionName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteByIDSender sends the DeleteByID request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationDefinitionsClient) DeleteByIDSender(req *http.Request) (future ApplicationDefinitionsDeleteByIDFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ApplicationDefinitionsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsDeleteByIDFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("managedapplications.ApplicationDefinitionsDeleteByIDFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteByIDResponder handles the response to the DeleteByID request. The method always
// closes the http.Response Body.
func (client ApplicationDefinitionsClient) DeleteByIDResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the managed application definition.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applicationDefinitionName - the name of the managed application definition.
func (client ApplicationDefinitionsClient) Get(ctx context.Context, resourceGroupName string, applicationDefinitionName string) (result ApplicationDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationDefinitionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\p{L}\._\(\)\w]+$`, Chain: nil}}},
		{TargetValue: applicationDefinitionName,
			Constraints: []validation.Constraint{{Target: "applicationDefinitionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationDefinitionName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplicationDefinitionsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, applicationDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationDefinitionsClient) GetPreparer(ctx context.Context, resourceGroupName string, applicationDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationDefinitionName": autorest.Encode("path", applicationDefinitionName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationDefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationDefinitionsClient) GetResponder(resp *http.Response) (result ApplicationDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByID gets the managed application definition.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applicationDefinitionName - the name of the managed application definition.
func (client ApplicationDefinitionsClient) GetByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string) (result ApplicationDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationDefinitionsClient.GetByID")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: applicationDefinitionName,
			Constraints: []validation.Constraint{{Target: "applicationDefinitionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationDefinitionName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplicationDefinitionsClient", "GetByID", err.Error())
	}

	req, err := client.GetByIDPreparer(ctx, resourceGroupName, applicationDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "GetByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "GetByID", resp, "Failure sending request")
		return
	}

	result, err = client.GetByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "GetByID", resp, "Failure responding to request")
		return
	}

	return
}

// GetByIDPreparer prepares the GetByID request.
func (client ApplicationDefinitionsClient) GetByIDPreparer(ctx context.Context, resourceGroupName string, applicationDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationDefinitionName": autorest.Encode("path", applicationDefinitionName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByIDSender sends the GetByID request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationDefinitionsClient) GetByIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetByIDResponder handles the response to the GetByID request. The method always
// closes the http.Response Body.
func (client ApplicationDefinitionsClient) GetByIDResponder(resp *http.Response) (result ApplicationDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup lists the managed application definitions in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
func (client ApplicationDefinitionsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ApplicationDefinitionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationDefinitionsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.adlr.Response.Response != nil {
				sc = result.adlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\p{L}\._\(\)\w]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplicationDefinitionsClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.adlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.adlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.adlr.hasNextLink() && result.adlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ApplicationDefinitionsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationDefinitionsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ApplicationDefinitionsClient) ListByResourceGroupResponder(resp *http.Response) (result ApplicationDefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ApplicationDefinitionsClient) listByResourceGroupNextResults(ctx context.Context, lastResults ApplicationDefinitionListResult) (result ApplicationDefinitionListResult, err error) {
	req, err := lastResults.applicationDefinitionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationDefinitionsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationDefinitionsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ApplicationDefinitionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationDefinitionsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}
