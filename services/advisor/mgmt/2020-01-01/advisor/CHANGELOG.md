# Unreleased Content

## Breaking Changes

### Struct Changes

#### Removed Struct Fields

1. SuppressionContract.autorest.Response

### Signature Changes

#### Funcs

1. SuppressionsClient.Create
	- Returns
		- From: SuppressionContract, error
		- To: SetObject, error
1. SuppressionsClient.CreateResponder
	- Returns
		- From: SuppressionContract, error
		- To: SetObject, error

## Additive Changes

### New Funcs

1. ResourceMetadata.MarshalJSON() ([]byte, error)

### Struct Changes

#### New Struct Fields

1. RecommendationProperties.Actions
1. RecommendationProperties.Description
1. RecommendationProperties.ExposedMetadataProperties
1. RecommendationProperties.Label
1. RecommendationProperties.LearnMoreLink
1. RecommendationProperties.PotentialBenefits
1. RecommendationProperties.Remediation
1. ResourceMetadata.Action
1. ResourceMetadata.Plural
1. ResourceMetadata.Singular
