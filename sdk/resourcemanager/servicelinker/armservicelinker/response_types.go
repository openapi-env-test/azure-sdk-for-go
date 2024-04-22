//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armservicelinker

// ConfigurationNamesClientListResponse contains the response from method ConfigurationNamesClient.List.
type ConfigurationNamesClientListResponse struct {
	ConfigurationNameResult
}

// ConnectorClientCreateDryrunResponse contains the response from method ConnectorClient.CreateDryrun.
type ConnectorClientCreateDryrunResponse struct {
	DryrunResource
}

// ConnectorClientCreateOrUpdateResponse contains the response from method ConnectorClient.CreateOrUpdate.
type ConnectorClientCreateOrUpdateResponse struct {
	LinkerResource
}

// ConnectorClientDeleteDryrunResponse contains the response from method ConnectorClient.DeleteDryrun.
type ConnectorClientDeleteDryrunResponse struct {
	// placeholder for future response values
}

// ConnectorClientDeleteResponse contains the response from method ConnectorClient.Delete.
type ConnectorClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectorClientGenerateConfigurationsResponse contains the response from method ConnectorClient.GenerateConfigurations.
type ConnectorClientGenerateConfigurationsResponse struct {
	ConfigurationResult
}

// ConnectorClientGetDryrunResponse contains the response from method ConnectorClient.GetDryrun.
type ConnectorClientGetDryrunResponse struct {
	DryrunResource
}

// ConnectorClientGetResponse contains the response from method ConnectorClient.Get.
type ConnectorClientGetResponse struct {
	LinkerResource
}

// ConnectorClientListDryrunResponse contains the response from method ConnectorClient.ListDryrun.
type ConnectorClientListDryrunResponse struct {
	DryrunList
}

// ConnectorClientListResponse contains the response from method ConnectorClient.List.
type ConnectorClientListResponse struct {
	ResourceList
}

// ConnectorClientUpdateDryrunResponse contains the response from method ConnectorClient.UpdateDryrun.
type ConnectorClientUpdateDryrunResponse struct {
	DryrunResource
}

// ConnectorClientUpdateResponse contains the response from method ConnectorClient.Update.
type ConnectorClientUpdateResponse struct {
	LinkerResource
}

// ConnectorClientValidateResponse contains the response from method ConnectorClient.Validate.
type ConnectorClientValidateResponse struct {
	ValidateOperationResult
}

// LinkerClientCreateOrUpdateResponse contains the response from method LinkerClient.CreateOrUpdate.
type LinkerClientCreateOrUpdateResponse struct {
	LinkerResource
}

// LinkerClientDeleteResponse contains the response from method LinkerClient.Delete.
type LinkerClientDeleteResponse struct {
	// placeholder for future response values
}

// LinkerClientGetResponse contains the response from method LinkerClient.Get.
type LinkerClientGetResponse struct {
	LinkerResource
}

// LinkerClientListConfigurationsResponse contains the response from method LinkerClient.ListConfigurations.
type LinkerClientListConfigurationsResponse struct {
	ConfigurationResult
}

// LinkerClientListResponse contains the response from method LinkerClient.List.
type LinkerClientListResponse struct {
	ResourceList
}

// LinkerClientUpdateResponse contains the response from method LinkerClient.Update.
type LinkerClientUpdateResponse struct {
	LinkerResource
}

// LinkerClientValidateResponse contains the response from method LinkerClient.Validate.
type LinkerClientValidateResponse struct {
	ValidateOperationResult
}

// LinkersClientCreateDryrunResponse contains the response from method LinkersClient.CreateDryrun.
type LinkersClientCreateDryrunResponse struct {
	DryrunResource
}

// LinkersClientDeleteDryrunResponse contains the response from method LinkersClient.DeleteDryrun.
type LinkersClientDeleteDryrunResponse struct {
	// placeholder for future response values
}

// LinkersClientGenerateConfigurationsResponse contains the response from method LinkersClient.GenerateConfigurations.
type LinkersClientGenerateConfigurationsResponse struct {
	ConfigurationResult
}

// LinkersClientGetDryrunResponse contains the response from method LinkersClient.GetDryrun.
type LinkersClientGetDryrunResponse struct {
	DryrunResource
}

// LinkersClientListDaprConfigurationsResponse contains the response from method LinkersClient.ListDaprConfigurations.
type LinkersClientListDaprConfigurationsResponse struct {
	DaprConfigurationList
}

// LinkersClientListDryrunResponse contains the response from method LinkersClient.ListDryrun.
type LinkersClientListDryrunResponse struct {
	DryrunList
}

// LinkersClientUpdateDryrunResponse contains the response from method LinkersClient.UpdateDryrun.
type LinkersClientUpdateDryrunResponse struct {
	DryrunResource
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}
