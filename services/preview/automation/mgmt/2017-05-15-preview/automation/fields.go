package automation

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

// FieldsClient is the automation Client
type FieldsClient struct {
	BaseClient
}

// NewFieldsClient creates an instance of the FieldsClient client.
func NewFieldsClient(subscriptionID string) FieldsClient {
	return NewFieldsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFieldsClientWithBaseURI creates an instance of the FieldsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewFieldsClientWithBaseURI(baseURI string, subscriptionID string) FieldsClient {
	return FieldsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByType retrieve a list of fields of a given type identified by module name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// moduleName - the name of module.
// typeName - the name of type.
func (client FieldsClient) ListByType(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, typeName string) (result TypeFieldListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FieldsClient.ListByType")
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
		return result, validation.NewError("automation.FieldsClient", "ListByType", err.Error())
	}

	req, err := client.ListByTypePreparer(ctx, resourceGroupName, automationAccountName, moduleName, typeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.FieldsClient", "ListByType", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByTypeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.FieldsClient", "ListByType", resp, "Failure sending request")
		return
	}

	result, err = client.ListByTypeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.FieldsClient", "ListByType", resp, "Failure responding to request")
		return
	}

	return
}

// ListByTypePreparer prepares the ListByType request.
func (client FieldsClient) ListByTypePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, typeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"moduleName":            autorest.Encode("path", moduleName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"typeName":              autorest.Encode("path", typeName),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/types/{typeName}/fields", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByTypeSender sends the ListByType request. The method will close the
// http.Response Body if it receives an error.
func (client FieldsClient) ListByTypeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByTypeResponder handles the response to the ListByType request. The method always
// closes the http.Response Body.
func (client FieldsClient) ListByTypeResponder(resp *http.Response) (result TypeFieldListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
