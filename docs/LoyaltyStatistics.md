# LoyaltyStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **time.Time** | Date at which data point was collected. | 
**TotalActivePoints** | **float32** | Total of active points for this loyalty program. | 
**TotalPendingPoints** | **float32** | Total of pending points for this loyalty program. | 
**TotalSpentPoints** | **float32** | Total of spent points for this loyalty program. | 
**TotalExpiredPoints** | **float32** | Total of expired points for this loyalty program. | 
**TotalMembers** | **float32** | Number of loyalty program members. | 
**NewMembers** | **float32** | Number of members who joined on this day. | 
**SpentPoints** | [**LoyaltyDashboardPointsBreakdown**](LoyaltyDashboardPointsBreakdown.md) |  | 
**EarnedPoints** | [**LoyaltyDashboardPointsBreakdown**](LoyaltyDashboardPointsBreakdown.md) |  | 

## Methods

### NewLoyaltyStatistics

`func NewLoyaltyStatistics(date time.Time, totalActivePoints float32, totalPendingPoints float32, totalSpentPoints float32, totalExpiredPoints float32, totalMembers float32, newMembers float32, spentPoints LoyaltyDashboardPointsBreakdown, earnedPoints LoyaltyDashboardPointsBreakdown, ) *LoyaltyStatistics`

NewLoyaltyStatistics instantiates a new LoyaltyStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyStatisticsWithDefaults

`func NewLoyaltyStatisticsWithDefaults() *LoyaltyStatistics`

NewLoyaltyStatisticsWithDefaults instantiates a new LoyaltyStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *LoyaltyStatistics) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *LoyaltyStatistics) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *LoyaltyStatistics) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetTotalActivePoints

`func (o *LoyaltyStatistics) GetTotalActivePoints() float32`

GetTotalActivePoints returns the TotalActivePoints field if non-nil, zero value otherwise.

### GetTotalActivePointsOk

`func (o *LoyaltyStatistics) GetTotalActivePointsOk() (*float32, bool)`

GetTotalActivePointsOk returns a tuple with the TotalActivePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActivePoints

`func (o *LoyaltyStatistics) SetTotalActivePoints(v float32)`

SetTotalActivePoints sets TotalActivePoints field to given value.


### GetTotalPendingPoints

`func (o *LoyaltyStatistics) GetTotalPendingPoints() float32`

GetTotalPendingPoints returns the TotalPendingPoints field if non-nil, zero value otherwise.

### GetTotalPendingPointsOk

`func (o *LoyaltyStatistics) GetTotalPendingPointsOk() (*float32, bool)`

GetTotalPendingPointsOk returns a tuple with the TotalPendingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPendingPoints

`func (o *LoyaltyStatistics) SetTotalPendingPoints(v float32)`

SetTotalPendingPoints sets TotalPendingPoints field to given value.


### GetTotalSpentPoints

`func (o *LoyaltyStatistics) GetTotalSpentPoints() float32`

GetTotalSpentPoints returns the TotalSpentPoints field if non-nil, zero value otherwise.

### GetTotalSpentPointsOk

`func (o *LoyaltyStatistics) GetTotalSpentPointsOk() (*float32, bool)`

GetTotalSpentPointsOk returns a tuple with the TotalSpentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpentPoints

`func (o *LoyaltyStatistics) SetTotalSpentPoints(v float32)`

SetTotalSpentPoints sets TotalSpentPoints field to given value.


### GetTotalExpiredPoints

`func (o *LoyaltyStatistics) GetTotalExpiredPoints() float32`

GetTotalExpiredPoints returns the TotalExpiredPoints field if non-nil, zero value otherwise.

### GetTotalExpiredPointsOk

`func (o *LoyaltyStatistics) GetTotalExpiredPointsOk() (*float32, bool)`

GetTotalExpiredPointsOk returns a tuple with the TotalExpiredPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExpiredPoints

`func (o *LoyaltyStatistics) SetTotalExpiredPoints(v float32)`

SetTotalExpiredPoints sets TotalExpiredPoints field to given value.


### GetTotalMembers

`func (o *LoyaltyStatistics) GetTotalMembers() float32`

GetTotalMembers returns the TotalMembers field if non-nil, zero value otherwise.

### GetTotalMembersOk

`func (o *LoyaltyStatistics) GetTotalMembersOk() (*float32, bool)`

GetTotalMembersOk returns a tuple with the TotalMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMembers

`func (o *LoyaltyStatistics) SetTotalMembers(v float32)`

SetTotalMembers sets TotalMembers field to given value.


### GetNewMembers

`func (o *LoyaltyStatistics) GetNewMembers() float32`

GetNewMembers returns the NewMembers field if non-nil, zero value otherwise.

### GetNewMembersOk

`func (o *LoyaltyStatistics) GetNewMembersOk() (*float32, bool)`

GetNewMembersOk returns a tuple with the NewMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMembers

`func (o *LoyaltyStatistics) SetNewMembers(v float32)`

SetNewMembers sets NewMembers field to given value.


### GetSpentPoints

`func (o *LoyaltyStatistics) GetSpentPoints() LoyaltyDashboardPointsBreakdown`

GetSpentPoints returns the SpentPoints field if non-nil, zero value otherwise.

### GetSpentPointsOk

`func (o *LoyaltyStatistics) GetSpentPointsOk() (*LoyaltyDashboardPointsBreakdown, bool)`

GetSpentPointsOk returns a tuple with the SpentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentPoints

`func (o *LoyaltyStatistics) SetSpentPoints(v LoyaltyDashboardPointsBreakdown)`

SetSpentPoints sets SpentPoints field to given value.


### GetEarnedPoints

`func (o *LoyaltyStatistics) GetEarnedPoints() LoyaltyDashboardPointsBreakdown`

GetEarnedPoints returns the EarnedPoints field if non-nil, zero value otherwise.

### GetEarnedPointsOk

`func (o *LoyaltyStatistics) GetEarnedPointsOk() (*LoyaltyDashboardPointsBreakdown, bool)`

GetEarnedPointsOk returns a tuple with the EarnedPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnedPoints

`func (o *LoyaltyStatistics) SetEarnedPoints(v LoyaltyDashboardPointsBreakdown)`

SetEarnedPoints sets EarnedPoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


