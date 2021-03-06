// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// RoleEligibilityScheduleInstancesClient contains the methods for the RoleEligibilityScheduleInstances group.
// Don't use this type directly, use NewRoleEligibilityScheduleInstancesClient() instead.
type RoleEligibilityScheduleInstancesClient struct {
	con *armcore.Connection
}

// NewRoleEligibilityScheduleInstancesClient creates a new instance of RoleEligibilityScheduleInstancesClient with the specified values.
func NewRoleEligibilityScheduleInstancesClient(con *armcore.Connection) *RoleEligibilityScheduleInstancesClient {
	return &RoleEligibilityScheduleInstancesClient{con: con}
}

// Get - Gets the specified role eligibility schedule instance.
// If the operation fails it returns the *CloudError error type.
func (client *RoleEligibilityScheduleInstancesClient) Get(ctx context.Context, scope string, roleEligibilityScheduleInstanceName string, options *RoleEligibilityScheduleInstancesGetOptions) (RoleEligibilityScheduleInstanceResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, roleEligibilityScheduleInstanceName, options)
	if err != nil {
		return RoleEligibilityScheduleInstanceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RoleEligibilityScheduleInstanceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RoleEligibilityScheduleInstanceResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleEligibilityScheduleInstancesClient) getCreateRequest(ctx context.Context, scope string, roleEligibilityScheduleInstanceName string, options *RoleEligibilityScheduleInstancesGetOptions) (*azcore.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleInstances/{roleEligibilityScheduleInstanceName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleEligibilityScheduleInstanceName == "" {
		return nil, errors.New("parameter roleEligibilityScheduleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleEligibilityScheduleInstanceName}", url.PathEscape(roleEligibilityScheduleInstanceName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleEligibilityScheduleInstancesClient) getHandleResponse(resp *azcore.Response) (RoleEligibilityScheduleInstanceResponse, error) {
	var val *RoleEligibilityScheduleInstance
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RoleEligibilityScheduleInstanceResponse{}, err
	}
	return RoleEligibilityScheduleInstanceResponse{RawResponse: resp.Response, RoleEligibilityScheduleInstance: val}, nil
}

// getHandleError handles the Get error response.
func (client *RoleEligibilityScheduleInstancesClient) getHandleError(resp *azcore.Response) error {
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

// ListForScope - Gets role eligibility schedule instances of a role eligibility schedule.
// If the operation fails it returns the *CloudError error type.
func (client *RoleEligibilityScheduleInstancesClient) ListForScope(scope string, options *RoleEligibilityScheduleInstancesListForScopeOptions) RoleEligibilityScheduleInstanceListResultPager {
	return &roleEligibilityScheduleInstanceListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listForScopeCreateRequest(ctx, scope, options)
		},
		responder: client.listForScopeHandleResponse,
		errorer:   client.listForScopeHandleError,
		advancer: func(ctx context.Context, resp RoleEligibilityScheduleInstanceListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RoleEligibilityScheduleInstanceListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleEligibilityScheduleInstancesClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleEligibilityScheduleInstancesListForScopeOptions) (*azcore.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleInstances"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-10-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleEligibilityScheduleInstancesClient) listForScopeHandleResponse(resp *azcore.Response) (RoleEligibilityScheduleInstanceListResultResponse, error) {
	var val *RoleEligibilityScheduleInstanceListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RoleEligibilityScheduleInstanceListResultResponse{}, err
	}
	return RoleEligibilityScheduleInstanceListResultResponse{RawResponse: resp.Response, RoleEligibilityScheduleInstanceListResult: val}, nil
}

// listForScopeHandleError handles the ListForScope error response.
func (client *RoleEligibilityScheduleInstancesClient) listForScopeHandleError(resp *azcore.Response) error {
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
