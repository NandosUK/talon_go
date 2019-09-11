/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package talon

type CodeGeneratorSettings struct {
	// Set of characters to be used when generating random part of code. Defaults to [A-Z, 0-9] (in terms of RegExp).
	ValidCharacters []string `json:"validCharacters"`
	// The pattern that will be used to generate coupon codes. The character `#` acts as a placeholder and will be replaced by a random character from the `validCharacters` set. 
	CouponPattern string `json:"couponPattern"`
}