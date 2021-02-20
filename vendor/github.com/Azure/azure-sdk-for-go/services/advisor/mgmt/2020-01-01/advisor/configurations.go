package advisor

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

// ConfigurationsClient is the REST APIs for Azure Advisor
type ConfigurationsClient struct {
	BaseClient
}

// NewConfigurationsClient creates an instance of the ConfigurationsClient client.
func NewConfigurationsClient(subscriptionID string) ConfigurationsClient {
	return NewConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConfigurationsClientWithBaseURI creates an instance of the ConfigurationsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsClient {
	return ConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateInResourceGroup sends the create in resource group request.
// Parameters:
// configContract - the Azure Advisor configuration data structure.
// resourceGroup - the name of the Azure resource group.
func (client ConfigurationsClient) CreateInResourceGroup(ctx context.Context, configContract ConfigData, resourceGroup string) (result ConfigData, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationsClient.CreateInResourceGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateInResourceGroupPreparer(ctx, configContract, resourceGroup)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "CreateInResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateInResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "CreateInResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.CreateInResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "CreateInResourceGroup", resp, "Failure responding to request")
		return
	}

	return
}

// CreateInResourceGroupPreparer prepares the CreateInResourceGroup request.
func (client ConfigurationsClient) CreateInResourceGroupPreparer(ctx context.Context, configContract ConfigData, resourceGroup string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationName": autorest.Encode("path", "default"),
		"resourceGroup":     autorest.Encode("path", resourceGroup),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Advisor/configurations/{configurationName}", pathParameters),
		autorest.WithJSON(configContract),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateInResourceGroupSender sends the CreateInResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationsClient) CreateInResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateInResourceGroupResponder handles the response to the CreateInResourceGroup request. The method always
// closes the http.Response Body.
func (client ConfigurationsClient) CreateInResourceGroupResponder(resp *http.Response) (result ConfigData, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateInSubscription create/Overwrite Azure Advisor configuration and also delete all configurations of contained
// resource groups.
// Parameters:
// configContract - the Azure Advisor configuration data structure.
func (client ConfigurationsClient) CreateInSubscription(ctx context.Context, configContract ConfigData) (result ConfigData, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationsClient.CreateInSubscription")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateInSubscriptionPreparer(ctx, configContract)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "CreateInSubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateInSubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "CreateInSubscription", resp, "Failure sending request")
		return
	}

	result, err = client.CreateInSubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "CreateInSubscription", resp, "Failure responding to request")
		return
	}

	return
}

// CreateInSubscriptionPreparer prepares the CreateInSubscription request.
func (client ConfigurationsClient) CreateInSubscriptionPreparer(ctx context.Context, configContract ConfigData) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationName": autorest.Encode("path", "default"),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/configurations/{configurationName}", pathParameters),
		autorest.WithJSON(configContract),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateInSubscriptionSender sends the CreateInSubscription request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationsClient) CreateInSubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateInSubscriptionResponder handles the response to the CreateInSubscription request. The method always
// closes the http.Response Body.
func (client ConfigurationsClient) CreateInSubscriptionResponder(resp *http.Response) (result ConfigData, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup sends the list by resource group request.
// Parameters:
// resourceGroup - the name of the Azure resource group.
func (client ConfigurationsClient) ListByResourceGroup(ctx context.Context, resourceGroup string) (result ConfigurationListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroup)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ConfigurationsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroup string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroup":  autorest.Encode("path", resourceGroup),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Advisor/configurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ConfigurationsClient) ListByResourceGroupResponder(resp *http.Response) (result ConfigurationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscription retrieve Azure Advisor configurations and also retrieve configurations of contained resource
// groups.
func (client ConfigurationsClient) ListBySubscription(ctx context.Context) (result ConfigurationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.clr.Response.Response != nil {
				sc = result.clr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.clr.hasNextLink() && result.clr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client ConfigurationsClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/configurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client ConfigurationsClient) ListBySubscriptionResponder(resp *http.Response) (result ConfigurationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client ConfigurationsClient) listBySubscriptionNextResults(ctx context.Context, lastResults ConfigurationListResult) (result ConfigurationListResult, err error) {
	req, err := lastResults.configurationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client ConfigurationsClient) ListBySubscriptionComplete(ctx context.Context) (result ConfigurationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx)
	return
}
