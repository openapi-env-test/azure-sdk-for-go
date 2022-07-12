//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicebus

const (
	moduleName    = "armservicebus"
	moduleVersion = "v0.4.0"
)

type AccessRights string

const (
	AccessRightsListen AccessRights = "Listen"
	AccessRightsManage AccessRights = "Manage"
	AccessRightsSend   AccessRights = "Send"
)

// PossibleAccessRightsValues returns the possible values for the AccessRights const type.
func PossibleAccessRightsValues() []AccessRights {
	return []AccessRights{
		AccessRightsListen,
		AccessRightsManage,
		AccessRightsSend,
	}
}

// ToPtr returns a *AccessRights pointing to the current value.
func (c AccessRights) ToPtr() *AccessRights {
	return &c
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// DefaultAction - Default Action for Network Rule Set
type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns the possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{
		DefaultActionAllow,
		DefaultActionDeny,
	}
}

// ToPtr returns a *DefaultAction pointing to the current value.
func (c DefaultAction) ToPtr() *DefaultAction {
	return &c
}

// EndPointProvisioningState - Provisioning state of the Private Endpoint Connection.
type EndPointProvisioningState string

const (
	EndPointProvisioningStateCanceled  EndPointProvisioningState = "Canceled"
	EndPointProvisioningStateCreating  EndPointProvisioningState = "Creating"
	EndPointProvisioningStateDeleting  EndPointProvisioningState = "Deleting"
	EndPointProvisioningStateFailed    EndPointProvisioningState = "Failed"
	EndPointProvisioningStateSucceeded EndPointProvisioningState = "Succeeded"
	EndPointProvisioningStateUpdating  EndPointProvisioningState = "Updating"
)

// PossibleEndPointProvisioningStateValues returns the possible values for the EndPointProvisioningState const type.
func PossibleEndPointProvisioningStateValues() []EndPointProvisioningState {
	return []EndPointProvisioningState{
		EndPointProvisioningStateCanceled,
		EndPointProvisioningStateCreating,
		EndPointProvisioningStateDeleting,
		EndPointProvisioningStateFailed,
		EndPointProvisioningStateSucceeded,
		EndPointProvisioningStateUpdating,
	}
}

// ToPtr returns a *EndPointProvisioningState pointing to the current value.
func (c EndPointProvisioningState) ToPtr() *EndPointProvisioningState {
	return &c
}

// EntityStatus - Entity status.
type EntityStatus string

const (
	EntityStatusActive          EntityStatus = "Active"
	EntityStatusDisabled        EntityStatus = "Disabled"
	EntityStatusRestoring       EntityStatus = "Restoring"
	EntityStatusSendDisabled    EntityStatus = "SendDisabled"
	EntityStatusReceiveDisabled EntityStatus = "ReceiveDisabled"
	EntityStatusCreating        EntityStatus = "Creating"
	EntityStatusDeleting        EntityStatus = "Deleting"
	EntityStatusRenaming        EntityStatus = "Renaming"
	EntityStatusUnknown         EntityStatus = "Unknown"
)

// PossibleEntityStatusValues returns the possible values for the EntityStatus const type.
func PossibleEntityStatusValues() []EntityStatus {
	return []EntityStatus{
		EntityStatusActive,
		EntityStatusDisabled,
		EntityStatusRestoring,
		EntityStatusSendDisabled,
		EntityStatusReceiveDisabled,
		EntityStatusCreating,
		EntityStatusDeleting,
		EntityStatusRenaming,
		EntityStatusUnknown,
	}
}

// ToPtr returns a *EntityStatus pointing to the current value.
func (c EntityStatus) ToPtr() *EntityStatus {
	return &c
}

// FilterType - Rule filter types
type FilterType string

const (
	FilterTypeCorrelationFilter FilterType = "CorrelationFilter"
	FilterTypeSQLFilter         FilterType = "SqlFilter"
)

// PossibleFilterTypeValues returns the possible values for the FilterType const type.
func PossibleFilterTypeValues() []FilterType {
	return []FilterType{
		FilterTypeCorrelationFilter,
		FilterTypeSQLFilter,
	}
}

// ToPtr returns a *FilterType pointing to the current value.
func (c FilterType) ToPtr() *FilterType {
	return &c
}

// KeySource - Enumerates the possible value of keySource for Encryption
type KeySource string

const (
	KeySourceMicrosoftKeyVault KeySource = "Microsoft.KeyVault"
)

// PossibleKeySourceValues returns the possible values for the KeySource const type.
func PossibleKeySourceValues() []KeySource {
	return []KeySource{
		KeySourceMicrosoftKeyVault,
	}
}

// ToPtr returns a *KeySource pointing to the current value.
func (c KeySource) ToPtr() *KeySource {
	return &c
}

// KeyType - The access key to regenerate.
type KeyType string

const (
	KeyTypePrimaryKey   KeyType = "PrimaryKey"
	KeyTypeSecondaryKey KeyType = "SecondaryKey"
)

// PossibleKeyTypeValues returns the possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypePrimaryKey,
		KeyTypeSecondaryKey,
	}
}

// ToPtr returns a *KeyType pointing to the current value.
func (c KeyType) ToPtr() *KeyType {
	return &c
}

// ManagedServiceIdentityType - Type of managed service identity.
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned, UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// ToPtr returns a *ManagedServiceIdentityType pointing to the current value.
func (c ManagedServiceIdentityType) ToPtr() *ManagedServiceIdentityType {
	return &c
}

type MigrationConfigurationName string

const (
	MigrationConfigurationNameDefault MigrationConfigurationName = "$default"
)

// PossibleMigrationConfigurationNameValues returns the possible values for the MigrationConfigurationName const type.
func PossibleMigrationConfigurationNameValues() []MigrationConfigurationName {
	return []MigrationConfigurationName{
		MigrationConfigurationNameDefault,
	}
}

// ToPtr returns a *MigrationConfigurationName pointing to the current value.
func (c MigrationConfigurationName) ToPtr() *MigrationConfigurationName {
	return &c
}

// NetworkRuleIPAction - The IP Filter Action
type NetworkRuleIPAction string

const (
	NetworkRuleIPActionAllow NetworkRuleIPAction = "Allow"
)

// PossibleNetworkRuleIPActionValues returns the possible values for the NetworkRuleIPAction const type.
func PossibleNetworkRuleIPActionValues() []NetworkRuleIPAction {
	return []NetworkRuleIPAction{
		NetworkRuleIPActionAllow,
	}
}

// ToPtr returns a *NetworkRuleIPAction pointing to the current value.
func (c NetworkRuleIPAction) ToPtr() *NetworkRuleIPAction {
	return &c
}

// PrivateLinkConnectionStatus - Status of the connection.
type PrivateLinkConnectionStatus string

const (
	PrivateLinkConnectionStatusApproved     PrivateLinkConnectionStatus = "Approved"
	PrivateLinkConnectionStatusDisconnected PrivateLinkConnectionStatus = "Disconnected"
	PrivateLinkConnectionStatusPending      PrivateLinkConnectionStatus = "Pending"
	PrivateLinkConnectionStatusRejected     PrivateLinkConnectionStatus = "Rejected"
)

// PossiblePrivateLinkConnectionStatusValues returns the possible values for the PrivateLinkConnectionStatus const type.
func PossiblePrivateLinkConnectionStatusValues() []PrivateLinkConnectionStatus {
	return []PrivateLinkConnectionStatus{
		PrivateLinkConnectionStatusApproved,
		PrivateLinkConnectionStatusDisconnected,
		PrivateLinkConnectionStatusPending,
		PrivateLinkConnectionStatusRejected,
	}
}

// ToPtr returns a *PrivateLinkConnectionStatus pointing to the current value.
func (c PrivateLinkConnectionStatus) ToPtr() *PrivateLinkConnectionStatus {
	return &c
}

// ProvisioningStateDR - Provisioning state of the Alias(Disaster Recovery configuration) - possible values 'Accepted' or
// 'Succeeded' or 'Failed'
type ProvisioningStateDR string

const (
	ProvisioningStateDRAccepted  ProvisioningStateDR = "Accepted"
	ProvisioningStateDRFailed    ProvisioningStateDR = "Failed"
	ProvisioningStateDRSucceeded ProvisioningStateDR = "Succeeded"
)

// PossibleProvisioningStateDRValues returns the possible values for the ProvisioningStateDR const type.
func PossibleProvisioningStateDRValues() []ProvisioningStateDR {
	return []ProvisioningStateDR{
		ProvisioningStateDRAccepted,
		ProvisioningStateDRFailed,
		ProvisioningStateDRSucceeded,
	}
}

// ToPtr returns a *ProvisioningStateDR pointing to the current value.
func (c ProvisioningStateDR) ToPtr() *ProvisioningStateDR {
	return &c
}

// PublicNetworkAccess - This determines if traffic is allowed over public network. By default it is enabled.
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled           PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled            PublicNetworkAccess = "Enabled"
	PublicNetworkAccessSecuredByPerimeter PublicNetworkAccess = "SecuredByPerimeter"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
		PublicNetworkAccessSecuredByPerimeter,
	}
}

// ToPtr returns a *PublicNetworkAccess pointing to the current value.
func (c PublicNetworkAccess) ToPtr() *PublicNetworkAccess {
	return &c
}

// PublicNetworkAccessFlag - This determines if traffic is allowed over public network. By default it is enabled.
type PublicNetworkAccessFlag string

const (
	PublicNetworkAccessFlagDisabled PublicNetworkAccessFlag = "Disabled"
	PublicNetworkAccessFlagEnabled  PublicNetworkAccessFlag = "Enabled"
)

// PossiblePublicNetworkAccessFlagValues returns the possible values for the PublicNetworkAccessFlag const type.
func PossiblePublicNetworkAccessFlagValues() []PublicNetworkAccessFlag {
	return []PublicNetworkAccessFlag{
		PublicNetworkAccessFlagDisabled,
		PublicNetworkAccessFlagEnabled,
	}
}

// ToPtr returns a *PublicNetworkAccessFlag pointing to the current value.
func (c PublicNetworkAccessFlag) ToPtr() *PublicNetworkAccessFlag {
	return &c
}

// RoleDisasterRecovery - role of namespace in GEO DR - possible values 'Primary' or 'PrimaryNotReplicating' or 'Secondary'
type RoleDisasterRecovery string

const (
	RoleDisasterRecoveryPrimary               RoleDisasterRecovery = "Primary"
	RoleDisasterRecoveryPrimaryNotReplicating RoleDisasterRecovery = "PrimaryNotReplicating"
	RoleDisasterRecoverySecondary             RoleDisasterRecovery = "Secondary"
)

// PossibleRoleDisasterRecoveryValues returns the possible values for the RoleDisasterRecovery const type.
func PossibleRoleDisasterRecoveryValues() []RoleDisasterRecovery {
	return []RoleDisasterRecovery{
		RoleDisasterRecoveryPrimary,
		RoleDisasterRecoveryPrimaryNotReplicating,
		RoleDisasterRecoverySecondary,
	}
}

// ToPtr returns a *RoleDisasterRecovery pointing to the current value.
func (c RoleDisasterRecovery) ToPtr() *RoleDisasterRecovery {
	return &c
}

// SKUName - Name of this SKU.
type SKUName string

const (
	SKUNameBasic    SKUName = "Basic"
	SKUNamePremium  SKUName = "Premium"
	SKUNameStandard SKUName = "Standard"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameBasic,
		SKUNamePremium,
		SKUNameStandard,
	}
}

// ToPtr returns a *SKUName pointing to the current value.
func (c SKUName) ToPtr() *SKUName {
	return &c
}

// SKUTier - The billing tier of this particular SKU.
type SKUTier string

const (
	SKUTierBasic    SKUTier = "Basic"
	SKUTierPremium  SKUTier = "Premium"
	SKUTierStandard SKUTier = "Standard"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierBasic,
		SKUTierPremium,
		SKUTierStandard,
	}
}

// ToPtr returns a *SKUTier pointing to the current value.
func (c SKUTier) ToPtr() *SKUTier {
	return &c
}

// TLSVersion - The minimum TLS version for the cluster to support, e.g. '1.2'
type TLSVersion string

const (
	TLSVersionOne0 TLSVersion = "1.0"
	TLSVersionOne1 TLSVersion = "1.1"
	TLSVersionOne2 TLSVersion = "1.2"
)

// PossibleTLSVersionValues returns the possible values for the TLSVersion const type.
func PossibleTLSVersionValues() []TLSVersion {
	return []TLSVersion{
		TLSVersionOne0,
		TLSVersionOne1,
		TLSVersionOne2,
	}
}

// ToPtr returns a *TLSVersion pointing to the current value.
func (c TLSVersion) ToPtr() *TLSVersion {
	return &c
}

// UnavailableReason - Specifies the reason for the unavailability of the service.
type UnavailableReason string

const (
	UnavailableReasonInvalidName                           UnavailableReason = "InvalidName"
	UnavailableReasonNameInLockdown                        UnavailableReason = "NameInLockdown"
	UnavailableReasonNameInUse                             UnavailableReason = "NameInUse"
	UnavailableReasonNone                                  UnavailableReason = "None"
	UnavailableReasonSubscriptionIsDisabled                UnavailableReason = "SubscriptionIsDisabled"
	UnavailableReasonTooManyNamespaceInCurrentSubscription UnavailableReason = "TooManyNamespaceInCurrentSubscription"
)

// PossibleUnavailableReasonValues returns the possible values for the UnavailableReason const type.
func PossibleUnavailableReasonValues() []UnavailableReason {
	return []UnavailableReason{
		UnavailableReasonInvalidName,
		UnavailableReasonNameInLockdown,
		UnavailableReasonNameInUse,
		UnavailableReasonNone,
		UnavailableReasonSubscriptionIsDisabled,
		UnavailableReasonTooManyNamespaceInCurrentSubscription,
	}
}

// ToPtr returns a *UnavailableReason pointing to the current value.
func (c UnavailableReason) ToPtr() *UnavailableReason {
	return &c
}
