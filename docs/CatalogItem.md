# CatalogItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | **time.Time** | The exact moment this entity was created. | 
**Sku** | **string** | The stock keeping unit of the item. | 
**Price** | Pointer to **float32** | Price of the item. | [optional] 
**Catalogid** | **int32** | The ID of the catalog the item belongs to. | 
**Version** | **int32** | The version of the catalog item. | 
**Attributes** | Pointer to [**[]ItemAttribute**](ItemAttribute.md) |  | [optional] 

## Methods

### NewCatalogItem

`func NewCatalogItem(id int32, created time.Time, sku string, catalogid int32, version int32, ) *CatalogItem`

NewCatalogItem instantiates a new CatalogItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemWithDefaults

`func NewCatalogItemWithDefaults() *CatalogItem`

NewCatalogItemWithDefaults instantiates a new CatalogItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogItem) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *CatalogItem) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CatalogItem) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CatalogItem) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetSku

`func (o *CatalogItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CatalogItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CatalogItem) SetSku(v string)`

SetSku sets Sku field to given value.


### GetPrice

`func (o *CatalogItem) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogItem) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CatalogItem) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CatalogItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCatalogid

`func (o *CatalogItem) GetCatalogid() int32`

GetCatalogid returns the Catalogid field if non-nil, zero value otherwise.

### GetCatalogidOk

`func (o *CatalogItem) GetCatalogidOk() (*int32, bool)`

GetCatalogidOk returns a tuple with the Catalogid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogid

`func (o *CatalogItem) SetCatalogid(v int32)`

SetCatalogid sets Catalogid field to given value.


### GetVersion

`func (o *CatalogItem) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalogItem) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CatalogItem) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetAttributes

`func (o *CatalogItem) GetAttributes() []ItemAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CatalogItem) GetAttributesOk() (*[]ItemAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CatalogItem) SetAttributes(v []ItemAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CatalogItem) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


