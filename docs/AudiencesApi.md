# \AudiencesApi

All URIs are relative to *https://goflink.europe-west1.talon.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAudienceV2**](AudiencesApi.md#CreateAudienceV2) | **Post** /v2/audiences | Create audience
[**DeleteAudienceMembershipsV2**](AudiencesApi.md#DeleteAudienceMembershipsV2) | **Delete** /v2/audiences/{audienceId}/memberships | Delete audience memberships
[**DeleteAudienceV2**](AudiencesApi.md#DeleteAudienceV2) | **Delete** /v2/audiences/{audienceId} | Delete audience
[**UpdateAudienceCustomersAttributes**](AudiencesApi.md#UpdateAudienceCustomersAttributes) | **Put** /v2/audience_customers/{audienceId}/attributes | Update profile attributes for all customers in audience
[**UpdateAudienceV2**](AudiencesApi.md#UpdateAudienceV2) | **Put** /v2/audiences/{audienceId} | Update audience name
[**UpdateCustomerProfileAudiences**](AudiencesApi.md#UpdateCustomerProfileAudiences) | **Post** /v2/customer_audiences | Update multiple customer profiles&#39; audiences



## CreateAudienceV2

> Audience CreateAudienceV2(ctx).Body(body).Execute()

Create audience



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
    body := *openapiclient.NewNewAudience("Travel audience") // NewAudience | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AudiencesApi.CreateAudienceV2(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesApi.CreateAudienceV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAudienceV2`: Audience
    fmt.Fprintf(os.Stdout, "Response from `AudiencesApi.CreateAudienceV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAudienceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewAudience**](NewAudience.md) |  | 

### Return type

[**Audience**](Audience.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAudienceMembershipsV2

> DeleteAudienceMembershipsV2(ctx, audienceId).Execute()

Delete audience memberships



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
    audienceId := int32(56) // int32 | The ID of the audience. You get it via the `id` property when [creating an audience](#operation/createAudienceV2).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AudiencesApi.DeleteAudienceMembershipsV2(context.Background(), audienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesApi.DeleteAudienceMembershipsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int32** | The ID of the audience. You get it via the &#x60;id&#x60; property when [creating an audience](#operation/createAudienceV2). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAudienceMembershipsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAudienceV2

> DeleteAudienceV2(ctx, audienceId).Execute()

Delete audience



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
    audienceId := int32(56) // int32 | The ID of the audience. You get it via the `id` property when [creating an audience](#operation/createAudienceV2).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AudiencesApi.DeleteAudienceV2(context.Background(), audienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesApi.DeleteAudienceV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int32** | The ID of the audience. You get it via the &#x60;id&#x60; property when [creating an audience](#operation/createAudienceV2). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAudienceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAudienceCustomersAttributes

> UpdateAudienceCustomersAttributes(ctx, audienceId).Body(body).Execute()

Update profile attributes for all customers in audience



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
    audienceId := int32(56) // int32 | The ID of the audience. You get it via the `id` property when [creating an audience](#operation/createAudienceV2).
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AudiencesApi.UpdateAudienceCustomersAttributes(context.Background(), audienceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesApi.UpdateAudienceCustomersAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int32** | The ID of the audience. You get it via the &#x60;id&#x60; property when [creating an audience](#operation/createAudienceV2). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAudienceCustomersAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

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


## UpdateAudienceV2

> Audience UpdateAudienceV2(ctx, audienceId).Body(body).Execute()

Update audience name



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
    audienceId := int32(56) // int32 | The ID of the audience. You get it via the `id` property when [creating an audience](#operation/createAudienceV2).
    body := *openapiclient.NewUpdateAudience("Travel audience") // UpdateAudience | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AudiencesApi.UpdateAudienceV2(context.Background(), audienceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesApi.UpdateAudienceV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAudienceV2`: Audience
    fmt.Fprintf(os.Stdout, "Response from `AudiencesApi.UpdateAudienceV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int32** | The ID of the audience. You get it via the &#x60;id&#x60; property when [creating an audience](#operation/createAudienceV2). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAudienceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateAudience**](UpdateAudience.md) |  | 

### Return type

[**Audience**](Audience.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerProfileAudiences

> UpdateCustomerProfileAudiences(ctx).Body(body).Execute()

Update multiple customer profiles' audiences



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
    body := *openapiclient.NewCustomerProfileAudienceRequest() // CustomerProfileAudienceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AudiencesApi.UpdateCustomerProfileAudiences(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesApi.UpdateCustomerProfileAudiences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerProfileAudiencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerProfileAudienceRequest**](CustomerProfileAudienceRequest.md) |  | 

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

