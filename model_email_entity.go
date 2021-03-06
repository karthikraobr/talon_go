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

// EmailEntity struct for EmailEntity
type EmailEntity struct {
	// The email address associated with your account.
	Email string `json:"email"`
}

// NewEmailEntity instantiates a new EmailEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailEntity(email string) *EmailEntity {
	this := EmailEntity{}
	this.Email = email
	return &this
}

// NewEmailEntityWithDefaults instantiates a new EmailEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailEntityWithDefaults() *EmailEntity {
	this := EmailEntity{}
	return &this
}

// GetEmail returns the Email field value
func (o *EmailEntity) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *EmailEntity) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *EmailEntity) SetEmail(v string) {
	o.Email = v
}

func (o EmailEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableEmailEntity struct {
	value *EmailEntity
	isSet bool
}

func (v NullableEmailEntity) Get() *EmailEntity {
	return v.value
}

func (v *NullableEmailEntity) Set(val *EmailEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailEntity(val *EmailEntity) *NullableEmailEntity {
	return &NullableEmailEntity{value: val, isSet: true}
}

func (v NullableEmailEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


