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

// MonitoringConfigClient is the client for the MonitoringConfig methods of the Databoxedge service.
type MonitoringConfigClient struct {
	BaseClient
}

// NewMonitoringConfigClient creates an instance of the MonitoringConfigClient client.
func NewMonitoringConfigClient(subscriptionID string) MonitoringConfigClient {
	return NewMonitoringConfigClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMonitoringConfigClientWithBaseURI creates an instance of the MonitoringConfigClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewMonitoringConfigClientWithBaseURI(baseURI string, subscriptionID string) MonitoringConfigClient {
	return MonitoringConfigClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate sends the create or update request.
// Parameters:
// deviceName - the device name.
// roleName - the role name.
// monitoringMetricConfiguration - the metric configuration.
// resourceGroupName - the resource group name.
func (client MonitoringConfigClient) CreateOrUpdate(ctx context.Context, deviceName string, roleName string, monitoringMetricConfiguration MonitoringMetricConfiguration, resourceGroupName string) (result MonitoringConfigCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitoringConfigClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: monitoringMetricConfiguration,
			Constraints: []validation.Constraint{{Target: "monitoringMetricConfiguration.MonitoringMetricConfigurationProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "monitoringMetricConfiguration.MonitoringMetricConfigurationProperties.MetricConfigurations", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("databoxedge.MonitoringConfigClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, deviceName, roleName, monitoringMetricConfiguration, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.MonitoringConfigClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.MonitoringConfigClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client MonitoringConfigClient) CreateOrUpdatePreparer(ctx context.Context, deviceName string, roleName string, monitoringMetricConfiguration MonitoringMetricConfiguration, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"roleName":          autorest.Encode("path", roleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/monitoringConfig/default", pathParameters),
		autorest.WithJSON(monitoringMetricConfiguration),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client MonitoringConfigClient) CreateOrUpdateSender(req *http.Request) (future MonitoringConfigCreateOrUpdateFuture, err error) {
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
func (client MonitoringConfigClient) CreateOrUpdateResponder(resp *http.Response) (result MonitoringMetricConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete sends the delete request.
// Parameters:
// deviceName - the device name.
// roleName - the role name.
// resourceGroupName - the resource group name.
func (client MonitoringConfigClient) Delete(ctx context.Context, deviceName string, roleName string, resourceGroupName string) (result MonitoringConfigDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitoringConfigClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, deviceName, roleName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.MonitoringConfigClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.MonitoringConfigClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client MonitoringConfigClient) DeletePreparer(ctx context.Context, deviceName string, roleName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"roleName":          autorest.Encode("path", roleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/monitoringConfig/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client MonitoringConfigClient) DeleteSender(req *http.Request) (future MonitoringConfigDeleteFuture, err error) {
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
func (client MonitoringConfigClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
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
// roleName - the role name.
// resourceGroupName - the resource group name.
func (client MonitoringConfigClient) Get(ctx context.Context, deviceName string, roleName string, resourceGroupName string) (result MonitoringMetricConfiguration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitoringConfigClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, deviceName, roleName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.MonitoringConfigClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databoxedge.MonitoringConfigClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.MonitoringConfigClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client MonitoringConfigClient) GetPreparer(ctx context.Context, deviceName string, roleName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"roleName":          autorest.Encode("path", roleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/monitoringConfig/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MonitoringConfigClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MonitoringConfigClient) GetResponder(resp *http.Response) (result MonitoringMetricConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
// Parameters:
// deviceName - the device name.
// roleName - the role name.
// resourceGroupName - the resource group name.
func (client MonitoringConfigClient) List(ctx context.Context, deviceName string, roleName string, resourceGroupName string) (result MonitoringMetricConfigurationListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitoringConfigClient.List")
		defer func() {
			sc := -1
			if result.mmcl.Response.Response != nil {
				sc = result.mmcl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, deviceName, roleName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.MonitoringConfigClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.mmcl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databoxedge.MonitoringConfigClient", "List", resp, "Failure sending request")
		return
	}

	result.mmcl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.MonitoringConfigClient", "List", resp, "Failure responding to request")
		return
	}
	if result.mmcl.hasNextLink() && result.mmcl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client MonitoringConfigClient) ListPreparer(ctx context.Context, deviceName string, roleName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"roleName":          autorest.Encode("path", roleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/monitoringConfig", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MonitoringConfigClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MonitoringConfigClient) ListResponder(resp *http.Response) (result MonitoringMetricConfigurationList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client MonitoringConfigClient) listNextResults(ctx context.Context, lastResults MonitoringMetricConfigurationList) (result MonitoringMetricConfigurationList, err error) {
	req, err := lastResults.monitoringMetricConfigurationListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "databoxedge.MonitoringConfigClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "databoxedge.MonitoringConfigClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.MonitoringConfigClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client MonitoringConfigClient) ListComplete(ctx context.Context, deviceName string, roleName string, resourceGroupName string) (result MonitoringMetricConfigurationListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitoringConfigClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, deviceName, roleName, resourceGroupName)
	return
}
