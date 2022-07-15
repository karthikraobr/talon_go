# \CouponsApi

All URIs are relative to *https://goflink.europe-west1.talon.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCouponReservation**](CouponsApi.md#CreateCouponReservation) | **Post** /v1/coupon_reservations/{couponValue} | Create coupon reservation
[**DeleteCouponReservation**](CouponsApi.md#DeleteCouponReservation) | **Delete** /v1/coupon_reservations/{couponValue} | Delete coupon reservations
[**GetReservedCustomers**](CouponsApi.md#GetReservedCustomers) | **Get** /v1/coupon_reservations/customerprofiles/{couponValue} | List customers that have this coupon reserved



## CreateCouponReservation

> Coupon CreateCouponReservation(ctx, couponValue).Body(body).Execute()

Create coupon reservation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    couponValue := "couponValue_example" // string | The code of the coupon.
    body := *openapiclient.NewCouponReservations([]string{"IntegrationIDs_example"}) // CouponReservations | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsApi.CreateCouponReservation(context.Background(), couponValue).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsApi.CreateCouponReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCouponReservation`: Coupon
    fmt.Fprintf(os.Stdout, "Response from `CouponsApi.CreateCouponReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponValue** | **string** | The code of the coupon. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CouponReservations**](CouponReservations.md) |  | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCouponReservation

> DeleteCouponReservation(ctx, couponValue).Body(body).Execute()

Delete coupon reservations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    couponValue := "couponValue_example" // string | The code of the coupon.
    body := *openapiclient.NewCouponReservations([]string{"IntegrationIDs_example"}) // CouponReservations | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsApi.DeleteCouponReservation(context.Background(), couponValue).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsApi.DeleteCouponReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponValue** | **string** | The code of the coupon. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCouponReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CouponReservations**](CouponReservations.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReservedCustomers

> GetReservedCustomers200Response GetReservedCustomers(ctx, couponValue).Execute()

List customers that have this coupon reserved



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    couponValue := "couponValue_example" // string | The code of the coupon.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsApi.GetReservedCustomers(context.Background(), couponValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsApi.GetReservedCustomers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReservedCustomers`: GetReservedCustomers200Response
    fmt.Fprintf(os.Stdout, "Response from `CouponsApi.GetReservedCustomers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponValue** | **string** | The code of the coupon. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReservedCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetReservedCustomers200Response**](GetReservedCustomers200Response.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

