# Unreleased

## Breaking Changes

### Struct Changes

#### Removed Struct Fields

1. ChargeSummary.Etag
1. ChargeSummary.Tags
1. CreditSummary.Etag
1. CreditSummary.Tags
1. EventSummary.Etag
1. EventSummary.Tags
1. LegacyChargeSummary.Etag
1. LegacyChargeSummary.Tags
1. LegacyReservationRecommendation.Etag
1. LotSummary.Etag
1. LotSummary.Tags
1. ModernChargeSummary.Etag
1. ModernChargeSummary.Tags
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

1. ChargeSummary.ETag
1. CreditSummary.ETag
1. EventSummary.ETag
1. LegacyChargeSummary.ETag
1. LegacyReservationRecommendation.ETag
1. LotSummary.ETag
1. ModernChargeSummary.ETag
1. ReservationDetail.ETag
1. ReservationRecommendation.ETag
1. ReservationSummary.ETag
