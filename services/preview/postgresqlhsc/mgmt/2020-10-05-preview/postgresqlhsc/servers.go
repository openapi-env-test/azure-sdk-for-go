package postgresqlhsc

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

// ServersClient is the the Microsoft Azure management API provides create, read, update, and delete functionality for
// Azure PostgreSQL Hyperscale (Citus) resources including server groups, servers, databases, firewall rules, VNET
// rules, security alert policies, log files and configurations.
type ServersClient struct {
	BaseClient
}

// NewServersClient creates an instance of the ServersClient client.
func NewServersClient(subscriptionID string) ServersClient {
	return NewServersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServersClientWithBaseURI creates an instance of the ServersClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewServersClientWithBaseURI(baseURI string, subscriptionID string) ServersClient {
	return ServersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets information about a server in server group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverGroupName - the name of the server group.
// serverName - the name of the server.
func (client ServersClient) Get(ctx context.Context, resourceGroupName string, serverGroupName string, serverName string) (result ServerGroupServer, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServersClient.Get")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: serverGroupName,
			Constraints: []validation.Constraint{{Target: "serverGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "serverGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: serverName,
			Constraints: []validation.Constraint{{Target: "serverName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "serverName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("postgresqlhsc.ServersClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serverGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ServersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ServersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ServersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServersClient) GetPreparer(ctx context.Context, resourceGroupName string, serverGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverGroupName":   autorest.Encode("path", serverGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-05-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}/servers/{serverName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServersClient) GetResponder(resp *http.Response) (result ServerGroupServer, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServerGroup lists servers of a server group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverGroupName - the name of the server group.
func (client ServersClient) ListByServerGroup(ctx context.Context, resourceGroupName string, serverGroupName string) (result ServerGroupServerListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServersClient.ListByServerGroup")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: serverGroupName,
			Constraints: []validation.Constraint{{Target: "serverGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "serverGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("postgresqlhsc.ServersClient", "ListByServerGroup", err.Error())
	}

	req, err := client.ListByServerGroupPreparer(ctx, resourceGroupName, serverGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ServersClient", "ListByServerGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ServersClient", "ListByServerGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServerGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ServersClient", "ListByServerGroup", resp, "Failure responding to request")
		return
	}

	return
}

// ListByServerGroupPreparer prepares the ListByServerGroup request.
func (client ServersClient) ListByServerGroupPreparer(ctx context.Context, resourceGroupName string, serverGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverGroupName":   autorest.Encode("path", serverGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-05-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}/servers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServerGroupSender sends the ListByServerGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ServersClient) ListByServerGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServerGroupResponder handles the response to the ListByServerGroup request. The method always
// closes the http.Response Body.
func (client ServersClient) ListByServerGroupResponder(resp *http.Response) (result ServerGroupServerListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
