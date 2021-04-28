package eventhub

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

// RegionsClient is the client for the Regions methods of the Eventhub service.
type RegionsClient struct {
	BaseClient
}

// NewRegionsClient creates an instance of the RegionsClient client.
func NewRegionsClient(subscriptionID string) RegionsClient {
	return NewRegionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRegionsClientWithBaseURI creates an instance of the RegionsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRegionsClientWithBaseURI(baseURI string, subscriptionID string) RegionsClient {
	return RegionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListBySku gets the available Regions for a given sku
// Parameters:
// sku - the sku type.
func (client RegionsClient) ListBySku(ctx context.Context, sku string) (result MessagingRegionsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegionsClient.ListBySku")
		defer func() {
			sc := -1
			if result.mrlr.Response.Response != nil {
				sc = result.mrlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: sku,
			Constraints: []validation.Constraint{{Target: "sku", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "sku", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("eventhub.RegionsClient", "ListBySku", err.Error())
	}

	result.fn = client.listBySkuNextResults
	req, err := client.ListBySkuPreparer(ctx, sku)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.RegionsClient", "ListBySku", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySkuSender(req)
	if err != nil {
		result.mrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventhub.RegionsClient", "ListBySku", resp, "Failure sending request")
		return
	}

	result.mrlr, err = client.ListBySkuResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.RegionsClient", "ListBySku", resp, "Failure responding to request")
		return
	}
	if result.mrlr.hasNextLink() && result.mrlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySkuPreparer prepares the ListBySku request.
func (client RegionsClient) ListBySkuPreparer(ctx context.Context, sku string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"sku":            autorest.Encode("path", sku),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.EventHub/sku/{sku}/regions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySkuSender sends the ListBySku request. The method will close the
// http.Response Body if it receives an error.
func (client RegionsClient) ListBySkuSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySkuResponder handles the response to the ListBySku request. The method always
// closes the http.Response Body.
func (client RegionsClient) ListBySkuResponder(resp *http.Response) (result MessagingRegionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySkuNextResults retrieves the next set of results, if any.
func (client RegionsClient) listBySkuNextResults(ctx context.Context, lastResults MessagingRegionsListResult) (result MessagingRegionsListResult, err error) {
	req, err := lastResults.messagingRegionsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventhub.RegionsClient", "listBySkuNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySkuSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventhub.RegionsClient", "listBySkuNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySkuResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.RegionsClient", "listBySkuNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySkuComplete enumerates all values, automatically crossing page boundaries as required.
func (client RegionsClient) ListBySkuComplete(ctx context.Context, sku string) (result MessagingRegionsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegionsClient.ListBySku")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySku(ctx, sku)
	return
}
