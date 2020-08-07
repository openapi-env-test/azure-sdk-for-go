package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
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

// InterfaceIPConfigurationsClient is the network Client
type InterfaceIPConfigurationsClient struct {
	BaseClient
}

// NewInterfaceIPConfigurationsClient creates an instance of the InterfaceIPConfigurationsClient client.
func NewInterfaceIPConfigurationsClient(subscriptionID string) InterfaceIPConfigurationsClient {
	return NewInterfaceIPConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInterfaceIPConfigurationsClientWithBaseURI creates an instance of the InterfaceIPConfigurationsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewInterfaceIPConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) InterfaceIPConfigurationsClient {
	return InterfaceIPConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the specified network interface ip configuration.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkInterfaceName - the name of the network interface.
// IPConfigurationName - the name of the ip configuration name.
func (client InterfaceIPConfigurationsClient) Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, IPConfigurationName string) (result InterfaceIPConfiguration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfaceIPConfigurationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, networkInterfaceName, IPConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client InterfaceIPConfigurationsClient) GetPreparer(ctx context.Context, resourceGroupName string, networkInterfaceName string, IPConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ipConfigurationName":  autorest.Encode("path", IPConfigurationName),
		"networkInterfaceName": autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client InterfaceIPConfigurationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InterfaceIPConfigurationsClient) GetResponder(resp *http.Response) (result InterfaceIPConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get all ip configurations in a network interface
// Parameters:
// resourceGroupName - the name of the resource group.
// networkInterfaceName - the name of the network interface.
func (client InterfaceIPConfigurationsClient) List(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result InterfaceIPConfigurationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfaceIPConfigurationsClient.List")
		defer func() {
			sc := -1
			if result.iiclr.Response.Response != nil {
				sc = result.iiclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, networkInterfaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.iiclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "List", resp, "Failure sending request")
		return
	}

	result.iiclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "List", resp, "Failure responding to request")
	}
	if result.iiclr.hasNextLink() && result.iiclr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client InterfaceIPConfigurationsClient) ListPreparer(ctx context.Context, resourceGroupName string, networkInterfaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName": autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client InterfaceIPConfigurationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client InterfaceIPConfigurationsClient) ListResponder(resp *http.Response) (result InterfaceIPConfigurationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client InterfaceIPConfigurationsClient) listNextResults(ctx context.Context, lastResults InterfaceIPConfigurationListResult) (result InterfaceIPConfigurationListResult, err error) {
	req, err := lastResults.interfaceIPConfigurationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client InterfaceIPConfigurationsClient) ListComplete(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result InterfaceIPConfigurationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfaceIPConfigurationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, networkInterfaceName)
	return
}
