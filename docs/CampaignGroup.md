# CampaignGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**Modified** | **time.Time** | The exact moment this entity was last modified. | 
**AccountId** | **int32** | The ID of the account that owns this entity. | 
**Name** | **string** | The name of this campaign group. | 
**Description** | Pointer to **string** | A longer description of the campaign group. | [optional] 
**SubscribedApplicationsIds** | Pointer to **[]int32** | A list of the IDs of the applications that this campaign group is enabled for. | [optional] 
**CampaignIds** | Pointer to **[]int32** | A list of the IDs of the campaigns that this campaign group owns. | [optional] 

## Methods

### NewCampaignGroup

`func NewCampaignGroup(id int32, created time.Time, modified time.Time, accountId int32, name string, ) *CampaignGroup`

NewCampaignGroup instantiates a new CampaignGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignGroupWithDefaults

`func NewCampaignGroupWithDefaults() *CampaignGroup`

NewCampaignGroupWithDefaults instantiates a new CampaignGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignGroup) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *CampaignGroup) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignGroup) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CampaignGroup) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *CampaignGroup) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CampaignGroup) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CampaignGroup) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetAccountId

`func (o *CampaignGroup) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CampaignGroup) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CampaignGroup) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *CampaignGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CampaignGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubscribedApplicationsIds

`func (o *CampaignGroup) GetSubscribedApplicationsIds() []int32`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *CampaignGroup) GetSubscribedApplicationsIdsOk() (*[]int32, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplicationsIds

`func (o *CampaignGroup) SetSubscribedApplicationsIds(v []int32)`

SetSubscribedApplicationsIds sets SubscribedApplicationsIds field to given value.

### HasSubscribedApplicationsIds

`func (o *CampaignGroup) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### GetCampaignIds

`func (o *CampaignGroup) GetCampaignIds() []int32`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *CampaignGroup) GetCampaignIdsOk() (*[]int32, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIds

`func (o *CampaignGroup) SetCampaignIds(v []int32)`

SetCampaignIds sets CampaignIds field to given value.

### HasCampaignIds

`func (o *CampaignGroup) HasCampaignIds() bool`

HasCampaignIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


