//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragemover

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// StorageMoversClient contains the methods for the StorageMovers group.
// Don't use this type directly, use NewStorageMoversClient() instead.
type StorageMoversClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewStorageMoversClient creates a new instance of StorageMoversClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewStorageMoversClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *StorageMoversClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &StorageMoversClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Creates or updates a top-level Storage Mover resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageMoverName - The name of the Storage Mover resource.
// options - StorageMoversClientCreateOrUpdateOptions contains the optional parameters for the StorageMoversClient.CreateOrUpdate
// method.
func (client *StorageMoversClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, storageMoverName string, storageMover StorageMover, options *StorageMoversClientCreateOrUpdateOptions) (StorageMoversClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, storageMoverName, storageMover, options)
	if err != nil {
		return StorageMoversClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageMoversClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageMoversClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *StorageMoversClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, storageMover StorageMover, options *StorageMoversClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, storageMover)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *StorageMoversClient) createOrUpdateHandleResponse(resp *http.Response) (StorageMoversClientCreateOrUpdateResponse, error) {
	result := StorageMoversClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageMover); err != nil {
		return StorageMoversClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes a Storage Mover resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageMoverName - The name of the Storage Mover resource.
// options - StorageMoversClientBeginDeleteOptions contains the optional parameters for the StorageMoversClient.BeginDelete
// method.
func (client *StorageMoversClient) BeginDelete(ctx context.Context, resourceGroupName string, storageMoverName string, options *StorageMoversClientBeginDeleteOptions) (StorageMoversClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, storageMoverName, options)
	if err != nil {
		return StorageMoversClientDeletePollerResponse{}, err
	}
	result := StorageMoversClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("StorageMoversClient.Delete", "location", resp, client.pl)
	if err != nil {
		return StorageMoversClientDeletePollerResponse{}, err
	}
	result.Poller = &StorageMoversClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a Storage Mover resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *StorageMoversClient) deleteOperation(ctx context.Context, resourceGroupName string, storageMoverName string, options *StorageMoversClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, storageMoverName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *StorageMoversClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, options *StorageMoversClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a Storage Mover resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageMoverName - The name of the Storage Mover resource.
// options - StorageMoversClientGetOptions contains the optional parameters for the StorageMoversClient.Get method.
func (client *StorageMoversClient) Get(ctx context.Context, resourceGroupName string, storageMoverName string, options *StorageMoversClientGetOptions) (StorageMoversClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, storageMoverName, options)
	if err != nil {
		return StorageMoversClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageMoversClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageMoversClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *StorageMoversClient) getCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, options *StorageMoversClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StorageMoversClient) getHandleResponse(resp *http.Response) (StorageMoversClientGetResponse, error) {
	result := StorageMoversClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageMover); err != nil {
		return StorageMoversClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all Storage Movers in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - StorageMoversClientListOptions contains the optional parameters for the StorageMoversClient.List method.
func (client *StorageMoversClient) List(resourceGroupName string, options *StorageMoversClientListOptions) *StorageMoversClientListPager {
	return &StorageMoversClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp StorageMoversClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.List.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *StorageMoversClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *StorageMoversClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *StorageMoversClient) listHandleResponse(resp *http.Response) (StorageMoversClientListResponse, error) {
	result := StorageMoversClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.List); err != nil {
		return StorageMoversClientListResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Lists all Storage Movers in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - StorageMoversClientListBySubscriptionOptions contains the optional parameters for the StorageMoversClient.ListBySubscription
// method.
func (client *StorageMoversClient) ListBySubscription(options *StorageMoversClientListBySubscriptionOptions) *StorageMoversClientListBySubscriptionPager {
	return &StorageMoversClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp StorageMoversClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.List.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *StorageMoversClient) listBySubscriptionCreateRequest(ctx context.Context, options *StorageMoversClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StorageMover/storageMovers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *StorageMoversClient) listBySubscriptionHandleResponse(resp *http.Response) (StorageMoversClientListBySubscriptionResponse, error) {
	result := StorageMoversClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.List); err != nil {
		return StorageMoversClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates properties for a Storage Mover resource. Properties not specified in the request body will be unchanged.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageMoverName - The name of the Storage Mover resource.
// options - StorageMoversClientUpdateOptions contains the optional parameters for the StorageMoversClient.Update method.
func (client *StorageMoversClient) Update(ctx context.Context, resourceGroupName string, storageMoverName string, storageMover UpdateParameters, options *StorageMoversClientUpdateOptions) (StorageMoversClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, storageMoverName, storageMover, options)
	if err != nil {
		return StorageMoversClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageMoversClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageMoversClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *StorageMoversClient) updateCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, storageMover UpdateParameters, options *StorageMoversClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, storageMover)
}

// updateHandleResponse handles the Update response.
func (client *StorageMoversClient) updateHandleResponse(resp *http.Response) (StorageMoversClientUpdateResponse, error) {
	result := StorageMoversClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageMover); err != nil {
		return StorageMoversClientUpdateResponse{}, err
	}
	return result, nil
}
