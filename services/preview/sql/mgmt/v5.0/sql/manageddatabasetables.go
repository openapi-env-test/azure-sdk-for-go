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

// ManagedDatabaseTablesClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type ManagedDatabaseTablesClient struct {
	BaseClient
}

// NewManagedDatabaseTablesClient creates an instance of the ManagedDatabaseTablesClient client.
func NewManagedDatabaseTablesClient(subscriptionID string) ManagedDatabaseTablesClient {
	return NewManagedDatabaseTablesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagedDatabaseTablesClientWithBaseURI creates an instance of the ManagedDatabaseTablesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewManagedDatabaseTablesClientWithBaseURI(baseURI string, subscriptionID string) ManagedDatabaseTablesClient {
	return ManagedDatabaseTablesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get managed database table
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// databaseName - the name of the database.
// schemaName - the name of the schema.
// tableName - the name of the table.
func (client ManagedDatabaseTablesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string) (result DatabaseTable, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedDatabaseTablesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseTablesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseTablesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseTablesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ManagedDatabaseTablesClient) GetPreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":        autorest.Encode("path", databaseName),
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"schemaName":          autorest.Encode("path", schemaName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
		"tableName":           autorest.Encode("path", tableName),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedDatabaseTablesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagedDatabaseTablesClient) GetResponder(resp *http.Response) (result DatabaseTable, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySchema list managed database tables
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// databaseName - the name of the database.
// schemaName - the name of the schema.
// filter - an OData filter expression that filters elements in the collection.
func (client ManagedDatabaseTablesClient) ListBySchema(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, filter string) (result DatabaseTableListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedDatabaseTablesClient.ListBySchema")
		defer func() {
			sc := -1
			if result.dtlr.Response.Response != nil {
				sc = result.dtlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySchemaNextResults
	req, err := client.ListBySchemaPreparer(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseTablesClient", "ListBySchema", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySchemaSender(req)
	if err != nil {
		result.dtlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseTablesClient", "ListBySchema", resp, "Failure sending request")
		return
	}

	result.dtlr, err = client.ListBySchemaResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseTablesClient", "ListBySchema", resp, "Failure responding to request")
		return
	}
	if result.dtlr.hasNextLink() && result.dtlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySchemaPreparer prepares the ListBySchema request.
func (client ManagedDatabaseTablesClient) ListBySchemaPreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":        autorest.Encode("path", databaseName),
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"schemaName":          autorest.Encode("path", schemaName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas/{schemaName}/tables", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySchemaSender sends the ListBySchema request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedDatabaseTablesClient) ListBySchemaSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySchemaResponder handles the response to the ListBySchema request. The method always
// closes the http.Response Body.
func (client ManagedDatabaseTablesClient) ListBySchemaResponder(resp *http.Response) (result DatabaseTableListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySchemaNextResults retrieves the next set of results, if any.
func (client ManagedDatabaseTablesClient) listBySchemaNextResults(ctx context.Context, lastResults DatabaseTableListResult) (result DatabaseTableListResult, err error) {
	req, err := lastResults.databaseTableListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ManagedDatabaseTablesClient", "listBySchemaNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySchemaSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ManagedDatabaseTablesClient", "listBySchemaNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySchemaResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseTablesClient", "listBySchemaNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySchemaComplete enumerates all values, automatically crossing page boundaries as required.
func (client ManagedDatabaseTablesClient) ListBySchemaComplete(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, filter string) (result DatabaseTableListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedDatabaseTablesClient.ListBySchema")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySchema(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, filter)
	return
}
