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

// CampaignEntity struct for CampaignEntity
type CampaignEntity struct {
	// The ID of the campaign that owns this entity.
	CampaignId int32 `json:"campaignId"`
}

// NewCampaignEntity instantiates a new CampaignEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignEntity(campaignId int32) *CampaignEntity {
	this := CampaignEntity{}
	this.CampaignId = campaignId
	return &this
}

// NewCampaignEntityWithDefaults instantiates a new CampaignEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignEntityWithDefaults() *CampaignEntity {
	this := CampaignEntity{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *CampaignEntity) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *CampaignEntity) GetCampaignIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *CampaignEntity) SetCampaignId(v int32) {
	o.CampaignId = v
}

func (o CampaignEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["campaignId"] = o.CampaignId
	}
	return json.Marshal(toSerialize)
}

type NullableCampaignEntity struct {
	value *CampaignEntity
	isSet bool
}

func (v NullableCampaignEntity) Get() *CampaignEntity {
	return v.value
}

func (v *NullableCampaignEntity) Set(val *CampaignEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignEntity(val *CampaignEntity) *NullableCampaignEntity {
	return &NullableCampaignEntity{value: val, isSet: true}
}

func (v NullableCampaignEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


