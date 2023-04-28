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
	"time"
)

// AddLoyaltyPoints Points to add.
type AddLoyaltyPoints struct {
	// Amount of loyalty points.
	Points float32 `json:"points"`
	// Name / reason for the point addition.
	Name *string `json:"name,omitempty"`
	// The time format is either: - `immediate` or, - an **integer** followed by one letter indicating the time unit.  Examples: `immediate`, `30s`, `40m`, `1h`, `5D`, `7W`, `10M`.  Available units:  - `s`: seconds - `m`: minutes - `h`: hours - `D`: days - `W`: weeks - `M`: months  You can round certain units up or down: - `_D` for rounding down days only. Signifies the start of the day. - `_U` for rounding up days, weeks and months. Signifies the end of the day, week, or month.  If passed, `validUntil` should be omitted.
	ValidityDuration *string `json:"validityDuration,omitempty"`
	// Date and time when points should expire. The value should be provided in RFC 3339 format. If passed, `validityDuration` should be omitted.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// The amount of time before the points are considered valid.  The time format is either: - `immediate` or, - an **integer** followed by one letter indicating the time unit.  Examples: `immediate`, `30s`, `40m`, `1h`, `5D`, `7W`, `10M`.  Available units:  - `s`: seconds - `m`: minutes - `h`: hours - `D`: days - `W`: weeks - `M`: months  You can round certain units up or down: - `_D` for rounding down days only. Signifies the start of the day. - `_U` for rounding up days, weeks and months. Signifies the end of the day, week, or month.
	PendingDuration *string `json:"pendingDuration,omitempty"`
	// Date and time after the points are considered valid. The value should be provided in RFC 3339 format. If passed, `pendingDuration` should be omitted.
	PendingUntil *time.Time `json:"pendingUntil,omitempty"`
	// ID of the subledger the points are added to. If there is no existing subledger with this ID, the subledger is created automatically.
	SubledgerId *string `json:"subledgerId,omitempty"`
	// ID of the Application that is connected to the loyalty program. It is displayed in your Talon.One deployment URL.
	ApplicationId *int32 `json:"applicationId,omitempty"`
}

// GetPoints returns the Points field value
func (o *AddLoyaltyPoints) GetPoints() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Points
}

// SetPoints sets field value
func (o *AddLoyaltyPoints) SetPoints(v float32) {
	o.Points = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AddLoyaltyPoints) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPoints) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AddLoyaltyPoints) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AddLoyaltyPoints) SetName(v string) {
	o.Name = &v
}

// GetValidityDuration returns the ValidityDuration field value if set, zero value otherwise.
func (o *AddLoyaltyPoints) GetValidityDuration() string {
	if o == nil || o.ValidityDuration == nil {
		var ret string
		return ret
	}
	return *o.ValidityDuration
}

// GetValidityDurationOk returns a tuple with the ValidityDuration field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPoints) GetValidityDurationOk() (string, bool) {
	if o == nil || o.ValidityDuration == nil {
		var ret string
		return ret, false
	}
	return *o.ValidityDuration, true
}

// HasValidityDuration returns a boolean if a field has been set.
func (o *AddLoyaltyPoints) HasValidityDuration() bool {
	if o != nil && o.ValidityDuration != nil {
		return true
	}

	return false
}

// SetValidityDuration gets a reference to the given string and assigns it to the ValidityDuration field.
func (o *AddLoyaltyPoints) SetValidityDuration(v string) {
	o.ValidityDuration = &v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *AddLoyaltyPoints) GetValidUntil() time.Time {
	if o == nil || o.ValidUntil == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPoints) GetValidUntilOk() (time.Time, bool) {
	if o == nil || o.ValidUntil == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *AddLoyaltyPoints) HasValidUntil() bool {
	if o != nil && o.ValidUntil != nil {
		return true
	}

	return false
}

// SetValidUntil gets a reference to the given time.Time and assigns it to the ValidUntil field.
func (o *AddLoyaltyPoints) SetValidUntil(v time.Time) {
	o.ValidUntil = &v
}

// GetPendingDuration returns the PendingDuration field value if set, zero value otherwise.
func (o *AddLoyaltyPoints) GetPendingDuration() string {
	if o == nil || o.PendingDuration == nil {
		var ret string
		return ret
	}
	return *o.PendingDuration
}

// GetPendingDurationOk returns a tuple with the PendingDuration field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPoints) GetPendingDurationOk() (string, bool) {
	if o == nil || o.PendingDuration == nil {
		var ret string
		return ret, false
	}
	return *o.PendingDuration, true
}

// HasPendingDuration returns a boolean if a field has been set.
func (o *AddLoyaltyPoints) HasPendingDuration() bool {
	if o != nil && o.PendingDuration != nil {
		return true
	}

	return false
}

// SetPendingDuration gets a reference to the given string and assigns it to the PendingDuration field.
func (o *AddLoyaltyPoints) SetPendingDuration(v string) {
	o.PendingDuration = &v
}

// GetPendingUntil returns the PendingUntil field value if set, zero value otherwise.
func (o *AddLoyaltyPoints) GetPendingUntil() time.Time {
	if o == nil || o.PendingUntil == nil {
		var ret time.Time
		return ret
	}
	return *o.PendingUntil
}

// GetPendingUntilOk returns a tuple with the PendingUntil field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPoints) GetPendingUntilOk() (time.Time, bool) {
	if o == nil || o.PendingUntil == nil {
		var ret time.Time
		return ret, false
	}
	return *o.PendingUntil, true
}

// HasPendingUntil returns a boolean if a field has been set.
func (o *AddLoyaltyPoints) HasPendingUntil() bool {
	if o != nil && o.PendingUntil != nil {
		return true
	}

	return false
}

// SetPendingUntil gets a reference to the given time.Time and assigns it to the PendingUntil field.
func (o *AddLoyaltyPoints) SetPendingUntil(v time.Time) {
	o.PendingUntil = &v
}

// GetSubledgerId returns the SubledgerId field value if set, zero value otherwise.
func (o *AddLoyaltyPoints) GetSubledgerId() string {
	if o == nil || o.SubledgerId == nil {
		var ret string
		return ret
	}
	return *o.SubledgerId
}

// GetSubledgerIdOk returns a tuple with the SubledgerId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPoints) GetSubledgerIdOk() (string, bool) {
	if o == nil || o.SubledgerId == nil {
		var ret string
		return ret, false
	}
	return *o.SubledgerId, true
}

// HasSubledgerId returns a boolean if a field has been set.
func (o *AddLoyaltyPoints) HasSubledgerId() bool {
	if o != nil && o.SubledgerId != nil {
		return true
	}

	return false
}

// SetSubledgerId gets a reference to the given string and assigns it to the SubledgerId field.
func (o *AddLoyaltyPoints) SetSubledgerId(v string) {
	o.SubledgerId = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *AddLoyaltyPoints) GetApplicationId() int32 {
	if o == nil || o.ApplicationId == nil {
		var ret int32
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPoints) GetApplicationIdOk() (int32, bool) {
	if o == nil || o.ApplicationId == nil {
		var ret int32
		return ret, false
	}
	return *o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *AddLoyaltyPoints) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.
func (o *AddLoyaltyPoints) SetApplicationId(v int32) {
	o.ApplicationId = &v
}

type NullableAddLoyaltyPoints struct {
	Value        AddLoyaltyPoints
	ExplicitNull bool
}

func (v NullableAddLoyaltyPoints) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAddLoyaltyPoints) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}