# Release History

## 2.0.0 (2023-10-30)
### Breaking Changes

- Type of `ErrorAdditionalInfo.Info` has been changed from `any` to `interface{}`
- Function `NewClientFactory` has been removed
- Function `*ClientFactory.NewLinkerClient` has been removed
- Function `*ClientFactory.NewOperationsClient` has been removed
- Function `*LinkerClient.BeginCreateOrUpdate` has been removed
- Function `*LinkerClient.BeginDelete` has been removed
- Function `*LinkerClient.Get` has been removed
- Function `*LinkerClient.NewListPager` has been removed
- Function `*LinkerClient.BeginUpdate` has been removed
- Struct `ClientFactory` has been removed
- Struct `LinkerClientCreateOrUpdateResponse` has been removed
- Struct `LinkerClientDeleteResponse` has been removed
- Struct `LinkerClientListResponse` has been removed
- Struct `LinkerClientUpdateResponse` has been removed

### Features Added

- New value `ClientTypeKafkaSpringBoot` added to type alias `ClientType`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).