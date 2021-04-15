package hdinsight

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

// ApplicationsClient is the hDInsight Management Client
type ApplicationsClient struct {
	BaseClient
}

// NewApplicationsClient creates an instance of the ApplicationsClient client.
func NewApplicationsClient(subscriptionID string) ApplicationsClient {
	return NewApplicationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationsClientWithBaseURI creates an instance of the ApplicationsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewApplicationsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationsClient {
	return ApplicationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates applications for the HDInsight cluster.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster.
// applicationName - the constant value for the application name.
// parameters - the application create request.
func (client ApplicationsClient) Create(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, parameters Application) (result ApplicationsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationsClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, resourceGroupName, clusterName, applicationName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ApplicationsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ApplicationsClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ApplicationsClient) CreatePreparer(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, parameters Application) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications/{applicationName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) CreateSender(req *http.Request) (future ApplicationsCreateFuture, err error) {
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

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) CreateResponder(resp *http.Response) (result Application, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified application on the HDInsight cluster.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster.
// applicationName - the constant value for the application name.
func (client ApplicationsClient) Delete(ctx context.Context, resourceGroupName string, clusterName string, applicationName string) (result ApplicationsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName, applicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ApplicationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ApplicationsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string, applicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications/{applicationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) DeleteSender(req *http.Request) (future ApplicationsDeleteFuture, err error) {
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
func (client ApplicationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets properties of the specified application.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster.
// applicationName - the constant value for the application name.
func (client ApplicationsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, applicationName string) (result Application, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName, applicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ApplicationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hdinsight.ApplicationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ApplicationsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationsClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string, applicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications/{applicationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) GetResponder(resp *http.Response) (result Application, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByCluster lists all of the applications for the HDInsight cluster.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster.
func (client ApplicationsClient) ListByCluster(ctx context.Context, resourceGroupName string, clusterName string) (result ApplicationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationsClient.ListByCluster")
		defer func() {
			sc := -1
			if result.alr.Response.Response != nil {
				sc = result.alr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByClusterNextResults
	req, err := client.ListByClusterPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ApplicationsClient", "ListByCluster", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByClusterSender(req)
	if err != nil {
		result.alr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hdinsight.ApplicationsClient", "ListByCluster", resp, "Failure sending request")
		return
	}

	result.alr, err = client.ListByClusterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ApplicationsClient", "ListByCluster", resp, "Failure responding to request")
		return
	}
	if result.alr.hasNextLink() && result.alr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByClusterPreparer prepares the ListByCluster request.
func (client ApplicationsClient) ListByClusterPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByClusterSender sends the ListByCluster request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) ListByClusterSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByClusterResponder handles the response to the ListByCluster request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) ListByClusterResponder(resp *http.Response) (result ApplicationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByClusterNextResults retrieves the next set of results, if any.
func (client ApplicationsClient) listByClusterNextResults(ctx context.Context, lastResults ApplicationListResult) (result ApplicationListResult, err error) {
	req, err := lastResults.applicationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hdinsight.ApplicationsClient", "listByClusterNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByClusterSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hdinsight.ApplicationsClient", "listByClusterNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByClusterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ApplicationsClient", "listByClusterNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByClusterComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationsClient) ListByClusterComplete(ctx context.Context, resourceGroupName string, clusterName string) (result ApplicationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationsClient.ListByCluster")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByCluster(ctx, resourceGroupName, clusterName)
	return
}
