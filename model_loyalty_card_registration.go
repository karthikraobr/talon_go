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

// LoyaltyCardRegistration struct for LoyaltyCardRegistration
type LoyaltyCardRegistration struct {
	// The integrationId of the customer profile.
	IntegrationId string `json:"integrationId"`
}

// NewLoyaltyCardRegistration instantiates a new LoyaltyCardRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoyaltyCardRegistration(integrationId string) *LoyaltyCardRegistration {
	this := LoyaltyCardRegistration{}
	this.IntegrationId = integrationId
	return &this
}

// NewLoyaltyCardRegistrationWithDefaults instantiates a new LoyaltyCardRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoyaltyCardRegistrationWithDefaults() *LoyaltyCardRegistration {
	this := LoyaltyCardRegistration{}
	return &this
}

// GetIntegrationId returns the IntegrationId field value
func (o *LoyaltyCardRegistration) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *LoyaltyCardRegistration) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *LoyaltyCardRegistration) SetIntegrationId(v string) {
	o.IntegrationId = v
}

func (o LoyaltyCardRegistration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["integrationId"] = o.IntegrationId
	}
	return json.Marshal(toSerialize)
}

type NullableLoyaltyCardRegistration struct {
	value *LoyaltyCardRegistration
	isSet bool
}

func (v NullableLoyaltyCardRegistration) Get() *LoyaltyCardRegistration {
	return v.value
}

func (v *NullableLoyaltyCardRegistration) Set(val *LoyaltyCardRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableLoyaltyCardRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableLoyaltyCardRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoyaltyCardRegistration(val *LoyaltyCardRegistration) *NullableLoyaltyCardRegistration {
	return &NullableLoyaltyCardRegistration{value: val, isSet: true}
}

func (v NullableLoyaltyCardRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoyaltyCardRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

