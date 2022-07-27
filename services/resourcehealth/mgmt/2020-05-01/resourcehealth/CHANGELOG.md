# Unreleased

## Breaking Changes

### Removed Funcs

1. ErrorResponseError.MarshalJSON() ([]byte, error)

### Struct Changes

#### Removed Structs

1. ErrorResponseError

#### Removed Struct Fields

1. AvailabilityStatusProperties.OccurredTime
1. AvailabilityStatusPropertiesRecentlyResolved.UnavailabilitySummary
1. AvailabilityStatusPropertiesRecentlyResolved.UnavailableOccurredTime
1. ErrorResponse.Error

## Additive Changes

### New Funcs

1. ErrorResponse.MarshalJSON() ([]byte, error)

### Struct Changes

#### New Struct Fields

1. AvailabilityStatusProperties.OccuredTime
1. AvailabilityStatusPropertiesRecentlyResolved.UnavailableOccuredTime
1. AvailabilityStatusPropertiesRecentlyResolved.UnavailableSummary
1. ErrorResponse.Code
1. ErrorResponse.Details
1. ErrorResponse.Message
