package managedapplications

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

// ApplicationArtifactName enumerates the values for application artifact name.
type ApplicationArtifactName string

const (
	// Authorizations ...
	Authorizations ApplicationArtifactName = "Authorizations"
	// CustomRoleDefinition ...
	CustomRoleDefinition ApplicationArtifactName = "CustomRoleDefinition"
	// NotSpecified ...
	NotSpecified ApplicationArtifactName = "NotSpecified"
	// ViewDefinition ...
	ViewDefinition ApplicationArtifactName = "ViewDefinition"
)

// PossibleApplicationArtifactNameValues returns an array of possible values for the ApplicationArtifactName const type.
func PossibleApplicationArtifactNameValues() []ApplicationArtifactName {
	return []ApplicationArtifactName{Authorizations, CustomRoleDefinition, NotSpecified, ViewDefinition}
}

// ApplicationArtifactType enumerates the values for application artifact type.
type ApplicationArtifactType string

const (
	// ApplicationArtifactTypeCustom ...
	ApplicationArtifactTypeCustom ApplicationArtifactType = "Custom"
	// ApplicationArtifactTypeNotSpecified ...
	ApplicationArtifactTypeNotSpecified ApplicationArtifactType = "NotSpecified"
	// ApplicationArtifactTypeTemplate ...
	ApplicationArtifactTypeTemplate ApplicationArtifactType = "Template"
)

// PossibleApplicationArtifactTypeValues returns an array of possible values for the ApplicationArtifactType const type.
func PossibleApplicationArtifactTypeValues() []ApplicationArtifactType {
	return []ApplicationArtifactType{ApplicationArtifactTypeCustom, ApplicationArtifactTypeNotSpecified, ApplicationArtifactTypeTemplate}
}

// ApplicationDefinitionArtifactName enumerates the values for application definition artifact name.
type ApplicationDefinitionArtifactName string

const (
	// ApplicationDefinitionArtifactNameApplicationResourceTemplate ...
	ApplicationDefinitionArtifactNameApplicationResourceTemplate ApplicationDefinitionArtifactName = "ApplicationResourceTemplate"
	// ApplicationDefinitionArtifactNameCreateUIDefinition ...
	ApplicationDefinitionArtifactNameCreateUIDefinition ApplicationDefinitionArtifactName = "CreateUiDefinition"
	// ApplicationDefinitionArtifactNameMainTemplateParameters ...
	ApplicationDefinitionArtifactNameMainTemplateParameters ApplicationDefinitionArtifactName = "MainTemplateParameters"
	// ApplicationDefinitionArtifactNameNotSpecified ...
	ApplicationDefinitionArtifactNameNotSpecified ApplicationDefinitionArtifactName = "NotSpecified"
)

// PossibleApplicationDefinitionArtifactNameValues returns an array of possible values for the ApplicationDefinitionArtifactName const type.
func PossibleApplicationDefinitionArtifactNameValues() []ApplicationDefinitionArtifactName {
	return []ApplicationDefinitionArtifactName{ApplicationDefinitionArtifactNameApplicationResourceTemplate, ApplicationDefinitionArtifactNameCreateUIDefinition, ApplicationDefinitionArtifactNameMainTemplateParameters, ApplicationDefinitionArtifactNameNotSpecified}
}

// ApplicationLockLevel enumerates the values for application lock level.
type ApplicationLockLevel string

const (
	// CanNotDelete ...
	CanNotDelete ApplicationLockLevel = "CanNotDelete"
	// None ...
	None ApplicationLockLevel = "None"
	// ReadOnly ...
	ReadOnly ApplicationLockLevel = "ReadOnly"
)

// PossibleApplicationLockLevelValues returns an array of possible values for the ApplicationLockLevel const type.
func PossibleApplicationLockLevelValues() []ApplicationLockLevel {
	return []ApplicationLockLevel{CanNotDelete, None, ReadOnly}
}

// ApplicationManagementMode enumerates the values for application management mode.
type ApplicationManagementMode string

const (
	// ApplicationManagementModeManaged ...
	ApplicationManagementModeManaged ApplicationManagementMode = "Managed"
	// ApplicationManagementModeNotSpecified ...
	ApplicationManagementModeNotSpecified ApplicationManagementMode = "NotSpecified"
	// ApplicationManagementModeUnmanaged ...
	ApplicationManagementModeUnmanaged ApplicationManagementMode = "Unmanaged"
)

// PossibleApplicationManagementModeValues returns an array of possible values for the ApplicationManagementMode const type.
func PossibleApplicationManagementModeValues() []ApplicationManagementMode {
	return []ApplicationManagementMode{ApplicationManagementModeManaged, ApplicationManagementModeNotSpecified, ApplicationManagementModeUnmanaged}
}

// DeploymentMode enumerates the values for deployment mode.
type DeploymentMode string

const (
	// DeploymentModeComplete ...
	DeploymentModeComplete DeploymentMode = "Complete"
	// DeploymentModeIncremental ...
	DeploymentModeIncremental DeploymentMode = "Incremental"
	// DeploymentModeNotSpecified ...
	DeploymentModeNotSpecified DeploymentMode = "NotSpecified"
)

// PossibleDeploymentModeValues returns an array of possible values for the DeploymentMode const type.
func PossibleDeploymentModeValues() []DeploymentMode {
	return []DeploymentMode{DeploymentModeComplete, DeploymentModeIncremental, DeploymentModeNotSpecified}
}

// JitApprovalMode enumerates the values for jit approval mode.
type JitApprovalMode string

const (
	// JitApprovalModeAutoApprove ...
	JitApprovalModeAutoApprove JitApprovalMode = "AutoApprove"
	// JitApprovalModeManualApprove ...
	JitApprovalModeManualApprove JitApprovalMode = "ManualApprove"
	// JitApprovalModeNotSpecified ...
	JitApprovalModeNotSpecified JitApprovalMode = "NotSpecified"
)

// PossibleJitApprovalModeValues returns an array of possible values for the JitApprovalMode const type.
func PossibleJitApprovalModeValues() []JitApprovalMode {
	return []JitApprovalMode{JitApprovalModeAutoApprove, JitApprovalModeManualApprove, JitApprovalModeNotSpecified}
}

// JitApproverType enumerates the values for jit approver type.
type JitApproverType string

const (
	// Group ...
	Group JitApproverType = "group"
	// User ...
	User JitApproverType = "user"
)

// PossibleJitApproverTypeValues returns an array of possible values for the JitApproverType const type.
func PossibleJitApproverTypeValues() []JitApproverType {
	return []JitApproverType{Group, User}
}

// JitRequestState enumerates the values for jit request state.
type JitRequestState string

const (
	// JitRequestStateApproved ...
	JitRequestStateApproved JitRequestState = "Approved"
	// JitRequestStateCanceled ...
	JitRequestStateCanceled JitRequestState = "Canceled"
	// JitRequestStateDenied ...
	JitRequestStateDenied JitRequestState = "Denied"
	// JitRequestStateExpired ...
	JitRequestStateExpired JitRequestState = "Expired"
	// JitRequestStateFailed ...
	JitRequestStateFailed JitRequestState = "Failed"
	// JitRequestStateNotSpecified ...
	JitRequestStateNotSpecified JitRequestState = "NotSpecified"
	// JitRequestStatePending ...
	JitRequestStatePending JitRequestState = "Pending"
	// JitRequestStateTimeout ...
	JitRequestStateTimeout JitRequestState = "Timeout"
)

// PossibleJitRequestStateValues returns an array of possible values for the JitRequestState const type.
func PossibleJitRequestStateValues() []JitRequestState {
	return []JitRequestState{JitRequestStateApproved, JitRequestStateCanceled, JitRequestStateDenied, JitRequestStateExpired, JitRequestStateFailed, JitRequestStateNotSpecified, JitRequestStatePending, JitRequestStateTimeout}
}

// JitSchedulingType enumerates the values for jit scheduling type.
type JitSchedulingType string

const (
	// JitSchedulingTypeNotSpecified ...
	JitSchedulingTypeNotSpecified JitSchedulingType = "NotSpecified"
	// JitSchedulingTypeOnce ...
	JitSchedulingTypeOnce JitSchedulingType = "Once"
	// JitSchedulingTypeRecurring ...
	JitSchedulingTypeRecurring JitSchedulingType = "Recurring"
)

// PossibleJitSchedulingTypeValues returns an array of possible values for the JitSchedulingType const type.
func PossibleJitSchedulingTypeValues() []JitSchedulingType {
	return []JitSchedulingType{JitSchedulingTypeNotSpecified, JitSchedulingTypeOnce, JitSchedulingTypeRecurring}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateAccepted ...
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled ...
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateCreated ...
	ProvisioningStateCreated ProvisioningState = "Created"
	// ProvisioningStateCreating ...
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleted ...
	ProvisioningStateDeleted ProvisioningState = "Deleted"
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateNotSpecified ...
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	// ProvisioningStateReady ...
	ProvisioningStateReady ProvisioningState = "Ready"
	// ProvisioningStateRunning ...
	ProvisioningStateRunning ProvisioningState = "Running"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating ...
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateAccepted, ProvisioningStateCanceled, ProvisioningStateCreated, ProvisioningStateCreating, ProvisioningStateDeleted, ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStateNotSpecified, ProvisioningStateReady, ProvisioningStateRunning, ProvisioningStateSucceeded, ProvisioningStateUpdating}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// ResourceIdentityTypeNone ...
	ResourceIdentityTypeNone ResourceIdentityType = "None"
	// ResourceIdentityTypeSystemAssigned ...
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	// ResourceIdentityTypeSystemAssignedUserAssigned ...
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	// ResourceIdentityTypeUserAssigned ...
	ResourceIdentityTypeUserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{ResourceIdentityTypeNone, ResourceIdentityTypeSystemAssigned, ResourceIdentityTypeSystemAssignedUserAssigned, ResourceIdentityTypeUserAssigned}
}
