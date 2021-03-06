# Meta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Campaigns** | Pointer to **map[string]interface{}** | Maps each evaluated campaign ID to a key-value list of that campaigns attributes. Campaigns without attributes will be omitted. | [optional] 
**Coupons** | Pointer to **map[string]interface{}** | Maps the coupon value to a key-value list of that coupons attributes. | [optional] 
**CouponRejectionReason** | Pointer to [**CouponRejectionReason**](CouponRejectionReason.md) |  | [optional] 
**ReferralRejectionReason** | Pointer to [**ReferralRejectionReason**](ReferralRejectionReason.md) |  | [optional] 
**Warnings** | Pointer to **map[string]interface{}** | Contains warnings about possible misuse. | [optional] 

## Methods

### NewMeta

`func NewMeta() *Meta`

NewMeta instantiates a new Meta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaWithDefaults

`func NewMetaWithDefaults() *Meta`

NewMetaWithDefaults instantiates a new Meta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaigns

`func (o *Meta) GetCampaigns() map[string]interface{}`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *Meta) GetCampaignsOk() (*map[string]interface{}, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *Meta) SetCampaigns(v map[string]interface{})`

SetCampaigns sets Campaigns field to given value.

### HasCampaigns

`func (o *Meta) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### GetCoupons

`func (o *Meta) GetCoupons() map[string]interface{}`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *Meta) GetCouponsOk() (*map[string]interface{}, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupons

`func (o *Meta) SetCoupons(v map[string]interface{})`

SetCoupons sets Coupons field to given value.

### HasCoupons

`func (o *Meta) HasCoupons() bool`

HasCoupons returns a boolean if a field has been set.

### GetCouponRejectionReason

`func (o *Meta) GetCouponRejectionReason() CouponRejectionReason`

GetCouponRejectionReason returns the CouponRejectionReason field if non-nil, zero value otherwise.

### GetCouponRejectionReasonOk

`func (o *Meta) GetCouponRejectionReasonOk() (*CouponRejectionReason, bool)`

GetCouponRejectionReasonOk returns a tuple with the CouponRejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponRejectionReason

`func (o *Meta) SetCouponRejectionReason(v CouponRejectionReason)`

SetCouponRejectionReason sets CouponRejectionReason field to given value.

### HasCouponRejectionReason

`func (o *Meta) HasCouponRejectionReason() bool`

HasCouponRejectionReason returns a boolean if a field has been set.

### GetReferralRejectionReason

`func (o *Meta) GetReferralRejectionReason() ReferralRejectionReason`

GetReferralRejectionReason returns the ReferralRejectionReason field if non-nil, zero value otherwise.

### GetReferralRejectionReasonOk

`func (o *Meta) GetReferralRejectionReasonOk() (*ReferralRejectionReason, bool)`

GetReferralRejectionReasonOk returns a tuple with the ReferralRejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralRejectionReason

`func (o *Meta) SetReferralRejectionReason(v ReferralRejectionReason)`

SetReferralRejectionReason sets ReferralRejectionReason field to given value.

### HasReferralRejectionReason

`func (o *Meta) HasReferralRejectionReason() bool`

HasReferralRejectionReason returns a boolean if a field has been set.

### GetWarnings

`func (o *Meta) GetWarnings() map[string]interface{}`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Meta) GetWarningsOk() (*map[string]interface{}, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Meta) SetWarnings(v map[string]interface{})`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *Meta) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


