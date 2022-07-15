# IntegrationEventV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known &#x60;profileId&#x60;, we recommend you use a guest &#x60;profileId&#x60;.  | [optional] 
**Type** | **string** | A string representing the event. Must not be a reserved event name. | 
**Attributes** | Pointer to **map[string]interface{}** | Arbitrary additional JSON data associated with the event. | [optional] 
**ResponseContent** | Pointer to **[]string** | Optional list of requested information to be present on the response related to the tracking custom event.  | [optional] 

## Methods

### NewIntegrationEventV2Request

`func NewIntegrationEventV2Request(type_ string, ) *IntegrationEventV2Request`

NewIntegrationEventV2Request instantiates a new IntegrationEventV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationEventV2RequestWithDefaults

`func NewIntegrationEventV2RequestWithDefaults() *IntegrationEventV2Request`

NewIntegrationEventV2RequestWithDefaults instantiates a new IntegrationEventV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *IntegrationEventV2Request) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *IntegrationEventV2Request) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *IntegrationEventV2Request) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *IntegrationEventV2Request) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetType

`func (o *IntegrationEventV2Request) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationEventV2Request) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationEventV2Request) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *IntegrationEventV2Request) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IntegrationEventV2Request) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IntegrationEventV2Request) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IntegrationEventV2Request) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetResponseContent

`func (o *IntegrationEventV2Request) GetResponseContent() []string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *IntegrationEventV2Request) GetResponseContentOk() (*[]string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *IntegrationEventV2Request) SetResponseContent(v []string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *IntegrationEventV2Request) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


