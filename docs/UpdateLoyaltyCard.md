# UpdateLoyaltyCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status of the loyalty card. Can be one of: [&#39;active&#39;, &#39;disabled&#39;]  | 

## Methods

### NewUpdateLoyaltyCard

`func NewUpdateLoyaltyCard(status string, ) *UpdateLoyaltyCard`

NewUpdateLoyaltyCard instantiates a new UpdateLoyaltyCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoyaltyCardWithDefaults

`func NewUpdateLoyaltyCardWithDefaults() *UpdateLoyaltyCard`

NewUpdateLoyaltyCardWithDefaults instantiates a new UpdateLoyaltyCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateLoyaltyCard) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateLoyaltyCard) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateLoyaltyCard) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


