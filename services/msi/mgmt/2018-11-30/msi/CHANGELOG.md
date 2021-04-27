# Unreleased Content

## Breaking Changes

### Removed Funcs

1. NewSystemAssignedIdentitiesClient(string) SystemAssignedIdentitiesClient
1. NewSystemAssignedIdentitiesClientWithBaseURI(string, string) SystemAssignedIdentitiesClient
1. SystemAssignedIdentitiesClient.GetByScope(context.Context, string) (SystemAssignedIdentity, error)
1. SystemAssignedIdentitiesClient.GetByScopePreparer(context.Context, string) (*http.Request, error)
1. SystemAssignedIdentitiesClient.GetByScopeResponder(*http.Response) (SystemAssignedIdentity, error)
1. SystemAssignedIdentitiesClient.GetByScopeSender(*http.Request) (*http.Response, error)

### Struct Changes

#### Removed Structs

1. SystemAssignedIdentitiesClient

#### Removed Struct Fields

1. SystemAssignedIdentity.autorest.Response
