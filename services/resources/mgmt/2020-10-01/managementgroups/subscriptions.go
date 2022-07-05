package managementgroups

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

// SubscriptionsClient is the the Azure Management Groups API enables consolidation of multiple
// subscriptions/resources into an organizational hierarchy and centrally
// manage access control, policies, alerting and reporting for those resources.
type SubscriptionsClient struct {
	BaseClient
}

// NewSubscriptionsClient creates an instance of the SubscriptionsClient client.
func NewSubscriptionsClient() SubscriptionsClient {
	return NewSubscriptionsClientWithBaseURI(DefaultBaseURI)
}

// NewSubscriptionsClientWithBaseURI creates an instance of the SubscriptionsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSubscriptionsClientWithBaseURI(baseURI string) SubscriptionsClient {
	return SubscriptionsClient{NewWithBaseURI(baseURI)}
}

// Create associates existing subscription with the management group.
// Parameters:
// groupID - management Group ID.
// subscriptionID - subscription ID.
// cacheControl - indicates whether the request should utilize any caches. Populate the header with 'no-cache'
// value to bypass existing caches.
func (client SubscriptionsClient) Create(ctx context.Context, groupID string, subscriptionID string, cacheControl string) (result SubscriptionUnderManagementGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, groupID, subscriptionID, cacheControl)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client SubscriptionsClient) CreatePreparer(ctx context.Context, groupID string, subscriptionID string, cacheControl string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":        autorest.Encode("path", groupID),
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2020-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}/subscriptions/{subscriptionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String("no-cache")))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) CreateResponder(resp *http.Response) (result SubscriptionUnderManagementGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete de-associates subscription from the management group.
// Parameters:
// groupID - management Group ID.
// subscriptionID - subscription ID.
// cacheControl - indicates whether the request should utilize any caches. Populate the header with 'no-cache'
// value to bypass existing caches.
func (client SubscriptionsClient) Delete(ctx context.Context, groupID string, subscriptionID string, cacheControl string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, groupID, subscriptionID, cacheControl)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SubscriptionsClient) DeletePreparer(ctx context.Context, groupID string, subscriptionID string, cacheControl string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":        autorest.Encode("path", groupID),
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2020-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}/subscriptions/{subscriptionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String("no-cache")))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetSubscription retrieves details about given subscription which is associated with the management group.
// Parameters:
// groupID - management Group ID.
// subscriptionID - subscription ID.
// cacheControl - indicates whether the request should utilize any caches. Populate the header with 'no-cache'
// value to bypass existing caches.
func (client SubscriptionsClient) GetSubscription(ctx context.Context, groupID string, subscriptionID string, cacheControl string) (result SubscriptionUnderManagementGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.GetSubscription")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetSubscriptionPreparer(ctx, groupID, subscriptionID, cacheControl)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "GetSubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "GetSubscription", resp, "Failure sending request")
		return
	}

	result, err = client.GetSubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "GetSubscription", resp, "Failure responding to request")
		return
	}

	return
}

// GetSubscriptionPreparer prepares the GetSubscription request.
func (client SubscriptionsClient) GetSubscriptionPreparer(ctx context.Context, groupID string, subscriptionID string, cacheControl string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":        autorest.Encode("path", groupID),
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2020-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}/subscriptions/{subscriptionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String("no-cache")))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSubscriptionSender sends the GetSubscription request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) GetSubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetSubscriptionResponder handles the response to the GetSubscription request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) GetSubscriptionResponder(resp *http.Response) (result SubscriptionUnderManagementGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSubscriptionsUnderManagementGroup retrieves details about all subscriptions which are associated with the
// management group.
// Parameters:
// groupID - management Group ID.
// skiptoken - page continuation token is only used if a previous operation returned a partial result.
// If a previous response contains a nextLink element, the value of the nextLink element will include a token
// parameter that specifies a starting point to use for subsequent calls.
func (client SubscriptionsClient) GetSubscriptionsUnderManagementGroup(ctx context.Context, groupID string, skiptoken string) (result ListSubscriptionUnderManagementGroupPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.GetSubscriptionsUnderManagementGroup")
		defer func() {
			sc := -1
			if result.lsumg.Response.Response != nil {
				sc = result.lsumg.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getSubscriptionsUnderManagementGroupNextResults
	req, err := client.GetSubscriptionsUnderManagementGroupPreparer(ctx, groupID, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "GetSubscriptionsUnderManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSubscriptionsUnderManagementGroupSender(req)
	if err != nil {
		result.lsumg.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "GetSubscriptionsUnderManagementGroup", resp, "Failure sending request")
		return
	}

	result.lsumg, err = client.GetSubscriptionsUnderManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "GetSubscriptionsUnderManagementGroup", resp, "Failure responding to request")
		return
	}
	if result.lsumg.hasNextLink() && result.lsumg.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetSubscriptionsUnderManagementGroupPreparer prepares the GetSubscriptionsUnderManagementGroup request.
func (client SubscriptionsClient) GetSubscriptionsUnderManagementGroupPreparer(ctx context.Context, groupID string, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId": autorest.Encode("path", groupID),
	}

	const APIVersion = "2020-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}/subscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSubscriptionsUnderManagementGroupSender sends the GetSubscriptionsUnderManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) GetSubscriptionsUnderManagementGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetSubscriptionsUnderManagementGroupResponder handles the response to the GetSubscriptionsUnderManagementGroup request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) GetSubscriptionsUnderManagementGroupResponder(resp *http.Response) (result ListSubscriptionUnderManagementGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getSubscriptionsUnderManagementGroupNextResults retrieves the next set of results, if any.
func (client SubscriptionsClient) getSubscriptionsUnderManagementGroupNextResults(ctx context.Context, lastResults ListSubscriptionUnderManagementGroup) (result ListSubscriptionUnderManagementGroup, err error) {
	req, err := lastResults.listSubscriptionUnderManagementGroupPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "getSubscriptionsUnderManagementGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetSubscriptionsUnderManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "getSubscriptionsUnderManagementGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetSubscriptionsUnderManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "getSubscriptionsUnderManagementGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetSubscriptionsUnderManagementGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client SubscriptionsClient) GetSubscriptionsUnderManagementGroupComplete(ctx context.Context, groupID string, skiptoken string) (result ListSubscriptionUnderManagementGroupIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.GetSubscriptionsUnderManagementGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetSubscriptionsUnderManagementGroup(ctx, groupID, skiptoken)
	return
}
