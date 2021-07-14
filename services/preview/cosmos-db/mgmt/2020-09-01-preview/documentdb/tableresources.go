package documentdb

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

// TableResourcesClient is the azure Cosmos DB Database Service Resource Provider REST API
type TableResourcesClient struct {
	BaseClient
}

// NewTableResourcesClient creates an instance of the TableResourcesClient client.
func NewTableResourcesClient(subscriptionID string) TableResourcesClient {
	return NewTableResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTableResourcesClientWithBaseURI creates an instance of the TableResourcesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTableResourcesClientWithBaseURI(baseURI string, subscriptionID string) TableResourcesClient {
	return TableResourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateUpdateTable create or update an Azure Cosmos DB Table
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - cosmos DB database account name.
// tableName - cosmos DB table name.
// createUpdateTableParameters - the parameters to provide for the current Table.
func (client TableResourcesClient) CreateUpdateTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, createUpdateTableParameters TableCreateUpdateParameters) (result TableResourcesCreateUpdateTableFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TableResourcesClient.CreateUpdateTable")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-z0-9]+(-[a-z0-9]+)*`, Chain: nil}}},
		{TargetValue: createUpdateTableParameters,
			Constraints: []validation.Constraint{{Target: "createUpdateTableParameters.TableCreateUpdateProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "createUpdateTableParameters.TableCreateUpdateProperties.Resource", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "createUpdateTableParameters.TableCreateUpdateProperties.Resource.ID", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("documentdb.TableResourcesClient", "CreateUpdateTable", err.Error())
	}

	req, err := client.CreateUpdateTablePreparer(ctx, resourceGroupName, accountName, tableName, createUpdateTableParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "CreateUpdateTable", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateUpdateTableSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "CreateUpdateTable", nil, "Failure sending request")
		return
	}

	return
}

// CreateUpdateTablePreparer prepares the CreateUpdateTable request.
func (client TableResourcesClient) CreateUpdateTablePreparer(ctx context.Context, resourceGroupName string, accountName string, tableName string, createUpdateTableParameters TableCreateUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tableName":         autorest.Encode("path", tableName),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}", pathParameters),
		autorest.WithJSON(createUpdateTableParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateUpdateTableSender sends the CreateUpdateTable request. The method will close the
// http.Response Body if it receives an error.
func (client TableResourcesClient) CreateUpdateTableSender(req *http.Request) (future TableResourcesCreateUpdateTableFuture, err error) {
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

// CreateUpdateTableResponder handles the response to the CreateUpdateTable request. The method always
// closes the http.Response Body.
func (client TableResourcesClient) CreateUpdateTableResponder(resp *http.Response) (result TableGetResults, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteTable deletes an existing Azure Cosmos DB Table.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - cosmos DB database account name.
// tableName - cosmos DB table name.
func (client TableResourcesClient) DeleteTable(ctx context.Context, resourceGroupName string, accountName string, tableName string) (result TableResourcesDeleteTableFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TableResourcesClient.DeleteTable")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-z0-9]+(-[a-z0-9]+)*`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.TableResourcesClient", "DeleteTable", err.Error())
	}

	req, err := client.DeleteTablePreparer(ctx, resourceGroupName, accountName, tableName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "DeleteTable", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteTableSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "DeleteTable", nil, "Failure sending request")
		return
	}

	return
}

// DeleteTablePreparer prepares the DeleteTable request.
func (client TableResourcesClient) DeleteTablePreparer(ctx context.Context, resourceGroupName string, accountName string, tableName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tableName":         autorest.Encode("path", tableName),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteTableSender sends the DeleteTable request. The method will close the
// http.Response Body if it receives an error.
func (client TableResourcesClient) DeleteTableSender(req *http.Request) (future TableResourcesDeleteTableFuture, err error) {
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

// DeleteTableResponder handles the response to the DeleteTable request. The method always
// closes the http.Response Body.
func (client TableResourcesClient) DeleteTableResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetTable gets the Tables under an existing Azure Cosmos DB database account with the provided name.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - cosmos DB database account name.
// tableName - cosmos DB table name.
func (client TableResourcesClient) GetTable(ctx context.Context, resourceGroupName string, accountName string, tableName string) (result TableGetResults, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TableResourcesClient.GetTable")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-z0-9]+(-[a-z0-9]+)*`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.TableResourcesClient", "GetTable", err.Error())
	}

	req, err := client.GetTablePreparer(ctx, resourceGroupName, accountName, tableName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "GetTable", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetTableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "GetTable", resp, "Failure sending request")
		return
	}

	result, err = client.GetTableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "GetTable", resp, "Failure responding to request")
		return
	}

	return
}

// GetTablePreparer prepares the GetTable request.
func (client TableResourcesClient) GetTablePreparer(ctx context.Context, resourceGroupName string, accountName string, tableName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tableName":         autorest.Encode("path", tableName),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetTableSender sends the GetTable request. The method will close the
// http.Response Body if it receives an error.
func (client TableResourcesClient) GetTableSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetTableResponder handles the response to the GetTable request. The method always
// closes the http.Response Body.
func (client TableResourcesClient) GetTableResponder(resp *http.Response) (result TableGetResults, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetTableThroughput gets the RUs per second of the Table under an existing Azure Cosmos DB database account with the
// provided name.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - cosmos DB database account name.
// tableName - cosmos DB table name.
func (client TableResourcesClient) GetTableThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string) (result ThroughputSettingsGetResults, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TableResourcesClient.GetTableThroughput")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-z0-9]+(-[a-z0-9]+)*`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.TableResourcesClient", "GetTableThroughput", err.Error())
	}

	req, err := client.GetTableThroughputPreparer(ctx, resourceGroupName, accountName, tableName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "GetTableThroughput", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetTableThroughputSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "GetTableThroughput", resp, "Failure sending request")
		return
	}

	result, err = client.GetTableThroughputResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "GetTableThroughput", resp, "Failure responding to request")
		return
	}

	return
}

// GetTableThroughputPreparer prepares the GetTableThroughput request.
func (client TableResourcesClient) GetTableThroughputPreparer(ctx context.Context, resourceGroupName string, accountName string, tableName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tableName":         autorest.Encode("path", tableName),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetTableThroughputSender sends the GetTableThroughput request. The method will close the
// http.Response Body if it receives an error.
func (client TableResourcesClient) GetTableThroughputSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetTableThroughputResponder handles the response to the GetTableThroughput request. The method always
// closes the http.Response Body.
func (client TableResourcesClient) GetTableThroughputResponder(resp *http.Response) (result ThroughputSettingsGetResults, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListTables lists the Tables under an existing Azure Cosmos DB database account.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - cosmos DB database account name.
func (client TableResourcesClient) ListTables(ctx context.Context, resourceGroupName string, accountName string) (result TableListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TableResourcesClient.ListTables")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-z0-9]+(-[a-z0-9]+)*`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.TableResourcesClient", "ListTables", err.Error())
	}

	req, err := client.ListTablesPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "ListTables", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListTablesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "ListTables", resp, "Failure sending request")
		return
	}

	result, err = client.ListTablesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "ListTables", resp, "Failure responding to request")
		return
	}

	return
}

// ListTablesPreparer prepares the ListTables request.
func (client TableResourcesClient) ListTablesPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListTablesSender sends the ListTables request. The method will close the
// http.Response Body if it receives an error.
func (client TableResourcesClient) ListTablesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListTablesResponder handles the response to the ListTables request. The method always
// closes the http.Response Body.
func (client TableResourcesClient) ListTablesResponder(resp *http.Response) (result TableListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// MigrateTableToAutoscale migrate an Azure Cosmos DB Table from manual throughput to autoscale
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - cosmos DB database account name.
// tableName - cosmos DB table name.
func (client TableResourcesClient) MigrateTableToAutoscale(ctx context.Context, resourceGroupName string, accountName string, tableName string) (result TableResourcesMigrateTableToAutoscaleFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TableResourcesClient.MigrateTableToAutoscale")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-z0-9]+(-[a-z0-9]+)*`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.TableResourcesClient", "MigrateTableToAutoscale", err.Error())
	}

	req, err := client.MigrateTableToAutoscalePreparer(ctx, resourceGroupName, accountName, tableName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "MigrateTableToAutoscale", nil, "Failure preparing request")
		return
	}

	result, err = client.MigrateTableToAutoscaleSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "MigrateTableToAutoscale", nil, "Failure sending request")
		return
	}

	return
}

// MigrateTableToAutoscalePreparer prepares the MigrateTableToAutoscale request.
func (client TableResourcesClient) MigrateTableToAutoscalePreparer(ctx context.Context, resourceGroupName string, accountName string, tableName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tableName":         autorest.Encode("path", tableName),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default/migrateToAutoscale", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// MigrateTableToAutoscaleSender sends the MigrateTableToAutoscale request. The method will close the
// http.Response Body if it receives an error.
func (client TableResourcesClient) MigrateTableToAutoscaleSender(req *http.Request) (future TableResourcesMigrateTableToAutoscaleFuture, err error) {
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

// MigrateTableToAutoscaleResponder handles the response to the MigrateTableToAutoscale request. The method always
// closes the http.Response Body.
func (client TableResourcesClient) MigrateTableToAutoscaleResponder(resp *http.Response) (result ThroughputSettingsGetResults, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// MigrateTableToManualThroughput migrate an Azure Cosmos DB Table from autoscale to manual throughput
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - cosmos DB database account name.
// tableName - cosmos DB table name.
func (client TableResourcesClient) MigrateTableToManualThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string) (result TableResourcesMigrateTableToManualThroughputFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TableResourcesClient.MigrateTableToManualThroughput")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-z0-9]+(-[a-z0-9]+)*`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.TableResourcesClient", "MigrateTableToManualThroughput", err.Error())
	}

	req, err := client.MigrateTableToManualThroughputPreparer(ctx, resourceGroupName, accountName, tableName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "MigrateTableToManualThroughput", nil, "Failure preparing request")
		return
	}

	result, err = client.MigrateTableToManualThroughputSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "MigrateTableToManualThroughput", nil, "Failure sending request")
		return
	}

	return
}

// MigrateTableToManualThroughputPreparer prepares the MigrateTableToManualThroughput request.
func (client TableResourcesClient) MigrateTableToManualThroughputPreparer(ctx context.Context, resourceGroupName string, accountName string, tableName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tableName":         autorest.Encode("path", tableName),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default/migrateToManualThroughput", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// MigrateTableToManualThroughputSender sends the MigrateTableToManualThroughput request. The method will close the
// http.Response Body if it receives an error.
func (client TableResourcesClient) MigrateTableToManualThroughputSender(req *http.Request) (future TableResourcesMigrateTableToManualThroughputFuture, err error) {
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

// MigrateTableToManualThroughputResponder handles the response to the MigrateTableToManualThroughput request. The method always
// closes the http.Response Body.
func (client TableResourcesClient) MigrateTableToManualThroughputResponder(resp *http.Response) (result ThroughputSettingsGetResults, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateTableThroughput update RUs per second of an Azure Cosmos DB Table
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - cosmos DB database account name.
// tableName - cosmos DB table name.
// updateThroughputParameters - the parameters to provide for the RUs per second of the current Table.
func (client TableResourcesClient) UpdateTableThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, updateThroughputParameters ThroughputSettingsUpdateParameters) (result TableResourcesUpdateTableThroughputFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TableResourcesClient.UpdateTableThroughput")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-z0-9]+(-[a-z0-9]+)*`, Chain: nil}}},
		{TargetValue: updateThroughputParameters,
			Constraints: []validation.Constraint{{Target: "updateThroughputParameters.ThroughputSettingsUpdateProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "updateThroughputParameters.ThroughputSettingsUpdateProperties.Resource", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "updateThroughputParameters.ThroughputSettingsUpdateProperties.Resource.AutoscaleSettings", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "updateThroughputParameters.ThroughputSettingsUpdateProperties.Resource.AutoscaleSettings.MaxThroughput", Name: validation.Null, Rule: true, Chain: nil}}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("documentdb.TableResourcesClient", "UpdateTableThroughput", err.Error())
	}

	req, err := client.UpdateTableThroughputPreparer(ctx, resourceGroupName, accountName, tableName, updateThroughputParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "UpdateTableThroughput", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateTableThroughputSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.TableResourcesClient", "UpdateTableThroughput", nil, "Failure sending request")
		return
	}

	return
}

// UpdateTableThroughputPreparer prepares the UpdateTableThroughput request.
func (client TableResourcesClient) UpdateTableThroughputPreparer(ctx context.Context, resourceGroupName string, accountName string, tableName string, updateThroughputParameters ThroughputSettingsUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tableName":         autorest.Encode("path", tableName),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default", pathParameters),
		autorest.WithJSON(updateThroughputParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateTableThroughputSender sends the UpdateTableThroughput request. The method will close the
// http.Response Body if it receives an error.
func (client TableResourcesClient) UpdateTableThroughputSender(req *http.Request) (future TableResourcesUpdateTableThroughputFuture, err error) {
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

// UpdateTableThroughputResponder handles the response to the UpdateTableThroughput request. The method always
// closes the http.Response Body.
func (client TableResourcesClient) UpdateTableThroughputResponder(resp *http.Response) (result ThroughputSettingsGetResults, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
