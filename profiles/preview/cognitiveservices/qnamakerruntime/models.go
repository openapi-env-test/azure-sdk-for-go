//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package qnamakerruntime

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v4.0/qnamakerruntime"

type ErrorCodeType = original.ErrorCodeType

const (
	BadArgument       ErrorCodeType = original.BadArgument
	EndpointKeysError ErrorCodeType = original.EndpointKeysError
	ExtractionFailure ErrorCodeType = original.ExtractionFailure
	Forbidden         ErrorCodeType = original.Forbidden
	KbNotFound        ErrorCodeType = original.KbNotFound
	NotFound          ErrorCodeType = original.NotFound
	OperationNotFound ErrorCodeType = original.OperationNotFound
	QnaRuntimeError   ErrorCodeType = original.QnaRuntimeError
	QuotaExceeded     ErrorCodeType = original.QuotaExceeded
	ServiceError      ErrorCodeType = original.ServiceError
	SKULimitExceeded  ErrorCodeType = original.SKULimitExceeded
	Unauthorized      ErrorCodeType = original.Unauthorized
	Unspecified       ErrorCodeType = original.Unspecified
	ValidationFailure ErrorCodeType = original.ValidationFailure
)

type StrictFiltersCompoundOperationType = original.StrictFiltersCompoundOperationType

const (
	AND StrictFiltersCompoundOperationType = original.AND
	OR  StrictFiltersCompoundOperationType = original.OR
)

type BaseClient = original.BaseClient
type ContextDTO = original.ContextDTO
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type ErrorResponseError = original.ErrorResponseError
type FeedbackRecordDTO = original.FeedbackRecordDTO
type FeedbackRecordsDTO = original.FeedbackRecordsDTO
type InnerErrorModel = original.InnerErrorModel
type MetadataDTO = original.MetadataDTO
type PromptDTO = original.PromptDTO
type PromptDTOQna = original.PromptDTOQna
type QnADTO = original.QnADTO
type QnADTOContext = original.QnADTOContext
type QnASearchResult = original.QnASearchResult
type QnASearchResultContext = original.QnASearchResultContext
type QnASearchResultList = original.QnASearchResultList
type QueryContextDTO = original.QueryContextDTO
type QueryDTO = original.QueryDTO
type QueryDTOContext = original.QueryDTOContext
type RuntimeClient = original.RuntimeClient

func New(runtimeEndpoint string) BaseClient {
	return original.New(runtimeEndpoint)
}
func NewRuntimeClient(runtimeEndpoint string) RuntimeClient {
	return original.NewRuntimeClient(runtimeEndpoint)
}
func NewWithoutDefaults(runtimeEndpoint string) BaseClient {
	return original.NewWithoutDefaults(runtimeEndpoint)
}
func PossibleErrorCodeTypeValues() []ErrorCodeType {
	return original.PossibleErrorCodeTypeValues()
}
func PossibleStrictFiltersCompoundOperationTypeValues() []StrictFiltersCompoundOperationType {
	return original.PossibleStrictFiltersCompoundOperationTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
