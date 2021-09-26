package compute

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

// DiskEncryptionSetsClient is the compute Client
type DiskEncryptionSetsClient struct {
	BaseClient
}

// NewDiskEncryptionSetsClient creates an instance of the DiskEncryptionSetsClient client.
func NewDiskEncryptionSetsClient(subscriptionID string) DiskEncryptionSetsClient {
	return NewDiskEncryptionSetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDiskEncryptionSetsClientWithBaseURI creates an instance of the DiskEncryptionSetsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewDiskEncryptionSetsClientWithBaseURI(baseURI string, subscriptionID string) DiskEncryptionSetsClient {
	return DiskEncryptionSetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a disk encryption set
// Parameters:
// resourceGroupName - the name of the resource group.
// diskEncryptionSetName - the name of the disk encryption set that is being created. The name can't be changed
// after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The
// maximum name length is 80 characters.
// diskEncryptionSet - disk encryption set object supplied in the body of the Put disk encryption set
// operation.
func (client DiskEncryptionSetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSet) (result DiskEncryptionSetsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskEncryptionSetsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: diskEncryptionSet,
			Constraints: []validation.Constraint{{Target: "diskEncryptionSet.EncryptionSetProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "diskEncryptionSet.EncryptionSetProperties.ActiveKey", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "diskEncryptionSet.EncryptionSetProperties.ActiveKey.SourceVault", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "diskEncryptionSet.EncryptionSetProperties.ActiveKey.KeyURL", Name: validation.Null, Rule: true, Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("compute.DiskEncryptionSetsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DiskEncryptionSetsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSet) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskEncryptionSetName": autorest.Encode("path", diskEncryptionSetName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}", pathParameters),
		autorest.WithJSON(diskEncryptionSet),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DiskEncryptionSetsClient) CreateOrUpdateSender(req *http.Request) (future DiskEncryptionSetsCreateOrUpdateFuture, err error) {
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
func (client DiskEncryptionSetsClient) CreateOrUpdateResponder(resp *http.Response) (result DiskEncryptionSet, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a disk encryption set.
// Parameters:
// resourceGroupName - the name of the resource group.
// diskEncryptionSetName - the name of the disk encryption set that is being created. The name can't be changed
// after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The
// maximum name length is 80 characters.
func (client DiskEncryptionSetsClient) Delete(ctx context.Context, resourceGroupName string, diskEncryptionSetName string) (result DiskEncryptionSetsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskEncryptionSetsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, diskEncryptionSetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DiskEncryptionSetsClient) DeletePreparer(ctx context.Context, resourceGroupName string, diskEncryptionSetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskEncryptionSetName": autorest.Encode("path", diskEncryptionSetName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DiskEncryptionSetsClient) DeleteSender(req *http.Request) (future DiskEncryptionSetsDeleteFuture, err error) {
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
func (client DiskEncryptionSetsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about a disk encryption set.
// Parameters:
// resourceGroupName - the name of the resource group.
// diskEncryptionSetName - the name of the disk encryption set that is being created. The name can't be changed
// after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The
// maximum name length is 80 characters.
func (client DiskEncryptionSetsClient) Get(ctx context.Context, resourceGroupName string, diskEncryptionSetName string) (result DiskEncryptionSet, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskEncryptionSetsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, diskEncryptionSetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DiskEncryptionSetsClient) GetPreparer(ctx context.Context, resourceGroupName string, diskEncryptionSetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskEncryptionSetName": autorest.Encode("path", diskEncryptionSetName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DiskEncryptionSetsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DiskEncryptionSetsClient) GetResponder(resp *http.Response) (result DiskEncryptionSet, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the disk encryption sets under a subscription.
func (client DiskEncryptionSetsClient) List(ctx context.Context) (result DiskEncryptionSetListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskEncryptionSetsClient.List")
		defer func() {
			sc := -1
			if result.desl.Response.Response != nil {
				sc = result.desl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.desl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "List", resp, "Failure sending request")
		return
	}

	result.desl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.desl.hasNextLink() && result.desl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client DiskEncryptionSetsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/diskEncryptionSets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DiskEncryptionSetsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DiskEncryptionSetsClient) ListResponder(resp *http.Response) (result DiskEncryptionSetList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DiskEncryptionSetsClient) listNextResults(ctx context.Context, lastResults DiskEncryptionSetList) (result DiskEncryptionSetList, err error) {
	req, err := lastResults.diskEncryptionSetListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DiskEncryptionSetsClient) ListComplete(ctx context.Context) (result DiskEncryptionSetListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskEncryptionSetsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListAssociatedResources lists all resources that are encrypted with this disk encryption set.
// Parameters:
// resourceGroupName - the name of the resource group.
// diskEncryptionSetName - the name of the disk encryption set that is being created. The name can't be changed
// after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The
// maximum name length is 80 characters.
func (client DiskEncryptionSetsClient) ListAssociatedResources(ctx context.Context, resourceGroupName string, diskEncryptionSetName string) (result ResourceURIListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskEncryptionSetsClient.ListAssociatedResources")
		defer func() {
			sc := -1
			if result.rul.Response.Response != nil {
				sc = result.rul.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAssociatedResourcesNextResults
	req, err := client.ListAssociatedResourcesPreparer(ctx, resourceGroupName, diskEncryptionSetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "ListAssociatedResources", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAssociatedResourcesSender(req)
	if err != nil {
		result.rul.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "ListAssociatedResources", resp, "Failure sending request")
		return
	}

	result.rul, err = client.ListAssociatedResourcesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "ListAssociatedResources", resp, "Failure responding to request")
		return
	}
	if result.rul.hasNextLink() && result.rul.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListAssociatedResourcesPreparer prepares the ListAssociatedResources request.
func (client DiskEncryptionSetsClient) ListAssociatedResourcesPreparer(ctx context.Context, resourceGroupName string, diskEncryptionSetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskEncryptionSetName": autorest.Encode("path", diskEncryptionSetName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}/associatedResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAssociatedResourcesSender sends the ListAssociatedResources request. The method will close the
// http.Response Body if it receives an error.
func (client DiskEncryptionSetsClient) ListAssociatedResourcesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAssociatedResourcesResponder handles the response to the ListAssociatedResources request. The method always
// closes the http.Response Body.
func (client DiskEncryptionSetsClient) ListAssociatedResourcesResponder(resp *http.Response) (result ResourceURIList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAssociatedResourcesNextResults retrieves the next set of results, if any.
func (client DiskEncryptionSetsClient) listAssociatedResourcesNextResults(ctx context.Context, lastResults ResourceURIList) (result ResourceURIList, err error) {
	req, err := lastResults.resourceURIListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "listAssociatedResourcesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAssociatedResourcesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "listAssociatedResourcesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAssociatedResourcesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "listAssociatedResourcesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAssociatedResourcesComplete enumerates all values, automatically crossing page boundaries as required.
func (client DiskEncryptionSetsClient) ListAssociatedResourcesComplete(ctx context.Context, resourceGroupName string, diskEncryptionSetName string) (result ResourceURIListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskEncryptionSetsClient.ListAssociatedResources")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAssociatedResources(ctx, resourceGroupName, diskEncryptionSetName)
	return
}

// ListByResourceGroup lists all the disk encryption sets under a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client DiskEncryptionSetsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result DiskEncryptionSetListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskEncryptionSetsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.desl.Response.Response != nil {
				sc = result.desl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.desl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.desl, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.desl.hasNextLink() && result.desl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client DiskEncryptionSetsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DiskEncryptionSetsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client DiskEncryptionSetsClient) ListByResourceGroupResponder(resp *http.Response) (result DiskEncryptionSetList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client DiskEncryptionSetsClient) listByResourceGroupNextResults(ctx context.Context, lastResults DiskEncryptionSetList) (result DiskEncryptionSetList, err error) {
	req, err := lastResults.diskEncryptionSetListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client DiskEncryptionSetsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result DiskEncryptionSetListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskEncryptionSetsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// Update updates (patches) a disk encryption set.
// Parameters:
// resourceGroupName - the name of the resource group.
// diskEncryptionSetName - the name of the disk encryption set that is being created. The name can't be changed
// after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The
// maximum name length is 80 characters.
// diskEncryptionSet - disk encryption set object supplied in the body of the Patch disk encryption set
// operation.
func (client DiskEncryptionSetsClient) Update(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSetUpdate) (result DiskEncryptionSetsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskEncryptionSetsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, diskEncryptionSetName, diskEncryptionSet)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskEncryptionSetsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DiskEncryptionSetsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, diskEncryptionSetName string, diskEncryptionSet DiskEncryptionSetUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskEncryptionSetName": autorest.Encode("path", diskEncryptionSetName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}", pathParameters),
		autorest.WithJSON(diskEncryptionSet),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DiskEncryptionSetsClient) UpdateSender(req *http.Request) (future DiskEncryptionSetsUpdateFuture, err error) {
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

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DiskEncryptionSetsClient) UpdateResponder(resp *http.Response) (result DiskEncryptionSet, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
