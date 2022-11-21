# Release History

## 0.9.0 (2022-11-21)
### Breaking Changes

- Type of `DefenderForServersAwsOfferingVMScannersConfiguration.ExclusionTags` has been changed from `interface{}` to `map[string]*string`
- Type of `ContactProperties.AlertNotifications` has been changed from `*AlertNotifications` to `*ContactPropertiesAlertNotifications`
- Const `AlertNotificationsOn` has been removed
- Const `AlertsToAdminsOn` has been removed
- Const `AlertsToAdminsOff` has been removed
- Const `AlertNotificationsOff` has been removed
- Type alias `AlertNotifications` has been removed
- Type alias `AlertsToAdmins` has been removed
- Function `PossibleAlertNotificationsValues` has been removed
- Function `*ContactsClient.Update` has been removed
- Function `PossibleAlertsToAdminsValues` has been removed
- Function `*AWSEnvironmentData.GetEnvironmentData` has been removed
- Struct `AWSEnvironmentData` has been removed
- Struct `ContactsClientUpdateOptions` has been removed
- Struct `ContactsClientUpdateResponse` has been removed
- Struct `DefenderFoDatabasesAwsOfferingArcAutoProvisioningServicePrincipalSecretMetadata` has been removed
- Struct `DefenderForDatabasesGcpOfferingArcAutoProvisioningConfiguration` has been removed
- Struct `DefenderForServersAwsOfferingArcAutoProvisioningServicePrincipalSecretMetadata` has been removed
- Struct `DefenderForServersGcpOfferingArcAutoProvisioningConfiguration` has been removed
- Field `ServicePrincipalSecretMetadata` of struct `DefenderFoDatabasesAwsOfferingArcAutoProvisioning` has been removed
- Field `Configuration` of struct `DefenderForServersGcpOfferingArcAutoProvisioning` has been removed
- Field `Configuration` of struct `DefenderForDatabasesGcpOfferingArcAutoProvisioning` has been removed
- Field `ServicePrincipalSecretMetadata` of struct `DefenderForServersAwsOfferingArcAutoProvisioning` has been removed
- Field `Email` of struct `ContactProperties` has been removed
- Field `AlertsToAdmins` of struct `ContactProperties` has been removed

### Features Added

- New const `OfferingTypeDefenderForDevOpsGithub`
- New const `ApplicationConditionOperatorEquals`
- New const `RolesContributor`
- New const `MinimalSeverityLow`
- New const `RolesAccountAdmin`
- New const `OfferingTypeDefenderCspmAws`
- New const `RolesServiceAdmin`
- New const `OfferingTypeDefenderForDevOpsAzureDevOps`
- New const `OfferingTypeDefenderCspmGcp`
- New const `MinimalSeverityHigh`
- New const `MinimalSeverityMedium`
- New const `RolesOwner`
- New type alias `MinimalSeverity`
- New type alias `Roles`
- New function `NewAPICollectionOffboardingClient(string, azcore.TokenCredential, *arm.ClientOptions) (*APICollectionOffboardingClient, error)`
- New function `*AwsEnvironmentData.GetEnvironmentData() *EnvironmentData`
- New function `*DefenderCspmGcpOffering.GetCloudOffering() *CloudOffering`
- New function `*DefenderForDevOpsAzureDevOpsOffering.GetCloudOffering() *CloudOffering`
- New function `*APICollectionClient.Get(context.Context, string, string, string, *APICollectionClientGetOptions) (APICollectionClientGetResponse, error)`
- New function `*DefenderForDevOpsGithubOffering.GetCloudOffering() *CloudOffering`
- New function `NewAPICollectionOnboardingClient(string, azcore.TokenCredential, *arm.ClientOptions) (*APICollectionOnboardingClient, error)`
- New function `NewAPICollectionClient(string, azcore.TokenCredential, *arm.ClientOptions) (*APICollectionClient, error)`
- New function `*APICollectionOnboardingClient.Create(context.Context, string, string, string, *APICollectionOnboardingClientCreateOptions) (APICollectionOnboardingClientCreateResponse, error)`
- New function `*APICollectionOffboardingClient.Delete(context.Context, string, string, string, *APICollectionOffboardingClientDeleteOptions) (APICollectionOffboardingClientDeleteResponse, error)`
- New function `PossibleMinimalSeverityValues() []MinimalSeverity`
- New function `PossibleRolesValues() []Roles`
- New function `*APICollectionClient.NewListPager(string, string, *APICollectionClientListOptions) *runtime.Pager[APICollectionClientListResponse]`
- New function `*DefenderCspmAwsOffering.GetCloudOffering() *CloudOffering`
- New struct `APICollectionClient`
- New struct `APICollectionClientGetOptions`
- New struct `APICollectionClientGetResponse`
- New struct `APICollectionClientListOptions`
- New struct `APICollectionClientListResponse`
- New struct `APICollectionOffboardingClient`
- New struct `APICollectionOffboardingClientDeleteOptions`
- New struct `APICollectionOffboardingClientDeleteResponse`
- New struct `APICollectionOnboardingClient`
- New struct `APICollectionOnboardingClientCreateOptions`
- New struct `APICollectionOnboardingClientCreateResponse`
- New struct `APICollectionProperties`
- New struct `APICollectionResponse`
- New struct `APICollectionResponseList`
- New struct `AwsEnvironmentData`
- New struct `ContactPropertiesAlertNotifications`
- New struct `ContactPropertiesNotificationsByRole`
- New struct `DefenderCspmAwsOffering`
- New struct `DefenderCspmAwsOfferingVMScanners`
- New struct `DefenderCspmAwsOfferingVMScannersConfiguration`
- New struct `DefenderCspmGcpOffering`
- New struct `DefenderFoDatabasesAwsOfferingRds`
- New struct `DefenderForDevOpsAzureDevOpsOffering`
- New struct `DefenderForDevOpsGithubOffering`
- New struct `ErrorDetail`
- New struct `ErrorResponse`
- New field `NotificationsByRole` in struct `ContactProperties`
- New field `Emails` in struct `ContactProperties`
- New field `Rds` in struct `DefenderFoDatabasesAwsOffering`


## 0.7.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.7.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).