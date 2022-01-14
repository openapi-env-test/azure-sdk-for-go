//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import "encoding/json"

func unmarshalChargeSummaryClassification(rawMsg json.RawMessage) (ChargeSummaryClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ChargeSummaryClassification
	switch m["kind"] {
	case string(ChargeSummaryKindLegacy):
		b = &LegacyChargeSummary{}
	case string(ChargeSummaryKindModern):
		b = &ModernChargeSummary{}
	default:
		b = &ChargeSummary{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalChargeSummaryClassificationArray(rawMsg json.RawMessage) ([]ChargeSummaryClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ChargeSummaryClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalChargeSummaryClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalLegacyReservationRecommendationPropertiesClassification(rawMsg json.RawMessage) (LegacyReservationRecommendationPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b LegacyReservationRecommendationPropertiesClassification
	switch m["scope"] {
	case "Shared":
		b = &LegacySharedScopeReservationRecommendationProperties{}
	case "Single":
		b = &LegacySingleScopeReservationRecommendationProperties{}
	default:
		b = &LegacyReservationRecommendationProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalLegacyReservationRecommendationPropertiesClassificationArray(rawMsg json.RawMessage) ([]LegacyReservationRecommendationPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]LegacyReservationRecommendationPropertiesClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalLegacyReservationRecommendationPropertiesClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalReservationRecommendationClassification(rawMsg json.RawMessage) (ReservationRecommendationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ReservationRecommendationClassification
	switch m["kind"] {
	case string(ReservationRecommendationKindLegacy):
		b = &LegacyReservationRecommendation{}
	case string(ReservationRecommendationKindModern):
		b = &ModernReservationRecommendation{}
	default:
		b = &ReservationRecommendation{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalReservationRecommendationClassificationArray(rawMsg json.RawMessage) ([]ReservationRecommendationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ReservationRecommendationClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalReservationRecommendationClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalUsageDetailClassification(rawMsg json.RawMessage) (UsageDetailClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b UsageDetailClassification
	switch m["kind"] {
	case string(UsageDetailsKindLegacy):
		b = &LegacyUsageDetail{}
	case string(UsageDetailsKindModern):
		b = &ModernUsageDetail{}
	default:
		b = &UsageDetail{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalUsageDetailClassificationArray(rawMsg json.RawMessage) ([]UsageDetailClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]UsageDetailClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalUsageDetailClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}
