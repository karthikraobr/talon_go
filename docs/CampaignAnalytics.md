# CampaignAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **time.Time** |  | 
**CampaignRevenue** | **float32** | Amount of revenue in this campaign (for coupon or discount sessions). | 
**TotalCampaignRevenue** | **float32** | Amount of revenue in this campaign since it began (for coupon or discount sessions). | 
**CampaignRefund** | **float32** | Amount of refunds in this campaign (for coupon or discount sessions). | 
**TotalCampaignRefund** | **float32** | Amount of refunds in this campaign since it began (for coupon or discount sessions). | 
**CampaignDiscountCosts** | **float32** | Amount of cost caused by discounts given in the campaign. | 
**TotalCampaignDiscountCosts** | **float32** | Amount of cost caused by discounts given in the campaign since it began. | 
**CampaignRefundedDiscounts** | **float32** | Amount of discounts rolledback due to refund in the campaign. | 
**TotalCampaignRefundedDiscounts** | **float32** | Amount of discounts rolledback due to refund in the campaign since it began. | 
**CampaignFreeItems** | **int32** | Amount of free items given in the campaign. | 
**TotalCampaignFreeItems** | **int32** | Amount of free items given in the campaign since it began. | 
**CouponRedemptions** | **int32** | Number of coupon redemptions in the campaign. | 
**TotalCouponRedemptions** | **int32** | Number of coupon redemptions in the campaign since it began. | 
**CouponRolledbackRedemptions** | **int32** | Number of coupon redemptions that have been rolled back (due to canceling closed session) in the campaign. | 
**TotalCouponRolledbackRedemptions** | **int32** | Number of coupon redemptions that have been rolled back (due to canceling closed session) in the campaign since it began. | 
**ReferralRedemptions** | **int32** | Number of referral redemptions in the campaign. | 
**TotalReferralRedemptions** | **int32** | Number of referral redemptions in the campaign since it began. | 
**CouponsCreated** | **int32** | Number of coupons created in the campaign by the rule engine. | 
**TotalCouponsCreated** | **int32** | Number of coupons created in the campaign by the rule engine since it began. | 
**ReferralsCreated** | **int32** | Number of referrals created in the campaign by the rule engine. | 
**TotalReferralsCreated** | **int32** | Number of referrals created in the campaign by the rule engine since it began. | 
**AddedLoyaltyPoints** | **float32** | Number of added loyalty points in the campaign in a specific interval. | 
**TotalAddedLoyaltyPoints** | **float32** | Number of added loyalty points in the campaign since it began. | 
**DeductedLoyaltyPoints** | **float32** | Number of deducted loyalty points in the campaign in a specific interval. | 
**TotalDeductedLoyaltyPoints** | **float32** | Number of deducted loyalty points in the campaign since it began. | 

## Methods

### NewCampaignAnalytics

`func NewCampaignAnalytics(date time.Time, campaignRevenue float32, totalCampaignRevenue float32, campaignRefund float32, totalCampaignRefund float32, campaignDiscountCosts float32, totalCampaignDiscountCosts float32, campaignRefundedDiscounts float32, totalCampaignRefundedDiscounts float32, campaignFreeItems int32, totalCampaignFreeItems int32, couponRedemptions int32, totalCouponRedemptions int32, couponRolledbackRedemptions int32, totalCouponRolledbackRedemptions int32, referralRedemptions int32, totalReferralRedemptions int32, couponsCreated int32, totalCouponsCreated int32, referralsCreated int32, totalReferralsCreated int32, addedLoyaltyPoints float32, totalAddedLoyaltyPoints float32, deductedLoyaltyPoints float32, totalDeductedLoyaltyPoints float32, ) *CampaignAnalytics`

NewCampaignAnalytics instantiates a new CampaignAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignAnalyticsWithDefaults

`func NewCampaignAnalyticsWithDefaults() *CampaignAnalytics`

NewCampaignAnalyticsWithDefaults instantiates a new CampaignAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CampaignAnalytics) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CampaignAnalytics) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CampaignAnalytics) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetCampaignRevenue

`func (o *CampaignAnalytics) GetCampaignRevenue() float32`

GetCampaignRevenue returns the CampaignRevenue field if non-nil, zero value otherwise.

### GetCampaignRevenueOk

`func (o *CampaignAnalytics) GetCampaignRevenueOk() (*float32, bool)`

GetCampaignRevenueOk returns a tuple with the CampaignRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignRevenue

`func (o *CampaignAnalytics) SetCampaignRevenue(v float32)`

SetCampaignRevenue sets CampaignRevenue field to given value.


### GetTotalCampaignRevenue

`func (o *CampaignAnalytics) GetTotalCampaignRevenue() float32`

GetTotalCampaignRevenue returns the TotalCampaignRevenue field if non-nil, zero value otherwise.

### GetTotalCampaignRevenueOk

`func (o *CampaignAnalytics) GetTotalCampaignRevenueOk() (*float32, bool)`

GetTotalCampaignRevenueOk returns a tuple with the TotalCampaignRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCampaignRevenue

`func (o *CampaignAnalytics) SetTotalCampaignRevenue(v float32)`

SetTotalCampaignRevenue sets TotalCampaignRevenue field to given value.


### GetCampaignRefund

`func (o *CampaignAnalytics) GetCampaignRefund() float32`

GetCampaignRefund returns the CampaignRefund field if non-nil, zero value otherwise.

### GetCampaignRefundOk

`func (o *CampaignAnalytics) GetCampaignRefundOk() (*float32, bool)`

GetCampaignRefundOk returns a tuple with the CampaignRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignRefund

`func (o *CampaignAnalytics) SetCampaignRefund(v float32)`

SetCampaignRefund sets CampaignRefund field to given value.


### GetTotalCampaignRefund

`func (o *CampaignAnalytics) GetTotalCampaignRefund() float32`

GetTotalCampaignRefund returns the TotalCampaignRefund field if non-nil, zero value otherwise.

### GetTotalCampaignRefundOk

`func (o *CampaignAnalytics) GetTotalCampaignRefundOk() (*float32, bool)`

GetTotalCampaignRefundOk returns a tuple with the TotalCampaignRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCampaignRefund

`func (o *CampaignAnalytics) SetTotalCampaignRefund(v float32)`

SetTotalCampaignRefund sets TotalCampaignRefund field to given value.


### GetCampaignDiscountCosts

`func (o *CampaignAnalytics) GetCampaignDiscountCosts() float32`

GetCampaignDiscountCosts returns the CampaignDiscountCosts field if non-nil, zero value otherwise.

### GetCampaignDiscountCostsOk

`func (o *CampaignAnalytics) GetCampaignDiscountCostsOk() (*float32, bool)`

GetCampaignDiscountCostsOk returns a tuple with the CampaignDiscountCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignDiscountCosts

`func (o *CampaignAnalytics) SetCampaignDiscountCosts(v float32)`

SetCampaignDiscountCosts sets CampaignDiscountCosts field to given value.


### GetTotalCampaignDiscountCosts

`func (o *CampaignAnalytics) GetTotalCampaignDiscountCosts() float32`

GetTotalCampaignDiscountCosts returns the TotalCampaignDiscountCosts field if non-nil, zero value otherwise.

### GetTotalCampaignDiscountCostsOk

`func (o *CampaignAnalytics) GetTotalCampaignDiscountCostsOk() (*float32, bool)`

GetTotalCampaignDiscountCostsOk returns a tuple with the TotalCampaignDiscountCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCampaignDiscountCosts

`func (o *CampaignAnalytics) SetTotalCampaignDiscountCosts(v float32)`

SetTotalCampaignDiscountCosts sets TotalCampaignDiscountCosts field to given value.


### GetCampaignRefundedDiscounts

`func (o *CampaignAnalytics) GetCampaignRefundedDiscounts() float32`

GetCampaignRefundedDiscounts returns the CampaignRefundedDiscounts field if non-nil, zero value otherwise.

### GetCampaignRefundedDiscountsOk

`func (o *CampaignAnalytics) GetCampaignRefundedDiscountsOk() (*float32, bool)`

GetCampaignRefundedDiscountsOk returns a tuple with the CampaignRefundedDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignRefundedDiscounts

`func (o *CampaignAnalytics) SetCampaignRefundedDiscounts(v float32)`

SetCampaignRefundedDiscounts sets CampaignRefundedDiscounts field to given value.


### GetTotalCampaignRefundedDiscounts

`func (o *CampaignAnalytics) GetTotalCampaignRefundedDiscounts() float32`

GetTotalCampaignRefundedDiscounts returns the TotalCampaignRefundedDiscounts field if non-nil, zero value otherwise.

### GetTotalCampaignRefundedDiscountsOk

`func (o *CampaignAnalytics) GetTotalCampaignRefundedDiscountsOk() (*float32, bool)`

GetTotalCampaignRefundedDiscountsOk returns a tuple with the TotalCampaignRefundedDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCampaignRefundedDiscounts

`func (o *CampaignAnalytics) SetTotalCampaignRefundedDiscounts(v float32)`

SetTotalCampaignRefundedDiscounts sets TotalCampaignRefundedDiscounts field to given value.


### GetCampaignFreeItems

`func (o *CampaignAnalytics) GetCampaignFreeItems() int32`

GetCampaignFreeItems returns the CampaignFreeItems field if non-nil, zero value otherwise.

### GetCampaignFreeItemsOk

`func (o *CampaignAnalytics) GetCampaignFreeItemsOk() (*int32, bool)`

GetCampaignFreeItemsOk returns a tuple with the CampaignFreeItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignFreeItems

`func (o *CampaignAnalytics) SetCampaignFreeItems(v int32)`

SetCampaignFreeItems sets CampaignFreeItems field to given value.


### GetTotalCampaignFreeItems

`func (o *CampaignAnalytics) GetTotalCampaignFreeItems() int32`

GetTotalCampaignFreeItems returns the TotalCampaignFreeItems field if non-nil, zero value otherwise.

### GetTotalCampaignFreeItemsOk

`func (o *CampaignAnalytics) GetTotalCampaignFreeItemsOk() (*int32, bool)`

GetTotalCampaignFreeItemsOk returns a tuple with the TotalCampaignFreeItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCampaignFreeItems

`func (o *CampaignAnalytics) SetTotalCampaignFreeItems(v int32)`

SetTotalCampaignFreeItems sets TotalCampaignFreeItems field to given value.


### GetCouponRedemptions

`func (o *CampaignAnalytics) GetCouponRedemptions() int32`

GetCouponRedemptions returns the CouponRedemptions field if non-nil, zero value otherwise.

### GetCouponRedemptionsOk

`func (o *CampaignAnalytics) GetCouponRedemptionsOk() (*int32, bool)`

GetCouponRedemptionsOk returns a tuple with the CouponRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponRedemptions

`func (o *CampaignAnalytics) SetCouponRedemptions(v int32)`

SetCouponRedemptions sets CouponRedemptions field to given value.


### GetTotalCouponRedemptions

`func (o *CampaignAnalytics) GetTotalCouponRedemptions() int32`

GetTotalCouponRedemptions returns the TotalCouponRedemptions field if non-nil, zero value otherwise.

### GetTotalCouponRedemptionsOk

`func (o *CampaignAnalytics) GetTotalCouponRedemptionsOk() (*int32, bool)`

GetTotalCouponRedemptionsOk returns a tuple with the TotalCouponRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCouponRedemptions

`func (o *CampaignAnalytics) SetTotalCouponRedemptions(v int32)`

SetTotalCouponRedemptions sets TotalCouponRedemptions field to given value.


### GetCouponRolledbackRedemptions

`func (o *CampaignAnalytics) GetCouponRolledbackRedemptions() int32`

GetCouponRolledbackRedemptions returns the CouponRolledbackRedemptions field if non-nil, zero value otherwise.

### GetCouponRolledbackRedemptionsOk

`func (o *CampaignAnalytics) GetCouponRolledbackRedemptionsOk() (*int32, bool)`

GetCouponRolledbackRedemptionsOk returns a tuple with the CouponRolledbackRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponRolledbackRedemptions

`func (o *CampaignAnalytics) SetCouponRolledbackRedemptions(v int32)`

SetCouponRolledbackRedemptions sets CouponRolledbackRedemptions field to given value.


### GetTotalCouponRolledbackRedemptions

`func (o *CampaignAnalytics) GetTotalCouponRolledbackRedemptions() int32`

GetTotalCouponRolledbackRedemptions returns the TotalCouponRolledbackRedemptions field if non-nil, zero value otherwise.

### GetTotalCouponRolledbackRedemptionsOk

`func (o *CampaignAnalytics) GetTotalCouponRolledbackRedemptionsOk() (*int32, bool)`

GetTotalCouponRolledbackRedemptionsOk returns a tuple with the TotalCouponRolledbackRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCouponRolledbackRedemptions

`func (o *CampaignAnalytics) SetTotalCouponRolledbackRedemptions(v int32)`

SetTotalCouponRolledbackRedemptions sets TotalCouponRolledbackRedemptions field to given value.


### GetReferralRedemptions

`func (o *CampaignAnalytics) GetReferralRedemptions() int32`

GetReferralRedemptions returns the ReferralRedemptions field if non-nil, zero value otherwise.

### GetReferralRedemptionsOk

`func (o *CampaignAnalytics) GetReferralRedemptionsOk() (*int32, bool)`

GetReferralRedemptionsOk returns a tuple with the ReferralRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralRedemptions

`func (o *CampaignAnalytics) SetReferralRedemptions(v int32)`

SetReferralRedemptions sets ReferralRedemptions field to given value.


### GetTotalReferralRedemptions

`func (o *CampaignAnalytics) GetTotalReferralRedemptions() int32`

GetTotalReferralRedemptions returns the TotalReferralRedemptions field if non-nil, zero value otherwise.

### GetTotalReferralRedemptionsOk

`func (o *CampaignAnalytics) GetTotalReferralRedemptionsOk() (*int32, bool)`

GetTotalReferralRedemptionsOk returns a tuple with the TotalReferralRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReferralRedemptions

`func (o *CampaignAnalytics) SetTotalReferralRedemptions(v int32)`

SetTotalReferralRedemptions sets TotalReferralRedemptions field to given value.


### GetCouponsCreated

`func (o *CampaignAnalytics) GetCouponsCreated() int32`

GetCouponsCreated returns the CouponsCreated field if non-nil, zero value otherwise.

### GetCouponsCreatedOk

`func (o *CampaignAnalytics) GetCouponsCreatedOk() (*int32, bool)`

GetCouponsCreatedOk returns a tuple with the CouponsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponsCreated

`func (o *CampaignAnalytics) SetCouponsCreated(v int32)`

SetCouponsCreated sets CouponsCreated field to given value.


### GetTotalCouponsCreated

`func (o *CampaignAnalytics) GetTotalCouponsCreated() int32`

GetTotalCouponsCreated returns the TotalCouponsCreated field if non-nil, zero value otherwise.

### GetTotalCouponsCreatedOk

`func (o *CampaignAnalytics) GetTotalCouponsCreatedOk() (*int32, bool)`

GetTotalCouponsCreatedOk returns a tuple with the TotalCouponsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCouponsCreated

`func (o *CampaignAnalytics) SetTotalCouponsCreated(v int32)`

SetTotalCouponsCreated sets TotalCouponsCreated field to given value.


### GetReferralsCreated

`func (o *CampaignAnalytics) GetReferralsCreated() int32`

GetReferralsCreated returns the ReferralsCreated field if non-nil, zero value otherwise.

### GetReferralsCreatedOk

`func (o *CampaignAnalytics) GetReferralsCreatedOk() (*int32, bool)`

GetReferralsCreatedOk returns a tuple with the ReferralsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralsCreated

`func (o *CampaignAnalytics) SetReferralsCreated(v int32)`

SetReferralsCreated sets ReferralsCreated field to given value.


### GetTotalReferralsCreated

`func (o *CampaignAnalytics) GetTotalReferralsCreated() int32`

GetTotalReferralsCreated returns the TotalReferralsCreated field if non-nil, zero value otherwise.

### GetTotalReferralsCreatedOk

`func (o *CampaignAnalytics) GetTotalReferralsCreatedOk() (*int32, bool)`

GetTotalReferralsCreatedOk returns a tuple with the TotalReferralsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReferralsCreated

`func (o *CampaignAnalytics) SetTotalReferralsCreated(v int32)`

SetTotalReferralsCreated sets TotalReferralsCreated field to given value.


### GetAddedLoyaltyPoints

`func (o *CampaignAnalytics) GetAddedLoyaltyPoints() float32`

GetAddedLoyaltyPoints returns the AddedLoyaltyPoints field if non-nil, zero value otherwise.

### GetAddedLoyaltyPointsOk

`func (o *CampaignAnalytics) GetAddedLoyaltyPointsOk() (*float32, bool)`

GetAddedLoyaltyPointsOk returns a tuple with the AddedLoyaltyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedLoyaltyPoints

`func (o *CampaignAnalytics) SetAddedLoyaltyPoints(v float32)`

SetAddedLoyaltyPoints sets AddedLoyaltyPoints field to given value.


### GetTotalAddedLoyaltyPoints

`func (o *CampaignAnalytics) GetTotalAddedLoyaltyPoints() float32`

GetTotalAddedLoyaltyPoints returns the TotalAddedLoyaltyPoints field if non-nil, zero value otherwise.

### GetTotalAddedLoyaltyPointsOk

`func (o *CampaignAnalytics) GetTotalAddedLoyaltyPointsOk() (*float32, bool)`

GetTotalAddedLoyaltyPointsOk returns a tuple with the TotalAddedLoyaltyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAddedLoyaltyPoints

`func (o *CampaignAnalytics) SetTotalAddedLoyaltyPoints(v float32)`

SetTotalAddedLoyaltyPoints sets TotalAddedLoyaltyPoints field to given value.


### GetDeductedLoyaltyPoints

`func (o *CampaignAnalytics) GetDeductedLoyaltyPoints() float32`

GetDeductedLoyaltyPoints returns the DeductedLoyaltyPoints field if non-nil, zero value otherwise.

### GetDeductedLoyaltyPointsOk

`func (o *CampaignAnalytics) GetDeductedLoyaltyPointsOk() (*float32, bool)`

GetDeductedLoyaltyPointsOk returns a tuple with the DeductedLoyaltyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeductedLoyaltyPoints

`func (o *CampaignAnalytics) SetDeductedLoyaltyPoints(v float32)`

SetDeductedLoyaltyPoints sets DeductedLoyaltyPoints field to given value.


### GetTotalDeductedLoyaltyPoints

`func (o *CampaignAnalytics) GetTotalDeductedLoyaltyPoints() float32`

GetTotalDeductedLoyaltyPoints returns the TotalDeductedLoyaltyPoints field if non-nil, zero value otherwise.

### GetTotalDeductedLoyaltyPointsOk

`func (o *CampaignAnalytics) GetTotalDeductedLoyaltyPointsOk() (*float32, bool)`

GetTotalDeductedLoyaltyPointsOk returns a tuple with the TotalDeductedLoyaltyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDeductedLoyaltyPoints

`func (o *CampaignAnalytics) SetTotalDeductedLoyaltyPoints(v float32)`

SetTotalDeductedLoyaltyPoints sets TotalDeductedLoyaltyPoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


