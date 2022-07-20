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

// ExtendedSQLPoolBlobAuditingPoliciesClient is the azure Synapse Analytics Management Client
type ExtendedSQLPoolBlobAuditingPoliciesClient struct {
	BaseClient
}

// NewExtendedSQLPoolBlobAuditingPoliciesClient creates an instance of the ExtendedSQLPoolBlobAuditingPoliciesClient
// client.
func NewExtendedSQLPoolBlobAuditingPoliciesClient(subscriptionID string) ExtendedSQLPoolBlobAuditingPoliciesClient {
	return NewExtendedSQLPoolBlobAuditingPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExtendedSQLPoolBlobAuditingPoliciesClientWithBaseURI creates an instance of the
// ExtendedSQLPoolBlobAuditingPoliciesClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewExtendedSQLPoolBlobAuditingPoliciesClientWithBaseURI(baseURI string, subscriptionID string) ExtendedSQLPoolBlobAuditingPoliciesClient {
	return ExtendedSQLPoolBlobAuditingPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an extended Sql pool's blob auditing policy.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// SQLPoolName - SQL pool name
// parameters - the extended Sql pool blob auditing policy.
func (client ExtendedSQLPoolBlobAuditingPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, parameters ExtendedSQLPoolBlobAuditingPolicy) (result ExtendedSQLPoolBlobAuditingPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtendedSQLPoolBlobAuditingPoliciesClient.CreateOrUpdate")
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
		return result, validation.NewError("synapse.ExtendedSQLPoolBlobAuditingPoliciesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.ExtendedSQLPoolBlobAuditingPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.ExtendedSQLPoolBlobAuditingPoliciesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.ExtendedSQLPoolBlobAuditingPoliciesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ExtendedSQLPoolBlobAuditingPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, parameters ExtendedSQLPoolBlobAuditingPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blobAuditingPolicyName": autorest.Encode("path", "default"),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sqlPoolName":            autorest.Encode("path", SQLPoolName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/extendedAuditingSettings/{blobAuditingPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ExtendedSQLPoolBlobAuditingPoliciesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ExtendedSQLPoolBlobAuditingPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result ExtendedSQLPoolBlobAuditingPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets an extended Sql pool's blob auditing policy.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// SQLPoolName - SQL pool name
func (client ExtendedSQLPoolBlobAuditingPoliciesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result ExtendedSQLPoolBlobAuditingPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtendedSQLPoolBlobAuditingPoliciesClient.Get")
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
		return result, validation.NewError("synapse.ExtendedSQLPoolBlobAuditingPoliciesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, SQLPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.ExtendedSQLPoolBlobAuditingPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.ExtendedSQLPoolBlobAuditingPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.ExtendedSQLPoolBlobAuditingPoliciesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExtendedSQLPoolBlobAuditingPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blobAuditingPolicyName": autorest.Encode("path", "default"),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sqlPoolName":            autorest.Encode("path", SQLPoolName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/extendedAuditingSettings/{blobAuditingPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExtendedSQLPoolBlobAuditingPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExtendedSQLPoolBlobAuditingPoliciesClient) GetResponder(resp *http.Response) (result ExtendedSQLPoolBlobAuditingPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySQLPool lists extended auditing settings of a Sql pool.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// SQLPoolName - SQL pool name
func (client ExtendedSQLPoolBlobAuditingPoliciesClient) ListBySQLPool(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result ExtendedSQLPoolBlobAuditingPolicyListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtendedSQLPoolBlobAuditingPoliciesClient.ListBySQLPool")
		defer func() {
			sc := -1
			if result.espbaplr.Response.Response != nil {
				sc = result.espbaplr.Response.Response.StatusCode
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
		return result, validation.NewError("synapse.ExtendedSQLPoolBlobAuditingPoliciesClient", "ListBySQLPool", err.Error())
	}

	result.fn = client.listBySQLPoolNextResults
	req, err := client.ListBySQLPoolPreparer(ctx, resourceGroupName, workspaceName, SQLPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.ExtendedSQLPoolBlobAuditingPoliciesClient", "ListBySQLPool", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySQLPoolSender(req)
	if err != nil {
		result.espbaplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.ExtendedSQLPoolBlobAuditingPoliciesClient", "ListBySQLPool", resp, "Failure sending request")
		return
	}

	result.espbaplr, err = client.ListBySQLPoolResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.ExtendedSQLPoolBlobAuditingPoliciesClient", "ListBySQLPool", resp, "Failure responding to request")
		return
	}
	if result.espbaplr.hasNextLink() && result.espbaplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySQLPoolPreparer prepares the ListBySQLPool request.
func (client ExtendedSQLPoolBlobAuditingPoliciesClient) ListBySQLPoolPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sqlPoolName":       autorest.Encode("path", SQLPoolName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/extendedAuditingSettings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySQLPoolSender sends the ListBySQLPool request. The method will close the
// http.Response Body if it receives an error.
func (client ExtendedSQLPoolBlobAuditingPoliciesClient) ListBySQLPoolSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySQLPoolResponder handles the response to the ListBySQLPool request. The method always
// closes the http.Response Body.
func (client ExtendedSQLPoolBlobAuditingPoliciesClient) ListBySQLPoolResponder(resp *http.Response) (result ExtendedSQLPoolBlobAuditingPolicyListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySQLPoolNextResults retrieves the next set of results, if any.
func (client ExtendedSQLPoolBlobAuditingPoliciesClient) listBySQLPoolNextResults(ctx context.Context, lastResults ExtendedSQLPoolBlobAuditingPolicyListResult) (result ExtendedSQLPoolBlobAuditingPolicyListResult, err error) {
	req, err := lastResults.extendedSQLPoolBlobAuditingPolicyListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "synapse.ExtendedSQLPoolBlobAuditingPoliciesClient", "listBySQLPoolNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySQLPoolSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "synapse.ExtendedSQLPoolBlobAuditingPoliciesClient", "listBySQLPoolNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySQLPoolResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.ExtendedSQLPoolBlobAuditingPoliciesClient", "listBySQLPoolNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySQLPoolComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExtendedSQLPoolBlobAuditingPoliciesClient) ListBySQLPoolComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result ExtendedSQLPoolBlobAuditingPolicyListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtendedSQLPoolBlobAuditingPoliciesClient.ListBySQLPool")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySQLPool(ctx, resourceGroupName, workspaceName, SQLPoolName)
	return
}
