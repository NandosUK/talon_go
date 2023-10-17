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

// Coupon
type Coupon struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The ID of the campaign that owns this entity.
	CampaignId int32 `json:"campaignId"`
	// The coupon code.
	Value string `json:"value"`
	// The number of times the coupon code can be redeemed. `0` means unlimited redemptions but any campaign usage limits will still apply.
	UsageLimit int32 `json:"usageLimit"`
	// The total discount value that the code can give. Typically used to represent a gift card value.
	DiscountLimit *float32 `json:"discountLimit,omitempty"`
	// The number of reservations that can be made with this coupon code.
	ReservationLimit *int32 `json:"reservationLimit,omitempty"`
	// Timestamp at which point the coupon becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Expiration date of the coupon. Coupon never expires if this is omitted, zero, or negative.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// Limits configuration for a coupon. These limits will override the limits set from the campaign.  **Note:** Only usable when creating a single coupon which is not tied to a specific recipient. Only per-profile limits are allowed to be configured.
	Limits *[]LimitConfig `json:"limits,omitempty"`
	// The number of times the coupon has been successfully redeemed.
	UsageCounter int32 `json:"usageCounter"`
	// The amount of discounts given on rules redeeming this coupon. Only usable if a coupon discount budget was set for this coupon.
	DiscountCounter *float32 `json:"discountCounter,omitempty"`
	// The remaining discount this coupon can give.
	DiscountRemainder *float32 `json:"discountRemainder,omitempty"`
	// The number of times this coupon has been reserved.
	ReservationCounter *float32 `json:"reservationCounter,omitempty"`
	// Custom attributes associated with this coupon.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// The integration ID of the referring customer (if any) for whom this coupon was created as an effect.
	ReferralId *int32 `json:"referralId,omitempty"`
	// The Integration ID of the customer that is allowed to redeem this coupon.
	RecipientIntegrationId *string `json:"recipientIntegrationId,omitempty"`
	// The ID of the Import which created this coupon.
	ImportId *int32 `json:"importId,omitempty"`
	// Defines the type of reservation: - `true`: The reservation is a soft reservation. Any customer can use the coupon. This is done via the [Create coupon reservation](https://docs.talon.one/integration-api#operation/createCouponReservation) endpoint. - `false`: The reservation is a hard reservation. Only the associated customer (`recipientIntegrationId`) can use the coupon. This is done via the Campaign Manager when you create a coupon for a given `recipientIntegrationId`, the [Create coupons](https://docs.talon.one/management-api#operation/createCoupons) endpoint or [Create coupons for multiple recipients](https://docs.talon.one/management-api#operation/createCouponsForMultipleRecipients) endpoint.
	Reservation *bool `json:"reservation,omitempty"`
	// The id of the batch the coupon belongs to.
	BatchId *string `json:"batchId,omitempty"`
	// Whether the reservation effect actually created a new reservation.
	IsReservationMandatory *bool `json:"isReservationMandatory,omitempty"`
	// An indication of whether the coupon is implicitly reserved for all customers.
	ImplicitlyReserved *bool `json:"implicitlyReserved,omitempty"`
}

// GetId returns the Id field value
func (o *Coupon) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *Coupon) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Coupon) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *Coupon) SetCreated(v time.Time) {
	o.Created = v
}

// GetCampaignId returns the CampaignId field value
func (o *Coupon) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// SetCampaignId sets field value
func (o *Coupon) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetValue returns the Value field value
func (o *Coupon) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *Coupon) SetValue(v string) {
	o.Value = v
}

// GetUsageLimit returns the UsageLimit field value
func (o *Coupon) GetUsageLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageLimit
}

// SetUsageLimit sets field value
func (o *Coupon) SetUsageLimit(v int32) {
	o.UsageLimit = v
}

// GetDiscountLimit returns the DiscountLimit field value if set, zero value otherwise.
func (o *Coupon) GetDiscountLimit() float32 {
	if o == nil || o.DiscountLimit == nil {
		var ret float32
		return ret
	}
	return *o.DiscountLimit
}

// GetDiscountLimitOk returns a tuple with the DiscountLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetDiscountLimitOk() (float32, bool) {
	if o == nil || o.DiscountLimit == nil {
		var ret float32
		return ret, false
	}
	return *o.DiscountLimit, true
}

// HasDiscountLimit returns a boolean if a field has been set.
func (o *Coupon) HasDiscountLimit() bool {
	if o != nil && o.DiscountLimit != nil {
		return true
	}

	return false
}

// SetDiscountLimit gets a reference to the given float32 and assigns it to the DiscountLimit field.
func (o *Coupon) SetDiscountLimit(v float32) {
	o.DiscountLimit = &v
}

// GetReservationLimit returns the ReservationLimit field value if set, zero value otherwise.
func (o *Coupon) GetReservationLimit() int32 {
	if o == nil || o.ReservationLimit == nil {
		var ret int32
		return ret
	}
	return *o.ReservationLimit
}

// GetReservationLimitOk returns a tuple with the ReservationLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetReservationLimitOk() (int32, bool) {
	if o == nil || o.ReservationLimit == nil {
		var ret int32
		return ret, false
	}
	return *o.ReservationLimit, true
}

// HasReservationLimit returns a boolean if a field has been set.
func (o *Coupon) HasReservationLimit() bool {
	if o != nil && o.ReservationLimit != nil {
		return true
	}

	return false
}

// SetReservationLimit gets a reference to the given int32 and assigns it to the ReservationLimit field.
func (o *Coupon) SetReservationLimit(v int32) {
	o.ReservationLimit = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Coupon) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetStartDateOk() (time.Time, bool) {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Coupon) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *Coupon) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *Coupon) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetExpiryDateOk() (time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *Coupon) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *Coupon) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *Coupon) GetLimits() []LimitConfig {
	if o == nil || o.Limits == nil {
		var ret []LimitConfig
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetLimitsOk() ([]LimitConfig, bool) {
	if o == nil || o.Limits == nil {
		var ret []LimitConfig
		return ret, false
	}
	return *o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *Coupon) HasLimits() bool {
	if o != nil && o.Limits != nil {
		return true
	}

	return false
}

// SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.
func (o *Coupon) SetLimits(v []LimitConfig) {
	o.Limits = &v
}

// GetUsageCounter returns the UsageCounter field value
func (o *Coupon) GetUsageCounter() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageCounter
}

// SetUsageCounter sets field value
func (o *Coupon) SetUsageCounter(v int32) {
	o.UsageCounter = v
}

// GetDiscountCounter returns the DiscountCounter field value if set, zero value otherwise.
func (o *Coupon) GetDiscountCounter() float32 {
	if o == nil || o.DiscountCounter == nil {
		var ret float32
		return ret
	}
	return *o.DiscountCounter
}

// GetDiscountCounterOk returns a tuple with the DiscountCounter field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetDiscountCounterOk() (float32, bool) {
	if o == nil || o.DiscountCounter == nil {
		var ret float32
		return ret, false
	}
	return *o.DiscountCounter, true
}

// HasDiscountCounter returns a boolean if a field has been set.
func (o *Coupon) HasDiscountCounter() bool {
	if o != nil && o.DiscountCounter != nil {
		return true
	}

	return false
}

// SetDiscountCounter gets a reference to the given float32 and assigns it to the DiscountCounter field.
func (o *Coupon) SetDiscountCounter(v float32) {
	o.DiscountCounter = &v
}

// GetDiscountRemainder returns the DiscountRemainder field value if set, zero value otherwise.
func (o *Coupon) GetDiscountRemainder() float32 {
	if o == nil || o.DiscountRemainder == nil {
		var ret float32
		return ret
	}
	return *o.DiscountRemainder
}

// GetDiscountRemainderOk returns a tuple with the DiscountRemainder field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetDiscountRemainderOk() (float32, bool) {
	if o == nil || o.DiscountRemainder == nil {
		var ret float32
		return ret, false
	}
	return *o.DiscountRemainder, true
}

// HasDiscountRemainder returns a boolean if a field has been set.
func (o *Coupon) HasDiscountRemainder() bool {
	if o != nil && o.DiscountRemainder != nil {
		return true
	}

	return false
}

// SetDiscountRemainder gets a reference to the given float32 and assigns it to the DiscountRemainder field.
func (o *Coupon) SetDiscountRemainder(v float32) {
	o.DiscountRemainder = &v
}

// GetReservationCounter returns the ReservationCounter field value if set, zero value otherwise.
func (o *Coupon) GetReservationCounter() float32 {
	if o == nil || o.ReservationCounter == nil {
		var ret float32
		return ret
	}
	return *o.ReservationCounter
}

// GetReservationCounterOk returns a tuple with the ReservationCounter field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetReservationCounterOk() (float32, bool) {
	if o == nil || o.ReservationCounter == nil {
		var ret float32
		return ret, false
	}
	return *o.ReservationCounter, true
}

// HasReservationCounter returns a boolean if a field has been set.
func (o *Coupon) HasReservationCounter() bool {
	if o != nil && o.ReservationCounter != nil {
		return true
	}

	return false
}

// SetReservationCounter gets a reference to the given float32 and assigns it to the ReservationCounter field.
func (o *Coupon) SetReservationCounter(v float32) {
	o.ReservationCounter = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Coupon) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Coupon) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *Coupon) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetReferralId returns the ReferralId field value if set, zero value otherwise.
func (o *Coupon) GetReferralId() int32 {
	if o == nil || o.ReferralId == nil {
		var ret int32
		return ret
	}
	return *o.ReferralId
}

// GetReferralIdOk returns a tuple with the ReferralId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetReferralIdOk() (int32, bool) {
	if o == nil || o.ReferralId == nil {
		var ret int32
		return ret, false
	}
	return *o.ReferralId, true
}

// HasReferralId returns a boolean if a field has been set.
func (o *Coupon) HasReferralId() bool {
	if o != nil && o.ReferralId != nil {
		return true
	}

	return false
}

// SetReferralId gets a reference to the given int32 and assigns it to the ReferralId field.
func (o *Coupon) SetReferralId(v int32) {
	o.ReferralId = &v
}

// GetRecipientIntegrationId returns the RecipientIntegrationId field value if set, zero value otherwise.
func (o *Coupon) GetRecipientIntegrationId() string {
	if o == nil || o.RecipientIntegrationId == nil {
		var ret string
		return ret
	}
	return *o.RecipientIntegrationId
}

// GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetRecipientIntegrationIdOk() (string, bool) {
	if o == nil || o.RecipientIntegrationId == nil {
		var ret string
		return ret, false
	}
	return *o.RecipientIntegrationId, true
}

// HasRecipientIntegrationId returns a boolean if a field has been set.
func (o *Coupon) HasRecipientIntegrationId() bool {
	if o != nil && o.RecipientIntegrationId != nil {
		return true
	}

	return false
}

// SetRecipientIntegrationId gets a reference to the given string and assigns it to the RecipientIntegrationId field.
func (o *Coupon) SetRecipientIntegrationId(v string) {
	o.RecipientIntegrationId = &v
}

// GetImportId returns the ImportId field value if set, zero value otherwise.
func (o *Coupon) GetImportId() int32 {
	if o == nil || o.ImportId == nil {
		var ret int32
		return ret
	}
	return *o.ImportId
}

// GetImportIdOk returns a tuple with the ImportId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetImportIdOk() (int32, bool) {
	if o == nil || o.ImportId == nil {
		var ret int32
		return ret, false
	}
	return *o.ImportId, true
}

// HasImportId returns a boolean if a field has been set.
func (o *Coupon) HasImportId() bool {
	if o != nil && o.ImportId != nil {
		return true
	}

	return false
}

// SetImportId gets a reference to the given int32 and assigns it to the ImportId field.
func (o *Coupon) SetImportId(v int32) {
	o.ImportId = &v
}

// GetReservation returns the Reservation field value if set, zero value otherwise.
func (o *Coupon) GetReservation() bool {
	if o == nil || o.Reservation == nil {
		var ret bool
		return ret
	}
	return *o.Reservation
}

// GetReservationOk returns a tuple with the Reservation field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetReservationOk() (bool, bool) {
	if o == nil || o.Reservation == nil {
		var ret bool
		return ret, false
	}
	return *o.Reservation, true
}

// HasReservation returns a boolean if a field has been set.
func (o *Coupon) HasReservation() bool {
	if o != nil && o.Reservation != nil {
		return true
	}

	return false
}

// SetReservation gets a reference to the given bool and assigns it to the Reservation field.
func (o *Coupon) SetReservation(v bool) {
	o.Reservation = &v
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *Coupon) GetBatchId() string {
	if o == nil || o.BatchId == nil {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetBatchIdOk() (string, bool) {
	if o == nil || o.BatchId == nil {
		var ret string
		return ret, false
	}
	return *o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *Coupon) HasBatchId() bool {
	if o != nil && o.BatchId != nil {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *Coupon) SetBatchId(v string) {
	o.BatchId = &v
}

// GetIsReservationMandatory returns the IsReservationMandatory field value if set, zero value otherwise.
func (o *Coupon) GetIsReservationMandatory() bool {
	if o == nil || o.IsReservationMandatory == nil {
		var ret bool
		return ret
	}
	return *o.IsReservationMandatory
}

// GetIsReservationMandatoryOk returns a tuple with the IsReservationMandatory field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetIsReservationMandatoryOk() (bool, bool) {
	if o == nil || o.IsReservationMandatory == nil {
		var ret bool
		return ret, false
	}
	return *o.IsReservationMandatory, true
}

// HasIsReservationMandatory returns a boolean if a field has been set.
func (o *Coupon) HasIsReservationMandatory() bool {
	if o != nil && o.IsReservationMandatory != nil {
		return true
	}

	return false
}

// SetIsReservationMandatory gets a reference to the given bool and assigns it to the IsReservationMandatory field.
func (o *Coupon) SetIsReservationMandatory(v bool) {
	o.IsReservationMandatory = &v
}

// GetImplicitlyReserved returns the ImplicitlyReserved field value if set, zero value otherwise.
func (o *Coupon) GetImplicitlyReserved() bool {
	if o == nil || o.ImplicitlyReserved == nil {
		var ret bool
		return ret
	}
	return *o.ImplicitlyReserved
}

// GetImplicitlyReservedOk returns a tuple with the ImplicitlyReserved field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetImplicitlyReservedOk() (bool, bool) {
	if o == nil || o.ImplicitlyReserved == nil {
		var ret bool
		return ret, false
	}
	return *o.ImplicitlyReserved, true
}

// HasImplicitlyReserved returns a boolean if a field has been set.
func (o *Coupon) HasImplicitlyReserved() bool {
	if o != nil && o.ImplicitlyReserved != nil {
		return true
	}

	return false
}

// SetImplicitlyReserved gets a reference to the given bool and assigns it to the ImplicitlyReserved field.
func (o *Coupon) SetImplicitlyReserved(v bool) {
	o.ImplicitlyReserved = &v
}

type NullableCoupon struct {
	Value        Coupon
	ExplicitNull bool
}

func (v NullableCoupon) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCoupon) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
