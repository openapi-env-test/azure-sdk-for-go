package network

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

// VpnSiteLinksClient is the network Client
type VpnSiteLinksClient struct {
	BaseClient
}

// NewVpnSiteLinksClient creates an instance of the VpnSiteLinksClient client.
func NewVpnSiteLinksClient(subscriptionID string) VpnSiteLinksClient {
	return NewVpnSiteLinksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVpnSiteLinksClientWithBaseURI creates an instance of the VpnSiteLinksClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewVpnSiteLinksClientWithBaseURI(baseURI string, subscriptionID string) VpnSiteLinksClient {
	return VpnSiteLinksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get retrieves the details of a VPN site link.
// Parameters:
// resourceGroupName - the resource group name of the VpnSite.
// vpnSiteName - the name of the VpnSite.
// vpnSiteLinkName - the name of the VpnSiteLink being retrieved.
func (client VpnSiteLinksClient) Get(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteLinkName string) (result VpnSiteLink, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnSiteLinksClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, vpnSiteName, vpnSiteLinkName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSiteLinksClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VpnSiteLinksClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSiteLinksClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VpnSiteLinksClient) GetPreparer(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteLinkName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vpnSiteLinkName":   autorest.Encode("path", vpnSiteLinkName),
		"vpnSiteName":       autorest.Encode("path", vpnSiteName),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}/vpnSiteLinks/{vpnSiteLinkName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VpnSiteLinksClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VpnSiteLinksClient) GetResponder(resp *http.Response) (result VpnSiteLink, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByVpnSite lists all the vpnSiteLinks in a resource group for a vpn site.
// Parameters:
// resourceGroupName - the resource group name of the VpnSite.
// vpnSiteName - the name of the VpnSite.
func (client VpnSiteLinksClient) ListByVpnSite(ctx context.Context, resourceGroupName string, vpnSiteName string) (result ListVpnSiteLinksResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnSiteLinksClient.ListByVpnSite")
		defer func() {
			sc := -1
			if result.lvslr.Response.Response != nil {
				sc = result.lvslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByVpnSiteNextResults
	req, err := client.ListByVpnSitePreparer(ctx, resourceGroupName, vpnSiteName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSiteLinksClient", "ListByVpnSite", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByVpnSiteSender(req)
	if err != nil {
		result.lvslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VpnSiteLinksClient", "ListByVpnSite", resp, "Failure sending request")
		return
	}

	result.lvslr, err = client.ListByVpnSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSiteLinksClient", "ListByVpnSite", resp, "Failure responding to request")
		return
	}
	if result.lvslr.hasNextLink() && result.lvslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByVpnSitePreparer prepares the ListByVpnSite request.
func (client VpnSiteLinksClient) ListByVpnSitePreparer(ctx context.Context, resourceGroupName string, vpnSiteName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vpnSiteName":       autorest.Encode("path", vpnSiteName),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}/vpnSiteLinks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByVpnSiteSender sends the ListByVpnSite request. The method will close the
// http.Response Body if it receives an error.
func (client VpnSiteLinksClient) ListByVpnSiteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByVpnSiteResponder handles the response to the ListByVpnSite request. The method always
// closes the http.Response Body.
func (client VpnSiteLinksClient) ListByVpnSiteResponder(resp *http.Response) (result ListVpnSiteLinksResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByVpnSiteNextResults retrieves the next set of results, if any.
func (client VpnSiteLinksClient) listByVpnSiteNextResults(ctx context.Context, lastResults ListVpnSiteLinksResult) (result ListVpnSiteLinksResult, err error) {
	req, err := lastResults.listVpnSiteLinksResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VpnSiteLinksClient", "listByVpnSiteNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByVpnSiteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VpnSiteLinksClient", "listByVpnSiteNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByVpnSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSiteLinksClient", "listByVpnSiteNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByVpnSiteComplete enumerates all values, automatically crossing page boundaries as required.
func (client VpnSiteLinksClient) ListByVpnSiteComplete(ctx context.Context, resourceGroupName string, vpnSiteName string) (result ListVpnSiteLinksResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnSiteLinksClient.ListByVpnSite")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByVpnSite(ctx, resourceGroupName, vpnSiteName)
	return
}
