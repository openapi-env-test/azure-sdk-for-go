# Unreleased Content

## Breaking Changes

### Removed Funcs

1. *UserAssignedIdentitiesListResultIterator.Next() error
1. *UserAssignedIdentitiesListResultIterator.NextWithContext(context.Context) error
1. *UserAssignedIdentitiesListResultPage.Next() error
1. *UserAssignedIdentitiesListResultPage.NextWithContext(context.Context) error
1. NewSystemAssignedIdentitiesClient(string) SystemAssignedIdentitiesClient
1. NewSystemAssignedIdentitiesClientWithBaseURI(string, string) SystemAssignedIdentitiesClient
1. NewUserAssignedIdentitiesListResultIterator(UserAssignedIdentitiesListResultPage) UserAssignedIdentitiesListResultIterator
1. NewUserAssignedIdentitiesListResultPage(UserAssignedIdentitiesListResult, func(context.Context, UserAssignedIdentitiesListResult) (UserAssignedIdentitiesListResult, error)) UserAssignedIdentitiesListResultPage
1. SystemAssignedIdentitiesClient.GetByScope(context.Context, string) (SystemAssignedIdentity, error)
1. SystemAssignedIdentitiesClient.GetByScopePreparer(context.Context, string) (*http.Request, error)
1. SystemAssignedIdentitiesClient.GetByScopeResponder(*http.Response) (SystemAssignedIdentity, error)
1. SystemAssignedIdentitiesClient.GetByScopeSender(*http.Request) (*http.Response, error)
1. UserAssignedIdentitiesClient.ListByResourceGroup(context.Context, string) (UserAssignedIdentitiesListResultPage, error)
1. UserAssignedIdentitiesClient.ListByResourceGroupComplete(context.Context, string) (UserAssignedIdentitiesListResultIterator, error)
1. UserAssignedIdentitiesClient.ListByResourceGroupPreparer(context.Context, string) (*http.Request, error)
1. UserAssignedIdentitiesClient.ListByResourceGroupResponder(*http.Response) (UserAssignedIdentitiesListResult, error)
1. UserAssignedIdentitiesClient.ListByResourceGroupSender(*http.Request) (*http.Response, error)
1. UserAssignedIdentitiesClient.ListBySubscription(context.Context) (UserAssignedIdentitiesListResultPage, error)
1. UserAssignedIdentitiesClient.ListBySubscriptionComplete(context.Context) (UserAssignedIdentitiesListResultIterator, error)
1. UserAssignedIdentitiesClient.ListBySubscriptionPreparer(context.Context) (*http.Request, error)
1. UserAssignedIdentitiesClient.ListBySubscriptionResponder(*http.Response) (UserAssignedIdentitiesListResult, error)
1. UserAssignedIdentitiesClient.ListBySubscriptionSender(*http.Request) (*http.Response, error)
1. UserAssignedIdentitiesListResult.IsEmpty() bool
1. UserAssignedIdentitiesListResultIterator.NotDone() bool
1. UserAssignedIdentitiesListResultIterator.Response() UserAssignedIdentitiesListResult
1. UserAssignedIdentitiesListResultIterator.Value() Identity
1. UserAssignedIdentitiesListResultPage.NotDone() bool
1. UserAssignedIdentitiesListResultPage.Response() UserAssignedIdentitiesListResult
1. UserAssignedIdentitiesListResultPage.Values() []Identity

### Struct Changes

#### Removed Structs

1. SystemAssignedIdentitiesClient
1. UserAssignedIdentitiesListResultIterator
1. UserAssignedIdentitiesListResultPage

#### Removed Struct Fields

1. SystemAssignedIdentity.autorest.Response
1. UserAssignedIdentitiesListResult.autorest.Response
