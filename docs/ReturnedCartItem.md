# ReturnedCartItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | **int32** | The index of the cart item in the provided customer session&#39;s &#x60;cartItems&#x60; property. | 
**Quantity** | Pointer to **int32** | Number of cart items to return. It is only available when [cart item flattening](https://docs.talon.one/docs/product/campaigns/campaign-evaluation/#flattened-cart-items) is enabled. If cart item flattening is disabled, the cart item can only be returned in its entirety.  | [optional] 

## Methods

### NewReturnedCartItem

`func NewReturnedCartItem(position int32, ) *ReturnedCartItem`

NewReturnedCartItem instantiates a new ReturnedCartItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnedCartItemWithDefaults

`func NewReturnedCartItemWithDefaults() *ReturnedCartItem`

NewReturnedCartItemWithDefaults instantiates a new ReturnedCartItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *ReturnedCartItem) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ReturnedCartItem) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ReturnedCartItem) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetQuantity

`func (o *ReturnedCartItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ReturnedCartItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ReturnedCartItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ReturnedCartItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


