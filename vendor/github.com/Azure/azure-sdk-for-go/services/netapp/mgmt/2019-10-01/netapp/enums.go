package netapp

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

// CheckNameResourceTypes enumerates the values for check name resource types.
type CheckNameResourceTypes string

const (
	// MicrosoftNetAppnetAppAccounts ...
	MicrosoftNetAppnetAppAccounts CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts"
	// MicrosoftNetAppnetAppAccountscapacityPools ...
	MicrosoftNetAppnetAppAccountscapacityPools CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools"
	// MicrosoftNetAppnetAppAccountscapacityPoolsvolumes ...
	MicrosoftNetAppnetAppAccountscapacityPoolsvolumes CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes"
	// MicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots ...
	MicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes/snapshots"
)

// PossibleCheckNameResourceTypesValues returns an array of possible values for the CheckNameResourceTypes const type.
func PossibleCheckNameResourceTypesValues() []CheckNameResourceTypes {
	return []CheckNameResourceTypes{MicrosoftNetAppnetAppAccounts, MicrosoftNetAppnetAppAccountscapacityPools, MicrosoftNetAppnetAppAccountscapacityPoolsvolumes, MicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots}
}

// EndpointType enumerates the values for endpoint type.
type EndpointType string

const (
	// Dst ...
	Dst EndpointType = "dst"
	// Src ...
	Src EndpointType = "src"
)

// PossibleEndpointTypeValues returns an array of possible values for the EndpointType const type.
func PossibleEndpointTypeValues() []EndpointType {
	return []EndpointType{Dst, Src}
}

// InAvailabilityReasonType enumerates the values for in availability reason type.
type InAvailabilityReasonType string

const (
	// AlreadyExists ...
	AlreadyExists InAvailabilityReasonType = "AlreadyExists"
	// Invalid ...
	Invalid InAvailabilityReasonType = "Invalid"
)

// PossibleInAvailabilityReasonTypeValues returns an array of possible values for the InAvailabilityReasonType const type.
func PossibleInAvailabilityReasonTypeValues() []InAvailabilityReasonType {
	return []InAvailabilityReasonType{AlreadyExists, Invalid}
}

// MirrorState enumerates the values for mirror state.
type MirrorState string

const (
	// Broken ...
	Broken MirrorState = "Broken"
	// Mirrored ...
	Mirrored MirrorState = "Mirrored"
	// Uninitialized ...
	Uninitialized MirrorState = "Uninitialized"
)

// PossibleMirrorStateValues returns an array of possible values for the MirrorState const type.
func PossibleMirrorStateValues() []MirrorState {
	return []MirrorState{Broken, Mirrored, Uninitialized}
}

// RelationshipStatus enumerates the values for relationship status.
type RelationshipStatus string

const (
	// Idle ...
	Idle RelationshipStatus = "Idle"
	// Transferring ...
	Transferring RelationshipStatus = "Transferring"
)

// PossibleRelationshipStatusValues returns an array of possible values for the RelationshipStatus const type.
func PossibleRelationshipStatusValues() []RelationshipStatus {
	return []RelationshipStatus{Idle, Transferring}
}

// ReplicationSchedule enumerates the values for replication schedule.
type ReplicationSchedule string

const (
	// OneZerominutely ...
	OneZerominutely ReplicationSchedule = "_10minutely"
	// Daily ...
	Daily ReplicationSchedule = "daily"
	// Hourly ...
	Hourly ReplicationSchedule = "hourly"
)

// PossibleReplicationScheduleValues returns an array of possible values for the ReplicationSchedule const type.
func PossibleReplicationScheduleValues() []ReplicationSchedule {
	return []ReplicationSchedule{OneZerominutely, Daily, Hourly}
}

// ServiceLevel enumerates the values for service level.
type ServiceLevel string

const (
	// Premium Premium service level
	Premium ServiceLevel = "Premium"
	// Standard Standard service level
	Standard ServiceLevel = "Standard"
	// Ultra Ultra service level
	Ultra ServiceLevel = "Ultra"
)

// PossibleServiceLevelValues returns an array of possible values for the ServiceLevel const type.
func PossibleServiceLevelValues() []ServiceLevel {
	return []ServiceLevel{Premium, Standard, Ultra}
}
