# Unreleased

## Breaking Changes

### Signature Changes

#### Struct Fields

1. OrchestrationServiceStateInput.ServiceName changed type from OrchestrationServiceNames to *string

## Additive Changes

### New Constants

1. AggregatedState.AggregatedStateFailed
1. AggregatedState.AggregatedStateInProgress
1. AggregatedState.AggregatedStateSucceeded
1. AggregatedState.AggregatedStateUnknown
1. GallerySharingPermissionTypes.GallerySharingPermissionTypesCommunity
1. SecurityProfileType.SecurityProfileTypeEncryptedVMGuestStateOnlyWithPmk
1. SecurityProfileType.SecurityProfileTypeEncryptedWithCmk
1. SecurityProfileType.SecurityProfileTypeEncryptedWithPmk
1. SharingUpdateOperationTypes.SharingUpdateOperationTypesEnableCommunity
1. State.StateFailed
1. State.StateInProgress
1. State.StateSucceeded
1. State.StateUnknown

### New Funcs

1. *CommunityGalleryImageListIterator.Next() error
1. *CommunityGalleryImageListIterator.NextWithContext(context.Context) error
1. *CommunityGalleryImageListPage.Next() error
1. *CommunityGalleryImageListPage.NextWithContext(context.Context) error
1. *CommunityGalleryImageVersionListIterator.Next() error
1. *CommunityGalleryImageVersionListIterator.NextWithContext(context.Context) error
1. *CommunityGalleryImageVersionListPage.Next() error
1. *CommunityGalleryImageVersionListPage.NextWithContext(context.Context) error
1. CommunityGalleryDataDiskImage.MarshalJSON() ([]byte, error)
1. CommunityGalleryDiskImage.MarshalJSON() ([]byte, error)
1. CommunityGalleryImageList.IsEmpty() bool
1. CommunityGalleryImageListIterator.NotDone() bool
1. CommunityGalleryImageListIterator.Response() CommunityGalleryImageList
1. CommunityGalleryImageListIterator.Value() CommunityGalleryImage
1. CommunityGalleryImageListPage.NotDone() bool
1. CommunityGalleryImageListPage.Response() CommunityGalleryImageList
1. CommunityGalleryImageListPage.Values() []CommunityGalleryImage
1. CommunityGalleryImageVersionList.IsEmpty() bool
1. CommunityGalleryImageVersionListIterator.NotDone() bool
1. CommunityGalleryImageVersionListIterator.Response() CommunityGalleryImageVersionList
1. CommunityGalleryImageVersionListIterator.Value() CommunityGalleryImageVersion
1. CommunityGalleryImageVersionListPage.NotDone() bool
1. CommunityGalleryImageVersionListPage.Response() CommunityGalleryImageVersionList
1. CommunityGalleryImageVersionListPage.Values() []CommunityGalleryImageVersion
1. CommunityGalleryImageVersionsClient.List(context.Context, string, string, string) (CommunityGalleryImageVersionListPage, error)
1. CommunityGalleryImageVersionsClient.ListComplete(context.Context, string, string, string) (CommunityGalleryImageVersionListIterator, error)
1. CommunityGalleryImageVersionsClient.ListPreparer(context.Context, string, string, string) (*http.Request, error)
1. CommunityGalleryImageVersionsClient.ListResponder(*http.Response) (CommunityGalleryImageVersionList, error)
1. CommunityGalleryImageVersionsClient.ListSender(*http.Request) (*http.Response, error)
1. CommunityGalleryImagesClient.List(context.Context, string, string) (CommunityGalleryImageListPage, error)
1. CommunityGalleryImagesClient.ListComplete(context.Context, string, string) (CommunityGalleryImageListIterator, error)
1. CommunityGalleryImagesClient.ListPreparer(context.Context, string, string) (*http.Request, error)
1. CommunityGalleryImagesClient.ListResponder(*http.Response) (CommunityGalleryImageList, error)
1. CommunityGalleryImagesClient.ListSender(*http.Request) (*http.Response, error)
1. CommunityGalleryInfo.MarshalJSON() ([]byte, error)
1. CommunityGalleryOSDiskImage.MarshalJSON() ([]byte, error)
1. NewCommunityGalleryImageListIterator(CommunityGalleryImageListPage) CommunityGalleryImageListIterator
1. NewCommunityGalleryImageListPage(CommunityGalleryImageList, func(context.Context, CommunityGalleryImageList) (CommunityGalleryImageList, error)) CommunityGalleryImageListPage
1. NewCommunityGalleryImageVersionListIterator(CommunityGalleryImageVersionListPage) CommunityGalleryImageVersionListIterator
1. NewCommunityGalleryImageVersionListPage(CommunityGalleryImageVersionList, func(context.Context, CommunityGalleryImageVersionList) (CommunityGalleryImageVersionList, error)) CommunityGalleryImageVersionListPage
1. PossibleAggregatedStateValues() []AggregatedState
1. PossibleSecurityProfileTypeValues() []SecurityProfileType
1. PossibleStateValues() []State

### Struct Changes

#### New Structs

1. CommunityGalleryDataDiskImage
1. CommunityGalleryDiskImage
1. CommunityGalleryImageList
1. CommunityGalleryImageListIterator
1. CommunityGalleryImageListPage
1. CommunityGalleryImageVersionList
1. CommunityGalleryImageVersionListIterator
1. CommunityGalleryImageVersionListPage
1. CommunityGalleryImageVersionStorageProfile
1. CommunityGalleryInfo
1. CommunityGalleryOSDiskImage
1. OSDiskImageSecurityProfile
1. RegionalSharingStatus
1. SharingStatus

#### New Struct Fields

1. CommunityGalleryImageProperties.Eula
1. CommunityGalleryImageProperties.PrivacyStatementURI
1. CommunityGalleryImageVersionProperties.ExcludeFromLatest
1. CommunityGalleryImageVersionProperties.StorageProfile
1. GalleryProperties.SharingStatus
1. OSDiskImageEncryption.SecurityProfile
1. SharingProfile.CommunityGalleryInfo
