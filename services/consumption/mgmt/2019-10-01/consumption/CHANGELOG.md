# Unreleased

## Breaking Changes

### Struct Changes

#### Removed Struct Fields

1. LegacyReservationRecommendation.Etag
1. ModernReservationRecommendation.Etag
1. ReservationDetail.Etag
1. ReservationRecommendation.Etag
1. ReservationRecommendationDetailsModel.ETag
1. ReservationSummary.Etag

### Signature Changes

#### Struct Fields

1. ModernReservationRecommendationProperties.Scope changed type from *string to *int32

## Additive Changes

### New Funcs

1. RevisedResource.MarshalJSON() ([]byte, error)

### Struct Changes

#### New Structs

1. RevisedResource

#### New Struct Fields

1. LegacyReservationRecommendation.ETag
1. ReservationDetail.ETag
1. ReservationRecommendation.ETag
1. ReservationSummary.ETag
