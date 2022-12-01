# Release History

## 3.0.0 (2022-12-01)
### Breaking Changes

- Const `OSSKUWindows2022` has been removed
- Const `OSSKUWindows2019` has been removed
- Function `*ManagedClustersClient.BeginRotateServiceAccountSigningKeys` has been removed
- Struct `ManagedClusterOIDCIssuerProfile` has been removed
- Struct `ManagedClusterStorageProfileBlobCSIDriver` has been removed
- Struct `ManagedClustersClientBeginRotateServiceAccountSigningKeysOptions` has been removed
- Struct `ManagedClustersClientRotateServiceAccountSigningKeysResponse` has been removed
- Field `BlobCSIDriver` of struct `ManagedClusterStorageProfile` has been removed
- Field `OidcIssuerProfile` of struct `ManagedClusterProperties` has been removed

### Features Added

- New const `ContainerServiceVMSizeTypesStandardDS2`
- New const `ContainerServiceVMSizeTypesStandardE6416SV3`
- New const `ContainerServiceVMSizeTypesStandardF16SV2`
- New const `ContainerServiceVMSizeTypesStandardDS14`
- New const `ContainerServiceVMSizeTypesStandardD8SV3`
- New const `ContainerServiceVMSizeTypesStandardE32SV3`
- New const `ContainerServiceVMSizeTypesStandardA10`
- New const `ContainerServiceVMSizeTypesStandardG4`
- New const `CountOne`
- New const `ContainerServiceVMSizeTypesStandardD4V2Promo`
- New const `ContainerServiceStorageProfileTypesStorageAccount`
- New const `ContainerServiceVMSizeTypesStandardD15V2`
- New const `ContainerServiceVMSizeTypesStandardD3`
- New const `ContainerServiceVMSizeTypesStandardD2V3`
- New const `ContainerServiceVMSizeTypesStandardDS4`
- New const `ContainerServiceVMSizeTypesStandardDS3V2Promo`
- New const `ContainerServiceVMSizeTypesStandardD4`
- New const `ContainerServiceVMSizeTypesStandardNV24`
- New const `ContainerServiceVMSizeTypesStandardA8V2`
- New const `ContainerServiceVMSizeTypesStandardNC12SV3`
- New const `ContainerServiceVMSizeTypesStandardE328SV3`
- New const `ContainerServiceVMSizeTypesStandardDS132V2`
- New const `ContainerServiceVMSizeTypesStandardD64V3`
- New const `ContainerServiceVMSizeTypesStandardA7`
- New const `ContainerServiceVMSizeTypesStandardE16V3`
- New const `ContainerServiceVMSizeTypesStandardD14V2Promo`
- New const `ContainerServiceVMSizeTypesStandardDS4V2Promo`
- New const `ContainerServiceVMSizeTypesStandardGS1`
- New const `ContainerServiceVMSizeTypesStandardNC24SV3`
- New const `ContainerServiceVMSizeTypesStandardDS1`
- New const `ContainerServiceVMSizeTypesStandardD2SV3`
- New const `ContainerServiceVMSizeTypesStandardGS48`
- New const `ContainerServiceVMSizeTypesStandardF4`
- New const `ContainerServiceVMSizeTypesStandardH16`
- New const `ContainerServiceVMSizeTypesStandardNC24RsV2`
- New const `ContainerServiceVMSizeTypesStandardDS11`
- New const `ContainerServiceVMSizeTypesStandardE16SV3`
- New const `ContainerServiceVMSizeTypesStandardDS4V2`
- New const `ContainerServiceVMSizeTypesStandardDS5V2`
- New const `ContainerServiceVMSizeTypesStandardB4Ms`
- New const `ContainerServiceVMSizeTypesStandardDS2V2`
- New const `ContainerServiceVMSizeTypesStandardNC24R`
- New const `ContainerServiceVMSizeTypesStandardD12V2Promo`
- New const `ContainerServiceVMSizeTypesStandardD1V2`
- New const `ContainerServiceVMSizeTypesStandardG2`
- New const `ContainerServiceVMSizeTypesStandardNV6`
- New const `ContainerServiceVMSizeTypesStandardA4V2`
- New const `ContainerServiceVMSizeTypesStandardND24S`
- New const `ContainerServiceVMSizeTypesStandardD16SV3`
- New const `ContainerServiceVMSizeTypesStandardG3`
- New const `ContainerServiceVMSizeTypesStandardM12832Ms`
- New const `ContainerServiceVMSizeTypesStandardGS2`
- New const `ContainerServiceVMSizeTypesStandardM6432Ms`
- New const `ContainerServiceVMSizeTypesStandardM6416Ms`
- New const `ContainerServiceVMSizeTypesStandardD11V2Promo`
- New const `CountThree`
- New const `ContainerServiceVMSizeTypesStandardDS11V2Promo`
- New const `ContainerServiceVMSizeTypesStandardDS12V2Promo`
- New const `ContainerServiceVMSizeTypesStandardA2V2`
- New const `ContainerServiceVMSizeTypesStandardD64SV3`
- New const `ContainerServiceVMSizeTypesStandardDS13`
- New const `ContainerServiceVMSizeTypesStandardF16`
- New const `ContainerServiceVMSizeTypesStandardD13`
- New const `ContainerServiceVMSizeTypesStandardH16R`
- New const `ContainerServiceVMSizeTypesStandardNC12SV2`
- New const `ContainerServiceVMSizeTypesStandardD12`
- New const `ContainerServiceVMSizeTypesStandardF4S`
- New const `ContainerServiceVMSizeTypesStandardNC6`
- New const `ContainerServiceVMSizeTypesStandardF4SV2`
- New const `ContainerServiceVMSizeTypesStandardE8SV3`
- New const `ContainerServiceVMSizeTypesStandardH8M`
- New const `ContainerServiceVMSizeTypesStandardNC24RsV3`
- New const `ContainerServiceVMSizeTypesStandardD2V2`
- New const `ContainerServiceVMSizeTypesStandardL16S`
- New const `ContainerServiceVMSizeTypesStandardA1V2`
- New const `ContainerServiceVMSizeTypesStandardDS148V2`
- New const `ContainerServiceVMSizeTypesStandardDS3V2`
- New const `ContainerServiceVMSizeTypesStandardDS14V2`
- New const `ContainerServiceVMSizeTypesStandardA3`
- New const `ContainerServiceVMSizeTypesStandardE2SV3`
- New const `ContainerServiceVMSizeTypesStandardDS5V2Promo`
- New const `ContainerServiceVMSizeTypesStandardH16M`
- New const `ContainerServiceVMSizeTypesStandardF2S`
- New const `ContainerServiceVMSizeTypesStandardDS12`
- New const `ContainerServiceVMSizeTypesStandardND12S`
- New const `ContainerServiceVMSizeTypesStandardND6S`
- New const `ContainerServiceVMSizeTypesStandardD12V2`
- New const `ContainerServiceVMSizeTypesStandardD4V3`
- New const `ContainerServiceVMSizeTypesStandardD11V2`
- New const `ContainerServiceVMSizeTypesStandardNC24SV2`
- New const `ContainerServiceVMSizeTypesStandardD14`
- New const `ContainerServiceVMSizeTypesStandardGS516`
- New const `ContainerServiceVMSizeTypesStandardM12864Ms`
- New const `ContainerServiceStorageProfileTypesManagedDisks`
- New const `ContainerServiceVMSizeTypesStandardD8V3`
- New const `ContainerServiceVMSizeTypesStandardM128S`
- New const `ContainerServiceVMSizeTypesStandardB2Ms`
- New const `ContainerServiceVMSizeTypesStandardA8`
- New const `ContainerServiceVMSizeTypesStandardD11`
- New const `ContainerServiceVMSizeTypesStandardA1`
- New const `ContainerServiceVMSizeTypesStandardH16Mr`
- New const `ContainerServiceVMSizeTypesStandardL32S`
- New const `ContainerServiceVMSizeTypesStandardNC24`
- New const `ContainerServiceVMSizeTypesStandardE3216SV3`
- New const `ContainerServiceVMSizeTypesStandardD13V2`
- New const `ContainerServiceVMSizeTypesStandardF2SV2`
- New const `ContainerServiceVMSizeTypesStandardM64Ms`
- New const `ContainerServiceVMSizeTypesStandardF1`
- New const `ContainerServiceVMSizeTypesStandardF32SV2`
- New const `ContainerServiceVMSizeTypesStandardD5V2Promo`
- New const `ContainerServiceVMSizeTypesStandardD32V3`
- New const `ContainerServiceVMSizeTypesStandardDS15V2`
- New const `ContainerServiceVMSizeTypesStandardF1S`
- New const `ContainerServiceVMSizeTypesStandardGS58`
- New const `ContainerServiceVMSizeTypesStandardB8Ms`
- New const `ContainerServiceVMSizeTypesStandardF8SV2`
- New const `ContainerServiceVMSizeTypesStandardM128Ms`
- New const `ContainerServiceVMSizeTypesStandardB2S`
- New const `ContainerServiceVMSizeTypesStandardD2V2Promo`
- New const `ContainerServiceVMSizeTypesStandardDS144V2`
- New const `ContainerServiceVMSizeTypesStandardDS12V2`
- New const `ContainerServiceVMSizeTypesStandardF16S`
- New const `ContainerServiceVMSizeTypesStandardD32SV3`
- New const `ContainerServiceVMSizeTypesStandardD3V2Promo`
- New const `ContainerServiceVMSizeTypesStandardNC12`
- New const `ContainerServiceVMSizeTypesStandardGS3`
- New const `ContainerServiceVMSizeTypesStandardF64SV2`
- New const `ContainerServiceVMSizeTypesStandardG1`
- New const `ContainerServiceVMSizeTypesStandardA6`
- New const `ContainerServiceVMSizeTypesStandardD5V2`
- New const `ContainerServiceVMSizeTypesStandardA2MV2`
- New const `ContainerServiceVMSizeTypesStandardF2`
- New const `ContainerServiceVMSizeTypesStandardE64V3`
- New const `ContainerServiceVMSizeTypesStandardD4V2`
- New const `ContainerServiceVMSizeTypesStandardND24Rs`
- New const `ContainerServiceVMSizeTypesStandardNV12`
- New const `ContainerServiceVMSizeTypesStandardM64S`
- New const `ContainerServiceVMSizeTypesStandardE32V3`
- New const `ContainerServiceVMSizeTypesStandardE4SV3`
- New const `ContainerServiceVMSizeTypesStandardA8MV2`
- New const `ContainerServiceVMSizeTypesStandardDS11V2`
- New const `ContainerServiceVMSizeTypesStandardD2`
- New const `ContainerServiceVMSizeTypesStandardDS13V2Promo`
- New const `ContainerServiceVMSizeTypesStandardNC6SV3`
- New const `ContainerServiceVMSizeTypesStandardA2`
- New const `ContainerServiceVMSizeTypesStandardDS13V2`
- New const `ContainerServiceVMSizeTypesStandardD16V3`
- New const `ContainerServiceVMSizeTypesStandardA9`
- New const `ContainerServiceVMSizeTypesStandardD1`
- New const `ContainerServiceVMSizeTypesStandardDS2V2Promo`
- New const `ContainerServiceVMSizeTypesStandardDS1V2`
- New const `ContainerServiceVMSizeTypesStandardGS44`
- New const `ContainerServiceVMSizeTypesStandardA11`
- New const `ContainerServiceVMSizeTypesStandardE8V3`
- New const `ContainerServiceVMSizeTypesStandardF8S`
- New const `ContainerServiceVMSizeTypesStandardNC6SV2`
- New const `ContainerServiceVMSizeTypesStandardL4S`
- New const `ContainerServiceVMSizeTypesStandardF72SV2`
- New const `ContainerServiceVMSizeTypesStandardF8`
- New const `ContainerServiceVMSizeTypesStandardA4`
- New const `ContainerServiceVMSizeTypesStandardH8`
- New const `ContainerServiceVMSizeTypesStandardD14V2`
- New const `ContainerServiceVMSizeTypesStandardE2V3`
- New const `ContainerServiceVMSizeTypesStandardD13V2Promo`
- New const `ContainerServiceVMSizeTypesStandardE6432SV3`
- New const `ContainerServiceVMSizeTypesStandardA5`
- New const `ContainerServiceVMSizeTypesStandardG5`
- New const `ContainerServiceVMSizeTypesStandardD3V2`
- New const `ContainerServiceVMSizeTypesStandardE64SV3`
- New const `ContainerServiceVMSizeTypesStandardGS5`
- New const `ContainerServiceVMSizeTypesStandardDS14V2Promo`
- New const `ContainerServiceVMSizeTypesStandardDS3`
- New const `ContainerServiceVMSizeTypesStandardDS134V2`
- New const `ContainerServiceVMSizeTypesStandardA4MV2`
- New const `ContainerServiceVMSizeTypesStandardE4V3`
- New const `CountFive`
- New const `ContainerServiceVMSizeTypesStandardGS4`
- New const `ContainerServiceVMSizeTypesStandardD4SV3`
- New const `ContainerServiceVMSizeTypesStandardL8S`
- New type alias `ContainerServiceStorageProfileTypes`
- New type alias `ContainerServiceVMSizeTypes`
- New type alias `Count`
- New function `PossibleCountValues() []Count`
- New function `PossibleContainerServiceVMSizeTypesValues() []ContainerServiceVMSizeTypes`
- New function `PossibleContainerServiceStorageProfileTypesValues() []ContainerServiceStorageProfileTypes`
- New struct `DiagnosticsProfile`
- New struct `MasterProfile`
- New struct `Resource`
- New struct `SubResource`
- New struct `TrackedResource`
- New struct `VMDiagnostics`


## 2.1.0 (2022-08-25)
### Features Added

- New const `OSSKUWindows2019`
- New const `OSSKUWindows2022`


## 2.0.0 (2022-07-22)
### Breaking Changes

- Struct `ManagedClusterSecurityProfileAzureDefender` has been removed
- Field `AzureDefender` of struct `ManagedClusterSecurityProfile` has been removed

### Features Added

- New const `KeyVaultNetworkAccessTypesPrivate`
- New const `NetworkPluginNone`
- New const `KeyVaultNetworkAccessTypesPublic`
- New function `PossibleKeyVaultNetworkAccessTypesValues() []KeyVaultNetworkAccessTypes`
- New struct `AzureKeyVaultKms`
- New struct `ManagedClusterSecurityProfileDefender`
- New struct `ManagedClusterSecurityProfileDefenderSecurityMonitoring`
- New field `HostGroupID` in struct `ManagedClusterAgentPoolProfileProperties`
- New field `HostGroupID` in struct `ManagedClusterAgentPoolProfile`
- New field `AzureKeyVaultKms` in struct `ManagedClusterSecurityProfile`
- New field `Defender` in struct `ManagedClusterSecurityProfile`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).