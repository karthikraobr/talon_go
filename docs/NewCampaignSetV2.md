# NewCampaignSetV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **int32** | The ID of the application that owns this entity. | 
**Version** | **int32** | Version of the campaign set. | 
**Set** | [**CampaignPrioritiesV2**](CampaignPrioritiesV2.md) |  | 

## Methods

### NewNewCampaignSetV2

`func NewNewCampaignSetV2(applicationId int32, version int32, set CampaignPrioritiesV2, ) *NewCampaignSetV2`

NewNewCampaignSetV2 instantiates a new NewCampaignSetV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewCampaignSetV2WithDefaults

`func NewNewCampaignSetV2WithDefaults() *NewCampaignSetV2`

NewNewCampaignSetV2WithDefaults instantiates a new NewCampaignSetV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *NewCampaignSetV2) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *NewCampaignSetV2) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *NewCampaignSetV2) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.


### GetVersion

`func (o *NewCampaignSetV2) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NewCampaignSetV2) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NewCampaignSetV2) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetSet

`func (o *NewCampaignSetV2) GetSet() CampaignPrioritiesV2`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *NewCampaignSetV2) GetSetOk() (*CampaignPrioritiesV2, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *NewCampaignSetV2) SetSet(v CampaignPrioritiesV2)`

SetSet sets Set field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


