# \CatalogsApi

All URIs are relative to *https://goflink.europe-west1.talon.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyncCatalog**](CatalogsApi.md#SyncCatalog) | **Put** /v1/catalogs/{catalogId}/sync | Sync cart item catalog



## SyncCatalog

> Catalog SyncCatalog(ctx, catalogId).Body(body).Execute()

Sync cart item catalog



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
    catalogId := int32(56) // int32 | 
    body := *openapiclient.NewCatalogSyncRequest([]openapiclient.CatalogAction{*openapiclient.NewCatalogAction("ADD", map[string]interface{}(123))}) // CatalogSyncRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogsApi.SyncCatalog(context.Background(), catalogId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogsApi.SyncCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncCatalog`: Catalog
    fmt.Fprintf(os.Stdout, "Response from `CatalogsApi.SyncCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CatalogSyncRequest**](CatalogSyncRequest.md) |  | 

### Return type

[**Catalog**](Catalog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

