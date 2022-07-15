# NewLoyaltyProgram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The display title for the Loyalty Program. | 
**Description** | Pointer to **string** | Description of our Loyalty Program. | [optional] 
**SubscribedApplications** | Pointer to **[]int32** | A list containing the IDs of all applications that are subscribed to this Loyalty Program. | [optional] 
**DefaultValidity** | **string** | Indicates the default duration after which new loyalty points should expire. The format is a number, followed by one letter indicating the unit; like &#39;1h&#39; or &#39;40m&#39;. | 
**DefaultPending** | **string** | Indicates the default duration for the pending time, after which points will be valid. The format is a number followed by a duration unit, like &#39;1h&#39; or &#39;40m&#39;. | 
**AllowSubledger** | **bool** | Indicates if this program supports subledgers inside the program. | 
**UsersPerCardLimit** | Pointer to **int32** | The max amount of user profiles with whom a card can be shared. This can be set to 0 for no limit. This property is only used when &#x60;cardBased&#x60; is &#x60;true&#x60;.  | [optional] 
**Name** | **string** | The internal name for the Loyalty Program. This is an immutable value. | 
**Tiers** | Pointer to [**[]NewLoyaltyTier**](NewLoyaltyTier.md) | The tiers in this loyalty program. | [optional] 
**Timezone** | **string** | A string containing an IANA timezone descriptor. | 
**CardBased** | **bool** | Defines the type of loyalty program: - &#x60;true&#x60;: the program is a card-based. - &#x60;false&#x60;: the program is profile-based.  | [default to false]

## Methods

### NewNewLoyaltyProgram

`func NewNewLoyaltyProgram(title string, defaultValidity string, defaultPending string, allowSubledger bool, name string, timezone string, cardBased bool, ) *NewLoyaltyProgram`

NewNewLoyaltyProgram instantiates a new NewLoyaltyProgram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewLoyaltyProgramWithDefaults

`func NewNewLoyaltyProgramWithDefaults() *NewLoyaltyProgram`

NewNewLoyaltyProgramWithDefaults instantiates a new NewLoyaltyProgram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NewLoyaltyProgram) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewLoyaltyProgram) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewLoyaltyProgram) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *NewLoyaltyProgram) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewLoyaltyProgram) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewLoyaltyProgram) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewLoyaltyProgram) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubscribedApplications

`func (o *NewLoyaltyProgram) GetSubscribedApplications() []int32`

GetSubscribedApplications returns the SubscribedApplications field if non-nil, zero value otherwise.

### GetSubscribedApplicationsOk

`func (o *NewLoyaltyProgram) GetSubscribedApplicationsOk() (*[]int32, bool)`

GetSubscribedApplicationsOk returns a tuple with the SubscribedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplications

`func (o *NewLoyaltyProgram) SetSubscribedApplications(v []int32)`

SetSubscribedApplications sets SubscribedApplications field to given value.

### HasSubscribedApplications

`func (o *NewLoyaltyProgram) HasSubscribedApplications() bool`

HasSubscribedApplications returns a boolean if a field has been set.

### GetDefaultValidity

`func (o *NewLoyaltyProgram) GetDefaultValidity() string`

GetDefaultValidity returns the DefaultValidity field if non-nil, zero value otherwise.

### GetDefaultValidityOk

`func (o *NewLoyaltyProgram) GetDefaultValidityOk() (*string, bool)`

GetDefaultValidityOk returns a tuple with the DefaultValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValidity

`func (o *NewLoyaltyProgram) SetDefaultValidity(v string)`

SetDefaultValidity sets DefaultValidity field to given value.


### GetDefaultPending

`func (o *NewLoyaltyProgram) GetDefaultPending() string`

GetDefaultPending returns the DefaultPending field if non-nil, zero value otherwise.

### GetDefaultPendingOk

`func (o *NewLoyaltyProgram) GetDefaultPendingOk() (*string, bool)`

GetDefaultPendingOk returns a tuple with the DefaultPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPending

`func (o *NewLoyaltyProgram) SetDefaultPending(v string)`

SetDefaultPending sets DefaultPending field to given value.


### GetAllowSubledger

`func (o *NewLoyaltyProgram) GetAllowSubledger() bool`

GetAllowSubledger returns the AllowSubledger field if non-nil, zero value otherwise.

### GetAllowSubledgerOk

`func (o *NewLoyaltyProgram) GetAllowSubledgerOk() (*bool, bool)`

GetAllowSubledgerOk returns a tuple with the AllowSubledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSubledger

`func (o *NewLoyaltyProgram) SetAllowSubledger(v bool)`

SetAllowSubledger sets AllowSubledger field to given value.


### GetUsersPerCardLimit

`func (o *NewLoyaltyProgram) GetUsersPerCardLimit() int32`

GetUsersPerCardLimit returns the UsersPerCardLimit field if non-nil, zero value otherwise.

### GetUsersPerCardLimitOk

`func (o *NewLoyaltyProgram) GetUsersPerCardLimitOk() (*int32, bool)`

GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersPerCardLimit

`func (o *NewLoyaltyProgram) SetUsersPerCardLimit(v int32)`

SetUsersPerCardLimit sets UsersPerCardLimit field to given value.

### HasUsersPerCardLimit

`func (o *NewLoyaltyProgram) HasUsersPerCardLimit() bool`

HasUsersPerCardLimit returns a boolean if a field has been set.

### GetName

`func (o *NewLoyaltyProgram) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewLoyaltyProgram) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewLoyaltyProgram) SetName(v string)`

SetName sets Name field to given value.


### GetTiers

`func (o *NewLoyaltyProgram) GetTiers() []NewLoyaltyTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *NewLoyaltyProgram) GetTiersOk() (*[]NewLoyaltyTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *NewLoyaltyProgram) SetTiers(v []NewLoyaltyTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *NewLoyaltyProgram) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetTimezone

`func (o *NewLoyaltyProgram) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *NewLoyaltyProgram) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *NewLoyaltyProgram) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetCardBased

`func (o *NewLoyaltyProgram) GetCardBased() bool`

GetCardBased returns the CardBased field if non-nil, zero value otherwise.

### GetCardBasedOk

`func (o *NewLoyaltyProgram) GetCardBasedOk() (*bool, bool)`

GetCardBasedOk returns a tuple with the CardBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBased

`func (o *NewLoyaltyProgram) SetCardBased(v bool)`

SetCardBased sets CardBased field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


