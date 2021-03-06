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

// FeatureFlag struct for FeatureFlag
type FeatureFlag struct {
	// The name for the featureflag.
	Name string `json:"name"`
	// The value for the featureflag.
	Value string `json:"value"`
	// The exact moment this entity was last created.
	Created *time.Time `json:"created,omitempty"`
	// The exact moment this entity was last modified.
	Modified *time.Time `json:"modified,omitempty"`
}

// NewFeatureFlag instantiates a new FeatureFlag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureFlag(name string, value string) *FeatureFlag {
	this := FeatureFlag{}
	this.Name = name
	this.Value = value
	return &this
}

// NewFeatureFlagWithDefaults instantiates a new FeatureFlag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureFlagWithDefaults() *FeatureFlag {
	this := FeatureFlag{}
	return &this
}

// GetName returns the Name field value
func (o *FeatureFlag) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FeatureFlag) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FeatureFlag) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *FeatureFlag) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FeatureFlag) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FeatureFlag) SetValue(v string) {
	o.Value = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *FeatureFlag) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlag) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *FeatureFlag) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *FeatureFlag) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *FeatureFlag) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlag) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *FeatureFlag) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *FeatureFlag) SetModified(v time.Time) {
	o.Modified = &v
}

func (o FeatureFlag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	return json.Marshal(toSerialize)
}

type NullableFeatureFlag struct {
	value *FeatureFlag
	isSet bool
}

func (v NullableFeatureFlag) Get() *FeatureFlag {
	return v.value
}

func (v *NullableFeatureFlag) Set(val *FeatureFlag) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureFlag) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureFlag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureFlag(val *FeatureFlag) *NullableFeatureFlag {
	return &NullableFeatureFlag{value: val, isSet: true}
}

func (v NullableFeatureFlag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureFlag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


