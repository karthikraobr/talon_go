/*
Integration API reference docs

Use the Integration API to push data to and retrieve data from Talon.One in real time. For more background information about this API, see [Integration API overview](/docs/dev/integration-api/overview)  For example, use this API to share shopping cart information as a session with Talon.One and evaluate promotion rules. You can also create custom events to track specific actions that do not fit into the session data model.  Ensure you [authenticate](#section/Authentication) to make requests to the API.  <div class=\"redoc-section\">   <p class=\"title\">Are you looking for a different API?</p>    If you need the API to:    - Interact with the Campaign Manager for backoffice operations, see [the Management API reference docs](https://docs.talon.one/management-api).   - Integrate with Talon.One from a CEP or CDP platform, see [the Third-party API reference docs](https://docs.talon.one/third-party-api).  </div>  # Authentication  <SecurityDefinitions /> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package talon

import (
	"encoding/json"
)

// CustomerAnalytics 
type CustomerAnalytics struct {
	// Total accepted coupons for this customer.
	AcceptedCoupons int32 `json:"acceptedCoupons"`
	// Total created coupons for this customer.
	CreatedCoupons int32 `json:"createdCoupons"`
	// Total free items given to this customer.
	FreeItems int32 `json:"freeItems"`
	// Total orders made by this customer.
	TotalOrders int32 `json:"totalOrders"`
	// Total orders made by this customer that had a discount.
	TotalDiscountedOrders int32 `json:"totalDiscountedOrders"`
	// Total Revenue across all closed sessions.
	TotalRevenue float32 `json:"totalRevenue"`
	// The sum of discounts that were given across all closed sessions.
	TotalDiscounts float32 `json:"totalDiscounts"`
}

// NewCustomerAnalytics instantiates a new CustomerAnalytics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAnalytics(acceptedCoupons int32, createdCoupons int32, freeItems int32, totalOrders int32, totalDiscountedOrders int32, totalRevenue float32, totalDiscounts float32) *CustomerAnalytics {
	this := CustomerAnalytics{}
	this.AcceptedCoupons = acceptedCoupons
	this.CreatedCoupons = createdCoupons
	this.FreeItems = freeItems
	this.TotalOrders = totalOrders
	this.TotalDiscountedOrders = totalDiscountedOrders
	this.TotalRevenue = totalRevenue
	this.TotalDiscounts = totalDiscounts
	return &this
}

// NewCustomerAnalyticsWithDefaults instantiates a new CustomerAnalytics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAnalyticsWithDefaults() *CustomerAnalytics {
	this := CustomerAnalytics{}
	return &this
}

// GetAcceptedCoupons returns the AcceptedCoupons field value
func (o *CustomerAnalytics) GetAcceptedCoupons() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AcceptedCoupons
}

// GetAcceptedCouponsOk returns a tuple with the AcceptedCoupons field value
// and a boolean to check if the value has been set.
func (o *CustomerAnalytics) GetAcceptedCouponsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptedCoupons, true
}

// SetAcceptedCoupons sets field value
func (o *CustomerAnalytics) SetAcceptedCoupons(v int32) {
	o.AcceptedCoupons = v
}

// GetCreatedCoupons returns the CreatedCoupons field value
func (o *CustomerAnalytics) GetCreatedCoupons() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedCoupons
}

// GetCreatedCouponsOk returns a tuple with the CreatedCoupons field value
// and a boolean to check if the value has been set.
func (o *CustomerAnalytics) GetCreatedCouponsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedCoupons, true
}

// SetCreatedCoupons sets field value
func (o *CustomerAnalytics) SetCreatedCoupons(v int32) {
	o.CreatedCoupons = v
}

// GetFreeItems returns the FreeItems field value
func (o *CustomerAnalytics) GetFreeItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FreeItems
}

// GetFreeItemsOk returns a tuple with the FreeItems field value
// and a boolean to check if the value has been set.
func (o *CustomerAnalytics) GetFreeItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FreeItems, true
}

// SetFreeItems sets field value
func (o *CustomerAnalytics) SetFreeItems(v int32) {
	o.FreeItems = v
}

// GetTotalOrders returns the TotalOrders field value
func (o *CustomerAnalytics) GetTotalOrders() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalOrders
}

// GetTotalOrdersOk returns a tuple with the TotalOrders field value
// and a boolean to check if the value has been set.
func (o *CustomerAnalytics) GetTotalOrdersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalOrders, true
}

// SetTotalOrders sets field value
func (o *CustomerAnalytics) SetTotalOrders(v int32) {
	o.TotalOrders = v
}

// GetTotalDiscountedOrders returns the TotalDiscountedOrders field value
func (o *CustomerAnalytics) GetTotalDiscountedOrders() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalDiscountedOrders
}

// GetTotalDiscountedOrdersOk returns a tuple with the TotalDiscountedOrders field value
// and a boolean to check if the value has been set.
func (o *CustomerAnalytics) GetTotalDiscountedOrdersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalDiscountedOrders, true
}

// SetTotalDiscountedOrders sets field value
func (o *CustomerAnalytics) SetTotalDiscountedOrders(v int32) {
	o.TotalDiscountedOrders = v
}

// GetTotalRevenue returns the TotalRevenue field value
func (o *CustomerAnalytics) GetTotalRevenue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalRevenue
}

// GetTotalRevenueOk returns a tuple with the TotalRevenue field value
// and a boolean to check if the value has been set.
func (o *CustomerAnalytics) GetTotalRevenueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRevenue, true
}

// SetTotalRevenue sets field value
func (o *CustomerAnalytics) SetTotalRevenue(v float32) {
	o.TotalRevenue = v
}

// GetTotalDiscounts returns the TotalDiscounts field value
func (o *CustomerAnalytics) GetTotalDiscounts() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalDiscounts
}

// GetTotalDiscountsOk returns a tuple with the TotalDiscounts field value
// and a boolean to check if the value has been set.
func (o *CustomerAnalytics) GetTotalDiscountsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalDiscounts, true
}

// SetTotalDiscounts sets field value
func (o *CustomerAnalytics) SetTotalDiscounts(v float32) {
	o.TotalDiscounts = v
}

func (o CustomerAnalytics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["acceptedCoupons"] = o.AcceptedCoupons
	}
	if true {
		toSerialize["createdCoupons"] = o.CreatedCoupons
	}
	if true {
		toSerialize["freeItems"] = o.FreeItems
	}
	if true {
		toSerialize["totalOrders"] = o.TotalOrders
	}
	if true {
		toSerialize["totalDiscountedOrders"] = o.TotalDiscountedOrders
	}
	if true {
		toSerialize["totalRevenue"] = o.TotalRevenue
	}
	if true {
		toSerialize["totalDiscounts"] = o.TotalDiscounts
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerAnalytics struct {
	value *CustomerAnalytics
	isSet bool
}

func (v NullableCustomerAnalytics) Get() *CustomerAnalytics {
	return v.value
}

func (v *NullableCustomerAnalytics) Set(val *CustomerAnalytics) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAnalytics) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAnalytics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAnalytics(val *CustomerAnalytics) *NullableCustomerAnalytics {
	return &NullableCustomerAnalytics{value: val, isSet: true}
}

func (v NullableCustomerAnalytics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAnalytics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


