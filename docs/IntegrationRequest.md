# IntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerSession** | [**NewCustomerSessionV2**](NewCustomerSessionV2.md) |  | 
**ResponseContent** | Pointer to **[]string** | Optional list of extra data that you want to get in the response. Use this property to get as much data as you need in one request instead of sending extra requests to other endpoints.  **Note:** &#x60;ruleFailureReasons&#x60; is always part of the response when the [Application type](https://docs.talon.one/docs/product/applications/overview#application-types) is &#x60;sandbox&#x60;.  | [optional] 

## Methods

### NewIntegrationRequest

`func NewIntegrationRequest(customerSession NewCustomerSessionV2, ) *IntegrationRequest`

NewIntegrationRequest instantiates a new IntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationRequestWithDefaults

`func NewIntegrationRequestWithDefaults() *IntegrationRequest`

NewIntegrationRequestWithDefaults instantiates a new IntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerSession

`func (o *IntegrationRequest) GetCustomerSession() NewCustomerSessionV2`

GetCustomerSession returns the CustomerSession field if non-nil, zero value otherwise.

### GetCustomerSessionOk

`func (o *IntegrationRequest) GetCustomerSessionOk() (*NewCustomerSessionV2, bool)`

GetCustomerSessionOk returns a tuple with the CustomerSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSession

`func (o *IntegrationRequest) SetCustomerSession(v NewCustomerSessionV2)`

SetCustomerSession sets CustomerSession field to given value.


### GetResponseContent

`func (o *IntegrationRequest) GetResponseContent() []string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *IntegrationRequest) GetResponseContentOk() (*[]string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *IntegrationRequest) SetResponseContent(v []string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *IntegrationRequest) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


