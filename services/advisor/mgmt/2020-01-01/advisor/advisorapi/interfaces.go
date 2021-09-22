package advisorapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/advisor/mgmt/2020-01-01/advisor"
	"github.com/Azure/go-autorest/autorest"
	"github.com/gofrs/uuid"
)

// RecommendationMetadataClientAPI contains the set of methods on the RecommendationMetadataClient type.
type RecommendationMetadataClientAPI interface {
	Get(ctx context.Context, name string) (result advisor.MetadataEntity, err error)
	List(ctx context.Context) (result advisor.MetadataEntityListResultPage, err error)
	ListComplete(ctx context.Context) (result advisor.MetadataEntityListResultIterator, err error)
}

var _ RecommendationMetadataClientAPI = (*advisor.RecommendationMetadataClient)(nil)

// ConfigurationsClientAPI contains the set of methods on the ConfigurationsClient type.
type ConfigurationsClientAPI interface {
	CreateInResourceGroup(ctx context.Context, configContract advisor.ConfigData, resourceGroup string) (result advisor.ConfigData, err error)
	CreateInSubscription(ctx context.Context, configContract advisor.ConfigData) (result advisor.ConfigData, err error)
	ListByResourceGroup(ctx context.Context, resourceGroup string) (result advisor.ConfigurationListResult, err error)
	ListBySubscription(ctx context.Context) (result advisor.ConfigurationListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result advisor.ConfigurationListResultIterator, err error)
}

var _ ConfigurationsClientAPI = (*advisor.ConfigurationsClient)(nil)

// RecommendationsClientAPI contains the set of methods on the RecommendationsClient type.
type RecommendationsClientAPI interface {
	Generate(ctx context.Context) (result autorest.Response, err error)
	Get(ctx context.Context, resourceURI string, recommendationID string) (result advisor.ResourceRecommendationBase, err error)
	GetGenerateStatus(ctx context.Context, operationID uuid.UUID) (result autorest.Response, err error)
	List(ctx context.Context, filter string, top *int32, skipToken string) (result advisor.ResourceRecommendationBaseListResultPage, err error)
	ListComplete(ctx context.Context, filter string, top *int32, skipToken string) (result advisor.ResourceRecommendationBaseListResultIterator, err error)
}

var _ RecommendationsClientAPI = (*advisor.RecommendationsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result advisor.OperationEntityListResultPage, err error)
	ListComplete(ctx context.Context) (result advisor.OperationEntityListResultIterator, err error)
}

var _ OperationsClientAPI = (*advisor.OperationsClient)(nil)

// SuppressionsClientAPI contains the set of methods on the SuppressionsClient type.
type SuppressionsClientAPI interface {
	Create(ctx context.Context, resourceURI string, recommendationID string, name string, suppressionContract advisor.SuppressionContract) (result advisor.SuppressionContract, err error)
	Delete(ctx context.Context, resourceURI string, recommendationID string, name string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceURI string, recommendationID string, name string) (result advisor.SuppressionContract, err error)
	List(ctx context.Context, top *int32, skipToken string) (result advisor.SuppressionContractListResultPage, err error)
	ListComplete(ctx context.Context, top *int32, skipToken string) (result advisor.SuppressionContractListResultIterator, err error)
}

var _ SuppressionsClientAPI = (*advisor.SuppressionsClient)(nil)
