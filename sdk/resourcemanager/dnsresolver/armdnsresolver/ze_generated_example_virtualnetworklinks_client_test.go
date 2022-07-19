//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/VirtualNetworkLink_Put.json
func ExampleVirtualNetworkLinksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdnsresolver.NewVirtualNetworkLinksClient("abdd4249-9f34-4cc6-8e42-c2e32110603e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"sampleResourceGroup",
		"sampleDnsForwardingRuleset",
		"sampleVirtualNetworkLink",
		armdnsresolver.VirtualNetworkLink{
			Properties: &armdnsresolver.VirtualNetworkLinkProperties{
				Metadata: map[string]*string{
					"additionalProp1": to.Ptr("value1"),
				},
				VirtualNetwork: &armdnsresolver.SubResource{
					ID: to.Ptr("/subscriptions/0403cfa9-9659-4f33-9f30-1f191c51d111/resourceGroups/sampleVnetResourceGroupName/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork"),
				},
			},
		},
		&armdnsresolver.VirtualNetworkLinksClientBeginCreateOrUpdateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/VirtualNetworkLink_Patch.json
func ExampleVirtualNetworkLinksClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdnsresolver.NewVirtualNetworkLinksClient("abdd4249-9f34-4cc6-8e42-c2e32110603e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"sampleResourceGroup",
		"sampleDnsForwardingRuleset",
		"sampleVirtualNetworkLink",
		armdnsresolver.VirtualNetworkLinkPatch{
			Properties: &armdnsresolver.VirtualNetworkLinkPatchProperties{
				Metadata: map[string]*string{
					"additionalProp1": to.Ptr("value1"),
				},
			},
		},
		&armdnsresolver.VirtualNetworkLinksClientBeginUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/VirtualNetworkLink_Delete.json
func ExampleVirtualNetworkLinksClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdnsresolver.NewVirtualNetworkLinksClient("abdd4249-9f34-4cc6-8e42-c2e32110603e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"sampleResourceGroup",
		"sampleDnsForwardingRuleset",
		"sampleVirtualNetworkLink",
		&armdnsresolver.VirtualNetworkLinksClientBeginDeleteOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/VirtualNetworkLink_Get.json
func ExampleVirtualNetworkLinksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdnsresolver.NewVirtualNetworkLinksClient("abdd4249-9f34-4cc6-8e42-c2e32110603e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"sampleResourceGroup",
		"sampleDnsForwardingRuleset",
		"sampleVirtualNetworkLink",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/VirtualNetworkLink_List.json
func ExampleVirtualNetworkLinksClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdnsresolver.NewVirtualNetworkLinksClient("abdd4249-9f34-4cc6-8e42-c2e32110603e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("sampleResourceGroup",
		"sampleDnsForwardingRuleset",
		&armdnsresolver.VirtualNetworkLinksClientListOptions{Top: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}