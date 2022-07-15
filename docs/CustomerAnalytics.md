# CustomerAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedCoupons** | **int32** | Total accepted coupons for this customer. | 
**CreatedCoupons** | **int32** | Total created coupons for this customer. | 
**FreeItems** | **int32** | Total free items given to this customer. | 
**TotalOrders** | **int32** | Total orders made by this customer. | 
**TotalDiscountedOrders** | **int32** | Total orders made by this customer that had a discount. | 
**TotalRevenue** | **float32** | Total Revenue across all closed sessions. | 
**TotalDiscounts** | **float32** | The sum of discounts that were given across all closed sessions. | 

## Methods

### NewCustomerAnalytics

`func NewCustomerAnalytics(acceptedCoupons int32, createdCoupons int32, freeItems int32, totalOrders int32, totalDiscountedOrders int32, totalRevenue float32, totalDiscounts float32, ) *CustomerAnalytics`

NewCustomerAnalytics instantiates a new CustomerAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAnalyticsWithDefaults

`func NewCustomerAnalyticsWithDefaults() *CustomerAnalytics`

NewCustomerAnalyticsWithDefaults instantiates a new CustomerAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedCoupons

`func (o *CustomerAnalytics) GetAcceptedCoupons() int32`

GetAcceptedCoupons returns the AcceptedCoupons field if non-nil, zero value otherwise.

### GetAcceptedCouponsOk

`func (o *CustomerAnalytics) GetAcceptedCouponsOk() (*int32, bool)`

GetAcceptedCouponsOk returns a tuple with the AcceptedCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCoupons

`func (o *CustomerAnalytics) SetAcceptedCoupons(v int32)`

SetAcceptedCoupons sets AcceptedCoupons field to given value.


### GetCreatedCoupons

`func (o *CustomerAnalytics) GetCreatedCoupons() int32`

GetCreatedCoupons returns the CreatedCoupons field if non-nil, zero value otherwise.

### GetCreatedCouponsOk

`func (o *CustomerAnalytics) GetCreatedCouponsOk() (*int32, bool)`

GetCreatedCouponsOk returns a tuple with the CreatedCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedCoupons

`func (o *CustomerAnalytics) SetCreatedCoupons(v int32)`

SetCreatedCoupons sets CreatedCoupons field to given value.


### GetFreeItems

`func (o *CustomerAnalytics) GetFreeItems() int32`

GetFreeItems returns the FreeItems field if non-nil, zero value otherwise.

### GetFreeItemsOk

`func (o *CustomerAnalytics) GetFreeItemsOk() (*int32, bool)`

GetFreeItemsOk returns a tuple with the FreeItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeItems

`func (o *CustomerAnalytics) SetFreeItems(v int32)`

SetFreeItems sets FreeItems field to given value.


### GetTotalOrders

`func (o *CustomerAnalytics) GetTotalOrders() int32`

GetTotalOrders returns the TotalOrders field if non-nil, zero value otherwise.

### GetTotalOrdersOk

`func (o *CustomerAnalytics) GetTotalOrdersOk() (*int32, bool)`

GetTotalOrdersOk returns a tuple with the TotalOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOrders

`func (o *CustomerAnalytics) SetTotalOrders(v int32)`

SetTotalOrders sets TotalOrders field to given value.


### GetTotalDiscountedOrders

`func (o *CustomerAnalytics) GetTotalDiscountedOrders() int32`

GetTotalDiscountedOrders returns the TotalDiscountedOrders field if non-nil, zero value otherwise.

### GetTotalDiscountedOrdersOk

`func (o *CustomerAnalytics) GetTotalDiscountedOrdersOk() (*int32, bool)`

GetTotalDiscountedOrdersOk returns a tuple with the TotalDiscountedOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountedOrders

`func (o *CustomerAnalytics) SetTotalDiscountedOrders(v int32)`

SetTotalDiscountedOrders sets TotalDiscountedOrders field to given value.


### GetTotalRevenue

`func (o *CustomerAnalytics) GetTotalRevenue() float32`

GetTotalRevenue returns the TotalRevenue field if non-nil, zero value otherwise.

### GetTotalRevenueOk

`func (o *CustomerAnalytics) GetTotalRevenueOk() (*float32, bool)`

GetTotalRevenueOk returns a tuple with the TotalRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRevenue

`func (o *CustomerAnalytics) SetTotalRevenue(v float32)`

SetTotalRevenue sets TotalRevenue field to given value.


### GetTotalDiscounts

`func (o *CustomerAnalytics) GetTotalDiscounts() float32`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *CustomerAnalytics) GetTotalDiscountsOk() (*float32, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscounts

`func (o *CustomerAnalytics) SetTotalDiscounts(v float32)`

SetTotalDiscounts sets TotalDiscounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


