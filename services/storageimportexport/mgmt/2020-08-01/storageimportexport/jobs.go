package storageimportexport

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

// JobsClient is the the Storage Import/Export Resource Provider API.
type JobsClient struct {
	BaseClient
}

// NewJobsClient creates an instance of the JobsClient client.
func NewJobsClient(subscriptionID string, acceptLanguage string) JobsClient {
	return NewJobsClientWithBaseURI(DefaultBaseURI, subscriptionID, acceptLanguage)
}

// NewJobsClientWithBaseURI creates an instance of the JobsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) JobsClient {
	return JobsClient{NewWithBaseURI(baseURI, subscriptionID, acceptLanguage)}
}

// Create creates a new job or updates an existing job in the specified subscription.
// Parameters:
// jobName - the name of the import/export job.
// resourceGroupName - the resource group name uniquely identifies the resource group within the user
// subscription.
// body - the parameters used for creating the job
// clientTenantID - the tenant ID of the client making the request.
func (client JobsClient) Create(ctx context.Context, jobName string, resourceGroupName string, body PutJobParameters, clientTenantID string) (result JobResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "body.Properties.ReturnAddress", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.Properties.ReturnAddress.RecipientName", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "body.Properties.ReturnAddress.StreetAddress1", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "body.Properties.ReturnAddress.City", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "body.Properties.ReturnAddress.PostalCode", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "body.Properties.ReturnAddress.CountryOrRegion", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "body.Properties.ReturnAddress.Phone", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "body.Properties.ReturnAddress.Email", Name: validation.Null, Rule: true, Chain: nil},
					}},
					{Target: "body.Properties.ReturnShipping", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "body.Properties.ReturnShipping.CarrierName", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "body.Properties.ReturnShipping.CarrierAccountNumber", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "body.Properties.DeliveryPackage", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "body.Properties.DeliveryPackage.CarrierName", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "body.Properties.DeliveryPackage.TrackingNumber", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "body.Properties.ReturnPackage", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "body.Properties.ReturnPackage.CarrierName", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "body.Properties.ReturnPackage.TrackingNumber", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "body.Properties.ReturnPackage.DriveCount", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "body.Properties.ReturnPackage.ShipDate", Name: validation.Null, Rule: true, Chain: nil},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("storageimportexport.JobsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, jobName, resourceGroupName, body, clientTenantID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client JobsClient) CreatePreparer(ctx context.Context, jobName string, resourceGroupName string, body PutJobParameters, clientTenantID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	if len(clientTenantID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-tenant-id", autorest.String(clientTenantID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client JobsClient) CreateResponder(resp *http.Response) (result JobResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an existing job. Only jobs in the Creating or Completed states can be deleted.
// Parameters:
// jobName - the name of the import/export job.
// resourceGroupName - the resource group name uniquely identifies the resource group within the user
// subscription.
func (client JobsClient) Delete(ctx context.Context, jobName string, resourceGroupName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, jobName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client JobsClient) DeletePreparer(ctx context.Context, jobName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client JobsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about an existing job.
// Parameters:
// jobName - the name of the import/export job.
// resourceGroupName - the resource group name uniquely identifies the resource group within the user
// subscription.
func (client JobsClient) Get(ctx context.Context, jobName string, resourceGroupName string) (result JobResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, jobName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobsClient) GetPreparer(ctx context.Context, jobName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobsClient) GetResponder(resp *http.Response) (result JobResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup returns all active and completed jobs in a resource group.
// Parameters:
// resourceGroupName - the resource group name uniquely identifies the resource group within the user
// subscription.
// top - an integer value that specifies how many jobs at most should be returned. The value cannot exceed 100.
// filter - can be used to restrict the results to certain conditions.
func (client JobsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, top *int32, filter string) (result ListJobsResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.ljr.Response.Response != nil {
				sc = result.ljr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.ljr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.ljr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.ljr.hasNextLink() && result.ljr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client JobsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client JobsClient) ListByResourceGroupResponder(resp *http.Response) (result ListJobsResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client JobsClient) listByResourceGroupNextResults(ctx context.Context, lastResults ListJobsResponse) (result ListJobsResponse, err error) {
	req, err := lastResults.listJobsResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, top *int32, filter string) (result ListJobsResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, top, filter)
	return
}

// ListBySubscription returns all active and completed jobs in a subscription.
// Parameters:
// top - an integer value that specifies how many jobs at most should be returned. The value cannot exceed 100.
// filter - can be used to restrict the results to certain conditions.
func (client JobsClient) ListBySubscription(ctx context.Context, top *int32, filter string) (result ListJobsResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.ljr.Response.Response != nil {
				sc = result.ljr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.ljr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.ljr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.ljr.hasNextLink() && result.ljr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client JobsClient) ListBySubscriptionPreparer(ctx context.Context, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ImportExport/jobs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client JobsClient) ListBySubscriptionResponder(resp *http.Response) (result ListJobsResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client JobsClient) listBySubscriptionNextResults(ctx context.Context, lastResults ListJobsResponse) (result ListJobsResponse, err error) {
	req, err := lastResults.listJobsResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobsClient) ListBySubscriptionComplete(ctx context.Context, top *int32, filter string) (result ListJobsResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, top, filter)
	return
}

// Update updates specific properties of a job. You can call this operation to notify the Import/Export service that
// the hard drives comprising the import or export job have been shipped to the Microsoft data center. It can also be
// used to cancel an existing job.
// Parameters:
// jobName - the name of the import/export job.
// resourceGroupName - the resource group name uniquely identifies the resource group within the user
// subscription.
// body - the parameters to update in the job
func (client JobsClient) Update(ctx context.Context, jobName string, resourceGroupName string, body UpdateJobParameters) (result JobResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, jobName, resourceGroupName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.JobsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client JobsClient) UpdatePreparer(ctx context.Context, jobName string, resourceGroupName string, body UpdateJobParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client JobsClient) UpdateResponder(resp *http.Response) (result JobResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
