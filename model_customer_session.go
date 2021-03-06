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

// CustomerSession 
type CustomerSession struct {
	// The integration ID set by your integration layer.
	IntegrationId string `json:"integrationId"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known `profileId`, we recommend you use a guest `profileId`. 
	ProfileId string `json:"profileId"`
	// Any coupon code entered.
	Coupon string `json:"coupon"`
	// Any referral code entered.
	Referral string `json:"referral"`
	// Indicates the current state of the session. Sessions can be created as `open` or `closed`. The state transitions are:  1. `open` → `closed` 2. `open` → `cancelled` 3. `closed` → `cancelled` or `partially_returned` 4. `partially_returned` → `cancelled`  For more information, see [Customer session states](/docs/dev/concepts/entities#customer-session). 
	State string `json:"state"`
	// Serialized JSON representation.
	CartItems []CartItem `json:"cartItems"`
	// Session custom identifiers that you can set limits on or use inside your rules.  For example, you can use IP addresses as identifiers to potentially identify devices and limit discounts abuse in case of customers creating multiple accounts. See the [tutorial](https://docs.talon.one/docs/dev/tutorials/using-identifiers/). 
	Identifiers []string `json:"identifiers,omitempty"`
	// The total sum of the cart in one session.
	Total float32 `json:"total"`
	// A key-value map of the sessions attributes. The potentially valid attributes are configured in your accounts developer settings. 
	Attributes map[string]interface{} `json:"attributes"`
	// Indicates whether this is the first session for the customer's profile. Will always be true for anonymous sessions.
	FirstSession bool `json:"firstSession"`
	// A map of labelled discount values, values will be in the same currency as the application associated with the session.
	Discounts map[string]float32 `json:"discounts"`
	// Timestamp of the most recent event received on this session.
	Updated time.Time `json:"updated"`
}

// NewCustomerSession instantiates a new CustomerSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSession(integrationId string, created time.Time, applicationId int32, profileId string, coupon string, referral string, state string, cartItems []CartItem, total float32, attributes map[string]interface{}, firstSession bool, discounts map[string]float32, updated time.Time) *CustomerSession {
	this := CustomerSession{}
	this.IntegrationId = integrationId
	this.Created = created
	this.ApplicationId = applicationId
	this.ProfileId = profileId
	this.Coupon = coupon
	this.Referral = referral
	this.State = state
	this.CartItems = cartItems
	this.Total = total
	this.Attributes = attributes
	this.FirstSession = firstSession
	this.Discounts = discounts
	this.Updated = updated
	return &this
}

// NewCustomerSessionWithDefaults instantiates a new CustomerSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSessionWithDefaults() *CustomerSession {
	this := CustomerSession{}
	var state string = "open"
	this.State = state
	return &this
}

// GetIntegrationId returns the IntegrationId field value
func (o *CustomerSession) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *CustomerSession) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *CustomerSession) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetCreated returns the Created field value
func (o *CustomerSession) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *CustomerSession) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *CustomerSession) SetCreated(v time.Time) {
	o.Created = v
}

// GetApplicationId returns the ApplicationId field value
func (o *CustomerSession) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *CustomerSession) GetApplicationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *CustomerSession) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetProfileId returns the ProfileId field value
func (o *CustomerSession) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value
// and a boolean to check if the value has been set.
func (o *CustomerSession) GetProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileId, true
}

// SetProfileId sets field value
func (o *CustomerSession) SetProfileId(v string) {
	o.ProfileId = v
}

// GetCoupon returns the Coupon field value
func (o *CustomerSession) GetCoupon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coupon
}

// GetCouponOk returns a tuple with the Coupon field value
// and a boolean to check if the value has been set.
func (o *CustomerSession) GetCouponOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coupon, true
}

// SetCoupon sets field value
func (o *CustomerSession) SetCoupon(v string) {
	o.Coupon = v
}

// GetReferral returns the Referral field value
func (o *CustomerSession) GetReferral() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Referral
}

// GetReferralOk returns a tuple with the Referral field value
// and a boolean to check if the value has been set.
func (o *CustomerSession) GetReferralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Referral, true
}

// SetReferral sets field value
func (o *CustomerSession) SetReferral(v string) {
	o.Referral = v
}

// GetState returns the State field value
func (o *CustomerSession) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CustomerSession) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CustomerSession) SetState(v string) {
	o.State = v
}

// GetCartItems returns the CartItems field value
func (o *CustomerSession) GetCartItems() []CartItem {
	if o == nil {
		var ret []CartItem
		return ret
	}

	return o.CartItems
}

// GetCartItemsOk returns a tuple with the CartItems field value
// and a boolean to check if the value has been set.
func (o *CustomerSession) GetCartItemsOk() ([]CartItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.CartItems, true
}

// SetCartItems sets field value
func (o *CustomerSession) SetCartItems(v []CartItem) {
	o.CartItems = v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise.
func (o *CustomerSession) GetIdentifiers() []string {
	if o == nil || o.Identifiers == nil {
		var ret []string
		return ret
	}
	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSession) GetIdentifiersOk() ([]string, bool) {
	if o == nil || o.Identifiers == nil {
		return nil, false
	}
	return o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *CustomerSession) HasIdentifiers() bool {
	if o != nil && o.Identifiers != nil {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given []string and assigns it to the Identifiers field.
func (o *CustomerSession) SetIdentifiers(v []string) {
	o.Identifiers = v
}

// GetTotal returns the Total field value
func (o *CustomerSession) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CustomerSession) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CustomerSession) SetTotal(v float32) {
	o.Total = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerSession) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerSession) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerSession) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetFirstSession returns the FirstSession field value
func (o *CustomerSession) GetFirstSession() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FirstSession
}

// GetFirstSessionOk returns a tuple with the FirstSession field value
// and a boolean to check if the value has been set.
func (o *CustomerSession) GetFirstSessionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstSession, true
}

// SetFirstSession sets field value
func (o *CustomerSession) SetFirstSession(v bool) {
	o.FirstSession = v
}

// GetDiscounts returns the Discounts field value
func (o *CustomerSession) GetDiscounts() map[string]float32 {
	if o == nil {
		var ret map[string]float32
		return ret
	}

	return o.Discounts
}

// GetDiscountsOk returns a tuple with the Discounts field value
// and a boolean to check if the value has been set.
func (o *CustomerSession) GetDiscountsOk() (*map[string]float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Discounts, true
}

// SetDiscounts sets field value
func (o *CustomerSession) SetDiscounts(v map[string]float32) {
	o.Discounts = v
}

// GetUpdated returns the Updated field value
func (o *CustomerSession) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *CustomerSession) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *CustomerSession) SetUpdated(v time.Time) {
	o.Updated = v
}

func (o CustomerSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["integrationId"] = o.IntegrationId
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if true {
		toSerialize["profileId"] = o.ProfileId
	}
	if true {
		toSerialize["coupon"] = o.Coupon
	}
	if true {
		toSerialize["referral"] = o.Referral
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["cartItems"] = o.CartItems
	}
	if o.Identifiers != nil {
		toSerialize["identifiers"] = o.Identifiers
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["firstSession"] = o.FirstSession
	}
	if true {
		toSerialize["discounts"] = o.Discounts
	}
	if true {
		toSerialize["updated"] = o.Updated
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerSession struct {
	value *CustomerSession
	isSet bool
}

func (v NullableCustomerSession) Get() *CustomerSession {
	return v.value
}

func (v *NullableCustomerSession) Set(val *CustomerSession) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSession) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSession(val *CustomerSession) *NullableCustomerSession {
	return &NullableCustomerSession{value: val, isSet: true}
}

func (v NullableCustomerSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


