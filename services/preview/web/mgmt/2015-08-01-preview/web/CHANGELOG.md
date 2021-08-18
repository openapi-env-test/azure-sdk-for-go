# Unreleased

## Breaking Changes

### Signature Changes

#### Funcs

1. ManagedApisClient.Get
	- Params
		- From: context.Context, string, string, *bool
		- To: context.Context, string, string, string, *bool
1. ManagedApisClient.GetPreparer
	- Params
		- From: context.Context, string, string, *bool
		- To: context.Context, string, string, string, *bool
