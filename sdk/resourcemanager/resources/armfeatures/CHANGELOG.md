# Release History

## 2.0.0 (2024-01-08)
### Breaking Changes

- Function `NewClientFactory` has been removed
- Function `*ClientFactory.NewClient` has been removed
- Function `*ClientFactory.NewFeatureClient` has been removed
- Function `*ClientFactory.NewSubscriptionFeatureRegistrationsClient` has been removed
- Function `dateTimeRFC3339.MarshalText` has been removed
- Function `*dateTimeRFC3339.Parse` has been removed
- Function `*dateTimeRFC3339.UnmarshalText` has been removed
- Struct `ClientFactory` has been removed

### Features Added

- New function `timeRFC3339.MarshalText() ([]byte, error)`
- New function `*timeRFC3339.Parse(string) error`
- New function `*timeRFC3339.UnmarshalText([]byte) error`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armfeatures` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).