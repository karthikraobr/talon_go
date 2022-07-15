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

// LimitCounter 
type LimitCounter struct {
	// The ID of the campaign that owns this entity.
	CampaignId int32 `json:"campaignId"`
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// Unique ID for this entity.
	Id int32 `json:"id"`
	// The limitable action of the limit counter.
	Action string `json:"action"`
	// The profile ID for which this limit counter is used.
	ProfileId *int32 `json:"profileId,omitempty"`
	// The profile integration ID for which this limit counter is used.
	ProfileIntegrationId *string `json:"profileIntegrationId,omitempty"`
	// The coupon ID for which this limit counter is used.
	CouponId *int32 `json:"couponId,omitempty"`
	// The coupon value for which this limit counter is used.
	CouponValue *string `json:"couponValue,omitempty"`
	// The referral ID for which this limit counter is used.
	ReferralId *int32 `json:"referralId,omitempty"`
	// The referral value for which this limit counter is used.
	ReferralValue *string `json:"referralValue,omitempty"`
	// The arbitrary identifier for which this limit counter is used.
	Identifier *int32 `json:"identifier,omitempty"`
	// The time period for which this limit counter is used.
	Period *string `json:"period,omitempty"`
	// The highest possible value for this limit counter.
	Limit float32 `json:"limit"`
	// The current value for this limit counter.
	Counter float32 `json:"counter"`
}

// NewLimitCounter instantiates a new LimitCounter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLimitCounter(campaignId int32, applicationId int32, accountId int32, id int32, action string, limit float32, counter float32) *LimitCounter {
	this := LimitCounter{}
	this.CampaignId = campaignId
	this.ApplicationId = applicationId
	this.AccountId = accountId
	this.Id = id
	this.Action = action
	this.Limit = limit
	this.Counter = counter
	return &this
}

// NewLimitCounterWithDefaults instantiates a new LimitCounter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimitCounterWithDefaults() *LimitCounter {
	this := LimitCounter{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *LimitCounter) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *LimitCounter) GetCampaignIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *LimitCounter) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetApplicationId returns the ApplicationId field value
func (o *LimitCounter) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *LimitCounter) GetApplicationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *LimitCounter) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetAccountId returns the AccountId field value
func (o *LimitCounter) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *LimitCounter) GetAccountIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *LimitCounter) SetAccountId(v int32) {
	o.AccountId = v
}

// GetId returns the Id field value
func (o *LimitCounter) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LimitCounter) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LimitCounter) SetId(v int32) {
	o.Id = v
}

// GetAction returns the Action field value
func (o *LimitCounter) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *LimitCounter) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *LimitCounter) SetAction(v string) {
	o.Action = v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *LimitCounter) GetProfileId() int32 {
	if o == nil || o.ProfileId == nil {
		var ret int32
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitCounter) GetProfileIdOk() (*int32, bool) {
	if o == nil || o.ProfileId == nil {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *LimitCounter) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given int32 and assigns it to the ProfileId field.
func (o *LimitCounter) SetProfileId(v int32) {
	o.ProfileId = &v
}

// GetProfileIntegrationId returns the ProfileIntegrationId field value if set, zero value otherwise.
func (o *LimitCounter) GetProfileIntegrationId() string {
	if o == nil || o.ProfileIntegrationId == nil {
		var ret string
		return ret
	}
	return *o.ProfileIntegrationId
}

// GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitCounter) GetProfileIntegrationIdOk() (*string, bool) {
	if o == nil || o.ProfileIntegrationId == nil {
		return nil, false
	}
	return o.ProfileIntegrationId, true
}

// HasProfileIntegrationId returns a boolean if a field has been set.
func (o *LimitCounter) HasProfileIntegrationId() bool {
	if o != nil && o.ProfileIntegrationId != nil {
		return true
	}

	return false
}

// SetProfileIntegrationId gets a reference to the given string and assigns it to the ProfileIntegrationId field.
func (o *LimitCounter) SetProfileIntegrationId(v string) {
	o.ProfileIntegrationId = &v
}

// GetCouponId returns the CouponId field value if set, zero value otherwise.
func (o *LimitCounter) GetCouponId() int32 {
	if o == nil || o.CouponId == nil {
		var ret int32
		return ret
	}
	return *o.CouponId
}

// GetCouponIdOk returns a tuple with the CouponId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitCounter) GetCouponIdOk() (*int32, bool) {
	if o == nil || o.CouponId == nil {
		return nil, false
	}
	return o.CouponId, true
}

// HasCouponId returns a boolean if a field has been set.
func (o *LimitCounter) HasCouponId() bool {
	if o != nil && o.CouponId != nil {
		return true
	}

	return false
}

// SetCouponId gets a reference to the given int32 and assigns it to the CouponId field.
func (o *LimitCounter) SetCouponId(v int32) {
	o.CouponId = &v
}

// GetCouponValue returns the CouponValue field value if set, zero value otherwise.
func (o *LimitCounter) GetCouponValue() string {
	if o == nil || o.CouponValue == nil {
		var ret string
		return ret
	}
	return *o.CouponValue
}

// GetCouponValueOk returns a tuple with the CouponValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitCounter) GetCouponValueOk() (*string, bool) {
	if o == nil || o.CouponValue == nil {
		return nil, false
	}
	return o.CouponValue, true
}

// HasCouponValue returns a boolean if a field has been set.
func (o *LimitCounter) HasCouponValue() bool {
	if o != nil && o.CouponValue != nil {
		return true
	}

	return false
}

// SetCouponValue gets a reference to the given string and assigns it to the CouponValue field.
func (o *LimitCounter) SetCouponValue(v string) {
	o.CouponValue = &v
}

// GetReferralId returns the ReferralId field value if set, zero value otherwise.
func (o *LimitCounter) GetReferralId() int32 {
	if o == nil || o.ReferralId == nil {
		var ret int32
		return ret
	}
	return *o.ReferralId
}

// GetReferralIdOk returns a tuple with the ReferralId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitCounter) GetReferralIdOk() (*int32, bool) {
	if o == nil || o.ReferralId == nil {
		return nil, false
	}
	return o.ReferralId, true
}

// HasReferralId returns a boolean if a field has been set.
func (o *LimitCounter) HasReferralId() bool {
	if o != nil && o.ReferralId != nil {
		return true
	}

	return false
}

// SetReferralId gets a reference to the given int32 and assigns it to the ReferralId field.
func (o *LimitCounter) SetReferralId(v int32) {
	o.ReferralId = &v
}

// GetReferralValue returns the ReferralValue field value if set, zero value otherwise.
func (o *LimitCounter) GetReferralValue() string {
	if o == nil || o.ReferralValue == nil {
		var ret string
		return ret
	}
	return *o.ReferralValue
}

// GetReferralValueOk returns a tuple with the ReferralValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitCounter) GetReferralValueOk() (*string, bool) {
	if o == nil || o.ReferralValue == nil {
		return nil, false
	}
	return o.ReferralValue, true
}

// HasReferralValue returns a boolean if a field has been set.
func (o *LimitCounter) HasReferralValue() bool {
	if o != nil && o.ReferralValue != nil {
		return true
	}

	return false
}

// SetReferralValue gets a reference to the given string and assigns it to the ReferralValue field.
func (o *LimitCounter) SetReferralValue(v string) {
	o.ReferralValue = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *LimitCounter) GetIdentifier() int32 {
	if o == nil || o.Identifier == nil {
		var ret int32
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitCounter) GetIdentifierOk() (*int32, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *LimitCounter) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given int32 and assigns it to the Identifier field.
func (o *LimitCounter) SetIdentifier(v int32) {
	o.Identifier = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *LimitCounter) GetPeriod() string {
	if o == nil || o.Period == nil {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitCounter) GetPeriodOk() (*string, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *LimitCounter) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *LimitCounter) SetPeriod(v string) {
	o.Period = &v
}

// GetLimit returns the Limit field value
func (o *LimitCounter) GetLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *LimitCounter) GetLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *LimitCounter) SetLimit(v float32) {
	o.Limit = v
}

// GetCounter returns the Counter field value
func (o *LimitCounter) GetCounter() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Counter
}

// GetCounterOk returns a tuple with the Counter field value
// and a boolean to check if the value has been set.
func (o *LimitCounter) GetCounterOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Counter, true
}

// SetCounter sets field value
func (o *LimitCounter) SetCounter(v float32) {
	o.Counter = v
}

func (o LimitCounter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["campaignId"] = o.CampaignId
	}
	if true {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["action"] = o.Action
	}
	if o.ProfileId != nil {
		toSerialize["profileId"] = o.ProfileId
	}
	if o.ProfileIntegrationId != nil {
		toSerialize["profileIntegrationId"] = o.ProfileIntegrationId
	}
	if o.CouponId != nil {
		toSerialize["couponId"] = o.CouponId
	}
	if o.CouponValue != nil {
		toSerialize["couponValue"] = o.CouponValue
	}
	if o.ReferralId != nil {
		toSerialize["referralId"] = o.ReferralId
	}
	if o.ReferralValue != nil {
		toSerialize["referralValue"] = o.ReferralValue
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["counter"] = o.Counter
	}
	return json.Marshal(toSerialize)
}

type NullableLimitCounter struct {
	value *LimitCounter
	isSet bool
}

func (v NullableLimitCounter) Get() *LimitCounter {
	return v.value
}

func (v *NullableLimitCounter) Set(val *LimitCounter) {
	v.value = val
	v.isSet = true
}

func (v NullableLimitCounter) IsSet() bool {
	return v.isSet
}

func (v *NullableLimitCounter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLimitCounter(val *LimitCounter) *NullableLimitCounter {
	return &NullableLimitCounter{value: val, isSet: true}
}

func (v NullableLimitCounter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLimitCounter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


