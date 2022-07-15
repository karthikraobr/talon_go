# LimitCounter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **int32** | The ID of the campaign that owns this entity. | 
**ApplicationId** | **int32** | The ID of the application that owns this entity. | 
**AccountId** | **int32** | The ID of the account that owns this entity. | 
**Id** | **int32** | Unique ID for this entity. | 
**Action** | **string** | The limitable action of the limit counter. | 
**ProfileId** | Pointer to **int32** | The profile ID for which this limit counter is used. | [optional] 
**ProfileIntegrationId** | Pointer to **string** | The profile integration ID for which this limit counter is used. | [optional] 
**CouponId** | Pointer to **int32** | The coupon ID for which this limit counter is used. | [optional] 
**CouponValue** | Pointer to **string** | The coupon value for which this limit counter is used. | [optional] 
**ReferralId** | Pointer to **int32** | The referral ID for which this limit counter is used. | [optional] 
**ReferralValue** | Pointer to **string** | The referral value for which this limit counter is used. | [optional] 
**Identifier** | Pointer to **int32** | The arbitrary identifier for which this limit counter is used. | [optional] 
**Period** | Pointer to **string** | The time period for which this limit counter is used. | [optional] 
**Limit** | **float32** | The highest possible value for this limit counter. | 
**Counter** | **float32** | The current value for this limit counter. | 

## Methods

### NewLimitCounter

`func NewLimitCounter(campaignId int32, applicationId int32, accountId int32, id int32, action string, limit float32, counter float32, ) *LimitCounter`

NewLimitCounter instantiates a new LimitCounter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitCounterWithDefaults

`func NewLimitCounterWithDefaults() *LimitCounter`

NewLimitCounterWithDefaults instantiates a new LimitCounter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *LimitCounter) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *LimitCounter) GetCampaignIdOk() (*int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *LimitCounter) SetCampaignId(v int32)`

SetCampaignId sets CampaignId field to given value.


### GetApplicationId

`func (o *LimitCounter) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *LimitCounter) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *LimitCounter) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.


### GetAccountId

`func (o *LimitCounter) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LimitCounter) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LimitCounter) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetId

`func (o *LimitCounter) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LimitCounter) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LimitCounter) SetId(v int32)`

SetId sets Id field to given value.


### GetAction

`func (o *LimitCounter) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LimitCounter) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *LimitCounter) SetAction(v string)`

SetAction sets Action field to given value.


### GetProfileId

`func (o *LimitCounter) GetProfileId() int32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *LimitCounter) GetProfileIdOk() (*int32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *LimitCounter) SetProfileId(v int32)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *LimitCounter) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileIntegrationId

`func (o *LimitCounter) GetProfileIntegrationId() string`

GetProfileIntegrationId returns the ProfileIntegrationId field if non-nil, zero value otherwise.

### GetProfileIntegrationIdOk

`func (o *LimitCounter) GetProfileIntegrationIdOk() (*string, bool)`

GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationId

`func (o *LimitCounter) SetProfileIntegrationId(v string)`

SetProfileIntegrationId sets ProfileIntegrationId field to given value.

### HasProfileIntegrationId

`func (o *LimitCounter) HasProfileIntegrationId() bool`

HasProfileIntegrationId returns a boolean if a field has been set.

### GetCouponId

`func (o *LimitCounter) GetCouponId() int32`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *LimitCounter) GetCouponIdOk() (*int32, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *LimitCounter) SetCouponId(v int32)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *LimitCounter) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### GetCouponValue

`func (o *LimitCounter) GetCouponValue() string`

GetCouponValue returns the CouponValue field if non-nil, zero value otherwise.

### GetCouponValueOk

`func (o *LimitCounter) GetCouponValueOk() (*string, bool)`

GetCouponValueOk returns a tuple with the CouponValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponValue

`func (o *LimitCounter) SetCouponValue(v string)`

SetCouponValue sets CouponValue field to given value.

### HasCouponValue

`func (o *LimitCounter) HasCouponValue() bool`

HasCouponValue returns a boolean if a field has been set.

### GetReferralId

`func (o *LimitCounter) GetReferralId() int32`

GetReferralId returns the ReferralId field if non-nil, zero value otherwise.

### GetReferralIdOk

`func (o *LimitCounter) GetReferralIdOk() (*int32, bool)`

GetReferralIdOk returns a tuple with the ReferralId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralId

`func (o *LimitCounter) SetReferralId(v int32)`

SetReferralId sets ReferralId field to given value.

### HasReferralId

`func (o *LimitCounter) HasReferralId() bool`

HasReferralId returns a boolean if a field has been set.

### GetReferralValue

`func (o *LimitCounter) GetReferralValue() string`

GetReferralValue returns the ReferralValue field if non-nil, zero value otherwise.

### GetReferralValueOk

`func (o *LimitCounter) GetReferralValueOk() (*string, bool)`

GetReferralValueOk returns a tuple with the ReferralValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralValue

`func (o *LimitCounter) SetReferralValue(v string)`

SetReferralValue sets ReferralValue field to given value.

### HasReferralValue

`func (o *LimitCounter) HasReferralValue() bool`

HasReferralValue returns a boolean if a field has been set.

### GetIdentifier

`func (o *LimitCounter) GetIdentifier() int32`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *LimitCounter) GetIdentifierOk() (*int32, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *LimitCounter) SetIdentifier(v int32)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *LimitCounter) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetPeriod

`func (o *LimitCounter) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *LimitCounter) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *LimitCounter) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *LimitCounter) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetLimit

`func (o *LimitCounter) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LimitCounter) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *LimitCounter) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetCounter

`func (o *LimitCounter) GetCounter() float32`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *LimitCounter) GetCounterOk() (*float32, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *LimitCounter) SetCounter(v float32)`

SetCounter sets Counter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


