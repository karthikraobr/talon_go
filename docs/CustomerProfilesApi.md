# \CustomerProfilesApi

All URIs are relative to *https://goflink.europe-west1.talon.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCustomerData**](CustomerProfilesApi.md#DeleteCustomerData) | **Delete** /v1/customer_data/{integrationId} | Delete customer&#39;s personal data
[**GetCustomerInventory**](CustomerProfilesApi.md#GetCustomerInventory) | **Get** /v1/customer_profiles/{integrationId}/inventory | List customer data
[**UpdateCustomerProfileV2**](CustomerProfilesApi.md#UpdateCustomerProfileV2) | **Put** /v2/customer_profiles/{integrationId} | Update customer profile
[**UpdateCustomerProfilesV2**](CustomerProfilesApi.md#UpdateCustomerProfilesV2) | **Put** /v2/customer_profiles | Update multiple customer profiles



## DeleteCustomerData

> DeleteCustomerData(ctx, integrationId).Execute()

Delete customer's personal data



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
    integrationId := "integrationId_example" // string | The integration ID of the customer profile. You can get the `integrationId` of a profile using: - A customer session integration Id with the [Update customer session endpoint](https://docs.talon.one/integration-api/#operation/updateCustomerSessionV2). - The Management API with the [List application's customers endpoint](https://docs.talon.one/management-api/#operation/getApplicationCustomers). 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerProfilesApi.DeleteCustomerData(context.Background(), integrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerProfilesApi.DeleteCustomerData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The integration ID of the customer profile. You can get the &#x60;integrationId&#x60; of a profile using: - A customer session integration Id with the [Update customer session endpoint](https://docs.talon.one/integration-api/#operation/updateCustomerSessionV2). - The Management API with the [List application&#39;s customers endpoint](https://docs.talon.one/management-api/#operation/getApplicationCustomers).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerInventory

> CustomerInventory GetCustomerInventory(ctx, integrationId).Profile(profile).Referrals(referrals).Coupons(coupons).Loyalty(loyalty).Giveaways(giveaways).LoyaltyProjectionEndDate(loyaltyProjectionEndDate).Execute()

List customer data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    integrationId := "integrationId_example" // string | The integration ID of the customer profile. You can get the `integrationId` of a profile using: - A customer session integration Id with the [Update customer session endpoint](https://docs.talon.one/integration-api/#operation/updateCustomerSessionV2). - The Management API with the [List application's customers endpoint](https://docs.talon.one/management-api/#operation/getApplicationCustomers). 
    profile := true // bool | Set to `true` to include customer profile information in the response. (optional)
    referrals := true // bool | Set to `true` to include referral information in the response. (optional)
    coupons := true // bool | Set to `true` to include coupon information in the response. (optional)
    loyalty := true // bool | Set to `true` to include loyalty information in the response. (optional)
    giveaways := true // bool | Set to `true` to include giveaways information in the response. (optional)
    loyaltyProjectionEndDate := time.Now() // time.Time | Set an end date to query the projected loyalty balances. You can project results up to 31 days from today. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerProfilesApi.GetCustomerInventory(context.Background(), integrationId).Profile(profile).Referrals(referrals).Coupons(coupons).Loyalty(loyalty).Giveaways(giveaways).LoyaltyProjectionEndDate(loyaltyProjectionEndDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerProfilesApi.GetCustomerInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerInventory`: CustomerInventory
    fmt.Fprintf(os.Stdout, "Response from `CustomerProfilesApi.GetCustomerInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The integration ID of the customer profile. You can get the &#x60;integrationId&#x60; of a profile using: - A customer session integration Id with the [Update customer session endpoint](https://docs.talon.one/integration-api/#operation/updateCustomerSessionV2). - The Management API with the [List application&#39;s customers endpoint](https://docs.talon.one/management-api/#operation/getApplicationCustomers).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profile** | **bool** | Set to &#x60;true&#x60; to include customer profile information in the response. | 
 **referrals** | **bool** | Set to &#x60;true&#x60; to include referral information in the response. | 
 **coupons** | **bool** | Set to &#x60;true&#x60; to include coupon information in the response. | 
 **loyalty** | **bool** | Set to &#x60;true&#x60; to include loyalty information in the response. | 
 **giveaways** | **bool** | Set to &#x60;true&#x60; to include giveaways information in the response. | 
 **loyaltyProjectionEndDate** | **time.Time** | Set an end date to query the projected loyalty balances. You can project results up to 31 days from today. | 

### Return type

[**CustomerInventory**](CustomerInventory.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerProfileV2

> IntegrationStateV2 UpdateCustomerProfileV2(ctx, integrationId).Body(body).RunRuleEngine(runRuleEngine).Dry(dry).Execute()

Update customer profile



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
    integrationId := "integrationId_example" // string | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier. 
    body := *openapiclient.NewCustomerProfileIntegrationRequestV2() // CustomerProfileIntegrationRequestV2 | 
    runRuleEngine := true // bool | Indicates whether to run the Rule Engine.  If `true`, the response includes: - The effects generated by the triggered campaigns are returned in the `effects` property. - The created coupons and referral objects.  If `false`: - The rules are not executed and the `effects` property is always empty. - The response time improves. - You cannot use `responseContent` in the body.  (optional) (default to false)
    dry := true // bool | Indicates whether to persist the changes. Changes are ignored when `dry=true`.  This property only works when `runRuleEngine=true`.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerProfilesApi.UpdateCustomerProfileV2(context.Background(), integrationId).Body(body).RunRuleEngine(runRuleEngine).Dry(dry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerProfilesApi.UpdateCustomerProfileV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomerProfileV2`: IntegrationStateV2
    fmt.Fprintf(os.Stdout, "Response from `CustomerProfilesApi.UpdateCustomerProfileV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerProfileV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CustomerProfileIntegrationRequestV2**](CustomerProfileIntegrationRequestV2.md) |  | 
 **runRuleEngine** | **bool** | Indicates whether to run the Rule Engine.  If &#x60;true&#x60;, the response includes: - The effects generated by the triggered campaigns are returned in the &#x60;effects&#x60; property. - The created coupons and referral objects.  If &#x60;false&#x60;: - The rules are not executed and the &#x60;effects&#x60; property is always empty. - The response time improves. - You cannot use &#x60;responseContent&#x60; in the body.  | [default to false]
 **dry** | **bool** | Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;.  This property only works when &#x60;runRuleEngine&#x3D;true&#x60;.  | 

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


## UpdateCustomerProfilesV2

> MultipleCustomerProfileIntegrationResponseV2 UpdateCustomerProfilesV2(ctx).Body(body).Silent(silent).Execute()

Update multiple customer profiles



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
    body := *openapiclient.NewMultipleCustomerProfileIntegrationRequest() // MultipleCustomerProfileIntegrationRequest | 
    silent := "silent_example" // string | Possible values: `yes` or `no`. - `yes`: Increases the perfomance of the API call by returning a 204 response. - `no`: Returns a 200 response that contains essential data such as the updated customer profiles and session-related information.  (optional) (default to "yes")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerProfilesApi.UpdateCustomerProfilesV2(context.Background()).Body(body).Silent(silent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerProfilesApi.UpdateCustomerProfilesV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomerProfilesV2`: MultipleCustomerProfileIntegrationResponseV2
    fmt.Fprintf(os.Stdout, "Response from `CustomerProfilesApi.UpdateCustomerProfilesV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerProfilesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MultipleCustomerProfileIntegrationRequest**](MultipleCustomerProfileIntegrationRequest.md) |  | 
 **silent** | **string** | Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the perfomance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains essential data such as the updated customer profiles and session-related information.  | [default to &quot;yes&quot;]

### Return type

[**MultipleCustomerProfileIntegrationResponseV2**](MultipleCustomerProfileIntegrationResponseV2.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

