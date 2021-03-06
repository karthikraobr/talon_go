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

// ApplicationSessionEntity struct for ApplicationSessionEntity
type ApplicationSessionEntity struct {
	// The globally unique Talon.One ID of the session where this entity was created.
	SessionId int32 `json:"sessionId"`
}

// NewApplicationSessionEntity instantiates a new ApplicationSessionEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationSessionEntity(sessionId int32) *ApplicationSessionEntity {
	this := ApplicationSessionEntity{}
	this.SessionId = sessionId
	return &this
}

// NewApplicationSessionEntityWithDefaults instantiates a new ApplicationSessionEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationSessionEntityWithDefaults() *ApplicationSessionEntity {
	this := ApplicationSessionEntity{}
	return &this
}

// GetSessionId returns the SessionId field value
func (o *ApplicationSessionEntity) GetSessionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *ApplicationSessionEntity) GetSessionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *ApplicationSessionEntity) SetSessionId(v int32) {
	o.SessionId = v
}

func (o ApplicationSessionEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sessionId"] = o.SessionId
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationSessionEntity struct {
	value *ApplicationSessionEntity
	isSet bool
}

func (v NullableApplicationSessionEntity) Get() *ApplicationSessionEntity {
	return v.value
}

func (v *NullableApplicationSessionEntity) Set(val *ApplicationSessionEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationSessionEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationSessionEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationSessionEntity(val *ApplicationSessionEntity) *NullableApplicationSessionEntity {
	return &NullableApplicationSessionEntity{value: val, isSet: true}
}

func (v NullableApplicationSessionEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationSessionEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


