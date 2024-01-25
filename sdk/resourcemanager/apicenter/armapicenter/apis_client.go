//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapicenter

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

// ApisClient contains the methods for the Apis group.
// Don't use this type directly, use NewApisClient() instead.
type ApisClient struct {
	host           string
	subscriptionID string
	workspaceName  string
	filter         *string
	apiName        string
	pl             runtime.Pipeline
}

// NewApisClient creates a new instance of ApisClient with the specified values.
// subscriptionID - The ID of the target subscription. The value must be an UUID.
// workspaceName - The name of the workspace.
// apiName - The name of the API.
// filter - OData filter parameter.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewApisClient(subscriptionID string, workspaceName string, apiName string, filter *string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApisClient, error) {
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
	client := &ApisClient{
		subscriptionID: subscriptionID,
		workspaceName:  workspaceName,
		filter:         filter,
		apiName:        apiName,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates new or updates existing API.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2024-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of Azure API Center service.
// payload - API entity.
// options - ApisClientCreateOrUpdateOptions contains the optional parameters for the ApisClient.CreateOrUpdate method.
func (client *ApisClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, payload API, options *ApisClientCreateOrUpdateOptions) (ApisClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, payload, options)
	if err != nil {
		return ApisClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApisClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ApisClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ApisClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, payload API, options *ApisClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.workspaceName == "" {
		return nil, errors.New("parameter client.workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(client.workspaceName))
	if client.apiName == "" {
		return nil, errors.New("parameter client.apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(client.apiName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, payload)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ApisClient) createOrUpdateHandleResponse(resp *http.Response) (ApisClientCreateOrUpdateResponse, error) {
	result := ApisClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.API); err != nil {
		return ApisClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes specified API.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2024-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of Azure API Center service.
// options - ApisClientDeleteOptions contains the optional parameters for the ApisClient.Delete method.
func (client *ApisClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, options *ApisClientDeleteOptions) (ApisClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return ApisClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApisClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ApisClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ApisClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApisClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ApisClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.workspaceName == "" {
		return nil, errors.New("parameter client.workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(client.workspaceName))
	if client.apiName == "" {
		return nil, errors.New("parameter client.apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(client.apiName))
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

// Get - Returns details of the API.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2024-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of Azure API Center service.
// options - ApisClientGetOptions contains the optional parameters for the ApisClient.Get method.
func (client *ApisClient) Get(ctx context.Context, resourceGroupName string, serviceName string, options *ApisClientGetOptions) (ApisClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return ApisClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApisClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApisClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApisClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ApisClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.workspaceName == "" {
		return nil, errors.New("parameter client.workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(client.workspaceName))
	if client.apiName == "" {
		return nil, errors.New("parameter client.apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(client.apiName))
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

// getHandleResponse handles the Get response.
func (client *ApisClient) getHandleResponse(resp *http.Response) (ApisClientGetResponse, error) {
	result := ApisClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.API); err != nil {
		return ApisClientGetResponse{}, err
	}
	return result, nil
}

// Head - Checks if specified API exists.
// Generated from API version 2024-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of Azure API Center service.
// options - ApisClientHeadOptions contains the optional parameters for the ApisClient.Head method.
func (client *ApisClient) Head(ctx context.Context, resourceGroupName string, serviceName string, options *ApisClientHeadOptions) (ApisClientHeadResponse, error) {
	req, err := client.headCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return ApisClientHeadResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApisClientHeadResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApisClientHeadResponse{}, runtime.NewResponseError(resp)
	}
	return ApisClientHeadResponse{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}, nil
}

// headCreateRequest creates the Head request.
func (client *ApisClient) headCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ApisClientHeadOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.workspaceName == "" {
		return nil, errors.New("parameter client.workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(client.workspaceName))
	if client.apiName == "" {
		return nil, errors.New("parameter client.apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(client.apiName))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListPager - Returns a collection of APIs.
// Generated from API version 2024-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of Azure API Center service.
// options - ApisClientListOptions contains the optional parameters for the ApisClient.List method.
func (client *ApisClient) NewListPager(resourceGroupName string, serviceName string, options *ApisClientListOptions) *runtime.Pager[ApisClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApisClientListResponse]{
		More: func(page ApisClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApisClientListResponse) (ApisClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApisClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ApisClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApisClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ApisClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ApisClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.workspaceName == "" {
		return nil, errors.New("parameter client.workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(client.workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	if client.filter != nil {
		reqQP.Set("$filter", *client.filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ApisClient) listHandleResponse(resp *http.Response) (ApisClientListResponse, error) {
	result := ApisClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APICollection); err != nil {
		return ApisClientListResponse{}, err
	}
	return result, nil
}
