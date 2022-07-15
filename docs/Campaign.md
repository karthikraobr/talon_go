# Campaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**ApplicationId** | **int32** | The ID of the application that owns this entity. | 
**UserId** | **int32** | The ID of the account that owns this entity. | 
**Name** | **string** | A user-facing name for this campaign. | 
**Description** | **string** | A detailed description of the campaign. | 
**StartTime** | Pointer to **time.Time** | Timestamp when the campaign will become active. | [optional] 
**EndTime** | Pointer to **time.Time** | Timestamp the campaign will become inactive. | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | Arbitrary properties associated with this campaign. | [optional] 
**State** | **string** | A disabled or archived campaign is not evaluated for rules or coupons.  | [default to "enabled"]
**ActiveRulesetId** | Pointer to **int32** | [ID of Ruleset](https://docs.talon.one/management-api/#operation/getRulesets) this campaign applies on customer session evaluation.  | [optional] 
**Tags** | **[]string** | A list of tags for the campaign. | 
**Features** | **[]string** | The features enabled in this campaign. | 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**ReferralSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**Limits** | [**[]LimitConfig**](LimitConfig.md) | The set of [budget limits](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets/) for this campaign.  | 
**CampaignGroups** | Pointer to **[]int32** | The IDs of the [campaign groups](https://docs.talon.one/docs/product/account/managing-campaign-groups/) this campaign belongs to.  | [optional] 
**CouponRedemptionCount** | Pointer to **int32** | Number of coupons redeemed in the campaign. | [optional] 
**ReferralRedemptionCount** | Pointer to **int32** | Number of referral codes redeemed in the campaign. | [optional] 
**DiscountCount** | Pointer to **float32** | Total amount of discounts redeemed in the campaign. | [optional] 
**DiscountEffectCount** | Pointer to **int32** | Total number of times discounts were redeemed in this campaign. | [optional] 
**CouponCreationCount** | Pointer to **int32** | Total number of coupons created by rules in this campaign. | [optional] 
**CustomEffectCount** | Pointer to **int32** | Total number of custom effects triggered by rules in this campaign. | [optional] 
**ReferralCreationCount** | Pointer to **int32** | Total number of referrals created by rules in this campaign. | [optional] 
**AddFreeItemEffectCount** | Pointer to **int32** | Total number of times triggering add free item effext is allowed in this campaign. | [optional] 
**AwardedGiveawaysCount** | Pointer to **int32** | Total number of giveaways awarded by rules in this campaign. | [optional] 
**CreatedLoyaltyPointsCount** | Pointer to **float32** | Total number of loyalty points created by rules in this campaign. | [optional] 
**CreatedLoyaltyPointsEffectCount** | Pointer to **int32** | Total number of loyalty point creation effects triggered by rules in this campaign. | [optional] 
**RedeemedLoyaltyPointsCount** | Pointer to **float32** | Total number of loyalty points redeemed by rules in this campaign. | [optional] 
**RedeemedLoyaltyPointsEffectCount** | Pointer to **int32** | Total number of loyalty point redemption effects triggered by rules in this campaign. | [optional] 
**CallApiEffectCount** | Pointer to **int32** | Total number of webhook triggered by rules in this campaign. | [optional] 
**LastActivity** | Pointer to **time.Time** | Timestamp of the most recent event received by this campaign. | [optional] 
**Updated** | Pointer to **time.Time** | Timestamp of the most recent update to the campaign&#39;s property. Updates to external entities used in this campaign are **not** registered by this property, such as collection or coupon updates.  | [optional] 
**CreatedBy** | Pointer to **string** | Name of the user who created this campaign if available. | [optional] 
**UpdatedBy** | Pointer to **string** | Name of the user who last updated this campaign if available. | [optional] 
**TemplateId** | Pointer to **int32** | The ID of the Campaign Template this Campaign was created from. | [optional] 

## Methods

### NewCampaign

`func NewCampaign(id int32, created time.Time, applicationId int32, userId int32, name string, description string, state string, tags []string, features []string, limits []LimitConfig, ) *Campaign`

NewCampaign instantiates a new Campaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignWithDefaults

`func NewCampaignWithDefaults() *Campaign`

NewCampaignWithDefaults instantiates a new Campaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Campaign) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Campaign) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Campaign) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *Campaign) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Campaign) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Campaign) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetApplicationId

`func (o *Campaign) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Campaign) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *Campaign) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.


### GetUserId

`func (o *Campaign) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Campaign) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Campaign) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetName

`func (o *Campaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Campaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Campaign) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Campaign) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Campaign) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Campaign) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStartTime

`func (o *Campaign) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Campaign) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Campaign) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Campaign) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *Campaign) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Campaign) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Campaign) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Campaign) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetAttributes

`func (o *Campaign) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Campaign) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Campaign) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Campaign) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetState

`func (o *Campaign) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Campaign) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Campaign) SetState(v string)`

SetState sets State field to given value.


### GetActiveRulesetId

`func (o *Campaign) GetActiveRulesetId() int32`

GetActiveRulesetId returns the ActiveRulesetId field if non-nil, zero value otherwise.

### GetActiveRulesetIdOk

`func (o *Campaign) GetActiveRulesetIdOk() (*int32, bool)`

GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesetId

`func (o *Campaign) SetActiveRulesetId(v int32)`

SetActiveRulesetId sets ActiveRulesetId field to given value.

### HasActiveRulesetId

`func (o *Campaign) HasActiveRulesetId() bool`

HasActiveRulesetId returns a boolean if a field has been set.

### GetTags

`func (o *Campaign) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Campaign) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Campaign) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetFeatures

`func (o *Campaign) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Campaign) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Campaign) SetFeatures(v []string)`

SetFeatures sets Features field to given value.


### GetCouponSettings

`func (o *Campaign) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *Campaign) GetCouponSettingsOk() (*CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponSettings

`func (o *Campaign) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings sets CouponSettings field to given value.

### HasCouponSettings

`func (o *Campaign) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### GetReferralSettings

`func (o *Campaign) GetReferralSettings() CodeGeneratorSettings`

GetReferralSettings returns the ReferralSettings field if non-nil, zero value otherwise.

### GetReferralSettingsOk

`func (o *Campaign) GetReferralSettingsOk() (*CodeGeneratorSettings, bool)`

GetReferralSettingsOk returns a tuple with the ReferralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralSettings

`func (o *Campaign) SetReferralSettings(v CodeGeneratorSettings)`

SetReferralSettings sets ReferralSettings field to given value.

### HasReferralSettings

`func (o *Campaign) HasReferralSettings() bool`

HasReferralSettings returns a boolean if a field has been set.

### GetLimits

`func (o *Campaign) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *Campaign) GetLimitsOk() (*[]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *Campaign) SetLimits(v []LimitConfig)`

SetLimits sets Limits field to given value.


### GetCampaignGroups

`func (o *Campaign) GetCampaignGroups() []int32`

GetCampaignGroups returns the CampaignGroups field if non-nil, zero value otherwise.

### GetCampaignGroupsOk

`func (o *Campaign) GetCampaignGroupsOk() (*[]int32, bool)`

GetCampaignGroupsOk returns a tuple with the CampaignGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignGroups

`func (o *Campaign) SetCampaignGroups(v []int32)`

SetCampaignGroups sets CampaignGroups field to given value.

### HasCampaignGroups

`func (o *Campaign) HasCampaignGroups() bool`

HasCampaignGroups returns a boolean if a field has been set.

### GetCouponRedemptionCount

`func (o *Campaign) GetCouponRedemptionCount() int32`

GetCouponRedemptionCount returns the CouponRedemptionCount field if non-nil, zero value otherwise.

### GetCouponRedemptionCountOk

`func (o *Campaign) GetCouponRedemptionCountOk() (*int32, bool)`

GetCouponRedemptionCountOk returns a tuple with the CouponRedemptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponRedemptionCount

`func (o *Campaign) SetCouponRedemptionCount(v int32)`

SetCouponRedemptionCount sets CouponRedemptionCount field to given value.

### HasCouponRedemptionCount

`func (o *Campaign) HasCouponRedemptionCount() bool`

HasCouponRedemptionCount returns a boolean if a field has been set.

### GetReferralRedemptionCount

`func (o *Campaign) GetReferralRedemptionCount() int32`

GetReferralRedemptionCount returns the ReferralRedemptionCount field if non-nil, zero value otherwise.

### GetReferralRedemptionCountOk

`func (o *Campaign) GetReferralRedemptionCountOk() (*int32, bool)`

GetReferralRedemptionCountOk returns a tuple with the ReferralRedemptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralRedemptionCount

`func (o *Campaign) SetReferralRedemptionCount(v int32)`

SetReferralRedemptionCount sets ReferralRedemptionCount field to given value.

### HasReferralRedemptionCount

`func (o *Campaign) HasReferralRedemptionCount() bool`

HasReferralRedemptionCount returns a boolean if a field has been set.

### GetDiscountCount

`func (o *Campaign) GetDiscountCount() float32`

GetDiscountCount returns the DiscountCount field if non-nil, zero value otherwise.

### GetDiscountCountOk

`func (o *Campaign) GetDiscountCountOk() (*float32, bool)`

GetDiscountCountOk returns a tuple with the DiscountCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCount

`func (o *Campaign) SetDiscountCount(v float32)`

SetDiscountCount sets DiscountCount field to given value.

### HasDiscountCount

`func (o *Campaign) HasDiscountCount() bool`

HasDiscountCount returns a boolean if a field has been set.

### GetDiscountEffectCount

`func (o *Campaign) GetDiscountEffectCount() int32`

GetDiscountEffectCount returns the DiscountEffectCount field if non-nil, zero value otherwise.

### GetDiscountEffectCountOk

`func (o *Campaign) GetDiscountEffectCountOk() (*int32, bool)`

GetDiscountEffectCountOk returns a tuple with the DiscountEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountEffectCount

`func (o *Campaign) SetDiscountEffectCount(v int32)`

SetDiscountEffectCount sets DiscountEffectCount field to given value.

### HasDiscountEffectCount

`func (o *Campaign) HasDiscountEffectCount() bool`

HasDiscountEffectCount returns a boolean if a field has been set.

### GetCouponCreationCount

`func (o *Campaign) GetCouponCreationCount() int32`

GetCouponCreationCount returns the CouponCreationCount field if non-nil, zero value otherwise.

### GetCouponCreationCountOk

`func (o *Campaign) GetCouponCreationCountOk() (*int32, bool)`

GetCouponCreationCountOk returns a tuple with the CouponCreationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCreationCount

`func (o *Campaign) SetCouponCreationCount(v int32)`

SetCouponCreationCount sets CouponCreationCount field to given value.

### HasCouponCreationCount

`func (o *Campaign) HasCouponCreationCount() bool`

HasCouponCreationCount returns a boolean if a field has been set.

### GetCustomEffectCount

`func (o *Campaign) GetCustomEffectCount() int32`

GetCustomEffectCount returns the CustomEffectCount field if non-nil, zero value otherwise.

### GetCustomEffectCountOk

`func (o *Campaign) GetCustomEffectCountOk() (*int32, bool)`

GetCustomEffectCountOk returns a tuple with the CustomEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEffectCount

`func (o *Campaign) SetCustomEffectCount(v int32)`

SetCustomEffectCount sets CustomEffectCount field to given value.

### HasCustomEffectCount

`func (o *Campaign) HasCustomEffectCount() bool`

HasCustomEffectCount returns a boolean if a field has been set.

### GetReferralCreationCount

`func (o *Campaign) GetReferralCreationCount() int32`

GetReferralCreationCount returns the ReferralCreationCount field if non-nil, zero value otherwise.

### GetReferralCreationCountOk

`func (o *Campaign) GetReferralCreationCountOk() (*int32, bool)`

GetReferralCreationCountOk returns a tuple with the ReferralCreationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralCreationCount

`func (o *Campaign) SetReferralCreationCount(v int32)`

SetReferralCreationCount sets ReferralCreationCount field to given value.

### HasReferralCreationCount

`func (o *Campaign) HasReferralCreationCount() bool`

HasReferralCreationCount returns a boolean if a field has been set.

### GetAddFreeItemEffectCount

`func (o *Campaign) GetAddFreeItemEffectCount() int32`

GetAddFreeItemEffectCount returns the AddFreeItemEffectCount field if non-nil, zero value otherwise.

### GetAddFreeItemEffectCountOk

`func (o *Campaign) GetAddFreeItemEffectCountOk() (*int32, bool)`

GetAddFreeItemEffectCountOk returns a tuple with the AddFreeItemEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddFreeItemEffectCount

`func (o *Campaign) SetAddFreeItemEffectCount(v int32)`

SetAddFreeItemEffectCount sets AddFreeItemEffectCount field to given value.

### HasAddFreeItemEffectCount

`func (o *Campaign) HasAddFreeItemEffectCount() bool`

HasAddFreeItemEffectCount returns a boolean if a field has been set.

### GetAwardedGiveawaysCount

`func (o *Campaign) GetAwardedGiveawaysCount() int32`

GetAwardedGiveawaysCount returns the AwardedGiveawaysCount field if non-nil, zero value otherwise.

### GetAwardedGiveawaysCountOk

`func (o *Campaign) GetAwardedGiveawaysCountOk() (*int32, bool)`

GetAwardedGiveawaysCountOk returns a tuple with the AwardedGiveawaysCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardedGiveawaysCount

`func (o *Campaign) SetAwardedGiveawaysCount(v int32)`

SetAwardedGiveawaysCount sets AwardedGiveawaysCount field to given value.

### HasAwardedGiveawaysCount

`func (o *Campaign) HasAwardedGiveawaysCount() bool`

HasAwardedGiveawaysCount returns a boolean if a field has been set.

### GetCreatedLoyaltyPointsCount

`func (o *Campaign) GetCreatedLoyaltyPointsCount() float32`

GetCreatedLoyaltyPointsCount returns the CreatedLoyaltyPointsCount field if non-nil, zero value otherwise.

### GetCreatedLoyaltyPointsCountOk

`func (o *Campaign) GetCreatedLoyaltyPointsCountOk() (*float32, bool)`

GetCreatedLoyaltyPointsCountOk returns a tuple with the CreatedLoyaltyPointsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedLoyaltyPointsCount

`func (o *Campaign) SetCreatedLoyaltyPointsCount(v float32)`

SetCreatedLoyaltyPointsCount sets CreatedLoyaltyPointsCount field to given value.

### HasCreatedLoyaltyPointsCount

`func (o *Campaign) HasCreatedLoyaltyPointsCount() bool`

HasCreatedLoyaltyPointsCount returns a boolean if a field has been set.

### GetCreatedLoyaltyPointsEffectCount

`func (o *Campaign) GetCreatedLoyaltyPointsEffectCount() int32`

GetCreatedLoyaltyPointsEffectCount returns the CreatedLoyaltyPointsEffectCount field if non-nil, zero value otherwise.

### GetCreatedLoyaltyPointsEffectCountOk

`func (o *Campaign) GetCreatedLoyaltyPointsEffectCountOk() (*int32, bool)`

GetCreatedLoyaltyPointsEffectCountOk returns a tuple with the CreatedLoyaltyPointsEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedLoyaltyPointsEffectCount

`func (o *Campaign) SetCreatedLoyaltyPointsEffectCount(v int32)`

SetCreatedLoyaltyPointsEffectCount sets CreatedLoyaltyPointsEffectCount field to given value.

### HasCreatedLoyaltyPointsEffectCount

`func (o *Campaign) HasCreatedLoyaltyPointsEffectCount() bool`

HasCreatedLoyaltyPointsEffectCount returns a boolean if a field has been set.

### GetRedeemedLoyaltyPointsCount

`func (o *Campaign) GetRedeemedLoyaltyPointsCount() float32`

GetRedeemedLoyaltyPointsCount returns the RedeemedLoyaltyPointsCount field if non-nil, zero value otherwise.

### GetRedeemedLoyaltyPointsCountOk

`func (o *Campaign) GetRedeemedLoyaltyPointsCountOk() (*float32, bool)`

GetRedeemedLoyaltyPointsCountOk returns a tuple with the RedeemedLoyaltyPointsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemedLoyaltyPointsCount

`func (o *Campaign) SetRedeemedLoyaltyPointsCount(v float32)`

SetRedeemedLoyaltyPointsCount sets RedeemedLoyaltyPointsCount field to given value.

### HasRedeemedLoyaltyPointsCount

`func (o *Campaign) HasRedeemedLoyaltyPointsCount() bool`

HasRedeemedLoyaltyPointsCount returns a boolean if a field has been set.

### GetRedeemedLoyaltyPointsEffectCount

`func (o *Campaign) GetRedeemedLoyaltyPointsEffectCount() int32`

GetRedeemedLoyaltyPointsEffectCount returns the RedeemedLoyaltyPointsEffectCount field if non-nil, zero value otherwise.

### GetRedeemedLoyaltyPointsEffectCountOk

`func (o *Campaign) GetRedeemedLoyaltyPointsEffectCountOk() (*int32, bool)`

GetRedeemedLoyaltyPointsEffectCountOk returns a tuple with the RedeemedLoyaltyPointsEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemedLoyaltyPointsEffectCount

`func (o *Campaign) SetRedeemedLoyaltyPointsEffectCount(v int32)`

SetRedeemedLoyaltyPointsEffectCount sets RedeemedLoyaltyPointsEffectCount field to given value.

### HasRedeemedLoyaltyPointsEffectCount

`func (o *Campaign) HasRedeemedLoyaltyPointsEffectCount() bool`

HasRedeemedLoyaltyPointsEffectCount returns a boolean if a field has been set.

### GetCallApiEffectCount

`func (o *Campaign) GetCallApiEffectCount() int32`

GetCallApiEffectCount returns the CallApiEffectCount field if non-nil, zero value otherwise.

### GetCallApiEffectCountOk

`func (o *Campaign) GetCallApiEffectCountOk() (*int32, bool)`

GetCallApiEffectCountOk returns a tuple with the CallApiEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallApiEffectCount

`func (o *Campaign) SetCallApiEffectCount(v int32)`

SetCallApiEffectCount sets CallApiEffectCount field to given value.

### HasCallApiEffectCount

`func (o *Campaign) HasCallApiEffectCount() bool`

HasCallApiEffectCount returns a boolean if a field has been set.

### GetLastActivity

`func (o *Campaign) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *Campaign) GetLastActivityOk() (*time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *Campaign) SetLastActivity(v time.Time)`

SetLastActivity sets LastActivity field to given value.

### HasLastActivity

`func (o *Campaign) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### GetUpdated

`func (o *Campaign) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Campaign) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Campaign) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Campaign) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Campaign) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Campaign) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Campaign) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Campaign) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Campaign) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Campaign) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Campaign) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Campaign) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetTemplateId

`func (o *Campaign) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *Campaign) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *Campaign) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *Campaign) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


