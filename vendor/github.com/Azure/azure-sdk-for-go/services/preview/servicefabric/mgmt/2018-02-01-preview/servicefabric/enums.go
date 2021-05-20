package servicefabric

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ArmUpgradeFailureAction enumerates the values for arm upgrade failure action.
type ArmUpgradeFailureAction string

const (
	// Manual Indicates that a manual repair will need to be performed by the administrator if the upgrade
	// fails. Service Fabric will not proceed to the next upgrade domain automatically.
	Manual ArmUpgradeFailureAction = "Manual"
	// Rollback Indicates that a rollback of the upgrade will be performed by Service Fabric if the upgrade
	// fails.
	Rollback ArmUpgradeFailureAction = "Rollback"
)

// PossibleArmUpgradeFailureActionValues returns an array of possible values for the ArmUpgradeFailureAction const type.
func PossibleArmUpgradeFailureActionValues() []ArmUpgradeFailureAction {
	return []ArmUpgradeFailureAction{Manual, Rollback}
}

// ClusterState enumerates the values for cluster state.
type ClusterState string

const (
	// AutoScale ...
	AutoScale ClusterState = "AutoScale"
	// BaselineUpgrade ...
	BaselineUpgrade ClusterState = "BaselineUpgrade"
	// Deploying ...
	Deploying ClusterState = "Deploying"
	// EnforcingClusterVersion ...
	EnforcingClusterVersion ClusterState = "EnforcingClusterVersion"
	// Ready ...
	Ready ClusterState = "Ready"
	// UpdatingInfrastructure ...
	UpdatingInfrastructure ClusterState = "UpdatingInfrastructure"
	// UpdatingUserCertificate ...
	UpdatingUserCertificate ClusterState = "UpdatingUserCertificate"
	// UpdatingUserConfiguration ...
	UpdatingUserConfiguration ClusterState = "UpdatingUserConfiguration"
	// UpgradeServiceUnreachable ...
	UpgradeServiceUnreachable ClusterState = "UpgradeServiceUnreachable"
	// WaitingForNodes ...
	WaitingForNodes ClusterState = "WaitingForNodes"
)

// PossibleClusterStateValues returns an array of possible values for the ClusterState const type.
func PossibleClusterStateValues() []ClusterState {
	return []ClusterState{AutoScale, BaselineUpgrade, Deploying, EnforcingClusterVersion, Ready, UpdatingInfrastructure, UpdatingUserCertificate, UpdatingUserConfiguration, UpgradeServiceUnreachable, WaitingForNodes}
}

// DurabilityLevel enumerates the values for durability level.
type DurabilityLevel string

const (
	// Bronze ...
	Bronze DurabilityLevel = "Bronze"
	// Gold ...
	Gold DurabilityLevel = "Gold"
	// Silver ...
	Silver DurabilityLevel = "Silver"
)

// PossibleDurabilityLevelValues returns an array of possible values for the DurabilityLevel const type.
func PossibleDurabilityLevelValues() []DurabilityLevel {
	return []DurabilityLevel{Bronze, Gold, Silver}
}

// Environment enumerates the values for environment.
type Environment string

const (
	// Linux ...
	Linux Environment = "Linux"
	// Windows ...
	Windows Environment = "Windows"
)

// PossibleEnvironmentValues returns an array of possible values for the Environment const type.
func PossibleEnvironmentValues() []Environment {
	return []Environment{Linux, Windows}
}

// MoveCost enumerates the values for move cost.
type MoveCost string

const (
	// High Specifies the move cost of the service as High. The value is 3.
	High MoveCost = "High"
	// Low Specifies the move cost of the service as Low. The value is 1.
	Low MoveCost = "Low"
	// Medium Specifies the move cost of the service as Medium. The value is 2.
	Medium MoveCost = "Medium"
	// Zero Zero move cost. This value is zero.
	Zero MoveCost = "Zero"
)

// PossibleMoveCostValues returns an array of possible values for the MoveCost const type.
func PossibleMoveCostValues() []MoveCost {
	return []MoveCost{High, Low, Medium, Zero}
}

// PartitionScheme enumerates the values for partition scheme.
type PartitionScheme string

const (
	// Invalid Indicates the partition kind is invalid. All Service Fabric enumerations have the invalid type.
	// The value is zero.
	Invalid PartitionScheme = "Invalid"
	// Named Indicates that the partition is based on string names, and is a NamedPartitionSchemeDescription
	// object. The value is 3
	Named PartitionScheme = "Named"
	// Singleton Indicates that the partition is based on string names, and is a
	// SingletonPartitionSchemeDescription object, The value is 1.
	Singleton PartitionScheme = "Singleton"
	// UniformInt64Range Indicates that the partition is based on Int64 key ranges, and is a
	// UniformInt64RangePartitionSchemeDescription object. The value is 2.
	UniformInt64Range PartitionScheme = "UniformInt64Range"
)

// PossiblePartitionSchemeValues returns an array of possible values for the PartitionScheme const type.
func PossiblePartitionSchemeValues() []PartitionScheme {
	return []PartitionScheme{Invalid, Named, Singleton, UniformInt64Range}
}

// PartitionSchemeBasicPartitionSchemeDescription enumerates the values for partition scheme basic partition
// scheme description.
type PartitionSchemeBasicPartitionSchemeDescription string

const (
	// PartitionSchemeNamed ...
	PartitionSchemeNamed PartitionSchemeBasicPartitionSchemeDescription = "Named"
	// PartitionSchemePartitionSchemeDescription ...
	PartitionSchemePartitionSchemeDescription PartitionSchemeBasicPartitionSchemeDescription = "PartitionSchemeDescription"
	// PartitionSchemeSingleton ...
	PartitionSchemeSingleton PartitionSchemeBasicPartitionSchemeDescription = "Singleton"
	// PartitionSchemeUniformInt64Range ...
	PartitionSchemeUniformInt64Range PartitionSchemeBasicPartitionSchemeDescription = "UniformInt64Range"
)

// PossiblePartitionSchemeBasicPartitionSchemeDescriptionValues returns an array of possible values for the PartitionSchemeBasicPartitionSchemeDescription const type.
func PossiblePartitionSchemeBasicPartitionSchemeDescriptionValues() []PartitionSchemeBasicPartitionSchemeDescription {
	return []PartitionSchemeBasicPartitionSchemeDescription{PartitionSchemeNamed, PartitionSchemePartitionSchemeDescription, PartitionSchemeSingleton, PartitionSchemeUniformInt64Range}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Canceled, Failed, Succeeded, Updating}
}

// ReliabilityLevel enumerates the values for reliability level.
type ReliabilityLevel string

const (
	// ReliabilityLevelBronze ...
	ReliabilityLevelBronze ReliabilityLevel = "Bronze"
	// ReliabilityLevelGold ...
	ReliabilityLevelGold ReliabilityLevel = "Gold"
	// ReliabilityLevelNone ...
	ReliabilityLevelNone ReliabilityLevel = "None"
	// ReliabilityLevelPlatinum ...
	ReliabilityLevelPlatinum ReliabilityLevel = "Platinum"
	// ReliabilityLevelSilver ...
	ReliabilityLevelSilver ReliabilityLevel = "Silver"
)

// PossibleReliabilityLevelValues returns an array of possible values for the ReliabilityLevel const type.
func PossibleReliabilityLevelValues() []ReliabilityLevel {
	return []ReliabilityLevel{ReliabilityLevelBronze, ReliabilityLevelGold, ReliabilityLevelNone, ReliabilityLevelPlatinum, ReliabilityLevelSilver}
}

// ReliabilityLevel1 enumerates the values for reliability level 1.
type ReliabilityLevel1 string

const (
	// ReliabilityLevel1Bronze ...
	ReliabilityLevel1Bronze ReliabilityLevel1 = "Bronze"
	// ReliabilityLevel1Gold ...
	ReliabilityLevel1Gold ReliabilityLevel1 = "Gold"
	// ReliabilityLevel1None ...
	ReliabilityLevel1None ReliabilityLevel1 = "None"
	// ReliabilityLevel1Platinum ...
	ReliabilityLevel1Platinum ReliabilityLevel1 = "Platinum"
	// ReliabilityLevel1Silver ...
	ReliabilityLevel1Silver ReliabilityLevel1 = "Silver"
)

// PossibleReliabilityLevel1Values returns an array of possible values for the ReliabilityLevel1 const type.
func PossibleReliabilityLevel1Values() []ReliabilityLevel1 {
	return []ReliabilityLevel1{ReliabilityLevel1Bronze, ReliabilityLevel1Gold, ReliabilityLevel1None, ReliabilityLevel1Platinum, ReliabilityLevel1Silver}
}

// ServiceCorrelationScheme enumerates the values for service correlation scheme.
type ServiceCorrelationScheme string

const (
	// ServiceCorrelationSchemeAffinity Indicates that this service has an affinity relationship with another
	// service. Provided for backwards compatibility, consider preferring the Aligned or NonAlignedAffinity
	// options. The value is 1.
	ServiceCorrelationSchemeAffinity ServiceCorrelationScheme = "Affinity"
	// ServiceCorrelationSchemeAlignedAffinity Aligned affinity ensures that the primaries of the partitions of
	// the affinitized services are collocated on the same nodes. This is the default and is the same as
	// selecting the Affinity scheme. The value is 2.
	ServiceCorrelationSchemeAlignedAffinity ServiceCorrelationScheme = "AlignedAffinity"
	// ServiceCorrelationSchemeInvalid An invalid correlation scheme. Cannot be used. The value is zero.
	ServiceCorrelationSchemeInvalid ServiceCorrelationScheme = "Invalid"
	// ServiceCorrelationSchemeNonAlignedAffinity Non-Aligned affinity guarantees that all replicas of each
	// service will be placed on the same nodes. Unlike Aligned Affinity, this does not guarantee that replicas
	// of particular role will be collocated. The value is 3.
	ServiceCorrelationSchemeNonAlignedAffinity ServiceCorrelationScheme = "NonAlignedAffinity"
)

// PossibleServiceCorrelationSchemeValues returns an array of possible values for the ServiceCorrelationScheme const type.
func PossibleServiceCorrelationSchemeValues() []ServiceCorrelationScheme {
	return []ServiceCorrelationScheme{ServiceCorrelationSchemeAffinity, ServiceCorrelationSchemeAlignedAffinity, ServiceCorrelationSchemeInvalid, ServiceCorrelationSchemeNonAlignedAffinity}
}

// ServiceKind enumerates the values for service kind.
type ServiceKind string

const (
	// ServiceKindInvalid Indicates the service kind is invalid. All Service Fabric enumerations have the
	// invalid type. The value is zero.
	ServiceKindInvalid ServiceKind = "Invalid"
	// ServiceKindStateful Uses Service Fabric to make its state or part of its state highly available and
	// reliable. The value is 2.
	ServiceKindStateful ServiceKind = "Stateful"
	// ServiceKindStateless Does not use Service Fabric to make its state highly available or reliable. The
	// value is 1.
	ServiceKindStateless ServiceKind = "Stateless"
)

// PossibleServiceKindValues returns an array of possible values for the ServiceKind const type.
func PossibleServiceKindValues() []ServiceKind {
	return []ServiceKind{ServiceKindInvalid, ServiceKindStateful, ServiceKindStateless}
}

// ServiceKindBasicServiceResourceProperties enumerates the values for service kind basic service resource
// properties.
type ServiceKindBasicServiceResourceProperties string

const (
	// ServiceKindServiceResourceProperties ...
	ServiceKindServiceResourceProperties ServiceKindBasicServiceResourceProperties = "ServiceResourceProperties"
	// ServiceKindStateful1 ...
	ServiceKindStateful1 ServiceKindBasicServiceResourceProperties = "Stateful"
	// ServiceKindStateless1 ...
	ServiceKindStateless1 ServiceKindBasicServiceResourceProperties = "Stateless"
)

// PossibleServiceKindBasicServiceResourcePropertiesValues returns an array of possible values for the ServiceKindBasicServiceResourceProperties const type.
func PossibleServiceKindBasicServiceResourcePropertiesValues() []ServiceKindBasicServiceResourceProperties {
	return []ServiceKindBasicServiceResourceProperties{ServiceKindServiceResourceProperties, ServiceKindStateful1, ServiceKindStateless1}
}

// ServiceKindBasicServiceResourceUpdateProperties enumerates the values for service kind basic service
// resource update properties.
type ServiceKindBasicServiceResourceUpdateProperties string

const (
	// ServiceKindBasicServiceResourceUpdatePropertiesServiceKindServiceResourceUpdateProperties ...
	ServiceKindBasicServiceResourceUpdatePropertiesServiceKindServiceResourceUpdateProperties ServiceKindBasicServiceResourceUpdateProperties = "ServiceResourceUpdateProperties"
	// ServiceKindBasicServiceResourceUpdatePropertiesServiceKindStateful ...
	ServiceKindBasicServiceResourceUpdatePropertiesServiceKindStateful ServiceKindBasicServiceResourceUpdateProperties = "Stateful"
	// ServiceKindBasicServiceResourceUpdatePropertiesServiceKindStateless ...
	ServiceKindBasicServiceResourceUpdatePropertiesServiceKindStateless ServiceKindBasicServiceResourceUpdateProperties = "Stateless"
)

// PossibleServiceKindBasicServiceResourceUpdatePropertiesValues returns an array of possible values for the ServiceKindBasicServiceResourceUpdateProperties const type.
func PossibleServiceKindBasicServiceResourceUpdatePropertiesValues() []ServiceKindBasicServiceResourceUpdateProperties {
	return []ServiceKindBasicServiceResourceUpdateProperties{ServiceKindBasicServiceResourceUpdatePropertiesServiceKindServiceResourceUpdateProperties, ServiceKindBasicServiceResourceUpdatePropertiesServiceKindStateful, ServiceKindBasicServiceResourceUpdatePropertiesServiceKindStateless}
}

// ServiceLoadMetricWeight enumerates the values for service load metric weight.
type ServiceLoadMetricWeight string

const (
	// ServiceLoadMetricWeightHigh Specifies the metric weight of the service load as High. The value is 3.
	ServiceLoadMetricWeightHigh ServiceLoadMetricWeight = "High"
	// ServiceLoadMetricWeightLow Specifies the metric weight of the service load as Low. The value is 1.
	ServiceLoadMetricWeightLow ServiceLoadMetricWeight = "Low"
	// ServiceLoadMetricWeightMedium Specifies the metric weight of the service load as Medium. The value is 2.
	ServiceLoadMetricWeightMedium ServiceLoadMetricWeight = "Medium"
	// ServiceLoadMetricWeightZero Disables resource balancing for this metric. This value is zero.
	ServiceLoadMetricWeightZero ServiceLoadMetricWeight = "Zero"
)

// PossibleServiceLoadMetricWeightValues returns an array of possible values for the ServiceLoadMetricWeight const type.
func PossibleServiceLoadMetricWeightValues() []ServiceLoadMetricWeight {
	return []ServiceLoadMetricWeight{ServiceLoadMetricWeightHigh, ServiceLoadMetricWeightLow, ServiceLoadMetricWeightMedium, ServiceLoadMetricWeightZero}
}

// ServicePlacementPolicyType enumerates the values for service placement policy type.
type ServicePlacementPolicyType string

const (
	// ServicePlacementPolicyTypeInvalid Indicates the type of the placement policy is invalid. All Service
	// Fabric enumerations have the invalid type. The value is zero.
	ServicePlacementPolicyTypeInvalid ServicePlacementPolicyType = "Invalid"
	// ServicePlacementPolicyTypeInvalidDomain Indicates that the ServicePlacementPolicyDescription is of type
	// ServicePlacementInvalidDomainPolicyDescription, which indicates that a particular fault or upgrade
	// domain cannot be used for placement of this service. The value is 1.
	ServicePlacementPolicyTypeInvalidDomain ServicePlacementPolicyType = "InvalidDomain"
	// ServicePlacementPolicyTypeNonPartiallyPlaceService Indicates that the ServicePlacementPolicyDescription
	// is of type ServicePlacementNonPartiallyPlaceServicePolicyDescription, which indicates that if possible
	// all replicas of a particular partition of the service should be placed atomically. The value is 5.
	ServicePlacementPolicyTypeNonPartiallyPlaceService ServicePlacementPolicyType = "NonPartiallyPlaceService"
	// ServicePlacementPolicyTypePreferredPrimaryDomain Indicates that the ServicePlacementPolicyDescription is
	// of type ServicePlacementPreferPrimaryDomainPolicyDescription, which indicates that if possible the
	// Primary replica for the partitions of the service should be located in a particular domain as an
	// optimization. The value is 3.
	ServicePlacementPolicyTypePreferredPrimaryDomain ServicePlacementPolicyType = "PreferredPrimaryDomain"
	// ServicePlacementPolicyTypeRequiredDomain Indicates that the ServicePlacementPolicyDescription is of type
	// ServicePlacementRequireDomainDistributionPolicyDescription indicating that the replicas of the service
	// must be placed in a specific domain. The value is 2.
	ServicePlacementPolicyTypeRequiredDomain ServicePlacementPolicyType = "RequiredDomain"
	// ServicePlacementPolicyTypeRequiredDomainDistribution Indicates that the
	// ServicePlacementPolicyDescription is of type ServicePlacementRequireDomainDistributionPolicyDescription,
	// indicating that the system will disallow placement of any two replicas from the same partition in the
	// same domain at any time. The value is 4.
	ServicePlacementPolicyTypeRequiredDomainDistribution ServicePlacementPolicyType = "RequiredDomainDistribution"
)

// PossibleServicePlacementPolicyTypeValues returns an array of possible values for the ServicePlacementPolicyType const type.
func PossibleServicePlacementPolicyTypeValues() []ServicePlacementPolicyType {
	return []ServicePlacementPolicyType{ServicePlacementPolicyTypeInvalid, ServicePlacementPolicyTypeInvalidDomain, ServicePlacementPolicyTypeNonPartiallyPlaceService, ServicePlacementPolicyTypePreferredPrimaryDomain, ServicePlacementPolicyTypeRequiredDomain, ServicePlacementPolicyTypeRequiredDomainDistribution}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeServicePlacementPolicyDescription ...
	TypeServicePlacementPolicyDescription Type = "ServicePlacementPolicyDescription"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeServicePlacementPolicyDescription}
}

// UpgradeMode enumerates the values for upgrade mode.
type UpgradeMode string

const (
	// UpgradeModeAutomatic ...
	UpgradeModeAutomatic UpgradeMode = "Automatic"
	// UpgradeModeManual ...
	UpgradeModeManual UpgradeMode = "Manual"
)

// PossibleUpgradeModeValues returns an array of possible values for the UpgradeMode const type.
func PossibleUpgradeModeValues() []UpgradeMode {
	return []UpgradeMode{UpgradeModeAutomatic, UpgradeModeManual}
}

// UpgradeMode1 enumerates the values for upgrade mode 1.
type UpgradeMode1 string

const (
	// UpgradeMode1Automatic ...
	UpgradeMode1Automatic UpgradeMode1 = "Automatic"
	// UpgradeMode1Manual ...
	UpgradeMode1Manual UpgradeMode1 = "Manual"
)

// PossibleUpgradeMode1Values returns an array of possible values for the UpgradeMode1 const type.
func PossibleUpgradeMode1Values() []UpgradeMode1 {
	return []UpgradeMode1{UpgradeMode1Automatic, UpgradeMode1Manual}
}

// X509StoreName enumerates the values for x509 store name.
type X509StoreName string

const (
	// AddressBook ...
	AddressBook X509StoreName = "AddressBook"
	// AuthRoot ...
	AuthRoot X509StoreName = "AuthRoot"
	// CertificateAuthority ...
	CertificateAuthority X509StoreName = "CertificateAuthority"
	// Disallowed ...
	Disallowed X509StoreName = "Disallowed"
	// My ...
	My X509StoreName = "My"
	// Root ...
	Root X509StoreName = "Root"
	// TrustedPeople ...
	TrustedPeople X509StoreName = "TrustedPeople"
	// TrustedPublisher ...
	TrustedPublisher X509StoreName = "TrustedPublisher"
)

// PossibleX509StoreNameValues returns an array of possible values for the X509StoreName const type.
func PossibleX509StoreNameValues() []X509StoreName {
	return []X509StoreName{AddressBook, AuthRoot, CertificateAuthority, Disallowed, My, Root, TrustedPeople, TrustedPublisher}
}

// X509StoreName1 enumerates the values for x509 store name 1.
type X509StoreName1 string

const (
	// X509StoreName1AddressBook ...
	X509StoreName1AddressBook X509StoreName1 = "AddressBook"
	// X509StoreName1AuthRoot ...
	X509StoreName1AuthRoot X509StoreName1 = "AuthRoot"
	// X509StoreName1CertificateAuthority ...
	X509StoreName1CertificateAuthority X509StoreName1 = "CertificateAuthority"
	// X509StoreName1Disallowed ...
	X509StoreName1Disallowed X509StoreName1 = "Disallowed"
	// X509StoreName1My ...
	X509StoreName1My X509StoreName1 = "My"
	// X509StoreName1Root ...
	X509StoreName1Root X509StoreName1 = "Root"
	// X509StoreName1TrustedPeople ...
	X509StoreName1TrustedPeople X509StoreName1 = "TrustedPeople"
	// X509StoreName1TrustedPublisher ...
	X509StoreName1TrustedPublisher X509StoreName1 = "TrustedPublisher"
)

// PossibleX509StoreName1Values returns an array of possible values for the X509StoreName1 const type.
func PossibleX509StoreName1Values() []X509StoreName1 {
	return []X509StoreName1{X509StoreName1AddressBook, X509StoreName1AuthRoot, X509StoreName1CertificateAuthority, X509StoreName1Disallowed, X509StoreName1My, X509StoreName1Root, X509StoreName1TrustedPeople, X509StoreName1TrustedPublisher}
}
