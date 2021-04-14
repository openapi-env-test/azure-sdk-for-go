package msiapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/msi/mgmt/2018-11-31/msi"
	"github.com/Azure/go-autorest/autorest"
)

// SystemAssignedIdentitiesClientAPI contains the set of methods on the SystemAssignedIdentitiesClient type.
type SystemAssignedIdentitiesClientAPI interface {
	GetByScope(ctx context.Context, scope string) (result msi.SystemAssignedIdentity, err error)
}

var _ SystemAssignedIdentitiesClientAPI = (*msi.SystemAssignedIdentitiesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result msi.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result msi.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*msi.OperationsClient)(nil)

// UserAssignedIdentitiesClientAPI contains the set of methods on the UserAssignedIdentitiesClient type.
type UserAssignedIdentitiesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters msi.Identity) (result msi.Identity, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result msi.Identity, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result msi.UserAssignedIdentitiesListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result msi.UserAssignedIdentitiesListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result msi.UserAssignedIdentitiesListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result msi.UserAssignedIdentitiesListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, parameters msi.IdentityUpdate) (result msi.Identity, err error)
}

var _ UserAssignedIdentitiesClientAPI = (*msi.UserAssignedIdentitiesClient)(nil)
