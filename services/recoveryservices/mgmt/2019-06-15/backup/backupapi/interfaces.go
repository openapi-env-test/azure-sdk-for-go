package backupapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/recoveryservices/mgmt/2019-06-15/backup"
	"github.com/Azure/go-autorest/autorest"
)

// ResourceVaultConfigsClientAPI contains the set of methods on the ResourceVaultConfigsClient type.
type ResourceVaultConfigsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string) (result backup.ResourceVaultConfigResource, err error)
	Put(ctx context.Context, vaultName string, resourceGroupName string, parameters backup.ResourceVaultConfigResource) (result backup.ResourceVaultConfigResource, err error)
	Update(ctx context.Context, vaultName string, resourceGroupName string, parameters backup.ResourceVaultConfigResource) (result backup.ResourceVaultConfigResource, err error)
}

var _ ResourceVaultConfigsClientAPI = (*backup.ResourceVaultConfigsClient)(nil)

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
	Trigger(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters backup.RestoreRequestResource) (result autorest.Response, err error)
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
