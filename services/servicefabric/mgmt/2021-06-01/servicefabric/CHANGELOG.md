# Unreleased

## Breaking Changes

### Removed Funcs

1. ClustersClient.ListByResourceGroup(context.Context, string) (ClusterListResult, error)
1. ClustersClient.ListByResourceGroupPreparer(context.Context, string) (*http.Request, error)
1. ClustersClient.ListByResourceGroupResponder(*http.Response) (ClusterListResult, error)
1. ClustersClient.ListByResourceGroupSender(*http.Request) (*http.Response, error)

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
