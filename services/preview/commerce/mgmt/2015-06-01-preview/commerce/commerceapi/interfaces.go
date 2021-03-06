package commerceapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/commerce/mgmt/2015-06-01-preview/commerce"
	"github.com/Azure/go-autorest/autorest/date"
)

// UsageAggregatesClientAPI contains the set of methods on the UsageAggregatesClient type.
type UsageAggregatesClientAPI interface {
	List(ctx context.Context, reportedStartTime date.Time, reportedEndTime date.Time, showDetails *bool, aggregationGranularity commerce.AggregationGranularity, continuationToken string) (result commerce.UsageAggregationListResultPage, err error)
	ListComplete(ctx context.Context, reportedStartTime date.Time, reportedEndTime date.Time, showDetails *bool, aggregationGranularity commerce.AggregationGranularity, continuationToken string) (result commerce.UsageAggregationListResultIterator, err error)
}

var _ UsageAggregatesClientAPI = (*commerce.UsageAggregatesClient)(nil)

// RateCardClientAPI contains the set of methods on the RateCardClient type.
type RateCardClientAPI interface {
	Get(ctx context.Context, filter string) (result commerce.ResourceRateCardInfo, err error)
}

var _ RateCardClientAPI = (*commerce.RateCardClient)(nil)
