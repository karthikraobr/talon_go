# CampaignStateChangedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Campaign** | [**Campaign**](Campaign.md) |  | 
**OldState** | **string** | The campaign&#39;s old state. Can be one of the following: [&#39;running&#39;, &#39;disabled&#39;, &#39;scheduled&#39;, &#39;expired&#39;, &#39;draft&#39;, &#39;archived&#39;]  | 
**NewState** | **string** | The campaign&#39;s new state. Can be one of the following: [&#39;running&#39;, &#39;disabled&#39;, &#39;scheduled&#39;, &#39;expired&#39;, &#39;draft&#39;, &#39;archived&#39;]  | 

## Methods

### NewCampaignStateChangedNotification

`func NewCampaignStateChangedNotification(campaign Campaign, oldState string, newState string, ) *CampaignStateChangedNotification`

NewCampaignStateChangedNotification instantiates a new CampaignStateChangedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignStateChangedNotificationWithDefaults

`func NewCampaignStateChangedNotificationWithDefaults() *CampaignStateChangedNotification`

NewCampaignStateChangedNotificationWithDefaults instantiates a new CampaignStateChangedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaign

`func (o *CampaignStateChangedNotification) GetCampaign() Campaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *CampaignStateChangedNotification) GetCampaignOk() (*Campaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *CampaignStateChangedNotification) SetCampaign(v Campaign)`

SetCampaign sets Campaign field to given value.


### GetOldState

`func (o *CampaignStateChangedNotification) GetOldState() string`

GetOldState returns the OldState field if non-nil, zero value otherwise.

### GetOldStateOk

`func (o *CampaignStateChangedNotification) GetOldStateOk() (*string, bool)`

GetOldStateOk returns a tuple with the OldState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldState

`func (o *CampaignStateChangedNotification) SetOldState(v string)`

SetOldState sets OldState field to given value.


### GetNewState

`func (o *CampaignStateChangedNotification) GetNewState() string`

GetNewState returns the NewState field if non-nil, zero value otherwise.

### GetNewStateOk

`func (o *CampaignStateChangedNotification) GetNewStateOk() (*string, bool)`

GetNewStateOk returns a tuple with the NewState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewState

`func (o *CampaignStateChangedNotification) SetNewState(v string)`

SetNewState sets NewState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


