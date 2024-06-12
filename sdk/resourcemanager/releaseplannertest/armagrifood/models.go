//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armagrifood

import "time"

// APIProperties - Api properties.
type APIProperties struct {
	// Interval in minutes for which the weather data for the api needs to be refreshed.
	APIFreshnessWindowInMinutes *int32 `json:"apiFreshnessWindowInMinutes,omitempty"`
}

// ArmAsyncOperation - Arm async operation class. Ref: https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/async-operations.
type ArmAsyncOperation struct {
	// Status of the async operation.
	Status *string `json:"status,omitempty"`
}

// CheckNameAvailabilityRequest - The check availability request body.
type CheckNameAvailabilityRequest struct {
	// The name of the resource for which availability needs to be checked.
	Name *string `json:"name,omitempty"`

	// The resource type.
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityResponse - The check availability result.
type CheckNameAvailabilityResponse struct {
	// Detailed reason why the given name is available.
	Message *string `json:"message,omitempty"`

	// Indicates if the resource name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`

	// The reason why the given name is not available.
	Reason *CheckNameAvailabilityReason `json:"reason,omitempty"`
}

// DetailedInformation - Model to capture detailed information for farmBeatsExtensions.
type DetailedInformation struct {
	// List of apiInputParameters.
	APIInputParameters []*string `json:"apiInputParameters,omitempty"`

	// ApiName available for the farmBeatsExtension.
	APIName *string `json:"apiName,omitempty"`

	// List of customParameters.
	CustomParameters []*string `json:"customParameters,omitempty"`

	// List of platformParameters.
	PlatformParameters []*string `json:"platformParameters,omitempty"`

	// Unit systems info for the data provider.
	UnitsSupported *UnitSystemsInfo `json:"unitsSupported,omitempty"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// Extension resource.
type Extension struct {
	// Extension resource properties.
	Properties *ExtensionProperties `json:"properties,omitempty"`

	// READ-ONLY; The ETag value to implement optimistic concurrency.
	ETag *string `json:"eTag,omitempty" azure:"ro"`

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ExtensionInstallationRequest - Extension Installation Request Body.
type ExtensionInstallationRequest struct {
	// Additional Api Properties.
	AdditionalAPIProperties map[string]*APIProperties `json:"additionalApiProperties,omitempty"`

	// Extension Version.
	ExtensionVersion *string `json:"extensionVersion,omitempty"`
}

// ExtensionListResponse - Paged response contains list of requested objects and a URL link to get the next set of results.
type ExtensionListResponse struct {
	// List of requested objects.
	Value []*Extension `json:"value,omitempty"`

	// READ-ONLY; Continuation link (absolute URI) to the next page of results in the list.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// ExtensionProperties - Extension resource properties.
type ExtensionProperties struct {
	// READ-ONLY; Additional api properties.
	AdditionalAPIProperties map[string]*APIProperties `json:"additionalApiProperties,omitempty" azure:"ro"`

	// READ-ONLY; Extension api docs link.
	ExtensionAPIDocsLink *string `json:"extensionApiDocsLink,omitempty" azure:"ro"`

	// READ-ONLY; Extension auth link.
	ExtensionAuthLink *string `json:"extensionAuthLink,omitempty" azure:"ro"`

	// READ-ONLY; Extension category. e.g. weather/sensor/satellite.
	ExtensionCategory *string `json:"extensionCategory,omitempty" azure:"ro"`

	// READ-ONLY; Extension Id.
	ExtensionID *string `json:"extensionId,omitempty" azure:"ro"`

	// READ-ONLY; Installed extension version.
	InstalledExtensionVersion *string `json:"installedExtensionVersion,omitempty" azure:"ro"`
}

// ExtensionsClientCreateOrUpdateOptions contains the optional parameters for the ExtensionsClient.CreateOrUpdate method.
type ExtensionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ExtensionsClientDeleteOptions contains the optional parameters for the ExtensionsClient.Delete method.
type ExtensionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ExtensionsClientGetOptions contains the optional parameters for the ExtensionsClient.Get method.
type ExtensionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ExtensionsClientListByFarmBeatsOptions contains the optional parameters for the ExtensionsClient.ListByFarmBeats method.
type ExtensionsClientListByFarmBeatsOptions struct {
	// Installed extension categories.
	ExtensionCategories []string
	// Installed extension ids.
	ExtensionIDs []string
	// Maximum number of items needed (inclusive). Minimum = 10, Maximum = 1000, Default value = 50.
	MaxPageSize *int32
	// Skip token for getting next set of results.
	SkipToken *string
}

// FarmBeats ARM Resource.
type FarmBeats struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Identity for the resource.
	Identity *Identity `json:"identity,omitempty"`

	// FarmBeats ARM Resource properties.
	Properties *FarmBeatsProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// FarmBeatsExtension - FarmBeats extension resource.
type FarmBeatsExtension struct {
	// FarmBeatsExtension properties.
	Properties *FarmBeatsExtensionProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// FarmBeatsExtensionListResponse - Paged response contains list of requested objects and a URL link to get the next set of
// results.
type FarmBeatsExtensionListResponse struct {
	// List of requested objects.
	Value []*FarmBeatsExtension `json:"value,omitempty"`

	// READ-ONLY; Continuation link (absolute URI) to the next page of results in the list.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// FarmBeatsExtensionProperties - FarmBeatsExtension properties.
type FarmBeatsExtensionProperties struct {
	// READ-ONLY; Textual description.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; Detailed information which shows summary of requested data. Used in descriptive get extension metadata call.
	// Information for weather category per api included are apisSupported, customParameters,
	// PlatformParameters and Units supported.
	DetailedInformation []*DetailedInformation `json:"detailedInformation,omitempty" azure:"ro"`

	// READ-ONLY; FarmBeatsExtension api docs link.
	ExtensionAPIDocsLink *string `json:"extensionApiDocsLink,omitempty" azure:"ro"`

	// READ-ONLY; FarmBeatsExtension auth link.
	ExtensionAuthLink *string `json:"extensionAuthLink,omitempty" azure:"ro"`

	// READ-ONLY; Category of the extension. e.g. weather/sensor/satellite.
	ExtensionCategory *string `json:"extensionCategory,omitempty" azure:"ro"`

	// READ-ONLY; FarmBeatsExtension ID.
	FarmBeatsExtensionID *string `json:"farmBeatsExtensionId,omitempty" azure:"ro"`

	// READ-ONLY; FarmBeatsExtension name.
	FarmBeatsExtensionName *string `json:"farmBeatsExtensionName,omitempty" azure:"ro"`

	// READ-ONLY; FarmBeatsExtension version.
	FarmBeatsExtensionVersion *string `json:"farmBeatsExtensionVersion,omitempty" azure:"ro"`

	// READ-ONLY; Publisher ID.
	PublisherID *string `json:"publisherId,omitempty" azure:"ro"`

	// READ-ONLY; Target ResourceType of the farmBeatsExtension.
	TargetResourceType *string `json:"targetResourceType,omitempty" azure:"ro"`
}

// FarmBeatsExtensionsClientGetOptions contains the optional parameters for the FarmBeatsExtensionsClient.Get method.
type FarmBeatsExtensionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// FarmBeatsExtensionsClientListOptions contains the optional parameters for the FarmBeatsExtensionsClient.List method.
type FarmBeatsExtensionsClientListOptions struct {
	// Extension categories.
	ExtensionCategories []string
	// FarmBeatsExtension ids.
	FarmBeatsExtensionIDs []string
	// FarmBeats extension names.
	FarmBeatsExtensionNames []string
	// Maximum number of items needed (inclusive). Minimum = 10, Maximum = 1000, Default value = 50.
	MaxPageSize *int32
	// Publisher ids.
	PublisherIDs []string
}

// FarmBeatsListResponse - Paged response contains list of requested objects and a URL link to get the next set of results.
type FarmBeatsListResponse struct {
	// List of requested objects.
	Value []*FarmBeats `json:"value,omitempty"`

	// READ-ONLY; Continuation link (absolute URI) to the next page of results in the list.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// FarmBeatsModelsClientBeginUpdateOptions contains the optional parameters for the FarmBeatsModelsClient.BeginUpdate method.
type FarmBeatsModelsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FarmBeatsModelsClientCreateOrUpdateOptions contains the optional parameters for the FarmBeatsModelsClient.CreateOrUpdate
// method.
type FarmBeatsModelsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// FarmBeatsModelsClientDeleteOptions contains the optional parameters for the FarmBeatsModelsClient.Delete method.
type FarmBeatsModelsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// FarmBeatsModelsClientGetOperationResultOptions contains the optional parameters for the FarmBeatsModelsClient.GetOperationResult
// method.
type FarmBeatsModelsClientGetOperationResultOptions struct {
	// placeholder for future optional parameters
}

// FarmBeatsModelsClientGetOptions contains the optional parameters for the FarmBeatsModelsClient.Get method.
type FarmBeatsModelsClientGetOptions struct {
	// placeholder for future optional parameters
}

// FarmBeatsModelsClientListByResourceGroupOptions contains the optional parameters for the FarmBeatsModelsClient.ListByResourceGroup
// method.
type FarmBeatsModelsClientListByResourceGroupOptions struct {
	// Maximum number of items needed (inclusive). Minimum = 10, Maximum = 1000, Default value = 50.
	MaxPageSize *int32
	// Continuation token for getting next set of results.
	SkipToken *string
}

// FarmBeatsModelsClientListBySubscriptionOptions contains the optional parameters for the FarmBeatsModelsClient.ListBySubscription
// method.
type FarmBeatsModelsClientListBySubscriptionOptions struct {
	// Maximum number of items needed (inclusive). Minimum = 10, Maximum = 1000, Default value = 50.
	MaxPageSize *int32
	// Skip token for getting next set of results.
	SkipToken *string
}

// FarmBeatsProperties - FarmBeats ARM Resource properties.
type FarmBeatsProperties struct {
	// Property to allow or block public traffic for an Azure FarmBeats resource.
	PublicNetworkAccess *PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// Sensor integration request model.
	SensorIntegration *SensorIntegration `json:"sensorIntegration,omitempty"`

	// READ-ONLY; Uri of the FarmBeats instance.
	InstanceURI *string `json:"instanceUri,omitempty" azure:"ro"`

	// READ-ONLY; The private endpoint connection resource.
	PrivateEndpointConnections *PrivateEndpointConnection `json:"privateEndpointConnections,omitempty" azure:"ro"`

	// READ-ONLY; FarmBeats instance provisioning state.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// FarmBeatsSolution - FarmBeats solution resource.
type FarmBeatsSolution struct {
	// FarmBeatsSolution properties.
	Properties *FarmBeatsSolutionProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// FarmBeatsSolutionListResponse - Paged response contains list of requested objects and a URL link to get the next set of
// results.
type FarmBeatsSolutionListResponse struct {
	// Continuation link (absolute URI) to the next page of results in the list.
	NextLink *string `json:"nextLink,omitempty"`

	// Token used in retrieving the next page. If null, there are no additional pages.
	SkipToken *string `json:"$skipToken,omitempty"`

	// List of requested objects.
	Value []*FarmBeatsSolution `json:"value,omitempty"`
}

// FarmBeatsSolutionProperties - FarmBeatsSolution properties.
type FarmBeatsSolutionProperties struct {
	MarketplaceOfferDetails *MarketplaceOfferDetails `json:"marketplaceOfferDetails,omitempty"`

	// READ-ONLY; Application id of the multi tenant application to be used by partner to access FarmBeats data.
	AccessFBApplicationID *string `json:"accessFBApplicationId,omitempty" azure:"ro"`

	// READ-ONLY; Application name of the multi tenant application to be used by partner to access FarmBeatsData.
	AccessFBApplicationName *string `json:"accessFBApplicationName,omitempty" azure:"ro"`

	// READ-ONLY; List of ActionIds needed to make the SaaS multi tenant application access relevant fb data.
	ActionIDs []*string `json:"actionIds,omitempty" azure:"ro"`

	// READ-ONLY; Gets scope of the FarmBeats data access that's required for processing solution request to partner. Example:
	// For gdd they might need weatherScope and satelliteScope.
	DataAccessScopes []*string `json:"dataAccessScopes,omitempty" azure:"ro"`

	// READ-ONLY; Gets example name: insight sample response Dictionary to capture all variations of computed results ingested
	// by partner.
	EvaluatedOutputsDictionary map[string]*SolutionEvaluatedOutput `json:"evaluatedOutputsDictionary,omitempty" azure:"ro"`

	// READ-ONLY; Gets scope of the FarmBeats related parameters that need to be validated in apiInputParameters. Example: For
	// if 'FarmHierarchy' is the input scope for 'WeatherScope' data access For working with
	// WeatherScope we need FarmHierarchy info implies 'farmerId', 'resourceId', 'resourceType' in request body.
	InputParametersValidationScopes []*ResourceParameter `json:"inputParametersValidationScopes,omitempty" azure:"ro"`

	// READ-ONLY; Gets apiVersion: Swagger Document Dictionary to capture all api versions of swagger exposed by partner to farmbeats.
	OpenAPISpecsDictionary map[string]interface{} `json:"openApiSpecsDictionary,omitempty" azure:"ro"`

	// READ-ONLY; Solution Partner Id.
	PartnerID *string `json:"partnerId,omitempty" azure:"ro"`

	// READ-ONLY; Solution Partner Tenant Id.
	PartnerTenantID *string `json:"partnerTenantId,omitempty" azure:"ro"`

	// READ-ONLY; Role Id of the SaaS multi tenant application to access relevant fb data.
	RoleID *string `json:"roleId,omitempty" azure:"ro"`

	// READ-ONLY; Role Name of the SaaS multi tenant application to access relevant fb data.
	RoleName *string `json:"roleName,omitempty" azure:"ro"`

	// READ-ONLY; Application id of the SaaS multi tenant application.
	SaaSApplicationID *string `json:"saaSApplicationId,omitempty" azure:"ro"`
}

// FarmBeatsUpdateProperties - FarmBeats ARM Resource properties.
type FarmBeatsUpdateProperties struct {
	// Property to allow or block public traffic for an Azure FarmBeats resource.
	PublicNetworkAccess *PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// Sensor integration request model.
	SensorIntegration *SensorIntegration `json:"sensorIntegration,omitempty"`
}

// FarmBeatsUpdateRequestModel - FarmBeats update request.
type FarmBeatsUpdateRequestModel struct {
	// Identity for the resource.
	Identity *Identity `json:"identity,omitempty"`

	// Geo-location where the resource lives.
	Location *string `json:"location,omitempty"`

	// FarmBeats ARM Resource properties.
	Properties *FarmBeatsUpdateProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// Identity for the resource.
type Identity struct {
	// The identity type.
	Type *string `json:"type,omitempty"`

	// READ-ONLY; The principal ID of resource identity. The value must be an UUID.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant ID of resource. The value must be an UUID.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

type Insight struct {
	CreatedDateTime      *time.Time `json:"createdDateTime,omitempty"`
	Description          *string    `json:"description,omitempty"`
	ETag                 *string    `json:"eTag,omitempty"`
	FarmerID             *string    `json:"farmerId,omitempty"`
	ID                   *string    `json:"id,omitempty"`
	InsightEndDateTime   *time.Time `json:"insightEndDateTime,omitempty"`
	InsightStartDateTime *time.Time `json:"insightStartDateTime,omitempty"`

	// Dictionary of
	Measures         map[string]*Measure `json:"measures,omitempty"`
	ModelID          *string             `json:"modelId,omitempty"`
	ModelVersion     *string             `json:"modelVersion,omitempty"`
	ModifiedDateTime *time.Time          `json:"modifiedDateTime,omitempty"`
	Name             *string             `json:"name,omitempty"`

	// Dictionary of
	Properties   map[string]interface{} `json:"properties,omitempty"`
	ResourceID   *string                `json:"resourceId,omitempty"`
	ResourceType *string                `json:"resourceType,omitempty"`
	Status       *string                `json:"status,omitempty"`
}

type InsightAttachment struct {
	CreatedDateTime  *time.Time `json:"createdDateTime,omitempty"`
	Description      *string    `json:"description,omitempty"`
	ETag             *string    `json:"eTag,omitempty"`
	FarmerID         *string    `json:"farmerId,omitempty"`
	FileLink         *string    `json:"fileLink,omitempty"`
	ID               *string    `json:"id,omitempty"`
	InsightID        *string    `json:"insightId,omitempty"`
	ModelID          *string    `json:"modelId,omitempty"`
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	Name             *string    `json:"name,omitempty"`
	OriginalFileName *string    `json:"originalFileName,omitempty"`
	ResourceID       *string    `json:"resourceId,omitempty"`
	ResourceType     *string    `json:"resourceType,omitempty"`
	Status           *string    `json:"status,omitempty"`
}

// LocationsClientCheckNameAvailabilityOptions contains the optional parameters for the LocationsClient.CheckNameAvailability
// method.
type LocationsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

type MarketplaceOfferDetails struct {
	PublisherID *string `json:"publisherId,omitempty"`
	SaasOfferID *string `json:"saasOfferId,omitempty"`
}

type Measure struct {
	Unit  *string  `json:"unit,omitempty"`
	Value *float64 `json:"value,omitempty"`
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpoint - The private endpoint resource.
type PrivateEndpoint struct {
	// READ-ONLY; The ARM identifier for private endpoint.
	ID *string `json:"id,omitempty" azure:"ro"`
}

// PrivateEndpointConnection - The private endpoint connection resource.
type PrivateEndpointConnection struct {
	// Resource properties.
	Properties *PrivateEndpointConnectionProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionListResult - List of private endpoint connections associated with the specified resource.
type PrivateEndpointConnectionListResult struct {
	// Array of private endpoint connections.
	Value []*PrivateEndpointConnection `json:"value,omitempty"`
}

// PrivateEndpointConnectionProperties - Properties of the private endpoint connection.
type PrivateEndpointConnectionProperties struct {
	// REQUIRED; A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `json:"privateLinkServiceConnectionState,omitempty"`

	// The private endpoint resource.
	PrivateEndpoint *PrivateEndpoint `json:"privateEndpoint,omitempty"`

	// READ-ONLY; The group ids for the private endpoint resource.
	GroupIDs []*string `json:"groupIds,omitempty" azure:"ro"`

	// READ-ONLY; The provisioning state of the private endpoint connection resource.
	ProvisioningState *PrivateEndpointConnectionProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
// method.
type PrivateEndpointConnectionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientCreateOrUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.CreateOrUpdate
// method.
type PrivateEndpointConnectionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
type PrivateEndpointConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientListByResourceOptions contains the optional parameters for the PrivateEndpointConnectionsClient.ListByResource
// method.
type PrivateEndpointConnectionsClientListByResourceOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResource - A private link resource.
type PrivateLinkResource struct {
	// Resource properties.
	Properties *PrivateLinkResourceProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateLinkResourceListResult - A list of private link resources.
type PrivateLinkResourceListResult struct {
	// Array of private link resources
	Value []*PrivateLinkResource `json:"value,omitempty"`
}

// PrivateLinkResourceProperties - Properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// The private link resource private link DNS zone name.
	RequiredZoneNames []*string `json:"requiredZoneNames,omitempty"`

	// READ-ONLY; The private link resource group id.
	GroupID *string `json:"groupId,omitempty" azure:"ro"`

	// READ-ONLY; The private link resource required member names.
	RequiredMembers []*string `json:"requiredMembers,omitempty" azure:"ro"`
}

// PrivateLinkResourcesClientGetOptions contains the optional parameters for the PrivateLinkResourcesClient.Get method.
type PrivateLinkResourcesClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientListByResourceOptions contains the optional parameters for the PrivateLinkResourcesClient.ListByResource
// method.
type PrivateLinkResourcesClientListByResourceOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkServiceConnectionState - A collection of information about the state of the connection between service consumer
// and provider.
type PrivateLinkServiceConnectionState struct {
	// A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string `json:"actionsRequired,omitempty"`

	// The reason for approval/rejection of the connection.
	Description *string `json:"description,omitempty"`

	// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *PrivateEndpointServiceConnectionStatus `json:"status,omitempty"`
}

type ResourceParameter struct {
	ResourceIDName *string `json:"resourceIdName,omitempty"`
	ResourceType   *string `json:"resourceType,omitempty"`
}

// SensorIntegration - Sensor integration request model.
type SensorIntegration struct {
	// Sensor integration enable state. Allowed values are True, None
	Enabled *string `json:"enabled,omitempty"`

	// Common error response for all Azure Resource Manager APIs to return error details for failed operations. (This also follows
	// the OData error response format.).
	ProvisioningInfo *ErrorResponse `json:"provisioningInfo,omitempty"`

	// READ-ONLY; Sensor integration instance provisioning state.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// Solution resource.
type Solution struct {
	// Solution resource properties.
	Properties *SolutionProperties `json:"properties,omitempty"`

	// READ-ONLY; The ETag value to implement optimistic concurrency.
	ETag *string `json:"eTag,omitempty" azure:"ro"`

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

type SolutionEvaluatedOutput struct {
	InsightAttachmentResponse *InsightAttachment `json:"insightAttachmentResponse,omitempty"`
	InsightResponse           *Insight           `json:"insightResponse,omitempty"`
}

// SolutionInstallationRequest - Solution Installation Request Body.
type SolutionInstallationRequest struct {
	// Solution resource properties.
	Properties *SolutionProperties `json:"properties,omitempty"`
}

// SolutionListResponse - Paged response contains list of requested objects and a URL link to get the next set of results.
type SolutionListResponse struct {
	// Continuation link (absolute URI) to the next page of results in the list.
	NextLink *string `json:"nextLink,omitempty"`

	// Token used in retrieving the next page. If null, there are no additional pages.
	SkipToken *string `json:"$skipToken,omitempty"`

	// List of requested objects.
	Value []*Solution `json:"value,omitempty"`
}

// SolutionProperties - Solution resource properties.
type SolutionProperties struct {
	// REQUIRED; SaaS application Publisher Id.
	MarketplacePublisherID *string `json:"marketplacePublisherId,omitempty"`

	// REQUIRED; SaaS application Offer Id.
	OfferID *string `json:"offerId,omitempty"`

	// REQUIRED; SaaS application Plan Id.
	PlanID *string `json:"planId,omitempty"`

	// REQUIRED; SaaS subscriptionId of the installed SaaS application.
	SaasSubscriptionID *string `json:"saasSubscriptionId,omitempty"`

	// REQUIRED; SaaS subscription name of the installed SaaS application.
	SaasSubscriptionName *string `json:"saasSubscriptionName,omitempty"`

	// REQUIRED; SaaS application Term Id.
	TermID *string `json:"termId,omitempty"`

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]interface{}

	// READ-ONLY; Partner Id of the Solution.
	PartnerID *string `json:"partnerId,omitempty" azure:"ro"`

	// READ-ONLY; Solution Id.
	SolutionID *string `json:"solutionId,omitempty" azure:"ro"`
}

// SolutionsClientCreateOrUpdateOptions contains the optional parameters for the SolutionsClient.CreateOrUpdate method.
type SolutionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// SolutionsClientDeleteOptions contains the optional parameters for the SolutionsClient.Delete method.
type SolutionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// SolutionsClientGetOptions contains the optional parameters for the SolutionsClient.Get method.
type SolutionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// SolutionsClientListOptions contains the optional parameters for the SolutionsClient.List method.
type SolutionsClientListOptions struct {
	// Ids of the resource.
	IDs []string
	// Maximum creation date of resource (inclusive).
	MaxCreatedDateTime *time.Time
	// Maximum last modified date of resource (inclusive).
	MaxLastModifiedDateTime *time.Time
	// Maximum number of items needed (inclusive). Minimum = 10, Maximum = 1000, Default value = 50.
	MaxPageSize *int32
	// Minimum creation date of resource (inclusive).
	MinCreatedDateTime *time.Time
	// Minimum last modified date of resource (inclusive).
	MinLastModifiedDateTime *time.Time
	// Names of the resource.
	Names []string
	// Filters on key-value pairs within the Properties object. eg. "{testKey} eq {testValue}".
	PropertyFilters []string
	// Skip token for getting next set of results.
	SkipToken *string
	// Installed Solution ids.
	SolutionIDs []string
	// Statuses of the resource.
	Statuses []string
}

// SolutionsDiscoverabilityClientGetOptions contains the optional parameters for the SolutionsDiscoverabilityClient.Get method.
type SolutionsDiscoverabilityClientGetOptions struct {
	// placeholder for future optional parameters
}

// SolutionsDiscoverabilityClientListOptions contains the optional parameters for the SolutionsDiscoverabilityClient.List
// method.
type SolutionsDiscoverabilityClientListOptions struct {
	// Ids of FarmBeats Solutions which the customer requests to fetch.
	FarmBeatsSolutionIDs []string
	// Names of FarmBeats Solutions which the customer requests to fetch.
	FarmBeatsSolutionNames []string
	// Maximum number of items needed (inclusive). Minimum = 10, Maximum = 1000, Default value = 50.
	MaxPageSize *int32
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// UnitSystemsInfo - Unit systems info for the data provider.
type UnitSystemsInfo struct {
	// REQUIRED; UnitSystem key sent as part of ProviderInput.
	Key *string `json:"key,omitempty"`

	// REQUIRED; List of unit systems supported by this data provider.
	Values []*string `json:"values,omitempty"`
}
