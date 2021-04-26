package elastic

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// Application ...
	Application CreatedByType = "Application"
	// Key ...
	Key CreatedByType = "Key"
	// ManagedIdentity ...
	ManagedIdentity CreatedByType = "ManagedIdentity"
	// User ...
	User CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{Application, Key, ManagedIdentity, User}
}

// DeploymentStatus enumerates the values for deployment status.
type DeploymentStatus string

const (
	// Healthy ...
	Healthy DeploymentStatus = "Healthy"
	// Unhealthy ...
	Unhealthy DeploymentStatus = "Unhealthy"
)

// PossibleDeploymentStatusValues returns an array of possible values for the DeploymentStatus const type.
func PossibleDeploymentStatusValues() []DeploymentStatus {
	return []DeploymentStatus{Healthy, Unhealthy}
}

// LiftrResourceCategories enumerates the values for liftr resource categories.
type LiftrResourceCategories string

const (
	// MonitorLogs ...
	MonitorLogs LiftrResourceCategories = "MonitorLogs"
	// Unknown ...
	Unknown LiftrResourceCategories = "Unknown"
)

// PossibleLiftrResourceCategoriesValues returns an array of possible values for the LiftrResourceCategories const type.
func PossibleLiftrResourceCategoriesValues() []LiftrResourceCategories {
	return []LiftrResourceCategories{MonitorLogs, Unknown}
}

// ManagedIdentityTypes enumerates the values for managed identity types.
type ManagedIdentityTypes string

const (
	// SystemAssigned ...
	SystemAssigned ManagedIdentityTypes = "SystemAssigned"
)

// PossibleManagedIdentityTypesValues returns an array of possible values for the ManagedIdentityTypes const type.
func PossibleManagedIdentityTypesValues() []ManagedIdentityTypes {
	return []ManagedIdentityTypes{SystemAssigned}
}

// MonitoringStatus enumerates the values for monitoring status.
type MonitoringStatus string

const (
	// Disabled ...
	Disabled MonitoringStatus = "Disabled"
	// Enabled ...
	Enabled MonitoringStatus = "Enabled"
)

// PossibleMonitoringStatusValues returns an array of possible values for the MonitoringStatus const type.
func PossibleMonitoringStatusValues() []MonitoringStatus {
	return []MonitoringStatus{Disabled, Enabled}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Accepted ...
	Accepted ProvisioningState = "Accepted"
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleted ...
	Deleted ProvisioningState = "Deleted"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// NotSpecified ...
	NotSpecified ProvisioningState = "NotSpecified"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Accepted, Canceled, Creating, Deleted, Deleting, Failed, NotSpecified, Succeeded, Updating}
}

// SendingLogs enumerates the values for sending logs.
type SendingLogs string

const (
	// False ...
	False SendingLogs = "False"
	// True ...
	True SendingLogs = "True"
)

// PossibleSendingLogsValues returns an array of possible values for the SendingLogs const type.
func PossibleSendingLogsValues() []SendingLogs {
	return []SendingLogs{False, True}
}

// TagAction enumerates the values for tag action.
type TagAction string

const (
	// Exclude ...
	Exclude TagAction = "Exclude"
	// Include ...
	Include TagAction = "Include"
)

// PossibleTagActionValues returns an array of possible values for the TagAction const type.
func PossibleTagActionValues() []TagAction {
	return []TagAction{Exclude, Include}
}
