package policy

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

// EnforcementMode enumerates the values for enforcement mode.
type EnforcementMode string

const (
	// Default The policy effect is enforced during resource creation or update.
	Default EnforcementMode = "Default"
	// DoNotEnforce The policy effect is not enforced during resource creation or update.
	DoNotEnforce EnforcementMode = "DoNotEnforce"
)

// PossibleEnforcementModeValues returns an array of possible values for the EnforcementMode const type.
func PossibleEnforcementModeValues() []EnforcementMode {
	return []EnforcementMode{Default, DoNotEnforce}
}

// ParameterType enumerates the values for parameter type.
type ParameterType string

const (
	// Array ...
	Array ParameterType = "Array"
	// Boolean ...
	Boolean ParameterType = "Boolean"
	// DateTime ...
	DateTime ParameterType = "DateTime"
	// Float ...
	Float ParameterType = "Float"
	// Integer ...
	Integer ParameterType = "Integer"
	// Object ...
	Object ParameterType = "Object"
	// String ...
	String ParameterType = "String"
)

// PossibleParameterTypeValues returns an array of possible values for the ParameterType const type.
func PossibleParameterTypeValues() []ParameterType {
	return []ParameterType{Array, Boolean, DateTime, Float, Integer, Object, String}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// None Indicates that no identity is associated with the resource or that the existing identity should be
	// removed.
	None ResourceIdentityType = "None"
	// SystemAssigned Indicates that a system assigned identity is associated with the resource.
	SystemAssigned ResourceIdentityType = "SystemAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{None, SystemAssigned}
}

// Type enumerates the values for type.
type Type string

const (
	// BuiltIn ...
	BuiltIn Type = "BuiltIn"
	// Custom ...
	Custom Type = "Custom"
	// NotSpecified ...
	NotSpecified Type = "NotSpecified"
	// Static ...
	Static Type = "Static"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{BuiltIn, Custom, NotSpecified, Static}
}
