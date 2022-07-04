//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armanalysisservices

const (
	moduleName    = "armanalysisservices"
	moduleVersion = "v0.2.1"
)

// ConnectionMode - How the read-write server's participation in the query pool is controlled.
// It can have the following values: * readOnly - indicates that the read-write server is intended not to participate in query
// operations
// * all - indicates that the read-write server can participate in query operations
// Specifying readOnly when capacity is 1 results in error.
type ConnectionMode string

const (
	ConnectionModeAll      ConnectionMode = "All"
	ConnectionModeReadOnly ConnectionMode = "ReadOnly"
)

// PossibleConnectionModeValues returns the possible values for the ConnectionMode const type.
func PossibleConnectionModeValues() []ConnectionMode {
	return []ConnectionMode{
		ConnectionModeAll,
		ConnectionModeReadOnly,
	}
}

// ToPtr returns a *ConnectionMode pointing to the current value.
func (c ConnectionMode) ToPtr() *ConnectionMode {
	return &c
}

// ManagedMode - The managed mode of the server (0 = not managed, 1 = managed).
type ManagedMode int32

const (
	ManagedModeZero ManagedMode = 0
	ManagedModeOne  ManagedMode = 1
)

// PossibleManagedModeValues returns the possible values for the ManagedMode const type.
func PossibleManagedModeValues() []ManagedMode {
	return []ManagedMode{
		ManagedModeZero,
		ManagedModeOne,
	}
}

// ToPtr returns a *ManagedMode pointing to the current value.
func (c ManagedMode) ToPtr() *ManagedMode {
	return &c
}

// ProvisioningState - The current deployment state of Analysis Services resource. The provisioningState is to indicate states
// for resource provisioning.
type ProvisioningState string

const (
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStatePaused       ProvisioningState = "Paused"
	ProvisioningStateSuspended    ProvisioningState = "Suspended"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
	ProvisioningStateSuspending   ProvisioningState = "Suspending"
	ProvisioningStatePausing      ProvisioningState = "Pausing"
	ProvisioningStateResuming     ProvisioningState = "Resuming"
	ProvisioningStatePreparing    ProvisioningState = "Preparing"
	ProvisioningStateScaling      ProvisioningState = "Scaling"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateDeleting,
		ProvisioningStateSucceeded,
		ProvisioningStateFailed,
		ProvisioningStatePaused,
		ProvisioningStateSuspended,
		ProvisioningStateProvisioning,
		ProvisioningStateUpdating,
		ProvisioningStateSuspending,
		ProvisioningStatePausing,
		ProvisioningStateResuming,
		ProvisioningStatePreparing,
		ProvisioningStateScaling,
	}
}

// ToPtr returns a *ProvisioningState pointing to the current value.
func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}

// SKUTier - The name of the Azure pricing tier to which the SKU applies.
type SKUTier string

const (
	SKUTierBasic       SKUTier = "Basic"
	SKUTierDevelopment SKUTier = "Development"
	SKUTierStandard    SKUTier = "Standard"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierBasic,
		SKUTierDevelopment,
		SKUTierStandard,
	}
}

// ToPtr returns a *SKUTier pointing to the current value.
func (c SKUTier) ToPtr() *SKUTier {
	return &c
}

// ServerMonitorMode - The server monitor mode for AS server
type ServerMonitorMode int32

const (
	ServerMonitorModeZero ServerMonitorMode = 0
	ServerMonitorModeOne  ServerMonitorMode = 1
)

// PossibleServerMonitorModeValues returns the possible values for the ServerMonitorMode const type.
func PossibleServerMonitorModeValues() []ServerMonitorMode {
	return []ServerMonitorMode{
		ServerMonitorModeZero,
		ServerMonitorModeOne,
	}
}

// ToPtr returns a *ServerMonitorMode pointing to the current value.
func (c ServerMonitorMode) ToPtr() *ServerMonitorMode {
	return &c
}

// State - The current state of Analysis Services resource. The state is to indicate more states outside of resource provisioning.
type State string

const (
	StateDeleting     State = "Deleting"
	StateFailed       State = "Failed"
	StatePaused       State = "Paused"
	StatePausing      State = "Pausing"
	StatePreparing    State = "Preparing"
	StateProvisioning State = "Provisioning"
	StateResuming     State = "Resuming"
	StateScaling      State = "Scaling"
	StateSucceeded    State = "Succeeded"
	StateSuspended    State = "Suspended"
	StateSuspending   State = "Suspending"
	StateUpdating     State = "Updating"
)

// PossibleStateValues returns the possible values for the State const type.
func PossibleStateValues() []State {
	return []State{
		StateDeleting,
		StateFailed,
		StatePaused,
		StatePausing,
		StatePreparing,
		StateProvisioning,
		StateResuming,
		StateScaling,
		StateSucceeded,
		StateSuspended,
		StateSuspending,
		StateUpdating,
	}
}

// ToPtr returns a *State pointing to the current value.
func (c State) ToPtr() *State {
	return &c
}
