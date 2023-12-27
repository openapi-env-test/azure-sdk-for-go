//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

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

// RouteMapsClient contains the methods for the RouteMaps group.
// Don't use this type directly, use NewRouteMapsClient() instead.
type RouteMapsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRouteMapsClient creates a new instance of RouteMapsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRouteMapsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RouteMapsClient, error) {
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
	client := &RouteMapsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a RouteMap if it doesn't exist else updates the existing one.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-06-01
// resourceGroupName - The resource group name of the RouteMap's resource group.
// virtualHubName - The name of the VirtualHub containing the RouteMap.
// routeMapName - The name of the RouteMap.
// routeMapParameters - Parameters supplied to create or update a RouteMap.
// options - RouteMapsClientBeginCreateOrUpdateOptions contains the optional parameters for the RouteMapsClient.BeginCreateOrUpdate
// method.
func (client *RouteMapsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, routeMapParameters RouteMap, options *RouteMapsClientBeginCreateOrUpdateOptions) (*runtime.Poller[RouteMapsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualHubName, routeMapName, routeMapParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[RouteMapsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[RouteMapsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates a RouteMap if it doesn't exist else updates the existing one.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-06-01
func (client *RouteMapsClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, routeMapParameters RouteMap, options *RouteMapsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualHubName, routeMapName, routeMapParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RouteMapsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, routeMapParameters RouteMap, options *RouteMapsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeMaps/{routeMapName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if routeMapName == "" {
		return nil, errors.New("parameter routeMapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeMapName}", url.PathEscape(routeMapName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, routeMapParameters)
}

// BeginDelete - Deletes a RouteMap.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-06-01
// resourceGroupName - The resource group name of the RouteMap's resource group.
// virtualHubName - The name of the VirtualHub containing the RouteMap.
// routeMapName - The name of the RouteMap.
// options - RouteMapsClientBeginDeleteOptions contains the optional parameters for the RouteMapsClient.BeginDelete method.
func (client *RouteMapsClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, options *RouteMapsClientBeginDeleteOptions) (*runtime.Poller[RouteMapsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualHubName, routeMapName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[RouteMapsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[RouteMapsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a RouteMap.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-06-01
func (client *RouteMapsClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, options *RouteMapsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualHubName, routeMapName, options)
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
func (client *RouteMapsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, options *RouteMapsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeMaps/{routeMapName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if routeMapName == "" {
		return nil, errors.New("parameter routeMapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeMapName}", url.PathEscape(routeMapName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves the details of a RouteMap.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-06-01
// resourceGroupName - The resource group name of the RouteMap's resource group.
// virtualHubName - The name of the VirtualHub containing the RouteMap.
// routeMapName - The name of the RouteMap.
// options - RouteMapsClientGetOptions contains the optional parameters for the RouteMapsClient.Get method.
func (client *RouteMapsClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, options *RouteMapsClientGetOptions) (RouteMapsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualHubName, routeMapName, options)
	if err != nil {
		return RouteMapsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RouteMapsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RouteMapsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RouteMapsClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, options *RouteMapsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeMaps/{routeMapName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if routeMapName == "" {
		return nil, errors.New("parameter routeMapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeMapName}", url.PathEscape(routeMapName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RouteMapsClient) getHandleResponse(resp *http.Response) (RouteMapsClientGetResponse, error) {
	result := RouteMapsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteMap); err != nil {
		return RouteMapsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Retrieves the details of all RouteMaps.
// Generated from API version 2023-06-01
// resourceGroupName - The resource group name of the RouteMap's resource group'.
// virtualHubName - The name of the VirtualHub containing the RouteMap.
// options - RouteMapsClientListOptions contains the optional parameters for the RouteMapsClient.List method.
func (client *RouteMapsClient) NewListPager(resourceGroupName string, virtualHubName string, options *RouteMapsClientListOptions) *runtime.Pager[RouteMapsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RouteMapsClientListResponse]{
		More: func(page RouteMapsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RouteMapsClientListResponse) (RouteMapsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, virtualHubName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RouteMapsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RouteMapsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RouteMapsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RouteMapsClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *RouteMapsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeMaps"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RouteMapsClient) listHandleResponse(resp *http.Response) (RouteMapsClientListResponse, error) {
	result := RouteMapsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListRouteMapsResult); err != nil {
		return RouteMapsClientListResponse{}, err
	}
	return result, nil
}
