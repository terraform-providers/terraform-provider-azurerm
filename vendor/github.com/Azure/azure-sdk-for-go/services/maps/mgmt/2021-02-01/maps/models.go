package maps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/maps/mgmt/2021-02-01/maps"

// Account an Azure resource which represents access to a suite of Maps REST APIs.
type Account struct {
	autorest.Response `json:"-"`
	// Sku - The SKU of this account.
	Sku *Sku `json:"sku,omitempty"`
	// Kind - Get or Set Kind property. Possible values include: 'KindGen1', 'KindGen2'
	Kind Kind `json:"kind,omitempty"`
	// SystemData - READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData `json:"systemData,omitempty"`
	// Properties - The map account properties.
	Properties *AccountProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Account.
func (a Account) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.Sku != nil {
		objectMap["sku"] = a.Sku
	}
	if a.Kind != "" {
		objectMap["kind"] = a.Kind
	}
	if a.Properties != nil {
		objectMap["properties"] = a.Properties
	}
	if a.Tags != nil {
		objectMap["tags"] = a.Tags
	}
	if a.Location != nil {
		objectMap["location"] = a.Location
	}
	return json.Marshal(objectMap)
}

// AccountKeys the set of keys which can be used to access the Maps REST APIs. Two keys are provided for
// key rotation without interruption.
type AccountKeys struct {
	autorest.Response `json:"-"`
	// PrimaryKeyLastUpdated - READ-ONLY; The last updated date and time of the primary key.
	PrimaryKeyLastUpdated *string `json:"primaryKeyLastUpdated,omitempty"`
	// PrimaryKey - READ-ONLY; The primary key for accessing the Maps REST APIs.
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// SecondaryKey - READ-ONLY; The secondary key for accessing the Maps REST APIs.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
	// SecondaryKeyLastUpdated - READ-ONLY; The last updated date and time of the secondary key.
	SecondaryKeyLastUpdated *string `json:"secondaryKeyLastUpdated,omitempty"`
}

// AccountProperties additional Map account properties
type AccountProperties struct {
	// UniqueID - READ-ONLY; A unique identifier for the maps account
	UniqueID *string `json:"uniqueId,omitempty"`
	// DisableLocalAuth - Allows toggle functionality on Azure Policy to disable Azure Maps local authentication support. This will disable Shared Keys authentication from any usage.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`
	// ProvisioningState - READ-ONLY; the state of the provisioning.
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountProperties.
func (ap AccountProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ap.DisableLocalAuth != nil {
		objectMap["disableLocalAuth"] = ap.DisableLocalAuth
	}
	return json.Marshal(objectMap)
}

// Accounts a list of Maps Accounts.
type Accounts struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; a Maps Account.
	Value *[]Account `json:"value,omitempty"`
	// NextLink - URL client should use to fetch the next page (per server side paging).
	// It's null for now, added for future use.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for Accounts.
func (a Accounts) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.NextLink != nil {
		objectMap["nextLink"] = a.NextLink
	}
	return json.Marshal(objectMap)
}

// AccountsIterator provides access to a complete listing of Account values.
type AccountsIterator struct {
	i    int
	page AccountsPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AccountsIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountsIterator.NextWithContext")
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
func (iter *AccountsIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AccountsIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter AccountsIterator) Response() Accounts {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AccountsIterator) Value() Account {
	if !iter.page.NotDone() {
		return Account{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the AccountsIterator type.
func NewAccountsIterator(page AccountsPage) AccountsIterator {
	return AccountsIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (a Accounts) IsEmpty() bool {
	return a.Value == nil || len(*a.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (a Accounts) hasNextLink() bool {
	return a.NextLink != nil && len(*a.NextLink) != 0
}

// accountsPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (a Accounts) accountsPreparer(ctx context.Context) (*http.Request, error) {
	if !a.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(a.NextLink)))
}

// AccountsPage contains a page of Account values.
type AccountsPage struct {
	fn func(context.Context, Accounts) (Accounts, error)
	a  Accounts
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AccountsPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountsPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.a)
		if err != nil {
			return err
		}
		page.a = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *AccountsPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AccountsPage) NotDone() bool {
	return !page.a.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page AccountsPage) Response() Accounts {
	return page.a
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page AccountsPage) Values() []Account {
	if page.a.IsEmpty() {
		return nil
	}
	return *page.a.Value
}

// Creates a new instance of the AccountsPage type.
func NewAccountsPage(cur Accounts, getNextPage func(context.Context, Accounts) (Accounts, error)) AccountsPage {
	return AccountsPage{
		fn: getNextPage,
		a:  cur,
	}
}

// AccountUpdateParameters parameters used to update an existing Maps Account.
type AccountUpdateParameters struct {
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string `json:"tags"`
	// Kind - Get or Set Kind property. Possible values include: 'KindGen1', 'KindGen2'
	Kind Kind `json:"kind,omitempty"`
	// Sku - The SKU of this account.
	Sku *Sku `json:"sku,omitempty"`
	// AccountProperties - The map account properties.
	*AccountProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountUpdateParameters.
func (aup AccountUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if aup.Tags != nil {
		objectMap["tags"] = aup.Tags
	}
	if aup.Kind != "" {
		objectMap["kind"] = aup.Kind
	}
	if aup.Sku != nil {
		objectMap["sku"] = aup.Sku
	}
	if aup.AccountProperties != nil {
		objectMap["properties"] = aup.AccountProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for AccountUpdateParameters struct.
func (aup *AccountUpdateParameters) UnmarshalJSON(body []byte) error {
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
				aup.Tags = tags
			}
		case "kind":
			if v != nil {
				var kind Kind
				err = json.Unmarshal(*v, &kind)
				if err != nil {
					return err
				}
				aup.Kind = kind
			}
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				aup.Sku = &sku
			}
		case "properties":
			if v != nil {
				var accountProperties AccountProperties
				err = json.Unmarshal(*v, &accountProperties)
				if err != nil {
					return err
				}
				aup.AccountProperties = &accountProperties
			}
		}
	}

	return nil
}

// AzureEntityResource the resource model definition for an Azure Resource Manager resource with an etag.
type AzureEntityResource struct {
	// Etag - READ-ONLY; Resource Etag.
	Etag *string `json:"etag,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Creator an Azure resource which represents Maps Creator product and provides ability to manage private
// location data.
type Creator struct {
	autorest.Response `json:"-"`
	// Properties - The Creator resource properties.
	Properties *CreatorProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Creator.
func (c Creator) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if c.Properties != nil {
		objectMap["properties"] = c.Properties
	}
	if c.Tags != nil {
		objectMap["tags"] = c.Tags
	}
	if c.Location != nil {
		objectMap["location"] = c.Location
	}
	return json.Marshal(objectMap)
}

// CreatorList a list of Creator resources.
type CreatorList struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; a Creator account.
	Value *[]Creator `json:"value,omitempty"`
	// NextLink - URL client should use to fetch the next page (per server side paging).
	// It's null for now, added for future use.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for CreatorList.
func (cl CreatorList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cl.NextLink != nil {
		objectMap["nextLink"] = cl.NextLink
	}
	return json.Marshal(objectMap)
}

// CreatorListIterator provides access to a complete listing of Creator values.
type CreatorListIterator struct {
	i    int
	page CreatorListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *CreatorListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CreatorListIterator.NextWithContext")
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
func (iter *CreatorListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter CreatorListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter CreatorListIterator) Response() CreatorList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter CreatorListIterator) Value() Creator {
	if !iter.page.NotDone() {
		return Creator{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the CreatorListIterator type.
func NewCreatorListIterator(page CreatorListPage) CreatorListIterator {
	return CreatorListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (cl CreatorList) IsEmpty() bool {
	return cl.Value == nil || len(*cl.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (cl CreatorList) hasNextLink() bool {
	return cl.NextLink != nil && len(*cl.NextLink) != 0
}

// creatorListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (cl CreatorList) creatorListPreparer(ctx context.Context) (*http.Request, error) {
	if !cl.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(cl.NextLink)))
}

// CreatorListPage contains a page of Creator values.
type CreatorListPage struct {
	fn func(context.Context, CreatorList) (CreatorList, error)
	cl CreatorList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *CreatorListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CreatorListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.cl)
		if err != nil {
			return err
		}
		page.cl = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *CreatorListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page CreatorListPage) NotDone() bool {
	return !page.cl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page CreatorListPage) Response() CreatorList {
	return page.cl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page CreatorListPage) Values() []Creator {
	if page.cl.IsEmpty() {
		return nil
	}
	return *page.cl.Value
}

// Creates a new instance of the CreatorListPage type.
func NewCreatorListPage(cur CreatorList, getNextPage func(context.Context, CreatorList) (CreatorList, error)) CreatorListPage {
	return CreatorListPage{
		fn: getNextPage,
		cl: cur,
	}
}

// CreatorProperties creator resource properties
type CreatorProperties struct {
	// ProvisioningState - READ-ONLY; The state of the resource provisioning, terminal states: Succeeded, Failed, Canceled
	ProvisioningState *string `json:"provisioningState,omitempty"`
	// StorageUnits - The storage units to be allocated. Integer values from 1 to 100, inclusive.
	StorageUnits *int32 `json:"storageUnits,omitempty"`
}

// MarshalJSON is the custom marshaler for CreatorProperties.
func (cp CreatorProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cp.StorageUnits != nil {
		objectMap["storageUnits"] = cp.StorageUnits
	}
	return json.Marshal(objectMap)
}

// CreatorUpdateParameters parameters used to update an existing Creator resource.
type CreatorUpdateParameters struct {
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string `json:"tags"`
	// CreatorProperties - Creator resource properties.
	*CreatorProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for CreatorUpdateParameters.
func (cup CreatorUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cup.Tags != nil {
		objectMap["tags"] = cup.Tags
	}
	if cup.CreatorProperties != nil {
		objectMap["properties"] = cup.CreatorProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for CreatorUpdateParameters struct.
func (cup *CreatorUpdateParameters) UnmarshalJSON(body []byte) error {
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
				cup.Tags = tags
			}
		case "properties":
			if v != nil {
				var creatorProperties CreatorProperties
				err = json.Unmarshal(*v, &creatorProperties)
				if err != nil {
					return err
				}
				cup.CreatorProperties = &creatorProperties
			}
		}
	}

	return nil
}

// Dimension dimension of map account, for example API Category, Api Name, Result Type, and Response Code.
type Dimension struct {
	// Name - Display name of dimension.
	Name *string `json:"name,omitempty"`
	// DisplayName - Display name of dimension.
	DisplayName *string `json:"displayName,omitempty"`
}

// ErrorAdditionalInfo the resource management error additional info.
type ErrorAdditionalInfo struct {
	// Type - READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty"`
	// Info - READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty"`
}

// ErrorDetail the error detail.
type ErrorDetail struct {
	// Code - READ-ONLY; The error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; The error message.
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; The error target.
	Target *string `json:"target,omitempty"`
	// Details - READ-ONLY; The error details.
	Details *[]ErrorDetail `json:"details,omitempty"`
	// AdditionalInfo - READ-ONLY; The error additional info.
	AdditionalInfo *[]ErrorAdditionalInfo `json:"additionalInfo,omitempty"`
}

// ErrorResponse common error response for all Azure Resource Manager APIs to return error details for
// failed operations. (This also follows the OData error response format.).
type ErrorResponse struct {
	// Error - The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// KeySpecification whether the operation refers to the primary or secondary key.
type KeySpecification struct {
	// KeyType - Whether the operation refers to the primary or secondary key. Possible values include: 'KeyTypePrimary', 'KeyTypeSecondary'
	KeyType KeyType `json:"keyType,omitempty"`
}

// MetricSpecification metric specification of operation.
type MetricSpecification struct {
	// Name - Name of metric specification.
	Name *string `json:"name,omitempty"`
	// DisplayName - Display name of metric specification.
	DisplayName *string `json:"displayName,omitempty"`
	// DisplayDescription - Display description of metric specification.
	DisplayDescription *string `json:"displayDescription,omitempty"`
	// Unit - Unit could be Count.
	Unit *string `json:"unit,omitempty"`
	// Dimensions - Dimensions of map account.
	Dimensions *[]Dimension `json:"dimensions,omitempty"`
	// AggregationType - Aggregation type could be Average.
	AggregationType *string `json:"aggregationType,omitempty"`
	// FillGapWithZero - The property to decide fill gap with zero or not.
	FillGapWithZero *bool `json:"fillGapWithZero,omitempty"`
	// Category - The category this metric specification belong to, could be Capacity.
	Category *string `json:"category,omitempty"`
	// ResourceIDDimensionNameOverride - Account Resource Id.
	ResourceIDDimensionNameOverride *string `json:"resourceIdDimensionNameOverride,omitempty"`
}

// OperationDetail operation detail payload
type OperationDetail struct {
	// Name - Name of the operation
	Name *string `json:"name,omitempty"`
	// IsDataAction - Indicates whether the operation is a data action
	IsDataAction *bool `json:"isDataAction,omitempty"`
	// Display - Display of the operation
	Display *OperationDisplay `json:"display,omitempty"`
	// Origin - Origin of the operation
	Origin *string `json:"origin,omitempty"`
	// OperationProperties - Properties of the operation
	*OperationProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationDetail.
func (od OperationDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if od.Name != nil {
		objectMap["name"] = od.Name
	}
	if od.IsDataAction != nil {
		objectMap["isDataAction"] = od.IsDataAction
	}
	if od.Display != nil {
		objectMap["display"] = od.Display
	}
	if od.Origin != nil {
		objectMap["origin"] = od.Origin
	}
	if od.OperationProperties != nil {
		objectMap["properties"] = od.OperationProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for OperationDetail struct.
func (od *OperationDetail) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				od.Name = &name
			}
		case "isDataAction":
			if v != nil {
				var isDataAction bool
				err = json.Unmarshal(*v, &isDataAction)
				if err != nil {
					return err
				}
				od.IsDataAction = &isDataAction
			}
		case "display":
			if v != nil {
				var display OperationDisplay
				err = json.Unmarshal(*v, &display)
				if err != nil {
					return err
				}
				od.Display = &display
			}
		case "origin":
			if v != nil {
				var origin string
				err = json.Unmarshal(*v, &origin)
				if err != nil {
					return err
				}
				od.Origin = &origin
			}
		case "properties":
			if v != nil {
				var operationProperties OperationProperties
				err = json.Unmarshal(*v, &operationProperties)
				if err != nil {
					return err
				}
				od.OperationProperties = &operationProperties
			}
		}
	}

	return nil
}

// OperationDisplay operation display payload
type OperationDisplay struct {
	// Provider - Resource provider of the operation
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource of the operation
	Resource *string `json:"resource,omitempty"`
	// Operation - Localized friendly name for the operation
	Operation *string `json:"operation,omitempty"`
	// Description - Localized friendly description for the operation
	Description *string `json:"description,omitempty"`
}

// OperationProperties properties of operation, include metric specifications.
type OperationProperties struct {
	// ServiceSpecification - One property of operation, include metric specifications.
	ServiceSpecification *ServiceSpecification `json:"serviceSpecification,omitempty"`
}

// Operations the set of operations available for Maps.
type Operations struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; An operation available for Maps.
	Value *[]OperationDetail `json:"value,omitempty"`
	// NextLink - URL client should use to fetch the next page (per server side paging).
	// It's null for now, added for future use.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for Operations.
func (o Operations) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if o.NextLink != nil {
		objectMap["nextLink"] = o.NextLink
	}
	return json.Marshal(objectMap)
}

// OperationsIterator provides access to a complete listing of OperationDetail values.
type OperationsIterator struct {
	i    int
	page OperationsPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationsIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationsIterator.NextWithContext")
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
func (iter *OperationsIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationsIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationsIterator) Response() Operations {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationsIterator) Value() OperationDetail {
	if !iter.page.NotDone() {
		return OperationDetail{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationsIterator type.
func NewOperationsIterator(page OperationsPage) OperationsIterator {
	return OperationsIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (o Operations) IsEmpty() bool {
	return o.Value == nil || len(*o.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (o Operations) hasNextLink() bool {
	return o.NextLink != nil && len(*o.NextLink) != 0
}

// operationsPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (o Operations) operationsPreparer(ctx context.Context) (*http.Request, error) {
	if !o.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(o.NextLink)))
}

// OperationsPage contains a page of OperationDetail values.
type OperationsPage struct {
	fn func(context.Context, Operations) (Operations, error)
	o  Operations
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationsPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationsPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.o)
		if err != nil {
			return err
		}
		page.o = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationsPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationsPage) NotDone() bool {
	return !page.o.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationsPage) Response() Operations {
	return page.o
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationsPage) Values() []OperationDetail {
	if page.o.IsEmpty() {
		return nil
	}
	return *page.o.Value
}

// Creates a new instance of the OperationsPage type.
func NewOperationsPage(cur Operations, getNextPage func(context.Context, Operations) (Operations, error)) OperationsPage {
	return OperationsPage{
		fn: getNextPage,
		o:  cur,
	}
}

// ProxyResource the resource model definition for a Azure Resource Manager proxy resource. It will not
// have tags and a location
type ProxyResource struct {
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Resource common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// ServiceSpecification one property of operation, include metric specifications.
type ServiceSpecification struct {
	// MetricSpecifications - Metric specifications of operation.
	MetricSpecifications *[]MetricSpecification `json:"metricSpecifications,omitempty"`
}

// Sku the SKU of the Maps Account.
type Sku struct {
	// Name - The name of the SKU, in standard format (such as S0). Possible values include: 'NameS0', 'NameS1', 'NameG2'
	Name Name `json:"name,omitempty"`
	// Tier - READ-ONLY; Gets the sku tier. This is based on the SKU name.
	Tier *string `json:"tier,omitempty"`
}

// MarshalJSON is the custom marshaler for Sku.
func (s Sku) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if s.Name != "" {
		objectMap["name"] = s.Name
	}
	return json.Marshal(objectMap)
}

// SystemData metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// CreatedBy - The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedByType - The type of identity that created the resource. Possible values include: 'CreatedByTypeUser', 'CreatedByTypeApplication', 'CreatedByTypeManagedIdentity', 'CreatedByTypeKey'
	CreatedByType CreatedByType `json:"createdByType,omitempty"`
	// CreatedAt - The timestamp of resource creation (UTC).
	CreatedAt *date.Time `json:"createdAt,omitempty"`
	// LastModifiedBy - The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// LastModifiedByType - The type of identity that last modified the resource. Possible values include: 'CreatedByTypeUser', 'CreatedByTypeApplication', 'CreatedByTypeManagedIdentity', 'CreatedByTypeKey'
	LastModifiedByType CreatedByType `json:"lastModifiedByType,omitempty"`
	// LastModifiedAt - The timestamp of resource last modification (UTC)
	LastModifiedAt *date.Time `json:"lastModifiedAt,omitempty"`
}

// TrackedResource the resource model definition for an Azure Resource Manager tracked top level resource
// which has 'tags' and a 'location'
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	return json.Marshal(objectMap)
}
