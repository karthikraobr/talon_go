# \ReferralsApi

All URIs are relative to *https://goflink.europe-west1.talon.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReferral**](ReferralsApi.md#CreateReferral) | **Post** /v1/referrals | Create referral code for an advocate
[**CreateReferralsForMultipleAdvocates**](ReferralsApi.md#CreateReferralsForMultipleAdvocates) | **Post** /v1/referrals_for_multiple_advocates | Create referral codes for multiple advocates



## CreateReferral

> Referral CreateReferral(ctx).Body(body).Execute()

Create referral code for an advocate



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
    body := *openapiclient.NewNewReferral(int32(78), "URNGV8294NV") // NewReferral | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReferralsApi.CreateReferral(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferralsApi.CreateReferral``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReferral`: Referral
    fmt.Fprintf(os.Stdout, "Response from `ReferralsApi.CreateReferral`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReferralRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewReferral**](NewReferral.md) |  | 

### Return type

[**Referral**](Referral.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReferralsForMultipleAdvocates

> CreateReferralsForMultipleAdvocates201Response CreateReferralsForMultipleAdvocates(ctx).Body(body).Silent(silent).Execute()

Create referral codes for multiple advocates



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
    body := *openapiclient.NewNewReferralsForMultipleAdvocates(int32(1), int32(45), []string{"AdvocateProfileIntegrationIds_example"}) // NewReferralsForMultipleAdvocates | 
    silent := "silent_example" // string | Possible values: `yes` or `no`. - `yes`: Increases the perfomance of the API call by returning a 204 response. - `no`: Returns a 200 response that contains essential data such as the updated customer profiles and session-related information.  (optional) (default to "yes")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReferralsApi.CreateReferralsForMultipleAdvocates(context.Background()).Body(body).Silent(silent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferralsApi.CreateReferralsForMultipleAdvocates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReferralsForMultipleAdvocates`: CreateReferralsForMultipleAdvocates201Response
    fmt.Fprintf(os.Stdout, "Response from `ReferralsApi.CreateReferralsForMultipleAdvocates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReferralsForMultipleAdvocatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewReferralsForMultipleAdvocates**](NewReferralsForMultipleAdvocates.md) |  | 
 **silent** | **string** | Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the perfomance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains essential data such as the updated customer profiles and session-related information.  | [default to &quot;yes&quot;]

### Return type

[**CreateReferralsForMultipleAdvocates201Response**](CreateReferralsForMultipleAdvocates201Response.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

