# CardLedgerTransactionLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | The exact moment the loyalty ledger transaction happened. | 
**ProgramId** | **int32** | ID of the loyalty program. | 
**CardIdentifier** | **string** | Identifier of the loyalty card. | 
**SessionId** | **string** | ID of the customer session. | 
**Type** | **string** | Type of transaction. Possible values are:   - &#x60;addition&#x60;: Points were added.   - &#x60;subtraction&#x60;: Points were subtracted.  | 
**Name** | **string** | Name or reason of the loyalty ledger transaction. | 
**StartDate** | **string** | Start date of the loyalty ledger entry. Possible values are:   - &#x60;immediate&#x60;: Points ere active immediately.   - &#x60;timestamp value&#x60;: Points are become active from certain date.  | 
**ExpiryDate** | **string** | Expiry date of the loyalty ledger entry. Possible values are:   - &#x60;unlimited&#x60;: Points have no expiration date.   - &#x60;timestamp value&#x60;: Points have certain expiration date.  | 
**SubledgerId** | **string** | ID of the subledger. | 
**Amount** | **float32** | Amount of loyalty points of the loyalty ledger transaction. | 
**Id** | **int32** | ID of the loyalty ledger transaction. | 

## Methods

### NewCardLedgerTransactionLogEntry

`func NewCardLedgerTransactionLogEntry(created time.Time, programId int32, cardIdentifier string, sessionId string, type_ string, name string, startDate string, expiryDate string, subledgerId string, amount float32, id int32, ) *CardLedgerTransactionLogEntry`

NewCardLedgerTransactionLogEntry instantiates a new CardLedgerTransactionLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardLedgerTransactionLogEntryWithDefaults

`func NewCardLedgerTransactionLogEntryWithDefaults() *CardLedgerTransactionLogEntry`

NewCardLedgerTransactionLogEntryWithDefaults instantiates a new CardLedgerTransactionLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CardLedgerTransactionLogEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CardLedgerTransactionLogEntry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CardLedgerTransactionLogEntry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetProgramId

`func (o *CardLedgerTransactionLogEntry) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *CardLedgerTransactionLogEntry) GetProgramIdOk() (*int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *CardLedgerTransactionLogEntry) SetProgramId(v int32)`

SetProgramId sets ProgramId field to given value.


### GetCardIdentifier

`func (o *CardLedgerTransactionLogEntry) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *CardLedgerTransactionLogEntry) GetCardIdentifierOk() (*string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIdentifier

`func (o *CardLedgerTransactionLogEntry) SetCardIdentifier(v string)`

SetCardIdentifier sets CardIdentifier field to given value.


### GetSessionId

`func (o *CardLedgerTransactionLogEntry) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CardLedgerTransactionLogEntry) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CardLedgerTransactionLogEntry) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetType

`func (o *CardLedgerTransactionLogEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CardLedgerTransactionLogEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CardLedgerTransactionLogEntry) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CardLedgerTransactionLogEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardLedgerTransactionLogEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardLedgerTransactionLogEntry) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *CardLedgerTransactionLogEntry) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardLedgerTransactionLogEntry) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CardLedgerTransactionLogEntry) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetExpiryDate

`func (o *CardLedgerTransactionLogEntry) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CardLedgerTransactionLogEntry) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *CardLedgerTransactionLogEntry) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.


### GetSubledgerId

`func (o *CardLedgerTransactionLogEntry) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *CardLedgerTransactionLogEntry) GetSubledgerIdOk() (*string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerId

`func (o *CardLedgerTransactionLogEntry) SetSubledgerId(v string)`

SetSubledgerId sets SubledgerId field to given value.


### GetAmount

`func (o *CardLedgerTransactionLogEntry) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardLedgerTransactionLogEntry) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CardLedgerTransactionLogEntry) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetId

`func (o *CardLedgerTransactionLogEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardLedgerTransactionLogEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardLedgerTransactionLogEntry) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


