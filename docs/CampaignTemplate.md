# CampaignTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**AccountId** | **int32** | The ID of the account that owns this entity. | 
**UserId** | **int32** | The ID of the account that owns this entity. | 
**Name** | **string** | The campaign template name. | 
**Description** | **string** | Customer-facing text that explains the objective of the template. | 
**Instructions** | **string** | Customer-facing text that explains how to use the template. For example, you can use this property to explain the available attributes of this template, and how they can be modified when a user uses this template to create a new campaign. | 
**CampaignAttributes** | Pointer to **map[string]interface{}** | The Campaign Attributes that Campaigns created from this template will have by default. | [optional] 
**CouponAttributes** | Pointer to **map[string]interface{}** | The Campaign Attributes that Coupons created from this template will have by default. | [optional] 
**State** | **string** | Only Campaign Templates in &#39;available&#39; state may be used to create Campaigns. | 
**ActiveRulesetId** | Pointer to **int32** | The ID of the Ruleset this Campaign Template will use. | [optional] 
**Tags** | Pointer to **[]string** | A list of tags for the campaign template. | [optional] 
**Features** | Pointer to **[]string** | A list of features for the campaign template. | [optional] 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**ReferralSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**Limits** | Pointer to [**[]TemplateLimitConfig**](TemplateLimitConfig.md) | The set of limits that will operate for this campaign template. | [optional] 
**TemplateParams** | Pointer to [**[]CampaignTemplateParams**](CampaignTemplateParams.md) | Template parameters are fields which can be used to replace values in a rule. | [optional] 
**ApplicationsIds** | **[]int32** | A list of the IDs of the applications that are subscribed to this campaign template. A list of the IDs of the applications that are subscribed to this campaign template. | 
**CampaignCollections** | Pointer to [**[]CampaignTemplateCollection**](CampaignTemplateCollection.md) | The campaign collections from the blueprint campaign for the template. | [optional] 
**Updated** | Pointer to **time.Time** | Timestamp of the most recent update to the campaign template or any of its elements. | [optional] 
**UpdatedBy** | Pointer to **string** | Name of the user who last updated this campaign template if available. | [optional] 
**ValidApplicationIds** | **[]int32** | The IDs of the applications that are related to this entity. | 

## Methods

### NewCampaignTemplate

`func NewCampaignTemplate(id int32, created time.Time, accountId int32, userId int32, name string, description string, instructions string, state string, applicationsIds []int32, validApplicationIds []int32, ) *CampaignTemplate`

NewCampaignTemplate instantiates a new CampaignTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignTemplateWithDefaults

`func NewCampaignTemplateWithDefaults() *CampaignTemplate`

NewCampaignTemplateWithDefaults instantiates a new CampaignTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignTemplate) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *CampaignTemplate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignTemplate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CampaignTemplate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetAccountId

`func (o *CampaignTemplate) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CampaignTemplate) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CampaignTemplate) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetUserId

`func (o *CampaignTemplate) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CampaignTemplate) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CampaignTemplate) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetName

`func (o *CampaignTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CampaignTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInstructions

`func (o *CampaignTemplate) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CampaignTemplate) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *CampaignTemplate) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.


### GetCampaignAttributes

`func (o *CampaignTemplate) GetCampaignAttributes() map[string]interface{}`

GetCampaignAttributes returns the CampaignAttributes field if non-nil, zero value otherwise.

### GetCampaignAttributesOk

`func (o *CampaignTemplate) GetCampaignAttributesOk() (*map[string]interface{}, bool)`

GetCampaignAttributesOk returns a tuple with the CampaignAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignAttributes

`func (o *CampaignTemplate) SetCampaignAttributes(v map[string]interface{})`

SetCampaignAttributes sets CampaignAttributes field to given value.

### HasCampaignAttributes

`func (o *CampaignTemplate) HasCampaignAttributes() bool`

HasCampaignAttributes returns a boolean if a field has been set.

### GetCouponAttributes

`func (o *CampaignTemplate) GetCouponAttributes() map[string]interface{}`

GetCouponAttributes returns the CouponAttributes field if non-nil, zero value otherwise.

### GetCouponAttributesOk

`func (o *CampaignTemplate) GetCouponAttributesOk() (*map[string]interface{}, bool)`

GetCouponAttributesOk returns a tuple with the CouponAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponAttributes

`func (o *CampaignTemplate) SetCouponAttributes(v map[string]interface{})`

SetCouponAttributes sets CouponAttributes field to given value.

### HasCouponAttributes

`func (o *CampaignTemplate) HasCouponAttributes() bool`

HasCouponAttributes returns a boolean if a field has been set.

### GetState

`func (o *CampaignTemplate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CampaignTemplate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CampaignTemplate) SetState(v string)`

SetState sets State field to given value.


### GetActiveRulesetId

`func (o *CampaignTemplate) GetActiveRulesetId() int32`

GetActiveRulesetId returns the ActiveRulesetId field if non-nil, zero value otherwise.

### GetActiveRulesetIdOk

`func (o *CampaignTemplate) GetActiveRulesetIdOk() (*int32, bool)`

GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesetId

`func (o *CampaignTemplate) SetActiveRulesetId(v int32)`

SetActiveRulesetId sets ActiveRulesetId field to given value.

### HasActiveRulesetId

`func (o *CampaignTemplate) HasActiveRulesetId() bool`

HasActiveRulesetId returns a boolean if a field has been set.

### GetTags

`func (o *CampaignTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CampaignTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CampaignTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CampaignTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFeatures

`func (o *CampaignTemplate) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CampaignTemplate) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CampaignTemplate) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CampaignTemplate) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetCouponSettings

`func (o *CampaignTemplate) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *CampaignTemplate) GetCouponSettingsOk() (*CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponSettings

`func (o *CampaignTemplate) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings sets CouponSettings field to given value.

### HasCouponSettings

`func (o *CampaignTemplate) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### GetReferralSettings

`func (o *CampaignTemplate) GetReferralSettings() CodeGeneratorSettings`

GetReferralSettings returns the ReferralSettings field if non-nil, zero value otherwise.

### GetReferralSettingsOk

`func (o *CampaignTemplate) GetReferralSettingsOk() (*CodeGeneratorSettings, bool)`

GetReferralSettingsOk returns a tuple with the ReferralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralSettings

`func (o *CampaignTemplate) SetReferralSettings(v CodeGeneratorSettings)`

SetReferralSettings sets ReferralSettings field to given value.

### HasReferralSettings

`func (o *CampaignTemplate) HasReferralSettings() bool`

HasReferralSettings returns a boolean if a field has been set.

### GetLimits

`func (o *CampaignTemplate) GetLimits() []TemplateLimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *CampaignTemplate) GetLimitsOk() (*[]TemplateLimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *CampaignTemplate) SetLimits(v []TemplateLimitConfig)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *CampaignTemplate) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetTemplateParams

`func (o *CampaignTemplate) GetTemplateParams() []CampaignTemplateParams`

GetTemplateParams returns the TemplateParams field if non-nil, zero value otherwise.

### GetTemplateParamsOk

`func (o *CampaignTemplate) GetTemplateParamsOk() (*[]CampaignTemplateParams, bool)`

GetTemplateParamsOk returns a tuple with the TemplateParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateParams

`func (o *CampaignTemplate) SetTemplateParams(v []CampaignTemplateParams)`

SetTemplateParams sets TemplateParams field to given value.

### HasTemplateParams

`func (o *CampaignTemplate) HasTemplateParams() bool`

HasTemplateParams returns a boolean if a field has been set.

### GetApplicationsIds

`func (o *CampaignTemplate) GetApplicationsIds() []int32`

GetApplicationsIds returns the ApplicationsIds field if non-nil, zero value otherwise.

### GetApplicationsIdsOk

`func (o *CampaignTemplate) GetApplicationsIdsOk() (*[]int32, bool)`

GetApplicationsIdsOk returns a tuple with the ApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationsIds

`func (o *CampaignTemplate) SetApplicationsIds(v []int32)`

SetApplicationsIds sets ApplicationsIds field to given value.


### GetCampaignCollections

`func (o *CampaignTemplate) GetCampaignCollections() []CampaignTemplateCollection`

GetCampaignCollections returns the CampaignCollections field if non-nil, zero value otherwise.

### GetCampaignCollectionsOk

`func (o *CampaignTemplate) GetCampaignCollectionsOk() (*[]CampaignTemplateCollection, bool)`

GetCampaignCollectionsOk returns a tuple with the CampaignCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignCollections

`func (o *CampaignTemplate) SetCampaignCollections(v []CampaignTemplateCollection)`

SetCampaignCollections sets CampaignCollections field to given value.

### HasCampaignCollections

`func (o *CampaignTemplate) HasCampaignCollections() bool`

HasCampaignCollections returns a boolean if a field has been set.

### GetUpdated

`func (o *CampaignTemplate) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CampaignTemplate) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CampaignTemplate) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CampaignTemplate) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *CampaignTemplate) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CampaignTemplate) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *CampaignTemplate) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *CampaignTemplate) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetValidApplicationIds

`func (o *CampaignTemplate) GetValidApplicationIds() []int32`

GetValidApplicationIds returns the ValidApplicationIds field if non-nil, zero value otherwise.

### GetValidApplicationIdsOk

`func (o *CampaignTemplate) GetValidApplicationIdsOk() (*[]int32, bool)`

GetValidApplicationIdsOk returns a tuple with the ValidApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidApplicationIds

`func (o *CampaignTemplate) SetValidApplicationIds(v []int32)`

SetValidApplicationIds sets ValidApplicationIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


