# \OutboxAPI

All URIs are relative to *https://emailengine.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1OutboxQueueid**](OutboxAPI.md#DeleteV1OutboxQueueid) | **Delete** /v1/outbox/{queueId} | Remove a message
[**GetV1Outbox**](OutboxAPI.md#GetV1Outbox) | **Get** /v1/outbox | List queued messages
[**GetV1OutboxQueueid**](OutboxAPI.md#GetV1OutboxQueueid) | **Get** /v1/outbox/{queueId} | Get queued message



## DeleteV1OutboxQueueid

> DeleteOutboxEntryResponse DeleteV1OutboxQueueid(ctx, queueId).XEeTimeout(xEeTimeout).Execute()

Remove a message



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
	queueId := "queueId_example" // string | Queue identifier for scheduled email
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutboxAPI.DeleteV1OutboxQueueid(context.Background(), queueId).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboxAPI.DeleteV1OutboxQueueid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1OutboxQueueid`: DeleteOutboxEntryResponse
	fmt.Fprintf(os.Stdout, "Response from `OutboxAPI.DeleteV1OutboxQueueid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** | Queue identifier for scheduled email | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1OutboxQueueidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

[**DeleteOutboxEntryResponse**](DeleteOutboxEntryResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1Outbox

> OutboxListResponse GetV1Outbox(ctx).XEeTimeout(xEeTimeout).Page(page).PageSize(pageSize).Execute()

List queued messages



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
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)
	page := int32(56) // int32 | Page number (zero indexed, so use 0 for first page) (optional) (default to 0)
	pageSize := int32(56) // int32 | How many entries per page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutboxAPI.GetV1Outbox(context.Background()).XEeTimeout(xEeTimeout).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboxAPI.GetV1Outbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1Outbox`: OutboxListResponse
	fmt.Fprintf(os.Stdout, "Response from `OutboxAPI.GetV1Outbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1OutboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **page** | **int32** | Page number (zero indexed, so use 0 for first page) | [default to 0]
 **pageSize** | **int32** | How many entries per page | [default to 20]

### Return type

[**OutboxListResponse**](OutboxListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1OutboxQueueid

> OutboxEntry GetV1OutboxQueueid(ctx, queueId).XEeTimeout(xEeTimeout).Execute()

Get queued message



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
	queueId := "queueId_example" // string | Queue identifier for scheduled email
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutboxAPI.GetV1OutboxQueueid(context.Background(), queueId).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboxAPI.GetV1OutboxQueueid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1OutboxQueueid`: OutboxEntry
	fmt.Fprintf(os.Stdout, "Response from `OutboxAPI.GetV1OutboxQueueid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** | Queue identifier for scheduled email | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1OutboxQueueidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

[**OutboxEntry**](OutboxEntry.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

