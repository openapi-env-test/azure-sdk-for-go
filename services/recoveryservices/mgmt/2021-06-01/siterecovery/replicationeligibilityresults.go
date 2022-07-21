package siterecovery

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

// ReplicationEligibilityResultsClient is the client for the ReplicationEligibilityResults methods of the Siterecovery
// service.
type ReplicationEligibilityResultsClient struct {
	BaseClient
}

// NewReplicationEligibilityResultsClient creates an instance of the ReplicationEligibilityResultsClient client.
func NewReplicationEligibilityResultsClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationEligibilityResultsClient {
	return NewReplicationEligibilityResultsClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationEligibilityResultsClientWithBaseURI creates an instance of the ReplicationEligibilityResultsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewReplicationEligibilityResultsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationEligibilityResultsClient {
	return ReplicationEligibilityResultsClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Get validates whether a given VM can be protected or not in which case returns list of errors.
// Parameters:
// virtualMachineName - virtual Machine name.
func (client ReplicationEligibilityResultsClient) Get(ctx context.Context, virtualMachineName string) (result ReplicationEligibilityResults, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationEligibilityResultsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, virtualMachineName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationEligibilityResultsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationEligibilityResultsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationEligibilityResultsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationEligibilityResultsClient) GetPreparer(ctx context.Context, virtualMachineName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualMachineName": autorest.Encode("path", virtualMachineName),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}/providers/Microsoft.RecoveryServices/replicationEligibilityResults/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationEligibilityResultsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationEligibilityResultsClient) GetResponder(resp *http.Response) (result ReplicationEligibilityResults, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List validates whether a given VM can be protected or not in which case returns list of errors.
// Parameters:
// virtualMachineName - virtual Machine name.
func (client ReplicationEligibilityResultsClient) List(ctx context.Context, virtualMachineName string) (result ReplicationEligibilityResultsCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationEligibilityResultsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, virtualMachineName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationEligibilityResultsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationEligibilityResultsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationEligibilityResultsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationEligibilityResultsClient) ListPreparer(ctx context.Context, virtualMachineName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualMachineName": autorest.Encode("path", virtualMachineName),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}/providers/Microsoft.RecoveryServices/replicationEligibilityResults", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationEligibilityResultsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationEligibilityResultsClient) ListResponder(resp *http.Response) (result ReplicationEligibilityResultsCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
