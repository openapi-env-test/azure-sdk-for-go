package migrate

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

// DatabaseInstancesClient is the migrate your workloads to Azure.
type DatabaseInstancesClient struct {
	BaseClient
}

// NewDatabaseInstancesClient creates an instance of the DatabaseInstancesClient client.
func NewDatabaseInstancesClient(subscriptionID string) DatabaseInstancesClient {
	return NewDatabaseInstancesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDatabaseInstancesClientWithBaseURI creates an instance of the DatabaseInstancesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewDatabaseInstancesClientWithBaseURI(baseURI string, subscriptionID string) DatabaseInstancesClient {
	return DatabaseInstancesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// EnumerateDatabaseInstances sends the enumerate database instances request.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that migrate project is part of.
// migrateProjectName - name of the Azure Migrate project.
// continuationToken - the continuation token.
// pageSize - the number of items to be returned in a single page. This value is honored only if it is less
// than the 100.
func (client DatabaseInstancesClient) EnumerateDatabaseInstances(ctx context.Context, resourceGroupName string, migrateProjectName string, continuationToken string, pageSize *int32) (result DatabaseInstanceCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseInstancesClient.EnumerateDatabaseInstances")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.EnumerateDatabaseInstancesPreparer(ctx, resourceGroupName, migrateProjectName, continuationToken, pageSize)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.DatabaseInstancesClient", "EnumerateDatabaseInstances", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnumerateDatabaseInstancesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.DatabaseInstancesClient", "EnumerateDatabaseInstances", resp, "Failure sending request")
		return
	}

	result, err = client.EnumerateDatabaseInstancesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.DatabaseInstancesClient", "EnumerateDatabaseInstances", resp, "Failure responding to request")
	}

	return
}

// EnumerateDatabaseInstancesPreparer prepares the EnumerateDatabaseInstances request.
func (client DatabaseInstancesClient) EnumerateDatabaseInstancesPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, continuationToken string, pageSize *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(continuationToken) > 0 {
		queryParameters["continuationToken"] = autorest.Encode("query", continuationToken)
	}
	if pageSize != nil {
		queryParameters["pageSize"] = autorest.Encode("query", *pageSize)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/databaseInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnumerateDatabaseInstancesSender sends the EnumerateDatabaseInstances request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseInstancesClient) EnumerateDatabaseInstancesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// EnumerateDatabaseInstancesResponder handles the response to the EnumerateDatabaseInstances request. The method always
// closes the http.Response Body.
func (client DatabaseInstancesClient) EnumerateDatabaseInstancesResponder(resp *http.Response) (result DatabaseInstanceCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDatabaseInstance sends the get database instance request.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that migrate project is part of.
// migrateProjectName - name of the Azure Migrate project.
// databaseInstanceName - unique name of a database instance in Azure migration hub.
func (client DatabaseInstancesClient) GetDatabaseInstance(ctx context.Context, resourceGroupName string, migrateProjectName string, databaseInstanceName string) (result DatabaseInstance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseInstancesClient.GetDatabaseInstance")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDatabaseInstancePreparer(ctx, resourceGroupName, migrateProjectName, databaseInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.DatabaseInstancesClient", "GetDatabaseInstance", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDatabaseInstanceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.DatabaseInstancesClient", "GetDatabaseInstance", resp, "Failure sending request")
		return
	}

	result, err = client.GetDatabaseInstanceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.DatabaseInstancesClient", "GetDatabaseInstance", resp, "Failure responding to request")
	}

	return
}

// GetDatabaseInstancePreparer prepares the GetDatabaseInstance request.
func (client DatabaseInstancesClient) GetDatabaseInstancePreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, databaseInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseInstanceName": autorest.Encode("path", databaseInstanceName),
		"migrateProjectName":   autorest.Encode("path", migrateProjectName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/databaseInstances/{databaseInstanceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDatabaseInstanceSender sends the GetDatabaseInstance request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseInstancesClient) GetDatabaseInstanceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetDatabaseInstanceResponder handles the response to the GetDatabaseInstance request. The method always
// closes the http.Response Body.
func (client DatabaseInstancesClient) GetDatabaseInstanceResponder(resp *http.Response) (result DatabaseInstance, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
