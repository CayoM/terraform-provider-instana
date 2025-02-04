package restapi_test

import (
	"encoding/json"
	. "github.com/gessnerfl/terraform-provider-instana/instana/restapi"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestShouldSuccessfullyUnmarshalWebsiteAlertConfig(t *testing.T) {
	thresholdValue := 5.0
	thresholdLastUpdate := int64(0)
	timeWindow := int64(600000)
	p90Aggregation := Percentile90Aggregation
	websiteAlertConfig := WebsiteAlertConfig{
		ID:              "website-alert-config-id",
		AlertChannelIDs: []string{"channel-2", "channel-1"},
		WebsiteID:       "website-id",
		Name:            "website-alert-config-name",
		Description:     "website-alert-config-description",
		Granularity:     Granularity600000,
		CustomerPayloadFields: []CustomPayloadField[StaticStringCustomPayloadFieldValue]{
			{
				Type:  StaticCustomPayloadType,
				Key:   "static-key",
				Value: StaticStringCustomPayloadFieldValue("static-value"),
			},
		},
		Rule: WebsiteAlertRule{
			AlertType:   "slowness",
			Aggregation: &p90Aggregation,
			MetricName:  "onLoadTime",
		},
		Severity:            SeverityCritical.GetAPIRepresentation(),
		TagFilterExpression: NewStringTagFilter(TagFilterEntitySource, "beacon.geo.country", EqualsOperator, "DE"),
		Threshold: Threshold{
			Type:        "staticThreshold",
			Operator:    ThresholdOperatorGreaterThan,
			LastUpdated: &thresholdLastUpdate,
			Value:       &thresholdValue,
		},
		TimeThreshold: WebsiteTimeThreshold{
			Type:       "violationsInSequence",
			TimeWindow: &timeWindow,
		},
		Triggering: true,
	}

	serializedJSON, _ := json.Marshal(websiteAlertConfig)

	result, err := NewWebsiteAlertConfigUnmarshaller().Unmarshal(serializedJSON)

	require.NoError(t, err)
	require.Equal(t, &websiteAlertConfig, result)
}

func TestShouldFailToUnmarshalWebsiteAlertConfigWhenResponseIsAJsonArray(t *testing.T) {
	response := `["foo","bar"]`

	_, err := NewWebsiteAlertConfigUnmarshaller().Unmarshal([]byte(response))

	require.Error(t, err)
}

func TestShouldReturnEmptyWebsiteAlertConfigWhenNoFieldOfResponseMatchesToModel(t *testing.T) {
	response := `{"foo" : "bar"}`
	config, err := NewWebsiteAlertConfigUnmarshaller().Unmarshal([]byte(response))

	require.NoError(t, err)
	require.Equal(t, &WebsiteAlertConfig{}, config)
}

func TestShouldFailToUnmarshalWebsiteAlertConfigWhenResponseIsNotAValidJson(t *testing.T) {
	response := `Invalid Data`

	_, err := NewWebsiteAlertConfigUnmarshaller().Unmarshal([]byte(response))

	require.Error(t, err)
}

func TestShouldFailToUnmarshalWebsiteAlertConfigWhenTagFilterIsNotValid(t *testing.T) {
	response := `
{
    "id": "12345",
    "name": "name",
    "description": "test-alert-description",
    "websiteId": "website-id",
    "severity": 5,
    "triggering": false,
    "tagFilterExpression": [ "foo", "bar"],
    "rule": {
      "alertType": "slowness",
      "aggregation": "P90",
      "metricName": "latency"
    },
    "threshold": {
      "type": "staticThreshold",
      "operator": ">=",
      "value": 5.0,
      "lastUpdated": 0
    },
    "alertChannelIds": [ "alert-channel-id-1", "alert-channel-id-2" ],
    "granularity": 600000,
    "timeThreshold": {
      "type": "violationsInSequence",
      "timeWindow": 600000
    },
    "customPayloadFields": [
		{
			"type": "staticString",
			"key": "test",
			"value": "test123"
      	}
	]
  }
`
	_, err := NewWebsiteAlertConfigUnmarshaller().Unmarshal([]byte(response))

	require.Error(t, err)
}
