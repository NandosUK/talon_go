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

// WillAwardGiveawayEffectProps The properties specific to the \"awardGiveaway\" effect when the session is not closed yet. This effect replaces \"awardGiveaway\" only when updating a session with any state other than \"closed\". This is to ensure no giveaway codes are leaked when they are still not guaranteed to be awarded.
type WillAwardGiveawayEffectProps struct {
	// The ID of the giveaways pool the code will be taken from.
	PoolId int32 `json:"poolId"`
	// The name of the giveaways pool the code will be taken from.
	PoolName string `json:"poolName"`
	// The integration ID of the profile that will be awarded the giveaway.
	RecipientIntegrationId string `json:"recipientIntegrationId"`
}

// GetPoolId returns the PoolId field value
func (o *WillAwardGiveawayEffectProps) GetPoolId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PoolId
}

// SetPoolId sets field value
func (o *WillAwardGiveawayEffectProps) SetPoolId(v int32) {
	o.PoolId = v
}

// GetPoolName returns the PoolName field value
func (o *WillAwardGiveawayEffectProps) GetPoolName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PoolName
}

// SetPoolName sets field value
func (o *WillAwardGiveawayEffectProps) SetPoolName(v string) {
	o.PoolName = v
}

// GetRecipientIntegrationId returns the RecipientIntegrationId field value
func (o *WillAwardGiveawayEffectProps) GetRecipientIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientIntegrationId
}

// SetRecipientIntegrationId sets field value
func (o *WillAwardGiveawayEffectProps) SetRecipientIntegrationId(v string) {
	o.RecipientIntegrationId = v
}

type NullableWillAwardGiveawayEffectProps struct {
	Value        WillAwardGiveawayEffectProps
	ExplicitNull bool
}

func (v NullableWillAwardGiveawayEffectProps) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableWillAwardGiveawayEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}