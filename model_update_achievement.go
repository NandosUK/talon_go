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

// UpdateAchievement struct for UpdateAchievement
type UpdateAchievement struct {
	// The internal name of the achievement used in API requests.
	Name *string `json:"name,omitempty"`
	// The display name for the achievement in the Campaign Manager.
	Title *string `json:"title,omitempty"`
	// A description of the achievement.
	Description *string `json:"description,omitempty"`
	// The required number of actions or the transactional milestone to complete the achievement.
	Target *float32 `json:"target,omitempty"`
	// The relative duration after which the achievement ends and resets for a particular customer profile.
	Period            *string    `json:"period,omitempty"`
	PeriodEndOverride *TimePoint `json:"periodEndOverride,omitempty"`
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateAchievement) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAchievement) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAchievement) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateAchievement) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UpdateAchievement) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAchievement) GetTitleOk() (string, bool) {
	if o == nil || o.Title == nil {
		var ret string
		return ret, false
	}
	return *o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateAchievement) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UpdateAchievement) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateAchievement) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAchievement) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateAchievement) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateAchievement) SetDescription(v string) {
	o.Description = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *UpdateAchievement) GetTarget() float32 {
	if o == nil || o.Target == nil {
		var ret float32
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAchievement) GetTargetOk() (float32, bool) {
	if o == nil || o.Target == nil {
		var ret float32
		return ret, false
	}
	return *o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *UpdateAchievement) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given float32 and assigns it to the Target field.
func (o *UpdateAchievement) SetTarget(v float32) {
	o.Target = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *UpdateAchievement) GetPeriod() string {
	if o == nil || o.Period == nil {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAchievement) GetPeriodOk() (string, bool) {
	if o == nil || o.Period == nil {
		var ret string
		return ret, false
	}
	return *o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *UpdateAchievement) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *UpdateAchievement) SetPeriod(v string) {
	o.Period = &v
}

// GetPeriodEndOverride returns the PeriodEndOverride field value if set, zero value otherwise.
func (o *UpdateAchievement) GetPeriodEndOverride() TimePoint {
	if o == nil || o.PeriodEndOverride == nil {
		var ret TimePoint
		return ret
	}
	return *o.PeriodEndOverride
}

// GetPeriodEndOverrideOk returns a tuple with the PeriodEndOverride field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAchievement) GetPeriodEndOverrideOk() (TimePoint, bool) {
	if o == nil || o.PeriodEndOverride == nil {
		var ret TimePoint
		return ret, false
	}
	return *o.PeriodEndOverride, true
}

// HasPeriodEndOverride returns a boolean if a field has been set.
func (o *UpdateAchievement) HasPeriodEndOverride() bool {
	if o != nil && o.PeriodEndOverride != nil {
		return true
	}

	return false
}

// SetPeriodEndOverride gets a reference to the given TimePoint and assigns it to the PeriodEndOverride field.
func (o *UpdateAchievement) SetPeriodEndOverride(v TimePoint) {
	o.PeriodEndOverride = &v
}

type NullableUpdateAchievement struct {
	Value        UpdateAchievement
	ExplicitNull bool
}

func (v NullableUpdateAchievement) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateAchievement) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}