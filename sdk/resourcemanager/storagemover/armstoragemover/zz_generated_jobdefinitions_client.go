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

// JobDefinitionsClient contains the methods for the JobDefinitions group.
// Don't use this type directly, use NewJobDefinitionsClient() instead.
type JobDefinitionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewJobDefinitionsClient creates a new instance of JobDefinitionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewJobDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *JobDefinitionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &JobDefinitionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Creates or updates a job definition resource, which contains configuration for a single unit of managed
// data transfer.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageMoverName - The name of the Storage Mover resource.
// projectName - The name of the project resource.
// jobDefinitionName - The name of the job definition resource.
// options - JobDefinitionsClientCreateOrUpdateOptions contains the optional parameters for the JobDefinitionsClient.CreateOrUpdate
// method.
func (client *JobDefinitionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, jobDefinition JobDefinition, options *JobDefinitionsClientCreateOrUpdateOptions) (JobDefinitionsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, storageMoverName, projectName, jobDefinitionName, jobDefinition, options)
	if err != nil {
		return JobDefinitionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobDefinitionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobDefinitionsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *JobDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, jobDefinition JobDefinition, options *JobDefinitionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/projects/{projectName}/jobDefinitions/{jobDefinitionName}"
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
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, jobDefinition)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *JobDefinitionsClient) createOrUpdateHandleResponse(resp *http.Response) (JobDefinitionsClientCreateOrUpdateResponse, error) {
	result := JobDefinitionsClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobDefinition); err != nil {
		return JobDefinitionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes a job definition resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageMoverName - The name of the Storage Mover resource.
// projectName - The name of the project resource.
// jobDefinitionName - The name of the job definition resource.
// options - JobDefinitionsClientBeginDeleteOptions contains the optional parameters for the JobDefinitionsClient.BeginDelete
// method.
func (client *JobDefinitionsClient) BeginDelete(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, options *JobDefinitionsClientBeginDeleteOptions) (JobDefinitionsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, storageMoverName, projectName, jobDefinitionName, options)
	if err != nil {
		return JobDefinitionsClientDeletePollerResponse{}, err
	}
	result := JobDefinitionsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("JobDefinitionsClient.Delete", "location", resp, client.pl)
	if err != nil {
		return JobDefinitionsClientDeletePollerResponse{}, err
	}
	result.Poller = &JobDefinitionsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a job definition resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *JobDefinitionsClient) deleteOperation(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, options *JobDefinitionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, storageMoverName, projectName, jobDefinitionName, options)
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
func (client *JobDefinitionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, options *JobDefinitionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/projects/{projectName}/jobDefinitions/{jobDefinitionName}"
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
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
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

// Get - Gets a job definition resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageMoverName - The name of the Storage Mover resource.
// projectName - The name of the project resource.
// jobDefinitionName - The name of the job definition resource.
// options - JobDefinitionsClientGetOptions contains the optional parameters for the JobDefinitionsClient.Get method.
func (client *JobDefinitionsClient) Get(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, options *JobDefinitionsClientGetOptions) (JobDefinitionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, storageMoverName, projectName, jobDefinitionName, options)
	if err != nil {
		return JobDefinitionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobDefinitionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JobDefinitionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, options *JobDefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/projects/{projectName}/jobDefinitions/{jobDefinitionName}"
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
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
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
func (client *JobDefinitionsClient) getHandleResponse(resp *http.Response) (JobDefinitionsClientGetResponse, error) {
	result := JobDefinitionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobDefinition); err != nil {
		return JobDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all job definitions in a project.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageMoverName - The name of the Storage Mover resource.
// projectName - The name of the project resource.
// options - JobDefinitionsClientListOptions contains the optional parameters for the JobDefinitionsClient.List method.
func (client *JobDefinitionsClient) List(resourceGroupName string, storageMoverName string, projectName string, options *JobDefinitionsClientListOptions) *JobDefinitionsClientListPager {
	return &JobDefinitionsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, storageMoverName, projectName, options)
		},
		advancer: func(ctx context.Context, resp JobDefinitionsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.JobDefinitionList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *JobDefinitionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, options *JobDefinitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/projects/{projectName}/jobDefinitions"
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
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
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
func (client *JobDefinitionsClient) listHandleResponse(resp *http.Response) (JobDefinitionsClientListResponse, error) {
	result := JobDefinitionsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobDefinitionList); err != nil {
		return JobDefinitionsClientListResponse{}, err
	}
	return result, nil
}

// StartJob - Requests an agent to start a new instance of this job definition, generating a new job run resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageMoverName - The name of the Storage Mover resource.
// projectName - The name of the project resource.
// jobDefinitionName - The name of the job definition resource.
// options - JobDefinitionsClientStartJobOptions contains the optional parameters for the JobDefinitionsClient.StartJob method.
func (client *JobDefinitionsClient) StartJob(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, options *JobDefinitionsClientStartJobOptions) (JobDefinitionsClientStartJobResponse, error) {
	req, err := client.startJobCreateRequest(ctx, resourceGroupName, storageMoverName, projectName, jobDefinitionName, options)
	if err != nil {
		return JobDefinitionsClientStartJobResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobDefinitionsClientStartJobResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobDefinitionsClientStartJobResponse{}, runtime.NewResponseError(resp)
	}
	return client.startJobHandleResponse(resp)
}

// startJobCreateRequest creates the StartJob request.
func (client *JobDefinitionsClient) startJobCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, options *JobDefinitionsClientStartJobOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/projects/{projectName}/jobDefinitions/{jobDefinitionName}/startJob"
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
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// startJobHandleResponse handles the StartJob response.
func (client *JobDefinitionsClient) startJobHandleResponse(resp *http.Response) (JobDefinitionsClientStartJobResponse, error) {
	result := JobDefinitionsClientStartJobResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobRunResourceID); err != nil {
		return JobDefinitionsClientStartJobResponse{}, err
	}
	return result, nil
}

// StopJob - Requests the agent of any active instance of this job definition to stop.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageMoverName - The name of the Storage Mover resource.
// projectName - The name of the project resource.
// jobDefinitionName - The name of the job definition resource.
// options - JobDefinitionsClientStopJobOptions contains the optional parameters for the JobDefinitionsClient.StopJob method.
func (client *JobDefinitionsClient) StopJob(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, options *JobDefinitionsClientStopJobOptions) (JobDefinitionsClientStopJobResponse, error) {
	req, err := client.stopJobCreateRequest(ctx, resourceGroupName, storageMoverName, projectName, jobDefinitionName, options)
	if err != nil {
		return JobDefinitionsClientStopJobResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobDefinitionsClientStopJobResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobDefinitionsClientStopJobResponse{}, runtime.NewResponseError(resp)
	}
	return client.stopJobHandleResponse(resp)
}

// stopJobCreateRequest creates the StopJob request.
func (client *JobDefinitionsClient) stopJobCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, options *JobDefinitionsClientStopJobOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/projects/{projectName}/jobDefinitions/{jobDefinitionName}/stopJob"
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
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// stopJobHandleResponse handles the StopJob response.
func (client *JobDefinitionsClient) stopJobHandleResponse(resp *http.Response) (JobDefinitionsClientStopJobResponse, error) {
	result := JobDefinitionsClientStopJobResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobRunResourceID); err != nil {
		return JobDefinitionsClientStopJobResponse{}, err
	}
	return result, nil
}

// Update - Updates properties for a job definition resource. Properties not specified in the request body will be unchanged.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageMoverName - The name of the Storage Mover resource.
// projectName - The name of the project resource.
// jobDefinitionName - The name of the job definition resource.
// options - JobDefinitionsClientUpdateOptions contains the optional parameters for the JobDefinitionsClient.Update method.
func (client *JobDefinitionsClient) Update(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, jobDefinition JobDefinitionUpdateParameters, options *JobDefinitionsClientUpdateOptions) (JobDefinitionsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, storageMoverName, projectName, jobDefinitionName, jobDefinition, options)
	if err != nil {
		return JobDefinitionsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobDefinitionsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobDefinitionsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *JobDefinitionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, jobDefinition JobDefinitionUpdateParameters, options *JobDefinitionsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/projects/{projectName}/jobDefinitions/{jobDefinitionName}"
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
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, jobDefinition)
}

// updateHandleResponse handles the Update response.
func (client *JobDefinitionsClient) updateHandleResponse(resp *http.Response) (JobDefinitionsClientUpdateResponse, error) {
	result := JobDefinitionsClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobDefinition); err != nil {
		return JobDefinitionsClientUpdateResponse{}, err
	}
	return result, nil
}
