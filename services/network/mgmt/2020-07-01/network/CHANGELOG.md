# Unreleased

## Breaking Changes

### Removed Funcs

1. *ApplicationSecurityGroupsDeleteFuture.UnmarshalJSON([]byte) error
1. *InterfacesDeleteFuture.UnmarshalJSON([]byte) error
1. ApplicationSecurityGroupsClient.Delete(context.Context, string, string) (ApplicationSecurityGroupsDeleteFuture, error)
1. ApplicationSecurityGroupsClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. ApplicationSecurityGroupsClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. ApplicationSecurityGroupsClient.DeleteSender(*http.Request) (ApplicationSecurityGroupsDeleteFuture, error)
1. InterfacesClient.Delete(context.Context, string, string) (InterfacesDeleteFuture, error)
1. InterfacesClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteSender(*http.Request) (InterfacesDeleteFuture, error)

### Struct Changes

#### Removed Structs

1. ApplicationSecurityGroupsDeleteFuture
1. InterfacesDeleteFuture

#### Removed Struct Fields

1. VirtualNetworkGatewayPropertiesFormat.ExtendedLocation
1. VirtualNetworkGatewayPropertiesFormat.VirtualNetworkExtendedLocationResourceID

### Signature Changes

#### Struct Fields

1. SubnetPropertiesFormat.PrivateEndpointNetworkPolicies changed type from *string to VirtualNetworkPrivateEndpointNetworkPolicies
1. SubnetPropertiesFormat.PrivateLinkServiceNetworkPolicies changed type from *string to VirtualNetworkPrivateLinkServiceNetworkPolicies

## Additive Changes

### New Constants

1. InterfaceNicType.InterfaceNicTypeElastic
1. InterfaceNicType.InterfaceNicTypeStandard
1. PublicIPAddressMigrationPhase.PublicIPAddressMigrationPhaseAbort
1. PublicIPAddressMigrationPhase.PublicIPAddressMigrationPhaseCommit
1. PublicIPAddressMigrationPhase.PublicIPAddressMigrationPhaseCommitted
1. PublicIPAddressMigrationPhase.PublicIPAddressMigrationPhaseNone
1. PublicIPAddressMigrationPhase.PublicIPAddressMigrationPhasePrepare
1. VirtualNetworkPrivateEndpointNetworkPolicies.VirtualNetworkPrivateEndpointNetworkPoliciesDisabled
1. VirtualNetworkPrivateEndpointNetworkPolicies.VirtualNetworkPrivateEndpointNetworkPoliciesEnabled
1. VirtualNetworkPrivateLinkServiceNetworkPolicies.VirtualNetworkPrivateLinkServiceNetworkPoliciesDisabled
1. VirtualNetworkPrivateLinkServiceNetworkPolicies.VirtualNetworkPrivateLinkServiceNetworkPoliciesEnabled

### New Funcs

1. *ApplicationSecurityGroupsDeleteABCFuture.UnmarshalJSON([]byte) error
1. *InterfacesDeleteABCDFuture.UnmarshalJSON([]byte) error
1. ApplicationSecurityGroupsClient.DeleteABC(context.Context, string, string) (ApplicationSecurityGroupsDeleteABCFuture, error)
1. ApplicationSecurityGroupsClient.DeleteABCPreparer(context.Context, string, string) (*http.Request, error)
1. ApplicationSecurityGroupsClient.DeleteABCResponder(*http.Response) (autorest.Response, error)
1. ApplicationSecurityGroupsClient.DeleteABCSender(*http.Request) (ApplicationSecurityGroupsDeleteABCFuture, error)
1. InterfacesClient.DeleteABCD(context.Context, string, string) (InterfacesDeleteABCDFuture, error)
1. InterfacesClient.DeleteABCDPreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteABCDResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteABCDSender(*http.Request) (InterfacesDeleteABCDFuture, error)
1. PossibleInterfaceNicTypeValues() []InterfaceNicType
1. PossiblePublicIPAddressMigrationPhaseValues() []PublicIPAddressMigrationPhase
1. PossibleVirtualNetworkPrivateEndpointNetworkPoliciesValues() []VirtualNetworkPrivateEndpointNetworkPolicies
1. PossibleVirtualNetworkPrivateLinkServiceNetworkPoliciesValues() []VirtualNetworkPrivateLinkServiceNetworkPolicies

### Struct Changes

#### New Structs

1. ApplicationSecurityGroupsDeleteABCFuture
1. InterfacesDeleteABCDFuture

#### New Struct Fields

1. AvailablePrivateEndpointType.DisplayName
1. IPAddressAvailabilityResult.IsPlatformReserved
1. InterfaceIPConfiguration.Type
1. InterfacePropertiesFormat.NicType
1. InterfacePropertiesFormat.PrivateLinkService
1. PublicIPAddressPropertiesFormat.LinkedPublicIPAddress
1. PublicIPAddressPropertiesFormat.MigrationPhase
1. PublicIPAddressPropertiesFormat.NatGateway
1. PublicIPPrefixPropertiesFormat.NatGateway
1. ServiceTagInformationPropertiesFormat.State
1. Subnet.Type
1. VirtualNetworkGateway.ExtendedLocation
1. VirtualNetworkGatewayPropertiesFormat.VNetExtendedLocationResourceID
1. VirtualNetworkPeering.Type
1. VirtualNetworkPeeringPropertiesFormat.ResourceGUID
