# LoyaltyProgramLedgers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The internal ID of loyalty program. | 
**Title** | **string** | Visible name of loyalty program. | 
**Name** | **string** | Internal name of loyalty program. | 
**Ledger** | [**LedgerInfo**](LedgerInfo.md) |  | 
**SubLedgers** | Pointer to [**map[string]LedgerInfo**](LedgerInfo.md) | A map containing information about each loyalty subledger. | [optional] 

## Methods

### NewLoyaltyProgramLedgers

`func NewLoyaltyProgramLedgers(id int32, title string, name string, ledger LedgerInfo, ) *LoyaltyProgramLedgers`

NewLoyaltyProgramLedgers instantiates a new LoyaltyProgramLedgers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyProgramLedgersWithDefaults

`func NewLoyaltyProgramLedgersWithDefaults() *LoyaltyProgramLedgers`

NewLoyaltyProgramLedgersWithDefaults instantiates a new LoyaltyProgramLedgers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoyaltyProgramLedgers) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoyaltyProgramLedgers) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoyaltyProgramLedgers) SetId(v int32)`

SetId sets Id field to given value.


### GetTitle

`func (o *LoyaltyProgramLedgers) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LoyaltyProgramLedgers) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LoyaltyProgramLedgers) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetName

`func (o *LoyaltyProgramLedgers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoyaltyProgramLedgers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoyaltyProgramLedgers) SetName(v string)`

SetName sets Name field to given value.


### GetLedger

`func (o *LoyaltyProgramLedgers) GetLedger() LedgerInfo`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *LoyaltyProgramLedgers) GetLedgerOk() (*LedgerInfo, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *LoyaltyProgramLedgers) SetLedger(v LedgerInfo)`

SetLedger sets Ledger field to given value.


### GetSubLedgers

`func (o *LoyaltyProgramLedgers) GetSubLedgers() map[string]LedgerInfo`

GetSubLedgers returns the SubLedgers field if non-nil, zero value otherwise.

### GetSubLedgersOk

`func (o *LoyaltyProgramLedgers) GetSubLedgersOk() (*map[string]LedgerInfo, bool)`

GetSubLedgersOk returns a tuple with the SubLedgers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLedgers

`func (o *LoyaltyProgramLedgers) SetSubLedgers(v map[string]LedgerInfo)`

SetSubLedgers sets SubLedgers field to given value.

### HasSubLedgers

`func (o *LoyaltyProgramLedgers) HasSubLedgers() bool`

HasSubLedgers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


