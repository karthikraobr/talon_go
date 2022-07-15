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

// UpdateCampaignTemplate struct for UpdateCampaignTemplate
type UpdateCampaignTemplate struct {
	// The campaign template name.
	Name string `json:"name"`
	// Customer-facing text that explains the objective of the template.
	Description string `json:"description"`
	// Customer-facing text that explains how to use the template. For example, you can use this property to explain the available attributes of this template, and how they can be modified when a user uses this template to create a new campaign.
	Instructions string `json:"instructions"`
	// The Campaign Attributes that Campaigns created from this template will have by default.
	CampaignAttributes map[string]interface{} `json:"campaignAttributes,omitempty"`
	// The Campaign Attributes that Coupons created from this template will have by default.
	CouponAttributes map[string]interface{} `json:"couponAttributes,omitempty"`
	// Only Campaign Templates in 'available' state may be used to create Campaigns.
	State string `json:"state"`
	// The ID of the Ruleset this Campaign Template will use.
	ActiveRulesetId *int32 `json:"activeRulesetId,omitempty"`
	// A list of tags for the campaign template.
	Tags []string `json:"tags,omitempty"`
	// A list of features for the campaign template.
	Features []string `json:"features,omitempty"`
	CouponSettings *CodeGeneratorSettings `json:"couponSettings,omitempty"`
	ReferralSettings *CodeGeneratorSettings `json:"referralSettings,omitempty"`
	// The set of limits that will operate for this campaign template.
	Limits []TemplateLimitConfig `json:"limits,omitempty"`
	// Template parameters are fields which can be used to replace values in a rule.
	TemplateParams []CampaignTemplateParams `json:"templateParams,omitempty"`
	// A list of the IDs of the applications that are subscribed to this campaign template.
	ApplicationsIds []int32 `json:"applicationsIds"`
	// The campaign collections from the blueprint campaign for the template.
	CampaignCollections []CampaignTemplateCollection `json:"campaignCollections,omitempty"`
}

// NewUpdateCampaignTemplate instantiates a new UpdateCampaignTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCampaignTemplate(name string, description string, instructions string, state string, applicationsIds []int32) *UpdateCampaignTemplate {
	this := UpdateCampaignTemplate{}
	this.Name = name
	this.Description = description
	this.Instructions = instructions
	this.State = state
	this.ApplicationsIds = applicationsIds
	return &this
}

// NewUpdateCampaignTemplateWithDefaults instantiates a new UpdateCampaignTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCampaignTemplateWithDefaults() *UpdateCampaignTemplate {
	this := UpdateCampaignTemplate{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateCampaignTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateCampaignTemplate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *UpdateCampaignTemplate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *UpdateCampaignTemplate) SetDescription(v string) {
	o.Description = v
}

// GetInstructions returns the Instructions field value
func (o *UpdateCampaignTemplate) GetInstructions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetInstructionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instructions, true
}

// SetInstructions sets field value
func (o *UpdateCampaignTemplate) SetInstructions(v string) {
	o.Instructions = v
}

// GetCampaignAttributes returns the CampaignAttributes field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetCampaignAttributes() map[string]interface{} {
	if o == nil || o.CampaignAttributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CampaignAttributes
}

// GetCampaignAttributesOk returns a tuple with the CampaignAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetCampaignAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.CampaignAttributes == nil {
		return nil, false
	}
	return o.CampaignAttributes, true
}

// HasCampaignAttributes returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasCampaignAttributes() bool {
	if o != nil && o.CampaignAttributes != nil {
		return true
	}

	return false
}

// SetCampaignAttributes gets a reference to the given map[string]interface{} and assigns it to the CampaignAttributes field.
func (o *UpdateCampaignTemplate) SetCampaignAttributes(v map[string]interface{}) {
	o.CampaignAttributes = v
}

// GetCouponAttributes returns the CouponAttributes field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetCouponAttributes() map[string]interface{} {
	if o == nil || o.CouponAttributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CouponAttributes
}

// GetCouponAttributesOk returns a tuple with the CouponAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetCouponAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.CouponAttributes == nil {
		return nil, false
	}
	return o.CouponAttributes, true
}

// HasCouponAttributes returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasCouponAttributes() bool {
	if o != nil && o.CouponAttributes != nil {
		return true
	}

	return false
}

// SetCouponAttributes gets a reference to the given map[string]interface{} and assigns it to the CouponAttributes field.
func (o *UpdateCampaignTemplate) SetCouponAttributes(v map[string]interface{}) {
	o.CouponAttributes = v
}

// GetState returns the State field value
func (o *UpdateCampaignTemplate) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *UpdateCampaignTemplate) SetState(v string) {
	o.State = v
}

// GetActiveRulesetId returns the ActiveRulesetId field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetActiveRulesetId() int32 {
	if o == nil || o.ActiveRulesetId == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesetId
}

// GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetActiveRulesetIdOk() (*int32, bool) {
	if o == nil || o.ActiveRulesetId == nil {
		return nil, false
	}
	return o.ActiveRulesetId, true
}

// HasActiveRulesetId returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasActiveRulesetId() bool {
	if o != nil && o.ActiveRulesetId != nil {
		return true
	}

	return false
}

// SetActiveRulesetId gets a reference to the given int32 and assigns it to the ActiveRulesetId field.
func (o *UpdateCampaignTemplate) SetActiveRulesetId(v int32) {
	o.ActiveRulesetId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateCampaignTemplate) SetTags(v []string) {
	o.Tags = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetFeatures() []string {
	if o == nil || o.Features == nil {
		var ret []string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetFeaturesOk() ([]string, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *UpdateCampaignTemplate) SetFeatures(v []string) {
	o.Features = v
}

// GetCouponSettings returns the CouponSettings field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetCouponSettings() CodeGeneratorSettings {
	if o == nil || o.CouponSettings == nil {
		var ret CodeGeneratorSettings
		return ret
	}
	return *o.CouponSettings
}

// GetCouponSettingsOk returns a tuple with the CouponSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetCouponSettingsOk() (*CodeGeneratorSettings, bool) {
	if o == nil || o.CouponSettings == nil {
		return nil, false
	}
	return o.CouponSettings, true
}

// HasCouponSettings returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasCouponSettings() bool {
	if o != nil && o.CouponSettings != nil {
		return true
	}

	return false
}

// SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.
func (o *UpdateCampaignTemplate) SetCouponSettings(v CodeGeneratorSettings) {
	o.CouponSettings = &v
}

// GetReferralSettings returns the ReferralSettings field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetReferralSettings() CodeGeneratorSettings {
	if o == nil || o.ReferralSettings == nil {
		var ret CodeGeneratorSettings
		return ret
	}
	return *o.ReferralSettings
}

// GetReferralSettingsOk returns a tuple with the ReferralSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetReferralSettingsOk() (*CodeGeneratorSettings, bool) {
	if o == nil || o.ReferralSettings == nil {
		return nil, false
	}
	return o.ReferralSettings, true
}

// HasReferralSettings returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasReferralSettings() bool {
	if o != nil && o.ReferralSettings != nil {
		return true
	}

	return false
}

// SetReferralSettings gets a reference to the given CodeGeneratorSettings and assigns it to the ReferralSettings field.
func (o *UpdateCampaignTemplate) SetReferralSettings(v CodeGeneratorSettings) {
	o.ReferralSettings = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetLimits() []TemplateLimitConfig {
	if o == nil || o.Limits == nil {
		var ret []TemplateLimitConfig
		return ret
	}
	return o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetLimitsOk() ([]TemplateLimitConfig, bool) {
	if o == nil || o.Limits == nil {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasLimits() bool {
	if o != nil && o.Limits != nil {
		return true
	}

	return false
}

// SetLimits gets a reference to the given []TemplateLimitConfig and assigns it to the Limits field.
func (o *UpdateCampaignTemplate) SetLimits(v []TemplateLimitConfig) {
	o.Limits = v
}

// GetTemplateParams returns the TemplateParams field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetTemplateParams() []CampaignTemplateParams {
	if o == nil || o.TemplateParams == nil {
		var ret []CampaignTemplateParams
		return ret
	}
	return o.TemplateParams
}

// GetTemplateParamsOk returns a tuple with the TemplateParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetTemplateParamsOk() ([]CampaignTemplateParams, bool) {
	if o == nil || o.TemplateParams == nil {
		return nil, false
	}
	return o.TemplateParams, true
}

// HasTemplateParams returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasTemplateParams() bool {
	if o != nil && o.TemplateParams != nil {
		return true
	}

	return false
}

// SetTemplateParams gets a reference to the given []CampaignTemplateParams and assigns it to the TemplateParams field.
func (o *UpdateCampaignTemplate) SetTemplateParams(v []CampaignTemplateParams) {
	o.TemplateParams = v
}

// GetApplicationsIds returns the ApplicationsIds field value
func (o *UpdateCampaignTemplate) GetApplicationsIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ApplicationsIds
}

// GetApplicationsIdsOk returns a tuple with the ApplicationsIds field value
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetApplicationsIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationsIds, true
}

// SetApplicationsIds sets field value
func (o *UpdateCampaignTemplate) SetApplicationsIds(v []int32) {
	o.ApplicationsIds = v
}

// GetCampaignCollections returns the CampaignCollections field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetCampaignCollections() []CampaignTemplateCollection {
	if o == nil || o.CampaignCollections == nil {
		var ret []CampaignTemplateCollection
		return ret
	}
	return o.CampaignCollections
}

// GetCampaignCollectionsOk returns a tuple with the CampaignCollections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetCampaignCollectionsOk() ([]CampaignTemplateCollection, bool) {
	if o == nil || o.CampaignCollections == nil {
		return nil, false
	}
	return o.CampaignCollections, true
}

// HasCampaignCollections returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasCampaignCollections() bool {
	if o != nil && o.CampaignCollections != nil {
		return true
	}

	return false
}

// SetCampaignCollections gets a reference to the given []CampaignTemplateCollection and assigns it to the CampaignCollections field.
func (o *UpdateCampaignTemplate) SetCampaignCollections(v []CampaignTemplateCollection) {
	o.CampaignCollections = v
}

func (o UpdateCampaignTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["instructions"] = o.Instructions
	}
	if o.CampaignAttributes != nil {
		toSerialize["campaignAttributes"] = o.CampaignAttributes
	}
	if o.CouponAttributes != nil {
		toSerialize["couponAttributes"] = o.CouponAttributes
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.ActiveRulesetId != nil {
		toSerialize["activeRulesetId"] = o.ActiveRulesetId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.CouponSettings != nil {
		toSerialize["couponSettings"] = o.CouponSettings
	}
	if o.ReferralSettings != nil {
		toSerialize["referralSettings"] = o.ReferralSettings
	}
	if o.Limits != nil {
		toSerialize["limits"] = o.Limits
	}
	if o.TemplateParams != nil {
		toSerialize["templateParams"] = o.TemplateParams
	}
	if true {
		toSerialize["applicationsIds"] = o.ApplicationsIds
	}
	if o.CampaignCollections != nil {
		toSerialize["campaignCollections"] = o.CampaignCollections
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateCampaignTemplate struct {
	value *UpdateCampaignTemplate
	isSet bool
}

func (v NullableUpdateCampaignTemplate) Get() *UpdateCampaignTemplate {
	return v.value
}

func (v *NullableUpdateCampaignTemplate) Set(val *UpdateCampaignTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCampaignTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCampaignTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCampaignTemplate(val *UpdateCampaignTemplate) *NullableUpdateCampaignTemplate {
	return &NullableUpdateCampaignTemplate{value: val, isSet: true}
}

func (v NullableUpdateCampaignTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCampaignTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


