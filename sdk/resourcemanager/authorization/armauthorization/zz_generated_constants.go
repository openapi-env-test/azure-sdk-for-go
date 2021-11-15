//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

const (
	module  = "armauthorization"
	version = "v0.1.1"
)

// AccessRecommendationType - The feature- generated recommendation shown to the reviewer.
type AccessRecommendationType string

const (
	AccessRecommendationTypeApprove         AccessRecommendationType = "Approve"
	AccessRecommendationTypeDeny            AccessRecommendationType = "Deny"
	AccessRecommendationTypeNoInfoAvailable AccessRecommendationType = "NoInfoAvailable"
)

// PossibleAccessRecommendationTypeValues returns the possible values for the AccessRecommendationType const type.
func PossibleAccessRecommendationTypeValues() []AccessRecommendationType {
	return []AccessRecommendationType{
		AccessRecommendationTypeApprove,
		AccessRecommendationTypeDeny,
		AccessRecommendationTypeNoInfoAvailable,
	}
}

// ToPtr returns a *AccessRecommendationType pointing to the current value.
func (c AccessRecommendationType) ToPtr() *AccessRecommendationType {
	return &c
}

// AccessReviewActorIdentityType - The identity type : user/servicePrincipal
type AccessReviewActorIdentityType string

const (
	AccessReviewActorIdentityTypeServicePrincipal AccessReviewActorIdentityType = "servicePrincipal"
	AccessReviewActorIdentityTypeUser             AccessReviewActorIdentityType = "user"
)

// PossibleAccessReviewActorIdentityTypeValues returns the possible values for the AccessReviewActorIdentityType const type.
func PossibleAccessReviewActorIdentityTypeValues() []AccessReviewActorIdentityType {
	return []AccessReviewActorIdentityType{
		AccessReviewActorIdentityTypeServicePrincipal,
		AccessReviewActorIdentityTypeUser,
	}
}

// ToPtr returns a *AccessReviewActorIdentityType pointing to the current value.
func (c AccessReviewActorIdentityType) ToPtr() *AccessReviewActorIdentityType {
	return &c
}

// AccessReviewApplyResult - The outcome of applying the decision.
type AccessReviewApplyResult string

const (
	AccessReviewApplyResultAppliedSuccessfully                  AccessReviewApplyResult = "AppliedSuccessfully"
	AccessReviewApplyResultAppliedSuccessfullyButObjectNotFound AccessReviewApplyResult = "AppliedSuccessfullyButObjectNotFound"
	AccessReviewApplyResultAppliedWithUnknownFailure            AccessReviewApplyResult = "AppliedWithUnknownFailure"
	AccessReviewApplyResultApplyNotSupported                    AccessReviewApplyResult = "ApplyNotSupported"
	AccessReviewApplyResultApplying                             AccessReviewApplyResult = "Applying"
	AccessReviewApplyResultNew                                  AccessReviewApplyResult = "New"
)

// PossibleAccessReviewApplyResultValues returns the possible values for the AccessReviewApplyResult const type.
func PossibleAccessReviewApplyResultValues() []AccessReviewApplyResult {
	return []AccessReviewApplyResult{
		AccessReviewApplyResultAppliedSuccessfully,
		AccessReviewApplyResultAppliedSuccessfullyButObjectNotFound,
		AccessReviewApplyResultAppliedWithUnknownFailure,
		AccessReviewApplyResultApplyNotSupported,
		AccessReviewApplyResultApplying,
		AccessReviewApplyResultNew,
	}
}

// ToPtr returns a *AccessReviewApplyResult pointing to the current value.
func (c AccessReviewApplyResult) ToPtr() *AccessReviewApplyResult {
	return &c
}

// AccessReviewInstanceStatus - This read-only field specifies the status of an access review instance.
type AccessReviewInstanceStatus string

const (
	AccessReviewInstanceStatusApplied       AccessReviewInstanceStatus = "Applied"
	AccessReviewInstanceStatusApplying      AccessReviewInstanceStatus = "Applying"
	AccessReviewInstanceStatusAutoReviewed  AccessReviewInstanceStatus = "AutoReviewed"
	AccessReviewInstanceStatusAutoReviewing AccessReviewInstanceStatus = "AutoReviewing"
	AccessReviewInstanceStatusCompleted     AccessReviewInstanceStatus = "Completed"
	AccessReviewInstanceStatusCompleting    AccessReviewInstanceStatus = "Completing"
	AccessReviewInstanceStatusInProgress    AccessReviewInstanceStatus = "InProgress"
	AccessReviewInstanceStatusInitializing  AccessReviewInstanceStatus = "Initializing"
	AccessReviewInstanceStatusNotStarted    AccessReviewInstanceStatus = "NotStarted"
	AccessReviewInstanceStatusScheduled     AccessReviewInstanceStatus = "Scheduled"
	AccessReviewInstanceStatusStarting      AccessReviewInstanceStatus = "Starting"
)

// PossibleAccessReviewInstanceStatusValues returns the possible values for the AccessReviewInstanceStatus const type.
func PossibleAccessReviewInstanceStatusValues() []AccessReviewInstanceStatus {
	return []AccessReviewInstanceStatus{
		AccessReviewInstanceStatusApplied,
		AccessReviewInstanceStatusApplying,
		AccessReviewInstanceStatusAutoReviewed,
		AccessReviewInstanceStatusAutoReviewing,
		AccessReviewInstanceStatusCompleted,
		AccessReviewInstanceStatusCompleting,
		AccessReviewInstanceStatusInProgress,
		AccessReviewInstanceStatusInitializing,
		AccessReviewInstanceStatusNotStarted,
		AccessReviewInstanceStatusScheduled,
		AccessReviewInstanceStatusStarting,
	}
}

// ToPtr returns a *AccessReviewInstanceStatus pointing to the current value.
func (c AccessReviewInstanceStatus) ToPtr() *AccessReviewInstanceStatus {
	return &c
}

// AccessReviewRecurrencePatternType - The recurrence type : weekly, monthly, etc.
type AccessReviewRecurrencePatternType string

const (
	AccessReviewRecurrencePatternTypeAbsoluteMonthly AccessReviewRecurrencePatternType = "absoluteMonthly"
	AccessReviewRecurrencePatternTypeWeekly          AccessReviewRecurrencePatternType = "weekly"
)

// PossibleAccessReviewRecurrencePatternTypeValues returns the possible values for the AccessReviewRecurrencePatternType const type.
func PossibleAccessReviewRecurrencePatternTypeValues() []AccessReviewRecurrencePatternType {
	return []AccessReviewRecurrencePatternType{
		AccessReviewRecurrencePatternTypeAbsoluteMonthly,
		AccessReviewRecurrencePatternTypeWeekly,
	}
}

// ToPtr returns a *AccessReviewRecurrencePatternType pointing to the current value.
func (c AccessReviewRecurrencePatternType) ToPtr() *AccessReviewRecurrencePatternType {
	return &c
}

// AccessReviewRecurrenceRangeType - The recurrence range type. The possible values are: endDate, noEnd, numbered.
type AccessReviewRecurrenceRangeType string

const (
	AccessReviewRecurrenceRangeTypeEndDate  AccessReviewRecurrenceRangeType = "endDate"
	AccessReviewRecurrenceRangeTypeNoEnd    AccessReviewRecurrenceRangeType = "noEnd"
	AccessReviewRecurrenceRangeTypeNumbered AccessReviewRecurrenceRangeType = "numbered"
)

// PossibleAccessReviewRecurrenceRangeTypeValues returns the possible values for the AccessReviewRecurrenceRangeType const type.
func PossibleAccessReviewRecurrenceRangeTypeValues() []AccessReviewRecurrenceRangeType {
	return []AccessReviewRecurrenceRangeType{
		AccessReviewRecurrenceRangeTypeEndDate,
		AccessReviewRecurrenceRangeTypeNoEnd,
		AccessReviewRecurrenceRangeTypeNumbered,
	}
}

// ToPtr returns a *AccessReviewRecurrenceRangeType pointing to the current value.
func (c AccessReviewRecurrenceRangeType) ToPtr() *AccessReviewRecurrenceRangeType {
	return &c
}

// AccessReviewResult - The decision on the approval step. This value is initially set to NotReviewed. Approvers can take action of Approve/Deny
type AccessReviewResult string

const (
	AccessReviewResultApprove     AccessReviewResult = "Approve"
	AccessReviewResultDeny        AccessReviewResult = "Deny"
	AccessReviewResultDontKnow    AccessReviewResult = "DontKnow"
	AccessReviewResultNotNotified AccessReviewResult = "NotNotified"
	AccessReviewResultNotReviewed AccessReviewResult = "NotReviewed"
)

// PossibleAccessReviewResultValues returns the possible values for the AccessReviewResult const type.
func PossibleAccessReviewResultValues() []AccessReviewResult {
	return []AccessReviewResult{
		AccessReviewResultApprove,
		AccessReviewResultDeny,
		AccessReviewResultDontKnow,
		AccessReviewResultNotNotified,
		AccessReviewResultNotReviewed,
	}
}

// ToPtr returns a *AccessReviewResult pointing to the current value.
func (c AccessReviewResult) ToPtr() *AccessReviewResult {
	return &c
}

// AccessReviewReviewerType - The identity type : user/servicePrincipal
type AccessReviewReviewerType string

const (
	AccessReviewReviewerTypeServicePrincipal AccessReviewReviewerType = "servicePrincipal"
	AccessReviewReviewerTypeUser             AccessReviewReviewerType = "user"
)

// PossibleAccessReviewReviewerTypeValues returns the possible values for the AccessReviewReviewerType const type.
func PossibleAccessReviewReviewerTypeValues() []AccessReviewReviewerType {
	return []AccessReviewReviewerType{
		AccessReviewReviewerTypeServicePrincipal,
		AccessReviewReviewerTypeUser,
	}
}

// ToPtr returns a *AccessReviewReviewerType pointing to the current value.
func (c AccessReviewReviewerType) ToPtr() *AccessReviewReviewerType {
	return &c
}

// AccessReviewScheduleDefinitionReviewersType - This field specifies the type of reviewers for a review. Usually for a review, reviewers are explicitly
// assigned. However, in some cases, the reviewers may not be assigned and instead be chosen
// dynamically. For example managers review or self review.
type AccessReviewScheduleDefinitionReviewersType string

const (
	AccessReviewScheduleDefinitionReviewersTypeAssigned AccessReviewScheduleDefinitionReviewersType = "Assigned"
	AccessReviewScheduleDefinitionReviewersTypeManagers AccessReviewScheduleDefinitionReviewersType = "Managers"
	AccessReviewScheduleDefinitionReviewersTypeSelf     AccessReviewScheduleDefinitionReviewersType = "Self"
)

// PossibleAccessReviewScheduleDefinitionReviewersTypeValues returns the possible values for the AccessReviewScheduleDefinitionReviewersType const type.
func PossibleAccessReviewScheduleDefinitionReviewersTypeValues() []AccessReviewScheduleDefinitionReviewersType {
	return []AccessReviewScheduleDefinitionReviewersType{
		AccessReviewScheduleDefinitionReviewersTypeAssigned,
		AccessReviewScheduleDefinitionReviewersTypeManagers,
		AccessReviewScheduleDefinitionReviewersTypeSelf,
	}
}

// ToPtr returns a *AccessReviewScheduleDefinitionReviewersType pointing to the current value.
func (c AccessReviewScheduleDefinitionReviewersType) ToPtr() *AccessReviewScheduleDefinitionReviewersType {
	return &c
}

// AccessReviewScheduleDefinitionStatus - This read-only field specifies the status of an accessReview.
type AccessReviewScheduleDefinitionStatus string

const (
	AccessReviewScheduleDefinitionStatusApplied       AccessReviewScheduleDefinitionStatus = "Applied"
	AccessReviewScheduleDefinitionStatusApplying      AccessReviewScheduleDefinitionStatus = "Applying"
	AccessReviewScheduleDefinitionStatusAutoReviewed  AccessReviewScheduleDefinitionStatus = "AutoReviewed"
	AccessReviewScheduleDefinitionStatusAutoReviewing AccessReviewScheduleDefinitionStatus = "AutoReviewing"
	AccessReviewScheduleDefinitionStatusCompleted     AccessReviewScheduleDefinitionStatus = "Completed"
	AccessReviewScheduleDefinitionStatusCompleting    AccessReviewScheduleDefinitionStatus = "Completing"
	AccessReviewScheduleDefinitionStatusInProgress    AccessReviewScheduleDefinitionStatus = "InProgress"
	AccessReviewScheduleDefinitionStatusInitializing  AccessReviewScheduleDefinitionStatus = "Initializing"
	AccessReviewScheduleDefinitionStatusNotStarted    AccessReviewScheduleDefinitionStatus = "NotStarted"
	AccessReviewScheduleDefinitionStatusScheduled     AccessReviewScheduleDefinitionStatus = "Scheduled"
	AccessReviewScheduleDefinitionStatusStarting      AccessReviewScheduleDefinitionStatus = "Starting"
)

// PossibleAccessReviewScheduleDefinitionStatusValues returns the possible values for the AccessReviewScheduleDefinitionStatus const type.
func PossibleAccessReviewScheduleDefinitionStatusValues() []AccessReviewScheduleDefinitionStatus {
	return []AccessReviewScheduleDefinitionStatus{
		AccessReviewScheduleDefinitionStatusApplied,
		AccessReviewScheduleDefinitionStatusApplying,
		AccessReviewScheduleDefinitionStatusAutoReviewed,
		AccessReviewScheduleDefinitionStatusAutoReviewing,
		AccessReviewScheduleDefinitionStatusCompleted,
		AccessReviewScheduleDefinitionStatusCompleting,
		AccessReviewScheduleDefinitionStatusInProgress,
		AccessReviewScheduleDefinitionStatusInitializing,
		AccessReviewScheduleDefinitionStatusNotStarted,
		AccessReviewScheduleDefinitionStatusScheduled,
		AccessReviewScheduleDefinitionStatusStarting,
	}
}

// ToPtr returns a *AccessReviewScheduleDefinitionStatus pointing to the current value.
func (c AccessReviewScheduleDefinitionStatus) ToPtr() *AccessReviewScheduleDefinitionStatus {
	return &c
}

// AccessReviewScopePrincipalType - The identity type user/servicePrincipal to review
type AccessReviewScopePrincipalType string

const (
	AccessReviewScopePrincipalTypeServicePrincipal AccessReviewScopePrincipalType = "servicePrincipal"
	AccessReviewScopePrincipalTypeUser             AccessReviewScopePrincipalType = "user"
)

// PossibleAccessReviewScopePrincipalTypeValues returns the possible values for the AccessReviewScopePrincipalType const type.
func PossibleAccessReviewScopePrincipalTypeValues() []AccessReviewScopePrincipalType {
	return []AccessReviewScopePrincipalType{
		AccessReviewScopePrincipalTypeServicePrincipal,
		AccessReviewScopePrincipalTypeUser,
	}
}

// ToPtr returns a *AccessReviewScopePrincipalType pointing to the current value.
func (c AccessReviewScopePrincipalType) ToPtr() *AccessReviewScopePrincipalType {
	return &c
}

// ApprovalMode - The type of rule
type ApprovalMode string

const (
	ApprovalModeNoApproval  ApprovalMode = "NoApproval"
	ApprovalModeParallel    ApprovalMode = "Parallel"
	ApprovalModeSerial      ApprovalMode = "Serial"
	ApprovalModeSingleStage ApprovalMode = "SingleStage"
)

// PossibleApprovalModeValues returns the possible values for the ApprovalMode const type.
func PossibleApprovalModeValues() []ApprovalMode {
	return []ApprovalMode{
		ApprovalModeNoApproval,
		ApprovalModeParallel,
		ApprovalModeSerial,
		ApprovalModeSingleStage,
	}
}

// ToPtr returns a *ApprovalMode pointing to the current value.
func (c ApprovalMode) ToPtr() *ApprovalMode {
	return &c
}

// AssignmentType - Assignment type of the role assignment schedule
type AssignmentType string

const (
	AssignmentTypeActivated AssignmentType = "Activated"
	AssignmentTypeAssigned  AssignmentType = "Assigned"
)

// PossibleAssignmentTypeValues returns the possible values for the AssignmentType const type.
func PossibleAssignmentTypeValues() []AssignmentType {
	return []AssignmentType{
		AssignmentTypeActivated,
		AssignmentTypeAssigned,
	}
}

// ToPtr returns a *AssignmentType pointing to the current value.
func (c AssignmentType) ToPtr() *AssignmentType {
	return &c
}

// DecisionTargetType - The type of decision target : User/ServicePrincipal
type DecisionTargetType string

const (
	DecisionTargetTypeServicePrincipal DecisionTargetType = "servicePrincipal"
	DecisionTargetTypeUser             DecisionTargetType = "user"
)

// PossibleDecisionTargetTypeValues returns the possible values for the DecisionTargetType const type.
func PossibleDecisionTargetTypeValues() []DecisionTargetType {
	return []DecisionTargetType{
		DecisionTargetTypeServicePrincipal,
		DecisionTargetTypeUser,
	}
}

// ToPtr returns a *DecisionTargetType pointing to the current value.
func (c DecisionTargetType) ToPtr() *DecisionTargetType {
	return &c
}

// DefaultDecisionType - This specifies the behavior for the autoReview feature when an access review completes.
type DefaultDecisionType string

const (
	DefaultDecisionTypeApprove        DefaultDecisionType = "Approve"
	DefaultDecisionTypeDeny           DefaultDecisionType = "Deny"
	DefaultDecisionTypeRecommendation DefaultDecisionType = "Recommendation"
)

// PossibleDefaultDecisionTypeValues returns the possible values for the DefaultDecisionType const type.
func PossibleDefaultDecisionTypeValues() []DefaultDecisionType {
	return []DefaultDecisionType{
		DefaultDecisionTypeApprove,
		DefaultDecisionTypeDeny,
		DefaultDecisionTypeRecommendation,
	}
}

// ToPtr returns a *DefaultDecisionType pointing to the current value.
func (c DefaultDecisionType) ToPtr() *DefaultDecisionType {
	return &c
}

// EnablementRules - The type of enable rules
type EnablementRules string

const (
	EnablementRulesJustification             EnablementRules = "Justification"
	EnablementRulesMultiFactorAuthentication EnablementRules = "MultiFactorAuthentication"
	EnablementRulesTicketing                 EnablementRules = "Ticketing"
)

// PossibleEnablementRulesValues returns the possible values for the EnablementRules const type.
func PossibleEnablementRulesValues() []EnablementRules {
	return []EnablementRules{
		EnablementRulesJustification,
		EnablementRulesMultiFactorAuthentication,
		EnablementRulesTicketing,
	}
}

// ToPtr returns a *EnablementRules pointing to the current value.
func (c EnablementRules) ToPtr() *EnablementRules {
	return &c
}

// MemberType - Membership type of the role assignment schedule
type MemberType string

const (
	MemberTypeDirect    MemberType = "Direct"
	MemberTypeGroup     MemberType = "Group"
	MemberTypeInherited MemberType = "Inherited"
)

// PossibleMemberTypeValues returns the possible values for the MemberType const type.
func PossibleMemberTypeValues() []MemberType {
	return []MemberType{
		MemberTypeDirect,
		MemberTypeGroup,
		MemberTypeInherited,
	}
}

// ToPtr returns a *MemberType pointing to the current value.
func (c MemberType) ToPtr() *MemberType {
	return &c
}

// NotificationDeliveryMechanism - The type of notification.
type NotificationDeliveryMechanism string

const (
	NotificationDeliveryMechanismEmail NotificationDeliveryMechanism = "Email"
)

// PossibleNotificationDeliveryMechanismValues returns the possible values for the NotificationDeliveryMechanism const type.
func PossibleNotificationDeliveryMechanismValues() []NotificationDeliveryMechanism {
	return []NotificationDeliveryMechanism{
		NotificationDeliveryMechanismEmail,
	}
}

// ToPtr returns a *NotificationDeliveryMechanism pointing to the current value.
func (c NotificationDeliveryMechanism) ToPtr() *NotificationDeliveryMechanism {
	return &c
}

// NotificationLevel - The notification level.
type NotificationLevel string

const (
	NotificationLevelAll      NotificationLevel = "All"
	NotificationLevelCritical NotificationLevel = "Critical"
	NotificationLevelNone     NotificationLevel = "None"
)

// PossibleNotificationLevelValues returns the possible values for the NotificationLevel const type.
func PossibleNotificationLevelValues() []NotificationLevel {
	return []NotificationLevel{
		NotificationLevelAll,
		NotificationLevelCritical,
		NotificationLevelNone,
	}
}

// ToPtr returns a *NotificationLevel pointing to the current value.
func (c NotificationLevel) ToPtr() *NotificationLevel {
	return &c
}

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

// ToPtr returns a *PrincipalType pointing to the current value.
func (c PrincipalType) ToPtr() *PrincipalType {
	return &c
}

// RecipientType - The recipient type.
type RecipientType string

const (
	RecipientTypeAdmin     RecipientType = "Admin"
	RecipientTypeApprover  RecipientType = "Approver"
	RecipientTypeRequestor RecipientType = "Requestor"
)

// PossibleRecipientTypeValues returns the possible values for the RecipientType const type.
func PossibleRecipientTypeValues() []RecipientType {
	return []RecipientType{
		RecipientTypeAdmin,
		RecipientTypeApprover,
		RecipientTypeRequestor,
	}
}

// ToPtr returns a *RecipientType pointing to the current value.
func (c RecipientType) ToPtr() *RecipientType {
	return &c
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

// ToPtr returns a *RequestType pointing to the current value.
func (c RequestType) ToPtr() *RequestType {
	return &c
}

// RoleManagementPolicyRuleType - The type of rule
type RoleManagementPolicyRuleType string

const (
	RoleManagementPolicyRuleTypeRoleManagementPolicyApprovalRule              RoleManagementPolicyRuleType = "RoleManagementPolicyApprovalRule"
	RoleManagementPolicyRuleTypeRoleManagementPolicyAuthenticationContextRule RoleManagementPolicyRuleType = "RoleManagementPolicyAuthenticationContextRule"
	RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule            RoleManagementPolicyRuleType = "RoleManagementPolicyEnablementRule"
	RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule            RoleManagementPolicyRuleType = "RoleManagementPolicyExpirationRule"
	RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule          RoleManagementPolicyRuleType = "RoleManagementPolicyNotificationRule"
)

// PossibleRoleManagementPolicyRuleTypeValues returns the possible values for the RoleManagementPolicyRuleType const type.
func PossibleRoleManagementPolicyRuleTypeValues() []RoleManagementPolicyRuleType {
	return []RoleManagementPolicyRuleType{
		RoleManagementPolicyRuleTypeRoleManagementPolicyApprovalRule,
		RoleManagementPolicyRuleTypeRoleManagementPolicyAuthenticationContextRule,
		RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule,
		RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule,
		RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule,
	}
}

// ToPtr returns a *RoleManagementPolicyRuleType pointing to the current value.
func (c RoleManagementPolicyRuleType) ToPtr() *RoleManagementPolicyRuleType {
	return &c
}

// Status - The status of the role assignment schedule.
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

// ToPtr returns a *Status pointing to the current value.
func (c Status) ToPtr() *Status {
	return &c
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

// ToPtr returns a *Type pointing to the current value.
func (c Type) ToPtr() *Type {
	return &c
}

// UserType - The type of user.
type UserType string

const (
	UserTypeGroup UserType = "Group"
	UserTypeUser  UserType = "User"
)

// PossibleUserTypeValues returns the possible values for the UserType const type.
func PossibleUserTypeValues() []UserType {
	return []UserType{
		UserTypeGroup,
		UserTypeUser,
	}
}

// ToPtr returns a *UserType pointing to the current value.
func (c UserType) ToPtr() *UserType {
	return &c
}
