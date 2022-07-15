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

// NewCollection 
type NewCollection struct {
	// A short description of the purpose of this collection.
	Description *string `json:"description,omitempty"`
	// A list of the IDs of the Applications where this collection is enabled.
	SubscribedApplicationsIds []int32 `json:"subscribedApplicationsIds,omitempty"`
	// The name of this collection.
	Name string `json:"name"`
}

// NewNewCollection instantiates a new NewCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewCollection(name string) *NewCollection {
	this := NewCollection{}
	this.Name = name
	return &this
}

// NewNewCollectionWithDefaults instantiates a new NewCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewCollectionWithDefaults() *NewCollection {
	this := NewCollection{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NewCollection) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewCollection) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NewCollection) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NewCollection) SetDescription(v string) {
	o.Description = &v
}

// GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field value if set, zero value otherwise.
func (o *NewCollection) GetSubscribedApplicationsIds() []int32 {
	if o == nil || o.SubscribedApplicationsIds == nil {
		var ret []int32
		return ret
	}
	return o.SubscribedApplicationsIds
}

// GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewCollection) GetSubscribedApplicationsIdsOk() ([]int32, bool) {
	if o == nil || o.SubscribedApplicationsIds == nil {
		return nil, false
	}
	return o.SubscribedApplicationsIds, true
}

// HasSubscribedApplicationsIds returns a boolean if a field has been set.
func (o *NewCollection) HasSubscribedApplicationsIds() bool {
	if o != nil && o.SubscribedApplicationsIds != nil {
		return true
	}

	return false
}

// SetSubscribedApplicationsIds gets a reference to the given []int32 and assigns it to the SubscribedApplicationsIds field.
func (o *NewCollection) SetSubscribedApplicationsIds(v []int32) {
	o.SubscribedApplicationsIds = v
}

// GetName returns the Name field value
func (o *NewCollection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NewCollection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NewCollection) SetName(v string) {
	o.Name = v
}

func (o NewCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.SubscribedApplicationsIds != nil {
		toSerialize["subscribedApplicationsIds"] = o.SubscribedApplicationsIds
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNewCollection struct {
	value *NewCollection
	isSet bool
}

func (v NullableNewCollection) Get() *NewCollection {
	return v.value
}

func (v *NullableNewCollection) Set(val *NewCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableNewCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableNewCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewCollection(val *NewCollection) *NullableNewCollection {
	return &NullableNewCollection{value: val, isSet: true}
}

func (v NullableNewCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


