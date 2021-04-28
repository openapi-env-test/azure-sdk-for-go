package peering

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

// ServiceProvidersClient is the peering Client
type ServiceProvidersClient struct {
	BaseClient
}

// NewServiceProvidersClient creates an instance of the ServiceProvidersClient client.
func NewServiceProvidersClient(subscriptionID string) ServiceProvidersClient {
	return NewServiceProvidersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServiceProvidersClientWithBaseURI creates an instance of the ServiceProvidersClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewServiceProvidersClientWithBaseURI(baseURI string, subscriptionID string) ServiceProvidersClient {
	return ServiceProvidersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists all of the available peering service locations for the specified kind of peering.
func (client ServiceProvidersClient) List(ctx context.Context) (result ServiceProviderListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceProvidersClient.List")
		defer func() {
			sc := -1
			if result.splr.Response.Response != nil {
				sc = result.splr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.ServiceProvidersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.splr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.ServiceProvidersClient", "List", resp, "Failure sending request")
		return
	}

	result.splr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.ServiceProvidersClient", "List", resp, "Failure responding to request")
		return
	}
	if result.splr.hasNextLink() && result.splr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ServiceProvidersClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringServiceProviders", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceProvidersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ServiceProvidersClient) ListResponder(resp *http.Response) (result ServiceProviderListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ServiceProvidersClient) listNextResults(ctx context.Context, lastResults ServiceProviderListResult) (result ServiceProviderListResult, err error) {
	req, err := lastResults.serviceProviderListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "peering.ServiceProvidersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "peering.ServiceProvidersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.ServiceProvidersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ServiceProvidersClient) ListComplete(ctx context.Context) (result ServiceProviderListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceProvidersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}
