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

// RestorableSQLResourcesClient is the client for the RestorableSQLResources methods of the Documentdb service.
type RestorableSQLResourcesClient struct {
	BaseClient
}

// NewRestorableSQLResourcesClient creates an instance of the RestorableSQLResourcesClient client.
func NewRestorableSQLResourcesClient(subscriptionID string) RestorableSQLResourcesClient {
	return NewRestorableSQLResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRestorableSQLResourcesClientWithBaseURI creates an instance of the RestorableSQLResourcesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewRestorableSQLResourcesClientWithBaseURI(baseURI string, subscriptionID string) RestorableSQLResourcesClient {
	return RestorableSQLResourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List return a list of database and container combo that exist on the account at the given timestamp and location.
// This helps in scenarios to validate what resources exist at given timestamp and location. This API requires
// 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/.../read' permission.
// Parameters:
// location - cosmos DB region, with spaces between words and each word capitalized.
// instanceID - the instanceId GUID of a restorable database account.
// restoreLocation - the location where the restorable resources are located.
// restoreTimestampInUtc - the timestamp when the restorable resources existed.
func (client RestorableSQLResourcesClient) List(ctx context.Context, location string, instanceID string, restoreLocation string, restoreTimestampInUtc string) (result RestorableSQLResourcesListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RestorableSQLResourcesClient.List")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.RestorableSQLResourcesClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, location, instanceID, restoreLocation, restoreTimestampInUtc)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableSQLResourcesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.RestorableSQLResourcesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableSQLResourcesClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client RestorableSQLResourcesClient) ListPreparer(ctx context.Context, location string, instanceID string, restoreLocation string, restoreTimestampInUtc string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":     autorest.Encode("path", instanceID),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(restoreLocation) > 0 {
		queryParameters["restoreLocation"] = autorest.Encode("query", restoreLocation)
	}
	if len(restoreTimestampInUtc) > 0 {
		queryParameters["restoreTimestampInUtc"] = autorest.Encode("query", restoreTimestampInUtc)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableSqlResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RestorableSQLResourcesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RestorableSQLResourcesClient) ListResponder(resp *http.Response) (result RestorableSQLResourcesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
