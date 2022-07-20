# Unreleased

## Breaking Changes

### Removed Constants

1. ResourceType.ResourceTypeMicrosoftComputeavailabilitySets
1. ResourceType.ResourceTypeMicrosoftComputediskEncryptionSets
1. ResourceType.ResourceTypeMicrosoftComputevirtualMachines
1. ResourceType.ResourceTypeMicrosoftKeyVaultvaults
1. ResourceType.ResourceTypeMicrosoftNetworkloadBalancers
1. ResourceType.ResourceTypeMicrosoftNetworknetworkInterfaces
1. ResourceType.ResourceTypeMicrosoftNetworknetworkSecurityGroups
1. ResourceType.ResourceTypeMicrosoftNetworkpublicIPAddresses
1. ResourceType.ResourceTypeMicrosoftNetworkvirtualNetworks
1. ResourceType.ResourceTypeMicrosoftSqlservers
1. ResourceType.ResourceTypeMicrosoftSqlserversdatabases
1. ResourceType.ResourceTypeMicrosoftSqlserverselasticPools
1. ResourceType.ResourceTypeResourceGroups
1. TargetAvailabilityZone.NA
1. TargetAvailabilityZone.One
1. TargetAvailabilityZone.Three
1. TargetAvailabilityZone.Two
1. ZoneRedundant.Disable
1. ZoneRedundant.Enable

### Removed Funcs

1. AvailabilitySetResourceSettings.AsAvailabilitySetResourceSettings() (*AvailabilitySetResourceSettings, bool)
1. AvailabilitySetResourceSettings.AsBasicResourceSettings() (BasicResourceSettings, bool)
1. AvailabilitySetResourceSettings.AsDiskEncryptionSetResourceSettings() (*DiskEncryptionSetResourceSettings, bool)
1. AvailabilitySetResourceSettings.AsKeyVaultResourceSettings() (*KeyVaultResourceSettings, bool)
1. AvailabilitySetResourceSettings.AsLoadBalancerResourceSettings() (*LoadBalancerResourceSettings, bool)
1. AvailabilitySetResourceSettings.AsNetworkInterfaceResourceSettings() (*NetworkInterfaceResourceSettings, bool)
1. AvailabilitySetResourceSettings.AsNetworkSecurityGroupResourceSettings() (*NetworkSecurityGroupResourceSettings, bool)
1. AvailabilitySetResourceSettings.AsPublicIPAddressResourceSettings() (*PublicIPAddressResourceSettings, bool)
1. AvailabilitySetResourceSettings.AsResourceGroupResourceSettings() (*ResourceGroupResourceSettings, bool)
1. AvailabilitySetResourceSettings.AsResourceSettings() (*ResourceSettings, bool)
1. AvailabilitySetResourceSettings.AsSQLDatabaseResourceSettings() (*SQLDatabaseResourceSettings, bool)
1. AvailabilitySetResourceSettings.AsSQLElasticPoolResourceSettings() (*SQLElasticPoolResourceSettings, bool)
1. AvailabilitySetResourceSettings.AsSQLServerResourceSettings() (*SQLServerResourceSettings, bool)
1. AvailabilitySetResourceSettings.AsVirtualMachineResourceSettings() (*VirtualMachineResourceSettings, bool)
1. AvailabilitySetResourceSettings.AsVirtualNetworkResourceSettings() (*VirtualNetworkResourceSettings, bool)
1. AvailabilitySetResourceSettings.MarshalJSON() ([]byte, error)
1. DiskEncryptionSetResourceSettings.AsAvailabilitySetResourceSettings() (*AvailabilitySetResourceSettings, bool)
1. DiskEncryptionSetResourceSettings.AsBasicResourceSettings() (BasicResourceSettings, bool)
1. DiskEncryptionSetResourceSettings.AsDiskEncryptionSetResourceSettings() (*DiskEncryptionSetResourceSettings, bool)
1. DiskEncryptionSetResourceSettings.AsKeyVaultResourceSettings() (*KeyVaultResourceSettings, bool)
1. DiskEncryptionSetResourceSettings.AsLoadBalancerResourceSettings() (*LoadBalancerResourceSettings, bool)
1. DiskEncryptionSetResourceSettings.AsNetworkInterfaceResourceSettings() (*NetworkInterfaceResourceSettings, bool)
1. DiskEncryptionSetResourceSettings.AsNetworkSecurityGroupResourceSettings() (*NetworkSecurityGroupResourceSettings, bool)
1. DiskEncryptionSetResourceSettings.AsPublicIPAddressResourceSettings() (*PublicIPAddressResourceSettings, bool)
1. DiskEncryptionSetResourceSettings.AsResourceGroupResourceSettings() (*ResourceGroupResourceSettings, bool)
1. DiskEncryptionSetResourceSettings.AsResourceSettings() (*ResourceSettings, bool)
1. DiskEncryptionSetResourceSettings.AsSQLDatabaseResourceSettings() (*SQLDatabaseResourceSettings, bool)
1. DiskEncryptionSetResourceSettings.AsSQLElasticPoolResourceSettings() (*SQLElasticPoolResourceSettings, bool)
1. DiskEncryptionSetResourceSettings.AsSQLServerResourceSettings() (*SQLServerResourceSettings, bool)
1. DiskEncryptionSetResourceSettings.AsVirtualMachineResourceSettings() (*VirtualMachineResourceSettings, bool)
1. DiskEncryptionSetResourceSettings.AsVirtualNetworkResourceSettings() (*VirtualNetworkResourceSettings, bool)
1. DiskEncryptionSetResourceSettings.MarshalJSON() ([]byte, error)
1. KeyVaultResourceSettings.AsAvailabilitySetResourceSettings() (*AvailabilitySetResourceSettings, bool)
1. KeyVaultResourceSettings.AsBasicResourceSettings() (BasicResourceSettings, bool)
1. KeyVaultResourceSettings.AsDiskEncryptionSetResourceSettings() (*DiskEncryptionSetResourceSettings, bool)
1. KeyVaultResourceSettings.AsKeyVaultResourceSettings() (*KeyVaultResourceSettings, bool)
1. KeyVaultResourceSettings.AsLoadBalancerResourceSettings() (*LoadBalancerResourceSettings, bool)
1. KeyVaultResourceSettings.AsNetworkInterfaceResourceSettings() (*NetworkInterfaceResourceSettings, bool)
1. KeyVaultResourceSettings.AsNetworkSecurityGroupResourceSettings() (*NetworkSecurityGroupResourceSettings, bool)
1. KeyVaultResourceSettings.AsPublicIPAddressResourceSettings() (*PublicIPAddressResourceSettings, bool)
1. KeyVaultResourceSettings.AsResourceGroupResourceSettings() (*ResourceGroupResourceSettings, bool)
1. KeyVaultResourceSettings.AsResourceSettings() (*ResourceSettings, bool)
1. KeyVaultResourceSettings.AsSQLDatabaseResourceSettings() (*SQLDatabaseResourceSettings, bool)
1. KeyVaultResourceSettings.AsSQLElasticPoolResourceSettings() (*SQLElasticPoolResourceSettings, bool)
1. KeyVaultResourceSettings.AsSQLServerResourceSettings() (*SQLServerResourceSettings, bool)
1. KeyVaultResourceSettings.AsVirtualMachineResourceSettings() (*VirtualMachineResourceSettings, bool)
1. KeyVaultResourceSettings.AsVirtualNetworkResourceSettings() (*VirtualNetworkResourceSettings, bool)
1. KeyVaultResourceSettings.MarshalJSON() ([]byte, error)
1. LoadBalancerResourceSettings.AsAvailabilitySetResourceSettings() (*AvailabilitySetResourceSettings, bool)
1. LoadBalancerResourceSettings.AsBasicResourceSettings() (BasicResourceSettings, bool)
1. LoadBalancerResourceSettings.AsDiskEncryptionSetResourceSettings() (*DiskEncryptionSetResourceSettings, bool)
1. LoadBalancerResourceSettings.AsKeyVaultResourceSettings() (*KeyVaultResourceSettings, bool)
1. LoadBalancerResourceSettings.AsLoadBalancerResourceSettings() (*LoadBalancerResourceSettings, bool)
1. LoadBalancerResourceSettings.AsNetworkInterfaceResourceSettings() (*NetworkInterfaceResourceSettings, bool)
1. LoadBalancerResourceSettings.AsNetworkSecurityGroupResourceSettings() (*NetworkSecurityGroupResourceSettings, bool)
1. LoadBalancerResourceSettings.AsPublicIPAddressResourceSettings() (*PublicIPAddressResourceSettings, bool)
1. LoadBalancerResourceSettings.AsResourceGroupResourceSettings() (*ResourceGroupResourceSettings, bool)
1. LoadBalancerResourceSettings.AsResourceSettings() (*ResourceSettings, bool)
1. LoadBalancerResourceSettings.AsSQLDatabaseResourceSettings() (*SQLDatabaseResourceSettings, bool)
1. LoadBalancerResourceSettings.AsSQLElasticPoolResourceSettings() (*SQLElasticPoolResourceSettings, bool)
1. LoadBalancerResourceSettings.AsSQLServerResourceSettings() (*SQLServerResourceSettings, bool)
1. LoadBalancerResourceSettings.AsVirtualMachineResourceSettings() (*VirtualMachineResourceSettings, bool)
1. LoadBalancerResourceSettings.AsVirtualNetworkResourceSettings() (*VirtualNetworkResourceSettings, bool)
1. LoadBalancerResourceSettings.MarshalJSON() ([]byte, error)
1. NetworkInterfaceResourceSettings.AsAvailabilitySetResourceSettings() (*AvailabilitySetResourceSettings, bool)
1. NetworkInterfaceResourceSettings.AsBasicResourceSettings() (BasicResourceSettings, bool)
1. NetworkInterfaceResourceSettings.AsDiskEncryptionSetResourceSettings() (*DiskEncryptionSetResourceSettings, bool)
1. NetworkInterfaceResourceSettings.AsKeyVaultResourceSettings() (*KeyVaultResourceSettings, bool)
1. NetworkInterfaceResourceSettings.AsLoadBalancerResourceSettings() (*LoadBalancerResourceSettings, bool)
1. NetworkInterfaceResourceSettings.AsNetworkInterfaceResourceSettings() (*NetworkInterfaceResourceSettings, bool)
1. NetworkInterfaceResourceSettings.AsNetworkSecurityGroupResourceSettings() (*NetworkSecurityGroupResourceSettings, bool)
1. NetworkInterfaceResourceSettings.AsPublicIPAddressResourceSettings() (*PublicIPAddressResourceSettings, bool)
1. NetworkInterfaceResourceSettings.AsResourceGroupResourceSettings() (*ResourceGroupResourceSettings, bool)
1. NetworkInterfaceResourceSettings.AsResourceSettings() (*ResourceSettings, bool)
1. NetworkInterfaceResourceSettings.AsSQLDatabaseResourceSettings() (*SQLDatabaseResourceSettings, bool)
1. NetworkInterfaceResourceSettings.AsSQLElasticPoolResourceSettings() (*SQLElasticPoolResourceSettings, bool)
1. NetworkInterfaceResourceSettings.AsSQLServerResourceSettings() (*SQLServerResourceSettings, bool)
1. NetworkInterfaceResourceSettings.AsVirtualMachineResourceSettings() (*VirtualMachineResourceSettings, bool)
1. NetworkInterfaceResourceSettings.AsVirtualNetworkResourceSettings() (*VirtualNetworkResourceSettings, bool)
1. NetworkInterfaceResourceSettings.MarshalJSON() ([]byte, error)
1. NetworkSecurityGroupResourceSettings.AsAvailabilitySetResourceSettings() (*AvailabilitySetResourceSettings, bool)
1. NetworkSecurityGroupResourceSettings.AsBasicResourceSettings() (BasicResourceSettings, bool)
1. NetworkSecurityGroupResourceSettings.AsDiskEncryptionSetResourceSettings() (*DiskEncryptionSetResourceSettings, bool)
1. NetworkSecurityGroupResourceSettings.AsKeyVaultResourceSettings() (*KeyVaultResourceSettings, bool)
1. NetworkSecurityGroupResourceSettings.AsLoadBalancerResourceSettings() (*LoadBalancerResourceSettings, bool)
1. NetworkSecurityGroupResourceSettings.AsNetworkInterfaceResourceSettings() (*NetworkInterfaceResourceSettings, bool)
1. NetworkSecurityGroupResourceSettings.AsNetworkSecurityGroupResourceSettings() (*NetworkSecurityGroupResourceSettings, bool)
1. NetworkSecurityGroupResourceSettings.AsPublicIPAddressResourceSettings() (*PublicIPAddressResourceSettings, bool)
1. NetworkSecurityGroupResourceSettings.AsResourceGroupResourceSettings() (*ResourceGroupResourceSettings, bool)
1. NetworkSecurityGroupResourceSettings.AsResourceSettings() (*ResourceSettings, bool)
1. NetworkSecurityGroupResourceSettings.AsSQLDatabaseResourceSettings() (*SQLDatabaseResourceSettings, bool)
1. NetworkSecurityGroupResourceSettings.AsSQLElasticPoolResourceSettings() (*SQLElasticPoolResourceSettings, bool)
1. NetworkSecurityGroupResourceSettings.AsSQLServerResourceSettings() (*SQLServerResourceSettings, bool)
1. NetworkSecurityGroupResourceSettings.AsVirtualMachineResourceSettings() (*VirtualMachineResourceSettings, bool)
1. NetworkSecurityGroupResourceSettings.AsVirtualNetworkResourceSettings() (*VirtualNetworkResourceSettings, bool)
1. NetworkSecurityGroupResourceSettings.MarshalJSON() ([]byte, error)
1. PossibleTargetAvailabilityZoneValues() []TargetAvailabilityZone
1. PossibleZoneRedundantValues() []ZoneRedundant
1. PublicIPAddressResourceSettings.AsAvailabilitySetResourceSettings() (*AvailabilitySetResourceSettings, bool)
1. PublicIPAddressResourceSettings.AsBasicResourceSettings() (BasicResourceSettings, bool)
1. PublicIPAddressResourceSettings.AsDiskEncryptionSetResourceSettings() (*DiskEncryptionSetResourceSettings, bool)
1. PublicIPAddressResourceSettings.AsKeyVaultResourceSettings() (*KeyVaultResourceSettings, bool)
1. PublicIPAddressResourceSettings.AsLoadBalancerResourceSettings() (*LoadBalancerResourceSettings, bool)
1. PublicIPAddressResourceSettings.AsNetworkInterfaceResourceSettings() (*NetworkInterfaceResourceSettings, bool)
1. PublicIPAddressResourceSettings.AsNetworkSecurityGroupResourceSettings() (*NetworkSecurityGroupResourceSettings, bool)
1. PublicIPAddressResourceSettings.AsPublicIPAddressResourceSettings() (*PublicIPAddressResourceSettings, bool)
1. PublicIPAddressResourceSettings.AsResourceGroupResourceSettings() (*ResourceGroupResourceSettings, bool)
1. PublicIPAddressResourceSettings.AsResourceSettings() (*ResourceSettings, bool)
1. PublicIPAddressResourceSettings.AsSQLDatabaseResourceSettings() (*SQLDatabaseResourceSettings, bool)
1. PublicIPAddressResourceSettings.AsSQLElasticPoolResourceSettings() (*SQLElasticPoolResourceSettings, bool)
1. PublicIPAddressResourceSettings.AsSQLServerResourceSettings() (*SQLServerResourceSettings, bool)
1. PublicIPAddressResourceSettings.AsVirtualMachineResourceSettings() (*VirtualMachineResourceSettings, bool)
1. PublicIPAddressResourceSettings.AsVirtualNetworkResourceSettings() (*VirtualNetworkResourceSettings, bool)
1. PublicIPAddressResourceSettings.MarshalJSON() ([]byte, error)
1. ResourceGroupResourceSettings.AsAvailabilitySetResourceSettings() (*AvailabilitySetResourceSettings, bool)
1. ResourceGroupResourceSettings.AsBasicResourceSettings() (BasicResourceSettings, bool)
1. ResourceGroupResourceSettings.AsDiskEncryptionSetResourceSettings() (*DiskEncryptionSetResourceSettings, bool)
1. ResourceGroupResourceSettings.AsKeyVaultResourceSettings() (*KeyVaultResourceSettings, bool)
1. ResourceGroupResourceSettings.AsLoadBalancerResourceSettings() (*LoadBalancerResourceSettings, bool)
1. ResourceGroupResourceSettings.AsNetworkInterfaceResourceSettings() (*NetworkInterfaceResourceSettings, bool)
1. ResourceGroupResourceSettings.AsNetworkSecurityGroupResourceSettings() (*NetworkSecurityGroupResourceSettings, bool)
1. ResourceGroupResourceSettings.AsPublicIPAddressResourceSettings() (*PublicIPAddressResourceSettings, bool)
1. ResourceGroupResourceSettings.AsResourceGroupResourceSettings() (*ResourceGroupResourceSettings, bool)
1. ResourceGroupResourceSettings.AsResourceSettings() (*ResourceSettings, bool)
1. ResourceGroupResourceSettings.AsSQLDatabaseResourceSettings() (*SQLDatabaseResourceSettings, bool)
1. ResourceGroupResourceSettings.AsSQLElasticPoolResourceSettings() (*SQLElasticPoolResourceSettings, bool)
1. ResourceGroupResourceSettings.AsSQLServerResourceSettings() (*SQLServerResourceSettings, bool)
1. ResourceGroupResourceSettings.AsVirtualMachineResourceSettings() (*VirtualMachineResourceSettings, bool)
1. ResourceGroupResourceSettings.AsVirtualNetworkResourceSettings() (*VirtualNetworkResourceSettings, bool)
1. ResourceGroupResourceSettings.MarshalJSON() ([]byte, error)
1. ResourceSettings.AsAvailabilitySetResourceSettings() (*AvailabilitySetResourceSettings, bool)
1. ResourceSettings.AsDiskEncryptionSetResourceSettings() (*DiskEncryptionSetResourceSettings, bool)
1. ResourceSettings.AsKeyVaultResourceSettings() (*KeyVaultResourceSettings, bool)
1. ResourceSettings.AsLoadBalancerResourceSettings() (*LoadBalancerResourceSettings, bool)
1. ResourceSettings.AsNetworkInterfaceResourceSettings() (*NetworkInterfaceResourceSettings, bool)
1. ResourceSettings.AsNetworkSecurityGroupResourceSettings() (*NetworkSecurityGroupResourceSettings, bool)
1. ResourceSettings.AsPublicIPAddressResourceSettings() (*PublicIPAddressResourceSettings, bool)
1. ResourceSettings.AsResourceGroupResourceSettings() (*ResourceGroupResourceSettings, bool)
1. ResourceSettings.AsSQLDatabaseResourceSettings() (*SQLDatabaseResourceSettings, bool)
1. ResourceSettings.AsSQLElasticPoolResourceSettings() (*SQLElasticPoolResourceSettings, bool)
1. ResourceSettings.AsSQLServerResourceSettings() (*SQLServerResourceSettings, bool)
1. ResourceSettings.AsVirtualMachineResourceSettings() (*VirtualMachineResourceSettings, bool)
1. ResourceSettings.AsVirtualNetworkResourceSettings() (*VirtualNetworkResourceSettings, bool)
1. SQLDatabaseResourceSettings.AsAvailabilitySetResourceSettings() (*AvailabilitySetResourceSettings, bool)
1. SQLDatabaseResourceSettings.AsBasicResourceSettings() (BasicResourceSettings, bool)
1. SQLDatabaseResourceSettings.AsDiskEncryptionSetResourceSettings() (*DiskEncryptionSetResourceSettings, bool)
1. SQLDatabaseResourceSettings.AsKeyVaultResourceSettings() (*KeyVaultResourceSettings, bool)
1. SQLDatabaseResourceSettings.AsLoadBalancerResourceSettings() (*LoadBalancerResourceSettings, bool)
1. SQLDatabaseResourceSettings.AsNetworkInterfaceResourceSettings() (*NetworkInterfaceResourceSettings, bool)
1. SQLDatabaseResourceSettings.AsNetworkSecurityGroupResourceSettings() (*NetworkSecurityGroupResourceSettings, bool)
1. SQLDatabaseResourceSettings.AsPublicIPAddressResourceSettings() (*PublicIPAddressResourceSettings, bool)
1. SQLDatabaseResourceSettings.AsResourceGroupResourceSettings() (*ResourceGroupResourceSettings, bool)
1. SQLDatabaseResourceSettings.AsResourceSettings() (*ResourceSettings, bool)
1. SQLDatabaseResourceSettings.AsSQLDatabaseResourceSettings() (*SQLDatabaseResourceSettings, bool)
1. SQLDatabaseResourceSettings.AsSQLElasticPoolResourceSettings() (*SQLElasticPoolResourceSettings, bool)
1. SQLDatabaseResourceSettings.AsSQLServerResourceSettings() (*SQLServerResourceSettings, bool)
1. SQLDatabaseResourceSettings.AsVirtualMachineResourceSettings() (*VirtualMachineResourceSettings, bool)
1. SQLDatabaseResourceSettings.AsVirtualNetworkResourceSettings() (*VirtualNetworkResourceSettings, bool)
1. SQLDatabaseResourceSettings.MarshalJSON() ([]byte, error)
1. SQLElasticPoolResourceSettings.AsAvailabilitySetResourceSettings() (*AvailabilitySetResourceSettings, bool)
1. SQLElasticPoolResourceSettings.AsBasicResourceSettings() (BasicResourceSettings, bool)
1. SQLElasticPoolResourceSettings.AsDiskEncryptionSetResourceSettings() (*DiskEncryptionSetResourceSettings, bool)
1. SQLElasticPoolResourceSettings.AsKeyVaultResourceSettings() (*KeyVaultResourceSettings, bool)
1. SQLElasticPoolResourceSettings.AsLoadBalancerResourceSettings() (*LoadBalancerResourceSettings, bool)
1. SQLElasticPoolResourceSettings.AsNetworkInterfaceResourceSettings() (*NetworkInterfaceResourceSettings, bool)
1. SQLElasticPoolResourceSettings.AsNetworkSecurityGroupResourceSettings() (*NetworkSecurityGroupResourceSettings, bool)
1. SQLElasticPoolResourceSettings.AsPublicIPAddressResourceSettings() (*PublicIPAddressResourceSettings, bool)
1. SQLElasticPoolResourceSettings.AsResourceGroupResourceSettings() (*ResourceGroupResourceSettings, bool)
1. SQLElasticPoolResourceSettings.AsResourceSettings() (*ResourceSettings, bool)
1. SQLElasticPoolResourceSettings.AsSQLDatabaseResourceSettings() (*SQLDatabaseResourceSettings, bool)
1. SQLElasticPoolResourceSettings.AsSQLElasticPoolResourceSettings() (*SQLElasticPoolResourceSettings, bool)
1. SQLElasticPoolResourceSettings.AsSQLServerResourceSettings() (*SQLServerResourceSettings, bool)
1. SQLElasticPoolResourceSettings.AsVirtualMachineResourceSettings() (*VirtualMachineResourceSettings, bool)
1. SQLElasticPoolResourceSettings.AsVirtualNetworkResourceSettings() (*VirtualNetworkResourceSettings, bool)
1. SQLElasticPoolResourceSettings.MarshalJSON() ([]byte, error)
1. SQLServerResourceSettings.AsAvailabilitySetResourceSettings() (*AvailabilitySetResourceSettings, bool)
1. SQLServerResourceSettings.AsBasicResourceSettings() (BasicResourceSettings, bool)
1. SQLServerResourceSettings.AsDiskEncryptionSetResourceSettings() (*DiskEncryptionSetResourceSettings, bool)
1. SQLServerResourceSettings.AsKeyVaultResourceSettings() (*KeyVaultResourceSettings, bool)
1. SQLServerResourceSettings.AsLoadBalancerResourceSettings() (*LoadBalancerResourceSettings, bool)
1. SQLServerResourceSettings.AsNetworkInterfaceResourceSettings() (*NetworkInterfaceResourceSettings, bool)
1. SQLServerResourceSettings.AsNetworkSecurityGroupResourceSettings() (*NetworkSecurityGroupResourceSettings, bool)
1. SQLServerResourceSettings.AsPublicIPAddressResourceSettings() (*PublicIPAddressResourceSettings, bool)
1. SQLServerResourceSettings.AsResourceGroupResourceSettings() (*ResourceGroupResourceSettings, bool)
1. SQLServerResourceSettings.AsResourceSettings() (*ResourceSettings, bool)
1. SQLServerResourceSettings.AsSQLDatabaseResourceSettings() (*SQLDatabaseResourceSettings, bool)
1. SQLServerResourceSettings.AsSQLElasticPoolResourceSettings() (*SQLElasticPoolResourceSettings, bool)
1. SQLServerResourceSettings.AsSQLServerResourceSettings() (*SQLServerResourceSettings, bool)
1. SQLServerResourceSettings.AsVirtualMachineResourceSettings() (*VirtualMachineResourceSettings, bool)
1. SQLServerResourceSettings.AsVirtualNetworkResourceSettings() (*VirtualNetworkResourceSettings, bool)
1. SQLServerResourceSettings.MarshalJSON() ([]byte, error)
1. VirtualMachineResourceSettings.AsAvailabilitySetResourceSettings() (*AvailabilitySetResourceSettings, bool)
1. VirtualMachineResourceSettings.AsBasicResourceSettings() (BasicResourceSettings, bool)
1. VirtualMachineResourceSettings.AsDiskEncryptionSetResourceSettings() (*DiskEncryptionSetResourceSettings, bool)
1. VirtualMachineResourceSettings.AsKeyVaultResourceSettings() (*KeyVaultResourceSettings, bool)
1. VirtualMachineResourceSettings.AsLoadBalancerResourceSettings() (*LoadBalancerResourceSettings, bool)
1. VirtualMachineResourceSettings.AsNetworkInterfaceResourceSettings() (*NetworkInterfaceResourceSettings, bool)
1. VirtualMachineResourceSettings.AsNetworkSecurityGroupResourceSettings() (*NetworkSecurityGroupResourceSettings, bool)
1. VirtualMachineResourceSettings.AsPublicIPAddressResourceSettings() (*PublicIPAddressResourceSettings, bool)
1. VirtualMachineResourceSettings.AsResourceGroupResourceSettings() (*ResourceGroupResourceSettings, bool)
1. VirtualMachineResourceSettings.AsResourceSettings() (*ResourceSettings, bool)
1. VirtualMachineResourceSettings.AsSQLDatabaseResourceSettings() (*SQLDatabaseResourceSettings, bool)
1. VirtualMachineResourceSettings.AsSQLElasticPoolResourceSettings() (*SQLElasticPoolResourceSettings, bool)
1. VirtualMachineResourceSettings.AsSQLServerResourceSettings() (*SQLServerResourceSettings, bool)
1. VirtualMachineResourceSettings.AsVirtualMachineResourceSettings() (*VirtualMachineResourceSettings, bool)
1. VirtualMachineResourceSettings.AsVirtualNetworkResourceSettings() (*VirtualNetworkResourceSettings, bool)
1. VirtualMachineResourceSettings.MarshalJSON() ([]byte, error)
1. VirtualNetworkResourceSettings.AsAvailabilitySetResourceSettings() (*AvailabilitySetResourceSettings, bool)
1. VirtualNetworkResourceSettings.AsBasicResourceSettings() (BasicResourceSettings, bool)
1. VirtualNetworkResourceSettings.AsDiskEncryptionSetResourceSettings() (*DiskEncryptionSetResourceSettings, bool)
1. VirtualNetworkResourceSettings.AsKeyVaultResourceSettings() (*KeyVaultResourceSettings, bool)
1. VirtualNetworkResourceSettings.AsLoadBalancerResourceSettings() (*LoadBalancerResourceSettings, bool)
1. VirtualNetworkResourceSettings.AsNetworkInterfaceResourceSettings() (*NetworkInterfaceResourceSettings, bool)
1. VirtualNetworkResourceSettings.AsNetworkSecurityGroupResourceSettings() (*NetworkSecurityGroupResourceSettings, bool)
1. VirtualNetworkResourceSettings.AsPublicIPAddressResourceSettings() (*PublicIPAddressResourceSettings, bool)
1. VirtualNetworkResourceSettings.AsResourceGroupResourceSettings() (*ResourceGroupResourceSettings, bool)
1. VirtualNetworkResourceSettings.AsResourceSettings() (*ResourceSettings, bool)
1. VirtualNetworkResourceSettings.AsSQLDatabaseResourceSettings() (*SQLDatabaseResourceSettings, bool)
1. VirtualNetworkResourceSettings.AsSQLElasticPoolResourceSettings() (*SQLElasticPoolResourceSettings, bool)
1. VirtualNetworkResourceSettings.AsSQLServerResourceSettings() (*SQLServerResourceSettings, bool)
1. VirtualNetworkResourceSettings.AsVirtualMachineResourceSettings() (*VirtualMachineResourceSettings, bool)
1. VirtualNetworkResourceSettings.AsVirtualNetworkResourceSettings() (*VirtualNetworkResourceSettings, bool)
1. VirtualNetworkResourceSettings.MarshalJSON() ([]byte, error)

### Struct Changes

#### Removed Structs

1. AvailabilitySetResourceSettings
1. AzureResourceReference
1. DiskEncryptionSetResourceSettings
1. KeyVaultResourceSettings
1. LBBackendAddressPoolResourceSettings
1. LBFrontendIPConfigurationResourceSettings
1. LoadBalancerBackendAddressPoolReference
1. LoadBalancerNatRuleReference
1. LoadBalancerResourceSettings
1. NetworkInterfaceResourceSettings
1. NetworkSecurityGroupResourceSettings
1. NicIPConfigurationResourceSettings
1. NsgReference
1. NsgSecurityRule
1. ProxyResourceReference
1. PublicIPAddressResourceSettings
1. PublicIPReference
1. ResourceGroupResourceSettings
1. SQLDatabaseResourceSettings
1. SQLElasticPoolResourceSettings
1. SQLServerResourceSettings
1. SubnetReference
1. SubnetResourceSettings
1. VirtualMachineResourceSettings
1. VirtualNetworkResourceSettings

#### Removed Struct Fields

1. MoveResourceProperties.SourceResourceSettings
1. ResourceSettings.TargetResourceName

## Additive Changes

### New Constants

1. MoveParadigm.CopyType
1. MoveParadigm.MigrateType

### New Funcs

1. NewSupportedResourceTypesForResourceMoverClient(string) SupportedResourceTypesForResourceMoverClient
1. NewSupportedResourceTypesForResourceMoverClientWithBaseURI(string, string) SupportedResourceTypesForResourceMoverClient
1. PossibleMoveParadigmValues() []MoveParadigm
1. ResourceTypeDefinitionForResourceMover.MarshalJSON() ([]byte, error)
1. SupportedResourceTypesForResourceMoverClient.Get(context.Context) (SupportedResourceTypesForResourceMover, error)
1. SupportedResourceTypesForResourceMoverClient.GetPreparer(context.Context) (*http.Request, error)
1. SupportedResourceTypesForResourceMoverClient.GetResponder(*http.Response) (SupportedResourceTypesForResourceMover, error)
1. SupportedResourceTypesForResourceMoverClient.GetSender(*http.Request) (*http.Response, error)

### Struct Changes

#### New Structs

1. ResourceTypeDefinitionForResourceMover
1. SupportedResourceTypesForResourceMover
1. SupportedResourceTypesForResourceMoverClient

#### New Struct Fields

1. BulkRemoveRequest.AddDependencyAutomatically
1. CommitRequest.AddDependencyAutomatically
1. DiscardRequest.AddDependencyAutomatically
1. MoveResourceProperties.TargetResourceName
1. PrepareRequest.AddDependencyAutomatically
1. ResourceMoveRequestType.AddDependencyAutomatically
1. ResourceSettings.SerializedResourceSettings
