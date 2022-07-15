# Return

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**ApplicationId** | **int32** | The ID of the application that owns this entity. | 
**AccountId** | **int32** | The ID of the account that owns this entity. | 
**ReturnedCartItems** | [**[]ReturnedCartItem**](ReturnedCartItem.md) | List of cart items to be returned. | 
**EventId** | **int32** | The event ID of that was generated for this return. | 
**SessionId** | **int32** | The internal ID of the session this return was requested on. | 
**SessionIntegrationId** | **string** | The integration ID of the session this return was requested on. | 
**ProfileId** | Pointer to **int32** | The internal ID of the profile this return was requested on. | [optional] 
**ProfileIntegrationId** | Pointer to **string** | The integration ID of the profile this return was requested on. | [optional] 
**CreatedBy** | Pointer to **int32** | ID of the user who requested this return. | [optional] 

## Methods

### NewReturn

`func NewReturn(id int32, created time.Time, applicationId int32, accountId int32, returnedCartItems []ReturnedCartItem, eventId int32, sessionId int32, sessionIntegrationId string, ) *Return`

NewReturn instantiates a new Return object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnWithDefaults

`func NewReturnWithDefaults() *Return`

NewReturnWithDefaults instantiates a new Return object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Return) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Return) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Return) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *Return) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Return) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Return) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetApplicationId

`func (o *Return) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Return) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *Return) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.


### GetAccountId

`func (o *Return) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Return) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Return) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetReturnedCartItems

`func (o *Return) GetReturnedCartItems() []ReturnedCartItem`

GetReturnedCartItems returns the ReturnedCartItems field if non-nil, zero value otherwise.

### GetReturnedCartItemsOk

`func (o *Return) GetReturnedCartItemsOk() (*[]ReturnedCartItem, bool)`

GetReturnedCartItemsOk returns a tuple with the ReturnedCartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedCartItems

`func (o *Return) SetReturnedCartItems(v []ReturnedCartItem)`

SetReturnedCartItems sets ReturnedCartItems field to given value.


### GetEventId

`func (o *Return) GetEventId() int32`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *Return) GetEventIdOk() (*int32, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *Return) SetEventId(v int32)`

SetEventId sets EventId field to given value.


### GetSessionId

`func (o *Return) GetSessionId() int32`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *Return) GetSessionIdOk() (*int32, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *Return) SetSessionId(v int32)`

SetSessionId sets SessionId field to given value.


### GetSessionIntegrationId

`func (o *Return) GetSessionIntegrationId() string`

GetSessionIntegrationId returns the SessionIntegrationId field if non-nil, zero value otherwise.

### GetSessionIntegrationIdOk

`func (o *Return) GetSessionIntegrationIdOk() (*string, bool)`

GetSessionIntegrationIdOk returns a tuple with the SessionIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIntegrationId

`func (o *Return) SetSessionIntegrationId(v string)`

SetSessionIntegrationId sets SessionIntegrationId field to given value.


### GetProfileId

`func (o *Return) GetProfileId() int32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *Return) GetProfileIdOk() (*int32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *Return) SetProfileId(v int32)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *Return) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileIntegrationId

`func (o *Return) GetProfileIntegrationId() string`

GetProfileIntegrationId returns the ProfileIntegrationId field if non-nil, zero value otherwise.

### GetProfileIntegrationIdOk

`func (o *Return) GetProfileIntegrationIdOk() (*string, bool)`

GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationId

`func (o *Return) SetProfileIntegrationId(v string)`

SetProfileIntegrationId sets ProfileIntegrationId field to given value.

### HasProfileIntegrationId

`func (o *Return) HasProfileIntegrationId() bool`

HasProfileIntegrationId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Return) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Return) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Return) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Return) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


