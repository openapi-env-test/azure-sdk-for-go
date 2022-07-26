//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// IncidentsClient contains the methods for the Incidents group.
// Don't use this type directly, use NewIncidentsClient() instead.
type IncidentsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIncidentsClient creates a new instance of IncidentsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIncidentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *IncidentsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &IncidentsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Creates or updates the incident.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// incidentID - Incident ID
// incident - The incident
// options - IncidentsClientCreateOrUpdateOptions contains the optional parameters for the IncidentsClient.CreateOrUpdate
// method.
func (client *IncidentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incident Incident, options *IncidentsClientCreateOrUpdateOptions) (IncidentsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, incident, options)
	if err != nil {
		return IncidentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IncidentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return IncidentsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IncidentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incident Incident, options *IncidentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, incident)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IncidentsClient) createOrUpdateHandleResponse(resp *http.Response) (IncidentsClientCreateOrUpdateResponse, error) {
	result := IncidentsClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Incident); err != nil {
		return IncidentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// CreateTeam - Creates a Microsoft team to investigate the incident by sharing information and insights between participants.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// incidentID - Incident ID
// teamProperties - Team properties
// options - IncidentsClientCreateTeamOptions contains the optional parameters for the IncidentsClient.CreateTeam method.
func (client *IncidentsClient) CreateTeam(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, teamProperties TeamProperties, options *IncidentsClientCreateTeamOptions) (IncidentsClientCreateTeamResponse, error) {
	req, err := client.createTeamCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, teamProperties, options)
	if err != nil {
		return IncidentsClientCreateTeamResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IncidentsClientCreateTeamResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IncidentsClientCreateTeamResponse{}, runtime.NewResponseError(resp)
	}
	return client.createTeamHandleResponse(resp)
}

// createTeamCreateRequest creates the CreateTeam request.
func (client *IncidentsClient) createTeamCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, teamProperties TeamProperties, options *IncidentsClientCreateTeamOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/createTeam"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, teamProperties)
}

// createTeamHandleResponse handles the CreateTeam response.
func (client *IncidentsClient) createTeamHandleResponse(resp *http.Response) (IncidentsClientCreateTeamResponse, error) {
	result := IncidentsClientCreateTeamResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TeamInformation); err != nil {
		return IncidentsClientCreateTeamResponse{}, err
	}
	return result, nil
}

// Delete - Delete the incident.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// incidentID - Incident ID
// options - IncidentsClientDeleteOptions contains the optional parameters for the IncidentsClient.Delete method.
func (client *IncidentsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientDeleteOptions) (IncidentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
	if err != nil {
		return IncidentsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IncidentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IncidentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return IncidentsClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IncidentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets an incident.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// incidentID - Incident ID
// options - IncidentsClientGetOptions contains the optional parameters for the IncidentsClient.Get method.
func (client *IncidentsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientGetOptions) (IncidentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
	if err != nil {
		return IncidentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IncidentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IncidentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IncidentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IncidentsClient) getHandleResponse(resp *http.Response) (IncidentsClientGetResponse, error) {
	result := IncidentsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Incident); err != nil {
		return IncidentsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all incidents.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - IncidentsClientListOptions contains the optional parameters for the IncidentsClient.List method.
func (client *IncidentsClient) List(resourceGroupName string, workspaceName string, options *IncidentsClientListOptions) *IncidentsClientListPager {
	return &IncidentsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
		},
		advancer: func(ctx context.Context, resp IncidentsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IncidentList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *IncidentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *IncidentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IncidentsClient) listHandleResponse(resp *http.Response) (IncidentsClientListResponse, error) {
	result := IncidentsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentList); err != nil {
		return IncidentsClientListResponse{}, err
	}
	return result, nil
}

// ListAlerts - Gets all incident alerts.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// incidentID - Incident ID
// options - IncidentsClientListAlertsOptions contains the optional parameters for the IncidentsClient.ListAlerts method.
func (client *IncidentsClient) ListAlerts(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListAlertsOptions) (IncidentsClientListAlertsResponse, error) {
	req, err := client.listAlertsCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
	if err != nil {
		return IncidentsClientListAlertsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IncidentsClientListAlertsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IncidentsClientListAlertsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listAlertsHandleResponse(resp)
}

// listAlertsCreateRequest creates the ListAlerts request.
func (client *IncidentsClient) listAlertsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListAlertsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/alerts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAlertsHandleResponse handles the ListAlerts response.
func (client *IncidentsClient) listAlertsHandleResponse(resp *http.Response) (IncidentsClientListAlertsResponse, error) {
	result := IncidentsClientListAlertsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentAlertList); err != nil {
		return IncidentsClientListAlertsResponse{}, err
	}
	return result, nil
}

// ListBookmarks - Gets all incident bookmarks.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// incidentID - Incident ID
// options - IncidentsClientListBookmarksOptions contains the optional parameters for the IncidentsClient.ListBookmarks method.
func (client *IncidentsClient) ListBookmarks(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListBookmarksOptions) (IncidentsClientListBookmarksResponse, error) {
	req, err := client.listBookmarksCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
	if err != nil {
		return IncidentsClientListBookmarksResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IncidentsClientListBookmarksResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IncidentsClientListBookmarksResponse{}, runtime.NewResponseError(resp)
	}
	return client.listBookmarksHandleResponse(resp)
}

// listBookmarksCreateRequest creates the ListBookmarks request.
func (client *IncidentsClient) listBookmarksCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListBookmarksOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/bookmarks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBookmarksHandleResponse handles the ListBookmarks response.
func (client *IncidentsClient) listBookmarksHandleResponse(resp *http.Response) (IncidentsClientListBookmarksResponse, error) {
	result := IncidentsClientListBookmarksResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentBookmarkList); err != nil {
		return IncidentsClientListBookmarksResponse{}, err
	}
	return result, nil
}

// ListEntities - Gets all incident related entities.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// incidentID - Incident ID
// options - IncidentsClientListEntitiesOptions contains the optional parameters for the IncidentsClient.ListEntities method.
func (client *IncidentsClient) ListEntities(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListEntitiesOptions) (IncidentsClientListEntitiesResponse, error) {
	req, err := client.listEntitiesCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
	if err != nil {
		return IncidentsClientListEntitiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IncidentsClientListEntitiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IncidentsClientListEntitiesResponse{}, runtime.NewResponseError(resp)
	}
	return client.listEntitiesHandleResponse(resp)
}

// listEntitiesCreateRequest creates the ListEntities request.
func (client *IncidentsClient) listEntitiesCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListEntitiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/entities"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listEntitiesHandleResponse handles the ListEntities response.
func (client *IncidentsClient) listEntitiesHandleResponse(resp *http.Response) (IncidentsClientListEntitiesResponse, error) {
	result := IncidentsClientListEntitiesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentEntitiesResponse); err != nil {
		return IncidentsClientListEntitiesResponse{}, err
	}
	return result, nil
}

// RunPlaybook - Triggers playbook on a specific incident
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - IncidentsClientRunPlaybookOptions contains the optional parameters for the IncidentsClient.RunPlaybook method.
func (client *IncidentsClient) RunPlaybook(ctx context.Context, resourceGroupName string, workspaceName string, incidentIdentifier string, options *IncidentsClientRunPlaybookOptions) (IncidentsClientRunPlaybookResponse, error) {
	req, err := client.runPlaybookCreateRequest(ctx, resourceGroupName, workspaceName, incidentIdentifier, options)
	if err != nil {
		return IncidentsClientRunPlaybookResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IncidentsClientRunPlaybookResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return IncidentsClientRunPlaybookResponse{}, runtime.NewResponseError(resp)
	}
	return client.runPlaybookHandleResponse(resp)
}

// runPlaybookCreateRequest creates the RunPlaybook request.
func (client *IncidentsClient) runPlaybookCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentIdentifier string, options *IncidentsClientRunPlaybookOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentIdentifier}/runPlaybook"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentIdentifier == "" {
		return nil, errors.New("parameter incidentIdentifier cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentIdentifier}", url.PathEscape(incidentIdentifier))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.RequestBody != nil {
		return req, runtime.MarshalAsJSON(req, *options.RequestBody)
	}
	return req, nil
}

// runPlaybookHandleResponse handles the RunPlaybook response.
func (client *IncidentsClient) runPlaybookHandleResponse(resp *http.Response) (IncidentsClientRunPlaybookResponse, error) {
	result := IncidentsClientRunPlaybookResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Object); err != nil {
		return IncidentsClientRunPlaybookResponse{}, err
	}
	return result, nil
}
