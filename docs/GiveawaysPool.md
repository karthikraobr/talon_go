# GiveawaysPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**AccountId** | **int32** | The ID of the account that owns this entity. | 
**Name** | **string** | The name of this giveaways pool. | 
**Description** | Pointer to **string** | The description of this giveaways pool. | [optional] 
**SubscribedApplicationsIds** | Pointer to **[]int32** | A list of the IDs of the applications that this giveaways pool is enabled for. | [optional] 
**Modified** | Pointer to **time.Time** | Timestamp of the most recent update to the giveaways pool. | [optional] 
**CreatedBy** | **int32** | ID of the user who created this giveaways pool. | 
**ModifiedBy** | Pointer to **int32** | ID of the user who last updated this giveaways pool if available. | [optional] 

## Methods

### NewGiveawaysPool

`func NewGiveawaysPool(id int32, created time.Time, accountId int32, name string, createdBy int32, ) *GiveawaysPool`

NewGiveawaysPool instantiates a new GiveawaysPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiveawaysPoolWithDefaults

`func NewGiveawaysPoolWithDefaults() *GiveawaysPool`

NewGiveawaysPoolWithDefaults instantiates a new GiveawaysPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GiveawaysPool) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GiveawaysPool) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GiveawaysPool) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *GiveawaysPool) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GiveawaysPool) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GiveawaysPool) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetAccountId

`func (o *GiveawaysPool) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GiveawaysPool) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GiveawaysPool) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *GiveawaysPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GiveawaysPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GiveawaysPool) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GiveawaysPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GiveawaysPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GiveawaysPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GiveawaysPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubscribedApplicationsIds

`func (o *GiveawaysPool) GetSubscribedApplicationsIds() []int32`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *GiveawaysPool) GetSubscribedApplicationsIdsOk() (*[]int32, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplicationsIds

`func (o *GiveawaysPool) SetSubscribedApplicationsIds(v []int32)`

SetSubscribedApplicationsIds sets SubscribedApplicationsIds field to given value.

### HasSubscribedApplicationsIds

`func (o *GiveawaysPool) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### GetModified

`func (o *GiveawaysPool) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *GiveawaysPool) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *GiveawaysPool) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *GiveawaysPool) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GiveawaysPool) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GiveawaysPool) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GiveawaysPool) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedBy

`func (o *GiveawaysPool) GetModifiedBy() int32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *GiveawaysPool) GetModifiedByOk() (*int32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *GiveawaysPool) SetModifiedBy(v int32)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *GiveawaysPool) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


