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

// CustomEffect 
type CustomEffect struct {
	// Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// The exact moment this entity was last modified.
	Modified time.Time `json:"modified"`
	// The IDs of the applications that are related to this entity.
	ApplicationIds []int32 `json:"applicationIds"`
	// The name of this effect.
	Name string `json:"name"`
	// The title of this effect.
	Title string `json:"title"`
	// The JSON payload of this effect.
	Payload string `json:"payload"`
	// The description of this effect.
	Description *string `json:"description,omitempty"`
	// Determines if this effect is active.
	Enabled bool `json:"enabled"`
	// Array of template argument definitions.
	Params []TemplateArgDef `json:"params,omitempty"`
	// ID of the user who last updated this effect if available.
	ModifiedBy *int32 `json:"modifiedBy,omitempty"`
	// ID of the user who created this effect.
	CreatedBy int32 `json:"createdBy"`
}

// NewCustomEffect instantiates a new CustomEffect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEffect(id int32, created time.Time, accountId int32, modified time.Time, applicationIds []int32, name string, title string, payload string, enabled bool, createdBy int32) *CustomEffect {
	this := CustomEffect{}
	this.Id = id
	this.Created = created
	this.AccountId = accountId
	this.Modified = modified
	this.ApplicationIds = applicationIds
	this.Name = name
	this.Title = title
	this.Payload = payload
	this.Enabled = enabled
	this.CreatedBy = createdBy
	return &this
}

// NewCustomEffectWithDefaults instantiates a new CustomEffect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEffectWithDefaults() *CustomEffect {
	this := CustomEffect{}
	return &this
}

// GetId returns the Id field value
func (o *CustomEffect) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomEffect) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomEffect) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *CustomEffect) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *CustomEffect) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *CustomEffect) SetCreated(v time.Time) {
	o.Created = v
}

// GetAccountId returns the AccountId field value
func (o *CustomEffect) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CustomEffect) GetAccountIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *CustomEffect) SetAccountId(v int32) {
	o.AccountId = v
}

// GetModified returns the Modified field value
func (o *CustomEffect) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *CustomEffect) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *CustomEffect) SetModified(v time.Time) {
	o.Modified = v
}

// GetApplicationIds returns the ApplicationIds field value
func (o *CustomEffect) GetApplicationIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ApplicationIds
}

// GetApplicationIdsOk returns a tuple with the ApplicationIds field value
// and a boolean to check if the value has been set.
func (o *CustomEffect) GetApplicationIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationIds, true
}

// SetApplicationIds sets field value
func (o *CustomEffect) SetApplicationIds(v []int32) {
	o.ApplicationIds = v
}

// GetName returns the Name field value
func (o *CustomEffect) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomEffect) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomEffect) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *CustomEffect) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CustomEffect) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CustomEffect) SetTitle(v string) {
	o.Title = v
}

// GetPayload returns the Payload field value
func (o *CustomEffect) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *CustomEffect) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *CustomEffect) SetPayload(v string) {
	o.Payload = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomEffect) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEffect) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomEffect) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomEffect) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *CustomEffect) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CustomEffect) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CustomEffect) SetEnabled(v bool) {
	o.Enabled = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *CustomEffect) GetParams() []TemplateArgDef {
	if o == nil || o.Params == nil {
		var ret []TemplateArgDef
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEffect) GetParamsOk() ([]TemplateArgDef, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *CustomEffect) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given []TemplateArgDef and assigns it to the Params field.
func (o *CustomEffect) SetParams(v []TemplateArgDef) {
	o.Params = v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *CustomEffect) GetModifiedBy() int32 {
	if o == nil || o.ModifiedBy == nil {
		var ret int32
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEffect) GetModifiedByOk() (*int32, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *CustomEffect) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given int32 and assigns it to the ModifiedBy field.
func (o *CustomEffect) SetModifiedBy(v int32) {
	o.ModifiedBy = &v
}

// GetCreatedBy returns the CreatedBy field value
func (o *CustomEffect) GetCreatedBy() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *CustomEffect) GetCreatedByOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *CustomEffect) SetCreatedBy(v int32) {
	o.CreatedBy = v
}

func (o CustomEffect) MarshalJSON() ([]byte, error) {
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
		toSerialize["modified"] = o.Modified
	}
	if true {
		toSerialize["applicationIds"] = o.ApplicationIds
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["payload"] = o.Payload
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	if o.ModifiedBy != nil {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	return json.Marshal(toSerialize)
}

type NullableCustomEffect struct {
	value *CustomEffect
	isSet bool
}

func (v NullableCustomEffect) Get() *CustomEffect {
	return v.value
}

func (v *NullableCustomEffect) Set(val *CustomEffect) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEffect) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEffect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEffect(val *CustomEffect) *NullableCustomEffect {
	return &NullableCustomEffect{value: val, isSet: true}
}

func (v NullableCustomEffect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEffect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


