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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// NetworkInterfaceLoadBalancersClient contains the methods for the NetworkInterfaceLoadBalancers group.
// Don't use this type directly, use NewNetworkInterfaceLoadBalancersClient() instead.
type NetworkInterfaceLoadBalancersClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewNetworkInterfaceLoadBalancersClient creates a new instance of NetworkInterfaceLoadBalancersClient with the specified values.
func NewNetworkInterfaceLoadBalancersClient(con *arm.Connection, subscriptionID string) *NetworkInterfaceLoadBalancersClient {
	return &NetworkInterfaceLoadBalancersClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// List - List all load balancers in a network interface.
// If the operation fails it returns the *CloudError error type.
func (client *NetworkInterfaceLoadBalancersClient) List(resourceGroupName string, networkInterfaceName string, options *NetworkInterfaceLoadBalancersListOptions) *NetworkInterfaceLoadBalancersListPager {
	return &NetworkInterfaceLoadBalancersListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, networkInterfaceName, options)
		},
		advancer: func(ctx context.Context, resp NetworkInterfaceLoadBalancersListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.NetworkInterfaceLoadBalancerListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *NetworkInterfaceLoadBalancersClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, options *NetworkInterfaceLoadBalancersListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/loadBalancers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkInterfaceName == "" {
		return nil, errors.New("parameter networkInterfaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listHandleResponse handles the List response.
func (client *NetworkInterfaceLoadBalancersClient) listHandleResponse(resp *http.Response) (NetworkInterfaceLoadBalancersListResponse, error) {
	result := NetworkInterfaceLoadBalancersListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkInterfaceLoadBalancerListResult); err != nil {
		return NetworkInterfaceLoadBalancersListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *NetworkInterfaceLoadBalancersClient) listHandleError(resp *http.Response) error {
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
