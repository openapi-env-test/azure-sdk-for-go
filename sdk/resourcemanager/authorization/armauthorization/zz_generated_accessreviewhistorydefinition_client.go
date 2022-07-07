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

// AccessReviewHistoryDefinitionClient contains the methods for the AccessReviewHistoryDefinition group.
// Don't use this type directly, use NewAccessReviewHistoryDefinitionClient() instead.
type AccessReviewHistoryDefinitionClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAccessReviewHistoryDefinitionClient creates a new instance of AccessReviewHistoryDefinitionClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAccessReviewHistoryDefinitionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AccessReviewHistoryDefinitionClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &AccessReviewHistoryDefinitionClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Create - Create a scheduled or one-time Access Review History Definition
// If the operation fails it returns an *azcore.ResponseError type.
// historyDefinitionID - The id of the access review history definition.
// properties - Access review history definition properties.
// options - AccessReviewHistoryDefinitionClientCreateOptions contains the optional parameters for the AccessReviewHistoryDefinitionClient.Create
// method.
func (client *AccessReviewHistoryDefinitionClient) Create(ctx context.Context, historyDefinitionID string, properties AccessReviewHistoryDefinitionProperties, options *AccessReviewHistoryDefinitionClientCreateOptions) (AccessReviewHistoryDefinitionClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, historyDefinitionID, properties, options)
	if err != nil {
		return AccessReviewHistoryDefinitionClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessReviewHistoryDefinitionClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessReviewHistoryDefinitionClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *AccessReviewHistoryDefinitionClient) createCreateRequest(ctx context.Context, historyDefinitionID string, properties AccessReviewHistoryDefinitionProperties, options *AccessReviewHistoryDefinitionClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewHistoryDefinitions/{historyDefinitionId}"
	if historyDefinitionID == "" {
		return nil, errors.New("parameter historyDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{historyDefinitionId}", url.PathEscape(historyDefinitionID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, properties)
}

// createHandleResponse handles the Create response.
func (client *AccessReviewHistoryDefinitionClient) createHandleResponse(resp *http.Response) (AccessReviewHistoryDefinitionClientCreateResponse, error) {
	result := AccessReviewHistoryDefinitionClientCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewHistoryDefinition); err != nil {
		return AccessReviewHistoryDefinitionClientCreateResponse{}, err
	}
	return result, nil
}

// DeleteByID - Delete an access review history definition
// If the operation fails it returns an *azcore.ResponseError type.
// historyDefinitionID - The id of the access review history definition.
// options - AccessReviewHistoryDefinitionClientDeleteByIDOptions contains the optional parameters for the AccessReviewHistoryDefinitionClient.DeleteByID
// method.
func (client *AccessReviewHistoryDefinitionClient) DeleteByID(ctx context.Context, historyDefinitionID string, options *AccessReviewHistoryDefinitionClientDeleteByIDOptions) (AccessReviewHistoryDefinitionClientDeleteByIDResponse, error) {
	req, err := client.deleteByIDCreateRequest(ctx, historyDefinitionID, options)
	if err != nil {
		return AccessReviewHistoryDefinitionClientDeleteByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessReviewHistoryDefinitionClientDeleteByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AccessReviewHistoryDefinitionClientDeleteByIDResponse{}, runtime.NewResponseError(resp)
	}
	return AccessReviewHistoryDefinitionClientDeleteByIDResponse{RawResponse: resp}, nil
}

// deleteByIDCreateRequest creates the DeleteByID request.
func (client *AccessReviewHistoryDefinitionClient) deleteByIDCreateRequest(ctx context.Context, historyDefinitionID string, options *AccessReviewHistoryDefinitionClientDeleteByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewHistoryDefinitions/{historyDefinitionId}"
	if historyDefinitionID == "" {
		return nil, errors.New("parameter historyDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{historyDefinitionId}", url.PathEscape(historyDefinitionID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
