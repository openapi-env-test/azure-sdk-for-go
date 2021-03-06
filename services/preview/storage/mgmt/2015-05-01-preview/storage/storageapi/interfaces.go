package storageapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/storage/mgmt/2015-05-01-preview/storage"
	"github.com/Azure/go-autorest/autorest"
)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	CheckNameAvailability(ctx context.Context, accountName storage.AccountCheckNameAvailabilityParameters) (result storage.CheckNameAvailabilityResult, err error)
	Create(ctx context.Context, resourceGroupName string, accountName string, parameters storage.AccountCreateParameters) (result storage.AccountsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	GetProperties(ctx context.Context, resourceGroupName string, accountName string) (result storage.Account, err error)
	List(ctx context.Context) (result storage.AccountListResultPage, err error)
	ListComplete(ctx context.Context) (result storage.AccountListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result storage.AccountListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result storage.AccountListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result storage.AccountKeys, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, accountName string, regenerateKey storage.AccountRegenerateKeyParameters) (result storage.AccountKeys, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, parameters storage.AccountUpdateParameters) (result storage.Account, err error)
}

var _ AccountsClientAPI = (*storage.AccountsClient)(nil)

// UsageClientAPI contains the set of methods on the UsageClient type.
type UsageClientAPI interface {
	List(ctx context.Context) (result storage.UsageListResult, err error)
}

var _ UsageClientAPI = (*storage.UsageClient)(nil)
