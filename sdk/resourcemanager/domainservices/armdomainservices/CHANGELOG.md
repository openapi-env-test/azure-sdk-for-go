# Release History

## 2.0.0 (2022-12-06)
### Breaking Changes

- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed

### Features Added

- New const `SyncScopeCloudOnly`
- New const `LdapSigningEnabled`
- New const `ChannelBindingDisabled`
- New const `ChannelBindingEnabled`
- New const `LdapSigningDisabled`
- New const `SyncScopeAll`
- New type alias `LdapSigning`
- New type alias `SyncScope`
- New type alias `ChannelBinding`
- New function `PossibleChannelBindingValues() []ChannelBinding`
- New function `PossibleSyncScopeValues() []SyncScope`
- New function `PossibleLdapSigningValues() []LdapSigning`
- New field `LdapSigning` in struct `DomainSecuritySettings`
- New field `ChannelBinding` in struct `DomainSecuritySettings`
- New field `SyncScope` in struct `DomainServiceProperties`
- New field `SyncApplicationID` in struct `DomainServiceProperties`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).