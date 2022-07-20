package backupapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/recoveryservices/mgmt/2021-01-01/backup"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	BMSPrepareDataMove(ctx context.Context, vaultName string, resourceGroupName string, parameters backup.PrepareDataMoveRequest) (result backup.BMSPrepareDataMoveFuture, err error)
	BMSTriggerDataMove(ctx context.Context, vaultName string, resourceGroupName string, parameters backup.TriggerDataMoveRequest) (result backup.BMSTriggerDataMoveFuture, err error)
	GetOperationStatus(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result backup.OperationStatus, err error)
	MoveRecoveryPoint(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters backup.MoveRPAcrossTiersRequest) (result backup.MoveRecoveryPointFuture, err error)
}

var _ BaseClientAPI = (*backup.BaseClient)(nil)

// ResourceVaultConfigsClientAPI contains the set of methods on the ResourceVaultConfigsClient type.
type ResourceVaultConfigsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string) (result backup.ResourceVaultConfigResource, err error)
	Put(ctx context.Context, vaultName string, resourceGroupName string, parameters backup.ResourceVaultConfigResource) (result backup.ResourceVaultConfigResource, err error)
	Update(ctx context.Context, vaultName string, resourceGroupName string, parameters backup.ResourceVaultConfigResource) (result backup.ResourceVaultConfigResource, err error)
}

var _ ResourceVaultConfigsClientAPI = (*backup.ResourceVaultConfigsClient)(nil)

// ResourceEncryptionConfigsClientAPI contains the set of methods on the ResourceEncryptionConfigsClient type.
type ResourceEncryptionConfigsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string) (result backup.ResourceEncryptionConfigResource, err error)
	Update(ctx context.Context, vaultName string, resourceGroupName string, parameters backup.ResourceEncryptionConfigResource) (result autorest.Response, err error)
}

var _ ResourceEncryptionConfigsClientAPI = (*backup.ResourceEncryptionConfigsClient)(nil)

// PrivateEndpointConnectionClientAPI contains the set of methods on the PrivateEndpointConnectionClient type.
type PrivateEndpointConnectionClientAPI interface {
	Delete(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string) (result backup.PrivateEndpointConnectionDeleteFuture, err error)
	Get(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string) (result backup.PrivateEndpointConnectionResource, err error)
	Put(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string, parameters backup.PrivateEndpointConnectionResource) (result backup.PrivateEndpointConnectionPutFuture, err error)
}

var _ PrivateEndpointConnectionClientAPI = (*backup.PrivateEndpointConnectionClient)(nil)

// PrivateEndpointClientAPI contains the set of methods on the PrivateEndpointClient type.
type PrivateEndpointClientAPI interface {
	GetOperationStatus(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string, operationID string) (result backup.OperationStatus, err error)
}

var _ PrivateEndpointClientAPI = (*backup.PrivateEndpointClient)(nil)

// BMSPrepareDataMoveOperationResultClientAPI contains the set of methods on the BMSPrepareDataMoveOperationResultClient type.
type BMSPrepareDataMoveOperationResultClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result backup.VaultStorageConfigOperationResultResponseModel, err error)
}

var _ BMSPrepareDataMoveOperationResultClientAPI = (*backup.BMSPrepareDataMoveOperationResultClient)(nil)

// ProtectedItemsClientAPI contains the set of methods on the ProtectedItemsClient type.
type ProtectedItemsClientAPI interface {
	CreateOrUpdate(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, parameters backup.ProtectedItemResource) (result backup.ProtectedItemResource, err error)
	Delete(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string) (result autorest.Response, err error)
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, filter string) (result backup.ProtectedItemResource, err error)
}

var _ ProtectedItemsClientAPI = (*backup.ProtectedItemsClient)(nil)

// ProtectedItemOperationResultsClientAPI contains the set of methods on the ProtectedItemOperationResultsClient type.
type ProtectedItemOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, operationID string) (result backup.ProtectedItemResource, err error)
}

var _ ProtectedItemOperationResultsClientAPI = (*backup.ProtectedItemOperationResultsClient)(nil)

// RecoveryPointsClientAPI contains the set of methods on the RecoveryPointsClient type.
type RecoveryPointsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string) (result backup.RecoveryPointResource, err error)
	List(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, filter string) (result backup.RecoveryPointResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, filter string) (result backup.RecoveryPointResourceListIterator, err error)
}

var _ RecoveryPointsClientAPI = (*backup.RecoveryPointsClient)(nil)

// RestoresClientAPI contains the set of methods on the RestoresClient type.
type RestoresClientAPI interface {
	Trigger(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters backup.RestoreRequestResource) (result backup.RestoresTriggerFuture, err error)
}

var _ RestoresClientAPI = (*backup.RestoresClient)(nil)

// PoliciesClientAPI contains the set of methods on the PoliciesClient type.
type PoliciesClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string) (result backup.ProtectionPolicyResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string) (result backup.ProtectionPolicyResourceListIterator, err error)
}

var _ PoliciesClientAPI = (*backup.PoliciesClient)(nil)

// ProtectionPoliciesClientAPI contains the set of methods on the ProtectionPoliciesClient type.
type ProtectionPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, vaultName string, resourceGroupName string, policyName string, parameters backup.ProtectionPolicyResource) (result backup.ProtectionPolicyResource, err error)
	Delete(ctx context.Context, vaultName string, resourceGroupName string, policyName string) (result backup.ProtectionPoliciesDeleteFuture, err error)
	Get(ctx context.Context, vaultName string, resourceGroupName string, policyName string) (result backup.ProtectionPolicyResource, err error)
}

var _ ProtectionPoliciesClientAPI = (*backup.ProtectionPoliciesClient)(nil)

// ProtectionPolicyOperationResultsClientAPI contains the set of methods on the ProtectionPolicyOperationResultsClient type.
type ProtectionPolicyOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string) (result backup.ProtectionPolicyResource, err error)
}

var _ ProtectionPolicyOperationResultsClientAPI = (*backup.ProtectionPolicyOperationResultsClient)(nil)

// JobsClientAPI contains the set of methods on the JobsClient type.
type JobsClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.JobResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.JobResourceListIterator, err error)
}

var _ JobsClientAPI = (*backup.JobsClient)(nil)

// JobDetailsClientAPI contains the set of methods on the JobDetailsClient type.
type JobDetailsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, jobName string) (result backup.JobResource, err error)
}

var _ JobDetailsClientAPI = (*backup.JobDetailsClient)(nil)

// JobCancellationsClientAPI contains the set of methods on the JobCancellationsClient type.
type JobCancellationsClientAPI interface {
	Trigger(ctx context.Context, vaultName string, resourceGroupName string, jobName string) (result autorest.Response, err error)
}

var _ JobCancellationsClientAPI = (*backup.JobCancellationsClient)(nil)

// JobOperationResultsClientAPI contains the set of methods on the JobOperationResultsClient type.
type JobOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, jobName string, operationID string) (result autorest.Response, err error)
}

var _ JobOperationResultsClientAPI = (*backup.JobOperationResultsClient)(nil)

// ExportJobsOperationResultsClientAPI contains the set of methods on the ExportJobsOperationResultsClient type.
type ExportJobsOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result backup.OperationResultInfoBaseResource, err error)
}

var _ ExportJobsOperationResultsClientAPI = (*backup.ExportJobsOperationResultsClient)(nil)

// JobsGroupClientAPI contains the set of methods on the JobsGroupClient type.
type JobsGroupClientAPI interface {
	Export(ctx context.Context, vaultName string, resourceGroupName string, filter string) (result autorest.Response, err error)
}

var _ JobsGroupClientAPI = (*backup.JobsGroupClient)(nil)

// ProtectedItemsGroupClientAPI contains the set of methods on the ProtectedItemsGroupClient type.
type ProtectedItemsGroupClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.ProtectedItemResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.ProtectedItemResourceListIterator, err error)
}

var _ ProtectedItemsGroupClientAPI = (*backup.ProtectedItemsGroupClient)(nil)

// OperationClientAPI contains the set of methods on the OperationClient type.
type OperationClientAPI interface {
	Validate(ctx context.Context, vaultName string, resourceGroupName string, parameters backup.BasicValidateOperationRequest) (result backup.ValidateOperationsResponse, err error)
}

var _ OperationClientAPI = (*backup.OperationClient)(nil)

// EnginesClientAPI contains the set of methods on the EnginesClient type.
type EnginesClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, backupEngineName string, filter string, skipToken string) (result backup.EngineBaseResource, err error)
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.EngineBaseResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.EngineBaseResourceListIterator, err error)
}

var _ EnginesClientAPI = (*backup.EnginesClient)(nil)

// ProtectionContainerRefreshOperationResultsClientAPI contains the set of methods on the ProtectionContainerRefreshOperationResultsClient type.
type ProtectionContainerRefreshOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, operationID string) (result autorest.Response, err error)
}

var _ ProtectionContainerRefreshOperationResultsClientAPI = (*backup.ProtectionContainerRefreshOperationResultsClient)(nil)

// ProtectableContainersClientAPI contains the set of methods on the ProtectableContainersClient type.
type ProtectableContainersClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, filter string) (result backup.ProtectableContainerResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, filter string) (result backup.ProtectableContainerResourceListIterator, err error)
}

var _ ProtectableContainersClientAPI = (*backup.ProtectableContainersClient)(nil)

// ProtectionContainersClientAPI contains the set of methods on the ProtectionContainersClient type.
type ProtectionContainersClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string) (result backup.ProtectionContainerResource, err error)
	Inquire(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, filter string) (result autorest.Response, err error)
	Refresh(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, filter string) (result autorest.Response, err error)
	Register(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, parameters backup.ProtectionContainerResource) (result backup.ProtectionContainerResource, err error)
	Unregister(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string) (result autorest.Response, err error)
}

var _ ProtectionContainersClientAPI = (*backup.ProtectionContainersClient)(nil)

// WorkloadItemsClientAPI contains the set of methods on the WorkloadItemsClient type.
type WorkloadItemsClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, filter string, skipToken string) (result backup.WorkloadItemResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, filter string, skipToken string) (result backup.WorkloadItemResourceListIterator, err error)
}

var _ WorkloadItemsClientAPI = (*backup.WorkloadItemsClient)(nil)

// ProtectionContainerOperationResultsClientAPI contains the set of methods on the ProtectionContainerOperationResultsClient type.
type ProtectionContainerOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, operationID string) (result backup.ProtectionContainerResource, err error)
}

var _ ProtectionContainerOperationResultsClientAPI = (*backup.ProtectionContainerOperationResultsClient)(nil)

// BackupsClientAPI contains the set of methods on the BackupsClient type.
type BackupsClientAPI interface {
	Trigger(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, parameters backup.RequestResource) (result autorest.Response, err error)
}

var _ BackupsClientAPI = (*backup.BackupsClient)(nil)

// ProtectedItemOperationStatusesClientAPI contains the set of methods on the ProtectedItemOperationStatusesClient type.
type ProtectedItemOperationStatusesClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, operationID string) (result backup.OperationStatus, err error)
}

var _ ProtectedItemOperationStatusesClientAPI = (*backup.ProtectedItemOperationStatusesClient)(nil)

// ItemLevelRecoveryConnectionsClientAPI contains the set of methods on the ItemLevelRecoveryConnectionsClient type.
type ItemLevelRecoveryConnectionsClientAPI interface {
	Provision(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters backup.ILRRequestResource) (result autorest.Response, err error)
	Revoke(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string) (result autorest.Response, err error)
}

var _ ItemLevelRecoveryConnectionsClientAPI = (*backup.ItemLevelRecoveryConnectionsClient)(nil)

// OperationResultsClientAPI contains the set of methods on the OperationResultsClient type.
type OperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result autorest.Response, err error)
}

var _ OperationResultsClientAPI = (*backup.OperationResultsClient)(nil)

// OperationStatusesClientAPI contains the set of methods on the OperationStatusesClient type.
type OperationStatusesClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result backup.OperationStatus, err error)
}

var _ OperationStatusesClientAPI = (*backup.OperationStatusesClient)(nil)

// ProtectionPolicyOperationStatusesClientAPI contains the set of methods on the ProtectionPolicyOperationStatusesClient type.
type ProtectionPolicyOperationStatusesClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string) (result backup.OperationStatus, err error)
}

var _ ProtectionPolicyOperationStatusesClientAPI = (*backup.ProtectionPolicyOperationStatusesClient)(nil)

// ProtectableItemsClientAPI contains the set of methods on the ProtectableItemsClient type.
type ProtectableItemsClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.WorkloadProtectableItemResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.WorkloadProtectableItemResourceListIterator, err error)
}

var _ ProtectableItemsClientAPI = (*backup.ProtectableItemsClient)(nil)

// ProtectionContainersGroupClientAPI contains the set of methods on the ProtectionContainersGroupClient type.
type ProtectionContainersGroupClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string) (result backup.ProtectionContainerResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string) (result backup.ProtectionContainerResourceListIterator, err error)
}

var _ ProtectionContainersGroupClientAPI = (*backup.ProtectionContainersGroupClient)(nil)

// SecurityPINsClientAPI contains the set of methods on the SecurityPINsClient type.
type SecurityPINsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string) (result backup.TokenInformation, err error)
}

var _ SecurityPINsClientAPI = (*backup.SecurityPINsClient)(nil)

// RecoveryPointsRecommendedForMoveClientAPI contains the set of methods on the RecoveryPointsRecommendedForMoveClient type.
type RecoveryPointsRecommendedForMoveClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, parameters backup.ListRecoveryPointsRecommendedForMoveRequest) (result backup.RecoveryPointResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, parameters backup.ListRecoveryPointsRecommendedForMoveRequest) (result backup.RecoveryPointResourceListIterator, err error)
}

var _ RecoveryPointsRecommendedForMoveClientAPI = (*backup.RecoveryPointsRecommendedForMoveClient)(nil)
