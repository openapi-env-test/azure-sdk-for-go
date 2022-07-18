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

// ProximityPlacementGroupsClient contains the methods for the ProximityPlacementGroups group.
// Don't use this type directly, use NewProximityPlacementGroupsClient() instead.
type ProximityPlacementGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProximityPlacementGroupsClient creates a new instance of ProximityPlacementGroupsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProximityPlacementGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ProximityPlacementGroupsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ProximityPlacementGroupsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Create or update a proximity placement group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// proximityPlacementGroupName - The name of the proximity placement group.
// parameters - Parameters supplied to the Create Proximity Placement Group operation.
// options - ProximityPlacementGroupsClientCreateOrUpdateOptions contains the optional parameters for the ProximityPlacementGroupsClient.CreateOrUpdate
// method.
func (client *ProximityPlacementGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroup, options *ProximityPlacementGroupsClientCreateOrUpdateOptions) (ProximityPlacementGroupsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, proximityPlacementGroupName, parameters, options)
	if err != nil {
		return ProximityPlacementGroupsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProximityPlacementGroupsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ProximityPlacementGroupsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProximityPlacementGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroup, options *ProximityPlacementGroupsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if proximityPlacementGroupName == "" {
		return nil, errors.New("parameter proximityPlacementGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{proximityPlacementGroupName}", url.PathEscape(proximityPlacementGroupName))
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

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ProximityPlacementGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (ProximityPlacementGroupsClientCreateOrUpdateResponse, error) {
	result := ProximityPlacementGroupsClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProximityPlacementGroup); err != nil {
		return ProximityPlacementGroupsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a proximity placement group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// proximityPlacementGroupName - The name of the proximity placement group.
// options - ProximityPlacementGroupsClientDeleteOptions contains the optional parameters for the ProximityPlacementGroupsClient.Delete
// method.
func (client *ProximityPlacementGroupsClient) Delete(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, options *ProximityPlacementGroupsClientDeleteOptions) (ProximityPlacementGroupsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, proximityPlacementGroupName, options)
	if err != nil {
		return ProximityPlacementGroupsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProximityPlacementGroupsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProximityPlacementGroupsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ProximityPlacementGroupsClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProximityPlacementGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, options *ProximityPlacementGroupsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if proximityPlacementGroupName == "" {
		return nil, errors.New("parameter proximityPlacementGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{proximityPlacementGroupName}", url.PathEscape(proximityPlacementGroupName))
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

// Get - Retrieves information about a proximity placement group .
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// proximityPlacementGroupName - The name of the proximity placement group.
// options - ProximityPlacementGroupsClientGetOptions contains the optional parameters for the ProximityPlacementGroupsClient.Get
// method.
func (client *ProximityPlacementGroupsClient) Get(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, options *ProximityPlacementGroupsClientGetOptions) (ProximityPlacementGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, proximityPlacementGroupName, options)
	if err != nil {
		return ProximityPlacementGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProximityPlacementGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProximityPlacementGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProximityPlacementGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, options *ProximityPlacementGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if proximityPlacementGroupName == "" {
		return nil, errors.New("parameter proximityPlacementGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{proximityPlacementGroupName}", url.PathEscape(proximityPlacementGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.IncludeColocationStatus != nil {
		reqQP.Set("includeColocationStatus", *options.IncludeColocationStatus)
	}
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProximityPlacementGroupsClient) getHandleResponse(resp *http.Response) (ProximityPlacementGroupsClientGetResponse, error) {
	result := ProximityPlacementGroupsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProximityPlacementGroup); err != nil {
		return ProximityPlacementGroupsClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists all proximity placement groups in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - ProximityPlacementGroupsClientListByResourceGroupOptions contains the optional parameters for the ProximityPlacementGroupsClient.ListByResourceGroup
// method.
func (client *ProximityPlacementGroupsClient) ListByResourceGroup(resourceGroupName string, options *ProximityPlacementGroupsClientListByResourceGroupOptions) *ProximityPlacementGroupsClientListByResourceGroupPager {
	return &ProximityPlacementGroupsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ProximityPlacementGroupsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ProximityPlacementGroupListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ProximityPlacementGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ProximityPlacementGroupsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups"
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
func (client *ProximityPlacementGroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (ProximityPlacementGroupsClientListByResourceGroupResponse, error) {
	result := ProximityPlacementGroupsClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProximityPlacementGroupListResult); err != nil {
		return ProximityPlacementGroupsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Lists all proximity placement groups in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ProximityPlacementGroupsClientListBySubscriptionOptions contains the optional parameters for the ProximityPlacementGroupsClient.ListBySubscription
// method.
func (client *ProximityPlacementGroupsClient) ListBySubscription(options *ProximityPlacementGroupsClientListBySubscriptionOptions) *ProximityPlacementGroupsClientListBySubscriptionPager {
	return &ProximityPlacementGroupsClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ProximityPlacementGroupsClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ProximityPlacementGroupListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ProximityPlacementGroupsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ProximityPlacementGroupsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/proximityPlacementGroups"
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

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ProximityPlacementGroupsClient) listBySubscriptionHandleResponse(resp *http.Response) (ProximityPlacementGroupsClientListBySubscriptionResponse, error) {
	result := ProximityPlacementGroupsClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProximityPlacementGroupListResult); err != nil {
		return ProximityPlacementGroupsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update a proximity placement group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// proximityPlacementGroupName - The name of the proximity placement group.
// parameters - Parameters supplied to the Update Proximity Placement Group operation.
// options - ProximityPlacementGroupsClientUpdateOptions contains the optional parameters for the ProximityPlacementGroupsClient.Update
// method.
func (client *ProximityPlacementGroupsClient) Update(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroupUpdate, options *ProximityPlacementGroupsClientUpdateOptions) (ProximityPlacementGroupsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, proximityPlacementGroupName, parameters, options)
	if err != nil {
		return ProximityPlacementGroupsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProximityPlacementGroupsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProximityPlacementGroupsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ProximityPlacementGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroupUpdate, options *ProximityPlacementGroupsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if proximityPlacementGroupName == "" {
		return nil, errors.New("parameter proximityPlacementGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{proximityPlacementGroupName}", url.PathEscape(proximityPlacementGroupName))
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

// updateHandleResponse handles the Update response.
func (client *ProximityPlacementGroupsClient) updateHandleResponse(resp *http.Response) (ProximityPlacementGroupsClientUpdateResponse, error) {
	result := ProximityPlacementGroupsClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProximityPlacementGroup); err != nil {
		return ProximityPlacementGroupsClientUpdateResponse{}, err
	}
	return result, nil
}
