# Unreleased

## Breaking Changes

### Removed Funcs

1. *ExpressRouteCircuitAuthorizationsDeleteFuture.UnmarshalJSON([]byte) error
1. *InterfacesDeleteFuture.UnmarshalJSON([]byte) error
1. *LoadBalancersDeleteFuture.UnmarshalJSON([]byte) error
1. *SecurityGroupsDeleteFuture.UnmarshalJSON([]byte) error
1. BaseClient.CheckDNSNameAvailability(context.Context, string, string) (DNSNameAvailabilityResult, error)
1. BaseClient.CheckDNSNameAvailabilityPreparer(context.Context, string, string) (*http.Request, error)
1. BaseClient.CheckDNSNameAvailabilityResponder(*http.Response) (DNSNameAvailabilityResult, error)
1. BaseClient.CheckDNSNameAvailabilitySender(*http.Request) (*http.Response, error)
1. ExpressRouteCircuitAuthorizationsClient.Delete(context.Context, string, string, string) (ExpressRouteCircuitAuthorizationsDeleteFuture, error)
1. ExpressRouteCircuitAuthorizationsClient.DeletePreparer(context.Context, string, string, string) (*http.Request, error)
1. ExpressRouteCircuitAuthorizationsClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. ExpressRouteCircuitAuthorizationsClient.DeleteSender(*http.Request) (ExpressRouteCircuitAuthorizationsDeleteFuture, error)
1. InterfacesClient.Delete(context.Context, string, string) (InterfacesDeleteFuture, error)
1. InterfacesClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteSender(*http.Request) (InterfacesDeleteFuture, error)
1. LoadBalancersClient.Delete(context.Context, string, string) (LoadBalancersDeleteFuture, error)
1. LoadBalancersClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. LoadBalancersClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. LoadBalancersClient.DeleteSender(*http.Request) (LoadBalancersDeleteFuture, error)
1. SecurityGroupsClient.Delete(context.Context, string, string) (SecurityGroupsDeleteFuture, error)
1. SecurityGroupsClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. SecurityGroupsClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. SecurityGroupsClient.DeleteSender(*http.Request) (SecurityGroupsDeleteFuture, error)

### Struct Changes

#### Removed Structs

1. ExpressRouteCircuitAuthorizationsDeleteFuture
1. InterfacesDeleteFuture
1. LoadBalancersDeleteFuture
1. SecurityGroupsDeleteFuture

## Additive Changes

### New Funcs

1. *ABC1Future.UnmarshalJSON([]byte) error
1. *ABCdFuture.UnmarshalJSON([]byte) error
1. *ABCeFuture.UnmarshalJSON([]byte) error
1. *LoadBalancersDeleteABCFuture.UnmarshalJSON([]byte) error
1. BaseClient.ABC(context.Context, string, string) (DNSNameAvailabilityResult, error)
1. BaseClient.ABC1(context.Context, string, string, string) (ABC1Future, error)
1. BaseClient.ABC1Preparer(context.Context, string, string, string) (*http.Request, error)
1. BaseClient.ABC1Responder(*http.Response) (autorest.Response, error)
1. BaseClient.ABC1Sender(*http.Request) (ABC1Future, error)
1. BaseClient.ABCPreparer(context.Context, string, string) (*http.Request, error)
1. BaseClient.ABCResponder(*http.Response) (DNSNameAvailabilityResult, error)
1. BaseClient.ABCSender(*http.Request) (*http.Response, error)
1. BaseClient.ABCd(context.Context, string, string) (ABCdFuture, error)
1. BaseClient.ABCdPreparer(context.Context, string, string) (*http.Request, error)
1. BaseClient.ABCdResponder(*http.Response) (autorest.Response, error)
1. BaseClient.ABCdSender(*http.Request) (ABCdFuture, error)
1. BaseClient.ABCe(context.Context, string, string) (ABCeFuture, error)
1. BaseClient.ABCePreparer(context.Context, string, string) (*http.Request, error)
1. BaseClient.ABCeResponder(*http.Response) (autorest.Response, error)
1. BaseClient.ABCeSender(*http.Request) (ABCeFuture, error)
1. LoadBalancersClient.DeleteABC(context.Context, string, string) (LoadBalancersDeleteABCFuture, error)
1. LoadBalancersClient.DeleteABCPreparer(context.Context, string, string) (*http.Request, error)
1. LoadBalancersClient.DeleteABCResponder(*http.Response) (autorest.Response, error)
1. LoadBalancersClient.DeleteABCSender(*http.Request) (LoadBalancersDeleteABCFuture, error)

### Struct Changes

#### New Structs

1. ABC1Future
1. ABCdFuture
1. ABCeFuture
1. LoadBalancersDeleteABCFuture
