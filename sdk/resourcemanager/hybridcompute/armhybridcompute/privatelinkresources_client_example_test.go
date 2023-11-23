//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/privateLinkScope/PrivateLinkScopePrivateLinkResource_ListGet.json
func ExamplePrivateLinkResourcesClient_NewListByPrivateLinkScopePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListByPrivateLinkScopePager("myResourceGroup", "myPrivateLinkScope", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PrivateLinkResourceListResult = armhybridcompute.PrivateLinkResourceListResult{
		// 	Value: []*armhybridcompute.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("hybridcompute"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/myPrivateLinkScope/privateLinkResources/hybridcompute"),
		// 			Properties: &armhybridcompute.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("hybridcompute"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("HybridCompute.ServerDP"),
		// 					to.Ptr("HybridCompute.K8sConfigurationDP"),
		// 					to.Ptr("HybridCompute.GuestConfigDP")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.his.arc.azure.com"),
		// 						to.Ptr("privatelink.kubernetesconfiguration.azure.com"),
		// 						to.Ptr("privatelink.Guestconfiguration.azure.com")},
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/privateLinkScope/PrivateLinkScopePrivateLinkResource_Get.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "myResourceGroup", "myPrivateLinkScope", "hybridcompute", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResource = armhybridcompute.PrivateLinkResource{
	// 	Name: to.Ptr("hybridcompute"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/privateLinkScopes/privateLinkResources"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/myPrivateLinkScope/privateLinkResources/hybridcompute"),
	// 	Properties: &armhybridcompute.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("hybridcompute"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("HybridCompute.Server"),
	// 			to.Ptr("HybridCompute.K8sConfiguration"),
	// 			to.Ptr("GuestConfig.DP")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.his.arc.azure.com"),
	// 				to.Ptr("privatelink.kubernetesconfiguration.azure.com"),
	// 				to.Ptr("privatelink.Guestconfiguration.azure.com")},
	// 			},
	// 		}
}
