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

// InlineResponse20028 struct for InlineResponse20028
type InlineResponse20028 struct {
	HasMore *bool                `json:"hasMore,omitempty"`
	Data    []ApplicationSession `json:"data"`
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *InlineResponse20028) GetHasMore() bool {
	if o == nil || o.HasMore == nil {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetHasMoreOk() (bool, bool) {
	if o == nil || o.HasMore == nil {
		var ret bool
		return ret, false
	}
	return *o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *InlineResponse20028) HasHasMore() bool {
	if o != nil && o.HasMore != nil {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *InlineResponse20028) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetData returns the Data field value
func (o *InlineResponse20028) GetData() []ApplicationSession {
	if o == nil {
		var ret []ApplicationSession
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *InlineResponse20028) SetData(v []ApplicationSession) {
	o.Data = v
}

type NullableInlineResponse20028 struct {
	Value        InlineResponse20028
	ExplicitNull bool
}

func (v NullableInlineResponse20028) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableInlineResponse20028) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
