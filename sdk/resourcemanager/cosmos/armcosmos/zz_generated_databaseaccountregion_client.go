//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// DatabaseAccountRegionClient contains the methods for the DatabaseAccountRegion group.
// Don't use this type directly, use NewDatabaseAccountRegionClient() instead.
type DatabaseAccountRegionClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDatabaseAccountRegionClient creates a new instance of DatabaseAccountRegionClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDatabaseAccountRegionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DatabaseAccountRegionClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &DatabaseAccountRegionClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// ListMetrics - Retrieves the metrics determined by the given filter for the given database account and region.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - Cosmos DB database account name.
// region - Cosmos DB region, with spaces between words and each word capitalized.
// filter - An OData filter expression that describes a subset of metrics to return. The parameters that can be filtered are
// name.value (name of the metric, can have an or of multiple names), startTime, endTime,
// and timeGrain. The supported operator is eq.
// options - DatabaseAccountRegionClientListMetricsOptions contains the optional parameters for the DatabaseAccountRegionClient.ListMetrics
// method.
func (client *DatabaseAccountRegionClient) ListMetrics(ctx context.Context, resourceGroupName string, accountName string, region string, filter string, options *DatabaseAccountRegionClientListMetricsOptions) (DatabaseAccountRegionClientListMetricsResponse, error) {
	req, err := client.listMetricsCreateRequest(ctx, resourceGroupName, accountName, region, filter, options)
	if err != nil {
		return DatabaseAccountRegionClientListMetricsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabaseAccountRegionClientListMetricsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabaseAccountRegionClientListMetricsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listMetricsHandleResponse(resp)
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *DatabaseAccountRegionClient) listMetricsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, region string, filter string, options *DatabaseAccountRegionClientListMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/metrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if region == "" {
		return nil, errors.New("parameter region cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{region}", url.PathEscape(region))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-15-preview")
	reqQP.Set("$filter", filter)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listMetricsHandleResponse handles the ListMetrics response.
func (client *DatabaseAccountRegionClient) listMetricsHandleResponse(resp *http.Response) (DatabaseAccountRegionClientListMetricsResponse, error) {
	result := DatabaseAccountRegionClientListMetricsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricListResult); err != nil {
		return DatabaseAccountRegionClientListMetricsResponse{}, err
	}
	return result, nil
}
