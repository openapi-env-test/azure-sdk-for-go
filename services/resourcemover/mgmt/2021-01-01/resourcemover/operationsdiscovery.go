package resourcemover

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

// OperationsDiscoveryClient is the a first party Azure service orchestrating the move of Azure resources from one
// Azure region to another.
type OperationsDiscoveryClient struct {
	BaseClient
}

// NewOperationsDiscoveryClient creates an instance of the OperationsDiscoveryClient client.
func NewOperationsDiscoveryClient(subscriptionID string) OperationsDiscoveryClient {
	return NewOperationsDiscoveryClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOperationsDiscoveryClientWithBaseURI creates an instance of the OperationsDiscoveryClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewOperationsDiscoveryClientWithBaseURI(baseURI string, subscriptionID string) OperationsDiscoveryClient {
	return OperationsDiscoveryClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get sends the get request.
func (client OperationsDiscoveryClient) Get(ctx context.Context) (result OperationsDiscoveryCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationsDiscoveryClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcemover.OperationsDiscoveryClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcemover.OperationsDiscoveryClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcemover.OperationsDiscoveryClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client OperationsDiscoveryClient) GetPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2021-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Migrate/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OperationsDiscoveryClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OperationsDiscoveryClient) GetResponder(resp *http.Response) (result OperationsDiscoveryCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
