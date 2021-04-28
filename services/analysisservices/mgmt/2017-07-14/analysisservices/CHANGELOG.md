# Unreleased Content

## Breaking Changes

### Removed Constants

1. Status.Live

### Removed Funcs

1. PossibleStatusValues() []Status

### Struct Changes

#### Removed Struct Fields

1. ErrorResponse.Code
1. ErrorResponse.Message

### Signature Changes

#### Struct Fields

1. GatewayListStatusLive.Status changed type from Status to *int32

## Additive Changes

### New Funcs

1. *OperationListResultIterator.Next() error
1. *OperationListResultIterator.NextWithContext(context.Context) error
1. *OperationListResultPage.Next() error
1. *OperationListResultPage.NextWithContext(context.Context) error
1. ErrorResponseError.MarshalJSON() ([]byte, error)
1. NewOperationListResultIterator(OperationListResultPage) OperationListResultIterator
1. NewOperationListResultPage(OperationListResult, func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage
1. NewOperationsClient(string) OperationsClient
1. NewOperationsClientWithBaseURI(string, string) OperationsClient
1. OperationDetail.MarshalJSON() ([]byte, error)
1. OperationDisplay.MarshalJSON() ([]byte, error)
1. OperationListResult.IsEmpty() bool
1. OperationListResultIterator.NotDone() bool
1. OperationListResultIterator.Response() OperationListResult
1. OperationListResultIterator.Value() OperationDetail
1. OperationListResultPage.NotDone() bool
1. OperationListResultPage.Response() OperationListResult
1. OperationListResultPage.Values() []OperationDetail
1. OperationsClient.List(context.Context) (OperationListResultPage, error)
1. OperationsClient.ListComplete(context.Context) (OperationListResultIterator, error)
1. OperationsClient.ListPreparer(context.Context) (*http.Request, error)
1. OperationsClient.ListResponder(*http.Response) (OperationListResult, error)
1. OperationsClient.ListSender(*http.Request) (*http.Response, error)

### Struct Changes

#### New Structs

1. ErrorAdditionalInfo
1. ErrorDetail
1. ErrorResponseError
1. LogSpecifications
1. MetricDimensions
1. MetricSpecifications
1. OperationDetail
1. OperationDetailProperties
1. OperationDetailPropertiesServiceSpecification
1. OperationDisplay
1. OperationListResult
1. OperationListResultIterator
1. OperationListResultPage
1. OperationsClient
1. OperationsErrorResponse

#### New Struct Fields

1. ErrorResponse.Error
1. ServerMutableProperties.ManagedMode
1. ServerMutableProperties.ServerMonitorMode
1. ServerProperties.ManagedMode
1. ServerProperties.ServerMonitorMode
1. SkuDetailsForExistingResource.ResourceType
