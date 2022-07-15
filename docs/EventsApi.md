# \EventsApi

All URIs are relative to *https://goflink.europe-west1.talon.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TrackEvent**](EventsApi.md#TrackEvent) | **Post** /v1/events | Track event
[**TrackEventV2**](EventsApi.md#TrackEventV2) | **Post** /v2/events | Track event V2



## TrackEvent

> IntegrationState TrackEvent(ctx).Body(body).Dry(dry).Execute()

Track event



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
    body := *openapiclient.NewNewEvent("pageViews", map[string]interface{}({"myAttribute":"myValue"}), "175KJPS947296") // NewEvent | 
    dry := true // bool | Indicates whether to persist the changes. Changes are ignored when `dry=true`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.TrackEvent(context.Background()).Body(body).Dry(dry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.TrackEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TrackEvent`: IntegrationState
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.TrackEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewEvent**](NewEvent.md) |  | 
 **dry** | **bool** | Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;. | 

### Return type

[**IntegrationState**](IntegrationState.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackEventV2

> IntegrationStateV2 TrackEventV2(ctx).Body(body).Silent(silent).Dry(dry).Execute()

Track event V2



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
    body := *openapiclient.NewIntegrationEventV2Request("email_opened") // IntegrationEventV2Request | 
    silent := "silent_example" // string | Possible values: `yes` or `no`. - `yes`: Increases the perfomance of the API call by returning a 204 response. - `no`: Returns a 200 response that contains essential data such as the updated customer profiles and session-related information.  (optional) (default to "yes")
    dry := true // bool | Indicates whether to persist the changes. Changes are ignored when `dry=true`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.TrackEventV2(context.Background()).Body(body).Silent(silent).Dry(dry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.TrackEventV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TrackEventV2`: IntegrationStateV2
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.TrackEventV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackEventV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IntegrationEventV2Request**](IntegrationEventV2Request.md) |  | 
 **silent** | **string** | Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the perfomance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains essential data such as the updated customer profiles and session-related information.  | [default to &quot;yes&quot;]
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

