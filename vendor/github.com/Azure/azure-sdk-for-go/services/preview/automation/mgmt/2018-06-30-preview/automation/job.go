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

// JobClient is the automation Client
type JobClient struct {
	BaseClient
}

// NewJobClient creates an instance of the JobClient client.
func NewJobClient(subscriptionID string) JobClient {
	return NewJobClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobClientWithBaseURI creates an instance of the JobClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewJobClientWithBaseURI(baseURI string, subscriptionID string) JobClient {
	return JobClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a job of the runbook.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// jobName - the job name.
// parameters - the parameters supplied to the create job operation.
// clientRequestID - identifies this specific client request.
func (client JobClient) Create(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, parameters JobCreateParameters, clientRequestID string) (result Job, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobClient.Create")
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
			Constraints: []validation.Constraint{{Target: "parameters.JobCreateProperties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.JobClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, automationAccountName, jobName, parameters, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client JobClient) CreatePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, parameters JobCreateParameters, clientRequestID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobName":               autorest.Encode("path", jobName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(clientRequestID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("clientRequestId", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client JobClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client JobClient) CreateResponder(resp *http.Response) (result Job, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get retrieve the job identified by job name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// jobName - the job name.
// clientRequestID - identifies this specific client request.
func (client JobClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, clientRequestID string) (result Job, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobClient.Get")
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
		return result, validation.NewError("automation.JobClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, automationAccountName, jobName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobClient) GetPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, clientRequestID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobName":               autorest.Encode("path", jobName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(clientRequestID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("clientRequestId", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobClient) GetResponder(resp *http.Response) (result Job, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetOutput retrieve the job output identified by job name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// jobName - the name of the job to be created.
// clientRequestID - identifies this specific client request.
func (client JobClient) GetOutput(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, clientRequestID string) (result ReadCloser, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobClient.GetOutput")
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
		return result, validation.NewError("automation.JobClient", "GetOutput", err.Error())
	}

	req, err := client.GetOutputPreparer(ctx, resourceGroupName, automationAccountName, jobName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "GetOutput", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOutputSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.JobClient", "GetOutput", resp, "Failure sending request")
		return
	}

	result, err = client.GetOutputResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "GetOutput", resp, "Failure responding to request")
	}

	return
}

// GetOutputPreparer prepares the GetOutput request.
func (client JobClient) GetOutputPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, clientRequestID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobName":               autorest.Encode("path", jobName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}/output", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(clientRequestID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("clientRequestId", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetOutputSender sends the GetOutput request. The method will close the
// http.Response Body if it receives an error.
func (client JobClient) GetOutputSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetOutputResponder handles the response to the GetOutput request. The method always
// closes the http.Response Body.
func (client JobClient) GetOutputResponder(resp *http.Response) (result ReadCloser, err error) {
	result.Value = &resp.Body
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	result.Response = autorest.Response{Response: resp}
	return
}

// GetRunbookContent retrieve the runbook content of the job identified by job name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// jobName - the job name.
// clientRequestID - identifies this specific client request.
func (client JobClient) GetRunbookContent(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, clientRequestID string) (result ReadCloser, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobClient.GetRunbookContent")
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
		return result, validation.NewError("automation.JobClient", "GetRunbookContent", err.Error())
	}

	req, err := client.GetRunbookContentPreparer(ctx, resourceGroupName, automationAccountName, jobName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "GetRunbookContent", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRunbookContentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.JobClient", "GetRunbookContent", resp, "Failure sending request")
		return
	}

	result, err = client.GetRunbookContentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "GetRunbookContent", resp, "Failure responding to request")
	}

	return
}

// GetRunbookContentPreparer prepares the GetRunbookContent request.
func (client JobClient) GetRunbookContentPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, clientRequestID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobName":               autorest.Encode("path", jobName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}/runbookContent", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(clientRequestID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("clientRequestId", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetRunbookContentSender sends the GetRunbookContent request. The method will close the
// http.Response Body if it receives an error.
func (client JobClient) GetRunbookContentSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetRunbookContentResponder handles the response to the GetRunbookContent request. The method always
// closes the http.Response Body.
func (client JobClient) GetRunbookContentResponder(resp *http.Response) (result ReadCloser, err error) {
	result.Value = &resp.Body
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAutomationAccount retrieve a list of jobs.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// filter - the filter to apply on the operation.
// clientRequestID - identifies this specific client request.
func (client JobClient) ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, filter string, clientRequestID string) (result JobListResultV2Page, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobClient.ListByAutomationAccount")
		defer func() {
			sc := -1
			if result.jlrv.Response.Response != nil {
				sc = result.jlrv.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.JobClient", "ListByAutomationAccount", err.Error())
	}

	result.fn = client.listByAutomationAccountNextResults
	req, err := client.ListByAutomationAccountPreparer(ctx, resourceGroupName, automationAccountName, filter, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "ListByAutomationAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.jlrv.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.JobClient", "ListByAutomationAccount", resp, "Failure sending request")
		return
	}

	result.jlrv, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "ListByAutomationAccount", resp, "Failure responding to request")
	}
	if result.jlrv.hasNextLink() && result.jlrv.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByAutomationAccountPreparer prepares the ListByAutomationAccount request.
func (client JobClient) ListByAutomationAccountPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, filter string, clientRequestID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(clientRequestID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("clientRequestId", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAutomationAccountSender sends the ListByAutomationAccount request. The method will close the
// http.Response Body if it receives an error.
func (client JobClient) ListByAutomationAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByAutomationAccountResponder handles the response to the ListByAutomationAccount request. The method always
// closes the http.Response Body.
func (client JobClient) ListByAutomationAccountResponder(resp *http.Response) (result JobListResultV2, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByAutomationAccountNextResults retrieves the next set of results, if any.
func (client JobClient) listByAutomationAccountNextResults(ctx context.Context, lastResults JobListResultV2) (result JobListResultV2, err error) {
	req, err := lastResults.jobListResultV2Preparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.JobClient", "listByAutomationAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.JobClient", "listByAutomationAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "listByAutomationAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByAutomationAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobClient) ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string, filter string, clientRequestID string) (result JobListResultV2Iterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobClient.ListByAutomationAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByAutomationAccount(ctx, resourceGroupName, automationAccountName, filter, clientRequestID)
	return
}

// Resume resume the job identified by jobName.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// jobName - the job name.
// clientRequestID - identifies this specific client request.
func (client JobClient) Resume(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, clientRequestID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobClient.Resume")
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
		return result, validation.NewError("automation.JobClient", "Resume", err.Error())
	}

	req, err := client.ResumePreparer(ctx, resourceGroupName, automationAccountName, jobName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Resume", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResumeSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Resume", resp, "Failure sending request")
		return
	}

	result, err = client.ResumeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Resume", resp, "Failure responding to request")
	}

	return
}

// ResumePreparer prepares the Resume request.
func (client JobClient) ResumePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, clientRequestID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobName":               autorest.Encode("path", jobName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}/resume", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(clientRequestID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("clientRequestId", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResumeSender sends the Resume request. The method will close the
// http.Response Body if it receives an error.
func (client JobClient) ResumeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ResumeResponder handles the response to the Resume request. The method always
// closes the http.Response Body.
func (client JobClient) ResumeResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop stop the job identified by jobName.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// jobName - the job name.
// clientRequestID - identifies this specific client request.
func (client JobClient) Stop(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, clientRequestID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobClient.Stop")
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
		return result, validation.NewError("automation.JobClient", "Stop", err.Error())
	}

	req, err := client.StopPreparer(ctx, resourceGroupName, automationAccountName, jobName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Stop", nil, "Failure preparing request")
		return
	}

	resp, err := client.StopSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Stop", resp, "Failure sending request")
		return
	}

	result, err = client.StopResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Stop", resp, "Failure responding to request")
	}

	return
}

// StopPreparer prepares the Stop request.
func (client JobClient) StopPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, clientRequestID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobName":               autorest.Encode("path", jobName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(clientRequestID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("clientRequestId", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client JobClient) StopSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client JobClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Suspend suspend the job identified by job name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// jobName - the job name.
// clientRequestID - identifies this specific client request.
func (client JobClient) Suspend(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, clientRequestID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobClient.Suspend")
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
		return result, validation.NewError("automation.JobClient", "Suspend", err.Error())
	}

	req, err := client.SuspendPreparer(ctx, resourceGroupName, automationAccountName, jobName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Suspend", nil, "Failure preparing request")
		return
	}

	resp, err := client.SuspendSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Suspend", resp, "Failure sending request")
		return
	}

	result, err = client.SuspendResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobClient", "Suspend", resp, "Failure responding to request")
	}

	return
}

// SuspendPreparer prepares the Suspend request.
func (client JobClient) SuspendPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, clientRequestID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobName":               autorest.Encode("path", jobName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}/suspend", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(clientRequestID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("clientRequestId", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SuspendSender sends the Suspend request. The method will close the
// http.Response Body if it receives an error.
func (client JobClient) SuspendSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// SuspendResponder handles the response to the Suspend request. The method always
// closes the http.Response Body.
func (client JobClient) SuspendResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
