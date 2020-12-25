package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ExpressRouteCircuitsClient is the network Client
type ExpressRouteCircuitsClient struct {
	BaseClient
}

// NewExpressRouteCircuitsClient creates an instance of the ExpressRouteCircuitsClient client.
func NewExpressRouteCircuitsClient(subscriptionID string) ExpressRouteCircuitsClient {
	return NewExpressRouteCircuitsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExpressRouteCircuitsClientWithBaseURI creates an instance of the ExpressRouteCircuitsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewExpressRouteCircuitsClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteCircuitsClient {
	return ExpressRouteCircuitsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an express route circuit.
// Parameters:
// resourceGroupName - the name of the resource group.
// circuitName - the name of the circuit.
// parameters - parameters supplied to the create or update express route circuit operation.
func (client ExpressRouteCircuitsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, parameters ExpressRouteCircuit) (result ExpressRouteCircuitsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCircuitsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, circuitName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ExpressRouteCircuitsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, circuitName string, parameters ExpressRouteCircuit) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       autorest.Encode("path", circuitName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) CreateOrUpdateSender(req *http.Request) (future ExpressRouteCircuitsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) CreateOrUpdateResponder(resp *http.Response) (result ExpressRouteCircuit, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified express route circuit.
// Parameters:
// resourceGroupName - the name of the resource group.
// circuitName - the name of the express route circuit.
func (client ExpressRouteCircuitsClient) Delete(ctx context.Context, resourceGroupName string, circuitName string) (result ExpressRouteCircuitsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCircuitsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, circuitName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExpressRouteCircuitsClient) DeletePreparer(ctx context.Context, resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       autorest.Encode("path", circuitName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) DeleteSender(req *http.Request) (future ExpressRouteCircuitsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified express route circuit.
// Parameters:
// resourceGroupName - the name of the resource group.
// circuitName - the name of express route circuit.
func (client ExpressRouteCircuitsClient) Get(ctx context.Context, resourceGroupName string, circuitName string) (result ExpressRouteCircuit, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCircuitsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, circuitName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExpressRouteCircuitsClient) GetPreparer(ctx context.Context, resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       autorest.Encode("path", circuitName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) GetResponder(resp *http.Response) (result ExpressRouteCircuit, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetPeeringStats gets all stats from an express route circuit in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// circuitName - the name of the express route circuit.
// peeringName - the name of the peering.
func (client ExpressRouteCircuitsClient) GetPeeringStats(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (result ExpressRouteCircuitStats, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCircuitsClient.GetPeeringStats")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPeeringStatsPreparer(ctx, resourceGroupName, circuitName, peeringName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "GetPeeringStats", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetPeeringStatsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "GetPeeringStats", resp, "Failure sending request")
		return
	}

	result, err = client.GetPeeringStatsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "GetPeeringStats", resp, "Failure responding to request")
		return
	}

	return
}

// GetPeeringStatsPreparer prepares the GetPeeringStats request.
func (client ExpressRouteCircuitsClient) GetPeeringStatsPreparer(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       autorest.Encode("path", circuitName),
		"peeringName":       autorest.Encode("path", peeringName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/stats", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetPeeringStatsSender sends the GetPeeringStats request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) GetPeeringStatsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetPeeringStatsResponder handles the response to the GetPeeringStats request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) GetPeeringStatsResponder(resp *http.Response) (result ExpressRouteCircuitStats, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetStats gets all the stats from an express route circuit in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// circuitName - the name of the express route circuit.
func (client ExpressRouteCircuitsClient) GetStats(ctx context.Context, resourceGroupName string, circuitName string) (result ExpressRouteCircuitStats, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCircuitsClient.GetStats")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetStatsPreparer(ctx, resourceGroupName, circuitName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "GetStats", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetStatsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "GetStats", resp, "Failure sending request")
		return
	}

	result, err = client.GetStatsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "GetStats", resp, "Failure responding to request")
		return
	}

	return
}

// GetStatsPreparer prepares the GetStats request.
func (client ExpressRouteCircuitsClient) GetStatsPreparer(ctx context.Context, resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       autorest.Encode("path", circuitName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/stats", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetStatsSender sends the GetStats request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) GetStatsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetStatsResponder handles the response to the GetStats request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) GetStatsResponder(resp *http.Response) (result ExpressRouteCircuitStats, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the express route circuits in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client ExpressRouteCircuitsClient) List(ctx context.Context, resourceGroupName string) (result ExpressRouteCircuitListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCircuitsClient.List")
		defer func() {
			sc := -1
			if result.erclr.Response.Response != nil {
				sc = result.erclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.erclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "List", resp, "Failure sending request")
		return
	}

	result.erclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.erclr.hasNextLink() && result.erclr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ExpressRouteCircuitsClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) ListResponder(resp *http.Response) (result ExpressRouteCircuitListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ExpressRouteCircuitsClient) listNextResults(ctx context.Context, lastResults ExpressRouteCircuitListResult) (result ExpressRouteCircuitListResult, err error) {
	req, err := lastResults.expressRouteCircuitListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExpressRouteCircuitsClient) ListComplete(ctx context.Context, resourceGroupName string) (result ExpressRouteCircuitListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCircuitsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName)
	return
}

// ListAll gets all the express route circuits in a subscription.
func (client ExpressRouteCircuitsClient) ListAll(ctx context.Context) (result ExpressRouteCircuitListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCircuitsClient.ListAll")
		defer func() {
			sc := -1
			if result.erclr.Response.Response != nil {
				sc = result.erclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAllNextResults
	req, err := client.ListAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.erclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListAll", resp, "Failure sending request")
		return
	}

	result.erclr, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListAll", resp, "Failure responding to request")
		return
	}
	if result.erclr.hasNextLink() && result.erclr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client ExpressRouteCircuitsClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCircuits", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) ListAllResponder(resp *http.Response) (result ExpressRouteCircuitListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllNextResults retrieves the next set of results, if any.
func (client ExpressRouteCircuitsClient) listAllNextResults(ctx context.Context, lastResults ExpressRouteCircuitListResult) (result ExpressRouteCircuitListResult, err error) {
	req, err := lastResults.expressRouteCircuitListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExpressRouteCircuitsClient) ListAllComplete(ctx context.Context) (result ExpressRouteCircuitListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCircuitsClient.ListAll")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAll(ctx)
	return
}

// ListArpTable gets the currently advertised ARP table associated with the express route circuit in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// circuitName - the name of the express route circuit.
// peeringName - the name of the peering.
// devicePath - the path of the device.
func (client ExpressRouteCircuitsClient) ListArpTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (result ExpressRouteCircuitsListArpTableFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCircuitsClient.ListArpTable")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListArpTablePreparer(ctx, resourceGroupName, circuitName, peeringName, devicePath)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListArpTable", nil, "Failure preparing request")
		return
	}

	result, err = client.ListArpTableSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListArpTable", result.Response(), "Failure sending request")
		return
	}

	return
}

// ListArpTablePreparer prepares the ListArpTable request.
func (client ExpressRouteCircuitsClient) ListArpTablePreparer(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       autorest.Encode("path", circuitName),
		"devicePath":        autorest.Encode("path", devicePath),
		"peeringName":       autorest.Encode("path", peeringName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/arpTables/{devicePath}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListArpTableSender sends the ListArpTable request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) ListArpTableSender(req *http.Request) (future ExpressRouteCircuitsListArpTableFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ListArpTableResponder handles the response to the ListArpTable request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) ListArpTableResponder(resp *http.Response) (result ExpressRouteCircuitsArpTableListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListRoutesTable gets the currently advertised routes table associated with the express route circuit in a resource
// group.
// Parameters:
// resourceGroupName - the name of the resource group.
// circuitName - the name of the express route circuit.
// peeringName - the name of the peering.
// devicePath - the path of the device.
func (client ExpressRouteCircuitsClient) ListRoutesTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (result ExpressRouteCircuitsListRoutesTableFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCircuitsClient.ListRoutesTable")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListRoutesTablePreparer(ctx, resourceGroupName, circuitName, peeringName, devicePath)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListRoutesTable", nil, "Failure preparing request")
		return
	}

	result, err = client.ListRoutesTableSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListRoutesTable", result.Response(), "Failure sending request")
		return
	}

	return
}

// ListRoutesTablePreparer prepares the ListRoutesTable request.
func (client ExpressRouteCircuitsClient) ListRoutesTablePreparer(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       autorest.Encode("path", circuitName),
		"devicePath":        autorest.Encode("path", devicePath),
		"peeringName":       autorest.Encode("path", peeringName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/routeTables/{devicePath}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListRoutesTableSender sends the ListRoutesTable request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) ListRoutesTableSender(req *http.Request) (future ExpressRouteCircuitsListRoutesTableFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ListRoutesTableResponder handles the response to the ListRoutesTable request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) ListRoutesTableResponder(resp *http.Response) (result ExpressRouteCircuitsRoutesTableListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListRoutesTableSummary gets the currently advertised routes table summary associated with the express route circuit
// in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// circuitName - the name of the express route circuit.
// peeringName - the name of the peering.
// devicePath - the path of the device.
func (client ExpressRouteCircuitsClient) ListRoutesTableSummary(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (result ExpressRouteCircuitsListRoutesTableSummaryFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCircuitsClient.ListRoutesTableSummary")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListRoutesTableSummaryPreparer(ctx, resourceGroupName, circuitName, peeringName, devicePath)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListRoutesTableSummary", nil, "Failure preparing request")
		return
	}

	result, err = client.ListRoutesTableSummarySender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListRoutesTableSummary", result.Response(), "Failure sending request")
		return
	}

	return
}

// ListRoutesTableSummaryPreparer prepares the ListRoutesTableSummary request.
func (client ExpressRouteCircuitsClient) ListRoutesTableSummaryPreparer(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       autorest.Encode("path", circuitName),
		"devicePath":        autorest.Encode("path", devicePath),
		"peeringName":       autorest.Encode("path", peeringName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/routeTablesSummary/{devicePath}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListRoutesTableSummarySender sends the ListRoutesTableSummary request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) ListRoutesTableSummarySender(req *http.Request) (future ExpressRouteCircuitsListRoutesTableSummaryFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ListRoutesTableSummaryResponder handles the response to the ListRoutesTableSummary request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) ListRoutesTableSummaryResponder(resp *http.Response) (result ExpressRouteCircuitsRoutesTableSummaryListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
