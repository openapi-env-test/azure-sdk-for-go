package resourcemover

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DependencyLevel enumerates the values for dependency level.
type DependencyLevel string

const (
	// Descendant ...
	Descendant DependencyLevel = "Descendant"
	// Direct ...
	Direct DependencyLevel = "Direct"
)

// PossibleDependencyLevelValues returns an array of possible values for the DependencyLevel const type.
func PossibleDependencyLevelValues() []DependencyLevel {
	return []DependencyLevel{Descendant, Direct}
}

// DependencyType enumerates the values for dependency type.
type DependencyType string

const (
	// RequiredForMove ...
	RequiredForMove DependencyType = "RequiredForMove"
	// RequiredForPrepare ...
	RequiredForPrepare DependencyType = "RequiredForPrepare"
)

// PossibleDependencyTypeValues returns an array of possible values for the DependencyType const type.
func PossibleDependencyTypeValues() []DependencyType {
	return []DependencyType{RequiredForMove, RequiredForPrepare}
}

// JobName enumerates the values for job name.
type JobName string

const (
	// InitialSync ...
	InitialSync JobName = "InitialSync"
)

// PossibleJobNameValues returns an array of possible values for the JobName const type.
func PossibleJobNameValues() []JobName {
	return []JobName{InitialSync}
}

// MoveParadigm enumerates the values for move paradigm.
type MoveParadigm string

const (
	// CopyType ...
	CopyType MoveParadigm = "CopyType"
	// MigrateType ...
	MigrateType MoveParadigm = "MigrateType"
)

// PossibleMoveParadigmValues returns an array of possible values for the MoveParadigm const type.
func PossibleMoveParadigmValues() []MoveParadigm {
	return []MoveParadigm{CopyType, MigrateType}
}

// MoveResourceInputType enumerates the values for move resource input type.
type MoveResourceInputType string

const (
	// MoveResourceID ...
	MoveResourceID MoveResourceInputType = "MoveResourceId"
	// MoveResourceSourceID ...
	MoveResourceSourceID MoveResourceInputType = "MoveResourceSourceId"
)

// PossibleMoveResourceInputTypeValues returns an array of possible values for the MoveResourceInputType const type.
func PossibleMoveResourceInputTypeValues() []MoveResourceInputType {
	return []MoveResourceInputType{MoveResourceID, MoveResourceSourceID}
}

// MoveState enumerates the values for move state.
type MoveState string

const (
	// AssignmentPending ...
	AssignmentPending MoveState = "AssignmentPending"
	// CommitFailed ...
	CommitFailed MoveState = "CommitFailed"
	// CommitInProgress ...
	CommitInProgress MoveState = "CommitInProgress"
	// CommitPending ...
	CommitPending MoveState = "CommitPending"
	// Committed ...
	Committed MoveState = "Committed"
	// DeleteSourcePending ...
	DeleteSourcePending MoveState = "DeleteSourcePending"
	// DiscardFailed ...
	DiscardFailed MoveState = "DiscardFailed"
	// DiscardInProgress ...
	DiscardInProgress MoveState = "DiscardInProgress"
	// MoveFailed ...
	MoveFailed MoveState = "MoveFailed"
	// MoveInProgress ...
	MoveInProgress MoveState = "MoveInProgress"
	// MovePending ...
	MovePending MoveState = "MovePending"
	// PrepareFailed ...
	PrepareFailed MoveState = "PrepareFailed"
	// PrepareInProgress ...
	PrepareInProgress MoveState = "PrepareInProgress"
	// PreparePending ...
	PreparePending MoveState = "PreparePending"
	// ResourceMoveCompleted ...
	ResourceMoveCompleted MoveState = "ResourceMoveCompleted"
)

// PossibleMoveStateValues returns an array of possible values for the MoveState const type.
func PossibleMoveStateValues() []MoveState {
	return []MoveState{AssignmentPending, CommitFailed, CommitInProgress, CommitPending, Committed, DeleteSourcePending, DiscardFailed, DiscardInProgress, MoveFailed, MoveInProgress, MovePending, PrepareFailed, PrepareInProgress, PreparePending, ResourceMoveCompleted}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Creating, Failed, Succeeded, Updating}
}

// ResolutionType enumerates the values for resolution type.
type ResolutionType string

const (
	// Automatic ...
	Automatic ResolutionType = "Automatic"
	// Manual ...
	Manual ResolutionType = "Manual"
)

// PossibleResolutionTypeValues returns an array of possible values for the ResolutionType const type.
func PossibleResolutionTypeValues() []ResolutionType {
	return []ResolutionType{Automatic, Manual}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// None ...
	None ResourceIdentityType = "None"
	// SystemAssigned ...
	SystemAssigned ResourceIdentityType = "SystemAssigned"
	// UserAssigned ...
	UserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{None, SystemAssigned, UserAssigned}
}

// ResourceType enumerates the values for resource type.
type ResourceType string

const (
	// ResourceTypeResourceSettings ...
	ResourceTypeResourceSettings ResourceType = "ResourceSettings"
)

// PossibleResourceTypeValues returns an array of possible values for the ResourceType const type.
func PossibleResourceTypeValues() []ResourceType {
	return []ResourceType{ResourceTypeResourceSettings}
}
