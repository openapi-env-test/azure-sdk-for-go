//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

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

// SecurityPartnerProvidersClient contains the methods for the SecurityPartnerProviders group.
// Don't use this type directly, use NewSecurityPartnerProvidersClient() instead.
type SecurityPartnerProvidersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSecurityPartnerProvidersClient creates a new instance of SecurityPartnerProvidersClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSecurityPartnerProvidersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SecurityPartnerProvidersClient, error) {
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
	client := &SecurityPartnerProvidersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the specified Security Partner Provider.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group.
// securityPartnerProviderName - The name of the Security Partner Provider.
// parameters - Parameters supplied to the create or update Security Partner Provider operation.
// options - SecurityPartnerProvidersClientBeginCreateOrUpdateOptions contains the optional parameters for the SecurityPartnerProvidersClient.BeginCreateOrUpdate
// method.
func (client *SecurityPartnerProvidersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters SecurityPartnerProvider, options *SecurityPartnerProvidersClientBeginCreateOrUpdateOptions) (*runtime.Poller[SecurityPartnerProvidersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, securityPartnerProviderName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SecurityPartnerProvidersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SecurityPartnerProvidersClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates the specified Security Partner Provider.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
func (client *SecurityPartnerProvidersClient) createOrUpdate(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters SecurityPartnerProvider, options *SecurityPartnerProvidersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, securityPartnerProviderName, parameters, options)
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
func (client *SecurityPartnerProvidersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters SecurityPartnerProvider, options *SecurityPartnerProvidersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders/{securityPartnerProviderName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityPartnerProviderName == "" {
		return nil, errors.New("parameter securityPartnerProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityPartnerProviderName}", url.PathEscape(securityPartnerProviderName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified Security Partner Provider.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group.
// securityPartnerProviderName - The name of the Security Partner Provider.
// options - SecurityPartnerProvidersClientBeginDeleteOptions contains the optional parameters for the SecurityPartnerProvidersClient.BeginDelete
// method.
func (client *SecurityPartnerProvidersClient) BeginDelete(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, options *SecurityPartnerProvidersClientBeginDeleteOptions) (*runtime.Poller[SecurityPartnerProvidersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, securityPartnerProviderName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SecurityPartnerProvidersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SecurityPartnerProvidersClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified Security Partner Provider.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
func (client *SecurityPartnerProvidersClient) deleteOperation(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, options *SecurityPartnerProvidersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, securityPartnerProviderName, options)
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
func (client *SecurityPartnerProvidersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, options *SecurityPartnerProvidersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders/{securityPartnerProviderName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityPartnerProviderName == "" {
		return nil, errors.New("parameter securityPartnerProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityPartnerProviderName}", url.PathEscape(securityPartnerProviderName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified Security Partner Provider.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group.
// securityPartnerProviderName - The name of the Security Partner Provider.
// options - SecurityPartnerProvidersClientGetOptions contains the optional parameters for the SecurityPartnerProvidersClient.Get
// method.
func (client *SecurityPartnerProvidersClient) Get(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, options *SecurityPartnerProvidersClientGetOptions) (SecurityPartnerProvidersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, securityPartnerProviderName, options)
	if err != nil {
		return SecurityPartnerProvidersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SecurityPartnerProvidersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SecurityPartnerProvidersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SecurityPartnerProvidersClient) getCreateRequest(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, options *SecurityPartnerProvidersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders/{securityPartnerProviderName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityPartnerProviderName == "" {
		return nil, errors.New("parameter securityPartnerProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityPartnerProviderName}", url.PathEscape(securityPartnerProviderName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SecurityPartnerProvidersClient) getHandleResponse(resp *http.Response) (SecurityPartnerProvidersClientGetResponse, error) {
	result := SecurityPartnerProvidersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityPartnerProvider); err != nil {
		return SecurityPartnerProvidersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all the Security Partner Providers in a subscription.
// Generated from API version 2022-07-01
// options - SecurityPartnerProvidersClientListOptions contains the optional parameters for the SecurityPartnerProvidersClient.List
// method.
func (client *SecurityPartnerProvidersClient) NewListPager(options *SecurityPartnerProvidersClientListOptions) *runtime.Pager[SecurityPartnerProvidersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecurityPartnerProvidersClientListResponse]{
		More: func(page SecurityPartnerProvidersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecurityPartnerProvidersClientListResponse) (SecurityPartnerProvidersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SecurityPartnerProvidersClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SecurityPartnerProvidersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SecurityPartnerProvidersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SecurityPartnerProvidersClient) listCreateRequest(ctx context.Context, options *SecurityPartnerProvidersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/securityPartnerProviders"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SecurityPartnerProvidersClient) listHandleResponse(resp *http.Response) (SecurityPartnerProvidersClientListResponse, error) {
	result := SecurityPartnerProvidersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityPartnerProviderListResult); err != nil {
		return SecurityPartnerProvidersClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all Security Partner Providers in a resource group.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group.
// options - SecurityPartnerProvidersClientListByResourceGroupOptions contains the optional parameters for the SecurityPartnerProvidersClient.ListByResourceGroup
// method.
func (client *SecurityPartnerProvidersClient) NewListByResourceGroupPager(resourceGroupName string, options *SecurityPartnerProvidersClientListByResourceGroupOptions) *runtime.Pager[SecurityPartnerProvidersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecurityPartnerProvidersClientListByResourceGroupResponse]{
		More: func(page SecurityPartnerProvidersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecurityPartnerProvidersClientListByResourceGroupResponse) (SecurityPartnerProvidersClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SecurityPartnerProvidersClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SecurityPartnerProvidersClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SecurityPartnerProvidersClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SecurityPartnerProvidersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SecurityPartnerProvidersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SecurityPartnerProvidersClient) listByResourceGroupHandleResponse(resp *http.Response) (SecurityPartnerProvidersClientListByResourceGroupResponse, error) {
	result := SecurityPartnerProvidersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityPartnerProviderListResult); err != nil {
		return SecurityPartnerProvidersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates tags of a Security Partner Provider resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group.
// securityPartnerProviderName - The name of the Security Partner Provider.
// parameters - Parameters supplied to update Security Partner Provider tags.
// options - SecurityPartnerProvidersClientUpdateTagsOptions contains the optional parameters for the SecurityPartnerProvidersClient.UpdateTags
// method.
func (client *SecurityPartnerProvidersClient) UpdateTags(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters TagsObject, options *SecurityPartnerProvidersClientUpdateTagsOptions) (SecurityPartnerProvidersClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, securityPartnerProviderName, parameters, options)
	if err != nil {
		return SecurityPartnerProvidersClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SecurityPartnerProvidersClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SecurityPartnerProvidersClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *SecurityPartnerProvidersClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters TagsObject, options *SecurityPartnerProvidersClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders/{securityPartnerProviderName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityPartnerProviderName == "" {
		return nil, errors.New("parameter securityPartnerProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityPartnerProviderName}", url.PathEscape(securityPartnerProviderName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *SecurityPartnerProvidersClient) updateTagsHandleResponse(resp *http.Response) (SecurityPartnerProvidersClientUpdateTagsResponse, error) {
	result := SecurityPartnerProvidersClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityPartnerProvider); err != nil {
		return SecurityPartnerProvidersClientUpdateTagsResponse{}, err
	}
	return result, nil
}
