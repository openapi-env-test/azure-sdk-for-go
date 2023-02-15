//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

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

// VPNSitesConfigurationClient contains the methods for the VPNSitesConfiguration group.
// Don't use this type directly, use NewVPNSitesConfigurationClient() instead.
type VPNSitesConfigurationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVPNSitesConfigurationClient creates a new instance of VPNSitesConfigurationClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVPNSitesConfigurationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VPNSitesConfigurationClient, error) {
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
	client := &VPNSitesConfigurationClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginDownload - Gives the sas-url to download the configurations for vpn-sites in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
// resourceGroupName - The resource group name.
// virtualWANName - The name of the VirtualWAN for which configuration of all vpn-sites is needed.
// request - Parameters supplied to download vpn-sites configuration.
// options - VPNSitesConfigurationClientBeginDownloadOptions contains the optional parameters for the VPNSitesConfigurationClient.BeginDownload
// method.
func (client *VPNSitesConfigurationClient) BeginDownload(ctx context.Context, resourceGroupName string, virtualWANName string, request GetVPNSitesConfigurationRequest, options *VPNSitesConfigurationClientBeginDownloadOptions) (*runtime.Poller[VPNSitesConfigurationClientDownloadResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.download(ctx, resourceGroupName, virtualWANName, request, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VPNSitesConfigurationClientDownloadResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VPNSitesConfigurationClientDownloadResponse](options.ResumeToken, client.pl, nil)
	}
}

// Download - Gives the sas-url to download the configurations for vpn-sites in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
func (client *VPNSitesConfigurationClient) download(ctx context.Context, resourceGroupName string, virtualWANName string, request GetVPNSitesConfigurationRequest, options *VPNSitesConfigurationClientBeginDownloadOptions) (*http.Response, error) {
	req, err := client.downloadCreateRequest(ctx, resourceGroupName, virtualWANName, request, options)
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

// downloadCreateRequest creates the Download request.
func (client *VPNSitesConfigurationClient) downloadCreateRequest(ctx context.Context, resourceGroupName string, virtualWANName string, request GetVPNSitesConfigurationRequest, options *VPNSitesConfigurationClientBeginDownloadOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/vpnConfiguration"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualWANName == "" {
		return nil, errors.New("parameter virtualWANName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualWANName}", url.PathEscape(virtualWANName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, request)
}
