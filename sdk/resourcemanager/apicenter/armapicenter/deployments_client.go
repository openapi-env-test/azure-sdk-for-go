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

// DeploymentsClient contains the methods for the Deployments group.
// Don't use this type directly, use NewDeploymentsClient() instead.
type DeploymentsClient struct {
	host           string
	subscriptionID string
	workspaceName  string
	apiName        string
	filter         *string
	deploymentName string
	pl             runtime.Pipeline
}

// NewDeploymentsClient creates a new instance of DeploymentsClient with the specified values.
// subscriptionID - The ID of the target subscription. The value must be an UUID.
// workspaceName - The name of the workspace.
// apiName - The name of the API.
// deploymentName - The name of the API deployment.
// filter - OData filter parameter.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDeploymentsClient(subscriptionID string, workspaceName string, apiName string, deploymentName string, filter *string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DeploymentsClient, error) {
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
	client := &DeploymentsClient{
		subscriptionID: subscriptionID,
		workspaceName:  workspaceName,
		apiName:        apiName,
		filter:         filter,
		deploymentName: deploymentName,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates new or updates existing API deployment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2024-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of Azure API Center service.
// payload - API deployment entity.
// options - DeploymentsClientCreateOrUpdateOptions contains the optional parameters for the DeploymentsClient.CreateOrUpdate
// method.
func (client *DeploymentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, payload Deployment, options *DeploymentsClientCreateOrUpdateOptions) (DeploymentsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, payload, options)
	if err != nil {
		return DeploymentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DeploymentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DeploymentsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DeploymentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, payload Deployment, options *DeploymentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/deployments/{deploymentName}"
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
	if client.deploymentName == "" {
		return nil, errors.New("parameter client.deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(client.deploymentName))
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
func (client *DeploymentsClient) createOrUpdateHandleResponse(resp *http.Response) (DeploymentsClientCreateOrUpdateResponse, error) {
	result := DeploymentsClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Deployment); err != nil {
		return DeploymentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes API deployment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2024-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of Azure API Center service.
// options - DeploymentsClientDeleteOptions contains the optional parameters for the DeploymentsClient.Delete method.
func (client *DeploymentsClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, options *DeploymentsClientDeleteOptions) (DeploymentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return DeploymentsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DeploymentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DeploymentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DeploymentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DeploymentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *DeploymentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/deployments/{deploymentName}"
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
	if client.deploymentName == "" {
		return nil, errors.New("parameter client.deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(client.deploymentName))
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

// Get - Returns details of the API deployment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2024-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of Azure API Center service.
// options - DeploymentsClientGetOptions contains the optional parameters for the DeploymentsClient.Get method.
func (client *DeploymentsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, options *DeploymentsClientGetOptions) (DeploymentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return DeploymentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DeploymentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeploymentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DeploymentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *DeploymentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/deployments/{deploymentName}"
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
	if client.deploymentName == "" {
		return nil, errors.New("parameter client.deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(client.deploymentName))
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
func (client *DeploymentsClient) getHandleResponse(resp *http.Response) (DeploymentsClientGetResponse, error) {
	result := DeploymentsClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Deployment); err != nil {
		return DeploymentsClientGetResponse{}, err
	}
	return result, nil
}

// Head - Checks if specified API deployment exists.
// Generated from API version 2024-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of Azure API Center service.
// options - DeploymentsClientHeadOptions contains the optional parameters for the DeploymentsClient.Head method.
func (client *DeploymentsClient) Head(ctx context.Context, resourceGroupName string, serviceName string, options *DeploymentsClientHeadOptions) (DeploymentsClientHeadResponse, error) {
	req, err := client.headCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return DeploymentsClientHeadResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DeploymentsClientHeadResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeploymentsClientHeadResponse{}, runtime.NewResponseError(resp)
	}
	return DeploymentsClientHeadResponse{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}, nil
}

// headCreateRequest creates the Head request.
func (client *DeploymentsClient) headCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *DeploymentsClientHeadOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/deployments/{deploymentName}"
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
	if client.deploymentName == "" {
		return nil, errors.New("parameter client.deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(client.deploymentName))
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

// NewListPager - Returns a collection of API deployments.
// Generated from API version 2024-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of Azure API Center service.
// options - DeploymentsClientListOptions contains the optional parameters for the DeploymentsClient.List method.
func (client *DeploymentsClient) NewListPager(resourceGroupName string, serviceName string, options *DeploymentsClientListOptions) *runtime.Pager[DeploymentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeploymentsClientListResponse]{
		More: func(page DeploymentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeploymentsClientListResponse) (DeploymentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DeploymentsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DeploymentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DeploymentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DeploymentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *DeploymentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/deployments"
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
	if client.filter != nil {
		reqQP.Set("$filter", *client.filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DeploymentsClient) listHandleResponse(resp *http.Response) (DeploymentsClientListResponse, error) {
	result := DeploymentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentCollection); err != nil {
		return DeploymentsClientListResponse{}, err
	}
	return result, nil
}
