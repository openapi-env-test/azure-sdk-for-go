# Unreleased Content

## Breaking Changes

### Signature Changes

#### Funcs

1. GalleriesClient.Get
	- Params
		- From: context.Context, string, string
		- To: context.Context, string, string, SelectPermissions
1. GalleriesClient.GetPreparer
	- Params
		- From: context.Context, string, string
		- To: context.Context, string, string, SelectPermissions

## Additive Changes

### New Constants

1. GallerySharingPermissionTypes.Groups
1. GallerySharingPermissionTypes.Private
1. SelectPermissions.Permissions
1. SharedToValues.Tenant
1. SharingProfileGroupTypes.AADTenants
1. SharingProfileGroupTypes.Subscriptions
1. SharingUpdateOperationTypes.Add
1. SharingUpdateOperationTypes.Remove
1. SharingUpdateOperationTypes.Reset

### New Funcs

1. *GallerySharingProfileUpdateFuture.UnmarshalJSON([]byte) error
1. *PirSharedGalleryResource.UnmarshalJSON([]byte) error
1. *SharedGallery.UnmarshalJSON([]byte) error
1. *SharedGalleryImage.UnmarshalJSON([]byte) error
1. *SharedGalleryImageListIterator.Next() error
1. *SharedGalleryImageListIterator.NextWithContext(context.Context) error
1. *SharedGalleryImageListPage.Next() error
1. *SharedGalleryImageListPage.NextWithContext(context.Context) error
1. *SharedGalleryImageVersion.UnmarshalJSON([]byte) error
1. *SharedGalleryImageVersionListIterator.Next() error
1. *SharedGalleryImageVersionListIterator.NextWithContext(context.Context) error
1. *SharedGalleryImageVersionListPage.Next() error
1. *SharedGalleryImageVersionListPage.NextWithContext(context.Context) error
1. *SharedGalleryListIterator.Next() error
1. *SharedGalleryListIterator.NextWithContext(context.Context) error
1. *SharedGalleryListPage.Next() error
1. *SharedGalleryListPage.NextWithContext(context.Context) error
1. GallerySharingProfileClient.Update(context.Context, string, string, SharingUpdate) (GallerySharingProfileUpdateFuture, error)
1. GallerySharingProfileClient.UpdatePreparer(context.Context, string, string, SharingUpdate) (*http.Request, error)
1. GallerySharingProfileClient.UpdateResponder(*http.Response) (SharingUpdate, error)
1. GallerySharingProfileClient.UpdateSender(*http.Request) (GallerySharingProfileUpdateFuture, error)
1. NewGallerySharingProfileClient(string) GallerySharingProfileClient
1. NewGallerySharingProfileClientWithBaseURI(string, string) GallerySharingProfileClient
1. NewSharedGalleriesClient(string) SharedGalleriesClient
1. NewSharedGalleriesClientWithBaseURI(string, string) SharedGalleriesClient
1. NewSharedGalleryImageListIterator(SharedGalleryImageListPage) SharedGalleryImageListIterator
1. NewSharedGalleryImageListPage(SharedGalleryImageList, func(context.Context, SharedGalleryImageList) (SharedGalleryImageList, error)) SharedGalleryImageListPage
1. NewSharedGalleryImageVersionListIterator(SharedGalleryImageVersionListPage) SharedGalleryImageVersionListIterator
1. NewSharedGalleryImageVersionListPage(SharedGalleryImageVersionList, func(context.Context, SharedGalleryImageVersionList) (SharedGalleryImageVersionList, error)) SharedGalleryImageVersionListPage
1. NewSharedGalleryImageVersionsClient(string) SharedGalleryImageVersionsClient
1. NewSharedGalleryImageVersionsClientWithBaseURI(string, string) SharedGalleryImageVersionsClient
1. NewSharedGalleryImagesClient(string) SharedGalleryImagesClient
1. NewSharedGalleryImagesClientWithBaseURI(string, string) SharedGalleryImagesClient
1. NewSharedGalleryListIterator(SharedGalleryListPage) SharedGalleryListIterator
1. NewSharedGalleryListPage(SharedGalleryList, func(context.Context, SharedGalleryList) (SharedGalleryList, error)) SharedGalleryListPage
1. PirSharedGalleryResource.MarshalJSON() ([]byte, error)
1. PossibleGallerySharingPermissionTypesValues() []GallerySharingPermissionTypes
1. PossibleSelectPermissionsValues() []SelectPermissions
1. PossibleSharedToValuesValues() []SharedToValues
1. PossibleSharingProfileGroupTypesValues() []SharingProfileGroupTypes
1. PossibleSharingUpdateOperationTypesValues() []SharingUpdateOperationTypes
1. SharedGalleriesClient.Get(context.Context, string, string) (SharedGallery, error)
1. SharedGalleriesClient.GetPreparer(context.Context, string, string) (*http.Request, error)
1. SharedGalleriesClient.GetResponder(*http.Response) (SharedGallery, error)
1. SharedGalleriesClient.GetSender(*http.Request) (*http.Response, error)
1. SharedGalleriesClient.List(context.Context, string, SharedToValues) (SharedGalleryListPage, error)
1. SharedGalleriesClient.ListComplete(context.Context, string, SharedToValues) (SharedGalleryListIterator, error)
1. SharedGalleriesClient.ListPreparer(context.Context, string, SharedToValues) (*http.Request, error)
1. SharedGalleriesClient.ListResponder(*http.Response) (SharedGalleryList, error)
1. SharedGalleriesClient.ListSender(*http.Request) (*http.Response, error)
1. SharedGallery.MarshalJSON() ([]byte, error)
1. SharedGalleryImage.MarshalJSON() ([]byte, error)
1. SharedGalleryImageList.IsEmpty() bool
1. SharedGalleryImageListIterator.NotDone() bool
1. SharedGalleryImageListIterator.Response() SharedGalleryImageList
1. SharedGalleryImageListIterator.Value() SharedGalleryImage
1. SharedGalleryImageListPage.NotDone() bool
1. SharedGalleryImageListPage.Response() SharedGalleryImageList
1. SharedGalleryImageListPage.Values() []SharedGalleryImage
1. SharedGalleryImageVersion.MarshalJSON() ([]byte, error)
1. SharedGalleryImageVersionList.IsEmpty() bool
1. SharedGalleryImageVersionListIterator.NotDone() bool
1. SharedGalleryImageVersionListIterator.Response() SharedGalleryImageVersionList
1. SharedGalleryImageVersionListIterator.Value() SharedGalleryImageVersion
1. SharedGalleryImageVersionListPage.NotDone() bool
1. SharedGalleryImageVersionListPage.Response() SharedGalleryImageVersionList
1. SharedGalleryImageVersionListPage.Values() []SharedGalleryImageVersion
1. SharedGalleryImageVersionsClient.Get(context.Context, string, string, string, string) (SharedGalleryImageVersion, error)
1. SharedGalleryImageVersionsClient.GetPreparer(context.Context, string, string, string, string) (*http.Request, error)
1. SharedGalleryImageVersionsClient.GetResponder(*http.Response) (SharedGalleryImageVersion, error)
1. SharedGalleryImageVersionsClient.GetSender(*http.Request) (*http.Response, error)
1. SharedGalleryImageVersionsClient.List(context.Context, string, string, string, SharedToValues) (SharedGalleryImageVersionListPage, error)
1. SharedGalleryImageVersionsClient.ListComplete(context.Context, string, string, string, SharedToValues) (SharedGalleryImageVersionListIterator, error)
1. SharedGalleryImageVersionsClient.ListPreparer(context.Context, string, string, string, SharedToValues) (*http.Request, error)
1. SharedGalleryImageVersionsClient.ListResponder(*http.Response) (SharedGalleryImageVersionList, error)
1. SharedGalleryImageVersionsClient.ListSender(*http.Request) (*http.Response, error)
1. SharedGalleryImagesClient.Get(context.Context, string, string, string) (SharedGalleryImage, error)
1. SharedGalleryImagesClient.GetPreparer(context.Context, string, string, string) (*http.Request, error)
1. SharedGalleryImagesClient.GetResponder(*http.Response) (SharedGalleryImage, error)
1. SharedGalleryImagesClient.GetSender(*http.Request) (*http.Response, error)
1. SharedGalleryImagesClient.List(context.Context, string, string, SharedToValues) (SharedGalleryImageListPage, error)
1. SharedGalleryImagesClient.ListComplete(context.Context, string, string, SharedToValues) (SharedGalleryImageListIterator, error)
1. SharedGalleryImagesClient.ListPreparer(context.Context, string, string, SharedToValues) (*http.Request, error)
1. SharedGalleryImagesClient.ListResponder(*http.Response) (SharedGalleryImageList, error)
1. SharedGalleryImagesClient.ListSender(*http.Request) (*http.Response, error)
1. SharedGalleryList.IsEmpty() bool
1. SharedGalleryListIterator.NotDone() bool
1. SharedGalleryListIterator.Response() SharedGalleryList
1. SharedGalleryListIterator.Value() SharedGallery
1. SharedGalleryListPage.NotDone() bool
1. SharedGalleryListPage.Response() SharedGalleryList
1. SharedGalleryListPage.Values() []SharedGallery
1. SharingProfile.MarshalJSON() ([]byte, error)

### Struct Changes

#### New Structs

1. GalleryImageFeature
1. GallerySharingProfileClient
1. GallerySharingProfileUpdateFuture
1. PirResource
1. PirSharedGalleryResource
1. SharedGalleriesClient
1. SharedGallery
1. SharedGalleryIdentifier
1. SharedGalleryImage
1. SharedGalleryImageList
1. SharedGalleryImageListIterator
1. SharedGalleryImageListPage
1. SharedGalleryImageProperties
1. SharedGalleryImageVersion
1. SharedGalleryImageVersionList
1. SharedGalleryImageVersionListIterator
1. SharedGalleryImageVersionListPage
1. SharedGalleryImageVersionProperties
1. SharedGalleryImageVersionsClient
1. SharedGalleryImagesClient
1. SharedGalleryList
1. SharedGalleryListIterator
1. SharedGalleryListPage
1. SharingProfile
1. SharingProfileGroup
1. SharingUpdate

#### New Struct Fields

1. GalleryArtifactVersionSource.URI
1. GalleryImageProperties.Features
1. GalleryProperties.SharingProfile
