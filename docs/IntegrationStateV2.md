# IntegrationStateV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerSession** | Pointer to [**CustomerSessionV2**](CustomerSessionV2.md) |  | [optional] 
**CustomerProfile** | Pointer to [**CustomerProfile**](CustomerProfile.md) |  | [optional] 
**Event** | Pointer to [**Event**](Event.md) |  | [optional] 
**Loyalty** | Pointer to [**Loyalty**](Loyalty.md) |  | [optional] 
**Referral** | Pointer to [**InventoryReferral**](InventoryReferral.md) |  | [optional] 
**Coupons** | Pointer to [**[]Coupon**](Coupon.md) |  | [optional] 
**TriggeredCampaigns** | Pointer to [**[]Campaign**](Campaign.md) |  | [optional] 
**Effects** | [**[]Effect**](Effect.md) |  | 
**RuleFailureReasons** | Pointer to [**[]RuleFailureReason**](RuleFailureReason.md) |  | [optional] 
**CreatedCoupons** | [**[]Coupon**](Coupon.md) |  | 
**CreatedReferrals** | [**[]Referral**](Referral.md) |  | 
**AwardedGiveaways** | Pointer to [**[]Giveaway**](Giveaway.md) |  | [optional] 
**Return** | Pointer to [**Return**](Return.md) |  | [optional] 
**PreviousReturns** | Pointer to [**[]Return**](Return.md) |  | [optional] 

## Methods

### NewIntegrationStateV2

`func NewIntegrationStateV2(effects []Effect, createdCoupons []Coupon, createdReferrals []Referral, ) *IntegrationStateV2`

NewIntegrationStateV2 instantiates a new IntegrationStateV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationStateV2WithDefaults

`func NewIntegrationStateV2WithDefaults() *IntegrationStateV2`

NewIntegrationStateV2WithDefaults instantiates a new IntegrationStateV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerSession

`func (o *IntegrationStateV2) GetCustomerSession() CustomerSessionV2`

GetCustomerSession returns the CustomerSession field if non-nil, zero value otherwise.

### GetCustomerSessionOk

`func (o *IntegrationStateV2) GetCustomerSessionOk() (*CustomerSessionV2, bool)`

GetCustomerSessionOk returns a tuple with the CustomerSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSession

`func (o *IntegrationStateV2) SetCustomerSession(v CustomerSessionV2)`

SetCustomerSession sets CustomerSession field to given value.

### HasCustomerSession

`func (o *IntegrationStateV2) HasCustomerSession() bool`

HasCustomerSession returns a boolean if a field has been set.

### GetCustomerProfile

`func (o *IntegrationStateV2) GetCustomerProfile() CustomerProfile`

GetCustomerProfile returns the CustomerProfile field if non-nil, zero value otherwise.

### GetCustomerProfileOk

`func (o *IntegrationStateV2) GetCustomerProfileOk() (*CustomerProfile, bool)`

GetCustomerProfileOk returns a tuple with the CustomerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfile

`func (o *IntegrationStateV2) SetCustomerProfile(v CustomerProfile)`

SetCustomerProfile sets CustomerProfile field to given value.

### HasCustomerProfile

`func (o *IntegrationStateV2) HasCustomerProfile() bool`

HasCustomerProfile returns a boolean if a field has been set.

### GetEvent

`func (o *IntegrationStateV2) GetEvent() Event`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *IntegrationStateV2) GetEventOk() (*Event, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *IntegrationStateV2) SetEvent(v Event)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *IntegrationStateV2) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetLoyalty

`func (o *IntegrationStateV2) GetLoyalty() Loyalty`

GetLoyalty returns the Loyalty field if non-nil, zero value otherwise.

### GetLoyaltyOk

`func (o *IntegrationStateV2) GetLoyaltyOk() (*Loyalty, bool)`

GetLoyaltyOk returns a tuple with the Loyalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyalty

`func (o *IntegrationStateV2) SetLoyalty(v Loyalty)`

SetLoyalty sets Loyalty field to given value.

### HasLoyalty

`func (o *IntegrationStateV2) HasLoyalty() bool`

HasLoyalty returns a boolean if a field has been set.

### GetReferral

`func (o *IntegrationStateV2) GetReferral() InventoryReferral`

GetReferral returns the Referral field if non-nil, zero value otherwise.

### GetReferralOk

`func (o *IntegrationStateV2) GetReferralOk() (*InventoryReferral, bool)`

GetReferralOk returns a tuple with the Referral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferral

`func (o *IntegrationStateV2) SetReferral(v InventoryReferral)`

SetReferral sets Referral field to given value.

### HasReferral

`func (o *IntegrationStateV2) HasReferral() bool`

HasReferral returns a boolean if a field has been set.

### GetCoupons

`func (o *IntegrationStateV2) GetCoupons() []Coupon`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *IntegrationStateV2) GetCouponsOk() (*[]Coupon, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupons

`func (o *IntegrationStateV2) SetCoupons(v []Coupon)`

SetCoupons sets Coupons field to given value.

### HasCoupons

`func (o *IntegrationStateV2) HasCoupons() bool`

HasCoupons returns a boolean if a field has been set.

### GetTriggeredCampaigns

`func (o *IntegrationStateV2) GetTriggeredCampaigns() []Campaign`

GetTriggeredCampaigns returns the TriggeredCampaigns field if non-nil, zero value otherwise.

### GetTriggeredCampaignsOk

`func (o *IntegrationStateV2) GetTriggeredCampaignsOk() (*[]Campaign, bool)`

GetTriggeredCampaignsOk returns a tuple with the TriggeredCampaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredCampaigns

`func (o *IntegrationStateV2) SetTriggeredCampaigns(v []Campaign)`

SetTriggeredCampaigns sets TriggeredCampaigns field to given value.

### HasTriggeredCampaigns

`func (o *IntegrationStateV2) HasTriggeredCampaigns() bool`

HasTriggeredCampaigns returns a boolean if a field has been set.

### GetEffects

`func (o *IntegrationStateV2) GetEffects() []Effect`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *IntegrationStateV2) GetEffectsOk() (*[]Effect, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *IntegrationStateV2) SetEffects(v []Effect)`

SetEffects sets Effects field to given value.


### GetRuleFailureReasons

`func (o *IntegrationStateV2) GetRuleFailureReasons() []RuleFailureReason`

GetRuleFailureReasons returns the RuleFailureReasons field if non-nil, zero value otherwise.

### GetRuleFailureReasonsOk

`func (o *IntegrationStateV2) GetRuleFailureReasonsOk() (*[]RuleFailureReason, bool)`

GetRuleFailureReasonsOk returns a tuple with the RuleFailureReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleFailureReasons

`func (o *IntegrationStateV2) SetRuleFailureReasons(v []RuleFailureReason)`

SetRuleFailureReasons sets RuleFailureReasons field to given value.

### HasRuleFailureReasons

`func (o *IntegrationStateV2) HasRuleFailureReasons() bool`

HasRuleFailureReasons returns a boolean if a field has been set.

### GetCreatedCoupons

`func (o *IntegrationStateV2) GetCreatedCoupons() []Coupon`

GetCreatedCoupons returns the CreatedCoupons field if non-nil, zero value otherwise.

### GetCreatedCouponsOk

`func (o *IntegrationStateV2) GetCreatedCouponsOk() (*[]Coupon, bool)`

GetCreatedCouponsOk returns a tuple with the CreatedCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedCoupons

`func (o *IntegrationStateV2) SetCreatedCoupons(v []Coupon)`

SetCreatedCoupons sets CreatedCoupons field to given value.


### GetCreatedReferrals

`func (o *IntegrationStateV2) GetCreatedReferrals() []Referral`

GetCreatedReferrals returns the CreatedReferrals field if non-nil, zero value otherwise.

### GetCreatedReferralsOk

`func (o *IntegrationStateV2) GetCreatedReferralsOk() (*[]Referral, bool)`

GetCreatedReferralsOk returns a tuple with the CreatedReferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedReferrals

`func (o *IntegrationStateV2) SetCreatedReferrals(v []Referral)`

SetCreatedReferrals sets CreatedReferrals field to given value.


### GetAwardedGiveaways

`func (o *IntegrationStateV2) GetAwardedGiveaways() []Giveaway`

GetAwardedGiveaways returns the AwardedGiveaways field if non-nil, zero value otherwise.

### GetAwardedGiveawaysOk

`func (o *IntegrationStateV2) GetAwardedGiveawaysOk() (*[]Giveaway, bool)`

GetAwardedGiveawaysOk returns a tuple with the AwardedGiveaways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardedGiveaways

`func (o *IntegrationStateV2) SetAwardedGiveaways(v []Giveaway)`

SetAwardedGiveaways sets AwardedGiveaways field to given value.

### HasAwardedGiveaways

`func (o *IntegrationStateV2) HasAwardedGiveaways() bool`

HasAwardedGiveaways returns a boolean if a field has been set.

### GetReturn

`func (o *IntegrationStateV2) GetReturn() Return`

GetReturn returns the Return field if non-nil, zero value otherwise.

### GetReturnOk

`func (o *IntegrationStateV2) GetReturnOk() (*Return, bool)`

GetReturnOk returns a tuple with the Return field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturn

`func (o *IntegrationStateV2) SetReturn(v Return)`

SetReturn sets Return field to given value.

### HasReturn

`func (o *IntegrationStateV2) HasReturn() bool`

HasReturn returns a boolean if a field has been set.

### GetPreviousReturns

`func (o *IntegrationStateV2) GetPreviousReturns() []Return`

GetPreviousReturns returns the PreviousReturns field if non-nil, zero value otherwise.

### GetPreviousReturnsOk

`func (o *IntegrationStateV2) GetPreviousReturnsOk() (*[]Return, bool)`

GetPreviousReturnsOk returns a tuple with the PreviousReturns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousReturns

`func (o *IntegrationStateV2) SetPreviousReturns(v []Return)`

SetPreviousReturns sets PreviousReturns field to given value.

### HasPreviousReturns

`func (o *IntegrationStateV2) HasPreviousReturns() bool`

HasPreviousReturns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


