# CampaignPrioritiesChangedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | [**Application**](Application.md) |  | 
**OldPriorities** | Pointer to **map[string][]int32** | Campaign IDs for each priority. The priority can be one of: [&#39;universal&#39;, &#39;stackable&#39;, &#39;exclusive&#39;]  | [optional] 
**Priorities** | **map[string][]int32** | Campaign IDs for each priority. The priority can be one of: [&#39;universal&#39;, &#39;stackable&#39;, &#39;exclusive&#39;]  | 

## Methods

### NewCampaignPrioritiesChangedNotification

`func NewCampaignPrioritiesChangedNotification(application Application, priorities map[string][]int32, ) *CampaignPrioritiesChangedNotification`

NewCampaignPrioritiesChangedNotification instantiates a new CampaignPrioritiesChangedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignPrioritiesChangedNotificationWithDefaults

`func NewCampaignPrioritiesChangedNotificationWithDefaults() *CampaignPrioritiesChangedNotification`

NewCampaignPrioritiesChangedNotificationWithDefaults instantiates a new CampaignPrioritiesChangedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *CampaignPrioritiesChangedNotification) GetApplication() Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *CampaignPrioritiesChangedNotification) GetApplicationOk() (*Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *CampaignPrioritiesChangedNotification) SetApplication(v Application)`

SetApplication sets Application field to given value.


### GetOldPriorities

`func (o *CampaignPrioritiesChangedNotification) GetOldPriorities() map[string][]int32`

GetOldPriorities returns the OldPriorities field if non-nil, zero value otherwise.

### GetOldPrioritiesOk

`func (o *CampaignPrioritiesChangedNotification) GetOldPrioritiesOk() (*map[string][]int32, bool)`

GetOldPrioritiesOk returns a tuple with the OldPriorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPriorities

`func (o *CampaignPrioritiesChangedNotification) SetOldPriorities(v map[string][]int32)`

SetOldPriorities sets OldPriorities field to given value.

### HasOldPriorities

`func (o *CampaignPrioritiesChangedNotification) HasOldPriorities() bool`

HasOldPriorities returns a boolean if a field has been set.

### GetPriorities

`func (o *CampaignPrioritiesChangedNotification) GetPriorities() map[string][]int32`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *CampaignPrioritiesChangedNotification) GetPrioritiesOk() (*map[string][]int32, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorities

`func (o *CampaignPrioritiesChangedNotification) SetPriorities(v map[string][]int32)`

SetPriorities sets Priorities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


