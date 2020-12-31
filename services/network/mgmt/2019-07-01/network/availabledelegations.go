package network

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

// AvailableDelegationsClient is the network Client
type AvailableDelegationsClient struct {
	BaseClient
}

// NewAvailableDelegationsClient creates an instance of the AvailableDelegationsClient client.
func NewAvailableDelegationsClient(subscriptionID string) AvailableDelegationsClient {
	return NewAvailableDelegationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAvailableDelegationsClientWithBaseURI creates an instance of the AvailableDelegationsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewAvailableDelegationsClientWithBaseURI(baseURI string, subscriptionID string) AvailableDelegationsClient {
	return AvailableDelegationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets all of the available subnet delegations for this subscription in this region.
// Parameters:
// location - the location of the subnet.
func (client AvailableDelegationsClient) List(ctx context.Context, location string) (result AvailableDelegationsResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailableDelegationsClient.List")
		defer func() {
			sc := -1
			if result.adr.Response.Response != nil {
				sc = result.adr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AvailableDelegationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.adr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.AvailableDelegationsClient", "List", resp, "Failure sending request")
		return
	}

	result.adr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AvailableDelegationsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.adr.hasNextLink() && result.adr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client AvailableDelegationsClient) ListPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availableDelegations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AvailableDelegationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AvailableDelegationsClient) ListResponder(resp *http.Response) (result AvailableDelegationsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AvailableDelegationsClient) listNextResults(ctx context.Context, lastResults AvailableDelegationsResult) (result AvailableDelegationsResult, err error) {
	req, err := lastResults.availableDelegationsResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.AvailableDelegationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.AvailableDelegationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AvailableDelegationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AvailableDelegationsClient) ListComplete(ctx context.Context, location string) (result AvailableDelegationsResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailableDelegationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, location)
	return
}
