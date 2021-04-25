# Unreleased Content

## Breaking Changes

### Removed Funcs

1. NetworkStatusClient.ListByLocation(context.Context, string, string, string) (NetworkStatusContract, error)
1. NetworkStatusClient.ListByLocationPreparer(context.Context, string, string, string) (*http.Request, error)
1. NetworkStatusClient.ListByLocationResponder(*http.Response) (NetworkStatusContract, error)
1. NetworkStatusClient.ListByLocationSender(*http.Request) (*http.Response, error)

### Struct Changes

#### Removed Structs

1. ListNetworkStatusContractByLocation

#### Removed Struct Fields

1. OperationResultContract.ActionLog
1. OperationResultContract.Error
1. OperationResultContract.ResultInfo
1. OperationResultContract.Started
1. OperationResultContract.Status
1. OperationResultContract.Updated
1. TenantConfigurationSyncStateContract.Branch
1. TenantConfigurationSyncStateContract.CommitID
1. TenantConfigurationSyncStateContract.ConfigurationChangeDate
1. TenantConfigurationSyncStateContract.IsExport
1. TenantConfigurationSyncStateContract.IsGitEnabled
1. TenantConfigurationSyncStateContract.IsSynced
1. TenantConfigurationSyncStateContract.SyncDate

### Signature Changes

#### Funcs

1. NetworkStatusClient.ListByService
	- Returns
		- From: ListNetworkStatusContractByLocation, error
		- To: autorest.Response, error
1. NetworkStatusClient.ListByServiceResponder
	- Returns
		- From: ListNetworkStatusContractByLocation, error
		- To: autorest.Response, error

#### Struct Fields

1. NetworkStatusContract.DNSServers changed type from *[]string to interface{}

## Additive Changes

### New Funcs

1. *OperationResultContract.UnmarshalJSON([]byte) error
1. *TenantConfigurationSyncStateContract.UnmarshalJSON([]byte) error
1. NetworkStatusClient.ListByLocationABC(context.Context, string, string, string) (NetworkStatusContract, error)
1. NetworkStatusClient.ListByLocationABCPreparer(context.Context, string, string, string) (*http.Request, error)
1. NetworkStatusClient.ListByLocationABCResponder(*http.Response) (NetworkStatusContract, error)
1. NetworkStatusClient.ListByLocationABCSender(*http.Request) (*http.Response, error)
1. OperationResultContractProperties.MarshalJSON() ([]byte, error)
1. TenantConfigurationSyncStateContract.MarshalJSON() ([]byte, error)

### Struct Changes

#### New Structs

1. OperationResultContractProperties
1. TenantConfigurationSyncStateContractProperties

#### New Struct Fields

1. OperationResultContract.*OperationResultContractProperties
1. OperationResultContract.Name
1. OperationResultContract.Type
1. TenantConfigurationSyncStateContract.*TenantConfigurationSyncStateContractProperties
