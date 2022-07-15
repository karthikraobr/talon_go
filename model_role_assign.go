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

// RoleAssign 
type RoleAssign struct {
	// An array of userIDs.
	Users []int32 `json:"users"`
	// An array of roleIDs.
	Roles []int32 `json:"roles"`
}

// NewRoleAssign instantiates a new RoleAssign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAssign(users []int32, roles []int32) *RoleAssign {
	this := RoleAssign{}
	this.Users = users
	this.Roles = roles
	return &this
}

// NewRoleAssignWithDefaults instantiates a new RoleAssign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAssignWithDefaults() *RoleAssign {
	this := RoleAssign{}
	return &this
}

// GetUsers returns the Users field value
func (o *RoleAssign) GetUsers() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *RoleAssign) GetUsersOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *RoleAssign) SetUsers(v []int32) {
	o.Users = v
}

// GetRoles returns the Roles field value
func (o *RoleAssign) GetRoles() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *RoleAssign) GetRolesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *RoleAssign) SetRoles(v []int32) {
	o.Roles = v
}

func (o RoleAssign) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["users"] = o.Users
	}
	if true {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableRoleAssign struct {
	value *RoleAssign
	isSet bool
}

func (v NullableRoleAssign) Get() *RoleAssign {
	return v.value
}

func (v *NullableRoleAssign) Set(val *RoleAssign) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssign) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssign(val *RoleAssign) *NullableRoleAssign {
	return &NullableRoleAssign{value: val, isSet: true}
}

func (v NullableRoleAssign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

