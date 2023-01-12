# Release History

## 0.8.0 (2023-01-12)
### Breaking Changes

- Const `ProvisioningStateCreating`, `ProvisioningStateDeleting`, `ProvisioningStateUpdating` from type alias `ProvisioningState` has been removed
- Type alias `PrivateEndpointConnectionProvisioningState` has been removed
- Type alias `PrivateEndpointServiceConnectionStatus` has been removed
- Type alias `PublicNetworkAccess` has been removed
- Function `*FarmBeatsModelsClient.GetOperationResult` has been removed
- Function `*FarmBeatsModelsClient.BeginUpdate` has been removed
- Function `NewPrivateEndpointConnectionsClient` has been removed
- Function `*PrivateEndpointConnectionsClient.CreateOrUpdate` has been removed
- Function `*PrivateEndpointConnectionsClient.BeginDelete` has been removed
- Function `*PrivateEndpointConnectionsClient.Get` has been removed
- Function `*PrivateEndpointConnectionsClient.NewListByResourcePager` has been removed
- Function `NewPrivateLinkResourcesClient` has been removed
- Function `*PrivateLinkResourcesClient.Get` has been removed
- Function `*PrivateLinkResourcesClient.NewListByResourcePager` has been removed
- Struct `ArmAsyncOperation` has been removed
- Struct `FarmBeatsUpdateProperties` has been removed
- Struct `Identity` has been removed
- Struct `PrivateEndpoint` has been removed
- Struct `PrivateEndpointConnection` has been removed
- Struct `PrivateEndpointConnectionListResult` has been removed
- Struct `PrivateEndpointConnectionProperties` has been removed
- Struct `PrivateEndpointConnectionsClient` has been removed
- Struct `PrivateEndpointConnectionsClientDeleteResponse` has been removed
- Struct `PrivateEndpointConnectionsClientListByResourceResponse` has been removed
- Struct `PrivateLinkResource` has been removed
- Struct `PrivateLinkResourceListResult` has been removed
- Struct `PrivateLinkResourceProperties` has been removed
- Struct `PrivateLinkResourcesClient` has been removed
- Struct `PrivateLinkResourcesClientListByResourceResponse` has been removed
- Struct `PrivateLinkServiceConnectionState` has been removed
- Struct `SensorIntegration` has been removed
- Field `Identity` of struct `FarmBeats` has been removed
- Field `PrivateEndpointConnections` of struct `FarmBeatsProperties` has been removed
- Field `ProvisioningState` of struct `FarmBeatsProperties` has been removed
- Field `PublicNetworkAccess` of struct `FarmBeatsProperties` has been removed
- Field `SensorIntegration` of struct `FarmBeatsProperties` has been removed
- Field `Identity` of struct `FarmBeatsUpdateRequestModel` has been removed
- Field `Properties` of struct `FarmBeatsUpdateRequestModel` has been removed
- Field `SystemData` of struct `ProxyResource` has been removed
- Field `SystemData` of struct `Resource` has been removed
- Field `SystemData` of struct `TrackedResource` has been removed

### Features Added

- New function `*FarmBeatsModelsClient.Update(context.Context, string, string, FarmBeatsUpdateRequestModel, *FarmBeatsModelsClientUpdateOptions) (FarmBeatsModelsClientUpdateResponse, error)`
- New field `ProvisioningStateSz` in struct `FarmBeatsProperties`


## 0.7.0 (2022-08-23)
### Breaking Changes

- Operation `FarmBeatsModelsClient.Update` has been changed to LRO, use `FarmBeatsModels.BeginUpdate` instead.

### Features Added

- Sensor, Private endpoint & managedIdentity support added.

## 0.6.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
