# Release History

## 2.0.0 (2023-12-27)
### Breaking Changes

- Type of `ApplicationDefinitionProperties.CreateUIDefinition` has been changed from `any` to `interface{}`
- Type of `ApplicationDefinitionProperties.MainTemplate` has been changed from `any` to `interface{}`
- Type of `ApplicationProperties.Parameters` has been changed from `any` to `interface{}`
- Type of `ApplicationProperties.Outputs` has been changed from `any` to `interface{}`
- Type of `ApplicationPropertiesPatchable.Parameters` has been changed from `any` to `interface{}`
- Type of `ApplicationPropertiesPatchable.Outputs` has been changed from `any` to `interface{}`
- Function `NewClientFactory` has been removed
- Function `*ClientFactory.NewApplicationClient` has been removed
- Function `*ClientFactory.NewApplicationDefinitionsClient` has been removed
- Function `*ClientFactory.NewApplicationsClient` has been removed
- Struct `ClientFactory` has been removed


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).