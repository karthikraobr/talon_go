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

// AttributesSettings Arbitrary settings associated with attributes.
type AttributesSettings struct {
	Mandatory *AttributesMandatory `json:"mandatory,omitempty"`
}

// NewAttributesSettings instantiates a new AttributesSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributesSettings() *AttributesSettings {
	this := AttributesSettings{}
	return &this
}

// NewAttributesSettingsWithDefaults instantiates a new AttributesSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributesSettingsWithDefaults() *AttributesSettings {
	this := AttributesSettings{}
	return &this
}

// GetMandatory returns the Mandatory field value if set, zero value otherwise.
func (o *AttributesSettings) GetMandatory() AttributesMandatory {
	if o == nil || o.Mandatory == nil {
		var ret AttributesMandatory
		return ret
	}
	return *o.Mandatory
}

// GetMandatoryOk returns a tuple with the Mandatory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributesSettings) GetMandatoryOk() (*AttributesMandatory, bool) {
	if o == nil || o.Mandatory == nil {
		return nil, false
	}
	return o.Mandatory, true
}

// HasMandatory returns a boolean if a field has been set.
func (o *AttributesSettings) HasMandatory() bool {
	if o != nil && o.Mandatory != nil {
		return true
	}

	return false
}

// SetMandatory gets a reference to the given AttributesMandatory and assigns it to the Mandatory field.
func (o *AttributesSettings) SetMandatory(v AttributesMandatory) {
	o.Mandatory = &v
}

func (o AttributesSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mandatory != nil {
		toSerialize["mandatory"] = o.Mandatory
	}
	return json.Marshal(toSerialize)
}

type NullableAttributesSettings struct {
	value *AttributesSettings
	isSet bool
}

func (v NullableAttributesSettings) Get() *AttributesSettings {
	return v.value
}

func (v *NullableAttributesSettings) Set(val *AttributesSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributesSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributesSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributesSettings(val *AttributesSettings) *NullableAttributesSettings {
	return &NullableAttributesSettings{value: val, isSet: true}
}

func (v NullableAttributesSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributesSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


