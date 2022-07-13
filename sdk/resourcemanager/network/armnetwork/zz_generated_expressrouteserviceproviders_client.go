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

// ExpressRouteServiceProvidersClient contains the methods for the ExpressRouteServiceProviders group.
// Don't use this type directly, use NewExpressRouteServiceProvidersClient() instead.
type ExpressRouteServiceProvidersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExpressRouteServiceProvidersClient creates a new instance of ExpressRouteServiceProvidersClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExpressRouteServiceProvidersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExpressRouteServiceProvidersClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ExpressRouteServiceProvidersClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// List - Gets all the available express route service providers.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ExpressRouteServiceProvidersClientListOptions contains the optional parameters for the ExpressRouteServiceProvidersClient.List
// method.
func (client *ExpressRouteServiceProvidersClient) List(options *ExpressRouteServiceProvidersClientListOptions) *ExpressRouteServiceProvidersClientListPager {
	return &ExpressRouteServiceProvidersClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ExpressRouteServiceProvidersClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteServiceProviderListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ExpressRouteServiceProvidersClient) listCreateRequest(ctx context.Context, options *ExpressRouteServiceProvidersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteServiceProviders"
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
func (client *ExpressRouteServiceProvidersClient) listHandleResponse(resp *http.Response) (ExpressRouteServiceProvidersClientListResponse, error) {
	result := ExpressRouteServiceProvidersClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteServiceProviderListResult); err != nil {
		return ExpressRouteServiceProvidersClientListResponse{}, err
	}
	return result, nil
}
