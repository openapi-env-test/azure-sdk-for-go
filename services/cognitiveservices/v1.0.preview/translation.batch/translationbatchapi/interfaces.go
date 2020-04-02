package translationbatchapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0.preview/translation.batch"
	"github.com/satori/go.uuid"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	BatchDocumentsSubmit(ctx context.Context, body *translationbatch.JobSubmissionBatchRequest) (result translationbatch.SetObject, err error)
	CancelOperation(ctx context.Context, ID uuid.UUID) (result translationbatch.SetObject, err error)
	GetDetailedDocumentStatus(ctx context.Context, ID uuid.UUID, top *int32, skip *int32) (result translationbatch.SetObject, err error)
	GetDocumentFormats(ctx context.Context) (result translationbatch.SetListDocumentFormat, err error)
	GetDocumentStatus(ctx context.Context, documentID int64) (result translationbatch.SetObject, err error)
	GetOperationStatus(ctx context.Context, ID uuid.UUID) (result translationbatch.SetObject, err error)
	GetStorageSourcesThatWeCurrentlySupport(ctx context.Context) (result translationbatch.SetListString, err error)
	SingleDocumentUpload(ctx context.Context, toParameter string, from string, category string) (result translationbatch.SetObject, err error)
}

var _ BaseClientAPI = (*translationbatch.BaseClient)(nil)
