package network

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

// ActiveSecurityAdminRulesClient is the network Client
type ActiveSecurityAdminRulesClient struct {
	BaseClient
}

// NewActiveSecurityAdminRulesClient creates an instance of the ActiveSecurityAdminRulesClient client.
func NewActiveSecurityAdminRulesClient(subscriptionID string) ActiveSecurityAdminRulesClient {
	return NewActiveSecurityAdminRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewActiveSecurityAdminRulesClientWithBaseURI creates an instance of the ActiveSecurityAdminRulesClient client using
// a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewActiveSecurityAdminRulesClientWithBaseURI(baseURI string, subscriptionID string) ActiveSecurityAdminRulesClient {
	return ActiveSecurityAdminRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists active security admin rules in a network manager.
// Parameters:
// parameters - active Configuration Parameter.
// resourceGroupName - the name of the resource group.
// networkManagerName - the name of the network manager.
func (client ActiveSecurityAdminRulesClient) List(ctx context.Context, parameters ActiveConfigurationParameter, resourceGroupName string, networkManagerName string) (result ActiveSecurityAdminRulesListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActiveSecurityAdminRulesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, parameters, resourceGroupName, networkManagerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ActiveSecurityAdminRulesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ActiveSecurityAdminRulesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ActiveSecurityAdminRulesClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ActiveSecurityAdminRulesClient) ListPreparer(ctx context.Context, parameters ActiveConfigurationParameter, resourceGroupName string, networkManagerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkManagerName": autorest.Encode("path", networkManagerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/listActiveSecurityAdminRules", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ActiveSecurityAdminRulesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ActiveSecurityAdminRulesClient) ListResponder(resp *http.Response) (result ActiveSecurityAdminRulesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
