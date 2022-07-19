package authorization

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

// AccessReviewDefaultSettingsClient is the client for the AccessReviewDefaultSettings methods of the Authorization
// service.
type AccessReviewDefaultSettingsClient struct {
	BaseClient
}

// NewAccessReviewDefaultSettingsClient creates an instance of the AccessReviewDefaultSettingsClient client.
func NewAccessReviewDefaultSettingsClient(subscriptionID string) AccessReviewDefaultSettingsClient {
	return NewAccessReviewDefaultSettingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAccessReviewDefaultSettingsClientWithBaseURI creates an instance of the AccessReviewDefaultSettingsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewAccessReviewDefaultSettingsClientWithBaseURI(baseURI string, subscriptionID string) AccessReviewDefaultSettingsClient {
	return AccessReviewDefaultSettingsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get access review default settings for the subscription
func (client AccessReviewDefaultSettingsClient) Get(ctx context.Context) (result AccessReviewDefaultSettings, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccessReviewDefaultSettingsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("authorization.AccessReviewDefaultSettingsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.AccessReviewDefaultSettingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.AccessReviewDefaultSettingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.AccessReviewDefaultSettingsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AccessReviewDefaultSettingsClient) GetPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleSettings/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AccessReviewDefaultSettingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AccessReviewDefaultSettingsClient) GetResponder(resp *http.Response) (result AccessReviewDefaultSettings, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Put get access review default settings for the subscription
// Parameters:
// properties - access review schedule settings.
func (client AccessReviewDefaultSettingsClient) Put(ctx context.Context, properties AccessReviewScheduleSettings) (result AccessReviewDefaultSettings, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccessReviewDefaultSettingsClient.Put")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("authorization.AccessReviewDefaultSettingsClient", "Put", err.Error())
	}

	req, err := client.PutPreparer(ctx, properties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.AccessReviewDefaultSettingsClient", "Put", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.AccessReviewDefaultSettingsClient", "Put", resp, "Failure sending request")
		return
	}

	result, err = client.PutResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.AccessReviewDefaultSettingsClient", "Put", resp, "Failure responding to request")
		return
	}

	return
}

// PutPreparer prepares the Put request.
func (client AccessReviewDefaultSettingsClient) PutPreparer(ctx context.Context, properties AccessReviewScheduleSettings) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleSettings/default", pathParameters),
		autorest.WithJSON(properties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutSender sends the Put request. The method will close the
// http.Response Body if it receives an error.
func (client AccessReviewDefaultSettingsClient) PutSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PutResponder handles the response to the Put request. The method always
// closes the http.Response Body.
func (client AccessReviewDefaultSettingsClient) PutResponder(resp *http.Response) (result AccessReviewDefaultSettings, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
