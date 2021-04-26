package securityapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v1.0/security"
	"github.com/Azure/go-autorest/autorest"
)

// SQLVulnerabilityAssessmentScansClientAPI contains the set of methods on the SQLVulnerabilityAssessmentScansClient type.
type SQLVulnerabilityAssessmentScansClientAPI interface {
	Get(ctx context.Context, scanID string, workspaceID string, APIVersion string, resourceID string) (result security.Scan, err error)
	List(ctx context.Context, workspaceID string, APIVersion string, resourceID string) (result security.Scans, err error)
}

var _ SQLVulnerabilityAssessmentScansClientAPI = (*security.SQLVulnerabilityAssessmentScansClient)(nil)

// SQLVulnerabilityAssessmentScanResultsClientAPI contains the set of methods on the SQLVulnerabilityAssessmentScanResultsClient type.
type SQLVulnerabilityAssessmentScanResultsClientAPI interface {
	Get(ctx context.Context, scanID string, scanResultID string, workspaceID string, APIVersion string, resourceID string) (result security.ScanResult, err error)
	List(ctx context.Context, scanID string, workspaceID string, APIVersion string, resourceID string) (result security.ScanResults, err error)
}

var _ SQLVulnerabilityAssessmentScanResultsClientAPI = (*security.SQLVulnerabilityAssessmentScanResultsClient)(nil)

// SQLVulnerabilityAssessmentBaselineRulesClientAPI contains the set of methods on the SQLVulnerabilityAssessmentBaselineRulesClient type.
type SQLVulnerabilityAssessmentBaselineRulesClientAPI interface {
	Add(ctx context.Context, workspaceID string, APIVersion string, resourceID string, body *security.RulesResultsInput) (result security.RulesResults, err error)
	CreateOrUpdate(ctx context.Context, ruleID string, workspaceID string, APIVersion string, resourceID string, body *security.RuleResultsInput) (result security.RuleResults, err error)
	Delete(ctx context.Context, ruleID string, workspaceID string, APIVersion string, resourceID string) (result autorest.Response, err error)
	Get(ctx context.Context, ruleID string, workspaceID string, APIVersion string, resourceID string) (result security.RuleResults, err error)
	List(ctx context.Context, workspaceID string, APIVersion string, resourceID string) (result security.RulesResults, err error)
}

var _ SQLVulnerabilityAssessmentBaselineRulesClientAPI = (*security.SQLVulnerabilityAssessmentBaselineRulesClient)(nil)

// SecureScoresClientAPI contains the set of methods on the SecureScoresClient type.
type SecureScoresClientAPI interface {
	Get(ctx context.Context, secureScoreName string) (result security.SecureScoreItem, err error)
	List(ctx context.Context) (result security.SecureScoresListPage, err error)
	ListComplete(ctx context.Context) (result security.SecureScoresListIterator, err error)
}

var _ SecureScoresClientAPI = (*security.SecureScoresClient)(nil)

// SecureScoreControlsClientAPI contains the set of methods on the SecureScoreControlsClient type.
type SecureScoreControlsClientAPI interface {
	List(ctx context.Context, expand security.ExpandControlsEnum) (result security.SecureScoreControlListPage, err error)
	ListComplete(ctx context.Context, expand security.ExpandControlsEnum) (result security.SecureScoreControlListIterator, err error)
	ListBySecureScore(ctx context.Context, secureScoreName string, expand security.ExpandControlsEnum) (result security.SecureScoreControlListPage, err error)
	ListBySecureScoreComplete(ctx context.Context, secureScoreName string, expand security.ExpandControlsEnum) (result security.SecureScoreControlListIterator, err error)
}

var _ SecureScoreControlsClientAPI = (*security.SecureScoreControlsClient)(nil)

// SecureScoreControlDefinitionsClientAPI contains the set of methods on the SecureScoreControlDefinitionsClient type.
type SecureScoreControlDefinitionsClientAPI interface {
	List(ctx context.Context) (result security.SecureScoreControlDefinitionListPage, err error)
	ListComplete(ctx context.Context) (result security.SecureScoreControlDefinitionListIterator, err error)
	ListBySubscription(ctx context.Context) (result security.SecureScoreControlDefinitionListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result security.SecureScoreControlDefinitionListIterator, err error)
}

var _ SecureScoreControlDefinitionsClientAPI = (*security.SecureScoreControlDefinitionsClient)(nil)

// ConnectorsClientAPI contains the set of methods on the ConnectorsClient type.
type ConnectorsClientAPI interface {
	CreateOrUpdate(ctx context.Context, connectorName string, connectorSetting security.ConnectorSetting) (result security.ConnectorSetting, err error)
	Delete(ctx context.Context, connectorName string) (result autorest.Response, err error)
	Get(ctx context.Context, connectorName string) (result security.ConnectorSetting, err error)
	List(ctx context.Context) (result security.ConnectorSettingListPage, err error)
	ListComplete(ctx context.Context) (result security.ConnectorSettingListIterator, err error)
}

var _ ConnectorsClientAPI = (*security.ConnectorsClient)(nil)

// AutomationsClientAPI contains the set of methods on the AutomationsClient type.
type AutomationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, automationName string, automation security.Automation) (result security.Automation, err error)
	Delete(ctx context.Context, resourceGroupName string, automationName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, automationName string) (result security.Automation, err error)
	List(ctx context.Context) (result security.AutomationListPage, err error)
	ListComplete(ctx context.Context) (result security.AutomationListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result security.AutomationListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result security.AutomationListIterator, err error)
	Validate(ctx context.Context, resourceGroupName string, automationName string, automation security.Automation) (result security.AutomationValidationStatus, err error)
}

var _ AutomationsClientAPI = (*security.AutomationsClient)(nil)

// SubAssessmentsClientAPI contains the set of methods on the SubAssessmentsClient type.
type SubAssessmentsClientAPI interface {
	Get(ctx context.Context, scope string, assessmentName string, subAssessmentName string) (result security.SubAssessment, err error)
	List(ctx context.Context, scope string, assessmentName string) (result security.SubAssessmentListPage, err error)
	ListComplete(ctx context.Context, scope string, assessmentName string) (result security.SubAssessmentListIterator, err error)
	ListAll(ctx context.Context, scope string) (result security.SubAssessmentListPage, err error)
	ListAllComplete(ctx context.Context, scope string) (result security.SubAssessmentListIterator, err error)
}

var _ SubAssessmentsClientAPI = (*security.SubAssessmentsClient)(nil)

// RegulatoryComplianceStandardsClientAPI contains the set of methods on the RegulatoryComplianceStandardsClient type.
type RegulatoryComplianceStandardsClientAPI interface {
	Get(ctx context.Context, regulatoryComplianceStandardName string) (result security.RegulatoryComplianceStandard, err error)
	List(ctx context.Context, filter string) (result security.RegulatoryComplianceStandardListPage, err error)
	ListComplete(ctx context.Context, filter string) (result security.RegulatoryComplianceStandardListIterator, err error)
}

var _ RegulatoryComplianceStandardsClientAPI = (*security.RegulatoryComplianceStandardsClient)(nil)

// RegulatoryComplianceControlsClientAPI contains the set of methods on the RegulatoryComplianceControlsClient type.
type RegulatoryComplianceControlsClientAPI interface {
	Get(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string) (result security.RegulatoryComplianceControl, err error)
	List(ctx context.Context, regulatoryComplianceStandardName string, filter string) (result security.RegulatoryComplianceControlListPage, err error)
	ListComplete(ctx context.Context, regulatoryComplianceStandardName string, filter string) (result security.RegulatoryComplianceControlListIterator, err error)
}

var _ RegulatoryComplianceControlsClientAPI = (*security.RegulatoryComplianceControlsClient)(nil)

// RegulatoryComplianceAssessmentsClientAPI contains the set of methods on the RegulatoryComplianceAssessmentsClient type.
type RegulatoryComplianceAssessmentsClientAPI interface {
	Get(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, regulatoryComplianceAssessmentName string) (result security.RegulatoryComplianceAssessment, err error)
	List(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, filter string) (result security.RegulatoryComplianceAssessmentListPage, err error)
	ListComplete(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, filter string) (result security.RegulatoryComplianceAssessmentListIterator, err error)
}

var _ RegulatoryComplianceAssessmentsClientAPI = (*security.RegulatoryComplianceAssessmentsClient)(nil)

// PricingsClientAPI contains the set of methods on the PricingsClient type.
type PricingsClientAPI interface {
	CreateOrUpdateResourceGroupPricing(ctx context.Context, resourceGroupName string, pricingName string, pricing security.Pricing) (result security.Pricing, err error)
	GetResourceGroupPricing(ctx context.Context, resourceGroupName string, pricingName string) (result security.Pricing, err error)
	GetSubscriptionPricing(ctx context.Context, pricingName string) (result security.Pricing, err error)
	List(ctx context.Context) (result security.PricingListPage, err error)
	ListComplete(ctx context.Context) (result security.PricingListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result security.PricingListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result security.PricingListIterator, err error)
	UpdateSubscriptionPricing(ctx context.Context, pricingName string, pricing security.Pricing) (result security.Pricing, err error)
}

var _ PricingsClientAPI = (*security.PricingsClient)(nil)

// ContactsClientAPI contains the set of methods on the ContactsClient type.
type ContactsClientAPI interface {
	Create(ctx context.Context, securityContactName string, securityContact security.Contact) (result security.Contact, err error)
	Delete(ctx context.Context, securityContactName string) (result autorest.Response, err error)
	Get(ctx context.Context, securityContactName string) (result security.Contact, err error)
	List(ctx context.Context) (result security.ContactListPage, err error)
	ListComplete(ctx context.Context) (result security.ContactListIterator, err error)
	Update(ctx context.Context, securityContactName string, securityContact security.Contact) (result security.Contact, err error)
}

var _ ContactsClientAPI = (*security.ContactsClient)(nil)

// WorkspaceSettingsClientAPI contains the set of methods on the WorkspaceSettingsClient type.
type WorkspaceSettingsClientAPI interface {
	Create(ctx context.Context, workspaceSettingName string, workspaceSetting security.WorkspaceSetting) (result security.WorkspaceSetting, err error)
	Delete(ctx context.Context, workspaceSettingName string) (result autorest.Response, err error)
	Get(ctx context.Context, workspaceSettingName string) (result security.WorkspaceSetting, err error)
	List(ctx context.Context) (result security.WorkspaceSettingListPage, err error)
	ListComplete(ctx context.Context) (result security.WorkspaceSettingListIterator, err error)
	Update(ctx context.Context, workspaceSettingName string, workspaceSetting security.WorkspaceSetting) (result security.WorkspaceSetting, err error)
}

var _ WorkspaceSettingsClientAPI = (*security.WorkspaceSettingsClient)(nil)

// AutoProvisioningSettingsClientAPI contains the set of methods on the AutoProvisioningSettingsClient type.
type AutoProvisioningSettingsClientAPI interface {
	Create(ctx context.Context, settingName string, setting security.AutoProvisioningSetting) (result security.AutoProvisioningSetting, err error)
	Get(ctx context.Context, settingName string) (result security.AutoProvisioningSetting, err error)
	List(ctx context.Context) (result security.AutoProvisioningSettingListPage, err error)
	ListComplete(ctx context.Context) (result security.AutoProvisioningSettingListIterator, err error)
}

var _ AutoProvisioningSettingsClientAPI = (*security.AutoProvisioningSettingsClient)(nil)

// CompliancesClientAPI contains the set of methods on the CompliancesClient type.
type CompliancesClientAPI interface {
	Get(ctx context.Context, scope string, complianceName string) (result security.Compliance, err error)
	List(ctx context.Context, scope string) (result security.ComplianceListPage, err error)
	ListComplete(ctx context.Context, scope string) (result security.ComplianceListIterator, err error)
}

var _ CompliancesClientAPI = (*security.CompliancesClient)(nil)

// AdvancedThreatProtectionClientAPI contains the set of methods on the AdvancedThreatProtectionClient type.
type AdvancedThreatProtectionClientAPI interface {
	Create(ctx context.Context, resourceID string, advancedThreatProtectionSetting security.AdvancedThreatProtectionSetting) (result security.AdvancedThreatProtectionSetting, err error)
	Get(ctx context.Context, resourceID string) (result security.AdvancedThreatProtectionSetting, err error)
}

var _ AdvancedThreatProtectionClientAPI = (*security.AdvancedThreatProtectionClient)(nil)

// DeviceSecurityGroupsClientAPI contains the set of methods on the DeviceSecurityGroupsClient type.
type DeviceSecurityGroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceID string, deviceSecurityGroupName string, deviceSecurityGroup security.DeviceSecurityGroup) (result security.DeviceSecurityGroup, err error)
	Delete(ctx context.Context, resourceID string, deviceSecurityGroupName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceID string, deviceSecurityGroupName string) (result security.DeviceSecurityGroup, err error)
	List(ctx context.Context, resourceID string) (result security.DeviceSecurityGroupListPage, err error)
	ListComplete(ctx context.Context, resourceID string) (result security.DeviceSecurityGroupListIterator, err error)
}

var _ DeviceSecurityGroupsClientAPI = (*security.DeviceSecurityGroupsClient)(nil)

// SettingsClientAPI contains the set of methods on the SettingsClient type.
type SettingsClientAPI interface {
	Get(ctx context.Context, settingName string) (result security.SettingModel, err error)
	List(ctx context.Context) (result security.SettingsListPage, err error)
	ListComplete(ctx context.Context) (result security.SettingsListIterator, err error)
	Update(ctx context.Context, settingName string, setting security.BasicSetting) (result security.SettingModel, err error)
}

var _ SettingsClientAPI = (*security.SettingsClient)(nil)

// InformationProtectionPoliciesClientAPI contains the set of methods on the InformationProtectionPoliciesClient type.
type InformationProtectionPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, scope string, informationProtectionPolicyName string, informationProtectionPolicy security.InformationProtectionPolicy) (result security.InformationProtectionPolicy, err error)
	Get(ctx context.Context, scope string, informationProtectionPolicyName string) (result security.InformationProtectionPolicy, err error)
	List(ctx context.Context, scope string) (result security.InformationProtectionPolicyListPage, err error)
	ListComplete(ctx context.Context, scope string) (result security.InformationProtectionPolicyListIterator, err error)
}

var _ InformationProtectionPoliciesClientAPI = (*security.InformationProtectionPoliciesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result security.OperationListPage, err error)
	ListComplete(ctx context.Context) (result security.OperationListIterator, err error)
}

var _ OperationsClientAPI = (*security.OperationsClient)(nil)

// LocationsClientAPI contains the set of methods on the LocationsClient type.
type LocationsClientAPI interface {
	Get(ctx context.Context) (result security.AscLocation, err error)
	List(ctx context.Context) (result security.AscLocationListPage, err error)
	ListComplete(ctx context.Context) (result security.AscLocationListIterator, err error)
}

var _ LocationsClientAPI = (*security.LocationsClient)(nil)

// TasksClientAPI contains the set of methods on the TasksClient type.
type TasksClientAPI interface {
	GetResourceGroupLevelTask(ctx context.Context, resourceGroupName string, taskName string) (result security.Task, err error)
	GetSubscriptionLevelTask(ctx context.Context, taskName string) (result security.Task, err error)
	List(ctx context.Context, filter string) (result security.TaskListPage, err error)
	ListComplete(ctx context.Context, filter string) (result security.TaskListIterator, err error)
	ListByHomeRegion(ctx context.Context, filter string) (result security.TaskListPage, err error)
	ListByHomeRegionComplete(ctx context.Context, filter string) (result security.TaskListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string) (result security.TaskListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string) (result security.TaskListIterator, err error)
	UpdateResourceGroupLevelTaskState(ctx context.Context, resourceGroupName string, taskName string, taskUpdateActionType string) (result autorest.Response, err error)
	UpdateSubscriptionLevelTaskState(ctx context.Context, taskName string, taskUpdateActionType string) (result autorest.Response, err error)
}

var _ TasksClientAPI = (*security.TasksClient)(nil)

// AlertsClientAPI contains the set of methods on the AlertsClient type.
type AlertsClientAPI interface {
	GetResourceGroupLevelAlerts(ctx context.Context, alertName string, resourceGroupName string) (result security.Alert, err error)
	GetSubscriptionLevelAlert(ctx context.Context, alertName string) (result security.Alert, err error)
	List(ctx context.Context, filter string, selectParameter string, expand string) (result security.AlertListPage, err error)
	ListComplete(ctx context.Context, filter string, selectParameter string, expand string) (result security.AlertListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, selectParameter string, expand string) (result security.AlertListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, selectParameter string, expand string) (result security.AlertListIterator, err error)
	ListResourceGroupLevelAlertsByRegion(ctx context.Context, resourceGroupName string, filter string, selectParameter string, expand string) (result security.AlertListPage, err error)
	ListResourceGroupLevelAlertsByRegionComplete(ctx context.Context, resourceGroupName string, filter string, selectParameter string, expand string) (result security.AlertListIterator, err error)
	ListSubscriptionLevelAlertsByRegion(ctx context.Context, filter string, selectParameter string, expand string) (result security.AlertListPage, err error)
	ListSubscriptionLevelAlertsByRegionComplete(ctx context.Context, filter string, selectParameter string, expand string) (result security.AlertListIterator, err error)
	UpdateResourceGroupLevelAlertStateToDismiss(ctx context.Context, alertName string, resourceGroupName string) (result autorest.Response, err error)
	UpdateResourceGroupLevelAlertStateToReactivate(ctx context.Context, alertName string, resourceGroupName string) (result autorest.Response, err error)
	UpdateSubscriptionLevelAlertStateToDismiss(ctx context.Context, alertName string) (result autorest.Response, err error)
	UpdateSubscriptionLevelAlertStateToReactivate(ctx context.Context, alertName string) (result autorest.Response, err error)
}

var _ AlertsClientAPI = (*security.AlertsClient)(nil)

// DiscoveredSecuritySolutionsClientAPI contains the set of methods on the DiscoveredSecuritySolutionsClient type.
type DiscoveredSecuritySolutionsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, discoveredSecuritySolutionName string) (result security.DiscoveredSecuritySolution, err error)
	List(ctx context.Context) (result security.DiscoveredSecuritySolutionListPage, err error)
	ListComplete(ctx context.Context) (result security.DiscoveredSecuritySolutionListIterator, err error)
	ListByHomeRegion(ctx context.Context) (result security.DiscoveredSecuritySolutionListPage, err error)
	ListByHomeRegionComplete(ctx context.Context) (result security.DiscoveredSecuritySolutionListIterator, err error)
}

var _ DiscoveredSecuritySolutionsClientAPI = (*security.DiscoveredSecuritySolutionsClient)(nil)

// JitNetworkAccessPoliciesClientAPI contains the set of methods on the JitNetworkAccessPoliciesClient type.
type JitNetworkAccessPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, jitNetworkAccessPolicyName string, body security.JitNetworkAccessPolicy) (result security.JitNetworkAccessPolicy, err error)
	Delete(ctx context.Context, resourceGroupName string, jitNetworkAccessPolicyName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, jitNetworkAccessPolicyName string) (result security.JitNetworkAccessPolicy, err error)
	Initiate(ctx context.Context, resourceGroupName string, jitNetworkAccessPolicyName string, body security.JitNetworkAccessPolicyInitiateRequest) (result security.JitNetworkAccessRequest, err error)
	List(ctx context.Context) (result security.JitNetworkAccessPoliciesListPage, err error)
	ListComplete(ctx context.Context) (result security.JitNetworkAccessPoliciesListIterator, err error)
	ListByRegion(ctx context.Context) (result security.JitNetworkAccessPoliciesListPage, err error)
	ListByRegionComplete(ctx context.Context) (result security.JitNetworkAccessPoliciesListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result security.JitNetworkAccessPoliciesListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result security.JitNetworkAccessPoliciesListIterator, err error)
	ListByResourceGroupAndRegion(ctx context.Context, resourceGroupName string) (result security.JitNetworkAccessPoliciesListPage, err error)
	ListByResourceGroupAndRegionComplete(ctx context.Context, resourceGroupName string) (result security.JitNetworkAccessPoliciesListIterator, err error)
}

var _ JitNetworkAccessPoliciesClientAPI = (*security.JitNetworkAccessPoliciesClient)(nil)

// AdaptiveApplicationControlsClientAPI contains the set of methods on the AdaptiveApplicationControlsClient type.
type AdaptiveApplicationControlsClientAPI interface {
	Delete(ctx context.Context, groupName string) (result autorest.Response, err error)
	Get(ctx context.Context, groupName string) (result security.AppWhitelistingGroup, err error)
	List(ctx context.Context, includePathRecommendations *bool, summary *bool) (result security.AppWhitelistingGroups, err error)
	Put(ctx context.Context, groupName string, body security.AppWhitelistingPutGroupData) (result security.AppWhitelistingGroup, err error)
}

var _ AdaptiveApplicationControlsClientAPI = (*security.AdaptiveApplicationControlsClient)(nil)

// ExternalSecuritySolutionsClientAPI contains the set of methods on the ExternalSecuritySolutionsClient type.
type ExternalSecuritySolutionsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, externalSecuritySolutionsName string) (result security.ExternalSecuritySolutionModel, err error)
	List(ctx context.Context) (result security.ExternalSecuritySolutionListPage, err error)
	ListComplete(ctx context.Context) (result security.ExternalSecuritySolutionListIterator, err error)
	ListByHomeRegion(ctx context.Context) (result security.ExternalSecuritySolutionListPage, err error)
	ListByHomeRegionComplete(ctx context.Context) (result security.ExternalSecuritySolutionListIterator, err error)
}

var _ ExternalSecuritySolutionsClientAPI = (*security.ExternalSecuritySolutionsClient)(nil)

// TopologyClientAPI contains the set of methods on the TopologyClient type.
type TopologyClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, topologyResourceName string) (result security.TopologyResource, err error)
	List(ctx context.Context) (result security.TopologyListPage, err error)
	ListComplete(ctx context.Context) (result security.TopologyListIterator, err error)
	ListByHomeRegion(ctx context.Context) (result security.TopologyListPage, err error)
	ListByHomeRegionComplete(ctx context.Context) (result security.TopologyListIterator, err error)
}

var _ TopologyClientAPI = (*security.TopologyClient)(nil)

// AllowedConnectionsClientAPI contains the set of methods on the AllowedConnectionsClient type.
type AllowedConnectionsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, connectionType security.ConnectionType) (result security.AllowedConnectionsResource, err error)
	List(ctx context.Context) (result security.AllowedConnectionsListPage, err error)
	ListComplete(ctx context.Context) (result security.AllowedConnectionsListIterator, err error)
	ListByHomeRegion(ctx context.Context) (result security.AllowedConnectionsListPage, err error)
	ListByHomeRegionComplete(ctx context.Context) (result security.AllowedConnectionsListIterator, err error)
}

var _ AllowedConnectionsClientAPI = (*security.AllowedConnectionsClient)(nil)

// AdaptiveNetworkHardeningsClientAPI contains the set of methods on the AdaptiveNetworkHardeningsClient type.
type AdaptiveNetworkHardeningsClientAPI interface {
	Enforce(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, adaptiveNetworkHardeningResourceName string, body security.AdaptiveNetworkHardeningEnforceRequest) (result security.AdaptiveNetworkHardeningsEnforceFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, adaptiveNetworkHardeningResourceName string) (result security.AdaptiveNetworkHardening, err error)
	ListByExtendedResource(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string) (result security.AdaptiveNetworkHardeningsListPage, err error)
	ListByExtendedResourceComplete(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string) (result security.AdaptiveNetworkHardeningsListIterator, err error)
}

var _ AdaptiveNetworkHardeningsClientAPI = (*security.AdaptiveNetworkHardeningsClient)(nil)

// AlertsSuppressionRulesClientAPI contains the set of methods on the AlertsSuppressionRulesClient type.
type AlertsSuppressionRulesClientAPI interface {
	Delete(ctx context.Context, alertsSuppressionRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, alertsSuppressionRuleName string) (result security.AlertsSuppressionRule, err error)
	List(ctx context.Context, alertType string) (result security.AlertsSuppressionRulesListPage, err error)
	ListComplete(ctx context.Context, alertType string) (result security.AlertsSuppressionRulesListIterator, err error)
	Update(ctx context.Context, alertsSuppressionRuleName string, alertsSuppressionRule security.AlertsSuppressionRule) (result security.AlertsSuppressionRule, err error)
}

var _ AlertsSuppressionRulesClientAPI = (*security.AlertsSuppressionRulesClient)(nil)
