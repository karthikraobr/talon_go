# LoyaltyProgramBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentBalance** | **float32** | Sum of currently active points. | 
**PendingBalance** | **float32** | Sum of pending points. | 
**ExpiredBalance** | **float32** | Sum of expired points. | 
**SpentBalance** | **float32** | Sum of spent points. | 
**TentativeCurrentBalance** | **float32** | Sum of currently active points, including points added and deducted in open sessions. | 

## Methods

### NewLoyaltyProgramBalance

`func NewLoyaltyProgramBalance(currentBalance float32, pendingBalance float32, expiredBalance float32, spentBalance float32, tentativeCurrentBalance float32, ) *LoyaltyProgramBalance`

NewLoyaltyProgramBalance instantiates a new LoyaltyProgramBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyProgramBalanceWithDefaults

`func NewLoyaltyProgramBalanceWithDefaults() *LoyaltyProgramBalance`

NewLoyaltyProgramBalanceWithDefaults instantiates a new LoyaltyProgramBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentBalance

`func (o *LoyaltyProgramBalance) GetCurrentBalance() float32`

GetCurrentBalance returns the CurrentBalance field if non-nil, zero value otherwise.

### GetCurrentBalanceOk

`func (o *LoyaltyProgramBalance) GetCurrentBalanceOk() (*float32, bool)`

GetCurrentBalanceOk returns a tuple with the CurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalance

`func (o *LoyaltyProgramBalance) SetCurrentBalance(v float32)`

SetCurrentBalance sets CurrentBalance field to given value.


### GetPendingBalance

`func (o *LoyaltyProgramBalance) GetPendingBalance() float32`

GetPendingBalance returns the PendingBalance field if non-nil, zero value otherwise.

### GetPendingBalanceOk

`func (o *LoyaltyProgramBalance) GetPendingBalanceOk() (*float32, bool)`

GetPendingBalanceOk returns a tuple with the PendingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingBalance

`func (o *LoyaltyProgramBalance) SetPendingBalance(v float32)`

SetPendingBalance sets PendingBalance field to given value.


### GetExpiredBalance

`func (o *LoyaltyProgramBalance) GetExpiredBalance() float32`

GetExpiredBalance returns the ExpiredBalance field if non-nil, zero value otherwise.

### GetExpiredBalanceOk

`func (o *LoyaltyProgramBalance) GetExpiredBalanceOk() (*float32, bool)`

GetExpiredBalanceOk returns a tuple with the ExpiredBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredBalance

`func (o *LoyaltyProgramBalance) SetExpiredBalance(v float32)`

SetExpiredBalance sets ExpiredBalance field to given value.


### GetSpentBalance

`func (o *LoyaltyProgramBalance) GetSpentBalance() float32`

GetSpentBalance returns the SpentBalance field if non-nil, zero value otherwise.

### GetSpentBalanceOk

`func (o *LoyaltyProgramBalance) GetSpentBalanceOk() (*float32, bool)`

GetSpentBalanceOk returns a tuple with the SpentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentBalance

`func (o *LoyaltyProgramBalance) SetSpentBalance(v float32)`

SetSpentBalance sets SpentBalance field to given value.


### GetTentativeCurrentBalance

`func (o *LoyaltyProgramBalance) GetTentativeCurrentBalance() float32`

GetTentativeCurrentBalance returns the TentativeCurrentBalance field if non-nil, zero value otherwise.

### GetTentativeCurrentBalanceOk

`func (o *LoyaltyProgramBalance) GetTentativeCurrentBalanceOk() (*float32, bool)`

GetTentativeCurrentBalanceOk returns a tuple with the TentativeCurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTentativeCurrentBalance

`func (o *LoyaltyProgramBalance) SetTentativeCurrentBalance(v float32)`

SetTentativeCurrentBalance sets TentativeCurrentBalance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


