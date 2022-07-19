// Package peering implements the Azure ARM Peering service API version 2022-06-01.
//
// Peering Client
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

const (
	// DefaultBaseURI is the default URI used for the service Peering
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Peering.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// CheckServiceProviderAvailability checks if the peering service provider is present within 1000 miles of customer's
// location
// Parameters:
// checkServiceProviderAvailabilityInput - the CheckServiceProviderAvailabilityInput indicating customer
// location and service provider.
func (client BaseClient) CheckServiceProviderAvailability(ctx context.Context, checkServiceProviderAvailabilityInput CheckServiceProviderAvailabilityInput) (result String, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.CheckServiceProviderAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CheckServiceProviderAvailabilityPreparer(ctx, checkServiceProviderAvailabilityInput)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.BaseClient", "CheckServiceProviderAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckServiceProviderAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.BaseClient", "CheckServiceProviderAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckServiceProviderAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.BaseClient", "CheckServiceProviderAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckServiceProviderAvailabilityPreparer prepares the CheckServiceProviderAvailability request.
func (client BaseClient) CheckServiceProviderAvailabilityPreparer(ctx context.Context, checkServiceProviderAvailabilityInput CheckServiceProviderAvailabilityInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/checkServiceProviderAvailability", pathParameters),
		autorest.WithJSON(checkServiceProviderAvailabilityInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckServiceProviderAvailabilitySender sends the CheckServiceProviderAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) CheckServiceProviderAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckServiceProviderAvailabilityResponder handles the response to the CheckServiceProviderAvailability request. The method always
// closes the http.Response Body.
func (client BaseClient) CheckServiceProviderAvailabilityResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
