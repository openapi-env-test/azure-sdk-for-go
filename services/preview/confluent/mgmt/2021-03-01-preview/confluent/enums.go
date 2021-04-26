package confluent

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

// ProvisionState enumerates the values for provision state.
type ProvisionState string

const (
	// Accepted ...
	Accepted ProvisionState = "Accepted"
	// Canceled ...
	Canceled ProvisionState = "Canceled"
	// Creating ...
	Creating ProvisionState = "Creating"
	// Deleted ...
	Deleted ProvisionState = "Deleted"
	// Deleting ...
	Deleting ProvisionState = "Deleting"
	// Failed ...
	Failed ProvisionState = "Failed"
	// NotSpecified ...
	NotSpecified ProvisionState = "NotSpecified"
	// Succeeded ...
	Succeeded ProvisionState = "Succeeded"
	// Updating ...
	Updating ProvisionState = "Updating"
)

// PossibleProvisionStateValues returns an array of possible values for the ProvisionState const type.
func PossibleProvisionStateValues() []ProvisionState {
	return []ProvisionState{Accepted, Canceled, Creating, Deleted, Deleting, Failed, NotSpecified, Succeeded, Updating}
}

// SaaSOfferStatus enumerates the values for saa s offer status.
type SaaSOfferStatus string

const (
	// SaaSOfferStatusFailed ...
	SaaSOfferStatusFailed SaaSOfferStatus = "Failed"
	// SaaSOfferStatusInProgress ...
	SaaSOfferStatusInProgress SaaSOfferStatus = "InProgress"
	// SaaSOfferStatusPendingFulfillmentStart ...
	SaaSOfferStatusPendingFulfillmentStart SaaSOfferStatus = "PendingFulfillmentStart"
	// SaaSOfferStatusReinstated ...
	SaaSOfferStatusReinstated SaaSOfferStatus = "Reinstated"
	// SaaSOfferStatusStarted ...
	SaaSOfferStatusStarted SaaSOfferStatus = "Started"
	// SaaSOfferStatusSubscribed ...
	SaaSOfferStatusSubscribed SaaSOfferStatus = "Subscribed"
	// SaaSOfferStatusSucceeded ...
	SaaSOfferStatusSucceeded SaaSOfferStatus = "Succeeded"
	// SaaSOfferStatusSuspended ...
	SaaSOfferStatusSuspended SaaSOfferStatus = "Suspended"
	// SaaSOfferStatusUnsubscribed ...
	SaaSOfferStatusUnsubscribed SaaSOfferStatus = "Unsubscribed"
	// SaaSOfferStatusUpdating ...
	SaaSOfferStatusUpdating SaaSOfferStatus = "Updating"
)

// PossibleSaaSOfferStatusValues returns an array of possible values for the SaaSOfferStatus const type.
func PossibleSaaSOfferStatusValues() []SaaSOfferStatus {
	return []SaaSOfferStatus{SaaSOfferStatusFailed, SaaSOfferStatusInProgress, SaaSOfferStatusPendingFulfillmentStart, SaaSOfferStatusReinstated, SaaSOfferStatusStarted, SaaSOfferStatusSubscribed, SaaSOfferStatusSucceeded, SaaSOfferStatusSuspended, SaaSOfferStatusUnsubscribed, SaaSOfferStatusUpdating}
}
