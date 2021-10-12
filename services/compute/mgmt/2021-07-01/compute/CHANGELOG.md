# Unreleased

## Breaking Changes

### Removed Funcs

1. *GalleriesDeleteFuture.UnmarshalJSON([]byte) error
1. GalleriesClient.Delete(context.Context, string, string) (GalleriesDeleteFuture, error)
1. GalleriesClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. GalleriesClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. GalleriesClient.DeleteSender(*http.Request) (GalleriesDeleteFuture, error)
1. SharedGalleryImagesClient.Get(context.Context, string, string, string) (SharedGalleryImage, error)
1. SharedGalleryImagesClient.GetPreparer(context.Context, string, string, string) (*http.Request, error)
1. SharedGalleryImagesClient.GetResponder(*http.Response) (SharedGalleryImage, error)
1. SharedGalleryImagesClient.GetSender(*http.Request) (*http.Response, error)

### Struct Changes

#### Removed Structs

1. GalleriesDeleteFuture

#### Removed Struct Fields

1. SharedGalleryImage.autorest.Response
