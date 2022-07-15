# Change

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**UserId** | **int32** | The ID of the account that owns this entity. | 
**ApplicationId** | Pointer to **int32** | ID of application associated with change. | [optional] 
**Entity** | **string** | API endpoint on which the change was initiated. | 
**Old** | Pointer to **map[string]interface{}** | Resource before the change occurred. | [optional] 
**New** | Pointer to **map[string]interface{}** | Resource after the change occurred. | [optional] 

## Methods

### NewChange

`func NewChange(id int32, created time.Time, userId int32, entity string, ) *Change`

NewChange instantiates a new Change object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeWithDefaults

`func NewChangeWithDefaults() *Change`

NewChangeWithDefaults instantiates a new Change object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Change) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Change) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Change) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *Change) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Change) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Change) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUserId

`func (o *Change) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Change) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Change) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetApplicationId

`func (o *Change) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Change) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *Change) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *Change) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetEntity

`func (o *Change) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Change) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *Change) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetOld

`func (o *Change) GetOld() map[string]interface{}`

GetOld returns the Old field if non-nil, zero value otherwise.

### GetOldOk

`func (o *Change) GetOldOk() (*map[string]interface{}, bool)`

GetOldOk returns a tuple with the Old field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOld

`func (o *Change) SetOld(v map[string]interface{})`

SetOld sets Old field to given value.

### HasOld

`func (o *Change) HasOld() bool`

HasOld returns a boolean if a field has been set.

### GetNew

`func (o *Change) GetNew() map[string]interface{}`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *Change) GetNewOk() (*map[string]interface{}, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *Change) SetNew(v map[string]interface{})`

SetNew sets New field to given value.

### HasNew

`func (o *Change) HasNew() bool`

HasNew returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


