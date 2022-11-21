//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// APICollectionClient contains the methods for the APICollection group.
// Don't use this type directly, use NewAPICollectionClient() instead.
type APICollectionClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAPICollectionClient creates a new instance of APICollectionClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAPICollectionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*APICollectionClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &APICollectionClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets an Azure API Management API if it has been onboarded to Defender for APIs. If an Azure API Management API is
// onboarded to Defender for APIs, the system will monitor the operations within the
// Azure API Management API for intrusive behaviors and provide alerts for attacks that have been detected.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-20-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// apiCollectionID - A string representing the apiCollections resource within the Microsoft.Security provider namespace. This
// string matches the Azure API Management API name.
// options - APICollectionClientGetOptions contains the optional parameters for the APICollectionClient.Get method.
func (client *APICollectionClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiCollectionID string, options *APICollectionClientGetOptions) (APICollectionClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, apiCollectionID, options)
	if err != nil {
		return APICollectionClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APICollectionClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APICollectionClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *APICollectionClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiCollectionID string, options *APICollectionClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/providers/Microsoft.Security/apiCollections/{apiCollectionId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiCollectionID == "" {
		return nil, errors.New("parameter apiCollectionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiCollectionId}", url.PathEscape(apiCollectionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *APICollectionClient) getHandleResponse(resp *http.Response) (APICollectionClientGetResponse, error) {
	result := APICollectionClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APICollectionResponse); err != nil {
		return APICollectionClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of Azure API Management APIs that have been onboarded to Defender for APIs. If an Azure API
// Management API is onboarded to Defender for APIs, the system will monitor the operations within
// the Azure API Management API for intrusive behaviors and provide alerts for attacks that have been detected.
// Generated from API version 2022-11-20-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceName - The name of the API Management service.
// options - APICollectionClientListOptions contains the optional parameters for the APICollectionClient.List method.
func (client *APICollectionClient) NewListPager(resourceGroupName string, serviceName string, options *APICollectionClientListOptions) *runtime.Pager[APICollectionClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[APICollectionClientListResponse]{
		More: func(page APICollectionClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *APICollectionClientListResponse) (APICollectionClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return APICollectionClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return APICollectionClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return APICollectionClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *APICollectionClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *APICollectionClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/providers/Microsoft.Security/apiCollections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *APICollectionClient) listHandleResponse(resp *http.Response) (APICollectionClientListResponse, error) {
	result := APICollectionClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APICollectionResponseList); err != nil {
		return APICollectionClientListResponse{}, err
	}
	return result, nil
}
