//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservices

// ClientCapabilitiesResponse contains the response from method Client.Capabilities.
type ClientCapabilitiesResponse struct {
	CapabilitiesResponse
}

// ClientCheckNameAvailabilityResponse contains the response from method Client.CheckNameAvailability.
type ClientCheckNameAvailabilityResponse struct {
	CheckNameAvailabilityResult
}

// OperationsClientGetOperationResultResponse contains the response from method OperationsClient.GetOperationResult.
type OperationsClientGetOperationResultResponse struct {
	Vault
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	ClientDiscoveryResponse
}

// OperationsClientOperationStatusGetResponse contains the response from method OperationsClient.OperationStatusGet.
type OperationsClientOperationStatusGetResponse struct {
	OperationResource
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResource
}

// PrivateLinkResourcesClientListResponse contains the response from method PrivateLinkResourcesClient.List.
type PrivateLinkResourcesClientListResponse struct {
	PrivateLinkResources
}

// RegisteredIdentitiesClientDeleteResponse contains the response from method RegisteredIdentitiesClient.Delete.
type RegisteredIdentitiesClientDeleteResponse struct {
	// placeholder for future response values
}

// ReplicationUsagesClientListResponse contains the response from method ReplicationUsagesClient.List.
type ReplicationUsagesClientListResponse struct {
	ReplicationUsageList
}

// UsagesClientListByVaultsResponse contains the response from method UsagesClient.ListByVaults.
type UsagesClientListByVaultsResponse struct {
	VaultUsageList
}

// VaultCertificatesClientCreateResponse contains the response from method VaultCertificatesClient.Create.
type VaultCertificatesClientCreateResponse struct {
	VaultCertificateResponse
}

// VaultExtendedInfoClientCreateOrUpdateResponse contains the response from method VaultExtendedInfoClient.CreateOrUpdate.
type VaultExtendedInfoClientCreateOrUpdateResponse struct {
	VaultExtendedInfoResource
}

// VaultExtendedInfoClientGetResponse contains the response from method VaultExtendedInfoClient.Get.
type VaultExtendedInfoClientGetResponse struct {
	VaultExtendedInfoResource
}

// VaultExtendedInfoClientUpdateResponse contains the response from method VaultExtendedInfoClient.Update.
type VaultExtendedInfoClientUpdateResponse struct {
	VaultExtendedInfoResource
}

// VaultsClientCreateOrUpdateResponse contains the response from method VaultsClient.CreateOrUpdate.
type VaultsClientCreateOrUpdateResponse struct {
	Vault
}

// VaultsClientDeleteResponse contains the response from method VaultsClient.Delete.
type VaultsClientDeleteResponse struct {
	// placeholder for future response values
}

// VaultsClientGetResponse contains the response from method VaultsClient.Get.
type VaultsClientGetResponse struct {
	Vault
}

// VaultsClientListByResourceGroupResponse contains the response from method VaultsClient.ListByResourceGroup.
type VaultsClientListByResourceGroupResponse struct {
	VaultList
}

// VaultsClientListBySubscriptionIDResponse contains the response from method VaultsClient.ListBySubscriptionID.
type VaultsClientListBySubscriptionIDResponse struct {
	VaultList
}

// VaultsClientUpdateResponse contains the response from method VaultsClient.Update.
type VaultsClientUpdateResponse struct {
	Vault
}
