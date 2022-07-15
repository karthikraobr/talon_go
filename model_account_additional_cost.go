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

// AccountAdditionalCost 
type AccountAdditionalCost struct {
	// Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// The additional cost name that will be used in API requests and Talang. E.g. if `name == \"shipping\"` then you would set the shipping additional cost by including an `additionalCosts.shipping` property in your request payload.
	Name string `json:"name"`
	// The human-readable name for the additional cost that will be shown in the Campaign Manager. Like `name`, the combination of entity and title must also be unique.
	Title string `json:"title"`
	// A description of this additional cost.
	Description string `json:"description"`
	// A list of the IDs of the applications that are subscribed to this additional cost.
	SubscribedApplicationsIds []int32 `json:"subscribedApplicationsIds,omitempty"`
	// The type of additional cost. The following options can be chosen: - `session`: Additional cost will be added per session, - `item`: Additional cost will be added per item, - `both`: Additional cost will be added per item and session. 
	Type *string `json:"type,omitempty"`
}

// NewAccountAdditionalCost instantiates a new AccountAdditionalCost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAdditionalCost(id int32, created time.Time, accountId int32, name string, title string, description string) *AccountAdditionalCost {
	this := AccountAdditionalCost{}
	this.Id = id
	this.Created = created
	this.AccountId = accountId
	this.Name = name
	this.Title = title
	this.Description = description
	var type_ string = "session"
	this.Type = &type_
	return &this
}

// NewAccountAdditionalCostWithDefaults instantiates a new AccountAdditionalCost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAdditionalCostWithDefaults() *AccountAdditionalCost {
	this := AccountAdditionalCost{}
	var type_ string = "session"
	this.Type = &type_
	return &this
}

// GetId returns the Id field value
func (o *AccountAdditionalCost) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountAdditionalCost) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountAdditionalCost) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *AccountAdditionalCost) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *AccountAdditionalCost) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *AccountAdditionalCost) SetCreated(v time.Time) {
	o.Created = v
}

// GetAccountId returns the AccountId field value
func (o *AccountAdditionalCost) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AccountAdditionalCost) GetAccountIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AccountAdditionalCost) SetAccountId(v int32) {
	o.AccountId = v
}

// GetName returns the Name field value
func (o *AccountAdditionalCost) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountAdditionalCost) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountAdditionalCost) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *AccountAdditionalCost) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AccountAdditionalCost) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *AccountAdditionalCost) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *AccountAdditionalCost) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AccountAdditionalCost) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AccountAdditionalCost) SetDescription(v string) {
	o.Description = v
}

// GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field value if set, zero value otherwise.
func (o *AccountAdditionalCost) GetSubscribedApplicationsIds() []int32 {
	if o == nil || o.SubscribedApplicationsIds == nil {
		var ret []int32
		return ret
	}
	return o.SubscribedApplicationsIds
}

// GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAdditionalCost) GetSubscribedApplicationsIdsOk() ([]int32, bool) {
	if o == nil || o.SubscribedApplicationsIds == nil {
		return nil, false
	}
	return o.SubscribedApplicationsIds, true
}

// HasSubscribedApplicationsIds returns a boolean if a field has been set.
func (o *AccountAdditionalCost) HasSubscribedApplicationsIds() bool {
	if o != nil && o.SubscribedApplicationsIds != nil {
		return true
	}

	return false
}

// SetSubscribedApplicationsIds gets a reference to the given []int32 and assigns it to the SubscribedApplicationsIds field.
func (o *AccountAdditionalCost) SetSubscribedApplicationsIds(v []int32) {
	o.SubscribedApplicationsIds = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccountAdditionalCost) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAdditionalCost) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccountAdditionalCost) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccountAdditionalCost) SetType(v string) {
	o.Type = &v
}

func (o AccountAdditionalCost) MarshalJSON() ([]byte, error) {
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
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.SubscribedApplicationsIds != nil {
		toSerialize["subscribedApplicationsIds"] = o.SubscribedApplicationsIds
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableAccountAdditionalCost struct {
	value *AccountAdditionalCost
	isSet bool
}

func (v NullableAccountAdditionalCost) Get() *AccountAdditionalCost {
	return v.value
}

func (v *NullableAccountAdditionalCost) Set(val *AccountAdditionalCost) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAdditionalCost) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAdditionalCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAdditionalCost(val *AccountAdditionalCost) *NullableAccountAdditionalCost {
	return &NullableAccountAdditionalCost{value: val, isSet: true}
}

func (v NullableAccountAdditionalCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAdditionalCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


