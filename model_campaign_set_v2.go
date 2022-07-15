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

// CampaignSetV2 
type CampaignSetV2 struct {
	// Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// Version of the campaign set.
	Version int32 `json:"version"`
	Set CampaignPrioritiesV2 `json:"set"`
}

// NewCampaignSetV2 instantiates a new CampaignSetV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignSetV2(id int32, created time.Time, applicationId int32, version int32, set CampaignPrioritiesV2) *CampaignSetV2 {
	this := CampaignSetV2{}
	this.Id = id
	this.Created = created
	this.ApplicationId = applicationId
	this.Version = version
	this.Set = set
	return &this
}

// NewCampaignSetV2WithDefaults instantiates a new CampaignSetV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignSetV2WithDefaults() *CampaignSetV2 {
	this := CampaignSetV2{}
	return &this
}

// GetId returns the Id field value
func (o *CampaignSetV2) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CampaignSetV2) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CampaignSetV2) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *CampaignSetV2) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *CampaignSetV2) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *CampaignSetV2) SetCreated(v time.Time) {
	o.Created = v
}

// GetApplicationId returns the ApplicationId field value
func (o *CampaignSetV2) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *CampaignSetV2) GetApplicationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *CampaignSetV2) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetVersion returns the Version field value
func (o *CampaignSetV2) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CampaignSetV2) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *CampaignSetV2) SetVersion(v int32) {
	o.Version = v
}

// GetSet returns the Set field value
func (o *CampaignSetV2) GetSet() CampaignPrioritiesV2 {
	if o == nil {
		var ret CampaignPrioritiesV2
		return ret
	}

	return o.Set
}

// GetSetOk returns a tuple with the Set field value
// and a boolean to check if the value has been set.
func (o *CampaignSetV2) GetSetOk() (*CampaignPrioritiesV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Set, true
}

// SetSet sets field value
func (o *CampaignSetV2) SetSet(v CampaignPrioritiesV2) {
	o.Set = v
}

func (o CampaignSetV2) MarshalJSON() ([]byte, error) {
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

type NullableCampaignSetV2 struct {
	value *CampaignSetV2
	isSet bool
}

func (v NullableCampaignSetV2) Get() *CampaignSetV2 {
	return v.value
}

func (v *NullableCampaignSetV2) Set(val *CampaignSetV2) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignSetV2) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignSetV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignSetV2(val *CampaignSetV2) *NullableCampaignSetV2 {
	return &NullableCampaignSetV2{value: val, isSet: true}
}

func (v NullableCampaignSetV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignSetV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

