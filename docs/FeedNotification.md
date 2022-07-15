# FeedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the feed notification. | 
**Created** | **time.Time** | Timestamp of the moment this feed notification was created. | 
**Updated** | **time.Time** | Timestamp of the moment this feed notification was last updated. | 
**ArticleUrl** | **string** | URL to the feed notification in the help center. | 
**Type** | **string** | The type of the feed notification. | 
**Body** | **string** | Body of the feed notification. | 

## Methods

### NewFeedNotification

`func NewFeedNotification(title string, created time.Time, updated time.Time, articleUrl string, type_ string, body string, ) *FeedNotification`

NewFeedNotification instantiates a new FeedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedNotificationWithDefaults

`func NewFeedNotificationWithDefaults() *FeedNotification`

NewFeedNotificationWithDefaults instantiates a new FeedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *FeedNotification) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FeedNotification) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FeedNotification) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCreated

`func (o *FeedNotification) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FeedNotification) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FeedNotification) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *FeedNotification) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FeedNotification) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FeedNotification) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetArticleUrl

`func (o *FeedNotification) GetArticleUrl() string`

GetArticleUrl returns the ArticleUrl field if non-nil, zero value otherwise.

### GetArticleUrlOk

`func (o *FeedNotification) GetArticleUrlOk() (*string, bool)`

GetArticleUrlOk returns a tuple with the ArticleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleUrl

`func (o *FeedNotification) SetArticleUrl(v string)`

SetArticleUrl sets ArticleUrl field to given value.


### GetType

`func (o *FeedNotification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeedNotification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeedNotification) SetType(v string)`

SetType sets Type field to given value.


### GetBody

`func (o *FeedNotification) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *FeedNotification) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *FeedNotification) SetBody(v string)`

SetBody sets Body field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


