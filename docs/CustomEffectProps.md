# CustomEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectId** | **int32** | The ID of the custom effect that was triggered. | 
**Name** | **string** | The type of the custom effect. | 
**Payload** | **map[string]interface{}** | The JSON payload of the custom effect. | 

## Methods

### NewCustomEffectProps

`func NewCustomEffectProps(effectId int32, name string, payload map[string]interface{}, ) *CustomEffectProps`

NewCustomEffectProps instantiates a new CustomEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEffectPropsWithDefaults

`func NewCustomEffectPropsWithDefaults() *CustomEffectProps`

NewCustomEffectPropsWithDefaults instantiates a new CustomEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectId

`func (o *CustomEffectProps) GetEffectId() int32`

GetEffectId returns the EffectId field if non-nil, zero value otherwise.

### GetEffectIdOk

`func (o *CustomEffectProps) GetEffectIdOk() (*int32, bool)`

GetEffectIdOk returns a tuple with the EffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectId

`func (o *CustomEffectProps) SetEffectId(v int32)`

SetEffectId sets EffectId field to given value.


### GetName

`func (o *CustomEffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEffectProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEffectProps) SetName(v string)`

SetName sets Name field to given value.


### GetPayload

`func (o *CustomEffectProps) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CustomEffectProps) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CustomEffectProps) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


