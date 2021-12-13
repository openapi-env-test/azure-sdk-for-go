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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// VirtualMachineExtensionImagesClient contains the methods for the VirtualMachineExtensionImages group.
// Don't use this type directly, use NewVirtualMachineExtensionImagesClient() instead.
type VirtualMachineExtensionImagesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVirtualMachineExtensionImagesClient creates a new instance of VirtualMachineExtensionImagesClient with the specified values.
func NewVirtualMachineExtensionImagesClient(con *arm.Connection, subscriptionID string) *VirtualMachineExtensionImagesClient {
	return &VirtualMachineExtensionImagesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Gets a virtual machine extension image.
// If the operation fails it returns a generic error.
func (client *VirtualMachineExtensionImagesClient) Get(ctx context.Context, location string, publisherName string, typeParam string, version string, options *VirtualMachineExtensionImagesGetOptions) (VirtualMachineExtensionImagesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, publisherName, typeParam, version, options)
	if err != nil {
		return VirtualMachineExtensionImagesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineExtensionImagesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineExtensionImagesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineExtensionImagesClient) getCreateRequest(ctx context.Context, location string, publisherName string, typeParam string, version string, options *VirtualMachineExtensionImagesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types/{type}/versions/{version}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if typeParam == "" {
		return nil, errors.New("parameter typeParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{type}", url.PathEscape(typeParam))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
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
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineExtensionImagesClient) getHandleResponse(resp *http.Response) (VirtualMachineExtensionImagesGetResponse, error) {
	result := VirtualMachineExtensionImagesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineExtensionImage); err != nil {
		return VirtualMachineExtensionImagesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VirtualMachineExtensionImagesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListTypes - Gets a list of virtual machine extension image types.
// If the operation fails it returns a generic error.
func (client *VirtualMachineExtensionImagesClient) ListTypes(ctx context.Context, location string, publisherName string, options *VirtualMachineExtensionImagesListTypesOptions) (VirtualMachineExtensionImagesListTypesResponse, error) {
	req, err := client.listTypesCreateRequest(ctx, location, publisherName, options)
	if err != nil {
		return VirtualMachineExtensionImagesListTypesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineExtensionImagesListTypesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineExtensionImagesListTypesResponse{}, client.listTypesHandleError(resp)
	}
	return client.listTypesHandleResponse(resp)
}

// listTypesCreateRequest creates the ListTypes request.
func (client *VirtualMachineExtensionImagesClient) listTypesCreateRequest(ctx context.Context, location string, publisherName string, options *VirtualMachineExtensionImagesListTypesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
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
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listTypesHandleResponse handles the ListTypes response.
func (client *VirtualMachineExtensionImagesClient) listTypesHandleResponse(resp *http.Response) (VirtualMachineExtensionImagesListTypesResponse, error) {
	result := VirtualMachineExtensionImagesListTypesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineExtensionImageArray); err != nil {
		return VirtualMachineExtensionImagesListTypesResponse{}, err
	}
	return result, nil
}

// listTypesHandleError handles the ListTypes error response.
func (client *VirtualMachineExtensionImagesClient) listTypesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListVersions - Gets a list of virtual machine extension image versions.
// If the operation fails it returns a generic error.
func (client *VirtualMachineExtensionImagesClient) ListVersions(ctx context.Context, location string, publisherName string, typeParam string, options *VirtualMachineExtensionImagesListVersionsOptions) (VirtualMachineExtensionImagesListVersionsResponse, error) {
	req, err := client.listVersionsCreateRequest(ctx, location, publisherName, typeParam, options)
	if err != nil {
		return VirtualMachineExtensionImagesListVersionsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineExtensionImagesListVersionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineExtensionImagesListVersionsResponse{}, client.listVersionsHandleError(resp)
	}
	return client.listVersionsHandleResponse(resp)
}

// listVersionsCreateRequest creates the ListVersions request.
func (client *VirtualMachineExtensionImagesClient) listVersionsCreateRequest(ctx context.Context, location string, publisherName string, typeParam string, options *VirtualMachineExtensionImagesListVersionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types/{type}/versions"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if typeParam == "" {
		return nil, errors.New("parameter typeParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{type}", url.PathEscape(typeParam))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listVersionsHandleResponse handles the ListVersions response.
func (client *VirtualMachineExtensionImagesClient) listVersionsHandleResponse(resp *http.Response) (VirtualMachineExtensionImagesListVersionsResponse, error) {
	result := VirtualMachineExtensionImagesListVersionsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineExtensionImageArray); err != nil {
		return VirtualMachineExtensionImagesListVersionsResponse{}, err
	}
	return result, nil
}

// listVersionsHandleError handles the ListVersions error response.
func (client *VirtualMachineExtensionImagesClient) listVersionsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
