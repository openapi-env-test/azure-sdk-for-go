//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// AccessReviewHistoryDefinitionInstancesClientListPager provides operations for iterating over paged responses.
type AccessReviewHistoryDefinitionInstancesClientListPager struct {
	client    *AccessReviewHistoryDefinitionInstancesClient
	current   AccessReviewHistoryDefinitionInstancesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccessReviewHistoryDefinitionInstancesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccessReviewHistoryDefinitionInstancesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccessReviewHistoryDefinitionInstancesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccessReviewHistoryDefinitionInstanceListResult.NextLink == nil || len(*p.current.AccessReviewHistoryDefinitionInstanceListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccessReviewHistoryDefinitionInstancesClientListResponse page.
func (p *AccessReviewHistoryDefinitionInstancesClientListPager) PageResponse() AccessReviewHistoryDefinitionInstancesClientListResponse {
	return p.current
}

// AccessReviewHistoryDefinitionsClientListPager provides operations for iterating over paged responses.
type AccessReviewHistoryDefinitionsClientListPager struct {
	client    *AccessReviewHistoryDefinitionsClient
	current   AccessReviewHistoryDefinitionsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccessReviewHistoryDefinitionsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccessReviewHistoryDefinitionsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccessReviewHistoryDefinitionsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccessReviewHistoryDefinitionListResult.NextLink == nil || len(*p.current.AccessReviewHistoryDefinitionListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccessReviewHistoryDefinitionsClientListResponse page.
func (p *AccessReviewHistoryDefinitionsClientListPager) PageResponse() AccessReviewHistoryDefinitionsClientListResponse {
	return p.current
}

// AccessReviewInstanceContactedReviewersClientListPager provides operations for iterating over paged responses.
type AccessReviewInstanceContactedReviewersClientListPager struct {
	client    *AccessReviewInstanceContactedReviewersClient
	current   AccessReviewInstanceContactedReviewersClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccessReviewInstanceContactedReviewersClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccessReviewInstanceContactedReviewersClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccessReviewInstanceContactedReviewersClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccessReviewContactedReviewerListResult.NextLink == nil || len(*p.current.AccessReviewContactedReviewerListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccessReviewInstanceContactedReviewersClientListResponse page.
func (p *AccessReviewInstanceContactedReviewersClientListPager) PageResponse() AccessReviewInstanceContactedReviewersClientListResponse {
	return p.current
}

// AccessReviewInstanceDecisionsClientListPager provides operations for iterating over paged responses.
type AccessReviewInstanceDecisionsClientListPager struct {
	client    *AccessReviewInstanceDecisionsClient
	current   AccessReviewInstanceDecisionsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccessReviewInstanceDecisionsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccessReviewInstanceDecisionsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccessReviewInstanceDecisionsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccessReviewDecisionListResult.NextLink == nil || len(*p.current.AccessReviewDecisionListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccessReviewInstanceDecisionsClientListResponse page.
func (p *AccessReviewInstanceDecisionsClientListPager) PageResponse() AccessReviewInstanceDecisionsClientListResponse {
	return p.current
}

// AccessReviewInstanceMyDecisionsClientListPager provides operations for iterating over paged responses.
type AccessReviewInstanceMyDecisionsClientListPager struct {
	client    *AccessReviewInstanceMyDecisionsClient
	current   AccessReviewInstanceMyDecisionsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccessReviewInstanceMyDecisionsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccessReviewInstanceMyDecisionsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccessReviewInstanceMyDecisionsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccessReviewDecisionListResult.NextLink == nil || len(*p.current.AccessReviewDecisionListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccessReviewInstanceMyDecisionsClientListResponse page.
func (p *AccessReviewInstanceMyDecisionsClientListPager) PageResponse() AccessReviewInstanceMyDecisionsClientListResponse {
	return p.current
}

// AccessReviewInstancesAssignedForMyApprovalClientListPager provides operations for iterating over paged responses.
type AccessReviewInstancesAssignedForMyApprovalClientListPager struct {
	client    *AccessReviewInstancesAssignedForMyApprovalClient
	current   AccessReviewInstancesAssignedForMyApprovalClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccessReviewInstancesAssignedForMyApprovalClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccessReviewInstancesAssignedForMyApprovalClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccessReviewInstancesAssignedForMyApprovalClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccessReviewInstanceListResult.NextLink == nil || len(*p.current.AccessReviewInstanceListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccessReviewInstancesAssignedForMyApprovalClientListResponse page.
func (p *AccessReviewInstancesAssignedForMyApprovalClientListPager) PageResponse() AccessReviewInstancesAssignedForMyApprovalClientListResponse {
	return p.current
}

// AccessReviewInstancesClientListPager provides operations for iterating over paged responses.
type AccessReviewInstancesClientListPager struct {
	client    *AccessReviewInstancesClient
	current   AccessReviewInstancesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccessReviewInstancesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccessReviewInstancesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccessReviewInstancesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccessReviewInstanceListResult.NextLink == nil || len(*p.current.AccessReviewInstanceListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccessReviewInstancesClientListResponse page.
func (p *AccessReviewInstancesClientListPager) PageResponse() AccessReviewInstancesClientListResponse {
	return p.current
}

// AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListPager provides operations for iterating over paged responses.
type AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListPager struct {
	client    *AccessReviewScheduleDefinitionsAssignedForMyApprovalClient
	current   AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccessReviewScheduleDefinitionListResult.NextLink == nil || len(*p.current.AccessReviewScheduleDefinitionListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse page.
func (p *AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListPager) PageResponse() AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse {
	return p.current
}

// AccessReviewScheduleDefinitionsClientListPager provides operations for iterating over paged responses.
type AccessReviewScheduleDefinitionsClientListPager struct {
	client    *AccessReviewScheduleDefinitionsClient
	current   AccessReviewScheduleDefinitionsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccessReviewScheduleDefinitionsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccessReviewScheduleDefinitionsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccessReviewScheduleDefinitionsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccessReviewScheduleDefinitionListResult.NextLink == nil || len(*p.current.AccessReviewScheduleDefinitionListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccessReviewScheduleDefinitionsClientListResponse page.
func (p *AccessReviewScheduleDefinitionsClientListPager) PageResponse() AccessReviewScheduleDefinitionsClientListResponse {
	return p.current
}

// OperationsClientListPager provides operations for iterating over paged responses.
type OperationsClientListPager struct {
	client    *OperationsClient
	current   OperationsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsClientListResponse page.
func (p *OperationsClientListPager) PageResponse() OperationsClientListResponse {
	return p.current
}

// TenantLevelAccessReviewInstanceContactedReviewersClientListPager provides operations for iterating over paged responses.
type TenantLevelAccessReviewInstanceContactedReviewersClientListPager struct {
	client    *TenantLevelAccessReviewInstanceContactedReviewersClient
	current   TenantLevelAccessReviewInstanceContactedReviewersClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, TenantLevelAccessReviewInstanceContactedReviewersClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *TenantLevelAccessReviewInstanceContactedReviewersClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *TenantLevelAccessReviewInstanceContactedReviewersClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccessReviewContactedReviewerListResult.NextLink == nil || len(*p.current.AccessReviewContactedReviewerListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current TenantLevelAccessReviewInstanceContactedReviewersClientListResponse page.
func (p *TenantLevelAccessReviewInstanceContactedReviewersClientListPager) PageResponse() TenantLevelAccessReviewInstanceContactedReviewersClientListResponse {
	return p.current
}
