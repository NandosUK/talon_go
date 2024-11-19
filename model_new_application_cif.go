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

// NewApplicationCif struct for NewApplicationCif
type NewApplicationCif struct {
	// The name of the Application cart item filter used in API requests.
	Name string `json:"name"`
	// A short description of the Application cart item filter.
	Description *string `json:"description,omitempty"`
	// The ID of the expression that the Application cart item filter uses.
	ActiveExpressionId *int32 `json:"activeExpressionId,omitempty"`
	// The ID of the user who last updated the Application cart item filter.
	ModifiedBy *int32 `json:"modifiedBy,omitempty"`
	// The ID of the user who created the Application cart item filter.
	CreatedBy *int32 `json:"createdBy,omitempty"`
	// Timestamp of the most recent update to the Application cart item filter.
	Modified *time.Time `json:"modified,omitempty"`
}

// GetName returns the Name field value
func (o *NewApplicationCif) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *NewApplicationCif) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NewApplicationCif) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewApplicationCif) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NewApplicationCif) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NewApplicationCif) SetDescription(v string) {
	o.Description = &v
}

// GetActiveExpressionId returns the ActiveExpressionId field value if set, zero value otherwise.
func (o *NewApplicationCif) GetActiveExpressionId() int32 {
	if o == nil || o.ActiveExpressionId == nil {
		var ret int32
		return ret
	}
	return *o.ActiveExpressionId
}

// GetActiveExpressionIdOk returns a tuple with the ActiveExpressionId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewApplicationCif) GetActiveExpressionIdOk() (int32, bool) {
	if o == nil || o.ActiveExpressionId == nil {
		var ret int32
		return ret, false
	}
	return *o.ActiveExpressionId, true
}

// HasActiveExpressionId returns a boolean if a field has been set.
func (o *NewApplicationCif) HasActiveExpressionId() bool {
	if o != nil && o.ActiveExpressionId != nil {
		return true
	}

	return false
}

// SetActiveExpressionId gets a reference to the given int32 and assigns it to the ActiveExpressionId field.
func (o *NewApplicationCif) SetActiveExpressionId(v int32) {
	o.ActiveExpressionId = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *NewApplicationCif) GetModifiedBy() int32 {
	if o == nil || o.ModifiedBy == nil {
		var ret int32
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewApplicationCif) GetModifiedByOk() (int32, bool) {
	if o == nil || o.ModifiedBy == nil {
		var ret int32
		return ret, false
	}
	return *o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *NewApplicationCif) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given int32 and assigns it to the ModifiedBy field.
func (o *NewApplicationCif) SetModifiedBy(v int32) {
	o.ModifiedBy = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *NewApplicationCif) GetCreatedBy() int32 {
	if o == nil || o.CreatedBy == nil {
		var ret int32
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewApplicationCif) GetCreatedByOk() (int32, bool) {
	if o == nil || o.CreatedBy == nil {
		var ret int32
		return ret, false
	}
	return *o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *NewApplicationCif) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.
func (o *NewApplicationCif) SetCreatedBy(v int32) {
	o.CreatedBy = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *NewApplicationCif) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewApplicationCif) GetModifiedOk() (time.Time, bool) {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret, false
	}
	return *o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *NewApplicationCif) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *NewApplicationCif) SetModified(v time.Time) {
	o.Modified = &v
}

type NullableNewApplicationCif struct {
	Value        NewApplicationCif
	ExplicitNull bool
}

func (v NullableNewApplicationCif) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewApplicationCif) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}