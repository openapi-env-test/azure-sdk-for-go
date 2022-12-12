# Release History

## 1.2.0 (2022-12-12)
### Features Added

- New const `ProvisioningStateCanceled`
- New const `LanguageExtensionImageNamePython3912IncludeDeepLearning`
- New const `LanguageExtensionImageNamePython365`
- New const `LanguageExtensionImageNameR`
- New const `LanguageExtensionImageNamePython3912`
- New const `DataConnectionKindCosmosDb`
- New const `LanguageExtensionImageNamePython3108`
- New type alias `LanguageExtensionImageName`
- New function `NewSKUsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SKUsClient, error)`
- New function `PossibleLanguageExtensionImageNameValues() []LanguageExtensionImageName`
- New function `*SKUsClient.NewListPager(string, *SKUsClientListOptions) *runtime.Pager[SKUsClientListResponse]`
- New function `*CosmosDbDataConnection.GetDataConnection() *DataConnection`
- New struct `CosmosDbDataConnection`
- New struct `CosmosDbDataConnectionProperties`
- New struct `ProxyResource`
- New struct `Resource`
- New struct `ResourceSKUCapabilities`
- New struct `ResourceSKUZoneDetails`
- New struct `SKUsClient`
- New struct `SKUsClientListOptions`
- New struct `SKUsClientListResponse`
- New struct `TrackedResource`
- New field `LanguageExtensionImageName` in struct `LanguageExtension`
- New field `ZoneDetails` in struct `SKULocationInfoItem`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).