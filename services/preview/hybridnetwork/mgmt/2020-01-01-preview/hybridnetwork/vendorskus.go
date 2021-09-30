package hybridnetwork

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

// VendorSkusClient is the client for the VendorSkus methods of the Hybridnetwork service.
type VendorSkusClient struct {
	BaseClient
}

// NewVendorSkusClient creates an instance of the VendorSkusClient client.
func NewVendorSkusClient(subscriptionID string) VendorSkusClient {
	return NewVendorSkusClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVendorSkusClientWithBaseURI creates an instance of the VendorSkusClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewVendorSkusClientWithBaseURI(baseURI string, subscriptionID string) VendorSkusClient {
	return VendorSkusClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a sku
// Parameters:
// vendorName - the name of the vendor.
// skuName - the name of the sku.
// parameters - parameters supplied to the create or update sku operation.
func (client VendorSkusClient) CreateOrUpdate(ctx context.Context, vendorName string, skuName string, parameters VendorSku) (result VendorSkusCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VendorSkusClient.CreateOrUpdate")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybridnetwork.VendorSkusClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, vendorName, skuName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorSkusClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorSkusClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VendorSkusClient) CreateOrUpdatePreparer(ctx context.Context, vendorName string, skuName string, parameters VendorSku) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"skuName":        autorest.Encode("path", skuName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"vendorName":     autorest.Encode("path", vendorName),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}/vendorSkus/{skuName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VendorSkusClient) CreateOrUpdateSender(req *http.Request) (future VendorSkusCreateOrUpdateFuture, err error) {
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
func (client VendorSkusClient) CreateOrUpdateResponder(resp *http.Response) (result VendorSku, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified sku.
// Parameters:
// vendorName - the name of the vendor.
// skuName - the name of the sku.
func (client VendorSkusClient) Delete(ctx context.Context, vendorName string, skuName string) (result VendorSkusDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VendorSkusClient.Delete")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybridnetwork.VendorSkusClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, vendorName, skuName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorSkusClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorSkusClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VendorSkusClient) DeletePreparer(ctx context.Context, vendorName string, skuName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"skuName":        autorest.Encode("path", skuName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"vendorName":     autorest.Encode("path", vendorName),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}/vendorSkus/{skuName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VendorSkusClient) DeleteSender(req *http.Request) (future VendorSkusDeleteFuture, err error) {
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
func (client VendorSkusClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified sku.
// Parameters:
// vendorName - the name of the vendor.
// skuName - the name of the sku.
func (client VendorSkusClient) Get(ctx context.Context, vendorName string, skuName string) (result VendorSku, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VendorSkusClient.Get")
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
		return result, validation.NewError("hybridnetwork.VendorSkusClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, vendorName, skuName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorSkusClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorSkusClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorSkusClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VendorSkusClient) GetPreparer(ctx context.Context, vendorName string, skuName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"skuName":        autorest.Encode("path", skuName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"vendorName":     autorest.Encode("path", vendorName),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}/vendorSkus/{skuName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VendorSkusClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VendorSkusClient) GetResponder(resp *http.Response) (result VendorSku, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the skus of a vendor.
// Parameters:
// vendorName - the name of the vendor.
func (client VendorSkusClient) List(ctx context.Context, vendorName string) (result VendorSkuListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VendorSkusClient.List")
		defer func() {
			sc := -1
			if result.vslr.Response.Response != nil {
				sc = result.vslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybridnetwork.VendorSkusClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, vendorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorSkusClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.vslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorSkusClient", "List", resp, "Failure sending request")
		return
	}

	result.vslr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorSkusClient", "List", resp, "Failure responding to request")
		return
	}
	if result.vslr.hasNextLink() && result.vslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client VendorSkusClient) ListPreparer(ctx context.Context, vendorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"vendorName":     autorest.Encode("path", vendorName),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}/vendorSkus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VendorSkusClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VendorSkusClient) ListResponder(resp *http.Response) (result VendorSkuListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VendorSkusClient) listNextResults(ctx context.Context, lastResults VendorSkuListResult) (result VendorSkuListResult, err error) {
	req, err := lastResults.vendorSkuListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hybridnetwork.VendorSkusClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hybridnetwork.VendorSkusClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorSkusClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VendorSkusClient) ListComplete(ctx context.Context, vendorName string) (result VendorSkuListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VendorSkusClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, vendorName)
	return
}
