//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappservice

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

// WorkflowVersionsClient contains the methods for the WorkflowVersions group.
// Don't use this type directly, use NewWorkflowVersionsClient() instead.
type WorkflowVersionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkflowVersionsClient creates a new instance of WorkflowVersionsClient with the specified values.
// subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkflowVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkflowVersionsClient, error) {
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
	client := &WorkflowVersionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets a workflow version.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-12-01
// resourceGroupName - Name of the resource group to which the resource belongs.
// name - Site name.
// workflowName - The workflow name.
// versionID - The workflow versionId.
// options - WorkflowVersionsClientGetOptions contains the optional parameters for the WorkflowVersionsClient.Get method.
func (client *WorkflowVersionsClient) Get(ctx context.Context, resourceGroupName string, name string, workflowName string, versionID string, options *WorkflowVersionsClientGetOptions) (WorkflowVersionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, workflowName, versionID, options)
	if err != nil {
		return WorkflowVersionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowVersionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkflowVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, workflowName string, versionID string, options *WorkflowVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{workflowName}/versions/{versionId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if versionID == "" {
		return nil, errors.New("parameter versionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionId}", url.PathEscape(versionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkflowVersionsClient) getHandleResponse(resp *http.Response) (WorkflowVersionsClientGetResponse, error) {
	result := WorkflowVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowVersion); err != nil {
		return WorkflowVersionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of workflow versions.
// Generated from API version 2023-12-01
// resourceGroupName - Name of the resource group to which the resource belongs.
// name - Site name.
// workflowName - The workflow name.
// options - WorkflowVersionsClientListOptions contains the optional parameters for the WorkflowVersionsClient.List method.
func (client *WorkflowVersionsClient) NewListPager(resourceGroupName string, name string, workflowName string, options *WorkflowVersionsClientListOptions) *runtime.Pager[WorkflowVersionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkflowVersionsClientListResponse]{
		More: func(page WorkflowVersionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkflowVersionsClientListResponse) (WorkflowVersionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, name, workflowName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkflowVersionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkflowVersionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkflowVersionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *WorkflowVersionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, name string, workflowName string, options *WorkflowVersionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{workflowName}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkflowVersionsClient) listHandleResponse(resp *http.Response) (WorkflowVersionsClientListResponse, error) {
	result := WorkflowVersionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowVersionListResult); err != nil {
		return WorkflowVersionsClientListResponse{}, err
	}
	return result, nil
}
