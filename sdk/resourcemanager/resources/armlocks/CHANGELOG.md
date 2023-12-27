# Release History

## 2.0.0 (2023-12-27)
### Breaking Changes

- Type of `ErrorAdditionalInfo.Info` has been changed from `any` to `interface{}`
- Function `NewClientFactory` has been removed
- Function `*ClientFactory.NewAuthorizationOperationsClient` has been removed
- Function `*ClientFactory.NewManagementLocksClient` has been removed
- Function `dateTimeRFC3339.MarshalText` has been removed
- Function `*dateTimeRFC3339.Parse` has been removed
- Function `*dateTimeRFC3339.UnmarshalText` has been removed
- Struct `ClientFactory` has been removed

### Features Added

- New function `timeRFC3339.MarshalText() ([]byte, error)`
- New function `*timeRFC3339.Parse(string) error`
- New function `*timeRFC3339.UnmarshalText([]byte) error`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlocks` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).