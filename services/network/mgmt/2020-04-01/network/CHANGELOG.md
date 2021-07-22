# Unreleased

## Breaking Changes

### Removed Funcs

1. *ApplicationGatewaysDeleteFuture.UnmarshalJSON([]byte) error
1. *ApplicationSecurityGroupsDeleteFuture.UnmarshalJSON([]byte) error
1. *InterfacesDeleteFuture.UnmarshalJSON([]byte) error
1. ApplicationGatewaysClient.Delete(context.Context, string, string) (ApplicationGatewaysDeleteFuture, error)
1. ApplicationGatewaysClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. ApplicationGatewaysClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. ApplicationGatewaysClient.DeleteSender(*http.Request) (ApplicationGatewaysDeleteFuture, error)
1. ApplicationSecurityGroupsClient.Delete(context.Context, string, string) (ApplicationSecurityGroupsDeleteFuture, error)
1. ApplicationSecurityGroupsClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. ApplicationSecurityGroupsClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. ApplicationSecurityGroupsClient.DeleteSender(*http.Request) (ApplicationSecurityGroupsDeleteFuture, error)
1. AvailableServiceAliasesClient.List(context.Context, string) (AvailableServiceAliasesResultPage, error)
1. AvailableServiceAliasesClient.ListComplete(context.Context, string) (AvailableServiceAliasesResultIterator, error)
1. AvailableServiceAliasesClient.ListPreparer(context.Context, string) (*http.Request, error)
1. AvailableServiceAliasesClient.ListResponder(*http.Response) (AvailableServiceAliasesResult, error)
1. AvailableServiceAliasesClient.ListSender(*http.Request) (*http.Response, error)
1. InterfacesClient.Delete(context.Context, string, string) (InterfacesDeleteFuture, error)
1. InterfacesClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteSender(*http.Request) (InterfacesDeleteFuture, error)

### Struct Changes

#### Removed Structs

1. ApplicationGatewaysDeleteFuture
1. ApplicationSecurityGroupsDeleteFuture
1. InterfacesDeleteFuture

## Additive Changes

### New Funcs

1. *ApplicationGatewaysDeleteABCFuture.UnmarshalJSON([]byte) error
1. *ApplicationSecurityGroupsDeleteABCFuture.UnmarshalJSON([]byte) error
1. *InterfacesDeleteABCDFuture.UnmarshalJSON([]byte) error
1. ApplicationGatewaysClient.DeleteABC(context.Context, string, string) (ApplicationGatewaysDeleteABCFuture, error)
1. ApplicationGatewaysClient.DeleteABCPreparer(context.Context, string, string) (*http.Request, error)
1. ApplicationGatewaysClient.DeleteABCResponder(*http.Response) (autorest.Response, error)
1. ApplicationGatewaysClient.DeleteABCSender(*http.Request) (ApplicationGatewaysDeleteABCFuture, error)
1. ApplicationSecurityGroupsClient.DeleteABC(context.Context, string, string) (ApplicationSecurityGroupsDeleteABCFuture, error)
1. ApplicationSecurityGroupsClient.DeleteABCPreparer(context.Context, string, string) (*http.Request, error)
1. ApplicationSecurityGroupsClient.DeleteABCResponder(*http.Response) (autorest.Response, error)
1. ApplicationSecurityGroupsClient.DeleteABCSender(*http.Request) (ApplicationSecurityGroupsDeleteABCFuture, error)
1. AvailableServiceAliasesClient.ListABC(context.Context, string) (AvailableServiceAliasesResultPage, error)
1. AvailableServiceAliasesClient.ListABCComplete(context.Context, string) (AvailableServiceAliasesResultIterator, error)
1. AvailableServiceAliasesClient.ListABCPreparer(context.Context, string) (*http.Request, error)
1. AvailableServiceAliasesClient.ListABCResponder(*http.Response) (AvailableServiceAliasesResult, error)
1. AvailableServiceAliasesClient.ListABCSender(*http.Request) (*http.Response, error)
1. InterfacesClient.DeleteABCD(context.Context, string, string) (InterfacesDeleteABCDFuture, error)
1. InterfacesClient.DeleteABCDPreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteABCDResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteABCDSender(*http.Request) (InterfacesDeleteABCDFuture, error)

### Struct Changes

#### New Structs

1. ApplicationGatewaysDeleteABCFuture
1. ApplicationSecurityGroupsDeleteABCFuture
1. InterfacesDeleteABCDFuture
