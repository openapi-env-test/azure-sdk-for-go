# Unreleased Content

## Breaking Changes

### Removed Funcs

1. MetricNamespacesClient.List(context.Context, string, string) (MetricNamespaceCollection, error)
1. MetricNamespacesClient.ListPreparer(context.Context, string, string) (*http.Request, error)
1. MetricNamespacesClient.ListResponder(*http.Response) (MetricNamespaceCollection, error)
1. MetricNamespacesClient.ListSender(*http.Request) (*http.Response, error)
1. NewMetricNamespacesClient(string) MetricNamespacesClient
1. NewMetricNamespacesClientWithBaseURI(string, string) MetricNamespacesClient

### Struct Changes

#### Removed Structs

1. MetricNamespace
1. MetricNamespaceCollection
1. MetricNamespaceName
1. MetricNamespacesClient

#### Removed Struct Fields

1. MetricAlertResourcePatch.*MetricAlertProperties

### Signature Changes

#### Const Types

1. High changed type from DynamicThresholdSensitivity to BaselineSensitivity
1. Low changed type from DynamicThresholdSensitivity to BaselineSensitivity
1. Medium changed type from DynamicThresholdSensitivity to BaselineSensitivity

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
1. DynamicThresholdSensitivity.DynamicThresholdSensitivityHigh
1. DynamicThresholdSensitivity.DynamicThresholdSensitivityLow
1. DynamicThresholdSensitivity.DynamicThresholdSensitivityMedium

### New Funcs

1. *MetricAlertPropertiesPatch.UnmarshalJSON([]byte) error
1. BaselinesClient.List(context.Context, string, string, string, string, *string, string, string, string, ResultType) (MetricBaselinesResponse, error)
1. BaselinesClient.ListPreparer(context.Context, string, string, string, string, *string, string, string, string, ResultType) (*http.Request, error)
1. BaselinesClient.ListResponder(*http.Response) (MetricBaselinesResponse, error)
1. BaselinesClient.ListSender(*http.Request) (*http.Response, error)
1. MetricAlertPropertiesPatch.MarshalJSON() ([]byte, error)
1. NewBaselinesClient(string) BaselinesClient
1. NewBaselinesClientWithBaseURI(string, string) BaselinesClient
1. PossibleAggregationTypeEnumValues() []AggregationTypeEnum
1. PossibleBaselineSensitivityValues() []BaselineSensitivity

### Struct Changes

#### New Structs

1. BaselineMetadata
1. BaselinesClient
1. ErrorContract
1. MetricAlertPropertiesPatch
1. MetricBaselinesResponse
1. MetricSingleDimension
1. SingleBaseline
1. SingleMetricBaseline
1. TimeSeriesBaseline

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
