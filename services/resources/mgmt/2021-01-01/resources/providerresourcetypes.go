package resources

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

// ProviderResourceTypesClient is the provides operations for working with resources and resource groups.
type ProviderResourceTypesClient struct {
	BaseClient
}

// NewProviderResourceTypesClient creates an instance of the ProviderResourceTypesClient client.
func NewProviderResourceTypesClient(subscriptionID string) ProviderResourceTypesClient {
	return NewProviderResourceTypesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProviderResourceTypesClientWithBaseURI creates an instance of the ProviderResourceTypesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewProviderResourceTypesClientWithBaseURI(baseURI string, subscriptionID string) ProviderResourceTypesClient {
	return ProviderResourceTypesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List list the resource types for a specified resource provider.
// Parameters:
// resourceProviderNamespace - the namespace of the resource provider.
// expand - the $expand query parameter. For example, to include property aliases in response, use
// $expand=resourceTypes/aliases.
func (client ProviderResourceTypesClient) List(ctx context.Context, resourceProviderNamespace string, expand string) (result ProviderResourceTypeListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderResourceTypesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceProviderNamespace, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProviderResourceTypesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.ProviderResourceTypesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProviderResourceTypesClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ProviderResourceTypesClient) ListPreparer(ctx context.Context, resourceProviderNamespace string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}/resourceTypes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderResourceTypesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProviderResourceTypesClient) ListResponder(resp *http.Response) (result ProviderResourceTypeListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
