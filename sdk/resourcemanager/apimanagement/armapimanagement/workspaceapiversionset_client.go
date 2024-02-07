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
	"strconv"
	"strings"
)

// WorkspaceAPIVersionSetClient contains the methods for the WorkspaceAPIVersionSet group.
// Don't use this type directly, use NewWorkspaceAPIVersionSetClient() instead.
type WorkspaceAPIVersionSetClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkspaceAPIVersionSetClient creates a new instance of WorkspaceAPIVersionSetClient with the specified values.
// subscriptionID - The ID of the target subscription. The value must be an UUID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkspaceAPIVersionSetClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceAPIVersionSetClient, error) {
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
	client := &WorkspaceAPIVersionSetClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or Updates a Api Version Set.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
// versionSetID - Api Version Set identifier. Must be unique in the current API Management service instance.
// parameters - Create or update parameters.
// options - WorkspaceAPIVersionSetClientCreateOrUpdateOptions contains the optional parameters for the WorkspaceAPIVersionSetClient.CreateOrUpdate
// method.
func (client *WorkspaceAPIVersionSetClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, versionSetID string, parameters APIVersionSetContract, options *WorkspaceAPIVersionSetClientCreateOrUpdateOptions) (WorkspaceAPIVersionSetClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, versionSetID, parameters, options)
	if err != nil {
		return WorkspaceAPIVersionSetClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceAPIVersionSetClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return WorkspaceAPIVersionSetClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceAPIVersionSetClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, versionSetID string, parameters APIVersionSetContract, options *WorkspaceAPIVersionSetClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/apiVersionSets/{versionSetId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if versionSetID == "" {
		return nil, errors.New("parameter versionSetID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionSetId}", url.PathEscape(versionSetID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WorkspaceAPIVersionSetClient) createOrUpdateHandleResponse(resp *http.Response) (WorkspaceAPIVersionSetClientCreateOrUpdateResponse, error) {
	result := WorkspaceAPIVersionSetClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIVersionSetContract); err != nil {
		return WorkspaceAPIVersionSetClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes specific Api Version Set.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
// versionSetID - Api Version Set identifier. Must be unique in the current API Management service instance.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// options - WorkspaceAPIVersionSetClientDeleteOptions contains the optional parameters for the WorkspaceAPIVersionSetClient.Delete
// method.
func (client *WorkspaceAPIVersionSetClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, versionSetID string, ifMatch string, options *WorkspaceAPIVersionSetClientDeleteOptions) (WorkspaceAPIVersionSetClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, versionSetID, ifMatch, options)
	if err != nil {
		return WorkspaceAPIVersionSetClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceAPIVersionSetClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return WorkspaceAPIVersionSetClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return WorkspaceAPIVersionSetClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceAPIVersionSetClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, versionSetID string, ifMatch string, options *WorkspaceAPIVersionSetClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/apiVersionSets/{versionSetId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if versionSetID == "" {
		return nil, errors.New("parameter versionSetID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionSetId}", url.PathEscape(versionSetID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the details of the Api Version Set specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
// versionSetID - Api Version Set identifier. Must be unique in the current API Management service instance.
// options - WorkspaceAPIVersionSetClientGetOptions contains the optional parameters for the WorkspaceAPIVersionSetClient.Get
// method.
func (client *WorkspaceAPIVersionSetClient) Get(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, versionSetID string, options *WorkspaceAPIVersionSetClientGetOptions) (WorkspaceAPIVersionSetClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, versionSetID, options)
	if err != nil {
		return WorkspaceAPIVersionSetClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceAPIVersionSetClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceAPIVersionSetClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceAPIVersionSetClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, versionSetID string, options *WorkspaceAPIVersionSetClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/apiVersionSets/{versionSetId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if versionSetID == "" {
		return nil, errors.New("parameter versionSetID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionSetId}", url.PathEscape(versionSetID))
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

// getHandleResponse handles the Get response.
func (client *WorkspaceAPIVersionSetClient) getHandleResponse(resp *http.Response) (WorkspaceAPIVersionSetClientGetResponse, error) {
	result := WorkspaceAPIVersionSetClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIVersionSetContract); err != nil {
		return WorkspaceAPIVersionSetClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the Api Version Set specified by its identifier.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
// versionSetID - Api Version Set identifier. Must be unique in the current API Management service instance.
// options - WorkspaceAPIVersionSetClientGetEntityTagOptions contains the optional parameters for the WorkspaceAPIVersionSetClient.GetEntityTag
// method.
func (client *WorkspaceAPIVersionSetClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, versionSetID string, options *WorkspaceAPIVersionSetClientGetEntityTagOptions) (WorkspaceAPIVersionSetClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, versionSetID, options)
	if err != nil {
		return WorkspaceAPIVersionSetClientGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceAPIVersionSetClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceAPIVersionSetClientGetEntityTagResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *WorkspaceAPIVersionSetClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, versionSetID string, options *WorkspaceAPIVersionSetClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/apiVersionSets/{versionSetId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if versionSetID == "" {
		return nil, errors.New("parameter versionSetID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionSetId}", url.PathEscape(versionSetID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *WorkspaceAPIVersionSetClient) getEntityTagHandleResponse(resp *http.Response) (WorkspaceAPIVersionSetClientGetEntityTagResponse, error) {
	result := WorkspaceAPIVersionSetClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
	return result, nil
}

// NewListByServicePager - Lists a collection of API Version Sets in the specified workspace with a service instance.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
// options - WorkspaceAPIVersionSetClientListByServiceOptions contains the optional parameters for the WorkspaceAPIVersionSetClient.ListByService
// method.
func (client *WorkspaceAPIVersionSetClient) NewListByServicePager(resourceGroupName string, serviceName string, workspaceID string, options *WorkspaceAPIVersionSetClientListByServiceOptions) *runtime.Pager[WorkspaceAPIVersionSetClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceAPIVersionSetClientListByServiceResponse]{
		More: func(page WorkspaceAPIVersionSetClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceAPIVersionSetClientListByServiceResponse) (WorkspaceAPIVersionSetClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkspaceAPIVersionSetClientListByServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkspaceAPIVersionSetClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspaceAPIVersionSetClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *WorkspaceAPIVersionSetClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, options *WorkspaceAPIVersionSetClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/apiVersionSets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *WorkspaceAPIVersionSetClient) listByServiceHandleResponse(resp *http.Response) (WorkspaceAPIVersionSetClientListByServiceResponse, error) {
	result := WorkspaceAPIVersionSetClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIVersionSetCollection); err != nil {
		return WorkspaceAPIVersionSetClientListByServiceResponse{}, err
	}
	return result, nil
}

// Update - Updates the details of the Api VersionSet specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
// versionSetID - Api Version Set identifier. Must be unique in the current API Management service instance.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// parameters - Update parameters.
// options - WorkspaceAPIVersionSetClientUpdateOptions contains the optional parameters for the WorkspaceAPIVersionSetClient.Update
// method.
func (client *WorkspaceAPIVersionSetClient) Update(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, versionSetID string, ifMatch string, parameters APIVersionSetUpdateParameters, options *WorkspaceAPIVersionSetClientUpdateOptions) (WorkspaceAPIVersionSetClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, versionSetID, ifMatch, parameters, options)
	if err != nil {
		return WorkspaceAPIVersionSetClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceAPIVersionSetClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceAPIVersionSetClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *WorkspaceAPIVersionSetClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, versionSetID string, ifMatch string, parameters APIVersionSetUpdateParameters, options *WorkspaceAPIVersionSetClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/apiVersionSets/{versionSetId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if versionSetID == "" {
		return nil, errors.New("parameter versionSetID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionSetId}", url.PathEscape(versionSetID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *WorkspaceAPIVersionSetClient) updateHandleResponse(resp *http.Response) (WorkspaceAPIVersionSetClientUpdateResponse, error) {
	result := WorkspaceAPIVersionSetClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIVersionSetContract); err != nil {
		return WorkspaceAPIVersionSetClientUpdateResponse{}, err
	}
	return result, nil
}
