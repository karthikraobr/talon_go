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

// MultipleCustomerProfileIntegrationResponseV2 struct for MultipleCustomerProfileIntegrationResponseV2
type MultipleCustomerProfileIntegrationResponseV2 struct {
	IntegrationStates []IntegrationStateV2 `json:"integrationStates,omitempty"`
}

// NewMultipleCustomerProfileIntegrationResponseV2 instantiates a new MultipleCustomerProfileIntegrationResponseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleCustomerProfileIntegrationResponseV2() *MultipleCustomerProfileIntegrationResponseV2 {
	this := MultipleCustomerProfileIntegrationResponseV2{}
	return &this
}

// NewMultipleCustomerProfileIntegrationResponseV2WithDefaults instantiates a new MultipleCustomerProfileIntegrationResponseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleCustomerProfileIntegrationResponseV2WithDefaults() *MultipleCustomerProfileIntegrationResponseV2 {
	this := MultipleCustomerProfileIntegrationResponseV2{}
	return &this
}

// GetIntegrationStates returns the IntegrationStates field value if set, zero value otherwise.
func (o *MultipleCustomerProfileIntegrationResponseV2) GetIntegrationStates() []IntegrationStateV2 {
	if o == nil || o.IntegrationStates == nil {
		var ret []IntegrationStateV2
		return ret
	}
	return o.IntegrationStates
}

// GetIntegrationStatesOk returns a tuple with the IntegrationStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleCustomerProfileIntegrationResponseV2) GetIntegrationStatesOk() ([]IntegrationStateV2, bool) {
	if o == nil || o.IntegrationStates == nil {
		return nil, false
	}
	return o.IntegrationStates, true
}

// HasIntegrationStates returns a boolean if a field has been set.
func (o *MultipleCustomerProfileIntegrationResponseV2) HasIntegrationStates() bool {
	if o != nil && o.IntegrationStates != nil {
		return true
	}

	return false
}

// SetIntegrationStates gets a reference to the given []IntegrationStateV2 and assigns it to the IntegrationStates field.
func (o *MultipleCustomerProfileIntegrationResponseV2) SetIntegrationStates(v []IntegrationStateV2) {
	o.IntegrationStates = v
}

func (o MultipleCustomerProfileIntegrationResponseV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IntegrationStates != nil {
		toSerialize["integrationStates"] = o.IntegrationStates
	}
	return json.Marshal(toSerialize)
}

type NullableMultipleCustomerProfileIntegrationResponseV2 struct {
	value *MultipleCustomerProfileIntegrationResponseV2
	isSet bool
}

func (v NullableMultipleCustomerProfileIntegrationResponseV2) Get() *MultipleCustomerProfileIntegrationResponseV2 {
	return v.value
}

func (v *NullableMultipleCustomerProfileIntegrationResponseV2) Set(val *MultipleCustomerProfileIntegrationResponseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleCustomerProfileIntegrationResponseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleCustomerProfileIntegrationResponseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleCustomerProfileIntegrationResponseV2(val *MultipleCustomerProfileIntegrationResponseV2) *NullableMultipleCustomerProfileIntegrationResponseV2 {
	return &NullableMultipleCustomerProfileIntegrationResponseV2{value: val, isSet: true}
}

func (v NullableMultipleCustomerProfileIntegrationResponseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleCustomerProfileIntegrationResponseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


