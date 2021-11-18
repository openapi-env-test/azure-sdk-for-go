package contosocdnapi

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
	"github.com/Azure/azure-sdk-for-go/services/contoso_cdn/mgmt/2020-09-01/contoso_cdn"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckNameAvailability(ctx context.Context, checkNameAvailabilityInput contosocdn.CheckNameAvailabilityInput) (result contosocdn.CheckNameAvailabilityOutput, err error)
	CheckNameAvailabilityWithSubscription(ctx context.Context, checkNameAvailabilityInput contosocdn.CheckNameAvailabilityInput) (result contosocdn.CheckNameAvailabilityOutput, err error)
	ValidateProbe(ctx context.Context, validateProbeInput contosocdn.ValidateProbeInput) (result contosocdn.ValidateProbeOutput, err error)
}

var _ BaseClientAPI = (*contosocdn.BaseClient)(nil)

// ProfilesClientAPI contains the set of methods on the ProfilesClient type.
type ProfilesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, profile contosocdn.Profile) (result contosocdn.ProfilesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.ProfilesDeleteFuture, err error)
	GenerateSsoURI(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.SsoURI, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.Profile, err error)
	List(ctx context.Context) (result contosocdn.ProfileListResultPage, err error)
	ListComplete(ctx context.Context) (result contosocdn.ProfileListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result contosocdn.ProfileListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result contosocdn.ProfileListResultIterator, err error)
	ListResourceUsage(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.ResourceUsageListResultPage, err error)
	ListResourceUsageComplete(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.ResourceUsageListResultIterator, err error)
	ListSupportedOptimizationTypes(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.SupportedOptimizationTypesListResult, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, profileUpdateParameters contosocdn.ProfileUpdateParameters) (result contosocdn.ProfilesUpdateFuture, err error)
}

var _ ProfilesClientAPI = (*contosocdn.ProfilesClient)(nil)

// EndpointsClientAPI contains the set of methods on the EndpointsClient type.
type EndpointsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpoint contosocdn.Endpoint) (result contosocdn.EndpointsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.EndpointsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.Endpoint, err error)
	ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.EndpointListResultPage, err error)
	ListByProfileComplete(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.EndpointListResultIterator, err error)
	ListResourceUsage(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.ResourceUsageListResultPage, err error)
	ListResourceUsageComplete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.ResourceUsageListResultIterator, err error)
	LoadContent(ctx context.Context, resourceGroupName string, profileName string, endpointName string, contentFilePaths contosocdn.LoadParameters) (result contosocdn.EndpointsLoadContentFuture, err error)
	PurgeContent(ctx context.Context, resourceGroupName string, profileName string, endpointName string, contentFilePaths contosocdn.PurgeParameters) (result contosocdn.EndpointsPurgeContentFuture, err error)
	Start(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.EndpointsStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.EndpointsStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpointUpdateProperties contosocdn.EndpointUpdateParameters) (result contosocdn.EndpointsUpdateFuture, err error)
	ValidateCustomDomain(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainProperties contosocdn.ValidateCustomDomainInput) (result contosocdn.ValidateCustomDomainOutput, err error)
}

var _ EndpointsClientAPI = (*contosocdn.EndpointsClient)(nil)

// OriginsClientAPI contains the set of methods on the OriginsClient type.
type OriginsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, origin contosocdn.Origin) (result contosocdn.OriginsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string) (result contosocdn.OriginsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string) (result contosocdn.Origin, err error)
	ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.OriginListResultPage, err error)
	ListByEndpointComplete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.OriginListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, originUpdateProperties contosocdn.OriginUpdateParameters) (result contosocdn.OriginsUpdateFuture, err error)
}

var _ OriginsClientAPI = (*contosocdn.OriginsClient)(nil)

// OriginGroupsClientAPI contains the set of methods on the OriginGroupsClient type.
type OriginGroupsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, originGroup contosocdn.OriginGroup) (result contosocdn.OriginGroupsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string) (result contosocdn.OriginGroupsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string) (result contosocdn.OriginGroup, err error)
	ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.OriginGroupListResultPage, err error)
	ListByEndpointComplete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.OriginGroupListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, originGroupUpdateProperties contosocdn.OriginGroupUpdateParameters) (result contosocdn.OriginGroupsUpdateFuture, err error)
}

var _ OriginGroupsClientAPI = (*contosocdn.OriginGroupsClient)(nil)

// CustomDomainsClientAPI contains the set of methods on the CustomDomainsClient type.
type CustomDomainsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, customDomainProperties contosocdn.CustomDomainParameters) (result contosocdn.CustomDomainsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (result contosocdn.CustomDomainsDeleteFuture, err error)
	DisableCustomHTTPS(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (result contosocdn.CustomDomain, err error)
	EnableCustomHTTPS(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, customDomainHTTPSParameters *contosocdn.BasicCustomDomainHTTPSParameters) (result contosocdn.CustomDomain, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (result contosocdn.CustomDomain, err error)
	ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.CustomDomainListResultPage, err error)
	ListByEndpointComplete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.CustomDomainListResultIterator, err error)
}

var _ CustomDomainsClientAPI = (*contosocdn.CustomDomainsClient)(nil)

// ResourceUsageClientAPI contains the set of methods on the ResourceUsageClient type.
type ResourceUsageClientAPI interface {
	List(ctx context.Context) (result contosocdn.ResourceUsageListResultPage, err error)
	ListComplete(ctx context.Context) (result contosocdn.ResourceUsageListResultIterator, err error)
}

var _ ResourceUsageClientAPI = (*contosocdn.ResourceUsageClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result contosocdn.OperationsListResultPage, err error)
	ListComplete(ctx context.Context) (result contosocdn.OperationsListResultIterator, err error)
}

var _ OperationsClientAPI = (*contosocdn.OperationsClient)(nil)

// EdgeNodesClientAPI contains the set of methods on the EdgeNodesClient type.
type EdgeNodesClientAPI interface {
	List(ctx context.Context) (result contosocdn.EdgenodeResultPage, err error)
	ListComplete(ctx context.Context) (result contosocdn.EdgenodeResultIterator, err error)
}

var _ EdgeNodesClientAPI = (*contosocdn.EdgeNodesClient)(nil)

// AFDProfilesClientAPI contains the set of methods on the AFDProfilesClient type.
type AFDProfilesClientAPI interface {
	CheckHostNameAvailability(ctx context.Context, resourceGroupName string, profileName string, checkHostNameAvailabilityInput contosocdn.ValidateCustomDomainInput) (result contosocdn.ValidateCustomDomainOutput, err error)
	ListResourceUsage(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.UsagesListResultPage, err error)
	ListResourceUsageComplete(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.UsagesListResultIterator, err error)
}

var _ AFDProfilesClientAPI = (*contosocdn.AFDProfilesClient)(nil)

// AFDCustomDomainsClientAPI contains the set of methods on the AFDCustomDomainsClient type.
type AFDCustomDomainsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomain contosocdn.AFDDomain) (result contosocdn.AFDCustomDomainsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, customDomainName string) (result contosocdn.AFDCustomDomainsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, customDomainName string) (result contosocdn.AFDDomain, err error)
	ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.AFDDomainListResultPage, err error)
	ListByProfileComplete(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.AFDDomainListResultIterator, err error)
	RefreshValidationToken(ctx context.Context, resourceGroupName string, profileName string, customDomainName string) (result contosocdn.AFDCustomDomainsRefreshValidationTokenFuture, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomainUpdateProperties contosocdn.AFDDomainUpdateParameters) (result contosocdn.AFDCustomDomainsUpdateFuture, err error)
}

var _ AFDCustomDomainsClientAPI = (*contosocdn.AFDCustomDomainsClient)(nil)

// AFDEndpointsClientAPI contains the set of methods on the AFDEndpointsClient type.
type AFDEndpointsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpoint contosocdn.AFDEndpoint) (result contosocdn.AFDEndpointsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.AFDEndpointsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.AFDEndpoint, err error)
	ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.AFDEndpointListResultPage, err error)
	ListByProfileComplete(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.AFDEndpointListResultIterator, err error)
	ListResourceUsage(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.UsagesListResultPage, err error)
	ListResourceUsageComplete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.UsagesListResultIterator, err error)
	PurgeContent(ctx context.Context, resourceGroupName string, profileName string, endpointName string, contents contosocdn.AfdPurgeParameters) (result contosocdn.AFDEndpointsPurgeContentFuture, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpointUpdateProperties contosocdn.AFDEndpointUpdateParameters) (result contosocdn.AFDEndpointsUpdateFuture, err error)
	ValidateCustomDomain(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainProperties contosocdn.ValidateCustomDomainInput) (result contosocdn.ValidateCustomDomainOutput, err error)
}

var _ AFDEndpointsClientAPI = (*contosocdn.AFDEndpointsClient)(nil)

// AFDOriginGroupsClientAPI contains the set of methods on the AFDOriginGroupsClient type.
type AFDOriginGroupsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originGroup contosocdn.AFDOriginGroup) (result contosocdn.AFDOriginGroupsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, originGroupName string) (result contosocdn.AFDOriginGroupsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, originGroupName string) (result contosocdn.AFDOriginGroup, err error)
	ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.AFDOriginGroupListResultPage, err error)
	ListByProfileComplete(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.AFDOriginGroupListResultIterator, err error)
	ListResourceUsage(ctx context.Context, resourceGroupName string, profileName string, originGroupName string) (result contosocdn.UsagesListResultPage, err error)
	ListResourceUsageComplete(ctx context.Context, resourceGroupName string, profileName string, originGroupName string) (result contosocdn.UsagesListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originGroupUpdateProperties contosocdn.AFDOriginGroupUpdateParameters) (result contosocdn.AFDOriginGroupsUpdateFuture, err error)
}

var _ AFDOriginGroupsClientAPI = (*contosocdn.AFDOriginGroupsClient)(nil)

// AFDOriginsClientAPI contains the set of methods on the AFDOriginsClient type.
type AFDOriginsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originName string, origin contosocdn.AFDOrigin) (result contosocdn.AFDOriginsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originName string) (result contosocdn.AFDOriginsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originName string) (result contosocdn.AFDOrigin, err error)
	ListByOriginGroup(ctx context.Context, resourceGroupName string, profileName string, originGroupName string) (result contosocdn.AFDOriginListResultPage, err error)
	ListByOriginGroupComplete(ctx context.Context, resourceGroupName string, profileName string, originGroupName string) (result contosocdn.AFDOriginListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originName string, originUpdateProperties contosocdn.AFDOriginUpdateParameters) (result contosocdn.AFDOriginsUpdateFuture, err error)
}

var _ AFDOriginsClientAPI = (*contosocdn.AFDOriginsClient)(nil)

// RoutesClientAPI contains the set of methods on the RoutesClient type.
type RoutesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, route contosocdn.Route) (result contosocdn.RoutesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string) (result contosocdn.RoutesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string) (result contosocdn.Route, err error)
	ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.RouteListResultPage, err error)
	ListByEndpointComplete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result contosocdn.RouteListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, routeUpdateProperties contosocdn.RouteUpdateParameters) (result contosocdn.RoutesUpdateFuture, err error)
}

var _ RoutesClientAPI = (*contosocdn.RoutesClient)(nil)

// RuleSetsClientAPI contains the set of methods on the RuleSetsClient type.
type RuleSetsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, ruleSet contosocdn.RuleSet) (result contosocdn.RuleSetsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string) (result contosocdn.RuleSetsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string) (result contosocdn.RuleSet, err error)
	ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.RuleSetListResultPage, err error)
	ListByProfileComplete(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.RuleSetListResultIterator, err error)
	ListResourceUsage(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string) (result contosocdn.UsagesListResultPage, err error)
	ListResourceUsageComplete(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string) (result contosocdn.UsagesListResultIterator, err error)
}

var _ RuleSetsClientAPI = (*contosocdn.RuleSetsClient)(nil)

// RulesClientAPI contains the set of methods on the RulesClient type.
type RulesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, ruleName string, rule contosocdn.Rule) (result contosocdn.RulesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, ruleName string) (result contosocdn.RulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, ruleName string) (result contosocdn.Rule, err error)
	ListByRuleSet(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string) (result contosocdn.RuleListResultPage, err error)
	ListByRuleSetComplete(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string) (result contosocdn.RuleListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, ruleName string, ruleUpdateProperties contosocdn.RuleUpdateParameters) (result contosocdn.RulesUpdateFuture, err error)
}

var _ RulesClientAPI = (*contosocdn.RulesClient)(nil)

// SecurityPoliciesClientAPI contains the set of methods on the SecurityPoliciesClient type.
type SecurityPoliciesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, securityPolicyName string, securityPolicy contosocdn.SecurityPolicy) (result contosocdn.SecurityPoliciesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, securityPolicyName string) (result contosocdn.SecurityPoliciesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, securityPolicyName string) (result contosocdn.SecurityPolicy, err error)
	ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.SecurityPolicyListResultPage, err error)
	ListByProfileComplete(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.SecurityPolicyListResultIterator, err error)
	Patch(ctx context.Context, resourceGroupName string, profileName string, securityPolicyName string, securityPolicyParameters contosocdn.SecurityPolicyWebApplicationFirewallParameters) (result contosocdn.SecurityPoliciesPatchFuture, err error)
}

var _ SecurityPoliciesClientAPI = (*contosocdn.SecurityPoliciesClient)(nil)

// SecretsClientAPI contains the set of methods on the SecretsClient type.
type SecretsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, secretName string, secret contosocdn.Secret) (result contosocdn.SecretsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, secretName string) (result contosocdn.SecretsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, secretName string) (result contosocdn.Secret, err error)
	ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.SecretListResultPage, err error)
	ListByProfileComplete(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.SecretListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, secretName string, secretProperties contosocdn.SecretProperties) (result contosocdn.SecretsUpdateFuture, err error)
}

var _ SecretsClientAPI = (*contosocdn.SecretsClient)(nil)

// ValidateClientAPI contains the set of methods on the ValidateClient type.
type ValidateClientAPI interface {
	SecretMethod(ctx context.Context, validateSecretInput contosocdn.ValidateSecretInput) (result contosocdn.ValidateSecretOutput, err error)
}

var _ ValidateClientAPI = (*contosocdn.ValidateClient)(nil)

// LogAnalyticsClientAPI contains the set of methods on the LogAnalyticsClient type.
type LogAnalyticsClientAPI interface {
	GetLogAnalyticsLocations(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.ContinentsResponse, err error)
	GetLogAnalyticsMetrics(ctx context.Context, resourceGroupName string, profileName string, metrics []string, dateTimeBegin date.Time, dateTimeEnd date.Time, granularity string, groupBy []string, continents []string, countryOrRegions []string, customDomains []string, protocols []string) (result contosocdn.MetricsResponse, err error)
	GetLogAnalyticsRankings(ctx context.Context, resourceGroupName string, profileName string, rankings []string, metrics []string, maxRanking float64, dateTimeBegin date.Time, dateTimeEnd date.Time, customDomains []string) (result contosocdn.RankingsResponse, err error)
	GetLogAnalyticsResources(ctx context.Context, resourceGroupName string, profileName string) (result contosocdn.ResourcesResponse, err error)
	GetWafLogAnalyticsMetrics(ctx context.Context, resourceGroupName string, profileName string, metrics []string, dateTimeBegin date.Time, dateTimeEnd date.Time, granularity string, actions []string, groupBy []string, ruleTypes []string) (result contosocdn.WafMetricsResponse, err error)
	GetWafLogAnalyticsRankings(ctx context.Context, resourceGroupName string, profileName string, metrics []string, dateTimeBegin date.Time, dateTimeEnd date.Time, maxRanking float64, rankings []string, actions []string, ruleTypes []string) (result contosocdn.WafRankingsResponse, err error)
}

var _ LogAnalyticsClientAPI = (*contosocdn.LogAnalyticsClient)(nil)

// PoliciesClientAPI contains the set of methods on the PoliciesClient type.
type PoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, policyName string, cdnWebApplicationFirewallPolicy contosocdn.CdnWebApplicationFirewallPolicy) (result contosocdn.PoliciesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, policyName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, policyName string) (result contosocdn.CdnWebApplicationFirewallPolicy, err error)
	List(ctx context.Context, resourceGroupName string) (result contosocdn.CdnWebApplicationFirewallPolicyListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result contosocdn.CdnWebApplicationFirewallPolicyListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, policyName string, cdnWebApplicationFirewallPolicyPatchParameters contosocdn.CdnWebApplicationFirewallPolicyPatchParameters) (result contosocdn.PoliciesUpdateFuture, err error)
}

var _ PoliciesClientAPI = (*contosocdn.PoliciesClient)(nil)

// ManagedRuleSetsClientAPI contains the set of methods on the ManagedRuleSetsClient type.
type ManagedRuleSetsClientAPI interface {
	List(ctx context.Context) (result contosocdn.ManagedRuleSetDefinitionListPage, err error)
	ListComplete(ctx context.Context) (result contosocdn.ManagedRuleSetDefinitionListIterator, err error)
}

var _ ManagedRuleSetsClientAPI = (*contosocdn.ManagedRuleSetsClient)(nil)
