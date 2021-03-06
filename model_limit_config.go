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

// LimitConfig struct for LimitConfig
type LimitConfig struct {
	// The limitable action to which this limit applies. For example: - `setDiscount` - `setDiscountEffect` - `redeemCoupon` - `createCoupon` 
	Action string `json:"action"`
	// The value to set for the limit.
	Limit float32 `json:"limit"`
	// The period on which the budget limit recurs.
	Period *string `json:"period,omitempty"`
	// The entity that this limit applies to.
	Entities []string `json:"entities"`
}

// NewLimitConfig instantiates a new LimitConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLimitConfig(action string, limit float32, entities []string) *LimitConfig {
	this := LimitConfig{}
	this.Action = action
	this.Limit = limit
	this.Entities = entities
	return &this
}

// NewLimitConfigWithDefaults instantiates a new LimitConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimitConfigWithDefaults() *LimitConfig {
	this := LimitConfig{}
	return &this
}

// GetAction returns the Action field value
func (o *LimitConfig) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *LimitConfig) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *LimitConfig) SetAction(v string) {
	o.Action = v
}

// GetLimit returns the Limit field value
func (o *LimitConfig) GetLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *LimitConfig) GetLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *LimitConfig) SetLimit(v float32) {
	o.Limit = v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *LimitConfig) GetPeriod() string {
	if o == nil || o.Period == nil {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitConfig) GetPeriodOk() (*string, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *LimitConfig) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *LimitConfig) SetPeriod(v string) {
	o.Period = &v
}

// GetEntities returns the Entities field value
func (o *LimitConfig) GetEntities() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
func (o *LimitConfig) GetEntitiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entities, true
}

// SetEntities sets field value
func (o *LimitConfig) SetEntities(v []string) {
	o.Entities = v
}

func (o LimitConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if true {
		toSerialize["entities"] = o.Entities
	}
	return json.Marshal(toSerialize)
}

type NullableLimitConfig struct {
	value *LimitConfig
	isSet bool
}

func (v NullableLimitConfig) Get() *LimitConfig {
	return v.value
}

func (v *NullableLimitConfig) Set(val *LimitConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableLimitConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableLimitConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLimitConfig(val *LimitConfig) *NullableLimitConfig {
	return &NullableLimitConfig{value: val, isSet: true}
}

func (v NullableLimitConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLimitConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


