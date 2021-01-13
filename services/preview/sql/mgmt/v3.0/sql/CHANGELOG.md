Generated from https://github.com/Azure/azure-rest-api-specs/tree/../../../../../azure-rest-api-specs/specification/sql/resource-manager/readme.md tag: `package-composite-v3`

Code generator 


## Breaking Changes

### Removed Constants

1. PrimaryAggregationType.Average
1. PrimaryAggregationType.Maximum
1. PrimaryAggregationType.Minimum
1. PrimaryAggregationType.Total
1. StorageKeyType.SharedAccessKey
1. StorageKeyType.StorageAccessKey

### Removed Funcs

1. *ImportExportOperationResult.UnmarshalJSON([]byte) error
1. *ManagedDatabaseRestoreDetailsResult.UnmarshalJSON([]byte) error
1. DatabasesClient.Export(context.Context, string, string, string, ExportDatabaseDefinition) (DatabasesExportFuture, error)
1. DatabasesClient.ExportPreparer(context.Context, string, string, string, ExportDatabaseDefinition) (*http.Request, error)
1. DatabasesClient.ExportResponder(*http.Response) (ImportExportOperationResult, error)
1. DatabasesClient.ExportSender(*http.Request) (DatabasesExportFuture, error)
1. ImportExportClient.Import(context.Context, string, string, string, ImportExistingDatabaseDefinition) (ImportExportImportFuture, error)
1. ImportExportClient.ImportPreparer(context.Context, string, string, string, ImportExistingDatabaseDefinition) (*http.Request, error)
1. ImportExportClient.ImportResponder(*http.Response) (ImportExportOperationResult, error)
1. ImportExportClient.ImportSender(*http.Request) (ImportExportImportFuture, error)
1. ImportExportOperationResult.MarshalJSON() ([]byte, error)
1. ManagedDatabaseRestoreDetailsClient.Get(context.Context, string, string, string) (ManagedDatabaseRestoreDetailsResult, error)
1. ManagedDatabaseRestoreDetailsClient.GetPreparer(context.Context, string, string, string) (*http.Request, error)
1. ManagedDatabaseRestoreDetailsClient.GetResponder(*http.Response) (ManagedDatabaseRestoreDetailsResult, error)
1. ManagedDatabaseRestoreDetailsClient.GetSender(*http.Request) (*http.Response, error)
1. ManagedDatabaseRestoreDetailsResult.MarshalJSON() ([]byte, error)
1. NewImportExportClient(string) ImportExportClient
1. NewImportExportClientWithBaseURI(string, string) ImportExportClient
1. NewManagedDatabaseRestoreDetailsClient(string) ManagedDatabaseRestoreDetailsClient
1. NewManagedDatabaseRestoreDetailsClientWithBaseURI(string, string) ManagedDatabaseRestoreDetailsClient
1. PossibleStorageKeyTypeValues() []StorageKeyType
1. ServersClient.ImportDatabase(context.Context, string, string, ImportNewDatabaseDefinition) (ServersImportDatabaseFuture, error)
1. ServersClient.ImportDatabasePreparer(context.Context, string, string, ImportNewDatabaseDefinition) (*http.Request, error)
1. ServersClient.ImportDatabaseResponder(*http.Response) (ImportExportOperationResult, error)
1. ServersClient.ImportDatabaseSender(*http.Request) (ServersImportDatabaseFuture, error)

## Struct Changes

### Removed Structs

1. DatabasesExportFuture
1. ExportDatabaseDefinition
1. ImportExistingDatabaseDefinition
1. ImportExportClient
1. ImportExportImportFuture
1. ImportExportOperationResult
1. ImportExportOperationResultProperties
1. ImportNewDatabaseDefinition
1. ManagedDatabaseRestoreDetailsClient
1. ManagedDatabaseRestoreDetailsProperties
1. ManagedDatabaseRestoreDetailsResult
1. NetworkIsolationSettings
1. PrivateEndpointConnectionRequestStatus
1. ServersImportDatabaseFuture

## Signature Changes

### Const Types

1. Count changed type from PrimaryAggregationType to QueryMetricUnitType
1. None changed type from PrimaryAggregationType to IdentityType

### New Constants

1. AggregationFunctionType.Avg
1. AggregationFunctionType.Max
1. AggregationFunctionType.Min
1. AggregationFunctionType.Stdev
1. AggregationFunctionType.Sum
1. IdentityType.UserAssigned
1. MetricType.CPU
1. MetricType.Dtu
1. MetricType.Duration
1. MetricType.Io
1. MetricType.LogIo
1. PrimaryAggregationType.PrimaryAggregationTypeAverage
1. PrimaryAggregationType.PrimaryAggregationTypeCount
1. PrimaryAggregationType.PrimaryAggregationTypeMaximum
1. PrimaryAggregationType.PrimaryAggregationTypeMinimum
1. PrimaryAggregationType.PrimaryAggregationTypeNone
1. PrimaryAggregationType.PrimaryAggregationTypeTotal
1. QueryMetricUnitType.KB
1. QueryMetricUnitType.Microseconds
1. QueryMetricUnitType.Percentage
1. QueryTimeGrainType.P1D
1. QueryTimeGrainType.PT1H

### New Funcs

1. *TopQueriesListResultIterator.Next() error
1. *TopQueriesListResultIterator.NextWithContext(context.Context) error
1. *TopQueriesListResultPage.Next() error
1. *TopQueriesListResultPage.NextWithContext(context.Context) error
1. ManagedInstancePrivateEndpointConnectionProperties.MarshalJSON() ([]byte, error)
1. ManagedInstancesClient.ListByManagedInstance(context.Context, string, string, *int32, string, string, string, QueryTimeGrainType, AggregationFunctionType, MetricType) (TopQueriesListResultPage, error)
1. ManagedInstancesClient.ListByManagedInstanceComplete(context.Context, string, string, *int32, string, string, string, QueryTimeGrainType, AggregationFunctionType, MetricType) (TopQueriesListResultIterator, error)
1. ManagedInstancesClient.ListByManagedInstancePreparer(context.Context, string, string, *int32, string, string, string, QueryTimeGrainType, AggregationFunctionType, MetricType) (*http.Request, error)
1. ManagedInstancesClient.ListByManagedInstanceResponder(*http.Response) (TopQueriesListResult, error)
1. ManagedInstancesClient.ListByManagedInstanceSender(*http.Request) (*http.Response, error)
1. NewTopQueriesListResultIterator(TopQueriesListResultPage) TopQueriesListResultIterator
1. NewTopQueriesListResultPage(TopQueriesListResult, func(context.Context, TopQueriesListResult) (TopQueriesListResult, error)) TopQueriesListResultPage
1. PossibleAggregationFunctionTypeValues() []AggregationFunctionType
1. PossibleMetricTypeValues() []MetricType
1. PossibleQueryMetricUnitTypeValues() []QueryMetricUnitType
1. PossibleQueryTimeGrainTypeValues() []QueryTimeGrainType
1. QueryMetricInterval.MarshalJSON() ([]byte, error)
1. QueryStatisticsProperties.MarshalJSON() ([]byte, error)
1. TopQueries.MarshalJSON() ([]byte, error)
1. TopQueriesListResult.IsEmpty() bool
1. TopQueriesListResultIterator.NotDone() bool
1. TopQueriesListResultIterator.Response() TopQueriesListResult
1. TopQueriesListResultIterator.Value() TopQueries
1. TopQueriesListResultPage.NotDone() bool
1. TopQueriesListResultPage.Response() TopQueriesListResult
1. TopQueriesListResultPage.Values() []TopQueries

## Struct Changes

### New Structs

1. ManagedInstancePecProperty
1. ManagedInstancePrivateEndpointConnectionProperties
1. ManagedInstancePrivateEndpointProperty
1. QueryMetricInterval
1. QueryMetricProperties
1. QueryStatisticsProperties
1. TopQueries
1. TopQueriesListResult
1. TopQueriesListResultIterator
1. TopQueriesListResultPage

### New Struct Fields

1. ManagedInstanceProperties.PrivateEndpointConnections
1. ManagedInstanceProperties.ZoneRedundant
1. ManagedInstanceUpdate.Identity
