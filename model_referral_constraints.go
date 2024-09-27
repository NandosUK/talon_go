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

// ReferralConstraints struct for ReferralConstraints
type ReferralConstraints struct {
	// Timestamp at which point the referral code becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Expiration date of the referral code. Referral never expires if this is omitted.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// The number of times a referral code can be used. `0` means no limit but any campaign usage limits will still apply.
	UsageLimit *int32 `json:"usageLimit,omitempty"`
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ReferralConstraints) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReferralConstraints) GetStartDateOk() (time.Time, bool) {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ReferralConstraints) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *ReferralConstraints) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *ReferralConstraints) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReferralConstraints) GetExpiryDateOk() (time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *ReferralConstraints) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *ReferralConstraints) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetUsageLimit returns the UsageLimit field value if set, zero value otherwise.
func (o *ReferralConstraints) GetUsageLimit() int32 {
	if o == nil || o.UsageLimit == nil {
		var ret int32
		return ret
	}
	return *o.UsageLimit
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReferralConstraints) GetUsageLimitOk() (int32, bool) {
	if o == nil || o.UsageLimit == nil {
		var ret int32
		return ret, false
	}
	return *o.UsageLimit, true
}

// HasUsageLimit returns a boolean if a field has been set.
func (o *ReferralConstraints) HasUsageLimit() bool {
	if o != nil && o.UsageLimit != nil {
		return true
	}

	return false
}

// SetUsageLimit gets a reference to the given int32 and assigns it to the UsageLimit field.
func (o *ReferralConstraints) SetUsageLimit(v int32) {
	o.UsageLimit = &v
}

type NullableReferralConstraints struct {
	Value        ReferralConstraints
	ExplicitNull bool
}

func (v NullableReferralConstraints) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableReferralConstraints) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
