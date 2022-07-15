# UpdateCampaignTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**ApplicationsIds** | **[]int32** | A list of the IDs of the applications that are subscribed to this campaign template. | 
**CampaignCollections** | Pointer to [**[]CampaignTemplateCollection**](CampaignTemplateCollection.md) | The campaign collections from the blueprint campaign for the template. | [optional] 

## Methods

### NewUpdateCampaignTemplate

`func NewUpdateCampaignTemplate(name string, description string, instructions string, state string, applicationsIds []int32, ) *UpdateCampaignTemplate`

NewUpdateCampaignTemplate instantiates a new UpdateCampaignTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCampaignTemplateWithDefaults

`func NewUpdateCampaignTemplateWithDefaults() *UpdateCampaignTemplate`

NewUpdateCampaignTemplateWithDefaults instantiates a new UpdateCampaignTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCampaignTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCampaignTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCampaignTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateCampaignTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCampaignTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCampaignTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInstructions

`func (o *UpdateCampaignTemplate) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *UpdateCampaignTemplate) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *UpdateCampaignTemplate) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.


### GetCampaignAttributes

`func (o *UpdateCampaignTemplate) GetCampaignAttributes() map[string]interface{}`

GetCampaignAttributes returns the CampaignAttributes field if non-nil, zero value otherwise.

### GetCampaignAttributesOk

`func (o *UpdateCampaignTemplate) GetCampaignAttributesOk() (*map[string]interface{}, bool)`

GetCampaignAttributesOk returns a tuple with the CampaignAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignAttributes

`func (o *UpdateCampaignTemplate) SetCampaignAttributes(v map[string]interface{})`

SetCampaignAttributes sets CampaignAttributes field to given value.

### HasCampaignAttributes

`func (o *UpdateCampaignTemplate) HasCampaignAttributes() bool`

HasCampaignAttributes returns a boolean if a field has been set.

### GetCouponAttributes

`func (o *UpdateCampaignTemplate) GetCouponAttributes() map[string]interface{}`

GetCouponAttributes returns the CouponAttributes field if non-nil, zero value otherwise.

### GetCouponAttributesOk

`func (o *UpdateCampaignTemplate) GetCouponAttributesOk() (*map[string]interface{}, bool)`

GetCouponAttributesOk returns a tuple with the CouponAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponAttributes

`func (o *UpdateCampaignTemplate) SetCouponAttributes(v map[string]interface{})`

SetCouponAttributes sets CouponAttributes field to given value.

### HasCouponAttributes

`func (o *UpdateCampaignTemplate) HasCouponAttributes() bool`

HasCouponAttributes returns a boolean if a field has been set.

### GetState

`func (o *UpdateCampaignTemplate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateCampaignTemplate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateCampaignTemplate) SetState(v string)`

SetState sets State field to given value.


### GetActiveRulesetId

`func (o *UpdateCampaignTemplate) GetActiveRulesetId() int32`

GetActiveRulesetId returns the ActiveRulesetId field if non-nil, zero value otherwise.

### GetActiveRulesetIdOk

`func (o *UpdateCampaignTemplate) GetActiveRulesetIdOk() (*int32, bool)`

GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesetId

`func (o *UpdateCampaignTemplate) SetActiveRulesetId(v int32)`

SetActiveRulesetId sets ActiveRulesetId field to given value.

### HasActiveRulesetId

`func (o *UpdateCampaignTemplate) HasActiveRulesetId() bool`

HasActiveRulesetId returns a boolean if a field has been set.

### GetTags

`func (o *UpdateCampaignTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateCampaignTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateCampaignTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateCampaignTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFeatures

`func (o *UpdateCampaignTemplate) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *UpdateCampaignTemplate) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *UpdateCampaignTemplate) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *UpdateCampaignTemplate) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetCouponSettings

`func (o *UpdateCampaignTemplate) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *UpdateCampaignTemplate) GetCouponSettingsOk() (*CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponSettings

`func (o *UpdateCampaignTemplate) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings sets CouponSettings field to given value.

### HasCouponSettings

`func (o *UpdateCampaignTemplate) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### GetReferralSettings

`func (o *UpdateCampaignTemplate) GetReferralSettings() CodeGeneratorSettings`

GetReferralSettings returns the ReferralSettings field if non-nil, zero value otherwise.

### GetReferralSettingsOk

`func (o *UpdateCampaignTemplate) GetReferralSettingsOk() (*CodeGeneratorSettings, bool)`

GetReferralSettingsOk returns a tuple with the ReferralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralSettings

`func (o *UpdateCampaignTemplate) SetReferralSettings(v CodeGeneratorSettings)`

SetReferralSettings sets ReferralSettings field to given value.

### HasReferralSettings

`func (o *UpdateCampaignTemplate) HasReferralSettings() bool`

HasReferralSettings returns a boolean if a field has been set.

### GetLimits

`func (o *UpdateCampaignTemplate) GetLimits() []TemplateLimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *UpdateCampaignTemplate) GetLimitsOk() (*[]TemplateLimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *UpdateCampaignTemplate) SetLimits(v []TemplateLimitConfig)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *UpdateCampaignTemplate) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetTemplateParams

`func (o *UpdateCampaignTemplate) GetTemplateParams() []CampaignTemplateParams`

GetTemplateParams returns the TemplateParams field if non-nil, zero value otherwise.

### GetTemplateParamsOk

`func (o *UpdateCampaignTemplate) GetTemplateParamsOk() (*[]CampaignTemplateParams, bool)`

GetTemplateParamsOk returns a tuple with the TemplateParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateParams

`func (o *UpdateCampaignTemplate) SetTemplateParams(v []CampaignTemplateParams)`

SetTemplateParams sets TemplateParams field to given value.

### HasTemplateParams

`func (o *UpdateCampaignTemplate) HasTemplateParams() bool`

HasTemplateParams returns a boolean if a field has been set.

### GetApplicationsIds

`func (o *UpdateCampaignTemplate) GetApplicationsIds() []int32`

GetApplicationsIds returns the ApplicationsIds field if non-nil, zero value otherwise.

### GetApplicationsIdsOk

`func (o *UpdateCampaignTemplate) GetApplicationsIdsOk() (*[]int32, bool)`

GetApplicationsIdsOk returns a tuple with the ApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationsIds

`func (o *UpdateCampaignTemplate) SetApplicationsIds(v []int32)`

SetApplicationsIds sets ApplicationsIds field to given value.


### GetCampaignCollections

`func (o *UpdateCampaignTemplate) GetCampaignCollections() []CampaignTemplateCollection`

GetCampaignCollections returns the CampaignCollections field if non-nil, zero value otherwise.

### GetCampaignCollectionsOk

`func (o *UpdateCampaignTemplate) GetCampaignCollectionsOk() (*[]CampaignTemplateCollection, bool)`

GetCampaignCollectionsOk returns a tuple with the CampaignCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignCollections

`func (o *UpdateCampaignTemplate) SetCampaignCollections(v []CampaignTemplateCollection)`

SetCampaignCollections sets CampaignCollections field to given value.

### HasCampaignCollections

`func (o *UpdateCampaignTemplate) HasCampaignCollections() bool`

HasCampaignCollections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


