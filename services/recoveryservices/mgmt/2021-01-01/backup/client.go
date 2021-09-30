// Package backup implements the Azure ARM Backup service API version .
//
// Open API 2.0 Specs for Azure RecoveryServices Backup service
package backup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

const (
	// DefaultBaseURI is the default URI used for the service Backup
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Backup.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// BMSPrepareDataMove prepares source vault for Data Move operation
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// parameters - prepare data move request
func (client BaseClient) BMSPrepareDataMove(ctx context.Context, vaultName string, resourceGroupName string, parameters PrepareDataMoveRequest) (result BMSPrepareDataMoveFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.BMSPrepareDataMove")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.TargetResourceID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.TargetRegion", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("backup.BaseClient", "BMSPrepareDataMove", err.Error())
	}

	req, err := client.BMSPrepareDataMovePreparer(ctx, vaultName, resourceGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.BaseClient", "BMSPrepareDataMove", nil, "Failure preparing request")
		return
	}

	result, err = client.BMSPrepareDataMoveSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.BaseClient", "BMSPrepareDataMove", result.Response(), "Failure sending request")
		return
	}

	return
}

// BMSPrepareDataMovePreparer prepares the BMSPrepareDataMove request.
func (client BaseClient) BMSPrepareDataMovePreparer(ctx context.Context, vaultName string, resourceGroupName string, parameters PrepareDataMoveRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupstorageconfig/vaultstorageconfig/prepareDataMove", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// BMSPrepareDataMoveSender sends the BMSPrepareDataMove request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) BMSPrepareDataMoveSender(req *http.Request) (future BMSPrepareDataMoveFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// BMSPrepareDataMoveResponder handles the response to the BMSPrepareDataMove request. The method always
// closes the http.Response Body.
func (client BaseClient) BMSPrepareDataMoveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// BMSTriggerDataMove triggers Data Move Operation on target vault
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// parameters - trigger data move request
func (client BaseClient) BMSTriggerDataMove(ctx context.Context, vaultName string, resourceGroupName string, parameters TriggerDataMoveRequest) (result BMSTriggerDataMoveFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.BMSTriggerDataMove")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.SourceResourceID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.SourceRegion", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.CorrelationID", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("backup.BaseClient", "BMSTriggerDataMove", err.Error())
	}

	req, err := client.BMSTriggerDataMovePreparer(ctx, vaultName, resourceGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.BaseClient", "BMSTriggerDataMove", nil, "Failure preparing request")
		return
	}

	result, err = client.BMSTriggerDataMoveSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.BaseClient", "BMSTriggerDataMove", result.Response(), "Failure sending request")
		return
	}

	return
}

// BMSTriggerDataMovePreparer prepares the BMSTriggerDataMove request.
func (client BaseClient) BMSTriggerDataMovePreparer(ctx context.Context, vaultName string, resourceGroupName string, parameters TriggerDataMoveRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupstorageconfig/vaultstorageconfig/triggerDataMove", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// BMSTriggerDataMoveSender sends the BMSTriggerDataMove request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) BMSTriggerDataMoveSender(req *http.Request) (future BMSTriggerDataMoveFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// BMSTriggerDataMoveResponder handles the response to the BMSTriggerDataMove request. The method always
// closes the http.Response Body.
func (client BaseClient) BMSTriggerDataMoveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetOperationStatus fetches operation status for data move operation on vault
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
func (client BaseClient) GetOperationStatus(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result OperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetOperationStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetOperationStatusPreparer(ctx, vaultName, resourceGroupName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.BaseClient", "GetOperationStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOperationStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.BaseClient", "GetOperationStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetOperationStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.BaseClient", "GetOperationStatus", resp, "Failure responding to request")
		return
	}

	return
}

// GetOperationStatusPreparer prepares the GetOperationStatus request.
func (client BaseClient) GetOperationStatusPreparer(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupstorageconfig/vaultstorageconfig/operationStatus/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetOperationStatusSender sends the GetOperationStatus request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetOperationStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetOperationStatusResponder handles the response to the GetOperationStatus request. The method always
// closes the http.Response Body.
func (client BaseClient) GetOperationStatusResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// MoveRecoveryPoint sends the move recovery point request.
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// parameters - move Resource Across Tiers Request
func (client BaseClient) MoveRecoveryPoint(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters MoveRPAcrossTiersRequest) (result MoveRecoveryPointFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.MoveRecoveryPoint")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.MoveRecoveryPointPreparer(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, recoveryPointID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.BaseClient", "MoveRecoveryPoint", nil, "Failure preparing request")
		return
	}

	result, err = client.MoveRecoveryPointSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.BaseClient", "MoveRecoveryPoint", result.Response(), "Failure sending request")
		return
	}

	return
}

// MoveRecoveryPointPreparer prepares the MoveRecoveryPoint request.
func (client BaseClient) MoveRecoveryPointPreparer(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters MoveRPAcrossTiersRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":     autorest.Encode("path", containerName),
		"fabricName":        autorest.Encode("path", fabricName),
		"protectedItemName": autorest.Encode("path", protectedItemName),
		"recoveryPointId":   autorest.Encode("path", recoveryPointID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints/{recoveryPointId}/move", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// MoveRecoveryPointSender sends the MoveRecoveryPoint request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) MoveRecoveryPointSender(req *http.Request) (future MoveRecoveryPointFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// MoveRecoveryPointResponder handles the response to the MoveRecoveryPoint request. The method always
// closes the http.Response Body.
func (client BaseClient) MoveRecoveryPointResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
