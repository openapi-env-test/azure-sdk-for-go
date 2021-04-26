# Unreleased Content

## Breaking Changes

### Removed Constants

1. Origin.System
1. Origin.Usersystem

### Struct Changes

#### Removed Structs

1. ErrorDefinition

#### Removed Struct Fields

1. ErrorResponse.AdditionalInfo
1. ErrorResponse.Code
1. ErrorResponse.Details
1. ErrorResponse.Message
1. ErrorResponse.Target

### Signature Changes

#### Const Types

1. User changed type from Origin to CreatedByType

## Additive Changes

### New Constants

1. CheckNameAvailabilityReason.AlreadyExists
1. CheckNameAvailabilityReason.Invalid
1. CreatedByType.Application
1. CreatedByType.Key
1. CreatedByType.ManagedIdentity
1. Origin.OriginSystem
1. Origin.OriginUser
1. Origin.OriginUsersystem
1. ResourceIdentityType.None
1. ResourceIdentityType.SystemAssigned

### New Funcs

1. BaseClient.CheckNameAvailability(context.Context, CheckNameAvailabilityRequest) (CheckNameAvailabilityResponse, error)
1. BaseClient.CheckNameAvailabilityPreparer(context.Context, CheckNameAvailabilityRequest) (*http.Request, error)
1. BaseClient.CheckNameAvailabilityResponder(*http.Response) (CheckNameAvailabilityResponse, error)
1. BaseClient.CheckNameAvailabilitySender(*http.Request) (*http.Response, error)
1. Identity.MarshalJSON() ([]byte, error)
1. PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason
1. PossibleCreatedByTypeValues() []CreatedByType
1. PossibleResourceIdentityTypeValues() []ResourceIdentityType

### Struct Changes

#### New Structs

1. CheckNameAvailabilityRequest
1. CheckNameAvailabilityResponse
1. ErrorDetail
1. Identity
1. SystemData

#### New Struct Fields

1. Account.Identity
1. Account.SystemData
1. AccountUpdate.Identity
1. ErrorResponse.Error
1. Instance.SystemData
