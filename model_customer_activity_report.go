/*
Integration API reference docs

Use the Integration API to push data to and retrieve data from Talon.One in real time. For more background information about this API, see [Integration API overview](/docs/dev/integration-api/overview)  For example, use this API to share shopping cart information as a session with Talon.One and evaluate promotion rules. You can also create custom events to track specific actions that do not fit into the session data model.  Ensure you [authenticate](#section/Authentication) to make requests to the API.  <div class=\"redoc-section\">   <p class=\"title\">Are you looking for a different API?</p>    If you need the API to:    - Interact with the Campaign Manager for backoffice operations, see [the Management API reference docs](https://docs.talon.one/management-api).   - Integrate with Talon.One from a CEP or CDP platform, see [the Third-party API reference docs](https://docs.talon.one/third-party-api).  </div>  # Authentication  <SecurityDefinitions /> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package talon

import (
	"encoding/json"
	"time"
)

// CustomerActivityReport 
type CustomerActivityReport struct {
	// The integration ID set by your integration layer.
	IntegrationId string `json:"integrationId"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The name for this customer profile.
	Name string `json:"name"`
	// The internal Talon.One ID of the customer.
	CustomerId int32 `json:"customerId"`
	// The last activity of the customer.
	LastActivity *time.Time `json:"lastActivity,omitempty"`
	// Number of coupon redemptions in all customer campaigns.
	CouponRedemptions int32 `json:"couponRedemptions"`
	// Number of coupon use attempts in all customer campaigns.
	CouponUseAttempts int32 `json:"couponUseAttempts"`
	// Number of failed coupon use attempts in all customer campaigns.
	CouponFailedAttempts int32 `json:"couponFailedAttempts"`
	// Number of accrued discounts in all customer campaigns.
	AccruedDiscounts float32 `json:"accruedDiscounts"`
	// Amount of accrued revenue in all customer campaigns.
	AccruedRevenue float32 `json:"accruedRevenue"`
	// Number of orders in all customer campaigns.
	TotalOrders int32 `json:"totalOrders"`
	// Number of orders without coupon used in all customer campaigns.
	TotalOrdersNoCoupon int32 `json:"totalOrdersNoCoupon"`
	// The name of the campaign this customer belongs to.
	CampaignName string `json:"campaignName"`
}

// NewCustomerActivityReport instantiates a new CustomerActivityReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerActivityReport(integrationId string, created time.Time, name string, customerId int32, couponRedemptions int32, couponUseAttempts int32, couponFailedAttempts int32, accruedDiscounts float32, accruedRevenue float32, totalOrders int32, totalOrdersNoCoupon int32, campaignName string) *CustomerActivityReport {
	this := CustomerActivityReport{}
	this.IntegrationId = integrationId
	this.Created = created
	this.Name = name
	this.CustomerId = customerId
	this.CouponRedemptions = couponRedemptions
	this.CouponUseAttempts = couponUseAttempts
	this.CouponFailedAttempts = couponFailedAttempts
	this.AccruedDiscounts = accruedDiscounts
	this.AccruedRevenue = accruedRevenue
	this.TotalOrders = totalOrders
	this.TotalOrdersNoCoupon = totalOrdersNoCoupon
	this.CampaignName = campaignName
	return &this
}

// NewCustomerActivityReportWithDefaults instantiates a new CustomerActivityReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerActivityReportWithDefaults() *CustomerActivityReport {
	this := CustomerActivityReport{}
	return &this
}

// GetIntegrationId returns the IntegrationId field value
func (o *CustomerActivityReport) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *CustomerActivityReport) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *CustomerActivityReport) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetCreated returns the Created field value
func (o *CustomerActivityReport) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *CustomerActivityReport) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *CustomerActivityReport) SetCreated(v time.Time) {
	o.Created = v
}

// GetName returns the Name field value
func (o *CustomerActivityReport) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomerActivityReport) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomerActivityReport) SetName(v string) {
	o.Name = v
}

// GetCustomerId returns the CustomerId field value
func (o *CustomerActivityReport) GetCustomerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CustomerActivityReport) GetCustomerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CustomerActivityReport) SetCustomerId(v int32) {
	o.CustomerId = v
}

// GetLastActivity returns the LastActivity field value if set, zero value otherwise.
func (o *CustomerActivityReport) GetLastActivity() time.Time {
	if o == nil || o.LastActivity == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActivity
}

// GetLastActivityOk returns a tuple with the LastActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerActivityReport) GetLastActivityOk() (*time.Time, bool) {
	if o == nil || o.LastActivity == nil {
		return nil, false
	}
	return o.LastActivity, true
}

// HasLastActivity returns a boolean if a field has been set.
func (o *CustomerActivityReport) HasLastActivity() bool {
	if o != nil && o.LastActivity != nil {
		return true
	}

	return false
}

// SetLastActivity gets a reference to the given time.Time and assigns it to the LastActivity field.
func (o *CustomerActivityReport) SetLastActivity(v time.Time) {
	o.LastActivity = &v
}

// GetCouponRedemptions returns the CouponRedemptions field value
func (o *CustomerActivityReport) GetCouponRedemptions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CouponRedemptions
}

// GetCouponRedemptionsOk returns a tuple with the CouponRedemptions field value
// and a boolean to check if the value has been set.
func (o *CustomerActivityReport) GetCouponRedemptionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CouponRedemptions, true
}

// SetCouponRedemptions sets field value
func (o *CustomerActivityReport) SetCouponRedemptions(v int32) {
	o.CouponRedemptions = v
}

// GetCouponUseAttempts returns the CouponUseAttempts field value
func (o *CustomerActivityReport) GetCouponUseAttempts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CouponUseAttempts
}

// GetCouponUseAttemptsOk returns a tuple with the CouponUseAttempts field value
// and a boolean to check if the value has been set.
func (o *CustomerActivityReport) GetCouponUseAttemptsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CouponUseAttempts, true
}

// SetCouponUseAttempts sets field value
func (o *CustomerActivityReport) SetCouponUseAttempts(v int32) {
	o.CouponUseAttempts = v
}

// GetCouponFailedAttempts returns the CouponFailedAttempts field value
func (o *CustomerActivityReport) GetCouponFailedAttempts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CouponFailedAttempts
}

// GetCouponFailedAttemptsOk returns a tuple with the CouponFailedAttempts field value
// and a boolean to check if the value has been set.
func (o *CustomerActivityReport) GetCouponFailedAttemptsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CouponFailedAttempts, true
}

// SetCouponFailedAttempts sets field value
func (o *CustomerActivityReport) SetCouponFailedAttempts(v int32) {
	o.CouponFailedAttempts = v
}

// GetAccruedDiscounts returns the AccruedDiscounts field value
func (o *CustomerActivityReport) GetAccruedDiscounts() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AccruedDiscounts
}

// GetAccruedDiscountsOk returns a tuple with the AccruedDiscounts field value
// and a boolean to check if the value has been set.
func (o *CustomerActivityReport) GetAccruedDiscountsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccruedDiscounts, true
}

// SetAccruedDiscounts sets field value
func (o *CustomerActivityReport) SetAccruedDiscounts(v float32) {
	o.AccruedDiscounts = v
}

// GetAccruedRevenue returns the AccruedRevenue field value
func (o *CustomerActivityReport) GetAccruedRevenue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AccruedRevenue
}

// GetAccruedRevenueOk returns a tuple with the AccruedRevenue field value
// and a boolean to check if the value has been set.
func (o *CustomerActivityReport) GetAccruedRevenueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccruedRevenue, true
}

// SetAccruedRevenue sets field value
func (o *CustomerActivityReport) SetAccruedRevenue(v float32) {
	o.AccruedRevenue = v
}

// GetTotalOrders returns the TotalOrders field value
func (o *CustomerActivityReport) GetTotalOrders() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalOrders
}

// GetTotalOrdersOk returns a tuple with the TotalOrders field value
// and a boolean to check if the value has been set.
func (o *CustomerActivityReport) GetTotalOrdersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalOrders, true
}

// SetTotalOrders sets field value
func (o *CustomerActivityReport) SetTotalOrders(v int32) {
	o.TotalOrders = v
}

// GetTotalOrdersNoCoupon returns the TotalOrdersNoCoupon field value
func (o *CustomerActivityReport) GetTotalOrdersNoCoupon() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalOrdersNoCoupon
}

// GetTotalOrdersNoCouponOk returns a tuple with the TotalOrdersNoCoupon field value
// and a boolean to check if the value has been set.
func (o *CustomerActivityReport) GetTotalOrdersNoCouponOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalOrdersNoCoupon, true
}

// SetTotalOrdersNoCoupon sets field value
func (o *CustomerActivityReport) SetTotalOrdersNoCoupon(v int32) {
	o.TotalOrdersNoCoupon = v
}

// GetCampaignName returns the CampaignName field value
func (o *CustomerActivityReport) GetCampaignName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignName
}

// GetCampaignNameOk returns a tuple with the CampaignName field value
// and a boolean to check if the value has been set.
func (o *CustomerActivityReport) GetCampaignNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignName, true
}

// SetCampaignName sets field value
func (o *CustomerActivityReport) SetCampaignName(v string) {
	o.CampaignName = v
}

func (o CustomerActivityReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["integrationId"] = o.IntegrationId
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["customerId"] = o.CustomerId
	}
	if o.LastActivity != nil {
		toSerialize["lastActivity"] = o.LastActivity
	}
	if true {
		toSerialize["couponRedemptions"] = o.CouponRedemptions
	}
	if true {
		toSerialize["couponUseAttempts"] = o.CouponUseAttempts
	}
	if true {
		toSerialize["couponFailedAttempts"] = o.CouponFailedAttempts
	}
	if true {
		toSerialize["accruedDiscounts"] = o.AccruedDiscounts
	}
	if true {
		toSerialize["accruedRevenue"] = o.AccruedRevenue
	}
	if true {
		toSerialize["totalOrders"] = o.TotalOrders
	}
	if true {
		toSerialize["totalOrdersNoCoupon"] = o.TotalOrdersNoCoupon
	}
	if true {
		toSerialize["campaignName"] = o.CampaignName
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerActivityReport struct {
	value *CustomerActivityReport
	isSet bool
}

func (v NullableCustomerActivityReport) Get() *CustomerActivityReport {
	return v.value
}

func (v *NullableCustomerActivityReport) Set(val *CustomerActivityReport) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerActivityReport) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerActivityReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerActivityReport(val *CustomerActivityReport) *NullableCustomerActivityReport {
	return &NullableCustomerActivityReport{value: val, isSet: true}
}

func (v NullableCustomerActivityReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerActivityReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


