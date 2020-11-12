package iotcentral

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral"

// App the IoT Central application.
type App struct {
	autorest.Response `json:"-"`
	// AppProperties - The common properties of an IoT Central application.
	*AppProperties `json:"properties,omitempty"`
	// Sku - A valid instance SKU.
	Sku *AppSkuInfo `json:"sku,omitempty"`
	// ID - READ-ONLY; The ARM resource identifier.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The ARM resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The resource type.
	Type *string `json:"type,omitempty"`
	// Location - The resource location.
	Location *string `json:"location,omitempty"`
	// Tags - The resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for App.
func (a App) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.AppProperties != nil {
		objectMap["properties"] = a.AppProperties
	}
	if a.Sku != nil {
		objectMap["sku"] = a.Sku
	}
	if a.Location != nil {
		objectMap["location"] = a.Location
	}
	if a.Tags != nil {
		objectMap["tags"] = a.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for App struct.
func (a *App) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var appProperties AppProperties
				err = json.Unmarshal(*v, &appProperties)
				if err != nil {
					return err
				}
				a.AppProperties = &appProperties
			}
		case "sku":
			if v != nil {
				var sku AppSkuInfo
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				a.Sku = &sku
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				a.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				a.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				a.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				a.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				a.Tags = tags
			}
		}
	}

	return nil
}

// AppAvailabilityInfo the properties indicating whether a given IoT Central application name or subdomain is
// available.
type AppAvailabilityInfo struct {
	autorest.Response `json:"-"`
	// NameAvailable - READ-ONLY; The value which indicates whether the provided name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - READ-ONLY; The reason for unavailability.
	Reason *string `json:"reason,omitempty"`
	// Message - READ-ONLY; The detailed reason message.
	Message *string `json:"message,omitempty"`
}

// AppListResult a list of IoT Central Applications with a next link.
type AppListResult struct {
	autorest.Response `json:"-"`
	// NextLink - The link used to get the next page of IoT Central Applications.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - A list of IoT Central Applications.
	Value *[]App `json:"value,omitempty"`
}

// AppListResultIterator provides access to a complete listing of App values.
type AppListResultIterator struct {
	i    int
	page AppListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AppListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AppListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *AppListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AppListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter AppListResultIterator) Response() AppListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AppListResultIterator) Value() App {
	if !iter.page.NotDone() {
		return App{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the AppListResultIterator type.
func NewAppListResultIterator(page AppListResultPage) AppListResultIterator {
	return AppListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (alr AppListResult) IsEmpty() bool {
	return alr.Value == nil || len(*alr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (alr AppListResult) hasNextLink() bool {
	return alr.NextLink != nil && len(*alr.NextLink) != 0
}

// appListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (alr AppListResult) appListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !alr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(alr.NextLink)))
}

// AppListResultPage contains a page of App values.
type AppListResultPage struct {
	fn  func(context.Context, AppListResult) (AppListResult, error)
	alr AppListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AppListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AppListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.alr)
		if err != nil {
			return err
		}
		page.alr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *AppListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AppListResultPage) NotDone() bool {
	return !page.alr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page AppListResultPage) Response() AppListResult {
	return page.alr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page AppListResultPage) Values() []App {
	if page.alr.IsEmpty() {
		return nil
	}
	return *page.alr.Value
}

// Creates a new instance of the AppListResultPage type.
func NewAppListResultPage(getNextPage func(context.Context, AppListResult) (AppListResult, error)) AppListResultPage {
	return AppListResultPage{fn: getNextPage}
}

// AppPatch the description of the IoT Central application.
type AppPatch struct {
	// Tags - Instance tags
	Tags map[string]*string `json:"tags"`
	// Sku - A valid instance SKU.
	Sku *AppSkuInfo `json:"sku,omitempty"`
	// AppProperties - The common properties of an IoT Central application.
	*AppProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for AppPatch.
func (ap AppPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ap.Tags != nil {
		objectMap["tags"] = ap.Tags
	}
	if ap.Sku != nil {
		objectMap["sku"] = ap.Sku
	}
	if ap.AppProperties != nil {
		objectMap["properties"] = ap.AppProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for AppPatch struct.
func (ap *AppPatch) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				ap.Tags = tags
			}
		case "sku":
			if v != nil {
				var sku AppSkuInfo
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				ap.Sku = &sku
			}
		case "properties":
			if v != nil {
				var appProperties AppProperties
				err = json.Unmarshal(*v, &appProperties)
				if err != nil {
					return err
				}
				ap.AppProperties = &appProperties
			}
		}
	}

	return nil
}

// AppProperties the properties of an IoT Central application.
type AppProperties struct {
	// ApplicationID - READ-ONLY; The ID of the application.
	ApplicationID *string `json:"applicationId,omitempty"`
	// DisplayName - The display name of the application.
	DisplayName *string `json:"displayName,omitempty"`
	// Subdomain - The subdomain of the application.
	Subdomain *string `json:"subdomain,omitempty"`
	// Template - The ID of the application template, which is a blueprint that defines the characteristics and behaviors of an application. Optional; if not specified, defaults to a blank blueprint and allows the application to be defined from scratch.
	Template *string `json:"template,omitempty"`
}

// MarshalJSON is the custom marshaler for AppProperties.
func (ap AppProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ap.DisplayName != nil {
		objectMap["displayName"] = ap.DisplayName
	}
	if ap.Subdomain != nil {
		objectMap["subdomain"] = ap.Subdomain
	}
	if ap.Template != nil {
		objectMap["template"] = ap.Template
	}
	return json.Marshal(objectMap)
}

// AppsCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type AppsCreateOrUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *AppsCreateOrUpdateFuture) Result(client AppsClient) (a App, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("iotcentral.AppsCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if a.Response.Response, err = future.GetResult(sender); err == nil && a.Response.Response.StatusCode != http.StatusNoContent {
		a, err = client.CreateOrUpdateResponder(a.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "iotcentral.AppsCreateOrUpdateFuture", "Result", a.Response.Response, "Failure responding to request")
		}
	}
	return
}

// AppsDeleteFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type AppsDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *AppsDeleteFuture) Result(client AppsClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("iotcentral.AppsDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// AppSkuInfo information about the SKU of the IoT Central application.
type AppSkuInfo struct {
	// Name - The name of the SKU. Possible values include: 'F1', 'S1', 'ST0', 'ST1', 'ST2'
	Name AppSku `json:"name,omitempty"`
}

// AppsUpdateFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type AppsUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *AppsUpdateFuture) Result(client AppsClient) (a App, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("iotcentral.AppsUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if a.Response.Response, err = future.GetResult(sender); err == nil && a.Response.Response.StatusCode != http.StatusNoContent {
		a, err = client.UpdateResponder(a.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "iotcentral.AppsUpdateFuture", "Result", a.Response.Response, "Failure responding to request")
		}
	}
	return
}

// AppTemplate ioT Central Application Template.
type AppTemplate struct {
	// ManifestID - READ-ONLY; The ID of the template.
	ManifestID *string `json:"manifestId,omitempty"`
	// ManifestVersion - READ-ONLY; The version of the template.
	ManifestVersion *string `json:"manifestVersion,omitempty"`
	// AppTemplateName - READ-ONLY; The name of the template.
	AppTemplateName *string `json:"appTemplateName,omitempty"`
	// Title - READ-ONLY; The title of the template.
	Title *string `json:"title,omitempty"`
	// Order - READ-ONLY; The order of the template in the templates list.
	Order *float64 `json:"order,omitempty"`
	// Description - READ-ONLY; The description of the template.
	Description *string `json:"description,omitempty"`
}

// AppTemplatesResult a list of IoT Central Application Templates with a next link.
type AppTemplatesResult struct {
	autorest.Response `json:"-"`
	// NextLink - The link used to get the next page of IoT Central application templates.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - READ-ONLY; A list of IoT Central Application Templates.
	Value *[]AppTemplate `json:"value,omitempty"`
}

// MarshalJSON is the custom marshaler for AppTemplatesResult.
func (atr AppTemplatesResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if atr.NextLink != nil {
		objectMap["nextLink"] = atr.NextLink
	}
	return json.Marshal(objectMap)
}

// AppTemplatesResultIterator provides access to a complete listing of AppTemplate values.
type AppTemplatesResultIterator struct {
	i    int
	page AppTemplatesResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AppTemplatesResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AppTemplatesResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *AppTemplatesResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AppTemplatesResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter AppTemplatesResultIterator) Response() AppTemplatesResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AppTemplatesResultIterator) Value() AppTemplate {
	if !iter.page.NotDone() {
		return AppTemplate{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the AppTemplatesResultIterator type.
func NewAppTemplatesResultIterator(page AppTemplatesResultPage) AppTemplatesResultIterator {
	return AppTemplatesResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (atr AppTemplatesResult) IsEmpty() bool {
	return atr.Value == nil || len(*atr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (atr AppTemplatesResult) hasNextLink() bool {
	return atr.NextLink != nil && len(*atr.NextLink) != 0
}

// appTemplatesResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (atr AppTemplatesResult) appTemplatesResultPreparer(ctx context.Context) (*http.Request, error) {
	if !atr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(atr.NextLink)))
}

// AppTemplatesResultPage contains a page of AppTemplate values.
type AppTemplatesResultPage struct {
	fn  func(context.Context, AppTemplatesResult) (AppTemplatesResult, error)
	atr AppTemplatesResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AppTemplatesResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AppTemplatesResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.atr)
		if err != nil {
			return err
		}
		page.atr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *AppTemplatesResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AppTemplatesResultPage) NotDone() bool {
	return !page.atr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page AppTemplatesResultPage) Response() AppTemplatesResult {
	return page.atr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page AppTemplatesResultPage) Values() []AppTemplate {
	if page.atr.IsEmpty() {
		return nil
	}
	return *page.atr.Value
}

// Creates a new instance of the AppTemplatesResultPage type.
func NewAppTemplatesResultPage(getNextPage func(context.Context, AppTemplatesResult) (AppTemplatesResult, error)) AppTemplatesResultPage {
	return AppTemplatesResultPage{fn: getNextPage}
}

// CloudError error details.
type CloudError struct {
	// CloudErrorBody - Error response body.
	*CloudErrorBody `json:"error,omitempty"`
}

// MarshalJSON is the custom marshaler for CloudError.
func (ce CloudError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ce.CloudErrorBody != nil {
		objectMap["error"] = ce.CloudErrorBody
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for CloudError struct.
func (ce *CloudError) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "error":
			if v != nil {
				var cloudErrorBody CloudErrorBody
				err = json.Unmarshal(*v, &cloudErrorBody)
				if err != nil {
					return err
				}
				ce.CloudErrorBody = &cloudErrorBody
			}
		}
	}

	return nil
}

// CloudErrorBody details of error response.
type CloudErrorBody struct {
	// Code - READ-ONLY; The error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; The error message.
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; The target of the particular error.
	Target *string `json:"target,omitempty"`
	// Details - A list of additional details about the error.
	Details *[]CloudErrorBody `json:"details,omitempty"`
}

// MarshalJSON is the custom marshaler for CloudErrorBody.
func (ceb CloudErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ceb.Details != nil {
		objectMap["details"] = ceb.Details
	}
	return json.Marshal(objectMap)
}

// Operation ioT Central REST API operation
type Operation struct {
	// Name - READ-ONLY; Operation name: {provider}/{resource}/{read | write | action | delete}
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// MarshalJSON is the custom marshaler for Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if o.Display != nil {
		objectMap["display"] = o.Display
	}
	return json.Marshal(objectMap)
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - READ-ONLY; Service provider: Microsoft IoT Central
	Provider *string `json:"provider,omitempty"`
	// Resource - READ-ONLY; Resource Type: IoT Central
	Resource *string `json:"resource,omitempty"`
	// Operation - READ-ONLY; Name of the operation
	Operation *string `json:"operation,omitempty"`
	// Description - READ-ONLY; Friendly description for the operation,
	Description *string `json:"description,omitempty"`
}

// OperationInputs input values.
type OperationInputs struct {
	// Name - The name of the IoT Central application instance to check.
	Name *string `json:"name,omitempty"`
	// Type - The type of the IoT Central resource to query.
	Type *string `json:"type,omitempty"`
}

// OperationListResult a list of IoT Central operations. It contains a list of operations and a URL link to get
// the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// NextLink - The link used to get the next page of IoT Central description objects.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - READ-ONLY; A list of operations supported by the Microsoft.IoTCentral resource provider.
	Value *[]Operation `json:"value,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationListResult.
func (olr OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if olr.NextLink != nil {
		objectMap["nextLink"] = olr.NextLink
	}
	return json.Marshal(objectMap)
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListResultIterator type.
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return OperationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (olr OperationListResult) hasNextLink() bool {
	return olr.NextLink != nil && len(*olr.NextLink) != 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !olr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.olr)
		if err != nil {
			return err
		}
		page.olr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Creates a new instance of the OperationListResultPage type.
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{fn: getNextPage}
}

// Resource the common properties of an ARM resource.
type Resource struct {
	// ID - READ-ONLY; The ARM resource identifier.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The ARM resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The resource type.
	Type *string `json:"type,omitempty"`
	// Location - The resource location.
	Location *string `json:"location,omitempty"`
	// Tags - The resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}
