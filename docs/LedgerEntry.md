# LedgerEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**ProfileId** | **string** | ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known &#x60;profileId&#x60;, we recommend you use a guest &#x60;profileId&#x60;.  | 
**AccountId** | **int32** | The ID of the Talon.One account that owns this profile. | 
**LoyaltyProgramId** | **int32** | ID of the ledger. | 
**EventId** | **int32** | ID of the related event. | 
**Amount** | **int32** | Amount of loyalty points. | 
**Reason** | **string** | reason for awarding/deducting points. | 
**ExpiryDate** | **time.Time** | Expiry date of the points. | 
**ReferenceId** | Pointer to **int32** | The ID of the balancing ledgerEntry. | [optional] 

## Methods

### NewLedgerEntry

`func NewLedgerEntry(id int32, created time.Time, profileId string, accountId int32, loyaltyProgramId int32, eventId int32, amount int32, reason string, expiryDate time.Time, ) *LedgerEntry`

NewLedgerEntry instantiates a new LedgerEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerEntryWithDefaults

`func NewLedgerEntryWithDefaults() *LedgerEntry`

NewLedgerEntryWithDefaults instantiates a new LedgerEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LedgerEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LedgerEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LedgerEntry) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *LedgerEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LedgerEntry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LedgerEntry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetProfileId

`func (o *LedgerEntry) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *LedgerEntry) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *LedgerEntry) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetAccountId

`func (o *LedgerEntry) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LedgerEntry) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LedgerEntry) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetLoyaltyProgramId

`func (o *LedgerEntry) GetLoyaltyProgramId() int32`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *LedgerEntry) GetLoyaltyProgramIdOk() (*int32, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramId

`func (o *LedgerEntry) SetLoyaltyProgramId(v int32)`

SetLoyaltyProgramId sets LoyaltyProgramId field to given value.


### GetEventId

`func (o *LedgerEntry) GetEventId() int32`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *LedgerEntry) GetEventIdOk() (*int32, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *LedgerEntry) SetEventId(v int32)`

SetEventId sets EventId field to given value.


### GetAmount

`func (o *LedgerEntry) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LedgerEntry) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *LedgerEntry) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetReason

`func (o *LedgerEntry) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LedgerEntry) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *LedgerEntry) SetReason(v string)`

SetReason sets Reason field to given value.


### GetExpiryDate

`func (o *LedgerEntry) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *LedgerEntry) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *LedgerEntry) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.


### GetReferenceId

`func (o *LedgerEntry) GetReferenceId() int32`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *LedgerEntry) GetReferenceIdOk() (*int32, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *LedgerEntry) SetReferenceId(v int32)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *LedgerEntry) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


