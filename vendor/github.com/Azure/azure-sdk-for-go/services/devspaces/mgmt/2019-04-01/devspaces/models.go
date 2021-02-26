package devspaces

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
const fqdn = "github.com/Azure/azure-sdk-for-go/services/devspaces/mgmt/2019-04-01/devspaces"

// ContainerHostMapping container host mapping object specifying the Container host resource ID and its
// associated Controller resource.
type ContainerHostMapping struct {
	autorest.Response `json:"-"`
	// ContainerHostResourceID - ARM ID of the Container Host resource
	ContainerHostResourceID *string `json:"containerHostResourceId,omitempty"`
	// MappedControllerResourceID - READ-ONLY; ARM ID of the mapped Controller resource
	MappedControllerResourceID *string `json:"mappedControllerResourceId,omitempty"`
}

// MarshalJSON is the custom marshaler for ContainerHostMapping.
func (chm ContainerHostMapping) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if chm.ContainerHostResourceID != nil {
		objectMap["containerHostResourceId"] = chm.ContainerHostResourceID
	}
	return json.Marshal(objectMap)
}

// Controller ...
type Controller struct {
	autorest.Response     `json:"-"`
	*ControllerProperties `json:"properties,omitempty"`
	Sku                   *Sku `json:"sku,omitempty"`
	// Tags - Tags for the Azure resource.
	Tags map[string]*string `json:"tags"`
	// Location - Region where the Azure resource is located.
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Controller.
func (c Controller) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if c.ControllerProperties != nil {
		objectMap["properties"] = c.ControllerProperties
	}
	if c.Sku != nil {
		objectMap["sku"] = c.Sku
	}
	if c.Tags != nil {
		objectMap["tags"] = c.Tags
	}
	if c.Location != nil {
		objectMap["location"] = c.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Controller struct.
func (c *Controller) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var controllerProperties ControllerProperties
				err = json.Unmarshal(*v, &controllerProperties)
				if err != nil {
					return err
				}
				c.ControllerProperties = &controllerProperties
			}
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				c.Sku = &sku
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

// ControllerConnectionDetails ...
type ControllerConnectionDetails struct {
	OrchestratorSpecificConnectionDetails BasicOrchestratorSpecificConnectionDetails `json:"orchestratorSpecificConnectionDetails,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for ControllerConnectionDetails struct.
func (ccd *ControllerConnectionDetails) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "orchestratorSpecificConnectionDetails":
			if v != nil {
				orchestratorSpecificConnectionDetails, err := unmarshalBasicOrchestratorSpecificConnectionDetails(*v)
				if err != nil {
					return err
				}
				ccd.OrchestratorSpecificConnectionDetails = orchestratorSpecificConnectionDetails
			}
		}
	}

	return nil
}

// ControllerConnectionDetailsList ...
type ControllerConnectionDetailsList struct {
	autorest.Response `json:"-"`
	// ConnectionDetailsList - List of Azure Dev Spaces Controller connection details.
	ConnectionDetailsList *[]ControllerConnectionDetails `json:"connectionDetailsList,omitempty"`
}

// ControllerList ...
type ControllerList struct {
	autorest.Response `json:"-"`
	// Value - List of Azure Dev Spaces Controllers.
	Value *[]Controller `json:"value,omitempty"`
	// NextLink - READ-ONLY; The URI that can be used to request the next page for list of Azure Dev Spaces Controllers.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for ControllerList.
func (cl ControllerList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cl.Value != nil {
		objectMap["value"] = cl.Value
	}
	return json.Marshal(objectMap)
}

// ControllerListIterator provides access to a complete listing of Controller values.
type ControllerListIterator struct {
	i    int
	page ControllerListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ControllerListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ControllerListIterator.NextWithContext")
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
func (iter *ControllerListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ControllerListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ControllerListIterator) Response() ControllerList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ControllerListIterator) Value() Controller {
	if !iter.page.NotDone() {
		return Controller{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ControllerListIterator type.
func NewControllerListIterator(page ControllerListPage) ControllerListIterator {
	return ControllerListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (cl ControllerList) IsEmpty() bool {
	return cl.Value == nil || len(*cl.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (cl ControllerList) hasNextLink() bool {
	return cl.NextLink != nil && len(*cl.NextLink) != 0
}

// controllerListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (cl ControllerList) controllerListPreparer(ctx context.Context) (*http.Request, error) {
	if !cl.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(cl.NextLink)))
}

// ControllerListPage contains a page of Controller values.
type ControllerListPage struct {
	fn func(context.Context, ControllerList) (ControllerList, error)
	cl ControllerList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ControllerListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ControllerListPage.NextWithContext")
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
func (page *ControllerListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ControllerListPage) NotDone() bool {
	return !page.cl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ControllerListPage) Response() ControllerList {
	return page.cl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ControllerListPage) Values() []Controller {
	if page.cl.IsEmpty() {
		return nil
	}
	return *page.cl.Value
}

// Creates a new instance of the ControllerListPage type.
func NewControllerListPage(cur ControllerList, getNextPage func(context.Context, ControllerList) (ControllerList, error)) ControllerListPage {
	return ControllerListPage{
		fn: getNextPage,
		cl: cur,
	}
}

// ControllerProperties ...
type ControllerProperties struct {
	// ProvisioningState - READ-ONLY; Provisioning state of the Azure Dev Spaces Controller. Possible values include: 'Succeeded', 'Failed', 'Canceled', 'Updating', 'Creating', 'Deleting', 'Deleted'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// HostSuffix - READ-ONLY; DNS suffix for public endpoints running in the Azure Dev Spaces Controller.
	HostSuffix *string `json:"hostSuffix,omitempty"`
	// DataPlaneFqdn - READ-ONLY; DNS name for accessing DataPlane services
	DataPlaneFqdn *string `json:"dataPlaneFqdn,omitempty"`
	// TargetContainerHostAPIServerFqdn - READ-ONLY; DNS of the target container host's API server
	TargetContainerHostAPIServerFqdn *string `json:"targetContainerHostApiServerFqdn,omitempty"`
	// TargetContainerHostResourceID - Resource ID of the target container host
	TargetContainerHostResourceID *string `json:"targetContainerHostResourceId,omitempty"`
	// TargetContainerHostCredentialsBase64 - Credentials of the target container host (base64).
	TargetContainerHostCredentialsBase64 *string `json:"targetContainerHostCredentialsBase64,omitempty"`
}

// MarshalJSON is the custom marshaler for ControllerProperties.
func (cp ControllerProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cp.TargetContainerHostResourceID != nil {
		objectMap["targetContainerHostResourceId"] = cp.TargetContainerHostResourceID
	}
	if cp.TargetContainerHostCredentialsBase64 != nil {
		objectMap["targetContainerHostCredentialsBase64"] = cp.TargetContainerHostCredentialsBase64
	}
	return json.Marshal(objectMap)
}

// ControllersCreateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ControllersCreateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(ControllersClient) (Controller, error)
}

// ControllersDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ControllersDeleteFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(ControllersClient) (autorest.Response, error)
}

// ControllerUpdateParameters parameters for updating an Azure Dev Spaces Controller.
type ControllerUpdateParameters struct {
	// Tags - Tags for the Azure Dev Spaces Controller.
	Tags                                  map[string]*string `json:"tags"`
	*ControllerUpdateParametersProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for ControllerUpdateParameters.
func (cup ControllerUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cup.Tags != nil {
		objectMap["tags"] = cup.Tags
	}
	if cup.ControllerUpdateParametersProperties != nil {
		objectMap["properties"] = cup.ControllerUpdateParametersProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ControllerUpdateParameters struct.
func (cup *ControllerUpdateParameters) UnmarshalJSON(body []byte) error {
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
				var controllerUpdateParametersProperties ControllerUpdateParametersProperties
				err = json.Unmarshal(*v, &controllerUpdateParametersProperties)
				if err != nil {
					return err
				}
				cup.ControllerUpdateParametersProperties = &controllerUpdateParametersProperties
			}
		}
	}

	return nil
}

// ControllerUpdateParametersProperties ...
type ControllerUpdateParametersProperties struct {
	// TargetContainerHostCredentialsBase64 - Credentials of the target container host (base64).
	TargetContainerHostCredentialsBase64 *string `json:"targetContainerHostCredentialsBase64,omitempty"`
}

// ErrorDetails ...
type ErrorDetails struct {
	// Code - READ-ONLY; Status code for the error.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; Error message describing the error in detail.
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; The target of the particular error.
	Target *string `json:"target,omitempty"`
}

// ErrorResponse error response indicates that the service is not able to process the incoming request. The
// reason is provided in the error message.
type ErrorResponse struct {
	// Error - The details of the error.
	Error *ErrorDetails `json:"error,omitempty"`
}

// KubernetesConnectionDetails contains information used to connect to a Kubernetes cluster
type KubernetesConnectionDetails struct {
	// KubeConfig - Gets the kubeconfig for the cluster.
	KubeConfig *string `json:"kubeConfig,omitempty"`
	// InstanceType - Possible values include: 'InstanceTypeOrchestratorSpecificConnectionDetails', 'InstanceTypeKubernetes'
	InstanceType InstanceType `json:"instanceType,omitempty"`
}

// MarshalJSON is the custom marshaler for KubernetesConnectionDetails.
func (kcd KubernetesConnectionDetails) MarshalJSON() ([]byte, error) {
	kcd.InstanceType = InstanceTypeKubernetes
	objectMap := make(map[string]interface{})
	if kcd.KubeConfig != nil {
		objectMap["kubeConfig"] = kcd.KubeConfig
	}
	if kcd.InstanceType != "" {
		objectMap["instanceType"] = kcd.InstanceType
	}
	return json.Marshal(objectMap)
}

// AsKubernetesConnectionDetails is the BasicOrchestratorSpecificConnectionDetails implementation for KubernetesConnectionDetails.
func (kcd KubernetesConnectionDetails) AsKubernetesConnectionDetails() (*KubernetesConnectionDetails, bool) {
	return &kcd, true
}

// AsOrchestratorSpecificConnectionDetails is the BasicOrchestratorSpecificConnectionDetails implementation for KubernetesConnectionDetails.
func (kcd KubernetesConnectionDetails) AsOrchestratorSpecificConnectionDetails() (*OrchestratorSpecificConnectionDetails, bool) {
	return nil, false
}

// AsBasicOrchestratorSpecificConnectionDetails is the BasicOrchestratorSpecificConnectionDetails implementation for KubernetesConnectionDetails.
func (kcd KubernetesConnectionDetails) AsBasicOrchestratorSpecificConnectionDetails() (BasicOrchestratorSpecificConnectionDetails, bool) {
	return &kcd, true
}

// ListConnectionDetailsParameters parameters for listing connection details of an Azure Dev Spaces
// Controller.
type ListConnectionDetailsParameters struct {
	// TargetContainerHostResourceID - Resource ID of the target container host mapped to the Azure Dev Spaces Controller.
	TargetContainerHostResourceID *string `json:"targetContainerHostResourceId,omitempty"`
}

// BasicOrchestratorSpecificConnectionDetails base class for types that supply values used to connect to container
// orchestrators
type BasicOrchestratorSpecificConnectionDetails interface {
	AsKubernetesConnectionDetails() (*KubernetesConnectionDetails, bool)
	AsOrchestratorSpecificConnectionDetails() (*OrchestratorSpecificConnectionDetails, bool)
}

// OrchestratorSpecificConnectionDetails base class for types that supply values used to connect to container
// orchestrators
type OrchestratorSpecificConnectionDetails struct {
	// InstanceType - Possible values include: 'InstanceTypeOrchestratorSpecificConnectionDetails', 'InstanceTypeKubernetes'
	InstanceType InstanceType `json:"instanceType,omitempty"`
}

func unmarshalBasicOrchestratorSpecificConnectionDetails(body []byte) (BasicOrchestratorSpecificConnectionDetails, error) {
	var m map[string]interface{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	switch m["instanceType"] {
	case string(InstanceTypeKubernetes):
		var kcd KubernetesConnectionDetails
		err := json.Unmarshal(body, &kcd)
		return kcd, err
	default:
		var oscd OrchestratorSpecificConnectionDetails
		err := json.Unmarshal(body, &oscd)
		return oscd, err
	}
}
func unmarshalBasicOrchestratorSpecificConnectionDetailsArray(body []byte) ([]BasicOrchestratorSpecificConnectionDetails, error) {
	var rawMessages []*json.RawMessage
	err := json.Unmarshal(body, &rawMessages)
	if err != nil {
		return nil, err
	}

	oscdArray := make([]BasicOrchestratorSpecificConnectionDetails, len(rawMessages))

	for index, rawMessage := range rawMessages {
		oscd, err := unmarshalBasicOrchestratorSpecificConnectionDetails(*rawMessage)
		if err != nil {
			return nil, err
		}
		oscdArray[index] = oscd
	}
	return oscdArray, nil
}

// MarshalJSON is the custom marshaler for OrchestratorSpecificConnectionDetails.
func (oscd OrchestratorSpecificConnectionDetails) MarshalJSON() ([]byte, error) {
	oscd.InstanceType = InstanceTypeOrchestratorSpecificConnectionDetails
	objectMap := make(map[string]interface{})
	if oscd.InstanceType != "" {
		objectMap["instanceType"] = oscd.InstanceType
	}
	return json.Marshal(objectMap)
}

// AsKubernetesConnectionDetails is the BasicOrchestratorSpecificConnectionDetails implementation for OrchestratorSpecificConnectionDetails.
func (oscd OrchestratorSpecificConnectionDetails) AsKubernetesConnectionDetails() (*KubernetesConnectionDetails, bool) {
	return nil, false
}

// AsOrchestratorSpecificConnectionDetails is the BasicOrchestratorSpecificConnectionDetails implementation for OrchestratorSpecificConnectionDetails.
func (oscd OrchestratorSpecificConnectionDetails) AsOrchestratorSpecificConnectionDetails() (*OrchestratorSpecificConnectionDetails, bool) {
	return &oscd, true
}

// AsBasicOrchestratorSpecificConnectionDetails is the BasicOrchestratorSpecificConnectionDetails implementation for OrchestratorSpecificConnectionDetails.
func (oscd OrchestratorSpecificConnectionDetails) AsBasicOrchestratorSpecificConnectionDetails() (BasicOrchestratorSpecificConnectionDetails, bool) {
	return &oscd, true
}

// Resource an Azure resource.
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
}

// ResourceProviderOperationDefinition ...
type ResourceProviderOperationDefinition struct {
	// Name - Resource provider operation name.
	Name    *string                           `json:"name,omitempty"`
	Display *ResourceProviderOperationDisplay `json:"display,omitempty"`
}

// ResourceProviderOperationDisplay ...
type ResourceProviderOperationDisplay struct {
	// Provider - Name of the resource provider.
	Provider *string `json:"provider,omitempty"`
	// Resource - Name of the resource type.
	Resource *string `json:"resource,omitempty"`
	// Operation - Name of the resource provider operation.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of the resource provider operation.
	Description *string `json:"description,omitempty"`
}

// ResourceProviderOperationList ...
type ResourceProviderOperationList struct {
	autorest.Response `json:"-"`
	// Value - Resource provider operations list.
	Value *[]ResourceProviderOperationDefinition `json:"value,omitempty"`
	// NextLink - READ-ONLY; The URI that can be used to request the next page for list of Azure operations.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceProviderOperationList.
func (rpol ResourceProviderOperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rpol.Value != nil {
		objectMap["value"] = rpol.Value
	}
	return json.Marshal(objectMap)
}

// ResourceProviderOperationListIterator provides access to a complete listing of
// ResourceProviderOperationDefinition values.
type ResourceProviderOperationListIterator struct {
	i    int
	page ResourceProviderOperationListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ResourceProviderOperationListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceProviderOperationListIterator.NextWithContext")
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
func (iter *ResourceProviderOperationListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ResourceProviderOperationListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ResourceProviderOperationListIterator) Response() ResourceProviderOperationList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ResourceProviderOperationListIterator) Value() ResourceProviderOperationDefinition {
	if !iter.page.NotDone() {
		return ResourceProviderOperationDefinition{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ResourceProviderOperationListIterator type.
func NewResourceProviderOperationListIterator(page ResourceProviderOperationListPage) ResourceProviderOperationListIterator {
	return ResourceProviderOperationListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rpol ResourceProviderOperationList) IsEmpty() bool {
	return rpol.Value == nil || len(*rpol.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (rpol ResourceProviderOperationList) hasNextLink() bool {
	return rpol.NextLink != nil && len(*rpol.NextLink) != 0
}

// resourceProviderOperationListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rpol ResourceProviderOperationList) resourceProviderOperationListPreparer(ctx context.Context) (*http.Request, error) {
	if !rpol.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rpol.NextLink)))
}

// ResourceProviderOperationListPage contains a page of ResourceProviderOperationDefinition values.
type ResourceProviderOperationListPage struct {
	fn   func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)
	rpol ResourceProviderOperationList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ResourceProviderOperationListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceProviderOperationListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.rpol)
		if err != nil {
			return err
		}
		page.rpol = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ResourceProviderOperationListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ResourceProviderOperationListPage) NotDone() bool {
	return !page.rpol.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ResourceProviderOperationListPage) Response() ResourceProviderOperationList {
	return page.rpol
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ResourceProviderOperationListPage) Values() []ResourceProviderOperationDefinition {
	if page.rpol.IsEmpty() {
		return nil
	}
	return *page.rpol.Value
}

// Creates a new instance of the ResourceProviderOperationListPage type.
func NewResourceProviderOperationListPage(cur ResourceProviderOperationList, getNextPage func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)) ResourceProviderOperationListPage {
	return ResourceProviderOperationListPage{
		fn:   getNextPage,
		rpol: cur,
	}
}

// Sku model representing SKU for Azure Dev Spaces Controller.
type Sku struct {
	// Name - The name of the SKU for Azure Dev Spaces Controller.
	Name *string `json:"name,omitempty"`
	// Tier - The tier of the SKU for Azure Dev Spaces Controller. Possible values include: 'Standard'
	Tier SkuTier `json:"tier,omitempty"`
}

// TrackedResource the resource model definition for a ARM tracked top level resource.
type TrackedResource struct {
	// Tags - Tags for the Azure resource.
	Tags map[string]*string `json:"tags"`
	// Location - Region where the Azure resource is located.
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource.
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
