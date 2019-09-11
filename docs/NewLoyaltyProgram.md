# NewLoyaltyProgram

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The internal name for the Loyalty Program. This is an immutable value. | [default to null]
**Title** | **string** | The display title for the Loyalty Program. | [default to null]
**Description** | **string** | Description of our Loyalty Program. | [optional] [default to null]
**SubscribedApplications** | **[]int32** | A list containing the IDs of all applications that are subscribed to this Loyalty Program. | [optional] [default to null]
**DefaultValidity** | **string** | Indicates the default duration after which new loyalty points should expire. The format is a number, followed by one letter indicating the unit; like &#39;1h&#39; or &#39;40m&#39; or &#39;30d&#39;. | [default to null]
**AllowSubledger** | **bool** | Indicates if this program supports subledgers inside the program | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

