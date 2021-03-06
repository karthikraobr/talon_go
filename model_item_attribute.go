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

// ItemAttribute 
type ItemAttribute struct {
	// The ID of the attribute of the item.
	Attributeid int32 `json:"attributeid"`
	// The name of the attribute.
	Name string `json:"name"`
	// The value of the attribute.
	Value map[string]interface{} `json:"value"`
}

// NewItemAttribute instantiates a new ItemAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemAttribute(attributeid int32, name string, value map[string]interface{}) *ItemAttribute {
	this := ItemAttribute{}
	this.Attributeid = attributeid
	this.Name = name
	this.Value = value
	return &this
}

// NewItemAttributeWithDefaults instantiates a new ItemAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemAttributeWithDefaults() *ItemAttribute {
	this := ItemAttribute{}
	return &this
}

// GetAttributeid returns the Attributeid field value
func (o *ItemAttribute) GetAttributeid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Attributeid
}

// GetAttributeidOk returns a tuple with the Attributeid field value
// and a boolean to check if the value has been set.
func (o *ItemAttribute) GetAttributeidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributeid, true
}

// SetAttributeid sets field value
func (o *ItemAttribute) SetAttributeid(v int32) {
	o.Attributeid = v
}

// GetName returns the Name field value
func (o *ItemAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ItemAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ItemAttribute) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *ItemAttribute) GetValue() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ItemAttribute) GetValueOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *ItemAttribute) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o ItemAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attributeid"] = o.Attributeid
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableItemAttribute struct {
	value *ItemAttribute
	isSet bool
}

func (v NullableItemAttribute) Get() *ItemAttribute {
	return v.value
}

func (v *NullableItemAttribute) Set(val *ItemAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableItemAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableItemAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemAttribute(val *ItemAttribute) *NullableItemAttribute {
	return &NullableItemAttribute{value: val, isSet: true}
}

func (v NullableItemAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


