# ProfileAudiencesChanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adds** | **[]int32** | The IDs of the audiences for the customer to join. | 
**Deletes** | **[]int32** | The IDs of the audiences for the customer to leave. | 

## Methods

### NewProfileAudiencesChanges

`func NewProfileAudiencesChanges(adds []int32, deletes []int32, ) *ProfileAudiencesChanges`

NewProfileAudiencesChanges instantiates a new ProfileAudiencesChanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileAudiencesChangesWithDefaults

`func NewProfileAudiencesChangesWithDefaults() *ProfileAudiencesChanges`

NewProfileAudiencesChangesWithDefaults instantiates a new ProfileAudiencesChanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdds

`func (o *ProfileAudiencesChanges) GetAdds() []int32`

GetAdds returns the Adds field if non-nil, zero value otherwise.

### GetAddsOk

`func (o *ProfileAudiencesChanges) GetAddsOk() (*[]int32, bool)`

GetAddsOk returns a tuple with the Adds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdds

`func (o *ProfileAudiencesChanges) SetAdds(v []int32)`

SetAdds sets Adds field to given value.


### GetDeletes

`func (o *ProfileAudiencesChanges) GetDeletes() []int32`

GetDeletes returns the Deletes field if non-nil, zero value otherwise.

### GetDeletesOk

`func (o *ProfileAudiencesChanges) GetDeletesOk() (*[]int32, bool)`

GetDeletesOk returns a tuple with the Deletes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletes

`func (o *ProfileAudiencesChanges) SetDeletes(v []int32)`

SetDeletes sets Deletes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


