// Package graphrbac implements the Azure ARM Graphrbac service API version 1.6.
//
// The Graph RBAC Management Client
package graphrbac

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseURI is the default URI used for the service Graphrbac
	DefaultBaseURI = "https://graph.windows.net"
)

// BaseClient is the base client for Graphrbac.
type BaseClient struct {
	autorest.Client
	BaseURI  string
	TenantID string
}

// New creates an instance of the BaseClient client.
func New(tenantID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, tenantID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, tenantID string) BaseClient {
	return BaseClient{
		Client:   autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:  baseURI,
		TenantID: tenantID,
	}
}
