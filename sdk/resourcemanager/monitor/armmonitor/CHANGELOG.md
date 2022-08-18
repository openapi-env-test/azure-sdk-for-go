# Release History

## 0.8.0 (2022-08-18)
### Breaking Changes

- Function `*DataCollectionRuleAssociationsClient.Create` parameter(s) have been changed from `(context.Context, string, string, *DataCollectionRuleAssociationsClientCreateOptions)` to `(context.Context, string, string, DataCollectionRuleAssociationProxyOnlyResource, *DataCollectionRuleAssociationsClientCreateOptions)`
- Function `*DataCollectionEndpointsClient.Update` parameter(s) have been changed from `(context.Context, string, string, *DataCollectionEndpointsClientUpdateOptions)` to `(context.Context, string, string, ResourceForUpdate, *DataCollectionEndpointsClientUpdateOptions)`
- Function `*DataCollectionEndpointsClient.Create` parameter(s) have been changed from `(context.Context, string, string, *DataCollectionEndpointsClientCreateOptions)` to `(context.Context, string, string, DataCollectionEndpointResource, *DataCollectionEndpointsClientCreateOptions)`
- Function `*DataCollectionRulesClient.Update` parameter(s) have been changed from `(context.Context, string, string, *DataCollectionRulesClientUpdateOptions)` to `(context.Context, string, string, ResourceForUpdate, *DataCollectionRulesClientUpdateOptions)`
- Function `*DataCollectionRulesClient.Create` parameter(s) have been changed from `(context.Context, string, string, *DataCollectionRulesClientCreateOptions)` to `(context.Context, string, string, DataCollectionRuleResource, *DataCollectionRulesClientCreateOptions)`
- Type of `OperationStatus.Error` has been changed from `*ErrorResponseCommon` to `*ErrorDetail`
- Type of `PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState` has been changed from `*PrivateLinkServiceConnectionStateProperty` to `*PrivateLinkServiceConnectionState`
- Type of `PrivateEndpointConnectionProperties.PrivateEndpoint` has been changed from `*PrivateEndpointProperty` to `*PrivateEndpoint`
- Type of `PrivateEndpointConnectionProperties.ProvisioningState` has been changed from `*string` to `*PrivateEndpointConnectionProvisioningState`
- Function `*PrivateLinkResourcesClient.NewListByPrivateLinkScopePager` has been removed
- Function `*DiagnosticSettingsCategoryClient.List` has been removed
- Function `*DiagnosticSettingsClient.List` has been removed
- Function `*PrivateEndpointConnectionsClient.NewListByPrivateLinkScopePager` has been removed
- Struct `ErrorResponseCommon` has been removed
- Struct `PrivateEndpointProperty` has been removed
- Struct `PrivateLinkScopesResource` has been removed
- Struct `PrivateLinkServiceConnectionStateProperty` has been removed
- Struct `ProxyOnlyResource` has been removed
- Field `NextLink` of struct `PrivateLinkResourceListResult` has been removed
- Field `Body` of struct `DataCollectionRulesClientCreateOptions` has been removed
- Field `Identity` of struct `ActionGroupResource` has been removed
- Field `Kind` of struct `ActionGroupResource` has been removed
- Field `Body` of struct `DataCollectionEndpointsClientUpdateOptions` has been removed
- Field `Identity` of struct `AzureResource` has been removed
- Field `Kind` of struct `AzureResource` has been removed
- Field `Body` of struct `DataCollectionEndpointsClientCreateOptions` has been removed
- Field `Body` of struct `DataCollectionRuleAssociationsClientCreateOptions` has been removed
- Field `NextLink` of struct `PrivateEndpointConnectionListResult` has been removed
- Field `Body` of struct `DataCollectionRulesClientUpdateOptions` has been removed

### Features Added

- New const `AccessModeOpen`
- New const `PredictiveAutoscalePolicyScaleModeForecastOnly`
- New const `PrivateEndpointServiceConnectionStatusApproved`
- New const `PrivateEndpointConnectionProvisioningStateDeleting`
- New const `PrivateEndpointConnectionProvisioningStateSucceeded`
- New const `PredictiveAutoscalePolicyScaleModeDisabled`
- New const `PrivateEndpointServiceConnectionStatusRejected`
- New const `PrivateEndpointConnectionProvisioningStateFailed`
- New const `PrivateEndpointServiceConnectionStatusPending`
- New const `PrivateEndpointConnectionProvisioningStateCreating`
- New const `AccessModePrivateOnly`
- New const `PredictiveAutoscalePolicyScaleModeEnabled`
- New function `PossibleAccessModeValues() []AccessMode`
- New function `*PrivateLinkResourcesClient.ListByPrivateLinkScope(context.Context, string, string, *PrivateLinkResourcesClientListByPrivateLinkScopeOptions) (PrivateLinkResourcesClientListByPrivateLinkScopeResponse, error)`
- New function `*DiagnosticSettingsClient.NewListPager(string, *DiagnosticSettingsClientListOptions) *runtime.Pager[DiagnosticSettingsClientListResponse]`
- New function `PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState`
- New function `PossiblePredictiveAutoscalePolicyScaleModeValues() []PredictiveAutoscalePolicyScaleMode`
- New function `PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus`
- New function `*PredictiveMetricClient.Get(context.Context, string, string, string, string, string, string, string, *PredictiveMetricClientGetOptions) (PredictiveMetricClientGetResponse, error)`
- New function `*PrivateEndpointConnectionsClient.ListByPrivateLinkScope(context.Context, string, string, *PrivateEndpointConnectionsClientListByPrivateLinkScopeOptions) (PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse, error)`
- New function `NewPredictiveMetricClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PredictiveMetricClient, error)`
- New function `*DiagnosticSettingsCategoryClient.NewListPager(string, *DiagnosticSettingsCategoryClientListOptions) *runtime.Pager[DiagnosticSettingsCategoryClientListResponse]`
- New struct `AccessModeSettings`
- New struct `AccessModeSettingsExclusion`
- New struct `AutoscaleErrorResponse`
- New struct `AutoscaleErrorResponseError`
- New struct `DefaultErrorResponse`
- New struct `PredictiveAutoscalePolicy`
- New struct `PredictiveMetricClient`
- New struct `PredictiveMetricClientGetOptions`
- New struct `PredictiveMetricClientGetResponse`
- New struct `PredictiveResponse`
- New struct `PredictiveValue`
- New struct `PrivateEndpoint`
- New struct `PrivateLinkServiceConnectionState`
- New struct `ProxyResourceAutoGenerated`
- New struct `ResourceAutoGenerated`
- New struct `ResourceAutoGenerated2`
- New struct `ResourceAutoGenerated3`
- New struct `ResourceAutoGenerated4`
- New struct `TrackedResource`
- New field `SystemData` in struct `DiagnosticSettingsCategoryResource`
- New anonymous field `TestNotificationDetailsResponse` in struct `ActionGroupsClientPostTestNotificationsResponse`
- New field `SystemData` in struct `Resource`
- New field `PredictiveAutoscalePolicy` in struct `AutoscaleSetting`
- New field `CategoryGroups` in struct `DiagnosticSettingsCategory`
- New field `SystemData` in struct `DiagnosticSettingsResource`
- New field `MarketplacePartnerID` in struct `DiagnosticSettings`
- New anonymous field `TestNotificationDetailsResponse` in struct `ActionGroupsClientCreateNotificationsAtActionGroupResourceLevelResponse`
- New field `CategoryGroup` in struct `LogSettings`
- New field `AccessModeSettings` in struct `AzureMonitorPrivateLinkScopeProperties`
- New field `SystemData` in struct `ScopedResource`
- New field `RequiredZoneNames` in struct `PrivateLinkResourceProperties`
- New field `SystemData` in struct `AzureMonitorPrivateLinkScope`
- New anonymous field `TestNotificationDetailsResponse` in struct `ActionGroupsClientCreateNotificationsAtResourceGroupLevelResponse`
- New field `SystemData` in struct `AutoscaleSettingResource`


## 0.7.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.7.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).