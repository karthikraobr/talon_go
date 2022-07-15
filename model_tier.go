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

// Tier struct for Tier
type Tier struct {
	// The internal ID of the tier.
	Id int32 `json:"id"`
	// The name of the tier.
	Name string `json:"name"`
}

// NewTier instantiates a new Tier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTier(id int32, name string) *Tier {
	this := Tier{}
	this.Id = id
	this.Name = name
	return &this
}

// NewTierWithDefaults instantiates a new Tier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTierWithDefaults() *Tier {
	this := Tier{}
	return &this
}

// GetId returns the Id field value
func (o *Tier) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Tier) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Tier) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Tier) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tier) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tier) SetName(v string) {
	o.Name = v
}

func (o Tier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableTier struct {
	value *Tier
	isSet bool
}

func (v NullableTier) Get() *Tier {
	return v.value
}

func (v *NullableTier) Set(val *Tier) {
	v.value = val
	v.isSet = true
}

func (v NullableTier) IsSet() bool {
	return v.isSet
}

func (v *NullableTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTier(val *Tier) *NullableTier {
	return &NullableTier{value: val, isSet: true}
}

func (v NullableTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


