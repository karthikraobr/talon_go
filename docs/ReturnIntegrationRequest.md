# ReturnIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Return** | [**NewReturn**](NewReturn.md) |  | 
**ResponseContent** | Pointer to **[]string** | Optional list of extra data that you want to get in the response. Use this property to get as much data as you need in one request instead of sending extra requests to other endpoints.  **Note:** &#x60;ruleFailureReasons&#x60; is always part of the response when the [Application type](https://docs.talon.one/docs/product/applications/overview#application-types) is &#x60;sandbox&#x60;.  | [optional] 

## Methods

### NewReturnIntegrationRequest

`func NewReturnIntegrationRequest(return_ NewReturn, ) *ReturnIntegrationRequest`

NewReturnIntegrationRequest instantiates a new ReturnIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnIntegrationRequestWithDefaults

`func NewReturnIntegrationRequestWithDefaults() *ReturnIntegrationRequest`

NewReturnIntegrationRequestWithDefaults instantiates a new ReturnIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturn

`func (o *ReturnIntegrationRequest) GetReturn() NewReturn`

GetReturn returns the Return field if non-nil, zero value otherwise.

### GetReturnOk

`func (o *ReturnIntegrationRequest) GetReturnOk() (*NewReturn, bool)`

GetReturnOk returns a tuple with the Return field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturn

`func (o *ReturnIntegrationRequest) SetReturn(v NewReturn)`

SetReturn sets Return field to given value.


### GetResponseContent

`func (o *ReturnIntegrationRequest) GetResponseContent() []string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *ReturnIntegrationRequest) GetResponseContentOk() (*[]string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *ReturnIntegrationRequest) SetResponseContent(v []string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *ReturnIntegrationRequest) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


