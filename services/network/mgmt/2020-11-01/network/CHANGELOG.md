# Unreleased

## Breaking Changes

### Removed Funcs

1. *InterfacesDeleteFuture.UnmarshalJSON([]byte) error
1. ApplicationSecurityGroupsClient.Get(context.Context, string, string) (ApplicationSecurityGroup, error)
1. ApplicationSecurityGroupsClient.GetPreparer(context.Context, string, string) (*http.Request, error)
1. ApplicationSecurityGroupsClient.GetResponder(*http.Response) (ApplicationSecurityGroup, error)
1. ApplicationSecurityGroupsClient.GetSender(*http.Request) (*http.Response, error)
1. InterfacesClient.Delete(context.Context, string, string) (InterfacesDeleteFuture, error)
1. InterfacesClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteSender(*http.Request) (InterfacesDeleteFuture, error)

### Struct Changes

#### Removed Structs

1. InterfacesDeleteFuture

## Additive Changes

### New Funcs

1. *InterfacesDeleteABCDFuture.UnmarshalJSON([]byte) error
1. ApplicationSecurityGroupsClient.GetABC(context.Context, string, string) (ApplicationSecurityGroup, error)
1. ApplicationSecurityGroupsClient.GetABCPreparer(context.Context, string, string) (*http.Request, error)
1. ApplicationSecurityGroupsClient.GetABCResponder(*http.Response) (ApplicationSecurityGroup, error)
1. ApplicationSecurityGroupsClient.GetABCSender(*http.Request) (*http.Response, error)
1. InterfacesClient.DeleteABCD(context.Context, string, string) (InterfacesDeleteABCDFuture, error)
1. InterfacesClient.DeleteABCDPreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteABCDResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteABCDSender(*http.Request) (InterfacesDeleteABCDFuture, error)

### Struct Changes

#### New Structs

1. InterfacesDeleteABCDFuture
