/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you access the Campaign Manager at `https://yourbaseurl.talon.one/`, the URL for the [updateCustomerSessionV2](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint is `https://yourbaseurl.talon.one/v2/customer_sessions/{Id}`
 *
 * API version:
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// AnalyticsDataPointWithTrendAndInfluencedRate struct for AnalyticsDataPointWithTrendAndInfluencedRate
type AnalyticsDataPointWithTrendAndInfluencedRate struct {
	Value          float32 `json:"value"`
	InfluencedRate float32 `json:"influencedRate"`
	Trend          float32 `json:"trend"`
}

// GetValue returns the Value field value
func (o *AnalyticsDataPointWithTrendAndInfluencedRate) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *AnalyticsDataPointWithTrendAndInfluencedRate) SetValue(v float32) {
	o.Value = v
}

// GetInfluencedRate returns the InfluencedRate field value
func (o *AnalyticsDataPointWithTrendAndInfluencedRate) GetInfluencedRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.InfluencedRate
}

// SetInfluencedRate sets field value
func (o *AnalyticsDataPointWithTrendAndInfluencedRate) SetInfluencedRate(v float32) {
	o.InfluencedRate = v
}

// GetTrend returns the Trend field value
func (o *AnalyticsDataPointWithTrendAndInfluencedRate) GetTrend() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Trend
}

// SetTrend sets field value
func (o *AnalyticsDataPointWithTrendAndInfluencedRate) SetTrend(v float32) {
	o.Trend = v
}

type NullableAnalyticsDataPointWithTrendAndInfluencedRate struct {
	Value        AnalyticsDataPointWithTrendAndInfluencedRate
	ExplicitNull bool
}

func (v NullableAnalyticsDataPointWithTrendAndInfluencedRate) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAnalyticsDataPointWithTrendAndInfluencedRate) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}