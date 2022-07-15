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

// ApplicationApiHealth Report of health of the API connection of an application.
type ApplicationApiHealth struct {
	// One-word summary of the health of the API connection of an application.
	Summary string `json:"summary"`
	// time of last request relevant to the API health test.
	LastUsed time.Time `json:"lastUsed"`
}

// NewApplicationApiHealth instantiates a new ApplicationApiHealth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationApiHealth(summary string, lastUsed time.Time) *ApplicationApiHealth {
	this := ApplicationApiHealth{}
	this.Summary = summary
	this.LastUsed = lastUsed
	return &this
}

// NewApplicationApiHealthWithDefaults instantiates a new ApplicationApiHealth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationApiHealthWithDefaults() *ApplicationApiHealth {
	this := ApplicationApiHealth{}
	return &this
}

// GetSummary returns the Summary field value
func (o *ApplicationApiHealth) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *ApplicationApiHealth) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *ApplicationApiHealth) SetSummary(v string) {
	o.Summary = v
}

// GetLastUsed returns the LastUsed field value
func (o *ApplicationApiHealth) GetLastUsed() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUsed
}

// GetLastUsedOk returns a tuple with the LastUsed field value
// and a boolean to check if the value has been set.
func (o *ApplicationApiHealth) GetLastUsedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUsed, true
}

// SetLastUsed sets field value
func (o *ApplicationApiHealth) SetLastUsed(v time.Time) {
	o.LastUsed = v
}

func (o ApplicationApiHealth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["summary"] = o.Summary
	}
	if true {
		toSerialize["lastUsed"] = o.LastUsed
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationApiHealth struct {
	value *ApplicationApiHealth
	isSet bool
}

func (v NullableApplicationApiHealth) Get() *ApplicationApiHealth {
	return v.value
}

func (v *NullableApplicationApiHealth) Set(val *ApplicationApiHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationApiHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationApiHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationApiHealth(val *ApplicationApiHealth) *NullableApplicationApiHealth {
	return &NullableApplicationApiHealth{value: val, isSet: true}
}

func (v NullableApplicationApiHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationApiHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


