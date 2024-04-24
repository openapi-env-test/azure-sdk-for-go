//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RestorePointCollectionsClient contains the methods for the RestorePointCollections group.
// Don't use this type directly, use NewRestorePointCollectionsClient() instead.
type RestorePointCollectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRestorePointCollectionsClient creates a new instance of RestorePointCollectionsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRestorePointCollectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RestorePointCollectionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &RestorePointCollectionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - The operation to create or update the restore point collection. Please refer to https://aka.ms/RestorePoints
// for more details. When updating a restore point collection, only tags may be modified.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2024-03-01
// resourceGroupName - The name of the resource group.
// restorePointCollectionName - The name of the restore point collection.
// parameters - Parameters supplied to the Create or Update restore point collection operation.
// options - RestorePointCollectionsClientCreateOrUpdateOptions contains the optional parameters for the RestorePointCollectionsClient.CreateOrUpdate
// method.
func (client *RestorePointCollectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, restorePointCollectionName string, parameters RestorePointCollection, options *RestorePointCollectionsClientCreateOrUpdateOptions) (RestorePointCollectionsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, restorePointCollectionName, parameters, options)
	if err != nil {
		return RestorePointCollectionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RestorePointCollectionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return RestorePointCollectionsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RestorePointCollectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, parameters RestorePointCollection, options *RestorePointCollectionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RestorePointCollectionsClient) createOrUpdateHandleResponse(resp *http.Response) (RestorePointCollectionsClientCreateOrUpdateResponse, error) {
	result := RestorePointCollectionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorePointCollection); err != nil {
		return RestorePointCollectionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - The operation to delete the restore point collection. This operation will also delete all the contained restore
// points.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2024-03-01
// resourceGroupName - The name of the resource group.
// restorePointCollectionName - The name of the Restore Point Collection.
// options - RestorePointCollectionsClientBeginDeleteOptions contains the optional parameters for the RestorePointCollectionsClient.BeginDelete
// method.
func (client *RestorePointCollectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, restorePointCollectionName string, options *RestorePointCollectionsClientBeginDeleteOptions) (*runtime.Poller[RestorePointCollectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, restorePointCollectionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[RestorePointCollectionsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[RestorePointCollectionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - The operation to delete the restore point collection. This operation will also delete all the contained restore
// points.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2024-03-01
func (client *RestorePointCollectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, restorePointCollectionName string, options *RestorePointCollectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, restorePointCollectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RestorePointCollectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, options *RestorePointCollectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The operation to get the restore point collection.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2024-03-01
// resourceGroupName - The name of the resource group.
// restorePointCollectionName - The name of the restore point collection.
// options - RestorePointCollectionsClientGetOptions contains the optional parameters for the RestorePointCollectionsClient.Get
// method.
func (client *RestorePointCollectionsClient) Get(ctx context.Context, resourceGroupName string, restorePointCollectionName string, options *RestorePointCollectionsClientGetOptions) (RestorePointCollectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, restorePointCollectionName, options)
	if err != nil {
		return RestorePointCollectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RestorePointCollectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RestorePointCollectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RestorePointCollectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, options *RestorePointCollectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RestorePointCollectionsClient) getHandleResponse(resp *http.Response) (RestorePointCollectionsClientGetResponse, error) {
	result := RestorePointCollectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorePointCollection); err != nil {
		return RestorePointCollectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of restore point collections in a resource group.
// Generated from API version 2024-03-01
// resourceGroupName - The name of the resource group.
// options - RestorePointCollectionsClientListOptions contains the optional parameters for the RestorePointCollectionsClient.List
// method.
func (client *RestorePointCollectionsClient) NewListPager(resourceGroupName string, options *RestorePointCollectionsClientListOptions) *runtime.Pager[RestorePointCollectionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RestorePointCollectionsClientListResponse]{
		More: func(page RestorePointCollectionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RestorePointCollectionsClientListResponse) (RestorePointCollectionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RestorePointCollectionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RestorePointCollectionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RestorePointCollectionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RestorePointCollectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *RestorePointCollectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections"
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
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorePointCollectionsClient) listHandleResponse(resp *http.Response) (RestorePointCollectionsClientListResponse, error) {
	result := RestorePointCollectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorePointCollectionListResult); err != nil {
		return RestorePointCollectionsClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets the list of restore point collections in the subscription. Use nextLink property in the response
// to get the next page of restore point collections. Do this till nextLink is not null to fetch all
// the restore point collections.
// Generated from API version 2024-03-01
// options - RestorePointCollectionsClientListAllOptions contains the optional parameters for the RestorePointCollectionsClient.ListAll
// method.
func (client *RestorePointCollectionsClient) NewListAllPager(options *RestorePointCollectionsClientListAllOptions) *runtime.Pager[RestorePointCollectionsClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[RestorePointCollectionsClientListAllResponse]{
		More: func(page RestorePointCollectionsClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RestorePointCollectionsClientListAllResponse) (RestorePointCollectionsClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RestorePointCollectionsClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RestorePointCollectionsClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RestorePointCollectionsClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *RestorePointCollectionsClient) listAllCreateRequest(ctx context.Context, options *RestorePointCollectionsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/restorePointCollections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *RestorePointCollectionsClient) listAllHandleResponse(resp *http.Response) (RestorePointCollectionsClientListAllResponse, error) {
	result := RestorePointCollectionsClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorePointCollectionListResult); err != nil {
		return RestorePointCollectionsClientListAllResponse{}, err
	}
	return result, nil
}

// Update - The operation to update the restore point collection.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2024-03-01
// resourceGroupName - The name of the resource group.
// restorePointCollectionName - The name of the restore point collection.
// parameters - Parameters supplied to the Update restore point collection operation.
// options - RestorePointCollectionsClientUpdateOptions contains the optional parameters for the RestorePointCollectionsClient.Update
// method.
func (client *RestorePointCollectionsClient) Update(ctx context.Context, resourceGroupName string, restorePointCollectionName string, parameters RestorePointCollectionUpdate, options *RestorePointCollectionsClientUpdateOptions) (RestorePointCollectionsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, restorePointCollectionName, parameters, options)
	if err != nil {
		return RestorePointCollectionsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RestorePointCollectionsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RestorePointCollectionsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *RestorePointCollectionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, parameters RestorePointCollectionUpdate, options *RestorePointCollectionsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *RestorePointCollectionsClient) updateHandleResponse(resp *http.Response) (RestorePointCollectionsClientUpdateResponse, error) {
	result := RestorePointCollectionsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorePointCollection); err != nil {
		return RestorePointCollectionsClientUpdateResponse{}, err
	}
	return result, nil
}
