# LoyaltyCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**ProgramID** | **int32** | The ID of the loyalty program that owns this entity. | 
**Status** | **string** | Status of the loyalty card. Can be one of: [&#39;active&#39;, &#39;disabled&#39;]  | 
**Identifier** | **string** | The alphanumeric identifier of the loyalty card. | 
**UsersPerCardLimit** | **int32** | The max amount of user profiles a card can be shared with. 0 means unlimited.  | 
**Profiles** | Pointer to [**[]LoyaltyCardProfileRegistration**](LoyaltyCardProfileRegistration.md) | Integration IDs of the customers associated with the card. | [optional] 
**Ledger** | Pointer to [**LedgerInfo**](LedgerInfo.md) |  | [optional] 
**Subledgers** | Pointer to [**map[string]LedgerInfo**](LedgerInfo.md) | Displays point balances of the card in the subledgers of the loyalty program. | [optional] 
**Modified** | Pointer to **time.Time** | Timestamp of the most recent update of the loyalty card. | [optional] 

## Methods

### NewLoyaltyCard

`func NewLoyaltyCard(id int32, created time.Time, programID int32, status string, identifier string, usersPerCardLimit int32, ) *LoyaltyCard`

NewLoyaltyCard instantiates a new LoyaltyCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyCardWithDefaults

`func NewLoyaltyCardWithDefaults() *LoyaltyCard`

NewLoyaltyCardWithDefaults instantiates a new LoyaltyCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoyaltyCard) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoyaltyCard) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoyaltyCard) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *LoyaltyCard) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoyaltyCard) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LoyaltyCard) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetProgramID

`func (o *LoyaltyCard) GetProgramID() int32`

GetProgramID returns the ProgramID field if non-nil, zero value otherwise.

### GetProgramIDOk

`func (o *LoyaltyCard) GetProgramIDOk() (*int32, bool)`

GetProgramIDOk returns a tuple with the ProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramID

`func (o *LoyaltyCard) SetProgramID(v int32)`

SetProgramID sets ProgramID field to given value.


### GetStatus

`func (o *LoyaltyCard) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoyaltyCard) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LoyaltyCard) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIdentifier

`func (o *LoyaltyCard) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *LoyaltyCard) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *LoyaltyCard) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetUsersPerCardLimit

`func (o *LoyaltyCard) GetUsersPerCardLimit() int32`

GetUsersPerCardLimit returns the UsersPerCardLimit field if non-nil, zero value otherwise.

### GetUsersPerCardLimitOk

`func (o *LoyaltyCard) GetUsersPerCardLimitOk() (*int32, bool)`

GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersPerCardLimit

`func (o *LoyaltyCard) SetUsersPerCardLimit(v int32)`

SetUsersPerCardLimit sets UsersPerCardLimit field to given value.


### GetProfiles

`func (o *LoyaltyCard) GetProfiles() []LoyaltyCardProfileRegistration`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *LoyaltyCard) GetProfilesOk() (*[]LoyaltyCardProfileRegistration, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *LoyaltyCard) SetProfiles(v []LoyaltyCardProfileRegistration)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *LoyaltyCard) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetLedger

`func (o *LoyaltyCard) GetLedger() LedgerInfo`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *LoyaltyCard) GetLedgerOk() (*LedgerInfo, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *LoyaltyCard) SetLedger(v LedgerInfo)`

SetLedger sets Ledger field to given value.

### HasLedger

`func (o *LoyaltyCard) HasLedger() bool`

HasLedger returns a boolean if a field has been set.

### GetSubledgers

`func (o *LoyaltyCard) GetSubledgers() map[string]LedgerInfo`

GetSubledgers returns the Subledgers field if non-nil, zero value otherwise.

### GetSubledgersOk

`func (o *LoyaltyCard) GetSubledgersOk() (*map[string]LedgerInfo, bool)`

GetSubledgersOk returns a tuple with the Subledgers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgers

`func (o *LoyaltyCard) SetSubledgers(v map[string]LedgerInfo)`

SetSubledgers sets Subledgers field to given value.

### HasSubledgers

`func (o *LoyaltyCard) HasSubledgers() bool`

HasSubledgers returns a boolean if a field has been set.

### GetModified

`func (o *LoyaltyCard) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *LoyaltyCard) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *LoyaltyCard) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *LoyaltyCard) HasModified() bool`

HasModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


