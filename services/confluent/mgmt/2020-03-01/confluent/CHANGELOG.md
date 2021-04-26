# Unreleased Content

## Breaking Changes

### Removed Funcs

1. *AgreementResourceListResponseIterator.Next() error
1. *AgreementResourceListResponseIterator.NextWithContext(context.Context) error
1. *AgreementResourceListResponsePage.Next() error
1. *AgreementResourceListResponsePage.NextWithContext(context.Context) error
1. *OperationListResultIterator.Next() error
1. *OperationListResultIterator.NextWithContext(context.Context) error
1. *OperationListResultPage.Next() error
1. *OperationListResultPage.NextWithContext(context.Context) error
1. AgreementResourceListResponse.IsEmpty() bool
1. AgreementResourceListResponseIterator.NotDone() bool
1. AgreementResourceListResponseIterator.Response() AgreementResourceListResponse
1. AgreementResourceListResponseIterator.Value() AgreementResource
1. AgreementResourceListResponsePage.NotDone() bool
1. AgreementResourceListResponsePage.Response() AgreementResourceListResponse
1. AgreementResourceListResponsePage.Values() []AgreementResource
1. MarketplaceAgreementsClient.List(context.Context) (AgreementResourceListResponsePage, error)
1. MarketplaceAgreementsClient.ListComplete(context.Context) (AgreementResourceListResponseIterator, error)
1. MarketplaceAgreementsClient.ListPreparer(context.Context) (*http.Request, error)
1. MarketplaceAgreementsClient.ListResponder(*http.Response) (AgreementResourceListResponse, error)
1. MarketplaceAgreementsClient.ListSender(*http.Request) (*http.Response, error)
1. NewAgreementResourceListResponseIterator(AgreementResourceListResponsePage) AgreementResourceListResponseIterator
1. NewAgreementResourceListResponsePage(AgreementResourceListResponse, func(context.Context, AgreementResourceListResponse) (AgreementResourceListResponse, error)) AgreementResourceListResponsePage
1. NewOperationListResultIterator(OperationListResultPage) OperationListResultIterator
1. NewOperationListResultPage(OperationListResult, func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage
1. NewOrganizationOperationsClient(string) OrganizationOperationsClient
1. NewOrganizationOperationsClientWithBaseURI(string, string) OrganizationOperationsClient
1. OperationListResult.IsEmpty() bool
1. OperationListResultIterator.NotDone() bool
1. OperationListResultIterator.Response() OperationListResult
1. OperationListResultIterator.Value() OperationResult
1. OperationListResultPage.NotDone() bool
1. OperationListResultPage.Response() OperationListResult
1. OperationListResultPage.Values() []OperationResult
1. OrganizationOperationsClient.List(context.Context) (OperationListResultPage, error)
1. OrganizationOperationsClient.ListComplete(context.Context) (OperationListResultIterator, error)
1. OrganizationOperationsClient.ListPreparer(context.Context) (*http.Request, error)
1. OrganizationOperationsClient.ListResponder(*http.Response) (OperationListResult, error)
1. OrganizationOperationsClient.ListSender(*http.Request) (*http.Response, error)

### Struct Changes

#### Removed Structs

1. AgreementResourceListResponseIterator
1. AgreementResourceListResponsePage
1. OperationListResultIterator
1. OperationListResultPage
1. OrganizationOperationsClient

#### Removed Struct Fields

1. AgreementResourceListResponse.autorest.Response
1. OperationListResult.autorest.Response

## Additive Changes

### New Funcs

1. OfferDetail.MarshalJSON() ([]byte, error)
1. OperationListResult.MarshalJSON() ([]byte, error)
1. OrganizationResourcePropertiesOfferDetail.MarshalJSON() ([]byte, error)
