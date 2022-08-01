//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappconfiguration

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ConfigurationStoresClientCreatePoller provides polling facilities until the operation reaches a terminal state.
type ConfigurationStoresClientCreatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ConfigurationStoresClientCreatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ConfigurationStoresClientCreatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ConfigurationStoresClientCreateResponse will be returned.
func (p *ConfigurationStoresClientCreatePoller) FinalResponse(ctx context.Context) (ConfigurationStoresClientCreateResponse, error) {
	respType := ConfigurationStoresClientCreateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ConfigurationStore)
	if err != nil {
		return ConfigurationStoresClientCreateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ConfigurationStoresClientCreatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ConfigurationStoresClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type ConfigurationStoresClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ConfigurationStoresClientDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ConfigurationStoresClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ConfigurationStoresClientDeleteResponse will be returned.
func (p *ConfigurationStoresClientDeletePoller) FinalResponse(ctx context.Context) (ConfigurationStoresClientDeleteResponse, error) {
	respType := ConfigurationStoresClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ConfigurationStoresClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ConfigurationStoresClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ConfigurationStoresClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ConfigurationStoresClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ConfigurationStoresClientUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ConfigurationStoresClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ConfigurationStoresClientUpdateResponse will be returned.
func (p *ConfigurationStoresClientUpdatePoller) FinalResponse(ctx context.Context) (ConfigurationStoresClientUpdateResponse, error) {
	respType := ConfigurationStoresClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ConfigurationStore)
	if err != nil {
		return ConfigurationStoresClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ConfigurationStoresClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
