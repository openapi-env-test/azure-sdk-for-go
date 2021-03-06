package sql

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// DatabaseRecommendedActionsClient is the the Azure SQL Database management API provides a RESTful set of web services
// that interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve,
// update, and delete databases.
type DatabaseRecommendedActionsClient struct {
	BaseClient
}

// NewDatabaseRecommendedActionsClient creates an instance of the DatabaseRecommendedActionsClient client.
func NewDatabaseRecommendedActionsClient(subscriptionID string) DatabaseRecommendedActionsClient {
	return NewDatabaseRecommendedActionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDatabaseRecommendedActionsClientWithBaseURI creates an instance of the DatabaseRecommendedActionsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewDatabaseRecommendedActionsClientWithBaseURI(baseURI string, subscriptionID string) DatabaseRecommendedActionsClient {
	return DatabaseRecommendedActionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a database recommended action.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// advisorName - the name of the Database Advisor.
// recommendedActionName - the name of Database Recommended Action.
func (client DatabaseRecommendedActionsClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string, recommendedActionName string) (result RecommendedAction, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseRecommendedActionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, databaseName, advisorName, recommendedActionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseRecommendedActionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DatabaseRecommendedActionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseRecommendedActionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DatabaseRecommendedActionsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string, recommendedActionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"advisorName":           autorest.Encode("path", advisorName),
		"databaseName":          autorest.Encode("path", databaseName),
		"recommendedActionName": autorest.Encode("path", recommendedActionName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"serverName":            autorest.Encode("path", serverName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advisors/{advisorName}/recommendedActions/{recommendedActionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseRecommendedActionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DatabaseRecommendedActionsClient) GetResponder(resp *http.Response) (result RecommendedAction, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDatabaseAdvisor gets list of Database Recommended Actions.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// advisorName - the name of the Database Advisor.
func (client DatabaseRecommendedActionsClient) ListByDatabaseAdvisor(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string) (result ListRecommendedAction, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseRecommendedActionsClient.ListByDatabaseAdvisor")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByDatabaseAdvisorPreparer(ctx, resourceGroupName, serverName, databaseName, advisorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseRecommendedActionsClient", "ListByDatabaseAdvisor", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDatabaseAdvisorSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DatabaseRecommendedActionsClient", "ListByDatabaseAdvisor", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDatabaseAdvisorResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseRecommendedActionsClient", "ListByDatabaseAdvisor", resp, "Failure responding to request")
		return
	}

	return
}

// ListByDatabaseAdvisorPreparer prepares the ListByDatabaseAdvisor request.
func (client DatabaseRecommendedActionsClient) ListByDatabaseAdvisorPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"advisorName":       autorest.Encode("path", advisorName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advisors/{advisorName}/recommendedActions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDatabaseAdvisorSender sends the ListByDatabaseAdvisor request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseRecommendedActionsClient) ListByDatabaseAdvisorSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDatabaseAdvisorResponder handles the response to the ListByDatabaseAdvisor request. The method always
// closes the http.Response Body.
func (client DatabaseRecommendedActionsClient) ListByDatabaseAdvisorResponder(resp *http.Response) (result ListRecommendedAction, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a database recommended action.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// advisorName - the name of the Database Advisor.
// recommendedActionName - the name of Database Recommended Action.
// parameters - the requested recommended action resource state.
func (client DatabaseRecommendedActionsClient) Update(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string, recommendedActionName string, parameters RecommendedAction) (result RecommendedAction, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseRecommendedActionsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, serverName, databaseName, advisorName, recommendedActionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseRecommendedActionsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DatabaseRecommendedActionsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseRecommendedActionsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DatabaseRecommendedActionsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string, recommendedActionName string, parameters RecommendedAction) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"advisorName":           autorest.Encode("path", advisorName),
		"databaseName":          autorest.Encode("path", databaseName),
		"recommendedActionName": autorest.Encode("path", recommendedActionName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"serverName":            autorest.Encode("path", serverName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Kind = nil
	parameters.Location = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advisors/{advisorName}/recommendedActions/{recommendedActionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseRecommendedActionsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DatabaseRecommendedActionsClient) UpdateResponder(resp *http.Response) (result RecommendedAction, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
