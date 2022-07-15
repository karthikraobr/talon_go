# CollectionWithoutPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**AccountId** | **int32** | The ID of the account that owns this entity. | 
**Modified** | **time.Time** | The exact moment this entity was last modified. | 
**Description** | Pointer to **string** | A short description of the purpose of this collection. | [optional] 
**SubscribedApplicationsIds** | Pointer to **[]int32** | A list of the IDs of the Applications where this collection is enabled. | [optional] 
**Name** | **string** | The name of this collection. | 
**ModifiedBy** | Pointer to **int32** | ID of the user who last updated this effect if available. | [optional] 
**CreatedBy** | **int32** | ID of the user who created this effect. | 
**ApplicationId** | Pointer to **int32** | The ID of the Application that owns this entity. | [optional] 
**CampaignId** | Pointer to **int32** | The ID of the campaign that owns this entity. | [optional] 

## Methods

### NewCollectionWithoutPayload

`func NewCollectionWithoutPayload(id int32, created time.Time, accountId int32, modified time.Time, name string, createdBy int32, ) *CollectionWithoutPayload`

NewCollectionWithoutPayload instantiates a new CollectionWithoutPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionWithoutPayloadWithDefaults

`func NewCollectionWithoutPayloadWithDefaults() *CollectionWithoutPayload`

NewCollectionWithoutPayloadWithDefaults instantiates a new CollectionWithoutPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CollectionWithoutPayload) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionWithoutPayload) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionWithoutPayload) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *CollectionWithoutPayload) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CollectionWithoutPayload) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CollectionWithoutPayload) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetAccountId

`func (o *CollectionWithoutPayload) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CollectionWithoutPayload) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CollectionWithoutPayload) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetModified

`func (o *CollectionWithoutPayload) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CollectionWithoutPayload) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CollectionWithoutPayload) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetDescription

`func (o *CollectionWithoutPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionWithoutPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionWithoutPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CollectionWithoutPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubscribedApplicationsIds

`func (o *CollectionWithoutPayload) GetSubscribedApplicationsIds() []int32`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *CollectionWithoutPayload) GetSubscribedApplicationsIdsOk() (*[]int32, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplicationsIds

`func (o *CollectionWithoutPayload) SetSubscribedApplicationsIds(v []int32)`

SetSubscribedApplicationsIds sets SubscribedApplicationsIds field to given value.

### HasSubscribedApplicationsIds

`func (o *CollectionWithoutPayload) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### GetName

`func (o *CollectionWithoutPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectionWithoutPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectionWithoutPayload) SetName(v string)`

SetName sets Name field to given value.


### GetModifiedBy

`func (o *CollectionWithoutPayload) GetModifiedBy() int32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *CollectionWithoutPayload) GetModifiedByOk() (*int32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *CollectionWithoutPayload) SetModifiedBy(v int32)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *CollectionWithoutPayload) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CollectionWithoutPayload) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CollectionWithoutPayload) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CollectionWithoutPayload) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.


### GetApplicationId

`func (o *CollectionWithoutPayload) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CollectionWithoutPayload) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CollectionWithoutPayload) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *CollectionWithoutPayload) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetCampaignId

`func (o *CollectionWithoutPayload) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CollectionWithoutPayload) GetCampaignIdOk() (*int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CollectionWithoutPayload) SetCampaignId(v int32)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *CollectionWithoutPayload) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


