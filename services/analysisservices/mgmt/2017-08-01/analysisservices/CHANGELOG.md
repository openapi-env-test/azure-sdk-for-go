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
1. OperationStatus.Error changed type from *ErrorResponse to *GatewayError

## Additive Changes

### Struct Changes

#### New Structs

1. ErrorAdditionalInfo
1. LogSpecifications
1. MetricDimensions
1. MetricSpecifications
1. OperationProperties
1. OperationPropertiesServiceSpecification

#### New Struct Fields

1. ErrorResponse.Error
1. GatewayError.AdditionalInfo
1. GatewayError.Details
1. GatewayError.HTTPStatusCode
1. GatewayError.SubCode
1. GatewayError.Target
1. GatewayError.TimeStamp
1. Operation.Origin
1. Operation.Properties
1. OperationDisplay.Description
1. ServerMutableProperties.ManagedMode
1. ServerMutableProperties.ServerMonitorMode
1. ServerProperties.ManagedMode
1. ServerProperties.ServerMonitorMode
1. ServerProperties.Sku
1. SkuDetailsForExistingResource.ResourceType
