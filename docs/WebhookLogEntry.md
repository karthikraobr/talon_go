# WebhookLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | UUID reference of the webhook request. | 
**IntegrationRequestUuid** | **string** | UUID reference of the integration request linked to this webhook request. | 
**WebhookId** | **int32** | ID of the webhook that triggered the request. | 
**ApplicationId** | Pointer to **int32** | ID of the application that triggered the webhook. | [optional] 
**Url** | **string** | Target url of request | 
**Request** | **string** | Request message | 
**Response** | Pointer to **string** | Response message | [optional] 
**Status** | Pointer to **int32** | HTTP status code of response. | [optional] 
**RequestTime** | **time.Time** | Timestamp of request | 
**ResponseTime** | Pointer to **time.Time** | Timestamp of response | [optional] 

## Methods

### NewWebhookLogEntry

`func NewWebhookLogEntry(id string, integrationRequestUuid string, webhookId int32, url string, request string, requestTime time.Time, ) *WebhookLogEntry`

NewWebhookLogEntry instantiates a new WebhookLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookLogEntryWithDefaults

`func NewWebhookLogEntryWithDefaults() *WebhookLogEntry`

NewWebhookLogEntryWithDefaults instantiates a new WebhookLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookLogEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookLogEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookLogEntry) SetId(v string)`

SetId sets Id field to given value.


### GetIntegrationRequestUuid

`func (o *WebhookLogEntry) GetIntegrationRequestUuid() string`

GetIntegrationRequestUuid returns the IntegrationRequestUuid field if non-nil, zero value otherwise.

### GetIntegrationRequestUuidOk

`func (o *WebhookLogEntry) GetIntegrationRequestUuidOk() (*string, bool)`

GetIntegrationRequestUuidOk returns a tuple with the IntegrationRequestUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationRequestUuid

`func (o *WebhookLogEntry) SetIntegrationRequestUuid(v string)`

SetIntegrationRequestUuid sets IntegrationRequestUuid field to given value.


### GetWebhookId

`func (o *WebhookLogEntry) GetWebhookId() int32`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookLogEntry) GetWebhookIdOk() (*int32, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookLogEntry) SetWebhookId(v int32)`

SetWebhookId sets WebhookId field to given value.


### GetApplicationId

`func (o *WebhookLogEntry) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *WebhookLogEntry) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *WebhookLogEntry) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *WebhookLogEntry) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookLogEntry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookLogEntry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookLogEntry) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRequest

`func (o *WebhookLogEntry) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *WebhookLogEntry) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *WebhookLogEntry) SetRequest(v string)`

SetRequest sets Request field to given value.


### GetResponse

`func (o *WebhookLogEntry) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *WebhookLogEntry) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *WebhookLogEntry) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *WebhookLogEntry) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookLogEntry) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookLogEntry) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookLogEntry) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookLogEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequestTime

`func (o *WebhookLogEntry) GetRequestTime() time.Time`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *WebhookLogEntry) GetRequestTimeOk() (*time.Time, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *WebhookLogEntry) SetRequestTime(v time.Time)`

SetRequestTime sets RequestTime field to given value.


### GetResponseTime

`func (o *WebhookLogEntry) GetResponseTime() time.Time`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *WebhookLogEntry) GetResponseTimeOk() (*time.Time, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTime

`func (o *WebhookLogEntry) SetResponseTime(v time.Time)`

SetResponseTime sets ResponseTime field to given value.

### HasResponseTime

`func (o *WebhookLogEntry) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


