# Effect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **int32** | The ID of the campaign that triggered this effect. | 
**RulesetId** | **int32** | The ID of the ruleset that was active in the campaign when this effect was triggered. | 
**RuleIndex** | **int32** | The position of the rule that triggered this effect within the ruleset. | 
**RuleName** | **string** | The name of the rule that triggered this effect. | 
**EffectType** | **string** | The type of effect that was triggered. | 
**TriggeredByCoupon** | Pointer to **int32** | The ID of the coupon that was being evaluated when this effect was triggered. | [optional] 
**Props** | **map[string]interface{}** |  | 

## Methods

### NewEffect

`func NewEffect(campaignId int32, rulesetId int32, ruleIndex int32, ruleName string, effectType string, props map[string]interface{}, ) *Effect`

NewEffect instantiates a new Effect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectWithDefaults

`func NewEffectWithDefaults() *Effect`

NewEffectWithDefaults instantiates a new Effect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *Effect) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *Effect) GetCampaignIdOk() (*int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *Effect) SetCampaignId(v int32)`

SetCampaignId sets CampaignId field to given value.


### GetRulesetId

`func (o *Effect) GetRulesetId() int32`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *Effect) GetRulesetIdOk() (*int32, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetId

`func (o *Effect) SetRulesetId(v int32)`

SetRulesetId sets RulesetId field to given value.


### GetRuleIndex

`func (o *Effect) GetRuleIndex() int32`

GetRuleIndex returns the RuleIndex field if non-nil, zero value otherwise.

### GetRuleIndexOk

`func (o *Effect) GetRuleIndexOk() (*int32, bool)`

GetRuleIndexOk returns a tuple with the RuleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIndex

`func (o *Effect) SetRuleIndex(v int32)`

SetRuleIndex sets RuleIndex field to given value.


### GetRuleName

`func (o *Effect) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *Effect) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *Effect) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.


### GetEffectType

`func (o *Effect) GetEffectType() string`

GetEffectType returns the EffectType field if non-nil, zero value otherwise.

### GetEffectTypeOk

`func (o *Effect) GetEffectTypeOk() (*string, bool)`

GetEffectTypeOk returns a tuple with the EffectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectType

`func (o *Effect) SetEffectType(v string)`

SetEffectType sets EffectType field to given value.


### GetTriggeredByCoupon

`func (o *Effect) GetTriggeredByCoupon() int32`

GetTriggeredByCoupon returns the TriggeredByCoupon field if non-nil, zero value otherwise.

### GetTriggeredByCouponOk

`func (o *Effect) GetTriggeredByCouponOk() (*int32, bool)`

GetTriggeredByCouponOk returns a tuple with the TriggeredByCoupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredByCoupon

`func (o *Effect) SetTriggeredByCoupon(v int32)`

SetTriggeredByCoupon sets TriggeredByCoupon field to given value.

### HasTriggeredByCoupon

`func (o *Effect) HasTriggeredByCoupon() bool`

HasTriggeredByCoupon returns a boolean if a field has been set.

### GetProps

`func (o *Effect) GetProps() map[string]interface{}`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *Effect) GetPropsOk() (*map[string]interface{}, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *Effect) SetProps(v map[string]interface{})`

SetProps sets Props field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


