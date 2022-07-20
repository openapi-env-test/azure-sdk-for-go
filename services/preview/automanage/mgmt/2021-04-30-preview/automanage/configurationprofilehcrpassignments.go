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

// ConfigurationProfileHCRPAssignmentsClient is the automanage Client
type ConfigurationProfileHCRPAssignmentsClient struct {
	BaseClient
}

// NewConfigurationProfileHCRPAssignmentsClient creates an instance of the ConfigurationProfileHCRPAssignmentsClient
// client.
func NewConfigurationProfileHCRPAssignmentsClient(subscriptionID string) ConfigurationProfileHCRPAssignmentsClient {
	return NewConfigurationProfileHCRPAssignmentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConfigurationProfileHCRPAssignmentsClientWithBaseURI creates an instance of the
// ConfigurationProfileHCRPAssignmentsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewConfigurationProfileHCRPAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationProfileHCRPAssignmentsClient {
	return ConfigurationProfileHCRPAssignmentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates an association between a ARC machine and Automanage configuration profile
// Parameters:
// parameters - parameters supplied to the create or update configuration profile assignment.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// machineName - the name of the Arc machine.
// configurationProfileAssignmentName - name of the configuration profile assignment. Only default is
// supported.
func (client ConfigurationProfileHCRPAssignmentsClient) CreateOrUpdate(ctx context.Context, parameters ConfigurationProfileAssignment, resourceGroupName string, machineName string, configurationProfileAssignmentName string) (result ConfigurationProfileAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationProfileHCRPAssignmentsClient.CreateOrUpdate")
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
		return result, validation.NewError("automanage.ConfigurationProfileHCRPAssignmentsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, parameters, resourceGroupName, machineName, configurationProfileAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileHCRPAssignmentsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileHCRPAssignmentsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileHCRPAssignmentsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ConfigurationProfileHCRPAssignmentsClient) CreateOrUpdatePreparer(ctx context.Context, parameters ConfigurationProfileAssignment, resourceGroupName string, machineName string, configurationProfileAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationProfileAssignmentName": autorest.Encode("path", configurationProfileAssignmentName),
		"machineName":                        autorest.Encode("path", machineName),
		"resourceGroupName":                  autorest.Encode("path", resourceGroupName),
		"subscriptionId":                     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-04-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationProfileHCRPAssignmentsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ConfigurationProfileHCRPAssignmentsClient) CreateOrUpdateResponder(resp *http.Response) (result ConfigurationProfileAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a configuration profile assignment
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// machineName - the name of the Arc machine.
// configurationProfileAssignmentName - name of the configuration profile assignment
func (client ConfigurationProfileHCRPAssignmentsClient) Delete(ctx context.Context, resourceGroupName string, machineName string, configurationProfileAssignmentName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationProfileHCRPAssignmentsClient.Delete")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automanage.ConfigurationProfileHCRPAssignmentsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, machineName, configurationProfileAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileHCRPAssignmentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileHCRPAssignmentsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileHCRPAssignmentsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ConfigurationProfileHCRPAssignmentsClient) DeletePreparer(ctx context.Context, resourceGroupName string, machineName string, configurationProfileAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationProfileAssignmentName": autorest.Encode("path", configurationProfileAssignmentName),
		"machineName":                        autorest.Encode("path", machineName),
		"resourceGroupName":                  autorest.Encode("path", resourceGroupName),
		"subscriptionId":                     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-04-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationProfileHCRPAssignmentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ConfigurationProfileHCRPAssignmentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get information about a configuration profile assignment
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// machineName - the name of the Arc machine.
// configurationProfileAssignmentName - the configuration profile assignment name.
func (client ConfigurationProfileHCRPAssignmentsClient) Get(ctx context.Context, resourceGroupName string, machineName string, configurationProfileAssignmentName string) (result ConfigurationProfileAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationProfileHCRPAssignmentsClient.Get")
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
		return result, validation.NewError("automanage.ConfigurationProfileHCRPAssignmentsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, machineName, configurationProfileAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileHCRPAssignmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileHCRPAssignmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfileHCRPAssignmentsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConfigurationProfileHCRPAssignmentsClient) GetPreparer(ctx context.Context, resourceGroupName string, machineName string, configurationProfileAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationProfileAssignmentName": autorest.Encode("path", configurationProfileAssignmentName),
		"machineName":                        autorest.Encode("path", machineName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationProfileHCRPAssignmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConfigurationProfileHCRPAssignmentsClient) GetResponder(resp *http.Response) (result ConfigurationProfileAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
