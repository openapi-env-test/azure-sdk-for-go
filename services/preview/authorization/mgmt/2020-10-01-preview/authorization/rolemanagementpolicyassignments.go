package authorization

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RoleManagementPolicyAssignmentsClient is the client for the RoleManagementPolicyAssignments methods of the
// Authorization service.
type RoleManagementPolicyAssignmentsClient struct {
	BaseClient
}

// NewRoleManagementPolicyAssignmentsClient creates an instance of the RoleManagementPolicyAssignmentsClient client.
func NewRoleManagementPolicyAssignmentsClient(subscriptionID string) RoleManagementPolicyAssignmentsClient {
	return NewRoleManagementPolicyAssignmentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRoleManagementPolicyAssignmentsClientWithBaseURI creates an instance of the RoleManagementPolicyAssignmentsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewRoleManagementPolicyAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) RoleManagementPolicyAssignmentsClient {
	return RoleManagementPolicyAssignmentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a role management policy assignment
// Parameters:
// scope - the scope of the role management policy assignment to upsert.
// roleManagementPolicyAssignmentName - the name of format {guid_guid} the role management policy assignment to
// upsert.
// parameters - parameters for the role management policy assignment.
func (client RoleManagementPolicyAssignmentsClient) Create(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, parameters RoleManagementPolicyAssignment) (result RoleManagementPolicyAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleManagementPolicyAssignmentsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, scope, roleManagementPolicyAssignmentName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPolicyAssignmentsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPolicyAssignmentsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPolicyAssignmentsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client RoleManagementPolicyAssignmentsClient) CreatePreparer(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, parameters RoleManagementPolicyAssignment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleManagementPolicyAssignmentName": autorest.Encode("path", roleManagementPolicyAssignmentName),
		"scope":                              scope,
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.ID = nil
	parameters.Name = nil
	parameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleManagementPolicyAssignments/{roleManagementPolicyAssignmentName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client RoleManagementPolicyAssignmentsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client RoleManagementPolicyAssignmentsClient) CreateResponder(resp *http.Response) (result RoleManagementPolicyAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a role management policy assignment
// Parameters:
// scope - the scope of the role management policy assignment to delete.
// roleManagementPolicyAssignmentName - the name of format {guid_guid} the role management policy assignment to
// delete.
func (client RoleManagementPolicyAssignmentsClient) Delete(ctx context.Context, scope string, roleManagementPolicyAssignmentName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleManagementPolicyAssignmentsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, scope, roleManagementPolicyAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPolicyAssignmentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPolicyAssignmentsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPolicyAssignmentsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RoleManagementPolicyAssignmentsClient) DeletePreparer(ctx context.Context, scope string, roleManagementPolicyAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleManagementPolicyAssignmentName": autorest.Encode("path", roleManagementPolicyAssignmentName),
		"scope":                              scope,
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleManagementPolicyAssignments/{roleManagementPolicyAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RoleManagementPolicyAssignmentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RoleManagementPolicyAssignmentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the specified role management policy assignment for a resource scope
// Parameters:
// scope - the scope of the role management policy.
// roleManagementPolicyAssignmentName - the name of format {guid_guid} the role management policy assignment to
// get.
func (client RoleManagementPolicyAssignmentsClient) Get(ctx context.Context, scope string, roleManagementPolicyAssignmentName string) (result RoleManagementPolicyAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleManagementPolicyAssignmentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, scope, roleManagementPolicyAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPolicyAssignmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPolicyAssignmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPolicyAssignmentsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client RoleManagementPolicyAssignmentsClient) GetPreparer(ctx context.Context, scope string, roleManagementPolicyAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleManagementPolicyAssignmentName": autorest.Encode("path", roleManagementPolicyAssignmentName),
		"scope":                              scope,
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleManagementPolicyAssignments/{roleManagementPolicyAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RoleManagementPolicyAssignmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RoleManagementPolicyAssignmentsClient) GetResponder(resp *http.Response) (result RoleManagementPolicyAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListForScope gets role management assignment policies for a resource scope.
// Parameters:
// scope - the scope of the role management policy.
func (client RoleManagementPolicyAssignmentsClient) ListForScope(ctx context.Context, scope string) (result RoleManagementPolicyAssignmentListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleManagementPolicyAssignmentsClient.ListForScope")
		defer func() {
			sc := -1
			if result.rmpalr.Response.Response != nil {
				sc = result.rmpalr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listForScopeNextResults
	req, err := client.ListForScopePreparer(ctx, scope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPolicyAssignmentsClient", "ListForScope", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListForScopeSender(req)
	if err != nil {
		result.rmpalr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPolicyAssignmentsClient", "ListForScope", resp, "Failure sending request")
		return
	}

	result.rmpalr, err = client.ListForScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPolicyAssignmentsClient", "ListForScope", resp, "Failure responding to request")
		return
	}
	if result.rmpalr.hasNextLink() && result.rmpalr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListForScopePreparer prepares the ListForScope request.
func (client RoleManagementPolicyAssignmentsClient) ListForScopePreparer(ctx context.Context, scope string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleManagementPolicyAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListForScopeSender sends the ListForScope request. The method will close the
// http.Response Body if it receives an error.
func (client RoleManagementPolicyAssignmentsClient) ListForScopeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListForScopeResponder handles the response to the ListForScope request. The method always
// closes the http.Response Body.
func (client RoleManagementPolicyAssignmentsClient) ListForScopeResponder(resp *http.Response) (result RoleManagementPolicyAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listForScopeNextResults retrieves the next set of results, if any.
func (client RoleManagementPolicyAssignmentsClient) listForScopeNextResults(ctx context.Context, lastResults RoleManagementPolicyAssignmentListResult) (result RoleManagementPolicyAssignmentListResult, err error) {
	req, err := lastResults.roleManagementPolicyAssignmentListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.RoleManagementPolicyAssignmentsClient", "listForScopeNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListForScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization.RoleManagementPolicyAssignmentsClient", "listForScopeNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListForScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleManagementPolicyAssignmentsClient", "listForScopeNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListForScopeComplete enumerates all values, automatically crossing page boundaries as required.
func (client RoleManagementPolicyAssignmentsClient) ListForScopeComplete(ctx context.Context, scope string) (result RoleManagementPolicyAssignmentListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleManagementPolicyAssignmentsClient.ListForScope")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListForScope(ctx, scope)
	return
}
