package enterpriseknowledgegraphservice

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Creating, Deleting, Failed, Succeeded}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// F0 ...
	F0 SkuName = "F0"
	// S1 ...
	S1 SkuName = "S1"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{F0, S1}
}
