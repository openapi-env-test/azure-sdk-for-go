//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// CapacityReservationGroupsClient contains the methods for the CapacityReservationGroups group.
// Don't use this type directly, use NewCapacityReservationGroupsClient() instead.
type CapacityReservationGroupsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewCapacityReservationGroupsClient creates a new instance of CapacityReservationGroupsClient with the specified values.
func NewCapacityReservationGroupsClient(con *arm.Connection, subscriptionID string) *CapacityReservationGroupsClient {
	return &CapacityReservationGroupsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - The operation to create or update a capacity reservation group. When updating a capacity reservation group, only tags may be modified.
// Please refer to https://aka.ms/CapacityReservation for more
// details.
// If the operation fails it returns the *CloudError error type.
func (client *CapacityReservationGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, parameters CapacityReservationGroup, options *CapacityReservationGroupsCreateOrUpdateOptions) (CapacityReservationGroupsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, capacityReservationGroupName, parameters, options)
	if err != nil {
		return CapacityReservationGroupsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CapacityReservationGroupsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return CapacityReservationGroupsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CapacityReservationGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, parameters CapacityReservationGroup, options *CapacityReservationGroupsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityReservationGroupName == "" {
		return nil, errors.New("parameter capacityReservationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityReservationGroupName}", url.PathEscape(capacityReservationGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CapacityReservationGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (CapacityReservationGroupsCreateOrUpdateResponse, error) {
	result := CapacityReservationGroupsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityReservationGroup); err != nil {
		return CapacityReservationGroupsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *CapacityReservationGroupsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - The operation to delete a capacity reservation group. This operation is allowed only if all the associated resources are disassociated from
// the reservation group and all capacity reservations under
// the reservation group have also been deleted. Please refer to https://aka.ms/CapacityReservation for more details.
// If the operation fails it returns the *CloudError error type.
func (client *CapacityReservationGroupsClient) Delete(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, options *CapacityReservationGroupsDeleteOptions) (CapacityReservationGroupsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, capacityReservationGroupName, options)
	if err != nil {
		return CapacityReservationGroupsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CapacityReservationGroupsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return CapacityReservationGroupsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return CapacityReservationGroupsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CapacityReservationGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, options *CapacityReservationGroupsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityReservationGroupName == "" {
		return nil, errors.New("parameter capacityReservationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityReservationGroupName}", url.PathEscape(capacityReservationGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *CapacityReservationGroupsClient) deleteHandleError(resp *http.Response) error {
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

// Get - The operation that retrieves information about a capacity reservation group.
// If the operation fails it returns the *CloudError error type.
func (client *CapacityReservationGroupsClient) Get(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, options *CapacityReservationGroupsGetOptions) (CapacityReservationGroupsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, capacityReservationGroupName, options)
	if err != nil {
		return CapacityReservationGroupsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CapacityReservationGroupsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CapacityReservationGroupsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CapacityReservationGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, options *CapacityReservationGroupsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityReservationGroupName == "" {
		return nil, errors.New("parameter capacityReservationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityReservationGroupName}", url.PathEscape(capacityReservationGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CapacityReservationGroupsClient) getHandleResponse(resp *http.Response) (CapacityReservationGroupsGetResponse, error) {
	result := CapacityReservationGroupsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityReservationGroup); err != nil {
		return CapacityReservationGroupsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *CapacityReservationGroupsClient) getHandleError(resp *http.Response) error {
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

// ListByResourceGroup - Lists all of the capacity reservation groups in the specified resource group. Use the nextLink property in the response to get
// the next page of capacity reservation groups.
// If the operation fails it returns the *CloudError error type.
func (client *CapacityReservationGroupsClient) ListByResourceGroup(resourceGroupName string, options *CapacityReservationGroupsListByResourceGroupOptions) *CapacityReservationGroupsListByResourceGroupPager {
	return &CapacityReservationGroupsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp CapacityReservationGroupsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CapacityReservationGroupListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CapacityReservationGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CapacityReservationGroupsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CapacityReservationGroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (CapacityReservationGroupsListByResourceGroupResponse, error) {
	result := CapacityReservationGroupsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityReservationGroupListResult); err != nil {
		return CapacityReservationGroupsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *CapacityReservationGroupsClient) listByResourceGroupHandleError(resp *http.Response) error {
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

// ListBySubscription - Lists all of the capacity reservation groups in the subscription. Use the nextLink property in the response to get the next page
// of capacity reservation groups.
// If the operation fails it returns the *CloudError error type.
func (client *CapacityReservationGroupsClient) ListBySubscription(options *CapacityReservationGroupsListBySubscriptionOptions) *CapacityReservationGroupsListBySubscriptionPager {
	return &CapacityReservationGroupsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp CapacityReservationGroupsListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CapacityReservationGroupListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CapacityReservationGroupsClient) listBySubscriptionCreateRequest(ctx context.Context, options *CapacityReservationGroupsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/capacityReservationGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CapacityReservationGroupsClient) listBySubscriptionHandleResponse(resp *http.Response) (CapacityReservationGroupsListBySubscriptionResponse, error) {
	result := CapacityReservationGroupsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityReservationGroupListResult); err != nil {
		return CapacityReservationGroupsListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *CapacityReservationGroupsClient) listBySubscriptionHandleError(resp *http.Response) error {
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

// Update - The operation to update a capacity reservation group. When updating a capacity reservation group, only tags may be modified.
// If the operation fails it returns the *CloudError error type.
func (client *CapacityReservationGroupsClient) Update(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, parameters CapacityReservationGroupUpdate, options *CapacityReservationGroupsUpdateOptions) (CapacityReservationGroupsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, capacityReservationGroupName, parameters, options)
	if err != nil {
		return CapacityReservationGroupsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CapacityReservationGroupsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CapacityReservationGroupsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *CapacityReservationGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, parameters CapacityReservationGroupUpdate, options *CapacityReservationGroupsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityReservationGroupName == "" {
		return nil, errors.New("parameter capacityReservationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityReservationGroupName}", url.PathEscape(capacityReservationGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *CapacityReservationGroupsClient) updateHandleResponse(resp *http.Response) (CapacityReservationGroupsUpdateResponse, error) {
	result := CapacityReservationGroupsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityReservationGroup); err != nil {
		return CapacityReservationGroupsUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *CapacityReservationGroupsClient) updateHandleError(resp *http.Response) error {
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
