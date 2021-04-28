package securityandcompliance

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

// PrivateEndpointConnectionsAdtAPIClient is the security And Compliance Client
type PrivateEndpointConnectionsAdtAPIClient struct {
	BaseClient
}

// NewPrivateEndpointConnectionsAdtAPIClient creates an instance of the PrivateEndpointConnectionsAdtAPIClient client.
func NewPrivateEndpointConnectionsAdtAPIClient(subscriptionID string) PrivateEndpointConnectionsAdtAPIClient {
	return NewPrivateEndpointConnectionsAdtAPIClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateEndpointConnectionsAdtAPIClientWithBaseURI creates an instance of the
// PrivateEndpointConnectionsAdtAPIClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPrivateEndpointConnectionsAdtAPIClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsAdtAPIClient {
	return PrivateEndpointConnectionsAdtAPIClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate update the state of the specified private endpoint connection associated with the service.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// resourceName - the name of the service instance.
// privateEndpointConnectionName - the name of the private endpoint connection associated with the Azure
// resource
// properties - the private endpoint connection properties.
func (client PrivateEndpointConnectionsAdtAPIClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, properties PrivateEndpointConnection) (result PrivateEndpointConnectionsAdtAPICreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsAdtAPIClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: properties,
			Constraints: []validation.Constraint{{Target: "properties.PrivateEndpointConnectionProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "properties.PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("securityandcompliance.PrivateEndpointConnectionsAdtAPIClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, resourceName, privateEndpointConnectionName, properties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateEndpointConnectionsAdtAPIClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateEndpointConnectionsAdtAPIClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PrivateEndpointConnectionsAdtAPIClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, properties PrivateEndpointConnection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"resourceName":                  autorest.Encode("path", resourceName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-08"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForO365ManagementActivityAPI/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}", pathParameters),
		autorest.WithJSON(properties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsAdtAPIClient) CreateOrUpdateSender(req *http.Request) (future PrivateEndpointConnectionsAdtAPICreateOrUpdateFuture, err error) {
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

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsAdtAPIClient) CreateOrUpdateResponder(resp *http.Response) (result PrivateEndpointConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a private endpoint connection.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// resourceName - the name of the service instance.
// privateEndpointConnectionName - the name of the private endpoint connection associated with the Azure
// resource
func (client PrivateEndpointConnectionsAdtAPIClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string) (result PrivateEndpointConnectionsAdtAPIDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsAdtAPIClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityandcompliance.PrivateEndpointConnectionsAdtAPIClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName, privateEndpointConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateEndpointConnectionsAdtAPIClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateEndpointConnectionsAdtAPIClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PrivateEndpointConnectionsAdtAPIClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"resourceName":                  autorest.Encode("path", resourceName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-08"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForO365ManagementActivityAPI/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsAdtAPIClient) DeleteSender(req *http.Request) (future PrivateEndpointConnectionsAdtAPIDeleteFuture, err error) {
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

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsAdtAPIClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified private endpoint connection associated with the service.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// resourceName - the name of the service instance.
// privateEndpointConnectionName - the name of the private endpoint connection associated with the Azure
// resource
func (client PrivateEndpointConnectionsAdtAPIClient) Get(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string) (result PrivateEndpointConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsAdtAPIClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityandcompliance.PrivateEndpointConnectionsAdtAPIClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName, privateEndpointConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateEndpointConnectionsAdtAPIClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateEndpointConnectionsAdtAPIClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateEndpointConnectionsAdtAPIClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrivateEndpointConnectionsAdtAPIClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"resourceName":                  autorest.Encode("path", resourceName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-08"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForO365ManagementActivityAPI/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsAdtAPIClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsAdtAPIClient) GetResponder(resp *http.Response) (result PrivateEndpointConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByService lists all private endpoint connections for a service.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// resourceName - the name of the service instance.
func (client PrivateEndpointConnectionsAdtAPIClient) ListByService(ctx context.Context, resourceGroupName string, resourceName string) (result PrivateEndpointConnectionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsAdtAPIClient.ListByService")
		defer func() {
			sc := -1
			if result.peclr.Response.Response != nil {
				sc = result.peclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityandcompliance.PrivateEndpointConnectionsAdtAPIClient", "ListByService", err.Error())
	}

	result.fn = client.listByServiceNextResults
	req, err := client.ListByServicePreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateEndpointConnectionsAdtAPIClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.peclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateEndpointConnectionsAdtAPIClient", "ListByService", resp, "Failure sending request")
		return
	}

	result.peclr, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateEndpointConnectionsAdtAPIClient", "ListByService", resp, "Failure responding to request")
		return
	}
	if result.peclr.hasNextLink() && result.peclr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client PrivateEndpointConnectionsAdtAPIClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-08"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForO365ManagementActivityAPI/{resourceName}/privateEndpointConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsAdtAPIClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsAdtAPIClient) ListByServiceResponder(resp *http.Response) (result PrivateEndpointConnectionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServiceNextResults retrieves the next set of results, if any.
func (client PrivateEndpointConnectionsAdtAPIClient) listByServiceNextResults(ctx context.Context, lastResults PrivateEndpointConnectionListResult) (result PrivateEndpointConnectionListResult, err error) {
	req, err := lastResults.privateEndpointConnectionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "securityandcompliance.PrivateEndpointConnectionsAdtAPIClient", "listByServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "securityandcompliance.PrivateEndpointConnectionsAdtAPIClient", "listByServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateEndpointConnectionsAdtAPIClient", "listByServiceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateEndpointConnectionsAdtAPIClient) ListByServiceComplete(ctx context.Context, resourceGroupName string, resourceName string) (result PrivateEndpointConnectionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsAdtAPIClient.ListByService")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByService(ctx, resourceGroupName, resourceName)
	return
}
