package apimanagement

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// APIType enumerates the values for api type.
type APIType string

const (
	// HTTP ...
	HTTP APIType = "http"
	// Soap ...
	Soap APIType = "soap"
)

// PossibleAPITypeValues returns an array of possible values for the APIType const type.
func PossibleAPITypeValues() []APIType {
	return []APIType{HTTP, Soap}
}

// AsyncOperationStatus enumerates the values for async operation status.
type AsyncOperationStatus string

const (
	// Failed ...
	Failed AsyncOperationStatus = "Failed"
	// InProgress ...
	InProgress AsyncOperationStatus = "InProgress"
	// Started ...
	Started AsyncOperationStatus = "Started"
	// Succeeded ...
	Succeeded AsyncOperationStatus = "Succeeded"
)

// PossibleAsyncOperationStatusValues returns an array of possible values for the AsyncOperationStatus const type.
func PossibleAsyncOperationStatusValues() []AsyncOperationStatus {
	return []AsyncOperationStatus{Failed, InProgress, Started, Succeeded}
}

// AuthorizationMethod enumerates the values for authorization method.
type AuthorizationMethod string

const (
	// DELETE ...
	DELETE AuthorizationMethod = "DELETE"
	// GET ...
	GET AuthorizationMethod = "GET"
	// HEAD ...
	HEAD AuthorizationMethod = "HEAD"
	// OPTIONS ...
	OPTIONS AuthorizationMethod = "OPTIONS"
	// PATCH ...
	PATCH AuthorizationMethod = "PATCH"
	// POST ...
	POST AuthorizationMethod = "POST"
	// PUT ...
	PUT AuthorizationMethod = "PUT"
	// TRACE ...
	TRACE AuthorizationMethod = "TRACE"
)

// PossibleAuthorizationMethodValues returns an array of possible values for the AuthorizationMethod const type.
func PossibleAuthorizationMethodValues() []AuthorizationMethod {
	return []AuthorizationMethod{DELETE, GET, HEAD, OPTIONS, PATCH, POST, PUT, TRACE}
}

// BackendProtocol enumerates the values for backend protocol.
type BackendProtocol string

const (
	// BackendProtocolHTTP The Backend is a RESTful service.
	BackendProtocolHTTP BackendProtocol = "http"
	// BackendProtocolSoap The Backend is a SOAP service.
	BackendProtocolSoap BackendProtocol = "soap"
)

// PossibleBackendProtocolValues returns an array of possible values for the BackendProtocol const type.
func PossibleBackendProtocolValues() []BackendProtocol {
	return []BackendProtocol{BackendProtocolHTTP, BackendProtocolSoap}
}

// BearerTokenSendingMethod enumerates the values for bearer token sending method.
type BearerTokenSendingMethod string

const (
	// AuthorizationHeader ...
	AuthorizationHeader BearerTokenSendingMethod = "authorizationHeader"
	// Query ...
	Query BearerTokenSendingMethod = "query"
)

// PossibleBearerTokenSendingMethodValues returns an array of possible values for the BearerTokenSendingMethod const type.
func PossibleBearerTokenSendingMethodValues() []BearerTokenSendingMethod {
	return []BearerTokenSendingMethod{AuthorizationHeader, Query}
}

// BearerTokenSendingMethods enumerates the values for bearer token sending methods.
type BearerTokenSendingMethods string

const (
	// BearerTokenSendingMethodsAuthorizationHeader Access token will be transmitted in the Authorization
	// header using Bearer schema
	BearerTokenSendingMethodsAuthorizationHeader BearerTokenSendingMethods = "authorizationHeader"
	// BearerTokenSendingMethodsQuery Access token will be transmitted as query parameters.
	BearerTokenSendingMethodsQuery BearerTokenSendingMethods = "query"
)

// PossibleBearerTokenSendingMethodsValues returns an array of possible values for the BearerTokenSendingMethods const type.
func PossibleBearerTokenSendingMethodsValues() []BearerTokenSendingMethods {
	return []BearerTokenSendingMethods{BearerTokenSendingMethodsAuthorizationHeader, BearerTokenSendingMethodsQuery}
}

// ClientAuthenticationMethod enumerates the values for client authentication method.
type ClientAuthenticationMethod string

const (
	// Basic Basic Client Authentication method.
	Basic ClientAuthenticationMethod = "Basic"
	// Body Body based Authentication method.
	Body ClientAuthenticationMethod = "Body"
)

// PossibleClientAuthenticationMethodValues returns an array of possible values for the ClientAuthenticationMethod const type.
func PossibleClientAuthenticationMethodValues() []ClientAuthenticationMethod {
	return []ClientAuthenticationMethod{Basic, Body}
}

// Confirmation enumerates the values for confirmation.
type Confirmation string

const (
	// Invite Send an e-mail inviting the user to sign-up and complete registration.
	Invite Confirmation = "invite"
	// Signup Send an e-mail to the user confirming they have successfully signed up.
	Signup Confirmation = "signup"
)

// PossibleConfirmationValues returns an array of possible values for the Confirmation const type.
func PossibleConfirmationValues() []Confirmation {
	return []Confirmation{Invite, Signup}
}

// ConnectivityStatusType enumerates the values for connectivity status type.
type ConnectivityStatusType string

const (
	// Failure ...
	Failure ConnectivityStatusType = "failure"
	// Initializing ...
	Initializing ConnectivityStatusType = "initializing"
	// Success ...
	Success ConnectivityStatusType = "success"
)

// PossibleConnectivityStatusTypeValues returns an array of possible values for the ConnectivityStatusType const type.
func PossibleConnectivityStatusTypeValues() []ConnectivityStatusType {
	return []ConnectivityStatusType{Failure, Initializing, Success}
}

// ContentFormat enumerates the values for content format.
type ContentFormat string

const (
	// SwaggerJSON The contents are inline and Content Type is a OpenApi 2.0 Document.
	SwaggerJSON ContentFormat = "swagger-json"
	// SwaggerLinkJSON The Open Api 2.0 document is hosted on a publicly accessible internet address.
	SwaggerLinkJSON ContentFormat = "swagger-link-json"
	// WadlLinkJSON The WADL document is hosted on a publicly accessible internet address.
	WadlLinkJSON ContentFormat = "wadl-link-json"
	// WadlXML The contents are inline and Content type is a WADL document.
	WadlXML ContentFormat = "wadl-xml"
	// Wsdl The contents are inline and the document is a WSDL/Soap document.
	Wsdl ContentFormat = "wsdl"
	// WsdlLink The WSDL document is hosted on a publicly accessible internet address.
	WsdlLink ContentFormat = "wsdl-link"
)

// PossibleContentFormatValues returns an array of possible values for the ContentFormat const type.
func PossibleContentFormatValues() []ContentFormat {
	return []ContentFormat{SwaggerJSON, SwaggerLinkJSON, WadlLinkJSON, WadlXML, Wsdl, WsdlLink}
}

// ExportFormat enumerates the values for export format.
type ExportFormat string

const (
	// ExportFormatSwagger Export the Api Definition in OpenApi Specification 2.0 format to the Storage Blob.
	ExportFormatSwagger ExportFormat = "swagger-link"
	// ExportFormatWadl Export the Api Definition in WADL Schema to Storage Blob.
	ExportFormatWadl ExportFormat = "wadl-link"
	// ExportFormatWsdl Export the Api Definition in WSDL Schema to Storage Blob. This is only supported for
	// APIs of Type `soap`
	ExportFormatWsdl ExportFormat = "wsdl-link"
)

// PossibleExportFormatValues returns an array of possible values for the ExportFormat const type.
func PossibleExportFormatValues() []ExportFormat {
	return []ExportFormat{ExportFormatSwagger, ExportFormatWadl, ExportFormatWsdl}
}

// GrantType enumerates the values for grant type.
type GrantType string

const (
	// AuthorizationCode Authorization Code Grant flow as described
	// https://tools.ietf.org/html/rfc6749#section-4.1.
	AuthorizationCode GrantType = "authorizationCode"
	// ClientCredentials Client Credentials Grant flow as described
	// https://tools.ietf.org/html/rfc6749#section-4.4.
	ClientCredentials GrantType = "clientCredentials"
	// Implicit Implicit Code Grant flow as described https://tools.ietf.org/html/rfc6749#section-4.2.
	Implicit GrantType = "implicit"
	// ResourceOwnerPassword Resource Owner Password Grant flow as described
	// https://tools.ietf.org/html/rfc6749#section-4.3.
	ResourceOwnerPassword GrantType = "resourceOwnerPassword"
)

// PossibleGrantTypeValues returns an array of possible values for the GrantType const type.
func PossibleGrantTypeValues() []GrantType {
	return []GrantType{AuthorizationCode, ClientCredentials, Implicit, ResourceOwnerPassword}
}

// GroupType enumerates the values for group type.
type GroupType string

const (
	// Custom ...
	Custom GroupType = "custom"
	// External ...
	External GroupType = "external"
	// System ...
	System GroupType = "system"
)

// PossibleGroupTypeValues returns an array of possible values for the GroupType const type.
func PossibleGroupTypeValues() []GroupType {
	return []GroupType{Custom, External, System}
}

// HostnameType enumerates the values for hostname type.
type HostnameType string

const (
	// Management ...
	Management HostnameType = "Management"
	// Portal ...
	Portal HostnameType = "Portal"
	// Proxy ...
	Proxy HostnameType = "Proxy"
	// Scm ...
	Scm HostnameType = "Scm"
)

// PossibleHostnameTypeValues returns an array of possible values for the HostnameType const type.
func PossibleHostnameTypeValues() []HostnameType {
	return []HostnameType{Management, Portal, Proxy, Scm}
}

// IdentityProviderType enumerates the values for identity provider type.
type IdentityProviderType string

const (
	// Aad Azure Active Directory as Identity provider.
	Aad IdentityProviderType = "aad"
	// AadB2C Azure Active Directory B2C as Identity provider.
	AadB2C IdentityProviderType = "aadB2C"
	// Facebook Facebook as Identity provider.
	Facebook IdentityProviderType = "facebook"
	// Google Google as Identity provider.
	Google IdentityProviderType = "google"
	// Microsoft Microsoft Live as Identity provider.
	Microsoft IdentityProviderType = "microsoft"
	// Twitter Twitter as Identity provider.
	Twitter IdentityProviderType = "twitter"
)

// PossibleIdentityProviderTypeValues returns an array of possible values for the IdentityProviderType const type.
func PossibleIdentityProviderTypeValues() []IdentityProviderType {
	return []IdentityProviderType{Aad, AadB2C, Facebook, Google, Microsoft, Twitter}
}

// KeyType enumerates the values for key type.
type KeyType string

const (
	// Primary ...
	Primary KeyType = "primary"
	// Secondary ...
	Secondary KeyType = "secondary"
)

// PossibleKeyTypeValues returns an array of possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{Primary, Secondary}
}

// LoggerType enumerates the values for logger type.
type LoggerType string

const (
	// ApplicationInsights Azure Application Insights as log destination.
	ApplicationInsights LoggerType = "applicationInsights"
	// AzureEventHub Azure Event Hub as log destination.
	AzureEventHub LoggerType = "azureEventHub"
)

// PossibleLoggerTypeValues returns an array of possible values for the LoggerType const type.
func PossibleLoggerTypeValues() []LoggerType {
	return []LoggerType{ApplicationInsights, AzureEventHub}
}

// NameAvailabilityReason enumerates the values for name availability reason.
type NameAvailabilityReason string

const (
	// AlreadyExists ...
	AlreadyExists NameAvailabilityReason = "AlreadyExists"
	// Invalid ...
	Invalid NameAvailabilityReason = "Invalid"
	// Valid ...
	Valid NameAvailabilityReason = "Valid"
)

// PossibleNameAvailabilityReasonValues returns an array of possible values for the NameAvailabilityReason const type.
func PossibleNameAvailabilityReasonValues() []NameAvailabilityReason {
	return []NameAvailabilityReason{AlreadyExists, Invalid, Valid}
}

// NotificationName enumerates the values for notification name.
type NotificationName string

const (
	// AccountClosedPublisher The following email recipients and users will receive email notifications when
	// developer closes his account.
	AccountClosedPublisher NotificationName = "AccountClosedPublisher"
	// BCC The following recipients will receive blind carbon copies of all emails sent to developers.
	BCC NotificationName = "BCC"
	// NewApplicationNotificationMessage The following email recipients and users will receive email
	// notifications when new applications are submitted to the application gallery.
	NewApplicationNotificationMessage NotificationName = "NewApplicationNotificationMessage"
	// NewIssuePublisherNotificationMessage The following email recipients and users will receive email
	// notifications when a new issue or comment is submitted on the developer portal.
	NewIssuePublisherNotificationMessage NotificationName = "NewIssuePublisherNotificationMessage"
	// PurchasePublisherNotificationMessage The following email recipients and users will receive email
	// notifications about new API product subscriptions.
	PurchasePublisherNotificationMessage NotificationName = "PurchasePublisherNotificationMessage"
	// QuotaLimitApproachingPublisherNotificationMessage The following email recipients and users will receive
	// email notifications when subscription usage gets close to usage quota.
	QuotaLimitApproachingPublisherNotificationMessage NotificationName = "QuotaLimitApproachingPublisherNotificationMessage"
	// RequestPublisherNotificationMessage The following email recipients and users will receive email
	// notifications about subscription requests for API products requiring approval.
	RequestPublisherNotificationMessage NotificationName = "RequestPublisherNotificationMessage"
)

// PossibleNotificationNameValues returns an array of possible values for the NotificationName const type.
func PossibleNotificationNameValues() []NotificationName {
	return []NotificationName{AccountClosedPublisher, BCC, NewApplicationNotificationMessage, NewIssuePublisherNotificationMessage, PurchasePublisherNotificationMessage, QuotaLimitApproachingPublisherNotificationMessage, RequestPublisherNotificationMessage}
}

// PolicyContentFormat enumerates the values for policy content format.
type PolicyContentFormat string

const (
	// Rawxml The contents are inline and Content type is a non XML encoded policy document.
	Rawxml PolicyContentFormat = "rawxml"
	// RawxmlLink The policy document is not Xml encoded and is hosted on a http endpoint accessible from the
	// API Management service.
	RawxmlLink PolicyContentFormat = "rawxml-link"
	// XML The contents are inline and Content type is an XML document.
	XML PolicyContentFormat = "xml"
	// XMLLink The policy XML document is hosted on a http endpoint accessible from the API Management service.
	XMLLink PolicyContentFormat = "xml-link"
)

// PossiblePolicyContentFormatValues returns an array of possible values for the PolicyContentFormat const type.
func PossiblePolicyContentFormatValues() []PolicyContentFormat {
	return []PolicyContentFormat{Rawxml, RawxmlLink, XML, XMLLink}
}

// PolicyScopeContract enumerates the values for policy scope contract.
type PolicyScopeContract string

const (
	// PolicyScopeContractAll ...
	PolicyScopeContractAll PolicyScopeContract = "All"
	// PolicyScopeContractAPI ...
	PolicyScopeContractAPI PolicyScopeContract = "Api"
	// PolicyScopeContractOperation ...
	PolicyScopeContractOperation PolicyScopeContract = "Operation"
	// PolicyScopeContractProduct ...
	PolicyScopeContractProduct PolicyScopeContract = "Product"
	// PolicyScopeContractTenant ...
	PolicyScopeContractTenant PolicyScopeContract = "Tenant"
)

// PossiblePolicyScopeContractValues returns an array of possible values for the PolicyScopeContract const type.
func PossiblePolicyScopeContractValues() []PolicyScopeContract {
	return []PolicyScopeContract{PolicyScopeContractAll, PolicyScopeContractAPI, PolicyScopeContractOperation, PolicyScopeContractProduct, PolicyScopeContractTenant}
}

// ProductState enumerates the values for product state.
type ProductState string

const (
	// NotPublished ...
	NotPublished ProductState = "notPublished"
	// Published ...
	Published ProductState = "published"
)

// PossibleProductStateValues returns an array of possible values for the ProductState const type.
func PossibleProductStateValues() []ProductState {
	return []ProductState{NotPublished, Published}
}

// Protocol enumerates the values for protocol.
type Protocol string

const (
	// ProtocolHTTP ...
	ProtocolHTTP Protocol = "http"
	// ProtocolHTTPS ...
	ProtocolHTTPS Protocol = "https"
)

// PossibleProtocolValues returns an array of possible values for the Protocol const type.
func PossibleProtocolValues() []Protocol {
	return []Protocol{ProtocolHTTP, ProtocolHTTPS}
}

// ResourceSkuCapacityScaleType enumerates the values for resource sku capacity scale type.
type ResourceSkuCapacityScaleType string

const (
	// Automatic ...
	Automatic ResourceSkuCapacityScaleType = "automatic"
	// Manual ...
	Manual ResourceSkuCapacityScaleType = "manual"
	// None ...
	None ResourceSkuCapacityScaleType = "none"
)

// PossibleResourceSkuCapacityScaleTypeValues returns an array of possible values for the ResourceSkuCapacityScaleType const type.
func PossibleResourceSkuCapacityScaleTypeValues() []ResourceSkuCapacityScaleType {
	return []ResourceSkuCapacityScaleType{Automatic, Manual, None}
}

// SkuType enumerates the values for sku type.
type SkuType string

const (
	// SkuTypeBasic Basic SKU of Api Management.
	SkuTypeBasic SkuType = "Basic"
	// SkuTypeDeveloper Developer SKU of Api Management.
	SkuTypeDeveloper SkuType = "Developer"
	// SkuTypePremium Premium SKU of Api Management.
	SkuTypePremium SkuType = "Premium"
	// SkuTypeStandard Standard SKU of Api Management.
	SkuTypeStandard SkuType = "Standard"
)

// PossibleSkuTypeValues returns an array of possible values for the SkuType const type.
func PossibleSkuTypeValues() []SkuType {
	return []SkuType{SkuTypeBasic, SkuTypeDeveloper, SkuTypePremium, SkuTypeStandard}
}

// SoapAPIType enumerates the values for soap api type.
type SoapAPIType string

const (
	// SoapPassThrough Imports the Soap API having a SOAP front end.
	SoapPassThrough SoapAPIType = "soap"
	// SoapToRest Imports a SOAP API having a RESTful front end.
	SoapToRest SoapAPIType = "http"
)

// PossibleSoapAPITypeValues returns an array of possible values for the SoapAPIType const type.
func PossibleSoapAPITypeValues() []SoapAPIType {
	return []SoapAPIType{SoapPassThrough, SoapToRest}
}

// State enumerates the values for state.
type State string

const (
	// Closed The issue was closed.
	Closed State = "closed"
	// Open The issue is opened.
	Open State = "open"
	// Proposed The issue is proposed.
	Proposed State = "proposed"
	// Removed The issue was removed.
	Removed State = "removed"
	// Resolved The issue is now resolved.
	Resolved State = "resolved"
)

// PossibleStateValues returns an array of possible values for the State const type.
func PossibleStateValues() []State {
	return []State{Closed, Open, Proposed, Removed, Resolved}
}

// StoreName enumerates the values for store name.
type StoreName string

const (
	// CertificateAuthority ...
	CertificateAuthority StoreName = "CertificateAuthority"
	// Root ...
	Root StoreName = "Root"
)

// PossibleStoreNameValues returns an array of possible values for the StoreName const type.
func PossibleStoreNameValues() []StoreName {
	return []StoreName{CertificateAuthority, Root}
}

// SubscriptionState enumerates the values for subscription state.
type SubscriptionState string

const (
	// Active ...
	Active SubscriptionState = "active"
	// Cancelled ...
	Cancelled SubscriptionState = "cancelled"
	// Expired ...
	Expired SubscriptionState = "expired"
	// Rejected ...
	Rejected SubscriptionState = "rejected"
	// Submitted ...
	Submitted SubscriptionState = "submitted"
	// Suspended ...
	Suspended SubscriptionState = "suspended"
)

// PossibleSubscriptionStateValues returns an array of possible values for the SubscriptionState const type.
func PossibleSubscriptionStateValues() []SubscriptionState {
	return []SubscriptionState{Active, Cancelled, Expired, Rejected, Submitted, Suspended}
}

// TemplateName enumerates the values for template name.
type TemplateName string

const (
	// AccountClosedDeveloper ...
	AccountClosedDeveloper TemplateName = "accountClosedDeveloper"
	// ApplicationApprovedNotificationMessage ...
	ApplicationApprovedNotificationMessage TemplateName = "applicationApprovedNotificationMessage"
	// ConfirmSignUpIdentityDefault ...
	ConfirmSignUpIdentityDefault TemplateName = "confirmSignUpIdentityDefault"
	// EmailChangeIdentityDefault ...
	EmailChangeIdentityDefault TemplateName = "emailChangeIdentityDefault"
	// InviteUserNotificationMessage ...
	InviteUserNotificationMessage TemplateName = "inviteUserNotificationMessage"
	// NewCommentNotificationMessage ...
	NewCommentNotificationMessage TemplateName = "newCommentNotificationMessage"
	// NewDeveloperNotificationMessage ...
	NewDeveloperNotificationMessage TemplateName = "newDeveloperNotificationMessage"
	// NewIssueNotificationMessage ...
	NewIssueNotificationMessage TemplateName = "newIssueNotificationMessage"
	// PasswordResetByAdminNotificationMessage ...
	PasswordResetByAdminNotificationMessage TemplateName = "passwordResetByAdminNotificationMessage"
	// PasswordResetIdentityDefault ...
	PasswordResetIdentityDefault TemplateName = "passwordResetIdentityDefault"
	// PurchaseDeveloperNotificationMessage ...
	PurchaseDeveloperNotificationMessage TemplateName = "purchaseDeveloperNotificationMessage"
	// QuotaLimitApproachingDeveloperNotificationMessage ...
	QuotaLimitApproachingDeveloperNotificationMessage TemplateName = "quotaLimitApproachingDeveloperNotificationMessage"
	// RejectDeveloperNotificationMessage ...
	RejectDeveloperNotificationMessage TemplateName = "rejectDeveloperNotificationMessage"
	// RequestDeveloperNotificationMessage ...
	RequestDeveloperNotificationMessage TemplateName = "requestDeveloperNotificationMessage"
)

// PossibleTemplateNameValues returns an array of possible values for the TemplateName const type.
func PossibleTemplateNameValues() []TemplateName {
	return []TemplateName{AccountClosedDeveloper, ApplicationApprovedNotificationMessage, ConfirmSignUpIdentityDefault, EmailChangeIdentityDefault, InviteUserNotificationMessage, NewCommentNotificationMessage, NewDeveloperNotificationMessage, NewIssueNotificationMessage, PasswordResetByAdminNotificationMessage, PasswordResetIdentityDefault, PurchaseDeveloperNotificationMessage, QuotaLimitApproachingDeveloperNotificationMessage, RejectDeveloperNotificationMessage, RequestDeveloperNotificationMessage}
}

// UserState enumerates the values for user state.
type UserState string

const (
	// UserStateActive User state is active.
	UserStateActive UserState = "active"
	// UserStateBlocked User is blocked. Blocked users cannot authenticate at developer portal or call API.
	UserStateBlocked UserState = "blocked"
	// UserStateDeleted User account is closed. All identities and related entities are removed.
	UserStateDeleted UserState = "deleted"
	// UserStatePending User account is pending. Requires identity confirmation before it can be made active.
	UserStatePending UserState = "pending"
)

// PossibleUserStateValues returns an array of possible values for the UserState const type.
func PossibleUserStateValues() []UserState {
	return []UserState{UserStateActive, UserStateBlocked, UserStateDeleted, UserStatePending}
}

// VersioningScheme enumerates the values for versioning scheme.
type VersioningScheme string

const (
	// VersioningSchemeHeader The API Version is passed in a HTTP header.
	VersioningSchemeHeader VersioningScheme = "Header"
	// VersioningSchemeQuery The API Version is passed in a query parameter.
	VersioningSchemeQuery VersioningScheme = "Query"
	// VersioningSchemeSegment The API Version is passed in a path segment.
	VersioningSchemeSegment VersioningScheme = "Segment"
)

// PossibleVersioningSchemeValues returns an array of possible values for the VersioningScheme const type.
func PossibleVersioningSchemeValues() []VersioningScheme {
	return []VersioningScheme{VersioningSchemeHeader, VersioningSchemeQuery, VersioningSchemeSegment}
}

// VersioningScheme1 enumerates the values for versioning scheme 1.
type VersioningScheme1 string

const (
	// VersioningScheme1Header ...
	VersioningScheme1Header VersioningScheme1 = "Header"
	// VersioningScheme1Query ...
	VersioningScheme1Query VersioningScheme1 = "Query"
	// VersioningScheme1Segment ...
	VersioningScheme1Segment VersioningScheme1 = "Segment"
)

// PossibleVersioningScheme1Values returns an array of possible values for the VersioningScheme1 const type.
func PossibleVersioningScheme1Values() []VersioningScheme1 {
	return []VersioningScheme1{VersioningScheme1Header, VersioningScheme1Query, VersioningScheme1Segment}
}

// VirtualNetworkType enumerates the values for virtual network type.
type VirtualNetworkType string

const (
	// VirtualNetworkTypeExternal The service is part of Virtual Network and it is accessible from Internet.
	VirtualNetworkTypeExternal VirtualNetworkType = "External"
	// VirtualNetworkTypeInternal The service is part of Virtual Network and it is only accessible from within
	// the virtual network.
	VirtualNetworkTypeInternal VirtualNetworkType = "Internal"
	// VirtualNetworkTypeNone The service is not part of any Virtual Network.
	VirtualNetworkTypeNone VirtualNetworkType = "None"
)

// PossibleVirtualNetworkTypeValues returns an array of possible values for the VirtualNetworkType const type.
func PossibleVirtualNetworkTypeValues() []VirtualNetworkType {
	return []VirtualNetworkType{VirtualNetworkTypeExternal, VirtualNetworkTypeInternal, VirtualNetworkTypeNone}
}
