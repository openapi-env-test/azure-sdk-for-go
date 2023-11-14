# Release History

## 3.0.0 (2023-11-14)
### Breaking Changes

- Type of `AzureCliScriptProperties.Outputs` has been changed from `map[string]any` to `map[string]interface{}`
- Type of `AzurePowerShellScriptProperties.Outputs` has been changed from `map[string]any` to `map[string]interface{}`
- Type of `ErrorAdditionalInfo.Info` has been changed from `any` to `interface{}`
- Function `NewClientFactory` has been removed
- Function `*ClientFactory.NewClient` has been removed
- Struct `ClientFactory` has been removed


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentscripts` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).