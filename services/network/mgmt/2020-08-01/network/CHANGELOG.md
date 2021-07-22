# Unreleased

## Breaking Changes

### Removed Funcs

1. *InterfacesDeleteFuture.UnmarshalJSON([]byte) error
1. InterfacesClient.Delete(context.Context, string, string) (InterfacesDeleteFuture, error)
1. InterfacesClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteSender(*http.Request) (InterfacesDeleteFuture, error)

### Struct Changes

#### Removed Structs

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

1. InterfaceMigrationPhase.InterfaceMigrationPhaseAbort
1. InterfaceMigrationPhase.InterfaceMigrationPhaseCommit
1. InterfaceMigrationPhase.InterfaceMigrationPhaseCommitted
1. InterfaceMigrationPhase.InterfaceMigrationPhaseNone
1. InterfaceMigrationPhase.InterfaceMigrationPhasePrepare
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

1. *InterfacesDeleteABCDFuture.UnmarshalJSON([]byte) error
1. InterfacesClient.DeleteABCD(context.Context, string, string) (InterfacesDeleteABCDFuture, error)
1. InterfacesClient.DeleteABCDPreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteABCDResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteABCDSender(*http.Request) (InterfacesDeleteABCDFuture, error)
1. PossibleInterfaceMigrationPhaseValues() []InterfaceMigrationPhase
1. PossibleInterfaceNicTypeValues() []InterfaceNicType
1. PossiblePublicIPAddressMigrationPhaseValues() []PublicIPAddressMigrationPhase
1. PossibleVirtualNetworkPrivateEndpointNetworkPoliciesValues() []VirtualNetworkPrivateEndpointNetworkPolicies
1. PossibleVirtualNetworkPrivateLinkServiceNetworkPoliciesValues() []VirtualNetworkPrivateLinkServiceNetworkPolicies
1. ProxyResource.MarshalJSON() ([]byte, error)

### Struct Changes

#### New Structs

1. InterfacesDeleteABCDFuture
1. ProxyResource

#### New Struct Fields

1. Delegation.Type
1. InterfaceIPConfiguration.Type
1. InterfacePropertiesFormat.MigrationPhase
1. InterfacePropertiesFormat.NicType
1. InterfacePropertiesFormat.PrivateLinkService
1. PublicIPAddressPropertiesFormat.LinkedPublicIPAddress
1. PublicIPAddressPropertiesFormat.MigrationPhase
1. PublicIPAddressPropertiesFormat.NatGateway
1. PublicIPAddressPropertiesFormat.ServicePublicIPAddress
1. PublicIPPrefixPropertiesFormat.NatGateway
1. Subnet.Type
1. SubnetPropertiesFormat.ApplicationGatewayIPConfigurations
1. VirtualNetworkGateway.ExtendedLocation
1. VirtualNetworkGatewayPropertiesFormat.VNetExtendedLocationResourceID
1. VirtualNetworkPeering.Type
1. VirtualNetworkPeeringPropertiesFormat.DoNotVerifyRemoteGateways
1. VirtualNetworkPeeringPropertiesFormat.ResourceGUID
