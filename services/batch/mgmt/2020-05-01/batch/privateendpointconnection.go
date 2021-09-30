package batch

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

// PrivateEndpointConnectionClient is the client for the PrivateEndpointConnection methods of the Batch service.
type PrivateEndpointConnectionClient struct {
	BaseClient
}

// NewPrivateEndpointConnectionClient creates an instance of the PrivateEndpointConnectionClient client.
func NewPrivateEndpointConnectionClient(subscriptionID string) PrivateEndpointConnectionClient {
	return NewPrivateEndpointConnectionClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateEndpointConnectionClientWithBaseURI creates an instance of the PrivateEndpointConnectionClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewPrivateEndpointConnectionClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionClient {
	return PrivateEndpointConnectionClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets information about the specified private endpoint connection.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// privateEndpointConnectionName - the private endpoint connection name. This must be unique within the
// account.
func (client PrivateEndpointConnectionClient) Get(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string) (result PrivateEndpointConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+$`, Chain: nil}}},
		{TargetValue: privateEndpointConnectionName,
			Constraints: []validation.Constraint{{Target: "privateEndpointConnectionName", Name: validation.MaxLength, Rule: 101, Chain: nil},
				{Target: "privateEndpointConnectionName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "privateEndpointConnectionName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+\.?[a-fA-F0-9-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.PrivateEndpointConnectionClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, privateEndpointConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PrivateEndpointConnectionClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.PrivateEndpointConnectionClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PrivateEndpointConnectionClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrivateEndpointConnectionClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":                   autorest.Encode("path", accountName),
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionClient) GetResponder(resp *http.Response) (result PrivateEndpointConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBatchAccount lists all of the private endpoint connections in the specified account.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// maxresults - the maximum number of items to return in the response.
func (client PrivateEndpointConnectionClient) ListByBatchAccount(ctx context.Context, resourceGroupName string, accountName string, maxresults *int32) (result ListPrivateEndpointConnectionsResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionClient.ListByBatchAccount")
		defer func() {
			sc := -1
			if result.lpecr.Response.Response != nil {
				sc = result.lpecr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.PrivateEndpointConnectionClient", "ListByBatchAccount", err.Error())
	}

	result.fn = client.listByBatchAccountNextResults
	req, err := client.ListByBatchAccountPreparer(ctx, resourceGroupName, accountName, maxresults)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PrivateEndpointConnectionClient", "ListByBatchAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBatchAccountSender(req)
	if err != nil {
		result.lpecr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.PrivateEndpointConnectionClient", "ListByBatchAccount", resp, "Failure sending request")
		return
	}

	result.lpecr, err = client.ListByBatchAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PrivateEndpointConnectionClient", "ListByBatchAccount", resp, "Failure responding to request")
		return
	}
	if result.lpecr.hasNextLink() && result.lpecr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByBatchAccountPreparer prepares the ListByBatchAccount request.
func (client PrivateEndpointConnectionClient) ListByBatchAccountPreparer(ctx context.Context, resourceGroupName string, accountName string, maxresults *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if maxresults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxresults)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/privateEndpointConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBatchAccountSender sends the ListByBatchAccount request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionClient) ListByBatchAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByBatchAccountResponder handles the response to the ListByBatchAccount request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionClient) ListByBatchAccountResponder(resp *http.Response) (result ListPrivateEndpointConnectionsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBatchAccountNextResults retrieves the next set of results, if any.
func (client PrivateEndpointConnectionClient) listByBatchAccountNextResults(ctx context.Context, lastResults ListPrivateEndpointConnectionsResult) (result ListPrivateEndpointConnectionsResult, err error) {
	req, err := lastResults.listPrivateEndpointConnectionsResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.PrivateEndpointConnectionClient", "listByBatchAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBatchAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.PrivateEndpointConnectionClient", "listByBatchAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBatchAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PrivateEndpointConnectionClient", "listByBatchAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBatchAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateEndpointConnectionClient) ListByBatchAccountComplete(ctx context.Context, resourceGroupName string, accountName string, maxresults *int32) (result ListPrivateEndpointConnectionsResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionClient.ListByBatchAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBatchAccount(ctx, resourceGroupName, accountName, maxresults)
	return
}

// Update updates the properties of an existing private endpoint connection.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// privateEndpointConnectionName - the private endpoint connection name. This must be unique within the
// account.
// parameters - privateEndpointConnection properties that should be updated. Properties that are supplied will
// be updated, any property not supplied will be unchanged.
// ifMatch - the state (ETag) version of the private endpoint connection to update. This value can be omitted
// or set to "*" to apply the operation unconditionally.
func (client PrivateEndpointConnectionClient) Update(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string, parameters PrivateEndpointConnection, ifMatch string) (result PrivateEndpointConnectionUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+$`, Chain: nil}}},
		{TargetValue: privateEndpointConnectionName,
			Constraints: []validation.Constraint{{Target: "privateEndpointConnectionName", Name: validation.MaxLength, Rule: 101, Chain: nil},
				{Target: "privateEndpointConnectionName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "privateEndpointConnectionName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+\.?[a-fA-F0-9-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.PrivateEndpointConnectionClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, privateEndpointConnectionName, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PrivateEndpointConnectionClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PrivateEndpointConnectionClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client PrivateEndpointConnectionClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string, parameters PrivateEndpointConnection, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":                   autorest.Encode("path", accountName),
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionClient) UpdateSender(req *http.Request) (future PrivateEndpointConnectionUpdateFuture, err error) {
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

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionClient) UpdateResponder(resp *http.Response) (result PrivateEndpointConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
