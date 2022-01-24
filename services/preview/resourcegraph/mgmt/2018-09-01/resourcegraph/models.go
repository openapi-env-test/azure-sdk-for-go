package resourcegraph

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/resourcegraph/mgmt/2018-09-01/resourcegraph"

// Column query result column descriptor.
type Column struct {
	// Name - Column name.
	Name *string `json:"name,omitempty"`
	// Type - Column data type. Possible values include: 'String', 'Integer', 'Number', 'Boolean', 'Object'
	Type ColumnDataType `json:"type,omitempty"`
}

// DateTimeInterval an interval in time specifying the date and time for the inclusive start and exclusive
// end, i.e. `[start, end)`.
type DateTimeInterval struct {
	// Start - A datetime indicating the inclusive/closed start of the time interval, i.e. `[`**`start`**`, end)`. Specifying a `start` that occurs chronologically after `end` will result in an error.
	Start *date.Time `json:"start,omitempty"`
	// End - A datetime indicating the exclusive/open end of the time interval, i.e. `[start, `**`end`**`)`. Specifying an `end` that occurs chronologically before `start` will result in an error.
	End *date.Time `json:"end,omitempty"`
}

// Error error details.
type Error struct {
	// Code - Error code identifying the specific error.
	Code *string `json:"code,omitempty"`
	// Message - A human readable error message.
	Message *string `json:"message,omitempty"`
	// Details - Error details
	Details *[]ErrorDetails `json:"details,omitempty"`
}

// ErrorDetails ...
type ErrorDetails struct {
	// AdditionalProperties - Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]interface{} `json:""`
	// Code - Error code identifying the specific error.
	Code *string `json:"code,omitempty"`
	// Message - A human readable error message.
	Message *string `json:"message,omitempty"`
}

// MarshalJSON is the custom marshaler for ErrorDetails.
func (ed ErrorDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ed.Code != nil {
		objectMap["code"] = ed.Code
	}
	if ed.Message != nil {
		objectMap["message"] = ed.Message
	}
	for k, v := range ed.AdditionalProperties {
		objectMap[k] = v
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ErrorDetails struct.
func (ed *ErrorDetails) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		default:
			if v != nil {
				var additionalProperties interface{}
				err = json.Unmarshal(*v, &additionalProperties)
				if err != nil {
					return err
				}
				if ed.AdditionalProperties == nil {
					ed.AdditionalProperties = make(map[string]interface{})
				}
				ed.AdditionalProperties[k] = additionalProperties
			}
		case "code":
			if v != nil {
				var code string
				err = json.Unmarshal(*v, &code)
				if err != nil {
					return err
				}
				ed.Code = &code
			}
		case "message":
			if v != nil {
				var message string
				err = json.Unmarshal(*v, &message)
				if err != nil {
					return err
				}
				ed.Message = &message
			}
		}
	}

	return nil
}

// ErrorFieldContract error Field contract.
type ErrorFieldContract struct {
	// Code - Property level error code.
	Code *string `json:"code,omitempty"`
	// Message - Human-readable representation of property-level error.
	Message *string `json:"message,omitempty"`
	// Target - Property name.
	Target *string `json:"target,omitempty"`
}

// ErrorResponse an error response from the API.
type ErrorResponse struct {
	// Error - Error information.
	Error *Error `json:"error,omitempty"`
}

// BasicFacet a facet containing additional statistics on the response of a query. Can be either FacetResult or
// FacetError.
type BasicFacet interface {
	AsFacetResult() (*FacetResult, bool)
	AsFacetError() (*FacetError, bool)
	AsFacet() (*Facet, bool)
}

// Facet a facet containing additional statistics on the response of a query. Can be either FacetResult or
// FacetError.
type Facet struct {
	// Expression - Facet expression, same as in the corresponding facet request.
	Expression *string `json:"expression,omitempty"`
	// ResultType - Possible values include: 'ResultTypeFacet', 'ResultTypeFacetResult', 'ResultTypeFacetError'
	ResultType ResultType `json:"resultType,omitempty"`
}

func unmarshalBasicFacet(body []byte) (BasicFacet, error) {
	var m map[string]interface{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	switch m["resultType"] {
	case string(ResultTypeFacetResult):
		var fr FacetResult
		err := json.Unmarshal(body, &fr)
		return fr, err
	case string(ResultTypeFacetError):
		var fe FacetError
		err := json.Unmarshal(body, &fe)
		return fe, err
	default:
		var f Facet
		err := json.Unmarshal(body, &f)
		return f, err
	}
}
func unmarshalBasicFacetArray(body []byte) ([]BasicFacet, error) {
	var rawMessages []*json.RawMessage
	err := json.Unmarshal(body, &rawMessages)
	if err != nil {
		return nil, err
	}

	fArray := make([]BasicFacet, len(rawMessages))

	for index, rawMessage := range rawMessages {
		f, err := unmarshalBasicFacet(*rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

// MarshalJSON is the custom marshaler for Facet.
func (f Facet) MarshalJSON() ([]byte, error) {
	f.ResultType = ResultTypeFacet
	objectMap := make(map[string]interface{})
	if f.Expression != nil {
		objectMap["expression"] = f.Expression
	}
	if f.ResultType != "" {
		objectMap["resultType"] = f.ResultType
	}
	return json.Marshal(objectMap)
}

// AsFacetResult is the BasicFacet implementation for Facet.
func (f Facet) AsFacetResult() (*FacetResult, bool) {
	return nil, false
}

// AsFacetError is the BasicFacet implementation for Facet.
func (f Facet) AsFacetError() (*FacetError, bool) {
	return nil, false
}

// AsFacet is the BasicFacet implementation for Facet.
func (f Facet) AsFacet() (*Facet, bool) {
	return &f, true
}

// AsBasicFacet is the BasicFacet implementation for Facet.
func (f Facet) AsBasicFacet() (BasicFacet, bool) {
	return &f, true
}

// FacetError a facet whose execution resulted in an error.
type FacetError struct {
	// Errors - An array containing detected facet errors with details.
	Errors *[]ErrorDetails `json:"errors,omitempty"`
	// Expression - Facet expression, same as in the corresponding facet request.
	Expression *string `json:"expression,omitempty"`
	// ResultType - Possible values include: 'ResultTypeFacet', 'ResultTypeFacetResult', 'ResultTypeFacetError'
	ResultType ResultType `json:"resultType,omitempty"`
}

// MarshalJSON is the custom marshaler for FacetError.
func (fe FacetError) MarshalJSON() ([]byte, error) {
	fe.ResultType = ResultTypeFacetError
	objectMap := make(map[string]interface{})
	if fe.Errors != nil {
		objectMap["errors"] = fe.Errors
	}
	if fe.Expression != nil {
		objectMap["expression"] = fe.Expression
	}
	if fe.ResultType != "" {
		objectMap["resultType"] = fe.ResultType
	}
	return json.Marshal(objectMap)
}

// AsFacetResult is the BasicFacet implementation for FacetError.
func (fe FacetError) AsFacetResult() (*FacetResult, bool) {
	return nil, false
}

// AsFacetError is the BasicFacet implementation for FacetError.
func (fe FacetError) AsFacetError() (*FacetError, bool) {
	return &fe, true
}

// AsFacet is the BasicFacet implementation for FacetError.
func (fe FacetError) AsFacet() (*Facet, bool) {
	return nil, false
}

// AsBasicFacet is the BasicFacet implementation for FacetError.
func (fe FacetError) AsBasicFacet() (BasicFacet, bool) {
	return &fe, true
}

// FacetRequest a request to compute additional statistics (facets) over the query results.
type FacetRequest struct {
	// Expression - The column or list of columns to summarize by
	Expression *string `json:"expression,omitempty"`
	// Options - The options for facet evaluation
	Options *FacetRequestOptions `json:"options,omitempty"`
}

// FacetRequestOptions the options for facet evaluation
type FacetRequestOptions struct {
	// SortBy - The column name or query expression to sort on. Defaults to count if not present.
	SortBy *string `json:"sortBy,omitempty"`
	// SortOrder - The sorting order by the selected column (count by default). Possible values include: 'Asc', 'Desc'
	SortOrder FacetSortOrder `json:"sortOrder,omitempty"`
	// Filter - Specifies the filter condition for the 'where' clause which will be run on main query's result, just before the actual faceting.
	Filter *string `json:"filter,omitempty"`
	// Top - The maximum number of facet rows that should be returned.
	Top *int32 `json:"$top,omitempty"`
}

// FacetResult successfully executed facet containing additional statistics on the response of a query.
type FacetResult struct {
	// TotalRecords - Number of total records in the facet results.
	TotalRecords *int64 `json:"totalRecords,omitempty"`
	// Count - Number of records returned in the facet response.
	Count *int32 `json:"count,omitempty"`
	// Data - A table containing the desired facets. Only present if the facet is valid.
	Data *Table `json:"data,omitempty"`
	// Expression - Facet expression, same as in the corresponding facet request.
	Expression *string `json:"expression,omitempty"`
	// ResultType - Possible values include: 'ResultTypeFacet', 'ResultTypeFacetResult', 'ResultTypeFacetError'
	ResultType ResultType `json:"resultType,omitempty"`
}

// MarshalJSON is the custom marshaler for FacetResult.
func (fr FacetResult) MarshalJSON() ([]byte, error) {
	fr.ResultType = ResultTypeFacetResult
	objectMap := make(map[string]interface{})
	if fr.TotalRecords != nil {
		objectMap["totalRecords"] = fr.TotalRecords
	}
	if fr.Count != nil {
		objectMap["count"] = fr.Count
	}
	if fr.Data != nil {
		objectMap["data"] = fr.Data
	}
	if fr.Expression != nil {
		objectMap["expression"] = fr.Expression
	}
	if fr.ResultType != "" {
		objectMap["resultType"] = fr.ResultType
	}
	return json.Marshal(objectMap)
}

// AsFacetResult is the BasicFacet implementation for FacetResult.
func (fr FacetResult) AsFacetResult() (*FacetResult, bool) {
	return &fr, true
}

// AsFacetError is the BasicFacet implementation for FacetResult.
func (fr FacetResult) AsFacetError() (*FacetError, bool) {
	return nil, false
}

// AsFacet is the BasicFacet implementation for FacetResult.
func (fr FacetResult) AsFacet() (*Facet, bool) {
	return nil, false
}

// AsBasicFacet is the BasicFacet implementation for FacetResult.
func (fr FacetResult) AsBasicFacet() (BasicFacet, bool) {
	return &fr, true
}

// GraphQueryError error message body that will indicate why the operation failed.
type GraphQueryError struct {
	// Code - Service-defined error code. This code serves as a sub-status for the HTTP error code specified in the response.
	Code *string `json:"code,omitempty"`
	// Message - Human-readable representation of the error.
	Message *string `json:"message,omitempty"`
	// Details - The list of invalid fields send in request, in case of validation error.
	Details *[]ErrorFieldContract `json:"details,omitempty"`
}

// GraphQueryListResult graph query list result.
type GraphQueryListResult struct {
	autorest.Response `json:"-"`
	// NextLink - URL to fetch the next set of queries.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - READ-ONLY; An array of graph queries.
	Value *[]GraphQueryResource `json:"value,omitempty"`
}

// MarshalJSON is the custom marshaler for GraphQueryListResult.
func (gqlr GraphQueryListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if gqlr.NextLink != nil {
		objectMap["nextLink"] = gqlr.NextLink
	}
	return json.Marshal(objectMap)
}

// GraphQueryListResultIterator provides access to a complete listing of GraphQueryResource values.
type GraphQueryListResultIterator struct {
	i    int
	page GraphQueryListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *GraphQueryListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GraphQueryListResultIterator.NextWithContext")
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
func (iter *GraphQueryListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter GraphQueryListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter GraphQueryListResultIterator) Response() GraphQueryListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter GraphQueryListResultIterator) Value() GraphQueryResource {
	if !iter.page.NotDone() {
		return GraphQueryResource{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the GraphQueryListResultIterator type.
func NewGraphQueryListResultIterator(page GraphQueryListResultPage) GraphQueryListResultIterator {
	return GraphQueryListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (gqlr GraphQueryListResult) IsEmpty() bool {
	return gqlr.Value == nil || len(*gqlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (gqlr GraphQueryListResult) hasNextLink() bool {
	return gqlr.NextLink != nil && len(*gqlr.NextLink) != 0
}

// graphQueryListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (gqlr GraphQueryListResult) graphQueryListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !gqlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(gqlr.NextLink)))
}

// GraphQueryListResultPage contains a page of GraphQueryResource values.
type GraphQueryListResultPage struct {
	fn   func(context.Context, GraphQueryListResult) (GraphQueryListResult, error)
	gqlr GraphQueryListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *GraphQueryListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GraphQueryListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.gqlr)
		if err != nil {
			return err
		}
		page.gqlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *GraphQueryListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page GraphQueryListResultPage) NotDone() bool {
	return !page.gqlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page GraphQueryListResultPage) Response() GraphQueryListResult {
	return page.gqlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page GraphQueryListResultPage) Values() []GraphQueryResource {
	if page.gqlr.IsEmpty() {
		return nil
	}
	return *page.gqlr.Value
}

// Creates a new instance of the GraphQueryListResultPage type.
func NewGraphQueryListResultPage(cur GraphQueryListResult, getNextPage func(context.Context, GraphQueryListResult) (GraphQueryListResult, error)) GraphQueryListResultPage {
	return GraphQueryListResultPage{
		fn:   getNextPage,
		gqlr: cur,
	}
}

// GraphQueryProperties properties that contain a graph query.
type GraphQueryProperties struct {
	// TimeModified - READ-ONLY; Date and time in UTC of the last modification that was made to this graph query definition.
	TimeModified *date.Time `json:"timeModified,omitempty"`
	// Description - The description of a graph query.
	Description *string `json:"description,omitempty"`
	// Query - KQL query that will be graph.
	Query *string `json:"query,omitempty"`
	// ResultKind - READ-ONLY; Enum indicating a type of graph query. Possible values include: 'Basic'
	ResultKind ResultKind `json:"resultKind,omitempty"`
}

// MarshalJSON is the custom marshaler for GraphQueryProperties.
func (gqp GraphQueryProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if gqp.Description != nil {
		objectMap["description"] = gqp.Description
	}
	if gqp.Query != nil {
		objectMap["query"] = gqp.Query
	}
	return json.Marshal(objectMap)
}

// GraphQueryPropertiesUpdateParameters properties that contain a workbook for PATCH operation.
type GraphQueryPropertiesUpdateParameters struct {
	// Description - The description of a graph query.
	Description *string `json:"description,omitempty"`
	// Query - KQL query that will be graph.
	Query *string `json:"query,omitempty"`
}

// GraphQueryResource graph Query entity definition.
type GraphQueryResource struct {
	autorest.Response `json:"-"`
	// GraphQueryProperties - Metadata describing a graph query for an Azure resource.
	*GraphQueryProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Azure resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Azure resource name. This is GUID value. The display name should be assigned within properties field.
	Name *string `json:"name,omitempty"`
	// Location - The location of the resource
	Location *string `json:"location,omitempty"`
	// Type - READ-ONLY; Azure resource type
	Type *string `json:"type,omitempty"`
	// Etag - This will be used to handle Optimistic Concurrency. If not present, it will always overwrite the existing resource without checking conflict.
	Etag *string `json:"etag,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for GraphQueryResource.
func (gqr GraphQueryResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if gqr.GraphQueryProperties != nil {
		objectMap["properties"] = gqr.GraphQueryProperties
	}
	if gqr.Location != nil {
		objectMap["location"] = gqr.Location
	}
	if gqr.Etag != nil {
		objectMap["etag"] = gqr.Etag
	}
	if gqr.Tags != nil {
		objectMap["tags"] = gqr.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for GraphQueryResource struct.
func (gqr *GraphQueryResource) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var graphQueryProperties GraphQueryProperties
				err = json.Unmarshal(*v, &graphQueryProperties)
				if err != nil {
					return err
				}
				gqr.GraphQueryProperties = &graphQueryProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				gqr.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				gqr.Name = &name
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				gqr.Location = &location
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				gqr.Type = &typeVar
			}
		case "etag":
			if v != nil {
				var etag string
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				gqr.Etag = &etag
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				gqr.Tags = tags
			}
		}
	}

	return nil
}

// GraphQueryUpdateParameters the parameters that can be provided when updating workbook properties
// properties.
type GraphQueryUpdateParameters struct {
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
	// Etag - This will be used to handle Optimistic Concurrency. If not present, it will always overwrite the existing resource without checking conflict.
	Etag *string `json:"etag,omitempty"`
	// GraphQueryPropertiesUpdateParameters - Metadata describing a graph query for an Azure resource.
	*GraphQueryPropertiesUpdateParameters `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for GraphQueryUpdateParameters.
func (gqup GraphQueryUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if gqup.Tags != nil {
		objectMap["tags"] = gqup.Tags
	}
	if gqup.Etag != nil {
		objectMap["etag"] = gqup.Etag
	}
	if gqup.GraphQueryPropertiesUpdateParameters != nil {
		objectMap["properties"] = gqup.GraphQueryPropertiesUpdateParameters
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for GraphQueryUpdateParameters struct.
func (gqup *GraphQueryUpdateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				gqup.Tags = tags
			}
		case "etag":
			if v != nil {
				var etag string
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				gqup.Etag = &etag
			}
		case "properties":
			if v != nil {
				var graphQueryPropertiesUpdateParameters GraphQueryPropertiesUpdateParameters
				err = json.Unmarshal(*v, &graphQueryPropertiesUpdateParameters)
				if err != nil {
					return err
				}
				gqup.GraphQueryPropertiesUpdateParameters = &graphQueryPropertiesUpdateParameters
			}
		}
	}

	return nil
}

// Operation resource Graph REST API operation definition.
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - Display metadata associated with the operation.
	Display *OperationDisplay `json:"display,omitempty"`
	// Origin - The origin of operations.
	Origin *string `json:"origin,omitempty"`
}

// OperationDisplay display metadata associated with the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft Resource Graph.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Type of operation: get, read, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description for the operation.
	Description *string `json:"description,omitempty"`
}

// OperationListResult result of the request to list Resource Graph operations. It contains a list of
// operations and a URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of Resource Graph operations supported by the Resource Graph resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - The link used to get the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
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
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
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

// OperationListResultPage contains a page of Operation values.
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
func (page OperationListResultPage) Values() []Operation {
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

// QueryRequest describes a query to be executed.
type QueryRequest struct {
	// Subscriptions - Azure subscriptions against which to execute the query.
	Subscriptions *[]string `json:"subscriptions,omitempty"`
	// Query - The resources query.
	Query *string `json:"query,omitempty"`
	// Options - The query evaluation options
	Options *QueryRequestOptions `json:"options,omitempty"`
	// Facets - An array of facet requests to be computed against the query result.
	Facets *[]FacetRequest `json:"facets,omitempty"`
}

// QueryRequestOptions the options for query evaluation
type QueryRequestOptions struct {
	// SkipToken - Continuation token for pagination, capturing the next page size and offset, as well as the context of the query.
	SkipToken *string `json:"$skipToken,omitempty"`
	// Top - The maximum number of rows that the query should return. Overrides the page size when ```$skipToken``` property is present.
	Top *int32 `json:"$top,omitempty"`
	// Skip - The number of rows to skip from the beginning of the results. Overrides the next page offset when ```$skipToken``` property is present.
	Skip *int32 `json:"$skip,omitempty"`
}

// QueryResponse query result.
type QueryResponse struct {
	autorest.Response `json:"-"`
	// TotalRecords - Number of total records matching the query.
	TotalRecords *int64 `json:"totalRecords,omitempty"`
	// Count - Number of records returned in the current response. In the case of paging, this is the number of records in the current page.
	Count *int64 `json:"count,omitempty"`
	// ResultTruncated - Indicates whether the query results are truncated. Possible values include: 'True', 'False'
	ResultTruncated ResultTruncated `json:"resultTruncated,omitempty"`
	// SkipToken - When present, the value can be passed to a subsequent query call (together with the same query and subscriptions used in the current request) to retrieve the next page of data.
	SkipToken *string `json:"$skipToken,omitempty"`
	// Data - Query output in tabular format.
	Data *Table `json:"data,omitempty"`
	// Facets - Query facets.
	Facets *[]BasicFacet `json:"facets,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for QueryResponse struct.
func (qr *QueryResponse) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "totalRecords":
			if v != nil {
				var totalRecords int64
				err = json.Unmarshal(*v, &totalRecords)
				if err != nil {
					return err
				}
				qr.TotalRecords = &totalRecords
			}
		case "count":
			if v != nil {
				var count int64
				err = json.Unmarshal(*v, &count)
				if err != nil {
					return err
				}
				qr.Count = &count
			}
		case "resultTruncated":
			if v != nil {
				var resultTruncated ResultTruncated
				err = json.Unmarshal(*v, &resultTruncated)
				if err != nil {
					return err
				}
				qr.ResultTruncated = resultTruncated
			}
		case "$skipToken":
			if v != nil {
				var skipToken string
				err = json.Unmarshal(*v, &skipToken)
				if err != nil {
					return err
				}
				qr.SkipToken = &skipToken
			}
		case "data":
			if v != nil {
				var data Table
				err = json.Unmarshal(*v, &data)
				if err != nil {
					return err
				}
				qr.Data = &data
			}
		case "facets":
			if v != nil {
				facets, err := unmarshalBasicFacetArray(*v)
				if err != nil {
					return err
				}
				qr.Facets = &facets
			}
		}
	}

	return nil
}

// Resource an azure resource object
type Resource struct {
	// ID - READ-ONLY; Azure resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Azure resource name. This is GUID value. The display name should be assigned within properties field.
	Name *string `json:"name,omitempty"`
	// Location - The location of the resource
	Location *string `json:"location,omitempty"`
	// Type - READ-ONLY; Azure resource type
	Type *string `json:"type,omitempty"`
	// Etag - This will be used to handle Optimistic Concurrency. If not present, it will always overwrite the existing resource without checking conflict.
	Etag *string `json:"etag,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Etag != nil {
		objectMap["etag"] = r.Etag
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// ResourceChangeData data on a specific change, represented by a pair of before and after resource
// snapshots.
type ResourceChangeData struct {
	autorest.Response `json:"-"`
	// ChangeID - The change ID. Valid and unique within the specified resource only.
	ChangeID *string `json:"changeId,omitempty"`
	// BeforeSnapshot - The snapshot before the change.
	BeforeSnapshot *ResourceChangeDataBeforeSnapshot `json:"beforeSnapshot,omitempty"`
	// AfterSnapshot - The snapshot after the change.
	AfterSnapshot *ResourceChangeDataAfterSnapshot `json:"afterSnapshot,omitempty"`
}

// ResourceChangeDataAfterSnapshot the snapshot after the change.
type ResourceChangeDataAfterSnapshot struct {
	// Timestamp - The time when the snapshot was created.
	// The snapshot timestamp provides an approximation as to when a modification to a resource was detected.  There can be a difference between the actual modification time and the detection time.  This is due to differences in how operations that modify a resource are processed, versus how operation that record resource snapshots are processed.
	Timestamp *date.Time `json:"timestamp,omitempty"`
	// Content - The resource snapshot content (in resourceChangeDetails response only).
	Content interface{} `json:"content,omitempty"`
}

// ResourceChangeDataBeforeSnapshot the snapshot before the change.
type ResourceChangeDataBeforeSnapshot struct {
	// Timestamp - The time when the snapshot was created.
	// The snapshot timestamp provides an approximation as to when a modification to a resource was detected.  There can be a difference between the actual modification time and the detection time.  This is due to differences in how operations that modify a resource are processed, versus how operation that record resource snapshots are processed.
	Timestamp *date.Time `json:"timestamp,omitempty"`
	// Content - The resource snapshot content (in resourceChangeDetails response only).
	Content interface{} `json:"content,omitempty"`
}

// ResourceChangeDetailsRequestParameters the parameters for a specific change details request.
type ResourceChangeDetailsRequestParameters struct {
	// ResourceID - Specifies the resource for a change details request.
	ResourceID *string `json:"resourceId,omitempty"`
	// ChangeID - Specifies the change ID.
	ChangeID *string `json:"changeId,omitempty"`
}

// ResourceChangeList a list of changes associated with a resource over a specific time interval.
type ResourceChangeList struct {
	autorest.Response `json:"-"`
	// Changes - The pageable value returned by the operation, i.e. a list of changes to the resource.
	// - The list is ordered from the most recent changes to the least recent changes.
	// - This list will be empty if there were no changes during the requested interval.
	// - The `Before` snapshot timestamp value of the oldest change can be outside of the specified time interval.
	Changes *[]ResourceChangeData `json:"changes,omitempty"`
	// SkipToken - Skip token that encodes the skip information while executing the current request
	SkipToken interface{} `json:"$skipToken,omitempty"`
}

// ResourceChangesRequestParameters the parameters for a specific changes request.
type ResourceChangesRequestParameters struct {
	// ResourceID - Specifies the resource for a changes request.
	ResourceID *string `json:"resourceId,omitempty"`
	// Interval - Specifies the date and time interval for a changes request.
	Interval *ResourceChangesRequestParametersInterval `json:"interval,omitempty"`
	// SkipToken - Acts as the continuation token for paged responses.
	SkipToken *string `json:"$skipToken,omitempty"`
	// Top - The maximum number of changes the client can accept in a paged response.
	Top *int32 `json:"$top,omitempty"`
}

// ResourceChangesRequestParametersInterval specifies the date and time interval for a changes request.
type ResourceChangesRequestParametersInterval struct {
	// Start - A datetime indicating the inclusive/closed start of the time interval, i.e. `[`**`start`**`, end)`. Specifying a `start` that occurs chronologically after `end` will result in an error.
	Start *date.Time `json:"start,omitempty"`
	// End - A datetime indicating the exclusive/open end of the time interval, i.e. `[start, `**`end`**`)`. Specifying an `end` that occurs chronologically before `start` will result in an error.
	End *date.Time `json:"end,omitempty"`
}

// ResourceSnapshotData data on a specific resource snapshot.
type ResourceSnapshotData struct {
	// Timestamp - The time when the snapshot was created.
	// The snapshot timestamp provides an approximation as to when a modification to a resource was detected.  There can be a difference between the actual modification time and the detection time.  This is due to differences in how operations that modify a resource are processed, versus how operation that record resource snapshots are processed.
	Timestamp *date.Time `json:"timestamp,omitempty"`
	// Content - The resource snapshot content (in resourceChangeDetails response only).
	Content interface{} `json:"content,omitempty"`
}

// Table query output in tabular format.
type Table struct {
	// Columns - Query result column descriptors.
	Columns *[]Column `json:"columns,omitempty"`
	// Rows - Query result rows.
	Rows *[][]interface{} `json:"rows,omitempty"`
}
