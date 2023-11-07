# Release History

## 4.0.0-beta.1 (2023-11-07)
### Breaking Changes

- Type of `ErrorAdditionalInfo.Info` has been changed from `any` to `interface{}`
- Const `ProtectedItemStateBackupsSuspended` from type alias `ProtectedItemState` has been removed
- Const `ProtectionStateBackupsSuspended` from type alias `ProtectionState` has been removed
- Const `SoftDeleteFeatureStateAlwaysON` from type alias `SoftDeleteFeatureState` has been removed
- Type alias `TargetDiskNetworkAccessOption` has been removed
- Type alias `VaultSubResourceType` has been removed
- Function `*AzureVMWorkloadSAPHanaHSRProtectableItem.GetAzureVMWorkloadProtectableItem` has been removed
- Function `*AzureVMWorkloadSAPHanaHSRProtectableItem.GetWorkloadProtectableItem` has been removed
- Function `NewClientFactory` has been removed
- Function `*ClientFactory.NewBMSPrepareDataMoveOperationResultClient` has been removed
- Function `*ClientFactory.NewBackupEnginesClient` has been removed
- Function `*ClientFactory.NewBackupJobsClient` has been removed
- Function `*ClientFactory.NewBackupOperationResultsClient` has been removed
- Function `*ClientFactory.NewBackupOperationStatusesClient` has been removed
- Function `*ClientFactory.NewBackupPoliciesClient` has been removed
- Function `*ClientFactory.NewBackupProtectableItemsClient` has been removed
- Function `*ClientFactory.NewBackupProtectedItemsClient` has been removed
- Function `*ClientFactory.NewBackupProtectionContainersClient` has been removed
- Function `*ClientFactory.NewBackupProtectionIntentClient` has been removed
- Function `*ClientFactory.NewBackupResourceEncryptionConfigsClient` has been removed
- Function `*ClientFactory.NewBackupResourceStorageConfigsNonCRRClient` has been removed
- Function `*ClientFactory.NewBackupResourceVaultConfigsClient` has been removed
- Function `*ClientFactory.NewBackupStatusClient` has been removed
- Function `*ClientFactory.NewBackupUsageSummariesClient` has been removed
- Function `*ClientFactory.NewBackupWorkloadItemsClient` has been removed
- Function `*ClientFactory.NewBackupsClient` has been removed
- Function `*ClientFactory.NewClient` has been removed
- Function `*ClientFactory.NewDeletedProtectionContainersClient` has been removed
- Function `*ClientFactory.NewExportJobsOperationResultsClient` has been removed
- Function `*ClientFactory.NewFeatureSupportClient` has been removed
- Function `*ClientFactory.NewItemLevelRecoveryConnectionsClient` has been removed
- Function `*ClientFactory.NewJobCancellationsClient` has been removed
- Function `*ClientFactory.NewJobDetailsClient` has been removed
- Function `*ClientFactory.NewJobOperationResultsClient` has been removed
- Function `*ClientFactory.NewJobsClient` has been removed
- Function `*ClientFactory.NewOperationClient` has been removed
- Function `*ClientFactory.NewOperationsClient` has been removed
- Function `*ClientFactory.NewPrivateEndpointClient` has been removed
- Function `*ClientFactory.NewPrivateEndpointConnectionClient` has been removed
- Function `*ClientFactory.NewProtectableContainersClient` has been removed
- Function `*ClientFactory.NewProtectedItemOperationResultsClient` has been removed
- Function `*ClientFactory.NewProtectedItemOperationStatusesClient` has been removed
- Function `*ClientFactory.NewProtectedItemsClient` has been removed
- Function `*ClientFactory.NewProtectionContainerOperationResultsClient` has been removed
- Function `*ClientFactory.NewProtectionContainerRefreshOperationResultsClient` has been removed
- Function `*ClientFactory.NewProtectionContainersClient` has been removed
- Function `*ClientFactory.NewProtectionIntentClient` has been removed
- Function `*ClientFactory.NewProtectionPoliciesClient` has been removed
- Function `*ClientFactory.NewProtectionPolicyOperationResultsClient` has been removed
- Function `*ClientFactory.NewProtectionPolicyOperationStatusesClient` has been removed
- Function `*ClientFactory.NewRecoveryPointsClient` has been removed
- Function `*ClientFactory.NewRecoveryPointsRecommendedForMoveClient` has been removed
- Function `*ClientFactory.NewResourceGuardProxiesClient` has been removed
- Function `*ClientFactory.NewResourceGuardProxyClient` has been removed
- Function `*ClientFactory.NewRestoresClient` has been removed
- Function `*ClientFactory.NewSecurityPINsClient` has been removed
- Function `*ClientFactory.NewValidateOperationClient` has been removed
- Function `*ClientFactory.NewValidateOperationResultsClient` has been removed
- Function `*ClientFactory.NewValidateOperationStatusesClient` has been removed
- Function `NewDeletedProtectionContainersClient` has been removed
- Function `*DeletedProtectionContainersClient.NewListPager` has been removed
- Struct `AzureVMWorkloadSAPHanaHSRProtectableItem` has been removed
- Struct `ClientFactory` has been removed
- Struct `DeletedProtectionContainersClientListResponse` has been removed
- Struct `ExtendedLocation` has been removed
- Struct `RecoveryPointProperties` has been removed
- Struct `SecuredVMDetails` has been removed
- Struct `TargetDiskNetworkAccessSettings` has been removed
- Field `RecoveryPointProperties` of struct `AzureFileShareRecoveryPoint` has been removed
- Field `SoftDeleteRetentionPeriodInDays` of struct `AzureFileshareProtectedItem` has been removed
- Field `SoftDeleteRetentionPeriodInDays` of struct `AzureIaaSClassicComputeVMProtectedItem` has been removed
- Field `SoftDeleteRetentionPeriodInDays` of struct `AzureIaaSComputeVMProtectedItem` has been removed
- Field `SoftDeleteRetentionPeriodInDays` of struct `AzureIaaSVMProtectedItem` has been removed
- Field `SoftDeleteRetentionPeriodInDays` of struct `AzureSQLProtectedItem` has been removed
- Field `NodesList` of struct `AzureVMWorkloadProtectedItem` has been removed
- Field `SoftDeleteRetentionPeriodInDays` of struct `AzureVMWorkloadProtectedItem` has been removed
- Field `NodesList` of struct `AzureVMWorkloadSAPAseDatabaseProtectedItem` has been removed
- Field `SoftDeleteRetentionPeriodInDays` of struct `AzureVMWorkloadSAPAseDatabaseProtectedItem` has been removed
- Field `IsProtectable` of struct `AzureVMWorkloadSAPAseSystemProtectableItem` has been removed
- Field `IsProtectable` of struct `AzureVMWorkloadSAPHanaDBInstance` has been removed
- Field `NodesList` of struct `AzureVMWorkloadSAPHanaDBInstanceProtectedItem` has been removed
- Field `SoftDeleteRetentionPeriodInDays` of struct `AzureVMWorkloadSAPHanaDBInstanceProtectedItem` has been removed
- Field `IsProtectable` of struct `AzureVMWorkloadSAPHanaDatabaseProtectableItem` has been removed
- Field `NodesList` of struct `AzureVMWorkloadSAPHanaDatabaseProtectedItem` has been removed
- Field `SoftDeleteRetentionPeriodInDays` of struct `AzureVMWorkloadSAPHanaDatabaseProtectedItem` has been removed
- Field `IsProtectable` of struct `AzureVMWorkloadSAPHanaSystemProtectableItem` has been removed
- Field `IsProtectable` of struct `AzureVMWorkloadSQLAvailabilityGroupProtectableItem` has been removed
- Field `NodesList` of struct `AzureVMWorkloadSQLAvailabilityGroupProtectableItem` has been removed
- Field `IsProtectable` of struct `AzureVMWorkloadSQLDatabaseProtectableItem` has been removed
- Field `NodesList` of struct `AzureVMWorkloadSQLDatabaseProtectedItem` has been removed
- Field `SoftDeleteRetentionPeriodInDays` of struct `AzureVMWorkloadSQLDatabaseProtectedItem` has been removed
- Field `IsProtectable` of struct `AzureVMWorkloadSQLInstanceProtectableItem` has been removed
- Field `RecoveryPointProperties` of struct `AzureWorkloadPointInTimeRecoveryPoint` has been removed
- Field `RecoveryPointProperties` of struct `AzureWorkloadRecoveryPoint` has been removed
- Field `RecoveryPointProperties` of struct `AzureWorkloadSAPHanaPointInTimeRecoveryPoint` has been removed
- Field `RecoveryPointProperties` of struct `AzureWorkloadSAPHanaRecoveryPoint` has been removed
- Field `RecoveryPointProperties` of struct `AzureWorkloadSQLPointInTimeRecoveryPoint` has been removed
- Field `RecoveryPointProperties` of struct `AzureWorkloadSQLRecoveryPoint` has been removed
- Field `IncludeSoftDeletedRP` of struct `BMSRPQueryObject` has been removed
- Field `SoftDeleteRetentionPeriodInDays` of struct `BackupResourceVaultConfig` has been removed
- Field `AcquireStorageAccountLock` of struct `BackupStatusResponse` has been removed
- Field `ProtectedItemsCount` of struct `BackupStatusResponse` has been removed
- Field `SoftDeleteRetentionPeriodInDays` of struct `DPMProtectedItem` has been removed
- Field `SourceResourceID` of struct `DistributedNodesInfo` has been removed
- Field `SoftDeleteRetentionPeriodInDays` of struct `GenericProtectedItem` has been removed
- Field `RecoveryPointProperties` of struct `GenericRecoveryPoint` has been removed
- Field `IsPrivateAccessEnabledOnAnyDisk` of struct `IaasVMRecoveryPoint` has been removed
- Field `RecoveryPointProperties` of struct `IaasVMRecoveryPoint` has been removed
- Field `SecurityType` of struct `IaasVMRecoveryPoint` has been removed
- Field `ExtendedLocation` of struct `IaasVMRestoreRequest` has been removed
- Field `SecuredVMDetails` of struct `IaasVMRestoreRequest` has been removed
- Field `TargetDiskNetworkAccessSettings` of struct `IaasVMRestoreRequest` has been removed
- Field `ExtendedLocation` of struct `IaasVMRestoreWithRehydrationRequest` has been removed
- Field `SecuredVMDetails` of struct `IaasVMRestoreWithRehydrationRequest` has been removed
- Field `TargetDiskNetworkAccessSettings` of struct `IaasVMRestoreWithRehydrationRequest` has been removed
- Field `ProtectableItemCount` of struct `InquiryValidation` has been removed
- Field `SoftDeleteRetentionPeriodInDays` of struct `MabFileFolderProtectedItem` has been removed
- Field `GroupIDs` of struct `PrivateEndpointConnection` has been removed
- Field `ActionsRequired` of struct `PrivateLinkServiceConnectionState` has been removed
- Field `SoftDeleteRetentionPeriodInDays` of struct `ProtectedItem` has been removed

### Features Added

- New function `*AzureVMWorkloadSAPHanaHSR.GetAzureVMWorkloadProtectableItem() *AzureVMWorkloadProtectableItem`
- New function `*AzureVMWorkloadSAPHanaHSR.GetWorkloadProtectableItem() *WorkloadProtectableItem`
- New struct `AzureVMWorkloadSAPHanaHSR`
- New field `ActionRequired` in struct `PrivateLinkServiceConnectionState`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).