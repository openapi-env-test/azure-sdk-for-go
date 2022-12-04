//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhealthcareapis

import "encoding/json"

func unmarshalAnalyticsConnectorDataDestinationClassification(rawMsg json.RawMessage) (AnalyticsConnectorDataDestinationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AnalyticsConnectorDataDestinationClassification
	switch m["type"] {
	case string(AnalyticsConnectorDataDestinationTypeDatalake):
		b = &AnalyticsConnectorDataLakeDataDestination{}
	default:
		b = &AnalyticsConnectorDataDestination{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAnalyticsConnectorDataSourceClassification(rawMsg json.RawMessage) (AnalyticsConnectorDataSourceClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AnalyticsConnectorDataSourceClassification
	switch m["type"] {
	case string(AnalyticsConnectorDataSourceTypeFhirservice):
		b = &AnalyticsConnectorFhirServiceDataSource{}
	default:
		b = &AnalyticsConnectorDataSource{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAnalyticsConnectorMappingClassification(rawMsg json.RawMessage) (AnalyticsConnectorMappingClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AnalyticsConnectorMappingClassification
	switch m["type"] {
	case string(AnalyticsConnectorMappingTypeFhirToParquet):
		b = &AnalyticsConnectorFhirToParquetMapping{}
	default:
		b = &AnalyticsConnectorMapping{}
	}
	return b, json.Unmarshal(rawMsg, b)
}
