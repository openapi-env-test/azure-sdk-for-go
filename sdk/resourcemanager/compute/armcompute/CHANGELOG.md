# Release History

## 0.2.0 (2021-12-10)
### Breaking Changes

- Function `RestorePointProvisioningDetails.MarshalJSON` has been removed
- Function `*AvailabilitySetsClient.CreateOrUpdate` has been removed
- Function `*RestorePointProvisioningDetails.UnmarshalJSON` has been removed
- Struct `AvailabilitySetsCreateOrUpdateOptions` has been removed
- Struct `AvailabilitySetsCreateOrUpdateResponse` has been removed
- Struct `AvailabilitySetsCreateOrUpdateResult` has been removed
- Struct `RestorePointProvisioningDetails` has been removed
- Field `ProvisioningDetails` of struct `RestorePointProperties` has been removed

### New Content

- New function `NewCommunityGalleryImageVersionsClient(*arm.Connection, string) *CommunityGalleryImageVersionsClient`
- New function `CommunityGalleryImageProperties.MarshalJSON() ([]byte, error)`
- New function `*CommunityGalleryImagesClient.Get(context.Context, string, string, string, *CommunityGalleryImagesGetOptions) (CommunityGalleryImagesGetResponse, error)`
- New function `NewCommunityGalleryImagesClient(*arm.Connection, string) *CommunityGalleryImagesClient`
- New function `NewCommunityGalleriesClient(*arm.Connection, string) *CommunityGalleriesClient`
- New function `CommunityGalleryImageVersionProperties.MarshalJSON() ([]byte, error)`
- New function `*CommunityGalleryImageProperties.UnmarshalJSON([]byte) error`
- New function `*CommunityGalleryImageVersionProperties.UnmarshalJSON([]byte) error`
- New function `*RestorePointProperties.UnmarshalJSON([]byte) error`
- New function `*CommunityGalleriesClient.Get(context.Context, string, string, *CommunityGalleriesGetOptions) (CommunityGalleriesGetResponse, error)`
- New function `*CommunityGalleryImageVersionsClient.Get(context.Context, string, string, string, string, *CommunityGalleryImageVersionsGetOptions) (CommunityGalleryImageVersionsGetResponse, error)`
- New struct `CommunityGalleriesClient`
- New struct `CommunityGalleriesGetOptions`
- New struct `CommunityGalleriesGetResponse`
- New struct `CommunityGalleriesGetResult`
- New struct `CommunityGallery`
- New struct `CommunityGalleryIdentifier`
- New struct `CommunityGalleryImage`
- New struct `CommunityGalleryImageProperties`
- New struct `CommunityGalleryImageVersion`
- New struct `CommunityGalleryImageVersionProperties`
- New struct `CommunityGalleryImageVersionsClient`
- New struct `CommunityGalleryImageVersionsGetOptions`
- New struct `CommunityGalleryImageVersionsGetResponse`
- New struct `CommunityGalleryImageVersionsGetResult`
- New struct `CommunityGalleryImagesClient`
- New struct `CommunityGalleryImagesGetOptions`
- New struct `CommunityGalleryImagesGetResponse`
- New struct `CommunityGalleryImagesGetResult`
- New struct `PirCommunityGalleryResource`
- New field `TimeCreated` in struct `RestorePointProperties`

Total 12 breaking change(s), 50 additive change(s).


## 0.1.0 (2021-09-29)
- To better align with the Azure SDK guidelines (https://azure.github.io/azure-sdk/general_introduction.html), we have decided to change the module path to "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute". Therefore, we are deprecating the old module path (which is "github.com/Azure/azure-sdk-for-go/sdk/compute/armcompute") to avoid confusion. 