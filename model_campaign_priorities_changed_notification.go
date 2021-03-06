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

// CampaignPrioritiesChangedNotification Notification about an Application whose campaigns' priorities changed.
type CampaignPrioritiesChangedNotification struct {
	Application Application `json:"application"`
	// Campaign IDs for each priority. The priority can be one of: ['universal', 'stackable', 'exclusive'] 
	OldPriorities *map[string][]int32 `json:"oldPriorities,omitempty"`
	// Campaign IDs for each priority. The priority can be one of: ['universal', 'stackable', 'exclusive'] 
	Priorities map[string][]int32 `json:"priorities"`
}

// NewCampaignPrioritiesChangedNotification instantiates a new CampaignPrioritiesChangedNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignPrioritiesChangedNotification(application Application, priorities map[string][]int32) *CampaignPrioritiesChangedNotification {
	this := CampaignPrioritiesChangedNotification{}
	this.Application = application
	this.Priorities = priorities
	return &this
}

// NewCampaignPrioritiesChangedNotificationWithDefaults instantiates a new CampaignPrioritiesChangedNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignPrioritiesChangedNotificationWithDefaults() *CampaignPrioritiesChangedNotification {
	this := CampaignPrioritiesChangedNotification{}
	return &this
}

// GetApplication returns the Application field value
func (o *CampaignPrioritiesChangedNotification) GetApplication() Application {
	if o == nil {
		var ret Application
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *CampaignPrioritiesChangedNotification) GetApplicationOk() (*Application, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value
func (o *CampaignPrioritiesChangedNotification) SetApplication(v Application) {
	o.Application = v
}

// GetOldPriorities returns the OldPriorities field value if set, zero value otherwise.
func (o *CampaignPrioritiesChangedNotification) GetOldPriorities() map[string][]int32 {
	if o == nil || o.OldPriorities == nil {
		var ret map[string][]int32
		return ret
	}
	return *o.OldPriorities
}

// GetOldPrioritiesOk returns a tuple with the OldPriorities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignPrioritiesChangedNotification) GetOldPrioritiesOk() (*map[string][]int32, bool) {
	if o == nil || o.OldPriorities == nil {
		return nil, false
	}
	return o.OldPriorities, true
}

// HasOldPriorities returns a boolean if a field has been set.
func (o *CampaignPrioritiesChangedNotification) HasOldPriorities() bool {
	if o != nil && o.OldPriorities != nil {
		return true
	}

	return false
}

// SetOldPriorities gets a reference to the given map[string][]int32 and assigns it to the OldPriorities field.
func (o *CampaignPrioritiesChangedNotification) SetOldPriorities(v map[string][]int32) {
	o.OldPriorities = &v
}

// GetPriorities returns the Priorities field value
func (o *CampaignPrioritiesChangedNotification) GetPriorities() map[string][]int32 {
	if o == nil {
		var ret map[string][]int32
		return ret
	}

	return o.Priorities
}

// GetPrioritiesOk returns a tuple with the Priorities field value
// and a boolean to check if the value has been set.
func (o *CampaignPrioritiesChangedNotification) GetPrioritiesOk() (*map[string][]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priorities, true
}

// SetPriorities sets field value
func (o *CampaignPrioritiesChangedNotification) SetPriorities(v map[string][]int32) {
	o.Priorities = v
}

func (o CampaignPrioritiesChangedNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["application"] = o.Application
	}
	if o.OldPriorities != nil {
		toSerialize["oldPriorities"] = o.OldPriorities
	}
	if true {
		toSerialize["priorities"] = o.Priorities
	}
	return json.Marshal(toSerialize)
}

type NullableCampaignPrioritiesChangedNotification struct {
	value *CampaignPrioritiesChangedNotification
	isSet bool
}

func (v NullableCampaignPrioritiesChangedNotification) Get() *CampaignPrioritiesChangedNotification {
	return v.value
}

func (v *NullableCampaignPrioritiesChangedNotification) Set(val *CampaignPrioritiesChangedNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignPrioritiesChangedNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignPrioritiesChangedNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignPrioritiesChangedNotification(val *CampaignPrioritiesChangedNotification) *NullableCampaignPrioritiesChangedNotification {
	return &NullableCampaignPrioritiesChangedNotification{value: val, isSet: true}
}

func (v NullableCampaignPrioritiesChangedNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignPrioritiesChangedNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


