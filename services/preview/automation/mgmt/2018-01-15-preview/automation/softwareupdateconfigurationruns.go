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
	"github.com/satori/go.uuid"
	"net/http"
)

// SoftwareUpdateConfigurationRunsClient is the automation Client
type SoftwareUpdateConfigurationRunsClient struct {
	BaseClient
}

// NewSoftwareUpdateConfigurationRunsClient creates an instance of the SoftwareUpdateConfigurationRunsClient client.
func NewSoftwareUpdateConfigurationRunsClient(subscriptionID string) SoftwareUpdateConfigurationRunsClient {
	return NewSoftwareUpdateConfigurationRunsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSoftwareUpdateConfigurationRunsClientWithBaseURI creates an instance of the SoftwareUpdateConfigurationRunsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewSoftwareUpdateConfigurationRunsClientWithBaseURI(baseURI string, subscriptionID string) SoftwareUpdateConfigurationRunsClient {
	return SoftwareUpdateConfigurationRunsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetByID get a single software update configuration Run by Id.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// softwareUpdateConfigurationRunID - the Id of the software update configuration run.
// clientRequestID - identifies this specific client request.
func (client SoftwareUpdateConfigurationRunsClient) GetByID(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationRunID uuid.UUID, clientRequestID string) (result SoftwareUpdateConfigurationRun, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SoftwareUpdateConfigurationRunsClient.GetByID")
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
		return result, validation.NewError("automation.SoftwareUpdateConfigurationRunsClient", "GetByID", err.Error())
	}

	req, err := client.GetByIDPreparer(ctx, resourceGroupName, automationAccountName, softwareUpdateConfigurationRunID, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationRunsClient", "GetByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationRunsClient", "GetByID", resp, "Failure sending request")
		return
	}

	result, err = client.GetByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationRunsClient", "GetByID", resp, "Failure responding to request")
	}

	return
}

// GetByIDPreparer prepares the GetByID request.
func (client SoftwareUpdateConfigurationRunsClient) GetByIDPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationRunID uuid.UUID, clientRequestID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName":            autorest.Encode("path", automationAccountName),
		"resourceGroupName":                autorest.Encode("path", resourceGroupName),
		"softwareUpdateConfigurationRunId": autorest.Encode("path", softwareUpdateConfigurationRunID),
		"subscriptionId":                   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurationRuns/{softwareUpdateConfigurationRunId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(clientRequestID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("clientRequestId", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByIDSender sends the GetByID request. The method will close the
// http.Response Body if it receives an error.
func (client SoftwareUpdateConfigurationRunsClient) GetByIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetByIDResponder handles the response to the GetByID request. The method always
// closes the http.Response Body.
func (client SoftwareUpdateConfigurationRunsClient) GetByIDResponder(resp *http.Response) (result SoftwareUpdateConfigurationRun, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List return list of software update configuration runs
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// clientRequestID - identifies this specific client request.
// filter - the filter to apply on the operation. You can use the following filters: 'properties/osType',
// 'properties/status', 'properties/startTime', and 'properties/softwareUpdateConfiguration/name'
// skip - number of entries you skip before returning results
// top - maximum number of entries returned in the results collection
func (client SoftwareUpdateConfigurationRunsClient) List(ctx context.Context, resourceGroupName string, automationAccountName string, clientRequestID string, filter string, skip string, top string) (result SoftwareUpdateConfigurationRunListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SoftwareUpdateConfigurationRunsClient.List")
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
		return result, validation.NewError("automation.SoftwareUpdateConfigurationRunsClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, resourceGroupName, automationAccountName, clientRequestID, filter, skip, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationRunsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationRunsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SoftwareUpdateConfigurationRunsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SoftwareUpdateConfigurationRunsClient) ListPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, clientRequestID string, filter string, skip string, top string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(skip) > 0 {
		queryParameters["$skip"] = autorest.Encode("query", skip)
	}
	if len(top) > 0 {
		queryParameters["$top"] = autorest.Encode("query", top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurationRuns", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(clientRequestID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("clientRequestId", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SoftwareUpdateConfigurationRunsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SoftwareUpdateConfigurationRunsClient) ListResponder(resp *http.Response) (result SoftwareUpdateConfigurationRunListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
