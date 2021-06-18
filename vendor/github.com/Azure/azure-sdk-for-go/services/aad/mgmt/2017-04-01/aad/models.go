package aad

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/aad/mgmt/2017-04-01/aad"

// DiagnosticSettings the diagnostic settings.
type DiagnosticSettings struct {
	// StorageAccountID - The resource ID of the storage account to which you would like to send Diagnostic Logs.
	StorageAccountID *string `json:"storageAccountId,omitempty"`
	// ServiceBusRuleID - The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
	ServiceBusRuleID *string `json:"serviceBusRuleId,omitempty"`
	// WorkspaceID - The workspace ID (resource ID of a Log Analytics workspace) for a Log Analytics workspace to which you would like to send Diagnostic Logs. Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
	WorkspaceID *string `json:"workspaceId,omitempty"`
	// EventHubAuthorizationRuleID - The resource Id for the event hub authorization rule.
	EventHubAuthorizationRuleID *string `json:"eventHubAuthorizationRuleId,omitempty"`
	// EventHubName - The name of the event hub. If none is specified, the default event hub will be selected.
	EventHubName *string `json:"eventHubName,omitempty"`
	// Logs - The list of logs settings.
	Logs *[]LogSettings `json:"logs,omitempty"`
}

// DiagnosticSettingsCategory the diagnostic settings Category.
type DiagnosticSettingsCategory struct {
	// CategoryType - The type of the diagnostic settings category. Possible values include: 'Logs'
	CategoryType CategoryType `json:"categoryType,omitempty"`
}

// DiagnosticSettingsCategoryResource the diagnostic settings category resource.
type DiagnosticSettingsCategoryResource struct {
	// DiagnosticSettingsCategory - The properties of a Diagnostic Settings Category.
	*DiagnosticSettingsCategory `json:"properties,omitempty"`
	// ID - READ-ONLY; Azure resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Azure resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Azure resource type
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for DiagnosticSettingsCategoryResource.
func (dscr DiagnosticSettingsCategoryResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dscr.DiagnosticSettingsCategory != nil {
		objectMap["properties"] = dscr.DiagnosticSettingsCategory
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for DiagnosticSettingsCategoryResource struct.
func (dscr *DiagnosticSettingsCategoryResource) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var diagnosticSettingsCategory DiagnosticSettingsCategory
				err = json.Unmarshal(*v, &diagnosticSettingsCategory)
				if err != nil {
					return err
				}
				dscr.DiagnosticSettingsCategory = &diagnosticSettingsCategory
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				dscr.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				dscr.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				dscr.Type = &typeVar
			}
		}
	}

	return nil
}

// DiagnosticSettingsCategoryResourceCollection represents a collection of diagnostic setting category
// resources.
type DiagnosticSettingsCategoryResourceCollection struct {
	autorest.Response `json:"-"`
	// Value - The collection of diagnostic settings category resources.
	Value *[]DiagnosticSettingsCategoryResource `json:"value,omitempty"`
}

// DiagnosticSettingsResource the diagnostic setting resource.
type DiagnosticSettingsResource struct {
	autorest.Response `json:"-"`
	// DiagnosticSettings - Properties of a Diagnostic Settings Resource.
	*DiagnosticSettings `json:"properties,omitempty"`
	// ID - READ-ONLY; Azure resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Azure resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Azure resource type
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for DiagnosticSettingsResource.
func (dsr DiagnosticSettingsResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dsr.DiagnosticSettings != nil {
		objectMap["properties"] = dsr.DiagnosticSettings
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for DiagnosticSettingsResource struct.
func (dsr *DiagnosticSettingsResource) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var diagnosticSettings DiagnosticSettings
				err = json.Unmarshal(*v, &diagnosticSettings)
				if err != nil {
					return err
				}
				dsr.DiagnosticSettings = &diagnosticSettings
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				dsr.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				dsr.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				dsr.Type = &typeVar
			}
		}
	}

	return nil
}

// DiagnosticSettingsResourceCollection represents a collection of alert rule resources.
type DiagnosticSettingsResourceCollection struct {
	autorest.Response `json:"-"`
	// Value - The collection of diagnostic settings resources.
	Value *[]DiagnosticSettingsResource `json:"value,omitempty"`
}

// Display contains the localized display information for this particular operation / action. These value
// will be used by several clients for (1) custom role definitions for RBAC; (2) complex query filters for
// the event service; and (3) audit history / records for management operations.
type Display struct {
	// Publisher - The publisher. The localized friendly form of the resource publisher name.
	Publisher *string `json:"publisher,omitempty"`
	// Provider - The provider. The localized friendly form of the resource provider name – it is expected to also include the publisher/company responsible. It should use Title Casing and begin with "Microsoft" for 1st party services. e.g. "Microsoft Monitoring Insights" or "Microsoft Compute."
	Provider *string `json:"provider,omitempty"`
	// Resource - The resource. The localized friendly form of the resource related to this action/operation – it should match the public documentation for the resource provider. It should use Title Casing. This value should be unique for a particular URL type (e.g. nested types should *not* reuse their parent’s display.resource field). e.g. "Virtual Machines" or "Scheduler Job Collections", or "Virtual Machine VM Sizes" or "Scheduler Jobs"
	Resource *string `json:"resource,omitempty"`
	// Operation - The operation. The localized friendly name for the operation, as it should be shown to the user. It should be concise (to fit in drop downs) but clear (i.e. self-documenting). It should use Title Casing. Prescriptive guidance: Read Create or Update Delete 'ActionName'
	Operation *string `json:"operation,omitempty"`
	// Description - The description. The localized friendly description for the operation, as it should be shown to the user. It should be thorough, yet concise – it will be used in tool tips and detailed views. Prescriptive guidance for namespaces: Read any 'display.provider' resource Create or Update any 'display.provider' resource Delete any 'display.provider' resource Perform any other action on any 'display.provider' resource Prescriptive guidance for namespaces: Read any 'display.resource' Create or Update any 'display.resource' Delete any 'display.resource' 'ActionName' any 'display.resources'
	Description *string `json:"description,omitempty"`
}

// ErrorDefinition error definition.
type ErrorDefinition struct {
	// Code - READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; Description of the error.
	Message *string `json:"message,omitempty"`
	// Details - READ-ONLY; Internal error details.
	Details *[]ErrorDefinition `json:"details,omitempty"`
}

// MarshalJSON is the custom marshaler for ErrorDefinition.
func (ed ErrorDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ErrorResponse error response.
type ErrorResponse struct {
	// Error - The error details.
	Error *ErrorDefinition `json:"error,omitempty"`
}

// LogSettings part of MultiTenantDiagnosticSettings. Specifies the settings for a particular log.
type LogSettings struct {
	// Category - Name of a Diagnostic Log category for a resource type this setting is applied to. To obtain the list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation. Possible values include: 'AuditLogs', 'SignInLogs'
	Category Category `json:"category,omitempty"`
	// Enabled - A value indicating whether this log is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// RetentionPolicy - The retention policy for this log.
	RetentionPolicy *RetentionPolicy `json:"retentionPolicy,omitempty"`
}

// OperationsDiscovery operations discovery class.
type OperationsDiscovery struct {
	// Name - Name of the API. The name of the operation being performed on this particular object. It should match the action name that appears in RBAC / the event service. Examples of operations include: * Microsoft.Compute/virtualMachine/capture/action * Microsoft.Compute/virtualMachine/restart/action * Microsoft.Compute/virtualMachine/write * Microsoft.Compute/virtualMachine/read * Microsoft.Compute/virtualMachine/delete Each action should include, in order: (1) Resource Provider Namespace (2) Type hierarchy for which the action applies (e.g. server/databases for a SQL Azure database) (3) Read, Write, Action or Delete indicating which type applies. If it is a PUT/PATCH on a collection or named value, Write should be used. If it is a GET, Read should be used. If it is a DELETE, Delete should be used. If it is a POST, Action should be used.
	Name *string `json:"name,omitempty"`
	// Display - Object type
	Display *Display `json:"display,omitempty"`
	// Origin - Origin. The intended executor of the operation; governs the display of the operation in the RBAC UX and the audit logs UX. Default value is "user,system"
	Origin *string `json:"origin,omitempty"`
	// Properties - Properties. Reserved for future use.
	Properties interface{} `json:"properties,omitempty"`
}

// OperationsDiscoveryCollection collection of ClientDiscovery details.
type OperationsDiscoveryCollection struct {
	autorest.Response `json:"-"`
	// Value - The ClientDiscovery details.
	Value *[]OperationsDiscovery `json:"value,omitempty"`
}

// ProxyOnlyResource a proxy only azure resource object.
type ProxyOnlyResource struct {
	// ID - READ-ONLY; Azure resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Azure resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Azure resource type
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for ProxyOnlyResource.
func (por ProxyOnlyResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// RetentionPolicy specifies the retention policy for the log.
type RetentionPolicy struct {
	// Enabled - A value indicating whether the retention policy is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Days - The number of days for the retention in days. A value of 0 will retain the events indefinitely.
	Days *int32 `json:"days,omitempty"`
}
