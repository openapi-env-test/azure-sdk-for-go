# Unreleased

## Breaking Changes

### Struct Changes

#### Removed Struct Fields

1. BaseClient.SubscriptionID2

### Signature Changes

#### Funcs

1. New
	- Params
		- From: string
		- To: string, string
1. NewOperationsClient
	- Params
		- From: string
		- To: string, string
1. NewOperationsClientWithBaseURI
	- Params
		- From: string, string
		- To: string, string, string
1. NewPolicyEventsClient
	- Params
		- From: string
		- To: string, string
1. NewPolicyEventsClientWithBaseURI
	- Params
		- From: string, string
		- To: string, string, string
1. NewPolicyMetadataClient
	- Params
		- From: string
		- To: string, string
1. NewPolicyMetadataClientWithBaseURI
	- Params
		- From: string, string
		- To: string, string, string
1. NewPolicyRestrictionsClient
	- Params
		- From: string
		- To: string, string
1. NewPolicyRestrictionsClientWithBaseURI
	- Params
		- From: string, string
		- To: string, string, string
1. NewPolicyStatesClient
	- Params
		- From: string
		- To: string, string
1. NewPolicyStatesClientWithBaseURI
	- Params
		- From: string, string
		- To: string, string, string
1. NewPolicyTrackedResourcesClient
	- Params
		- From: string
		- To: string, string
1. NewPolicyTrackedResourcesClientWithBaseURI
	- Params
		- From: string, string
		- To: string, string, string
1. NewRemediationsClient
	- Params
		- From: string
		- To: string, string
1. NewRemediationsClientWithBaseURI
	- Params
		- From: string, string
		- To: string, string, string
1. NewWithBaseURI
	- Params
		- From: string, string
		- To: string, string, string
1. PolicyEventsClient.ListQueryResultsForPolicyDefinition
	- Params
		- From: context.Context, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyEventsClient.ListQueryResultsForPolicyDefinitionComplete
	- Params
		- From: context.Context, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyEventsClient.ListQueryResultsForPolicyDefinitionPreparer
	- Params
		- From: context.Context, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyEventsClient.ListQueryResultsForPolicySetDefinition
	- Params
		- From: context.Context, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyEventsClient.ListQueryResultsForPolicySetDefinitionComplete
	- Params
		- From: context.Context, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyEventsClient.ListQueryResultsForPolicySetDefinitionPreparer
	- Params
		- From: context.Context, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyEventsClient.ListQueryResultsForResourceGroup
	- Params
		- From: context.Context, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyEventsClient.ListQueryResultsForResourceGroupComplete
	- Params
		- From: context.Context, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyEventsClient.ListQueryResultsForResourceGroupLevelPolicyAssignment
	- Params
		- From: context.Context, string, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyEventsClient.ListQueryResultsForResourceGroupLevelPolicyAssignmentComplete
	- Params
		- From: context.Context, string, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyEventsClient.ListQueryResultsForResourceGroupLevelPolicyAssignmentPreparer
	- Params
		- From: context.Context, string, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyEventsClient.ListQueryResultsForResourceGroupPreparer
	- Params
		- From: context.Context, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyEventsClient.ListQueryResultsForSubscription
	- Params
		- From: context.Context, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyEventsClient.ListQueryResultsForSubscriptionComplete
	- Params
		- From: context.Context, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyEventsClient.ListQueryResultsForSubscriptionLevelPolicyAssignment
	- Params
		- From: context.Context, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyEventsClient.ListQueryResultsForSubscriptionLevelPolicyAssignmentComplete
	- Params
		- From: context.Context, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyEventsClient.ListQueryResultsForSubscriptionLevelPolicyAssignmentPreparer
	- Params
		- From: context.Context, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyEventsClient.ListQueryResultsForSubscriptionPreparer
	- Params
		- From: context.Context, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyRestrictionsClient.CheckAtResourceGroupScope
	- Params
		- From: context.Context, string, string, CheckRestrictionsRequest
		- To: context.Context, string, CheckRestrictionsRequest
1. PolicyRestrictionsClient.CheckAtResourceGroupScopePreparer
	- Params
		- From: context.Context, string, string, CheckRestrictionsRequest
		- To: context.Context, string, CheckRestrictionsRequest
1. PolicyRestrictionsClient.CheckAtSubscriptionScope
	- Params
		- From: context.Context, string, CheckRestrictionsRequest
		- To: context.Context, CheckRestrictionsRequest
1. PolicyRestrictionsClient.CheckAtSubscriptionScopePreparer
	- Params
		- From: context.Context, string, CheckRestrictionsRequest
		- To: context.Context, CheckRestrictionsRequest
1. PolicyStatesClient.ListQueryResultsForPolicyDefinition
	- Params
		- From: context.Context, PolicyStatesResource, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.ListQueryResultsForPolicyDefinitionComplete
	- Params
		- From: context.Context, PolicyStatesResource, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.ListQueryResultsForPolicyDefinitionPreparer
	- Params
		- From: context.Context, PolicyStatesResource, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.ListQueryResultsForPolicySetDefinition
	- Params
		- From: context.Context, PolicyStatesResource, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.ListQueryResultsForPolicySetDefinitionComplete
	- Params
		- From: context.Context, PolicyStatesResource, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.ListQueryResultsForPolicySetDefinitionPreparer
	- Params
		- From: context.Context, PolicyStatesResource, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.ListQueryResultsForResourceGroup
	- Params
		- From: context.Context, PolicyStatesResource, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.ListQueryResultsForResourceGroupComplete
	- Params
		- From: context.Context, PolicyStatesResource, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.ListQueryResultsForResourceGroupLevelPolicyAssignment
	- Params
		- From: context.Context, PolicyStatesResource, string, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.ListQueryResultsForResourceGroupLevelPolicyAssignmentComplete
	- Params
		- From: context.Context, PolicyStatesResource, string, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.ListQueryResultsForResourceGroupLevelPolicyAssignmentPreparer
	- Params
		- From: context.Context, PolicyStatesResource, string, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.ListQueryResultsForResourceGroupPreparer
	- Params
		- From: context.Context, PolicyStatesResource, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.ListQueryResultsForSubscription
	- Params
		- From: context.Context, PolicyStatesResource, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.ListQueryResultsForSubscriptionComplete
	- Params
		- From: context.Context, PolicyStatesResource, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.ListQueryResultsForSubscriptionLevelPolicyAssignment
	- Params
		- From: context.Context, PolicyStatesResource, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.ListQueryResultsForSubscriptionLevelPolicyAssignmentComplete
	- Params
		- From: context.Context, PolicyStatesResource, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.ListQueryResultsForSubscriptionLevelPolicyAssignmentPreparer
	- Params
		- From: context.Context, PolicyStatesResource, string, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, string, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.ListQueryResultsForSubscriptionPreparer
	- Params
		- From: context.Context, PolicyStatesResource, string, *int32, string, string, *date.Time, *date.Time, string, string, string
		- To: context.Context, PolicyStatesResource, *int32, string, string, *date.Time, *date.Time, string, string, string
1. PolicyStatesClient.SummarizeForPolicyDefinition
	- Params
		- From: context.Context, string, string, *int32, *date.Time, *date.Time, string
		- To: context.Context, string, *int32, *date.Time, *date.Time, string
1. PolicyStatesClient.SummarizeForPolicyDefinitionPreparer
	- Params
		- From: context.Context, string, string, *int32, *date.Time, *date.Time, string
		- To: context.Context, string, *int32, *date.Time, *date.Time, string
1. PolicyStatesClient.SummarizeForPolicySetDefinition
	- Params
		- From: context.Context, string, string, *int32, *date.Time, *date.Time, string
		- To: context.Context, string, *int32, *date.Time, *date.Time, string
1. PolicyStatesClient.SummarizeForPolicySetDefinitionPreparer
	- Params
		- From: context.Context, string, string, *int32, *date.Time, *date.Time, string
		- To: context.Context, string, *int32, *date.Time, *date.Time, string
1. PolicyStatesClient.SummarizeForResourceGroup
	- Params
		- From: context.Context, string, string, *int32, *date.Time, *date.Time, string
		- To: context.Context, string, *int32, *date.Time, *date.Time, string
1. PolicyStatesClient.SummarizeForResourceGroupLevelPolicyAssignment
	- Params
		- From: context.Context, string, string, string, *int32, *date.Time, *date.Time, string
		- To: context.Context, string, string, *int32, *date.Time, *date.Time, string
1. PolicyStatesClient.SummarizeForResourceGroupLevelPolicyAssignmentPreparer
	- Params
		- From: context.Context, string, string, string, *int32, *date.Time, *date.Time, string
		- To: context.Context, string, string, *int32, *date.Time, *date.Time, string
1. PolicyStatesClient.SummarizeForResourceGroupPreparer
	- Params
		- From: context.Context, string, string, *int32, *date.Time, *date.Time, string
		- To: context.Context, string, *int32, *date.Time, *date.Time, string
1. PolicyStatesClient.SummarizeForSubscription
	- Params
		- From: context.Context, string, *int32, *date.Time, *date.Time, string
		- To: context.Context, *int32, *date.Time, *date.Time, string
1. PolicyStatesClient.SummarizeForSubscriptionLevelPolicyAssignment
	- Params
		- From: context.Context, string, string, *int32, *date.Time, *date.Time, string
		- To: context.Context, string, *int32, *date.Time, *date.Time, string
1. PolicyStatesClient.SummarizeForSubscriptionLevelPolicyAssignmentPreparer
	- Params
		- From: context.Context, string, string, *int32, *date.Time, *date.Time, string
		- To: context.Context, string, *int32, *date.Time, *date.Time, string
1. PolicyStatesClient.SummarizeForSubscriptionPreparer
	- Params
		- From: context.Context, string, *int32, *date.Time, *date.Time, string
		- To: context.Context, *int32, *date.Time, *date.Time, string
1. PolicyStatesClient.TriggerResourceGroupEvaluation
	- Params
		- From: context.Context, string, string
		- To: context.Context, string
1. PolicyStatesClient.TriggerResourceGroupEvaluationPreparer
	- Params
		- From: context.Context, string, string
		- To: context.Context, string
1. PolicyStatesClient.TriggerSubscriptionEvaluation
	- Params
		- From: context.Context, string
		- To: context.Context
1. PolicyStatesClient.TriggerSubscriptionEvaluationPreparer
	- Params
		- From: context.Context, string
		- To: context.Context
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
1. RemediationsClient.CancelAtResourceGroup
	- Params
		- From: context.Context, string, string, string
		- To: context.Context, string, string
1. RemediationsClient.CancelAtResourceGroupPreparer
	- Params
		- From: context.Context, string, string, string
		- To: context.Context, string, string
1. RemediationsClient.CancelAtSubscription
	- Params
		- From: context.Context, string, string
		- To: context.Context, string
1. RemediationsClient.CancelAtSubscriptionPreparer
	- Params
		- From: context.Context, string, string
		- To: context.Context, string
1. RemediationsClient.CreateOrUpdateAtResourceGroup
	- Params
		- From: context.Context, string, string, string, Remediation
		- To: context.Context, string, string, Remediation
1. RemediationsClient.CreateOrUpdateAtResourceGroupPreparer
	- Params
		- From: context.Context, string, string, string, Remediation
		- To: context.Context, string, string, Remediation
1. RemediationsClient.CreateOrUpdateAtSubscription
	- Params
		- From: context.Context, string, string, Remediation
		- To: context.Context, string, Remediation
1. RemediationsClient.CreateOrUpdateAtSubscriptionPreparer
	- Params
		- From: context.Context, string, string, Remediation
		- To: context.Context, string, Remediation
1. RemediationsClient.DeleteAtResourceGroup
	- Params
		- From: context.Context, string, string, string
		- To: context.Context, string, string
1. RemediationsClient.DeleteAtResourceGroupPreparer
	- Params
		- From: context.Context, string, string, string
		- To: context.Context, string, string
1. RemediationsClient.DeleteAtSubscription
	- Params
		- From: context.Context, string, string
		- To: context.Context, string
1. RemediationsClient.DeleteAtSubscriptionPreparer
	- Params
		- From: context.Context, string, string
		- To: context.Context, string
1. RemediationsClient.GetAtResourceGroup
	- Params
		- From: context.Context, string, string, string
		- To: context.Context, string, string
1. RemediationsClient.GetAtResourceGroupPreparer
	- Params
		- From: context.Context, string, string, string
		- To: context.Context, string, string
1. RemediationsClient.GetAtSubscription
	- Params
		- From: context.Context, string, string
		- To: context.Context, string
1. RemediationsClient.GetAtSubscriptionPreparer
	- Params
		- From: context.Context, string, string
		- To: context.Context, string
1. RemediationsClient.ListDeploymentsAtResourceGroup
	- Params
		- From: context.Context, string, string, string, *int32
		- To: context.Context, string, string, *int32
1. RemediationsClient.ListDeploymentsAtResourceGroupComplete
	- Params
		- From: context.Context, string, string, string, *int32
		- To: context.Context, string, string, *int32
1. RemediationsClient.ListDeploymentsAtResourceGroupPreparer
	- Params
		- From: context.Context, string, string, string, *int32
		- To: context.Context, string, string, *int32
1. RemediationsClient.ListDeploymentsAtSubscription
	- Params
		- From: context.Context, string, string, *int32
		- To: context.Context, string, *int32
1. RemediationsClient.ListDeploymentsAtSubscriptionComplete
	- Params
		- From: context.Context, string, string, *int32
		- To: context.Context, string, *int32
1. RemediationsClient.ListDeploymentsAtSubscriptionPreparer
	- Params
		- From: context.Context, string, string, *int32
		- To: context.Context, string, *int32
1. RemediationsClient.ListForResourceGroup
	- Params
		- From: context.Context, string, string, *int32, string
		- To: context.Context, string, *int32, string
1. RemediationsClient.ListForResourceGroupComplete
	- Params
		- From: context.Context, string, string, *int32, string
		- To: context.Context, string, *int32, string
1. RemediationsClient.ListForResourceGroupPreparer
	- Params
		- From: context.Context, string, string, *int32, string
		- To: context.Context, string, *int32, string
1. RemediationsClient.ListForSubscription
	- Params
		- From: context.Context, string, *int32, string
		- To: context.Context, *int32, string
1. RemediationsClient.ListForSubscriptionComplete
	- Params
		- From: context.Context, string, *int32, string
		- To: context.Context, *int32, string
1. RemediationsClient.ListForSubscriptionPreparer
	- Params
		- From: context.Context, string, *int32, string
		- To: context.Context, *int32, string

## Additive Changes

### Struct Changes

#### New Struct Fields

1. BaseClient.SubscriptionID
1. BaseClient.SubscriptionID1
