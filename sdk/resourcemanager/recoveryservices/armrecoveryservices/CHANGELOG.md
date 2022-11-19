# Release History

## 2.0.0 (2022-11-19)
### Breaking Changes

- Struct `CloudError` has been removed

### Features Added

- New const `ImmutabilityStateUnlocked`
- New const `VaultSubResourceTypeAzureSiteRecovery`
- New const `ImmutabilityStateLocked`
- New const `ImmutabilityStateDisabled`
- New const `VaultSubResourceTypeAzureBackup`
- New const `StandardTierStorageRedundancyLocallyRedundant`
- New const `StandardTierStorageRedundancyZoneRedundant`
- New const `VaultSubResourceTypeAzureBackupSecondary`
- New const `CrossRegionRestoreDisabled`
- New const `CrossRegionRestoreEnabled`
- New const `PublicNetworkAccessEnabled`
- New const `PublicNetworkAccessDisabled`
- New const `StandardTierStorageRedundancyGeoRedundant`
- New type alias `VaultSubResourceType`
- New type alias `CrossRegionRestore`
- New type alias `ImmutabilityState`
- New type alias `PublicNetworkAccess`
- New type alias `StandardTierStorageRedundancy`
- New function `PossibleStandardTierStorageRedundancyValues() []StandardTierStorageRedundancy`
- New function `PossiblePublicNetworkAccessValues() []PublicNetworkAccess`
- New function `PossibleCrossRegionRestoreValues() []CrossRegionRestore`
- New function `*Client.Capabilities(context.Context, string, ResourceCapabilities, *ClientCapabilitiesOptions) (ClientCapabilitiesResponse, error)`
- New function `PossibleVaultSubResourceTypeValues() []VaultSubResourceType`
- New function `PossibleImmutabilityStateValues() []ImmutabilityState`
- New struct `CapabilitiesProperties`
- New struct `CapabilitiesResponse`
- New struct `CapabilitiesResponseProperties`
- New struct `ClientCapabilitiesOptions`
- New struct `ClientCapabilitiesResponse`
- New struct `DNSZone`
- New struct `DNSZoneResponse`
- New struct `ImmutabilitySettings`
- New struct `ResourceCapabilities`
- New struct `ResourceCapabilitiesBase`
- New struct `SecuritySettings`
- New struct `VaultPropertiesRedundancySettings`
- New field `AADAudience` in struct `ResourceCertificateAndAADDetails`
- New field `SecuritySettings` in struct `VaultProperties`
- New field `RedundancySettings` in struct `VaultProperties`
- New field `PublicNetworkAccess` in struct `VaultProperties`
- New field `GroupIDs` in struct `PrivateEndpointConnection`


## 1.1.0 (2022-07-22)
### Features Added

- New const `StandardTierStorageRedundancyGeoRedundant`
- New const `StandardTierStorageRedundancyZoneRedundant`
- New const `CrossRegionRestoreEnabled`
- New const `CrossRegionRestoreDisabled`
- New const `StandardTierStorageRedundancyLocallyRedundant`
- New function `PossibleCrossRegionRestoreValues() []CrossRegionRestore`
- New function `PossibleStandardTierStorageRedundancyValues() []StandardTierStorageRedundancy`
- New struct `VaultPropertiesRedundancySettings`
- New field `RedundancySettings` in struct `VaultProperties`
- New field `AADAudience` in struct `ResourceCertificateAndAADDetails`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).