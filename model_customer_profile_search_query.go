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

// CustomerProfileSearchQuery struct for CustomerProfileSearchQuery
type CustomerProfileSearchQuery struct {
	// Properties to match against a customer profile. All provided attributes will be exactly matched against profile attributes.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	IntegrationIDs []string `json:"integrationIDs,omitempty"`
	ProfileIDs []int32 `json:"profileIDs,omitempty"`
}

// NewCustomerProfileSearchQuery instantiates a new CustomerProfileSearchQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerProfileSearchQuery() *CustomerProfileSearchQuery {
	this := CustomerProfileSearchQuery{}
	return &this
}

// NewCustomerProfileSearchQueryWithDefaults instantiates a new CustomerProfileSearchQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerProfileSearchQueryWithDefaults() *CustomerProfileSearchQuery {
	this := CustomerProfileSearchQuery{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CustomerProfileSearchQuery) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerProfileSearchQuery) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CustomerProfileSearchQuery) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *CustomerProfileSearchQuery) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetIntegrationIDs returns the IntegrationIDs field value if set, zero value otherwise.
func (o *CustomerProfileSearchQuery) GetIntegrationIDs() []string {
	if o == nil || o.IntegrationIDs == nil {
		var ret []string
		return ret
	}
	return o.IntegrationIDs
}

// GetIntegrationIDsOk returns a tuple with the IntegrationIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerProfileSearchQuery) GetIntegrationIDsOk() ([]string, bool) {
	if o == nil || o.IntegrationIDs == nil {
		return nil, false
	}
	return o.IntegrationIDs, true
}

// HasIntegrationIDs returns a boolean if a field has been set.
func (o *CustomerProfileSearchQuery) HasIntegrationIDs() bool {
	if o != nil && o.IntegrationIDs != nil {
		return true
	}

	return false
}

// SetIntegrationIDs gets a reference to the given []string and assigns it to the IntegrationIDs field.
func (o *CustomerProfileSearchQuery) SetIntegrationIDs(v []string) {
	o.IntegrationIDs = v
}

// GetProfileIDs returns the ProfileIDs field value if set, zero value otherwise.
func (o *CustomerProfileSearchQuery) GetProfileIDs() []int32 {
	if o == nil || o.ProfileIDs == nil {
		var ret []int32
		return ret
	}
	return o.ProfileIDs
}

// GetProfileIDsOk returns a tuple with the ProfileIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerProfileSearchQuery) GetProfileIDsOk() ([]int32, bool) {
	if o == nil || o.ProfileIDs == nil {
		return nil, false
	}
	return o.ProfileIDs, true
}

// HasProfileIDs returns a boolean if a field has been set.
func (o *CustomerProfileSearchQuery) HasProfileIDs() bool {
	if o != nil && o.ProfileIDs != nil {
		return true
	}

	return false
}

// SetProfileIDs gets a reference to the given []int32 and assigns it to the ProfileIDs field.
func (o *CustomerProfileSearchQuery) SetProfileIDs(v []int32) {
	o.ProfileIDs = v
}

func (o CustomerProfileSearchQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.IntegrationIDs != nil {
		toSerialize["integrationIDs"] = o.IntegrationIDs
	}
	if o.ProfileIDs != nil {
		toSerialize["profileIDs"] = o.ProfileIDs
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerProfileSearchQuery struct {
	value *CustomerProfileSearchQuery
	isSet bool
}

func (v NullableCustomerProfileSearchQuery) Get() *CustomerProfileSearchQuery {
	return v.value
}

func (v *NullableCustomerProfileSearchQuery) Set(val *CustomerProfileSearchQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerProfileSearchQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerProfileSearchQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerProfileSearchQuery(val *CustomerProfileSearchQuery) *NullableCustomerProfileSearchQuery {
	return &NullableCustomerProfileSearchQuery{value: val, isSet: true}
}

func (v NullableCustomerProfileSearchQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerProfileSearchQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

