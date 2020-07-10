package datashare

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

// ConsumerInvitationsClient is the creates a Microsoft.DataShare management client.
type ConsumerInvitationsClient struct {
	BaseClient
}

// NewConsumerInvitationsClient creates an instance of the ConsumerInvitationsClient client.
func NewConsumerInvitationsClient(subscriptionID string) ConsumerInvitationsClient {
	return NewConsumerInvitationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConsumerInvitationsClientWithBaseURI creates an instance of the ConsumerInvitationsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewConsumerInvitationsClientWithBaseURI(baseURI string, subscriptionID string) ConsumerInvitationsClient {
	return ConsumerInvitationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get an invitation
// Parameters:
// location - location of the invitation
// invitationID - an invitation id
func (client ConsumerInvitationsClient) Get(ctx context.Context, location string, invitationID string) (result ConsumerInvitation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConsumerInvitationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, location, invitationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.ConsumerInvitationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.ConsumerInvitationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.ConsumerInvitationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConsumerInvitationsClient) GetPreparer(ctx context.Context, location string, invitationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"invitationId": autorest.Encode("path", invitationID),
		"location":     autorest.Encode("path", location),
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.DataShare/locations/{location}/consumerInvitations/{invitationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConsumerInvitationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConsumerInvitationsClient) GetResponder(resp *http.Response) (result ConsumerInvitation, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListInvitations lists invitations
// Parameters:
// skipToken - the continuation token
func (client ConsumerInvitationsClient) ListInvitations(ctx context.Context, skipToken string) (result ConsumerInvitationListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConsumerInvitationsClient.ListInvitations")
		defer func() {
			sc := -1
			if result.cil.Response.Response != nil {
				sc = result.cil.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listInvitationsNextResults
	req, err := client.ListInvitationsPreparer(ctx, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.ConsumerInvitationsClient", "ListInvitations", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListInvitationsSender(req)
	if err != nil {
		result.cil.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.ConsumerInvitationsClient", "ListInvitations", resp, "Failure sending request")
		return
	}

	result.cil, err = client.ListInvitationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.ConsumerInvitationsClient", "ListInvitations", resp, "Failure responding to request")
	}

	return
}

// ListInvitationsPreparer prepares the ListInvitations request.
func (client ConsumerInvitationsClient) ListInvitationsPreparer(ctx context.Context, skipToken string) (*http.Request, error) {
	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.DataShare/ListInvitations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListInvitationsSender sends the ListInvitations request. The method will close the
// http.Response Body if it receives an error.
func (client ConsumerInvitationsClient) ListInvitationsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListInvitationsResponder handles the response to the ListInvitations request. The method always
// closes the http.Response Body.
func (client ConsumerInvitationsClient) ListInvitationsResponder(resp *http.Response) (result ConsumerInvitationList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listInvitationsNextResults retrieves the next set of results, if any.
func (client ConsumerInvitationsClient) listInvitationsNextResults(ctx context.Context, lastResults ConsumerInvitationList) (result ConsumerInvitationList, err error) {
	req, err := lastResults.consumerInvitationListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datashare.ConsumerInvitationsClient", "listInvitationsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListInvitationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datashare.ConsumerInvitationsClient", "listInvitationsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListInvitationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.ConsumerInvitationsClient", "listInvitationsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListInvitationsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ConsumerInvitationsClient) ListInvitationsComplete(ctx context.Context, skipToken string) (result ConsumerInvitationListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConsumerInvitationsClient.ListInvitations")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListInvitations(ctx, skipToken)
	return
}

// RejectInvitation reject an invitation
// Parameters:
// location - location of the invitation
// invitation - an invitation payload
func (client ConsumerInvitationsClient) RejectInvitation(ctx context.Context, location string, invitation ConsumerInvitation) (result ConsumerInvitation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConsumerInvitationsClient.RejectInvitation")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: invitation,
			Constraints: []validation.Constraint{{Target: "invitation.ConsumerInvitationProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "invitation.ConsumerInvitationProperties.InvitationID", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("datashare.ConsumerInvitationsClient", "RejectInvitation", err.Error())
	}

	req, err := client.RejectInvitationPreparer(ctx, location, invitation)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.ConsumerInvitationsClient", "RejectInvitation", nil, "Failure preparing request")
		return
	}

	resp, err := client.RejectInvitationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.ConsumerInvitationsClient", "RejectInvitation", resp, "Failure sending request")
		return
	}

	result, err = client.RejectInvitationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.ConsumerInvitationsClient", "RejectInvitation", resp, "Failure responding to request")
	}

	return
}

// RejectInvitationPreparer prepares the RejectInvitation request.
func (client ConsumerInvitationsClient) RejectInvitationPreparer(ctx context.Context, location string, invitation ConsumerInvitation) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location": autorest.Encode("path", location),
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.DataShare/locations/{location}/RejectInvitation", pathParameters),
		autorest.WithJSON(invitation),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RejectInvitationSender sends the RejectInvitation request. The method will close the
// http.Response Body if it receives an error.
func (client ConsumerInvitationsClient) RejectInvitationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RejectInvitationResponder handles the response to the RejectInvitation request. The method always
// closes the http.Response Body.
func (client ConsumerInvitationsClient) RejectInvitationResponder(resp *http.Response) (result ConsumerInvitation, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
