// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package keyvault

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessPolicyUpdateKind = original.AccessPolicyUpdateKind

const (
	Add     AccessPolicyUpdateKind = original.Add
	Remove  AccessPolicyUpdateKind = original.Remove
	Replace AccessPolicyUpdateKind = original.Replace
)

type CertificatePermissions = original.CertificatePermissions

const (
	Backup         CertificatePermissions = original.Backup
	Create         CertificatePermissions = original.Create
	Delete         CertificatePermissions = original.Delete
	Deleteissuers  CertificatePermissions = original.Deleteissuers
	Get            CertificatePermissions = original.Get
	Getissuers     CertificatePermissions = original.Getissuers
	Import         CertificatePermissions = original.Import
	List           CertificatePermissions = original.List
	Listissuers    CertificatePermissions = original.Listissuers
	Managecontacts CertificatePermissions = original.Managecontacts
	Manageissuers  CertificatePermissions = original.Manageissuers
	Purge          CertificatePermissions = original.Purge
	Recover        CertificatePermissions = original.Recover
	Restore        CertificatePermissions = original.Restore
	Setissuers     CertificatePermissions = original.Setissuers
	Update         CertificatePermissions = original.Update
)

type CreateMode = original.CreateMode

const (
	CreateModeDefault CreateMode = original.CreateModeDefault
	CreateModeRecover CreateMode = original.CreateModeRecover
)

type KeyPermissions = original.KeyPermissions

const (
	KeyPermissionsBackup    KeyPermissions = original.KeyPermissionsBackup
	KeyPermissionsCreate    KeyPermissions = original.KeyPermissionsCreate
	KeyPermissionsDecrypt   KeyPermissions = original.KeyPermissionsDecrypt
	KeyPermissionsDelete    KeyPermissions = original.KeyPermissionsDelete
	KeyPermissionsEncrypt   KeyPermissions = original.KeyPermissionsEncrypt
	KeyPermissionsGet       KeyPermissions = original.KeyPermissionsGet
	KeyPermissionsImport    KeyPermissions = original.KeyPermissionsImport
	KeyPermissionsList      KeyPermissions = original.KeyPermissionsList
	KeyPermissionsPurge     KeyPermissions = original.KeyPermissionsPurge
	KeyPermissionsRecover   KeyPermissions = original.KeyPermissionsRecover
	KeyPermissionsRestore   KeyPermissions = original.KeyPermissionsRestore
	KeyPermissionsSign      KeyPermissions = original.KeyPermissionsSign
	KeyPermissionsUnwrapKey KeyPermissions = original.KeyPermissionsUnwrapKey
	KeyPermissionsUpdate    KeyPermissions = original.KeyPermissionsUpdate
	KeyPermissionsVerify    KeyPermissions = original.KeyPermissionsVerify
	KeyPermissionsWrapKey   KeyPermissions = original.KeyPermissionsWrapKey
)

type NetworkRuleAction = original.NetworkRuleAction

const (
	Allow NetworkRuleAction = original.Allow
	Deny  NetworkRuleAction = original.Deny
)

type NetworkRuleBypassOptions = original.NetworkRuleBypassOptions

const (
	AzureServices NetworkRuleBypassOptions = original.AzureServices
	None          NetworkRuleBypassOptions = original.None
)

type PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningState

const (
	Creating  PrivateEndpointConnectionProvisioningState = original.Creating
	Deleting  PrivateEndpointConnectionProvisioningState = original.Deleting
	Failed    PrivateEndpointConnectionProvisioningState = original.Failed
	Succeeded PrivateEndpointConnectionProvisioningState = original.Succeeded
)

type PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatus

const (
	Approved     PrivateEndpointServiceConnectionStatus = original.Approved
	Disconnected PrivateEndpointServiceConnectionStatus = original.Disconnected
	Pending      PrivateEndpointServiceConnectionStatus = original.Pending
	Rejected     PrivateEndpointServiceConnectionStatus = original.Rejected
)

type Reason = original.Reason

const (
	AccountNameInvalid Reason = original.AccountNameInvalid
	AlreadyExists      Reason = original.AlreadyExists
)

type SecretPermissions = original.SecretPermissions

const (
	SecretPermissionsBackup  SecretPermissions = original.SecretPermissionsBackup
	SecretPermissionsDelete  SecretPermissions = original.SecretPermissionsDelete
	SecretPermissionsGet     SecretPermissions = original.SecretPermissionsGet
	SecretPermissionsList    SecretPermissions = original.SecretPermissionsList
	SecretPermissionsPurge   SecretPermissions = original.SecretPermissionsPurge
	SecretPermissionsRecover SecretPermissions = original.SecretPermissionsRecover
	SecretPermissionsRestore SecretPermissions = original.SecretPermissionsRestore
	SecretPermissionsSet     SecretPermissions = original.SecretPermissionsSet
)

type SkuName = original.SkuName

const (
	Premium  SkuName = original.Premium
	Standard SkuName = original.Standard
)

type StoragePermissions = original.StoragePermissions

const (
	StoragePermissionsBackup        StoragePermissions = original.StoragePermissionsBackup
	StoragePermissionsDelete        StoragePermissions = original.StoragePermissionsDelete
	StoragePermissionsDeletesas     StoragePermissions = original.StoragePermissionsDeletesas
	StoragePermissionsGet           StoragePermissions = original.StoragePermissionsGet
	StoragePermissionsGetsas        StoragePermissions = original.StoragePermissionsGetsas
	StoragePermissionsList          StoragePermissions = original.StoragePermissionsList
	StoragePermissionsListsas       StoragePermissions = original.StoragePermissionsListsas
	StoragePermissionsPurge         StoragePermissions = original.StoragePermissionsPurge
	StoragePermissionsRecover       StoragePermissions = original.StoragePermissionsRecover
	StoragePermissionsRegeneratekey StoragePermissions = original.StoragePermissionsRegeneratekey
	StoragePermissionsRestore       StoragePermissions = original.StoragePermissionsRestore
	StoragePermissionsSet           StoragePermissions = original.StoragePermissionsSet
	StoragePermissionsSetsas        StoragePermissions = original.StoragePermissionsSetsas
	StoragePermissionsUpdate        StoragePermissions = original.StoragePermissionsUpdate
)

type AccessPolicyEntry = original.AccessPolicyEntry
type BaseClient = original.BaseClient
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type DeletedVault = original.DeletedVault
type DeletedVaultListResult = original.DeletedVaultListResult
type DeletedVaultListResultIterator = original.DeletedVaultListResultIterator
type DeletedVaultListResultPage = original.DeletedVaultListResultPage
type DeletedVaultProperties = original.DeletedVaultProperties
type IPRule = original.IPRule
type LogSpecification = original.LogSpecification
type NetworkRuleSet = original.NetworkRuleSet
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type Permissions = original.Permissions
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type Resource = original.Resource
type ResourceListResult = original.ResourceListResult
type ResourceListResultIterator = original.ResourceListResultIterator
type ResourceListResultPage = original.ResourceListResultPage
type ServiceSpecification = original.ServiceSpecification
type Sku = original.Sku
type Vault = original.Vault
type VaultAccessPolicyParameters = original.VaultAccessPolicyParameters
type VaultAccessPolicyProperties = original.VaultAccessPolicyProperties
type VaultCheckNameAvailabilityParameters = original.VaultCheckNameAvailabilityParameters
type VaultCreateOrUpdateParameters = original.VaultCreateOrUpdateParameters
type VaultListResult = original.VaultListResult
type VaultListResultIterator = original.VaultListResultIterator
type VaultListResultPage = original.VaultListResultPage
type VaultPatchParameters = original.VaultPatchParameters
type VaultPatchProperties = original.VaultPatchProperties
type VaultProperties = original.VaultProperties
type VaultsClient = original.VaultsClient
type VaultsCreateOrUpdateFuture = original.VaultsCreateOrUpdateFuture
type VaultsPurgeDeletedFuture = original.VaultsPurgeDeletedFuture
type VirtualNetworkRule = original.VirtualNetworkRule

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewDeletedVaultListResultIterator(page DeletedVaultListResultPage) DeletedVaultListResultIterator {
	return original.NewDeletedVaultListResultIterator(page)
}
func NewDeletedVaultListResultPage(getNextPage func(context.Context, DeletedVaultListResult) (DeletedVaultListResult, error)) DeletedVaultListResultPage {
	return original.NewDeletedVaultListResultPage(getNextPage)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceListResultIterator(page ResourceListResultPage) ResourceListResultIterator {
	return original.NewResourceListResultIterator(page)
}
func NewResourceListResultPage(getNextPage func(context.Context, ResourceListResult) (ResourceListResult, error)) ResourceListResultPage {
	return original.NewResourceListResultPage(getNextPage)
}
func NewVaultListResultIterator(page VaultListResultPage) VaultListResultIterator {
	return original.NewVaultListResultIterator(page)
}
func NewVaultListResultPage(getNextPage func(context.Context, VaultListResult) (VaultListResult, error)) VaultListResultPage {
	return original.NewVaultListResultPage(getNextPage)
}
func NewVaultsClient(subscriptionID string) VaultsClient {
	return original.NewVaultsClient(subscriptionID)
}
func NewVaultsClientWithBaseURI(baseURI string, subscriptionID string) VaultsClient {
	return original.NewVaultsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessPolicyUpdateKindValues() []AccessPolicyUpdateKind {
	return original.PossibleAccessPolicyUpdateKindValues()
}
func PossibleCertificatePermissionsValues() []CertificatePermissions {
	return original.PossibleCertificatePermissionsValues()
}
func PossibleCreateModeValues() []CreateMode {
	return original.PossibleCreateModeValues()
}
func PossibleKeyPermissionsValues() []KeyPermissions {
	return original.PossibleKeyPermissionsValues()
}
func PossibleNetworkRuleActionValues() []NetworkRuleAction {
	return original.PossibleNetworkRuleActionValues()
}
func PossibleNetworkRuleBypassOptionsValues() []NetworkRuleBypassOptions {
	return original.PossibleNetworkRuleBypassOptionsValues()
}
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return original.PossiblePrivateEndpointConnectionProvisioningStateValues()
}
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return original.PossiblePrivateEndpointServiceConnectionStatusValues()
}
func PossibleReasonValues() []Reason {
	return original.PossibleReasonValues()
}
func PossibleSecretPermissionsValues() []SecretPermissions {
	return original.PossibleSecretPermissionsValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleStoragePermissionsValues() []StoragePermissions {
	return original.PossibleStoragePermissionsValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
