package reservations

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

// ExchangeClient is the client for the Exchange methods of the Reservations service.
type ExchangeClient struct {
	BaseClient
}

// NewExchangeClient creates an instance of the ExchangeClient client.
func NewExchangeClient() ExchangeClient {
	return NewExchangeClientWithBaseURI(DefaultBaseURI)
}

// NewExchangeClientWithBaseURI creates an instance of the ExchangeClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewExchangeClientWithBaseURI(baseURI string) ExchangeClient {
	return ExchangeClient{NewWithBaseURI(baseURI)}
}

// Post returns one or more `Reservations` in exchange for one or more `Reservation` purchases.
// Parameters:
// body - request containing the refunds and purchases that need to be executed.
func (client ExchangeClient) Post(ctx context.Context, body ExchangeRequest) (result ExchangePostFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExchangeClient.Post")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PostPreparer(ctx, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.ExchangeClient", "Post", nil, "Failure preparing request")
		return
	}

	result, err = client.PostSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.ExchangeClient", "Post", result.Response(), "Failure sending request")
		return
	}

	return
}

// PostPreparer prepares the Post request.
func (client ExchangeClient) PostPreparer(ctx context.Context, body ExchangeRequest) (*http.Request, error) {
	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Capacity/exchange"),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PostSender sends the Post request. The method will close the
// http.Response Body if it receives an error.
func (client ExchangeClient) PostSender(req *http.Request) (future ExchangePostFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// PostResponder handles the response to the Post request. The method always
// closes the http.Response Body.
func (client ExchangeClient) PostResponder(resp *http.Response) (result ExchangeOperationResultResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
