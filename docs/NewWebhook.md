# NewWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationIds** | **[]int32** | The IDs of the applications that are related to this entity. | 
**Title** | **string** | Friendly title for this webhook. | 
**Verb** | **string** | API method for this webhook. | 
**Url** | **string** | API url (supports templating using parameters) for this webhook. | 
**Headers** | **[]string** | List of API HTTP headers for this webhook. | 
**Payload** | Pointer to **string** | API payload (supports templating using parameters) for this webhook. | [optional] 
**Params** | [**[]TemplateArgDef**](TemplateArgDef.md) | Array of template argument definitions. | 
**Enabled** | **bool** | Enables or disables webhook from showing in rule builder. | 

## Methods

### NewNewWebhook

`func NewNewWebhook(applicationIds []int32, title string, verb string, url string, headers []string, params []TemplateArgDef, enabled bool, ) *NewWebhook`

NewNewWebhook instantiates a new NewWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewWebhookWithDefaults

`func NewNewWebhookWithDefaults() *NewWebhook`

NewNewWebhookWithDefaults instantiates a new NewWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationIds

`func (o *NewWebhook) GetApplicationIds() []int32`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *NewWebhook) GetApplicationIdsOk() (*[]int32, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *NewWebhook) SetApplicationIds(v []int32)`

SetApplicationIds sets ApplicationIds field to given value.


### GetTitle

`func (o *NewWebhook) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewWebhook) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewWebhook) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVerb

`func (o *NewWebhook) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *NewWebhook) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *NewWebhook) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetUrl

`func (o *NewWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NewWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NewWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *NewWebhook) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *NewWebhook) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *NewWebhook) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.


### GetPayload

`func (o *NewWebhook) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *NewWebhook) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *NewWebhook) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *NewWebhook) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetParams

`func (o *NewWebhook) GetParams() []TemplateArgDef`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *NewWebhook) GetParamsOk() (*[]TemplateArgDef, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *NewWebhook) SetParams(v []TemplateArgDef)`

SetParams sets Params field to given value.


### GetEnabled

`func (o *NewWebhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NewWebhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NewWebhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


