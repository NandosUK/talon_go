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

// UpdateCouponBatch
type UpdateCouponBatch struct {
	// The number of times the coupon code can be redeemed. `0` means unlimited redemptions but any campaign usage limits will still apply.
	UsageLimit *int32 `json:"usageLimit,omitempty"`
	// The total discount value that the code can give. Typically used to represent a gift card value.
	DiscountLimit *float32 `json:"discountLimit,omitempty"`
	// The number of reservations that can be made with this coupon code.
	ReservationLimit *int32 `json:"reservationLimit,omitempty"`
	// Timestamp at which point the coupon becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Expiration date of the coupon. Coupon never expires if this is omitted, zero, or negative.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// Optional property to set the value of custom coupon attributes. They are defined in the Campaign Manager, see [Managing attributes](https://docs.talon.one/docs/product/account/dev-tools/managing-attributes).  Coupon attributes can also be set to _mandatory_ in your Application [settings](https://docs.talon.one/docs/product/applications/using-attributes#making-attributes-mandatory). If your Application uses mandatory attributes, you must use this property to set their value.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// The ID of the batch the coupon(s) belong to.
	BatchID *string `json:"batchID,omitempty"`
}

// GetUsageLimit returns the UsageLimit field value if set, zero value otherwise.
func (o *UpdateCouponBatch) GetUsageLimit() int32 {
	if o == nil || o.UsageLimit == nil {
		var ret int32
		return ret
	}
	return *o.UsageLimit
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCouponBatch) GetUsageLimitOk() (int32, bool) {
	if o == nil || o.UsageLimit == nil {
		var ret int32
		return ret, false
	}
	return *o.UsageLimit, true
}

// HasUsageLimit returns a boolean if a field has been set.
func (o *UpdateCouponBatch) HasUsageLimit() bool {
	if o != nil && o.UsageLimit != nil {
		return true
	}

	return false
}

// SetUsageLimit gets a reference to the given int32 and assigns it to the UsageLimit field.
func (o *UpdateCouponBatch) SetUsageLimit(v int32) {
	o.UsageLimit = &v
}

// GetDiscountLimit returns the DiscountLimit field value if set, zero value otherwise.
func (o *UpdateCouponBatch) GetDiscountLimit() float32 {
	if o == nil || o.DiscountLimit == nil {
		var ret float32
		return ret
	}
	return *o.DiscountLimit
}

// GetDiscountLimitOk returns a tuple with the DiscountLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCouponBatch) GetDiscountLimitOk() (float32, bool) {
	if o == nil || o.DiscountLimit == nil {
		var ret float32
		return ret, false
	}
	return *o.DiscountLimit, true
}

// HasDiscountLimit returns a boolean if a field has been set.
func (o *UpdateCouponBatch) HasDiscountLimit() bool {
	if o != nil && o.DiscountLimit != nil {
		return true
	}

	return false
}

// SetDiscountLimit gets a reference to the given float32 and assigns it to the DiscountLimit field.
func (o *UpdateCouponBatch) SetDiscountLimit(v float32) {
	o.DiscountLimit = &v
}

// GetReservationLimit returns the ReservationLimit field value if set, zero value otherwise.
func (o *UpdateCouponBatch) GetReservationLimit() int32 {
	if o == nil || o.ReservationLimit == nil {
		var ret int32
		return ret
	}
	return *o.ReservationLimit
}

// GetReservationLimitOk returns a tuple with the ReservationLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCouponBatch) GetReservationLimitOk() (int32, bool) {
	if o == nil || o.ReservationLimit == nil {
		var ret int32
		return ret, false
	}
	return *o.ReservationLimit, true
}

// HasReservationLimit returns a boolean if a field has been set.
func (o *UpdateCouponBatch) HasReservationLimit() bool {
	if o != nil && o.ReservationLimit != nil {
		return true
	}

	return false
}

// SetReservationLimit gets a reference to the given int32 and assigns it to the ReservationLimit field.
func (o *UpdateCouponBatch) SetReservationLimit(v int32) {
	o.ReservationLimit = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *UpdateCouponBatch) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCouponBatch) GetStartDateOk() (time.Time, bool) {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *UpdateCouponBatch) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *UpdateCouponBatch) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *UpdateCouponBatch) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCouponBatch) GetExpiryDateOk() (time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *UpdateCouponBatch) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *UpdateCouponBatch) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UpdateCouponBatch) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCouponBatch) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UpdateCouponBatch) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *UpdateCouponBatch) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetBatchID returns the BatchID field value if set, zero value otherwise.
func (o *UpdateCouponBatch) GetBatchID() string {
	if o == nil || o.BatchID == nil {
		var ret string
		return ret
	}
	return *o.BatchID
}

// GetBatchIDOk returns a tuple with the BatchID field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCouponBatch) GetBatchIDOk() (string, bool) {
	if o == nil || o.BatchID == nil {
		var ret string
		return ret, false
	}
	return *o.BatchID, true
}

// HasBatchID returns a boolean if a field has been set.
func (o *UpdateCouponBatch) HasBatchID() bool {
	if o != nil && o.BatchID != nil {
		return true
	}

	return false
}

// SetBatchID gets a reference to the given string and assigns it to the BatchID field.
func (o *UpdateCouponBatch) SetBatchID(v string) {
	o.BatchID = &v
}

type NullableUpdateCouponBatch struct {
	Value        UpdateCouponBatch
	ExplicitNull bool
}

func (v NullableUpdateCouponBatch) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateCouponBatch) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
