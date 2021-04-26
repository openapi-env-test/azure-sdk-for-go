package dataprotectionapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/dataprotection/mgmt/2021-02-01-preview/dataprotection"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckFeatureSupport(ctx context.Context, location string, parameters dataprotection.BasicFeatureValidationRequestBase) (result dataprotection.FeatureValidationResponseBaseModel, err error)
	GetOperationResultPatch(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result dataprotection.BackupVaultResource, err error)
	GetOperationStatus(ctx context.Context, location string, operationID string) (result dataprotection.OperationResource, err error)
}

var _ BaseClientAPI = (*dataprotection.BaseClient)(nil)

// BackupVaultsClientAPI contains the set of methods on the BackupVaultsClient type.
type BackupVaultsClientAPI interface {
	CheckNameAvailability(ctx context.Context, resourceGroupName string, location string, parameters dataprotection.CheckNameAvailabilityRequest) (result dataprotection.CheckNameAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, vaultName string, resourceGroupName string, parameters dataprotection.BackupVaultResource) (result dataprotection.BackupVaultsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, vaultName string, resourceGroupName string) (result autorest.Response, err error)
	Get(ctx context.Context, vaultName string, resourceGroupName string) (result dataprotection.BackupVaultResource, err error)
	GetResourcesInResourceGroup(ctx context.Context, resourceGroupName string) (result dataprotection.BackupVaultResourceListPage, err error)
	GetResourcesInResourceGroupComplete(ctx context.Context, resourceGroupName string) (result dataprotection.BackupVaultResourceListIterator, err error)
	GetResourcesInSubscription(ctx context.Context) (result dataprotection.BackupVaultResourceListPage, err error)
	GetResourcesInSubscriptionComplete(ctx context.Context) (result dataprotection.BackupVaultResourceListIterator, err error)
	Patch(ctx context.Context, vaultName string, resourceGroupName string, parameters dataprotection.PatchResourceRequestInput) (result dataprotection.BackupVaultsPatchFuture, err error)
}

var _ BackupVaultsClientAPI = (*dataprotection.BackupVaultsClient)(nil)

// OperationResultClientAPI contains the set of methods on the OperationResultClient type.
type OperationResultClientAPI interface {
	Get(ctx context.Context, operationID string, location string) (result dataprotection.OperationJobExtendedInfo, err error)
}

var _ OperationResultClientAPI = (*dataprotection.OperationResultClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result dataprotection.ClientDiscoveryResponsePage, err error)
	ListComplete(ctx context.Context) (result dataprotection.ClientDiscoveryResponseIterator, err error)
}

var _ OperationsClientAPI = (*dataprotection.OperationsClient)(nil)

// BackupPoliciesClientAPI contains the set of methods on the BackupPoliciesClient type.
type BackupPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, vaultName string, resourceGroupName string, backupPolicyName string, parameters dataprotection.BaseBackupPolicyResource) (result dataprotection.BaseBackupPolicyResource, err error)
	Delete(ctx context.Context, vaultName string, resourceGroupName string, backupPolicyName string) (result autorest.Response, err error)
	Get(ctx context.Context, vaultName string, resourceGroupName string, backupPolicyName string) (result dataprotection.BaseBackupPolicyResource, err error)
	List(ctx context.Context, vaultName string, resourceGroupName string) (result dataprotection.BaseBackupPolicyResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string) (result dataprotection.BaseBackupPolicyResourceListIterator, err error)
}

var _ BackupPoliciesClientAPI = (*dataprotection.BackupPoliciesClient)(nil)

// BackupInstancesClientAPI contains the set of methods on the BackupInstancesClient type.
type BackupInstancesClientAPI interface {
	AdhocBackup(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, parameters dataprotection.TriggerBackupRequest) (result dataprotection.BackupInstancesAdhocBackupFuture, err error)
	CreateOrUpdate(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, parameters dataprotection.BackupInstanceResource) (result dataprotection.BackupInstancesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string) (result dataprotection.BackupInstancesDeleteFuture, err error)
	Get(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string) (result dataprotection.BackupInstanceResource, err error)
	List(ctx context.Context, vaultName string, resourceGroupName string) (result dataprotection.BackupInstanceResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string) (result dataprotection.BackupInstanceResourceListIterator, err error)
	TriggerRehydrate(ctx context.Context, resourceGroupName string, vaultName string, parameters dataprotection.AzureBackupRehydrationRequest, backupInstanceName string) (result dataprotection.BackupInstancesTriggerRehydrateFuture, err error)
	TriggerRestore(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, parameters dataprotection.BasicAzureBackupRestoreRequest) (result dataprotection.BackupInstancesTriggerRestoreFuture, err error)
	ValidateForBackup(ctx context.Context, vaultName string, resourceGroupName string, parameters dataprotection.ValidateForBackupRequest) (result dataprotection.BackupInstancesValidateForBackupFuture, err error)
	ValidateRestore(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, parameters dataprotection.ValidateRestoreRequestObject) (result dataprotection.BackupInstancesValidateRestoreFuture, err error)
}

var _ BackupInstancesClientAPI = (*dataprotection.BackupInstancesClient)(nil)

// RecoveryPointsClientAPI contains the set of methods on the RecoveryPointsClient type.
type RecoveryPointsClientAPI interface {
	GetList(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, filter string, skipToken string) (result dataprotection.AzureBackupRecoveryPointResourceListPage, err error)
	GetListComplete(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, filter string, skipToken string) (result dataprotection.AzureBackupRecoveryPointResourceListIterator, err error)
}

var _ RecoveryPointsClientAPI = (*dataprotection.RecoveryPointsClient)(nil)

// RecoveryPointClientAPI contains the set of methods on the RecoveryPointClient type.
type RecoveryPointClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, recoveryPointID string) (result dataprotection.AzureBackupRecoveryPointResource, err error)
}

var _ RecoveryPointClientAPI = (*dataprotection.RecoveryPointClient)(nil)

// JobsClientAPI contains the set of methods on the JobsClient type.
type JobsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, vaultName string) (result dataprotection.AzureBackupJobResourceListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, vaultName string) (result dataprotection.AzureBackupJobResourceListIterator, err error)
}

var _ JobsClientAPI = (*dataprotection.JobsClient)(nil)

// FindRestorableTimeRangesClientAPI contains the set of methods on the FindRestorableTimeRangesClient type.
type FindRestorableTimeRangesClientAPI interface {
	Post(ctx context.Context, vaultName string, resourceGroupName string, backupInstances string, parameters dataprotection.AzureBackupFindRestorableTimeRangesRequest) (result dataprotection.AzureBackupFindRestorableTimeRangesResponseResource, err error)
}

var _ FindRestorableTimeRangesClientAPI = (*dataprotection.FindRestorableTimeRangesClient)(nil)

// JobClientAPI contains the set of methods on the JobClient type.
type JobClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, vaultName string, jobID string) (result dataprotection.AzureBackupJobResource, err error)
}

var _ JobClientAPI = (*dataprotection.JobClient)(nil)

// ExportJobsClientAPI contains the set of methods on the ExportJobsClient type.
type ExportJobsClientAPI interface {
	Trigger(ctx context.Context, resourceGroupName string, vaultName string) (result dataprotection.ExportJobsTriggerFuture, err error)
}

var _ ExportJobsClientAPI = (*dataprotection.ExportJobsClient)(nil)

// ExportJobsOperationResultClientAPI contains the set of methods on the ExportJobsOperationResultClient type.
type ExportJobsOperationResultClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, vaultName string, operationID string) (result dataprotection.ExportJobsResult, err error)
}

var _ ExportJobsOperationResultClientAPI = (*dataprotection.ExportJobsOperationResultClient)(nil)
