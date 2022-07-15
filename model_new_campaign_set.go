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

// NewCampaignSet 
type NewCampaignSet struct {
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// Version of the campaign set.
	Version int32 `json:"version"`
	Set CampaignSetBranchNode `json:"set"`
}

// NewNewCampaignSet instantiates a new NewCampaignSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewCampaignSet(applicationId int32, version int32, set CampaignSetBranchNode) *NewCampaignSet {
	this := NewCampaignSet{}
	this.ApplicationId = applicationId
	this.Version = version
	this.Set = set
	return &this
}

// NewNewCampaignSetWithDefaults instantiates a new NewCampaignSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewCampaignSetWithDefaults() *NewCampaignSet {
	this := NewCampaignSet{}
	return &this
}

// GetApplicationId returns the ApplicationId field value
func (o *NewCampaignSet) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *NewCampaignSet) GetApplicationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *NewCampaignSet) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetVersion returns the Version field value
func (o *NewCampaignSet) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *NewCampaignSet) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *NewCampaignSet) SetVersion(v int32) {
	o.Version = v
}

// GetSet returns the Set field value
func (o *NewCampaignSet) GetSet() CampaignSetBranchNode {
	if o == nil {
		var ret CampaignSetBranchNode
		return ret
	}

	return o.Set
}

// GetSetOk returns a tuple with the Set field value
// and a boolean to check if the value has been set.
func (o *NewCampaignSet) GetSetOk() (*CampaignSetBranchNode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Set, true
}

// SetSet sets field value
func (o *NewCampaignSet) SetSet(v CampaignSetBranchNode) {
	o.Set = v
}

func (o NewCampaignSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["set"] = o.Set
	}
	return json.Marshal(toSerialize)
}

type NullableNewCampaignSet struct {
	value *NewCampaignSet
	isSet bool
}

func (v NullableNewCampaignSet) Get() *NewCampaignSet {
	return v.value
}

func (v *NullableNewCampaignSet) Set(val *NewCampaignSet) {
	v.value = val
	v.isSet = true
}

func (v NullableNewCampaignSet) IsSet() bool {
	return v.isSet
}

func (v *NullableNewCampaignSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewCampaignSet(val *NewCampaignSet) *NullableNewCampaignSet {
	return &NullableNewCampaignSet{value: val, isSet: true}
}

func (v NullableNewCampaignSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewCampaignSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


