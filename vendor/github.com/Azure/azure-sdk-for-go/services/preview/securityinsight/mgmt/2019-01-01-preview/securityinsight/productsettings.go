package securityinsight

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

// ProductSettingsClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights) resource provider
type ProductSettingsClient struct {
	BaseClient
}

// NewProductSettingsClient creates an instance of the ProductSettingsClient client.
func NewProductSettingsClient(subscriptionID string) ProductSettingsClient {
	return NewProductSettingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProductSettingsClientWithBaseURI creates an instance of the ProductSettingsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewProductSettingsClientWithBaseURI(baseURI string, subscriptionID string) ProductSettingsClient {
	return ProductSettingsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete delete setting of the product.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// settingsName - the setting name. Supports- EyesOn
func (client ProductSettingsClient) Delete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, settingsName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductSettingsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.ProductSettingsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, settingsName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ProductSettingsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "securityinsight.ProductSettingsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ProductSettingsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ProductSettingsClient) DeletePreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, settingsName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"settingsName":                        autorest.Encode("path", settingsName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/settings/{settingsName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ProductSettingsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ProductSettingsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a setting.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// settingsName - the setting name. Supports- EyesOn
func (client ProductSettingsClient) Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, settingsName string) (result SettingsModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductSettingsClient.Get")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.ProductSettingsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, settingsName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ProductSettingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.ProductSettingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ProductSettingsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProductSettingsClient) GetPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, settingsName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"settingsName":                        autorest.Encode("path", settingsName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/settings/{settingsName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProductSettingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProductSettingsClient) GetResponder(resp *http.Response) (result SettingsModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAll list of all the settings
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
func (client ProductSettingsClient) GetAll(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result SettingList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductSettingsClient.GetAll")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.ProductSettingsClient", "GetAll", err.Error())
	}

	req, err := client.GetAllPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ProductSettingsClient", "GetAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.ProductSettingsClient", "GetAll", resp, "Failure sending request")
		return
	}

	result, err = client.GetAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ProductSettingsClient", "GetAll", resp, "Failure responding to request")
	}

	return
}

// GetAllPreparer prepares the GetAll request.
func (client ProductSettingsClient) GetAllPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/settings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAllSender sends the GetAll request. The method will close the
// http.Response Body if it receives an error.
func (client ProductSettingsClient) GetAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetAllResponder handles the response to the GetAll request. The method always
// closes the http.Response Body.
func (client ProductSettingsClient) GetAllResponder(resp *http.Response) (result SettingList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates setting.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// settingsName - the setting name. Supports- EyesOn
// settings - the setting
func (client ProductSettingsClient) Update(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, settingsName string, settings BasicSettings) (result SettingsModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductSettingsClient.Update")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.ProductSettingsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, settingsName, settings)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ProductSettingsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.ProductSettingsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.ProductSettingsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ProductSettingsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, settingsName string, settings BasicSettings) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"settingsName":                        autorest.Encode("path", settingsName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/settings/{settingsName}", pathParameters),
		autorest.WithJSON(settings),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ProductSettingsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ProductSettingsClient) UpdateResponder(resp *http.Response) (result SettingsModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
