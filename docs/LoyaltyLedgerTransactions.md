# LoyaltyLedgerTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | Pointer to [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | List of transaction entries from a loyalty ledger. | [optional] 

## Methods

### NewLoyaltyLedgerTransactions

`func NewLoyaltyLedgerTransactions() *LoyaltyLedgerTransactions`

NewLoyaltyLedgerTransactions instantiates a new LoyaltyLedgerTransactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyLedgerTransactionsWithDefaults

`func NewLoyaltyLedgerTransactionsWithDefaults() *LoyaltyLedgerTransactions`

NewLoyaltyLedgerTransactionsWithDefaults instantiates a new LoyaltyLedgerTransactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *LoyaltyLedgerTransactions) GetTransactions() []LoyaltyLedgerEntry`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *LoyaltyLedgerTransactions) GetTransactionsOk() (*[]LoyaltyLedgerEntry, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *LoyaltyLedgerTransactions) SetTransactions(v []LoyaltyLedgerEntry)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *LoyaltyLedgerTransactions) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


