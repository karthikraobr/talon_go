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

// NewAccount struct for NewAccount
type NewAccount struct {
	CompanyName string `json:"companyName"`
}

// NewNewAccount instantiates a new NewAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewAccount(companyName string) *NewAccount {
	this := NewAccount{}
	this.CompanyName = companyName
	return &this
}

// NewNewAccountWithDefaults instantiates a new NewAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewAccountWithDefaults() *NewAccount {
	this := NewAccount{}
	return &this
}

// GetCompanyName returns the CompanyName field value
func (o *NewAccount) GetCompanyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value
// and a boolean to check if the value has been set.
func (o *NewAccount) GetCompanyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyName, true
}

// SetCompanyName sets field value
func (o *NewAccount) SetCompanyName(v string) {
	o.CompanyName = v
}

func (o NewAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["companyName"] = o.CompanyName
	}
	return json.Marshal(toSerialize)
}

type NullableNewAccount struct {
	value *NewAccount
	isSet bool
}

func (v NullableNewAccount) Get() *NewAccount {
	return v.value
}

func (v *NullableNewAccount) Set(val *NewAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableNewAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableNewAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewAccount(val *NewAccount) *NullableNewAccount {
	return &NullableNewAccount{value: val, isSet: true}
}

func (v NullableNewAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


