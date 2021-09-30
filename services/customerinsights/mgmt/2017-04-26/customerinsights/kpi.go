package customerinsights

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

// KpiClient is the the Azure Customer Insights management API provides a RESTful set of web services that interact
// with Azure Customer Insights service to manage your resources. The API has entities that capture the relationship
// between an end user and the Azure Customer Insights service.
type KpiClient struct {
	BaseClient
}

// NewKpiClient creates an instance of the KpiClient client.
func NewKpiClient(subscriptionID string) KpiClient {
	return NewKpiClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewKpiClientWithBaseURI creates an instance of the KpiClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewKpiClientWithBaseURI(baseURI string, subscriptionID string) KpiClient {
	return KpiClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a KPI or updates an existing KPI in the hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// kpiName - the name of the KPI.
// parameters - parameters supplied to the create/update KPI operation.
func (client KpiClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, kpiName string, parameters KpiResourceFormat) (result KpiCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KpiClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: kpiName,
			Constraints: []validation.Constraint{{Target: "kpiName", Name: validation.MaxLength, Rule: 512, Chain: nil},
				{Target: "kpiName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "kpiName", Name: validation.Pattern, Rule: `^[a-zA-Z][a-zA-Z0-9_]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.KpiDefinition", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.KpiDefinition.EntityTypeName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.KpiDefinition.Expression", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.KpiDefinition.ThresHolds", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.KpiDefinition.ThresHolds.LowerLimit", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.KpiDefinition.ThresHolds.UpperLimit", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.KpiDefinition.ThresHolds.IncreasingKpi", Name: validation.Null, Rule: true, Chain: nil},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("customerinsights.KpiClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, hubName, kpiName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.KpiClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.KpiClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client KpiClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, hubName string, kpiName string, parameters KpiResourceFormat) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"kpiName":           autorest.Encode("path", kpiName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/kpi/{kpiName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client KpiClient) CreateOrUpdateSender(req *http.Request) (future KpiCreateOrUpdateFuture, err error) {
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
func (client KpiClient) CreateOrUpdateResponder(resp *http.Response) (result KpiResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a KPI in the hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// kpiName - the name of the KPI.
func (client KpiClient) Delete(ctx context.Context, resourceGroupName string, hubName string, kpiName string) (result KpiDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KpiClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, hubName, kpiName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.KpiClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.KpiClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client KpiClient) DeletePreparer(ctx context.Context, resourceGroupName string, hubName string, kpiName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"kpiName":           autorest.Encode("path", kpiName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/kpi/{kpiName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client KpiClient) DeleteSender(req *http.Request) (future KpiDeleteFuture, err error) {
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
func (client KpiClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a KPI in the hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// kpiName - the name of the KPI.
func (client KpiClient) Get(ctx context.Context, resourceGroupName string, hubName string, kpiName string) (result KpiResourceFormat, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KpiClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, hubName, kpiName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.KpiClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.KpiClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.KpiClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client KpiClient) GetPreparer(ctx context.Context, resourceGroupName string, hubName string, kpiName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"kpiName":           autorest.Encode("path", kpiName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/kpi/{kpiName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client KpiClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client KpiClient) GetResponder(resp *http.Response) (result KpiResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByHub gets all the KPIs in the specified hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
func (client KpiClient) ListByHub(ctx context.Context, resourceGroupName string, hubName string) (result KpiListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KpiClient.ListByHub")
		defer func() {
			sc := -1
			if result.klr.Response.Response != nil {
				sc = result.klr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByHubNextResults
	req, err := client.ListByHubPreparer(ctx, resourceGroupName, hubName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.KpiClient", "ListByHub", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.klr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.KpiClient", "ListByHub", resp, "Failure sending request")
		return
	}

	result.klr, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.KpiClient", "ListByHub", resp, "Failure responding to request")
		return
	}
	if result.klr.hasNextLink() && result.klr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByHubPreparer prepares the ListByHub request.
func (client KpiClient) ListByHubPreparer(ctx context.Context, resourceGroupName string, hubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/kpi", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByHubSender sends the ListByHub request. The method will close the
// http.Response Body if it receives an error.
func (client KpiClient) ListByHubSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByHubResponder handles the response to the ListByHub request. The method always
// closes the http.Response Body.
func (client KpiClient) ListByHubResponder(resp *http.Response) (result KpiListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByHubNextResults retrieves the next set of results, if any.
func (client KpiClient) listByHubNextResults(ctx context.Context, lastResults KpiListResult) (result KpiListResult, err error) {
	req, err := lastResults.kpiListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "customerinsights.KpiClient", "listByHubNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "customerinsights.KpiClient", "listByHubNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.KpiClient", "listByHubNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByHubComplete enumerates all values, automatically crossing page boundaries as required.
func (client KpiClient) ListByHubComplete(ctx context.Context, resourceGroupName string, hubName string) (result KpiListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KpiClient.ListByHub")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByHub(ctx, resourceGroupName, hubName)
	return
}

// Reprocess reprocesses the Kpi values of the specified KPI.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// kpiName - the name of the KPI.
func (client KpiClient) Reprocess(ctx context.Context, resourceGroupName string, hubName string, kpiName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KpiClient.Reprocess")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ReprocessPreparer(ctx, resourceGroupName, hubName, kpiName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.KpiClient", "Reprocess", nil, "Failure preparing request")
		return
	}

	resp, err := client.ReprocessSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "customerinsights.KpiClient", "Reprocess", resp, "Failure sending request")
		return
	}

	result, err = client.ReprocessResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.KpiClient", "Reprocess", resp, "Failure responding to request")
		return
	}

	return
}

// ReprocessPreparer prepares the Reprocess request.
func (client KpiClient) ReprocessPreparer(ctx context.Context, resourceGroupName string, hubName string, kpiName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"kpiName":           autorest.Encode("path", kpiName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/kpi/{kpiName}/reprocess", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ReprocessSender sends the Reprocess request. The method will close the
// http.Response Body if it receives an error.
func (client KpiClient) ReprocessSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ReprocessResponder handles the response to the Reprocess request. The method always
// closes the http.Response Body.
func (client KpiClient) ReprocessResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
