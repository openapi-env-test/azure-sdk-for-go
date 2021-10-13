# Unreleased

## Breaking Changes

### Removed Constants

1. ServiceTypeBasicServiceResourceProperties.ServiceTypeBasicServiceResourcePropertiesServiceTypeDataTransferServiceResourceProperties
1. ServiceTypeBasicServiceResourceProperties.ServiceTypeBasicServiceResourcePropertiesServiceTypeSQLDedicatedGatewayServiceResourceProperties

### Signature Changes

#### Funcs

1. ServiceClient.Create
	- Params
		- From: context.Context, string, string, string, ServiceResource
		- To: context.Context, string, string, string, ServiceResourceCreateUpdateParameters
1. ServiceClient.CreatePreparer
	- Params
		- From: context.Context, string, string, string, ServiceResource
		- To: context.Context, string, string, string, ServiceResourceCreateUpdateParameters

## Additive Changes

### New Constants

1. ServiceTypeBasicServiceResourceProperties.ServiceTypeBasicServiceResourcePropertiesServiceTypeDataTransfer
1. ServiceTypeBasicServiceResourceProperties.ServiceTypeBasicServiceResourcePropertiesServiceTypeSQLDedicatedGateway

### New Funcs

1. *ServiceResourceCreateUpdateParameters.UnmarshalJSON([]byte) error
1. DatabaseAccountsClient.Test(context.Context) (DatabaseAccountsListResult, error)
1. DatabaseAccountsClient.TestPreparer(context.Context) (*http.Request, error)
1. DatabaseAccountsClient.TestResponder(*http.Response) (DatabaseAccountsListResult, error)
1. DatabaseAccountsClient.TestSender(*http.Request) (*http.Response, error)
1. ServiceResourceCreateUpdateParameters.MarshalJSON() ([]byte, error)

### Struct Changes

#### New Structs

1. ServiceResourceCreateUpdateParameters
1. ServiceResourceCreateUpdateProperties
