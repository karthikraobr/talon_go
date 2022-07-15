# AccountLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LiveApplications** | **int32** | Total number of allowed live applications in the account. | 
**SandboxApplications** | **int32** | Total number of allowed sandbox applications in the account. | 
**ActiveCampaigns** | **int32** | Total number of allowed active campaigns in live applications in the account. | 
**Coupons** | **int32** | Total number of allowed coupons in the account. | 
**ReferralCodes** | **int32** | Total number of allowed referral codes in the account. | 
**ActiveRules** | **int32** | Total number of allowed active rulesets in the account. | 
**LiveLoyaltyPrograms** | **int32** | Total number of allowed live loyalty programs in the account. | 
**SandboxLoyaltyPrograms** | **int32** | Total number of allowed sandbox loyalty programs in the account. | 
**Webhooks** | **int32** | Total number of allowed webhooks in the account. | 
**Users** | **int32** | Total number of allowed users in the account. | 
**ApiVolume** | **int32** | Allowed volume of API requests to the account. | 
**PromotionTypes** | **[]string** | Array of promotion types that are employed in the account. | 

## Methods

### NewAccountLimits

`func NewAccountLimits(liveApplications int32, sandboxApplications int32, activeCampaigns int32, coupons int32, referralCodes int32, activeRules int32, liveLoyaltyPrograms int32, sandboxLoyaltyPrograms int32, webhooks int32, users int32, apiVolume int32, promotionTypes []string, ) *AccountLimits`

NewAccountLimits instantiates a new AccountLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountLimitsWithDefaults

`func NewAccountLimitsWithDefaults() *AccountLimits`

NewAccountLimitsWithDefaults instantiates a new AccountLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLiveApplications

`func (o *AccountLimits) GetLiveApplications() int32`

GetLiveApplications returns the LiveApplications field if non-nil, zero value otherwise.

### GetLiveApplicationsOk

`func (o *AccountLimits) GetLiveApplicationsOk() (*int32, bool)`

GetLiveApplicationsOk returns a tuple with the LiveApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveApplications

`func (o *AccountLimits) SetLiveApplications(v int32)`

SetLiveApplications sets LiveApplications field to given value.


### GetSandboxApplications

`func (o *AccountLimits) GetSandboxApplications() int32`

GetSandboxApplications returns the SandboxApplications field if non-nil, zero value otherwise.

### GetSandboxApplicationsOk

`func (o *AccountLimits) GetSandboxApplicationsOk() (*int32, bool)`

GetSandboxApplicationsOk returns a tuple with the SandboxApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxApplications

`func (o *AccountLimits) SetSandboxApplications(v int32)`

SetSandboxApplications sets SandboxApplications field to given value.


### GetActiveCampaigns

`func (o *AccountLimits) GetActiveCampaigns() int32`

GetActiveCampaigns returns the ActiveCampaigns field if non-nil, zero value otherwise.

### GetActiveCampaignsOk

`func (o *AccountLimits) GetActiveCampaignsOk() (*int32, bool)`

GetActiveCampaignsOk returns a tuple with the ActiveCampaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCampaigns

`func (o *AccountLimits) SetActiveCampaigns(v int32)`

SetActiveCampaigns sets ActiveCampaigns field to given value.


### GetCoupons

`func (o *AccountLimits) GetCoupons() int32`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *AccountLimits) GetCouponsOk() (*int32, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupons

`func (o *AccountLimits) SetCoupons(v int32)`

SetCoupons sets Coupons field to given value.


### GetReferralCodes

`func (o *AccountLimits) GetReferralCodes() int32`

GetReferralCodes returns the ReferralCodes field if non-nil, zero value otherwise.

### GetReferralCodesOk

`func (o *AccountLimits) GetReferralCodesOk() (*int32, bool)`

GetReferralCodesOk returns a tuple with the ReferralCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralCodes

`func (o *AccountLimits) SetReferralCodes(v int32)`

SetReferralCodes sets ReferralCodes field to given value.


### GetActiveRules

`func (o *AccountLimits) GetActiveRules() int32`

GetActiveRules returns the ActiveRules field if non-nil, zero value otherwise.

### GetActiveRulesOk

`func (o *AccountLimits) GetActiveRulesOk() (*int32, bool)`

GetActiveRulesOk returns a tuple with the ActiveRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRules

`func (o *AccountLimits) SetActiveRules(v int32)`

SetActiveRules sets ActiveRules field to given value.


### GetLiveLoyaltyPrograms

`func (o *AccountLimits) GetLiveLoyaltyPrograms() int32`

GetLiveLoyaltyPrograms returns the LiveLoyaltyPrograms field if non-nil, zero value otherwise.

### GetLiveLoyaltyProgramsOk

`func (o *AccountLimits) GetLiveLoyaltyProgramsOk() (*int32, bool)`

GetLiveLoyaltyProgramsOk returns a tuple with the LiveLoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveLoyaltyPrograms

`func (o *AccountLimits) SetLiveLoyaltyPrograms(v int32)`

SetLiveLoyaltyPrograms sets LiveLoyaltyPrograms field to given value.


### GetSandboxLoyaltyPrograms

`func (o *AccountLimits) GetSandboxLoyaltyPrograms() int32`

GetSandboxLoyaltyPrograms returns the SandboxLoyaltyPrograms field if non-nil, zero value otherwise.

### GetSandboxLoyaltyProgramsOk

`func (o *AccountLimits) GetSandboxLoyaltyProgramsOk() (*int32, bool)`

GetSandboxLoyaltyProgramsOk returns a tuple with the SandboxLoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxLoyaltyPrograms

`func (o *AccountLimits) SetSandboxLoyaltyPrograms(v int32)`

SetSandboxLoyaltyPrograms sets SandboxLoyaltyPrograms field to given value.


### GetWebhooks

`func (o *AccountLimits) GetWebhooks() int32`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *AccountLimits) GetWebhooksOk() (*int32, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *AccountLimits) SetWebhooks(v int32)`

SetWebhooks sets Webhooks field to given value.


### GetUsers

`func (o *AccountLimits) GetUsers() int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AccountLimits) GetUsersOk() (*int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AccountLimits) SetUsers(v int32)`

SetUsers sets Users field to given value.


### GetApiVolume

`func (o *AccountLimits) GetApiVolume() int32`

GetApiVolume returns the ApiVolume field if non-nil, zero value otherwise.

### GetApiVolumeOk

`func (o *AccountLimits) GetApiVolumeOk() (*int32, bool)`

GetApiVolumeOk returns a tuple with the ApiVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVolume

`func (o *AccountLimits) SetApiVolume(v int32)`

SetApiVolume sets ApiVolume field to given value.


### GetPromotionTypes

`func (o *AccountLimits) GetPromotionTypes() []string`

GetPromotionTypes returns the PromotionTypes field if non-nil, zero value otherwise.

### GetPromotionTypesOk

`func (o *AccountLimits) GetPromotionTypesOk() (*[]string, bool)`

GetPromotionTypesOk returns a tuple with the PromotionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionTypes

`func (o *AccountLimits) SetPromotionTypes(v []string)`

SetPromotionTypes sets PromotionTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


