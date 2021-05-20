package adpapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/adp/mgmt/2021-02-01-preview/adp"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result adp.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result adp.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*adp.OperationsClient)(nil)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	CheckNameAvailability(ctx context.Context, parameters adp.AccountCheckNameAvailabilityParameters) (result adp.CheckNameAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, parameters *adp.Account) (result adp.AccountsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result adp.AccountsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result adp.Account, err error)
	List(ctx context.Context) (result adp.AccountListPage, err error)
	ListComplete(ctx context.Context) (result adp.AccountListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result adp.AccountListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result adp.AccountListIterator, err error)
	UpdateABDC(ctx context.Context, resourceGroupName string, accountName string, parameters *adp.AccountPatch) (result adp.AccountsUpdateABDCFuture, err error)
}

var _ AccountsClientAPI = (*adp.AccountsClient)(nil)

// DataPoolsClientAPI contains the set of methods on the DataPoolsClient type.
type DataPoolsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string, parameters *adp.DataPool) (result adp.DataPoolsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string) (result adp.DataPoolsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string) (result adp.DataPool, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result adp.DataPoolListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, accountName string) (result adp.DataPoolListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string, parameters *adp.DataPoolPatch) (result adp.DataPoolsUpdateFuture, err error)
}

var _ DataPoolsClientAPI = (*adp.DataPoolsClient)(nil)
