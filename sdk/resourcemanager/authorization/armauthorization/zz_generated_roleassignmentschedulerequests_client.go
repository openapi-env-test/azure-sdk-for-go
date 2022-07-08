//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// RoleAssignmentScheduleRequestsClient contains the methods for the RoleAssignmentScheduleRequests group.
// Don't use this type directly, use NewRoleAssignmentScheduleRequestsClient() instead.
type RoleAssignmentScheduleRequestsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewRoleAssignmentScheduleRequestsClient creates a new instance of RoleAssignmentScheduleRequestsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRoleAssignmentScheduleRequestsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *RoleAssignmentScheduleRequestsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &RoleAssignmentScheduleRequestsClient{
		host: string(cp.Endpoint),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Cancel - Cancels a pending role assignment schedule request.
// If the operation fails it returns an *azcore.ResponseError type.
// scope - The scope of the role assignment request to cancel.
// roleAssignmentScheduleRequestName - The name of the role assignment request to cancel.
// options - RoleAssignmentScheduleRequestsClientCancelOptions contains the optional parameters for the RoleAssignmentScheduleRequestsClient.Cancel
// method.
func (client *RoleAssignmentScheduleRequestsClient) Cancel(ctx context.Context, scope string, roleAssignmentScheduleRequestName string, options *RoleAssignmentScheduleRequestsClientCancelOptions) (RoleAssignmentScheduleRequestsClientCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, scope, roleAssignmentScheduleRequestName, options)
	if err != nil {
		return RoleAssignmentScheduleRequestsClientCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentScheduleRequestsClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentScheduleRequestsClientCancelResponse{}, runtime.NewResponseError(resp)
	}
	return RoleAssignmentScheduleRequestsClientCancelResponse{RawResponse: resp}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *RoleAssignmentScheduleRequestsClient) cancelCreateRequest(ctx context.Context, scope string, roleAssignmentScheduleRequestName string, options *RoleAssignmentScheduleRequestsClientCancelOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/{roleAssignmentScheduleRequestName}/cancel"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentScheduleRequestName == "" {
		return nil, errors.New("parameter roleAssignmentScheduleRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentScheduleRequestName}", url.PathEscape(roleAssignmentScheduleRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Create - Creates a role assignment schedule request.
// If the operation fails it returns an *azcore.ResponseError type.
// scope - The scope of the role assignment schedule request to create. The scope can be any REST resource instance. For example,
// use '/providers/Microsoft.Subscription/subscriptions/{subscription-id}/' for a
// subscription, '/providers/Microsoft.Subscription/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}'
// for a resource group, and
// '/providers/Microsoft.Subscription/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}'
// for a resource.
// roleAssignmentScheduleRequestName - A GUID for the role assignment to create. The name must be unique and different for
// each role assignment.
// parameters - Parameters for the role assignment schedule request.
// options - RoleAssignmentScheduleRequestsClientCreateOptions contains the optional parameters for the RoleAssignmentScheduleRequestsClient.Create
// method.
func (client *RoleAssignmentScheduleRequestsClient) Create(ctx context.Context, scope string, roleAssignmentScheduleRequestName string, parameters RoleAssignmentScheduleRequest, options *RoleAssignmentScheduleRequestsClientCreateOptions) (RoleAssignmentScheduleRequestsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, scope, roleAssignmentScheduleRequestName, parameters, options)
	if err != nil {
		return RoleAssignmentScheduleRequestsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentScheduleRequestsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return RoleAssignmentScheduleRequestsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *RoleAssignmentScheduleRequestsClient) createCreateRequest(ctx context.Context, scope string, roleAssignmentScheduleRequestName string, parameters RoleAssignmentScheduleRequest, options *RoleAssignmentScheduleRequestsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/{roleAssignmentScheduleRequestName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentScheduleRequestName == "" {
		return nil, errors.New("parameter roleAssignmentScheduleRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentScheduleRequestName}", url.PathEscape(roleAssignmentScheduleRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *RoleAssignmentScheduleRequestsClient) createHandleResponse(resp *http.Response) (RoleAssignmentScheduleRequestsClientCreateResponse, error) {
	result := RoleAssignmentScheduleRequestsClientCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentScheduleRequest); err != nil {
		return RoleAssignmentScheduleRequestsClientCreateResponse{}, err
	}
	return result, nil
}

// Get - Get the specified role assignment schedule request.
// If the operation fails it returns an *azcore.ResponseError type.
// scope - The scope of the role assignment schedule request.
// roleAssignmentScheduleRequestName - The name (guid) of the role assignment schedule request to get.
// options - RoleAssignmentScheduleRequestsClientGetOptions contains the optional parameters for the RoleAssignmentScheduleRequestsClient.Get
// method.
func (client *RoleAssignmentScheduleRequestsClient) Get(ctx context.Context, scope string, roleAssignmentScheduleRequestName string, options *RoleAssignmentScheduleRequestsClientGetOptions) (RoleAssignmentScheduleRequestsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, roleAssignmentScheduleRequestName, options)
	if err != nil {
		return RoleAssignmentScheduleRequestsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentScheduleRequestsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentScheduleRequestsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleAssignmentScheduleRequestsClient) getCreateRequest(ctx context.Context, scope string, roleAssignmentScheduleRequestName string, options *RoleAssignmentScheduleRequestsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/{roleAssignmentScheduleRequestName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentScheduleRequestName == "" {
		return nil, errors.New("parameter roleAssignmentScheduleRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentScheduleRequestName}", url.PathEscape(roleAssignmentScheduleRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleAssignmentScheduleRequestsClient) getHandleResponse(resp *http.Response) (RoleAssignmentScheduleRequestsClientGetResponse, error) {
	result := RoleAssignmentScheduleRequestsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentScheduleRequest); err != nil {
		return RoleAssignmentScheduleRequestsClientGetResponse{}, err
	}
	return result, nil
}

// ListForScope - Gets role assignment schedule requests for a scope.
// If the operation fails it returns an *azcore.ResponseError type.
// scope - The scope of the role assignments schedule requests.
// options - RoleAssignmentScheduleRequestsClientListForScopeOptions contains the optional parameters for the RoleAssignmentScheduleRequestsClient.ListForScope
// method.
func (client *RoleAssignmentScheduleRequestsClient) ListForScope(scope string, options *RoleAssignmentScheduleRequestsClientListForScopeOptions) *RoleAssignmentScheduleRequestsClientListForScopePager {
	return &RoleAssignmentScheduleRequestsClientListForScopePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listForScopeCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp RoleAssignmentScheduleRequestsClientListForScopeResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RoleAssignmentScheduleRequestListResult.NextLink)
		},
	}
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleAssignmentScheduleRequestsClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleAssignmentScheduleRequestsClientListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignmentScheduleRequests"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2022-06-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleAssignmentScheduleRequestsClient) listForScopeHandleResponse(resp *http.Response) (RoleAssignmentScheduleRequestsClientListForScopeResponse, error) {
	result := RoleAssignmentScheduleRequestsClientListForScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentScheduleRequestListResult); err != nil {
		return RoleAssignmentScheduleRequestsClientListForScopeResponse{}, err
	}
	return result, nil
}

// Validate - Validates a new role assignment schedule request.
// If the operation fails it returns an *azcore.ResponseError type.
// scope - The scope of the role assignment request to validate.
// roleAssignmentScheduleRequestName - The name of the role assignment request to validate.
// parameters - Parameters for the role assignment schedule request.
// options - RoleAssignmentScheduleRequestsClientValidateOptions contains the optional parameters for the RoleAssignmentScheduleRequestsClient.Validate
// method.
func (client *RoleAssignmentScheduleRequestsClient) Validate(ctx context.Context, scope string, roleAssignmentScheduleRequestName string, parameters RoleAssignmentScheduleRequest, options *RoleAssignmentScheduleRequestsClientValidateOptions) (RoleAssignmentScheduleRequestsClientValidateResponse, error) {
	req, err := client.validateCreateRequest(ctx, scope, roleAssignmentScheduleRequestName, parameters, options)
	if err != nil {
		return RoleAssignmentScheduleRequestsClientValidateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentScheduleRequestsClientValidateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentScheduleRequestsClientValidateResponse{}, runtime.NewResponseError(resp)
	}
	return client.validateHandleResponse(resp)
}

// validateCreateRequest creates the Validate request.
func (client *RoleAssignmentScheduleRequestsClient) validateCreateRequest(ctx context.Context, scope string, roleAssignmentScheduleRequestName string, parameters RoleAssignmentScheduleRequest, options *RoleAssignmentScheduleRequestsClientValidateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/{roleAssignmentScheduleRequestName}/validate"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentScheduleRequestName == "" {
		return nil, errors.New("parameter roleAssignmentScheduleRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentScheduleRequestName}", url.PathEscape(roleAssignmentScheduleRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// validateHandleResponse handles the Validate response.
func (client *RoleAssignmentScheduleRequestsClient) validateHandleResponse(resp *http.Response) (RoleAssignmentScheduleRequestsClientValidateResponse, error) {
	result := RoleAssignmentScheduleRequestsClientValidateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentScheduleRequest); err != nil {
		return RoleAssignmentScheduleRequestsClientValidateResponse{}, err
	}
	return result, nil
}
