package mariadbapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/mariadb/mgmt/2018-06-01/mariadb"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CreateRecommendedActionSession(ctx context.Context, resourceGroupName string, serverName string, advisorName string, databaseName string) (result mariadb.CreateRecommendedActionSessionFuture, err error)
	ResetQueryPerformanceInsightData(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.QueryPerformanceInsightResetDataResult, err error)
}

var _ BaseClientAPI = (*mariadb.BaseClient)(nil)

// ServersClientAPI contains the set of methods on the ServersClient type.
type ServersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, serverName string, parameters mariadb.ServerForCreate) (result mariadb.ServersCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.ServersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.Server, err error)
	List(ctx context.Context) (result mariadb.ServerListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result mariadb.ServerListResult, err error)
	Restart(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.ServersRestartFuture, err error)
	Update(ctx context.Context, resourceGroupName string, serverName string, parameters mariadb.ServerUpdateParameters) (result mariadb.ServersUpdateFuture, err error)
}

var _ ServersClientAPI = (*mariadb.ServersClient)(nil)

// ReplicasClientAPI contains the set of methods on the ReplicasClient type.
type ReplicasClientAPI interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.ServerListResult, err error)
}

var _ ReplicasClientAPI = (*mariadb.ReplicasClient)(nil)

// FirewallRulesClientAPI contains the set of methods on the FirewallRulesClient type.
type FirewallRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters mariadb.FirewallRule) (result mariadb.FirewallRulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (result mariadb.FirewallRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (result mariadb.FirewallRule, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.FirewallRuleListResult, err error)
}

var _ FirewallRulesClientAPI = (*mariadb.FirewallRulesClient)(nil)

// VirtualNetworkRulesClientAPI contains the set of methods on the VirtualNetworkRulesClient type.
type VirtualNetworkRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, parameters mariadb.VirtualNetworkRule) (result mariadb.VirtualNetworkRulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string) (result mariadb.VirtualNetworkRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string) (result mariadb.VirtualNetworkRule, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.VirtualNetworkRuleListResultPage, err error)
	ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.VirtualNetworkRuleListResultIterator, err error)
}

var _ VirtualNetworkRulesClientAPI = (*mariadb.VirtualNetworkRulesClient)(nil)

// DatabasesClientAPI contains the set of methods on the DatabasesClient type.
type DatabasesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters mariadb.Database) (result mariadb.DatabasesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result mariadb.DatabasesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result mariadb.Database, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.DatabaseListResult, err error)
}

var _ DatabasesClientAPI = (*mariadb.DatabasesClient)(nil)

// ConfigurationsClientAPI contains the set of methods on the ConfigurationsClient type.
type ConfigurationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, configurationName string, parameters mariadb.Configuration) (result mariadb.ConfigurationsCreateOrUpdateFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, configurationName string) (result mariadb.Configuration, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.ConfigurationListResult, err error)
}

var _ ConfigurationsClientAPI = (*mariadb.ConfigurationsClient)(nil)

// ServerParametersClientAPI contains the set of methods on the ServerParametersClient type.
type ServerParametersClientAPI interface {
	ListUpdateConfigurations(ctx context.Context, resourceGroupName string, serverName string, value mariadb.ConfigurationListResult) (result mariadb.ServerParametersListUpdateConfigurationsFuture, err error)
}

var _ ServerParametersClientAPI = (*mariadb.ServerParametersClient)(nil)

// LogFilesClientAPI contains the set of methods on the LogFilesClient type.
type LogFilesClientAPI interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.LogFileListResult, err error)
}

var _ LogFilesClientAPI = (*mariadb.LogFilesClient)(nil)

// RecoverableServersClientAPI contains the set of methods on the RecoverableServersClient type.
type RecoverableServersClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.RecoverableServerResource, err error)
}

var _ RecoverableServersClientAPI = (*mariadb.RecoverableServersClient)(nil)

// ServerBasedPerformanceTierClientAPI contains the set of methods on the ServerBasedPerformanceTierClient type.
type ServerBasedPerformanceTierClientAPI interface {
	List(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.PerformanceTierListResult, err error)
}

var _ ServerBasedPerformanceTierClientAPI = (*mariadb.ServerBasedPerformanceTierClient)(nil)

// LocationBasedPerformanceTierClientAPI contains the set of methods on the LocationBasedPerformanceTierClient type.
type LocationBasedPerformanceTierClientAPI interface {
	List(ctx context.Context, locationName string) (result mariadb.PerformanceTierListResult, err error)
}

var _ LocationBasedPerformanceTierClientAPI = (*mariadb.LocationBasedPerformanceTierClient)(nil)

// CheckNameAvailabilityClientAPI contains the set of methods on the CheckNameAvailabilityClient type.
type CheckNameAvailabilityClientAPI interface {
	Execute(ctx context.Context, nameAvailabilityRequest mariadb.NameAvailabilityRequest) (result mariadb.NameAvailability, err error)
}

var _ CheckNameAvailabilityClientAPI = (*mariadb.CheckNameAvailabilityClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result mariadb.OperationListResult, err error)
}

var _ OperationsClientAPI = (*mariadb.OperationsClient)(nil)

// QueryTextsClientAPI contains the set of methods on the QueryTextsClient type.
type QueryTextsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serverName string, queryID string) (result mariadb.QueryText, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string, queryIds []string) (result mariadb.QueryTextsResultListPage, err error)
	ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string, queryIds []string) (result mariadb.QueryTextsResultListIterator, err error)
}

var _ QueryTextsClientAPI = (*mariadb.QueryTextsClient)(nil)

// TopQueryStatisticsClientAPI contains the set of methods on the TopQueryStatisticsClient type.
type TopQueryStatisticsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serverName string, queryStatisticID string) (result mariadb.QueryStatistic, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string, parameters mariadb.TopQueryStatisticsInput) (result mariadb.TopQueryStatisticsResultListPage, err error)
	ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string, parameters mariadb.TopQueryStatisticsInput) (result mariadb.TopQueryStatisticsResultListIterator, err error)
}

var _ TopQueryStatisticsClientAPI = (*mariadb.TopQueryStatisticsClient)(nil)

// WaitStatisticsClientAPI contains the set of methods on the WaitStatisticsClient type.
type WaitStatisticsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serverName string, waitStatisticsID string) (result mariadb.WaitStatistic, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string, parameters mariadb.WaitStatisticsInput) (result mariadb.WaitStatisticsResultListPage, err error)
	ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string, parameters mariadb.WaitStatisticsInput) (result mariadb.WaitStatisticsResultListIterator, err error)
}

var _ WaitStatisticsClientAPI = (*mariadb.WaitStatisticsClient)(nil)

// AdvisorsClientAPI contains the set of methods on the AdvisorsClient type.
type AdvisorsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serverName string, advisorName string) (result mariadb.Advisor, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.AdvisorsResultListPage, err error)
	ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.AdvisorsResultListIterator, err error)
}

var _ AdvisorsClientAPI = (*mariadb.AdvisorsClient)(nil)

// RecommendedActionsClientAPI contains the set of methods on the RecommendedActionsClient type.
type RecommendedActionsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serverName string, advisorName string, recommendedActionName string) (result mariadb.RecommendationAction, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string, advisorName string, sessionID string) (result mariadb.RecommendationActionsResultListPage, err error)
	ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string, advisorName string, sessionID string) (result mariadb.RecommendationActionsResultListIterator, err error)
}

var _ RecommendedActionsClientAPI = (*mariadb.RecommendedActionsClient)(nil)

// LocationBasedRecommendedActionSessionsOperationStatusClientAPI contains the set of methods on the LocationBasedRecommendedActionSessionsOperationStatusClient type.
type LocationBasedRecommendedActionSessionsOperationStatusClientAPI interface {
	Get(ctx context.Context, locationName string, operationID string) (result mariadb.RecommendedActionSessionsOperationStatus, err error)
}

var _ LocationBasedRecommendedActionSessionsOperationStatusClientAPI = (*mariadb.LocationBasedRecommendedActionSessionsOperationStatusClient)(nil)

// LocationBasedRecommendedActionSessionsResultClientAPI contains the set of methods on the LocationBasedRecommendedActionSessionsResultClient type.
type LocationBasedRecommendedActionSessionsResultClientAPI interface {
	List(ctx context.Context, locationName string, operationID string) (result mariadb.RecommendationActionsResultListPage, err error)
	ListComplete(ctx context.Context, locationName string, operationID string) (result mariadb.RecommendationActionsResultListIterator, err error)
}

var _ LocationBasedRecommendedActionSessionsResultClientAPI = (*mariadb.LocationBasedRecommendedActionSessionsResultClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, privateEndpointConnectionName string, parameters mariadb.PrivateEndpointConnection) (result mariadb.PrivateEndpointConnectionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, privateEndpointConnectionName string) (result mariadb.PrivateEndpointConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, privateEndpointConnectionName string) (result mariadb.PrivateEndpointConnection, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.PrivateEndpointConnectionListResultPage, err error)
	ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.PrivateEndpointConnectionListResultIterator, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, serverName string, privateEndpointConnectionName string, parameters mariadb.TagsObject) (result mariadb.PrivateEndpointConnectionsUpdateTagsFuture, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*mariadb.PrivateEndpointConnectionsClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serverName string, groupName string) (result mariadb.PrivateLinkResource, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.PrivateLinkResourceListResultPage, err error)
	ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.PrivateLinkResourceListResultIterator, err error)
}

var _ PrivateLinkResourcesClientAPI = (*mariadb.PrivateLinkResourcesClient)(nil)

// ServerSecurityAlertPoliciesClientAPI contains the set of methods on the ServerSecurityAlertPoliciesClient type.
type ServerSecurityAlertPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters mariadb.ServerSecurityAlertPolicy) (result mariadb.ServerSecurityAlertPoliciesCreateOrUpdateFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.ServerSecurityAlertPolicy, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.ServerSecurityAlertPolicyListResultPage, err error)
	ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.ServerSecurityAlertPolicyListResultIterator, err error)
}

var _ ServerSecurityAlertPoliciesClientAPI = (*mariadb.ServerSecurityAlertPoliciesClient)(nil)
