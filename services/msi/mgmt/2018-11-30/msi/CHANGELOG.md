# Unreleased Content

## Breaking Changes

### Removed Funcs

1. UserAssignedIdentitiesClient.ListByResourceGroup(context.Context, string) (UserAssignedIdentitiesListResultPage, error)
1. UserAssignedIdentitiesClient.ListByResourceGroupComplete(context.Context, string) (UserAssignedIdentitiesListResultIterator, error)
1. UserAssignedIdentitiesClient.ListByResourceGroupPreparer(context.Context, string) (*http.Request, error)
1. UserAssignedIdentitiesClient.ListByResourceGroupResponder(*http.Response) (UserAssignedIdentitiesListResult, error)
1. UserAssignedIdentitiesClient.ListByResourceGroupSender(*http.Request) (*http.Response, error)
