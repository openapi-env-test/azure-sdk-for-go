package subscriptionsapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-11-01/subscriptions"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckResourceName(ctx context.Context, resourceNameDefinition *subscriptions.ResourceName) (result subscriptions.CheckResourceNameResult, err error)
}

var _ BaseClientAPI = (*subscriptions.BaseClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result subscriptions.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result subscriptions.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*subscriptions.OperationsClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CheckZonePeers(ctx context.Context, subscriptionID string, parameters subscriptions.CheckZonePeersRequest) (result subscriptions.CheckZonePeersResult, err error)
	Get(ctx context.Context, subscriptionID string) (result subscriptions.Subscription, err error)
	List(ctx context.Context) (result subscriptions.ListResultPage, err error)
	ListComplete(ctx context.Context) (result subscriptions.ListResultIterator, err error)
	ListLocations(ctx context.Context, subscriptionID string) (result subscriptions.LocationListResult, err error)
}

var _ ClientAPI = (*subscriptions.Client)(nil)

// TenantsClientAPI contains the set of methods on the TenantsClient type.
type TenantsClientAPI interface {
	List(ctx context.Context) (result subscriptions.TenantListResultPage, err error)
	ListComplete(ctx context.Context) (result subscriptions.TenantListResultIterator, err error)
}

var _ TenantsClientAPI = (*subscriptions.TenantsClient)(nil)
