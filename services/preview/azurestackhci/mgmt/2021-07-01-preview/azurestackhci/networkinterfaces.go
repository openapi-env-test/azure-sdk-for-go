package azurestackhci

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

// NetworkinterfacesClient is the azure Stack HCI management service
type NetworkinterfacesClient struct {
	BaseClient
}

// NewNetworkinterfacesClient creates an instance of the NetworkinterfacesClient client.
func NewNetworkinterfacesClient(subscriptionID string) NetworkinterfacesClient {
	return NewNetworkinterfacesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewNetworkinterfacesClientWithBaseURI creates an instance of the NetworkinterfacesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewNetworkinterfacesClientWithBaseURI(baseURI string, subscriptionID string) NetworkinterfacesClient {
	return NetworkinterfacesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate sends the create or update request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
func (client NetworkinterfacesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkinterfacesName string, networkinterfaces Networkinterfaces) (result NetworkinterfacesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkinterfacesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: networkinterfacesName,
			Constraints: []validation.Constraint{{Target: "networkinterfacesName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "networkinterfacesName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("azurestackhci.NetworkinterfacesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, networkinterfacesName, networkinterfaces)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client NetworkinterfacesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, networkinterfacesName string, networkinterfaces Networkinterfaces) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkinterfacesName": autorest.Encode("path", networkinterfacesName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/networkinterfaces/{networkinterfacesName}", pathParameters),
		autorest.WithJSON(networkinterfaces),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkinterfacesClient) CreateOrUpdateSender(req *http.Request) (future NetworkinterfacesCreateOrUpdateFuture, err error) {
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
func (client NetworkinterfacesClient) CreateOrUpdateResponder(resp *http.Response) (result Networkinterfaces, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete sends the delete request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
func (client NetworkinterfacesClient) Delete(ctx context.Context, resourceGroupName string, networkinterfacesName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkinterfacesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: networkinterfacesName,
			Constraints: []validation.Constraint{{Target: "networkinterfacesName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "networkinterfacesName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("azurestackhci.NetworkinterfacesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, networkinterfacesName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client NetworkinterfacesClient) DeletePreparer(ctx context.Context, resourceGroupName string, networkinterfacesName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkinterfacesName": autorest.Encode("path", networkinterfacesName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/networkinterfaces/{networkinterfacesName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkinterfacesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client NetworkinterfacesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByResourceGroup sends the list by resource group request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
func (client NetworkinterfacesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result NetworkinterfacesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkinterfacesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.nlr.Response.Response != nil {
				sc = result.nlr.Response.Response.StatusCode
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
		return result, validation.NewError("azurestackhci.NetworkinterfacesClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.nlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.nlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.nlr.hasNextLink() && result.nlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client NetworkinterfacesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/networkinterfaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkinterfacesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client NetworkinterfacesClient) ListByResourceGroupResponder(resp *http.Response) (result NetworkinterfacesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client NetworkinterfacesClient) listByResourceGroupNextResults(ctx context.Context, lastResults NetworkinterfacesListResult) (result NetworkinterfacesListResult, err error) {
	req, err := lastResults.networkinterfacesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client NetworkinterfacesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result NetworkinterfacesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkinterfacesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// ListBySubscription sends the list by subscription request.
func (client NetworkinterfacesClient) ListBySubscription(ctx context.Context) (result NetworkinterfacesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkinterfacesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.nlr.Response.Response != nil {
				sc = result.nlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("azurestackhci.NetworkinterfacesClient", "ListBySubscription", err.Error())
	}

	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.nlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.nlr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.nlr.hasNextLink() && result.nlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client NetworkinterfacesClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AzureStackHCI/networkinterfaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkinterfacesClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client NetworkinterfacesClient) ListBySubscriptionResponder(resp *http.Response) (result NetworkinterfacesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client NetworkinterfacesClient) listBySubscriptionNextResults(ctx context.Context, lastResults NetworkinterfacesListResult) (result NetworkinterfacesListResult, err error) {
	req, err := lastResults.networkinterfacesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client NetworkinterfacesClient) ListBySubscriptionComplete(ctx context.Context) (result NetworkinterfacesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkinterfacesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx)
	return
}

// Retrieve sends the retrieve request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
func (client NetworkinterfacesClient) Retrieve(ctx context.Context, resourceGroupName string, networkinterfacesName string) (result Networkinterfaces, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkinterfacesClient.Retrieve")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: networkinterfacesName,
			Constraints: []validation.Constraint{{Target: "networkinterfacesName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "networkinterfacesName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("azurestackhci.NetworkinterfacesClient", "Retrieve", err.Error())
	}

	req, err := client.RetrievePreparer(ctx, resourceGroupName, networkinterfacesName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "Retrieve", nil, "Failure preparing request")
		return
	}

	resp, err := client.RetrieveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "Retrieve", resp, "Failure sending request")
		return
	}

	result, err = client.RetrieveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "Retrieve", resp, "Failure responding to request")
		return
	}

	return
}

// RetrievePreparer prepares the Retrieve request.
func (client NetworkinterfacesClient) RetrievePreparer(ctx context.Context, resourceGroupName string, networkinterfacesName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkinterfacesName": autorest.Encode("path", networkinterfacesName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/networkinterfaces/{networkinterfacesName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RetrieveSender sends the Retrieve request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkinterfacesClient) RetrieveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RetrieveResponder handles the response to the Retrieve request. The method always
// closes the http.Response Body.
func (client NetworkinterfacesClient) RetrieveResponder(resp *http.Response) (result Networkinterfaces, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update sends the update request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
func (client NetworkinterfacesClient) Update(ctx context.Context, resourceGroupName string, networkinterfacesName string, networkinterfaces NetworkinterfacesPatch) (result Networkinterfaces, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkinterfacesClient.Update")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: networkinterfacesName,
			Constraints: []validation.Constraint{{Target: "networkinterfacesName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "networkinterfacesName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("azurestackhci.NetworkinterfacesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, networkinterfacesName, networkinterfaces)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestackhci.NetworkinterfacesClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client NetworkinterfacesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, networkinterfacesName string, networkinterfaces NetworkinterfacesPatch) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkinterfacesName": autorest.Encode("path", networkinterfacesName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/networkinterfaces/{networkinterfacesName}", pathParameters),
		autorest.WithJSON(networkinterfaces),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkinterfacesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client NetworkinterfacesClient) UpdateResponder(resp *http.Response) (result Networkinterfaces, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
