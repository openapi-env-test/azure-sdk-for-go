package automation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccountState enumerates the values for account state.
type AccountState string

const (
	// AccountStateOk ...
	AccountStateOk AccountState = "Ok"
	// AccountStateSuspended ...
	AccountStateSuspended AccountState = "Suspended"
	// AccountStateUnavailable ...
	AccountStateUnavailable AccountState = "Unavailable"
)

// PossibleAccountStateValues returns an array of possible values for the AccountState const type.
func PossibleAccountStateValues() []AccountState {
	return []AccountState{AccountStateOk, AccountStateSuspended, AccountStateUnavailable}
}

// AgentRegistrationKeyName enumerates the values for agent registration key name.
type AgentRegistrationKeyName string

const (
	// AgentRegistrationKeyNamePrimary ...
	AgentRegistrationKeyNamePrimary AgentRegistrationKeyName = "primary"
	// AgentRegistrationKeyNameSecondary ...
	AgentRegistrationKeyNameSecondary AgentRegistrationKeyName = "secondary"
)

// PossibleAgentRegistrationKeyNameValues returns an array of possible values for the AgentRegistrationKeyName const type.
func PossibleAgentRegistrationKeyNameValues() []AgentRegistrationKeyName {
	return []AgentRegistrationKeyName{AgentRegistrationKeyNamePrimary, AgentRegistrationKeyNameSecondary}
}

// ContentSourceType enumerates the values for content source type.
type ContentSourceType string

const (
	// ContentSourceTypeEmbeddedContent ...
	ContentSourceTypeEmbeddedContent ContentSourceType = "embeddedContent"
	// ContentSourceTypeURI ...
	ContentSourceTypeURI ContentSourceType = "uri"
)

// PossibleContentSourceTypeValues returns an array of possible values for the ContentSourceType const type.
func PossibleContentSourceTypeValues() []ContentSourceType {
	return []ContentSourceType{ContentSourceTypeEmbeddedContent, ContentSourceTypeURI}
}

// CountType enumerates the values for count type.
type CountType string

const (
	// CountTypeNodeconfiguration ...
	CountTypeNodeconfiguration CountType = "nodeconfiguration"
	// CountTypeStatus ...
	CountTypeStatus CountType = "status"
)

// PossibleCountTypeValues returns an array of possible values for the CountType const type.
func PossibleCountTypeValues() []CountType {
	return []CountType{CountTypeNodeconfiguration, CountTypeStatus}
}

// DscConfigurationProvisioningState enumerates the values for dsc configuration provisioning state.
type DscConfigurationProvisioningState string

const (
	// DscConfigurationProvisioningStateSucceeded ...
	DscConfigurationProvisioningStateSucceeded DscConfigurationProvisioningState = "Succeeded"
)

// PossibleDscConfigurationProvisioningStateValues returns an array of possible values for the DscConfigurationProvisioningState const type.
func PossibleDscConfigurationProvisioningStateValues() []DscConfigurationProvisioningState {
	return []DscConfigurationProvisioningState{DscConfigurationProvisioningStateSucceeded}
}

// DscConfigurationState enumerates the values for dsc configuration state.
type DscConfigurationState string

const (
	// DscConfigurationStateEdit ...
	DscConfigurationStateEdit DscConfigurationState = "Edit"
	// DscConfigurationStateNew ...
	DscConfigurationStateNew DscConfigurationState = "New"
	// DscConfigurationStatePublished ...
	DscConfigurationStatePublished DscConfigurationState = "Published"
)

// PossibleDscConfigurationStateValues returns an array of possible values for the DscConfigurationState const type.
func PossibleDscConfigurationStateValues() []DscConfigurationState {
	return []DscConfigurationState{DscConfigurationStateEdit, DscConfigurationStateNew, DscConfigurationStatePublished}
}

// EncryptionKeySourceType enumerates the values for encryption key source type.
type EncryptionKeySourceType string

const (
	// EncryptionKeySourceTypeMicrosoftAutomation ...
	EncryptionKeySourceTypeMicrosoftAutomation EncryptionKeySourceType = "Microsoft.Automation"
	// EncryptionKeySourceTypeMicrosoftKeyvault ...
	EncryptionKeySourceTypeMicrosoftKeyvault EncryptionKeySourceType = "Microsoft.Keyvault"
)

// PossibleEncryptionKeySourceTypeValues returns an array of possible values for the EncryptionKeySourceType const type.
func PossibleEncryptionKeySourceTypeValues() []EncryptionKeySourceType {
	return []EncryptionKeySourceType{EncryptionKeySourceTypeMicrosoftAutomation, EncryptionKeySourceTypeMicrosoftKeyvault}
}

// GroupTypeEnum enumerates the values for group type enum.
type GroupTypeEnum string

const (
	// GroupTypeEnumSystem ...
	GroupTypeEnumSystem GroupTypeEnum = "System"
	// GroupTypeEnumUser ...
	GroupTypeEnumUser GroupTypeEnum = "User"
)

// PossibleGroupTypeEnumValues returns an array of possible values for the GroupTypeEnum const type.
func PossibleGroupTypeEnumValues() []GroupTypeEnum {
	return []GroupTypeEnum{GroupTypeEnumSystem, GroupTypeEnumUser}
}

// HTTPStatusCode enumerates the values for http status code.
type HTTPStatusCode string

const (
	// HTTPStatusCodeAccepted ...
	HTTPStatusCodeAccepted HTTPStatusCode = "Accepted"
	// HTTPStatusCodeAmbiguous ...
	HTTPStatusCodeAmbiguous HTTPStatusCode = "Ambiguous"
	// HTTPStatusCodeBadGateway ...
	HTTPStatusCodeBadGateway HTTPStatusCode = "BadGateway"
	// HTTPStatusCodeBadRequest ...
	HTTPStatusCodeBadRequest HTTPStatusCode = "BadRequest"
	// HTTPStatusCodeConflict ...
	HTTPStatusCodeConflict HTTPStatusCode = "Conflict"
	// HTTPStatusCodeContinue ...
	HTTPStatusCodeContinue HTTPStatusCode = "Continue"
	// HTTPStatusCodeCreated ...
	HTTPStatusCodeCreated HTTPStatusCode = "Created"
	// HTTPStatusCodeExpectationFailed ...
	HTTPStatusCodeExpectationFailed HTTPStatusCode = "ExpectationFailed"
	// HTTPStatusCodeForbidden ...
	HTTPStatusCodeForbidden HTTPStatusCode = "Forbidden"
	// HTTPStatusCodeFound ...
	HTTPStatusCodeFound HTTPStatusCode = "Found"
	// HTTPStatusCodeGatewayTimeout ...
	HTTPStatusCodeGatewayTimeout HTTPStatusCode = "GatewayTimeout"
	// HTTPStatusCodeGone ...
	HTTPStatusCodeGone HTTPStatusCode = "Gone"
	// HTTPStatusCodeHTTPVersionNotSupported ...
	HTTPStatusCodeHTTPVersionNotSupported HTTPStatusCode = "HttpVersionNotSupported"
	// HTTPStatusCodeInternalServerError ...
	HTTPStatusCodeInternalServerError HTTPStatusCode = "InternalServerError"
	// HTTPStatusCodeLengthRequired ...
	HTTPStatusCodeLengthRequired HTTPStatusCode = "LengthRequired"
	// HTTPStatusCodeMethodNotAllowed ...
	HTTPStatusCodeMethodNotAllowed HTTPStatusCode = "MethodNotAllowed"
	// HTTPStatusCodeMoved ...
	HTTPStatusCodeMoved HTTPStatusCode = "Moved"
	// HTTPStatusCodeMovedPermanently ...
	HTTPStatusCodeMovedPermanently HTTPStatusCode = "MovedPermanently"
	// HTTPStatusCodeMultipleChoices ...
	HTTPStatusCodeMultipleChoices HTTPStatusCode = "MultipleChoices"
	// HTTPStatusCodeNoContent ...
	HTTPStatusCodeNoContent HTTPStatusCode = "NoContent"
	// HTTPStatusCodeNonAuthoritativeInformation ...
	HTTPStatusCodeNonAuthoritativeInformation HTTPStatusCode = "NonAuthoritativeInformation"
	// HTTPStatusCodeNotAcceptable ...
	HTTPStatusCodeNotAcceptable HTTPStatusCode = "NotAcceptable"
	// HTTPStatusCodeNotFound ...
	HTTPStatusCodeNotFound HTTPStatusCode = "NotFound"
	// HTTPStatusCodeNotImplemented ...
	HTTPStatusCodeNotImplemented HTTPStatusCode = "NotImplemented"
	// HTTPStatusCodeNotModified ...
	HTTPStatusCodeNotModified HTTPStatusCode = "NotModified"
	// HTTPStatusCodeOK ...
	HTTPStatusCodeOK HTTPStatusCode = "OK"
	// HTTPStatusCodePartialContent ...
	HTTPStatusCodePartialContent HTTPStatusCode = "PartialContent"
	// HTTPStatusCodePaymentRequired ...
	HTTPStatusCodePaymentRequired HTTPStatusCode = "PaymentRequired"
	// HTTPStatusCodePreconditionFailed ...
	HTTPStatusCodePreconditionFailed HTTPStatusCode = "PreconditionFailed"
	// HTTPStatusCodeProxyAuthenticationRequired ...
	HTTPStatusCodeProxyAuthenticationRequired HTTPStatusCode = "ProxyAuthenticationRequired"
	// HTTPStatusCodeRedirect ...
	HTTPStatusCodeRedirect HTTPStatusCode = "Redirect"
	// HTTPStatusCodeRedirectKeepVerb ...
	HTTPStatusCodeRedirectKeepVerb HTTPStatusCode = "RedirectKeepVerb"
	// HTTPStatusCodeRedirectMethod ...
	HTTPStatusCodeRedirectMethod HTTPStatusCode = "RedirectMethod"
	// HTTPStatusCodeRequestedRangeNotSatisfiable ...
	HTTPStatusCodeRequestedRangeNotSatisfiable HTTPStatusCode = "RequestedRangeNotSatisfiable"
	// HTTPStatusCodeRequestEntityTooLarge ...
	HTTPStatusCodeRequestEntityTooLarge HTTPStatusCode = "RequestEntityTooLarge"
	// HTTPStatusCodeRequestTimeout ...
	HTTPStatusCodeRequestTimeout HTTPStatusCode = "RequestTimeout"
	// HTTPStatusCodeRequestURITooLong ...
	HTTPStatusCodeRequestURITooLong HTTPStatusCode = "RequestUriTooLong"
	// HTTPStatusCodeResetContent ...
	HTTPStatusCodeResetContent HTTPStatusCode = "ResetContent"
	// HTTPStatusCodeSeeOther ...
	HTTPStatusCodeSeeOther HTTPStatusCode = "SeeOther"
	// HTTPStatusCodeServiceUnavailable ...
	HTTPStatusCodeServiceUnavailable HTTPStatusCode = "ServiceUnavailable"
	// HTTPStatusCodeSwitchingProtocols ...
	HTTPStatusCodeSwitchingProtocols HTTPStatusCode = "SwitchingProtocols"
	// HTTPStatusCodeTemporaryRedirect ...
	HTTPStatusCodeTemporaryRedirect HTTPStatusCode = "TemporaryRedirect"
	// HTTPStatusCodeUnauthorized ...
	HTTPStatusCodeUnauthorized HTTPStatusCode = "Unauthorized"
	// HTTPStatusCodeUnsupportedMediaType ...
	HTTPStatusCodeUnsupportedMediaType HTTPStatusCode = "UnsupportedMediaType"
	// HTTPStatusCodeUnused ...
	HTTPStatusCodeUnused HTTPStatusCode = "Unused"
	// HTTPStatusCodeUpgradeRequired ...
	HTTPStatusCodeUpgradeRequired HTTPStatusCode = "UpgradeRequired"
	// HTTPStatusCodeUseProxy ...
	HTTPStatusCodeUseProxy HTTPStatusCode = "UseProxy"
)

// PossibleHTTPStatusCodeValues returns an array of possible values for the HTTPStatusCode const type.
func PossibleHTTPStatusCodeValues() []HTTPStatusCode {
	return []HTTPStatusCode{HTTPStatusCodeAccepted, HTTPStatusCodeAmbiguous, HTTPStatusCodeBadGateway, HTTPStatusCodeBadRequest, HTTPStatusCodeConflict, HTTPStatusCodeContinue, HTTPStatusCodeCreated, HTTPStatusCodeExpectationFailed, HTTPStatusCodeForbidden, HTTPStatusCodeFound, HTTPStatusCodeGatewayTimeout, HTTPStatusCodeGone, HTTPStatusCodeHTTPVersionNotSupported, HTTPStatusCodeInternalServerError, HTTPStatusCodeLengthRequired, HTTPStatusCodeMethodNotAllowed, HTTPStatusCodeMoved, HTTPStatusCodeMovedPermanently, HTTPStatusCodeMultipleChoices, HTTPStatusCodeNoContent, HTTPStatusCodeNonAuthoritativeInformation, HTTPStatusCodeNotAcceptable, HTTPStatusCodeNotFound, HTTPStatusCodeNotImplemented, HTTPStatusCodeNotModified, HTTPStatusCodeOK, HTTPStatusCodePartialContent, HTTPStatusCodePaymentRequired, HTTPStatusCodePreconditionFailed, HTTPStatusCodeProxyAuthenticationRequired, HTTPStatusCodeRedirect, HTTPStatusCodeRedirectKeepVerb, HTTPStatusCodeRedirectMethod, HTTPStatusCodeRequestedRangeNotSatisfiable, HTTPStatusCodeRequestEntityTooLarge, HTTPStatusCodeRequestTimeout, HTTPStatusCodeRequestURITooLong, HTTPStatusCodeResetContent, HTTPStatusCodeSeeOther, HTTPStatusCodeServiceUnavailable, HTTPStatusCodeSwitchingProtocols, HTTPStatusCodeTemporaryRedirect, HTTPStatusCodeUnauthorized, HTTPStatusCodeUnsupportedMediaType, HTTPStatusCodeUnused, HTTPStatusCodeUpgradeRequired, HTTPStatusCodeUseProxy}
}

// JobProvisioningState enumerates the values for job provisioning state.
type JobProvisioningState string

const (
	// JobProvisioningStateFailed ...
	JobProvisioningStateFailed JobProvisioningState = "Failed"
	// JobProvisioningStateProcessing ...
	JobProvisioningStateProcessing JobProvisioningState = "Processing"
	// JobProvisioningStateSucceeded ...
	JobProvisioningStateSucceeded JobProvisioningState = "Succeeded"
	// JobProvisioningStateSuspended ...
	JobProvisioningStateSuspended JobProvisioningState = "Suspended"
)

// PossibleJobProvisioningStateValues returns an array of possible values for the JobProvisioningState const type.
func PossibleJobProvisioningStateValues() []JobProvisioningState {
	return []JobProvisioningState{JobProvisioningStateFailed, JobProvisioningStateProcessing, JobProvisioningStateSucceeded, JobProvisioningStateSuspended}
}

// JobStatus enumerates the values for job status.
type JobStatus string

const (
	// JobStatusActivating ...
	JobStatusActivating JobStatus = "Activating"
	// JobStatusBlocked ...
	JobStatusBlocked JobStatus = "Blocked"
	// JobStatusCompleted ...
	JobStatusCompleted JobStatus = "Completed"
	// JobStatusDisconnected ...
	JobStatusDisconnected JobStatus = "Disconnected"
	// JobStatusFailed ...
	JobStatusFailed JobStatus = "Failed"
	// JobStatusNew ...
	JobStatusNew JobStatus = "New"
	// JobStatusRemoving ...
	JobStatusRemoving JobStatus = "Removing"
	// JobStatusResuming ...
	JobStatusResuming JobStatus = "Resuming"
	// JobStatusRunning ...
	JobStatusRunning JobStatus = "Running"
	// JobStatusStopped ...
	JobStatusStopped JobStatus = "Stopped"
	// JobStatusStopping ...
	JobStatusStopping JobStatus = "Stopping"
	// JobStatusSuspended ...
	JobStatusSuspended JobStatus = "Suspended"
	// JobStatusSuspending ...
	JobStatusSuspending JobStatus = "Suspending"
)

// PossibleJobStatusValues returns an array of possible values for the JobStatus const type.
func PossibleJobStatusValues() []JobStatus {
	return []JobStatus{JobStatusActivating, JobStatusBlocked, JobStatusCompleted, JobStatusDisconnected, JobStatusFailed, JobStatusNew, JobStatusRemoving, JobStatusResuming, JobStatusRunning, JobStatusStopped, JobStatusStopping, JobStatusSuspended, JobStatusSuspending}
}

// JobStreamType enumerates the values for job stream type.
type JobStreamType string

const (
	// JobStreamTypeAny ...
	JobStreamTypeAny JobStreamType = "Any"
	// JobStreamTypeDebug ...
	JobStreamTypeDebug JobStreamType = "Debug"
	// JobStreamTypeError ...
	JobStreamTypeError JobStreamType = "Error"
	// JobStreamTypeOutput ...
	JobStreamTypeOutput JobStreamType = "Output"
	// JobStreamTypeProgress ...
	JobStreamTypeProgress JobStreamType = "Progress"
	// JobStreamTypeVerbose ...
	JobStreamTypeVerbose JobStreamType = "Verbose"
	// JobStreamTypeWarning ...
	JobStreamTypeWarning JobStreamType = "Warning"
)

// PossibleJobStreamTypeValues returns an array of possible values for the JobStreamType const type.
func PossibleJobStreamTypeValues() []JobStreamType {
	return []JobStreamType{JobStreamTypeAny, JobStreamTypeDebug, JobStreamTypeError, JobStreamTypeOutput, JobStreamTypeProgress, JobStreamTypeVerbose, JobStreamTypeWarning}
}

// KeyName enumerates the values for key name.
type KeyName string

const (
	// KeyNamePrimary ...
	KeyNamePrimary KeyName = "Primary"
	// KeyNameSecondary ...
	KeyNameSecondary KeyName = "Secondary"
)

// PossibleKeyNameValues returns an array of possible values for the KeyName const type.
func PossibleKeyNameValues() []KeyName {
	return []KeyName{KeyNamePrimary, KeyNameSecondary}
}

// KeyPermissions enumerates the values for key permissions.
type KeyPermissions string

const (
	// KeyPermissionsFull ...
	KeyPermissionsFull KeyPermissions = "Full"
	// KeyPermissionsRead ...
	KeyPermissionsRead KeyPermissions = "Read"
)

// PossibleKeyPermissionsValues returns an array of possible values for the KeyPermissions const type.
func PossibleKeyPermissionsValues() []KeyPermissions {
	return []KeyPermissions{KeyPermissionsFull, KeyPermissionsRead}
}

// LinuxUpdateClasses enumerates the values for linux update classes.
type LinuxUpdateClasses string

const (
	// LinuxUpdateClassesCritical ...
	LinuxUpdateClassesCritical LinuxUpdateClasses = "Critical"
	// LinuxUpdateClassesOther ...
	LinuxUpdateClassesOther LinuxUpdateClasses = "Other"
	// LinuxUpdateClassesSecurity ...
	LinuxUpdateClassesSecurity LinuxUpdateClasses = "Security"
	// LinuxUpdateClassesUnclassified ...
	LinuxUpdateClassesUnclassified LinuxUpdateClasses = "Unclassified"
)

// PossibleLinuxUpdateClassesValues returns an array of possible values for the LinuxUpdateClasses const type.
func PossibleLinuxUpdateClassesValues() []LinuxUpdateClasses {
	return []LinuxUpdateClasses{LinuxUpdateClassesCritical, LinuxUpdateClassesOther, LinuxUpdateClassesSecurity, LinuxUpdateClassesUnclassified}
}

// ModuleProvisioningState enumerates the values for module provisioning state.
type ModuleProvisioningState string

const (
	// ModuleProvisioningStateActivitiesStored ...
	ModuleProvisioningStateActivitiesStored ModuleProvisioningState = "ActivitiesStored"
	// ModuleProvisioningStateCancelled ...
	ModuleProvisioningStateCancelled ModuleProvisioningState = "Cancelled"
	// ModuleProvisioningStateConnectionTypeImported ...
	ModuleProvisioningStateConnectionTypeImported ModuleProvisioningState = "ConnectionTypeImported"
	// ModuleProvisioningStateContentDownloaded ...
	ModuleProvisioningStateContentDownloaded ModuleProvisioningState = "ContentDownloaded"
	// ModuleProvisioningStateContentRetrieved ...
	ModuleProvisioningStateContentRetrieved ModuleProvisioningState = "ContentRetrieved"
	// ModuleProvisioningStateContentStored ...
	ModuleProvisioningStateContentStored ModuleProvisioningState = "ContentStored"
	// ModuleProvisioningStateContentValidated ...
	ModuleProvisioningStateContentValidated ModuleProvisioningState = "ContentValidated"
	// ModuleProvisioningStateCreated ...
	ModuleProvisioningStateCreated ModuleProvisioningState = "Created"
	// ModuleProvisioningStateCreating ...
	ModuleProvisioningStateCreating ModuleProvisioningState = "Creating"
	// ModuleProvisioningStateFailed ...
	ModuleProvisioningStateFailed ModuleProvisioningState = "Failed"
	// ModuleProvisioningStateModuleDataStored ...
	ModuleProvisioningStateModuleDataStored ModuleProvisioningState = "ModuleDataStored"
	// ModuleProvisioningStateModuleImportRunbookComplete ...
	ModuleProvisioningStateModuleImportRunbookComplete ModuleProvisioningState = "ModuleImportRunbookComplete"
	// ModuleProvisioningStateRunningImportModuleRunbook ...
	ModuleProvisioningStateRunningImportModuleRunbook ModuleProvisioningState = "RunningImportModuleRunbook"
	// ModuleProvisioningStateStartingImportModuleRunbook ...
	ModuleProvisioningStateStartingImportModuleRunbook ModuleProvisioningState = "StartingImportModuleRunbook"
	// ModuleProvisioningStateSucceeded ...
	ModuleProvisioningStateSucceeded ModuleProvisioningState = "Succeeded"
	// ModuleProvisioningStateUpdating ...
	ModuleProvisioningStateUpdating ModuleProvisioningState = "Updating"
)

// PossibleModuleProvisioningStateValues returns an array of possible values for the ModuleProvisioningState const type.
func PossibleModuleProvisioningStateValues() []ModuleProvisioningState {
	return []ModuleProvisioningState{ModuleProvisioningStateActivitiesStored, ModuleProvisioningStateCancelled, ModuleProvisioningStateConnectionTypeImported, ModuleProvisioningStateContentDownloaded, ModuleProvisioningStateContentRetrieved, ModuleProvisioningStateContentStored, ModuleProvisioningStateContentValidated, ModuleProvisioningStateCreated, ModuleProvisioningStateCreating, ModuleProvisioningStateFailed, ModuleProvisioningStateModuleDataStored, ModuleProvisioningStateModuleImportRunbookComplete, ModuleProvisioningStateRunningImportModuleRunbook, ModuleProvisioningStateStartingImportModuleRunbook, ModuleProvisioningStateSucceeded, ModuleProvisioningStateUpdating}
}

// OperatingSystemType enumerates the values for operating system type.
type OperatingSystemType string

const (
	// OperatingSystemTypeLinux ...
	OperatingSystemTypeLinux OperatingSystemType = "Linux"
	// OperatingSystemTypeWindows ...
	OperatingSystemTypeWindows OperatingSystemType = "Windows"
)

// PossibleOperatingSystemTypeValues returns an array of possible values for the OperatingSystemType const type.
func PossibleOperatingSystemTypeValues() []OperatingSystemType {
	return []OperatingSystemType{OperatingSystemTypeLinux, OperatingSystemTypeWindows}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateCompleted ...
	ProvisioningStateCompleted ProvisioningState = "Completed"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateRunning ...
	ProvisioningStateRunning ProvisioningState = "Running"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateCompleted, ProvisioningStateFailed, ProvisioningStateRunning}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// ResourceIdentityTypeNone ...
	ResourceIdentityTypeNone ResourceIdentityType = "None"
	// ResourceIdentityTypeSystemAssigned ...
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	// ResourceIdentityTypeSystemAssignedUserAssigned ...
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	// ResourceIdentityTypeUserAssigned ...
	ResourceIdentityTypeUserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{ResourceIdentityTypeNone, ResourceIdentityTypeSystemAssigned, ResourceIdentityTypeSystemAssignedUserAssigned, ResourceIdentityTypeUserAssigned}
}

// RunbookProvisioningState enumerates the values for runbook provisioning state.
type RunbookProvisioningState string

const (
	// RunbookProvisioningStateSucceeded ...
	RunbookProvisioningStateSucceeded RunbookProvisioningState = "Succeeded"
)

// PossibleRunbookProvisioningStateValues returns an array of possible values for the RunbookProvisioningState const type.
func PossibleRunbookProvisioningStateValues() []RunbookProvisioningState {
	return []RunbookProvisioningState{RunbookProvisioningStateSucceeded}
}

// RunbookState enumerates the values for runbook state.
type RunbookState string

const (
	// RunbookStateEdit ...
	RunbookStateEdit RunbookState = "Edit"
	// RunbookStateNew ...
	RunbookStateNew RunbookState = "New"
	// RunbookStatePublished ...
	RunbookStatePublished RunbookState = "Published"
)

// PossibleRunbookStateValues returns an array of possible values for the RunbookState const type.
func PossibleRunbookStateValues() []RunbookState {
	return []RunbookState{RunbookStateEdit, RunbookStateNew, RunbookStatePublished}
}

// RunbookTypeEnum enumerates the values for runbook type enum.
type RunbookTypeEnum string

const (
	// RunbookTypeEnumGraph ...
	RunbookTypeEnumGraph RunbookTypeEnum = "Graph"
	// RunbookTypeEnumGraphPowerShell ...
	RunbookTypeEnumGraphPowerShell RunbookTypeEnum = "GraphPowerShell"
	// RunbookTypeEnumGraphPowerShellWorkflow ...
	RunbookTypeEnumGraphPowerShellWorkflow RunbookTypeEnum = "GraphPowerShellWorkflow"
	// RunbookTypeEnumPowerShell ...
	RunbookTypeEnumPowerShell RunbookTypeEnum = "PowerShell"
	// RunbookTypeEnumPowerShellWorkflow ...
	RunbookTypeEnumPowerShellWorkflow RunbookTypeEnum = "PowerShellWorkflow"
	// RunbookTypeEnumPython2 ...
	RunbookTypeEnumPython2 RunbookTypeEnum = "Python2"
	// RunbookTypeEnumPython3 ...
	RunbookTypeEnumPython3 RunbookTypeEnum = "Python3"
	// RunbookTypeEnumScript ...
	RunbookTypeEnumScript RunbookTypeEnum = "Script"
)

// PossibleRunbookTypeEnumValues returns an array of possible values for the RunbookTypeEnum const type.
func PossibleRunbookTypeEnumValues() []RunbookTypeEnum {
	return []RunbookTypeEnum{RunbookTypeEnumGraph, RunbookTypeEnumGraphPowerShell, RunbookTypeEnumGraphPowerShellWorkflow, RunbookTypeEnumPowerShell, RunbookTypeEnumPowerShellWorkflow, RunbookTypeEnumPython2, RunbookTypeEnumPython3, RunbookTypeEnumScript}
}

// ScheduleDay enumerates the values for schedule day.
type ScheduleDay string

const (
	// ScheduleDayFriday ...
	ScheduleDayFriday ScheduleDay = "Friday"
	// ScheduleDayMonday ...
	ScheduleDayMonday ScheduleDay = "Monday"
	// ScheduleDaySaturday ...
	ScheduleDaySaturday ScheduleDay = "Saturday"
	// ScheduleDaySunday ...
	ScheduleDaySunday ScheduleDay = "Sunday"
	// ScheduleDayThursday ...
	ScheduleDayThursday ScheduleDay = "Thursday"
	// ScheduleDayTuesday ...
	ScheduleDayTuesday ScheduleDay = "Tuesday"
	// ScheduleDayWednesday ...
	ScheduleDayWednesday ScheduleDay = "Wednesday"
)

// PossibleScheduleDayValues returns an array of possible values for the ScheduleDay const type.
func PossibleScheduleDayValues() []ScheduleDay {
	return []ScheduleDay{ScheduleDayFriday, ScheduleDayMonday, ScheduleDaySaturday, ScheduleDaySunday, ScheduleDayThursday, ScheduleDayTuesday, ScheduleDayWednesday}
}

// ScheduleFrequency enumerates the values for schedule frequency.
type ScheduleFrequency string

const (
	// ScheduleFrequencyDay ...
	ScheduleFrequencyDay ScheduleFrequency = "Day"
	// ScheduleFrequencyHour ...
	ScheduleFrequencyHour ScheduleFrequency = "Hour"
	// ScheduleFrequencyMinute The minimum allowed interval for Minute schedules is 15 minutes.
	ScheduleFrequencyMinute ScheduleFrequency = "Minute"
	// ScheduleFrequencyMonth ...
	ScheduleFrequencyMonth ScheduleFrequency = "Month"
	// ScheduleFrequencyOneTime ...
	ScheduleFrequencyOneTime ScheduleFrequency = "OneTime"
	// ScheduleFrequencyWeek ...
	ScheduleFrequencyWeek ScheduleFrequency = "Week"
)

// PossibleScheduleFrequencyValues returns an array of possible values for the ScheduleFrequency const type.
func PossibleScheduleFrequencyValues() []ScheduleFrequency {
	return []ScheduleFrequency{ScheduleFrequencyDay, ScheduleFrequencyHour, ScheduleFrequencyMinute, ScheduleFrequencyMonth, ScheduleFrequencyOneTime, ScheduleFrequencyWeek}
}

// SkuNameEnum enumerates the values for sku name enum.
type SkuNameEnum string

const (
	// SkuNameEnumBasic ...
	SkuNameEnumBasic SkuNameEnum = "Basic"
	// SkuNameEnumFree ...
	SkuNameEnumFree SkuNameEnum = "Free"
)

// PossibleSkuNameEnumValues returns an array of possible values for the SkuNameEnum const type.
func PossibleSkuNameEnumValues() []SkuNameEnum {
	return []SkuNameEnum{SkuNameEnumBasic, SkuNameEnumFree}
}

// SourceType enumerates the values for source type.
type SourceType string

const (
	// SourceTypeGitHub ...
	SourceTypeGitHub SourceType = "GitHub"
	// SourceTypeVsoGit ...
	SourceTypeVsoGit SourceType = "VsoGit"
	// SourceTypeVsoTfvc ...
	SourceTypeVsoTfvc SourceType = "VsoTfvc"
)

// PossibleSourceTypeValues returns an array of possible values for the SourceType const type.
func PossibleSourceTypeValues() []SourceType {
	return []SourceType{SourceTypeGitHub, SourceTypeVsoGit, SourceTypeVsoTfvc}
}

// StreamType enumerates the values for stream type.
type StreamType string

const (
	// StreamTypeError ...
	StreamTypeError StreamType = "Error"
	// StreamTypeOutput ...
	StreamTypeOutput StreamType = "Output"
)

// PossibleStreamTypeValues returns an array of possible values for the StreamType const type.
func PossibleStreamTypeValues() []StreamType {
	return []StreamType{StreamTypeError, StreamTypeOutput}
}

// SyncType enumerates the values for sync type.
type SyncType string

const (
	// SyncTypeFullSync ...
	SyncTypeFullSync SyncType = "FullSync"
	// SyncTypePartialSync ...
	SyncTypePartialSync SyncType = "PartialSync"
)

// PossibleSyncTypeValues returns an array of possible values for the SyncType const type.
func PossibleSyncTypeValues() []SyncType {
	return []SyncType{SyncTypeFullSync, SyncTypePartialSync}
}

// TagOperators enumerates the values for tag operators.
type TagOperators string

const (
	// TagOperatorsAll ...
	TagOperatorsAll TagOperators = "All"
	// TagOperatorsAny ...
	TagOperatorsAny TagOperators = "Any"
)

// PossibleTagOperatorsValues returns an array of possible values for the TagOperators const type.
func PossibleTagOperatorsValues() []TagOperators {
	return []TagOperators{TagOperatorsAll, TagOperatorsAny}
}

// TokenType enumerates the values for token type.
type TokenType string

const (
	// TokenTypeOauth ...
	TokenTypeOauth TokenType = "Oauth"
	// TokenTypePersonalAccessToken ...
	TokenTypePersonalAccessToken TokenType = "PersonalAccessToken"
)

// PossibleTokenTypeValues returns an array of possible values for the TokenType const type.
func PossibleTokenTypeValues() []TokenType {
	return []TokenType{TokenTypeOauth, TokenTypePersonalAccessToken}
}

// WindowsUpdateClasses enumerates the values for windows update classes.
type WindowsUpdateClasses string

const (
	// WindowsUpdateClassesCritical ...
	WindowsUpdateClassesCritical WindowsUpdateClasses = "Critical"
	// WindowsUpdateClassesDefinition ...
	WindowsUpdateClassesDefinition WindowsUpdateClasses = "Definition"
	// WindowsUpdateClassesFeaturePack ...
	WindowsUpdateClassesFeaturePack WindowsUpdateClasses = "FeaturePack"
	// WindowsUpdateClassesSecurity ...
	WindowsUpdateClassesSecurity WindowsUpdateClasses = "Security"
	// WindowsUpdateClassesServicePack ...
	WindowsUpdateClassesServicePack WindowsUpdateClasses = "ServicePack"
	// WindowsUpdateClassesTools ...
	WindowsUpdateClassesTools WindowsUpdateClasses = "Tools"
	// WindowsUpdateClassesUnclassified ...
	WindowsUpdateClassesUnclassified WindowsUpdateClasses = "Unclassified"
	// WindowsUpdateClassesUpdateRollup ...
	WindowsUpdateClassesUpdateRollup WindowsUpdateClasses = "UpdateRollup"
	// WindowsUpdateClassesUpdates ...
	WindowsUpdateClassesUpdates WindowsUpdateClasses = "Updates"
)

// PossibleWindowsUpdateClassesValues returns an array of possible values for the WindowsUpdateClasses const type.
func PossibleWindowsUpdateClassesValues() []WindowsUpdateClasses {
	return []WindowsUpdateClasses{WindowsUpdateClassesCritical, WindowsUpdateClassesDefinition, WindowsUpdateClassesFeaturePack, WindowsUpdateClassesSecurity, WindowsUpdateClassesServicePack, WindowsUpdateClassesTools, WindowsUpdateClassesUnclassified, WindowsUpdateClassesUpdateRollup, WindowsUpdateClassesUpdates}
}
