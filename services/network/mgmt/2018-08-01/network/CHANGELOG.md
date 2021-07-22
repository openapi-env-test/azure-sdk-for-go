# Unreleased

## Breaking Changes

### Removed Funcs

1. *ApplicationGatewaysDeleteFuture.UnmarshalJSON([]byte) error
1. *DdosProtectionPlansDeleteFuture.UnmarshalJSON([]byte) error
1. *InterfacesDeleteFuture.UnmarshalJSON([]byte) error
1. *ProfilesDeleteFuture.UnmarshalJSON([]byte) error
1. ApplicationGatewaysClient.Delete(context.Context, string, string) (ApplicationGatewaysDeleteFuture, error)
1. ApplicationGatewaysClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. ApplicationGatewaysClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. ApplicationGatewaysClient.DeleteSender(*http.Request) (ApplicationGatewaysDeleteFuture, error)
1. DdosProtectionPlansClient.Delete(context.Context, string, string) (DdosProtectionPlansDeleteFuture, error)
1. DdosProtectionPlansClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. DdosProtectionPlansClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. DdosProtectionPlansClient.DeleteSender(*http.Request) (DdosProtectionPlansDeleteFuture, error)
1. InterfacesClient.Delete(context.Context, string, string) (InterfacesDeleteFuture, error)
1. InterfacesClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteSender(*http.Request) (InterfacesDeleteFuture, error)
1. ProfilesClient.Delete(context.Context, string, string) (ProfilesDeleteFuture, error)
1. ProfilesClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. ProfilesClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. ProfilesClient.DeleteSender(*http.Request) (ProfilesDeleteFuture, error)

### Struct Changes

#### Removed Structs

1. ApplicationGatewaysDeleteFuture
1. DdosProtectionPlansDeleteFuture
1. InterfacesDeleteFuture
1. ProfilesDeleteFuture

## Additive Changes

### New Funcs

1. *ApplicationGatewaysDeleteABCFuture.UnmarshalJSON([]byte) error
1. *DdosProtectionPlansDeleteABCFuture.UnmarshalJSON([]byte) error
1. *InterfacesDeleteABCDFuture.UnmarshalJSON([]byte) error
1. *ProfilesDeleteABCFuture.UnmarshalJSON([]byte) error
1. ApplicationGatewaysClient.DeleteABC(context.Context, string, string) (ApplicationGatewaysDeleteABCFuture, error)
1. ApplicationGatewaysClient.DeleteABCPreparer(context.Context, string, string) (*http.Request, error)
1. ApplicationGatewaysClient.DeleteABCResponder(*http.Response) (autorest.Response, error)
1. ApplicationGatewaysClient.DeleteABCSender(*http.Request) (ApplicationGatewaysDeleteABCFuture, error)
1. DdosProtectionPlansClient.DeleteABC(context.Context, string, string) (DdosProtectionPlansDeleteABCFuture, error)
1. DdosProtectionPlansClient.DeleteABCPreparer(context.Context, string, string) (*http.Request, error)
1. DdosProtectionPlansClient.DeleteABCResponder(*http.Response) (autorest.Response, error)
1. DdosProtectionPlansClient.DeleteABCSender(*http.Request) (DdosProtectionPlansDeleteABCFuture, error)
1. InterfacesClient.DeleteABCD(context.Context, string, string) (InterfacesDeleteABCDFuture, error)
1. InterfacesClient.DeleteABCDPreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteABCDResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteABCDSender(*http.Request) (InterfacesDeleteABCDFuture, error)
1. ProfilesClient.DeleteABC(context.Context, string, string) (ProfilesDeleteABCFuture, error)
1. ProfilesClient.DeleteABCPreparer(context.Context, string, string) (*http.Request, error)
1. ProfilesClient.DeleteABCResponder(*http.Response) (autorest.Response, error)
1. ProfilesClient.DeleteABCSender(*http.Request) (ProfilesDeleteABCFuture, error)

### Struct Changes

#### New Structs

1. ApplicationGatewaysDeleteABCFuture
1. DdosProtectionPlansDeleteABCFuture
1. InterfacesDeleteABCDFuture
1. ProfilesDeleteABCFuture
