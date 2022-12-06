# Release History

## 2.2.0-beta.1 (2022-12-06)
### Features Added

- New const `ContainerGroupPriorityRegular`
- New const `ContainerGroupPrioritySpot`
- New const `ContainerGroupSKUConfidential`
- New type alias `ContainerGroupPriority`
- New function `PossibleContainerGroupPriorityValues() []ContainerGroupPriority`
- New struct `ConfidentialComputeProperties`
- New field `ConfidentialComputeProperties` in struct `ContainerGroupPropertiesProperties`
- New field `Priority` in struct `ContainerGroupPropertiesProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).