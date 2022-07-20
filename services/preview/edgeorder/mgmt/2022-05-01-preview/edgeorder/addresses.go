package edgeorder

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

// AddressesClient is the edge Order API's
type AddressesClient struct {
	BaseClient
}

// NewAddressesClient creates an instance of the AddressesClient client.
func NewAddressesClient(subscriptionID string) AddressesClient {
	return NewAddressesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAddressesClientWithBaseURI creates an instance of the AddressesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAddressesClientWithBaseURI(baseURI string, subscriptionID string) AddressesClient {
	return AddressesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a new address with the specified parameters. Existing address cannot be updated with this API and
// should
// instead be updated with the Update address API.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// addressName - the name of the address Resource within the specified resource group. address names must be
// between 3 and 24 characters in length and use any alphanumeric and underscore only.
// addressResource - address details from request body.
func (client AddressesClient) Create(ctx context.Context, resourceGroupName string, addressName string, addressResource AddressResource) (result AddressesCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddressesClient.Create")
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
		{TargetValue: addressName,
			Constraints: []validation.Constraint{{Target: "addressName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "addressName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "addressName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}},
		{TargetValue: addressResource,
			Constraints: []validation.Constraint{{Target: "addressResource.AddressProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "addressResource.AddressProperties.ShippingAddress", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "addressResource.AddressProperties.ShippingAddress.StreetAddress1", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "addressResource.AddressProperties.ShippingAddress.Country", Name: validation.Null, Rule: true, Chain: nil},
					}},
					{Target: "addressResource.AddressProperties.ContactDetails", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "addressResource.AddressProperties.ContactDetails.ContactName", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "addressResource.AddressProperties.ContactDetails.Phone", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "addressResource.AddressProperties.ContactDetails.EmailList", Name: validation.Null, Rule: true, Chain: nil},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("edgeorder.AddressesClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, addressName, addressResource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client AddressesClient) CreatePreparer(ctx context.Context, resourceGroupName string, addressName string, addressResource AddressResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"addressName":       autorest.Encode("path", addressName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	addressResource.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/addresses/{addressName}", pathParameters),
		autorest.WithJSON(addressResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client AddressesClient) CreateSender(req *http.Request) (future AddressesCreateFuture, err error) {
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

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client AddressesClient) CreateResponder(resp *http.Response) (result AddressResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete an address.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// addressName - the name of the address Resource within the specified resource group. address names must be
// between 3 and 24 characters in length and use any alphanumeric and underscore only.
func (client AddressesClient) Delete(ctx context.Context, resourceGroupName string, addressName string) (result AddressesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddressesClient.Delete")
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
		{TargetValue: addressName,
			Constraints: []validation.Constraint{{Target: "addressName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "addressName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "addressName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("edgeorder.AddressesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, addressName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AddressesClient) DeletePreparer(ctx context.Context, resourceGroupName string, addressName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"addressName":       autorest.Encode("path", addressName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/addresses/{addressName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AddressesClient) DeleteSender(req *http.Request) (future AddressesDeleteFuture, err error) {
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
func (client AddressesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get information about the specified address.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// addressName - the name of the address Resource within the specified resource group. address names must be
// between 3 and 24 characters in length and use any alphanumeric and underscore only.
func (client AddressesClient) Get(ctx context.Context, resourceGroupName string, addressName string) (result AddressResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddressesClient.Get")
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
		{TargetValue: addressName,
			Constraints: []validation.Constraint{{Target: "addressName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "addressName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "addressName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("edgeorder.AddressesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, addressName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AddressesClient) GetPreparer(ctx context.Context, resourceGroupName string, addressName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"addressName":       autorest.Encode("path", addressName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/addresses/{addressName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AddressesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AddressesClient) GetResponder(resp *http.Response) (result AddressResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup list all the addresses available under the given resource group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// filter - $filter is supported to filter based on shipping address properties. Filter supports only equals
// operation.
// skipToken - $skipToken is supported on Get list of addresses, which provides the next page in the list of
// addresses.
// top - $top is supported on fetching list of resources. $top=10 means that the first 10 items in the list
// will be returned to the API caller.
func (client AddressesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, skipToken string, top *int32) (result AddressResourceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddressesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.arl.Response.Response != nil {
				sc = result.arl.Response.Response.StatusCode
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
		return result, validation.NewError("edgeorder.AddressesClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, filter, skipToken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.arl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.arl, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.arl.hasNextLink() && result.arl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client AddressesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string, skipToken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/addresses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AddressesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client AddressesClient) ListByResourceGroupResponder(resp *http.Response) (result AddressResourceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client AddressesClient) listByResourceGroupNextResults(ctx context.Context, lastResults AddressResourceList) (result AddressResourceList, err error) {
	req, err := lastResults.addressResourceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client AddressesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, skipToken string, top *int32) (result AddressResourceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddressesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, filter, skipToken, top)
	return
}

// ListBySubscription list all the addresses available under the subscription.
// Parameters:
// filter - $filter is supported to filter based on shipping address properties. Filter supports only equals
// operation.
// skipToken - $skipToken is supported on Get list of addresses, which provides the next page in the list of
// addresses.
// top - $top is supported on fetching list of resources. $top=10 means that the first 10 items in the list
// will be returned to the API caller.
func (client AddressesClient) ListBySubscription(ctx context.Context, filter string, skipToken string, top *int32) (result AddressResourceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddressesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.arl.Response.Response != nil {
				sc = result.arl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("edgeorder.AddressesClient", "ListBySubscription", err.Error())
	}

	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, filter, skipToken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.arl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.arl, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.arl.hasNextLink() && result.arl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client AddressesClient) ListBySubscriptionPreparer(ctx context.Context, filter string, skipToken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.EdgeOrder/addresses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client AddressesClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client AddressesClient) ListBySubscriptionResponder(resp *http.Response) (result AddressResourceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client AddressesClient) listBySubscriptionNextResults(ctx context.Context, lastResults AddressResourceList) (result AddressResourceList, err error) {
	req, err := lastResults.addressResourceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client AddressesClient) ListBySubscriptionComplete(ctx context.Context, filter string, skipToken string, top *int32) (result AddressResourceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddressesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, filter, skipToken, top)
	return
}

// Update update the properties of an existing address.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// addressName - the name of the address Resource within the specified resource group. address names must be
// between 3 and 24 characters in length and use any alphanumeric and underscore only.
// addressUpdateParameter - address update parameters from request body.
// ifMatch - defines the If-Match condition. The patch will be performed only if the ETag of the job on the
// server matches this value.
func (client AddressesClient) Update(ctx context.Context, resourceGroupName string, addressName string, addressUpdateParameter AddressUpdateParameter, ifMatch string) (result AddressesUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddressesClient.Update")
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
		{TargetValue: addressName,
			Constraints: []validation.Constraint{{Target: "addressName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "addressName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "addressName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("edgeorder.AddressesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, addressName, addressUpdateParameter, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "edgeorder.AddressesClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client AddressesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, addressName string, addressUpdateParameter AddressUpdateParameter, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"addressName":       autorest.Encode("path", addressName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/addresses/{addressName}", pathParameters),
		autorest.WithJSON(addressUpdateParameter),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client AddressesClient) UpdateSender(req *http.Request) (future AddressesUpdateFuture, err error) {
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
func (client AddressesClient) UpdateResponder(resp *http.Response) (result AddressResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
