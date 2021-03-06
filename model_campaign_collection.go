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

// CampaignCollection 
type CampaignCollection struct {
	// Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// The exact moment this entity was last modified.
	Modified time.Time `json:"modified"`
	// A short description of the purpose of this collection.
	Description *string `json:"description,omitempty"`
	// The name of this collection.
	Name string `json:"name"`
	// ID of the user who last updated this effect if available.
	ModifiedBy *int32 `json:"modifiedBy,omitempty"`
	// ID of the user who created this effect.
	CreatedBy int32 `json:"createdBy"`
	// The ID of the Application that owns this entity.
	ApplicationId *int32 `json:"applicationId,omitempty"`
	// The ID of the campaign that owns this entity.
	CampaignId *int32 `json:"campaignId,omitempty"`
	// The content of the collection.
	Payload []string `json:"payload,omitempty"`
}

// NewCampaignCollection instantiates a new CampaignCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignCollection(id int32, created time.Time, accountId int32, modified time.Time, name string, createdBy int32) *CampaignCollection {
	this := CampaignCollection{}
	this.Id = id
	this.Created = created
	this.AccountId = accountId
	this.Modified = modified
	this.Name = name
	this.CreatedBy = createdBy
	return &this
}

// NewCampaignCollectionWithDefaults instantiates a new CampaignCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignCollectionWithDefaults() *CampaignCollection {
	this := CampaignCollection{}
	return &this
}

// GetId returns the Id field value
func (o *CampaignCollection) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CampaignCollection) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CampaignCollection) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *CampaignCollection) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *CampaignCollection) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *CampaignCollection) SetCreated(v time.Time) {
	o.Created = v
}

// GetAccountId returns the AccountId field value
func (o *CampaignCollection) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CampaignCollection) GetAccountIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *CampaignCollection) SetAccountId(v int32) {
	o.AccountId = v
}

// GetModified returns the Modified field value
func (o *CampaignCollection) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *CampaignCollection) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *CampaignCollection) SetModified(v time.Time) {
	o.Modified = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CampaignCollection) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCollection) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CampaignCollection) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CampaignCollection) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *CampaignCollection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CampaignCollection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CampaignCollection) SetName(v string) {
	o.Name = v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *CampaignCollection) GetModifiedBy() int32 {
	if o == nil || o.ModifiedBy == nil {
		var ret int32
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCollection) GetModifiedByOk() (*int32, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *CampaignCollection) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given int32 and assigns it to the ModifiedBy field.
func (o *CampaignCollection) SetModifiedBy(v int32) {
	o.ModifiedBy = &v
}

// GetCreatedBy returns the CreatedBy field value
func (o *CampaignCollection) GetCreatedBy() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *CampaignCollection) GetCreatedByOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *CampaignCollection) SetCreatedBy(v int32) {
	o.CreatedBy = v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *CampaignCollection) GetApplicationId() int32 {
	if o == nil || o.ApplicationId == nil {
		var ret int32
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCollection) GetApplicationIdOk() (*int32, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *CampaignCollection) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.
func (o *CampaignCollection) SetApplicationId(v int32) {
	o.ApplicationId = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *CampaignCollection) GetCampaignId() int32 {
	if o == nil || o.CampaignId == nil {
		var ret int32
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCollection) GetCampaignIdOk() (*int32, bool) {
	if o == nil || o.CampaignId == nil {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *CampaignCollection) HasCampaignId() bool {
	if o != nil && o.CampaignId != nil {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.
func (o *CampaignCollection) SetCampaignId(v int32) {
	o.CampaignId = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *CampaignCollection) GetPayload() []string {
	if o == nil || o.Payload == nil {
		var ret []string
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCollection) GetPayloadOk() ([]string, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *CampaignCollection) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given []string and assigns it to the Payload field.
func (o *CampaignCollection) SetPayload(v []string) {
	o.Payload = v
}

func (o CampaignCollection) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ModifiedBy != nil {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.ApplicationId != nil {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if o.CampaignId != nil {
		toSerialize["campaignId"] = o.CampaignId
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	return json.Marshal(toSerialize)
}

type NullableCampaignCollection struct {
	value *CampaignCollection
	isSet bool
}

func (v NullableCampaignCollection) Get() *CampaignCollection {
	return v.value
}

func (v *NullableCampaignCollection) Set(val *CampaignCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignCollection(val *CampaignCollection) *NullableCampaignCollection {
	return &NullableCampaignCollection{value: val, isSet: true}
}

func (v NullableCampaignCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


