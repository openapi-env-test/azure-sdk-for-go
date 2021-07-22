# Unreleased

## Breaking Changes

### Removed Funcs

1. *ApplicationGatewaysDeleteFuture.UnmarshalJSON([]byte) error
1. *InterfacesDeleteFuture.UnmarshalJSON([]byte) error
1. ApplicationGatewaysClient.Delete(context.Context, string, string) (ApplicationGatewaysDeleteFuture, error)
1. ApplicationGatewaysClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. ApplicationGatewaysClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. ApplicationGatewaysClient.DeleteSender(*http.Request) (ApplicationGatewaysDeleteFuture, error)
1. InterfacesClient.Delete(context.Context, string, string) (InterfacesDeleteFuture, error)
1. InterfacesClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteSender(*http.Request) (InterfacesDeleteFuture, error)

### Struct Changes

#### Removed Structs

1. ApplicationGatewaysDeleteFuture
1. InterfacesDeleteFuture

## Additive Changes

### New Funcs

1. *ApplicationGatewaysDeleteABCFuture.UnmarshalJSON([]byte) error
1. *InterfacesDeleteABCDFuture.UnmarshalJSON([]byte) error
1. ApplicationGatewaysClient.DeleteABC(context.Context, string, string) (ApplicationGatewaysDeleteABCFuture, error)
1. ApplicationGatewaysClient.DeleteABCPreparer(context.Context, string, string) (*http.Request, error)
1. ApplicationGatewaysClient.DeleteABCResponder(*http.Response) (autorest.Response, error)
1. ApplicationGatewaysClient.DeleteABCSender(*http.Request) (ApplicationGatewaysDeleteABCFuture, error)
1. InterfacesClient.DeleteABCD(context.Context, string, string) (InterfacesDeleteABCDFuture, error)
1. InterfacesClient.DeleteABCDPreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteABCDResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteABCDSender(*http.Request) (InterfacesDeleteABCDFuture, error)

### Struct Changes

#### New Structs

1. ApplicationGatewaysDeleteABCFuture
1. InterfacesDeleteABCDFuture
