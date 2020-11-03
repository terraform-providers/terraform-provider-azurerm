package azurestackhci

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/azurestackhci/mgmt/2020-10-01/azurestackhci"

// AvailableOperations available operations of the service
type AvailableOperations struct {
	autorest.Response `json:"-"`
	// Value - Collection of available operation details
	Value *[]OperationDetail `json:"value,omitempty"`
	// NextLink - URL client should use to fetch the next page (per server side paging).
	// It's null for now, added for future use.
	NextLink *string `json:"nextLink,omitempty"`
}

// AzureEntityResource the resource model definition for a Azure Resource Manager resource with an etag.
type AzureEntityResource struct {
	// Etag - READ-ONLY; Resource Etag.
	Etag *string `json:"etag,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// Cluster cluster details.
type Cluster struct {
	autorest.Response `json:"-"`
	// ClusterProperties - Cluster properties.
	*ClusterProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Cluster.
func (c Cluster) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if c.ClusterProperties != nil {
		objectMap["properties"] = c.ClusterProperties
	}
	if c.Tags != nil {
		objectMap["tags"] = c.Tags
	}
	if c.Location != nil {
		objectMap["location"] = c.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Cluster struct.
func (c *Cluster) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var clusterProperties ClusterProperties
				err = json.Unmarshal(*v, &clusterProperties)
				if err != nil {
					return err
				}
				c.ClusterProperties = &clusterProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				c.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				c.Location = &location
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				c.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				c.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				c.Type = &typeVar
			}
		}
	}

	return nil
}

// ClusterList list of clusters.
type ClusterList struct {
	autorest.Response `json:"-"`
	// Value - List of clusters.
	Value *[]Cluster `json:"value,omitempty"`
	// NextLink - READ-ONLY; Link to the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for ClusterList.
func (cl ClusterList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cl.Value != nil {
		objectMap["value"] = cl.Value
	}
	return json.Marshal(objectMap)
}

// ClusterListIterator provides access to a complete listing of Cluster values.
type ClusterListIterator struct {
	i    int
	page ClusterListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ClusterListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClusterListIterator.NextWithContext")
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
func (iter *ClusterListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ClusterListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ClusterListIterator) Response() ClusterList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ClusterListIterator) Value() Cluster {
	if !iter.page.NotDone() {
		return Cluster{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ClusterListIterator type.
func NewClusterListIterator(page ClusterListPage) ClusterListIterator {
	return ClusterListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (cl ClusterList) IsEmpty() bool {
	return cl.Value == nil || len(*cl.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (cl ClusterList) hasNextLink() bool {
	return cl.NextLink != nil && len(*cl.NextLink) != 0
}

// clusterListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (cl ClusterList) clusterListPreparer(ctx context.Context) (*http.Request, error) {
	if !cl.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(cl.NextLink)))
}

// ClusterListPage contains a page of Cluster values.
type ClusterListPage struct {
	fn func(context.Context, ClusterList) (ClusterList, error)
	cl ClusterList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ClusterListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClusterListPage.NextWithContext")
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
func (page *ClusterListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ClusterListPage) NotDone() bool {
	return !page.cl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ClusterListPage) Response() ClusterList {
	return page.cl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ClusterListPage) Values() []Cluster {
	if page.cl.IsEmpty() {
		return nil
	}
	return *page.cl.Value
}

// Creates a new instance of the ClusterListPage type.
func NewClusterListPage(getNextPage func(context.Context, ClusterList) (ClusterList, error)) ClusterListPage {
	return ClusterListPage{fn: getNextPage}
}

// ClusterNode cluster node details.
type ClusterNode struct {
	// Name - READ-ONLY; Name of the cluster node.
	Name *string `json:"name,omitempty"`
	// ID - READ-ONLY; Id of the node in the cluster.
	ID *float64 `json:"id,omitempty"`
	// Manufacturer - READ-ONLY; Manufacturer of the cluster node hardware.
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Model - READ-ONLY; Model name of the cluster node hardware.
	Model *string `json:"model,omitempty"`
	// OsName - READ-ONLY; Operating system running on the cluster node.
	OsName *string `json:"osName,omitempty"`
	// OsVersion - READ-ONLY; Version of the operating system running on the cluster node.
	OsVersion *string `json:"osVersion,omitempty"`
	// SerialNumber - READ-ONLY; Immutable id of the cluster node.
	SerialNumber *string `json:"serialNumber,omitempty"`
	// CoreCount - READ-ONLY; Number of physical cores on the cluster node.
	CoreCount *float64 `json:"coreCount,omitempty"`
	// MemoryInGiB - READ-ONLY; Total available memory on the cluster node (in GiB).
	MemoryInGiB *float64 `json:"memoryInGiB,omitempty"`
}

// ClusterProperties cluster properties.
type ClusterProperties struct {
	// ProvisioningState - READ-ONLY; Provisioning state. Possible values include: 'Succeeded', 'Failed', 'Canceled', 'Accepted', 'Provisioning'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// Status - READ-ONLY; Status of the cluster agent. Possible values include: 'NotYetRegistered', 'ConnectedRecently', 'NotConnectedRecently', 'Disconnected', 'Error'
	Status Status `json:"status,omitempty"`
	// CloudID - READ-ONLY; Unique, immutable resource id.
	CloudID *string `json:"cloudId,omitempty"`
	// AadClientID - App id of cluster AAD identity.
	AadClientID *string `json:"aadClientId,omitempty"`
	// AadTenantID - Tenant id of cluster AAD identity.
	AadTenantID *string `json:"aadTenantId,omitempty"`
	// ReportedProperties - Properties reported by cluster agent.
	ReportedProperties *ClusterReportedProperties `json:"reportedProperties,omitempty"`
	// TrialDaysRemaining - READ-ONLY; Number of days remaining in the trial period.
	TrialDaysRemaining *float64 `json:"trialDaysRemaining,omitempty"`
	// BillingModel - READ-ONLY; Type of billing applied to the resource.
	BillingModel *string `json:"billingModel,omitempty"`
	// RegistrationTimestamp - READ-ONLY; First cluster sync timestamp.
	RegistrationTimestamp *date.Time `json:"registrationTimestamp,omitempty"`
	// LastSyncTimestamp - READ-ONLY; Most recent cluster sync timestamp.
	LastSyncTimestamp *date.Time `json:"lastSyncTimestamp,omitempty"`
	// LastBillingTimestamp - READ-ONLY; Most recent billing meter timestamp.
	LastBillingTimestamp *date.Time `json:"lastBillingTimestamp,omitempty"`
}

// MarshalJSON is the custom marshaler for ClusterProperties.
func (cp ClusterProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cp.AadClientID != nil {
		objectMap["aadClientId"] = cp.AadClientID
	}
	if cp.AadTenantID != nil {
		objectMap["aadTenantId"] = cp.AadTenantID
	}
	if cp.ReportedProperties != nil {
		objectMap["reportedProperties"] = cp.ReportedProperties
	}
	return json.Marshal(objectMap)
}

// ClusterReportedProperties properties reported by cluster agent.
type ClusterReportedProperties struct {
	// ClusterName - READ-ONLY; Name of the on-prem cluster connected to this resource.
	ClusterName *string `json:"clusterName,omitempty"`
	// ClusterID - READ-ONLY; Unique id generated by the on-prem cluster.
	ClusterID *string `json:"clusterId,omitempty"`
	// ClusterVersion - READ-ONLY; Version of the cluster software.
	ClusterVersion *string `json:"clusterVersion,omitempty"`
	// Nodes - READ-ONLY; List of nodes reported by the cluster.
	Nodes *[]ClusterNode `json:"nodes,omitempty"`
	// LastUpdated - READ-ONLY; Last time the cluster reported the data.
	LastUpdated *date.Time `json:"lastUpdated,omitempty"`
}

// ClusterUpdate cluster details to update.
type ClusterUpdate struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ClusterUpdate.
func (cu ClusterUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cu.Tags != nil {
		objectMap["tags"] = cu.Tags
	}
	return json.Marshal(objectMap)
}

// ErrorAdditionalInfo the resource management error additional info.
type ErrorAdditionalInfo struct {
	// Type - READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty"`
	// Info - READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty"`
}

// ErrorResponse the resource management error response.
type ErrorResponse struct {
	// Error - The error object.
	Error *ErrorResponseError `json:"error,omitempty"`
}

// ErrorResponseError the error object.
type ErrorResponseError struct {
	// Code - READ-ONLY; The error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; The error message.
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; The error target.
	Target *string `json:"target,omitempty"`
	// Details - READ-ONLY; The error details.
	Details *[]ErrorResponse `json:"details,omitempty"`
	// AdditionalInfo - READ-ONLY; The error additional info.
	AdditionalInfo *[]ErrorAdditionalInfo `json:"additionalInfo,omitempty"`
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
	// Properties - Properties of the operation
	Properties interface{} `json:"properties,omitempty"`
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

// ProxyResource the resource model definition for a ARM proxy resource. It will have everything other than
// required location and tags
type ProxyResource struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// Resource the resource model definition for a ARM tracked top level resource
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// TrackedResource the resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
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
