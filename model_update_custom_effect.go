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

// UpdateCustomEffect 
type UpdateCustomEffect struct {
	// The IDs of the applications that are related to this entity.
	ApplicationIds []int32 `json:"applicationIds"`
	// The name of this effect.
	Name string `json:"name"`
	// The title of this effect.
	Title string `json:"title"`
	// The JSON payload of this effect.
	Payload string `json:"payload"`
	// The description of this effect.
	Description *string `json:"description,omitempty"`
	// Determines if this effect is active.
	Enabled bool `json:"enabled"`
	// Array of template argument definitions.
	Params []TemplateArgDef `json:"params,omitempty"`
}

// NewUpdateCustomEffect instantiates a new UpdateCustomEffect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCustomEffect(applicationIds []int32, name string, title string, payload string, enabled bool) *UpdateCustomEffect {
	this := UpdateCustomEffect{}
	this.ApplicationIds = applicationIds
	this.Name = name
	this.Title = title
	this.Payload = payload
	this.Enabled = enabled
	return &this
}

// NewUpdateCustomEffectWithDefaults instantiates a new UpdateCustomEffect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCustomEffectWithDefaults() *UpdateCustomEffect {
	this := UpdateCustomEffect{}
	return &this
}

// GetApplicationIds returns the ApplicationIds field value
func (o *UpdateCustomEffect) GetApplicationIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ApplicationIds
}

// GetApplicationIdsOk returns a tuple with the ApplicationIds field value
// and a boolean to check if the value has been set.
func (o *UpdateCustomEffect) GetApplicationIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationIds, true
}

// SetApplicationIds sets field value
func (o *UpdateCustomEffect) SetApplicationIds(v []int32) {
	o.ApplicationIds = v
}

// GetName returns the Name field value
func (o *UpdateCustomEffect) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateCustomEffect) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateCustomEffect) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *UpdateCustomEffect) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *UpdateCustomEffect) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *UpdateCustomEffect) SetTitle(v string) {
	o.Title = v
}

// GetPayload returns the Payload field value
func (o *UpdateCustomEffect) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *UpdateCustomEffect) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *UpdateCustomEffect) SetPayload(v string) {
	o.Payload = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateCustomEffect) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomEffect) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateCustomEffect) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateCustomEffect) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *UpdateCustomEffect) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UpdateCustomEffect) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *UpdateCustomEffect) SetEnabled(v bool) {
	o.Enabled = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *UpdateCustomEffect) GetParams() []TemplateArgDef {
	if o == nil || o.Params == nil {
		var ret []TemplateArgDef
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomEffect) GetParamsOk() ([]TemplateArgDef, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *UpdateCustomEffect) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given []TemplateArgDef and assigns it to the Params field.
func (o *UpdateCustomEffect) SetParams(v []TemplateArgDef) {
	o.Params = v
}

func (o UpdateCustomEffect) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["applicationIds"] = o.ApplicationIds
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["payload"] = o.Payload
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateCustomEffect struct {
	value *UpdateCustomEffect
	isSet bool
}

func (v NullableUpdateCustomEffect) Get() *UpdateCustomEffect {
	return v.value
}

func (v *NullableUpdateCustomEffect) Set(val *UpdateCustomEffect) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCustomEffect) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCustomEffect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCustomEffect(val *UpdateCustomEffect) *NullableUpdateCustomEffect {
	return &NullableUpdateCustomEffect{value: val, isSet: true}
}

func (v NullableUpdateCustomEffect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCustomEffect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


