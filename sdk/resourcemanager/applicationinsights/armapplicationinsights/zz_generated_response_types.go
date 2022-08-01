//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

import "net/http"

// WorkbooksClientCreateOrUpdateResponse contains the response from method WorkbooksClient.CreateOrUpdate.
type WorkbooksClientCreateOrUpdateResponse struct {
	WorkbooksClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkbooksClientCreateOrUpdateResult contains the result from method WorkbooksClient.CreateOrUpdate.
type WorkbooksClientCreateOrUpdateResult struct {
	Workbook
}

// WorkbooksClientDeleteResponse contains the response from method WorkbooksClient.Delete.
type WorkbooksClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkbooksClientGetResponse contains the response from method WorkbooksClient.Get.
type WorkbooksClientGetResponse struct {
	WorkbooksClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkbooksClientGetResult contains the result from method WorkbooksClient.Get.
type WorkbooksClientGetResult struct {
	Workbook
}

// WorkbooksClientListByResourceGroupResponse contains the response from method WorkbooksClient.ListByResourceGroup.
type WorkbooksClientListByResourceGroupResponse struct {
	WorkbooksClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkbooksClientListByResourceGroupResult contains the result from method WorkbooksClient.ListByResourceGroup.
type WorkbooksClientListByResourceGroupResult struct {
	WorkbooksListResult
}

// WorkbooksClientListBySubscriptionResponse contains the response from method WorkbooksClient.ListBySubscription.
type WorkbooksClientListBySubscriptionResponse struct {
	WorkbooksClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkbooksClientListBySubscriptionResult contains the result from method WorkbooksClient.ListBySubscription.
type WorkbooksClientListBySubscriptionResult struct {
	WorkbooksListResult
}

// WorkbooksClientRevisionGetResponse contains the response from method WorkbooksClient.RevisionGet.
type WorkbooksClientRevisionGetResponse struct {
	WorkbooksClientRevisionGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkbooksClientRevisionGetResult contains the result from method WorkbooksClient.RevisionGet.
type WorkbooksClientRevisionGetResult struct {
	Workbook
}

// WorkbooksClientRevisionsListResponse contains the response from method WorkbooksClient.RevisionsList.
type WorkbooksClientRevisionsListResponse struct {
	WorkbooksClientRevisionsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkbooksClientRevisionsListResult contains the result from method WorkbooksClient.RevisionsList.
type WorkbooksClientRevisionsListResult struct {
	WorkbooksListResult
}

// WorkbooksClientUpdateResponse contains the response from method WorkbooksClient.Update.
type WorkbooksClientUpdateResponse struct {
	WorkbooksClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// WorkbooksClientUpdateResult contains the result from method WorkbooksClient.Update.
type WorkbooksClientUpdateResult struct {
	Workbook
}
