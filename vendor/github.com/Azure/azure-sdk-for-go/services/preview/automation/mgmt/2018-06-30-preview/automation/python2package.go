package automation

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

// Python2PackageClient is the automation Client
type Python2PackageClient struct {
	BaseClient
}

// NewPython2PackageClient creates an instance of the Python2PackageClient client.
func NewPython2PackageClient(subscriptionID string) Python2PackageClient {
	return NewPython2PackageClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPython2PackageClientWithBaseURI creates an instance of the Python2PackageClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPython2PackageClientWithBaseURI(baseURI string, subscriptionID string) Python2PackageClient {
	return Python2PackageClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or Update the python 2 package identified by package name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// packageName - the name of python package.
// parameters - the create or update parameters for python package.
func (client Python2PackageClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, parameters PythonPackageCreateParameters) (result Module, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Python2PackageClient.CreateOrUpdate")
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.PythonPackageCreateProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.PythonPackageCreateProperties.ContentLink", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.PythonPackageCreateProperties.ContentLink.ContentHash", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.PythonPackageCreateProperties.ContentLink.ContentHash.Algorithm", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.PythonPackageCreateProperties.ContentLink.ContentHash.Value", Name: validation.Null, Rule: true, Chain: nil},
						}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("automation.Python2PackageClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, automationAccountName, packageName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.Python2PackageClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.Python2PackageClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.Python2PackageClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client Python2PackageClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, parameters PythonPackageCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"packageName":           autorest.Encode("path", packageName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages/{packageName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client Python2PackageClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client Python2PackageClient) CreateOrUpdateResponder(resp *http.Response) (result Module, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the python 2 package by name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// packageName - the python package name.
func (client Python2PackageClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Python2PackageClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.Python2PackageClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, automationAccountName, packageName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.Python2PackageClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automation.Python2PackageClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.Python2PackageClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client Python2PackageClient) DeletePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"packageName":           autorest.Encode("path", packageName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages/{packageName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client Python2PackageClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client Python2PackageClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve the python 2 package identified by package name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// packageName - the python package name.
func (client Python2PackageClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string) (result Module, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Python2PackageClient.Get")
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.Python2PackageClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, automationAccountName, packageName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.Python2PackageClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.Python2PackageClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.Python2PackageClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client Python2PackageClient) GetPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"packageName":           autorest.Encode("path", packageName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages/{packageName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client Python2PackageClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client Python2PackageClient) GetResponder(resp *http.Response) (result Module, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAutomationAccount retrieve a list of python 2 packages.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
func (client Python2PackageClient) ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string) (result ModuleListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Python2PackageClient.ListByAutomationAccount")
		defer func() {
			sc := -1
			if result.mlr.Response.Response != nil {
				sc = result.mlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.Python2PackageClient", "ListByAutomationAccount", err.Error())
	}

	result.fn = client.listByAutomationAccountNextResults
	req, err := client.ListByAutomationAccountPreparer(ctx, resourceGroupName, automationAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.Python2PackageClient", "ListByAutomationAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.mlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.Python2PackageClient", "ListByAutomationAccount", resp, "Failure sending request")
		return
	}

	result.mlr, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.Python2PackageClient", "ListByAutomationAccount", resp, "Failure responding to request")
		return
	}
	if result.mlr.hasNextLink() && result.mlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByAutomationAccountPreparer prepares the ListByAutomationAccount request.
func (client Python2PackageClient) ListByAutomationAccountPreparer(ctx context.Context, resourceGroupName string, automationAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAutomationAccountSender sends the ListByAutomationAccount request. The method will close the
// http.Response Body if it receives an error.
func (client Python2PackageClient) ListByAutomationAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByAutomationAccountResponder handles the response to the ListByAutomationAccount request. The method always
// closes the http.Response Body.
func (client Python2PackageClient) ListByAutomationAccountResponder(resp *http.Response) (result ModuleListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByAutomationAccountNextResults retrieves the next set of results, if any.
func (client Python2PackageClient) listByAutomationAccountNextResults(ctx context.Context, lastResults ModuleListResult) (result ModuleListResult, err error) {
	req, err := lastResults.moduleListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.Python2PackageClient", "listByAutomationAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.Python2PackageClient", "listByAutomationAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.Python2PackageClient", "listByAutomationAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByAutomationAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client Python2PackageClient) ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string) (result ModuleListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Python2PackageClient.ListByAutomationAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByAutomationAccount(ctx, resourceGroupName, automationAccountName)
	return
}

// Update update the python 2 package identified by package name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// packageName - the name of python package.
// parameters - the update parameters for python package.
func (client Python2PackageClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, parameters PythonPackageUpdateParameters) (result Module, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Python2PackageClient.Update")
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.Python2PackageClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, automationAccountName, packageName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.Python2PackageClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.Python2PackageClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.Python2PackageClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client Python2PackageClient) UpdatePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, parameters PythonPackageUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"packageName":           autorest.Encode("path", packageName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages/{packageName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client Python2PackageClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client Python2PackageClient) UpdateResponder(resp *http.Response) (result Module, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
