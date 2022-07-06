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

// NatRulesClient contains the methods for the NatRules group.
// Don't use this type directly, use NewNatRulesClient() instead.
type NatRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewNatRulesClient creates a new instance of NatRulesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewNatRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *NatRulesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &NatRulesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates a nat rule to a scalable vpn gateway if it doesn't exist else updates the existing nat rules.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnGateway.
// gatewayName - The name of the gateway.
// natRuleName - The name of the nat rule.
// natRuleParameters - Parameters supplied to create or Update a Nat Rule.
// options - NatRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the NatRulesClient.BeginCreateOrUpdate
// method.
func (client *NatRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, natRuleParameters VPNGatewayNatRule, options *NatRulesClientBeginCreateOrUpdateOptions) (NatRulesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, gatewayName, natRuleName, natRuleParameters, options)
	if err != nil {
		return NatRulesClientCreateOrUpdatePollerResponse{}, err
	}
	result := NatRulesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("NatRulesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return NatRulesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &NatRulesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a nat rule to a scalable vpn gateway if it doesn't exist else updates the existing nat rules.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *NatRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, natRuleParameters VPNGatewayNatRule, options *NatRulesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, gatewayName, natRuleName, natRuleParameters, options)
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
func (client *NatRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, natRuleParameters VPNGatewayNatRule, options *NatRulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, natRuleParameters)
}

// BeginDelete - Deletes a nat rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnGateway.
// gatewayName - The name of the gateway.
// natRuleName - The name of the nat rule.
// options - NatRulesClientBeginDeleteOptions contains the optional parameters for the NatRulesClient.BeginDelete method.
func (client *NatRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, options *NatRulesClientBeginDeleteOptions) (NatRulesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, gatewayName, natRuleName, options)
	if err != nil {
		return NatRulesClientDeletePollerResponse{}, err
	}
	result := NatRulesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("NatRulesClient.Delete", "location", resp, client.pl)
	if err != nil {
		return NatRulesClientDeletePollerResponse{}, err
	}
	result.Poller = &NatRulesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a nat rule.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *NatRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, options *NatRulesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, gatewayName, natRuleName, options)
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
func (client *NatRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, options *NatRulesClientBeginDeleteOptions) (*policy.Request, error) {
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

// Get - Retrieves the details of a nat ruleGet.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnGateway.
// gatewayName - The name of the gateway.
// natRuleName - The name of the nat rule.
// options - NatRulesClientGetOptions contains the optional parameters for the NatRulesClient.Get method.
func (client *NatRulesClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, options *NatRulesClientGetOptions) (NatRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, natRuleName, options)
	if err != nil {
		return NatRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NatRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NatRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NatRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, options *NatRulesClientGetOptions) (*policy.Request, error) {
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
func (client *NatRulesClient) getHandleResponse(resp *http.Response) (NatRulesClientGetResponse, error) {
	result := NatRulesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNGatewayNatRule); err != nil {
		return NatRulesClientGetResponse{}, err
	}
	return result, nil
}

// ListByVPNGateway - Retrieves all nat rules for a particular virtual wan vpn gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnGateway.
// gatewayName - The name of the gateway.
// options - NatRulesClientListByVPNGatewayOptions contains the optional parameters for the NatRulesClient.ListByVPNGateway
// method.
func (client *NatRulesClient) ListByVPNGateway(resourceGroupName string, gatewayName string, options *NatRulesClientListByVPNGatewayOptions) *NatRulesClientListByVPNGatewayPager {
	return &NatRulesClientListByVPNGatewayPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByVPNGatewayCreateRequest(ctx, resourceGroupName, gatewayName, options)
		},
		advancer: func(ctx context.Context, resp NatRulesClientListByVPNGatewayResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVPNGatewayNatRulesResult.NextLink)
		},
	}
}

// listByVPNGatewayCreateRequest creates the ListByVPNGateway request.
func (client *NatRulesClient) listByVPNGatewayCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *NatRulesClientListByVPNGatewayOptions) (*policy.Request, error) {
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

// listByVPNGatewayHandleResponse handles the ListByVPNGateway response.
func (client *NatRulesClient) listByVPNGatewayHandleResponse(resp *http.Response) (NatRulesClientListByVPNGatewayResponse, error) {
	result := NatRulesClientListByVPNGatewayResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNGatewayNatRulesResult); err != nil {
		return NatRulesClientListByVPNGatewayResponse{}, err
	}
	return result, nil
}
