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

// CouponValue struct for CouponValue
type CouponValue struct {
	// The coupon code.
	Value *string `json:"value,omitempty"`
}

// NewCouponValue instantiates a new CouponValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponValue() *CouponValue {
	this := CouponValue{}
	return &this
}

// NewCouponValueWithDefaults instantiates a new CouponValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponValueWithDefaults() *CouponValue {
	this := CouponValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CouponValue) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponValue) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CouponValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CouponValue) SetValue(v string) {
	o.Value = &v
}

func (o CouponValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCouponValue struct {
	value *CouponValue
	isSet bool
}

func (v NullableCouponValue) Get() *CouponValue {
	return v.value
}

func (v *NullableCouponValue) Set(val *CouponValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponValue(val *CouponValue) *NullableCouponValue {
	return &NullableCouponValue{value: val, isSet: true}
}

func (v NullableCouponValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

