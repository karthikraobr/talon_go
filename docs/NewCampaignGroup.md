# NewCampaignGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this campaign group. | 
**Description** | Pointer to **string** | A longer description of the campaign group. | [optional] 
**SubscribedApplicationsIds** | Pointer to **[]int32** | A list of the IDs of the applications that this campaign group is enabled for. | [optional] 
**CampaignIds** | Pointer to **[]int32** | A list of the IDs of the campaigns that this campaign group owns. | [optional] 

## Methods

### NewNewCampaignGroup

`func NewNewCampaignGroup(name string, ) *NewCampaignGroup`

NewNewCampaignGroup instantiates a new NewCampaignGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewCampaignGroupWithDefaults

`func NewNewCampaignGroupWithDefaults() *NewCampaignGroup`

NewNewCampaignGroupWithDefaults instantiates a new NewCampaignGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewCampaignGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewCampaignGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewCampaignGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NewCampaignGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewCampaignGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewCampaignGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewCampaignGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubscribedApplicationsIds

`func (o *NewCampaignGroup) GetSubscribedApplicationsIds() []int32`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *NewCampaignGroup) GetSubscribedApplicationsIdsOk() (*[]int32, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplicationsIds

`func (o *NewCampaignGroup) SetSubscribedApplicationsIds(v []int32)`

SetSubscribedApplicationsIds sets SubscribedApplicationsIds field to given value.

### HasSubscribedApplicationsIds

`func (o *NewCampaignGroup) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### GetCampaignIds

`func (o *NewCampaignGroup) GetCampaignIds() []int32`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *NewCampaignGroup) GetCampaignIdsOk() (*[]int32, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIds

`func (o *NewCampaignGroup) SetCampaignIds(v []int32)`

SetCampaignIds sets CampaignIds field to given value.

### HasCampaignIds

`func (o *NewCampaignGroup) HasCampaignIds() bool`

HasCampaignIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


