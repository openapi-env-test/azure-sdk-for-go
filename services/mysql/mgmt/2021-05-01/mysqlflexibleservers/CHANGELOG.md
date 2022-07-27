# Unreleased

## Additive Changes

### New Constants

1. DataEncryptionType.DataEncryptionTypeAzureKeyVault
1. DataEncryptionType.DataEncryptionTypeSystemManaged
1. ManagedServiceIdentityType.ManagedServiceIdentityTypeUserAssigned

### New Funcs

1. CheckNameAvailabilityWithoutLocationClient.Execute(context.Context, NameAvailabilityRequest) (NameAvailability, error)
1. CheckNameAvailabilityWithoutLocationClient.ExecutePreparer(context.Context, NameAvailabilityRequest) (*http.Request, error)
1. CheckNameAvailabilityWithoutLocationClient.ExecuteResponder(*http.Response) (NameAvailability, error)
1. CheckNameAvailabilityWithoutLocationClient.ExecuteSender(*http.Request) (*http.Response, error)
1. Identity.MarshalJSON() ([]byte, error)
1. NewCheckNameAvailabilityWithoutLocationClient(string) CheckNameAvailabilityWithoutLocationClient
1. NewCheckNameAvailabilityWithoutLocationClientWithBaseURI(string, string) CheckNameAvailabilityWithoutLocationClient
1. PossibleDataEncryptionTypeValues() []DataEncryptionType
1. PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType
1. UserAssignedIdentity.MarshalJSON() ([]byte, error)

### Struct Changes

#### New Structs

1. CheckNameAvailabilityWithoutLocationClient
1. DataEncryption
1. Identity
1. UserAssignedIdentity

#### New Struct Fields

1. Server.Identity
1. ServerForUpdate.Identity
1. ServerProperties.DataEncryption
1. ServerPropertiesForUpdate.DataEncryption
1. VirtualNetworkSubnetUsageResult.Location
1. VirtualNetworkSubnetUsageResult.SubscriptionID
