/*
Integration API reference docs

Use the Integration API to push data to and retrieve data from Talon.One in real time. For more background information about this API, see [Integration API overview](/docs/dev/integration-api/overview)  For example, use this API to share shopping cart information as a session with Talon.One and evaluate promotion rules. You can also create custom events to track specific actions that do not fit into the session data model.  Ensure you [authenticate](#section/Authentication) to make requests to the API.  <div class=\"redoc-section\">   <p class=\"title\">Are you looking for a different API?</p>    If you need the API to:    - Interact with the Campaign Manager for backoffice operations, see [the Management API reference docs](https://docs.talon.one/management-api).   - Integrate with Talon.One from a CEP or CDP platform, see [the Third-party API reference docs](https://docs.talon.one/third-party-api).  </div>  # Authentication  <SecurityDefinitions /> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package talon

import (
	"encoding/json"
	"time"
)

// Export 
type Export struct {
	// Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// The ID of the account that owns this entity.
	UserId int32 `json:"userId"`
	// The name of the entity that was exported.
	Entity string `json:"entity"`
	// Map of keys and values that were used to filter the exported rows.
	Filter map[string]interface{} `json:"filter"`
}

// NewExport instantiates a new Export object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExport(id int32, created time.Time, accountId int32, userId int32, entity string, filter map[string]interface{}) *Export {
	this := Export{}
	this.Id = id
	this.Created = created
	this.AccountId = accountId
	this.UserId = userId
	this.Entity = entity
	this.Filter = filter
	return &this
}

// NewExportWithDefaults instantiates a new Export object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportWithDefaults() *Export {
	this := Export{}
	return &this
}

// GetId returns the Id field value
func (o *Export) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Export) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Export) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Export) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Export) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Export) SetCreated(v time.Time) {
	o.Created = v
}

// GetAccountId returns the AccountId field value
func (o *Export) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Export) GetAccountIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Export) SetAccountId(v int32) {
	o.AccountId = v
}

// GetUserId returns the UserId field value
func (o *Export) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *Export) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *Export) SetUserId(v int32) {
	o.UserId = v
}

// GetEntity returns the Entity field value
func (o *Export) GetEntity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value
// and a boolean to check if the value has been set.
func (o *Export) GetEntityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entity, true
}

// SetEntity sets field value
func (o *Export) SetEntity(v string) {
	o.Entity = v
}

// GetFilter returns the Filter field value
func (o *Export) GetFilter() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *Export) GetFilterOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter, true
}

// SetFilter sets field value
func (o *Export) SetFilter(v map[string]interface{}) {
	o.Filter = v
}

func (o Export) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["userId"] = o.UserId
	}
	if true {
		toSerialize["entity"] = o.Entity
	}
	if true {
		toSerialize["filter"] = o.Filter
	}
	return json.Marshal(toSerialize)
}

type NullableExport struct {
	value *Export
	isSet bool
}

func (v NullableExport) Get() *Export {
	return v.value
}

func (v *NullableExport) Set(val *Export) {
	v.value = val
	v.isSet = true
}

func (v NullableExport) IsSet() bool {
	return v.isSet
}

func (v *NullableExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExport(val *Export) *NullableExport {
	return &NullableExport{value: val, isSet: true}
}

func (v NullableExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


