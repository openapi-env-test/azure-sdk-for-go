package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
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

// SecurityGroupsClient is the network Client
type SecurityGroupsClient struct {
	BaseClient
}

// NewSecurityGroupsClient creates an instance of the SecurityGroupsClient client.
func NewSecurityGroupsClient(subscriptionID string) SecurityGroupsClient {
	return NewSecurityGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSecurityGroupsClientWithBaseURI creates an instance of the SecurityGroupsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSecurityGroupsClientWithBaseURI(baseURI string, subscriptionID string) SecurityGroupsClient {
	return SecurityGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a network security group in the specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkSecurityGroupName - the name of the network security group.
// parameters - parameters supplied to the create or update network security group operation.
func (client SecurityGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters SecurityGroup) (result SecurityGroupsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityGroupsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, networkSecurityGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SecurityGroupsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters SecurityGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkSecurityGroupName": autorest.Encode("path", networkSecurityGroupName),
		"resourceGroupName":        autorest.Encode("path", resourceGroupName),
		"subscriptionId":           autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityGroupsClient) CreateOrUpdateSender(req *http.Request) (future SecurityGroupsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SecurityGroupsClient) CreateOrUpdateResponder(resp *http.Response) (result SecurityGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified network security group.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkSecurityGroupName - the name of the network security group.
func (client SecurityGroupsClient) Delete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string) (result SecurityGroupsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityGroupsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, networkSecurityGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SecurityGroupsClient) DeletePreparer(ctx context.Context, resourceGroupName string, networkSecurityGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkSecurityGroupName": autorest.Encode("path", networkSecurityGroupName),
		"resourceGroupName":        autorest.Encode("path", resourceGroupName),
		"subscriptionId":           autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityGroupsClient) DeleteSender(req *http.Request) (future SecurityGroupsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SecurityGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified network security group.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkSecurityGroupName - the name of the network security group.
// expand - expands referenced resources.
func (client SecurityGroupsClient) Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, expand string) (result SecurityGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityGroupsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, networkSecurityGroupName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SecurityGroupsClient) GetPreparer(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkSecurityGroupName": autorest.Encode("path", networkSecurityGroupName),
		"resourceGroupName":        autorest.Encode("path", resourceGroupName),
		"subscriptionId":           autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SecurityGroupsClient) GetResponder(resp *http.Response) (result SecurityGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all network security groups in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client SecurityGroupsClient) List(ctx context.Context, resourceGroupName string) (result SecurityGroupListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityGroupsClient.List")
		defer func() {
			sc := -1
			if result.sglr.Response.Response != nil {
				sc = result.sglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.sglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "List", resp, "Failure sending request")
		return
	}

	result.sglr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.sglr.hasNextLink() && result.sglr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client SecurityGroupsClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityGroupsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SecurityGroupsClient) ListResponder(resp *http.Response) (result SecurityGroupListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client SecurityGroupsClient) listNextResults(ctx context.Context, lastResults SecurityGroupListResult) (result SecurityGroupListResult, err error) {
	req, err := lastResults.securityGroupListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SecurityGroupsClient) ListComplete(ctx context.Context, resourceGroupName string) (result SecurityGroupListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityGroupsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName)
	return
}

// ListAll gets all network security groups in a subscription.
func (client SecurityGroupsClient) ListAll(ctx context.Context) (result SecurityGroupListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityGroupsClient.ListAll")
		defer func() {
			sc := -1
			if result.sglr.Response.Response != nil {
				sc = result.sglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAllNextResults
	req, err := client.ListAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.sglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "ListAll", resp, "Failure sending request")
		return
	}

	result.sglr, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "ListAll", resp, "Failure responding to request")
		return
	}
	if result.sglr.hasNextLink() && result.sglr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client SecurityGroupsClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkSecurityGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityGroupsClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client SecurityGroupsClient) ListAllResponder(resp *http.Response) (result SecurityGroupListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllNextResults retrieves the next set of results, if any.
func (client SecurityGroupsClient) listAllNextResults(ctx context.Context, lastResults SecurityGroupListResult) (result SecurityGroupListResult, err error) {
	req, err := lastResults.securityGroupListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "listAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "listAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "listAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client SecurityGroupsClient) ListAllComplete(ctx context.Context) (result SecurityGroupListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityGroupsClient.ListAll")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAll(ctx)
	return
}

// UpdateTags updates a network security group tags.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkSecurityGroupName - the name of the network security group.
// parameters - parameters supplied to update network security group tags.
func (client SecurityGroupsClient) UpdateTags(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters TagsObject) (result SecurityGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityGroupsClient.UpdateTags")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateTagsPreparer(ctx, resourceGroupName, networkSecurityGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "UpdateTags", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateTagsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "UpdateTags", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateTagsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityGroupsClient", "UpdateTags", resp, "Failure responding to request")
		return
	}

	return
}

// UpdateTagsPreparer prepares the UpdateTags request.
func (client SecurityGroupsClient) UpdateTagsPreparer(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters TagsObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkSecurityGroupName": autorest.Encode("path", networkSecurityGroupName),
		"resourceGroupName":        autorest.Encode("path", resourceGroupName),
		"subscriptionId":           autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateTagsSender sends the UpdateTags request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityGroupsClient) UpdateTagsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateTagsResponder handles the response to the UpdateTags request. The method always
// closes the http.Response Body.
func (client SecurityGroupsClient) UpdateTagsResponder(resp *http.Response) (result SecurityGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
