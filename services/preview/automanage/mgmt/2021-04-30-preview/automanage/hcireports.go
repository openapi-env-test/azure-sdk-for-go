package automanage

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

// HCIReportsClient is the automanage Client
type HCIReportsClient struct {
	BaseClient
}

// NewHCIReportsClient creates an instance of the HCIReportsClient client.
func NewHCIReportsClient(subscriptionID string) HCIReportsClient {
	return NewHCIReportsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewHCIReportsClientWithBaseURI creates an instance of the HCIReportsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewHCIReportsClientWithBaseURI(baseURI string, subscriptionID string) HCIReportsClient {
	return HCIReportsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get information about a report associated with a configuration profile assignment run
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterName - the name of the Arc machine.
// configurationProfileAssignmentName - the configuration profile assignment name.
// reportName - the report name.
func (client HCIReportsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, configurationProfileAssignmentName string, reportName string) (result Report, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HCIReportsClient.Get")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automanage.HCIReportsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName, configurationProfileAssignmentName, reportName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.HCIReportsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automanage.HCIReportsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.HCIReportsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client HCIReportsClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string, configurationProfileAssignmentName string, reportName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":                        autorest.Encode("path", clusterName),
		"configurationProfileAssignmentName": autorest.Encode("path", configurationProfileAssignmentName),
		"reportName":                         autorest.Encode("path", reportName),
		"resourceGroupName":                  autorest.Encode("path", resourceGroupName),
		"subscriptionId":                     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-04-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHci/clusters/{clusterName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}/reports/{reportName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client HCIReportsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client HCIReportsClient) GetResponder(resp *http.Response) (result Report, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByConfigurationProfileAssignments retrieve a list of reports within a given configuration profile assignment
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterName - the name of the Arc machine.
// configurationProfileAssignmentName - the configuration profile assignment name.
func (client HCIReportsClient) ListByConfigurationProfileAssignments(ctx context.Context, resourceGroupName string, clusterName string, configurationProfileAssignmentName string) (result ReportList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HCIReportsClient.ListByConfigurationProfileAssignments")
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
		return result, validation.NewError("automanage.HCIReportsClient", "ListByConfigurationProfileAssignments", err.Error())
	}

	req, err := client.ListByConfigurationProfileAssignmentsPreparer(ctx, resourceGroupName, clusterName, configurationProfileAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.HCIReportsClient", "ListByConfigurationProfileAssignments", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByConfigurationProfileAssignmentsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automanage.HCIReportsClient", "ListByConfigurationProfileAssignments", resp, "Failure sending request")
		return
	}

	result, err = client.ListByConfigurationProfileAssignmentsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.HCIReportsClient", "ListByConfigurationProfileAssignments", resp, "Failure responding to request")
		return
	}

	return
}

// ListByConfigurationProfileAssignmentsPreparer prepares the ListByConfigurationProfileAssignments request.
func (client HCIReportsClient) ListByConfigurationProfileAssignmentsPreparer(ctx context.Context, resourceGroupName string, clusterName string, configurationProfileAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":                        autorest.Encode("path", clusterName),
		"configurationProfileAssignmentName": autorest.Encode("path", configurationProfileAssignmentName),
		"resourceGroupName":                  autorest.Encode("path", resourceGroupName),
		"subscriptionId":                     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-04-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHci/clusters/{clusterName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}/reports", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByConfigurationProfileAssignmentsSender sends the ListByConfigurationProfileAssignments request. The method will close the
// http.Response Body if it receives an error.
func (client HCIReportsClient) ListByConfigurationProfileAssignmentsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByConfigurationProfileAssignmentsResponder handles the response to the ListByConfigurationProfileAssignments request. The method always
// closes the http.Response Body.
func (client HCIReportsClient) ListByConfigurationProfileAssignmentsResponder(resp *http.Response) (result ReportList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
