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

// UpdateCollection struct for UpdateCollection
type UpdateCollection struct {
	// A short description of the purpose of this collection.
	Description *string `json:"description,omitempty"`
	// A list of the IDs of the Applications where this collection is enabled.
	SubscribedApplicationsIds []int32 `json:"subscribedApplicationsIds,omitempty"`
}

// NewUpdateCollection instantiates a new UpdateCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCollection() *UpdateCollection {
	this := UpdateCollection{}
	return &this
}

// NewUpdateCollectionWithDefaults instantiates a new UpdateCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCollectionWithDefaults() *UpdateCollection {
	this := UpdateCollection{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateCollection) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCollection) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateCollection) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateCollection) SetDescription(v string) {
	o.Description = &v
}

// GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field value if set, zero value otherwise.
func (o *UpdateCollection) GetSubscribedApplicationsIds() []int32 {
	if o == nil || o.SubscribedApplicationsIds == nil {
		var ret []int32
		return ret
	}
	return o.SubscribedApplicationsIds
}

// GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCollection) GetSubscribedApplicationsIdsOk() ([]int32, bool) {
	if o == nil || o.SubscribedApplicationsIds == nil {
		return nil, false
	}
	return o.SubscribedApplicationsIds, true
}

// HasSubscribedApplicationsIds returns a boolean if a field has been set.
func (o *UpdateCollection) HasSubscribedApplicationsIds() bool {
	if o != nil && o.SubscribedApplicationsIds != nil {
		return true
	}

	return false
}

// SetSubscribedApplicationsIds gets a reference to the given []int32 and assigns it to the SubscribedApplicationsIds field.
func (o *UpdateCollection) SetSubscribedApplicationsIds(v []int32) {
	o.SubscribedApplicationsIds = v
}

func (o UpdateCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.SubscribedApplicationsIds != nil {
		toSerialize["subscribedApplicationsIds"] = o.SubscribedApplicationsIds
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateCollection struct {
	value *UpdateCollection
	isSet bool
}

func (v NullableUpdateCollection) Get() *UpdateCollection {
	return v.value
}

func (v *NullableUpdateCollection) Set(val *UpdateCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCollection(val *UpdateCollection) *NullableUpdateCollection {
	return &NullableUpdateCollection{value: val, isSet: true}
}

func (v NullableUpdateCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


