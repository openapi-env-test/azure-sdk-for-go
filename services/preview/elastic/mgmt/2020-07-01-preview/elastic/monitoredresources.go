package elastic

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

// MonitoredResourcesClient is the client for the MonitoredResources methods of the Elastic service.
type MonitoredResourcesClient struct {
	BaseClient
}

// NewMonitoredResourcesClient creates an instance of the MonitoredResourcesClient client.
func NewMonitoredResourcesClient(subscriptionID string) MonitoredResourcesClient {
	return NewMonitoredResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMonitoredResourcesClientWithBaseURI creates an instance of the MonitoredResourcesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewMonitoredResourcesClientWithBaseURI(baseURI string, subscriptionID string) MonitoredResourcesClient {
	return MonitoredResourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List sends the list request.
// Parameters:
// resourceGroupName - the name of the resource group to which the Elastic resource belongs.
// monitorName - monitor resource name
func (client MonitoredResourcesClient) List(ctx context.Context, resourceGroupName string, monitorName string) (result MonitoredResourceListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitoredResourcesClient.List")
		defer func() {
			sc := -1
			if result.mrlr.Response.Response != nil {
				sc = result.mrlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, monitorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.MonitoredResourcesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.mrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "elastic.MonitoredResourcesClient", "List", resp, "Failure sending request")
		return
	}

	result.mrlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.MonitoredResourcesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.mrlr.hasNextLink() && result.mrlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client MonitoredResourcesClient) ListPreparer(ctx context.Context, resourceGroupName string, monitorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorName":       autorest.Encode("path", monitorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/listMonitoredResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MonitoredResourcesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MonitoredResourcesClient) ListResponder(resp *http.Response) (result MonitoredResourceListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client MonitoredResourcesClient) listNextResults(ctx context.Context, lastResults MonitoredResourceListResponse) (result MonitoredResourceListResponse, err error) {
	req, err := lastResults.monitoredResourceListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "elastic.MonitoredResourcesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "elastic.MonitoredResourcesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.MonitoredResourcesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client MonitoredResourcesClient) ListComplete(ctx context.Context, resourceGroupName string, monitorName string) (result MonitoredResourceListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitoredResourcesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, monitorName)
	return
}
