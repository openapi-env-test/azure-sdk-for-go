package storagepool

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DiskPoolTier enumerates the values for disk pool tier.
type DiskPoolTier string

const (
	// Basic ...
	Basic DiskPoolTier = "Basic"
	// Premium ...
	Premium DiskPoolTier = "Premium"
	// Standard ...
	Standard DiskPoolTier = "Standard"
)

// PossibleDiskPoolTierValues returns an array of possible values for the DiskPoolTier const type.
func PossibleDiskPoolTierValues() []DiskPoolTier {
	return []DiskPoolTier{Basic, Premium, Standard}
}

// OperationalStatus enumerates the values for operational status.
type OperationalStatus string

const (
	// Healthy ...
	Healthy OperationalStatus = "Healthy"
	// Invalid ...
	Invalid OperationalStatus = "Invalid"
	// Running ...
	Running OperationalStatus = "Running"
	// Stopped ...
	Stopped OperationalStatus = "Stopped"
	// Stoppeddeallocated ...
	Stoppeddeallocated OperationalStatus = "Stopped (deallocated)"
	// Unhealthy ...
	Unhealthy OperationalStatus = "Unhealthy"
	// Unknown ...
	Unknown OperationalStatus = "Unknown"
	// Updating ...
	Updating OperationalStatus = "Updating"
)

// PossibleOperationalStatusValues returns an array of possible values for the OperationalStatus const type.
func PossibleOperationalStatusValues() []OperationalStatus {
	return []OperationalStatus{Healthy, Invalid, Running, Stopped, Stoppeddeallocated, Unhealthy, Unknown, Updating}
}

// ProvisioningStates enumerates the values for provisioning states.
type ProvisioningStates string

const (
	// ProvisioningStatesCanceled ...
	ProvisioningStatesCanceled ProvisioningStates = "Canceled"
	// ProvisioningStatesCreating ...
	ProvisioningStatesCreating ProvisioningStates = "Creating"
	// ProvisioningStatesDeleting ...
	ProvisioningStatesDeleting ProvisioningStates = "Deleting"
	// ProvisioningStatesFailed ...
	ProvisioningStatesFailed ProvisioningStates = "Failed"
	// ProvisioningStatesInvalid ...
	ProvisioningStatesInvalid ProvisioningStates = "Invalid"
	// ProvisioningStatesPending ...
	ProvisioningStatesPending ProvisioningStates = "Pending"
	// ProvisioningStatesSucceeded ...
	ProvisioningStatesSucceeded ProvisioningStates = "Succeeded"
	// ProvisioningStatesUpdating ...
	ProvisioningStatesUpdating ProvisioningStates = "Updating"
)

// PossibleProvisioningStatesValues returns an array of possible values for the ProvisioningStates const type.
func PossibleProvisioningStatesValues() []ProvisioningStates {
	return []ProvisioningStates{ProvisioningStatesCanceled, ProvisioningStatesCreating, ProvisioningStatesDeleting, ProvisioningStatesFailed, ProvisioningStatesInvalid, ProvisioningStatesPending, ProvisioningStatesSucceeded, ProvisioningStatesUpdating}
}
