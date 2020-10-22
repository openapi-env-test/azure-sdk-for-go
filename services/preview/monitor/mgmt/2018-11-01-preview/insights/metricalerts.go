package insights

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
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
    "github.com/Azure/go-autorest/autorest/validation"
)

// MetricAlertsClient is the monitor Management Client
type MetricAlertsClient struct {
    BaseClient
}
// NewMetricAlertsClient creates an instance of the MetricAlertsClient client.
func NewMetricAlertsClient(subscriptionID string) MetricAlertsClient {
    return NewMetricAlertsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMetricAlertsClientWithBaseURI creates an instance of the MetricAlertsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewMetricAlertsClientWithBaseURI(baseURI string, subscriptionID string) MetricAlertsClient {
        return MetricAlertsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate create or update an metric alert definition.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // ruleName - the name of the rule.
        // parameters - the parameters of the rule to create or update.
func (client MetricAlertsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters MetricAlertResource) (result MetricAlertResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MetricAlertsClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: parameters,
         Constraints: []validation.Constraint{	{Target: "parameters.MetricAlertProperties", Name: validation.Null, Rule: true ,
        Chain: []validation.Constraint{	{Target: "parameters.MetricAlertProperties.Description", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "parameters.MetricAlertProperties.Severity", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "parameters.MetricAlertProperties.Enabled", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "parameters.MetricAlertProperties.EvaluationFrequency", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "parameters.MetricAlertProperties.WindowSize", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "parameters.MetricAlertProperties.Criteria", Name: validation.Null, Rule: true, Chain: nil },
        }}}}}); err != nil {
        return result, validation.NewError("insights.MetricAlertsClient", "CreateOrUpdate", err.Error())
        }

        req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, ruleName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        resp, err := client.CreateOrUpdateSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "CreateOrUpdate", resp, "Failure sending request")
        return
        }

        result, err = client.CreateOrUpdateResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "CreateOrUpdate", resp, "Failure responding to request")
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client MetricAlertsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, ruleName string, parameters MetricAlertResource) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "ruleName": autorest.Encode("path",ruleName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2018-03-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts/{ruleName}",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client MetricAlertsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
    // closes the http.Response Body.
    func (client MetricAlertsClient) CreateOrUpdateResponder(resp *http.Response) (result MetricAlertResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete delete an alert rule definition.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // ruleName - the name of the rule.
func (client MetricAlertsClient) Delete(ctx context.Context, resourceGroupName string, ruleName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MetricAlertsClient.Delete")
        defer func() {
            sc := -1
        if result.Response != nil {
        sc = result.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.DeletePreparer(ctx, resourceGroupName, ruleName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "Delete", nil , "Failure preparing request")
    return
    }

        resp, err := client.DeleteSender(req)
        if err != nil {
        result.Response = resp
        err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "Delete", resp, "Failure sending request")
        return
        }

        result, err = client.DeleteResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "Delete", resp, "Failure responding to request")
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client MetricAlertsClient) DeletePreparer(ctx context.Context, resourceGroupName string, ruleName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "ruleName": autorest.Encode("path",ruleName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2018-03-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts/{ruleName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client MetricAlertsClient) DeleteSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // DeleteResponder handles the response to the Delete request. The method always
    // closes the http.Response Body.
    func (client MetricAlertsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get retrieve an alert rule definition.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // ruleName - the name of the rule.
func (client MetricAlertsClient) Get(ctx context.Context, resourceGroupName string, ruleName string) (result MetricAlertResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MetricAlertsClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, resourceGroupName, ruleName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "Get", resp, "Failure responding to request")
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client MetricAlertsClient) GetPreparer(ctx context.Context, resourceGroupName string, ruleName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "ruleName": autorest.Encode("path",ruleName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2018-03-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts/{ruleName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client MetricAlertsClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client MetricAlertsClient) GetResponder(resp *http.Response) (result MetricAlertResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// ListByResourceGroup retrieve alert rule definitions in a resource group.
    // Parameters:
        // resourceGroupName - the name of the resource group.
func (client MetricAlertsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result MetricAlertResourceCollection, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MetricAlertsClient.ListByResourceGroup")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "ListByResourceGroup", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListByResourceGroupSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "ListByResourceGroup", resp, "Failure sending request")
        return
        }

        result, err = client.ListByResourceGroupResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "ListByResourceGroup", resp, "Failure responding to request")
        }

    return
}

    // ListByResourceGroupPreparer prepares the ListByResourceGroup request.
    func (client MetricAlertsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2018-03-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
    // http.Response Body if it receives an error.
    func (client MetricAlertsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
    // closes the http.Response Body.
    func (client MetricAlertsClient) ListByResourceGroupResponder(resp *http.Response) (result MetricAlertResourceCollection, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// ListBySubscription retrieve alert rule definitions in a subscription.
func (client MetricAlertsClient) ListBySubscription(ctx context.Context) (result MetricAlertResourceCollection, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MetricAlertsClient.ListBySubscription")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.ListBySubscriptionPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "ListBySubscription", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListBySubscriptionSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "ListBySubscription", resp, "Failure sending request")
        return
        }

        result, err = client.ListBySubscriptionResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "ListBySubscription", resp, "Failure responding to request")
        }

    return
}

    // ListBySubscriptionPreparer prepares the ListBySubscription request.
    func (client MetricAlertsClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2018-03-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Insights/metricAlerts",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListBySubscriptionSender sends the ListBySubscription request. The method will close the
    // http.Response Body if it receives an error.
    func (client MetricAlertsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
    // closes the http.Response Body.
    func (client MetricAlertsClient) ListBySubscriptionResponder(resp *http.Response) (result MetricAlertResourceCollection, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Update update an metric alert definition.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // ruleName - the name of the rule.
        // parameters - the parameters of the rule to update.
func (client MetricAlertsClient) Update(ctx context.Context, resourceGroupName string, ruleName string, parameters MetricAlertResourcePatch) (result MetricAlertResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MetricAlertsClient.Update")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.UpdatePreparer(ctx, resourceGroupName, ruleName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "Update", nil , "Failure preparing request")
    return
    }

        resp, err := client.UpdateSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "Update", resp, "Failure sending request")
        return
        }

        result, err = client.UpdateResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "insights.MetricAlertsClient", "Update", resp, "Failure responding to request")
        }

    return
}

    // UpdatePreparer prepares the Update request.
    func (client MetricAlertsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, ruleName string, parameters MetricAlertResourcePatch) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "ruleName": autorest.Encode("path",ruleName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2018-03-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPatch(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts/{ruleName}",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client MetricAlertsClient) UpdateSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // UpdateResponder handles the response to the Update request. The method always
    // closes the http.Response Body.
    func (client MetricAlertsClient) UpdateResponder(resp *http.Response) (result MetricAlertResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

