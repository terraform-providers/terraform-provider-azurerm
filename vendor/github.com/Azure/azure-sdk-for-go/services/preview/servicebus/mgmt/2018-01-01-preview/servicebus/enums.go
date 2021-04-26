package servicebus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccessRights enumerates the values for access rights.
type AccessRights string

const (
	// Listen ...
	Listen AccessRights = "Listen"
	// Manage ...
	Manage AccessRights = "Manage"
	// SendEnumValue ...
	SendEnumValue AccessRights = "Send"
)

// PossibleAccessRightsValues returns an array of possible values for the AccessRights const type.
func PossibleAccessRightsValues() []AccessRights {
	return []AccessRights{Listen, Manage, SendEnumValue}
}

// DefaultAction enumerates the values for default action.
type DefaultAction string

const (
	// Allow ...
	Allow DefaultAction = "Allow"
	// Deny ...
	Deny DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns an array of possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{Allow, Deny}
}

// EncodingCaptureDescription enumerates the values for encoding capture description.
type EncodingCaptureDescription string

const (
	// Avro ...
	Avro EncodingCaptureDescription = "Avro"
	// AvroDeflate ...
	AvroDeflate EncodingCaptureDescription = "AvroDeflate"
)

// PossibleEncodingCaptureDescriptionValues returns an array of possible values for the EncodingCaptureDescription const type.
func PossibleEncodingCaptureDescriptionValues() []EncodingCaptureDescription {
	return []EncodingCaptureDescription{Avro, AvroDeflate}
}

// EndPointProvisioningState enumerates the values for end point provisioning state.
type EndPointProvisioningState string

const (
	// Canceled ...
	Canceled EndPointProvisioningState = "Canceled"
	// Creating ...
	Creating EndPointProvisioningState = "Creating"
	// Deleting ...
	Deleting EndPointProvisioningState = "Deleting"
	// Failed ...
	Failed EndPointProvisioningState = "Failed"
	// Succeeded ...
	Succeeded EndPointProvisioningState = "Succeeded"
	// Updating ...
	Updating EndPointProvisioningState = "Updating"
)

// PossibleEndPointProvisioningStateValues returns an array of possible values for the EndPointProvisioningState const type.
func PossibleEndPointProvisioningStateValues() []EndPointProvisioningState {
	return []EndPointProvisioningState{Canceled, Creating, Deleting, Failed, Succeeded, Updating}
}

// EntityStatus enumerates the values for entity status.
type EntityStatus string

const (
	// EntityStatusActive ...
	EntityStatusActive EntityStatus = "Active"
	// EntityStatusCreating ...
	EntityStatusCreating EntityStatus = "Creating"
	// EntityStatusDeleting ...
	EntityStatusDeleting EntityStatus = "Deleting"
	// EntityStatusDisabled ...
	EntityStatusDisabled EntityStatus = "Disabled"
	// EntityStatusReceiveDisabled ...
	EntityStatusReceiveDisabled EntityStatus = "ReceiveDisabled"
	// EntityStatusRenaming ...
	EntityStatusRenaming EntityStatus = "Renaming"
	// EntityStatusRestoring ...
	EntityStatusRestoring EntityStatus = "Restoring"
	// EntityStatusSendDisabled ...
	EntityStatusSendDisabled EntityStatus = "SendDisabled"
	// EntityStatusUnknown ...
	EntityStatusUnknown EntityStatus = "Unknown"
)

// PossibleEntityStatusValues returns an array of possible values for the EntityStatus const type.
func PossibleEntityStatusValues() []EntityStatus {
	return []EntityStatus{EntityStatusActive, EntityStatusCreating, EntityStatusDeleting, EntityStatusDisabled, EntityStatusReceiveDisabled, EntityStatusRenaming, EntityStatusRestoring, EntityStatusSendDisabled, EntityStatusUnknown}
}

// FilterType enumerates the values for filter type.
type FilterType string

const (
	// FilterTypeCorrelationFilter ...
	FilterTypeCorrelationFilter FilterType = "CorrelationFilter"
	// FilterTypeSQLFilter ...
	FilterTypeSQLFilter FilterType = "SqlFilter"
)

// PossibleFilterTypeValues returns an array of possible values for the FilterType const type.
func PossibleFilterTypeValues() []FilterType {
	return []FilterType{FilterTypeCorrelationFilter, FilterTypeSQLFilter}
}

// IdentityType enumerates the values for identity type.
type IdentityType string

const (
	// SystemAssigned ...
	SystemAssigned IdentityType = "SystemAssigned"
)

// PossibleIdentityTypeValues returns an array of possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{SystemAssigned}
}

// IPAction enumerates the values for ip action.
type IPAction string

const (
	// Accept ...
	Accept IPAction = "Accept"
	// Reject ...
	Reject IPAction = "Reject"
)

// PossibleIPActionValues returns an array of possible values for the IPAction const type.
func PossibleIPActionValues() []IPAction {
	return []IPAction{Accept, Reject}
}

// KeySource enumerates the values for key source.
type KeySource string

const (
	// MicrosoftKeyVault ...
	MicrosoftKeyVault KeySource = "Microsoft.KeyVault"
)

// PossibleKeySourceValues returns an array of possible values for the KeySource const type.
func PossibleKeySourceValues() []KeySource {
	return []KeySource{MicrosoftKeyVault}
}

// KeyType enumerates the values for key type.
type KeyType string

const (
	// PrimaryKey ...
	PrimaryKey KeyType = "PrimaryKey"
	// SecondaryKey ...
	SecondaryKey KeyType = "SecondaryKey"
)

// PossibleKeyTypeValues returns an array of possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{PrimaryKey, SecondaryKey}
}

// NameSpaceType enumerates the values for name space type.
type NameSpaceType string

const (
	// EventHub ...
	EventHub NameSpaceType = "EventHub"
	// Messaging ...
	Messaging NameSpaceType = "Messaging"
	// Mixed ...
	Mixed NameSpaceType = "Mixed"
	// NotificationHub ...
	NotificationHub NameSpaceType = "NotificationHub"
	// Relay ...
	Relay NameSpaceType = "Relay"
)

// PossibleNameSpaceTypeValues returns an array of possible values for the NameSpaceType const type.
func PossibleNameSpaceTypeValues() []NameSpaceType {
	return []NameSpaceType{EventHub, Messaging, Mixed, NotificationHub, Relay}
}

// NetworkRuleIPAction enumerates the values for network rule ip action.
type NetworkRuleIPAction string

const (
	// NetworkRuleIPActionAllow ...
	NetworkRuleIPActionAllow NetworkRuleIPAction = "Allow"
)

// PossibleNetworkRuleIPActionValues returns an array of possible values for the NetworkRuleIPAction const type.
func PossibleNetworkRuleIPActionValues() []NetworkRuleIPAction {
	return []NetworkRuleIPAction{NetworkRuleIPActionAllow}
}

// PrivateLinkConnectionStatus enumerates the values for private link connection status.
type PrivateLinkConnectionStatus string

const (
	// Approved ...
	Approved PrivateLinkConnectionStatus = "Approved"
	// Disconnected ...
	Disconnected PrivateLinkConnectionStatus = "Disconnected"
	// Pending ...
	Pending PrivateLinkConnectionStatus = "Pending"
	// Rejected ...
	Rejected PrivateLinkConnectionStatus = "Rejected"
)

// PossiblePrivateLinkConnectionStatusValues returns an array of possible values for the PrivateLinkConnectionStatus const type.
func PossiblePrivateLinkConnectionStatusValues() []PrivateLinkConnectionStatus {
	return []PrivateLinkConnectionStatus{Approved, Disconnected, Pending, Rejected}
}

// ProvisioningStateDR enumerates the values for provisioning state dr.
type ProvisioningStateDR string

const (
	// ProvisioningStateDRAccepted ...
	ProvisioningStateDRAccepted ProvisioningStateDR = "Accepted"
	// ProvisioningStateDRFailed ...
	ProvisioningStateDRFailed ProvisioningStateDR = "Failed"
	// ProvisioningStateDRSucceeded ...
	ProvisioningStateDRSucceeded ProvisioningStateDR = "Succeeded"
)

// PossibleProvisioningStateDRValues returns an array of possible values for the ProvisioningStateDR const type.
func PossibleProvisioningStateDRValues() []ProvisioningStateDR {
	return []ProvisioningStateDR{ProvisioningStateDRAccepted, ProvisioningStateDRFailed, ProvisioningStateDRSucceeded}
}

// RoleDisasterRecovery enumerates the values for role disaster recovery.
type RoleDisasterRecovery string

const (
	// Primary ...
	Primary RoleDisasterRecovery = "Primary"
	// PrimaryNotReplicating ...
	PrimaryNotReplicating RoleDisasterRecovery = "PrimaryNotReplicating"
	// Secondary ...
	Secondary RoleDisasterRecovery = "Secondary"
)

// PossibleRoleDisasterRecoveryValues returns an array of possible values for the RoleDisasterRecovery const type.
func PossibleRoleDisasterRecoveryValues() []RoleDisasterRecovery {
	return []RoleDisasterRecovery{Primary, PrimaryNotReplicating, Secondary}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Basic ...
	Basic SkuName = "Basic"
	// Premium ...
	Premium SkuName = "Premium"
	// Standard ...
	Standard SkuName = "Standard"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{Basic, Premium, Standard}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// SkuTierBasic ...
	SkuTierBasic SkuTier = "Basic"
	// SkuTierPremium ...
	SkuTierPremium SkuTier = "Premium"
	// SkuTierStandard ...
	SkuTierStandard SkuTier = "Standard"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{SkuTierBasic, SkuTierPremium, SkuTierStandard}
}

// UnavailableReason enumerates the values for unavailable reason.
type UnavailableReason string

const (
	// InvalidName ...
	InvalidName UnavailableReason = "InvalidName"
	// NameInLockdown ...
	NameInLockdown UnavailableReason = "NameInLockdown"
	// NameInUse ...
	NameInUse UnavailableReason = "NameInUse"
	// None ...
	None UnavailableReason = "None"
	// SubscriptionIsDisabled ...
	SubscriptionIsDisabled UnavailableReason = "SubscriptionIsDisabled"
	// TooManyNamespaceInCurrentSubscription ...
	TooManyNamespaceInCurrentSubscription UnavailableReason = "TooManyNamespaceInCurrentSubscription"
)

// PossibleUnavailableReasonValues returns an array of possible values for the UnavailableReason const type.
func PossibleUnavailableReasonValues() []UnavailableReason {
	return []UnavailableReason{InvalidName, NameInLockdown, NameInUse, None, SubscriptionIsDisabled, TooManyNamespaceInCurrentSubscription}
}
