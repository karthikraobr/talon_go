# SamlLoginEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | ID of the SAML service. | 
**LoginURL** | **string** | Single Sign-On URL. | 

## Methods

### NewSamlLoginEndpoint

`func NewSamlLoginEndpoint(name string, loginURL string, ) *SamlLoginEndpoint`

NewSamlLoginEndpoint instantiates a new SamlLoginEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlLoginEndpointWithDefaults

`func NewSamlLoginEndpointWithDefaults() *SamlLoginEndpoint`

NewSamlLoginEndpointWithDefaults instantiates a new SamlLoginEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SamlLoginEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SamlLoginEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SamlLoginEndpoint) SetName(v string)`

SetName sets Name field to given value.


### GetLoginURL

`func (o *SamlLoginEndpoint) GetLoginURL() string`

GetLoginURL returns the LoginURL field if non-nil, zero value otherwise.

### GetLoginURLOk

`func (o *SamlLoginEndpoint) GetLoginURLOk() (*string, bool)`

GetLoginURLOk returns a tuple with the LoginURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginURL

`func (o *SamlLoginEndpoint) SetLoginURL(v string)`

SetLoginURL sets LoginURL field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


