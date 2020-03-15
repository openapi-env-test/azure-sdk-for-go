package resourcegraphapi

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
	"github.com/Azure/azure-sdk-for-go/services/resourcegraph/mgmt/2019-04-01/resourcegraph"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	Resources(ctx context.Context, query resourcegraph.QueryRequest) (result resourcegraph.QueryResponse, err error)
}

var _ BaseClientAPI = (*resourcegraph.BaseClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result resourcegraph.OperationListResult, err error)
}

var _ OperationsClientAPI = (*resourcegraph.OperationsClient)(nil)

// GraphQueryClientAPI contains the set of methods on the GraphQueryClient type.
type GraphQueryClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, properties resourcegraph.GraphQueryResource) (result resourcegraph.GraphQueryResource, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result resourcegraph.GraphQueryResource, err error)
	List(ctx context.Context, resourceGroupName string) (result resourcegraph.GraphQueryListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result resourcegraph.GraphQueryListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, body resourcegraph.GraphQueryUpdateParameters) (result resourcegraph.GraphQueryResource, err error)
}

var _ GraphQueryClientAPI = (*resourcegraph.GraphQueryClient)(nil)
