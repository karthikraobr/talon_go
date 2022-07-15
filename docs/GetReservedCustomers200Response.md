# GetReservedCustomers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | **int32** |  | 
**Data** | [**[]CustomerProfile**](CustomerProfile.md) |  | 

## Methods

### NewGetReservedCustomers200Response

`func NewGetReservedCustomers200Response(totalResultSize int32, data []CustomerProfile, ) *GetReservedCustomers200Response`

NewGetReservedCustomers200Response instantiates a new GetReservedCustomers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReservedCustomers200ResponseWithDefaults

`func NewGetReservedCustomers200ResponseWithDefaults() *GetReservedCustomers200Response`

NewGetReservedCustomers200ResponseWithDefaults instantiates a new GetReservedCustomers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *GetReservedCustomers200Response) GetTotalResultSize() int32`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *GetReservedCustomers200Response) GetTotalResultSizeOk() (*int32, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *GetReservedCustomers200Response) SetTotalResultSize(v int32)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *GetReservedCustomers200Response) GetData() []CustomerProfile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetReservedCustomers200Response) GetDataOk() (*[]CustomerProfile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetReservedCustomers200Response) SetData(v []CustomerProfile)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


