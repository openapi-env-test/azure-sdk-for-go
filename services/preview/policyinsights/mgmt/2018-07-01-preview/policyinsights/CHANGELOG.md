# Unreleased

## Breaking Changes

### Signature Changes

#### Funcs

1. New
	- Params
		- From: <none>
		- To: string
1. NewOperationsClient
	- Params
		- From: <none>
		- To: string
1. NewOperationsClientWithBaseURI
	- Params
		- From: string
		- To: string, string
1. NewPolicyEventsClient
	- Params
		- From: <none>
		- To: string
1. NewPolicyEventsClientWithBaseURI
	- Params
		- From: string
		- To: string, string
1. NewPolicyStatesClient
	- Params
		- From: <none>
		- To: string
1. NewPolicyStatesClientWithBaseURI
	- Params
		- From: string
		- To: string, string
1. NewPolicyTrackedResourcesClient
	- Params
		- From: <none>
		- To: string
1. NewPolicyTrackedResourcesClientWithBaseURI
	- Params
		- From: string
		- To: string, string
1. NewRemediationsClient
	- Params
		- From: <none>
		- To: string
1. NewRemediationsClientWithBaseURI
	- Params
		- From: string
		- To: string, string
1. NewWithBaseURI
	- Params
		- From: string
		- To: string, string
1. PolicyTrackedResourcesClient.ListQueryResultsForResourceGroup
	- Params
		- From: context.Context, string, string, *int32, string
		- To: context.Context, string, *int32, string
1. PolicyTrackedResourcesClient.ListQueryResultsForResourceGroupComplete
	- Params
		- From: context.Context, string, string, *int32, string
		- To: context.Context, string, *int32, string
1. PolicyTrackedResourcesClient.ListQueryResultsForResourceGroupPreparer
	- Params
		- From: context.Context, string, string, *int32, string
		- To: context.Context, string, *int32, string
1. PolicyTrackedResourcesClient.ListQueryResultsForSubscription
	- Params
		- From: context.Context, string, *int32, string
		- To: context.Context, *int32, string
1. PolicyTrackedResourcesClient.ListQueryResultsForSubscriptionComplete
	- Params
		- From: context.Context, string, *int32, string
		- To: context.Context, *int32, string
1. PolicyTrackedResourcesClient.ListQueryResultsForSubscriptionPreparer
	- Params
		- From: context.Context, string, *int32, string
		- To: context.Context, *int32, string

## Additive Changes

### Struct Changes

#### New Struct Fields

1. BaseClient.SubscriptionID
