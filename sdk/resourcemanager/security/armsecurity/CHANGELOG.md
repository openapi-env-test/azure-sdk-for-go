# Release History

## 0.10.0 (2022-12-06)
### Breaking Changes

- Function `NewGovernanceRulesClient` parameter(s) have been changed from `(string, azcore.TokenCredential, *arm.ClientOptions)` to `(string, string, azcore.TokenCredential, *arm.ClientOptions)`
- Function `*SubscriptionGovernanceRulesExecuteStatusClient.BeginGet` has been removed
- Function `*ConnectorGovernanceRulesExecuteStatusClient.BeginGet` has been removed
- Function `*ConnectorGovernanceRulesClient.Delete` has been removed
- Struct `ConnectorGovernanceRulesClientDeleteOptions` has been removed
- Struct `ConnectorGovernanceRulesExecuteStatusClientBeginGetOptions` has been removed
- Struct `SubscriptionGovernanceRulesExecuteStatusClientBeginGetOptions` has been removed

### Features Added

- New function `NewAPICollectionClient(string, azcore.TokenCredential, *arm.ClientOptions) (*APICollectionClient, error)`
- New function `*ManagementGroupGovernanceRulesClient.Get(context.Context, string, *ManagementGroupGovernanceRulesClientGetOptions) (ManagementGroupGovernanceRulesClientGetResponse, error)`
- New function `*ManagementGroupGovernanceRulesDeleteStatusClient.Get(context.Context, string, string, *ManagementGroupGovernanceRulesDeleteStatusClientGetOptions) (ManagementGroupGovernanceRulesDeleteStatusClientGetResponse, error)`
- New function `*APICollectionOnboardingClient.Create(context.Context, string, string, string, *APICollectionOnboardingClientCreateOptions) (APICollectionOnboardingClientCreateResponse, error)`
- New function `*ManagementGroupGovernanceRuleClient.NewListPager(*ManagementGroupGovernanceRuleClientListOptions) *runtime.Pager[ManagementGroupGovernanceRuleClientListResponse]`
- New function `*ManagementGroupGovernanceRulesClient.CreateOrUpdate(context.Context, string, GovernanceRule, *ManagementGroupGovernanceRulesClientCreateOrUpdateOptions) (ManagementGroupGovernanceRulesClientCreateOrUpdateResponse, error)`
- New function `*ConnectorGovernanceRulesClient.BeginDelete(context.Context, string, string, string, *ConnectorGovernanceRulesClientBeginDeleteOptions) (*runtime.Poller[ConnectorGovernanceRulesClientDeleteResponse], error)`
- New function `*ManagementGroupGovernanceRulesClient.BeginDelete(context.Context, string, *ManagementGroupGovernanceRulesClientBeginDeleteOptions) (*runtime.Poller[ManagementGroupGovernanceRulesClientDeleteResponse], error)`
- New function `NewAPICollectionOnboardingClient(string, azcore.TokenCredential, *arm.ClientOptions) (*APICollectionOnboardingClient, error)`
- New function `*GovernanceRulesClient.BeginRuleIDExecuteSingleManagementGroup(context.Context, string, *GovernanceRulesClientBeginRuleIDExecuteSingleManagementGroupOptions) (*runtime.Poller[GovernanceRulesClientRuleIDExecuteSingleManagementGroupResponse], error)`
- New function `*SubscriptionGovernanceRulesExecuteStatusClient.Get(context.Context, string, string, *SubscriptionGovernanceRulesExecuteStatusClientGetOptions) (SubscriptionGovernanceRulesExecuteStatusClientGetResponse, error)`
- New function `NewManagementGroupGovernanceRulesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagementGroupGovernanceRulesClient, error)`
- New function `*ManagementGroupGovernanceRulesExecuteStatusClient.Get(context.Context, string, string, *ManagementGroupGovernanceRulesExecuteStatusClientGetOptions) (ManagementGroupGovernanceRulesExecuteStatusClientGetResponse, error)`
- New function `NewAPICollectionOffboardingClient(string, azcore.TokenCredential, *arm.ClientOptions) (*APICollectionOffboardingClient, error)`
- New function `*APICollectionClient.NewListPager(string, string, *APICollectionClientListOptions) *runtime.Pager[APICollectionClientListResponse]`
- New function `NewManagementGroupGovernanceRulesDeleteStatusClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagementGroupGovernanceRulesDeleteStatusClient, error)`
- New function `*APICollectionClient.Get(context.Context, string, string, string, *APICollectionClientGetOptions) (APICollectionClientGetResponse, error)`
- New function `*APICollectionOffboardingClient.Delete(context.Context, string, string, string, *APICollectionOffboardingClientDeleteOptions) (APICollectionOffboardingClientDeleteResponse, error)`
- New function `NewManagementGroupGovernanceRulesExecuteStatusClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagementGroupGovernanceRulesExecuteStatusClient, error)`
- New function `NewManagementGroupGovernanceRuleClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagementGroupGovernanceRuleClient, error)`
- New function `*ConnectorGovernanceRulesExecuteStatusClient.Get(context.Context, string, string, string, string, *ConnectorGovernanceRulesExecuteStatusClientGetOptions) (ConnectorGovernanceRulesExecuteStatusClientGetResponse, error)`
- New struct `APICollectionClient`
- New struct `APICollectionClientGetOptions`
- New struct `APICollectionClientGetResponse`
- New struct `APICollectionClientListOptions`
- New struct `APICollectionClientListResponse`
- New struct `APICollectionOffboardingClient`
- New struct `APICollectionOffboardingClientDeleteOptions`
- New struct `APICollectionOffboardingClientDeleteResponse`
- New struct `APICollectionOnboardingClient`
- New struct `APICollectionOnboardingClientCreateOptions`
- New struct `APICollectionOnboardingClientCreateResponse`
- New struct `APICollectionProperties`
- New struct `APICollectionResponse`
- New struct `APICollectionResponseList`
- New struct `ConnectorGovernanceRulesClientBeginDeleteOptions`
- New struct `ConnectorGovernanceRulesExecuteStatusClientGetOptions`
- New struct `ErrorDetail`
- New struct `ErrorResponse`
- New struct `GovernanceRuleMetadata`
- New struct `GovernanceRulesClientBeginRuleIDExecuteSingleManagementGroupOptions`
- New struct `GovernanceRulesClientRuleIDExecuteSingleManagementGroupResponse`
- New struct `ManagementGroupGovernanceRuleClient`
- New struct `ManagementGroupGovernanceRuleClientListOptions`
- New struct `ManagementGroupGovernanceRuleClientListResponse`
- New struct `ManagementGroupGovernanceRulesClient`
- New struct `ManagementGroupGovernanceRulesClientBeginDeleteOptions`
- New struct `ManagementGroupGovernanceRulesClientCreateOrUpdateOptions`
- New struct `ManagementGroupGovernanceRulesClientCreateOrUpdateResponse`
- New struct `ManagementGroupGovernanceRulesClientDeleteResponse`
- New struct `ManagementGroupGovernanceRulesClientGetOptions`
- New struct `ManagementGroupGovernanceRulesClientGetResponse`
- New struct `ManagementGroupGovernanceRulesDeleteStatusClient`
- New struct `ManagementGroupGovernanceRulesDeleteStatusClientGetOptions`
- New struct `ManagementGroupGovernanceRulesDeleteStatusClientGetResponse`
- New struct `ManagementGroupGovernanceRulesExecuteStatusClient`
- New struct `ManagementGroupGovernanceRulesExecuteStatusClientGetOptions`
- New struct `ManagementGroupGovernanceRulesExecuteStatusClientGetResponse`
- New struct `SubscriptionGovernanceRulesExecuteStatusClientGetOptions`
- New field `TenantID` in struct `GovernanceRuleProperties`
- New field `Metadata` in struct `GovernanceRuleProperties`
- New field `ExcludedScopes` in struct `GovernanceRuleProperties`
- New field `IncludeMemberScopes` in struct `GovernanceRuleProperties`
- New field `Location` in struct `ConnectorGovernanceRulesExecuteStatusClientGetResponse`
- New field `Location` in struct `SubscriptionGovernanceRulesExecuteStatusClientGetResponse`


## 0.7.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.7.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).