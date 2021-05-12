# Unreleased Content

## Breaking Changes

### Struct Changes

#### Removed Struct Fields

1. MetricAlertResourcePatch.*MetricAlertProperties

### Signature Changes

#### Struct Fields

1. DynamicMetricCriteria.TimeAggregation changed type from interface{} to AggregationTypeEnum
1. MetricCriteria.TimeAggregation changed type from interface{} to AggregationTypeEnum
1. MultiMetricCriteria.TimeAggregation changed type from interface{} to AggregationTypeEnum

## Additive Changes

### New Constants

1. AggregationTypeEnum.AggregationTypeEnumAverage
1. AggregationTypeEnum.AggregationTypeEnumCount
1. AggregationTypeEnum.AggregationTypeEnumMaximum
1. AggregationTypeEnum.AggregationTypeEnumMinimum
1. AggregationTypeEnum.AggregationTypeEnumTotal
1. ConditionalOperator.ConditionalOperatorGreaterThanOrEqual
1. ConditionalOperator.ConditionalOperatorLessThanOrEqual
1. NamespaceClassification.Custom
1. NamespaceClassification.Platform
1. NamespaceClassification.Qos

### New Funcs

1. *MetricAlertPropertiesPatch.UnmarshalJSON([]byte) error
1. MetricAlertPropertiesPatch.MarshalJSON() ([]byte, error)
1. PossibleAggregationTypeEnumValues() []AggregationTypeEnum
1. PossibleNamespaceClassificationValues() []NamespaceClassification

### Struct Changes

#### New Structs

1. ErrorContract
1. MetricAlertPropertiesPatch

#### New Struct Fields

1. ActionGroupResource.Etag
1. ActionGroupResource.Kind
1. ActivityLogAlertResource.Etag
1. ActivityLogAlertResource.Kind
1. AlertRule.Action
1. AlertRule.ProvisioningState
1. AlertRuleResource.Etag
1. AlertRuleResource.Kind
1. AutoscaleSettingResource.Etag
1. AutoscaleSettingResource.Kind
1. Baseline.Timestamps
1. LogProfileResource.Etag
1. LogProfileResource.Kind
1. LogSearchRule.AutoMitigate
1. LogSearchRule.CreatedWithAPIVersion
1. LogSearchRule.DisplayName
1. LogSearchRule.IsLegacyLogAnalyticsRule
1. LogSearchRuleResource.Etag
1. LogSearchRuleResource.Kind
1. Metric.DisplayDescription
1. Metric.ErrorCode
1. MetricAlertProperties.IsMigrated
1. MetricAlertResource.Etag
1. MetricAlertResource.Kind
1. MetricAlertResourcePatch.*MetricAlertPropertiesPatch
1. MetricDefinition.Category
1. MetricDefinition.DisplayDescription
1. MetricNamespace.Classification
1. MetricTrigger.DividePerInstance
1. Resource.Etag
1. Resource.Kind
1. RuleDataSource.LegacyResourceID
1. RuleDataSource.MetricNamespace
1. RuleDataSource.ResourceLocation
1. RuleManagementEventDataSource.LegacyResourceID
1. RuleManagementEventDataSource.MetricNamespace
1. RuleManagementEventDataSource.ResourceLocation
1. RuleMetricDataSource.LegacyResourceID
1. RuleMetricDataSource.MetricNamespace
1. RuleMetricDataSource.ResourceLocation
