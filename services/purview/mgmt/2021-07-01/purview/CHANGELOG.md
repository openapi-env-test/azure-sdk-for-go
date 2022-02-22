# Unreleased

## Breaking Changes

### Removed Funcs

1. AccountsClient.Get(context.Context, string, string) (Account, error)
1. AccountsClient.GetPreparer(context.Context, string, string) (*http.Request, error)
1. AccountsClient.GetResponder(*http.Response) (Account, error)
1. AccountsClient.GetSender(*http.Request) (*http.Response, error)

## Additive Changes

### New Constants

1. Type.TypeNone
1. Type.TypeUserAssigned

### New Funcs

1. UserAssignedIdentity.MarshalJSON() ([]byte, error)

### Struct Changes

#### New Structs

1. AccountSkuModel
1. UserAssignedIdentity

#### New Struct Fields

1. AccountUpdateParameters.Identity
1. Identity.UserAssignedIdentities
