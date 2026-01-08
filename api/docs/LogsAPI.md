# \LogsAPI

All URIs are relative to *https://emailengine.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1LogsAccount**](LogsAPI.md#GetV1LogsAccount) | **Get** /v1/logs/{account} | Return IMAP logs for an account



## GetV1LogsAccount

> string GetV1LogsAccount(ctx, account).XEeTimeout(xEeTimeout).Execute()

Return IMAP logs for an account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | Unique identifier for the email account
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.GetV1LogsAccount(context.Background(), account).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetV1LogsAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1LogsAccount`: string
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.GetV1LogsAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1LogsAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

