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

// Referral 
type Referral struct {
	// Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// Timestamp at which point the referral code becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Expiry date of the referral code. Referral never expires if this is omitted, zero, or negative.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// The number of times a referral code can be used. `0` means no limit but any campaign usage limits will still apply. 
	UsageLimit int32 `json:"usageLimit"`
	// ID of the campaign from which the referral received the referral code.
	CampaignId int32 `json:"campaignId"`
	// The Integration ID of the Advocate's Profile.
	AdvocateProfileIntegrationId string `json:"advocateProfileIntegrationId"`
	// An optional Integration ID of the Friend's Profile.
	FriendProfileIntegrationId *string `json:"friendProfileIntegrationId,omitempty"`
	// Arbitrary properties associated with this item.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// The ID of the Import which created this referral.
	ImportId *int32 `json:"importId,omitempty"`
	// The referral code.
	Code string `json:"code"`
	// The number of times this referral code has been successfully used.
	UsageCounter int32 `json:"usageCounter"`
	// The ID of the batch the referrals belong to.
	BatchId *string `json:"batchId,omitempty"`
}

// NewReferral instantiates a new Referral object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferral(id int32, created time.Time, usageLimit int32, campaignId int32, advocateProfileIntegrationId string, code string, usageCounter int32) *Referral {
	this := Referral{}
	this.Id = id
	this.Created = created
	this.UsageLimit = usageLimit
	this.CampaignId = campaignId
	this.AdvocateProfileIntegrationId = advocateProfileIntegrationId
	this.Code = code
	this.UsageCounter = usageCounter
	return &this
}

// NewReferralWithDefaults instantiates a new Referral object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferralWithDefaults() *Referral {
	this := Referral{}
	return &this
}

// GetId returns the Id field value
func (o *Referral) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Referral) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Referral) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Referral) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Referral) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Referral) SetCreated(v time.Time) {
	o.Created = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Referral) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Referral) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Referral) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *Referral) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *Referral) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Referral) GetExpiryDateOk() (*time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *Referral) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *Referral) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetUsageLimit returns the UsageLimit field value
func (o *Referral) GetUsageLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageLimit
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value
// and a boolean to check if the value has been set.
func (o *Referral) GetUsageLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageLimit, true
}

// SetUsageLimit sets field value
func (o *Referral) SetUsageLimit(v int32) {
	o.UsageLimit = v
}

// GetCampaignId returns the CampaignId field value
func (o *Referral) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *Referral) GetCampaignIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *Referral) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetAdvocateProfileIntegrationId returns the AdvocateProfileIntegrationId field value
func (o *Referral) GetAdvocateProfileIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdvocateProfileIntegrationId
}

// GetAdvocateProfileIntegrationIdOk returns a tuple with the AdvocateProfileIntegrationId field value
// and a boolean to check if the value has been set.
func (o *Referral) GetAdvocateProfileIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdvocateProfileIntegrationId, true
}

// SetAdvocateProfileIntegrationId sets field value
func (o *Referral) SetAdvocateProfileIntegrationId(v string) {
	o.AdvocateProfileIntegrationId = v
}

// GetFriendProfileIntegrationId returns the FriendProfileIntegrationId field value if set, zero value otherwise.
func (o *Referral) GetFriendProfileIntegrationId() string {
	if o == nil || o.FriendProfileIntegrationId == nil {
		var ret string
		return ret
	}
	return *o.FriendProfileIntegrationId
}

// GetFriendProfileIntegrationIdOk returns a tuple with the FriendProfileIntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Referral) GetFriendProfileIntegrationIdOk() (*string, bool) {
	if o == nil || o.FriendProfileIntegrationId == nil {
		return nil, false
	}
	return o.FriendProfileIntegrationId, true
}

// HasFriendProfileIntegrationId returns a boolean if a field has been set.
func (o *Referral) HasFriendProfileIntegrationId() bool {
	if o != nil && o.FriendProfileIntegrationId != nil {
		return true
	}

	return false
}

// SetFriendProfileIntegrationId gets a reference to the given string and assigns it to the FriendProfileIntegrationId field.
func (o *Referral) SetFriendProfileIntegrationId(v string) {
	o.FriendProfileIntegrationId = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Referral) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Referral) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Referral) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *Referral) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetImportId returns the ImportId field value if set, zero value otherwise.
func (o *Referral) GetImportId() int32 {
	if o == nil || o.ImportId == nil {
		var ret int32
		return ret
	}
	return *o.ImportId
}

// GetImportIdOk returns a tuple with the ImportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Referral) GetImportIdOk() (*int32, bool) {
	if o == nil || o.ImportId == nil {
		return nil, false
	}
	return o.ImportId, true
}

// HasImportId returns a boolean if a field has been set.
func (o *Referral) HasImportId() bool {
	if o != nil && o.ImportId != nil {
		return true
	}

	return false
}

// SetImportId gets a reference to the given int32 and assigns it to the ImportId field.
func (o *Referral) SetImportId(v int32) {
	o.ImportId = &v
}

// GetCode returns the Code field value
func (o *Referral) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Referral) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Referral) SetCode(v string) {
	o.Code = v
}

// GetUsageCounter returns the UsageCounter field value
func (o *Referral) GetUsageCounter() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageCounter
}

// GetUsageCounterOk returns a tuple with the UsageCounter field value
// and a boolean to check if the value has been set.
func (o *Referral) GetUsageCounterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageCounter, true
}

// SetUsageCounter sets field value
func (o *Referral) SetUsageCounter(v int32) {
	o.UsageCounter = v
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *Referral) GetBatchId() string {
	if o == nil || o.BatchId == nil {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Referral) GetBatchIdOk() (*string, bool) {
	if o == nil || o.BatchId == nil {
		return nil, false
	}
	return o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *Referral) HasBatchId() bool {
	if o != nil && o.BatchId != nil {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *Referral) SetBatchId(v string) {
	o.BatchId = &v
}

func (o Referral) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.ExpiryDate != nil {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	if true {
		toSerialize["usageLimit"] = o.UsageLimit
	}
	if true {
		toSerialize["campaignId"] = o.CampaignId
	}
	if true {
		toSerialize["advocateProfileIntegrationId"] = o.AdvocateProfileIntegrationId
	}
	if o.FriendProfileIntegrationId != nil {
		toSerialize["friendProfileIntegrationId"] = o.FriendProfileIntegrationId
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.ImportId != nil {
		toSerialize["importId"] = o.ImportId
	}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["usageCounter"] = o.UsageCounter
	}
	if o.BatchId != nil {
		toSerialize["batchId"] = o.BatchId
	}
	return json.Marshal(toSerialize)
}

type NullableReferral struct {
	value *Referral
	isSet bool
}

func (v NullableReferral) Get() *Referral {
	return v.value
}

func (v *NullableReferral) Set(val *Referral) {
	v.value = val
	v.isSet = true
}

func (v NullableReferral) IsSet() bool {
	return v.isSet
}

func (v *NullableReferral) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferral(val *Referral) *NullableReferral {
	return &NullableReferral{value: val, isSet: true}
}

func (v NullableReferral) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferral) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

