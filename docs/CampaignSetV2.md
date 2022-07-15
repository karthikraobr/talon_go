# CampaignSetV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**ApplicationId** | **int32** | The ID of the application that owns this entity. | 
**Version** | **int32** | Version of the campaign set. | 
**Set** | [**CampaignPrioritiesV2**](CampaignPrioritiesV2.md) |  | 

## Methods

### NewCampaignSetV2

`func NewCampaignSetV2(id int32, created time.Time, applicationId int32, version int32, set CampaignPrioritiesV2, ) *CampaignSetV2`

NewCampaignSetV2 instantiates a new CampaignSetV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignSetV2WithDefaults

`func NewCampaignSetV2WithDefaults() *CampaignSetV2`

NewCampaignSetV2WithDefaults instantiates a new CampaignSetV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignSetV2) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignSetV2) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignSetV2) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *CampaignSetV2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignSetV2) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CampaignSetV2) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetApplicationId

`func (o *CampaignSetV2) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CampaignSetV2) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CampaignSetV2) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.


### GetVersion

`func (o *CampaignSetV2) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CampaignSetV2) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CampaignSetV2) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetSet

`func (o *CampaignSetV2) GetSet() CampaignPrioritiesV2`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *CampaignSetV2) GetSetOk() (*CampaignPrioritiesV2, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *CampaignSetV2) SetSet(v CampaignPrioritiesV2)`

SetSet sets Set field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


