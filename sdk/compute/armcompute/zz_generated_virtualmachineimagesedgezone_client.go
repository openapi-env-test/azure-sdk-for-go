// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// VirtualMachineImagesEdgeZoneClient contains the methods for the VirtualMachineImagesEdgeZone group.
// Don't use this type directly, use NewVirtualMachineImagesEdgeZoneClient() instead.
type VirtualMachineImagesEdgeZoneClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualMachineImagesEdgeZoneClient creates a new instance of VirtualMachineImagesEdgeZoneClient with the specified values.
func NewVirtualMachineImagesEdgeZoneClient(con *armcore.Connection, subscriptionID string) *VirtualMachineImagesEdgeZoneClient {
	return &VirtualMachineImagesEdgeZoneClient{con: con, subscriptionID: subscriptionID}
}

// Get - Gets a virtual machine image in an edge zone.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineImagesEdgeZoneClient) Get(ctx context.Context, location string, edgeZone string, publisherName string, offer string, skus string, version string, options *VirtualMachineImagesEdgeZoneGetOptions) (VirtualMachineImageResponse, error) {
	req, err := client.getCreateRequest(ctx, location, edgeZone, publisherName, offer, skus, version, options)
	if err != nil {
		return VirtualMachineImageResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImageResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineImageResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineImagesEdgeZoneClient) getCreateRequest(ctx context.Context, location string, edgeZone string, publisherName string, offer string, skus string, version string, options *VirtualMachineImagesEdgeZoneGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/edgeZones/{edgeZone}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions/{version}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if edgeZone == "" {
		return nil, errors.New("parameter edgeZone cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{edgeZone}", url.PathEscape(edgeZone))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if offer == "" {
		return nil, errors.New("parameter offer cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	if skus == "" {
		return nil, errors.New("parameter skus cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skus}", url.PathEscape(skus))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
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
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineImagesEdgeZoneClient) getHandleResponse(resp *azcore.Response) (VirtualMachineImageResponse, error) {
	var val *VirtualMachineImage
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualMachineImageResponse{}, err
	}
	return VirtualMachineImageResponse{RawResponse: resp.Response, VirtualMachineImage: val}, nil
}

// getHandleError handles the Get error response.
func (client *VirtualMachineImagesEdgeZoneClient) getHandleError(resp *azcore.Response) error {
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

// List - Gets a list of all virtual machine image versions for the specified location, edge zone, publisher, offer, and SKU.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineImagesEdgeZoneClient) List(ctx context.Context, location string, edgeZone string, publisherName string, offer string, skus string, options *VirtualMachineImagesEdgeZoneListOptions) (VirtualMachineImageResourceArrayResponse, error) {
	req, err := client.listCreateRequest(ctx, location, edgeZone, publisherName, offer, skus, options)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineImageResourceArrayResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *VirtualMachineImagesEdgeZoneClient) listCreateRequest(ctx context.Context, location string, edgeZone string, publisherName string, offer string, skus string, options *VirtualMachineImagesEdgeZoneListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/edgeZones/{edgeZone}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if edgeZone == "" {
		return nil, errors.New("parameter edgeZone cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{edgeZone}", url.PathEscape(edgeZone))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if offer == "" {
		return nil, errors.New("parameter offer cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	if skus == "" {
		return nil, errors.New("parameter skus cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skus}", url.PathEscape(skus))
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineImagesEdgeZoneClient) listHandleResponse(resp *azcore.Response) (VirtualMachineImageResourceArrayResponse, error) {
	var val []*VirtualMachineImageResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	return VirtualMachineImageResourceArrayResponse{RawResponse: resp.Response, VirtualMachineImageResourceArray: val}, nil
}

// listHandleError handles the List error response.
func (client *VirtualMachineImagesEdgeZoneClient) listHandleError(resp *azcore.Response) error {
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

// ListOffers - Gets a list of virtual machine image offers for the specified location, edge zone and publisher.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineImagesEdgeZoneClient) ListOffers(ctx context.Context, location string, edgeZone string, publisherName string, options *VirtualMachineImagesEdgeZoneListOffersOptions) (VirtualMachineImageResourceArrayResponse, error) {
	req, err := client.listOffersCreateRequest(ctx, location, edgeZone, publisherName, options)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineImageResourceArrayResponse{}, client.listOffersHandleError(resp)
	}
	return client.listOffersHandleResponse(resp)
}

// listOffersCreateRequest creates the ListOffers request.
func (client *VirtualMachineImagesEdgeZoneClient) listOffersCreateRequest(ctx context.Context, location string, edgeZone string, publisherName string, options *VirtualMachineImagesEdgeZoneListOffersOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/edgeZones/{edgeZone}/publishers/{publisherName}/artifacttypes/vmimage/offers"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if edgeZone == "" {
		return nil, errors.New("parameter edgeZone cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{edgeZone}", url.PathEscape(edgeZone))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
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
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listOffersHandleResponse handles the ListOffers response.
func (client *VirtualMachineImagesEdgeZoneClient) listOffersHandleResponse(resp *azcore.Response) (VirtualMachineImageResourceArrayResponse, error) {
	var val []*VirtualMachineImageResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	return VirtualMachineImageResourceArrayResponse{RawResponse: resp.Response, VirtualMachineImageResourceArray: val}, nil
}

// listOffersHandleError handles the ListOffers error response.
func (client *VirtualMachineImagesEdgeZoneClient) listOffersHandleError(resp *azcore.Response) error {
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

// ListPublishers - Gets a list of virtual machine image publishers for the specified Azure location and edge zone.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineImagesEdgeZoneClient) ListPublishers(ctx context.Context, location string, edgeZone string, options *VirtualMachineImagesEdgeZoneListPublishersOptions) (VirtualMachineImageResourceArrayResponse, error) {
	req, err := client.listPublishersCreateRequest(ctx, location, edgeZone, options)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineImageResourceArrayResponse{}, client.listPublishersHandleError(resp)
	}
	return client.listPublishersHandleResponse(resp)
}

// listPublishersCreateRequest creates the ListPublishers request.
func (client *VirtualMachineImagesEdgeZoneClient) listPublishersCreateRequest(ctx context.Context, location string, edgeZone string, options *VirtualMachineImagesEdgeZoneListPublishersOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/edgeZones/{edgeZone}/publishers"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if edgeZone == "" {
		return nil, errors.New("parameter edgeZone cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{edgeZone}", url.PathEscape(edgeZone))
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
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listPublishersHandleResponse handles the ListPublishers response.
func (client *VirtualMachineImagesEdgeZoneClient) listPublishersHandleResponse(resp *azcore.Response) (VirtualMachineImageResourceArrayResponse, error) {
	var val []*VirtualMachineImageResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	return VirtualMachineImageResourceArrayResponse{RawResponse: resp.Response, VirtualMachineImageResourceArray: val}, nil
}

// listPublishersHandleError handles the ListPublishers error response.
func (client *VirtualMachineImagesEdgeZoneClient) listPublishersHandleError(resp *azcore.Response) error {
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

// ListSKUs - Gets a list of virtual machine image SKUs for the specified location, edge zone, publisher, and offer.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineImagesEdgeZoneClient) ListSKUs(ctx context.Context, location string, edgeZone string, publisherName string, offer string, options *VirtualMachineImagesEdgeZoneListSKUsOptions) (VirtualMachineImageResourceArrayResponse, error) {
	req, err := client.listSKUsCreateRequest(ctx, location, edgeZone, publisherName, offer, options)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineImageResourceArrayResponse{}, client.listSKUsHandleError(resp)
	}
	return client.listSKUsHandleResponse(resp)
}

// listSKUsCreateRequest creates the ListSKUs request.
func (client *VirtualMachineImagesEdgeZoneClient) listSKUsCreateRequest(ctx context.Context, location string, edgeZone string, publisherName string, offer string, options *VirtualMachineImagesEdgeZoneListSKUsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/edgeZones/{edgeZone}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if edgeZone == "" {
		return nil, errors.New("parameter edgeZone cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{edgeZone}", url.PathEscape(edgeZone))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if offer == "" {
		return nil, errors.New("parameter offer cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
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
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listSKUsHandleResponse handles the ListSKUs response.
func (client *VirtualMachineImagesEdgeZoneClient) listSKUsHandleResponse(resp *azcore.Response) (VirtualMachineImageResourceArrayResponse, error) {
	var val []*VirtualMachineImageResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	return VirtualMachineImageResourceArrayResponse{RawResponse: resp.Response, VirtualMachineImageResourceArray: val}, nil
}

// listSKUsHandleError handles the ListSKUs error response.
func (client *VirtualMachineImagesEdgeZoneClient) listSKUsHandleError(resp *azcore.Response) error {
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
