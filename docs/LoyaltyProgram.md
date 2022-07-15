# LoyaltyProgram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of loyalty program. Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**Title** | **string** | The display title for the Loyalty Program. | 
**Description** | **string** | Description of our Loyalty Program. | 
**SubscribedApplications** | **[]int32** | A list containing the IDs of all applications that are subscribed to this Loyalty Program. | 
**DefaultValidity** | **string** | Indicates the default duration after which new loyalty points should expire. The format is a number, followed by one letter indicating the unit; like &#39;1h&#39; or &#39;40m&#39;. | 
**DefaultPending** | **string** | Indicates the default duration for the pending time, after which points will be valid. The format is a number followed by a duration unit, like &#39;1h&#39; or &#39;40m&#39;. | 
**AllowSubledger** | **bool** | Indicates if this program supports subledgers inside the program. | 
**UsersPerCardLimit** | Pointer to **int32** | The max amount of user profiles with whom a card can be shared. This can be set to 0 for no limit. This property is only used when &#x60;cardBased&#x60; is &#x60;true&#x60;.  | [optional] 
**AccountID** | **int32** | The ID of the Talon.One account that owns this program. | 
**Name** | **string** | The internal name for the Loyalty Program. This is an immutable value. | 
**Tiers** | Pointer to [**[]LoyaltyTier**](LoyaltyTier.md) | The tiers in this loyalty program. | [optional] 
**Timezone** | **string** | A string containing an IANA timezone descriptor. | 
**CardBased** | **bool** | Defines the type of loyalty program: - &#x60;true&#x60;: the program is a card-based. - &#x60;false&#x60;: the program is profile-based.  | [default to false]

## Methods

### NewLoyaltyProgram

`func NewLoyaltyProgram(id int32, created time.Time, title string, description string, subscribedApplications []int32, defaultValidity string, defaultPending string, allowSubledger bool, accountID int32, name string, timezone string, cardBased bool, ) *LoyaltyProgram`

NewLoyaltyProgram instantiates a new LoyaltyProgram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyProgramWithDefaults

`func NewLoyaltyProgramWithDefaults() *LoyaltyProgram`

NewLoyaltyProgramWithDefaults instantiates a new LoyaltyProgram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoyaltyProgram) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoyaltyProgram) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoyaltyProgram) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *LoyaltyProgram) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoyaltyProgram) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LoyaltyProgram) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetTitle

`func (o *LoyaltyProgram) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LoyaltyProgram) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LoyaltyProgram) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *LoyaltyProgram) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoyaltyProgram) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LoyaltyProgram) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSubscribedApplications

`func (o *LoyaltyProgram) GetSubscribedApplications() []int32`

GetSubscribedApplications returns the SubscribedApplications field if non-nil, zero value otherwise.

### GetSubscribedApplicationsOk

`func (o *LoyaltyProgram) GetSubscribedApplicationsOk() (*[]int32, bool)`

GetSubscribedApplicationsOk returns a tuple with the SubscribedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplications

`func (o *LoyaltyProgram) SetSubscribedApplications(v []int32)`

SetSubscribedApplications sets SubscribedApplications field to given value.


### GetDefaultValidity

`func (o *LoyaltyProgram) GetDefaultValidity() string`

GetDefaultValidity returns the DefaultValidity field if non-nil, zero value otherwise.

### GetDefaultValidityOk

`func (o *LoyaltyProgram) GetDefaultValidityOk() (*string, bool)`

GetDefaultValidityOk returns a tuple with the DefaultValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValidity

`func (o *LoyaltyProgram) SetDefaultValidity(v string)`

SetDefaultValidity sets DefaultValidity field to given value.


### GetDefaultPending

`func (o *LoyaltyProgram) GetDefaultPending() string`

GetDefaultPending returns the DefaultPending field if non-nil, zero value otherwise.

### GetDefaultPendingOk

`func (o *LoyaltyProgram) GetDefaultPendingOk() (*string, bool)`

GetDefaultPendingOk returns a tuple with the DefaultPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPending

`func (o *LoyaltyProgram) SetDefaultPending(v string)`

SetDefaultPending sets DefaultPending field to given value.


### GetAllowSubledger

`func (o *LoyaltyProgram) GetAllowSubledger() bool`

GetAllowSubledger returns the AllowSubledger field if non-nil, zero value otherwise.

### GetAllowSubledgerOk

`func (o *LoyaltyProgram) GetAllowSubledgerOk() (*bool, bool)`

GetAllowSubledgerOk returns a tuple with the AllowSubledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSubledger

`func (o *LoyaltyProgram) SetAllowSubledger(v bool)`

SetAllowSubledger sets AllowSubledger field to given value.


### GetUsersPerCardLimit

`func (o *LoyaltyProgram) GetUsersPerCardLimit() int32`

GetUsersPerCardLimit returns the UsersPerCardLimit field if non-nil, zero value otherwise.

### GetUsersPerCardLimitOk

`func (o *LoyaltyProgram) GetUsersPerCardLimitOk() (*int32, bool)`

GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersPerCardLimit

`func (o *LoyaltyProgram) SetUsersPerCardLimit(v int32)`

SetUsersPerCardLimit sets UsersPerCardLimit field to given value.

### HasUsersPerCardLimit

`func (o *LoyaltyProgram) HasUsersPerCardLimit() bool`

HasUsersPerCardLimit returns a boolean if a field has been set.

### GetAccountID

`func (o *LoyaltyProgram) GetAccountID() int32`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *LoyaltyProgram) GetAccountIDOk() (*int32, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *LoyaltyProgram) SetAccountID(v int32)`

SetAccountID sets AccountID field to given value.


### GetName

`func (o *LoyaltyProgram) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoyaltyProgram) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoyaltyProgram) SetName(v string)`

SetName sets Name field to given value.


### GetTiers

`func (o *LoyaltyProgram) GetTiers() []LoyaltyTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *LoyaltyProgram) GetTiersOk() (*[]LoyaltyTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *LoyaltyProgram) SetTiers(v []LoyaltyTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *LoyaltyProgram) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetTimezone

`func (o *LoyaltyProgram) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *LoyaltyProgram) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *LoyaltyProgram) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetCardBased

`func (o *LoyaltyProgram) GetCardBased() bool`

GetCardBased returns the CardBased field if non-nil, zero value otherwise.

### GetCardBasedOk

`func (o *LoyaltyProgram) GetCardBasedOk() (*bool, bool)`

GetCardBasedOk returns a tuple with the CardBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBased

`func (o *LoyaltyProgram) SetCardBased(v bool)`

SetCardBased sets CardBased field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


