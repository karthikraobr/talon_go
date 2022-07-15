# LoyaltyProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projections** | Pointer to [**[]LoyaltyProjectionData**](LoyaltyProjectionData.md) |  | [optional] 
**TotalExpiringPoints** | **float32** | Sum of points to be expired by the projection date set in the query parameter. | 
**TotalActivatingPoints** | **float32** | Sum of points to be active by the projection date set in the query parameter. | 

## Methods

### NewLoyaltyProjection

`func NewLoyaltyProjection(totalExpiringPoints float32, totalActivatingPoints float32, ) *LoyaltyProjection`

NewLoyaltyProjection instantiates a new LoyaltyProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyProjectionWithDefaults

`func NewLoyaltyProjectionWithDefaults() *LoyaltyProjection`

NewLoyaltyProjectionWithDefaults instantiates a new LoyaltyProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjections

`func (o *LoyaltyProjection) GetProjections() []LoyaltyProjectionData`

GetProjections returns the Projections field if non-nil, zero value otherwise.

### GetProjectionsOk

`func (o *LoyaltyProjection) GetProjectionsOk() (*[]LoyaltyProjectionData, bool)`

GetProjectionsOk returns a tuple with the Projections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjections

`func (o *LoyaltyProjection) SetProjections(v []LoyaltyProjectionData)`

SetProjections sets Projections field to given value.

### HasProjections

`func (o *LoyaltyProjection) HasProjections() bool`

HasProjections returns a boolean if a field has been set.

### GetTotalExpiringPoints

`func (o *LoyaltyProjection) GetTotalExpiringPoints() float32`

GetTotalExpiringPoints returns the TotalExpiringPoints field if non-nil, zero value otherwise.

### GetTotalExpiringPointsOk

`func (o *LoyaltyProjection) GetTotalExpiringPointsOk() (*float32, bool)`

GetTotalExpiringPointsOk returns a tuple with the TotalExpiringPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExpiringPoints

`func (o *LoyaltyProjection) SetTotalExpiringPoints(v float32)`

SetTotalExpiringPoints sets TotalExpiringPoints field to given value.


### GetTotalActivatingPoints

`func (o *LoyaltyProjection) GetTotalActivatingPoints() float32`

GetTotalActivatingPoints returns the TotalActivatingPoints field if non-nil, zero value otherwise.

### GetTotalActivatingPointsOk

`func (o *LoyaltyProjection) GetTotalActivatingPointsOk() (*float32, bool)`

GetTotalActivatingPointsOk returns a tuple with the TotalActivatingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActivatingPoints

`func (o *LoyaltyProjection) SetTotalActivatingPoints(v float32)`

SetTotalActivatingPoints sets TotalActivatingPoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


