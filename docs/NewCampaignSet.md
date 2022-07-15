# NewCampaignSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **int32** | The ID of the application that owns this entity. | 
**Version** | **int32** | Version of the campaign set. | 
**Set** | [**CampaignSetBranchNode**](CampaignSetBranchNode.md) |  | 

## Methods

### NewNewCampaignSet

`func NewNewCampaignSet(applicationId int32, version int32, set CampaignSetBranchNode, ) *NewCampaignSet`

NewNewCampaignSet instantiates a new NewCampaignSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewCampaignSetWithDefaults

`func NewNewCampaignSetWithDefaults() *NewCampaignSet`

NewNewCampaignSetWithDefaults instantiates a new NewCampaignSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *NewCampaignSet) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *NewCampaignSet) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *NewCampaignSet) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.


### GetVersion

`func (o *NewCampaignSet) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NewCampaignSet) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NewCampaignSet) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetSet

`func (o *NewCampaignSet) GetSet() CampaignSetBranchNode`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *NewCampaignSet) GetSetOk() (*CampaignSetBranchNode, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *NewCampaignSet) SetSet(v CampaignSetBranchNode)`

SetSet sets Set field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


