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

// SnapshotsClient contains the methods for the Snapshots group.
// Don't use this type directly, use NewSnapshotsClient() instead.
type SnapshotsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSnapshotsClient creates a new instance of SnapshotsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSnapshotsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SnapshotsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &SnapshotsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// snapshotName - The name of the snapshot that is being created. The name can't be changed after the snapshot is created.
// Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80
// characters.
// snapshot - Snapshot object supplied in the body of the Put disk operation.
// options - SnapshotsClientBeginCreateOrUpdateOptions contains the optional parameters for the SnapshotsClient.BeginCreateOrUpdate
// method.
func (client *SnapshotsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, snapshotName string, snapshot Snapshot, options *SnapshotsClientBeginCreateOrUpdateOptions) (SnapshotsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, snapshotName, snapshot, options)
	if err != nil {
		return SnapshotsClientCreateOrUpdatePollerResponse{}, err
	}
	result := SnapshotsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SnapshotsClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return SnapshotsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &SnapshotsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SnapshotsClient) createOrUpdate(ctx context.Context, resourceGroupName string, snapshotName string, snapshot Snapshot, options *SnapshotsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, snapshotName, snapshot, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SnapshotsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, snapshot Snapshot, options *SnapshotsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, snapshot)
}

// BeginDelete - Deletes a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// snapshotName - The name of the snapshot that is being created. The name can't be changed after the snapshot is created.
// Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80
// characters.
// options - SnapshotsClientBeginDeleteOptions contains the optional parameters for the SnapshotsClient.BeginDelete method.
func (client *SnapshotsClient) BeginDelete(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsClientBeginDeleteOptions) (SnapshotsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, snapshotName, options)
	if err != nil {
		return SnapshotsClientDeletePollerResponse{}, err
	}
	result := SnapshotsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SnapshotsClient.Delete", "", resp, client.pl)
	if err != nil {
		return SnapshotsClientDeletePollerResponse{}, err
	}
	result.Poller = &SnapshotsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SnapshotsClient) deleteOperation(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, snapshotName, options)
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
func (client *SnapshotsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets information about a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// snapshotName - The name of the snapshot that is being created. The name can't be changed after the snapshot is created.
// Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80
// characters.
// options - SnapshotsClientGetOptions contains the optional parameters for the SnapshotsClient.Get method.
func (client *SnapshotsClient) Get(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsClientGetOptions) (SnapshotsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, snapshotName, options)
	if err != nil {
		return SnapshotsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SnapshotsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SnapshotsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SnapshotsClient) getCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SnapshotsClient) getHandleResponse(resp *http.Response) (SnapshotsClientGetResponse, error) {
	result := SnapshotsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Snapshot); err != nil {
		return SnapshotsClientGetResponse{}, err
	}
	return result, nil
}

// BeginGrantAccess - Grants access to a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// snapshotName - The name of the snapshot that is being created. The name can't be changed after the snapshot is created.
// Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80
// characters.
// grantAccessData - Access data object supplied in the body of the get snapshot access operation.
// options - SnapshotsClientBeginGrantAccessOptions contains the optional parameters for the SnapshotsClient.BeginGrantAccess
// method.
func (client *SnapshotsClient) BeginGrantAccess(ctx context.Context, resourceGroupName string, snapshotName string, grantAccessData GrantAccessData, options *SnapshotsClientBeginGrantAccessOptions) (SnapshotsClientGrantAccessPollerResponse, error) {
	resp, err := client.grantAccess(ctx, resourceGroupName, snapshotName, grantAccessData, options)
	if err != nil {
		return SnapshotsClientGrantAccessPollerResponse{}, err
	}
	result := SnapshotsClientGrantAccessPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SnapshotsClient.GrantAccess", "location", resp, client.pl)
	if err != nil {
		return SnapshotsClientGrantAccessPollerResponse{}, err
	}
	result.Poller = &SnapshotsClientGrantAccessPoller{
		pt: pt,
	}
	return result, nil
}

// GrantAccess - Grants access to a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SnapshotsClient) grantAccess(ctx context.Context, resourceGroupName string, snapshotName string, grantAccessData GrantAccessData, options *SnapshotsClientBeginGrantAccessOptions) (*http.Response, error) {
	req, err := client.grantAccessCreateRequest(ctx, resourceGroupName, snapshotName, grantAccessData, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// grantAccessCreateRequest creates the GrantAccess request.
func (client *SnapshotsClient) grantAccessCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, grantAccessData GrantAccessData, options *SnapshotsClientBeginGrantAccessOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}/beginGetAccess"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, grantAccessData)
}

// List - Lists snapshots under a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - SnapshotsClientListOptions contains the optional parameters for the SnapshotsClient.List method.
func (client *SnapshotsClient) List(options *SnapshotsClientListOptions) *SnapshotsClientListPager {
	return &SnapshotsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SnapshotsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SnapshotList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SnapshotsClient) listCreateRequest(ctx context.Context, options *SnapshotsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/snapshots"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SnapshotsClient) listHandleResponse(resp *http.Response) (SnapshotsClientListResponse, error) {
	result := SnapshotsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SnapshotList); err != nil {
		return SnapshotsClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists snapshots under a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - SnapshotsClientListByResourceGroupOptions contains the optional parameters for the SnapshotsClient.ListByResourceGroup
// method.
func (client *SnapshotsClient) ListByResourceGroup(resourceGroupName string, options *SnapshotsClientListByResourceGroupOptions) *SnapshotsClientListByResourceGroupPager {
	return &SnapshotsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp SnapshotsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SnapshotList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SnapshotsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SnapshotsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots"
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
	reqQP.Set("api-version", "2022-07-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SnapshotsClient) listByResourceGroupHandleResponse(resp *http.Response) (SnapshotsClientListByResourceGroupResponse, error) {
	result := SnapshotsClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SnapshotList); err != nil {
		return SnapshotsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginRevokeAccess - Revokes access to a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// snapshotName - The name of the snapshot that is being created. The name can't be changed after the snapshot is created.
// Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80
// characters.
// options - SnapshotsClientBeginRevokeAccessOptions contains the optional parameters for the SnapshotsClient.BeginRevokeAccess
// method.
func (client *SnapshotsClient) BeginRevokeAccess(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsClientBeginRevokeAccessOptions) (SnapshotsClientRevokeAccessPollerResponse, error) {
	resp, err := client.revokeAccess(ctx, resourceGroupName, snapshotName, options)
	if err != nil {
		return SnapshotsClientRevokeAccessPollerResponse{}, err
	}
	result := SnapshotsClientRevokeAccessPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SnapshotsClient.RevokeAccess", "location", resp, client.pl)
	if err != nil {
		return SnapshotsClientRevokeAccessPollerResponse{}, err
	}
	result.Poller = &SnapshotsClientRevokeAccessPoller{
		pt: pt,
	}
	return result, nil
}

// RevokeAccess - Revokes access to a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SnapshotsClient) revokeAccess(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsClientBeginRevokeAccessOptions) (*http.Response, error) {
	req, err := client.revokeAccessCreateRequest(ctx, resourceGroupName, snapshotName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// revokeAccessCreateRequest creates the RevokeAccess request.
func (client *SnapshotsClient) revokeAccessCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsClientBeginRevokeAccessOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}/endGetAccess"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginUpdate - Updates (patches) a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// snapshotName - The name of the snapshot that is being created. The name can't be changed after the snapshot is created.
// Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80
// characters.
// snapshot - Snapshot object supplied in the body of the Patch snapshot operation.
// options - SnapshotsClientBeginUpdateOptions contains the optional parameters for the SnapshotsClient.BeginUpdate method.
func (client *SnapshotsClient) BeginUpdate(ctx context.Context, resourceGroupName string, snapshotName string, snapshot SnapshotUpdate, options *SnapshotsClientBeginUpdateOptions) (SnapshotsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, snapshotName, snapshot, options)
	if err != nil {
		return SnapshotsClientUpdatePollerResponse{}, err
	}
	result := SnapshotsClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SnapshotsClient.Update", "", resp, client.pl)
	if err != nil {
		return SnapshotsClientUpdatePollerResponse{}, err
	}
	result.Poller = &SnapshotsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates (patches) a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SnapshotsClient) update(ctx context.Context, resourceGroupName string, snapshotName string, snapshot SnapshotUpdate, options *SnapshotsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, snapshotName, snapshot, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *SnapshotsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, snapshot SnapshotUpdate, options *SnapshotsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, snapshot)
}
