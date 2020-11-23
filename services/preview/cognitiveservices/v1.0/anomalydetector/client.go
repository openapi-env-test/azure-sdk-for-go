// Package anomalydetector implements the Azure ARM Anomalydetector service API version 1.0.
//
// The Anomaly Detector API detects anomalies automatically in time series data. It supports two kinds of mode, one is
// for stateless using, another is for stateful using. In stateless mode, there are three functionalities. Entire
// Detect is for detecting the whole series with model trained by the time series, Last Detect is detecting last point
// with model trained by points before. ChangePoint Detect is for detecting trend changes in time series. In stateful
// mode, user can store time series, the stored time series will be used for detection anomalies. Under this mode, user
// can still use the above three functionalities by only giving a time range without preparing time series in client
// side. Besides the above three functionalities, stateful model also provide group based detection and labeling
// service. By leveraging labeling service user can provide labels for each detection result, these labels will be used
// for retuning or regenerating detection models. Inconsistency detection is a kind of group based detection, this
// detection will find inconsistency ones in a set of time series. By using anomaly detector service, business
// customers can discover incidents and establish a logic flow for root cause analysis.
package anomalydetector

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
	"crypto/tls"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// BaseClient is the base client for Anomalydetector.
type BaseClient struct {
	autorest.Client
	Endpoint string
}

// New creates an instance of the BaseClient client.
func New(endpoint string) BaseClient {
	return NewWithoutDefaults(endpoint)
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults(endpoint string) BaseClient {
	return BaseClient{
		Client:   autorest.NewClientWithOptions(autorest.ClientOptions{UserAgent: UserAgent(), Renegotiation: tls.RenegotiateFreelyAsClient}),
		Endpoint: endpoint,
	}
}

// BreakingChangeDetectEntireSeries this operation generates a model using an entire series, each point is detected
// with the same model. With this method, points before and after a certain point are used to determine whether it is
// an anomaly. The entire detection can give user an overall status of the time series.
// Parameters:
// body - time series points and period if needed. Advanced model parameters can also be set in the request.
func (client BaseClient) BreakingChangeDetectEntireSeries(ctx context.Context, body DetectRequest) (result EntireDetectResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.BreakingChangeDetectEntireSeries")
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
			Constraints: []validation.Constraint{{Target: "body.Series", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("anomalydetector.BaseClient", "BreakingChangeDetectEntireSeries", err.Error())
	}

	req, err := client.BreakingChangeDetectEntireSeriesPreparer(ctx, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "anomalydetector.BaseClient", "BreakingChangeDetectEntireSeries", nil, "Failure preparing request")
		return
	}

	resp, err := client.BreakingChangeDetectEntireSeriesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "anomalydetector.BaseClient", "BreakingChangeDetectEntireSeries", resp, "Failure sending request")
		return
	}

	result, err = client.BreakingChangeDetectEntireSeriesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "anomalydetector.BaseClient", "BreakingChangeDetectEntireSeries", resp, "Failure responding to request")
	}

	return
}

// BreakingChangeDetectEntireSeriesPreparer prepares the BreakingChangeDetectEntireSeries request.
func (client BaseClient) BreakingChangeDetectEntireSeriesPreparer(ctx context.Context, body DetectRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/anomalydetector/v1.0", urlParameters),
		autorest.WithPath("/timeseries/entire/detect"),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// BreakingChangeDetectEntireSeriesSender sends the BreakingChangeDetectEntireSeries request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) BreakingChangeDetectEntireSeriesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// BreakingChangeDetectEntireSeriesResponder handles the response to the BreakingChangeDetectEntireSeries request. The method always
// closes the http.Response Body.
func (client BaseClient) BreakingChangeDetectEntireSeriesResponder(resp *http.Response) (result EntireDetectResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DetectChangePoint evaluate change point score of every series point
// Parameters:
// body - time series points and granularity is needed. Advanced model parameters can also be set in the
// request if needed.
func (client BaseClient) DetectChangePoint(ctx context.Context, body ChangePointDetectRequest) (result ChangePointDetectResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.DetectChangePoint")
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
			Constraints: []validation.Constraint{{Target: "body.Series", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("anomalydetector.BaseClient", "DetectChangePoint", err.Error())
	}

	req, err := client.DetectChangePointPreparer(ctx, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "anomalydetector.BaseClient", "DetectChangePoint", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetectChangePointSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "anomalydetector.BaseClient", "DetectChangePoint", resp, "Failure sending request")
		return
	}

	result, err = client.DetectChangePointResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "anomalydetector.BaseClient", "DetectChangePoint", resp, "Failure responding to request")
	}

	return
}

// DetectChangePointPreparer prepares the DetectChangePoint request.
func (client BaseClient) DetectChangePointPreparer(ctx context.Context, body ChangePointDetectRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/anomalydetector/v1.0", urlParameters),
		autorest.WithPath("/timeseries/changepoint/detect"),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DetectChangePointSender sends the DetectChangePoint request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) DetectChangePointSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DetectChangePointResponder handles the response to the DetectChangePoint request. The method always
// closes the http.Response Body.
func (client BaseClient) DetectChangePointResponder(resp *http.Response) (result ChangePointDetectResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DetectLastPoint this operation generates a model using points before the latest one. With this method, only
// historical points are used to determine whether the target point is an anomaly. The latest point detecting operation
// matches the scenario of real-time monitoring of business metrics.
// Parameters:
// body - time series points and period if needed. Advanced model parameters can also be set in the request.
func (client BaseClient) DetectLastPoint(ctx context.Context, body DetectRequest) (result LastDetectResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.DetectLastPoint")
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
			Constraints: []validation.Constraint{{Target: "body.Series", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("anomalydetector.BaseClient", "DetectLastPoint", err.Error())
	}

	req, err := client.DetectLastPointPreparer(ctx, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "anomalydetector.BaseClient", "DetectLastPoint", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetectLastPointSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "anomalydetector.BaseClient", "DetectLastPoint", resp, "Failure sending request")
		return
	}

	result, err = client.DetectLastPointResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "anomalydetector.BaseClient", "DetectLastPoint", resp, "Failure responding to request")
	}

	return
}

// DetectLastPointPreparer prepares the DetectLastPoint request.
func (client BaseClient) DetectLastPointPreparer(ctx context.Context, body DetectRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/anomalydetector/v1.0", urlParameters),
		autorest.WithPath("/timeseries/last/detect"),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DetectLastPointSender sends the DetectLastPoint request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) DetectLastPointSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DetectLastPointResponder handles the response to the DetectLastPoint request. The method always
// closes the http.Response Body.
func (client BaseClient) DetectLastPointResponder(resp *http.Response) (result LastDetectResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
