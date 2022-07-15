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

// CampaignSet 
type CampaignSet struct {
	// Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// Version of the campaign set.
	Version int32 `json:"version"`
	Set CampaignSetBranchNode `json:"set"`
}

// NewCampaignSet instantiates a new CampaignSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignSet(id int32, created time.Time, applicationId int32, version int32, set CampaignSetBranchNode) *CampaignSet {
	this := CampaignSet{}
	this.Id = id
	this.Created = created
	this.ApplicationId = applicationId
	this.Version = version
	this.Set = set
	return &this
}

// NewCampaignSetWithDefaults instantiates a new CampaignSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignSetWithDefaults() *CampaignSet {
	this := CampaignSet{}
	return &this
}

// GetId returns the Id field value
func (o *CampaignSet) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CampaignSet) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CampaignSet) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *CampaignSet) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *CampaignSet) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *CampaignSet) SetCreated(v time.Time) {
	o.Created = v
}

// GetApplicationId returns the ApplicationId field value
func (o *CampaignSet) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *CampaignSet) GetApplicationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *CampaignSet) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetVersion returns the Version field value
func (o *CampaignSet) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CampaignSet) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *CampaignSet) SetVersion(v int32) {
	o.Version = v
}

// GetSet returns the Set field value
func (o *CampaignSet) GetSet() CampaignSetBranchNode {
	if o == nil {
		var ret CampaignSetBranchNode
		return ret
	}

	return o.Set
}

// GetSetOk returns a tuple with the Set field value
// and a boolean to check if the value has been set.
func (o *CampaignSet) GetSetOk() (*CampaignSetBranchNode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Set, true
}

// SetSet sets field value
func (o *CampaignSet) SetSet(v CampaignSetBranchNode) {
	o.Set = v
}

func (o CampaignSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created"] = o.Created
	}
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

type NullableCampaignSet struct {
	value *CampaignSet
	isSet bool
}

func (v NullableCampaignSet) Get() *CampaignSet {
	return v.value
}

func (v *NullableCampaignSet) Set(val *CampaignSet) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignSet) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignSet(val *CampaignSet) *NullableCampaignSet {
	return &NullableCampaignSet{value: val, isSet: true}
}

func (v NullableCampaignSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


