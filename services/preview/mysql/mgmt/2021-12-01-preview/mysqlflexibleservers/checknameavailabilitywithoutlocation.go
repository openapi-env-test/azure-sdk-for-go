package mysqlflexibleservers

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

// CheckNameAvailabilityWithoutLocationClient is the the Microsoft Azure management API provides create, read, update,
// and delete functionality for Azure MySQL resources including servers, databases, firewall rules, VNET rules, log
// files and configurations with new business model.
type CheckNameAvailabilityWithoutLocationClient struct {
	BaseClient
}

// NewCheckNameAvailabilityWithoutLocationClient creates an instance of the CheckNameAvailabilityWithoutLocationClient
// client.
func NewCheckNameAvailabilityWithoutLocationClient(subscriptionID string) CheckNameAvailabilityWithoutLocationClient {
	return NewCheckNameAvailabilityWithoutLocationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCheckNameAvailabilityWithoutLocationClientWithBaseURI creates an instance of the
// CheckNameAvailabilityWithoutLocationClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewCheckNameAvailabilityWithoutLocationClientWithBaseURI(baseURI string, subscriptionID string) CheckNameAvailabilityWithoutLocationClient {
	return CheckNameAvailabilityWithoutLocationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Execute check the availability of name for server
// Parameters:
// nameAvailabilityRequest - the required parameters for checking if server name is available.
func (client CheckNameAvailabilityWithoutLocationClient) Execute(ctx context.Context, nameAvailabilityRequest NameAvailabilityRequest) (result NameAvailability, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CheckNameAvailabilityWithoutLocationClient.Execute")
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
		{TargetValue: nameAvailabilityRequest,
			Constraints: []validation.Constraint{{Target: "nameAvailabilityRequest.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("mysqlflexibleservers.CheckNameAvailabilityWithoutLocationClient", "Execute", err.Error())
	}

	req, err := client.ExecutePreparer(ctx, nameAvailabilityRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysqlflexibleservers.CheckNameAvailabilityWithoutLocationClient", "Execute", nil, "Failure preparing request")
		return
	}

	resp, err := client.ExecuteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mysqlflexibleservers.CheckNameAvailabilityWithoutLocationClient", "Execute", resp, "Failure sending request")
		return
	}

	result, err = client.ExecuteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysqlflexibleservers.CheckNameAvailabilityWithoutLocationClient", "Execute", resp, "Failure responding to request")
		return
	}

	return
}

// ExecutePreparer prepares the Execute request.
func (client CheckNameAvailabilityWithoutLocationClient) ExecutePreparer(ctx context.Context, nameAvailabilityRequest NameAvailabilityRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/checkNameAvailability", pathParameters),
		autorest.WithJSON(nameAvailabilityRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExecuteSender sends the Execute request. The method will close the
// http.Response Body if it receives an error.
func (client CheckNameAvailabilityWithoutLocationClient) ExecuteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ExecuteResponder handles the response to the Execute request. The method always
// closes the http.Response Body.
func (client CheckNameAvailabilityWithoutLocationClient) ExecuteResponder(resp *http.Response) (result NameAvailability, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
