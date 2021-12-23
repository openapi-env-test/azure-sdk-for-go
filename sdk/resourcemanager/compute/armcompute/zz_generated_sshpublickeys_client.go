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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SSHPublicKeysClient contains the methods for the SSHPublicKeys group.
// Don't use this type directly, use NewSSHPublicKeysClient() instead.
type SSHPublicKeysClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSSHPublicKeysClient creates a new instance of SSHPublicKeysClient with the specified values.
func NewSSHPublicKeysClient(con *arm.Connection, subscriptionID string) *SSHPublicKeysClient {
	return &SSHPublicKeysClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Create - Creates a new SSH public key resource.
// If the operation fails it returns a generic error.
func (client *SSHPublicKeysClient) Create(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyResource, options *SSHPublicKeysCreateOptions) (SSHPublicKeysCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, sshPublicKeyName, parameters, options)
	if err != nil {
		return SSHPublicKeysCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SSHPublicKeysCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SSHPublicKeysCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *SSHPublicKeysClient) createCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyResource, options *SSHPublicKeysCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sshPublicKeyName == "" {
		return nil, errors.New("parameter sshPublicKeyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *SSHPublicKeysClient) createHandleResponse(resp *http.Response) (SSHPublicKeysCreateResponse, error) {
	result := SSHPublicKeysCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SSHPublicKeyResource); err != nil {
		return SSHPublicKeysCreateResponse{}, err
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *SSHPublicKeysClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Delete - Delete an SSH public key.
// If the operation fails it returns a generic error.
func (client *SSHPublicKeysClient) Delete(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysDeleteOptions) (SSHPublicKeysDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sshPublicKeyName, options)
	if err != nil {
		return SSHPublicKeysDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SSHPublicKeysDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SSHPublicKeysDeleteResponse{}, client.deleteHandleError(resp)
	}
	return SSHPublicKeysDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SSHPublicKeysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sshPublicKeyName == "" {
		return nil, errors.New("parameter sshPublicKeyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SSHPublicKeysClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// GenerateKeyPair - Generates and returns a public/private key pair and populates the SSH public key resource with the public key. The length of the key
// will be 3072 bits. This operation can only be performed once per
// SSH public key resource.
// If the operation fails it returns a generic error.
func (client *SSHPublicKeysClient) GenerateKeyPair(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysGenerateKeyPairOptions) (SSHPublicKeysGenerateKeyPairResponse, error) {
	req, err := client.generateKeyPairCreateRequest(ctx, resourceGroupName, sshPublicKeyName, options)
	if err != nil {
		return SSHPublicKeysGenerateKeyPairResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SSHPublicKeysGenerateKeyPairResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SSHPublicKeysGenerateKeyPairResponse{}, client.generateKeyPairHandleError(resp)
	}
	return client.generateKeyPairHandleResponse(resp)
}

// generateKeyPairCreateRequest creates the GenerateKeyPair request.
func (client *SSHPublicKeysClient) generateKeyPairCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysGenerateKeyPairOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}/generateKeyPair"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sshPublicKeyName == "" {
		return nil, errors.New("parameter sshPublicKeyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// generateKeyPairHandleResponse handles the GenerateKeyPair response.
func (client *SSHPublicKeysClient) generateKeyPairHandleResponse(resp *http.Response) (SSHPublicKeysGenerateKeyPairResponse, error) {
	result := SSHPublicKeysGenerateKeyPairResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SSHPublicKeyGenerateKeyPairResult); err != nil {
		return SSHPublicKeysGenerateKeyPairResponse{}, err
	}
	return result, nil
}

// generateKeyPairHandleError handles the GenerateKeyPair error response.
func (client *SSHPublicKeysClient) generateKeyPairHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Retrieves information about an SSH public key.
// If the operation fails it returns a generic error.
func (client *SSHPublicKeysClient) Get(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysGetOptions) (SSHPublicKeysGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, sshPublicKeyName, options)
	if err != nil {
		return SSHPublicKeysGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SSHPublicKeysGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SSHPublicKeysGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SSHPublicKeysClient) getCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sshPublicKeyName == "" {
		return nil, errors.New("parameter sshPublicKeyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SSHPublicKeysClient) getHandleResponse(resp *http.Response) (SSHPublicKeysGetResponse, error) {
	result := SSHPublicKeysGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SSHPublicKeyResource); err != nil {
		return SSHPublicKeysGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SSHPublicKeysClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByResourceGroup - Lists all of the SSH public keys in the specified resource group. Use the nextLink property in the response to get the next page
// of SSH public keys.
// If the operation fails it returns a generic error.
func (client *SSHPublicKeysClient) ListByResourceGroup(resourceGroupName string, options *SSHPublicKeysListByResourceGroupOptions) *SSHPublicKeysListByResourceGroupPager {
	return &SSHPublicKeysListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp SSHPublicKeysListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SSHPublicKeysGroupListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SSHPublicKeysClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SSHPublicKeysListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SSHPublicKeysClient) listByResourceGroupHandleResponse(resp *http.Response) (SSHPublicKeysListByResourceGroupResponse, error) {
	result := SSHPublicKeysListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SSHPublicKeysGroupListResult); err != nil {
		return SSHPublicKeysListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *SSHPublicKeysClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListBySubscription - Lists all of the SSH public keys in the subscription. Use the nextLink property in the response to get the next page of SSH public
// keys.
// If the operation fails it returns a generic error.
func (client *SSHPublicKeysClient) ListBySubscription(options *SSHPublicKeysListBySubscriptionOptions) *SSHPublicKeysListBySubscriptionPager {
	return &SSHPublicKeysListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SSHPublicKeysListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SSHPublicKeysGroupListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SSHPublicKeysClient) listBySubscriptionCreateRequest(ctx context.Context, options *SSHPublicKeysListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/sshPublicKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SSHPublicKeysClient) listBySubscriptionHandleResponse(resp *http.Response) (SSHPublicKeysListBySubscriptionResponse, error) {
	result := SSHPublicKeysListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SSHPublicKeysGroupListResult); err != nil {
		return SSHPublicKeysListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *SSHPublicKeysClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Update - Updates a new SSH public key resource.
// If the operation fails it returns a generic error.
func (client *SSHPublicKeysClient) Update(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyUpdateResource, options *SSHPublicKeysUpdateOptions) (SSHPublicKeysUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sshPublicKeyName, parameters, options)
	if err != nil {
		return SSHPublicKeysUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SSHPublicKeysUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SSHPublicKeysUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SSHPublicKeysClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyUpdateResource, options *SSHPublicKeysUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sshPublicKeyName == "" {
		return nil, errors.New("parameter sshPublicKeyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *SSHPublicKeysClient) updateHandleResponse(resp *http.Response) (SSHPublicKeysUpdateResponse, error) {
	result := SSHPublicKeysUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SSHPublicKeyResource); err != nil {
		return SSHPublicKeysUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *SSHPublicKeysClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
