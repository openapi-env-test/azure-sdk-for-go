# Release History

## 2.0.0 (2023-12-27)
### Breaking Changes

- Function `*DeploymentsClient.CalculateTemplateHash` parameter(s) have been changed from `(context.Context, any, *DeploymentsClientCalculateTemplateHashOptions)` to `(context.Context, interface{}, *DeploymentsClientCalculateTemplateHashOptions)`
- Type of `DeploymentExportResult.Template` has been changed from `any` to `interface{}`
- Type of `DeploymentProperties.Template` has been changed from `any` to `interface{}`
- Type of `DeploymentProperties.Parameters` has been changed from `any` to `map[string]*DeploymentParameter`
- Type of `DeploymentPropertiesExtended.Outputs` has been changed from `any` to `interface{}`
- Type of `DeploymentPropertiesExtended.Parameters` has been changed from `any` to `interface{}`
- Type of `DeploymentWhatIfProperties.Parameters` has been changed from `any` to `map[string]*DeploymentParameter`
- Type of `DeploymentWhatIfProperties.Template` has been changed from `any` to `interface{}`
- Type of `ErrorAdditionalInfo.Info` has been changed from `any` to `interface{}`
- Type of `GenericResource.Properties` has been changed from `any` to `interface{}`
- Type of `GenericResourceExpanded.Properties` has been changed from `any` to `interface{}`
- Type of `HTTPMessage.Content` has been changed from `any` to `interface{}`
- Type of `ResourceGroupExportResult.Template` has been changed from `any` to `interface{}`
- Type of `WhatIfChange.After` has been changed from `any` to `interface{}`
- Type of `WhatIfChange.Before` has been changed from `any` to `interface{}`
- Type of `WhatIfPropertyChange.Before` has been changed from `any` to `interface{}`
- Type of `WhatIfPropertyChange.After` has been changed from `any` to `interface{}`
- Function `NewClientFactory` has been removed
- Function `*ClientFactory.NewClient` has been removed
- Function `*ClientFactory.NewDeploymentOperationsClient` has been removed
- Function `*ClientFactory.NewDeploymentsClient` has been removed
- Function `*ClientFactory.NewOperationsClient` has been removed
- Function `*ClientFactory.NewProviderResourceTypesClient` has been removed
- Function `*ClientFactory.NewProvidersClient` has been removed
- Function `*ClientFactory.NewResourceGroupsClient` has been removed
- Function `*ClientFactory.NewTagsClient` has been removed
- Function `dateTimeRFC3339.MarshalText` has been removed
- Function `*dateTimeRFC3339.Parse` has been removed
- Function `*dateTimeRFC3339.UnmarshalText` has been removed
- Operation `*TagsClient.CreateOrUpdateAtScope` has been changed to LRO, use `*TagsClient.BeginCreateOrUpdateAtScope` instead.
- Operation `*TagsClient.DeleteAtScope` has been changed to LRO, use `*TagsClient.BeginDeleteAtScope` instead.
- Operation `*TagsClient.UpdateAtScope` has been changed to LRO, use `*TagsClient.BeginUpdateAtScope` instead.
- Struct `ClientFactory` has been removed

### Features Added

- New function `timeRFC3339.MarshalText() ([]byte, error)`
- New function `*timeRFC3339.Parse(string) error`
- New function `*timeRFC3339.UnmarshalText([]byte) error`
- New struct `DeploymentParameter`
- New struct `KeyVaultParameterReference`
- New struct `KeyVaultReference`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).