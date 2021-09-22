# Unreleased

## Breaking Changes

### Struct Changes

#### Removed Structs

1. SetObject

### Signature Changes

#### Funcs

1. RecommendationMetadataClient.Get
	- Returns
		- From: SetObject, error
		- To: MetadataEntity, error
1. RecommendationMetadataClient.GetResponder
	- Returns
		- From: SetObject, error
		- To: MetadataEntity, error

## Additive Changes

### Struct Changes

#### New Struct Fields

1. MetadataEntity.autorest.Response
