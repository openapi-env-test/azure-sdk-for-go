# Unreleased

## Breaking Changes

### Removed Funcs

1. AvailabilitySetsClient.CreateOrUpdate(context.Context, string, string, AvailabilitySet) (AvailabilitySet, error)
1. AvailabilitySetsClient.CreateOrUpdatePreparer(context.Context, string, string, AvailabilitySet) (*http.Request, error)
1. AvailabilitySetsClient.CreateOrUpdateResponder(*http.Response) (AvailabilitySet, error)
1. AvailabilitySetsClient.CreateOrUpdateSender(*http.Request) (*http.Response, error)

### Signature Changes

#### Struct Fields

1. OrchestrationServiceStateInput.ServiceName changed type from *string to OrchestrationServiceNames

## Additive Changes

### New Funcs

1. AvailabilitySetsClient.CreateOrUpdateA(context.Context, string, string, AvailabilitySet) (AvailabilitySet, error)
1. AvailabilitySetsClient.CreateOrUpdateAPreparer(context.Context, string, string, AvailabilitySet) (*http.Request, error)
1. AvailabilitySetsClient.CreateOrUpdateAResponder(*http.Response) (AvailabilitySet, error)
1. AvailabilitySetsClient.CreateOrUpdateASender(*http.Request) (*http.Response, error)
