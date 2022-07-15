# ApplicationEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **int32** | The ID of the application that owns this entity. | 

## Methods

### NewApplicationEntity

`func NewApplicationEntity(applicationId int32, ) *ApplicationEntity`

NewApplicationEntity instantiates a new ApplicationEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationEntityWithDefaults

`func NewApplicationEntityWithDefaults() *ApplicationEntity`

NewApplicationEntityWithDefaults instantiates a new ApplicationEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ApplicationEntity) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationEntity) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationEntity) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


