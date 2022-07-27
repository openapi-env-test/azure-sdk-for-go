package postgresqlflexibleservers

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

// LocationBasedCapabilitiesClient is the the Microsoft Azure management API provides create, read, update, and delete
// functionality for Azure PostgreSQL resources including servers, databases, firewall rules, VNET rules, security
// alert policies, log files and configurations with new business model.
type LocationBasedCapabilitiesClient struct {
	BaseClient
}

// NewLocationBasedCapabilitiesClient creates an instance of the LocationBasedCapabilitiesClient client.
func NewLocationBasedCapabilitiesClient(subscriptionID string) LocationBasedCapabilitiesClient {
	return NewLocationBasedCapabilitiesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLocationBasedCapabilitiesClientWithBaseURI creates an instance of the LocationBasedCapabilitiesClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewLocationBasedCapabilitiesClientWithBaseURI(baseURI string, subscriptionID string) LocationBasedCapabilitiesClient {
	return LocationBasedCapabilitiesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Execute get capabilities at specified location in a given subscription.
// Parameters:
// locationName - the name of the location.
func (client LocationBasedCapabilitiesClient) Execute(ctx context.Context, locationName string) (result CapabilitiesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationBasedCapabilitiesClient.Execute")
		defer func() {
			sc := -1
			if result.clr.Response.Response != nil {
				sc = result.clr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("postgresqlflexibleservers.LocationBasedCapabilitiesClient", "Execute", err.Error())
	}

	result.fn = client.executeNextResults
	req, err := client.ExecutePreparer(ctx, locationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlflexibleservers.LocationBasedCapabilitiesClient", "Execute", nil, "Failure preparing request")
		return
	}

	resp, err := client.ExecuteSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "postgresqlflexibleservers.LocationBasedCapabilitiesClient", "Execute", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ExecuteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlflexibleservers.LocationBasedCapabilitiesClient", "Execute", resp, "Failure responding to request")
		return
	}
	if result.clr.hasNextLink() && result.clr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ExecutePreparer prepares the Execute request.
func (client LocationBasedCapabilitiesClient) ExecutePreparer(ctx context.Context, locationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/locations/{locationName}/capabilities", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExecuteSender sends the Execute request. The method will close the
// http.Response Body if it receives an error.
func (client LocationBasedCapabilitiesClient) ExecuteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ExecuteResponder handles the response to the Execute request. The method always
// closes the http.Response Body.
func (client LocationBasedCapabilitiesClient) ExecuteResponder(resp *http.Response) (result CapabilitiesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// executeNextResults retrieves the next set of results, if any.
func (client LocationBasedCapabilitiesClient) executeNextResults(ctx context.Context, lastResults CapabilitiesListResult) (result CapabilitiesListResult, err error) {
	req, err := lastResults.capabilitiesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "postgresqlflexibleservers.LocationBasedCapabilitiesClient", "executeNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ExecuteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "postgresqlflexibleservers.LocationBasedCapabilitiesClient", "executeNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ExecuteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlflexibleservers.LocationBasedCapabilitiesClient", "executeNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ExecuteComplete enumerates all values, automatically crossing page boundaries as required.
func (client LocationBasedCapabilitiesClient) ExecuteComplete(ctx context.Context, locationName string) (result CapabilitiesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationBasedCapabilitiesClient.Execute")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.Execute(ctx, locationName)
	return
}
