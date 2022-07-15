# AccountAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | **int32** | Total number of applications in the account. | 
**LiveApplications** | **int32** | Total number of live applications in the account. | 
**SandboxApplications** | **int32** | Total number of sandbox applications in the account. | 
**Campaigns** | **int32** | Total number of campaigns in the account. | 
**ActiveCampaigns** | **int32** | Total number of active campaigns in the account. | 
**LiveActiveCampaigns** | **int32** | Total number of active campaigns in live applications in the account. | 
**Coupons** | **int32** | Total number of coupons in the account. | 
**ActiveCoupons** | **int32** | Total number of active coupons in the account. | 
**ExpiredCoupons** | **int32** | Total number of expired coupons in the account. | 
**ReferralCodes** | **int32** | Total number of referral codes in the account. | 
**ActiveReferralCodes** | **int32** | Total number of active referral codes in the account. | 
**ExpiredReferralCodes** | **int32** | Total number of expired referral codes in the account. | 
**ActiveRules** | **int32** | Total number of active rules in the account. | 
**Users** | **int32** | Total number of users in the account. | 
**Roles** | **int32** | Total number of roles in the account. | 
**CustomAttributes** | **int32** | Total number of custom attributes in the account. | 
**Webhooks** | **int32** | Total number of webhooks in the account. | 
**LoyaltyPrograms** | **int32** | Total number of all loyalty programs in the account. | 
**LiveLoyaltyPrograms** | **int32** | Total number of live loyalty programs in the account. | 

## Methods

### NewAccountAnalytics

`func NewAccountAnalytics(applications int32, liveApplications int32, sandboxApplications int32, campaigns int32, activeCampaigns int32, liveActiveCampaigns int32, coupons int32, activeCoupons int32, expiredCoupons int32, referralCodes int32, activeReferralCodes int32, expiredReferralCodes int32, activeRules int32, users int32, roles int32, customAttributes int32, webhooks int32, loyaltyPrograms int32, liveLoyaltyPrograms int32, ) *AccountAnalytics`

NewAccountAnalytics instantiates a new AccountAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAnalyticsWithDefaults

`func NewAccountAnalyticsWithDefaults() *AccountAnalytics`

NewAccountAnalyticsWithDefaults instantiates a new AccountAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *AccountAnalytics) GetApplications() int32`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *AccountAnalytics) GetApplicationsOk() (*int32, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *AccountAnalytics) SetApplications(v int32)`

SetApplications sets Applications field to given value.


### GetLiveApplications

`func (o *AccountAnalytics) GetLiveApplications() int32`

GetLiveApplications returns the LiveApplications field if non-nil, zero value otherwise.

### GetLiveApplicationsOk

`func (o *AccountAnalytics) GetLiveApplicationsOk() (*int32, bool)`

GetLiveApplicationsOk returns a tuple with the LiveApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveApplications

`func (o *AccountAnalytics) SetLiveApplications(v int32)`

SetLiveApplications sets LiveApplications field to given value.


### GetSandboxApplications

`func (o *AccountAnalytics) GetSandboxApplications() int32`

GetSandboxApplications returns the SandboxApplications field if non-nil, zero value otherwise.

### GetSandboxApplicationsOk

`func (o *AccountAnalytics) GetSandboxApplicationsOk() (*int32, bool)`

GetSandboxApplicationsOk returns a tuple with the SandboxApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxApplications

`func (o *AccountAnalytics) SetSandboxApplications(v int32)`

SetSandboxApplications sets SandboxApplications field to given value.


### GetCampaigns

`func (o *AccountAnalytics) GetCampaigns() int32`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *AccountAnalytics) GetCampaignsOk() (*int32, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *AccountAnalytics) SetCampaigns(v int32)`

SetCampaigns sets Campaigns field to given value.


### GetActiveCampaigns

`func (o *AccountAnalytics) GetActiveCampaigns() int32`

GetActiveCampaigns returns the ActiveCampaigns field if non-nil, zero value otherwise.

### GetActiveCampaignsOk

`func (o *AccountAnalytics) GetActiveCampaignsOk() (*int32, bool)`

GetActiveCampaignsOk returns a tuple with the ActiveCampaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCampaigns

`func (o *AccountAnalytics) SetActiveCampaigns(v int32)`

SetActiveCampaigns sets ActiveCampaigns field to given value.


### GetLiveActiveCampaigns

`func (o *AccountAnalytics) GetLiveActiveCampaigns() int32`

GetLiveActiveCampaigns returns the LiveActiveCampaigns field if non-nil, zero value otherwise.

### GetLiveActiveCampaignsOk

`func (o *AccountAnalytics) GetLiveActiveCampaignsOk() (*int32, bool)`

GetLiveActiveCampaignsOk returns a tuple with the LiveActiveCampaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveActiveCampaigns

`func (o *AccountAnalytics) SetLiveActiveCampaigns(v int32)`

SetLiveActiveCampaigns sets LiveActiveCampaigns field to given value.


### GetCoupons

`func (o *AccountAnalytics) GetCoupons() int32`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *AccountAnalytics) GetCouponsOk() (*int32, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupons

`func (o *AccountAnalytics) SetCoupons(v int32)`

SetCoupons sets Coupons field to given value.


### GetActiveCoupons

`func (o *AccountAnalytics) GetActiveCoupons() int32`

GetActiveCoupons returns the ActiveCoupons field if non-nil, zero value otherwise.

### GetActiveCouponsOk

`func (o *AccountAnalytics) GetActiveCouponsOk() (*int32, bool)`

GetActiveCouponsOk returns a tuple with the ActiveCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCoupons

`func (o *AccountAnalytics) SetActiveCoupons(v int32)`

SetActiveCoupons sets ActiveCoupons field to given value.


### GetExpiredCoupons

`func (o *AccountAnalytics) GetExpiredCoupons() int32`

GetExpiredCoupons returns the ExpiredCoupons field if non-nil, zero value otherwise.

### GetExpiredCouponsOk

`func (o *AccountAnalytics) GetExpiredCouponsOk() (*int32, bool)`

GetExpiredCouponsOk returns a tuple with the ExpiredCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredCoupons

`func (o *AccountAnalytics) SetExpiredCoupons(v int32)`

SetExpiredCoupons sets ExpiredCoupons field to given value.


### GetReferralCodes

`func (o *AccountAnalytics) GetReferralCodes() int32`

GetReferralCodes returns the ReferralCodes field if non-nil, zero value otherwise.

### GetReferralCodesOk

`func (o *AccountAnalytics) GetReferralCodesOk() (*int32, bool)`

GetReferralCodesOk returns a tuple with the ReferralCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralCodes

`func (o *AccountAnalytics) SetReferralCodes(v int32)`

SetReferralCodes sets ReferralCodes field to given value.


### GetActiveReferralCodes

`func (o *AccountAnalytics) GetActiveReferralCodes() int32`

GetActiveReferralCodes returns the ActiveReferralCodes field if non-nil, zero value otherwise.

### GetActiveReferralCodesOk

`func (o *AccountAnalytics) GetActiveReferralCodesOk() (*int32, bool)`

GetActiveReferralCodesOk returns a tuple with the ActiveReferralCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveReferralCodes

`func (o *AccountAnalytics) SetActiveReferralCodes(v int32)`

SetActiveReferralCodes sets ActiveReferralCodes field to given value.


### GetExpiredReferralCodes

`func (o *AccountAnalytics) GetExpiredReferralCodes() int32`

GetExpiredReferralCodes returns the ExpiredReferralCodes field if non-nil, zero value otherwise.

### GetExpiredReferralCodesOk

`func (o *AccountAnalytics) GetExpiredReferralCodesOk() (*int32, bool)`

GetExpiredReferralCodesOk returns a tuple with the ExpiredReferralCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredReferralCodes

`func (o *AccountAnalytics) SetExpiredReferralCodes(v int32)`

SetExpiredReferralCodes sets ExpiredReferralCodes field to given value.


### GetActiveRules

`func (o *AccountAnalytics) GetActiveRules() int32`

GetActiveRules returns the ActiveRules field if non-nil, zero value otherwise.

### GetActiveRulesOk

`func (o *AccountAnalytics) GetActiveRulesOk() (*int32, bool)`

GetActiveRulesOk returns a tuple with the ActiveRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRules

`func (o *AccountAnalytics) SetActiveRules(v int32)`

SetActiveRules sets ActiveRules field to given value.


### GetUsers

`func (o *AccountAnalytics) GetUsers() int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AccountAnalytics) GetUsersOk() (*int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AccountAnalytics) SetUsers(v int32)`

SetUsers sets Users field to given value.


### GetRoles

`func (o *AccountAnalytics) GetRoles() int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AccountAnalytics) GetRolesOk() (*int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AccountAnalytics) SetRoles(v int32)`

SetRoles sets Roles field to given value.


### GetCustomAttributes

`func (o *AccountAnalytics) GetCustomAttributes() int32`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *AccountAnalytics) GetCustomAttributesOk() (*int32, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *AccountAnalytics) SetCustomAttributes(v int32)`

SetCustomAttributes sets CustomAttributes field to given value.


### GetWebhooks

`func (o *AccountAnalytics) GetWebhooks() int32`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *AccountAnalytics) GetWebhooksOk() (*int32, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *AccountAnalytics) SetWebhooks(v int32)`

SetWebhooks sets Webhooks field to given value.


### GetLoyaltyPrograms

`func (o *AccountAnalytics) GetLoyaltyPrograms() int32`

GetLoyaltyPrograms returns the LoyaltyPrograms field if non-nil, zero value otherwise.

### GetLoyaltyProgramsOk

`func (o *AccountAnalytics) GetLoyaltyProgramsOk() (*int32, bool)`

GetLoyaltyProgramsOk returns a tuple with the LoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyPrograms

`func (o *AccountAnalytics) SetLoyaltyPrograms(v int32)`

SetLoyaltyPrograms sets LoyaltyPrograms field to given value.


### GetLiveLoyaltyPrograms

`func (o *AccountAnalytics) GetLiveLoyaltyPrograms() int32`

GetLiveLoyaltyPrograms returns the LiveLoyaltyPrograms field if non-nil, zero value otherwise.

### GetLiveLoyaltyProgramsOk

`func (o *AccountAnalytics) GetLiveLoyaltyProgramsOk() (*int32, bool)`

GetLiveLoyaltyProgramsOk returns a tuple with the LiveLoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveLoyaltyPrograms

`func (o *AccountAnalytics) SetLiveLoyaltyPrograms(v int32)`

SetLiveLoyaltyPrograms sets LiveLoyaltyPrograms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


