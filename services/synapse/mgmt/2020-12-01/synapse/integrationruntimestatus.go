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

// IntegrationRuntimeStatusClient is the azure Synapse Analytics Management Client
type IntegrationRuntimeStatusClient struct {
	BaseClient
}

// NewIntegrationRuntimeStatusClient creates an instance of the IntegrationRuntimeStatusClient client.
func NewIntegrationRuntimeStatusClient(subscriptionID string) IntegrationRuntimeStatusClient {
	return NewIntegrationRuntimeStatusClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIntegrationRuntimeStatusClientWithBaseURI creates an instance of the IntegrationRuntimeStatusClient client using
// a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewIntegrationRuntimeStatusClientWithBaseURI(baseURI string, subscriptionID string) IntegrationRuntimeStatusClient {
	return IntegrationRuntimeStatusClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get the integration runtime status
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// integrationRuntimeName - integration runtime name
func (client IntegrationRuntimeStatusClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (result IntegrationRuntimeStatusResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationRuntimeStatusClient.Get")
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
		return result, validation.NewError("synapse.IntegrationRuntimeStatusClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, integrationRuntimeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeStatusClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeStatusClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeStatusClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client IntegrationRuntimeStatusClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationRuntimeName": autorest.Encode("path", integrationRuntimeName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/getStatus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationRuntimeStatusClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IntegrationRuntimeStatusClient) GetResponder(resp *http.Response) (result IntegrationRuntimeStatusResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
