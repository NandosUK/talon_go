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

// NotificationTest struct for NotificationTest
type NotificationTest struct {
	// The returned http response.
	HttpResponse string `json:"httpResponse"`
	// The returned http status code.
	HttpStatus int32 `json:"httpStatus"`
}

// GetHttpResponse returns the HttpResponse field value
func (o *NotificationTest) GetHttpResponse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HttpResponse
}

// SetHttpResponse sets field value
func (o *NotificationTest) SetHttpResponse(v string) {
	o.HttpResponse = v
}

// GetHttpStatus returns the HttpStatus field value
func (o *NotificationTest) GetHttpStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HttpStatus
}

// SetHttpStatus sets field value
func (o *NotificationTest) SetHttpStatus(v int32) {
	o.HttpStatus = v
}

type NullableNotificationTest struct {
	Value        NotificationTest
	ExplicitNull bool
}

func (v NullableNotificationTest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNotificationTest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}