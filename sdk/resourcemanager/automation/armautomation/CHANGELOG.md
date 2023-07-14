# Release History

## 0.9.0 (2023-07-14)
### Breaking Changes

- Type of `EncryptionPropertiesIdentity.UserAssignedIdentity` has been changed from `any` to `interface{}`
- Type of `JobStreamProperties.Value` has been changed from `map[string]any` to `map[string]interface{}`
- Type of `NodeReportsClientGetContentResponse.Interface` has been changed from `any` to `interface{}`
- Type of `ScheduleCreateOrUpdateProperties.Interval` has been changed from `any` to `interface{}`
- Type of `ScheduleProperties.Interval` has been changed from `any` to `interface{}`
- Type of `SourceControlSyncJobStreamByIDProperties.Value` has been changed from `map[string]any` to `map[string]interface{}`
- Function `NewClientFactory` has been removed
- Function `*ClientFactory.NewAccountClient` has been removed
- Function `*ClientFactory.NewActivityClient` has been removed
- Function `*ClientFactory.NewAgentRegistrationInformationClient` has been removed
- Function `*ClientFactory.NewCertificateClient` has been removed
- Function `*ClientFactory.NewClient` has been removed
- Function `*ClientFactory.NewConnectionClient` has been removed
- Function `*ClientFactory.NewConnectionTypeClient` has been removed
- Function `*ClientFactory.NewCredentialClient` has been removed
- Function `*ClientFactory.NewDeletedAutomationAccountsClient` has been removed
- Function `*ClientFactory.NewDscCompilationJobClient` has been removed
- Function `*ClientFactory.NewDscCompilationJobStreamClient` has been removed
- Function `*ClientFactory.NewDscConfigurationClient` has been removed
- Function `*ClientFactory.NewDscNodeClient` has been removed
- Function `*ClientFactory.NewDscNodeConfigurationClient` has been removed
- Function `*ClientFactory.NewFieldsClient` has been removed
- Function `*ClientFactory.NewHybridRunbookWorkerGroupClient` has been removed
- Function `*ClientFactory.NewHybridRunbookWorkersClient` has been removed
- Function `*ClientFactory.NewJobClient` has been removed
- Function `*ClientFactory.NewJobScheduleClient` has been removed
- Function `*ClientFactory.NewJobStreamClient` has been removed
- Function `*ClientFactory.NewKeysClient` has been removed
- Function `*ClientFactory.NewLinkedWorkspaceClient` has been removed
- Function `*ClientFactory.NewModuleClient` has been removed
- Function `*ClientFactory.NewNodeCountInformationClient` has been removed
- Function `*ClientFactory.NewNodeReportsClient` has been removed
- Function `*ClientFactory.NewObjectDataTypesClient` has been removed
- Function `*ClientFactory.NewOperationsClient` has been removed
- Function `*ClientFactory.NewPrivateEndpointConnectionsClient` has been removed
- Function `*ClientFactory.NewPrivateLinkResourcesClient` has been removed
- Function `*ClientFactory.NewPython2PackageClient` has been removed
- Function `*ClientFactory.NewRunbookClient` has been removed
- Function `*ClientFactory.NewRunbookDraftClient` has been removed
- Function `*ClientFactory.NewScheduleClient` has been removed
- Function `*ClientFactory.NewSoftwareUpdateConfigurationMachineRunsClient` has been removed
- Function `*ClientFactory.NewSoftwareUpdateConfigurationRunsClient` has been removed
- Function `*ClientFactory.NewSoftwareUpdateConfigurationsClient` has been removed
- Function `*ClientFactory.NewSourceControlClient` has been removed
- Function `*ClientFactory.NewSourceControlSyncJobClient` has been removed
- Function `*ClientFactory.NewSourceControlSyncJobStreamsClient` has been removed
- Function `*ClientFactory.NewStatisticsClient` has been removed
- Function `*ClientFactory.NewTestJobClient` has been removed
- Function `*ClientFactory.NewTestJobStreamsClient` has been removed
- Function `*ClientFactory.NewUsagesClient` has been removed
- Function `*ClientFactory.NewVariableClient` has been removed
- Function `*ClientFactory.NewWatcherClient` has been removed
- Function `*ClientFactory.NewWebhookClient` has been removed
- Struct `ClientFactory` has been removed


## 0.7.0 (2022-07-12)
### Breaking Changes

- Function `*DscConfigurationClient.UpdateWithJSON` parameter(s) have been changed from `(context.Context, string, string, string, *DscConfigurationClientUpdateWithJSONOptions)` to `(context.Context, string, string, string, DscConfigurationUpdateParameters, *DscConfigurationClientUpdateWithJSONOptions)`
- Function `*DscConfigurationClient.UpdateWithText` parameter(s) have been changed from `(context.Context, string, string, string, *DscConfigurationClientUpdateWithTextOptions)` to `(context.Context, string, string, string, string, *DscConfigurationClientUpdateWithTextOptions)`
- Struct `HybridRunbookWorkerGroupUpdateParameters` has been removed
- Struct `HybridRunbookWorkerLegacy` has been removed
- Field `Parameters` of struct `DscConfigurationClientUpdateWithJSONOptions` has been removed
- Field `Credential` of struct `HybridRunbookWorkerGroup` has been removed
- Field `GroupType` of struct `HybridRunbookWorkerGroup` has been removed
- Field `HybridRunbookWorkers` of struct `HybridRunbookWorkerGroup` has been removed
- Field `Parameters` of struct `DscConfigurationClientUpdateWithTextOptions` has been removed
- Field `Credential` of struct `HybridRunbookWorkerGroupCreateOrUpdateParameters` has been removed

### Features Added

- New const `RunbookTypeEnumPython3`
- New const `RunbookTypeEnumPython2`
- New function `NewDeletedAutomationAccountsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DeletedAutomationAccountsClient, error)`
- New function `*DeletedAutomationAccountsClient.ListBySubscription(context.Context, *DeletedAutomationAccountsClientListBySubscriptionOptions) (DeletedAutomationAccountsClientListBySubscriptionResponse, error)`
- New struct `DeletedAutomationAccount`
- New struct `DeletedAutomationAccountListResult`
- New struct `DeletedAutomationAccountProperties`
- New struct `DeletedAutomationAccountsClient`
- New struct `DeletedAutomationAccountsClientListBySubscriptionOptions`
- New struct `DeletedAutomationAccountsClientListBySubscriptionResponse`
- New struct `HybridRunbookWorkerGroupCreateOrUpdateProperties`
- New struct `HybridRunbookWorkerGroupProperties`
- New field `Name` in struct `HybridRunbookWorkerGroupCreateOrUpdateParameters`
- New field `Properties` in struct `HybridRunbookWorkerGroupCreateOrUpdateParameters`
- New field `Properties` in struct `HybridRunbookWorkerGroup`


## 0.6.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).