# Unreleased

## Breaking Changes

### Removed Funcs

1. MarketplaceAgreementsClient.Create(context.Context, *AgreementResource) (AgreementResource, error)
1. MarketplaceAgreementsClient.CreatePreparer(context.Context, *AgreementResource) (*http.Request, error)
1. MarketplaceAgreementsClient.CreateResponder(*http.Response) (AgreementResource, error)
1. MarketplaceAgreementsClient.CreateSender(*http.Request) (*http.Response, error)

### Struct Changes

#### Removed Struct Fields

1. AgreementResource.autorest.Response

## Additive Changes

### New Funcs

1. OfferDetail.MarshalJSON() ([]byte, error)
1. OperationListResult.MarshalJSON() ([]byte, error)
1. OrganizationResourcePropertiesOfferDetail.MarshalJSON() ([]byte, error)
