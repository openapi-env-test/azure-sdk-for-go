//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// VPNSitesClient contains the methods for the VPNSites group.
// Don't use this type directly, use NewVPNSitesClient() instead.
type VPNSitesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVPNSitesClient creates a new instance of VPNSitesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVPNSitesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VPNSitesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &VPNSitesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates a VpnSite resource if it doesn't exist else updates the existing VpnSite.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnSite.
// vpnSiteName - The name of the VpnSite being created or updated.
// vpnSiteParameters - Parameters supplied to create or update VpnSite.
// options - VPNSitesClientBeginCreateOrUpdateOptions contains the optional parameters for the VPNSitesClient.BeginCreateOrUpdate
// method.
func (client *VPNSitesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters VPNSite, options *VPNSitesClientBeginCreateOrUpdateOptions) (VPNSitesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, vpnSiteName, vpnSiteParameters, options)
	if err != nil {
		return VPNSitesClientCreateOrUpdatePollerResponse{}, err
	}
	result := VPNSitesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VPNSitesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return VPNSitesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VPNSitesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a VpnSite resource if it doesn't exist else updates the existing VpnSite.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VPNSitesClient) createOrUpdate(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters VPNSite, options *VPNSitesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vpnSiteName, vpnSiteParameters, options)
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
func (client *VPNSitesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters VPNSite, options *VPNSitesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnSiteName == "" {
		return nil, errors.New("parameter vpnSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, vpnSiteParameters)
}

// BeginDelete - Deletes a VpnSite.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnSite.
// vpnSiteName - The name of the VpnSite being deleted.
// options - VPNSitesClientBeginDeleteOptions contains the optional parameters for the VPNSitesClient.BeginDelete method.
func (client *VPNSitesClient) BeginDelete(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VPNSitesClientBeginDeleteOptions) (VPNSitesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, vpnSiteName, options)
	if err != nil {
		return VPNSitesClientDeletePollerResponse{}, err
	}
	result := VPNSitesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VPNSitesClient.Delete", "location", resp, client.pl)
	if err != nil {
		return VPNSitesClientDeletePollerResponse{}, err
	}
	result.Poller = &VPNSitesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a VpnSite.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VPNSitesClient) deleteOperation(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VPNSitesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vpnSiteName, options)
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
func (client *VPNSitesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VPNSitesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnSiteName == "" {
		return nil, errors.New("parameter vpnSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Retrieves the details of a VPN site.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnSite.
// vpnSiteName - The name of the VpnSite being retrieved.
// options - VPNSitesClientGetOptions contains the optional parameters for the VPNSitesClient.Get method.
func (client *VPNSitesClient) Get(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VPNSitesClientGetOptions) (VPNSitesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vpnSiteName, options)
	if err != nil {
		return VPNSitesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VPNSitesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VPNSitesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VPNSitesClient) getCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VPNSitesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnSiteName == "" {
		return nil, errors.New("parameter vpnSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VPNSitesClient) getHandleResponse(resp *http.Response) (VPNSitesClientGetResponse, error) {
	result := VPNSitesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNSite); err != nil {
		return VPNSitesClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all the VpnSites in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - VPNSitesClientListOptions contains the optional parameters for the VPNSitesClient.List method.
func (client *VPNSitesClient) List(options *VPNSitesClientListOptions) *VPNSitesClientListPager {
	return &VPNSitesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp VPNSitesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVPNSitesResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VPNSitesClient) listCreateRequest(ctx context.Context, options *VPNSitesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnSites"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VPNSitesClient) listHandleResponse(resp *http.Response) (VPNSitesClientListResponse, error) {
	result := VPNSitesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNSitesResult); err != nil {
		return VPNSitesClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists all the vpnSites in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnSite.
// options - VPNSitesClientListByResourceGroupOptions contains the optional parameters for the VPNSitesClient.ListByResourceGroup
// method.
func (client *VPNSitesClient) ListByResourceGroup(resourceGroupName string, options *VPNSitesClientListByResourceGroupOptions) *VPNSitesClientListByResourceGroupPager {
	return &VPNSitesClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp VPNSitesClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVPNSitesResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VPNSitesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VPNSitesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VPNSitesClient) listByResourceGroupHandleResponse(resp *http.Response) (VPNSitesClientListByResourceGroupResponse, error) {
	result := VPNSitesClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNSitesResult); err != nil {
		return VPNSitesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates VpnSite tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnSite.
// vpnSiteName - The name of the VpnSite being updated.
// vpnSiteParameters - Parameters supplied to update VpnSite tags.
// options - VPNSitesClientUpdateTagsOptions contains the optional parameters for the VPNSitesClient.UpdateTags method.
func (client *VPNSitesClient) UpdateTags(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters TagsObject, options *VPNSitesClientUpdateTagsOptions) (VPNSitesClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, vpnSiteName, vpnSiteParameters, options)
	if err != nil {
		return VPNSitesClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VPNSitesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VPNSitesClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VPNSitesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters TagsObject, options *VPNSitesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnSiteName == "" {
		return nil, errors.New("parameter vpnSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, vpnSiteParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *VPNSitesClient) updateTagsHandleResponse(resp *http.Response) (VPNSitesClientUpdateTagsResponse, error) {
	result := VPNSitesClientUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNSite); err != nil {
		return VPNSitesClientUpdateTagsResponse{}, err
	}
	return result, nil
}
