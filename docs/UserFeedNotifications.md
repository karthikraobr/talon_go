# UserFeedNotifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdate** | **time.Time** | Timestamp of the last request for this list. | 
**Notifications** | [**[]FeedNotification**](FeedNotification.md) | List of all notifications to notify the user about. | 

## Methods

### NewUserFeedNotifications

`func NewUserFeedNotifications(lastUpdate time.Time, notifications []FeedNotification, ) *UserFeedNotifications`

NewUserFeedNotifications instantiates a new UserFeedNotifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFeedNotificationsWithDefaults

`func NewUserFeedNotificationsWithDefaults() *UserFeedNotifications`

NewUserFeedNotificationsWithDefaults instantiates a new UserFeedNotifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdate

`func (o *UserFeedNotifications) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *UserFeedNotifications) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *UserFeedNotifications) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.


### GetNotifications

`func (o *UserFeedNotifications) GetNotifications() []FeedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *UserFeedNotifications) GetNotificationsOk() (*[]FeedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *UserFeedNotifications) SetNotifications(v []FeedNotification)`

SetNotifications sets Notifications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


