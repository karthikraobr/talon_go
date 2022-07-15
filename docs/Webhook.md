# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**Modified** | **time.Time** | The exact moment this entity was last modified. | 
**ApplicationIds** | **[]int32** | The IDs of the applications that are related to this entity. The IDs of the applications that are related to this entity. | 
**Title** | **string** | Friendly title for this webhook. | 
**Verb** | **string** | API method for this webhook. | 
**Url** | **string** | API url (supports templating using parameters) for this webhook. | 
**Headers** | **[]string** | List of API HTTP headers for this webhook. | 
**Payload** | Pointer to **string** | API payload (supports templating using parameters) for this webhook. | [optional] 
**Params** | [**[]TemplateArgDef**](TemplateArgDef.md) | Array of template argument definitions. | 
**Enabled** | **bool** | Enables or disables webhook from showing in rule builder. | 

## Methods

### NewWebhook

`func NewWebhook(id int32, created time.Time, modified time.Time, applicationIds []int32, title string, verb string, url string, headers []string, params []TemplateArgDef, enabled bool, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Webhook) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Webhook) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *Webhook) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Webhook) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Webhook) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *Webhook) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Webhook) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Webhook) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetApplicationIds

`func (o *Webhook) GetApplicationIds() []int32`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *Webhook) GetApplicationIdsOk() (*[]int32, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *Webhook) SetApplicationIds(v []int32)`

SetApplicationIds sets ApplicationIds field to given value.


### GetTitle

`func (o *Webhook) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Webhook) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Webhook) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVerb

`func (o *Webhook) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *Webhook) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *Webhook) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *Webhook) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Webhook) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Webhook) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.


### GetPayload

`func (o *Webhook) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Webhook) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Webhook) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Webhook) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetParams

`func (o *Webhook) GetParams() []TemplateArgDef`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Webhook) GetParamsOk() (*[]TemplateArgDef, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Webhook) SetParams(v []TemplateArgDef)`

SetParams sets Params field to given value.


### GetEnabled

`func (o *Webhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Webhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Webhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


