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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VirtualMachineScaleSetVMRunCommandsClient contains the methods for the VirtualMachineScaleSetVMRunCommands group.
// Don't use this type directly, use NewVirtualMachineScaleSetVMRunCommandsClient() instead.
type VirtualMachineScaleSetVMRunCommandsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualMachineScaleSetVMRunCommandsClient creates a new instance of VirtualMachineScaleSetVMRunCommandsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualMachineScaleSetVMRunCommandsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualMachineScaleSetVMRunCommandsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &VirtualMachineScaleSetVMRunCommandsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - The operation to create or update the VMSS VM run command.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set.
// instanceID - The instance ID of the virtual machine.
// runCommandName - The name of the virtual machine run command.
// runCommand - Parameters supplied to the Create Virtual Machine RunCommand operation.
// options - VirtualMachineScaleSetVMRunCommandsClientBeginCreateOrUpdateOptions contains the optional parameters for the
// VirtualMachineScaleSetVMRunCommandsClient.BeginCreateOrUpdate method.
func (client *VirtualMachineScaleSetVMRunCommandsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineScaleSetVMRunCommandsClientBeginCreateOrUpdateOptions) (VirtualMachineScaleSetVMRunCommandsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, runCommand, options)
	if err != nil {
		return VirtualMachineScaleSetVMRunCommandsClientCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualMachineScaleSetVMRunCommandsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineScaleSetVMRunCommandsClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return VirtualMachineScaleSetVMRunCommandsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualMachineScaleSetVMRunCommandsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - The operation to create or update the VMSS VM run command.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachineScaleSetVMRunCommandsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineScaleSetVMRunCommandsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, runCommand, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualMachineScaleSetVMRunCommandsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineScaleSetVMRunCommandsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, runtime.MarshalAsJSON(req, runCommand)
}

// BeginDelete - The operation to delete the VMSS VM run command.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set.
// instanceID - The instance ID of the virtual machine.
// runCommandName - The name of the virtual machine run command.
// options - VirtualMachineScaleSetVMRunCommandsClientBeginDeleteOptions contains the optional parameters for the VirtualMachineScaleSetVMRunCommandsClient.BeginDelete
// method.
func (client *VirtualMachineScaleSetVMRunCommandsClient) BeginDelete(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsClientBeginDeleteOptions) (VirtualMachineScaleSetVMRunCommandsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, options)
	if err != nil {
		return VirtualMachineScaleSetVMRunCommandsClientDeletePollerResponse{}, err
	}
	result := VirtualMachineScaleSetVMRunCommandsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineScaleSetVMRunCommandsClient.Delete", "", resp, client.pl)
	if err != nil {
		return VirtualMachineScaleSetVMRunCommandsClientDeletePollerResponse{}, err
	}
	result.Poller = &VirtualMachineScaleSetVMRunCommandsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - The operation to delete the VMSS VM run command.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachineScaleSetVMRunCommandsClient) deleteOperation(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, options)
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
func (client *VirtualMachineScaleSetVMRunCommandsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// Get - The operation to get the VMSS VM run command.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set.
// instanceID - The instance ID of the virtual machine.
// runCommandName - The name of the virtual machine run command.
// options - VirtualMachineScaleSetVMRunCommandsClientGetOptions contains the optional parameters for the VirtualMachineScaleSetVMRunCommandsClient.Get
// method.
func (client *VirtualMachineScaleSetVMRunCommandsClient) Get(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsClientGetOptions) (VirtualMachineScaleSetVMRunCommandsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, options)
	if err != nil {
		return VirtualMachineScaleSetVMRunCommandsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineScaleSetVMRunCommandsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineScaleSetVMRunCommandsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineScaleSetVMRunCommandsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) getHandleResponse(resp *http.Response) (VirtualMachineScaleSetVMRunCommandsClientGetResponse, error) {
	result := VirtualMachineScaleSetVMRunCommandsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineRunCommand); err != nil {
		return VirtualMachineScaleSetVMRunCommandsClientGetResponse{}, err
	}
	return result, nil
}

// List - The operation to get all run commands of an instance in Virtual Machine Scaleset.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set.
// instanceID - The instance ID of the virtual machine.
// options - VirtualMachineScaleSetVMRunCommandsClientListOptions contains the optional parameters for the VirtualMachineScaleSetVMRunCommandsClient.List
// method.
func (client *VirtualMachineScaleSetVMRunCommandsClient) List(resourceGroupName string, vmScaleSetName string, instanceID string, options *VirtualMachineScaleSetVMRunCommandsClientListOptions) *VirtualMachineScaleSetVMRunCommandsClientListPager {
	return &VirtualMachineScaleSetVMRunCommandsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, options)
		},
		advancer: func(ctx context.Context, resp VirtualMachineScaleSetVMRunCommandsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualMachineRunCommandsListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualMachineScaleSetVMRunCommandsClient) listCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, options *VirtualMachineScaleSetVMRunCommandsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) listHandleResponse(resp *http.Response) (VirtualMachineScaleSetVMRunCommandsClientListResponse, error) {
	result := VirtualMachineScaleSetVMRunCommandsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineRunCommandsListResult); err != nil {
		return VirtualMachineScaleSetVMRunCommandsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The operation to update the VMSS VM run command.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set.
// instanceID - The instance ID of the virtual machine.
// runCommandName - The name of the virtual machine run command.
// runCommand - Parameters supplied to the Update Virtual Machine RunCommand operation.
// options - VirtualMachineScaleSetVMRunCommandsClientBeginUpdateOptions contains the optional parameters for the VirtualMachineScaleSetVMRunCommandsClient.BeginUpdate
// method.
func (client *VirtualMachineScaleSetVMRunCommandsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineScaleSetVMRunCommandsClientBeginUpdateOptions) (VirtualMachineScaleSetVMRunCommandsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, runCommand, options)
	if err != nil {
		return VirtualMachineScaleSetVMRunCommandsClientUpdatePollerResponse{}, err
	}
	result := VirtualMachineScaleSetVMRunCommandsClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineScaleSetVMRunCommandsClient.Update", "", resp, client.pl)
	if err != nil {
		return VirtualMachineScaleSetVMRunCommandsClientUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualMachineScaleSetVMRunCommandsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - The operation to update the VMSS VM run command.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachineScaleSetVMRunCommandsClient) update(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineScaleSetVMRunCommandsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, runCommand, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualMachineScaleSetVMRunCommandsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineScaleSetVMRunCommandsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, runtime.MarshalAsJSON(req, runCommand)
}
