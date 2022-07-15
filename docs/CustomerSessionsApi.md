# \CustomerSessionsApi

All URIs are relative to *https://goflink.europe-west1.talon.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomerSession**](CustomerSessionsApi.md#GetCustomerSession) | **Get** /v2/customer_sessions/{customerSessionId} | Get customer session
[**ReturnCartItems**](CustomerSessionsApi.md#ReturnCartItems) | **Post** /v2/customer_sessions/{customerSessionId}/returns | Return cart items
[**UpdateCustomerSessionV2**](CustomerSessionsApi.md#UpdateCustomerSessionV2) | **Put** /v2/customer_sessions/{customerSessionId} | Update customer session



## GetCustomerSession

> IntegrationCustomerSessionResponse GetCustomerSession(ctx, customerSessionId).Execute()

Get customer session



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
    customerSessionId := "customerSessionId_example" // string | The `integration ID` of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager's **Sessions** menu, or via the [List Application session endpoint](https://docs.talon.one/management-api/#operation/getApplicationSessions). 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerSessionsApi.GetCustomerSession(context.Background(), customerSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerSessionsApi.GetCustomerSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerSession`: IntegrationCustomerSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerSessionsApi.GetCustomerSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSessionId** | **string** | The &#x60;integration ID&#x60; of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager&#39;s **Sessions** menu, or via the [List Application session endpoint](https://docs.talon.one/management-api/#operation/getApplicationSessions).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationCustomerSessionResponse**](IntegrationCustomerSessionResponse.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnCartItems

> IntegrationStateV2 ReturnCartItems(ctx, customerSessionId).Body(body).Dry(dry).Execute()

Return cart items



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
    customerSessionId := "customerSessionId_example" // string | The `integration ID` of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager's **Sessions** menu, or via the [List Application session endpoint](https://docs.talon.one/management-api/#operation/getApplicationSessions). 
    body := *openapiclient.NewReturnIntegrationRequest(*openapiclient.NewNewReturn([]openapiclient.ReturnedCartItem{*openapiclient.NewReturnedCartItem(int32(2))})) // ReturnIntegrationRequest | 
    dry := true // bool | Indicates whether to persist the changes. Changes are ignored when `dry=true`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerSessionsApi.ReturnCartItems(context.Background(), customerSessionId).Body(body).Dry(dry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerSessionsApi.ReturnCartItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnCartItems`: IntegrationStateV2
    fmt.Fprintf(os.Stdout, "Response from `CustomerSessionsApi.ReturnCartItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSessionId** | **string** | The &#x60;integration ID&#x60; of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager&#39;s **Sessions** menu, or via the [List Application session endpoint](https://docs.talon.one/management-api/#operation/getApplicationSessions).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnCartItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ReturnIntegrationRequest**](ReturnIntegrationRequest.md) |  | 
 **dry** | **bool** | Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;. | 

### Return type

[**IntegrationStateV2**](IntegrationStateV2.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerSessionV2

> IntegrationStateV2 UpdateCustomerSessionV2(ctx, customerSessionId).Body(body).Dry(dry).Execute()

Update customer session



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
    customerSessionId := "customerSessionId_example" // string | The `integration ID` of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager's **Sessions** menu, or via the [List Application session endpoint](https://docs.talon.one/management-api/#operation/getApplicationSessions). 
    body := *openapiclient.NewIntegrationRequest(*openapiclient.NewNewCustomerSessionV2()) // IntegrationRequest | 
    dry := true // bool | Indicates whether to persist the changes. Changes are ignored when `dry=true`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerSessionsApi.UpdateCustomerSessionV2(context.Background(), customerSessionId).Body(body).Dry(dry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerSessionsApi.UpdateCustomerSessionV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomerSessionV2`: IntegrationStateV2
    fmt.Fprintf(os.Stdout, "Response from `CustomerSessionsApi.UpdateCustomerSessionV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSessionId** | **string** | The &#x60;integration ID&#x60; of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager&#39;s **Sessions** menu, or via the [List Application session endpoint](https://docs.talon.one/management-api/#operation/getApplicationSessions).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerSessionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IntegrationRequest**](IntegrationRequest.md) |  | 
 **dry** | **bool** | Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;. | 

### Return type

[**IntegrationStateV2**](IntegrationStateV2.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

