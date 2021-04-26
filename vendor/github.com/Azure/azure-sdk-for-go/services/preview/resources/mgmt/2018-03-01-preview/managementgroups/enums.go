package managementgroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// InheritedPermissions enumerates the values for inherited permissions.
type InheritedPermissions string

const (
	// Delete ...
	Delete InheritedPermissions = "delete"
	// Edit ...
	Edit InheritedPermissions = "edit"
	// Noaccess ...
	Noaccess InheritedPermissions = "noaccess"
	// View ...
	View InheritedPermissions = "view"
)

// PossibleInheritedPermissionsValues returns an array of possible values for the InheritedPermissions const type.
func PossibleInheritedPermissionsValues() []InheritedPermissions {
	return []InheritedPermissions{Delete, Edit, Noaccess, View}
}

// Permissions enumerates the values for permissions.
type Permissions string

const (
	// PermissionsDelete ...
	PermissionsDelete Permissions = "delete"
	// PermissionsEdit ...
	PermissionsEdit Permissions = "edit"
	// PermissionsNoaccess ...
	PermissionsNoaccess Permissions = "noaccess"
	// PermissionsView ...
	PermissionsView Permissions = "view"
)

// PossiblePermissionsValues returns an array of possible values for the Permissions const type.
func PossiblePermissionsValues() []Permissions {
	return []Permissions{PermissionsDelete, PermissionsEdit, PermissionsNoaccess, PermissionsView}
}

// Permissions1 enumerates the values for permissions 1.
type Permissions1 string

const (
	// Permissions1Delete ...
	Permissions1Delete Permissions1 = "delete"
	// Permissions1Edit ...
	Permissions1Edit Permissions1 = "edit"
	// Permissions1Noaccess ...
	Permissions1Noaccess Permissions1 = "noaccess"
	// Permissions1View ...
	Permissions1View Permissions1 = "view"
)

// PossiblePermissions1Values returns an array of possible values for the Permissions1 const type.
func PossiblePermissions1Values() []Permissions1 {
	return []Permissions1{Permissions1Delete, Permissions1Edit, Permissions1Noaccess, Permissions1View}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Updating}
}

// Reason enumerates the values for reason.
type Reason string

const (
	// AlreadyExists ...
	AlreadyExists Reason = "AlreadyExists"
	// Invalid ...
	Invalid Reason = "Invalid"
)

// PossibleReasonValues returns an array of possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{AlreadyExists, Invalid}
}

// Status enumerates the values for status.
type Status string

const (
	// Cancelled ...
	Cancelled Status = "Cancelled"
	// Completed ...
	Completed Status = "Completed"
	// Failed ...
	Failed Status = "Failed"
	// NotStarted ...
	NotStarted Status = "NotStarted"
	// NotStartedButGroupsExist ...
	NotStartedButGroupsExist Status = "NotStartedButGroupsExist"
	// Started ...
	Started Status = "Started"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{Cancelled, Completed, Failed, NotStarted, NotStartedButGroupsExist, Started}
}

// Type enumerates the values for type.
type Type string

const (
	// ProvidersMicrosoftManagementmanagementGroups ...
	ProvidersMicrosoftManagementmanagementGroups Type = "/providers/Microsoft.Management/managementGroups"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{ProvidersMicrosoftManagementmanagementGroups}
}

// Type1 enumerates the values for type 1.
type Type1 string

const (
	// Type1ProvidersMicrosoftManagementmanagementGroups ...
	Type1ProvidersMicrosoftManagementmanagementGroups Type1 = "/providers/Microsoft.Management/managementGroups"
	// Type1Subscriptions ...
	Type1Subscriptions Type1 = "/subscriptions"
)

// PossibleType1Values returns an array of possible values for the Type1 const type.
func PossibleType1Values() []Type1 {
	return []Type1{Type1ProvidersMicrosoftManagementmanagementGroups, Type1Subscriptions}
}

// Type2 enumerates the values for type 2.
type Type2 string

const (
	// Type2ProvidersMicrosoftManagementmanagementGroups ...
	Type2ProvidersMicrosoftManagementmanagementGroups Type2 = "/providers/Microsoft.Management/managementGroups"
	// Type2Subscriptions ...
	Type2Subscriptions Type2 = "/subscriptions"
)

// PossibleType2Values returns an array of possible values for the Type2 const type.
func PossibleType2Values() []Type2 {
	return []Type2{Type2ProvidersMicrosoftManagementmanagementGroups, Type2Subscriptions}
}
