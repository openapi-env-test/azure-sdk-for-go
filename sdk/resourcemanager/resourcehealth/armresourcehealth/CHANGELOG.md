# Release History

## 2.0.0 (2022-12-06)
### Breaking Changes

- Struct `ErrorResponseError` has been removed
- Field `Error` of struct `ErrorResponse` has been removed
- Field `OccurredTime` of struct `AvailabilityStatusProperties` has been removed
- Field `UnavailabilitySummary` of struct `AvailabilityStatusPropertiesRecentlyResolved` has been removed
- Field `UnavailableOccurredTime` of struct `AvailabilityStatusPropertiesRecentlyResolved` has been removed

### Features Added

- New const `EventStatusValuesActive`
- New const `EventLevelValuesWarning`
- New const `EventTypeValuesEmergingIssues`
- New const `EventSourceValuesResourceHealth`
- New const `CreatedByTypeUser`
- New const `LinkTypeValuesButton`
- New const `LinkTypeValuesHyperlink`
- New const `CreatedByTypeKey`
- New const `LevelValuesCritical`
- New const `EventSourceValuesServiceHealth`
- New const `EventTypeValuesServiceIssue`
- New const `EventLevelValuesInformational`
- New const `EventTypeValuesHealthAdvisory`
- New const `EventLevelValuesError`
- New const `CreatedByTypeApplication`
- New const `EventTypeValuesRCA`
- New const `EventTypeValuesSecurityAdvisory`
- New const `EventTypeValuesPlannedMaintenance`
- New const `EventLevelValuesCritical`
- New const `CreatedByTypeManagedIdentity`
- New const `LevelValuesWarning`
- New const `EventStatusValuesResolved`
- New type alias `EventSourceValues`
- New type alias `LevelValues`
- New type alias `LinkTypeValues`
- New type alias `EventStatusValues`
- New type alias `EventTypeValues`
- New type alias `CreatedByType`
- New type alias `EventLevelValues`
- New function `*EventClient.GetBySubscriptionIDAndTrackingID(context.Context, string, *EventClientGetBySubscriptionIDAndTrackingIDOptions) (EventClientGetBySubscriptionIDAndTrackingIDResponse, error)`
- New function `PossibleEventSourceValuesValues() []EventSourceValues`
- New function `PossibleEventLevelValuesValues() []EventLevelValues`
- New function `NewEventClient(string, azcore.TokenCredential, *arm.ClientOptions) (*EventClient, error)`
- New function `*EventClient.GetByTenantIDAndTrackingID(context.Context, string, *EventClientGetByTenantIDAndTrackingIDOptions) (EventClientGetByTenantIDAndTrackingIDResponse, error)`
- New function `PossibleCreatedByTypeValues() []CreatedByType`
- New function `NewEventsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*EventsClient, error)`
- New function `PossibleEventTypeValuesValues() []EventTypeValues`
- New function `NewImpactedResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ImpactedResourcesClient, error)`
- New function `PossibleLevelValuesValues() []LevelValues`
- New function `*ImpactedResourcesClient.Get(context.Context, string, string, *ImpactedResourcesClientGetOptions) (ImpactedResourcesClientGetResponse, error)`
- New function `PossibleLinkTypeValuesValues() []LinkTypeValues`
- New function `*EventsClient.NewListBySingleResourcePager(string, *EventsClientListBySingleResourceOptions) *runtime.Pager[EventsClientListBySingleResourceResponse]`
- New function `PossibleEventStatusValuesValues() []EventStatusValues`
- New function `*EventsClient.NewListByTenantIDPager(*EventsClientListByTenantIDOptions) *runtime.Pager[EventsClientListByTenantIDResponse]`
- New function `*EventsClient.NewListBySubscriptionIDPager(*EventsClientListBySubscriptionIDOptions) *runtime.Pager[EventsClientListBySubscriptionIDResponse]`
- New function `*ImpactedResourcesClient.NewListBySubscriptionIDAndEventIDPager(string, *ImpactedResourcesClientListBySubscriptionIDAndEventIDOptions) *runtime.Pager[ImpactedResourcesClientListBySubscriptionIDAndEventIDResponse]`
- New struct `Event`
- New struct `EventClient`
- New struct `EventClientGetBySubscriptionIDAndTrackingIDOptions`
- New struct `EventClientGetBySubscriptionIDAndTrackingIDResponse`
- New struct `EventClientGetByTenantIDAndTrackingIDOptions`
- New struct `EventClientGetByTenantIDAndTrackingIDResponse`
- New struct `EventImpactedResource`
- New struct `EventImpactedResourceListResult`
- New struct `EventImpactedResourceProperties`
- New struct `EventProperties`
- New struct `EventPropertiesAdditionalInformation`
- New struct `EventPropertiesArticle`
- New struct `EventPropertiesRecommendedActions`
- New struct `EventPropertiesRecommendedActionsItem`
- New struct `Events`
- New struct `EventsClient`
- New struct `EventsClientListBySingleResourceOptions`
- New struct `EventsClientListBySingleResourceResponse`
- New struct `EventsClientListBySubscriptionIDOptions`
- New struct `EventsClientListBySubscriptionIDResponse`
- New struct `EventsClientListByTenantIDOptions`
- New struct `EventsClientListByTenantIDResponse`
- New struct `Faq`
- New struct `Impact`
- New struct `ImpactedResourcesClient`
- New struct `ImpactedResourcesClientGetOptions`
- New struct `ImpactedResourcesClientGetResponse`
- New struct `ImpactedResourcesClientListBySubscriptionIDAndEventIDOptions`
- New struct `ImpactedResourcesClientListBySubscriptionIDAndEventIDResponse`
- New struct `ImpactedServiceRegion`
- New struct `KeyValueItem`
- New struct `Link`
- New struct `LinkDisplayText`
- New struct `ProxyResource`
- New struct `SystemData`
- New struct `Update`
- New field `SystemData` in struct `Resource`
- New field `UnavailableSummary` in struct `AvailabilityStatusPropertiesRecentlyResolved`
- New field `UnavailableOccuredTime` in struct `AvailabilityStatusPropertiesRecentlyResolved`
- New field `Code` in struct `ErrorResponse`
- New field `Details` in struct `ErrorResponse`
- New field `Message` in struct `ErrorResponse`
- New field `OccuredTime` in struct `AvailabilityStatusProperties`
- New field `SystemData` in struct `ImpactedResourceStatus`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).