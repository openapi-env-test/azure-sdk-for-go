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

// RestorableSQLDatabasesClient contains the methods for the RestorableSQLDatabases group.
// Don't use this type directly, use NewRestorableSQLDatabasesClient() instead.
type RestorableSQLDatabasesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRestorableSQLDatabasesClient creates a new instance of RestorableSQLDatabasesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRestorableSQLDatabasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RestorableSQLDatabasesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &RestorableSQLDatabasesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// List - Show the event feed of all mutations done on all the Azure Cosmos DB SQL databases under the restorable account.
// This helps in scenario where database was accidentally deleted to get the deletion
// time. This API requires 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/…/read' permission
// If the operation fails it returns an *azcore.ResponseError type.
// location - Cosmos DB region, with spaces between words and each word capitalized.
// instanceID - The instanceId GUID of a restorable database account.
// options - RestorableSQLDatabasesClientListOptions contains the optional parameters for the RestorableSQLDatabasesClient.List
// method.
func (client *RestorableSQLDatabasesClient) List(ctx context.Context, location string, instanceID string, options *RestorableSQLDatabasesClientListOptions) (RestorableSQLDatabasesClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, instanceID, options)
	if err != nil {
		return RestorableSQLDatabasesClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RestorableSQLDatabasesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RestorableSQLDatabasesClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *RestorableSQLDatabasesClient) listCreateRequest(ctx context.Context, location string, instanceID string, options *RestorableSQLDatabasesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableSqlDatabases"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorableSQLDatabasesClient) listHandleResponse(resp *http.Response) (RestorableSQLDatabasesClientListResponse, error) {
	result := RestorableSQLDatabasesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableSQLDatabasesListResult); err != nil {
		return RestorableSQLDatabasesClientListResponse{}, err
	}
	return result, nil
}
