//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization

const (
	moduleName    = "armauthorization"
	moduleVersion = "v3.0.0-beta.2"
)

// PrincipalType - The principal type of the assigned principal ID.
type PrincipalType string

const (
	PrincipalTypeDevice           PrincipalType = "Device"
	PrincipalTypeForeignGroup     PrincipalType = "ForeignGroup"
	PrincipalTypeGroup            PrincipalType = "Group"
	PrincipalTypeServicePrincipal PrincipalType = "ServicePrincipal"
	PrincipalTypeUser             PrincipalType = "User"
)

// PossiblePrincipalTypeValues returns the possible values for the PrincipalType const type.
func PossiblePrincipalTypeValues() []PrincipalType {
	return []PrincipalType{
		PrincipalTypeDevice,
		PrincipalTypeForeignGroup,
		PrincipalTypeGroup,
		PrincipalTypeServicePrincipal,
		PrincipalTypeUser,
	}
}

// RequestType - The type of the role assignment schedule request. Eg: SelfActivate, AdminAssign etc
type RequestType string

const (
	RequestTypeAdminAssign    RequestType = "AdminAssign"
	RequestTypeAdminExtend    RequestType = "AdminExtend"
	RequestTypeAdminRemove    RequestType = "AdminRemove"
	RequestTypeAdminRenew     RequestType = "AdminRenew"
	RequestTypeAdminUpdate    RequestType = "AdminUpdate"
	RequestTypeSelfActivate   RequestType = "SelfActivate"
	RequestTypeSelfDeactivate RequestType = "SelfDeactivate"
	RequestTypeSelfExtend     RequestType = "SelfExtend"
	RequestTypeSelfRenew      RequestType = "SelfRenew"
)

// PossibleRequestTypeValues returns the possible values for the RequestType const type.
func PossibleRequestTypeValues() []RequestType {
	return []RequestType{
		RequestTypeAdminAssign,
		RequestTypeAdminExtend,
		RequestTypeAdminRemove,
		RequestTypeAdminRenew,
		RequestTypeAdminUpdate,
		RequestTypeSelfActivate,
		RequestTypeSelfDeactivate,
		RequestTypeSelfExtend,
		RequestTypeSelfRenew,
	}
}

// Status - The status of the role assignment schedule request.
type Status string

const (
	StatusAccepted                    Status = "Accepted"
	StatusAdminApproved               Status = "AdminApproved"
	StatusAdminDenied                 Status = "AdminDenied"
	StatusCanceled                    Status = "Canceled"
	StatusDenied                      Status = "Denied"
	StatusFailed                      Status = "Failed"
	StatusFailedAsResourceIsLocked    Status = "FailedAsResourceIsLocked"
	StatusGranted                     Status = "Granted"
	StatusInvalid                     Status = "Invalid"
	StatusPendingAdminDecision        Status = "PendingAdminDecision"
	StatusPendingApproval             Status = "PendingApproval"
	StatusPendingApprovalProvisioning Status = "PendingApprovalProvisioning"
	StatusPendingEvaluation           Status = "PendingEvaluation"
	StatusPendingExternalProvisioning Status = "PendingExternalProvisioning"
	StatusPendingProvisioning         Status = "PendingProvisioning"
	StatusPendingRevocation           Status = "PendingRevocation"
	StatusPendingScheduleCreation     Status = "PendingScheduleCreation"
	StatusProvisioned                 Status = "Provisioned"
	StatusProvisioningStarted         Status = "ProvisioningStarted"
	StatusRevoked                     Status = "Revoked"
	StatusScheduleCreated             Status = "ScheduleCreated"
	StatusTimedOut                    Status = "TimedOut"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusAccepted,
		StatusAdminApproved,
		StatusAdminDenied,
		StatusCanceled,
		StatusDenied,
		StatusFailed,
		StatusFailedAsResourceIsLocked,
		StatusGranted,
		StatusInvalid,
		StatusPendingAdminDecision,
		StatusPendingApproval,
		StatusPendingApprovalProvisioning,
		StatusPendingEvaluation,
		StatusPendingExternalProvisioning,
		StatusPendingProvisioning,
		StatusPendingRevocation,
		StatusPendingScheduleCreation,
		StatusProvisioned,
		StatusProvisioningStarted,
		StatusRevoked,
		StatusScheduleCreated,
		StatusTimedOut,
	}
}

// Type - Type of the role assignment schedule expiration
type Type string

const (
	TypeAfterDateTime Type = "AfterDateTime"
	TypeAfterDuration Type = "AfterDuration"
	TypeNoExpiration  Type = "NoExpiration"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeAfterDateTime,
		TypeAfterDuration,
		TypeNoExpiration,
	}
}
