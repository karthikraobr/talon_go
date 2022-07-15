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

// AddFreeItemEffectProps The properties specific to the \"addFreeItem\" effect. This gets triggered whenever a validated rule contained an \"add free item\" effect.
type AddFreeItemEffectProps struct {
	// SKU of the item that needs to be added.
	Sku string `json:"sku"`
	// The name/description of the effect.
	Name string `json:"name"`
}

// NewAddFreeItemEffectProps instantiates a new AddFreeItemEffectProps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddFreeItemEffectProps(sku string, name string) *AddFreeItemEffectProps {
	this := AddFreeItemEffectProps{}
	this.Sku = sku
	this.Name = name
	return &this
}

// NewAddFreeItemEffectPropsWithDefaults instantiates a new AddFreeItemEffectProps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddFreeItemEffectPropsWithDefaults() *AddFreeItemEffectProps {
	this := AddFreeItemEffectProps{}
	return &this
}

// GetSku returns the Sku field value
func (o *AddFreeItemEffectProps) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *AddFreeItemEffectProps) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *AddFreeItemEffectProps) SetSku(v string) {
	o.Sku = v
}

// GetName returns the Name field value
func (o *AddFreeItemEffectProps) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddFreeItemEffectProps) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddFreeItemEffectProps) SetName(v string) {
	o.Name = v
}

func (o AddFreeItemEffectProps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sku"] = o.Sku
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableAddFreeItemEffectProps struct {
	value *AddFreeItemEffectProps
	isSet bool
}

func (v NullableAddFreeItemEffectProps) Get() *AddFreeItemEffectProps {
	return v.value
}

func (v *NullableAddFreeItemEffectProps) Set(val *AddFreeItemEffectProps) {
	v.value = val
	v.isSet = true
}

func (v NullableAddFreeItemEffectProps) IsSet() bool {
	return v.isSet
}

func (v *NullableAddFreeItemEffectProps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddFreeItemEffectProps(val *AddFreeItemEffectProps) *NullableAddFreeItemEffectProps {
	return &NullableAddFreeItemEffectProps{value: val, isSet: true}
}

func (v NullableAddFreeItemEffectProps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddFreeItemEffectProps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


