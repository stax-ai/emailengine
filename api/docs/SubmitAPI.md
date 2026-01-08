# \SubmitAPI

All URIs are relative to *https://emailengine.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1AccountAccountSubmit**](SubmitAPI.md#PostV1AccountAccountSubmit) | **Post** /v1/account/{account}/submit | Submit message for delivery



## PostV1AccountAccountSubmit

> SubmitMessageResponse PostV1AccountAccountSubmit(ctx, account).XEeTimeout(xEeTimeout).IdempotencyKey(idempotencyKey).DocumentStore(documentStore).UseStructuredFormat(useStructuredFormat).SubmitMessage(submitMessage).Execute()

Submit message for delivery



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
	idempotencyKey := "idempotencyKey_example" // string | Unique key to prevent duplicate processing of the same request (optional)
	documentStore := true // bool | If enabled then fetch email used as a reference template from the Document Store (optional) (default to false)
	useStructuredFormat := true // bool | For MS Graph accounts: If true, uses structured JSON format (respects from field for shared mailboxes, breaks calendar invites and special MIME types). If false, sends as raw MIME (preserves calendar invites, ignores from field). Default is false (raw MIME). (optional) (default to false)
	submitMessage := *openapiclient.NewSubmitMessage() // SubmitMessage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubmitAPI.PostV1AccountAccountSubmit(context.Background(), account).XEeTimeout(xEeTimeout).IdempotencyKey(idempotencyKey).DocumentStore(documentStore).UseStructuredFormat(useStructuredFormat).SubmitMessage(submitMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubmitAPI.PostV1AccountAccountSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1AccountAccountSubmit`: SubmitMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `SubmitAPI.PostV1AccountAccountSubmit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1AccountAccountSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **idempotencyKey** | **string** | Unique key to prevent duplicate processing of the same request | 
 **documentStore** | **bool** | If enabled then fetch email used as a reference template from the Document Store | [default to false]
 **useStructuredFormat** | **bool** | For MS Graph accounts: If true, uses structured JSON format (respects from field for shared mailboxes, breaks calendar invites and special MIME types). If false, sends as raw MIME (preserves calendar invites, ignores from field). Default is false (raw MIME). | [default to false]
 **submitMessage** | [**SubmitMessage**](SubmitMessage.md) |  | 

### Return type

[**SubmitMessageResponse**](SubmitMessageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

