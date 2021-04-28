# Unreleased Content

## Breaking Changes

### Removed Constants

1. SkuTier.Free

### Removed Funcs

1. DiskPoolProperties.MarshalJSON() ([]byte, error)
1. IscsiTargetProperties.MarshalJSON() ([]byte, error)
1. PossibleSkuTierValues() []SkuTier
1. TargetPortalGroup.MarshalJSON() ([]byte, error)

### Struct Changes

#### Removed Structs

1. Sku

#### Removed Struct Fields

1. DiskPool.Sku

### Signature Changes

#### Const Types

1. Basic changed type from SkuTier to DiskPoolTier
1. Premium changed type from SkuTier to DiskPoolTier
1. Standard changed type from SkuTier to DiskPoolTier

#### Funcs

1. DiskPoolsClient.CreateOrUpdate
	- Params
		- From: context.Context, string, string, DiskPool
		- To: context.Context, string, string, DiskPoolCreate
1. DiskPoolsClient.CreateOrUpdatePreparer
	- Params
		- From: context.Context, string, string, DiskPool
		- To: context.Context, string, string, DiskPoolCreate
1. DiskPoolsClient.Update
	- Params
		- From: context.Context, string, string, DiskPool
		- To: context.Context, string, string, DiskPoolUpdate
	- Returns
		- From: DiskPool, error
		- To: DiskPoolsUpdateFuture, error
1. DiskPoolsClient.UpdatePreparer
	- Params
		- From: context.Context, string, string, DiskPool
		- To: context.Context, string, string, DiskPoolUpdate
1. DiskPoolsClient.UpdateSender
	- Returns
		- From: *http.Response, error
		- To: DiskPoolsUpdateFuture, error
1. IscsiTargetsClient.CreateOrUpdate
	- Params
		- From: context.Context, string, string, string, IscsiTarget
		- To: context.Context, string, string, string, IscsiTargetCreate
1. IscsiTargetsClient.CreateOrUpdatePreparer
	- Params
		- From: context.Context, string, string, string, IscsiTarget
		- To: context.Context, string, string, string, IscsiTargetCreate

## Additive Changes

### New Constants

1. OperationalStatus.Running
1. OperationalStatus.Stopped
1. OperationalStatus.Stoppeddeallocated
1. OperationalStatus.Updating

### New Funcs

1. *DiskPoolCreate.UnmarshalJSON([]byte) error
1. *DiskPoolUpdate.UnmarshalJSON([]byte) error
1. *DiskPoolsUpdateFuture.UnmarshalJSON([]byte) error
1. *IscsiTargetCreate.UnmarshalJSON([]byte) error
1. *IscsiTargetUpdate.UnmarshalJSON([]byte) error
1. *IscsiTargetsUpdateFuture.UnmarshalJSON([]byte) error
1. DiskPoolCreate.MarshalJSON() ([]byte, error)
1. DiskPoolUpdate.MarshalJSON() ([]byte, error)
1. IscsiTargetCreate.MarshalJSON() ([]byte, error)
1. IscsiTargetUpdate.MarshalJSON() ([]byte, error)
1. IscsiTargetsClient.Update(context.Context, string, string, string, IscsiTargetUpdate) (IscsiTargetsUpdateFuture, error)
1. IscsiTargetsClient.UpdatePreparer(context.Context, string, string, string, IscsiTargetUpdate) (*http.Request, error)
1. IscsiTargetsClient.UpdateResponder(*http.Response) (IscsiTarget, error)
1. IscsiTargetsClient.UpdateSender(*http.Request) (IscsiTargetsUpdateFuture, error)
1. PossibleDiskPoolTierValues() []DiskPoolTier

### Struct Changes

#### New Structs

1. DiskPoolCreate
1. DiskPoolCreateProperties
1. DiskPoolUpdate
1. DiskPoolUpdateProperties
1. DiskPoolsUpdateFuture
1. IscsiTargetCreate
1. IscsiTargetCreateProperties
1. IscsiTargetUpdate
1. IscsiTargetUpdateProperties
1. IscsiTargetsUpdateFuture
1. TargetPortalGroupCreate
1. TargetPortalGroupUpdate

#### New Struct Fields

1. DiskPoolProperties.AdditionalCapabilities
1. DiskPoolProperties.Tier
