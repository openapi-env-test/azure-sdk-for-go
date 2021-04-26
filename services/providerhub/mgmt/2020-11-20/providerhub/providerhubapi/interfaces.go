package providerhubapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/providerhub/mgmt/2020-11-20/providerhub"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckinManifest(ctx context.Context, providerNamespace string, checkinManifestParams providerhub.CheckinManifestParams) (result providerhub.CheckinManifestInfo, err error)
	GenerateManifest(ctx context.Context, providerNamespace string) (result providerhub.ResourceProviderManifest, err error)
}

var _ BaseClientAPI = (*providerhub.BaseClient)(nil)

// CustomRolloutsClientAPI contains the set of methods on the CustomRolloutsClient type.
type CustomRolloutsClientAPI interface {
	CreateOrUpdate(ctx context.Context, providerNamespace string, rolloutName string, properties providerhub.CustomRollout) (result providerhub.CustomRollout, err error)
	Get(ctx context.Context, providerNamespace string, rolloutName string) (result providerhub.CustomRollout, err error)
	ListByProviderRegistration(ctx context.Context, providerNamespace string) (result providerhub.CustomRolloutArrayResponseWithContinuationPage, err error)
	ListByProviderRegistrationComplete(ctx context.Context, providerNamespace string) (result providerhub.CustomRolloutArrayResponseWithContinuationIterator, err error)
}

var _ CustomRolloutsClientAPI = (*providerhub.CustomRolloutsClient)(nil)

// DefaultRolloutsClientAPI contains the set of methods on the DefaultRolloutsClient type.
type DefaultRolloutsClientAPI interface {
	CreateOrUpdate(ctx context.Context, providerNamespace string, rolloutName string, properties providerhub.DefaultRollout) (result providerhub.DefaultRolloutsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, providerNamespace string, rolloutName string) (result autorest.Response, err error)
	Get(ctx context.Context, providerNamespace string, rolloutName string) (result providerhub.DefaultRollout, err error)
	ListByProviderRegistration(ctx context.Context, providerNamespace string) (result providerhub.DefaultRolloutArrayResponseWithContinuationPage, err error)
	ListByProviderRegistrationComplete(ctx context.Context, providerNamespace string) (result providerhub.DefaultRolloutArrayResponseWithContinuationIterator, err error)
	Stop(ctx context.Context, providerNamespace string, rolloutName string) (result autorest.Response, err error)
}

var _ DefaultRolloutsClientAPI = (*providerhub.DefaultRolloutsClient)(nil)

// NotificationRegistrationsClientAPI contains the set of methods on the NotificationRegistrationsClient type.
type NotificationRegistrationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, providerNamespace string, notificationRegistrationName string, properties providerhub.NotificationRegistration) (result providerhub.NotificationRegistration, err error)
	Delete(ctx context.Context, providerNamespace string, notificationRegistrationName string) (result autorest.Response, err error)
	Get(ctx context.Context, providerNamespace string, notificationRegistrationName string) (result providerhub.NotificationRegistration, err error)
	ListByProviderRegistration(ctx context.Context, providerNamespace string) (result providerhub.NotificationRegistrationArrayResponseWithContinuationPage, err error)
	ListByProviderRegistrationComplete(ctx context.Context, providerNamespace string) (result providerhub.NotificationRegistrationArrayResponseWithContinuationIterator, err error)
}

var _ NotificationRegistrationsClientAPI = (*providerhub.NotificationRegistrationsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, providerNamespace string, operationsPutContent providerhub.OperationsPutContent) (result providerhub.OperationsContent, err error)
	Delete(ctx context.Context, providerNamespace string) (result autorest.Response, err error)
	List(ctx context.Context) (result providerhub.OperationsDefinitionArrayResponseWithContinuationPage, err error)
	ListComplete(ctx context.Context) (result providerhub.OperationsDefinitionArrayResponseWithContinuationIterator, err error)
	ListByProviderRegistration(ctx context.Context, providerNamespace string) (result providerhub.ListOperationsDefinition, err error)
}

var _ OperationsClientAPI = (*providerhub.OperationsClient)(nil)

// ProviderRegistrationsClientAPI contains the set of methods on the ProviderRegistrationsClient type.
type ProviderRegistrationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, providerNamespace string, properties providerhub.ProviderRegistration) (result providerhub.ProviderRegistrationsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, providerNamespace string) (result autorest.Response, err error)
	GenerateOperations(ctx context.Context, providerNamespace string) (result providerhub.ListOperationsDefinition, err error)
	Get(ctx context.Context, providerNamespace string) (result providerhub.ProviderRegistration, err error)
	List(ctx context.Context) (result providerhub.ProviderRegistrationArrayResponseWithContinuationPage, err error)
	ListComplete(ctx context.Context) (result providerhub.ProviderRegistrationArrayResponseWithContinuationIterator, err error)
}

var _ ProviderRegistrationsClientAPI = (*providerhub.ProviderRegistrationsClient)(nil)

// ResourceTypeRegistrationsClientAPI contains the set of methods on the ResourceTypeRegistrationsClient type.
type ResourceTypeRegistrationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, providerNamespace string, resourceType string, properties providerhub.ResourceTypeRegistration) (result providerhub.ResourceTypeRegistrationsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, providerNamespace string, resourceType string) (result autorest.Response, err error)
	Get(ctx context.Context, providerNamespace string, resourceType string) (result providerhub.ResourceTypeRegistration, err error)
	ListByProviderRegistration(ctx context.Context, providerNamespace string) (result providerhub.ResourceTypeRegistrationArrayResponseWithContinuationPage, err error)
	ListByProviderRegistrationComplete(ctx context.Context, providerNamespace string) (result providerhub.ResourceTypeRegistrationArrayResponseWithContinuationIterator, err error)
}

var _ ResourceTypeRegistrationsClientAPI = (*providerhub.ResourceTypeRegistrationsClient)(nil)

// SkusClientAPI contains the set of methods on the SkusClient type.
type SkusClientAPI interface {
	CreateOrUpdate(ctx context.Context, providerNamespace string, resourceType string, sku string, properties providerhub.ResourceTypeSku) (result providerhub.SkuResource, err error)
	CreateOrUpdateNestedResourceTypeFirst(ctx context.Context, providerNamespace string, resourceType string, nestedResourceTypeFirst string, sku string, properties providerhub.ResourceTypeSku) (result providerhub.SkuResource, err error)
	CreateOrUpdateNestedResourceTypeSecond(ctx context.Context, providerNamespace string, resourceType string, nestedResourceTypeFirst string, nestedResourceTypeSecond string, sku string, properties providerhub.ResourceTypeSku) (result providerhub.SkuResource, err error)
	CreateOrUpdateNestedResourceTypeThird(ctx context.Context, providerNamespace string, resourceType string, nestedResourceTypeFirst string, nestedResourceTypeSecond string, nestedResourceTypeThird string, sku string, properties providerhub.ResourceTypeSku) (result providerhub.SkuResource, err error)
	Delete(ctx context.Context, providerNamespace string, resourceType string, sku string) (result autorest.Response, err error)
	DeleteNestedResourceTypeFirst(ctx context.Context, providerNamespace string, resourceType string, nestedResourceTypeFirst string, sku string) (result autorest.Response, err error)
	DeleteNestedResourceTypeSecond(ctx context.Context, providerNamespace string, resourceType string, nestedResourceTypeFirst string, nestedResourceTypeSecond string, sku string) (result autorest.Response, err error)
	DeleteNestedResourceTypeThird(ctx context.Context, providerNamespace string, resourceType string, nestedResourceTypeFirst string, nestedResourceTypeSecond string, nestedResourceTypeThird string, sku string) (result autorest.Response, err error)
	Get(ctx context.Context, providerNamespace string, resourceType string, sku string) (result providerhub.SkuResource, err error)
	GetNestedResourceTypeFirst(ctx context.Context, providerNamespace string, resourceType string, nestedResourceTypeFirst string, sku string) (result providerhub.SkuResource, err error)
	GetNestedResourceTypeSecond(ctx context.Context, providerNamespace string, resourceType string, nestedResourceTypeFirst string, nestedResourceTypeSecond string, sku string) (result providerhub.SkuResource, err error)
	GetNestedResourceTypeThird(ctx context.Context, providerNamespace string, resourceType string, nestedResourceTypeFirst string, nestedResourceTypeSecond string, nestedResourceTypeThird string, sku string) (result providerhub.SkuResource, err error)
	ListByResourceTypeRegistrations(ctx context.Context, providerNamespace string, resourceType string) (result providerhub.SkuResourceArrayResponseWithContinuationPage, err error)
	ListByResourceTypeRegistrationsComplete(ctx context.Context, providerNamespace string, resourceType string) (result providerhub.SkuResourceArrayResponseWithContinuationIterator, err error)
	ListByResourceTypeRegistrationsNestedResourceTypeFirst(ctx context.Context, providerNamespace string, resourceType string, nestedResourceTypeFirst string) (result providerhub.SkuResourceArrayResponseWithContinuationPage, err error)
	ListByResourceTypeRegistrationsNestedResourceTypeFirstComplete(ctx context.Context, providerNamespace string, resourceType string, nestedResourceTypeFirst string) (result providerhub.SkuResourceArrayResponseWithContinuationIterator, err error)
	ListByResourceTypeRegistrationsNestedResourceTypeSecond(ctx context.Context, providerNamespace string, resourceType string, nestedResourceTypeFirst string, nestedResourceTypeSecond string) (result providerhub.SkuResourceArrayResponseWithContinuationPage, err error)
	ListByResourceTypeRegistrationsNestedResourceTypeSecondComplete(ctx context.Context, providerNamespace string, resourceType string, nestedResourceTypeFirst string, nestedResourceTypeSecond string) (result providerhub.SkuResourceArrayResponseWithContinuationIterator, err error)
	ListByResourceTypeRegistrationsNestedResourceTypeThird(ctx context.Context, providerNamespace string, resourceType string, nestedResourceTypeFirst string, nestedResourceTypeSecond string, nestedResourceTypeThird string) (result providerhub.SkuResourceArrayResponseWithContinuationPage, err error)
	ListByResourceTypeRegistrationsNestedResourceTypeThirdComplete(ctx context.Context, providerNamespace string, resourceType string, nestedResourceTypeFirst string, nestedResourceTypeSecond string, nestedResourceTypeThird string) (result providerhub.SkuResourceArrayResponseWithContinuationIterator, err error)
}

var _ SkusClientAPI = (*providerhub.SkusClient)(nil)
