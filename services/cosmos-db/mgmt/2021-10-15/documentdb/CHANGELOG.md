# Unreleased

## Breaking Changes

### Signature Changes

#### Struct Fields

1. RestorableMongodbResourcesListResult.Value changed type from *[]DatabaseRestoreResource to *[]RestorableMongodbResourcesGetResult
1. RestorableSQLResourcesListResult.Value changed type from *[]DatabaseRestoreResource to *[]RestorableSQLResourcesGetResult

## Additive Changes

### New Funcs

1. RestorableMongodbResourcesGetResult.MarshalJSON() ([]byte, error)
1. RestorableSQLResourcesGetResult.MarshalJSON() ([]byte, error)

### Struct Changes

#### New Structs

1. RestorableMongodbResourcesGetResult
1. RestorableSQLResourcesGetResult
