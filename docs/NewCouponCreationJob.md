# NewCouponCreationJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageLimit** | **int32** | The number of times the coupon code can be redeemed. &#x60;0&#x60; means unlimited redemptions but any campaign usage limits will still apply.  | 
**DiscountLimit** | Pointer to **float32** | The amount of discounts that can be given with this coupon code.  | [optional] 
**StartDate** | Pointer to **time.Time** | Timestamp at which point the coupon becomes valid. | [optional] 
**ExpiryDate** | Pointer to **time.Time** | Expiry date of the coupon. Coupon never expires if this is omitted, zero, or negative. | [optional] 
**NumberOfCoupons** | **int32** | The number of new coupon codes to generate for the campaign. Must be between 20,001 and 5,000,000. | 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**Attributes** | **map[string]interface{}** | Arbitrary properties associated with coupons. | 

## Methods

### NewNewCouponCreationJob

`func NewNewCouponCreationJob(usageLimit int32, numberOfCoupons int32, attributes map[string]interface{}, ) *NewCouponCreationJob`

NewNewCouponCreationJob instantiates a new NewCouponCreationJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewCouponCreationJobWithDefaults

`func NewNewCouponCreationJobWithDefaults() *NewCouponCreationJob`

NewNewCouponCreationJobWithDefaults instantiates a new NewCouponCreationJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsageLimit

`func (o *NewCouponCreationJob) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *NewCouponCreationJob) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *NewCouponCreationJob) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.


### GetDiscountLimit

`func (o *NewCouponCreationJob) GetDiscountLimit() float32`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *NewCouponCreationJob) GetDiscountLimitOk() (*float32, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLimit

`func (o *NewCouponCreationJob) SetDiscountLimit(v float32)`

SetDiscountLimit sets DiscountLimit field to given value.

### HasDiscountLimit

`func (o *NewCouponCreationJob) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### GetStartDate

`func (o *NewCouponCreationJob) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *NewCouponCreationJob) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *NewCouponCreationJob) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *NewCouponCreationJob) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *NewCouponCreationJob) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *NewCouponCreationJob) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *NewCouponCreationJob) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *NewCouponCreationJob) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetNumberOfCoupons

`func (o *NewCouponCreationJob) GetNumberOfCoupons() int32`

GetNumberOfCoupons returns the NumberOfCoupons field if non-nil, zero value otherwise.

### GetNumberOfCouponsOk

`func (o *NewCouponCreationJob) GetNumberOfCouponsOk() (*int32, bool)`

GetNumberOfCouponsOk returns a tuple with the NumberOfCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCoupons

`func (o *NewCouponCreationJob) SetNumberOfCoupons(v int32)`

SetNumberOfCoupons sets NumberOfCoupons field to given value.


### GetCouponSettings

`func (o *NewCouponCreationJob) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *NewCouponCreationJob) GetCouponSettingsOk() (*CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponSettings

`func (o *NewCouponCreationJob) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings sets CouponSettings field to given value.

### HasCouponSettings

`func (o *NewCouponCreationJob) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### GetAttributes

`func (o *NewCouponCreationJob) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewCouponCreationJob) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NewCouponCreationJob) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


