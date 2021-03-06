// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// LoadBalancerOutboundRulesClient contains the methods for the LoadBalancerOutboundRules group.
// Don't use this type directly, use NewLoadBalancerOutboundRulesClient() instead.
type LoadBalancerOutboundRulesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewLoadBalancerOutboundRulesClient creates a new instance of LoadBalancerOutboundRulesClient with the specified values.
func NewLoadBalancerOutboundRulesClient(con *armcore.Connection, subscriptionID string) *LoadBalancerOutboundRulesClient {
	return &LoadBalancerOutboundRulesClient{con: con, subscriptionID: subscriptionID}
}

// Get - Gets the specified load balancer outbound rule.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancerOutboundRulesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, outboundRuleName string, options *LoadBalancerOutboundRulesGetOptions) (OutboundRuleResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, outboundRuleName, options)
	if err != nil {
		return OutboundRuleResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return OutboundRuleResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return OutboundRuleResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LoadBalancerOutboundRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, outboundRuleName string, options *LoadBalancerOutboundRulesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules/{outboundRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if outboundRuleName == "" {
		return nil, errors.New("parameter outboundRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outboundRuleName}", url.PathEscape(outboundRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LoadBalancerOutboundRulesClient) getHandleResponse(resp *azcore.Response) (OutboundRuleResponse, error) {
	var val *OutboundRule
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return OutboundRuleResponse{}, err
	}
	return OutboundRuleResponse{RawResponse: resp.Response, OutboundRule: val}, nil
}

// getHandleError handles the Get error response.
func (client *LoadBalancerOutboundRulesClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Gets all the outbound rules in a load balancer.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancerOutboundRulesClient) List(resourceGroupName string, loadBalancerName string, options *LoadBalancerOutboundRulesListOptions) LoadBalancerOutboundRuleListResultPager {
	return &loadBalancerOutboundRuleListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp LoadBalancerOutboundRuleListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerOutboundRuleListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *LoadBalancerOutboundRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancerOutboundRulesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LoadBalancerOutboundRulesClient) listHandleResponse(resp *azcore.Response) (LoadBalancerOutboundRuleListResultResponse, error) {
	var val *LoadBalancerOutboundRuleListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LoadBalancerOutboundRuleListResultResponse{}, err
	}
	return LoadBalancerOutboundRuleListResultResponse{RawResponse: resp.Response, LoadBalancerOutboundRuleListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *LoadBalancerOutboundRulesClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
