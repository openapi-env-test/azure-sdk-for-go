# Release History

## 1.1.0-beta.2 (2023-01-16)
### Breaking Changes

- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed

### Features Added

- New type alias `ActivationStatus` with values `ActivationStatusActive`, `ActivationStatusFailed`, `ActivationStatusNotActivated`, `ActivationStatusUnknown`
- New function `*ManagedHsmsClient.CheckMhsmNameAvailability(context.Context, CheckMhsmNameAvailabilityParameters, *ManagedHsmsClientCheckMhsmNameAvailabilityOptions) (ManagedHsmsClientCheckMhsmNameAvailabilityResponse, error)`
- New struct `CheckMhsmNameAvailabilityParameters`
- New struct `CheckMhsmNameAvailabilityResult`
- New struct `ManagedHSMSecurityDomainProperties`
- New field `Etag` in struct `MHSMPrivateEndpointConnectionItem`
- New field `ID` in struct `MHSMPrivateEndpointConnectionItem`
- New field `SecurityDomainProperties` in struct `ManagedHsmProperties`


## 1.1.0-beta.1 (2022-05-19)
### Features Added

- New const `KeyRotationPolicyActionTypeNotify`
- New const `JSONWebKeyOperationRelease`
- New const `KeyRotationPolicyActionTypeRotate`
- New const `KeyPermissionsRotate`
- New const `KeyPermissionsRelease`
- New const `KeyPermissionsSetrotationpolicy`
- New const `KeyPermissionsGetrotationpolicy`
- New function `PossibleKeyRotationPolicyActionTypeValues() []KeyRotationPolicyActionType`
- New function `*KeyReleasePolicy.UnmarshalJSON([]byte) error`
- New function `KeyReleasePolicy.MarshalJSON() ([]byte, error)`
- New function `RotationPolicy.MarshalJSON() ([]byte, error)`
- New struct `Action`
- New struct `KeyReleasePolicy`
- New struct `KeyRotationPolicyAttributes`
- New struct `LifetimeAction`
- New struct `RotationPolicy`
- New struct `Trigger`
- New field `ReleasePolicy` in struct `KeyProperties`
- New field `RotationPolicy` in struct `KeyProperties`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).