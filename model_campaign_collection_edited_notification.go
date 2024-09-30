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

// CampaignCollectionEditedNotification A notification regarding a collection that was edited.
type CampaignCollectionEditedNotification struct {
	Campaign   Campaign                 `json:"campaign"`
	Ruleset    *Ruleset                 `json:"ruleset,omitempty"`
	Collection CollectionWithoutPayload `json:"collection"`
}

// GetCampaign returns the Campaign field value
func (o *CampaignCollectionEditedNotification) GetCampaign() Campaign {
	if o == nil {
		var ret Campaign
		return ret
	}

	return o.Campaign
}

// SetCampaign sets field value
func (o *CampaignCollectionEditedNotification) SetCampaign(v Campaign) {
	o.Campaign = v
}

// GetRuleset returns the Ruleset field value if set, zero value otherwise.
func (o *CampaignCollectionEditedNotification) GetRuleset() Ruleset {
	if o == nil || o.Ruleset == nil {
		var ret Ruleset
		return ret
	}
	return *o.Ruleset
}

// GetRulesetOk returns a tuple with the Ruleset field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCollectionEditedNotification) GetRulesetOk() (Ruleset, bool) {
	if o == nil || o.Ruleset == nil {
		var ret Ruleset
		return ret, false
	}
	return *o.Ruleset, true
}

// HasRuleset returns a boolean if a field has been set.
func (o *CampaignCollectionEditedNotification) HasRuleset() bool {
	if o != nil && o.Ruleset != nil {
		return true
	}

	return false
}

// SetRuleset gets a reference to the given Ruleset and assigns it to the Ruleset field.
func (o *CampaignCollectionEditedNotification) SetRuleset(v Ruleset) {
	o.Ruleset = &v
}

// GetCollection returns the Collection field value
func (o *CampaignCollectionEditedNotification) GetCollection() CollectionWithoutPayload {
	if o == nil {
		var ret CollectionWithoutPayload
		return ret
	}

	return o.Collection
}

// SetCollection sets field value
func (o *CampaignCollectionEditedNotification) SetCollection(v CollectionWithoutPayload) {
	o.Collection = v
}

type NullableCampaignCollectionEditedNotification struct {
	Value        CampaignCollectionEditedNotification
	ExplicitNull bool
}

func (v NullableCampaignCollectionEditedNotification) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCampaignCollectionEditedNotification) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
