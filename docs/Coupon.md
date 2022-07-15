# Coupon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**CampaignId** | **int32** | The ID of the campaign that owns this entity. | 
**Value** | **string** | The coupon code. | 
**UsageLimit** | **int32** | The number of times the coupon code can be redeemed. &#x60;0&#x60; means unlimited redemptions but any campaign usage limits will still apply.  | 
**DiscountLimit** | Pointer to **float32** | The amount of discounts that can be given with this coupon code.  | [optional] 
**StartDate** | Pointer to **time.Time** | Timestamp at which point the coupon becomes valid. | [optional] 
**ExpiryDate** | Pointer to **time.Time** | Expiry date of the coupon. Coupon never expires if this is omitted, zero, or negative. | [optional] 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | Limits configuration for a coupon. These limits will override the limits set from the campaign.  **Note:** Only usable when creating a single coupon which is not tied to a specific recipient. Only per-profile limits are allowed to be configured.  | [optional] 
**UsageCounter** | **int32** | The number of times this coupon has been successfully used. | 
**DiscountCounter** | Pointer to **float32** | The amount of discounts given on rules redeeming this coupon. Only usable if a coupon discount budget was set for this coupon. | [optional] 
**DiscountRemainder** | Pointer to **float32** | The remaining discount this coupon can give. | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | Custom attributes associated with this coupon. | [optional] 
**ReferralId** | Pointer to **int32** | The integration ID of the referring customer (if any) for whom this coupon was created as an effect. | [optional] 
**RecipientIntegrationId** | Pointer to **string** | The Integration ID of the customer that is allowed to redeem this coupon. | [optional] 
**ImportId** | Pointer to **int32** | The ID of the Import which created this coupon. | [optional] 
**Reservation** | Pointer to **bool** | Defines the type of reservation: - &#x60;true&#x60;: The reservation is a soft reservation. Any customer can use the coupon. This is done via the [Create coupon reservation endpoint](/integration-api/#operation/createCouponReservation). - &#x60;false&#x60;: The reservation is a hard reservation. Only the associated customer (&#x60;recipientIntegrationId&#x60;) can use the coupon. This is done via the Campaign Manager when you create a coupon for a given &#x60;recipientIntegrationId&#x60;, the [Create coupons endpoint](/management-api/#operation/createCoupons) or [Create coupons for multiple recipients endpoint](/management-api/#operation/createCouponsForMultipleRecipients).  | [optional] [default to true]
**BatchId** | Pointer to **string** | The id of the batch the coupon belongs to. | [optional] 

## Methods

### NewCoupon

`func NewCoupon(id int32, created time.Time, campaignId int32, value string, usageLimit int32, usageCounter int32, ) *Coupon`

NewCoupon instantiates a new Coupon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponWithDefaults

`func NewCouponWithDefaults() *Coupon`

NewCouponWithDefaults instantiates a new Coupon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Coupon) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Coupon) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Coupon) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *Coupon) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Coupon) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Coupon) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCampaignId

`func (o *Coupon) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *Coupon) GetCampaignIdOk() (*int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *Coupon) SetCampaignId(v int32)`

SetCampaignId sets CampaignId field to given value.


### GetValue

`func (o *Coupon) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Coupon) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Coupon) SetValue(v string)`

SetValue sets Value field to given value.


### GetUsageLimit

`func (o *Coupon) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *Coupon) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *Coupon) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.


### GetDiscountLimit

`func (o *Coupon) GetDiscountLimit() float32`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *Coupon) GetDiscountLimitOk() (*float32, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLimit

`func (o *Coupon) SetDiscountLimit(v float32)`

SetDiscountLimit sets DiscountLimit field to given value.

### HasDiscountLimit

`func (o *Coupon) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### GetStartDate

`func (o *Coupon) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Coupon) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Coupon) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Coupon) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *Coupon) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Coupon) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *Coupon) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *Coupon) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetLimits

`func (o *Coupon) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *Coupon) GetLimitsOk() (*[]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *Coupon) SetLimits(v []LimitConfig)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *Coupon) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetUsageCounter

`func (o *Coupon) GetUsageCounter() int32`

GetUsageCounter returns the UsageCounter field if non-nil, zero value otherwise.

### GetUsageCounterOk

`func (o *Coupon) GetUsageCounterOk() (*int32, bool)`

GetUsageCounterOk returns a tuple with the UsageCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCounter

`func (o *Coupon) SetUsageCounter(v int32)`

SetUsageCounter sets UsageCounter field to given value.


### GetDiscountCounter

`func (o *Coupon) GetDiscountCounter() float32`

GetDiscountCounter returns the DiscountCounter field if non-nil, zero value otherwise.

### GetDiscountCounterOk

`func (o *Coupon) GetDiscountCounterOk() (*float32, bool)`

GetDiscountCounterOk returns a tuple with the DiscountCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCounter

`func (o *Coupon) SetDiscountCounter(v float32)`

SetDiscountCounter sets DiscountCounter field to given value.

### HasDiscountCounter

`func (o *Coupon) HasDiscountCounter() bool`

HasDiscountCounter returns a boolean if a field has been set.

### GetDiscountRemainder

`func (o *Coupon) GetDiscountRemainder() float32`

GetDiscountRemainder returns the DiscountRemainder field if non-nil, zero value otherwise.

### GetDiscountRemainderOk

`func (o *Coupon) GetDiscountRemainderOk() (*float32, bool)`

GetDiscountRemainderOk returns a tuple with the DiscountRemainder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRemainder

`func (o *Coupon) SetDiscountRemainder(v float32)`

SetDiscountRemainder sets DiscountRemainder field to given value.

### HasDiscountRemainder

`func (o *Coupon) HasDiscountRemainder() bool`

HasDiscountRemainder returns a boolean if a field has been set.

### GetAttributes

`func (o *Coupon) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Coupon) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Coupon) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Coupon) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetReferralId

`func (o *Coupon) GetReferralId() int32`

GetReferralId returns the ReferralId field if non-nil, zero value otherwise.

### GetReferralIdOk

`func (o *Coupon) GetReferralIdOk() (*int32, bool)`

GetReferralIdOk returns a tuple with the ReferralId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralId

`func (o *Coupon) SetReferralId(v int32)`

SetReferralId sets ReferralId field to given value.

### HasReferralId

`func (o *Coupon) HasReferralId() bool`

HasReferralId returns a boolean if a field has been set.

### GetRecipientIntegrationId

`func (o *Coupon) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *Coupon) GetRecipientIntegrationIdOk() (*string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientIntegrationId

`func (o *Coupon) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId sets RecipientIntegrationId field to given value.

### HasRecipientIntegrationId

`func (o *Coupon) HasRecipientIntegrationId() bool`

HasRecipientIntegrationId returns a boolean if a field has been set.

### GetImportId

`func (o *Coupon) GetImportId() int32`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *Coupon) GetImportIdOk() (*int32, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportId

`func (o *Coupon) SetImportId(v int32)`

SetImportId sets ImportId field to given value.

### HasImportId

`func (o *Coupon) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### GetReservation

`func (o *Coupon) GetReservation() bool`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *Coupon) GetReservationOk() (*bool, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *Coupon) SetReservation(v bool)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *Coupon) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetBatchId

`func (o *Coupon) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *Coupon) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *Coupon) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *Coupon) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


