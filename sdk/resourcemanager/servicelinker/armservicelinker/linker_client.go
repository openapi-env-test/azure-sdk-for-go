//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armservicelinker

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LinkerClient contains the methods for the Linker group.
// Don't use this type directly, use NewLinkerClient() instead.
type LinkerClient struct {
	host string
	pl   runtime.Pipeline
}

// NewLinkerClient creates a new instance of LinkerClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLinkerClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*LinkerClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &LinkerClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// ListConfigurations - list source configurations for a linker.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
// linkerName - The name Linker resource.
// options - LinkerClientListConfigurationsOptions contains the optional parameters for the LinkerClient.ListConfigurations
// method.
func (client *LinkerClient) ListConfigurations(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientListConfigurationsOptions) (LinkerClientListConfigurationsResponse, error) {
	req, err := client.listConfigurationsCreateRequest(ctx, resourceURI, linkerName, options)
	if err != nil {
		return LinkerClientListConfigurationsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkerClientListConfigurationsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkerClientListConfigurationsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listConfigurationsHandleResponse(resp)
}

// listConfigurationsCreateRequest creates the ListConfigurations request.
func (client *LinkerClient) listConfigurationsCreateRequest(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientListConfigurationsOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ServiceLinker/linkers/{linkerName}/listConfigurations"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if linkerName == "" {
		return nil, errors.New("parameter linkerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkerName}", url.PathEscape(linkerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listConfigurationsHandleResponse handles the ListConfigurations response.
func (client *LinkerClient) listConfigurationsHandleResponse(resp *http.Response) (LinkerClientListConfigurationsResponse, error) {
	result := LinkerClientListConfigurationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SourceConfigurationResult); err != nil {
		return LinkerClientListConfigurationsResponse{}, err
	}
	return result, nil
}

// BeginValidate - Validate a link.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
// linkerName - The name Linker resource.
// options - LinkerClientBeginValidateOptions contains the optional parameters for the LinkerClient.BeginValidate method.
func (client *LinkerClient) BeginValidate(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientBeginValidateOptions) (*runtime.Poller[LinkerClientValidateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.validate(ctx, resourceURI, linkerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[LinkerClientValidateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[LinkerClientValidateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Validate - Validate a link.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
func (client *LinkerClient) validate(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientBeginValidateOptions) (*http.Response, error) {
	req, err := client.validateCreateRequest(ctx, resourceURI, linkerName, options)
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
func (client *LinkerClient) validateCreateRequest(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientBeginValidateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ServiceLinker/linkers/{linkerName}/validateLinker"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if linkerName == "" {
		return nil, errors.New("parameter linkerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkerName}", url.PathEscape(linkerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
