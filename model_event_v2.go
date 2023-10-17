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

// EventV2
type EventV2 struct {
	// ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known `profileId`, we recommend you use a guest `profileId`.
	ProfileId *string `json:"profileId,omitempty"`
	// When using the `dry` query parameter, use this property to list the campaign to be evaluated by the Rule Engine.  These campaigns will be evaluated, even if they are disabled, allowing you to test specific campaigns before activating them.
	EvaluableCampaignIds *[]int32 `json:"evaluableCampaignIds,omitempty"`
	// A string representing the event name. Must not be a reserved event name. You create this value when you [create an attribute](https://docs.talon.one/docs/dev/concepts/entities/events#creating-a-custom-event) of type `event` in the Campaign Manager.
	Type string `json:"type"`
	// Arbitrary additional JSON properties associated with the event. They must be created in the Campaign Manager before setting them with this property. See [creating custom attributes](https://docs.talon.one/docs/product/account/dev-tools/managing-attributes#creating-a-custom-attribute).
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *EventV2) GetProfileId() string {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EventV2) GetProfileIdOk() (string, bool) {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret, false
	}
	return *o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *EventV2) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *EventV2) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetEvaluableCampaignIds returns the EvaluableCampaignIds field value if set, zero value otherwise.
func (o *EventV2) GetEvaluableCampaignIds() []int32 {
	if o == nil || o.EvaluableCampaignIds == nil {
		var ret []int32
		return ret
	}
	return *o.EvaluableCampaignIds
}

// GetEvaluableCampaignIdsOk returns a tuple with the EvaluableCampaignIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EventV2) GetEvaluableCampaignIdsOk() ([]int32, bool) {
	if o == nil || o.EvaluableCampaignIds == nil {
		var ret []int32
		return ret, false
	}
	return *o.EvaluableCampaignIds, true
}

// HasEvaluableCampaignIds returns a boolean if a field has been set.
func (o *EventV2) HasEvaluableCampaignIds() bool {
	if o != nil && o.EvaluableCampaignIds != nil {
		return true
	}

	return false
}

// SetEvaluableCampaignIds gets a reference to the given []int32 and assigns it to the EvaluableCampaignIds field.
func (o *EventV2) SetEvaluableCampaignIds(v []int32) {
	o.EvaluableCampaignIds = &v
}

// GetType returns the Type field value
func (o *EventV2) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *EventV2) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EventV2) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EventV2) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EventV2) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *EventV2) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

type NullableEventV2 struct {
	Value        EventV2
	ExplicitNull bool
}

func (v NullableEventV2) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableEventV2) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
