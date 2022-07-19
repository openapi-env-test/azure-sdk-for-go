//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// DomainWhoisClient contains the methods for the DomainWhois group.
// Don't use this type directly, use NewDomainWhoisClient() instead.
type DomainWhoisClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDomainWhoisClient creates a new instance of DomainWhoisClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDomainWhoisClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DomainWhoisClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &DomainWhoisClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Get whois information for a single domain name
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// domain - Domain name to be enriched
// options - DomainWhoisClientGetOptions contains the optional parameters for the DomainWhoisClient.Get method.
func (client *DomainWhoisClient) Get(ctx context.Context, resourceGroupName string, domain string, options *DomainWhoisClientGetOptions) (DomainWhoisClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, domain, options)
	if err != nil {
		return DomainWhoisClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DomainWhoisClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DomainWhoisClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DomainWhoisClient) getCreateRequest(ctx context.Context, resourceGroupName string, domain string, options *DomainWhoisClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityInsights/enrichment/domain/whois/"
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
	reqQP.Set("api-version", "2022-07-01-preview")
	reqQP.Set("domain", domain)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DomainWhoisClient) getHandleResponse(resp *http.Response) (DomainWhoisClientGetResponse, error) {
	result := DomainWhoisClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnrichmentDomainWhois); err != nil {
		return DomainWhoisClientGetResponse{}, err
	}
	return result, nil
}
