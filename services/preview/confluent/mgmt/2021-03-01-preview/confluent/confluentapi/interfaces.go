package confluentapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/confluent/mgmt/2021-03-01-preview/confluent"
)

// MarketplaceAgreementsClientAPI contains the set of methods on the MarketplaceAgreementsClient type.
type MarketplaceAgreementsClientAPI interface {
	Create(ctx context.Context, body *confluent.AgreementResource) (result confluent.AgreementResource, err error)
	List(ctx context.Context) (result confluent.AgreementResourceListResponsePage, err error)
	ListComplete(ctx context.Context) (result confluent.AgreementResourceListResponseIterator, err error)
}

var _ MarketplaceAgreementsClientAPI = (*confluent.MarketplaceAgreementsClient)(nil)

// OrganizationOperationsClientAPI contains the set of methods on the OrganizationOperationsClient type.
type OrganizationOperationsClientAPI interface {
	List(ctx context.Context) (result confluent.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result confluent.OperationListResultIterator, err error)
}

var _ OrganizationOperationsClientAPI = (*confluent.OrganizationOperationsClient)(nil)

// OrganizationClientAPI contains the set of methods on the OrganizationClient type.
type OrganizationClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, organizationName string) (result confluent.OrganizationDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, organizationName string) (result confluent.OrganizationResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result confluent.OrganizationResourceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result confluent.OrganizationResourceListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result confluent.OrganizationResourceListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result confluent.OrganizationResourceListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, organizationName string, body *confluent.OrganizationResourceUpdate) (result confluent.OrganizationResource, err error)
}

var _ OrganizationClientAPI = (*confluent.OrganizationClient)(nil)

// ValidationsClientAPI contains the set of methods on the ValidationsClient type.
type ValidationsClientAPI interface {
	ValidateOrganization(ctx context.Context, resourceGroupName string, organizationName string, body confluent.OrganizationResource) (result confluent.OrganizationResource, err error)
}

var _ ValidationsClientAPI = (*confluent.ValidationsClient)(nil)
