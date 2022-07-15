# Ruleset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**UserId** | **int32** | The ID of the account that owns this entity. | 
**Rules** | [**[]Rule**](Rule.md) | Set of rules to apply. | 
**Bindings** | [**[]Binding**](Binding.md) | An array that provides objects with variable names (name) and talang expressions to whose result they are bound (expression) during rule evaluation. The order of the evaluation is decided by the position in the array. | 
**RbVersion** | Pointer to **string** | The version of the rulebuilder used to create this ruleset. | [optional] 
**Activate** | Pointer to **bool** | Indicates whether this created ruleset should be activated for the campaign that owns it. | [optional] 
**CampaignId** | Pointer to **int32** | The ID of the campaign that owns this entity. | [optional] 
**TemplateId** | Pointer to **int32** | The ID of the campaign template that owns this entity. | [optional] 
**ActivatedAt** | Pointer to **time.Time** | Timestamp indicating when this Ruleset was activated. | [optional] 

## Methods

### NewRuleset

`func NewRuleset(id int32, created time.Time, userId int32, rules []Rule, bindings []Binding, ) *Ruleset`

NewRuleset instantiates a new Ruleset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesetWithDefaults

`func NewRulesetWithDefaults() *Ruleset`

NewRulesetWithDefaults instantiates a new Ruleset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Ruleset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ruleset) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ruleset) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *Ruleset) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ruleset) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Ruleset) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUserId

`func (o *Ruleset) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Ruleset) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Ruleset) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetRules

`func (o *Ruleset) GetRules() []Rule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Ruleset) GetRulesOk() (*[]Rule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Ruleset) SetRules(v []Rule)`

SetRules sets Rules field to given value.


### GetBindings

`func (o *Ruleset) GetBindings() []Binding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *Ruleset) GetBindingsOk() (*[]Binding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *Ruleset) SetBindings(v []Binding)`

SetBindings sets Bindings field to given value.


### GetRbVersion

`func (o *Ruleset) GetRbVersion() string`

GetRbVersion returns the RbVersion field if non-nil, zero value otherwise.

### GetRbVersionOk

`func (o *Ruleset) GetRbVersionOk() (*string, bool)`

GetRbVersionOk returns a tuple with the RbVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbVersion

`func (o *Ruleset) SetRbVersion(v string)`

SetRbVersion sets RbVersion field to given value.

### HasRbVersion

`func (o *Ruleset) HasRbVersion() bool`

HasRbVersion returns a boolean if a field has been set.

### GetActivate

`func (o *Ruleset) GetActivate() bool`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *Ruleset) GetActivateOk() (*bool, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *Ruleset) SetActivate(v bool)`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *Ruleset) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### GetCampaignId

`func (o *Ruleset) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *Ruleset) GetCampaignIdOk() (*int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *Ruleset) SetCampaignId(v int32)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *Ruleset) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetTemplateId

`func (o *Ruleset) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *Ruleset) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *Ruleset) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *Ruleset) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetActivatedAt

`func (o *Ruleset) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *Ruleset) GetActivatedAtOk() (*time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *Ruleset) SetActivatedAt(v time.Time)`

SetActivatedAt sets ActivatedAt field to given value.

### HasActivatedAt

`func (o *Ruleset) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


