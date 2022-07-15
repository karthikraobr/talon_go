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

// NewCampaignCollection 
type NewCampaignCollection struct {
	// A short description of the purpose of this collection.
	Description *string `json:"description,omitempty"`
	// The name of this collection.
	Name string `json:"name"`
}

// NewNewCampaignCollection instantiates a new NewCampaignCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewCampaignCollection(name string) *NewCampaignCollection {
	this := NewCampaignCollection{}
	this.Name = name
	return &this
}

// NewNewCampaignCollectionWithDefaults instantiates a new NewCampaignCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewCampaignCollectionWithDefaults() *NewCampaignCollection {
	this := NewCampaignCollection{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NewCampaignCollection) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaignCollection) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NewCampaignCollection) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NewCampaignCollection) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *NewCampaignCollection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NewCampaignCollection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NewCampaignCollection) SetName(v string) {
	o.Name = v
}

func (o NewCampaignCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNewCampaignCollection struct {
	value *NewCampaignCollection
	isSet bool
}

func (v NullableNewCampaignCollection) Get() *NewCampaignCollection {
	return v.value
}

func (v *NullableNewCampaignCollection) Set(val *NewCampaignCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableNewCampaignCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableNewCampaignCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewCampaignCollection(val *NewCampaignCollection) *NullableNewCampaignCollection {
	return &NullableNewCampaignCollection{value: val, isSet: true}
}

func (v NullableNewCampaignCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewCampaignCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


