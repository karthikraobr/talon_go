# Binding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A descriptive name for the value to be bound. | 
**Type** | Pointer to **string** | The kind of binding. Possible values are: - &#x60;bundle&#x60; - &#x60;cartItemFilter&#x60; - &#x60;subledgerBalance&#x60; - &#x60;templateParameter&#x60;  | [optional] 
**Expression** | **[]map[string]interface{}** | A Talang expression that will be evaluated and its result attached to the name of the binding. | 
**ValueType** | Pointer to **string** | Can be one of the following: - &#x60;string&#x60; - &#x60;number&#x60; - &#x60;boolean&#x60;  | [optional] 

## Methods

### NewBinding

`func NewBinding(name string, expression []map[string]interface{}, ) *Binding`

NewBinding instantiates a new Binding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindingWithDefaults

`func NewBindingWithDefaults() *Binding`

NewBindingWithDefaults instantiates a new Binding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Binding) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Binding) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Binding) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Binding) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Binding) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Binding) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Binding) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExpression

`func (o *Binding) GetExpression() []map[string]interface{}`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *Binding) GetExpressionOk() (*[]map[string]interface{}, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *Binding) SetExpression(v []map[string]interface{})`

SetExpression sets Expression field to given value.


### GetValueType

`func (o *Binding) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *Binding) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *Binding) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *Binding) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


