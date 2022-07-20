//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// ImagesClient contains the methods for the Images group.
// Don't use this type directly, use NewImagesClient() instead.
type ImagesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewImagesClient creates a new instance of ImagesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewImagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ImagesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ImagesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Create or update an image.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// imageName - The name of the image.
// parameters - Parameters supplied to the Create Image operation.
// options - ImagesClientBeginCreateOrUpdateOptions contains the optional parameters for the ImagesClient.BeginCreateOrUpdate
// method.
func (client *ImagesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, imageName string, parameters Image, options *ImagesClientBeginCreateOrUpdateOptions) (ImagesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, imageName, parameters, options)
	if err != nil {
		return ImagesClientCreateOrUpdatePollerResponse{}, err
	}
	result := ImagesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ImagesClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return ImagesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ImagesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update an image.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ImagesClient) createOrUpdate(ctx context.Context, resourceGroupName string, imageName string, parameters Image, options *ImagesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, imageName, parameters, options)
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
func (client *ImagesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, imageName string, parameters Image, options *ImagesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if imageName == "" {
		return nil, errors.New("parameter imageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageName}", url.PathEscape(imageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes an Image.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// imageName - The name of the image.
// options - ImagesClientBeginDeleteOptions contains the optional parameters for the ImagesClient.BeginDelete method.
func (client *ImagesClient) BeginDelete(ctx context.Context, resourceGroupName string, imageName string, options *ImagesClientBeginDeleteOptions) (ImagesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, imageName, options)
	if err != nil {
		return ImagesClientDeletePollerResponse{}, err
	}
	result := ImagesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ImagesClient.Delete", "", resp, client.pl)
	if err != nil {
		return ImagesClientDeletePollerResponse{}, err
	}
	result.Poller = &ImagesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes an Image.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ImagesClient) deleteOperation(ctx context.Context, resourceGroupName string, imageName string, options *ImagesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, imageName, options)
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
func (client *ImagesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, imageName string, options *ImagesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if imageName == "" {
		return nil, errors.New("parameter imageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageName}", url.PathEscape(imageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets an image.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// imageName - The name of the image.
// options - ImagesClientGetOptions contains the optional parameters for the ImagesClient.Get method.
func (client *ImagesClient) Get(ctx context.Context, resourceGroupName string, imageName string, options *ImagesClientGetOptions) (ImagesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, imageName, options)
	if err != nil {
		return ImagesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ImagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ImagesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ImagesClient) getCreateRequest(ctx context.Context, resourceGroupName string, imageName string, options *ImagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if imageName == "" {
		return nil, errors.New("parameter imageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageName}", url.PathEscape(imageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ImagesClient) getHandleResponse(resp *http.Response) (ImagesClientGetResponse, error) {
	result := ImagesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Image); err != nil {
		return ImagesClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets the list of Images in the subscription. Use nextLink property in the response to get the next page of Images.
// Do this till nextLink is null to fetch all the Images.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ImagesClientListOptions contains the optional parameters for the ImagesClient.List method.
func (client *ImagesClient) List(options *ImagesClientListOptions) *ImagesClientListPager {
	return &ImagesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ImagesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ImageListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ImagesClient) listCreateRequest(ctx context.Context, options *ImagesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/images"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ImagesClient) listHandleResponse(resp *http.Response) (ImagesClientListResponse, error) {
	result := ImagesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImageListResult); err != nil {
		return ImagesClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets the list of images under a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - ImagesClientListByResourceGroupOptions contains the optional parameters for the ImagesClient.ListByResourceGroup
// method.
func (client *ImagesClient) ListByResourceGroup(resourceGroupName string, options *ImagesClientListByResourceGroupOptions) *ImagesClientListByResourceGroupPager {
	return &ImagesClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ImagesClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ImageListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ImagesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ImagesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ImagesClient) listByResourceGroupHandleResponse(resp *http.Response) (ImagesClientListByResourceGroupResponse, error) {
	result := ImagesClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImageListResult); err != nil {
		return ImagesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update an image.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// imageName - The name of the image.
// parameters - Parameters supplied to the Update Image operation.
// options - ImagesClientBeginUpdateOptions contains the optional parameters for the ImagesClient.BeginUpdate method.
func (client *ImagesClient) BeginUpdate(ctx context.Context, resourceGroupName string, imageName string, parameters ImageUpdate, options *ImagesClientBeginUpdateOptions) (ImagesClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, imageName, parameters, options)
	if err != nil {
		return ImagesClientUpdatePollerResponse{}, err
	}
	result := ImagesClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ImagesClient.Update", "", resp, client.pl)
	if err != nil {
		return ImagesClientUpdatePollerResponse{}, err
	}
	result.Poller = &ImagesClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Update an image.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ImagesClient) update(ctx context.Context, resourceGroupName string, imageName string, parameters ImageUpdate, options *ImagesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, imageName, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *ImagesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, imageName string, parameters ImageUpdate, options *ImagesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if imageName == "" {
		return nil, errors.New("parameter imageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageName}", url.PathEscape(imageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
