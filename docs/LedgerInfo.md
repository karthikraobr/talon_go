# LedgerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentBalance** | **float32** | Sum of currently active points. | 
**PendingBalance** | **float32** | Sum of pending points. | 
**ExpiredBalance** | **float32** | Sum of expired points. | 
**SpentBalance** | **float32** | Sum of spent points. | 
**TentativeCurrentBalance** | **float32** | Sum of currently active points, including points added and deducted in open sessions. | 
**CurrentTier** | Pointer to [**Tier**](Tier.md) |  | [optional] 
**PointsToNextTier** | Pointer to **float32** | Points required to move up a tier. | [optional] 
**Projection** | Pointer to [**LoyaltyProjection**](LoyaltyProjection.md) |  | [optional] 

## Methods

### NewLedgerInfo

`func NewLedgerInfo(currentBalance float32, pendingBalance float32, expiredBalance float32, spentBalance float32, tentativeCurrentBalance float32, ) *LedgerInfo`

NewLedgerInfo instantiates a new LedgerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerInfoWithDefaults

`func NewLedgerInfoWithDefaults() *LedgerInfo`

NewLedgerInfoWithDefaults instantiates a new LedgerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentBalance

`func (o *LedgerInfo) GetCurrentBalance() float32`

GetCurrentBalance returns the CurrentBalance field if non-nil, zero value otherwise.

### GetCurrentBalanceOk

`func (o *LedgerInfo) GetCurrentBalanceOk() (*float32, bool)`

GetCurrentBalanceOk returns a tuple with the CurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalance

`func (o *LedgerInfo) SetCurrentBalance(v float32)`

SetCurrentBalance sets CurrentBalance field to given value.


### GetPendingBalance

`func (o *LedgerInfo) GetPendingBalance() float32`

GetPendingBalance returns the PendingBalance field if non-nil, zero value otherwise.

### GetPendingBalanceOk

`func (o *LedgerInfo) GetPendingBalanceOk() (*float32, bool)`

GetPendingBalanceOk returns a tuple with the PendingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingBalance

`func (o *LedgerInfo) SetPendingBalance(v float32)`

SetPendingBalance sets PendingBalance field to given value.


### GetExpiredBalance

`func (o *LedgerInfo) GetExpiredBalance() float32`

GetExpiredBalance returns the ExpiredBalance field if non-nil, zero value otherwise.

### GetExpiredBalanceOk

`func (o *LedgerInfo) GetExpiredBalanceOk() (*float32, bool)`

GetExpiredBalanceOk returns a tuple with the ExpiredBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredBalance

`func (o *LedgerInfo) SetExpiredBalance(v float32)`

SetExpiredBalance sets ExpiredBalance field to given value.


### GetSpentBalance

`func (o *LedgerInfo) GetSpentBalance() float32`

GetSpentBalance returns the SpentBalance field if non-nil, zero value otherwise.

### GetSpentBalanceOk

`func (o *LedgerInfo) GetSpentBalanceOk() (*float32, bool)`

GetSpentBalanceOk returns a tuple with the SpentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentBalance

`func (o *LedgerInfo) SetSpentBalance(v float32)`

SetSpentBalance sets SpentBalance field to given value.


### GetTentativeCurrentBalance

`func (o *LedgerInfo) GetTentativeCurrentBalance() float32`

GetTentativeCurrentBalance returns the TentativeCurrentBalance field if non-nil, zero value otherwise.

### GetTentativeCurrentBalanceOk

`func (o *LedgerInfo) GetTentativeCurrentBalanceOk() (*float32, bool)`

GetTentativeCurrentBalanceOk returns a tuple with the TentativeCurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTentativeCurrentBalance

`func (o *LedgerInfo) SetTentativeCurrentBalance(v float32)`

SetTentativeCurrentBalance sets TentativeCurrentBalance field to given value.


### GetCurrentTier

`func (o *LedgerInfo) GetCurrentTier() Tier`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *LedgerInfo) GetCurrentTierOk() (*Tier, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTier

`func (o *LedgerInfo) SetCurrentTier(v Tier)`

SetCurrentTier sets CurrentTier field to given value.

### HasCurrentTier

`func (o *LedgerInfo) HasCurrentTier() bool`

HasCurrentTier returns a boolean if a field has been set.

### GetPointsToNextTier

`func (o *LedgerInfo) GetPointsToNextTier() float32`

GetPointsToNextTier returns the PointsToNextTier field if non-nil, zero value otherwise.

### GetPointsToNextTierOk

`func (o *LedgerInfo) GetPointsToNextTierOk() (*float32, bool)`

GetPointsToNextTierOk returns a tuple with the PointsToNextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsToNextTier

`func (o *LedgerInfo) SetPointsToNextTier(v float32)`

SetPointsToNextTier sets PointsToNextTier field to given value.

### HasPointsToNextTier

`func (o *LedgerInfo) HasPointsToNextTier() bool`

HasPointsToNextTier returns a boolean if a field has been set.

### GetProjection

`func (o *LedgerInfo) GetProjection() LoyaltyProjection`

GetProjection returns the Projection field if non-nil, zero value otherwise.

### GetProjectionOk

`func (o *LedgerInfo) GetProjectionOk() (*LoyaltyProjection, bool)`

GetProjectionOk returns a tuple with the Projection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjection

`func (o *LedgerInfo) SetProjection(v LoyaltyProjection)`

SetProjection sets Projection field to given value.

### HasProjection

`func (o *LedgerInfo) HasProjection() bool`

HasProjection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


