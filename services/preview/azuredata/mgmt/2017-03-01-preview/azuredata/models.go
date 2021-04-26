package azuredata

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
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/azuredata/mgmt/2017-03-01-preview/azuredata"

// CloudError an error response from the Azure Data service.
type CloudError struct {
	// Error - null
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody an error response from the Batch service.
type CloudErrorBody struct {
	// Code - An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`
	// Message - A message describing the error, intended to be suitable for display in a user interface.
	Message *string `json:"message,omitempty"`
	// Target - The target of the particular error. For example, the name of the property in error.
	Target *string `json:"target,omitempty"`
	// Details - A list of additional details about the error.
	Details *[]CloudErrorBody `json:"details,omitempty"`
}

// Identity identity for the resource.
type Identity struct {
	// PrincipalID - READ-ONLY; The principal ID of resource identity.
	PrincipalID *string `json:"principalId,omitempty"`
	// TenantID - READ-ONLY; The tenant ID of resource.
	TenantID *string `json:"tenantId,omitempty"`
	// Type - The identity type. Possible values include: 'SystemAssigned'
	Type ResourceIdentityType `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Identity.
func (i Identity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if i.Type != "" {
		objectMap["type"] = i.Type
	}
	return json.Marshal(objectMap)
}

// ODataError information about an error.
type ODataError struct {
	// Code - A language-independent error name.
	Code *string `json:"code,omitempty"`
	// Message - The error message.
	Message *string `json:"message,omitempty"`
	// Target - The target of the error (for example, the name of the property in error).
	Target *string `json:"target,omitempty"`
	// Details - The error details.
	Details *[]ODataError `json:"details,omitempty"`
}

// Operation SQL REST API operation definition.
type Operation struct {
	// Name - READ-ONLY; The name of the operation being performed on this particular object.
	Name *string `json:"name,omitempty"`
	// Display - READ-ONLY; The localized display information for this particular operation / action.
	Display *OperationDisplay `json:"display,omitempty"`
	// Origin - READ-ONLY; The intended executor of the operation. Possible values include: 'OperationOriginUser', 'OperationOriginSystem'
	Origin OperationOrigin `json:"origin,omitempty"`
	// Properties - READ-ONLY; Additional descriptions for the operation.
	Properties map[string]interface{} `json:"properties"`
}

// MarshalJSON is the custom marshaler for Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// OperationDisplay display metadata associated with the operation.
type OperationDisplay struct {
	// Provider - READ-ONLY; The localized friendly form of the resource provider name.
	Provider *string `json:"provider,omitempty"`
	// Resource - READ-ONLY; The localized friendly form of the resource type related to this action/operation.
	Resource *string `json:"resource,omitempty"`
	// Operation - READ-ONLY; The localized friendly name for the operation.
	Operation *string `json:"operation,omitempty"`
	// Description - READ-ONLY; The localized friendly description for the operation.
	Description *string `json:"description,omitempty"`
}

// OperationListResult result of the request to list SQL operations.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; Array of results.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - READ-ONLY; Link to retrieve next page of results.
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

// Plan plan for the resource.
type Plan struct {
	// Name - A user defined name of the 3rd Party Artifact that is being procured.
	Name *string `json:"name,omitempty"`
	// Publisher - The publisher of the 3rd Party Artifact that is being bought. E.g. NewRelic
	Publisher *string `json:"publisher,omitempty"`
	// Product - The 3rd Party artifact that is being procured. E.g. NewRelic. Product maps to the OfferID specified for the artifact at the time of Data Market onboarding.
	Product *string `json:"product,omitempty"`
	// PromotionCode - A publisher provided promotion code as provisioned in Data Market for the said product/artifact.
	PromotionCode *string `json:"promotionCode,omitempty"`
	// Version - The version of the desired product/artifact.
	Version *string `json:"version,omitempty"`
}

// ProxyResource the resource model definition for a ARM proxy resource. It will have everything other than
// required location and tags
type ProxyResource struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// Resource ...
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// ResourceModelWithAllowedPropertySet the resource model definition containing the full set of allowed
// properties for a resource. Except properties bag, there cannot be a top level property outside of this
// set.
type ResourceModelWithAllowedPropertySet struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts..
	Type *string `json:"type,omitempty"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ManagedBy - The  fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy *string `json:"managedBy,omitempty"`
	// Kind - Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `json:"kind,omitempty"`
	// Etag - READ-ONLY; The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	Etag *string `json:"etag,omitempty"`
	// Tags - Resource tags.
	Tags     map[string]*string                           `json:"tags"`
	Identity *ResourceModelWithAllowedPropertySetIdentity `json:"identity,omitempty"`
	Sku      *ResourceModelWithAllowedPropertySetSku      `json:"sku,omitempty"`
	Plan     *ResourceModelWithAllowedPropertySetPlan     `json:"plan,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceModelWithAllowedPropertySet.
func (rmwaps ResourceModelWithAllowedPropertySet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rmwaps.Location != nil {
		objectMap["location"] = rmwaps.Location
	}
	if rmwaps.ManagedBy != nil {
		objectMap["managedBy"] = rmwaps.ManagedBy
	}
	if rmwaps.Kind != nil {
		objectMap["kind"] = rmwaps.Kind
	}
	if rmwaps.Tags != nil {
		objectMap["tags"] = rmwaps.Tags
	}
	if rmwaps.Identity != nil {
		objectMap["identity"] = rmwaps.Identity
	}
	if rmwaps.Sku != nil {
		objectMap["sku"] = rmwaps.Sku
	}
	if rmwaps.Plan != nil {
		objectMap["plan"] = rmwaps.Plan
	}
	return json.Marshal(objectMap)
}

// ResourceModelWithAllowedPropertySetIdentity ...
type ResourceModelWithAllowedPropertySetIdentity struct {
	// PrincipalID - READ-ONLY; The principal ID of resource identity.
	PrincipalID *string `json:"principalId,omitempty"`
	// TenantID - READ-ONLY; The tenant ID of resource.
	TenantID *string `json:"tenantId,omitempty"`
	// Type - The identity type. Possible values include: 'SystemAssigned'
	Type ResourceIdentityType `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceModelWithAllowedPropertySetIdentity.
func (rmwaps ResourceModelWithAllowedPropertySetIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rmwaps.Type != "" {
		objectMap["type"] = rmwaps.Type
	}
	return json.Marshal(objectMap)
}

// ResourceModelWithAllowedPropertySetPlan ...
type ResourceModelWithAllowedPropertySetPlan struct {
	// Name - A user defined name of the 3rd Party Artifact that is being procured.
	Name *string `json:"name,omitempty"`
	// Publisher - The publisher of the 3rd Party Artifact that is being bought. E.g. NewRelic
	Publisher *string `json:"publisher,omitempty"`
	// Product - The 3rd Party artifact that is being procured. E.g. NewRelic. Product maps to the OfferID specified for the artifact at the time of Data Market onboarding.
	Product *string `json:"product,omitempty"`
	// PromotionCode - A publisher provided promotion code as provisioned in Data Market for the said product/artifact.
	PromotionCode *string `json:"promotionCode,omitempty"`
	// Version - The version of the desired product/artifact.
	Version *string `json:"version,omitempty"`
}

// ResourceModelWithAllowedPropertySetSku ...
type ResourceModelWithAllowedPropertySetSku struct {
	// Name - The name of the SKU. Ex - P3. It is typically a letter+number code
	Name *string `json:"name,omitempty"`
	// Tier - This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT. Possible values include: 'Free', 'Basic', 'Standard', 'Premium'
	Tier SkuTier `json:"tier,omitempty"`
	// Size - The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code.
	Size *string `json:"size,omitempty"`
	// Family - If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string `json:"family,omitempty"`
	// Capacity - If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the resource this may be omitted.
	Capacity *int32 `json:"capacity,omitempty"`
}

// ResourceSku ...
type ResourceSku struct {
	Capacity *int32  `json:"capacity,omitempty"`
	Family   *string `json:"family,omitempty"`
	Name     *string `json:"name,omitempty"`
	Size     *string `json:"size,omitempty"`
	Tier     *string `json:"tier,omitempty"`
}

// Sku the resource model definition representing SKU
type Sku struct {
	// Name - The name of the SKU. Ex - P3. It is typically a letter+number code
	Name *string `json:"name,omitempty"`
	// Tier - This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT. Possible values include: 'Free', 'Basic', 'Standard', 'Premium'
	Tier SkuTier `json:"tier,omitempty"`
	// Size - The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code.
	Size *string `json:"size,omitempty"`
	// Family - If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string `json:"family,omitempty"`
	// Capacity - If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the resource this may be omitted.
	Capacity *int32 `json:"capacity,omitempty"`
}

// SQLServer a SQL server.
type SQLServer struct {
	autorest.Response `json:"-"`
	// SQLServerProperties - Resource properties.
	*SQLServerProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for SQLServer.
func (ss SQLServer) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ss.SQLServerProperties != nil {
		objectMap["properties"] = ss.SQLServerProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for SQLServer struct.
func (ss *SQLServer) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var SQLServerProperties SQLServerProperties
				err = json.Unmarshal(*v, &SQLServerProperties)
				if err != nil {
					return err
				}
				ss.SQLServerProperties = &SQLServerProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ss.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ss.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ss.Type = &typeVar
			}
		}
	}

	return nil
}

// SQLServerListResult a list of SQL servers.
type SQLServerListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; Array of results.
	Value *[]SQLServer `json:"value,omitempty"`
	// NextLink - READ-ONLY; Link to retrieve next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// SQLServerListResultIterator provides access to a complete listing of SQLServer values.
type SQLServerListResultIterator struct {
	i    int
	page SQLServerListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *SQLServerListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerListResultIterator.NextWithContext")
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
func (iter *SQLServerListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter SQLServerListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter SQLServerListResultIterator) Response() SQLServerListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter SQLServerListResultIterator) Value() SQLServer {
	if !iter.page.NotDone() {
		return SQLServer{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the SQLServerListResultIterator type.
func NewSQLServerListResultIterator(page SQLServerListResultPage) SQLServerListResultIterator {
	return SQLServerListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (sslr SQLServerListResult) IsEmpty() bool {
	return sslr.Value == nil || len(*sslr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (sslr SQLServerListResult) hasNextLink() bool {
	return sslr.NextLink != nil && len(*sslr.NextLink) != 0
}

// sQLServerListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (sslr SQLServerListResult) sQLServerListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !sslr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(sslr.NextLink)))
}

// SQLServerListResultPage contains a page of SQLServer values.
type SQLServerListResultPage struct {
	fn   func(context.Context, SQLServerListResult) (SQLServerListResult, error)
	sslr SQLServerListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *SQLServerListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.sslr)
		if err != nil {
			return err
		}
		page.sslr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *SQLServerListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page SQLServerListResultPage) NotDone() bool {
	return !page.sslr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page SQLServerListResultPage) Response() SQLServerListResult {
	return page.sslr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page SQLServerListResultPage) Values() []SQLServer {
	if page.sslr.IsEmpty() {
		return nil
	}
	return *page.sslr.Value
}

// Creates a new instance of the SQLServerListResultPage type.
func NewSQLServerListResultPage(cur SQLServerListResult, getNextPage func(context.Context, SQLServerListResult) (SQLServerListResult, error)) SQLServerListResultPage {
	return SQLServerListResultPage{
		fn:   getNextPage,
		sslr: cur,
	}
}

// SQLServerProperties the SQL server properties.
type SQLServerProperties struct {
	// Cores - Cores of the Sql Server.
	Cores *int32 `json:"cores,omitempty"`
	// Version - Version of the Sql Server.
	Version *string `json:"version,omitempty"`
	// Edition - Sql Server Edition.
	Edition *string `json:"edition,omitempty"`
	// RegistrationID - ID for Parent Sql Server Registration.
	RegistrationID *string `json:"registrationID,omitempty"`
	// PropertyBag - Sql Server Json Property Bag.
	PropertyBag *string `json:"propertyBag,omitempty"`
}

// SQLServerRegistration a SQL server registration.
type SQLServerRegistration struct {
	autorest.Response `json:"-"`
	// SQLServerRegistrationProperties - Resource properties.
	*SQLServerRegistrationProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// SystemData - READ-ONLY
	SystemData *SystemData `json:"systemData,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for SQLServerRegistration.
func (ssr SQLServerRegistration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ssr.SQLServerRegistrationProperties != nil {
		objectMap["properties"] = ssr.SQLServerRegistrationProperties
	}
	if ssr.Tags != nil {
		objectMap["tags"] = ssr.Tags
	}
	if ssr.Location != nil {
		objectMap["location"] = ssr.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for SQLServerRegistration struct.
func (ssr *SQLServerRegistration) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var SQLServerRegistrationProperties SQLServerRegistrationProperties
				err = json.Unmarshal(*v, &SQLServerRegistrationProperties)
				if err != nil {
					return err
				}
				ssr.SQLServerRegistrationProperties = &SQLServerRegistrationProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				ssr.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				ssr.Location = &location
			}
		case "systemData":
			if v != nil {
				var systemData SystemData
				err = json.Unmarshal(*v, &systemData)
				if err != nil {
					return err
				}
				ssr.SystemData = &systemData
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ssr.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ssr.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ssr.Type = &typeVar
			}
		}
	}

	return nil
}

// SQLServerRegistrationListResult server
type SQLServerRegistrationListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; Array of results.
	Value *[]SQLServerRegistration `json:"value,omitempty"`
	// NextLink - READ-ONLY; Link to retrieve next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// SQLServerRegistrationListResultIterator provides access to a complete listing of SQLServerRegistration
// values.
type SQLServerRegistrationListResultIterator struct {
	i    int
	page SQLServerRegistrationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *SQLServerRegistrationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerRegistrationListResultIterator.NextWithContext")
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
func (iter *SQLServerRegistrationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter SQLServerRegistrationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter SQLServerRegistrationListResultIterator) Response() SQLServerRegistrationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter SQLServerRegistrationListResultIterator) Value() SQLServerRegistration {
	if !iter.page.NotDone() {
		return SQLServerRegistration{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the SQLServerRegistrationListResultIterator type.
func NewSQLServerRegistrationListResultIterator(page SQLServerRegistrationListResultPage) SQLServerRegistrationListResultIterator {
	return SQLServerRegistrationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ssrlr SQLServerRegistrationListResult) IsEmpty() bool {
	return ssrlr.Value == nil || len(*ssrlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (ssrlr SQLServerRegistrationListResult) hasNextLink() bool {
	return ssrlr.NextLink != nil && len(*ssrlr.NextLink) != 0
}

// sQLServerRegistrationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ssrlr SQLServerRegistrationListResult) sQLServerRegistrationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !ssrlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ssrlr.NextLink)))
}

// SQLServerRegistrationListResultPage contains a page of SQLServerRegistration values.
type SQLServerRegistrationListResultPage struct {
	fn    func(context.Context, SQLServerRegistrationListResult) (SQLServerRegistrationListResult, error)
	ssrlr SQLServerRegistrationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *SQLServerRegistrationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerRegistrationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.ssrlr)
		if err != nil {
			return err
		}
		page.ssrlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *SQLServerRegistrationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page SQLServerRegistrationListResultPage) NotDone() bool {
	return !page.ssrlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page SQLServerRegistrationListResultPage) Response() SQLServerRegistrationListResult {
	return page.ssrlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page SQLServerRegistrationListResultPage) Values() []SQLServerRegistration {
	if page.ssrlr.IsEmpty() {
		return nil
	}
	return *page.ssrlr.Value
}

// Creates a new instance of the SQLServerRegistrationListResultPage type.
func NewSQLServerRegistrationListResultPage(cur SQLServerRegistrationListResult, getNextPage func(context.Context, SQLServerRegistrationListResult) (SQLServerRegistrationListResult, error)) SQLServerRegistrationListResultPage {
	return SQLServerRegistrationListResultPage{
		fn:    getNextPage,
		ssrlr: cur,
	}
}

// SQLServerRegistrationProperties the SQL server Registration properties.
type SQLServerRegistrationProperties struct {
	// SubscriptionID - Subscription Id
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// ResourceGroup - Resource Group Name
	ResourceGroup *string `json:"resourceGroup,omitempty"`
	// PropertyBag - Optional Properties as JSON string
	PropertyBag *string `json:"propertyBag,omitempty"`
}

// SQLServerRegistrationUpdate an update to a SQL Server Registration.
type SQLServerRegistrationUpdate struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for SQLServerRegistrationUpdate.
func (ssru SQLServerRegistrationUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ssru.Tags != nil {
		objectMap["tags"] = ssru.Tags
	}
	return json.Marshal(objectMap)
}

// SystemData read only system data
type SystemData struct {
	// CreatedBy - An identifier for the identity that created the resource
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedByType - The type of identity that created the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	CreatedByType IdentityType `json:"createdByType,omitempty"`
	// CreatedAt - The timestamp of resource creation (UTC)
	CreatedAt *date.Time `json:"createdAt,omitempty"`
	// LastModifiedBy - An identifier for the identity that last modified the resource
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// LastModifiedByType - The type of identity that last modified the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	LastModifiedByType IdentityType `json:"lastModifiedByType,omitempty"`
	// LastModifiedAt - The timestamp of resource last modification (UTC)
	LastModifiedAt *date.Time `json:"lastModifiedAt,omitempty"`
}

// TrackedResource the resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// SystemData - READ-ONLY
	SystemData *SystemData `json:"systemData,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	return json.Marshal(objectMap)
}
