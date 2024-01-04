# Release History

## 2.0.0 (2024-01-04)
### Breaking Changes

- Function `NewClientFactory` has been removed
- Function `*ClientFactory.NewAccountsClient` has been removed
- Function `*ClientFactory.NewDefaultAccountsClient` has been removed
- Function `*ClientFactory.NewOperationsClient` has been removed
- Function `*ClientFactory.NewPrivateEndpointConnectionsClient` has been removed
- Function `*ClientFactory.NewPrivateLinkResourcesClient` has been removed
- Function `dateTimeRFC3339.MarshalText` has been removed
- Function `*dateTimeRFC3339.Parse` has been removed
- Function `*dateTimeRFC3339.UnmarshalText` has been removed
- Struct `ClientFactory` has been removed
- Field `Count` of struct `OperationList` has been removed
- Field `Count` of struct `PrivateEndpointConnectionList` has been removed
- Field `Count` of struct `PrivateLinkResourceList` has been removed

### Features Added

- New type alias `AccountProvisioningState` with values `AccountProvisioningStateCanceled`, `AccountProvisioningStateCreating`, `AccountProvisioningStateDeleting`, `AccountProvisioningStateFailed`, `AccountProvisioningStateMoving`, `AccountProvisioningStateSoftDeleted`, `AccountProvisioningStateSoftDeleting`, `AccountProvisioningStateSucceeded`, `AccountProvisioningStateUnknown`, `AccountProvisioningStateUpdating`
- New type alias `CredentialsType` with values `CredentialsTypeNone`, `CredentialsTypeSystemAssigned`, `CredentialsTypeUserAssigned`
- New type alias `EventHubType` with values `EventHubTypeHook`, `EventHubTypeNotification`
- New type alias `EventStreamingState` with values `EventStreamingStateDisabled`, `EventStreamingStateEnabled`
- New type alias `EventStreamingType` with values `EventStreamingTypeAzure`, `EventStreamingTypeManaged`, `EventStreamingTypeNone`
- New type alias `ManagedEventHubState` with values `ManagedEventHubStateDisabled`, `ManagedEventHubStateEnabled`, `ManagedEventHubStateNotSpecified`
- New type alias `ManagedResourcesPublicNetworkAccess` with values `ManagedResourcesPublicNetworkAccessDisabled`, `ManagedResourcesPublicNetworkAccessEnabled`, `ManagedResourcesPublicNetworkAccessNotSpecified`
- New function `NewFeaturesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*FeaturesClient, error)`
- New function `*FeaturesClient.AccountGet(context.Context, string, string, BatchFeatureRequest, *FeaturesClientAccountGetOptions) (FeaturesClientAccountGetResponse, error)`
- New function `*FeaturesClient.SubscriptionGet(context.Context, string, BatchFeatureRequest, *FeaturesClientSubscriptionGetOptions) (FeaturesClientSubscriptionGetResponse, error)`
- New function `NewKafkaConfigurationsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*KafkaConfigurationsClient, error)`
- New function `*KafkaConfigurationsClient.CreateOrUpdate(context.Context, string, string, string, KafkaConfiguration, *KafkaConfigurationsClientCreateOrUpdateOptions) (KafkaConfigurationsClientCreateOrUpdateResponse, error)`
- New function `*KafkaConfigurationsClient.Delete(context.Context, string, string, string, *KafkaConfigurationsClientDeleteOptions) (KafkaConfigurationsClientDeleteResponse, error)`
- New function `*KafkaConfigurationsClient.Get(context.Context, string, string, string, *KafkaConfigurationsClientGetOptions) (KafkaConfigurationsClientGetResponse, error)`
- New function `*KafkaConfigurationsClient.NewListByAccountPager(string, string, *KafkaConfigurationsClientListByAccountOptions) *runtime.Pager[KafkaConfigurationsClientListByAccountResponse]`
- New function `NewUsagesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*UsagesClient, error)`
- New function `*UsagesClient.Get(context.Context, string, *UsagesClientGetOptions) (UsagesClientGetResponse, error)`
- New function `timeRFC3339.MarshalText() ([]byte, error)`
- New function `*timeRFC3339.Parse(string) error`
- New function `*timeRFC3339.UnmarshalText([]byte) error`
- New struct `AccountPropertiesAccountStatus`
- New struct `AccountStatus`
- New struct `AccountStatusErrorDetails`
- New struct `BatchFeatureRequest`
- New struct `BatchFeatureStatus`
- New struct `Credentials`
- New struct `FeaturesClient`
- New struct `KafkaConfiguration`
- New struct `KafkaConfigurationList`
- New struct `KafkaConfigurationProperties`
- New struct `KafkaConfigurationsClient`
- New struct `KafkaConfigurationsClientListByAccountResponse`
- New struct `ProxyResourceSystemData`
- New struct `QuotaName`
- New struct `Usage`
- New struct `UsageList`
- New struct `UsageName`
- New struct `UsagesClient`
- New field `AccountStatus` in struct `AccountProperties`
- New field `ManagedEventHubState` in struct `AccountProperties`
- New field `ManagedResourcesPublicNetworkAccess` in struct `AccountProperties`
- New field `SystemData` in struct `PrivateEndpointConnection`
- New field `SystemData` in struct `ProxyResource`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).