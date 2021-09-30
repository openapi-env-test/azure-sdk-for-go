//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package containerinstance

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/containerinstance/mgmt/2021-07-01/containerinstance"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ContainerGroupIPAddressType = original.ContainerGroupIPAddressType

const (
	ContainerGroupIPAddressTypePrivate ContainerGroupIPAddressType = original.ContainerGroupIPAddressTypePrivate
	ContainerGroupIPAddressTypePublic  ContainerGroupIPAddressType = original.ContainerGroupIPAddressTypePublic
)

type ContainerGroupNetworkProtocol = original.ContainerGroupNetworkProtocol

const (
	ContainerGroupNetworkProtocolTCP ContainerGroupNetworkProtocol = original.ContainerGroupNetworkProtocolTCP
	ContainerGroupNetworkProtocolUDP ContainerGroupNetworkProtocol = original.ContainerGroupNetworkProtocolUDP
)

type ContainerGroupRestartPolicy = original.ContainerGroupRestartPolicy

const (
	ContainerGroupRestartPolicyAlways    ContainerGroupRestartPolicy = original.ContainerGroupRestartPolicyAlways
	ContainerGroupRestartPolicyNever     ContainerGroupRestartPolicy = original.ContainerGroupRestartPolicyNever
	ContainerGroupRestartPolicyOnFailure ContainerGroupRestartPolicy = original.ContainerGroupRestartPolicyOnFailure
)

type ContainerGroupSku = original.ContainerGroupSku

const (
	ContainerGroupSkuDedicated ContainerGroupSku = original.ContainerGroupSkuDedicated
	ContainerGroupSkuStandard  ContainerGroupSku = original.ContainerGroupSkuStandard
)

type ContainerNetworkProtocol = original.ContainerNetworkProtocol

const (
	ContainerNetworkProtocolTCP ContainerNetworkProtocol = original.ContainerNetworkProtocolTCP
	ContainerNetworkProtocolUDP ContainerNetworkProtocol = original.ContainerNetworkProtocolUDP
)

type GpuSku = original.GpuSku

const (
	GpuSkuK80  GpuSku = original.GpuSkuK80
	GpuSkuP100 GpuSku = original.GpuSkuP100
	GpuSkuV100 GpuSku = original.GpuSkuV100
)

type LogAnalyticsLogType = original.LogAnalyticsLogType

const (
	LogAnalyticsLogTypeContainerInsights     LogAnalyticsLogType = original.LogAnalyticsLogTypeContainerInsights
	LogAnalyticsLogTypeContainerInstanceLogs LogAnalyticsLogType = original.LogAnalyticsLogTypeContainerInstanceLogs
)

type OperatingSystemTypes = original.OperatingSystemTypes

const (
	OperatingSystemTypesLinux   OperatingSystemTypes = original.OperatingSystemTypesLinux
	OperatingSystemTypesWindows OperatingSystemTypes = original.OperatingSystemTypesWindows
)

type OperationsOrigin = original.OperationsOrigin

const (
	OperationsOriginSystem OperationsOrigin = original.OperationsOriginSystem
	OperationsOriginUser   OperationsOrigin = original.OperationsOriginUser
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone                       ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssignedUserAssigned
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = original.ResourceIdentityTypeUserAssigned
)

type Scheme = original.Scheme

const (
	SchemeHTTP  Scheme = original.SchemeHTTP
	SchemeHTTPS Scheme = original.SchemeHTTPS
)

type AzureFileVolume = original.AzureFileVolume
type BaseClient = original.BaseClient
type CachedImages = original.CachedImages
type CachedImagesListResult = original.CachedImagesListResult
type CachedImagesListResultIterator = original.CachedImagesListResultIterator
type CachedImagesListResultPage = original.CachedImagesListResultPage
type Capabilities = original.Capabilities
type CapabilitiesCapabilities = original.CapabilitiesCapabilities
type CapabilitiesListResult = original.CapabilitiesListResult
type CapabilitiesListResultIterator = original.CapabilitiesListResultIterator
type CapabilitiesListResultPage = original.CapabilitiesListResultPage
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type Container = original.Container
type ContainerAttachResponse = original.ContainerAttachResponse
type ContainerExec = original.ContainerExec
type ContainerExecRequest = original.ContainerExecRequest
type ContainerExecRequestTerminalSize = original.ContainerExecRequestTerminalSize
type ContainerExecResponse = original.ContainerExecResponse
type ContainerGroup = original.ContainerGroup
type ContainerGroupDiagnostics = original.ContainerGroupDiagnostics
type ContainerGroupIdentity = original.ContainerGroupIdentity
type ContainerGroupIdentityUserAssignedIdentitiesValue = original.ContainerGroupIdentityUserAssignedIdentitiesValue
type ContainerGroupListResult = original.ContainerGroupListResult
type ContainerGroupListResultIterator = original.ContainerGroupListResultIterator
type ContainerGroupListResultPage = original.ContainerGroupListResultPage
type ContainerGroupProperties = original.ContainerGroupProperties
type ContainerGroupPropertiesInstanceView = original.ContainerGroupPropertiesInstanceView
type ContainerGroupSubnetID = original.ContainerGroupSubnetID
type ContainerGroupsClient = original.ContainerGroupsClient
type ContainerGroupsCreateOrUpdateFuture = original.ContainerGroupsCreateOrUpdateFuture
type ContainerGroupsDeleteFuture = original.ContainerGroupsDeleteFuture
type ContainerGroupsRestartFuture = original.ContainerGroupsRestartFuture
type ContainerGroupsStartFuture = original.ContainerGroupsStartFuture
type ContainerHTTPGet = original.ContainerHTTPGet
type ContainerPort = original.ContainerPort
type ContainerProbe = original.ContainerProbe
type ContainerProperties = original.ContainerProperties
type ContainerPropertiesInstanceView = original.ContainerPropertiesInstanceView
type ContainerState = original.ContainerState
type ContainersClient = original.ContainersClient
type DNSConfiguration = original.DNSConfiguration
type EncryptionProperties = original.EncryptionProperties
type EnvironmentVariable = original.EnvironmentVariable
type Event = original.Event
type GitRepoVolume = original.GitRepoVolume
type GpuResource = original.GpuResource
type HTTPHeader = original.HTTPHeader
type IPAddress = original.IPAddress
type ImageRegistryCredential = original.ImageRegistryCredential
type InitContainerDefinition = original.InitContainerDefinition
type InitContainerPropertiesDefinition = original.InitContainerPropertiesDefinition
type InitContainerPropertiesDefinitionInstanceView = original.InitContainerPropertiesDefinitionInstanceView
type ListString = original.ListString
type LocationClient = original.LocationClient
type LogAnalytics = original.LogAnalytics
type Logs = original.Logs
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type Port = original.Port
type Resource = original.Resource
type ResourceLimits = original.ResourceLimits
type ResourceRequests = original.ResourceRequests
type ResourceRequirements = original.ResourceRequirements
type Usage = original.Usage
type UsageListResult = original.UsageListResult
type UsageName = original.UsageName
type Volume = original.Volume
type VolumeMount = original.VolumeMount

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewCachedImagesListResultIterator(page CachedImagesListResultPage) CachedImagesListResultIterator {
	return original.NewCachedImagesListResultIterator(page)
}
func NewCachedImagesListResultPage(cur CachedImagesListResult, getNextPage func(context.Context, CachedImagesListResult) (CachedImagesListResult, error)) CachedImagesListResultPage {
	return original.NewCachedImagesListResultPage(cur, getNextPage)
}
func NewCapabilitiesListResultIterator(page CapabilitiesListResultPage) CapabilitiesListResultIterator {
	return original.NewCapabilitiesListResultIterator(page)
}
func NewCapabilitiesListResultPage(cur CapabilitiesListResult, getNextPage func(context.Context, CapabilitiesListResult) (CapabilitiesListResult, error)) CapabilitiesListResultPage {
	return original.NewCapabilitiesListResultPage(cur, getNextPage)
}
func NewContainerGroupListResultIterator(page ContainerGroupListResultPage) ContainerGroupListResultIterator {
	return original.NewContainerGroupListResultIterator(page)
}
func NewContainerGroupListResultPage(cur ContainerGroupListResult, getNextPage func(context.Context, ContainerGroupListResult) (ContainerGroupListResult, error)) ContainerGroupListResultPage {
	return original.NewContainerGroupListResultPage(cur, getNextPage)
}
func NewContainerGroupsClient(subscriptionID string) ContainerGroupsClient {
	return original.NewContainerGroupsClient(subscriptionID)
}
func NewContainerGroupsClientWithBaseURI(baseURI string, subscriptionID string) ContainerGroupsClient {
	return original.NewContainerGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewContainersClient(subscriptionID string) ContainersClient {
	return original.NewContainersClient(subscriptionID)
}
func NewContainersClientWithBaseURI(baseURI string, subscriptionID string) ContainersClient {
	return original.NewContainersClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocationClient(subscriptionID string) LocationClient {
	return original.NewLocationClient(subscriptionID)
}
func NewLocationClientWithBaseURI(baseURI string, subscriptionID string) LocationClient {
	return original.NewLocationClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleContainerGroupIPAddressTypeValues() []ContainerGroupIPAddressType {
	return original.PossibleContainerGroupIPAddressTypeValues()
}
func PossibleContainerGroupNetworkProtocolValues() []ContainerGroupNetworkProtocol {
	return original.PossibleContainerGroupNetworkProtocolValues()
}
func PossibleContainerGroupRestartPolicyValues() []ContainerGroupRestartPolicy {
	return original.PossibleContainerGroupRestartPolicyValues()
}
func PossibleContainerGroupSkuValues() []ContainerGroupSku {
	return original.PossibleContainerGroupSkuValues()
}
func PossibleContainerNetworkProtocolValues() []ContainerNetworkProtocol {
	return original.PossibleContainerNetworkProtocolValues()
}
func PossibleGpuSkuValues() []GpuSku {
	return original.PossibleGpuSkuValues()
}
func PossibleLogAnalyticsLogTypeValues() []LogAnalyticsLogType {
	return original.PossibleLogAnalyticsLogTypeValues()
}
func PossibleOperatingSystemTypesValues() []OperatingSystemTypes {
	return original.PossibleOperatingSystemTypesValues()
}
func PossibleOperationsOriginValues() []OperationsOrigin {
	return original.PossibleOperationsOriginValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleSchemeValues() []Scheme {
	return original.PossibleSchemeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
