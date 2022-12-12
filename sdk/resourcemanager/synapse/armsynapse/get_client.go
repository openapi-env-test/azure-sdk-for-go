//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse

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

// GetClient contains the methods for the Get group.
// Don't use this type directly, use NewGetClient() instead.
type GetClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGetClient creates a new instance of GetClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGetClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GetClient, error) {
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
	client := &GetClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// IntegrationRuntimeStart - Get an integration runtime start operation status
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// integrationRuntimeName - Integration runtime name
// integrationRuntimeOperationID - Integration runtime Operation Id
// options - GetClientIntegrationRuntimeStartOptions contains the optional parameters for the GetClient.IntegrationRuntimeStart
// method.
func (client *GetClient) IntegrationRuntimeStart(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, integrationRuntimeOperationID string, options *GetClientIntegrationRuntimeStartOptions) (GetClientIntegrationRuntimeStartResponse, error) {
	req, err := client.integrationRuntimeStartCreateRequest(ctx, resourceGroupName, workspaceName, integrationRuntimeName, integrationRuntimeOperationID, options)
	if err != nil {
		return GetClientIntegrationRuntimeStartResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GetClientIntegrationRuntimeStartResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GetClientIntegrationRuntimeStartResponse{}, runtime.NewResponseError(resp)
	}
	return client.integrationRuntimeStartHandleResponse(resp)
}

// integrationRuntimeStartCreateRequest creates the IntegrationRuntimeStart request.
func (client *GetClient) integrationRuntimeStartCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, integrationRuntimeOperationID string, options *GetClientIntegrationRuntimeStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/start/operationstatuses/{integrationRuntimeOperationId}"
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
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	if integrationRuntimeOperationID == "" {
		return nil, errors.New("parameter integrationRuntimeOperationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeOperationId}", url.PathEscape(integrationRuntimeOperationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// integrationRuntimeStartHandleResponse handles the IntegrationRuntimeStart response.
func (client *GetClient) integrationRuntimeStartHandleResponse(resp *http.Response) (GetClientIntegrationRuntimeStartResponse, error) {
	result := GetClientIntegrationRuntimeStartResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationRuntimeOperationStatus); err != nil {
		return GetClientIntegrationRuntimeStartResponse{}, err
	}
	return result, nil
}
