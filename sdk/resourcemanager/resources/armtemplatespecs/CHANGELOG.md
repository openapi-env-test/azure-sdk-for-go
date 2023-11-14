# Release History

## 2.0.0 (2023-11-14)
### Breaking Changes

- Type of `ErrorAdditionalInfo.Info` has been changed from `any` to `interface{}`
- Type of `LinkedTemplateArtifact.Template` has been changed from `any` to `interface{}`
- Type of `TemplateSpecProperties.Metadata` has been changed from `any` to `interface{}`
- Type of `TemplateSpecVersionProperties.Metadata` has been changed from `any` to `interface{}`
- Type of `TemplateSpecVersionProperties.UIFormDefinition` has been changed from `any` to `interface{}`
- Type of `TemplateSpecVersionProperties.MainTemplate` has been changed from `any` to `interface{}`
- Function `NewClientFactory` has been removed
- Function `*ClientFactory.NewClient` has been removed
- Function `*ClientFactory.NewTemplateSpecVersionsClient` has been removed
- Struct `ClientFactory` has been removed


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armtemplatespecs` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).