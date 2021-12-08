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
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// NatRulesClient contains the methods for the NatRules group.
// Don't use this type directly, use NewNatRulesClient() instead.
type NatRulesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewNatRulesClient creates a new instance of NatRulesClient with the specified values.
func NewNatRulesClient(con *arm.Connection, subscriptionID string) *NatRulesClient {
	return &NatRulesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates a nat rule to a scalable vpn gateway if it doesn't exist else updates the existing nat rules.
// If the operation fails it returns the *CloudError error type.
func (client *NatRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, natRuleParameters VPNGatewayNatRule, options *NatRulesBeginCreateOrUpdateOptions) (NatRulesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, gatewayName, natRuleName, natRuleParameters, options)
	if err != nil {
		return NatRulesCreateOrUpdatePollerResponse{}, err
	}
	result := NatRulesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("NatRulesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return NatRulesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &NatRulesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a nat rule to a scalable vpn gateway if it doesn't exist else updates the existing nat rules.
// If the operation fails it returns the *CloudError error type.
func (client *NatRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, natRuleParameters VPNGatewayNatRule, options *NatRulesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, gatewayName, natRuleName, natRuleParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NatRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, natRuleParameters VPNGatewayNatRule, options *NatRulesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/natRules/{natRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if natRuleName == "" {
		return nil, errors.New("parameter natRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natRuleName}", url.PathEscape(natRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, natRuleParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *NatRulesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes a nat rule.
// If the operation fails it returns the *CloudError error type.
func (client *NatRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, options *NatRulesBeginDeleteOptions) (NatRulesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, gatewayName, natRuleName, options)
	if err != nil {
		return NatRulesDeletePollerResponse{}, err
	}
	result := NatRulesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("NatRulesClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return NatRulesDeletePollerResponse{}, err
	}
	result.Poller = &NatRulesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a nat rule.
// If the operation fails it returns the *CloudError error type.
func (client *NatRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, options *NatRulesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, gatewayName, natRuleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NatRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, options *NatRulesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/natRules/{natRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if natRuleName == "" {
		return nil, errors.New("parameter natRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natRuleName}", url.PathEscape(natRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *NatRulesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Retrieves the details of a nat ruleGet.
// If the operation fails it returns the *CloudError error type.
func (client *NatRulesClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, options *NatRulesGetOptions) (NatRulesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, natRuleName, options)
	if err != nil {
		return NatRulesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NatRulesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NatRulesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NatRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, options *NatRulesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/natRules/{natRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if natRuleName == "" {
		return nil, errors.New("parameter natRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natRuleName}", url.PathEscape(natRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NatRulesClient) getHandleResponse(resp *http.Response) (NatRulesGetResponse, error) {
	result := NatRulesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNGatewayNatRule); err != nil {
		return NatRulesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *NatRulesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByVPNGateway - Retrieves all nat rules for a particular virtual wan vpn gateway.
// If the operation fails it returns the *CloudError error type.
func (client *NatRulesClient) ListByVPNGateway(resourceGroupName string, gatewayName string, options *NatRulesListByVPNGatewayOptions) *NatRulesListByVPNGatewayPager {
	return &NatRulesListByVPNGatewayPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByVPNGatewayCreateRequest(ctx, resourceGroupName, gatewayName, options)
		},
		advancer: func(ctx context.Context, resp NatRulesListByVPNGatewayResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVPNGatewayNatRulesResult.NextLink)
		},
	}
}

// listByVPNGatewayCreateRequest creates the ListByVPNGateway request.
func (client *NatRulesClient) listByVPNGatewayCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *NatRulesListByVPNGatewayOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/natRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByVPNGatewayHandleResponse handles the ListByVPNGateway response.
func (client *NatRulesClient) listByVPNGatewayHandleResponse(resp *http.Response) (NatRulesListByVPNGatewayResponse, error) {
	result := NatRulesListByVPNGatewayResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNGatewayNatRulesResult); err != nil {
		return NatRulesListByVPNGatewayResponse{}, err
	}
	return result, nil
}

// listByVPNGatewayHandleError handles the ListByVPNGateway error response.
func (client *NatRulesClient) listByVPNGatewayHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
