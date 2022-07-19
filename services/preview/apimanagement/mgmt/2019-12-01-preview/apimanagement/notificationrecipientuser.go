package apimanagement

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

// NotificationRecipientUserClient is the apiManagement Client
type NotificationRecipientUserClient struct {
	BaseClient
}

// NewNotificationRecipientUserClient creates an instance of the NotificationRecipientUserClient client.
func NewNotificationRecipientUserClient(subscriptionID string) NotificationRecipientUserClient {
	return NewNotificationRecipientUserClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewNotificationRecipientUserClientWithBaseURI creates an instance of the NotificationRecipientUserClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewNotificationRecipientUserClientWithBaseURI(baseURI string, subscriptionID string) NotificationRecipientUserClient {
	return NotificationRecipientUserClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckEntityExists determine if the Notification Recipient User is subscribed to the notification.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// notificationName - notification Name Identifier.
// userID - user identifier. Must be unique in the current API Management service instance.
func (client NotificationRecipientUserClient) CheckEntityExists(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NotificationRecipientUserClient.CheckEntityExists")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: userID,
			Constraints: []validation.Constraint{{Target: "userID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "userID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.NotificationRecipientUserClient", "CheckEntityExists", err.Error())
	}

	req, err := client.CheckEntityExistsPreparer(ctx, resourceGroupName, serviceName, notificationName, userID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.NotificationRecipientUserClient", "CheckEntityExists", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckEntityExistsSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.NotificationRecipientUserClient", "CheckEntityExists", resp, "Failure sending request")
		return
	}

	result, err = client.CheckEntityExistsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.NotificationRecipientUserClient", "CheckEntityExists", resp, "Failure responding to request")
		return
	}

	return
}

// CheckEntityExistsPreparer prepares the CheckEntityExists request.
func (client NotificationRecipientUserClient) CheckEntityExistsPreparer(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"notificationName":  autorest.Encode("path", notificationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userId":            autorest.Encode("path", userID),
	}

	const APIVersion = "2019-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientUsers/{userId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckEntityExistsSender sends the CheckEntityExists request. The method will close the
// http.Response Body if it receives an error.
func (client NotificationRecipientUserClient) CheckEntityExistsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckEntityExistsResponder handles the response to the CheckEntityExists request. The method always
// closes the http.Response Body.
func (client NotificationRecipientUserClient) CheckEntityExistsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate adds the API Management User to the list of Recipients for the Notification.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// notificationName - notification Name Identifier.
// userID - user identifier. Must be unique in the current API Management service instance.
func (client NotificationRecipientUserClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string) (result RecipientUserContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NotificationRecipientUserClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: userID,
			Constraints: []validation.Constraint{{Target: "userID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "userID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.NotificationRecipientUserClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, notificationName, userID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.NotificationRecipientUserClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.NotificationRecipientUserClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.NotificationRecipientUserClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client NotificationRecipientUserClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"notificationName":  autorest.Encode("path", notificationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userId":            autorest.Encode("path", userID),
	}

	const APIVersion = "2019-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientUsers/{userId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client NotificationRecipientUserClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client NotificationRecipientUserClient) CreateOrUpdateResponder(resp *http.Response) (result RecipientUserContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete removes the API Management user from the list of Notification.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// notificationName - notification Name Identifier.
// userID - user identifier. Must be unique in the current API Management service instance.
func (client NotificationRecipientUserClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NotificationRecipientUserClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: userID,
			Constraints: []validation.Constraint{{Target: "userID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "userID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.NotificationRecipientUserClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, notificationName, userID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.NotificationRecipientUserClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.NotificationRecipientUserClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.NotificationRecipientUserClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client NotificationRecipientUserClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"notificationName":  autorest.Encode("path", notificationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userId":            autorest.Encode("path", userID),
	}

	const APIVersion = "2019-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientUsers/{userId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client NotificationRecipientUserClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client NotificationRecipientUserClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByNotification gets the list of the Notification Recipient User subscribed to the notification.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// notificationName - notification Name Identifier.
func (client NotificationRecipientUserClient) ListByNotification(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName) (result RecipientUserCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NotificationRecipientUserClient.ListByNotification")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.NotificationRecipientUserClient", "ListByNotification", err.Error())
	}

	req, err := client.ListByNotificationPreparer(ctx, resourceGroupName, serviceName, notificationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.NotificationRecipientUserClient", "ListByNotification", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByNotificationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.NotificationRecipientUserClient", "ListByNotification", resp, "Failure sending request")
		return
	}

	result, err = client.ListByNotificationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.NotificationRecipientUserClient", "ListByNotification", resp, "Failure responding to request")
		return
	}

	return
}

// ListByNotificationPreparer prepares the ListByNotification request.
func (client NotificationRecipientUserClient) ListByNotificationPreparer(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"notificationName":  autorest.Encode("path", notificationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientUsers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByNotificationSender sends the ListByNotification request. The method will close the
// http.Response Body if it receives an error.
func (client NotificationRecipientUserClient) ListByNotificationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByNotificationResponder handles the response to the ListByNotification request. The method always
// closes the http.Response Body.
func (client NotificationRecipientUserClient) ListByNotificationResponder(resp *http.Response) (result RecipientUserCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
