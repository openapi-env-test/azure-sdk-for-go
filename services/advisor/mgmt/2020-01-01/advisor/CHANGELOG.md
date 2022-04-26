# Unreleased

## Breaking Changes

### Removed Funcs

1. RecommendationMetadataClient.Get(context.Context, string) (SetObject, error)
1. RecommendationMetadataClient.GetPreparer(context.Context, string) (*http.Request, error)
1. RecommendationMetadataClient.GetResponder(*http.Response) (SetObject, error)
1. RecommendationMetadataClient.GetSender(*http.Request) (*http.Response, error)

### Struct Changes

#### Removed Structs

1. SetObject

### Signature Changes

#### Funcs

1. SuppressionsClient.Get
	- Returns
		- From: SetObject, error
		- To: SuppressionContract, error
1. SuppressionsClient.GetResponder
	- Returns
		- From: SetObject, error
		- To: SuppressionContract, error

## Additive Changes

### New Funcs

1. NewRecommendationClient(string) RecommendationClient
1. NewRecommendationClientWithBaseURI(string, string) RecommendationClient
1. RecommendationClient.Get(context.Context, string) (MetadataEntity, error)
1. RecommendationClient.GetPreparer(context.Context, string) (*http.Request, error)
1. RecommendationClient.GetResponder(*http.Response) (MetadataEntity, error)
1. RecommendationClient.GetSender(*http.Request) (*http.Response, error)
1. ResourceMetadata.MarshalJSON() ([]byte, error)

### Struct Changes

#### New Structs

1. RecommendationClient

#### New Struct Fields

1. MetadataEntity.autorest.Response
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
