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

// MultipleNewAttribute struct for MultipleNewAttribute
type MultipleNewAttribute struct {
	Attributes []NewAttribute `json:"attributes,omitempty"`
}

// NewMultipleNewAttribute instantiates a new MultipleNewAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleNewAttribute() *MultipleNewAttribute {
	this := MultipleNewAttribute{}
	return &this
}

// NewMultipleNewAttributeWithDefaults instantiates a new MultipleNewAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleNewAttributeWithDefaults() *MultipleNewAttribute {
	this := MultipleNewAttribute{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MultipleNewAttribute) GetAttributes() []NewAttribute {
	if o == nil || o.Attributes == nil {
		var ret []NewAttribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleNewAttribute) GetAttributesOk() ([]NewAttribute, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MultipleNewAttribute) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []NewAttribute and assigns it to the Attributes field.
func (o *MultipleNewAttribute) SetAttributes(v []NewAttribute) {
	o.Attributes = v
}

func (o MultipleNewAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableMultipleNewAttribute struct {
	value *MultipleNewAttribute
	isSet bool
}

func (v NullableMultipleNewAttribute) Get() *MultipleNewAttribute {
	return v.value
}

func (v *NullableMultipleNewAttribute) Set(val *MultipleNewAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleNewAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleNewAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleNewAttribute(val *MultipleNewAttribute) *NullableMultipleNewAttribute {
	return &NullableMultipleNewAttribute{value: val, isSet: true}
}

func (v NullableMultipleNewAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleNewAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


