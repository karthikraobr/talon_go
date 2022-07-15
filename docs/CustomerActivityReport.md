# CustomerActivityReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | **string** | The integration ID set by your integration layer. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**Name** | **string** | The name for this customer profile. | 
**CustomerId** | **int32** | The internal Talon.One ID of the customer. | 
**LastActivity** | Pointer to **time.Time** | The last activity of the customer. | [optional] 
**CouponRedemptions** | **int32** | Number of coupon redemptions in all customer campaigns. | 
**CouponUseAttempts** | **int32** | Number of coupon use attempts in all customer campaigns. | 
**CouponFailedAttempts** | **int32** | Number of failed coupon use attempts in all customer campaigns. | 
**AccruedDiscounts** | **float32** | Number of accrued discounts in all customer campaigns. | 
**AccruedRevenue** | **float32** | Amount of accrued revenue in all customer campaigns. | 
**TotalOrders** | **int32** | Number of orders in all customer campaigns. | 
**TotalOrdersNoCoupon** | **int32** | Number of orders without coupon used in all customer campaigns. | 
**CampaignName** | **string** | The name of the campaign this customer belongs to. | 

## Methods

### NewCustomerActivityReport

`func NewCustomerActivityReport(integrationId string, created time.Time, name string, customerId int32, couponRedemptions int32, couponUseAttempts int32, couponFailedAttempts int32, accruedDiscounts float32, accruedRevenue float32, totalOrders int32, totalOrdersNoCoupon int32, campaignName string, ) *CustomerActivityReport`

NewCustomerActivityReport instantiates a new CustomerActivityReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerActivityReportWithDefaults

`func NewCustomerActivityReportWithDefaults() *CustomerActivityReport`

NewCustomerActivityReportWithDefaults instantiates a new CustomerActivityReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationId

`func (o *CustomerActivityReport) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CustomerActivityReport) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *CustomerActivityReport) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetCreated

`func (o *CustomerActivityReport) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomerActivityReport) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomerActivityReport) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetName

`func (o *CustomerActivityReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerActivityReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerActivityReport) SetName(v string)`

SetName sets Name field to given value.


### GetCustomerId

`func (o *CustomerActivityReport) GetCustomerId() int32`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerActivityReport) GetCustomerIdOk() (*int32, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerActivityReport) SetCustomerId(v int32)`

SetCustomerId sets CustomerId field to given value.


### GetLastActivity

`func (o *CustomerActivityReport) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *CustomerActivityReport) GetLastActivityOk() (*time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *CustomerActivityReport) SetLastActivity(v time.Time)`

SetLastActivity sets LastActivity field to given value.

### HasLastActivity

`func (o *CustomerActivityReport) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### GetCouponRedemptions

`func (o *CustomerActivityReport) GetCouponRedemptions() int32`

GetCouponRedemptions returns the CouponRedemptions field if non-nil, zero value otherwise.

### GetCouponRedemptionsOk

`func (o *CustomerActivityReport) GetCouponRedemptionsOk() (*int32, bool)`

GetCouponRedemptionsOk returns a tuple with the CouponRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponRedemptions

`func (o *CustomerActivityReport) SetCouponRedemptions(v int32)`

SetCouponRedemptions sets CouponRedemptions field to given value.


### GetCouponUseAttempts

`func (o *CustomerActivityReport) GetCouponUseAttempts() int32`

GetCouponUseAttempts returns the CouponUseAttempts field if non-nil, zero value otherwise.

### GetCouponUseAttemptsOk

`func (o *CustomerActivityReport) GetCouponUseAttemptsOk() (*int32, bool)`

GetCouponUseAttemptsOk returns a tuple with the CouponUseAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponUseAttempts

`func (o *CustomerActivityReport) SetCouponUseAttempts(v int32)`

SetCouponUseAttempts sets CouponUseAttempts field to given value.


### GetCouponFailedAttempts

`func (o *CustomerActivityReport) GetCouponFailedAttempts() int32`

GetCouponFailedAttempts returns the CouponFailedAttempts field if non-nil, zero value otherwise.

### GetCouponFailedAttemptsOk

`func (o *CustomerActivityReport) GetCouponFailedAttemptsOk() (*int32, bool)`

GetCouponFailedAttemptsOk returns a tuple with the CouponFailedAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponFailedAttempts

`func (o *CustomerActivityReport) SetCouponFailedAttempts(v int32)`

SetCouponFailedAttempts sets CouponFailedAttempts field to given value.


### GetAccruedDiscounts

`func (o *CustomerActivityReport) GetAccruedDiscounts() float32`

GetAccruedDiscounts returns the AccruedDiscounts field if non-nil, zero value otherwise.

### GetAccruedDiscountsOk

`func (o *CustomerActivityReport) GetAccruedDiscountsOk() (*float32, bool)`

GetAccruedDiscountsOk returns a tuple with the AccruedDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedDiscounts

`func (o *CustomerActivityReport) SetAccruedDiscounts(v float32)`

SetAccruedDiscounts sets AccruedDiscounts field to given value.


### GetAccruedRevenue

`func (o *CustomerActivityReport) GetAccruedRevenue() float32`

GetAccruedRevenue returns the AccruedRevenue field if non-nil, zero value otherwise.

### GetAccruedRevenueOk

`func (o *CustomerActivityReport) GetAccruedRevenueOk() (*float32, bool)`

GetAccruedRevenueOk returns a tuple with the AccruedRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedRevenue

`func (o *CustomerActivityReport) SetAccruedRevenue(v float32)`

SetAccruedRevenue sets AccruedRevenue field to given value.


### GetTotalOrders

`func (o *CustomerActivityReport) GetTotalOrders() int32`

GetTotalOrders returns the TotalOrders field if non-nil, zero value otherwise.

### GetTotalOrdersOk

`func (o *CustomerActivityReport) GetTotalOrdersOk() (*int32, bool)`

GetTotalOrdersOk returns a tuple with the TotalOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOrders

`func (o *CustomerActivityReport) SetTotalOrders(v int32)`

SetTotalOrders sets TotalOrders field to given value.


### GetTotalOrdersNoCoupon

`func (o *CustomerActivityReport) GetTotalOrdersNoCoupon() int32`

GetTotalOrdersNoCoupon returns the TotalOrdersNoCoupon field if non-nil, zero value otherwise.

### GetTotalOrdersNoCouponOk

`func (o *CustomerActivityReport) GetTotalOrdersNoCouponOk() (*int32, bool)`

GetTotalOrdersNoCouponOk returns a tuple with the TotalOrdersNoCoupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOrdersNoCoupon

`func (o *CustomerActivityReport) SetTotalOrdersNoCoupon(v int32)`

SetTotalOrdersNoCoupon sets TotalOrdersNoCoupon field to given value.


### GetCampaignName

`func (o *CustomerActivityReport) GetCampaignName() string`

GetCampaignName returns the CampaignName field if non-nil, zero value otherwise.

### GetCampaignNameOk

`func (o *CustomerActivityReport) GetCampaignNameOk() (*string, bool)`

GetCampaignNameOk returns a tuple with the CampaignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignName

`func (o *CustomerActivityReport) SetCampaignName(v string)`

SetCampaignName sets CampaignName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


