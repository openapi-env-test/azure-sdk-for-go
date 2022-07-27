//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// EntityQueryTemplatesClient contains the methods for the EntityQueryTemplates group.
// Don't use this type directly, use NewEntityQueryTemplatesClient() instead.
type EntityQueryTemplatesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewEntityQueryTemplatesClient creates a new instance of EntityQueryTemplatesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEntityQueryTemplatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *EntityQueryTemplatesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &EntityQueryTemplatesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Gets an entity query.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// entityQueryTemplateID - entity query template ID
// options - EntityQueryTemplatesClientGetOptions contains the optional parameters for the EntityQueryTemplatesClient.Get
// method.
func (client *EntityQueryTemplatesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, entityQueryTemplateID string, options *EntityQueryTemplatesClientGetOptions) (EntityQueryTemplatesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, entityQueryTemplateID, options)
	if err != nil {
		return EntityQueryTemplatesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EntityQueryTemplatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EntityQueryTemplatesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EntityQueryTemplatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, entityQueryTemplateID string, options *EntityQueryTemplatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entityQueryTemplates/{entityQueryTemplateId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if entityQueryTemplateID == "" {
		return nil, errors.New("parameter entityQueryTemplateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{entityQueryTemplateId}", url.PathEscape(entityQueryTemplateID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EntityQueryTemplatesClient) getHandleResponse(resp *http.Response) (EntityQueryTemplatesClientGetResponse, error) {
	result := EntityQueryTemplatesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return EntityQueryTemplatesClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all entity query templates.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - EntityQueryTemplatesClientListOptions contains the optional parameters for the EntityQueryTemplatesClient.List
// method.
func (client *EntityQueryTemplatesClient) List(resourceGroupName string, workspaceName string, options *EntityQueryTemplatesClientListOptions) *EntityQueryTemplatesClientListPager {
	return &EntityQueryTemplatesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
		},
		advancer: func(ctx context.Context, resp EntityQueryTemplatesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.EntityQueryTemplateList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *EntityQueryTemplatesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *EntityQueryTemplatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entityQueryTemplates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Kind != nil {
		reqQP.Set("kind", string(*options.Kind))
	}
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EntityQueryTemplatesClient) listHandleResponse(resp *http.Response) (EntityQueryTemplatesClientListResponse, error) {
	result := EntityQueryTemplatesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EntityQueryTemplateList); err != nil {
		return EntityQueryTemplatesClientListResponse{}, err
	}
	return result, nil
}
