package locks

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
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01/locks"

// LockLevel enumerates the values for lock level.
type LockLevel string

const (
	// CanNotDelete ...
	CanNotDelete LockLevel = "CanNotDelete"
	// NotSpecified ...
	NotSpecified LockLevel = "NotSpecified"
	// ReadOnly ...
	ReadOnly LockLevel = "ReadOnly"
)

// PossibleLockLevelValues returns an array of possible values for the LockLevel const type.
func PossibleLockLevelValues() []LockLevel {
	return []LockLevel{CanNotDelete, NotSpecified, ReadOnly}
}

// ManagementLockListResult the list of locks.
type ManagementLockListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of locks.
	Value *[]ManagementLockObject `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ManagementLockListResultIterator provides access to a complete listing of ManagementLockObject values.
type ManagementLockListResultIterator struct {
	i    int
	page ManagementLockListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ManagementLockListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagementLockListResultIterator.NextWithContext")
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
func (iter *ManagementLockListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ManagementLockListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ManagementLockListResultIterator) Response() ManagementLockListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ManagementLockListResultIterator) Value() ManagementLockObject {
	if !iter.page.NotDone() {
		return ManagementLockObject{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ManagementLockListResultIterator type.
func NewManagementLockListResultIterator(page ManagementLockListResultPage) ManagementLockListResultIterator {
	return ManagementLockListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (mllr ManagementLockListResult) IsEmpty() bool {
	return mllr.Value == nil || len(*mllr.Value) == 0
}

// managementLockListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (mllr ManagementLockListResult) managementLockListResultPreparer(ctx context.Context) (*http.Request, error) {
	if mllr.NextLink == nil || len(to.String(mllr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(mllr.NextLink)))
}

// ManagementLockListResultPage contains a page of ManagementLockObject values.
type ManagementLockListResultPage struct {
	fn   func(context.Context, ManagementLockListResult) (ManagementLockListResult, error)
	mllr ManagementLockListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ManagementLockListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagementLockListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.mllr)
	if err != nil {
		return err
	}
	page.mllr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ManagementLockListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ManagementLockListResultPage) NotDone() bool {
	return !page.mllr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ManagementLockListResultPage) Response() ManagementLockListResult {
	return page.mllr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ManagementLockListResultPage) Values() []ManagementLockObject {
	if page.mllr.IsEmpty() {
		return nil
	}
	return *page.mllr.Value
}

// Creates a new instance of the ManagementLockListResultPage type.
func NewManagementLockListResultPage(getNextPage func(context.Context, ManagementLockListResult) (ManagementLockListResult, error)) ManagementLockListResultPage {
	return ManagementLockListResultPage{fn: getNextPage}
}

// ManagementLockObject the lock information.
type ManagementLockObject struct {
	autorest.Response `json:"-"`
	// ManagementLockProperties - The properties of the lock.
	*ManagementLockProperties `json:"properties,omitempty"`
	// ID - The resource ID of the lock.
	ID *string `json:"id,omitempty"`
	// Type - The resource type of the lock - Microsoft.Authorization/locks.
	Type *string `json:"type,omitempty"`
	// Name - The name of the lock.
	Name *string `json:"name,omitempty"`
}

// MarshalJSON is the custom marshaler for ManagementLockObject.
func (mlo ManagementLockObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mlo.ManagementLockProperties != nil {
		objectMap["properties"] = mlo.ManagementLockProperties
	}
	if mlo.ID != nil {
		objectMap["id"] = mlo.ID
	}
	if mlo.Type != nil {
		objectMap["type"] = mlo.Type
	}
	if mlo.Name != nil {
		objectMap["name"] = mlo.Name
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ManagementLockObject struct.
func (mlo *ManagementLockObject) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var managementLockProperties ManagementLockProperties
				err = json.Unmarshal(*v, &managementLockProperties)
				if err != nil {
					return err
				}
				mlo.ManagementLockProperties = &managementLockProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				mlo.ID = &ID
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				mlo.Type = &typeVar
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				mlo.Name = &name
			}
		}
	}

	return nil
}

// ManagementLockOwner lock owner properties.
type ManagementLockOwner struct {
	// ApplicationID - The application ID of the lock owner.
	ApplicationID *string `json:"applicationId,omitempty"`
}

// ManagementLockProperties the lock properties.
type ManagementLockProperties struct {
	// Level - The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it. Possible values include: 'NotSpecified', 'CanNotDelete', 'ReadOnly'
	Level LockLevel `json:"level,omitempty"`
	// Notes - Notes about the lock. Maximum of 512 characters.
	Notes *string `json:"notes,omitempty"`
	// Owners - The owners of the lock.
	Owners *[]ManagementLockOwner `json:"owners,omitempty"`
}

// Operation microsoft.Authorization operation
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.Authorization
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of the request to list Microsoft.Authorization operations. It contains a list
// of operations and a URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of Microsoft.Authorization operations.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
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

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
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
	next, err := page.fn(ctx, page.olr)
	if err != nil {
		return err
	}
	page.olr = next
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
