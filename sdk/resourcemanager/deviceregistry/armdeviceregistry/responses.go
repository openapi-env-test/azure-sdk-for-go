//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeviceregistry

// AssetEndpointProfilesClientCreateOrReplaceResponse contains the response from method AssetEndpointProfilesClient.BeginCreateOrReplace.
type AssetEndpointProfilesClientCreateOrReplaceResponse struct {
	// Asset Endpoint Profile definition.
	AssetEndpointProfile
}

// AssetEndpointProfilesClientDeleteResponse contains the response from method AssetEndpointProfilesClient.BeginDelete.
type AssetEndpointProfilesClientDeleteResponse struct {
	// placeholder for future response values
}

// AssetEndpointProfilesClientGetResponse contains the response from method AssetEndpointProfilesClient.Get.
type AssetEndpointProfilesClientGetResponse struct {
	// Asset Endpoint Profile definition.
	AssetEndpointProfile
}

// AssetEndpointProfilesClientListByResourceGroupResponse contains the response from method AssetEndpointProfilesClient.NewListByResourceGroupPager.
type AssetEndpointProfilesClientListByResourceGroupResponse struct {
	// The response of a AssetEndpointProfile list operation.
	AssetEndpointProfileListResult
}

// AssetEndpointProfilesClientListBySubscriptionResponse contains the response from method AssetEndpointProfilesClient.NewListBySubscriptionPager.
type AssetEndpointProfilesClientListBySubscriptionResponse struct {
	// The response of a AssetEndpointProfile list operation.
	AssetEndpointProfileListResult
}

// AssetEndpointProfilesClientUpdateResponse contains the response from method AssetEndpointProfilesClient.BeginUpdate.
type AssetEndpointProfilesClientUpdateResponse struct {
	// Asset Endpoint Profile definition.
	AssetEndpointProfile
}

// AssetsClientCreateOrReplaceResponse contains the response from method AssetsClient.BeginCreateOrReplace.
type AssetsClientCreateOrReplaceResponse struct {
	// Asset definition.
	Asset
}

// AssetsClientDeleteResponse contains the response from method AssetsClient.BeginDelete.
type AssetsClientDeleteResponse struct {
	// placeholder for future response values
}

// AssetsClientGetResponse contains the response from method AssetsClient.Get.
type AssetsClientGetResponse struct {
	// Asset definition.
	Asset
}

// AssetsClientListByResourceGroupResponse contains the response from method AssetsClient.NewListByResourceGroupPager.
type AssetsClientListByResourceGroupResponse struct {
	// The response of a Asset list operation.
	AssetListResult
}

// AssetsClientListBySubscriptionResponse contains the response from method AssetsClient.NewListBySubscriptionPager.
type AssetsClientListBySubscriptionResponse struct {
	// The response of a Asset list operation.
	AssetListResult
}

// AssetsClientUpdateResponse contains the response from method AssetsClient.BeginUpdate.
type AssetsClientUpdateResponse struct {
	// Asset definition.
	Asset
}

// OperationStatusClientGetResponse contains the response from method OperationStatusClient.Get.
type OperationStatusClientGetResponse struct {
	// The current status of an async operation.
	OperationStatusResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}
