# Release History

## 0.9.0 (2023-11-14)
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
- Type alias `AssignmentScopeValidation` has been removed
- Type alias `OverrideKind` has been removed
- Type alias `SelectorKind` has been removed
- Function `NewClientFactory` has been removed
- Function `*ClientFactory.NewAssignmentsClient` has been removed
- Function `*ClientFactory.NewDataPolicyManifestsClient` has been removed
- Function `*ClientFactory.NewDefinitionsClient` has been removed
- Function `*ClientFactory.NewExemptionsClient` has been removed
- Function `*ClientFactory.NewSetDefinitionsClient` has been removed
- Function `*ClientFactory.NewVariableValuesClient` has been removed
- Function `*ClientFactory.NewVariablesClient` has been removed
- Function `*ExemptionsClient.Update` has been removed
- Function `NewVariableValuesClient` has been removed
- Function `*VariableValuesClient.CreateOrUpdate` has been removed
- Function `*VariableValuesClient.CreateOrUpdateAtManagementGroup` has been removed
- Function `*VariableValuesClient.Delete` has been removed
- Function `*VariableValuesClient.DeleteAtManagementGroup` has been removed
- Function `*VariableValuesClient.Get` has been removed
- Function `*VariableValuesClient.GetAtManagementGroup` has been removed
- Function `*VariableValuesClient.NewListForManagementGroupPager` has been removed
- Function `*VariableValuesClient.NewListPager` has been removed
- Function `NewVariablesClient` has been removed
- Function `*VariablesClient.CreateOrUpdate` has been removed
- Function `*VariablesClient.CreateOrUpdateAtManagementGroup` has been removed
- Function `*VariablesClient.Delete` has been removed
- Function `*VariablesClient.DeleteAtManagementGroup` has been removed
- Function `*VariablesClient.Get` has been removed
- Function `*VariablesClient.GetAtManagementGroup` has been removed
- Function `*VariablesClient.NewListForManagementGroupPager` has been removed
- Function `*VariablesClient.NewListPager` has been removed
- Struct `AssignmentUpdateProperties` has been removed
- Struct `ClientFactory` has been removed
- Struct `ExemptionUpdate` has been removed
- Struct `ExemptionUpdateProperties` has been removed
- Struct `Override` has been removed
- Struct `ResourceSelector` has been removed
- Struct `Selector` has been removed
- Struct `Variable` has been removed
- Struct `VariableColumn` has been removed
- Struct `VariableListResult` has been removed
- Struct `VariableProperties` has been removed
- Struct `VariableValue` has been removed
- Struct `VariableValueColumnValue` has been removed
- Struct `VariableValueListResult` has been removed
- Struct `VariableValueProperties` has been removed
- Struct `VariableValuesClientListForManagementGroupResponse` has been removed
- Struct `VariableValuesClientListResponse` has been removed
- Struct `VariablesClientListForManagementGroupResponse` has been removed
- Struct `VariablesClientListResponse` has been removed
- Field `Overrides` of struct `AssignmentProperties` has been removed
- Field `ResourceSelectors` of struct `AssignmentProperties` has been removed
- Field `Properties` of struct `AssignmentUpdate` has been removed
- Field `AssignmentScopeValidation` of struct `ExemptionProperties` has been removed
- Field `ResourceSelectors` of struct `ExemptionProperties` has been removed

### Features Added

- New struct `ErrorAdditionalInfo`
- New struct `ErrorResponse`


## 0.6.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).