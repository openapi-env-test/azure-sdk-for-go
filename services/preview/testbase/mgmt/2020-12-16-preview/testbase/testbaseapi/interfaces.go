package testbaseapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/testbase/mgmt/2020-12-16-preview/testbase"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	AccountGetFileUploadURL(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters *testbase.GetFileUploadURLParameters) (result testbase.FileUploadURLResponse, err error)
	CheckPackageNameAvailability(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters testbase.PackageCheckNameAvailabilityParameters) (result testbase.CheckNameAvailabilityResult, err error)
	PackageGetDownloadURL(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string) (result testbase.DownloadURLResponse, err error)
	TestResultGetDownloadURL(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, testResultName string) (result testbase.DownloadURLResponse, err error)
	TestResultGetVideoDownloadURL(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, testResultName string) (result testbase.DownloadURLResponse, err error)
}

var _ BaseClientAPI = (*testbase.BaseClient)(nil)

// SKUsClientAPI contains the set of methods on the SKUsClient type.
type SKUsClientAPI interface {
	List(ctx context.Context) (result testbase.AccountSKUListResultPage, err error)
	ListComplete(ctx context.Context) (result testbase.AccountSKUListResultIterator, err error)
}

var _ SKUsClientAPI = (*testbase.SKUsClient)(nil)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	ListByResourceGroup(ctx context.Context, resourceGroupName string, getDeleted *bool) (result testbase.AccountListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, getDeleted *bool) (result testbase.AccountListResultIterator, err error)
	ListBySubscription(ctx context.Context, getDeleted *bool) (result testbase.AccountListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context, getDeleted *bool) (result testbase.AccountListResultIterator, err error)
}

var _ AccountsClientAPI = (*testbase.AccountsClient)(nil)

// AccountClientAPI contains the set of methods on the AccountClient type.
type AccountClientAPI interface {
	Create(ctx context.Context, parameters testbase.AccountResource, resourceGroupName string, testBaseAccountName string, restore *bool) (result testbase.AccountCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, testBaseAccountName string) (result testbase.AccountDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, testBaseAccountName string) (result testbase.AccountResource, err error)
	Offboard(ctx context.Context, resourceGroupName string, testBaseAccountName string) (result testbase.AccountOffboardFuture, err error)
	Update(ctx context.Context, parameters testbase.AccountUpdateParameters, resourceGroupName string, testBaseAccountName string) (result testbase.AccountUpdateFuture, err error)
}

var _ AccountClientAPI = (*testbase.AccountClient)(nil)

// AccountUsageClientAPI contains the set of methods on the AccountUsageClient type.
type AccountUsageClientAPI interface {
	List(ctx context.Context, resourceGroupName string, testBaseAccountName string, filter string) (result testbase.AccountUsageDataListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, testBaseAccountName string, filter string) (result testbase.AccountUsageDataListIterator, err error)
}

var _ AccountUsageClientAPI = (*testbase.AccountUsageClient)(nil)

// AccountAvailableOSsClientAPI contains the set of methods on the AccountAvailableOSsClient type.
type AccountAvailableOSsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, testBaseAccountName string, osUpdateType testbase.OsUpdateType) (result testbase.AvailableOSListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, testBaseAccountName string, osUpdateType testbase.OsUpdateType) (result testbase.AvailableOSListResultIterator, err error)
}

var _ AccountAvailableOSsClientAPI = (*testbase.AccountAvailableOSsClient)(nil)

// AccountAvailableOSClientAPI contains the set of methods on the AccountAvailableOSClient type.
type AccountAvailableOSClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, availableOSResourceName string) (result testbase.AvailableOSResource, err error)
}

var _ AccountAvailableOSClientAPI = (*testbase.AccountAvailableOSClient)(nil)

// AccountFlightingRingsClientAPI contains the set of methods on the AccountFlightingRingsClient type.
type AccountFlightingRingsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, testBaseAccountName string) (result testbase.FlightingRingListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, testBaseAccountName string) (result testbase.FlightingRingListResultIterator, err error)
}

var _ AccountFlightingRingsClientAPI = (*testbase.AccountFlightingRingsClient)(nil)

// AccountFlightingRingClientAPI contains the set of methods on the AccountFlightingRingClient type.
type AccountFlightingRingClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, flightingRingResourceName string) (result testbase.FlightingRingResource, err error)
}

var _ AccountFlightingRingClientAPI = (*testbase.AccountFlightingRingClient)(nil)

// AccountTestTypesClientAPI contains the set of methods on the AccountTestTypesClient type.
type AccountTestTypesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, testBaseAccountName string) (result testbase.TestTypeListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, testBaseAccountName string) (result testbase.TestTypeListResultIterator, err error)
}

var _ AccountTestTypesClientAPI = (*testbase.AccountTestTypesClient)(nil)

// AccountTestTypeClientAPI contains the set of methods on the AccountTestTypeClient type.
type AccountTestTypeClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, testTypeResourceName string) (result testbase.TestTypeResource, err error)
}

var _ AccountTestTypeClientAPI = (*testbase.AccountTestTypeClient)(nil)

// PackagesClientAPI contains the set of methods on the PackagesClient type.
type PackagesClientAPI interface {
	ListByTestBaseAccount(ctx context.Context, resourceGroupName string, testBaseAccountName string) (result testbase.PackageListResultPage, err error)
	ListByTestBaseAccountComplete(ctx context.Context, resourceGroupName string, testBaseAccountName string) (result testbase.PackageListResultIterator, err error)
}

var _ PackagesClientAPI = (*testbase.PackagesClient)(nil)

// PackageClientAPI contains the set of methods on the PackageClient type.
type PackageClientAPI interface {
	Create(ctx context.Context, parameters testbase.PackageResource, resourceGroupName string, testBaseAccountName string, packageName string) (result testbase.PackageCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string) (result testbase.PackageDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string) (result testbase.PackageResource, err error)
	HardDelete(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string) (result testbase.PackageHardDeleteFuture, err error)
	Update(ctx context.Context, parameters testbase.PackageUpdateParameters, resourceGroupName string, testBaseAccountName string, packageName string) (result testbase.PackageUpdateFuture, err error)
}

var _ PackageClientAPI = (*testbase.PackageClient)(nil)

// TestSummariesClientAPI contains the set of methods on the TestSummariesClient type.
type TestSummariesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, testBaseAccountName string) (result testbase.TestSummaryListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, testBaseAccountName string) (result testbase.TestSummaryListResultIterator, err error)
}

var _ TestSummariesClientAPI = (*testbase.TestSummariesClient)(nil)

// TestSummaryClientAPI contains the set of methods on the TestSummaryClient type.
type TestSummaryClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, testSummaryName string) (result testbase.TestSummaryResource, err error)
}

var _ TestSummaryClientAPI = (*testbase.TestSummaryClient)(nil)

// TestResultsClientAPI contains the set of methods on the TestResultsClient type.
type TestResultsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, osUpdateType testbase.OsUpdateType, filter string) (result testbase.TestResultListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, osUpdateType testbase.OsUpdateType, filter string) (result testbase.TestResultListResultIterator, err error)
}

var _ TestResultsClientAPI = (*testbase.TestResultsClient)(nil)

// TestResultClientAPI contains the set of methods on the TestResultClient type.
type TestResultClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, testResultName string) (result testbase.TestResultResource, err error)
}

var _ TestResultClientAPI = (*testbase.TestResultClient)(nil)

// OSUpdatesClientAPI contains the set of methods on the OSUpdatesClient type.
type OSUpdatesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, osUpdateType testbase.OsUpdateType) (result testbase.OSUpdateListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, osUpdateType testbase.OsUpdateType) (result testbase.OSUpdateListResultIterator, err error)
}

var _ OSUpdatesClientAPI = (*testbase.OSUpdatesClient)(nil)

// OSUpdateClientAPI contains the set of methods on the OSUpdateClient type.
type OSUpdateClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, osUpdateResourceName string) (result testbase.OSUpdateResource, err error)
}

var _ OSUpdateClientAPI = (*testbase.OSUpdateClient)(nil)

// FavoriteProcessesClientAPI contains the set of methods on the FavoriteProcessesClient type.
type FavoriteProcessesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string) (result testbase.FavoriteProcessListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string) (result testbase.FavoriteProcessListResultIterator, err error)
}

var _ FavoriteProcessesClientAPI = (*testbase.FavoriteProcessesClient)(nil)

// FavoriteProcessClientAPI contains the set of methods on the FavoriteProcessClient type.
type FavoriteProcessClientAPI interface {
	Create(ctx context.Context, parameters testbase.FavoriteProcessResource, resourceGroupName string, testBaseAccountName string, packageName string, favoriteProcessResourceName string) (result testbase.FavoriteProcessResource, err error)
	Delete(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, favoriteProcessResourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, favoriteProcessResourceName string) (result testbase.FavoriteProcessResource, err error)
}

var _ FavoriteProcessClientAPI = (*testbase.FavoriteProcessClient)(nil)

// AnalysisResultsClientAPI contains the set of methods on the AnalysisResultsClient type.
type AnalysisResultsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, testResultName string, analysisResultType testbase.AnalysisResultType) (result testbase.AnalysisResultListResult, err error)
}

var _ AnalysisResultsClientAPI = (*testbase.AnalysisResultsClient)(nil)

// AnalysisResultClientAPI contains the set of methods on the AnalysisResultClient type.
type AnalysisResultClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, testResultName string, analysisResultName testbase.AnalysisResultName) (result testbase.AnalysisResultSingletonResource, err error)
}

var _ AnalysisResultClientAPI = (*testbase.AnalysisResultClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result testbase.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result testbase.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*testbase.OperationsClient)(nil)
