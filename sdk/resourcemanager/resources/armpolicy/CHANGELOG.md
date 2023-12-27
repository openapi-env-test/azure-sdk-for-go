# Release History

## 0.10.0 (2023-12-27)
### Breaking Changes

- Type of `AssignmentProperties.Metadata` has been changed from `any` to `interface{}`
- Type of `DataEffect.DetailsSchema` has been changed from `any` to `interface{}`
- Type of `DefinitionProperties.PolicyRule` has been changed from `any` to `interface{}`
- Type of `DefinitionProperties.Metadata` has been changed from `any` to `interface{}`
- Type of `ExemptionProperties.Metadata` has been changed from `any` to `interface{}`
- Type of `ParameterDefinitionsValue.AllowedValues` has been changed from `[]any` to `[]interface{}`
- Type of `ParameterDefinitionsValue.DefaultValue` has been changed from `any` to `interface{}`
- Type of `ParameterDefinitionsValueMetadata.AdditionalProperties` has been changed from `map[string]any` to `map[string]interface{}`
- Type of `ParameterValuesValue.Value` has been changed from `any` to `interface{}`
- Type of `SetDefinitionProperties.Metadata` has been changed from `any` to `interface{}`
- Type of `VariableValueColumnValue.ColumnValue` has been changed from `any` to `interface{}`
- Function `NewClientFactory` has been removed
- Function `*ClientFactory.NewAssignmentsClient` has been removed
- Function `*ClientFactory.NewDataPolicyManifestsClient` has been removed
- Function `*ClientFactory.NewDefinitionsClient` has been removed
- Function `*ClientFactory.NewExemptionsClient` has been removed
- Function `*ClientFactory.NewSetDefinitionsClient` has been removed
- Function `*ClientFactory.NewVariableValuesClient` has been removed
- Function `*ClientFactory.NewVariablesClient` has been removed
- Function `dateTimeRFC3339.MarshalText` has been removed
- Function `*dateTimeRFC3339.Parse` has been removed
- Function `*dateTimeRFC3339.UnmarshalText` has been removed
- Struct `ClientFactory` has been removed

### Features Added

- New function `NewDefinitionVersionsClient(string, string, string, azcore.TokenCredential, *arm.ClientOptions) (*DefinitionVersionsClient, error)`
- New function `*DefinitionVersionsClient.CreateOrUpdate(context.Context, DefinitionVersion, *DefinitionVersionsClientCreateOrUpdateOptions) (DefinitionVersionsClientCreateOrUpdateResponse, error)`
- New function `*DefinitionVersionsClient.CreateOrUpdateAtManagementGroup(context.Context, string, DefinitionVersion, *DefinitionVersionsClientCreateOrUpdateAtManagementGroupOptions) (DefinitionVersionsClientCreateOrUpdateAtManagementGroupResponse, error)`
- New function `*DefinitionVersionsClient.Delete(context.Context, *DefinitionVersionsClientDeleteOptions) (DefinitionVersionsClientDeleteResponse, error)`
- New function `*DefinitionVersionsClient.DeleteAtManagementGroup(context.Context, string, *DefinitionVersionsClientDeleteAtManagementGroupOptions) (DefinitionVersionsClientDeleteAtManagementGroupResponse, error)`
- New function `*DefinitionVersionsClient.Get(context.Context, *DefinitionVersionsClientGetOptions) (DefinitionVersionsClientGetResponse, error)`
- New function `*DefinitionVersionsClient.GetAtManagementGroup(context.Context, string, *DefinitionVersionsClientGetAtManagementGroupOptions) (DefinitionVersionsClientGetAtManagementGroupResponse, error)`
- New function `*DefinitionVersionsClient.GetBuiltIn(context.Context, *DefinitionVersionsClientGetBuiltInOptions) (DefinitionVersionsClientGetBuiltInResponse, error)`
- New function `*DefinitionVersionsClient.ListAll(context.Context, *DefinitionVersionsClientListAllOptions) (DefinitionVersionsClientListAllResponse, error)`
- New function `*DefinitionVersionsClient.ListAllAtManagementGroup(context.Context, string, *DefinitionVersionsClientListAllAtManagementGroupOptions) (DefinitionVersionsClientListAllAtManagementGroupResponse, error)`
- New function `*DefinitionVersionsClient.ListAllBuiltins(context.Context, *DefinitionVersionsClientListAllBuiltinsOptions) (DefinitionVersionsClientListAllBuiltinsResponse, error)`
- New function `*DefinitionVersionsClient.NewListBuiltInPager(*DefinitionVersionsClientListBuiltInOptions) *runtime.Pager[DefinitionVersionsClientListBuiltInResponse]`
- New function `*DefinitionVersionsClient.NewListByManagementGroupPager(string, *DefinitionVersionsClientListByManagementGroupOptions) *runtime.Pager[DefinitionVersionsClientListByManagementGroupResponse]`
- New function `*DefinitionVersionsClient.NewListPager(*DefinitionVersionsClientListOptions) *runtime.Pager[DefinitionVersionsClientListResponse]`
- New function `NewSetDefinitionVersionsClient(string, string, string, azcore.TokenCredential, *arm.ClientOptions) (*SetDefinitionVersionsClient, error)`
- New function `*SetDefinitionVersionsClient.CreateOrUpdate(context.Context, SetDefinitionVersion, *SetDefinitionVersionsClientCreateOrUpdateOptions) (SetDefinitionVersionsClientCreateOrUpdateResponse, error)`
- New function `*SetDefinitionVersionsClient.CreateOrUpdateAtManagementGroup(context.Context, string, SetDefinitionVersion, *SetDefinitionVersionsClientCreateOrUpdateAtManagementGroupOptions) (SetDefinitionVersionsClientCreateOrUpdateAtManagementGroupResponse, error)`
- New function `*SetDefinitionVersionsClient.Delete(context.Context, *SetDefinitionVersionsClientDeleteOptions) (SetDefinitionVersionsClientDeleteResponse, error)`
- New function `*SetDefinitionVersionsClient.DeleteAtManagementGroup(context.Context, string, *SetDefinitionVersionsClientDeleteAtManagementGroupOptions) (SetDefinitionVersionsClientDeleteAtManagementGroupResponse, error)`
- New function `*SetDefinitionVersionsClient.Get(context.Context, *SetDefinitionVersionsClientGetOptions) (SetDefinitionVersionsClientGetResponse, error)`
- New function `*SetDefinitionVersionsClient.GetAtManagementGroup(context.Context, string, *SetDefinitionVersionsClientGetAtManagementGroupOptions) (SetDefinitionVersionsClientGetAtManagementGroupResponse, error)`
- New function `*SetDefinitionVersionsClient.GetBuiltIn(context.Context, *SetDefinitionVersionsClientGetBuiltInOptions) (SetDefinitionVersionsClientGetBuiltInResponse, error)`
- New function `*SetDefinitionVersionsClient.ListAll(context.Context, *SetDefinitionVersionsClientListAllOptions) (SetDefinitionVersionsClientListAllResponse, error)`
- New function `*SetDefinitionVersionsClient.ListAllAtManagementGroup(context.Context, string, *SetDefinitionVersionsClientListAllAtManagementGroupOptions) (SetDefinitionVersionsClientListAllAtManagementGroupResponse, error)`
- New function `*SetDefinitionVersionsClient.ListAllBuiltins(context.Context, *SetDefinitionVersionsClientListAllBuiltinsOptions) (SetDefinitionVersionsClientListAllBuiltinsResponse, error)`
- New function `*SetDefinitionVersionsClient.NewListBuiltInPager(*SetDefinitionVersionsClientListBuiltInOptions) *runtime.Pager[SetDefinitionVersionsClientListBuiltInResponse]`
- New function `*SetDefinitionVersionsClient.NewListByManagementGroupPager(string, *SetDefinitionVersionsClientListByManagementGroupOptions) *runtime.Pager[SetDefinitionVersionsClientListByManagementGroupResponse]`
- New function `*SetDefinitionVersionsClient.NewListPager(*SetDefinitionVersionsClientListOptions) *runtime.Pager[SetDefinitionVersionsClientListResponse]`
- New function `timeRFC3339.MarshalText() ([]byte, error)`
- New function `*timeRFC3339.Parse(string) error`
- New function `*timeRFC3339.UnmarshalText([]byte) error`
- New struct `DefinitionVersion`
- New struct `DefinitionVersionListResult`
- New struct `DefinitionVersionProperties`
- New struct `DefinitionVersionsClient`
- New struct `DefinitionVersionsClientListBuiltInResponse`
- New struct `DefinitionVersionsClientListByManagementGroupResponse`
- New struct `DefinitionVersionsClientListResponse`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorResponse`
- New struct `SetDefinitionVersion`
- New struct `SetDefinitionVersionListResult`
- New struct `SetDefinitionVersionProperties`
- New struct `SetDefinitionVersionsClient`
- New struct `SetDefinitionVersionsClientListBuiltInResponse`
- New struct `SetDefinitionVersionsClientListByManagementGroupResponse`
- New struct `SetDefinitionVersionsClientListResponse`
- New field `Version` in struct `DefinitionProperties`
- New field `Versions` in struct `DefinitionProperties`
- New field `Schema` in struct `ParameterDefinitionsValue`
- New field `Version` in struct `SetDefinitionProperties`
- New field `Versions` in struct `SetDefinitionProperties`


## 0.6.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).