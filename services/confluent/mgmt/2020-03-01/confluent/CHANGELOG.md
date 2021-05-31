# Unreleased Content

## Breaking Changes

### Removed Funcs

1. *OrganizationDeleteFuture.UnmarshalJSON([]byte) error
1. OrganizationClient.Delete(context.Context, string, string) (OrganizationDeleteFuture, error)
1. OrganizationClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. OrganizationClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. OrganizationClient.DeleteSender(*http.Request) (OrganizationDeleteFuture, error)

### Struct Changes

#### Removed Structs

1. OrganizationDeleteFuture

## Additive Changes

### New Funcs

1. OfferDetail.MarshalJSON() ([]byte, error)
1. OperationListResult.MarshalJSON() ([]byte, error)
1. OrganizationResourcePropertiesOfferDetail.MarshalJSON() ([]byte, error)
