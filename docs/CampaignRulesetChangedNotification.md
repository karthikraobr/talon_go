# CampaignRulesetChangedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Campaign** | [**Campaign**](Campaign.md) |  | 
**OldRuleset** | Pointer to [**Ruleset**](Ruleset.md) |  | [optional] 
**Ruleset** | [**Ruleset**](Ruleset.md) |  | 

## Methods

### NewCampaignRulesetChangedNotification

`func NewCampaignRulesetChangedNotification(campaign Campaign, ruleset Ruleset, ) *CampaignRulesetChangedNotification`

NewCampaignRulesetChangedNotification instantiates a new CampaignRulesetChangedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignRulesetChangedNotificationWithDefaults

`func NewCampaignRulesetChangedNotificationWithDefaults() *CampaignRulesetChangedNotification`

NewCampaignRulesetChangedNotificationWithDefaults instantiates a new CampaignRulesetChangedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaign

`func (o *CampaignRulesetChangedNotification) GetCampaign() Campaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *CampaignRulesetChangedNotification) GetCampaignOk() (*Campaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *CampaignRulesetChangedNotification) SetCampaign(v Campaign)`

SetCampaign sets Campaign field to given value.


### GetOldRuleset

`func (o *CampaignRulesetChangedNotification) GetOldRuleset() Ruleset`

GetOldRuleset returns the OldRuleset field if non-nil, zero value otherwise.

### GetOldRulesetOk

`func (o *CampaignRulesetChangedNotification) GetOldRulesetOk() (*Ruleset, bool)`

GetOldRulesetOk returns a tuple with the OldRuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldRuleset

`func (o *CampaignRulesetChangedNotification) SetOldRuleset(v Ruleset)`

SetOldRuleset sets OldRuleset field to given value.

### HasOldRuleset

`func (o *CampaignRulesetChangedNotification) HasOldRuleset() bool`

HasOldRuleset returns a boolean if a field has been set.

### GetRuleset

`func (o *CampaignRulesetChangedNotification) GetRuleset() Ruleset`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *CampaignRulesetChangedNotification) GetRulesetOk() (*Ruleset, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *CampaignRulesetChangedNotification) SetRuleset(v Ruleset)`

SetRuleset sets Ruleset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


