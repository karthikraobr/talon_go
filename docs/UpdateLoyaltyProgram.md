# UpdateLoyaltyProgram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The display title for the Loyalty Program. | [optional] 
**Description** | Pointer to **string** | Description of our Loyalty Program. | [optional] 
**SubscribedApplications** | Pointer to **[]int32** | A list containing the IDs of all applications that are subscribed to this Loyalty Program. | [optional] 
**DefaultValidity** | Pointer to **string** | Indicates the default duration after which new loyalty points should expire. The format is a number, followed by one letter indicating the unit; like &#39;1h&#39; or &#39;40m&#39;. | [optional] 
**DefaultPending** | Pointer to **string** | Indicates the default duration for the pending time, after which points will be valid. The format is a number followed by a duration unit, like &#39;1h&#39; or &#39;40m&#39;. | [optional] 
**AllowSubledger** | Pointer to **bool** | Indicates if this program supports subledgers inside the program. | [optional] 
**UsersPerCardLimit** | Pointer to **int32** | The max amount of user profiles with whom a card can be shared. This can be set to 0 for no limit. This property is only used when &#x60;cardBased&#x60; is &#x60;true&#x60;.  | [optional] 
**Tiers** | Pointer to [**[]NewLoyaltyTier**](NewLoyaltyTier.md) | The tiers in this loyalty program. | [optional] 

## Methods

### NewUpdateLoyaltyProgram

`func NewUpdateLoyaltyProgram() *UpdateLoyaltyProgram`

NewUpdateLoyaltyProgram instantiates a new UpdateLoyaltyProgram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoyaltyProgramWithDefaults

`func NewUpdateLoyaltyProgramWithDefaults() *UpdateLoyaltyProgram`

NewUpdateLoyaltyProgramWithDefaults instantiates a new UpdateLoyaltyProgram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateLoyaltyProgram) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateLoyaltyProgram) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateLoyaltyProgram) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateLoyaltyProgram) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateLoyaltyProgram) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateLoyaltyProgram) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateLoyaltyProgram) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateLoyaltyProgram) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubscribedApplications

`func (o *UpdateLoyaltyProgram) GetSubscribedApplications() []int32`

GetSubscribedApplications returns the SubscribedApplications field if non-nil, zero value otherwise.

### GetSubscribedApplicationsOk

`func (o *UpdateLoyaltyProgram) GetSubscribedApplicationsOk() (*[]int32, bool)`

GetSubscribedApplicationsOk returns a tuple with the SubscribedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplications

`func (o *UpdateLoyaltyProgram) SetSubscribedApplications(v []int32)`

SetSubscribedApplications sets SubscribedApplications field to given value.

### HasSubscribedApplications

`func (o *UpdateLoyaltyProgram) HasSubscribedApplications() bool`

HasSubscribedApplications returns a boolean if a field has been set.

### GetDefaultValidity

`func (o *UpdateLoyaltyProgram) GetDefaultValidity() string`

GetDefaultValidity returns the DefaultValidity field if non-nil, zero value otherwise.

### GetDefaultValidityOk

`func (o *UpdateLoyaltyProgram) GetDefaultValidityOk() (*string, bool)`

GetDefaultValidityOk returns a tuple with the DefaultValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValidity

`func (o *UpdateLoyaltyProgram) SetDefaultValidity(v string)`

SetDefaultValidity sets DefaultValidity field to given value.

### HasDefaultValidity

`func (o *UpdateLoyaltyProgram) HasDefaultValidity() bool`

HasDefaultValidity returns a boolean if a field has been set.

### GetDefaultPending

`func (o *UpdateLoyaltyProgram) GetDefaultPending() string`

GetDefaultPending returns the DefaultPending field if non-nil, zero value otherwise.

### GetDefaultPendingOk

`func (o *UpdateLoyaltyProgram) GetDefaultPendingOk() (*string, bool)`

GetDefaultPendingOk returns a tuple with the DefaultPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPending

`func (o *UpdateLoyaltyProgram) SetDefaultPending(v string)`

SetDefaultPending sets DefaultPending field to given value.

### HasDefaultPending

`func (o *UpdateLoyaltyProgram) HasDefaultPending() bool`

HasDefaultPending returns a boolean if a field has been set.

### GetAllowSubledger

`func (o *UpdateLoyaltyProgram) GetAllowSubledger() bool`

GetAllowSubledger returns the AllowSubledger field if non-nil, zero value otherwise.

### GetAllowSubledgerOk

`func (o *UpdateLoyaltyProgram) GetAllowSubledgerOk() (*bool, bool)`

GetAllowSubledgerOk returns a tuple with the AllowSubledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSubledger

`func (o *UpdateLoyaltyProgram) SetAllowSubledger(v bool)`

SetAllowSubledger sets AllowSubledger field to given value.

### HasAllowSubledger

`func (o *UpdateLoyaltyProgram) HasAllowSubledger() bool`

HasAllowSubledger returns a boolean if a field has been set.

### GetUsersPerCardLimit

`func (o *UpdateLoyaltyProgram) GetUsersPerCardLimit() int32`

GetUsersPerCardLimit returns the UsersPerCardLimit field if non-nil, zero value otherwise.

### GetUsersPerCardLimitOk

`func (o *UpdateLoyaltyProgram) GetUsersPerCardLimitOk() (*int32, bool)`

GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersPerCardLimit

`func (o *UpdateLoyaltyProgram) SetUsersPerCardLimit(v int32)`

SetUsersPerCardLimit sets UsersPerCardLimit field to given value.

### HasUsersPerCardLimit

`func (o *UpdateLoyaltyProgram) HasUsersPerCardLimit() bool`

HasUsersPerCardLimit returns a boolean if a field has been set.

### GetTiers

`func (o *UpdateLoyaltyProgram) GetTiers() []NewLoyaltyTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *UpdateLoyaltyProgram) GetTiersOk() (*[]NewLoyaltyTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *UpdateLoyaltyProgram) SetTiers(v []NewLoyaltyTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *UpdateLoyaltyProgram) HasTiers() bool`

HasTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


