# Unreleased

## Breaking Changes

### Removed Funcs

1. AvailabilitySetsClient.CreateOrUpdate(context.Context, string, string, AvailabilitySet) (AvailabilitySet, error)
1. AvailabilitySetsClient.CreateOrUpdatePreparer(context.Context, string, string, AvailabilitySet) (*http.Request, error)
1. AvailabilitySetsClient.CreateOrUpdateResponder(*http.Response) (AvailabilitySet, error)
1. AvailabilitySetsClient.CreateOrUpdateSender(*http.Request) (*http.Response, error)

### Struct Changes

#### Removed Structs

1. RestorePointProvisioningDetails

#### Removed Struct Fields

1. RestorePointProperties.ProvisioningDetails

## Additive Changes

### Struct Changes

#### New Struct Fields

1. RestorePointProperties.TimeCreated
