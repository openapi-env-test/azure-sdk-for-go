package servicebus

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

// PremiumMessagingRegionsClient is the client for the PremiumMessagingRegions methods of the Servicebus service.
type PremiumMessagingRegionsClient struct {
	BaseClient
}

// NewPremiumMessagingRegionsClient creates an instance of the PremiumMessagingRegionsClient client.
func NewPremiumMessagingRegionsClient(subscriptionID string) PremiumMessagingRegionsClient {
	return NewPremiumMessagingRegionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPremiumMessagingRegionsClientWithBaseURI creates an instance of the PremiumMessagingRegionsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewPremiumMessagingRegionsClientWithBaseURI(baseURI string, subscriptionID string) PremiumMessagingRegionsClient {
	return PremiumMessagingRegionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets the available premium messaging regions for servicebus
func (client PremiumMessagingRegionsClient) List(ctx context.Context) (result PremiumMessagingRegionsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PremiumMessagingRegionsClient.List")
		defer func() {
			sc := -1
			if result.pmrlr.Response.Response != nil {
				sc = result.pmrlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.PremiumMessagingRegionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pmrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicebus.PremiumMessagingRegionsClient", "List", resp, "Failure sending request")
		return
	}

	result.pmrlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.PremiumMessagingRegionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.pmrlr.hasNextLink() && result.pmrlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client PremiumMessagingRegionsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceBus/premiumMessagingRegions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PremiumMessagingRegionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PremiumMessagingRegionsClient) ListResponder(resp *http.Response) (result PremiumMessagingRegionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client PremiumMessagingRegionsClient) listNextResults(ctx context.Context, lastResults PremiumMessagingRegionsListResult) (result PremiumMessagingRegionsListResult, err error) {
	req, err := lastResults.premiumMessagingRegionsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicebus.PremiumMessagingRegionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicebus.PremiumMessagingRegionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.PremiumMessagingRegionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client PremiumMessagingRegionsClient) ListComplete(ctx context.Context) (result PremiumMessagingRegionsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PremiumMessagingRegionsClient.List")
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
