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

// DataConnectorsCheckRequirementsClient contains the methods for the DataConnectorsCheckRequirements group.
// Don't use this type directly, use NewDataConnectorsCheckRequirementsClient() instead.
type DataConnectorsCheckRequirementsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDataConnectorsCheckRequirementsClient creates a new instance of DataConnectorsCheckRequirementsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDataConnectorsCheckRequirementsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DataConnectorsCheckRequirementsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &DataConnectorsCheckRequirementsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Post - Get requirements state for a data connector type.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// dataConnectorsCheckRequirements - The parameters for requirements check message
// options - DataConnectorsCheckRequirementsClientPostOptions contains the optional parameters for the DataConnectorsCheckRequirementsClient.Post
// method.
func (client *DataConnectorsCheckRequirementsClient) Post(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorsCheckRequirements DataConnectorsCheckRequirementsClassification, options *DataConnectorsCheckRequirementsClientPostOptions) (DataConnectorsCheckRequirementsClientPostResponse, error) {
	req, err := client.postCreateRequest(ctx, resourceGroupName, workspaceName, dataConnectorsCheckRequirements, options)
	if err != nil {
		return DataConnectorsCheckRequirementsClientPostResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataConnectorsCheckRequirementsClientPostResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataConnectorsCheckRequirementsClientPostResponse{}, runtime.NewResponseError(resp)
	}
	return client.postHandleResponse(resp)
}

// postCreateRequest creates the Post request.
func (client *DataConnectorsCheckRequirementsClient) postCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorsCheckRequirements DataConnectorsCheckRequirementsClassification, options *DataConnectorsCheckRequirementsClientPostOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectorsCheckRequirements"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, dataConnectorsCheckRequirements)
}

// postHandleResponse handles the Post response.
func (client *DataConnectorsCheckRequirementsClient) postHandleResponse(resp *http.Response) (DataConnectorsCheckRequirementsClientPostResponse, error) {
	result := DataConnectorsCheckRequirementsClientPostResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataConnectorRequirementsState); err != nil {
		return DataConnectorsCheckRequirementsClientPostResponse{}, err
	}
	return result, nil
}
