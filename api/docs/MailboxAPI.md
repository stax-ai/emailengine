# \MailboxAPI

All URIs are relative to *https://emailengine.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1AccountAccountMailbox**](MailboxAPI.md#DeleteV1AccountAccountMailbox) | **Delete** /v1/account/{account}/mailbox | Delete mailbox
[**GetV1AccountAccountMailboxes**](MailboxAPI.md#GetV1AccountAccountMailboxes) | **Get** /v1/account/{account}/mailboxes | List mailboxes
[**PostV1AccountAccountMailbox**](MailboxAPI.md#PostV1AccountAccountMailbox) | **Post** /v1/account/{account}/mailbox | Create mailbox
[**PutV1AccountAccountMailbox**](MailboxAPI.md#PutV1AccountAccountMailbox) | **Put** /v1/account/{account}/mailbox | Modify mailbox



## DeleteV1AccountAccountMailbox

> DeleteMailboxResponse DeleteV1AccountAccountMailbox(ctx, account).Path(path).XEeTimeout(xEeTimeout).Execute()

Delete mailbox



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
	path := "path_example" // string | Mailbox folder path to delete
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MailboxAPI.DeleteV1AccountAccountMailbox(context.Background(), account).Path(path).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxAPI.DeleteV1AccountAccountMailbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1AccountAccountMailbox`: DeleteMailboxResponse
	fmt.Fprintf(os.Stdout, "Response from `MailboxAPI.DeleteV1AccountAccountMailbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1AccountAccountMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | Mailbox folder path to delete | 
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

[**DeleteMailboxResponse**](DeleteMailboxResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1AccountAccountMailboxes

> MailboxesFilterResponse GetV1AccountAccountMailboxes(ctx, account).XEeTimeout(xEeTimeout).Counters(counters).Execute()

List mailboxes



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
	counters := true // bool | If true, then includes message counters in the response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MailboxAPI.GetV1AccountAccountMailboxes(context.Background(), account).XEeTimeout(xEeTimeout).Counters(counters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxAPI.GetV1AccountAccountMailboxes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1AccountAccountMailboxes`: MailboxesFilterResponse
	fmt.Fprintf(os.Stdout, "Response from `MailboxAPI.GetV1AccountAccountMailboxes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AccountAccountMailboxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **counters** | **bool** | If true, then includes message counters in the response | [default to false]

### Return type

[**MailboxesFilterResponse**](MailboxesFilterResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1AccountAccountMailbox

> CreateMailboxResponse PostV1AccountAccountMailbox(ctx, account).XEeTimeout(xEeTimeout).CreateMailbox(createMailbox).Execute()

Create mailbox



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
	createMailbox := *openapiclient.NewCreateMailbox() // CreateMailbox |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MailboxAPI.PostV1AccountAccountMailbox(context.Background(), account).XEeTimeout(xEeTimeout).CreateMailbox(createMailbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxAPI.PostV1AccountAccountMailbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1AccountAccountMailbox`: CreateMailboxResponse
	fmt.Fprintf(os.Stdout, "Response from `MailboxAPI.PostV1AccountAccountMailbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1AccountAccountMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **createMailbox** | [**CreateMailbox**](CreateMailbox.md) |  | 

### Return type

[**CreateMailboxResponse**](CreateMailboxResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1AccountAccountMailbox

> ModifyMailboxResponse PutV1AccountAccountMailbox(ctx, account).XEeTimeout(xEeTimeout).ModifyMailbox(modifyMailbox).Execute()

Modify mailbox



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
	modifyMailbox := *openapiclient.NewModifyMailbox("Path_example") // ModifyMailbox |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MailboxAPI.PutV1AccountAccountMailbox(context.Background(), account).XEeTimeout(xEeTimeout).ModifyMailbox(modifyMailbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailboxAPI.PutV1AccountAccountMailbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1AccountAccountMailbox`: ModifyMailboxResponse
	fmt.Fprintf(os.Stdout, "Response from `MailboxAPI.PutV1AccountAccountMailbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1AccountAccountMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **modifyMailbox** | [**ModifyMailbox**](ModifyMailbox.md) |  | 

### Return type

[**ModifyMailboxResponse**](ModifyMailboxResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

