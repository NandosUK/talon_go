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

// ScimBaseUser Schema definition for base user fields, provisioned using the SCIM protocol and used by Talon.One.
type ScimBaseUser struct {
	// Status of the user.
	Active *bool `json:"active,omitempty"`
	// Display name of the user.
	DisplayName *string `json:"displayName,omitempty"`
	// Unique identifier of the user. This is usually an email address.
	UserName *string           `json:"userName,omitempty"`
	Name     *ScimBaseUserName `json:"name,omitempty"`
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ScimBaseUser) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ScimBaseUser) GetActiveOk() (bool, bool) {
	if o == nil || o.Active == nil {
		var ret bool
		return ret, false
	}
	return *o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ScimBaseUser) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ScimBaseUser) SetActive(v bool) {
	o.Active = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ScimBaseUser) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ScimBaseUser) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ScimBaseUser) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ScimBaseUser) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *ScimBaseUser) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ScimBaseUser) GetUserNameOk() (string, bool) {
	if o == nil || o.UserName == nil {
		var ret string
		return ret, false
	}
	return *o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *ScimBaseUser) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *ScimBaseUser) SetUserName(v string) {
	o.UserName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScimBaseUser) GetName() ScimBaseUserName {
	if o == nil || o.Name == nil {
		var ret ScimBaseUserName
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ScimBaseUser) GetNameOk() (ScimBaseUserName, bool) {
	if o == nil || o.Name == nil {
		var ret ScimBaseUserName
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScimBaseUser) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given ScimBaseUserName and assigns it to the Name field.
func (o *ScimBaseUser) SetName(v ScimBaseUserName) {
	o.Name = &v
}

type NullableScimBaseUser struct {
	Value        ScimBaseUser
	ExplicitNull bool
}

func (v NullableScimBaseUser) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableScimBaseUser) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
