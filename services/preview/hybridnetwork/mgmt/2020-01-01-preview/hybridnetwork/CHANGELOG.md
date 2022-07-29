# Unreleased

## Breaking Changes

### Struct Changes

#### Removed Struct Fields

1. DevicePropertiesFormat.AzureStackEdge

## Additive Changes

### New Constants

1. DeviceType.DeviceTypeAzureStackEdge

### New Funcs

1. AzureStackEdgeFormat.AsAzureStackEdgeFormat() (*AzureStackEdgeFormat, bool)
1. AzureStackEdgeFormat.AsBasicDevicePropertiesFormat() (BasicDevicePropertiesFormat, bool)
1. AzureStackEdgeFormat.AsDevicePropertiesFormat() (*DevicePropertiesFormat, bool)
1. AzureStackEdgeFormat.MarshalJSON() ([]byte, error)
1. DevicePropertiesFormat.AsAzureStackEdgeFormat() (*AzureStackEdgeFormat, bool)

### Struct Changes

#### New Structs

1. AzureStackEdgeFormat
