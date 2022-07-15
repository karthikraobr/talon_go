# LoyaltyProjectionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **time.Time** | Specific date of projection. | 
**ExpiringPoints** | **float32** | Points that will be expired by the specified date. | 
**ActivatingPoints** | **float32** | Points that will be active by the specified date. | 
**ProjectedBalance** | **float32** | Current balance plus projected active points, minus expiring points. | 

## Methods

### NewLoyaltyProjectionData

`func NewLoyaltyProjectionData(date time.Time, expiringPoints float32, activatingPoints float32, projectedBalance float32, ) *LoyaltyProjectionData`

NewLoyaltyProjectionData instantiates a new LoyaltyProjectionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyProjectionDataWithDefaults

`func NewLoyaltyProjectionDataWithDefaults() *LoyaltyProjectionData`

NewLoyaltyProjectionDataWithDefaults instantiates a new LoyaltyProjectionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *LoyaltyProjectionData) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *LoyaltyProjectionData) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *LoyaltyProjectionData) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetExpiringPoints

`func (o *LoyaltyProjectionData) GetExpiringPoints() float32`

GetExpiringPoints returns the ExpiringPoints field if non-nil, zero value otherwise.

### GetExpiringPointsOk

`func (o *LoyaltyProjectionData) GetExpiringPointsOk() (*float32, bool)`

GetExpiringPointsOk returns a tuple with the ExpiringPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiringPoints

`func (o *LoyaltyProjectionData) SetExpiringPoints(v float32)`

SetExpiringPoints sets ExpiringPoints field to given value.


### GetActivatingPoints

`func (o *LoyaltyProjectionData) GetActivatingPoints() float32`

GetActivatingPoints returns the ActivatingPoints field if non-nil, zero value otherwise.

### GetActivatingPointsOk

`func (o *LoyaltyProjectionData) GetActivatingPointsOk() (*float32, bool)`

GetActivatingPointsOk returns a tuple with the ActivatingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatingPoints

`func (o *LoyaltyProjectionData) SetActivatingPoints(v float32)`

SetActivatingPoints sets ActivatingPoints field to given value.


### GetProjectedBalance

`func (o *LoyaltyProjectionData) GetProjectedBalance() float32`

GetProjectedBalance returns the ProjectedBalance field if non-nil, zero value otherwise.

### GetProjectedBalanceOk

`func (o *LoyaltyProjectionData) GetProjectedBalanceOk() (*float32, bool)`

GetProjectedBalanceOk returns a tuple with the ProjectedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedBalance

`func (o *LoyaltyProjectionData) SetProjectedBalance(v float32)`

SetProjectedBalance sets ProjectedBalance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


