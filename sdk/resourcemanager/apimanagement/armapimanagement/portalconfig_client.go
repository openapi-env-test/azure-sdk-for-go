//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapimanagement

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

// PortalConfigClient contains the methods for the PortalConfig group.
// Don't use this type directly, use NewPortalConfigClient() instead.
type PortalConfigClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPortalConfigClient creates a new instance of PortalConfigClient with the specified values.
// subscriptionID - The ID of the target subscription. The value must be an UUID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPortalConfigClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PortalConfigClient, error) {
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
	client := &PortalConfigClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update the developer portal configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// portalConfigID - Portal configuration identifier.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// parameters - Update the developer portal configuration.
// options - PortalConfigClientCreateOrUpdateOptions contains the optional parameters for the PortalConfigClient.CreateOrUpdate
// method.
func (client *PortalConfigClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, portalConfigID string, ifMatch string, parameters PortalConfigContract, options *PortalConfigClientCreateOrUpdateOptions) (PortalConfigClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, portalConfigID, ifMatch, parameters, options)
	if err != nil {
		return PortalConfigClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PortalConfigClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PortalConfigClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PortalConfigClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, portalConfigID string, ifMatch string, parameters PortalConfigContract, options *PortalConfigClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalconfigs/{portalConfigId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if portalConfigID == "" {
		return nil, errors.New("parameter portalConfigID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{portalConfigId}", url.PathEscape(portalConfigID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PortalConfigClient) createOrUpdateHandleResponse(resp *http.Response) (PortalConfigClientCreateOrUpdateResponse, error) {
	result := PortalConfigClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PortalConfigContract); err != nil {
		return PortalConfigClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Get the developer portal configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// portalConfigID - Portal configuration identifier.
// options - PortalConfigClientGetOptions contains the optional parameters for the PortalConfigClient.Get method.
func (client *PortalConfigClient) Get(ctx context.Context, resourceGroupName string, serviceName string, portalConfigID string, options *PortalConfigClientGetOptions) (PortalConfigClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, portalConfigID, options)
	if err != nil {
		return PortalConfigClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PortalConfigClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PortalConfigClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PortalConfigClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, portalConfigID string, options *PortalConfigClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalconfigs/{portalConfigId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if portalConfigID == "" {
		return nil, errors.New("parameter portalConfigID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{portalConfigId}", url.PathEscape(portalConfigID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PortalConfigClient) getHandleResponse(resp *http.Response) (PortalConfigClientGetResponse, error) {
	result := PortalConfigClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PortalConfigContract); err != nil {
		return PortalConfigClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the developer portal configuration.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// portalConfigID - Portal configuration identifier.
// options - PortalConfigClientGetEntityTagOptions contains the optional parameters for the PortalConfigClient.GetEntityTag
// method.
func (client *PortalConfigClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, portalConfigID string, options *PortalConfigClientGetEntityTagOptions) (PortalConfigClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, portalConfigID, options)
	if err != nil {
		return PortalConfigClientGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PortalConfigClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PortalConfigClientGetEntityTagResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *PortalConfigClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, portalConfigID string, options *PortalConfigClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalconfigs/{portalConfigId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if portalConfigID == "" {
		return nil, errors.New("parameter portalConfigID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{portalConfigId}", url.PathEscape(portalConfigID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *PortalConfigClient) getEntityTagHandleResponse(resp *http.Response) (PortalConfigClientGetEntityTagResponse, error) {
	result := PortalConfigClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
	return result, nil
}

// NewListByServicePager - Lists the developer portal configurations.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// options - PortalConfigClientListByServiceOptions contains the optional parameters for the PortalConfigClient.ListByService
// method.
func (client *PortalConfigClient) NewListByServicePager(resourceGroupName string, serviceName string, options *PortalConfigClientListByServiceOptions) *runtime.Pager[PortalConfigClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[PortalConfigClientListByServiceResponse]{
		More: func(page PortalConfigClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PortalConfigClientListByServiceResponse) (PortalConfigClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PortalConfigClientListByServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PortalConfigClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PortalConfigClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *PortalConfigClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *PortalConfigClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalconfigs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *PortalConfigClient) listByServiceHandleResponse(resp *http.Response) (PortalConfigClientListByServiceResponse, error) {
	result := PortalConfigClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PortalConfigCollection); err != nil {
		return PortalConfigClientListByServiceResponse{}, err
	}
	return result, nil
}

// Update - Update the developer portal configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// portalConfigID - Portal configuration identifier.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// parameters - Update the developer portal configuration.
// options - PortalConfigClientUpdateOptions contains the optional parameters for the PortalConfigClient.Update method.
func (client *PortalConfigClient) Update(ctx context.Context, resourceGroupName string, serviceName string, portalConfigID string, ifMatch string, parameters PortalConfigContract, options *PortalConfigClientUpdateOptions) (PortalConfigClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, portalConfigID, ifMatch, parameters, options)
	if err != nil {
		return PortalConfigClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PortalConfigClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PortalConfigClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *PortalConfigClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, portalConfigID string, ifMatch string, parameters PortalConfigContract, options *PortalConfigClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalconfigs/{portalConfigId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if portalConfigID == "" {
		return nil, errors.New("parameter portalConfigID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{portalConfigId}", url.PathEscape(portalConfigID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *PortalConfigClient) updateHandleResponse(resp *http.Response) (PortalConfigClientUpdateResponse, error) {
	result := PortalConfigClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PortalConfigContract); err != nil {
		return PortalConfigClientUpdateResponse{}, err
	}
	return result, nil
}
