//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataprotection

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

// OperationStatusBackupVaultContextClient contains the methods for the OperationStatusBackupVaultContext group.
// Don't use this type directly, use NewOperationStatusBackupVaultContextClient() instead.
type OperationStatusBackupVaultContextClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOperationStatusBackupVaultContextClient creates a new instance of OperationStatusBackupVaultContextClient with the specified values.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOperationStatusBackupVaultContextClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *OperationStatusBackupVaultContextClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &OperationStatusBackupVaultContextClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Gets the operation status for an operation over a BackupVault's context.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group where the backup vault is present.
// vaultName - The name of the backup vault.
// options - OperationStatusBackupVaultContextClientGetOptions contains the optional parameters for the OperationStatusBackupVaultContextClient.Get
// method.
func (client *OperationStatusBackupVaultContextClient) Get(ctx context.Context, resourceGroupName string, vaultName string, operationID string, options *OperationStatusBackupVaultContextClientGetOptions) (OperationStatusBackupVaultContextClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, operationID, options)
	if err != nil {
		return OperationStatusBackupVaultContextClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OperationStatusBackupVaultContextClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationStatusBackupVaultContextClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OperationStatusBackupVaultContextClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, operationID string, options *OperationStatusBackupVaultContextClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/operationStatus/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OperationStatusBackupVaultContextClient) getHandleResponse(resp *http.Response) (OperationStatusBackupVaultContextClientGetResponse, error) {
	result := OperationStatusBackupVaultContextClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationResource); err != nil {
		return OperationStatusBackupVaultContextClientGetResponse{}, err
	}
	return result, nil
}
