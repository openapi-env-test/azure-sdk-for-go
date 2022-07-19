package synapse

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

// WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient is the azure Synapse Analytics Management Client
type WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient struct {
	BaseClient
}

// NewWorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient creates an instance of the
// WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient client.
func NewWorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient(subscriptionID string) WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient {
	return NewWorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientWithBaseURI creates an instance of the
// WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientWithBaseURI(baseURI string, subscriptionID string) WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient {
	return WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get workspace managed sql server's minimal tls settings.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// dedicatedSQLminimalTLSSettingsName - the name of the dedicated sql minimal tls settings.
func (client WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, dedicatedSQLminimalTLSSettingsName string) (result DedicatedSQLminimalTLSSettings, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, dedicatedSQLminimalTLSSettingsName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, dedicatedSQLminimalTLSSettingsName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dedicatedSQLminimalTlsSettingsName": autorest.Encode("path", dedicatedSQLminimalTLSSettingsName),
		"resourceGroupName":                  autorest.Encode("path", resourceGroupName),
		"subscriptionId":                     autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                      autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/dedicatedSQLminimalTlsSettings/{dedicatedSQLminimalTlsSettingsName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) GetResponder(resp *http.Response) (result DedicatedSQLminimalTLSSettings, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list workspace managed sql server's minimal tls settings.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
func (client WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) List(ctx context.Context, resourceGroupName string, workspaceName string) (result DedicatedSQLminimalTLSSettingsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient.List")
		defer func() {
			sc := -1
			if result.dsltslr.Response.Response != nil {
				sc = result.dsltslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dsltslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient", "List", resp, "Failure sending request")
		return
	}

	result.dsltslr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.dsltslr.hasNextLink() && result.dsltslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) ListPreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/dedicatedSQLminimalTlsSettings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) ListResponder(resp *http.Response) (result DedicatedSQLminimalTLSSettingsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) listNextResults(ctx context.Context, lastResults DedicatedSQLminimalTLSSettingsListResult) (result DedicatedSQLminimalTLSSettingsListResult, err error) {
	req, err := lastResults.dedicatedSQLminimalTLSSettingsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result DedicatedSQLminimalTLSSettingsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, workspaceName)
	return
}

// Update update workspace managed sql server's minimal tls settings.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// parameters - minimal tls settings
func (client WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) Update(ctx context.Context, resourceGroupName string, workspaceName string, parameters DedicatedSQLminimalTLSSettings) (result WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, workspaceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, parameters DedicatedSQLminimalTLSSettings) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dedicatedSQLminimalTlsSettingsName": autorest.Encode("path", "default"),
		"resourceGroupName":                  autorest.Encode("path", resourceGroupName),
		"subscriptionId":                     autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                      autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Location = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/dedicatedSQLminimalTlsSettings/{dedicatedSQLminimalTlsSettingsName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) UpdateSender(req *http.Request) (future WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsUpdateFuture, err error) {
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
func (client WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) UpdateResponder(resp *http.Response) (result DedicatedSQLminimalTLSSettings, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
