# Release History

## 0.10.0 (2024-04-24)
### Breaking Changes

- Type of `ErrorAdditionalInfo.Info` has been changed from `any` to `interface{}`
- Const `PublicNetworkAccessHybrid` from type alias `PublicNetworkAccess` has been removed
- Function `NewClientFactory` has been removed
- Function `*ClientFactory.NewExtensionsClient` has been removed
- Function `*ClientFactory.NewFarmBeatsExtensionsClient` has been removed
- Function `*ClientFactory.NewFarmBeatsModelsClient` has been removed
- Function `*ClientFactory.NewLocationsClient` has been removed
- Function `*ClientFactory.NewOperationsClient` has been removed
- Function `*ClientFactory.NewPrivateEndpointConnectionsClient` has been removed
- Function `*ClientFactory.NewPrivateLinkResourcesClient` has been removed
- Function `*ExtensionsClient.Create` has been removed
- Function `*ExtensionsClient.NewListByFarmBeatsPager` has been removed
- Function `*ExtensionsClient.Update` has been removed
- Function `NewFarmBeatsExtensionsClient` has been removed
- Function `*FarmBeatsExtensionsClient.Get` has been removed
- Function `*FarmBeatsExtensionsClient.NewListPager` has been removed
- Function `NewFarmBeatsModelsClient` has been removed
- Function `*FarmBeatsModelsClient.CreateOrUpdate` has been removed
- Function `*FarmBeatsModelsClient.Delete` has been removed
- Function `*FarmBeatsModelsClient.Get` has been removed
- Function `*FarmBeatsModelsClient.GetOperationResult` has been removed
- Function `*FarmBeatsModelsClient.NewListByResourceGroupPager` has been removed
- Function `*FarmBeatsModelsClient.NewListBySubscriptionPager` has been removed
- Function `*FarmBeatsModelsClient.BeginUpdate` has been removed
- Function `NewLocationsClient` has been removed
- Function `*LocationsClient.CheckNameAvailability` has been removed
- Function `*PrivateEndpointConnectionsClient.NewListByResourcePager` has been removed
- Function `*PrivateLinkResourcesClient.NewListByResourcePager` has been removed
- Function `dateTimeRFC3339.MarshalText` has been removed
- Function `*dateTimeRFC3339.Parse` has been removed
- Function `*dateTimeRFC3339.UnmarshalText` has been removed
- Struct `ClientFactory` has been removed
- Struct `ExtensionsClientListByFarmBeatsResponse` has been removed
- Struct `FarmBeats` has been removed
- Struct `FarmBeatsExtension` has been removed
- Struct `FarmBeatsExtensionListResponse` has been removed
- Struct `FarmBeatsExtensionProperties` has been removed
- Struct `FarmBeatsExtensionsClientListResponse` has been removed
- Struct `FarmBeatsListResponse` has been removed
- Struct `FarmBeatsModelsClientListByResourceGroupResponse` has been removed
- Struct `FarmBeatsModelsClientListBySubscriptionResponse` has been removed
- Struct `FarmBeatsModelsClientUpdateResponse` has been removed
- Struct `FarmBeatsProperties` has been removed
- Struct `FarmBeatsUpdateProperties` has been removed
- Struct `FarmBeatsUpdateRequestModel` has been removed

### Features Added

- New value `ProvisioningStateRunning` added to type alias `ProvisioningState`
- New value `PublicNetworkAccessDisabled` added to type alias `PublicNetworkAccess`
- New type alias `AuthCredentialsKind` with values `AuthCredentialsKindAPIKeyAuthCredentials`, `AuthCredentialsKindOAuthClientCredentials`
- New function `*APIKeyAuthCredentials.GetAuthCredentials() *AuthCredentials`
- New function `*AuthCredentials.GetAuthCredentials() *AuthCredentials`
- New function `NewCheckNameAvailabilityClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CheckNameAvailabilityClient, error)`
- New function `*CheckNameAvailabilityClient.CheckNameAvailability(context.Context, CheckNameAvailabilityRequest, *CheckNameAvailabilityClientCheckNameAvailabilityOptions) (CheckNameAvailabilityClientCheckNameAvailabilityResponse, error)`
- New function `NewDataConnectorsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DataConnectorsClient, error)`
- New function `*DataConnectorsClient.CreateOrUpdate(context.Context, string, string, string, DataConnector, *DataConnectorsClientCreateOrUpdateOptions) (DataConnectorsClientCreateOrUpdateResponse, error)`
- New function `*DataConnectorsClient.Delete(context.Context, string, string, string, *DataConnectorsClientDeleteOptions) (DataConnectorsClientDeleteResponse, error)`
- New function `*DataConnectorsClient.Get(context.Context, string, string, string, *DataConnectorsClientGetOptions) (DataConnectorsClientGetResponse, error)`
- New function `*DataConnectorsClient.NewListPager(string, string, *DataConnectorsClientListOptions) *runtime.Pager[DataConnectorsClientListResponse]`
- New function `NewDataManagerForAgricultureExtensionsClient(azcore.TokenCredential, *arm.ClientOptions) (*DataManagerForAgricultureExtensionsClient, error)`
- New function `*DataManagerForAgricultureExtensionsClient.Get(context.Context, string, *DataManagerForAgricultureExtensionsClientGetOptions) (DataManagerForAgricultureExtensionsClientGetResponse, error)`
- New function `*DataManagerForAgricultureExtensionsClient.NewListPager(*DataManagerForAgricultureExtensionsClientListOptions) *runtime.Pager[DataManagerForAgricultureExtensionsClientListResponse]`
- New function `NewDataManagerForAgricultureResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DataManagerForAgricultureResourcesClient, error)`
- New function `*DataManagerForAgricultureResourcesClient.CreateOrUpdate(context.Context, string, string, DataManagerForAgriculture, *DataManagerForAgricultureResourcesClientCreateOrUpdateOptions) (DataManagerForAgricultureResourcesClientCreateOrUpdateResponse, error)`
- New function `*DataManagerForAgricultureResourcesClient.Delete(context.Context, string, string, *DataManagerForAgricultureResourcesClientDeleteOptions) (DataManagerForAgricultureResourcesClientDeleteResponse, error)`
- New function `*DataManagerForAgricultureResourcesClient.Get(context.Context, string, string, *DataManagerForAgricultureResourcesClientGetOptions) (DataManagerForAgricultureResourcesClientGetResponse, error)`
- New function `*DataManagerForAgricultureResourcesClient.NewListByResourceGroupPager(string, *DataManagerForAgricultureResourcesClientListByResourceGroupOptions) *runtime.Pager[DataManagerForAgricultureResourcesClientListByResourceGroupResponse]`
- New function `*DataManagerForAgricultureResourcesClient.NewListBySubscriptionPager(*DataManagerForAgricultureResourcesClientListBySubscriptionOptions) *runtime.Pager[DataManagerForAgricultureResourcesClientListBySubscriptionResponse]`
- New function `*DataManagerForAgricultureResourcesClient.BeginUpdate(context.Context, string, string, DataManagerForAgricultureUpdateRequestModel, *DataManagerForAgricultureResourcesClientBeginUpdateOptions) (*runtime.Poller[DataManagerForAgricultureResourcesClientUpdateResponse], error)`
- New function `*ExtensionsClient.CreateOrUpdate(context.Context, string, string, string, *ExtensionsClientCreateOrUpdateOptions) (ExtensionsClientCreateOrUpdateResponse, error)`
- New function `*ExtensionsClient.NewListByDataManagerForAgriculturePager(string, string, *ExtensionsClientListByDataManagerForAgricultureOptions) *runtime.Pager[ExtensionsClientListByDataManagerForAgricultureResponse]`
- New function `*OAuthClientCredentials.GetAuthCredentials() *AuthCredentials`
- New function `NewOperationResultsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*OperationResultsClient, error)`
- New function `*OperationResultsClient.Get(context.Context, string, string, *OperationResultsClientGetOptions) (OperationResultsClientGetResponse, error)`
- New function `*PrivateEndpointConnectionsClient.ListByResource(context.Context, string, string, *PrivateEndpointConnectionsClientListByResourceOptions) (PrivateEndpointConnectionsClientListByResourceResponse, error)`
- New function `*PrivateLinkResourcesClient.ListByResource(context.Context, string, string, *PrivateLinkResourcesClientListByResourceOptions) (PrivateLinkResourcesClientListByResourceResponse, error)`
- New function `NewSolutionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SolutionsClient, error)`
- New function `*SolutionsClient.CreateOrUpdate(context.Context, string, string, string, *SolutionsClientCreateOrUpdateOptions) (SolutionsClientCreateOrUpdateResponse, error)`
- New function `*SolutionsClient.Delete(context.Context, string, string, string, *SolutionsClientDeleteOptions) (SolutionsClientDeleteResponse, error)`
- New function `*SolutionsClient.Get(context.Context, string, string, string, *SolutionsClientGetOptions) (SolutionsClientGetResponse, error)`
- New function `*SolutionsClient.NewListPager(string, string, *SolutionsClientListOptions) *runtime.Pager[SolutionsClientListResponse]`
- New function `NewSolutionsDiscoverabilityClient(azcore.TokenCredential, *arm.ClientOptions) (*SolutionsDiscoverabilityClient, error)`
- New function `*SolutionsDiscoverabilityClient.Get(context.Context, string, *SolutionsDiscoverabilityClientGetOptions) (SolutionsDiscoverabilityClientGetResponse, error)`
- New function `*SolutionsDiscoverabilityClient.NewListPager(*SolutionsDiscoverabilityClientListOptions) *runtime.Pager[SolutionsDiscoverabilityClientListResponse]`
- New function `timeRFC3339.MarshalText() ([]byte, error)`
- New function `*timeRFC3339.Parse(string) error`
- New function `*timeRFC3339.UnmarshalText([]byte) error`
- New struct `APIKeyAuthCredentials`
- New struct `APIProperties`
- New struct `ArmAsyncOperationError`
- New struct `CheckNameAvailabilityClient`
- New struct `DataConnector`
- New struct `DataConnectorListResponse`
- New struct `DataConnectorProperties`
- New struct `DataConnectorsClient`
- New struct `DataConnectorsClientListResponse`
- New struct `DataManagerForAgriculture`
- New struct `DataManagerForAgricultureExtension`
- New struct `DataManagerForAgricultureExtensionListResponse`
- New struct `DataManagerForAgricultureExtensionProperties`
- New struct `DataManagerForAgricultureExtensionsClient`
- New struct `DataManagerForAgricultureExtensionsClientListResponse`
- New struct `DataManagerForAgricultureListResponse`
- New struct `DataManagerForAgricultureProperties`
- New struct `DataManagerForAgricultureResourcesClient`
- New struct `DataManagerForAgricultureResourcesClientListByResourceGroupResponse`
- New struct `DataManagerForAgricultureResourcesClientListBySubscriptionResponse`
- New struct `DataManagerForAgricultureResourcesClientUpdateResponse`
- New struct `DataManagerForAgricultureSolution`
- New struct `DataManagerForAgricultureSolutionListResponse`
- New struct `DataManagerForAgricultureSolutionProperties`
- New struct `DataManagerForAgricultureUpdateProperties`
- New struct `DataManagerForAgricultureUpdateRequestModel`
- New struct `ExtensionInstallationRequest`
- New struct `ExtensionsClientListByDataManagerForAgricultureResponse`
- New struct `KeyVaultProperties`
- New struct `MarketplaceOfferDetails`
- New struct `OAuthClientCredentials`
- New struct `OperationResultsClient`
- New struct `Solution`
- New struct `SolutionListResponse`
- New struct `SolutionProperties`
- New struct `SolutionsClient`
- New struct `SolutionsClientListResponse`
- New struct `SolutionsDiscoverabilityClient`
- New struct `SolutionsDiscoverabilityClientListResponse`
- New field `Error` in struct `ArmAsyncOperation`
- New field `APIDefaultInputParameters` in struct `DetailedInformation`
- New field `APIDocsLink` in struct `DetailedInformation`
- New field `APIType` in struct `DetailedInformation`
- New field `SkipToken` in struct `ExtensionListResponse`
- New field `AdditionalAPIProperties` in struct `ExtensionProperties`
- New field `GroupIDs` in struct `PrivateEndpointConnectionProperties`


## 0.7.0 (2022-08-23)
### Breaking Changes

- Operation `FarmBeatsModelsClient.Update` has been changed to LRO, use `FarmBeatsModels.BeginUpdate` instead.

### Features Added

- Sensor, Private endpoint & managedIdentity support added.

## 0.6.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
