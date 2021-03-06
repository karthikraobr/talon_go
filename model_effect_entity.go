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

// EffectEntity Definition of all properties that are present on all effects, independent of their type.
type EffectEntity struct {
	// The ID of the campaign that triggered this effect.
	CampaignId int32 `json:"campaignId"`
	// The ID of the ruleset that was active in the campaign when this effect was triggered.
	RulesetId int32 `json:"rulesetId"`
	// The position of the rule that triggered this effect within the ruleset.
	RuleIndex int32 `json:"ruleIndex"`
	// The name of the rule that triggered this effect.
	RuleName string `json:"ruleName"`
	// The type of effect that was triggered.
	EffectType string `json:"effectType"`
	// The ID of the coupon that was being evaluated when this effect was triggered.
	TriggeredByCoupon *int32 `json:"triggeredByCoupon,omitempty"`
}

// NewEffectEntity instantiates a new EffectEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEffectEntity(campaignId int32, rulesetId int32, ruleIndex int32, ruleName string, effectType string) *EffectEntity {
	this := EffectEntity{}
	this.CampaignId = campaignId
	this.RulesetId = rulesetId
	this.RuleIndex = ruleIndex
	this.RuleName = ruleName
	this.EffectType = effectType
	return &this
}

// NewEffectEntityWithDefaults instantiates a new EffectEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEffectEntityWithDefaults() *EffectEntity {
	this := EffectEntity{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *EffectEntity) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *EffectEntity) GetCampaignIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *EffectEntity) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetRulesetId returns the RulesetId field value
func (o *EffectEntity) GetRulesetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RulesetId
}

// GetRulesetIdOk returns a tuple with the RulesetId field value
// and a boolean to check if the value has been set.
func (o *EffectEntity) GetRulesetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RulesetId, true
}

// SetRulesetId sets field value
func (o *EffectEntity) SetRulesetId(v int32) {
	o.RulesetId = v
}

// GetRuleIndex returns the RuleIndex field value
func (o *EffectEntity) GetRuleIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RuleIndex
}

// GetRuleIndexOk returns a tuple with the RuleIndex field value
// and a boolean to check if the value has been set.
func (o *EffectEntity) GetRuleIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleIndex, true
}

// SetRuleIndex sets field value
func (o *EffectEntity) SetRuleIndex(v int32) {
	o.RuleIndex = v
}

// GetRuleName returns the RuleName field value
func (o *EffectEntity) GetRuleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value
// and a boolean to check if the value has been set.
func (o *EffectEntity) GetRuleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleName, true
}

// SetRuleName sets field value
func (o *EffectEntity) SetRuleName(v string) {
	o.RuleName = v
}

// GetEffectType returns the EffectType field value
func (o *EffectEntity) GetEffectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EffectType
}

// GetEffectTypeOk returns a tuple with the EffectType field value
// and a boolean to check if the value has been set.
func (o *EffectEntity) GetEffectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectType, true
}

// SetEffectType sets field value
func (o *EffectEntity) SetEffectType(v string) {
	o.EffectType = v
}

// GetTriggeredByCoupon returns the TriggeredByCoupon field value if set, zero value otherwise.
func (o *EffectEntity) GetTriggeredByCoupon() int32 {
	if o == nil || o.TriggeredByCoupon == nil {
		var ret int32
		return ret
	}
	return *o.TriggeredByCoupon
}

// GetTriggeredByCouponOk returns a tuple with the TriggeredByCoupon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectEntity) GetTriggeredByCouponOk() (*int32, bool) {
	if o == nil || o.TriggeredByCoupon == nil {
		return nil, false
	}
	return o.TriggeredByCoupon, true
}

// HasTriggeredByCoupon returns a boolean if a field has been set.
func (o *EffectEntity) HasTriggeredByCoupon() bool {
	if o != nil && o.TriggeredByCoupon != nil {
		return true
	}

	return false
}

// SetTriggeredByCoupon gets a reference to the given int32 and assigns it to the TriggeredByCoupon field.
func (o *EffectEntity) SetTriggeredByCoupon(v int32) {
	o.TriggeredByCoupon = &v
}

func (o EffectEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["campaignId"] = o.CampaignId
	}
	if true {
		toSerialize["rulesetId"] = o.RulesetId
	}
	if true {
		toSerialize["ruleIndex"] = o.RuleIndex
	}
	if true {
		toSerialize["ruleName"] = o.RuleName
	}
	if true {
		toSerialize["effectType"] = o.EffectType
	}
	if o.TriggeredByCoupon != nil {
		toSerialize["triggeredByCoupon"] = o.TriggeredByCoupon
	}
	return json.Marshal(toSerialize)
}

type NullableEffectEntity struct {
	value *EffectEntity
	isSet bool
}

func (v NullableEffectEntity) Get() *EffectEntity {
	return v.value
}

func (v *NullableEffectEntity) Set(val *EffectEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableEffectEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableEffectEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEffectEntity(val *EffectEntity) *NullableEffectEntity {
	return &NullableEffectEntity{value: val, isSet: true}
}

func (v NullableEffectEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEffectEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


