/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// NewRole
type NewRole struct {
	// Name of the role
	Name string `json:"name"`
	// Description of the role
	Description *string `json:"description,omitempty"`
	// Role Policy this should be a stringified blob of json
	Acl string `json:"acl"`
	// An array of user identifiers
	Users []int32 `json:"users"`
}

// GetName returns the Name field value
func (o *NewRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *NewRole) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NewRole) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewRole) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NewRole) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NewRole) SetDescription(v string) {
	o.Description = &v
}

// GetAcl returns the Acl field value
func (o *NewRole) GetAcl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Acl
}

// SetAcl sets field value
func (o *NewRole) SetAcl(v string) {
	o.Acl = v
}

// GetUsers returns the Users field value
func (o *NewRole) GetUsers() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Users
}

// SetUsers sets field value
func (o *NewRole) SetUsers(v []int32) {
	o.Users = v
}

type NullableNewRole struct {
	Value        NewRole
	ExplicitNull bool
}

func (v NullableNewRole) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewRole) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
