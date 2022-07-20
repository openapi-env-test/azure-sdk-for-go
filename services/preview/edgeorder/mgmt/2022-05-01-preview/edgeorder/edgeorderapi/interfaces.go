package edgeorderapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/edgeorder/mgmt/2022-05-01-preview/edgeorder"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result edgeorder.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result edgeorder.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*edgeorder.OperationsClient)(nil)

// AddressesClientAPI contains the set of methods on the AddressesClient type.
type AddressesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, addressName string, addressResource edgeorder.AddressResource) (result edgeorder.AddressesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, addressName string) (result edgeorder.AddressesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, addressName string) (result edgeorder.AddressResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, skipToken string, top *int32) (result edgeorder.AddressResourceListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, skipToken string, top *int32) (result edgeorder.AddressResourceListIterator, err error)
	ListBySubscription(ctx context.Context, filter string, skipToken string, top *int32) (result edgeorder.AddressResourceListPage, err error)
	ListBySubscriptionComplete(ctx context.Context, filter string, skipToken string, top *int32) (result edgeorder.AddressResourceListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, addressName string, addressUpdateParameter edgeorder.AddressUpdateParameter, ifMatch string) (result edgeorder.AddressesUpdateFuture, err error)
}

var _ AddressesClientAPI = (*edgeorder.AddressesClient)(nil)

// ProductsAndConfigurationsClientAPI contains the set of methods on the ProductsAndConfigurationsClient type.
type ProductsAndConfigurationsClientAPI interface {
	ListConfigurations(ctx context.Context, configurationsRequest edgeorder.ConfigurationsRequest, skipToken string) (result edgeorder.ConfigurationsPage, err error)
	ListConfigurationsComplete(ctx context.Context, configurationsRequest edgeorder.ConfigurationsRequest, skipToken string) (result edgeorder.ConfigurationsIterator, err error)
	ListProductFamilies(ctx context.Context, productFamiliesRequest edgeorder.ProductFamiliesRequest, expand string, skipToken string) (result edgeorder.ProductFamiliesPage, err error)
	ListProductFamiliesComplete(ctx context.Context, productFamiliesRequest edgeorder.ProductFamiliesRequest, expand string, skipToken string) (result edgeorder.ProductFamiliesIterator, err error)
	ListProductFamiliesMetadata(ctx context.Context, skipToken string) (result edgeorder.ProductFamiliesMetadataPage, err error)
	ListProductFamiliesMetadataComplete(ctx context.Context, skipToken string) (result edgeorder.ProductFamiliesMetadataIterator, err error)
}

var _ ProductsAndConfigurationsClientAPI = (*edgeorder.ProductsAndConfigurationsClient)(nil)

// OrderItemsClientAPI contains the set of methods on the OrderItemsClient type.
type OrderItemsClientAPI interface {
	Cancel(ctx context.Context, resourceGroupName string, orderItemName string, cancellationReason edgeorder.CancellationReason) (result autorest.Response, err error)
	Create(ctx context.Context, resourceGroupName string, orderItemName string, orderItemResource edgeorder.OrderItemResource) (result edgeorder.OrderItemsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, orderItemName string) (result edgeorder.OrderItemsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, orderItemName string, expand string) (result edgeorder.OrderItemResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, expand string, skipToken string, top *int32) (result edgeorder.OrderItemResourceListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, expand string, skipToken string, top *int32) (result edgeorder.OrderItemResourceListIterator, err error)
	ListBySubscription(ctx context.Context, filter string, expand string, skipToken string, top *int32) (result edgeorder.OrderItemResourceListPage, err error)
	ListBySubscriptionComplete(ctx context.Context, filter string, expand string, skipToken string, top *int32) (result edgeorder.OrderItemResourceListIterator, err error)
	Return(ctx context.Context, resourceGroupName string, orderItemName string, returnOrderItemDetails edgeorder.ReturnOrderItemDetails) (result edgeorder.OrderItemsReturnFuture, err error)
	Update(ctx context.Context, resourceGroupName string, orderItemName string, orderItemUpdateParameter edgeorder.OrderItemUpdateParameter, ifMatch string) (result edgeorder.OrderItemsUpdateFuture, err error)
}

var _ OrderItemsClientAPI = (*edgeorder.OrderItemsClient)(nil)

// OrdersClientAPI contains the set of methods on the OrdersClient type.
type OrdersClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, location string, orderName string) (result edgeorder.OrderResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, skipToken string, top *int32) (result edgeorder.OrderResourceListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, skipToken string, top *int32) (result edgeorder.OrderResourceListIterator, err error)
	ListBySubscription(ctx context.Context, skipToken string, top *int32) (result edgeorder.OrderResourceListPage, err error)
	ListBySubscriptionComplete(ctx context.Context, skipToken string, top *int32) (result edgeorder.OrderResourceListIterator, err error)
}

var _ OrdersClientAPI = (*edgeorder.OrdersClient)(nil)
