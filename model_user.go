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

// User
type User struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The time this entity was last modified.
	Modified time.Time `json:"modified"`
	// The email address associated with the user profile.
	Email string `json:"email"`
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// Name of the user.
	Name string `json:"name"`
	// State of the user.
	State string `json:"state"`
	// Invitation token of the user.  **Note**: If the user has already accepted their invitation, this is `null`.
	InviteToken string `json:"inviteToken"`
	// Indicates whether the user is an `admin`.
	IsAdmin *bool `json:"isAdmin,omitempty"`
	// Access level of the user.
	Policy map[string]interface{} `json:"policy"`
	// A list of the IDs of the roles assigned to the user.
	Roles *[]int32 `json:"roles,omitempty"`
	// Authentication method for this user.
	AuthMethod *string `json:"authMethod,omitempty"`
	// Application notifications that the user is subscribed to.
	ApplicationNotificationSubscriptions *map[string]interface{} `json:"applicationNotificationSubscriptions,omitempty"`
	// Timestamp when the user last signed in to Talon.One.
	LastSignedIn *time.Time `json:"lastSignedIn,omitempty"`
	// Timestamp of the user's last activity after signing in to Talon.One.
	LastAccessed *time.Time `json:"lastAccessed,omitempty"`
	// Timestamp when the user was notified for feed.
	LatestFeedTimestamp *time.Time `json:"latestFeedTimestamp,omitempty"`
	// Additional user attributes, created and used by external identity providers.
	AdditionalAttributes *map[string]interface{} `json:"additionalAttributes,omitempty"`
}

// GetId returns the Id field value
func (o *User) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *User) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *User) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *User) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *User) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// SetModified sets field value
func (o *User) SetModified(v time.Time) {
	o.Modified = v
}

// GetEmail returns the Email field value
func (o *User) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// SetEmail sets field value
func (o *User) SetEmail(v string) {
	o.Email = v
}

// GetAccountId returns the AccountId field value
func (o *User) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// SetAccountId sets field value
func (o *User) SetAccountId(v int32) {
	o.AccountId = v
}

// GetName returns the Name field value
func (o *User) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *User) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *User) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// SetState sets field value
func (o *User) SetState(v string) {
	o.State = v
}

// GetInviteToken returns the InviteToken field value
func (o *User) GetInviteToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InviteToken
}

// SetInviteToken sets field value
func (o *User) SetInviteToken(v string) {
	o.InviteToken = v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *User) GetIsAdmin() bool {
	if o == nil || o.IsAdmin == nil {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIsAdminOk() (bool, bool) {
	if o == nil || o.IsAdmin == nil {
		var ret bool
		return ret, false
	}
	return *o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *User) HasIsAdmin() bool {
	if o != nil && o.IsAdmin != nil {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *User) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

// GetPolicy returns the Policy field value
func (o *User) GetPolicy() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Policy
}

// SetPolicy sets field value
func (o *User) SetPolicy(v map[string]interface{}) {
	o.Policy = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *User) GetRoles() []int32 {
	if o == nil || o.Roles == nil {
		var ret []int32
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetRolesOk() ([]int32, bool) {
	if o == nil || o.Roles == nil {
		var ret []int32
		return ret, false
	}
	return *o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *User) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []int32 and assigns it to the Roles field.
func (o *User) SetRoles(v []int32) {
	o.Roles = &v
}

// GetAuthMethod returns the AuthMethod field value if set, zero value otherwise.
func (o *User) GetAuthMethod() string {
	if o == nil || o.AuthMethod == nil {
		var ret string
		return ret
	}
	return *o.AuthMethod
}

// GetAuthMethodOk returns a tuple with the AuthMethod field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAuthMethodOk() (string, bool) {
	if o == nil || o.AuthMethod == nil {
		var ret string
		return ret, false
	}
	return *o.AuthMethod, true
}

// HasAuthMethod returns a boolean if a field has been set.
func (o *User) HasAuthMethod() bool {
	if o != nil && o.AuthMethod != nil {
		return true
	}

	return false
}

// SetAuthMethod gets a reference to the given string and assigns it to the AuthMethod field.
func (o *User) SetAuthMethod(v string) {
	o.AuthMethod = &v
}

// GetApplicationNotificationSubscriptions returns the ApplicationNotificationSubscriptions field value if set, zero value otherwise.
func (o *User) GetApplicationNotificationSubscriptions() map[string]interface{} {
	if o == nil || o.ApplicationNotificationSubscriptions == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ApplicationNotificationSubscriptions
}

// GetApplicationNotificationSubscriptionsOk returns a tuple with the ApplicationNotificationSubscriptions field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetApplicationNotificationSubscriptionsOk() (map[string]interface{}, bool) {
	if o == nil || o.ApplicationNotificationSubscriptions == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.ApplicationNotificationSubscriptions, true
}

// HasApplicationNotificationSubscriptions returns a boolean if a field has been set.
func (o *User) HasApplicationNotificationSubscriptions() bool {
	if o != nil && o.ApplicationNotificationSubscriptions != nil {
		return true
	}

	return false
}

// SetApplicationNotificationSubscriptions gets a reference to the given map[string]interface{} and assigns it to the ApplicationNotificationSubscriptions field.
func (o *User) SetApplicationNotificationSubscriptions(v map[string]interface{}) {
	o.ApplicationNotificationSubscriptions = &v
}

// GetLastSignedIn returns the LastSignedIn field value if set, zero value otherwise.
func (o *User) GetLastSignedIn() time.Time {
	if o == nil || o.LastSignedIn == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSignedIn
}

// GetLastSignedInOk returns a tuple with the LastSignedIn field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastSignedInOk() (time.Time, bool) {
	if o == nil || o.LastSignedIn == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastSignedIn, true
}

// HasLastSignedIn returns a boolean if a field has been set.
func (o *User) HasLastSignedIn() bool {
	if o != nil && o.LastSignedIn != nil {
		return true
	}

	return false
}

// SetLastSignedIn gets a reference to the given time.Time and assigns it to the LastSignedIn field.
func (o *User) SetLastSignedIn(v time.Time) {
	o.LastSignedIn = &v
}

// GetLastAccessed returns the LastAccessed field value if set, zero value otherwise.
func (o *User) GetLastAccessed() time.Time {
	if o == nil || o.LastAccessed == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAccessed
}

// GetLastAccessedOk returns a tuple with the LastAccessed field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastAccessedOk() (time.Time, bool) {
	if o == nil || o.LastAccessed == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastAccessed, true
}

// HasLastAccessed returns a boolean if a field has been set.
func (o *User) HasLastAccessed() bool {
	if o != nil && o.LastAccessed != nil {
		return true
	}

	return false
}

// SetLastAccessed gets a reference to the given time.Time and assigns it to the LastAccessed field.
func (o *User) SetLastAccessed(v time.Time) {
	o.LastAccessed = &v
}

// GetLatestFeedTimestamp returns the LatestFeedTimestamp field value if set, zero value otherwise.
func (o *User) GetLatestFeedTimestamp() time.Time {
	if o == nil || o.LatestFeedTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.LatestFeedTimestamp
}

// GetLatestFeedTimestampOk returns a tuple with the LatestFeedTimestamp field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLatestFeedTimestampOk() (time.Time, bool) {
	if o == nil || o.LatestFeedTimestamp == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LatestFeedTimestamp, true
}

// HasLatestFeedTimestamp returns a boolean if a field has been set.
func (o *User) HasLatestFeedTimestamp() bool {
	if o != nil && o.LatestFeedTimestamp != nil {
		return true
	}

	return false
}

// SetLatestFeedTimestamp gets a reference to the given time.Time and assigns it to the LatestFeedTimestamp field.
func (o *User) SetLatestFeedTimestamp(v time.Time) {
	o.LatestFeedTimestamp = &v
}

// GetAdditionalAttributes returns the AdditionalAttributes field value if set, zero value otherwise.
func (o *User) GetAdditionalAttributes() map[string]interface{} {
	if o == nil || o.AdditionalAttributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AdditionalAttributes
}

// GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAdditionalAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.AdditionalAttributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.AdditionalAttributes, true
}

// HasAdditionalAttributes returns a boolean if a field has been set.
func (o *User) HasAdditionalAttributes() bool {
	if o != nil && o.AdditionalAttributes != nil {
		return true
	}

	return false
}

// SetAdditionalAttributes gets a reference to the given map[string]interface{} and assigns it to the AdditionalAttributes field.
func (o *User) SetAdditionalAttributes(v map[string]interface{}) {
	o.AdditionalAttributes = &v
}

type NullableUser struct {
	Value        User
	ExplicitNull bool
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
