# CampaignActivationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | **[]int32** | The list of IDs of the users who will receive the activation request. | 

## Methods

### NewCampaignActivationRequest

`func NewCampaignActivationRequest(userIds []int32, ) *CampaignActivationRequest`

NewCampaignActivationRequest instantiates a new CampaignActivationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignActivationRequestWithDefaults

`func NewCampaignActivationRequestWithDefaults() *CampaignActivationRequest`

NewCampaignActivationRequestWithDefaults instantiates a new CampaignActivationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIds

`func (o *CampaignActivationRequest) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *CampaignActivationRequest) GetUserIdsOk() (*[]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *CampaignActivationRequest) SetUserIds(v []int32)`

SetUserIds sets UserIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


