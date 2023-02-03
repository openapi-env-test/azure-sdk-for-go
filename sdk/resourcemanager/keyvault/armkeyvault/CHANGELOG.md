# Release History

## 2.0.0 (2023-02-03)
### Breaking Changes

- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed

### Features Added

- New value `JSONWebKeyOperationRelease` added to type alias `JSONWebKeyOperation`
- New value `KeyPermissionsGetrotationpolicy`, `KeyPermissionsRelease`, `KeyPermissionsRotate`, `KeyPermissionsSetrotationpolicy` added to type alias `KeyPermissions`
- New type alias `ActivationStatus` with values `ActivationStatusActive`, `ActivationStatusFailed`, `ActivationStatusNotActivated`, `ActivationStatusUnknown`
- New type alias `KeyRotationPolicyActionType` with values `KeyRotationPolicyActionTypeNotify`, `KeyRotationPolicyActionTypeRotate`
- New function `NewManagedHsmKeysClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagedHsmKeysClient, error)`
- New function `*ManagedHsmKeysClient.CreateIfNotExist(context.Context, string, string, string, ManagedHsmKeyCreateParameters, *ManagedHsmKeysClientCreateIfNotExistOptions) (ManagedHsmKeysClientCreateIfNotExistResponse, error)`
- New function `*ManagedHsmKeysClient.Get(context.Context, string, string, string, *ManagedHsmKeysClientGetOptions) (ManagedHsmKeysClientGetResponse, error)`
- New function `*ManagedHsmKeysClient.GetVersion(context.Context, string, string, string, string, *ManagedHsmKeysClientGetVersionOptions) (ManagedHsmKeysClientGetVersionResponse, error)`
- New function `*ManagedHsmKeysClient.NewListPager(string, string, *ManagedHsmKeysClientListOptions) *runtime.Pager[ManagedHsmKeysClientListResponse]`
- New function `*ManagedHsmKeysClient.NewListVersionsPager(string, string, string, *ManagedHsmKeysClientListVersionsOptions) *runtime.Pager[ManagedHsmKeysClientListVersionsResponse]`
- New function `*ManagedHsmsClient.CheckMhsmNameAvailability(context.Context, CheckMhsmNameAvailabilityParameters, *ManagedHsmsClientCheckMhsmNameAvailabilityOptions) (ManagedHsmsClientCheckMhsmNameAvailabilityResponse, error)`
- New struct `Action`
- New struct `CheckMhsmNameAvailabilityParameters`
- New struct `CheckMhsmNameAvailabilityResult`
- New struct `KeyReleasePolicy`
- New struct `KeyRotationPolicyAttributes`
- New struct `LifetimeAction`
- New struct `ManagedHSMSecurityDomainProperties`
- New struct `ManagedHsmAction`
- New struct `ManagedHsmKey`
- New struct `ManagedHsmKeyAttributes`
- New struct `ManagedHsmKeyCreateParameters`
- New struct `ManagedHsmKeyListResult`
- New struct `ManagedHsmKeyProperties`
- New struct `ManagedHsmKeyReleasePolicy`
- New struct `ManagedHsmKeyRotationPolicyAttributes`
- New struct `ManagedHsmKeysClient`
- New struct `ManagedHsmKeysClientListResponse`
- New struct `ManagedHsmKeysClientListVersionsResponse`
- New struct `ManagedHsmLifetimeAction`
- New struct `ManagedHsmRotationPolicy`
- New struct `ManagedHsmTrigger`
- New struct `ProxyResourceWithoutSystemData`
- New struct `RotationPolicy`
- New struct `Trigger`
- New field `ReleasePolicy` in struct `KeyProperties`
- New field `RotationPolicy` in struct `KeyProperties`
- New field `Etag` in struct `MHSMPrivateEndpointConnectionItem`
- New field `ID` in struct `MHSMPrivateEndpointConnectionItem`
- New field `SecurityDomainProperties` in struct `ManagedHsmProperties`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).