# Unreleased

## Breaking Changes

### Signature Changes

#### Funcs

1. ProviderClient.GetAvailableStacks
	- Params
		- From: context.Context, string
		- To: context.Context, ProviderOsTypeSelected
1. ProviderClient.GetAvailableStacksComplete
	- Params
		- From: context.Context, string
		- To: context.Context, ProviderOsTypeSelected
1. ProviderClient.GetAvailableStacksOnPrem
	- Params
		- From: context.Context, string
		- To: context.Context, ProviderOsTypeSelected
1. ProviderClient.GetAvailableStacksOnPremComplete
	- Params
		- From: context.Context, string
		- To: context.Context, ProviderOsTypeSelected
1. ProviderClient.GetAvailableStacksOnPremPreparer
	- Params
		- From: context.Context, string
		- To: context.Context, ProviderOsTypeSelected
1. ProviderClient.GetAvailableStacksPreparer
	- Params
		- From: context.Context, string
		- To: context.Context, ProviderOsTypeSelected
1. ProviderClient.GetFunctionAppStacks
	- Params
		- From: context.Context, string
		- To: context.Context, ProviderStackOsType
1. ProviderClient.GetFunctionAppStacksComplete
	- Params
		- From: context.Context, string
		- To: context.Context, ProviderStackOsType
1. ProviderClient.GetFunctionAppStacksForLocation
	- Params
		- From: context.Context, string, string
		- To: context.Context, string, ProviderStackOsType
1. ProviderClient.GetFunctionAppStacksForLocationComplete
	- Params
		- From: context.Context, string, string
		- To: context.Context, string, ProviderStackOsType
1. ProviderClient.GetFunctionAppStacksForLocationPreparer
	- Params
		- From: context.Context, string, string
		- To: context.Context, string, ProviderStackOsType
1. ProviderClient.GetFunctionAppStacksPreparer
	- Params
		- From: context.Context, string
		- To: context.Context, ProviderStackOsType
1. ProviderClient.GetWebAppStacks
	- Params
		- From: context.Context, string
		- To: context.Context, ProviderStackOsType
1. ProviderClient.GetWebAppStacksComplete
	- Params
		- From: context.Context, string
		- To: context.Context, ProviderStackOsType
1. ProviderClient.GetWebAppStacksForLocation
	- Params
		- From: context.Context, string, string
		- To: context.Context, string, ProviderStackOsType
1. ProviderClient.GetWebAppStacksForLocationComplete
	- Params
		- From: context.Context, string, string
		- To: context.Context, string, ProviderStackOsType
1. ProviderClient.GetWebAppStacksForLocationPreparer
	- Params
		- From: context.Context, string, string
		- To: context.Context, string, ProviderStackOsType
1. ProviderClient.GetWebAppStacksPreparer
	- Params
		- From: context.Context, string
		- To: context.Context, ProviderStackOsType

#### Struct Fields

1. AppServiceCertificateOrderPatchResourceProperties.AppServiceCertificateNotRenewableReasons changed type from *[]string to *[]ResourceNotRenewableReason
1. AppServiceCertificateOrderProperties.AppServiceCertificateNotRenewableReasons changed type from *[]string to *[]ResourceNotRenewableReason

## Additive Changes

### New Constants

1. ProviderOsTypeSelected.ProviderOsTypeSelectedAll
1. ProviderOsTypeSelected.ProviderOsTypeSelectedLinux
1. ProviderOsTypeSelected.ProviderOsTypeSelectedLinuxFunctions
1. ProviderOsTypeSelected.ProviderOsTypeSelectedWindows
1. ProviderOsTypeSelected.ProviderOsTypeSelectedWindowsFunctions
1. ProviderStackOsType.ProviderStackOsTypeAll
1. ProviderStackOsType.ProviderStackOsTypeLinux
1. ProviderStackOsType.ProviderStackOsTypeWindows
1. ResourceNotRenewableReason.ResourceNotRenewableReasonExpirationNotInRenewalTimeRange
1. ResourceNotRenewableReason.ResourceNotRenewableReasonRegistrationStatusNotSupportedForRenewal
1. ResourceNotRenewableReason.ResourceNotRenewableReasonSubscriptionNotActive

### New Funcs

1. AppsClient.CreateOneDeployOperation(context.Context, string, string) (SetObject, error)
1. AppsClient.CreateOneDeployOperationPreparer(context.Context, string, string) (*http.Request, error)
1. AppsClient.CreateOneDeployOperationResponder(*http.Response) (SetObject, error)
1. AppsClient.CreateOneDeployOperationSender(*http.Request) (*http.Response, error)
1. AppsClient.GetAuthSettingsV2WithoutSecretsSlot(context.Context, string, string, string) (SiteAuthSettingsV2, error)
1. AppsClient.GetAuthSettingsV2WithoutSecretsSlotPreparer(context.Context, string, string, string) (*http.Request, error)
1. AppsClient.GetAuthSettingsV2WithoutSecretsSlotResponder(*http.Response) (SiteAuthSettingsV2, error)
1. AppsClient.GetAuthSettingsV2WithoutSecretsSlotSender(*http.Request) (*http.Response, error)
1. AppsClient.GetOneDeployStatus(context.Context, string, string) (SetObject, error)
1. AppsClient.GetOneDeployStatusPreparer(context.Context, string, string) (*http.Request, error)
1. AppsClient.GetOneDeployStatusResponder(*http.Response) (SetObject, error)
1. AppsClient.GetOneDeployStatusSender(*http.Request) (*http.Response, error)
1. PossibleProviderOsTypeSelectedValues() []ProviderOsTypeSelected
1. PossibleProviderStackOsTypeValues() []ProviderStackOsType
1. PossibleResourceNotRenewableReasonValues() []ResourceNotRenewableReason

### Struct Changes

#### New Struct Fields

1. AppServicePlanPatchResourceProperties.NumberOfWorkers
1. AppServicePlanProperties.NumberOfWorkers
