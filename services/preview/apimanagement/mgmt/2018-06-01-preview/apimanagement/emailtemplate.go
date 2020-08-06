package apimanagement

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

// EmailTemplateClient is the apiManagement Client
type EmailTemplateClient struct {
	BaseClient
}

// NewEmailTemplateClient creates an instance of the EmailTemplateClient client.
func NewEmailTemplateClient(subscriptionID string) EmailTemplateClient {
	return NewEmailTemplateClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEmailTemplateClientWithBaseURI creates an instance of the EmailTemplateClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewEmailTemplateClientWithBaseURI(baseURI string, subscriptionID string) EmailTemplateClient {
	return EmailTemplateClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate updates an Email Template.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// templateName - email Template Name Identifier.
// parameters - email Template update parameters.
// ifMatch - eTag of the Entity. Not required when creating an entity, but required when updating an entity.
func (client EmailTemplateClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, parameters EmailTemplateUpdateParameters, ifMatch string) (result EmailTemplateContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EmailTemplateClient.CreateOrUpdate")
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
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.EmailTemplateUpdateParameterProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.EmailTemplateUpdateParameterProperties.Subject", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.EmailTemplateUpdateParameterProperties.Subject", Name: validation.MaxLength, Rule: 1000, Chain: nil},
						{Target: "parameters.EmailTemplateUpdateParameterProperties.Subject", Name: validation.MinLength, Rule: 1, Chain: nil},
					}},
					{Target: "parameters.EmailTemplateUpdateParameterProperties.Body", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.EmailTemplateUpdateParameterProperties.Body", Name: validation.MinLength, Rule: 1, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("apimanagement.EmailTemplateClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, templateName, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client EmailTemplateClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, parameters EmailTemplateUpdateParameters, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"templateName":      autorest.Encode("path", templateName),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client EmailTemplateClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client EmailTemplateClient) CreateOrUpdateResponder(resp *http.Response) (result EmailTemplateContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete reset the Email Template to default template provided by the API Management service instance.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// templateName - email Template Name Identifier.
// ifMatch - eTag of the Entity. ETag should match the current entity state from the header response of the GET
// request or it should be * for unconditional update.
func (client EmailTemplateClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, ifMatch string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EmailTemplateClient.Delete")
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
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.EmailTemplateClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, templateName, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client EmailTemplateClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"templateName":      autorest.Encode("path", templateName),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client EmailTemplateClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client EmailTemplateClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of the email template specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// templateName - email Template Name Identifier.
func (client EmailTemplateClient) Get(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName) (result EmailTemplateContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EmailTemplateClient.Get")
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
		return result, validation.NewError("apimanagement.EmailTemplateClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, templateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client EmailTemplateClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"templateName":      autorest.Encode("path", templateName),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client EmailTemplateClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client EmailTemplateClient) GetResponder(resp *http.Response) (result EmailTemplateContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetEntityTag gets the entity state (Etag) version of the email template specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// templateName - email Template Name Identifier.
func (client EmailTemplateClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EmailTemplateClient.GetEntityTag")
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
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.EmailTemplateClient", "GetEntityTag", err.Error())
	}

	req, err := client.GetEntityTagPreparer(ctx, resourceGroupName, serviceName, templateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "GetEntityTag", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEntityTagSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "GetEntityTag", resp, "Failure sending request")
		return
	}

	result, err = client.GetEntityTagResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "GetEntityTag", resp, "Failure responding to request")
	}

	return
}

// GetEntityTagPreparer prepares the GetEntityTag request.
func (client EmailTemplateClient) GetEntityTagPreparer(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"templateName":      autorest.Encode("path", templateName),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEntityTagSender sends the GetEntityTag request. The method will close the
// http.Response Body if it receives an error.
func (client EmailTemplateClient) GetEntityTagSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetEntityTagResponder handles the response to the GetEntityTag request. The method always
// closes the http.Response Body.
func (client EmailTemplateClient) GetEntityTagResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByService lists a collection of properties defined within a service instance.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// filter - | Field       | Supported operators    | Supported functions               |
// |-------------|------------------------|-----------------------------------|
//
// |name | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith|
// top - number of records to return.
// skip - number of records to skip.
func (client EmailTemplateClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result EmailTemplateCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EmailTemplateClient.ListByService")
		defer func() {
			sc := -1
			if result.etc.Response.Response != nil {
				sc = result.etc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.EmailTemplateClient", "ListByService", err.Error())
	}

	result.fn = client.listByServiceNextResults
	req, err := client.ListByServicePreparer(ctx, resourceGroupName, serviceName, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.etc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "ListByService", resp, "Failure sending request")
		return
	}

	result.etc, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "ListByService", resp, "Failure responding to request")
	}
	if result.etc.hasNextLink() && result.etc.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client EmailTemplateClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client EmailTemplateClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client EmailTemplateClient) ListByServiceResponder(resp *http.Response) (result EmailTemplateCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServiceNextResults retrieves the next set of results, if any.
func (client EmailTemplateClient) listByServiceNextResults(ctx context.Context, lastResults EmailTemplateCollection) (result EmailTemplateCollection, err error) {
	req, err := lastResults.emailTemplateCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "listByServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "listByServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "listByServiceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client EmailTemplateClient) ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result EmailTemplateCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EmailTemplateClient.ListByService")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByService(ctx, resourceGroupName, serviceName, filter, top, skip)
	return
}

// Update updates the specific Email Template.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// templateName - email Template Name Identifier.
// parameters - update parameters.
func (client EmailTemplateClient) Update(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, parameters EmailTemplateUpdateParameters) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EmailTemplateClient.Update")
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
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.EmailTemplateClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, serviceName, templateName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.EmailTemplateClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client EmailTemplateClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, parameters EmailTemplateUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"templateName":      autorest.Encode("path", templateName),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client EmailTemplateClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client EmailTemplateClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
