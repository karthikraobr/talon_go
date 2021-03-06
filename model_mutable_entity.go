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

// MutableEntity struct for MutableEntity
type MutableEntity struct {
	// The exact moment this entity was last modified.
	Modified time.Time `json:"modified"`
}

// NewMutableEntity instantiates a new MutableEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutableEntity(modified time.Time) *MutableEntity {
	this := MutableEntity{}
	this.Modified = modified
	return &this
}

// NewMutableEntityWithDefaults instantiates a new MutableEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutableEntityWithDefaults() *MutableEntity {
	this := MutableEntity{}
	return &this
}

// GetModified returns the Modified field value
func (o *MutableEntity) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *MutableEntity) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *MutableEntity) SetModified(v time.Time) {
	o.Modified = v
}

func (o MutableEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["modified"] = o.Modified
	}
	return json.Marshal(toSerialize)
}

type NullableMutableEntity struct {
	value *MutableEntity
	isSet bool
}

func (v NullableMutableEntity) Get() *MutableEntity {
	return v.value
}

func (v *NullableMutableEntity) Set(val *MutableEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableMutableEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableMutableEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutableEntity(val *MutableEntity) *NullableMutableEntity {
	return &NullableMutableEntity{value: val, isSet: true}
}

func (v NullableMutableEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutableEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


