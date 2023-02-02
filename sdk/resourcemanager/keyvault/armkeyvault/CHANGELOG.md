# Release History

## 2.0.0 (2023-02-02)
### Breaking Changes

- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed

### Features Added

- New value `JSONWebKeyOperationRelease` added to type alias `JSONWebKeyOperation`
- New value `KeyPermissionsGetrotationpolicy`, `KeyPermissionsRelease`, `KeyPermissionsRotate`, `KeyPermissionsSetrotationpolicy` added to type alias `KeyPermissions`
- New type alias `ActivationStatus` with values `ActivationStatusActive`, `ActivationStatusFailed`, `ActivationStatusNotActivated`, `ActivationStatusUnknown`
- New type alias `KeyRotationPolicyActionType` with values `KeyRotationPolicyActionTypeNotify`, `KeyRotationPolicyActionTypeRotate`
- New function `*ManagedHsmsClient.CheckMhsmNameAvailability(context.Context, CheckMhsmNameAvailabilityParameters, *ManagedHsmsClientCheckMhsmNameAvailabilityOptions) (ManagedHsmsClientCheckMhsmNameAvailabilityResponse, error)`
- New struct `Action`
- New struct `CheckMhsmNameAvailabilityParameters`
- New struct `CheckMhsmNameAvailabilityResult`
- New struct `KeyReleasePolicy`
- New struct `KeyRotationPolicyAttributes`
- New struct `LifetimeAction`
- New struct `ManagedHSMSecurityDomainProperties`
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