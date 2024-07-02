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

// TransferLoyaltyCard struct for TransferLoyaltyCard
type TransferLoyaltyCard struct {
	// The alphanumeric identifier of the loyalty card.
	NewCardIdentifier string `json:"newCardIdentifier"`
	// Reason for transferring and blocking the loyalty card.
	BlockReason string `json:"blockReason"`
}

// GetNewCardIdentifier returns the NewCardIdentifier field value
func (o *TransferLoyaltyCard) GetNewCardIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewCardIdentifier
}

// SetNewCardIdentifier sets field value
func (o *TransferLoyaltyCard) SetNewCardIdentifier(v string) {
	o.NewCardIdentifier = v
}

// GetBlockReason returns the BlockReason field value
func (o *TransferLoyaltyCard) GetBlockReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockReason
}

// SetBlockReason sets field value
func (o *TransferLoyaltyCard) SetBlockReason(v string) {
	o.BlockReason = v
}

type NullableTransferLoyaltyCard struct {
	Value        TransferLoyaltyCard
	ExplicitNull bool
}

func (v NullableTransferLoyaltyCard) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableTransferLoyaltyCard) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
