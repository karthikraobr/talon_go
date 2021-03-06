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

// NewCampaign struct for NewCampaign
type NewCampaign struct {
	// A user-facing name for this campaign.
	Name string `json:"name"`
	// A detailed description of the campaign.
	Description *string `json:"description,omitempty"`
	// Timestamp when the campaign will become active.
	StartTime *time.Time `json:"startTime,omitempty"`
	// Timestamp the campaign will become inactive.
	EndTime *time.Time `json:"endTime,omitempty"`
	// Arbitrary properties associated with this campaign.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// A disabled or archived campaign is not evaluated for rules or coupons. 
	State string `json:"state"`
	// [ID of Ruleset](https://docs.talon.one/management-api/#operation/getRulesets) this campaign applies on customer session evaluation. 
	ActiveRulesetId *int32 `json:"activeRulesetId,omitempty"`
	// A list of tags for the campaign.
	Tags []string `json:"tags"`
	// The features enabled in this campaign.
	Features []string `json:"features"`
	CouponSettings *CodeGeneratorSettings `json:"couponSettings,omitempty"`
	ReferralSettings *CodeGeneratorSettings `json:"referralSettings,omitempty"`
	// The set of [budget limits](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets/) for this campaign. 
	Limits []LimitConfig `json:"limits"`
	// The IDs of the [campaign groups](https://docs.talon.one/docs/product/account/managing-campaign-groups/) this campaign belongs to. 
	CampaignGroups []int32 `json:"campaignGroups,omitempty"`
}

// NewNewCampaign instantiates a new NewCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewCampaign(name string, state string, tags []string, features []string, limits []LimitConfig) *NewCampaign {
	this := NewCampaign{}
	this.Name = name
	this.State = state
	this.Tags = tags
	this.Features = features
	this.Limits = limits
	return &this
}

// NewNewCampaignWithDefaults instantiates a new NewCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewCampaignWithDefaults() *NewCampaign {
	this := NewCampaign{}
	var state string = "enabled"
	this.State = state
	return &this
}

// GetName returns the Name field value
func (o *NewCampaign) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NewCampaign) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NewCampaign) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NewCampaign) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NewCampaign) SetDescription(v string) {
	o.Description = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *NewCampaign) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *NewCampaign) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *NewCampaign) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *NewCampaign) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *NewCampaign) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *NewCampaign) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NewCampaign) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NewCampaign) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *NewCampaign) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetState returns the State field value
func (o *NewCampaign) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *NewCampaign) SetState(v string) {
	o.State = v
}

// GetActiveRulesetId returns the ActiveRulesetId field value if set, zero value otherwise.
func (o *NewCampaign) GetActiveRulesetId() int32 {
	if o == nil || o.ActiveRulesetId == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesetId
}

// GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetActiveRulesetIdOk() (*int32, bool) {
	if o == nil || o.ActiveRulesetId == nil {
		return nil, false
	}
	return o.ActiveRulesetId, true
}

// HasActiveRulesetId returns a boolean if a field has been set.
func (o *NewCampaign) HasActiveRulesetId() bool {
	if o != nil && o.ActiveRulesetId != nil {
		return true
	}

	return false
}

// SetActiveRulesetId gets a reference to the given int32 and assigns it to the ActiveRulesetId field.
func (o *NewCampaign) SetActiveRulesetId(v int32) {
	o.ActiveRulesetId = &v
}

// GetTags returns the Tags field value
func (o *NewCampaign) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *NewCampaign) SetTags(v []string) {
	o.Tags = v
}

// GetFeatures returns the Features field value
func (o *NewCampaign) GetFeatures() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetFeaturesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *NewCampaign) SetFeatures(v []string) {
	o.Features = v
}

// GetCouponSettings returns the CouponSettings field value if set, zero value otherwise.
func (o *NewCampaign) GetCouponSettings() CodeGeneratorSettings {
	if o == nil || o.CouponSettings == nil {
		var ret CodeGeneratorSettings
		return ret
	}
	return *o.CouponSettings
}

// GetCouponSettingsOk returns a tuple with the CouponSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetCouponSettingsOk() (*CodeGeneratorSettings, bool) {
	if o == nil || o.CouponSettings == nil {
		return nil, false
	}
	return o.CouponSettings, true
}

// HasCouponSettings returns a boolean if a field has been set.
func (o *NewCampaign) HasCouponSettings() bool {
	if o != nil && o.CouponSettings != nil {
		return true
	}

	return false
}

// SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.
func (o *NewCampaign) SetCouponSettings(v CodeGeneratorSettings) {
	o.CouponSettings = &v
}

// GetReferralSettings returns the ReferralSettings field value if set, zero value otherwise.
func (o *NewCampaign) GetReferralSettings() CodeGeneratorSettings {
	if o == nil || o.ReferralSettings == nil {
		var ret CodeGeneratorSettings
		return ret
	}
	return *o.ReferralSettings
}

// GetReferralSettingsOk returns a tuple with the ReferralSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetReferralSettingsOk() (*CodeGeneratorSettings, bool) {
	if o == nil || o.ReferralSettings == nil {
		return nil, false
	}
	return o.ReferralSettings, true
}

// HasReferralSettings returns a boolean if a field has been set.
func (o *NewCampaign) HasReferralSettings() bool {
	if o != nil && o.ReferralSettings != nil {
		return true
	}

	return false
}

// SetReferralSettings gets a reference to the given CodeGeneratorSettings and assigns it to the ReferralSettings field.
func (o *NewCampaign) SetReferralSettings(v CodeGeneratorSettings) {
	o.ReferralSettings = &v
}

// GetLimits returns the Limits field value
func (o *NewCampaign) GetLimits() []LimitConfig {
	if o == nil {
		var ret []LimitConfig
		return ret
	}

	return o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetLimitsOk() ([]LimitConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limits, true
}

// SetLimits sets field value
func (o *NewCampaign) SetLimits(v []LimitConfig) {
	o.Limits = v
}

// GetCampaignGroups returns the CampaignGroups field value if set, zero value otherwise.
func (o *NewCampaign) GetCampaignGroups() []int32 {
	if o == nil || o.CampaignGroups == nil {
		var ret []int32
		return ret
	}
	return o.CampaignGroups
}

// GetCampaignGroupsOk returns a tuple with the CampaignGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetCampaignGroupsOk() ([]int32, bool) {
	if o == nil || o.CampaignGroups == nil {
		return nil, false
	}
	return o.CampaignGroups, true
}

// HasCampaignGroups returns a boolean if a field has been set.
func (o *NewCampaign) HasCampaignGroups() bool {
	if o != nil && o.CampaignGroups != nil {
		return true
	}

	return false
}

// SetCampaignGroups gets a reference to the given []int32 and assigns it to the CampaignGroups field.
func (o *NewCampaign) SetCampaignGroups(v []int32) {
	o.CampaignGroups = v
}

func (o NewCampaign) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}
	if o.EndTime != nil {
		toSerialize["endTime"] = o.EndTime
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.ActiveRulesetId != nil {
		toSerialize["activeRulesetId"] = o.ActiveRulesetId
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["features"] = o.Features
	}
	if o.CouponSettings != nil {
		toSerialize["couponSettings"] = o.CouponSettings
	}
	if o.ReferralSettings != nil {
		toSerialize["referralSettings"] = o.ReferralSettings
	}
	if true {
		toSerialize["limits"] = o.Limits
	}
	if o.CampaignGroups != nil {
		toSerialize["campaignGroups"] = o.CampaignGroups
	}
	return json.Marshal(toSerialize)
}

type NullableNewCampaign struct {
	value *NewCampaign
	isSet bool
}

func (v NullableNewCampaign) Get() *NewCampaign {
	return v.value
}

func (v *NullableNewCampaign) Set(val *NewCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableNewCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableNewCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewCampaign(val *NewCampaign) *NullableNewCampaign {
	return &NullableNewCampaign{value: val, isSet: true}
}

func (v NullableNewCampaign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


