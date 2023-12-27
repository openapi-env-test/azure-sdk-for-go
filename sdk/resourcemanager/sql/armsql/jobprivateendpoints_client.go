//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql

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

// JobPrivateEndpointsClient contains the methods for the JobPrivateEndpoints group.
// Don't use this type directly, use NewJobPrivateEndpointsClient() instead.
type JobPrivateEndpointsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewJobPrivateEndpointsClient creates a new instance of JobPrivateEndpointsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewJobPrivateEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*JobPrivateEndpointsClient, error) {
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
	client := &JobPrivateEndpointsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// privateEndpointName - The name of the private endpoint.
// parameters - The requested private endpoint state.
// options - JobPrivateEndpointsClientBeginCreateOrUpdateOptions contains the optional parameters for the JobPrivateEndpointsClient.BeginCreateOrUpdate
// method.
func (client *JobPrivateEndpointsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, privateEndpointName string, parameters JobPrivateEndpoint, options *JobPrivateEndpointsClientBeginCreateOrUpdateOptions) (*runtime.Poller[JobPrivateEndpointsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, jobAgentName, privateEndpointName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[JobPrivateEndpointsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[JobPrivateEndpointsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-05-01-preview
func (client *JobPrivateEndpointsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, privateEndpointName string, parameters JobPrivateEndpoint, options *JobPrivateEndpointsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, privateEndpointName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *JobPrivateEndpointsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, privateEndpointName string, parameters JobPrivateEndpoint, options *JobPrivateEndpointsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/privateEndpoints/{privateEndpointName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if privateEndpointName == "" {
		return nil, errors.New("parameter privateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// privateEndpointName - The name of the private endpoint to delete.
// options - JobPrivateEndpointsClientBeginDeleteOptions contains the optional parameters for the JobPrivateEndpointsClient.BeginDelete
// method.
func (client *JobPrivateEndpointsClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, privateEndpointName string, options *JobPrivateEndpointsClientBeginDeleteOptions) (*runtime.Poller[JobPrivateEndpointsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, jobAgentName, privateEndpointName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[JobPrivateEndpointsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[JobPrivateEndpointsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-05-01-preview
func (client *JobPrivateEndpointsClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, privateEndpointName string, options *JobPrivateEndpointsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, privateEndpointName, options)
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
func (client *JobPrivateEndpointsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, privateEndpointName string, options *JobPrivateEndpointsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/privateEndpoints/{privateEndpointName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if privateEndpointName == "" {
		return nil, errors.New("parameter privateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// privateEndpointName - The name of the private endpoint to get.
// options - JobPrivateEndpointsClientGetOptions contains the optional parameters for the JobPrivateEndpointsClient.Get method.
func (client *JobPrivateEndpointsClient) Get(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, privateEndpointName string, options *JobPrivateEndpointsClientGetOptions) (JobPrivateEndpointsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, privateEndpointName, options)
	if err != nil {
		return JobPrivateEndpointsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobPrivateEndpointsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobPrivateEndpointsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JobPrivateEndpointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, privateEndpointName string, options *JobPrivateEndpointsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/privateEndpoints/{privateEndpointName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if privateEndpointName == "" {
		return nil, errors.New("parameter privateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
func (client *JobPrivateEndpointsClient) getHandleResponse(resp *http.Response) (JobPrivateEndpointsClientGetResponse, error) {
	result := JobPrivateEndpointsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobPrivateEndpoint); err != nil {
		return JobPrivateEndpointsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAgentPager - Gets a list of job agent private endpoints.
// Generated from API version 2023-05-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// jobAgentName - The name of the job agent.
// options - JobPrivateEndpointsClientListByAgentOptions contains the optional parameters for the JobPrivateEndpointsClient.ListByAgent
// method.
func (client *JobPrivateEndpointsClient) NewListByAgentPager(resourceGroupName string, serverName string, jobAgentName string, options *JobPrivateEndpointsClientListByAgentOptions) *runtime.Pager[JobPrivateEndpointsClientListByAgentResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobPrivateEndpointsClientListByAgentResponse]{
		More: func(page JobPrivateEndpointsClientListByAgentResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobPrivateEndpointsClientListByAgentResponse) (JobPrivateEndpointsClientListByAgentResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAgentCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JobPrivateEndpointsClientListByAgentResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return JobPrivateEndpointsClientListByAgentResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobPrivateEndpointsClientListByAgentResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAgentHandleResponse(resp)
		},
	})
}

// listByAgentCreateRequest creates the ListByAgent request.
func (client *JobPrivateEndpointsClient) listByAgentCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, options *JobPrivateEndpointsClientListByAgentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/privateEndpoints"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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

// listByAgentHandleResponse handles the ListByAgent response.
func (client *JobPrivateEndpointsClient) listByAgentHandleResponse(resp *http.Response) (JobPrivateEndpointsClientListByAgentResponse, error) {
	result := JobPrivateEndpointsClientListByAgentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobPrivateEndpointListResult); err != nil {
		return JobPrivateEndpointsClientListByAgentResponse{}, err
	}
	return result, nil
}
