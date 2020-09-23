package sql

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

// JobCredentialsClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type JobCredentialsClient struct {
	BaseClient
}

// NewJobCredentialsClient creates an instance of the JobCredentialsClient client.
func NewJobCredentialsClient(subscriptionID string) JobCredentialsClient {
	return NewJobCredentialsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobCredentialsClientWithBaseURI creates an instance of the JobCredentialsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewJobCredentialsClientWithBaseURI(baseURI string, subscriptionID string) JobCredentialsClient {
	return JobCredentialsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a job credential.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
// credentialName - the name of the credential.
// parameters - the requested job credential state.
func (client JobCredentialsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, credentialName string, parameters JobCredential) (result JobCredential, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobCredentialsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.JobCredentialProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.JobCredentialProperties.Username", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.JobCredentialProperties.Password", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("sql.JobCredentialsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, jobAgentName, credentialName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobCredentialsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.JobCredentialsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobCredentialsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client JobCredentialsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, credentialName string, parameters JobCredential) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"credentialName":    autorest.Encode("path", credentialName),
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/credentials/{credentialName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client JobCredentialsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client JobCredentialsClient) CreateOrUpdateResponder(resp *http.Response) (result JobCredential, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a job credential.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
// credentialName - the name of the credential.
func (client JobCredentialsClient) Delete(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, credentialName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobCredentialsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, jobAgentName, credentialName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobCredentialsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sql.JobCredentialsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobCredentialsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client JobCredentialsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, credentialName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"credentialName":    autorest.Encode("path", credentialName),
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/credentials/{credentialName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client JobCredentialsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client JobCredentialsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a jobs credential.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
// credentialName - the name of the credential.
func (client JobCredentialsClient) Get(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, credentialName string) (result JobCredential, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobCredentialsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, jobAgentName, credentialName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobCredentialsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.JobCredentialsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobCredentialsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobCredentialsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, credentialName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"credentialName":    autorest.Encode("path", credentialName),
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/credentials/{credentialName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobCredentialsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobCredentialsClient) GetResponder(resp *http.Response) (result JobCredential, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAgent gets a list of jobs credentials.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
func (client JobCredentialsClient) ListByAgent(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string) (result JobCredentialListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobCredentialsClient.ListByAgent")
		defer func() {
			sc := -1
			if result.jclr.Response.Response != nil {
				sc = result.jclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByAgentNextResults
	req, err := client.ListByAgentPreparer(ctx, resourceGroupName, serverName, jobAgentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobCredentialsClient", "ListByAgent", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAgentSender(req)
	if err != nil {
		result.jclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.JobCredentialsClient", "ListByAgent", resp, "Failure sending request")
		return
	}

	result.jclr, err = client.ListByAgentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobCredentialsClient", "ListByAgent", resp, "Failure responding to request")
	}
	if result.jclr.hasNextLink() && result.jclr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByAgentPreparer prepares the ListByAgent request.
func (client JobCredentialsClient) ListByAgentPreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/credentials", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAgentSender sends the ListByAgent request. The method will close the
// http.Response Body if it receives an error.
func (client JobCredentialsClient) ListByAgentSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByAgentResponder handles the response to the ListByAgent request. The method always
// closes the http.Response Body.
func (client JobCredentialsClient) ListByAgentResponder(resp *http.Response) (result JobCredentialListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByAgentNextResults retrieves the next set of results, if any.
func (client JobCredentialsClient) listByAgentNextResults(ctx context.Context, lastResults JobCredentialListResult) (result JobCredentialListResult, err error) {
	req, err := lastResults.jobCredentialListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.JobCredentialsClient", "listByAgentNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAgentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.JobCredentialsClient", "listByAgentNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAgentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobCredentialsClient", "listByAgentNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByAgentComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobCredentialsClient) ListByAgentComplete(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string) (result JobCredentialListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobCredentialsClient.ListByAgent")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByAgent(ctx, resourceGroupName, serverName, jobAgentName)
	return
}
