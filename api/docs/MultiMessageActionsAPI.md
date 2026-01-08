# \MultiMessageActionsAPI

All URIs are relative to *https://emailengine.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1AccountAccountMessages**](MultiMessageActionsAPI.md#PutV1AccountAccountMessages) | **Put** /v1/account/{account}/messages | Update messages
[**PutV1AccountAccountMessagesDelete**](MultiMessageActionsAPI.md#PutV1AccountAccountMessagesDelete) | **Put** /v1/account/{account}/messages/delete | Delete messages
[**PutV1AccountAccountMessagesMove**](MultiMessageActionsAPI.md#PutV1AccountAccountMessagesMove) | **Put** /v1/account/{account}/messages/move | Move messages



## PutV1AccountAccountMessages

> MessageUpdateResponse PutV1AccountAccountMessages(ctx, account).Path(path).XEeTimeout(xEeTimeout).MessagesUpdateRequest(messagesUpdateRequest).Execute()

Update messages



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
	path := "path_example" // string | Mailbox folder path. Can use special use labels like \"\\Sent\". Special value \"\\All\" is available for Gmail IMAP, Gmail API, MS Graph API accounts.
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)
	messagesUpdateRequest := *openapiclient.NewMessagesUpdateRequest(*openapiclient.NewSearchQuery()) // MessagesUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiMessageActionsAPI.PutV1AccountAccountMessages(context.Background(), account).Path(path).XEeTimeout(xEeTimeout).MessagesUpdateRequest(messagesUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiMessageActionsAPI.PutV1AccountAccountMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1AccountAccountMessages`: MessageUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiMessageActionsAPI.PutV1AccountAccountMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1AccountAccountMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | Mailbox folder path. Can use special use labels like \&quot;\\Sent\&quot;. Special value \&quot;\\All\&quot; is available for Gmail IMAP, Gmail API, MS Graph API accounts. | 
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **messagesUpdateRequest** | [**MessagesUpdateRequest**](MessagesUpdateRequest.md) |  | 

### Return type

[**MessageUpdateResponse**](MessageUpdateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1AccountAccountMessagesDelete

> MessagesDeleteResponse PutV1AccountAccountMessagesDelete(ctx, account).Path(path).XEeTimeout(xEeTimeout).Force(force).MessagesDeleteRequest(messagesDeleteRequest).Execute()

Delete messages



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
	path := "path_example" // string | Mailbox folder path. Can use special use labels like \"\\Sent\". Special value \"\\All\" is available for Gmail IMAP, Gmail API, MS Graph API accounts.
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)
	force := true // bool | Delete messages even if not in Trash (optional) (default to false)
	messagesDeleteRequest := *openapiclient.NewMessagesDeleteRequest(*openapiclient.NewSearchQuery()) // MessagesDeleteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiMessageActionsAPI.PutV1AccountAccountMessagesDelete(context.Background(), account).Path(path).XEeTimeout(xEeTimeout).Force(force).MessagesDeleteRequest(messagesDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiMessageActionsAPI.PutV1AccountAccountMessagesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1AccountAccountMessagesDelete`: MessagesDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiMessageActionsAPI.PutV1AccountAccountMessagesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1AccountAccountMessagesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | Mailbox folder path. Can use special use labels like \&quot;\\Sent\&quot;. Special value \&quot;\\All\&quot; is available for Gmail IMAP, Gmail API, MS Graph API accounts. | 
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **force** | **bool** | Delete messages even if not in Trash | [default to false]
 **messagesDeleteRequest** | [**MessagesDeleteRequest**](MessagesDeleteRequest.md) |  | 

### Return type

[**MessagesDeleteResponse**](MessagesDeleteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1AccountAccountMessagesMove

> MessagesMoveResponse PutV1AccountAccountMessagesMove(ctx, account).Path(path).XEeTimeout(xEeTimeout).MessagesMoveRequest(messagesMoveRequest).Execute()

Move messages



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
	path := "path_example" // string | Mailbox folder path. Can use special use labels like \"\\Sent\". Special value \"\\All\" is available for Gmail IMAP, Gmail API, MS Graph API accounts.
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)
	messagesMoveRequest := *openapiclient.NewMessagesMoveRequest("Path_example", *openapiclient.NewSearchQuery()) // MessagesMoveRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiMessageActionsAPI.PutV1AccountAccountMessagesMove(context.Background(), account).Path(path).XEeTimeout(xEeTimeout).MessagesMoveRequest(messagesMoveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiMessageActionsAPI.PutV1AccountAccountMessagesMove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1AccountAccountMessagesMove`: MessagesMoveResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiMessageActionsAPI.PutV1AccountAccountMessagesMove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1AccountAccountMessagesMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | Mailbox folder path. Can use special use labels like \&quot;\\Sent\&quot;. Special value \&quot;\\All\&quot; is available for Gmail IMAP, Gmail API, MS Graph API accounts. | 
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **messagesMoveRequest** | [**MessagesMoveRequest**](MessagesMoveRequest.md) |  | 

### Return type

[**MessagesMoveResponse**](MessagesMoveResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

