package databricksapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/databricks/mgmt/2021-04-01-preview/databricks"
)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, parameters databricks.Workspace, resourceGroupName string, workspaceName string) (result databricks.WorkspacesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string) (result databricks.WorkspacesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result databricks.Workspace, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result databricks.WorkspaceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result databricks.WorkspaceListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result databricks.WorkspaceListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result databricks.WorkspaceListResultIterator, err error)
	Update(ctx context.Context, parameters databricks.WorkspaceUpdate, resourceGroupName string, workspaceName string) (result databricks.WorkspacesUpdateFuture, err error)
}

var _ WorkspacesClientAPI = (*databricks.WorkspacesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result databricks.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result databricks.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*databricks.OperationsClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workspaceName string, groupID string) (result databricks.GroupIDInformation, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result databricks.PrivateLinkResourcesListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result databricks.PrivateLinkResourcesListIterator, err error)
}

var _ PrivateLinkResourcesClientAPI = (*databricks.PrivateLinkResourcesClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, privateEndpointConnection databricks.PrivateEndpointConnection) (result databricks.PrivateEndpointConnectionsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string) (result databricks.PrivateEndpointConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string) (result databricks.PrivateEndpointConnection, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result databricks.PrivateEndpointConnectionsListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result databricks.PrivateEndpointConnectionsListIterator, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*databricks.PrivateEndpointConnectionsClient)(nil)

// OutboundNetworkDependenciesEndpointsClientAPI contains the set of methods on the OutboundNetworkDependenciesEndpointsClient type.
type OutboundNetworkDependenciesEndpointsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result databricks.ListOutboundEnvironmentEndpoint, err error)
}

var _ OutboundNetworkDependenciesEndpointsClientAPI = (*databricks.OutboundNetworkDependenciesEndpointsClient)(nil)

// VNetPeeringClientAPI contains the set of methods on the VNetPeeringClient type.
type VNetPeeringClientAPI interface {
	CreateOrUpdate(ctx context.Context, virtualNetworkPeeringParameters databricks.VirtualNetworkPeering, resourceGroupName string, workspaceName string, peeringName string) (result databricks.VNetPeeringCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, peeringName string) (result databricks.VNetPeeringDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, peeringName string) (result databricks.VirtualNetworkPeering, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result databricks.VirtualNetworkPeeringListPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result databricks.VirtualNetworkPeeringListIterator, err error)
}

var _ VNetPeeringClientAPI = (*databricks.VNetPeeringClient)(nil)
