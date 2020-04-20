package automation

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

// WatcherClient is the automation Client
type WatcherClient struct {
	BaseClient
}

// NewWatcherClient creates an instance of the WatcherClient client.
func NewWatcherClient(subscriptionID string) WatcherClient {
	return NewWatcherClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWatcherClientWithBaseURI creates an instance of the WatcherClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWatcherClientWithBaseURI(baseURI string, subscriptionID string) WatcherClient {
	return WatcherClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create the watcher identified by watcher name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// watcherName - the watcher name.
// parameters - the create or update parameters for watcher.
func (client WatcherClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string, parameters Watcher) (result Watcher, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WatcherClient.CreateOrUpdate")
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.WatcherClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, automationAccountName, watcherName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client WatcherClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string, parameters Watcher) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"watcherName":           autorest.Encode("path", watcherName),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/watchers/{watcherName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client WatcherClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client WatcherClient) CreateOrUpdateResponder(resp *http.Response) (result Watcher, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the watcher by name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// watcherName - the watcher name.
func (client WatcherClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WatcherClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.WatcherClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, automationAccountName, watcherName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client WatcherClient) DeletePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"watcherName":           autorest.Encode("path", watcherName),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/watchers/{watcherName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client WatcherClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client WatcherClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve the watcher identified by watcher name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// watcherName - the watcher name.
func (client WatcherClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string) (result Watcher, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WatcherClient.Get")
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.WatcherClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, automationAccountName, watcherName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WatcherClient) GetPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"watcherName":           autorest.Encode("path", watcherName),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/watchers/{watcherName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WatcherClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WatcherClient) GetResponder(resp *http.Response) (result Watcher, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAutomationAccount retrieve a list of watchers.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// filter - the filter to apply on the operation.
func (client WatcherClient) ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result WatcherListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WatcherClient.ListByAutomationAccount")
		defer func() {
			sc := -1
			if result.wlr.Response.Response != nil {
				sc = result.wlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.WatcherClient", "ListByAutomationAccount", err.Error())
	}

	result.fn = client.listByAutomationAccountNextResults
	req, err := client.ListByAutomationAccountPreparer(ctx, resourceGroupName, automationAccountName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "ListByAutomationAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.wlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "ListByAutomationAccount", resp, "Failure sending request")
		return
	}

	result.wlr, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "ListByAutomationAccount", resp, "Failure responding to request")
	}

	return
}

// ListByAutomationAccountPreparer prepares the ListByAutomationAccount request.
func (client WatcherClient) ListByAutomationAccountPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/watchers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAutomationAccountSender sends the ListByAutomationAccount request. The method will close the
// http.Response Body if it receives an error.
func (client WatcherClient) ListByAutomationAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByAutomationAccountResponder handles the response to the ListByAutomationAccount request. The method always
// closes the http.Response Body.
func (client WatcherClient) ListByAutomationAccountResponder(resp *http.Response) (result WatcherListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByAutomationAccountNextResults retrieves the next set of results, if any.
func (client WatcherClient) listByAutomationAccountNextResults(ctx context.Context, lastResults WatcherListResult) (result WatcherListResult, err error) {
	req, err := lastResults.watcherListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.WatcherClient", "listByAutomationAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.WatcherClient", "listByAutomationAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "listByAutomationAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByAutomationAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client WatcherClient) ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result WatcherListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WatcherClient.ListByAutomationAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByAutomationAccount(ctx, resourceGroupName, automationAccountName, filter)
	return
}

// Start resume the watcher identified by watcher name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// watcherName - the watcher name.
func (client WatcherClient) Start(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WatcherClient.Start")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.WatcherClient", "Start", err.Error())
	}

	req, err := client.StartPreparer(ctx, resourceGroupName, automationAccountName, watcherName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "Start", nil, "Failure preparing request")
		return
	}

	resp, err := client.StartSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "Start", resp, "Failure sending request")
		return
	}

	result, err = client.StartResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "Start", resp, "Failure responding to request")
	}

	return
}

// StartPreparer prepares the Start request.
func (client WatcherClient) StartPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"watcherName":           autorest.Encode("path", watcherName),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/watchers/{watcherName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client WatcherClient) StartSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client WatcherClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop resume the watcher identified by watcher name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// watcherName - the watcher name.
func (client WatcherClient) Stop(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WatcherClient.Stop")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.WatcherClient", "Stop", err.Error())
	}

	req, err := client.StopPreparer(ctx, resourceGroupName, automationAccountName, watcherName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "Stop", nil, "Failure preparing request")
		return
	}

	resp, err := client.StopSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "Stop", resp, "Failure sending request")
		return
	}

	result, err = client.StopResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "Stop", resp, "Failure responding to request")
	}

	return
}

// StopPreparer prepares the Stop request.
func (client WatcherClient) StopPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"watcherName":           autorest.Encode("path", watcherName),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/watchers/{watcherName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client WatcherClient) StopSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client WatcherClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update update the watcher identified by watcher name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// watcherName - the watcher name.
// parameters - the update parameters for watcher.
func (client WatcherClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string, parameters WatcherUpdateParameters) (result Watcher, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WatcherClient.Update")
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.WatcherClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, automationAccountName, watcherName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.WatcherClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client WatcherClient) UpdatePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string, parameters WatcherUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"watcherName":           autorest.Encode("path", watcherName),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/watchers/{watcherName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client WatcherClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client WatcherClient) UpdateResponder(resp *http.Response) (result Watcher, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
