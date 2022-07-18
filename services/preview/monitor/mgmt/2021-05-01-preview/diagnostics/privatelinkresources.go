package diagnostics

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

// PrivateLinkResourcesClient is the monitor Management Client
type PrivateLinkResourcesClient struct {
	BaseClient
}

// NewPrivateLinkResourcesClient creates an instance of the PrivateLinkResourcesClient client.
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return NewPrivateLinkResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateLinkResourcesClientWithBaseURI creates an instance of the PrivateLinkResourcesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return PrivateLinkResourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the private link resources that need to be created for a Azure Monitor PrivateLinkScope.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// scopeName - the name of the Azure Monitor PrivateLinkScope resource.
// groupName - the name of the private link resource.
func (client PrivateLinkResourcesClient) Get(ctx context.Context, resourceGroupName string, scopeName string, groupName string) (result PrivateLinkResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkResourcesClient.Get")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("diagnostics.PrivateLinkResourcesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, scopeName, groupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.PrivateLinkResourcesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "diagnostics.PrivateLinkResourcesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.PrivateLinkResourcesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrivateLinkResourcesClient) GetPreparer(ctx context.Context, resourceGroupName string, scopeName string, groupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupName":         autorest.Encode("path", groupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"scopeName":         autorest.Encode("path", scopeName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/privateLinkScopes/{scopeName}/privateLinkResources/{groupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkResourcesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrivateLinkResourcesClient) GetResponder(resp *http.Response) (result PrivateLinkResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByPrivateLinkScope gets the private link resources that need to be created for a Azure Monitor PrivateLinkScope.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// scopeName - the name of the Azure Monitor PrivateLinkScope resource.
func (client PrivateLinkResourcesClient) ListByPrivateLinkScope(ctx context.Context, resourceGroupName string, scopeName string) (result PrivateLinkResourceListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkResourcesClient.ListByPrivateLinkScope")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("diagnostics.PrivateLinkResourcesClient", "ListByPrivateLinkScope", err.Error())
	}

	req, err := client.ListByPrivateLinkScopePreparer(ctx, resourceGroupName, scopeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.PrivateLinkResourcesClient", "ListByPrivateLinkScope", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByPrivateLinkScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "diagnostics.PrivateLinkResourcesClient", "ListByPrivateLinkScope", resp, "Failure sending request")
		return
	}

	result, err = client.ListByPrivateLinkScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.PrivateLinkResourcesClient", "ListByPrivateLinkScope", resp, "Failure responding to request")
		return
	}

	return
}

// ListByPrivateLinkScopePreparer prepares the ListByPrivateLinkScope request.
func (client PrivateLinkResourcesClient) ListByPrivateLinkScopePreparer(ctx context.Context, resourceGroupName string, scopeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"scopeName":         autorest.Encode("path", scopeName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/privateLinkScopes/{scopeName}/privateLinkResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByPrivateLinkScopeSender sends the ListByPrivateLinkScope request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkResourcesClient) ListByPrivateLinkScopeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByPrivateLinkScopeResponder handles the response to the ListByPrivateLinkScope request. The method always
// closes the http.Response Body.
func (client PrivateLinkResourcesClient) ListByPrivateLinkScopeResponder(resp *http.Response) (result PrivateLinkResourceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
