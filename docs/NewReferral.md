# NewReferral

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **time.Time** | Timestamp at which point the referral code becomes valid. | [optional] 
**ExpiryDate** | Pointer to **time.Time** | Expiry date of the referral code. Referral never expires if this is omitted, zero, or negative. | [optional] 
**UsageLimit** | Pointer to **int32** | The number of times a referral code can be used. &#x60;0&#x60; means no limit but any campaign usage limits will still apply.  | [optional] 
**CampaignId** | **int32** | ID of the campaign from which the referral received the referral code. | 
**AdvocateProfileIntegrationId** | **string** | The Integration ID of the Advocate&#39;s Profile. | 
**FriendProfileIntegrationId** | Pointer to **string** | An optional Integration ID of the Friend&#39;s Profile. | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | Arbitrary properties associated with this item. | [optional] 

## Methods

### NewNewReferral

`func NewNewReferral(campaignId int32, advocateProfileIntegrationId string, ) *NewReferral`

NewNewReferral instantiates a new NewReferral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewReferralWithDefaults

`func NewNewReferralWithDefaults() *NewReferral`

NewNewReferralWithDefaults instantiates a new NewReferral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *NewReferral) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *NewReferral) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *NewReferral) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *NewReferral) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *NewReferral) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *NewReferral) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *NewReferral) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *NewReferral) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetUsageLimit

`func (o *NewReferral) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *NewReferral) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *NewReferral) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *NewReferral) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetCampaignId

`func (o *NewReferral) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *NewReferral) GetCampaignIdOk() (*int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *NewReferral) SetCampaignId(v int32)`

SetCampaignId sets CampaignId field to given value.


### GetAdvocateProfileIntegrationId

`func (o *NewReferral) GetAdvocateProfileIntegrationId() string`

GetAdvocateProfileIntegrationId returns the AdvocateProfileIntegrationId field if non-nil, zero value otherwise.

### GetAdvocateProfileIntegrationIdOk

`func (o *NewReferral) GetAdvocateProfileIntegrationIdOk() (*string, bool)`

GetAdvocateProfileIntegrationIdOk returns a tuple with the AdvocateProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvocateProfileIntegrationId

`func (o *NewReferral) SetAdvocateProfileIntegrationId(v string)`

SetAdvocateProfileIntegrationId sets AdvocateProfileIntegrationId field to given value.


### GetFriendProfileIntegrationId

`func (o *NewReferral) GetFriendProfileIntegrationId() string`

GetFriendProfileIntegrationId returns the FriendProfileIntegrationId field if non-nil, zero value otherwise.

### GetFriendProfileIntegrationIdOk

`func (o *NewReferral) GetFriendProfileIntegrationIdOk() (*string, bool)`

GetFriendProfileIntegrationIdOk returns a tuple with the FriendProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendProfileIntegrationId

`func (o *NewReferral) SetFriendProfileIntegrationId(v string)`

SetFriendProfileIntegrationId sets FriendProfileIntegrationId field to given value.

### HasFriendProfileIntegrationId

`func (o *NewReferral) HasFriendProfileIntegrationId() bool`

HasFriendProfileIntegrationId returns a boolean if a field has been set.

### GetAttributes

`func (o *NewReferral) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewReferral) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NewReferral) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NewReferral) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


