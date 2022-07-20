package billing

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

// SubscriptionsAliasesClient is the billing Client
type SubscriptionsAliasesClient struct {
	BaseClient
}

// NewSubscriptionsAliasesClient creates an instance of the SubscriptionsAliasesClient client.
func NewSubscriptionsAliasesClient() SubscriptionsAliasesClient {
	return NewSubscriptionsAliasesClientWithBaseURI(DefaultBaseURI)
}

// NewSubscriptionsAliasesClientWithBaseURI creates an instance of the SubscriptionsAliasesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewSubscriptionsAliasesClientWithBaseURI(baseURI string) SubscriptionsAliasesClient {
	return SubscriptionsAliasesClient{NewWithBaseURI(baseURI)}
}

// CreateOrUpdate creates or updates a billing subscription by its alias ID.  The operation is supported for seat based
// billing subscriptions.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// aliasName - the ID that uniquely identifies a subscription alias.
// parameters - new or updated billing subscription alias.
func (client SubscriptionsAliasesClient) CreateOrUpdate(ctx context.Context, billingAccountName string, aliasName string, parameters SubscriptionAlias) (result SubscriptionsAliasesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsAliasesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, billingAccountName, aliasName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsAliasesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsAliasesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SubscriptionsAliasesClient) CreateOrUpdatePreparer(ctx context.Context, billingAccountName string, aliasName string, parameters SubscriptionAlias) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"aliasName":          autorest.Encode("path", aliasName),
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2021-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingSubscriptionAliases/{aliasName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsAliasesClient) CreateOrUpdateSender(req *http.Request) (future SubscriptionsAliasesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
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
func (client SubscriptionsAliasesClient) CreateOrUpdateResponder(resp *http.Response) (result SubscriptionAlias, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a subscription by its alias ID.  The operation is supported for seat based billing subscriptions.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// aliasName - the ID that uniquely identifies a subscription alias.
func (client SubscriptionsAliasesClient) Get(ctx context.Context, billingAccountName string, aliasName string) (result SubscriptionAlias, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsAliasesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, billingAccountName, aliasName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsAliasesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsAliasesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsAliasesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SubscriptionsAliasesClient) GetPreparer(ctx context.Context, billingAccountName string, aliasName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"aliasName":          autorest.Encode("path", aliasName),
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2021-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingSubscriptionAliases/{aliasName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsAliasesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SubscriptionsAliasesClient) GetResponder(resp *http.Response) (result SubscriptionAlias, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBillingAccount lists the subscription aliases for a billing account. The operation is supported for seat based
// billing subscriptions.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
func (client SubscriptionsAliasesClient) ListByBillingAccount(ctx context.Context, billingAccountName string) (result SubscriptionAliasListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsAliasesClient.ListByBillingAccount")
		defer func() {
			sc := -1
			if result.salr.Response.Response != nil {
				sc = result.salr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingAccountNextResults
	req, err := client.ListByBillingAccountPreparer(ctx, billingAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsAliasesClient", "ListByBillingAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.salr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsAliasesClient", "ListByBillingAccount", resp, "Failure sending request")
		return
	}

	result.salr, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsAliasesClient", "ListByBillingAccount", resp, "Failure responding to request")
		return
	}
	if result.salr.hasNextLink() && result.salr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByBillingAccountPreparer prepares the ListByBillingAccount request.
func (client SubscriptionsAliasesClient) ListByBillingAccountPreparer(ctx context.Context, billingAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2021-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingSubscriptionAliases", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingAccountSender sends the ListByBillingAccount request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsAliasesClient) ListByBillingAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingAccountResponder handles the response to the ListByBillingAccount request. The method always
// closes the http.Response Body.
func (client SubscriptionsAliasesClient) ListByBillingAccountResponder(resp *http.Response) (result SubscriptionAliasListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingAccountNextResults retrieves the next set of results, if any.
func (client SubscriptionsAliasesClient) listByBillingAccountNextResults(ctx context.Context, lastResults SubscriptionAliasListResult) (result SubscriptionAliasListResult, err error) {
	req, err := lastResults.subscriptionAliasListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.SubscriptionsAliasesClient", "listByBillingAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.SubscriptionsAliasesClient", "listByBillingAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.SubscriptionsAliasesClient", "listByBillingAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client SubscriptionsAliasesClient) ListByBillingAccountComplete(ctx context.Context, billingAccountName string) (result SubscriptionAliasListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsAliasesClient.ListByBillingAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingAccount(ctx, billingAccountName)
	return
}
