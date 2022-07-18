package desktopvirtualization

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

// MsixImagesClient is the client for the MsixImages methods of the Desktopvirtualization service.
type MsixImagesClient struct {
	BaseClient
}

// NewMsixImagesClient creates an instance of the MsixImagesClient client.
func NewMsixImagesClient(subscriptionID string) MsixImagesClient {
	return NewMsixImagesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMsixImagesClientWithBaseURI creates an instance of the MsixImagesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewMsixImagesClientWithBaseURI(baseURI string, subscriptionID string) MsixImagesClient {
	return MsixImagesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Expand expands and Lists MSIX packages in an Image, given the Image Path.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// hostPoolName - the name of the host pool within the specified resource group
// msixImageURI - object containing URI to MSIX Image
func (client MsixImagesClient) Expand(ctx context.Context, resourceGroupName string, hostPoolName string, msixImageURI MSIXImageURI) (result ExpandMsixImageListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MsixImagesClient.Expand")
		defer func() {
			sc := -1
			if result.emil.Response.Response != nil {
				sc = result.emil.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: hostPoolName,
			Constraints: []validation.Constraint{{Target: "hostPoolName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "hostPoolName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.MsixImagesClient", "Expand", err.Error())
	}

	result.fn = client.expandNextResults
	req, err := client.ExpandPreparer(ctx, resourceGroupName, hostPoolName, msixImageURI)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.MsixImagesClient", "Expand", nil, "Failure preparing request")
		return
	}

	resp, err := client.ExpandSender(req)
	if err != nil {
		result.emil.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "desktopvirtualization.MsixImagesClient", "Expand", resp, "Failure sending request")
		return
	}

	result.emil, err = client.ExpandResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.MsixImagesClient", "Expand", resp, "Failure responding to request")
		return
	}
	if result.emil.hasNextLink() && result.emil.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ExpandPreparer prepares the Expand request.
func (client MsixImagesClient) ExpandPreparer(ctx context.Context, resourceGroupName string, hostPoolName string, msixImageURI MSIXImageURI) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostPoolName":      autorest.Encode("path", hostPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/expandMsixImage", pathParameters),
		autorest.WithJSON(msixImageURI),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExpandSender sends the Expand request. The method will close the
// http.Response Body if it receives an error.
func (client MsixImagesClient) ExpandSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ExpandResponder handles the response to the Expand request. The method always
// closes the http.Response Body.
func (client MsixImagesClient) ExpandResponder(resp *http.Response) (result ExpandMsixImageList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// expandNextResults retrieves the next set of results, if any.
func (client MsixImagesClient) expandNextResults(ctx context.Context, lastResults ExpandMsixImageList) (result ExpandMsixImageList, err error) {
	req, err := lastResults.expandMsixImageListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "desktopvirtualization.MsixImagesClient", "expandNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ExpandSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "desktopvirtualization.MsixImagesClient", "expandNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ExpandResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.MsixImagesClient", "expandNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ExpandComplete enumerates all values, automatically crossing page boundaries as required.
func (client MsixImagesClient) ExpandComplete(ctx context.Context, resourceGroupName string, hostPoolName string, msixImageURI MSIXImageURI) (result ExpandMsixImageListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MsixImagesClient.Expand")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.Expand(ctx, resourceGroupName, hostPoolName, msixImageURI)
	return
}
