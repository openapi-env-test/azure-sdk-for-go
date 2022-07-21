//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragemover

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

// AgentsClient contains the methods for the Agents group.
// Don't use this type directly, use NewAgentsClient() instead.
type AgentsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAgentsClient creates a new instance of AgentsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAgentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AgentsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &AgentsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Creates or updates an agent resource, which references a hybrid compute machine that can run jobs.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageMoverName - The name of the Storage Mover resource.
// agentName - The name of the agent resource.
// options - AgentsClientCreateOrUpdateOptions contains the optional parameters for the AgentsClient.CreateOrUpdate method.
func (client *AgentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, agent Agent, options *AgentsClientCreateOrUpdateOptions) (AgentsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, storageMoverName, agentName, agent, options)
	if err != nil {
		return AgentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AgentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AgentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, agent Agent, options *AgentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/agents/{agentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if agentName == "" {
		return nil, errors.New("parameter agentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentName}", url.PathEscape(agentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, agent)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AgentsClient) createOrUpdateHandleResponse(resp *http.Response) (AgentsClientCreateOrUpdateResponse, error) {
	result := AgentsClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Agent); err != nil {
		return AgentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes an agent resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageMoverName - The name of the Storage Mover resource.
// agentName - The name of the agent resource.
// options - AgentsClientBeginDeleteOptions contains the optional parameters for the AgentsClient.BeginDelete method.
func (client *AgentsClient) BeginDelete(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, options *AgentsClientBeginDeleteOptions) (AgentsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, storageMoverName, agentName, options)
	if err != nil {
		return AgentsClientDeletePollerResponse{}, err
	}
	result := AgentsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AgentsClient.Delete", "location", resp, client.pl)
	if err != nil {
		return AgentsClientDeletePollerResponse{}, err
	}
	result.Poller = &AgentsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes an agent resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AgentsClient) deleteOperation(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, options *AgentsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, storageMoverName, agentName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AgentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, options *AgentsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/agents/{agentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if agentName == "" {
		return nil, errors.New("parameter agentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentName}", url.PathEscape(agentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets an agent resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageMoverName - The name of the Storage Mover resource.
// agentName - The name of the agent resource.
// options - AgentsClientGetOptions contains the optional parameters for the AgentsClient.Get method.
func (client *AgentsClient) Get(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, options *AgentsClientGetOptions) (AgentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, storageMoverName, agentName, options)
	if err != nil {
		return AgentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AgentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AgentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, options *AgentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/agents/{agentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if agentName == "" {
		return nil, errors.New("parameter agentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentName}", url.PathEscape(agentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AgentsClient) getHandleResponse(resp *http.Response) (AgentsClientGetResponse, error) {
	result := AgentsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Agent); err != nil {
		return AgentsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all agents in a Storage Mover.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageMoverName - The name of the Storage Mover resource.
// options - AgentsClientListOptions contains the optional parameters for the AgentsClient.List method.
func (client *AgentsClient) List(resourceGroupName string, storageMoverName string, options *AgentsClientListOptions) *AgentsClientListPager {
	return &AgentsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, storageMoverName, options)
		},
		advancer: func(ctx context.Context, resp AgentsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AgentList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AgentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, options *AgentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/agents"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AgentsClient) listHandleResponse(resp *http.Response) (AgentsClientListResponse, error) {
	result := AgentsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentList); err != nil {
		return AgentsClientListResponse{}, err
	}
	return result, nil
}

// Update - Creates or updates an agent resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageMoverName - The name of the Storage Mover resource.
// agentName - The name of the agent resource.
// options - AgentsClientUpdateOptions contains the optional parameters for the AgentsClient.Update method.
func (client *AgentsClient) Update(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, agent AgentUpdateParameters, options *AgentsClientUpdateOptions) (AgentsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, storageMoverName, agentName, agent, options)
	if err != nil {
		return AgentsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AgentsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AgentsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, agent AgentUpdateParameters, options *AgentsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/agents/{agentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if agentName == "" {
		return nil, errors.New("parameter agentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentName}", url.PathEscape(agentName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, agent)
}

// updateHandleResponse handles the Update response.
func (client *AgentsClient) updateHandleResponse(resp *http.Response) (AgentsClientUpdateResponse, error) {
	result := AgentsClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Agent); err != nil {
		return AgentsClientUpdateResponse{}, err
	}
	return result, nil
}
