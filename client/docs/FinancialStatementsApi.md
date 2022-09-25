# \FinancialStatementsApi

All URIs are relative to *https://financialmodelingprep.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIncomeStatements**](FinancialStatementsApi.md#GetIncomeStatements) | **Get** /income-statement/{symbol} | Get Income Statements



## GetIncomeStatements

> IncomeStatementListResponse GetIncomeStatements(ctx, symbol).Limit(limit).Period(period).Apikey(apikey).Datatype(datatype).Execute()

Get Income Statements



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
    symbol := "symbol_example" // string | 
    limit := int32(56) // int32 | Number of results to return. (optional)
    period := "period_example" // string | Specify a quarterly income statement. Default is annual. (optional)
    apikey := "apikey_example" // string | API key for authn / authz (optional)
    datatype := "datatype_example" // string | Format of data to return (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FinancialStatementsApi.GetIncomeStatements(context.Background(), symbol).Limit(limit).Period(period).Apikey(apikey).Datatype(datatype).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FinancialStatementsApi.GetIncomeStatements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncomeStatements`: IncomeStatementListResponse
    fmt.Fprintf(os.Stdout, "Response from `FinancialStatementsApi.GetIncomeStatements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncomeStatementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return. | 
 **period** | **string** | Specify a quarterly income statement. Default is annual. | 
 **apikey** | **string** | API key for authn / authz | 
 **datatype** | **string** | Format of data to return | 

### Return type

[**IncomeStatementListResponse**](IncomeStatementListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

