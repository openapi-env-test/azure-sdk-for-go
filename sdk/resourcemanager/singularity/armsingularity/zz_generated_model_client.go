//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsingularity

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

// ModelClient contains the methods for the Model group.
// Don't use this type directly, use NewModelClient() instead.
type ModelClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewModelClient creates a new instance of ModelClient with the specified values.
// subscriptionID - The customer subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewModelClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ModelClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ModelClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates a model resource with the specified name, description and properties. If a model with the
// same name exists, then it is updated with the specified description and properties.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// accountName - The name of the Singularity account.
// modelName - The name of the Singularity model.
// body - Singularity model information.
// options - ModelClientBeginCreateOrUpdateOptions contains the optional parameters for the ModelClient.BeginCreateOrUpdate
// method.
func (client *ModelClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, modelName string, body ModelResourceDescription, options *ModelClientBeginCreateOrUpdateOptions) (ModelClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, accountName, modelName, body, options)
	if err != nil {
		return ModelClientCreateOrUpdatePollerResponse{}, err
	}
	result := ModelClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ModelClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return ModelClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ModelClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a model resource with the specified name, description and properties. If a model with the same
// name exists, then it is updated with the specified description and properties.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ModelClient) createOrUpdate(ctx context.Context, resourceGroupName string, accountName string, modelName string, body ModelResourceDescription, options *ModelClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, modelName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ModelClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, modelName string, body ModelResourceDescription, options *ModelClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Singularity/accounts/{accountName}/models/{modelName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", accountName)
	urlPath = strings.ReplaceAll(urlPath, "{modelName}", modelName)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Deletes the model resource identified by the name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// accountName - The name of the Singularity account.
// modelName - The name of the Singularity model.
// options - ModelClientBeginDeleteOptions contains the optional parameters for the ModelClient.BeginDelete method.
func (client *ModelClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, modelName string, options *ModelClientBeginDeleteOptions) (ModelClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, modelName, options)
	if err != nil {
		return ModelClientDeletePollerResponse{}, err
	}
	result := ModelClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ModelClient.Delete", "", resp, client.pl)
	if err != nil {
		return ModelClientDeletePollerResponse{}, err
	}
	result.Poller = &ModelClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the model resource identified by the name.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ModelClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, modelName string, options *ModelClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, modelName, options)
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
func (client *ModelClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, modelName string, options *ModelClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Singularity/accounts/{accountName}/models/{modelName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", accountName)
	urlPath = strings.ReplaceAll(urlPath, "{modelName}", modelName)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the information about the model resource with the given name. The information include the description and other
// properties of the model.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// accountName - The name of the Singularity account.
// modelName - The name of the Singularity model.
// options - ModelClientGetOptions contains the optional parameters for the ModelClient.Get method.
func (client *ModelClient) Get(ctx context.Context, resourceGroupName string, accountName string, modelName string, options *ModelClientGetOptions) (ModelClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, modelName, options)
	if err != nil {
		return ModelClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ModelClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ModelClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ModelClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, modelName string, options *ModelClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Singularity/accounts/{accountName}/models/{modelName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", accountName)
	urlPath = strings.ReplaceAll(urlPath, "{modelName}", modelName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ModelClient) getHandleResponse(resp *http.Response) (ModelClientGetResponse, error) {
	result := ModelClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ModelResourceDescription); err != nil {
		return ModelClientGetResponse{}, err
	}
	return result, nil
}

// ListByAccount - Gets the information about all the model resources in a given account. The information include the description
// and other properties of the model.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// accountName - The name of the Singularity account.
// options - ModelClientListByAccountOptions contains the optional parameters for the ModelClient.ListByAccount method.
func (client *ModelClient) ListByAccount(resourceGroupName string, accountName string, options *ModelClientListByAccountOptions) *ModelClientListByAccountPager {
	return &ModelClientListByAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByAccountCreateRequest(ctx, resourceGroupName, accountName, options)
		},
		advancer: func(ctx context.Context, resp ModelClientListByAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ModelResourceDescriptionList.NextLink)
		},
	}
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *ModelClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *ModelClientListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Singularity/accounts/{accountName}/models"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", accountName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.ContinuationToken != nil {
		unencodedParams = append(unencodedParams, "continuationToken="+*options.ContinuationToken)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *ModelClient) listByAccountHandleResponse(resp *http.Response) (ModelClientListByAccountResponse, error) {
	result := ModelClientListByAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ModelResourceDescriptionList); err != nil {
		return ModelClientListByAccountResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a model resource with the specified name, description and properties. If a model with the same name
// exists, then it is updated with the specified description and properties.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// accountName - The name of the Singularity account.
// modelName - The name of the Singularity model.
// body - Singularity model information.
// options - ModelClientBeginUpdateOptions contains the optional parameters for the ModelClient.BeginUpdate method.
func (client *ModelClient) BeginUpdate(ctx context.Context, resourceGroupName string, accountName string, modelName string, body ModelResourcePatchDescription, options *ModelClientBeginUpdateOptions) (ModelClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, accountName, modelName, body, options)
	if err != nil {
		return ModelClientUpdatePollerResponse{}, err
	}
	result := ModelClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ModelClient.Update", "", resp, client.pl)
	if err != nil {
		return ModelClientUpdatePollerResponse{}, err
	}
	result.Poller = &ModelClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates a model resource with the specified name, description and properties. If a model with the same name exists,
// then it is updated with the specified description and properties.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ModelClient) update(ctx context.Context, resourceGroupName string, accountName string, modelName string, body ModelResourcePatchDescription, options *ModelClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, modelName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ModelClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, modelName string, body ModelResourcePatchDescription, options *ModelClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Singularity/accounts/{accountName}/models/{modelName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", accountName)
	urlPath = strings.ReplaceAll(urlPath, "{modelName}", modelName)
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginValidate - Checks if the model resource properties are valid
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// accountName - The name of the Singularity account.
// modelName - The name of the Singularity model.
// body - Singularity model information.
// options - ModelClientBeginValidateOptions contains the optional parameters for the ModelClient.BeginValidate method.
func (client *ModelClient) BeginValidate(ctx context.Context, resourceGroupName string, accountName string, modelName string, body ModelResourceDescription, options *ModelClientBeginValidateOptions) (ModelClientValidatePollerResponse, error) {
	resp, err := client.validate(ctx, resourceGroupName, accountName, modelName, body, options)
	if err != nil {
		return ModelClientValidatePollerResponse{}, err
	}
	result := ModelClientValidatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ModelClient.Validate", "", resp, client.pl)
	if err != nil {
		return ModelClientValidatePollerResponse{}, err
	}
	result.Poller = &ModelClientValidatePoller{
		pt: pt,
	}
	return result, nil
}

// Validate - Checks if the model resource properties are valid
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ModelClient) validate(ctx context.Context, resourceGroupName string, accountName string, modelName string, body ModelResourceDescription, options *ModelClientBeginValidateOptions) (*http.Response, error) {
	req, err := client.validateCreateRequest(ctx, resourceGroupName, accountName, modelName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// validateCreateRequest creates the Validate request.
func (client *ModelClient) validateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, modelName string, body ModelResourceDescription, options *ModelClientBeginValidateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Singularity/accounts/{accountName}/models/{modelName}/validate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", accountName)
	urlPath = strings.ReplaceAll(urlPath, "{modelName}", modelName)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}
