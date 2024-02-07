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

// WorkspaceAPISchemaClient contains the methods for the WorkspaceAPISchema group.
// Don't use this type directly, use NewWorkspaceAPISchemaClient() instead.
type WorkspaceAPISchemaClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkspaceAPISchemaClient creates a new instance of WorkspaceAPISchemaClient with the specified values.
// subscriptionID - The ID of the target subscription. The value must be an UUID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkspaceAPISchemaClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceAPISchemaClient, error) {
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
	client := &WorkspaceAPISchemaClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates schema configuration for the API.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
// apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
// ;rev=n as a suffix where n is the revision number.
// schemaID - Schema id identifier. Must be unique in the current API Management service instance.
// parameters - The schema contents to apply.
// options - WorkspaceAPISchemaClientBeginCreateOrUpdateOptions contains the optional parameters for the WorkspaceAPISchemaClient.BeginCreateOrUpdate
// method.
func (client *WorkspaceAPISchemaClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, schemaID string, parameters SchemaContract, options *WorkspaceAPISchemaClientBeginCreateOrUpdateOptions) (*runtime.Poller[WorkspaceAPISchemaClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, workspaceID, apiID, schemaID, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[WorkspaceAPISchemaClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[WorkspaceAPISchemaClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates schema configuration for the API.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-05-01-preview
func (client *WorkspaceAPISchemaClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, schemaID string, parameters SchemaContract, options *WorkspaceAPISchemaClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, apiID, schemaID, parameters, options)
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
func (client *WorkspaceAPISchemaClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, schemaID string, parameters SchemaContract, options *WorkspaceAPISchemaClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/apis/{apiId}/schemas/{schemaId}"
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
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if schemaID == "" {
		return nil, errors.New("parameter schemaID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaId}", url.PathEscape(schemaID))
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

// Delete - Deletes the schema configuration at the Api.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
// apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
// ;rev=n as a suffix where n is the revision number.
// schemaID - Schema id identifier. Must be unique in the current API Management service instance.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// options - WorkspaceAPISchemaClientDeleteOptions contains the optional parameters for the WorkspaceAPISchemaClient.Delete
// method.
func (client *WorkspaceAPISchemaClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, schemaID string, ifMatch string, options *WorkspaceAPISchemaClientDeleteOptions) (WorkspaceAPISchemaClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, apiID, schemaID, ifMatch, options)
	if err != nil {
		return WorkspaceAPISchemaClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceAPISchemaClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return WorkspaceAPISchemaClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return WorkspaceAPISchemaClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceAPISchemaClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, schemaID string, ifMatch string, options *WorkspaceAPISchemaClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/apis/{apiId}/schemas/{schemaId}"
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
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if schemaID == "" {
		return nil, errors.New("parameter schemaID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaId}", url.PathEscape(schemaID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Force != nil {
		reqQP.Set("force", strconv.FormatBool(*options.Force))
	}
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the schema configuration at the API level.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
// apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
// ;rev=n as a suffix where n is the revision number.
// schemaID - Schema id identifier. Must be unique in the current API Management service instance.
// options - WorkspaceAPISchemaClientGetOptions contains the optional parameters for the WorkspaceAPISchemaClient.Get method.
func (client *WorkspaceAPISchemaClient) Get(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, schemaID string, options *WorkspaceAPISchemaClientGetOptions) (WorkspaceAPISchemaClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, apiID, schemaID, options)
	if err != nil {
		return WorkspaceAPISchemaClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceAPISchemaClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceAPISchemaClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceAPISchemaClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, schemaID string, options *WorkspaceAPISchemaClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/apis/{apiId}/schemas/{schemaId}"
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
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if schemaID == "" {
		return nil, errors.New("parameter schemaID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaId}", url.PathEscape(schemaID))
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
func (client *WorkspaceAPISchemaClient) getHandleResponse(resp *http.Response) (WorkspaceAPISchemaClientGetResponse, error) {
	result := WorkspaceAPISchemaClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.SchemaContract); err != nil {
		return WorkspaceAPISchemaClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the schema specified by its identifier.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
// apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
// ;rev=n as a suffix where n is the revision number.
// schemaID - Schema id identifier. Must be unique in the current API Management service instance.
// options - WorkspaceAPISchemaClientGetEntityTagOptions contains the optional parameters for the WorkspaceAPISchemaClient.GetEntityTag
// method.
func (client *WorkspaceAPISchemaClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, schemaID string, options *WorkspaceAPISchemaClientGetEntityTagOptions) (WorkspaceAPISchemaClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, apiID, schemaID, options)
	if err != nil {
		return WorkspaceAPISchemaClientGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceAPISchemaClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceAPISchemaClientGetEntityTagResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *WorkspaceAPISchemaClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, schemaID string, options *WorkspaceAPISchemaClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/apis/{apiId}/schemas/{schemaId}"
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
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if schemaID == "" {
		return nil, errors.New("parameter schemaID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaId}", url.PathEscape(schemaID))
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
func (client *WorkspaceAPISchemaClient) getEntityTagHandleResponse(resp *http.Response) (WorkspaceAPISchemaClientGetEntityTagResponse, error) {
	result := WorkspaceAPISchemaClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
	return result, nil
}

// NewListByAPIPager - Get the schema configuration at the API level.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
// apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
// ;rev=n as a suffix where n is the revision number.
// options - WorkspaceAPISchemaClientListByAPIOptions contains the optional parameters for the WorkspaceAPISchemaClient.ListByAPI
// method.
func (client *WorkspaceAPISchemaClient) NewListByAPIPager(resourceGroupName string, serviceName string, workspaceID string, apiID string, options *WorkspaceAPISchemaClientListByAPIOptions) *runtime.Pager[WorkspaceAPISchemaClientListByAPIResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceAPISchemaClientListByAPIResponse]{
		More: func(page WorkspaceAPISchemaClientListByAPIResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceAPISchemaClientListByAPIResponse) (WorkspaceAPISchemaClientListByAPIResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAPICreateRequest(ctx, resourceGroupName, serviceName, workspaceID, apiID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkspaceAPISchemaClientListByAPIResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkspaceAPISchemaClientListByAPIResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspaceAPISchemaClientListByAPIResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAPIHandleResponse(resp)
		},
	})
}

// listByAPICreateRequest creates the ListByAPI request.
func (client *WorkspaceAPISchemaClient) listByAPICreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, options *WorkspaceAPISchemaClientListByAPIOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/apis/{apiId}/schemas"
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
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
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

// listByAPIHandleResponse handles the ListByAPI response.
func (client *WorkspaceAPISchemaClient) listByAPIHandleResponse(resp *http.Response) (WorkspaceAPISchemaClientListByAPIResponse, error) {
	result := WorkspaceAPISchemaClientListByAPIResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SchemaCollection); err != nil {
		return WorkspaceAPISchemaClientListByAPIResponse{}, err
	}
	return result, nil
}
