package confluent

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/confluent/mgmt/2021-03-01-preview/confluent"

// AgreementProperties terms properties for Marketplace and Confluent.
type AgreementProperties struct {
	// Publisher - Publisher identifier string.
	Publisher *string `json:"publisher,omitempty"`
	// Product - Product identifier string.
	Product *string `json:"product,omitempty"`
	// Plan - Plan identifier string.
	Plan *string `json:"plan,omitempty"`
	// LicenseTextLink - Link to HTML with Microsoft and Publisher terms.
	LicenseTextLink *string `json:"licenseTextLink,omitempty"`
	// PrivacyPolicyLink - Link to the privacy policy of the publisher.
	PrivacyPolicyLink *string `json:"privacyPolicyLink,omitempty"`
	// RetrieveDatetime - Date and time in UTC of when the terms were accepted. This is empty if Accepted is false.
	RetrieveDatetime *date.Time `json:"retrieveDatetime,omitempty"`
	// Signature - Terms signature.
	Signature *string `json:"signature,omitempty"`
	// Accepted - If any version of the terms have been accepted, otherwise false.
	Accepted *bool `json:"accepted,omitempty"`
}

// AgreementResource agreement Terms definition
type AgreementResource struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The ARM id of the resource.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the agreement.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the agreement.
	Type *string `json:"type,omitempty"`
	// SystemData - READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty"`
	// AgreementProperties - Represents the properties of the resource.
	*AgreementProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for AgreementResource.
func (ar AgreementResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ar.AgreementProperties != nil {
		objectMap["properties"] = ar.AgreementProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for AgreementResource struct.
func (ar *AgreementResource) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ar.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ar.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ar.Type = &typeVar
			}
		case "systemData":
			if v != nil {
				var systemData SystemData
				err = json.Unmarshal(*v, &systemData)
				if err != nil {
					return err
				}
				ar.SystemData = &systemData
			}
		case "properties":
			if v != nil {
				var agreementProperties AgreementProperties
				err = json.Unmarshal(*v, &agreementProperties)
				if err != nil {
					return err
				}
				ar.AgreementProperties = &agreementProperties
			}
		}
	}

	return nil
}

// AgreementResourceListResponse response of a list operation.
type AgreementResourceListResponse struct {
	autorest.Response `json:"-"`
	// Value - Results of a list operation.
	Value *[]AgreementResource `json:"value,omitempty"`
	// NextLink - Link to the next set of results, if any.
	NextLink *string `json:"nextLink,omitempty"`
}

// AgreementResourceListResponseIterator provides access to a complete listing of AgreementResource values.
type AgreementResourceListResponseIterator struct {
	i    int
	page AgreementResourceListResponsePage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AgreementResourceListResponseIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AgreementResourceListResponseIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *AgreementResourceListResponseIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AgreementResourceListResponseIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter AgreementResourceListResponseIterator) Response() AgreementResourceListResponse {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AgreementResourceListResponseIterator) Value() AgreementResource {
	if !iter.page.NotDone() {
		return AgreementResource{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the AgreementResourceListResponseIterator type.
func NewAgreementResourceListResponseIterator(page AgreementResourceListResponsePage) AgreementResourceListResponseIterator {
	return AgreementResourceListResponseIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (arlr AgreementResourceListResponse) IsEmpty() bool {
	return arlr.Value == nil || len(*arlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (arlr AgreementResourceListResponse) hasNextLink() bool {
	return arlr.NextLink != nil && len(*arlr.NextLink) != 0
}

// agreementResourceListResponsePreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (arlr AgreementResourceListResponse) agreementResourceListResponsePreparer(ctx context.Context) (*http.Request, error) {
	if !arlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(arlr.NextLink)))
}

// AgreementResourceListResponsePage contains a page of AgreementResource values.
type AgreementResourceListResponsePage struct {
	fn   func(context.Context, AgreementResourceListResponse) (AgreementResourceListResponse, error)
	arlr AgreementResourceListResponse
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AgreementResourceListResponsePage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AgreementResourceListResponsePage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.arlr)
		if err != nil {
			return err
		}
		page.arlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *AgreementResourceListResponsePage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AgreementResourceListResponsePage) NotDone() bool {
	return !page.arlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page AgreementResourceListResponsePage) Response() AgreementResourceListResponse {
	return page.arlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page AgreementResourceListResponsePage) Values() []AgreementResource {
	if page.arlr.IsEmpty() {
		return nil
	}
	return *page.arlr.Value
}

// Creates a new instance of the AgreementResourceListResponsePage type.
func NewAgreementResourceListResponsePage(cur AgreementResourceListResponse, getNextPage func(context.Context, AgreementResourceListResponse) (AgreementResourceListResponse, error)) AgreementResourceListResponsePage {
	return AgreementResourceListResponsePage{
		fn:   getNextPage,
		arlr: cur,
	}
}

// ErrorResponseBody response body of Error
type ErrorResponseBody struct {
	// Code - READ-ONLY; Error code
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; Error message
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; Error target
	Target *string `json:"target,omitempty"`
	// Details - READ-ONLY; Error detail
	Details *[]ErrorResponseBody `json:"details,omitempty"`
}

// MarshalJSON is the custom marshaler for ErrorResponseBody.
func (erb ErrorResponseBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// OfferDetail confluent Offer detail
type OfferDetail struct {
	// PublisherID - Publisher Id
	PublisherID *string `json:"publisherId,omitempty"`
	// ID - Offer Id
	ID *string `json:"id,omitempty"`
	// PlanID - Offer Plan Id
	PlanID *string `json:"planId,omitempty"`
	// PlanName - Offer Plan Name
	PlanName *string `json:"planName,omitempty"`
	// TermUnit - Offer Plan Term unit
	TermUnit *string `json:"termUnit,omitempty"`
	// Status - SaaS Offer Status. Possible values include: 'SaaSOfferStatusStarted', 'SaaSOfferStatusPendingFulfillmentStart', 'SaaSOfferStatusInProgress', 'SaaSOfferStatusSubscribed', 'SaaSOfferStatusSuspended', 'SaaSOfferStatusReinstated', 'SaaSOfferStatusSucceeded', 'SaaSOfferStatusFailed', 'SaaSOfferStatusUnsubscribed', 'SaaSOfferStatusUpdating'
	Status SaaSOfferStatus `json:"status,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.Confluent
	Provider *string `json:"provider,omitempty"`
	// Resource - Type on which the operation is performed, e.g., 'clusters'.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type, e.g., read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of the operation, e.g., 'Write confluent'.
	Description *string `json:"description,omitempty"`
}

// OperationListResult result of GET request to list Confluent operations.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of Confluent operations supported by the Microsoft.Confluent provider.
	Value *[]OperationResult `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of OperationResult values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() OperationResult {
	if !iter.page.NotDone() {
		return OperationResult{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListResultIterator type.
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return OperationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (olr OperationListResult) hasNextLink() bool {
	return olr.NextLink != nil && len(*olr.NextLink) != 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !olr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of OperationResult values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.olr)
		if err != nil {
			return err
		}
		page.olr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []OperationResult {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Creates a new instance of the OperationListResultPage type.
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{
		fn:  getNextPage,
		olr: cur,
	}
}

// OperationResult an Confluent REST API operation.
type OperationResult struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
	// IsDataAction - Indicates whether the operation is a data action
	IsDataAction *bool `json:"isDataAction,omitempty"`
}

// OrganizationCreateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type OrganizationCreateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(OrganizationClient) (OrganizationResource, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *OrganizationCreateFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for OrganizationCreateFuture.Result.
func (future *OrganizationCreateFuture) result(client OrganizationClient) (or OrganizationResource, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confluent.OrganizationCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		or.Response.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("confluent.OrganizationCreateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if or.Response.Response, err = future.GetResult(sender); err == nil && or.Response.Response.StatusCode != http.StatusNoContent {
		or, err = client.CreateResponder(or.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "confluent.OrganizationCreateFuture", "Result", or.Response.Response, "Failure responding to request")
		}
	}
	return
}

// OrganizationDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type OrganizationDeleteFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(OrganizationClient) (autorest.Response, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *OrganizationDeleteFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for OrganizationDeleteFuture.Result.
func (future *OrganizationDeleteFuture) result(client OrganizationClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "confluent.OrganizationDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		ar.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("confluent.OrganizationDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// OrganizationResource organization resource.
type OrganizationResource struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The ARM id of the resource.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
	// SystemData - READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty"`
	// OrganizationResourceProperties - Organization resource properties
	*OrganizationResourceProperties `json:"properties,omitempty"`
	// Tags - Organization resource tags
	Tags map[string]*string `json:"tags"`
	// Location - Location of Organization resource
	Location *string `json:"location,omitempty"`
}

// MarshalJSON is the custom marshaler for OrganizationResource.
func (or OrganizationResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if or.OrganizationResourceProperties != nil {
		objectMap["properties"] = or.OrganizationResourceProperties
	}
	if or.Tags != nil {
		objectMap["tags"] = or.Tags
	}
	if or.Location != nil {
		objectMap["location"] = or.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for OrganizationResource struct.
func (or *OrganizationResource) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				or.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				or.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				or.Type = &typeVar
			}
		case "systemData":
			if v != nil {
				var systemData SystemData
				err = json.Unmarshal(*v, &systemData)
				if err != nil {
					return err
				}
				or.SystemData = &systemData
			}
		case "properties":
			if v != nil {
				var organizationResourceProperties OrganizationResourceProperties
				err = json.Unmarshal(*v, &organizationResourceProperties)
				if err != nil {
					return err
				}
				or.OrganizationResourceProperties = &organizationResourceProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				or.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				or.Location = &location
			}
		}
	}

	return nil
}

// OrganizationResourceListResult the response of a list operation.
type OrganizationResourceListResult struct {
	autorest.Response `json:"-"`
	// Value - Result of a list operation.
	Value *[]OrganizationResource `json:"value,omitempty"`
	// NextLink - Link to the next set of results, if any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OrganizationResourceListResultIterator provides access to a complete listing of OrganizationResource
// values.
type OrganizationResourceListResultIterator struct {
	i    int
	page OrganizationResourceListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OrganizationResourceListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrganizationResourceListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OrganizationResourceListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OrganizationResourceListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OrganizationResourceListResultIterator) Response() OrganizationResourceListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OrganizationResourceListResultIterator) Value() OrganizationResource {
	if !iter.page.NotDone() {
		return OrganizationResource{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OrganizationResourceListResultIterator type.
func NewOrganizationResourceListResultIterator(page OrganizationResourceListResultPage) OrganizationResourceListResultIterator {
	return OrganizationResourceListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (orlr OrganizationResourceListResult) IsEmpty() bool {
	return orlr.Value == nil || len(*orlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (orlr OrganizationResourceListResult) hasNextLink() bool {
	return orlr.NextLink != nil && len(*orlr.NextLink) != 0
}

// organizationResourceListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (orlr OrganizationResourceListResult) organizationResourceListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !orlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(orlr.NextLink)))
}

// OrganizationResourceListResultPage contains a page of OrganizationResource values.
type OrganizationResourceListResultPage struct {
	fn   func(context.Context, OrganizationResourceListResult) (OrganizationResourceListResult, error)
	orlr OrganizationResourceListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OrganizationResourceListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrganizationResourceListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.orlr)
		if err != nil {
			return err
		}
		page.orlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OrganizationResourceListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OrganizationResourceListResultPage) NotDone() bool {
	return !page.orlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OrganizationResourceListResultPage) Response() OrganizationResourceListResult {
	return page.orlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OrganizationResourceListResultPage) Values() []OrganizationResource {
	if page.orlr.IsEmpty() {
		return nil
	}
	return *page.orlr.Value
}

// Creates a new instance of the OrganizationResourceListResultPage type.
func NewOrganizationResourceListResultPage(cur OrganizationResourceListResult, getNextPage func(context.Context, OrganizationResourceListResult) (OrganizationResourceListResult, error)) OrganizationResourceListResultPage {
	return OrganizationResourceListResultPage{
		fn:   getNextPage,
		orlr: cur,
	}
}

// OrganizationResourceProperties organization resource property
type OrganizationResourceProperties struct {
	// CreatedTime - READ-ONLY; The creation time of the resource.
	CreatedTime *date.Time `json:"createdTime,omitempty"`
	// ProvisioningState - READ-ONLY; Provision states for confluent RP. Possible values include: 'Accepted', 'Creating', 'Updating', 'Deleting', 'Succeeded', 'Failed', 'Canceled', 'Deleted', 'NotSpecified'
	ProvisioningState ProvisionState `json:"provisioningState,omitempty"`
	// OrganizationID - READ-ONLY; Id of the Confluent organization.
	OrganizationID *string `json:"organizationId,omitempty"`
	// SsoURL - READ-ONLY; SSO url for the Confluent organization.
	SsoURL *string `json:"ssoUrl,omitempty"`
	// OfferDetail - Confluent offer detail
	OfferDetail *OfferDetail `json:"offerDetail,omitempty"`
	// UserDetail - Subscriber detail
	UserDetail *UserDetail `json:"userDetail,omitempty"`
}

// MarshalJSON is the custom marshaler for OrganizationResourceProperties.
func (orp OrganizationResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if orp.OfferDetail != nil {
		objectMap["offerDetail"] = orp.OfferDetail
	}
	if orp.UserDetail != nil {
		objectMap["userDetail"] = orp.UserDetail
	}
	return json.Marshal(objectMap)
}

// OrganizationResourceUpdate organization Resource update
type OrganizationResourceUpdate struct {
	// Tags - ARM resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for OrganizationResourceUpdate.
func (oru OrganizationResourceUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if oru.Tags != nil {
		objectMap["tags"] = oru.Tags
	}
	return json.Marshal(objectMap)
}

// ResourceProviderDefaultErrorResponse default error response for resource provider
type ResourceProviderDefaultErrorResponse struct {
	// Error - READ-ONLY; Response body of Error
	Error *ErrorResponseBody `json:"error,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceProviderDefaultErrorResponse.
func (rpder ResourceProviderDefaultErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// SystemData metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// CreatedBy - The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedByType - The type of identity that created the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	CreatedByType CreatedByType `json:"createdByType,omitempty"`
	// CreatedAt - The timestamp of resource creation (UTC).
	CreatedAt *date.Time `json:"createdAt,omitempty"`
	// LastModifiedBy - The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// LastModifiedByType - The type of identity that last modified the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	LastModifiedByType CreatedByType `json:"lastModifiedByType,omitempty"`
	// LastModifiedAt - The timestamp of resource last modification (UTC)
	LastModifiedAt *date.Time `json:"lastModifiedAt,omitempty"`
}

// UserDetail subscriber detail
type UserDetail struct {
	// FirstName - First name
	FirstName *string `json:"firstName,omitempty"`
	// LastName - Last name
	LastName *string `json:"lastName,omitempty"`
	// EmailAddress - Email address
	EmailAddress *string `json:"emailAddress,omitempty"`
}
