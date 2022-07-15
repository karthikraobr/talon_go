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

// CampaignStateChangedNotification A notification regarding a campaign whose state changed.
type CampaignStateChangedNotification struct {
	Campaign Campaign `json:"campaign"`
	// The campaign's old state. Can be one of the following: ['running', 'disabled', 'scheduled', 'expired', 'draft', 'archived'] 
	OldState string `json:"oldState"`
	// The campaign's new state. Can be one of the following: ['running', 'disabled', 'scheduled', 'expired', 'draft', 'archived'] 
	NewState string `json:"newState"`
}

// NewCampaignStateChangedNotification instantiates a new CampaignStateChangedNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignStateChangedNotification(campaign Campaign, oldState string, newState string) *CampaignStateChangedNotification {
	this := CampaignStateChangedNotification{}
	this.Campaign = campaign
	this.OldState = oldState
	this.NewState = newState
	return &this
}

// NewCampaignStateChangedNotificationWithDefaults instantiates a new CampaignStateChangedNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignStateChangedNotificationWithDefaults() *CampaignStateChangedNotification {
	this := CampaignStateChangedNotification{}
	return &this
}

// GetCampaign returns the Campaign field value
func (o *CampaignStateChangedNotification) GetCampaign() Campaign {
	if o == nil {
		var ret Campaign
		return ret
	}

	return o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value
// and a boolean to check if the value has been set.
func (o *CampaignStateChangedNotification) GetCampaignOk() (*Campaign, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaign, true
}

// SetCampaign sets field value
func (o *CampaignStateChangedNotification) SetCampaign(v Campaign) {
	o.Campaign = v
}

// GetOldState returns the OldState field value
func (o *CampaignStateChangedNotification) GetOldState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OldState
}

// GetOldStateOk returns a tuple with the OldState field value
// and a boolean to check if the value has been set.
func (o *CampaignStateChangedNotification) GetOldStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldState, true
}

// SetOldState sets field value
func (o *CampaignStateChangedNotification) SetOldState(v string) {
	o.OldState = v
}

// GetNewState returns the NewState field value
func (o *CampaignStateChangedNotification) GetNewState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewState
}

// GetNewStateOk returns a tuple with the NewState field value
// and a boolean to check if the value has been set.
func (o *CampaignStateChangedNotification) GetNewStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewState, true
}

// SetNewState sets field value
func (o *CampaignStateChangedNotification) SetNewState(v string) {
	o.NewState = v
}

func (o CampaignStateChangedNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["campaign"] = o.Campaign
	}
	if true {
		toSerialize["oldState"] = o.OldState
	}
	if true {
		toSerialize["newState"] = o.NewState
	}
	return json.Marshal(toSerialize)
}

type NullableCampaignStateChangedNotification struct {
	value *CampaignStateChangedNotification
	isSet bool
}

func (v NullableCampaignStateChangedNotification) Get() *CampaignStateChangedNotification {
	return v.value
}

func (v *NullableCampaignStateChangedNotification) Set(val *CampaignStateChangedNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignStateChangedNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignStateChangedNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignStateChangedNotification(val *CampaignStateChangedNotification) *NullableCampaignStateChangedNotification {
	return &NullableCampaignStateChangedNotification{value: val, isSet: true}
}

func (v NullableCampaignStateChangedNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignStateChangedNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

