# Unreleased Content

## Breaking Changes

### Removed Funcs

1. *ConfigurationStoresDeleteFuture.UnmarshalJSON([]byte) error
1. *ConfigurationStoresUpdateFuture.UnmarshalJSON([]byte) error
1. ConfigurationStoresClient.Delete(context.Context, string, string) (ConfigurationStoresDeleteFuture, error)
1. ConfigurationStoresClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. ConfigurationStoresClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. ConfigurationStoresClient.DeleteSender(*http.Request) (ConfigurationStoresDeleteFuture, error)
1. ConfigurationStoresClient.Update(context.Context, string, string, ConfigurationStoreUpdateParameters) (ConfigurationStoresUpdateFuture, error)
1. ConfigurationStoresClient.UpdatePreparer(context.Context, string, string, ConfigurationStoreUpdateParameters) (*http.Request, error)
1. ConfigurationStoresClient.UpdateResponder(*http.Response) (ConfigurationStore, error)
1. ConfigurationStoresClient.UpdateSender(*http.Request) (ConfigurationStoresUpdateFuture, error)

### Struct Changes

#### Removed Structs

1. ConfigurationStoresDeleteFuture
1. ConfigurationStoresUpdateFuture

## Additive Changes

### New Funcs

1. *ConfigurationStoresDeleteABCFuture.UnmarshalJSON([]byte) error
1. ConfigurationStoresClient.DeleteABC(context.Context, string, string) (ConfigurationStoresDeleteABCFuture, error)
1. ConfigurationStoresClient.DeleteABCPreparer(context.Context, string, string) (*http.Request, error)
1. ConfigurationStoresClient.DeleteABCResponder(*http.Response) (autorest.Response, error)
1. ConfigurationStoresClient.DeleteABCSender(*http.Request) (ConfigurationStoresDeleteABCFuture, error)

### Struct Changes

#### New Structs

1. ConfigurationStoresDeleteABCFuture
