package hybridkubernetes

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

// OperationsClient is the azure Connected Cluster Resource Provider API for adopting any Kubernetes Cluster
type OperationsClient struct {
	BaseClient
}

// NewOperationsClient creates an instance of the OperationsClient client.
func NewOperationsClient(subscriptionID string) OperationsClient {
	return NewOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOperationsClientWithBaseURI creates an instance of the OperationsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return OperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get lists all of the available API operations for Connected Cluster resource.
func (client OperationsClient) Get(ctx context.Context) (result OperationListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationsClient.Get")
		defer func() {
			sc := -1
			if result.ol.Response.Response != nil {
				sc = result.ol.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getNextResults
	req, err := client.GetPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridkubernetes.OperationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.ol.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybridkubernetes.OperationsClient", "Get", resp, "Failure sending request")
		return
	}

	result.ol, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridkubernetes.OperationsClient", "Get", resp, "Failure responding to request")
		return
	}
	if result.ol.hasNextLink() && result.ol.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client OperationsClient) GetPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2022-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Kubernetes/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OperationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OperationsClient) GetResponder(resp *http.Response) (result OperationList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getNextResults retrieves the next set of results, if any.
func (client OperationsClient) getNextResults(ctx context.Context, lastResults OperationList) (result OperationList, err error) {
	req, err := lastResults.operationListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hybridkubernetes.OperationsClient", "getNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hybridkubernetes.OperationsClient", "getNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridkubernetes.OperationsClient", "getNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetComplete enumerates all values, automatically crossing page boundaries as required.
func (client OperationsClient) GetComplete(ctx context.Context) (result OperationListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationsClient.Get")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.Get(ctx)
	return
}
