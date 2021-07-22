# Unreleased

## Breaking Changes

### Removed Funcs

1. *InterfacesDeleteFuture.UnmarshalJSON([]byte) error
1. *PublicIPAddressesDeleteFuture.UnmarshalJSON([]byte) error
1. InterfacesClient.Delete(context.Context, string, string) (InterfacesDeleteFuture, error)
1. InterfacesClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteSender(*http.Request) (InterfacesDeleteFuture, error)
1. PublicIPAddressesClient.Delete(context.Context, string, string) (PublicIPAddressesDeleteFuture, error)
1. PublicIPAddressesClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. PublicIPAddressesClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. PublicIPAddressesClient.DeleteSender(*http.Request) (PublicIPAddressesDeleteFuture, error)

### Struct Changes

#### Removed Structs

1. InterfacesDeleteFuture
1. PublicIPAddressesDeleteFuture

## Additive Changes

### New Funcs

1. *InterfacesDeleteABCDFuture.UnmarshalJSON([]byte) error
1. *PublicIPAddressesDeleteABCFuture.UnmarshalJSON([]byte) error
1. InterfacesClient.DeleteABCD(context.Context, string, string) (InterfacesDeleteABCDFuture, error)
1. InterfacesClient.DeleteABCDPreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteABCDResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteABCDSender(*http.Request) (InterfacesDeleteABCDFuture, error)
1. PublicIPAddressesClient.DeleteABC(context.Context, string, string) (PublicIPAddressesDeleteABCFuture, error)
1. PublicIPAddressesClient.DeleteABCPreparer(context.Context, string, string) (*http.Request, error)
1. PublicIPAddressesClient.DeleteABCResponder(*http.Response) (autorest.Response, error)
1. PublicIPAddressesClient.DeleteABCSender(*http.Request) (PublicIPAddressesDeleteABCFuture, error)

### Struct Changes

#### New Structs

1. InterfacesDeleteABCDFuture
1. PublicIPAddressesDeleteABCFuture
