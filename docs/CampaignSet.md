# CampaignSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**ApplicationId** | **int32** | The ID of the application that owns this entity. | 
**Version** | **int32** | Version of the campaign set. | 
**Set** | [**CampaignSetBranchNode**](CampaignSetBranchNode.md) |  | 

## Methods

### NewCampaignSet

`func NewCampaignSet(id int32, created time.Time, applicationId int32, version int32, set CampaignSetBranchNode, ) *CampaignSet`

NewCampaignSet instantiates a new CampaignSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignSetWithDefaults

`func NewCampaignSetWithDefaults() *CampaignSet`

NewCampaignSetWithDefaults instantiates a new CampaignSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignSet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignSet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignSet) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *CampaignSet) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignSet) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CampaignSet) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetApplicationId

`func (o *CampaignSet) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CampaignSet) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CampaignSet) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.


### GetVersion

`func (o *CampaignSet) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CampaignSet) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CampaignSet) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetSet

`func (o *CampaignSet) GetSet() CampaignSetBranchNode`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *CampaignSet) GetSetOk() (*CampaignSetBranchNode, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *CampaignSet) SetSet(v CampaignSetBranchNode)`

SetSet sets Set field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


