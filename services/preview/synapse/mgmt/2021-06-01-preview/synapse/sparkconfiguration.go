package synapse

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

// SparkConfigurationClient is the azure Synapse Analytics Management Client
type SparkConfigurationClient struct {
	BaseClient
}

// NewSparkConfigurationClient creates an instance of the SparkConfigurationClient client.
func NewSparkConfigurationClient(subscriptionID string) SparkConfigurationClient {
	return NewSparkConfigurationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSparkConfigurationClientWithBaseURI creates an instance of the SparkConfigurationClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewSparkConfigurationClientWithBaseURI(baseURI string, subscriptionID string) SparkConfigurationClient {
	return SparkConfigurationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get SparkConfiguration by name in a workspace.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// sparkConfigurationName - sparkConfiguration name
// workspaceName - the name of the workspace.
func (client SparkConfigurationClient) Get(ctx context.Context, resourceGroupName string, sparkConfigurationName string, workspaceName string) (result SparkConfigurationResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SparkConfigurationClient.Get")
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
		return result, validation.NewError("synapse.SparkConfigurationClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, sparkConfigurationName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SparkConfigurationClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.SparkConfigurationClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SparkConfigurationClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SparkConfigurationClient) GetPreparer(ctx context.Context, resourceGroupName string, sparkConfigurationName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sparkConfigurationName": autorest.Encode("path", sparkConfigurationName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sparkconfigurations/{sparkConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SparkConfigurationClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SparkConfigurationClient) GetResponder(resp *http.Response) (result SparkConfigurationResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
