package testbase

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Action enumerates the values for action.
type Action string

const (
	// Close ...
	Close Action = "Close"
	// Custom ...
	Custom Action = "Custom"
	// Install ...
	Install Action = "Install"
	// Launch ...
	Launch Action = "Launch"
	// Uninstall ...
	Uninstall Action = "Uninstall"
)

// PossibleActionValues returns an array of possible values for the Action const type.
func PossibleActionValues() []Action {
	return []Action{Close, Custom, Install, Launch, Uninstall}
}

// AnalysisResultName enumerates the values for analysis result name.
type AnalysisResultName string

const (
	// CPURegression ...
	CPURegression AnalysisResultName = "cpuRegression"
	// CPUUtilization ...
	CPUUtilization AnalysisResultName = "cpuUtilization"
	// MemoryRegression ...
	MemoryRegression AnalysisResultName = "memoryRegression"
	// MemoryUtilization ...
	MemoryUtilization AnalysisResultName = "memoryUtilization"
	// Reliability ...
	Reliability AnalysisResultName = "reliability"
	// ScriptExecution ...
	ScriptExecution AnalysisResultName = "scriptExecution"
)

// PossibleAnalysisResultNameValues returns an array of possible values for the AnalysisResultName const type.
func PossibleAnalysisResultNameValues() []AnalysisResultName {
	return []AnalysisResultName{CPURegression, CPUUtilization, MemoryRegression, MemoryUtilization, Reliability, ScriptExecution}
}

// AnalysisResultType enumerates the values for analysis result type.
type AnalysisResultType string

const (
	// AnalysisResultTypeCPURegression ...
	AnalysisResultTypeCPURegression AnalysisResultType = "CPURegression"
	// AnalysisResultTypeCPUUtilization ...
	AnalysisResultTypeCPUUtilization AnalysisResultType = "CPUUtilization"
	// AnalysisResultTypeMemoryRegression ...
	AnalysisResultTypeMemoryRegression AnalysisResultType = "MemoryRegression"
	// AnalysisResultTypeMemoryUtilization ...
	AnalysisResultTypeMemoryUtilization AnalysisResultType = "MemoryUtilization"
	// AnalysisResultTypeReliability ...
	AnalysisResultTypeReliability AnalysisResultType = "Reliability"
	// AnalysisResultTypeScriptExecution ...
	AnalysisResultTypeScriptExecution AnalysisResultType = "ScriptExecution"
)

// PossibleAnalysisResultTypeValues returns an array of possible values for the AnalysisResultType const type.
func PossibleAnalysisResultTypeValues() []AnalysisResultType {
	return []AnalysisResultType{AnalysisResultTypeCPURegression, AnalysisResultTypeCPUUtilization, AnalysisResultTypeMemoryRegression, AnalysisResultTypeMemoryUtilization, AnalysisResultTypeReliability, AnalysisResultTypeScriptExecution}
}

// AnalysisResultTypeBasicAnalysisResultSingletonResourceProperties enumerates the values for analysis result
// type basic analysis result singleton resource properties.
type AnalysisResultTypeBasicAnalysisResultSingletonResourceProperties string

const (
	// AnalysisResultTypeAnalysisResultSingletonResourceProperties ...
	AnalysisResultTypeAnalysisResultSingletonResourceProperties AnalysisResultTypeBasicAnalysisResultSingletonResourceProperties = "AnalysisResultSingletonResourceProperties"
	// AnalysisResultTypeCPURegression1 ...
	AnalysisResultTypeCPURegression1 AnalysisResultTypeBasicAnalysisResultSingletonResourceProperties = "CPURegression"
	// AnalysisResultTypeCPUUtilization1 ...
	AnalysisResultTypeCPUUtilization1 AnalysisResultTypeBasicAnalysisResultSingletonResourceProperties = "CPUUtilization"
	// AnalysisResultTypeMemoryRegression1 ...
	AnalysisResultTypeMemoryRegression1 AnalysisResultTypeBasicAnalysisResultSingletonResourceProperties = "MemoryRegression"
	// AnalysisResultTypeMemoryUtilization1 ...
	AnalysisResultTypeMemoryUtilization1 AnalysisResultTypeBasicAnalysisResultSingletonResourceProperties = "MemoryUtilization"
	// AnalysisResultTypeReliability1 ...
	AnalysisResultTypeReliability1 AnalysisResultTypeBasicAnalysisResultSingletonResourceProperties = "Reliability"
	// AnalysisResultTypeScriptExecution1 ...
	AnalysisResultTypeScriptExecution1 AnalysisResultTypeBasicAnalysisResultSingletonResourceProperties = "ScriptExecution"
)

// PossibleAnalysisResultTypeBasicAnalysisResultSingletonResourcePropertiesValues returns an array of possible values for the AnalysisResultTypeBasicAnalysisResultSingletonResourceProperties const type.
func PossibleAnalysisResultTypeBasicAnalysisResultSingletonResourcePropertiesValues() []AnalysisResultTypeBasicAnalysisResultSingletonResourceProperties {
	return []AnalysisResultTypeBasicAnalysisResultSingletonResourceProperties{AnalysisResultTypeAnalysisResultSingletonResourceProperties, AnalysisResultTypeCPURegression1, AnalysisResultTypeCPUUtilization1, AnalysisResultTypeMemoryRegression1, AnalysisResultTypeMemoryUtilization1, AnalysisResultTypeReliability1, AnalysisResultTypeScriptExecution1}
}

// AnalysisStatus enumerates the values for analysis status.
type AnalysisStatus string

const (
	// Available ...
	Available AnalysisStatus = "Available"
	// Completed ...
	Completed AnalysisStatus = "Completed"
	// Failed ...
	Failed AnalysisStatus = "Failed"
	// InProgress ...
	InProgress AnalysisStatus = "InProgress"
	// None ...
	None AnalysisStatus = "None"
	// NotAvailable ...
	NotAvailable AnalysisStatus = "NotAvailable"
	// Succeeded ...
	Succeeded AnalysisStatus = "Succeeded"
)

// PossibleAnalysisStatusValues returns an array of possible values for the AnalysisStatus const type.
func PossibleAnalysisStatusValues() []AnalysisStatus {
	return []AnalysisStatus{Available, Completed, Failed, InProgress, None, NotAvailable, Succeeded}
}

// ContentType enumerates the values for content type.
type ContentType string

const (
	// File ...
	File ContentType = "File"
	// Inline ...
	Inline ContentType = "Inline"
	// Path ...
	Path ContentType = "Path"
)

// PossibleContentTypeValues returns an array of possible values for the ContentType const type.
func PossibleContentTypeValues() []ContentType {
	return []ContentType{File, Inline, Path}
}

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

// ExecutionStatus enumerates the values for execution status.
type ExecutionStatus string

const (
	// ExecutionStatusCompleted ...
	ExecutionStatusCompleted ExecutionStatus = "Completed"
	// ExecutionStatusFailed ...
	ExecutionStatusFailed ExecutionStatus = "Failed"
	// ExecutionStatusIncomplete ...
	ExecutionStatusIncomplete ExecutionStatus = "Incomplete"
	// ExecutionStatusInProgress ...
	ExecutionStatusInProgress ExecutionStatus = "InProgress"
	// ExecutionStatusNone ...
	ExecutionStatusNone ExecutionStatus = "None"
	// ExecutionStatusNotExecuted ...
	ExecutionStatusNotExecuted ExecutionStatus = "NotExecuted"
	// ExecutionStatusProcessing ...
	ExecutionStatusProcessing ExecutionStatus = "Processing"
	// ExecutionStatusSucceeded ...
	ExecutionStatusSucceeded ExecutionStatus = "Succeeded"
)

// PossibleExecutionStatusValues returns an array of possible values for the ExecutionStatus const type.
func PossibleExecutionStatusValues() []ExecutionStatus {
	return []ExecutionStatus{ExecutionStatusCompleted, ExecutionStatusFailed, ExecutionStatusIncomplete, ExecutionStatusInProgress, ExecutionStatusNone, ExecutionStatusNotExecuted, ExecutionStatusProcessing, ExecutionStatusSucceeded}
}

// Grade enumerates the values for grade.
type Grade string

const (
	// GradeFail ...
	GradeFail Grade = "Fail"
	// GradeNone ...
	GradeNone Grade = "None"
	// GradeNotAvailable ...
	GradeNotAvailable Grade = "NotAvailable"
	// GradePass ...
	GradePass Grade = "Pass"
)

// PossibleGradeValues returns an array of possible values for the Grade const type.
func PossibleGradeValues() []Grade {
	return []Grade{GradeFail, GradeNone, GradeNotAvailable, GradePass}
}

// OsUpdateType enumerates the values for os update type.
type OsUpdateType string

const (
	// FeatureUpdate ...
	FeatureUpdate OsUpdateType = "FeatureUpdate"
	// SecurityUpdate ...
	SecurityUpdate OsUpdateType = "SecurityUpdate"
)

// PossibleOsUpdateTypeValues returns an array of possible values for the OsUpdateType const type.
func PossibleOsUpdateTypeValues() []OsUpdateType {
	return []OsUpdateType{FeatureUpdate, SecurityUpdate}
}

// PackageStatus enumerates the values for package status.
type PackageStatus string

const (
	// Deleted ...
	Deleted PackageStatus = "Deleted"
	// Error ...
	Error PackageStatus = "Error"
	// PreValidationCheckPass ...
	PreValidationCheckPass PackageStatus = "PreValidationCheckPass"
	// Ready ...
	Ready PackageStatus = "Ready"
	// Registered ...
	Registered PackageStatus = "Registered"
	// Unknown ...
	Unknown PackageStatus = "Unknown"
	// ValidatingPackage ...
	ValidatingPackage PackageStatus = "ValidatingPackage"
	// ValidationLongerThanUsual ...
	ValidationLongerThanUsual PackageStatus = "ValidationLongerThanUsual"
	// VerifyingPackage ...
	VerifyingPackage PackageStatus = "VerifyingPackage"
)

// PossiblePackageStatusValues returns an array of possible values for the PackageStatus const type.
func PossiblePackageStatusValues() []PackageStatus {
	return []PackageStatus{Deleted, Error, PreValidationCheckPass, Ready, Registered, Unknown, ValidatingPackage, ValidationLongerThanUsual, VerifyingPackage}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateCancelled ...
	ProvisioningStateCancelled ProvisioningState = "Cancelled"
	// ProvisioningStateCreating ...
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating ...
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateCancelled, ProvisioningStateCreating, ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStateSucceeded, ProvisioningStateUpdating}
}

// Reason enumerates the values for reason.
type Reason string

const (
	// AlreadyExists ...
	AlreadyExists Reason = "AlreadyExists"
	// Invalid ...
	Invalid Reason = "Invalid"
)

// PossibleReasonValues returns an array of possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{AlreadyExists, Invalid}
}

// TestType enumerates the values for test type.
type TestType string

const (
	// FunctionalTest ...
	FunctionalTest TestType = "FunctionalTest"
	// OutOfBoxTest ...
	OutOfBoxTest TestType = "OutOfBoxTest"
)

// PossibleTestTypeValues returns an array of possible values for the TestType const type.
func PossibleTestTypeValues() []TestType {
	return []TestType{FunctionalTest, OutOfBoxTest}
}

// Tier enumerates the values for tier.
type Tier string

const (
	// Basic ...
	Basic Tier = "Basic"
	// Standard ...
	Standard Tier = "Standard"
)

// PossibleTierValues returns an array of possible values for the Tier const type.
func PossibleTierValues() []Tier {
	return []Tier{Basic, Standard}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeFeatureUpdate ...
	TypeFeatureUpdate Type = "FeatureUpdate"
	// TypeSecurityUpdate ...
	TypeSecurityUpdate Type = "SecurityUpdate"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeFeatureUpdate, TypeSecurityUpdate}
}

// ValidationRunStatus enumerates the values for validation run status.
type ValidationRunStatus string

const (
	// ValidationRunStatusFailed ...
	ValidationRunStatusFailed ValidationRunStatus = "Failed"
	// ValidationRunStatusPassed ...
	ValidationRunStatusPassed ValidationRunStatus = "Passed"
	// ValidationRunStatusPending ...
	ValidationRunStatusPending ValidationRunStatus = "Pending"
	// ValidationRunStatusUnknown ...
	ValidationRunStatusUnknown ValidationRunStatus = "Unknown"
)

// PossibleValidationRunStatusValues returns an array of possible values for the ValidationRunStatus const type.
func PossibleValidationRunStatusValues() []ValidationRunStatus {
	return []ValidationRunStatus{ValidationRunStatusFailed, ValidationRunStatusPassed, ValidationRunStatusPending, ValidationRunStatusUnknown}
}
