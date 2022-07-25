# Unreleased

## Breaking Changes

### Signature Changes

#### Funcs

1. ClusterVersionsClient.GetByEnvironment
	- Params
		- From: context.Context, string, string, string
		- To: context.Context, string, ClusterVersionsEnvironment, string
1. ClusterVersionsClient.GetByEnvironmentPreparer
	- Params
		- From: context.Context, string, string, string
		- To: context.Context, string, ClusterVersionsEnvironment, string
1. ClusterVersionsClient.ListByEnvironment
	- Params
		- From: context.Context, string, string
		- To: context.Context, string, ClusterVersionsEnvironment
1. ClusterVersionsClient.ListByEnvironmentPreparer
	- Params
		- From: context.Context, string, string
		- To: context.Context, string, ClusterVersionsEnvironment

## Additive Changes

### New Constants

1. ClusterVersionsEnvironment.ClusterVersionsEnvironmentLinux
1. ClusterVersionsEnvironment.ClusterVersionsEnvironmentWindows

### New Funcs

1. PossibleClusterVersionsEnvironmentValues() []ClusterVersionsEnvironment
