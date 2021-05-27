package maps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// CreatedByTypeApplication ...
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey ...
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity ...
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser ...
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{CreatedByTypeApplication, CreatedByTypeKey, CreatedByTypeManagedIdentity, CreatedByTypeUser}
}

// KeyType enumerates the values for key type.
type KeyType string

const (
	// KeyTypePrimary ...
	KeyTypePrimary KeyType = "primary"
	// KeyTypeSecondary ...
	KeyTypeSecondary KeyType = "secondary"
)

// PossibleKeyTypeValues returns an array of possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{KeyTypePrimary, KeyTypeSecondary}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// KindGen1 ...
	KindGen1 Kind = "Gen1"
	// KindGen2 ...
	KindGen2 Kind = "Gen2"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{KindGen1, KindGen2}
}

// Name enumerates the values for name.
type Name string

const (
	// NameG2 ...
	NameG2 Name = "G2"
	// NameS0 ...
	NameS0 Name = "S0"
	// NameS1 ...
	NameS1 Name = "S1"
)

// PossibleNameValues returns an array of possible values for the Name const type.
func PossibleNameValues() []Name {
	return []Name{NameG2, NameS0, NameS1}
}
