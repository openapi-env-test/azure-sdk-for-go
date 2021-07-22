# Unreleased

## Breaking Changes

### Removed Funcs

1. *ApplicationSecurityGroupsDeleteFuture.UnmarshalJSON([]byte) error
1. *InterfacesDeleteFuture.UnmarshalJSON([]byte) error
1. ApplicationSecurityGroupsClient.Delete(context.Context, string, string) (ApplicationSecurityGroupsDeleteFuture, error)
1. ApplicationSecurityGroupsClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. ApplicationSecurityGroupsClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. ApplicationSecurityGroupsClient.DeleteSender(*http.Request) (ApplicationSecurityGroupsDeleteFuture, error)
1. InterfacesClient.Delete(context.Context, string, string) (InterfacesDeleteFuture, error)
1. InterfacesClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteSender(*http.Request) (InterfacesDeleteFuture, error)

### Struct Changes

#### Removed Structs

1. ApplicationSecurityGroupsDeleteFuture
1. InterfacesDeleteFuture

## Additive Changes

### New Funcs

1. *ApplicationSecurityGroupsDeleteABCFuture.UnmarshalJSON([]byte) error
1. *InterfacesDeleteABCDFuture.UnmarshalJSON([]byte) error
1. ApplicationSecurityGroupsClient.DeleteABC(context.Context, string, string) (ApplicationSecurityGroupsDeleteABCFuture, error)
1. ApplicationSecurityGroupsClient.DeleteABCPreparer(context.Context, string, string) (*http.Request, error)
1. ApplicationSecurityGroupsClient.DeleteABCResponder(*http.Response) (autorest.Response, error)
1. ApplicationSecurityGroupsClient.DeleteABCSender(*http.Request) (ApplicationSecurityGroupsDeleteABCFuture, error)
1. InterfacesClient.DeleteABCD(context.Context, string, string) (InterfacesDeleteABCDFuture, error)
1. InterfacesClient.DeleteABCDPreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteABCDResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteABCDSender(*http.Request) (InterfacesDeleteABCDFuture, error)

### Struct Changes

#### New Structs

1. ApplicationSecurityGroupsDeleteABCFuture
1. InterfacesDeleteABCDFuture
