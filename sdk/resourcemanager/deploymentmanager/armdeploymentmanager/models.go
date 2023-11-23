//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentmanager

import "time"

// APIKeyAuthentication - ApiKey authentication gives a name and a value that can be included in either the request header
// or query parameters.
type APIKeyAuthentication struct {
	// REQUIRED; The location of the authentication key/value pair in the request.
	In *RestAuthLocation

	// REQUIRED; The key name of the authentication key/value pair.
	Name *string

	// REQUIRED; The authentication type.
	Type *RestAuthType

	// REQUIRED; The value of the authentication key/value pair.
	Value *string
}

// GetRestRequestAuthentication implements the RestRequestAuthenticationClassification interface for type APIKeyAuthentication.
func (a *APIKeyAuthentication) GetRestRequestAuthentication() *RestRequestAuthentication {
	return &RestRequestAuthentication{
		Type: a.Type,
	}
}

// ArtifactSource - The resource that defines the source location where the artifacts are located.
type ArtifactSource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The properties that define the artifact source.
	Properties *ArtifactSourceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ArtifactSourceProperties - The properties that define the artifact source.
type ArtifactSourceProperties struct {
	// REQUIRED; The authentication method to use to access the artifact source.
	Authentication AuthenticationClassification

	// REQUIRED; The type of artifact source used.
	SourceType *string

	// The path from the location that the 'authentication' property [say, a SAS URI to the blob container] refers to, to the
	// location of the artifacts. This can be used to differentiate different versions
	// of the artifacts. Or, different types of artifacts like binaries or templates. The location referenced by the authentication
	// property concatenated with this optional artifactRoot path forms the
	// artifact source location where the artifacts are expected to be found.
	ArtifactRoot *string
}

// ArtifactSourcePropertiesAutoGenerated - The properties that define the source location where the artifacts are located.
type ArtifactSourcePropertiesAutoGenerated struct {
	// REQUIRED; The authentication method to use to access the artifact source.
	Authentication AuthenticationClassification

	// REQUIRED; The type of artifact source used.
	SourceType *string

	// The path from the location that the 'authentication' property [say, a SAS URI to the blob container] refers to, to the
	// location of the artifacts. This can be used to differentiate different versions
	// of the artifacts. Or, different types of artifacts like binaries or templates. The location referenced by the authentication
	// property concatenated with this optional artifactRoot path forms the
	// artifact source location where the artifacts are expected to be found.
	ArtifactRoot *string
}

// Authentication - Defines the authentication method and properties to access the artifacts.
type Authentication struct {
	// REQUIRED; The authentication type
	Type *string
}

// GetAuthentication implements the AuthenticationClassification interface for type Authentication.
func (a *Authentication) GetAuthentication() *Authentication { return a }

// CloudErrorBody - Detailed error information of any failure.
type CloudErrorBody struct {
	// More detailed error information.
	Details []*CloudErrorBody

	// Error target
	Target *string

	// READ-ONLY; Error code string.
	Code *string

	// READ-ONLY; Descriptive error information.
	Message *string
}

// HealthCheckStepAttributes - The attributes for the health check step.
type HealthCheckStepAttributes struct {
	// REQUIRED; The duration in ISO 8601 format for which the resource is expected to be continuously healthy. If maxElasticDuration
	// is specified, healthy state duration is enforced after the detection of first
	// healthy signal.
	HealthyStateDuration *string

	// REQUIRED; The type of health check.
	Type *string

	// The duration in ISO 8601 format for which the health check waits for the resource to become healthy. Health check fails
	// if it doesn't. Health check starts to enforce healthyStateDuration once resource
	// becomes healthy.
	MaxElasticDuration *string

	// The duration in ISO 8601 format for which health check waits idly without any checks.
	WaitDuration *string
}

// GetHealthCheckStepAttributes implements the HealthCheckStepAttributesClassification interface for type HealthCheckStepAttributes.
func (h *HealthCheckStepAttributes) GetHealthCheckStepAttributes() *HealthCheckStepAttributes {
	return h
}

// HealthCheckStepProperties - Defines the properties of a health check step.
type HealthCheckStepProperties struct {
	// REQUIRED; The health check step attributes
	Attributes HealthCheckStepAttributesClassification

	// REQUIRED; The type of step.
	StepType *StepType
}

// GetStepProperties implements the StepPropertiesClassification interface for type HealthCheckStepProperties.
func (h *HealthCheckStepProperties) GetStepProperties() *StepProperties {
	return &StepProperties{
		StepType: h.StepType,
	}
}

// Identity for the resource.
type Identity struct {
	// REQUIRED; The list of identities.
	IdentityIDs []*string

	// REQUIRED; The identity type.
	Type *string
}

// Message - Supplementary contextual messages during a rollout.
type Message struct {
	// READ-ONLY; The actual message text.
	Message *string

	// READ-ONLY; Time in UTC this message was provided.
	TimeStamp *time.Time
}

// Operation - Represents an operation that can be performed on the service.
type Operation struct {
	// The display name of the operation.
	Display *OperationDetail

	// The name of the operation.
	Name *string

	// The origin of the operation.
	Origin *string

	// The properties of the operation.
	Properties any
}

// OperationDetail - The detail about an operation.
type OperationDetail struct {
	// The description of the operation.
	Description *string

	// The name of the operation.
	Operation *string

	// The name of the provider that supports the operation.
	Provider *string

	// The resource type on which this operation can be performed.
	Resource *string
}

// OperationsList - The operations response.
type OperationsList struct {
	// The list of supported operations
	Value *Operation
}

// PrePostStep - The properties that define a step.
type PrePostStep struct {
	// REQUIRED; The resource Id of the step to be run.
	StepID *string
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ResourceOperation - Individual resource operation information.
type ResourceOperation struct {
	// Name of the resource as specified in the artifacts. For ARM resources, this is the name of the resource specified in the
	// template.
	ResourceName *string

	// Type of the resource as specified in the artifacts. For ARM resources, this is the type of the resource specified in the
	// template.
	ResourceType *string

	// READ-ONLY; Unique identifier of the operation. For ARM resources, this is the operationId obtained from ARM service.
	OperationID *string

	// READ-ONLY; State of the resource deployment. For ARM resources, this is the current provisioning state of the resource.
	ProvisioningState *string

	// READ-ONLY; Http status code of the operation.
	StatusCode *string

	// READ-ONLY; Descriptive information of the resource operation.
	StatusMessage *string
}

// RestHealthCheck - A REST based health check
type RestHealthCheck struct {
	// REQUIRED; A unique name for this check.
	Name *string

	// REQUIRED; The request to the health provider.
	Request *RestRequest

	// The expected response from the health provider. If no expected response is provided, the default is to expect the received
	// response to have an HTTP status code of 200 OK.
	Response *RestResponse
}

// RestHealthCheckStepAttributes - Defines the REST health check step properties.
type RestHealthCheckStepAttributes struct {
	// REQUIRED; The duration in ISO 8601 format for which the resource is expected to be continuously healthy. If maxElasticDuration
	// is specified, healthy state duration is enforced after the detection of first
	// healthy signal.
	HealthyStateDuration *string

	// REQUIRED; The type of health check.
	Type *string

	// The duration in ISO 8601 format for which the health check waits for the resource to become healthy. Health check fails
	// if it doesn't. Health check starts to enforce healthyStateDuration once resource
	// becomes healthy.
	MaxElasticDuration *string

	// The REST health check parameters.
	Properties *RestParameters

	// The duration in ISO 8601 format for which health check waits idly without any checks.
	WaitDuration *string
}

// GetHealthCheckStepAttributes implements the HealthCheckStepAttributesClassification interface for type RestHealthCheckStepAttributes.
func (r *RestHealthCheckStepAttributes) GetHealthCheckStepAttributes() *HealthCheckStepAttributes {
	return &HealthCheckStepAttributes{
		HealthyStateDuration: r.HealthyStateDuration,
		MaxElasticDuration:   r.MaxElasticDuration,
		Type:                 r.Type,
		WaitDuration:         r.WaitDuration,
	}
}

// RestParameters - The parameters for the REST health check.
type RestParameters struct {
	// REQUIRED; The list of checks that form the health check step.
	HealthChecks []*RestHealthCheck
}

// RestRequest - The properties that make up a REST request
type RestRequest struct {
	// REQUIRED; The authentication information required in the request to the health provider.
	Authentication RestRequestAuthenticationClassification

	// REQUIRED; The HTTP method to use for the request.
	Method *RestRequestMethod

	// REQUIRED; The HTTP URI to use for the request.
	URI *string
}

// RestRequestAuthentication - The authentication information required in the REST health check request to the health provider.
type RestRequestAuthentication struct {
	// REQUIRED; The authentication type.
	Type *RestAuthType
}

// GetRestRequestAuthentication implements the RestRequestAuthenticationClassification interface for type RestRequestAuthentication.
func (r *RestRequestAuthentication) GetRestRequestAuthentication() *RestRequestAuthentication {
	return r
}

// RestResponse - The properties that make up the expected REST response
type RestResponse struct {
	// The regular expressions to match the response content with.
	Regex *RestResponseRegex

	// The HTTP status codes expected in a successful health check response. The response is expected to match one of the given
	// status codes. If no expected status codes are provided, default expected status
	// code is 200 OK.
	SuccessStatusCodes []*string
}

// RestResponseRegex - The regular expressions to match the response content with.
type RestResponseRegex struct {
	// Indicates whether any or all of the expressions should match with the response content.
	MatchQuantifier *RestMatchQuantifier

	// The list of regular expressions.
	Matches []*string
}

// Rollout - Defines the rollout.
type Rollout struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Identity for the resource.
	Identity *Identity

	// The properties that define a rollout.
	Properties *RolloutProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// RolloutIdentityAuthentication - RolloutIdentity uses the user-assigned managed identity authentication context specified
// in the Identity property during rollout creation.
type RolloutIdentityAuthentication struct {
	// REQUIRED; The authentication type.
	Type *RestAuthType
}

// GetRestRequestAuthentication implements the RestRequestAuthenticationClassification interface for type RolloutIdentityAuthentication.
func (r *RolloutIdentityAuthentication) GetRestRequestAuthentication() *RestRequestAuthentication {
	return &RestRequestAuthentication{
		Type: r.Type,
	}
}

// RolloutOperationInfo - Detailed runtime information of the rollout.
type RolloutOperationInfo struct {
	// READ-ONLY; The start time of the rollout in UTC. This property will not be set if the rollout has not completed yet.
	EndTime *time.Time

	// READ-ONLY; The detailed error information for any failure.
	Error *CloudErrorBody

	// READ-ONLY; The ordinal count of the number of retry attempts on a rollout. 0 if no retries of the rollout have been performed.
	// If the rollout is updated with a PUT, this count is reset to 0.
	RetryAttempt *int32

	// READ-ONLY; True, if all steps that succeeded on the previous run/attempt were chosen to be skipped in this retry attempt.
	// False, otherwise.
	SkipSucceededOnRetry *bool

	// READ-ONLY; The start time of the rollout in UTC.
	StartTime *time.Time
}

// RolloutProperties - The properties that define a rollout.
type RolloutProperties struct {
	// REQUIRED; The version of the build being deployed.
	BuildVersion *string

	// REQUIRED; The list of step groups that define the orchestration.
	StepGroups []*StepGroup

	// REQUIRED; The resource Id of the service topology from which service units are being referenced in step groups to be deployed.
	TargetServiceTopologyID *string

	// The reference to the artifact source resource Id where the payload is located.
	ArtifactSourceID *string

	// READ-ONLY; Operational information of the rollout.
	OperationInfo *RolloutOperationInfo

	// READ-ONLY; The detailed information on the services being deployed.
	Services []*Service

	// READ-ONLY; The current status of the rollout.
	Status *string

	// READ-ONLY; The cardinal count of total number of retries performed on the rollout at a given time.
	TotalRetryAttempts *int32
}

// RolloutPropertiesAutoGenerated - Defines the properties of a rollout.
type RolloutPropertiesAutoGenerated struct {
	// READ-ONLY; Operational information of the rollout.
	OperationInfo *RolloutOperationInfo

	// READ-ONLY; The detailed information on the services being deployed.
	Services []*Service

	// READ-ONLY; The current status of the rollout.
	Status *string

	// READ-ONLY; The cardinal count of total number of retries performed on the rollout at a given time.
	TotalRetryAttempts *int32
}

// RolloutRequest - Defines the PUT rollout request body.
type RolloutRequest struct {
	// REQUIRED; Identity for the resource.
	Identity *Identity

	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; Defines the properties that make up a rollout request.
	Properties *RolloutRequestProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// RolloutRequestProperties - The properties for defining a rollout.
type RolloutRequestProperties struct {
	// REQUIRED; The version of the build being deployed.
	BuildVersion *string

	// REQUIRED; The list of step groups that define the orchestration.
	StepGroups []*StepGroup

	// REQUIRED; The resource Id of the service topology from which service units are being referenced in step groups to be deployed.
	TargetServiceTopologyID *string

	// The reference to the artifact source resource Id where the payload is located.
	ArtifactSourceID *string
}

// RolloutStep - Defines a specific step on a target service unit.
type RolloutStep struct {
	// REQUIRED; Name of the step.
	Name *string

	// The step group the current step is part of.
	StepGroup *string

	// READ-ONLY; Supplementary informative messages during rollout.
	Messages []*Message

	// READ-ONLY; Detailed information of specific action execution.
	OperationInfo *StepOperationInfo

	// READ-ONLY; Set of resource operations that were performed, if any, on an Azure resource.
	ResourceOperations []*ResourceOperation

	// READ-ONLY; Current state of the step.
	Status *string
}

// SasAuthentication - Defines the properties to access the artifacts using an Azure Storage SAS URI.
type SasAuthentication struct {
	// REQUIRED; The authentication type
	Type *string

	// The SAS properties
	Properties *SasProperties
}

// GetAuthentication implements the AuthenticationClassification interface for type SasAuthentication.
func (s *SasAuthentication) GetAuthentication() *Authentication {
	return &Authentication{
		Type: s.Type,
	}
}

// SasProperties - The properties that define SAS authentication.
type SasProperties struct {
	// REQUIRED; The SAS URI to the Azure Storage blob container. Any offset from the root of the container to where the artifacts
	// are located can be defined in the artifactRoot.
	SasURI *string
}

// Service - Defines a service.
type Service struct {
	// REQUIRED; The Azure location to which the resources in the service belong to or should be deployed to.
	TargetLocation *string

	// REQUIRED; The subscription to which the resources in the service belong to or should be deployed to.
	TargetSubscriptionID *string

	// Name of the service.
	Name *string

	// The detailed information about the units that make up the service.
	ServiceUnits []*ServiceUnit
}

// ServiceProperties - The properties of a service.
type ServiceProperties struct {
	// REQUIRED; The Azure location to which the resources in the service belong to or should be deployed to.
	TargetLocation *string

	// REQUIRED; The subscription to which the resources in the service belong to or should be deployed to.
	TargetSubscriptionID *string
}

// ServiceResource - The resource representation of a service in a service topology.
type ServiceResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The properties that define a service in a service topology.
	Properties *ServiceResourceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ServiceResourceProperties - The properties that define a service in a service topology.
type ServiceResourceProperties struct {
	// REQUIRED; The Azure location to which the resources in the service belong to or should be deployed to.
	TargetLocation *string

	// REQUIRED; The subscription to which the resources in the service belong to or should be deployed to.
	TargetSubscriptionID *string
}

// ServiceTopologyProperties - The properties of a service topology.
type ServiceTopologyProperties struct {
	// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
	ArtifactSourceID *string
}

// ServiceTopologyResource - The resource representation of a service topology.
type ServiceTopologyResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The properties that define the service topology.
	Properties *ServiceTopologyResourceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ServiceTopologyResourceProperties - The properties that define the service topology.
type ServiceTopologyResourceProperties struct {
	// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
	ArtifactSourceID *string
}

// ServiceUnit - Defines a service unit.
type ServiceUnit struct {
	// REQUIRED; Describes the type of ARM deployment to be performed on the resource.
	DeploymentMode *DeploymentMode

	// REQUIRED; The Azure Resource Group to which the resources in the service unit belong to or should be deployed to.
	TargetResourceGroup *string

	// The artifacts for the service unit.
	Artifacts *ServiceUnitArtifacts

	// Name of the service unit.
	Name *string

	// Detailed step information, if present.
	Steps []*RolloutStep
}

// ServiceUnitArtifacts - Defines the artifacts of a service unit.
type ServiceUnitArtifacts struct {
	// The path to the ARM parameters file relative to the artifact source.
	ParametersArtifactSourceRelativePath *string

	// The full URI of the ARM parameters file with the SAS token.
	ParametersURI *string

	// The path to the ARM template file relative to the artifact source.
	TemplateArtifactSourceRelativePath *string

	// The full URI of the ARM template file with the SAS token.
	TemplateURI *string
}

// ServiceUnitProperties - Defines the properties of a service unit.
type ServiceUnitProperties struct {
	// REQUIRED; Describes the type of ARM deployment to be performed on the resource.
	DeploymentMode *DeploymentMode

	// REQUIRED; The Azure Resource Group to which the resources in the service unit belong to or should be deployed to.
	TargetResourceGroup *string

	// The artifacts for the service unit.
	Artifacts *ServiceUnitArtifacts
}

// ServiceUnitResource - Represents the response of a service unit resource.
type ServiceUnitResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The properties that define the service unit.
	Properties *ServiceUnitResourceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ServiceUnitResourceProperties - The properties that define the service unit.
type ServiceUnitResourceProperties struct {
	// REQUIRED; Describes the type of ARM deployment to be performed on the resource.
	DeploymentMode *DeploymentMode

	// REQUIRED; The Azure Resource Group to which the resources in the service unit belong to or should be deployed to.
	TargetResourceGroup *string

	// The artifacts for the service unit.
	Artifacts *ServiceUnitArtifacts
}

// StepGroup - The properties that define a Step group in a rollout.
type StepGroup struct {
	// REQUIRED; The resource Id of service unit to be deployed. The service unit should be from the service topology referenced
	// in targetServiceTopologyId
	DeploymentTargetID *string

	// REQUIRED; The name of the step group.
	Name *string

	// The list of step group names on which this step group depends on.
	DependsOnStepGroups []*string

	// The list of steps to be run after deploying the target.
	PostDeploymentSteps []*PrePostStep

	// The list of steps to be run before deploying the target.
	PreDeploymentSteps []*PrePostStep
}

// StepOperationInfo - Detailed information of a specific step run.
type StepOperationInfo struct {
	// The errors, if any, for the action.
	Error *CloudErrorBody

	// READ-ONLY; Unique identifier to track the request for ARM-based resources.
	CorrelationID *string

	// READ-ONLY; The name of the ARM deployment initiated as part of the step.
	DeploymentName *string

	// READ-ONLY; End time of the action in UTC.
	EndTime *time.Time

	// READ-ONLY; Last time in UTC this operation was updated.
	LastUpdatedTime *time.Time

	// READ-ONLY; Start time of the action in UTC.
	StartTime *time.Time
}

// StepProperties - The properties of a step resource.
type StepProperties struct {
	// REQUIRED; The type of step.
	StepType *StepType
}

// GetStepProperties implements the StepPropertiesClassification interface for type StepProperties.
func (s *StepProperties) GetStepProperties() *StepProperties { return s }

// StepResource - The resource representation of a rollout step.
type StepResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The properties that define the step.
	Properties StepPropertiesClassification

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// WaitStepAttributes - The parameters for the wait step.
type WaitStepAttributes struct {
	// REQUIRED; The duration in ISO 8601 format of how long the wait should be.
	Duration *string
}

// WaitStepProperties - Defines the properties of a Wait step.
type WaitStepProperties struct {
	// REQUIRED; The Wait attributes
	Attributes *WaitStepAttributes

	// REQUIRED; The type of step.
	StepType *StepType
}

// GetStepProperties implements the StepPropertiesClassification interface for type WaitStepProperties.
func (w *WaitStepProperties) GetStepProperties() *StepProperties {
	return &StepProperties{
		StepType: w.StepType,
	}
}
