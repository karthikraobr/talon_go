# ApplicationApiHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | **string** | One-word summary of the health of the API connection of an application. | 
**LastUsed** | **time.Time** | time of last request relevant to the API health test. | 

## Methods

### NewApplicationApiHealth

`func NewApplicationApiHealth(summary string, lastUsed time.Time, ) *ApplicationApiHealth`

NewApplicationApiHealth instantiates a new ApplicationApiHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationApiHealthWithDefaults

`func NewApplicationApiHealthWithDefaults() *ApplicationApiHealth`

NewApplicationApiHealthWithDefaults instantiates a new ApplicationApiHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *ApplicationApiHealth) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ApplicationApiHealth) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ApplicationApiHealth) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetLastUsed

`func (o *ApplicationApiHealth) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *ApplicationApiHealth) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *ApplicationApiHealth) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


