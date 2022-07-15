# NotificationWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**Modified** | **time.Time** | The exact moment this entity was last modified. | 
**ApplicationId** | **int32** | The ID of the application that owns this entity. | 
**Url** | **string** | API url for this notification webhook. | 
**Headers** | **[]string** | List of API HTTP headers for this notification webhook. | 

## Methods

### NewNotificationWebhook

`func NewNotificationWebhook(id int32, created time.Time, modified time.Time, applicationId int32, url string, headers []string, ) *NotificationWebhook`

NewNotificationWebhook instantiates a new NotificationWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWebhookWithDefaults

`func NewNotificationWebhookWithDefaults() *NotificationWebhook`

NewNotificationWebhookWithDefaults instantiates a new NotificationWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationWebhook) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationWebhook) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationWebhook) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *NotificationWebhook) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NotificationWebhook) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NotificationWebhook) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *NotificationWebhook) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *NotificationWebhook) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *NotificationWebhook) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetApplicationId

`func (o *NotificationWebhook) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *NotificationWebhook) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *NotificationWebhook) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.


### GetUrl

`func (o *NotificationWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *NotificationWebhook) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *NotificationWebhook) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *NotificationWebhook) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


