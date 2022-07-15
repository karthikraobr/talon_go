# WebhookActivationLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationRequestUuid** | **string** | UUID reference of the integration request that triggered the effect with the webhook. | 
**WebhookId** | **int32** | ID of the webhook that triggered the request. | 
**ApplicationId** | **int32** | ID of the application that triggered the webhook. | 
**CampaignId** | **int32** | ID of the campaign that triggered the webhook. | 
**Created** | **time.Time** | Timestamp of request | 

## Methods

### NewWebhookActivationLogEntry

`func NewWebhookActivationLogEntry(integrationRequestUuid string, webhookId int32, applicationId int32, campaignId int32, created time.Time, ) *WebhookActivationLogEntry`

NewWebhookActivationLogEntry instantiates a new WebhookActivationLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookActivationLogEntryWithDefaults

`func NewWebhookActivationLogEntryWithDefaults() *WebhookActivationLogEntry`

NewWebhookActivationLogEntryWithDefaults instantiates a new WebhookActivationLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationRequestUuid

`func (o *WebhookActivationLogEntry) GetIntegrationRequestUuid() string`

GetIntegrationRequestUuid returns the IntegrationRequestUuid field if non-nil, zero value otherwise.

### GetIntegrationRequestUuidOk

`func (o *WebhookActivationLogEntry) GetIntegrationRequestUuidOk() (*string, bool)`

GetIntegrationRequestUuidOk returns a tuple with the IntegrationRequestUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationRequestUuid

`func (o *WebhookActivationLogEntry) SetIntegrationRequestUuid(v string)`

SetIntegrationRequestUuid sets IntegrationRequestUuid field to given value.


### GetWebhookId

`func (o *WebhookActivationLogEntry) GetWebhookId() int32`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookActivationLogEntry) GetWebhookIdOk() (*int32, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookActivationLogEntry) SetWebhookId(v int32)`

SetWebhookId sets WebhookId field to given value.


### GetApplicationId

`func (o *WebhookActivationLogEntry) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *WebhookActivationLogEntry) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *WebhookActivationLogEntry) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.


### GetCampaignId

`func (o *WebhookActivationLogEntry) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *WebhookActivationLogEntry) GetCampaignIdOk() (*int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *WebhookActivationLogEntry) SetCampaignId(v int32)`

SetCampaignId sets CampaignId field to given value.


### GetCreated

`func (o *WebhookActivationLogEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebhookActivationLogEntry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WebhookActivationLogEntry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


