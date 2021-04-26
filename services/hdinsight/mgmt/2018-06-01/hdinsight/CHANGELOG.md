# Unreleased Content

## Breaking Changes

### Struct Changes

#### Removed Structs

1. OperationResource

#### Removed Struct Fields

1. CapabilitiesResult.VMSizeFilters
1. CapabilitiesResult.VMSizes
1. Extension.autorest.Response
1. VMSizeCompatibilityFilter.Vmsizes

### Signature Changes

#### Funcs

1. ExtensionsClient.Get
	- Returns
		- From: Extension, error
		- To: ClusterMonitoringResponse, error
1. ExtensionsClient.GetResponder
	- Returns
		- From: Extension, error
		- To: ClusterMonitoringResponse, error

#### Struct Fields

1. Usage.CurrentValue changed type from *int32 to *int64
1. Usage.Limit changed type from *int32 to *int64
1. VersionSpec.IsDefault changed type from *string to *bool

## Additive Changes

### New Funcs

1. *ClustersUpdateIdentityCertificateFuture.UnmarshalJSON([]byte) error
1. ApplicationGetHTTPSEndpoint.MarshalJSON() ([]byte, error)
1. ApplicationsClient.GetAzureAsyncOperationStatus(context.Context, string, string, string, string) (AsyncOperationResult, error)
1. ApplicationsClient.GetAzureAsyncOperationStatusPreparer(context.Context, string, string, string, string) (*http.Request, error)
1. ApplicationsClient.GetAzureAsyncOperationStatusResponder(*http.Response) (AsyncOperationResult, error)
1. ApplicationsClient.GetAzureAsyncOperationStatusSender(*http.Request) (*http.Response, error)
1. ClusterCreateRequestValidationParameters.MarshalJSON() ([]byte, error)
1. ClustersClient.GetAzureAsyncOperationStatus(context.Context, string, string, string) (AsyncOperationResult, error)
1. ClustersClient.GetAzureAsyncOperationStatusPreparer(context.Context, string, string, string) (*http.Request, error)
1. ClustersClient.GetAzureAsyncOperationStatusResponder(*http.Response) (AsyncOperationResult, error)
1. ClustersClient.GetAzureAsyncOperationStatusSender(*http.Request) (*http.Response, error)
1. ClustersClient.UpdateIdentityCertificate(context.Context, string, string, UpdateClusterIdentityCertificateParameters) (ClustersUpdateIdentityCertificateFuture, error)
1. ClustersClient.UpdateIdentityCertificatePreparer(context.Context, string, string, UpdateClusterIdentityCertificateParameters) (*http.Request, error)
1. ClustersClient.UpdateIdentityCertificateResponder(*http.Response) (autorest.Response, error)
1. ClustersClient.UpdateIdentityCertificateSender(*http.Request) (ClustersUpdateIdentityCertificateFuture, error)
1. ExtensionsClient.GetAzureAsyncOperationStatus(context.Context, string, string, string, string) (AsyncOperationResult, error)
1. ExtensionsClient.GetAzureAsyncOperationStatusPreparer(context.Context, string, string, string, string) (*http.Request, error)
1. ExtensionsClient.GetAzureAsyncOperationStatusResponder(*http.Response) (AsyncOperationResult, error)
1. ExtensionsClient.GetAzureAsyncOperationStatusSender(*http.Request) (*http.Response, error)
1. KafkaRestProperties.MarshalJSON() ([]byte, error)
1. LocationsClient.CheckNameAvailability(context.Context, string, NameAvailabilityCheckRequestParameters) (NameAvailabilityCheckResult, error)
1. LocationsClient.CheckNameAvailabilityPreparer(context.Context, string, NameAvailabilityCheckRequestParameters) (*http.Request, error)
1. LocationsClient.CheckNameAvailabilityResponder(*http.Response) (NameAvailabilityCheckResult, error)
1. LocationsClient.CheckNameAvailabilitySender(*http.Request) (*http.Response, error)
1. LocationsClient.GetAzureAsyncOperationStatus(context.Context, string, string) (AsyncOperationResult, error)
1. LocationsClient.GetAzureAsyncOperationStatusPreparer(context.Context, string, string) (*http.Request, error)
1. LocationsClient.GetAzureAsyncOperationStatusResponder(*http.Response) (AsyncOperationResult, error)
1. LocationsClient.GetAzureAsyncOperationStatusSender(*http.Request) (*http.Response, error)
1. LocationsClient.ValidateClusterCreateRequest(context.Context, string, ClusterCreateRequestValidationParameters) (ClusterCreateValidationResult, error)
1. LocationsClient.ValidateClusterCreateRequestPreparer(context.Context, string, ClusterCreateRequestValidationParameters) (*http.Request, error)
1. LocationsClient.ValidateClusterCreateRequestResponder(*http.Response) (ClusterCreateValidationResult, error)
1. LocationsClient.ValidateClusterCreateRequestSender(*http.Request) (*http.Response, error)
1. NameAvailabilityCheckResult.MarshalJSON() ([]byte, error)
1. ScriptActionsClient.GetExecutionAsyncOperationStatus(context.Context, string, string, string) (AsyncOperationResult, error)
1. ScriptActionsClient.GetExecutionAsyncOperationStatusPreparer(context.Context, string, string, string) (*http.Request, error)
1. ScriptActionsClient.GetExecutionAsyncOperationStatusResponder(*http.Response) (AsyncOperationResult, error)
1. ScriptActionsClient.GetExecutionAsyncOperationStatusSender(*http.Request) (*http.Response, error)
1. VirtualMachinesClient.GetAsyncOperationStatus(context.Context, string, string, string) (AsyncOperationResult, error)
1. VirtualMachinesClient.GetAsyncOperationStatusPreparer(context.Context, string, string, string) (*http.Request, error)
1. VirtualMachinesClient.GetAsyncOperationStatusResponder(*http.Response) (AsyncOperationResult, error)
1. VirtualMachinesClient.GetAsyncOperationStatusSender(*http.Request) (*http.Response, error)

### Struct Changes

#### New Structs

1. AaddsResourceDetails
1. AsyncOperationResult
1. ClusterCreateRequestValidationParameters
1. ClusterCreateValidationResult
1. ClustersUpdateIdentityCertificateFuture
1. NameAvailabilityCheckRequestParameters
1. NameAvailabilityCheckResult
1. UpdateClusterIdentityCertificateParameters
1. ValidationErrorInfo

#### New Struct Fields

1. ApplicationGetHTTPSEndpoint.PrivateIPAddress
1. CapabilitiesResult.VmsizeFilters
1. CapabilitiesResult.Vmsizes
1. KafkaRestProperties.ConfigurationOverride
1. Role.VMGroupName
1. StorageAccount.Fileshare
1. StorageAccount.Saskey
1. VMSizeCompatibilityFilter.ComputeIsolationSupported
1. VMSizeCompatibilityFilter.ESPApplied
1. VMSizeCompatibilityFilter.OsType
1. VMSizeCompatibilityFilter.VMSizes
