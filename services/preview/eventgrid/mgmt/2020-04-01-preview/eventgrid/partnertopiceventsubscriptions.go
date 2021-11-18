package eventgrid

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

// PartnerTopicEventSubscriptionsClient is the azure EventGrid Management Client
type PartnerTopicEventSubscriptionsClient struct {
	BaseClient
}

// NewPartnerTopicEventSubscriptionsClient creates an instance of the PartnerTopicEventSubscriptionsClient client.
func NewPartnerTopicEventSubscriptionsClient(subscriptionID string) PartnerTopicEventSubscriptionsClient {
	return NewPartnerTopicEventSubscriptionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPartnerTopicEventSubscriptionsClientWithBaseURI creates an instance of the PartnerTopicEventSubscriptionsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewPartnerTopicEventSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) PartnerTopicEventSubscriptionsClient {
	return PartnerTopicEventSubscriptionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate asynchronously creates or updates an event subscription of a partner topic with the specified
// parameters. Existing event subscriptions will be updated with this API.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerTopicName - name of the partner topic.
// eventSubscriptionName - name of the event subscription to be created. Event subscription names must be
// between 3 and 100 characters in length and use alphanumeric letters only.
// eventSubscriptionInfo - event subscription properties containing the destination and filter information.
func (client PartnerTopicEventSubscriptionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription) (result PartnerTopicEventSubscriptionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerTopicEventSubscriptionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, eventSubscriptionInfo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PartnerTopicEventSubscriptionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventSubscriptionName": autorest.Encode("path", eventSubscriptionName),
		"partnerTopicName":      autorest.Encode("path", partnerTopicName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	eventSubscriptionInfo.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions/{eventSubscriptionName}", pathParameters),
		autorest.WithJSON(eventSubscriptionInfo),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerTopicEventSubscriptionsClient) CreateOrUpdateSender(req *http.Request) (future PartnerTopicEventSubscriptionsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client PartnerTopicEventSubscriptionsClient) CreateOrUpdateResponder(resp *http.Response) (result EventSubscription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete an event subscription of a partner topic.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerTopicName - name of the partner topic.
// eventSubscriptionName - name of the event subscription to be created. Event subscription names must be
// between 3 and 100 characters in length and use alphanumeric letters only.
func (client PartnerTopicEventSubscriptionsClient) Delete(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string) (result PartnerTopicEventSubscriptionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerTopicEventSubscriptionsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PartnerTopicEventSubscriptionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventSubscriptionName": autorest.Encode("path", eventSubscriptionName),
		"partnerTopicName":      autorest.Encode("path", partnerTopicName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions/{eventSubscriptionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerTopicEventSubscriptionsClient) DeleteSender(req *http.Request) (future PartnerTopicEventSubscriptionsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PartnerTopicEventSubscriptionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get an event subscription of a partner topic.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerTopicName - name of the partner topic.
// eventSubscriptionName - name of the event subscription to be found. Event subscription names must be between
// 3 and 100 characters in length and use alphanumeric letters only.
func (client PartnerTopicEventSubscriptionsClient) Get(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string) (result EventSubscription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerTopicEventSubscriptionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PartnerTopicEventSubscriptionsClient) GetPreparer(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventSubscriptionName": autorest.Encode("path", eventSubscriptionName),
		"partnerTopicName":      autorest.Encode("path", partnerTopicName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions/{eventSubscriptionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerTopicEventSubscriptionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PartnerTopicEventSubscriptionsClient) GetResponder(resp *http.Response) (result EventSubscription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetFullURL get the full endpoint URL for an event subscription of a partner topic.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerTopicName - name of the partner topic.
// eventSubscriptionName - name of the event subscription to be created. Event subscription names must be
// between 3 and 100 characters in length and use alphanumeric letters only.
func (client PartnerTopicEventSubscriptionsClient) GetFullURL(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string) (result EventSubscriptionFullURL, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerTopicEventSubscriptionsClient.GetFullURL")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetFullURLPreparer(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "GetFullURL", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetFullURLSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "GetFullURL", resp, "Failure sending request")
		return
	}

	result, err = client.GetFullURLResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "GetFullURL", resp, "Failure responding to request")
		return
	}

	return
}

// GetFullURLPreparer prepares the GetFullURL request.
func (client PartnerTopicEventSubscriptionsClient) GetFullURLPreparer(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventSubscriptionName": autorest.Encode("path", eventSubscriptionName),
		"partnerTopicName":      autorest.Encode("path", partnerTopicName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions/{eventSubscriptionName}/getFullUrl", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetFullURLSender sends the GetFullURL request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerTopicEventSubscriptionsClient) GetFullURLSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetFullURLResponder handles the response to the GetFullURL request. The method always
// closes the http.Response Body.
func (client PartnerTopicEventSubscriptionsClient) GetFullURLResponder(resp *http.Response) (result EventSubscriptionFullURL, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByPartnerTopic list event subscriptions that belong to a specific partner topic.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerTopicName - name of the partner topic.
// filter - the query used to filter the search results using OData syntax. Filtering is permitted on the
// 'name' property only and with limited number of OData operations. These operations are: the 'contains'
// function as well as the following logical operations: not, and, or, eq (for equal), and ne (for not equal).
// No arithmetic operations are supported. The following is a valid filter example: $filter=contains(namE,
// 'PATTERN') and name ne 'PATTERN-1'. The following is not a valid filter example: $filter=location eq
// 'westus'.
// top - the number of results to return per page for the list operation. Valid range for top parameter is 1 to
// 100. If not specified, the default number of results to be returned is 20 items per page.
func (client PartnerTopicEventSubscriptionsClient) ListByPartnerTopic(ctx context.Context, resourceGroupName string, partnerTopicName string, filter string, top *int32) (result EventSubscriptionsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerTopicEventSubscriptionsClient.ListByPartnerTopic")
		defer func() {
			sc := -1
			if result.eslr.Response.Response != nil {
				sc = result.eslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByPartnerTopicNextResults
	req, err := client.ListByPartnerTopicPreparer(ctx, resourceGroupName, partnerTopicName, filter, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "ListByPartnerTopic", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByPartnerTopicSender(req)
	if err != nil {
		result.eslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "ListByPartnerTopic", resp, "Failure sending request")
		return
	}

	result.eslr, err = client.ListByPartnerTopicResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "ListByPartnerTopic", resp, "Failure responding to request")
		return
	}
	if result.eslr.hasNextLink() && result.eslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByPartnerTopicPreparer prepares the ListByPartnerTopic request.
func (client PartnerTopicEventSubscriptionsClient) ListByPartnerTopicPreparer(ctx context.Context, resourceGroupName string, partnerTopicName string, filter string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"partnerTopicName":  autorest.Encode("path", partnerTopicName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByPartnerTopicSender sends the ListByPartnerTopic request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerTopicEventSubscriptionsClient) ListByPartnerTopicSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByPartnerTopicResponder handles the response to the ListByPartnerTopic request. The method always
// closes the http.Response Body.
func (client PartnerTopicEventSubscriptionsClient) ListByPartnerTopicResponder(resp *http.Response) (result EventSubscriptionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByPartnerTopicNextResults retrieves the next set of results, if any.
func (client PartnerTopicEventSubscriptionsClient) listByPartnerTopicNextResults(ctx context.Context, lastResults EventSubscriptionsListResult) (result EventSubscriptionsListResult, err error) {
	req, err := lastResults.eventSubscriptionsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "listByPartnerTopicNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByPartnerTopicSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "listByPartnerTopicNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByPartnerTopicResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "listByPartnerTopicNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByPartnerTopicComplete enumerates all values, automatically crossing page boundaries as required.
func (client PartnerTopicEventSubscriptionsClient) ListByPartnerTopicComplete(ctx context.Context, resourceGroupName string, partnerTopicName string, filter string, top *int32) (result EventSubscriptionsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerTopicEventSubscriptionsClient.ListByPartnerTopic")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByPartnerTopic(ctx, resourceGroupName, partnerTopicName, filter, top)
	return
}

// Update update event subscription of a partner topic.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerTopicName - name of the partner topic.
// eventSubscriptionName - name of the event subscription to be created. Event subscription names must be
// between 3 and 100 characters in length and use alphanumeric letters only.
// eventSubscriptionUpdateParameters - updated event subscription information.
func (client PartnerTopicEventSubscriptionsClient) Update(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters) (result PartnerTopicEventSubscriptionsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PartnerTopicEventSubscriptionsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, eventSubscriptionUpdateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PartnerTopicEventSubscriptionsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client PartnerTopicEventSubscriptionsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventSubscriptionName": autorest.Encode("path", eventSubscriptionName),
		"partnerTopicName":      autorest.Encode("path", partnerTopicName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions/{eventSubscriptionName}", pathParameters),
		autorest.WithJSON(eventSubscriptionUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client PartnerTopicEventSubscriptionsClient) UpdateSender(req *http.Request) (future PartnerTopicEventSubscriptionsUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client PartnerTopicEventSubscriptionsClient) UpdateResponder(resp *http.Response) (result EventSubscription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
