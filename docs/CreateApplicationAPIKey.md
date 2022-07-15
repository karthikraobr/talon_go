# CreateApplicationAPIKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title for API Key. | 
**Expires** | **time.Time** | The date the API key expired. | 
**Platform** | Pointer to **string** | The third-party platform the API key is valid for. Use &#x60;none&#x60; for a generic API key to be used from your own integration layer.  | [optional] 

## Methods

### NewCreateApplicationAPIKey

`func NewCreateApplicationAPIKey(title string, expires time.Time, ) *CreateApplicationAPIKey`

NewCreateApplicationAPIKey instantiates a new CreateApplicationAPIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationAPIKeyWithDefaults

`func NewCreateApplicationAPIKeyWithDefaults() *CreateApplicationAPIKey`

NewCreateApplicationAPIKeyWithDefaults instantiates a new CreateApplicationAPIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateApplicationAPIKey) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateApplicationAPIKey) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateApplicationAPIKey) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetExpires

`func (o *CreateApplicationAPIKey) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CreateApplicationAPIKey) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *CreateApplicationAPIKey) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.


### GetPlatform

`func (o *CreateApplicationAPIKey) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateApplicationAPIKey) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateApplicationAPIKey) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CreateApplicationAPIKey) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


