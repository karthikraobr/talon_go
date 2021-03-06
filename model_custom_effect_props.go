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

// CustomEffectProps Effect containing custom payload.
type CustomEffectProps struct {
	// The ID of the custom effect that was triggered.
	EffectId int32 `json:"effectId"`
	// The type of the custom effect.
	Name string `json:"name"`
	// The JSON payload of the custom effect.
	Payload map[string]interface{} `json:"payload"`
}

// NewCustomEffectProps instantiates a new CustomEffectProps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEffectProps(effectId int32, name string, payload map[string]interface{}) *CustomEffectProps {
	this := CustomEffectProps{}
	this.EffectId = effectId
	this.Name = name
	this.Payload = payload
	return &this
}

// NewCustomEffectPropsWithDefaults instantiates a new CustomEffectProps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEffectPropsWithDefaults() *CustomEffectProps {
	this := CustomEffectProps{}
	return &this
}

// GetEffectId returns the EffectId field value
func (o *CustomEffectProps) GetEffectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EffectId
}

// GetEffectIdOk returns a tuple with the EffectId field value
// and a boolean to check if the value has been set.
func (o *CustomEffectProps) GetEffectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectId, true
}

// SetEffectId sets field value
func (o *CustomEffectProps) SetEffectId(v int32) {
	o.EffectId = v
}

// GetName returns the Name field value
func (o *CustomEffectProps) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomEffectProps) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomEffectProps) SetName(v string) {
	o.Name = v
}

// GetPayload returns the Payload field value
func (o *CustomEffectProps) GetPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *CustomEffectProps) GetPayloadOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payload, true
}

// SetPayload sets field value
func (o *CustomEffectProps) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

func (o CustomEffectProps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["effectId"] = o.EffectId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["payload"] = o.Payload
	}
	return json.Marshal(toSerialize)
}

type NullableCustomEffectProps struct {
	value *CustomEffectProps
	isSet bool
}

func (v NullableCustomEffectProps) Get() *CustomEffectProps {
	return v.value
}

func (v *NullableCustomEffectProps) Set(val *CustomEffectProps) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEffectProps) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEffectProps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEffectProps(val *CustomEffectProps) *NullableCustomEffectProps {
	return &NullableCustomEffectProps{value: val, isSet: true}
}

func (v NullableCustomEffectProps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEffectProps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


