# Unreleased

## Breaking Changes

### Removed Funcs

1. *BastionHostsDeleteFuture.UnmarshalJSON([]byte) error
1. *InterfacesDeleteFuture.UnmarshalJSON([]byte) error
1. BastionHostsClient.Delete(context.Context, string, string) (BastionHostsDeleteFuture, error)
1. BastionHostsClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. BastionHostsClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. BastionHostsClient.DeleteSender(*http.Request) (BastionHostsDeleteFuture, error)
1. InterfacesClient.Delete(context.Context, string, string) (InterfacesDeleteFuture, error)
1. InterfacesClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteSender(*http.Request) (InterfacesDeleteFuture, error)
1. WebCategoriesClient.Get(context.Context, string, string) (AzureWebCategory, error)
1. WebCategoriesClient.GetPreparer(context.Context, string, string) (*http.Request, error)
1. WebCategoriesClient.GetResponder(*http.Response) (AzureWebCategory, error)
1. WebCategoriesClient.GetSender(*http.Request) (*http.Response, error)

### Struct Changes

#### Removed Structs

1. BastionHostsDeleteFuture
1. InterfacesDeleteFuture

## Additive Changes

### New Funcs

1. *BastionHostsDeleteABCFuture.UnmarshalJSON([]byte) error
1. *InterfacesDeleteABCDFuture.UnmarshalJSON([]byte) error
1. BastionHostsClient.DeleteABC(context.Context, string, string) (BastionHostsDeleteABCFuture, error)
1. BastionHostsClient.DeleteABCPreparer(context.Context, string, string) (*http.Request, error)
1. BastionHostsClient.DeleteABCResponder(*http.Response) (autorest.Response, error)
1. BastionHostsClient.DeleteABCSender(*http.Request) (BastionHostsDeleteABCFuture, error)
1. InterfacesClient.DeleteABCD(context.Context, string, string) (InterfacesDeleteABCDFuture, error)
1. InterfacesClient.DeleteABCDPreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteABCDResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteABCDSender(*http.Request) (InterfacesDeleteABCDFuture, error)
1. WebCategoriesClient.GetABC(context.Context, string, string) (AzureWebCategory, error)
1. WebCategoriesClient.GetABCPreparer(context.Context, string, string) (*http.Request, error)
1. WebCategoriesClient.GetABCResponder(*http.Response) (AzureWebCategory, error)
1. WebCategoriesClient.GetABCSender(*http.Request) (*http.Response, error)

### Struct Changes

#### New Structs

1. BastionHostsDeleteABCFuture
1. InterfacesDeleteABCDFuture
