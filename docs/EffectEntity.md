# EffectEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **int32** | The ID of the campaign that triggered this effect. | 
**RulesetId** | **int32** | The ID of the ruleset that was active in the campaign when this effect was triggered. | 
**RuleIndex** | **int32** | The position of the rule that triggered this effect within the ruleset. | 
**RuleName** | **string** | The name of the rule that triggered this effect. | 
**EffectType** | **string** | The type of effect that was triggered. | 
**TriggeredByCoupon** | Pointer to **int32** | The ID of the coupon that was being evaluated when this effect was triggered. | [optional] 

## Methods

### NewEffectEntity

`func NewEffectEntity(campaignId int32, rulesetId int32, ruleIndex int32, ruleName string, effectType string, ) *EffectEntity`

NewEffectEntity instantiates a new EffectEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectEntityWithDefaults

`func NewEffectEntityWithDefaults() *EffectEntity`

NewEffectEntityWithDefaults instantiates a new EffectEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *EffectEntity) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *EffectEntity) GetCampaignIdOk() (*int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *EffectEntity) SetCampaignId(v int32)`

SetCampaignId sets CampaignId field to given value.


### GetRulesetId

`func (o *EffectEntity) GetRulesetId() int32`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *EffectEntity) GetRulesetIdOk() (*int32, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetId

`func (o *EffectEntity) SetRulesetId(v int32)`

SetRulesetId sets RulesetId field to given value.


### GetRuleIndex

`func (o *EffectEntity) GetRuleIndex() int32`

GetRuleIndex returns the RuleIndex field if non-nil, zero value otherwise.

### GetRuleIndexOk

`func (o *EffectEntity) GetRuleIndexOk() (*int32, bool)`

GetRuleIndexOk returns a tuple with the RuleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIndex

`func (o *EffectEntity) SetRuleIndex(v int32)`

SetRuleIndex sets RuleIndex field to given value.


### GetRuleName

`func (o *EffectEntity) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *EffectEntity) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *EffectEntity) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.


### GetEffectType

`func (o *EffectEntity) GetEffectType() string`

GetEffectType returns the EffectType field if non-nil, zero value otherwise.

### GetEffectTypeOk

`func (o *EffectEntity) GetEffectTypeOk() (*string, bool)`

GetEffectTypeOk returns a tuple with the EffectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectType

`func (o *EffectEntity) SetEffectType(v string)`

SetEffectType sets EffectType field to given value.


### GetTriggeredByCoupon

`func (o *EffectEntity) GetTriggeredByCoupon() int32`

GetTriggeredByCoupon returns the TriggeredByCoupon field if non-nil, zero value otherwise.

### GetTriggeredByCouponOk

`func (o *EffectEntity) GetTriggeredByCouponOk() (*int32, bool)`

GetTriggeredByCouponOk returns a tuple with the TriggeredByCoupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredByCoupon

`func (o *EffectEntity) SetTriggeredByCoupon(v int32)`

SetTriggeredByCoupon sets TriggeredByCoupon field to given value.

### HasTriggeredByCoupon

`func (o *EffectEntity) HasTriggeredByCoupon() bool`

HasTriggeredByCoupon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


