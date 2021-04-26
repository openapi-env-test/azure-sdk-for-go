package databoxedge

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

// StorageAccountsClient is the client for the StorageAccounts methods of the Databoxedge service.
type StorageAccountsClient struct {
	BaseClient
}

// NewStorageAccountsClient creates an instance of the StorageAccountsClient client.
func NewStorageAccountsClient(subscriptionID string) StorageAccountsClient {
	return NewStorageAccountsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewStorageAccountsClientWithBaseURI creates an instance of the StorageAccountsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewStorageAccountsClientWithBaseURI(baseURI string, subscriptionID string) StorageAccountsClient {
	return StorageAccountsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate sends the create or update request.
// Parameters:
// deviceName - the device name.
// storageAccountName - the StorageAccount name.
// storageAccount - the StorageAccount properties.
// resourceGroupName - the resource group name.
func (client StorageAccountsClient) CreateOrUpdate(ctx context.Context, deviceName string, storageAccountName string, storageAccount StorageAccount, resourceGroupName string) (result StorageAccountsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: storageAccount,
			Constraints: []validation.Constraint{{Target: "storageAccount.StorageAccountProperties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("databoxedge.StorageAccountsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, deviceName, storageAccountName, storageAccount, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client StorageAccountsClient) CreateOrUpdatePreparer(ctx context.Context, deviceName string, storageAccountName string, storageAccount StorageAccount, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":         autorest.Encode("path", deviceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"storageAccountName": autorest.Encode("path", storageAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}", pathParameters),
		autorest.WithJSON(storageAccount),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountsClient) CreateOrUpdateSender(req *http.Request) (future StorageAccountsCreateOrUpdateFuture, err error) {
	var resp *http.Response
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

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client StorageAccountsClient) CreateOrUpdateResponder(resp *http.Response) (result StorageAccount, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the StorageAccount on the Data Box Edge/Data Box Gateway device.
// Parameters:
// deviceName - the device name.
// storageAccountName - the StorageAccount name.
// resourceGroupName - the resource group name.
func (client StorageAccountsClient) Delete(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string) (result StorageAccountsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, deviceName, storageAccountName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client StorageAccountsClient) DeletePreparer(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":         autorest.Encode("path", deviceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"storageAccountName": autorest.Encode("path", storageAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountsClient) DeleteSender(req *http.Request) (future StorageAccountsDeleteFuture, err error) {
	var resp *http.Response
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

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client StorageAccountsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get sends the get request.
// Parameters:
// deviceName - the device name.
// storageAccountName - the storage account name.
// resourceGroupName - the resource group name.
func (client StorageAccountsClient) Get(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string) (result StorageAccount, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, deviceName, storageAccountName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client StorageAccountsClient) GetPreparer(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":         autorest.Encode("path", deviceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"storageAccountName": autorest.Encode("path", storageAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client StorageAccountsClient) GetResponder(resp *http.Response) (result StorageAccount, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDataBoxEdgeDevice sends the list by data box edge device request.
// Parameters:
// deviceName - the device name.
// resourceGroupName - the resource group name.
func (client StorageAccountsClient) ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result StorageAccountListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountsClient.ListByDataBoxEdgeDevice")
		defer func() {
			sc := -1
			if result.sal.Response.Response != nil {
				sc = result.sal.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByDataBoxEdgeDeviceNextResults
	req, err := client.ListByDataBoxEdgeDevicePreparer(ctx, deviceName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountsClient", "ListByDataBoxEdgeDevice", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDataBoxEdgeDeviceSender(req)
	if err != nil {
		result.sal.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountsClient", "ListByDataBoxEdgeDevice", resp, "Failure sending request")
		return
	}

	result.sal, err = client.ListByDataBoxEdgeDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountsClient", "ListByDataBoxEdgeDevice", resp, "Failure responding to request")
		return
	}
	if result.sal.hasNextLink() && result.sal.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByDataBoxEdgeDevicePreparer prepares the ListByDataBoxEdgeDevice request.
func (client StorageAccountsClient) ListByDataBoxEdgeDevicePreparer(ctx context.Context, deviceName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDataBoxEdgeDeviceSender sends the ListByDataBoxEdgeDevice request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountsClient) ListByDataBoxEdgeDeviceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDataBoxEdgeDeviceResponder handles the response to the ListByDataBoxEdgeDevice request. The method always
// closes the http.Response Body.
func (client StorageAccountsClient) ListByDataBoxEdgeDeviceResponder(resp *http.Response) (result StorageAccountList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDataBoxEdgeDeviceNextResults retrieves the next set of results, if any.
func (client StorageAccountsClient) listByDataBoxEdgeDeviceNextResults(ctx context.Context, lastResults StorageAccountList) (result StorageAccountList, err error) {
	req, err := lastResults.storageAccountListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "databoxedge.StorageAccountsClient", "listByDataBoxEdgeDeviceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDataBoxEdgeDeviceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "databoxedge.StorageAccountsClient", "listByDataBoxEdgeDeviceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDataBoxEdgeDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountsClient", "listByDataBoxEdgeDeviceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByDataBoxEdgeDeviceComplete enumerates all values, automatically crossing page boundaries as required.
func (client StorageAccountsClient) ListByDataBoxEdgeDeviceComplete(ctx context.Context, deviceName string, resourceGroupName string) (result StorageAccountListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountsClient.ListByDataBoxEdgeDevice")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDataBoxEdgeDevice(ctx, deviceName, resourceGroupName)
	return
}
