# Unreleased

## Breaking Changes

### Removed Funcs

1. AvailabilitySetsClient.Update(context.Context, string, string, AvailabilitySetUpdate) (AvailabilitySet, error)
1. AvailabilitySetsClient.UpdatePreparer(context.Context, string, string, AvailabilitySetUpdate) (*http.Request, error)
1. AvailabilitySetsClient.UpdateResponder(*http.Response) (AvailabilitySet, error)
1. AvailabilitySetsClient.UpdateSender(*http.Request) (*http.Response, error)
