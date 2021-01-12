package contosocdn

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AFDCustomDomainsClient is the cdn Management Client
type AFDCustomDomainsClient struct {
	BaseClient
}

// NewAFDCustomDomainsClient creates an instance of the AFDCustomDomainsClient client.
func NewAFDCustomDomainsClient(subscriptionID string, subscriptionID1 string) AFDCustomDomainsClient {
	return NewAFDCustomDomainsClientWithBaseURI(DefaultBaseURI, subscriptionID, subscriptionID1)
}

// NewAFDCustomDomainsClientWithBaseURI creates an instance of the AFDCustomDomainsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewAFDCustomDomainsClientWithBaseURI(baseURI string, subscriptionID string, subscriptionID1 string) AFDCustomDomainsClient {
	return AFDCustomDomainsClient{NewWithBaseURI(baseURI, subscriptionID, subscriptionID1)}
}

// Create creates a new domain within the specified profile.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the CDN profile which is unique within the resource group.
// customDomainName - name of the domain under the profile which is unique globally
// customDomain - domain properties
func (client AFDCustomDomainsClient) Create(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomain AFDDomain) (result AFDCustomDomainsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AFDCustomDomainsClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: customDomain,
			Constraints: []validation.Constraint{{Target: "customDomain.AFDDomainProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "customDomain.AFDDomainProperties.HostName", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("contosocdn.AFDCustomDomainsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, profileName, customDomainName, customDomain)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client AFDCustomDomainsClient) CreatePreparer(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomain AFDDomain) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customDomainName":  autorest.Encode("path", customDomainName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}", pathParameters),
		autorest.WithJSON(customDomain),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client AFDCustomDomainsClient) CreateSender(req *http.Request) (future AFDCustomDomainsCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client AFDCustomDomainsClient) (ad AFDDomain, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsCreateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("contosocdn.AFDCustomDomainsCreateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if ad.Response.Response, err = future.GetResult(sender); err == nil && ad.Response.Response.StatusCode != http.StatusNoContent {
			ad, err = client.CreateResponder(ad.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsCreateFuture", "Result", ad.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client AFDCustomDomainsClient) CreateResponder(resp *http.Response) (result AFDDomain, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an existing AzureFrontDoor domain with the specified domain name under the specified subscription,
// resource group and profile.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the CDN profile which is unique within the resource group.
// customDomainName - name of the domain under the profile which is unique globally.
func (client AFDCustomDomainsClient) Delete(ctx context.Context, resourceGroupName string, profileName string, customDomainName string) (result AFDCustomDomainsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AFDCustomDomainsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("contosocdn.AFDCustomDomainsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, profileName, customDomainName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AFDCustomDomainsClient) DeletePreparer(ctx context.Context, resourceGroupName string, profileName string, customDomainName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customDomainName":  autorest.Encode("path", customDomainName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AFDCustomDomainsClient) DeleteSender(req *http.Request) (future AFDCustomDomainsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client AFDCustomDomainsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("contosocdn.AFDCustomDomainsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AFDCustomDomainsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an existing AzureFrontDoor domain with the specified domain name under the specified subscription, resource
// group and profile.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the CDN profile which is unique within the resource group.
// customDomainName - name of the domain under the profile which is unique globally.
func (client AFDCustomDomainsClient) Get(ctx context.Context, resourceGroupName string, profileName string, customDomainName string) (result AFDDomain, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AFDCustomDomainsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("contosocdn.AFDCustomDomainsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, profileName, customDomainName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AFDCustomDomainsClient) GetPreparer(ctx context.Context, resourceGroupName string, profileName string, customDomainName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customDomainName":  autorest.Encode("path", customDomainName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AFDCustomDomainsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AFDCustomDomainsClient) GetResponder(resp *http.Response) (result AFDDomain, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByProfile lists existing AzureFrontDoor domains.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the CDN profile which is unique within the resource group.
func (client AFDCustomDomainsClient) ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result AFDDomainListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AFDCustomDomainsClient.ListByProfile")
		defer func() {
			sc := -1
			if result.adlr.Response.Response != nil {
				sc = result.adlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("contosocdn.AFDCustomDomainsClient", "ListByProfile", err.Error())
	}

	result.fn = client.listByProfileNextResults
	req, err := client.ListByProfilePreparer(ctx, resourceGroupName, profileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsClient", "ListByProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByProfileSender(req)
	if err != nil {
		result.adlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsClient", "ListByProfile", resp, "Failure sending request")
		return
	}

	result.adlr, err = client.ListByProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsClient", "ListByProfile", resp, "Failure responding to request")
		return
	}
	if result.adlr.hasNextLink() && result.adlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByProfilePreparer prepares the ListByProfile request.
func (client AFDCustomDomainsClient) ListByProfilePreparer(ctx context.Context, resourceGroupName string, profileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByProfileSender sends the ListByProfile request. The method will close the
// http.Response Body if it receives an error.
func (client AFDCustomDomainsClient) ListByProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByProfileResponder handles the response to the ListByProfile request. The method always
// closes the http.Response Body.
func (client AFDCustomDomainsClient) ListByProfileResponder(resp *http.Response) (result AFDDomainListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByProfileNextResults retrieves the next set of results, if any.
func (client AFDCustomDomainsClient) listByProfileNextResults(ctx context.Context, lastResults AFDDomainListResult) (result AFDDomainListResult, err error) {
	req, err := lastResults.aFDDomainListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsClient", "listByProfileNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsClient", "listByProfileNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsClient", "listByProfileNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByProfileComplete enumerates all values, automatically crossing page boundaries as required.
func (client AFDCustomDomainsClient) ListByProfileComplete(ctx context.Context, resourceGroupName string, profileName string) (result AFDDomainListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AFDCustomDomainsClient.ListByProfile")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByProfile(ctx, resourceGroupName, profileName)
	return
}

// RefreshValidationToken updates the domain validation token.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the CDN profile which is unique within the resource group.
// customDomainName - name of the domain under the profile which is unique globally.
func (client AFDCustomDomainsClient) RefreshValidationToken(ctx context.Context, resourceGroupName string, profileName string, customDomainName string) (result AFDCustomDomainsRefreshValidationTokenFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AFDCustomDomainsClient.RefreshValidationToken")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("contosocdn.AFDCustomDomainsClient", "RefreshValidationToken", err.Error())
	}

	req, err := client.RefreshValidationTokenPreparer(ctx, resourceGroupName, profileName, customDomainName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsClient", "RefreshValidationToken", nil, "Failure preparing request")
		return
	}

	result, err = client.RefreshValidationTokenSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsClient", "RefreshValidationToken", nil, "Failure sending request")
		return
	}

	return
}

// RefreshValidationTokenPreparer prepares the RefreshValidationToken request.
func (client AFDCustomDomainsClient) RefreshValidationTokenPreparer(ctx context.Context, resourceGroupName string, profileName string, customDomainName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customDomainName":  autorest.Encode("path", customDomainName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}/refreshValidationToken", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RefreshValidationTokenSender sends the RefreshValidationToken request. The method will close the
// http.Response Body if it receives an error.
func (client AFDCustomDomainsClient) RefreshValidationTokenSender(req *http.Request) (future AFDCustomDomainsRefreshValidationTokenFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client AFDCustomDomainsClient) (vt ValidationToken, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsRefreshValidationTokenFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("contosocdn.AFDCustomDomainsRefreshValidationTokenFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if vt.Response.Response, err = future.GetResult(sender); err == nil && vt.Response.Response.StatusCode != http.StatusNoContent {
			vt, err = client.RefreshValidationTokenResponder(vt.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsRefreshValidationTokenFuture", "Result", vt.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// RefreshValidationTokenResponder handles the response to the RefreshValidationToken request. The method always
// closes the http.Response Body.
func (client AFDCustomDomainsClient) RefreshValidationTokenResponder(resp *http.Response) (result ValidationToken, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates an existing domain within a profile.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the CDN profile which is unique within the resource group.
// customDomainName - name of the domain under the profile which is unique globally
// customDomainUpdateProperties - domain properties
func (client AFDCustomDomainsClient) Update(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomainUpdateProperties AFDDomainUpdateParameters) (result AFDCustomDomainsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AFDCustomDomainsClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("contosocdn.AFDCustomDomainsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, profileName, customDomainName, customDomainUpdateProperties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsClient", "Update", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client AFDCustomDomainsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomainUpdateProperties AFDDomainUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customDomainName":  autorest.Encode("path", customDomainName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}", pathParameters),
		autorest.WithJSON(customDomainUpdateProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client AFDCustomDomainsClient) UpdateSender(req *http.Request) (future AFDCustomDomainsUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client AFDCustomDomainsClient) (ad AFDDomain, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("contosocdn.AFDCustomDomainsUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if ad.Response.Response, err = future.GetResult(sender); err == nil && ad.Response.Response.StatusCode != http.StatusNoContent {
			ad, err = client.UpdateResponder(ad.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "contosocdn.AFDCustomDomainsUpdateFuture", "Result", ad.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client AFDCustomDomainsClient) UpdateResponder(resp *http.Response) (result AFDDomain, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
