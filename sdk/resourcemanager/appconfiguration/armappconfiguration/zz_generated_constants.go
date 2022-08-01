//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappconfiguration

const (
	moduleName    = "armappconfiguration"
	moduleVersion = "v0.3.0"
)

// ConfigurationResourceType - The resource type to check for name availability.
type ConfigurationResourceType string

const (
	ConfigurationResourceTypeMicrosoftAppConfigurationConfigurationStores ConfigurationResourceType = "Microsoft.AppConfiguration/configurationStores"
)

// PossibleConfigurationResourceTypeValues returns the possible values for the ConfigurationResourceType const type.
func PossibleConfigurationResourceTypeValues() []ConfigurationResourceType {
	return []ConfigurationResourceType{
		ConfigurationResourceTypeMicrosoftAppConfigurationConfigurationStores,
	}
}

// ToPtr returns a *ConfigurationResourceType pointing to the current value.
func (c ConfigurationResourceType) ToPtr() *ConfigurationResourceType {
	return &c
}

// ProvisioningState - The provisioning state of the configuration store.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ToPtr returns a *ProvisioningState pointing to the current value.
func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}
