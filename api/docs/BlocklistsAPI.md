# \BlocklistsAPI

All URIs are relative to *https://emailengine.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1BlocklistListid**](BlocklistsAPI.md#DeleteV1BlocklistListid) | **Delete** /v1/blocklist/{listId} | Remove from blocklist
[**GetV1BlocklistListid**](BlocklistsAPI.md#GetV1BlocklistListid) | **Get** /v1/blocklist/{listId} | List blocklist entries
[**GetV1Blocklists**](BlocklistsAPI.md#GetV1Blocklists) | **Get** /v1/blocklists | List blocklists
[**PostV1BlocklistListid**](BlocklistsAPI.md#PostV1BlocklistListid) | **Post** /v1/blocklist/{listId} | Add to blocklist



## DeleteV1BlocklistListid

> DeleteBlocklistResponse DeleteV1BlocklistListid(ctx, listId).Recipient(recipient).XEeTimeout(xEeTimeout).Execute()

Remove from blocklist



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
	listId := "listId_example" // string | List ID. Must use a subdomain name format. Lists are registered ad-hoc, so a new identifier defines a new list.
	recipient := "recipient_example" // string | Email address to remove from the list
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocklistsAPI.DeleteV1BlocklistListid(context.Background(), listId).Recipient(recipient).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocklistsAPI.DeleteV1BlocklistListid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1BlocklistListid`: DeleteBlocklistResponse
	fmt.Fprintf(os.Stdout, "Response from `BlocklistsAPI.DeleteV1BlocklistListid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | List ID. Must use a subdomain name format. Lists are registered ad-hoc, so a new identifier defines a new list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1BlocklistListidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recipient** | **string** | Email address to remove from the list | 
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

[**DeleteBlocklistResponse**](DeleteBlocklistResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1BlocklistListid

> BlocklistListResponse GetV1BlocklistListid(ctx, listId).XEeTimeout(xEeTimeout).Page(page).PageSize(pageSize).Execute()

List blocklist entries



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
	listId := "listId_example" // string | List ID. Must use a subdomain name format. Lists are registered ad-hoc, so a new identifier defines a new list.
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)
	page := int32(56) // int32 | Page number (zero indexed, so use 0 for first page) (optional) (default to 0)
	pageSize := int32(56) // int32 | How many entries per page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocklistsAPI.GetV1BlocklistListid(context.Background(), listId).XEeTimeout(xEeTimeout).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocklistsAPI.GetV1BlocklistListid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1BlocklistListid`: BlocklistListResponse
	fmt.Fprintf(os.Stdout, "Response from `BlocklistsAPI.GetV1BlocklistListid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | List ID. Must use a subdomain name format. Lists are registered ad-hoc, so a new identifier defines a new list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1BlocklistListidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **page** | **int32** | Page number (zero indexed, so use 0 for first page) | [default to 0]
 **pageSize** | **int32** | How many entries per page | [default to 20]

### Return type

[**BlocklistListResponse**](BlocklistListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1Blocklists

> BlocklistsResponse GetV1Blocklists(ctx).XEeTimeout(xEeTimeout).Page(page).PageSize(pageSize).Execute()

List blocklists



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
	resp, r, err := apiClient.BlocklistsAPI.GetV1Blocklists(context.Background()).XEeTimeout(xEeTimeout).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocklistsAPI.GetV1Blocklists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1Blocklists`: BlocklistsResponse
	fmt.Fprintf(os.Stdout, "Response from `BlocklistsAPI.GetV1Blocklists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1BlocklistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **page** | **int32** | Page number (zero indexed, so use 0 for first page) | [default to 0]
 **pageSize** | **int32** | How many entries per page | [default to 20]

### Return type

[**BlocklistsResponse**](BlocklistsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1BlocklistListid

> BlocklistListAddResponse PostV1BlocklistListid(ctx, listId).XEeTimeout(xEeTimeout).BlocklistListAddPayload(blocklistListAddPayload).Execute()

Add to blocklist



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
	listId := "listId_example" // string | List ID. Must use a subdomain name format. Lists are registered ad-hoc, so a new identifier defines a new list.
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)
	blocklistListAddPayload := *openapiclient.NewBlocklistListAddPayload("Account_example", "Recipient_example") // BlocklistListAddPayload |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocklistsAPI.PostV1BlocklistListid(context.Background(), listId).XEeTimeout(xEeTimeout).BlocklistListAddPayload(blocklistListAddPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocklistsAPI.PostV1BlocklistListid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1BlocklistListid`: BlocklistListAddResponse
	fmt.Fprintf(os.Stdout, "Response from `BlocklistsAPI.PostV1BlocklistListid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | List ID. Must use a subdomain name format. Lists are registered ad-hoc, so a new identifier defines a new list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1BlocklistListidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **blocklistListAddPayload** | [**BlocklistListAddPayload**](BlocklistListAddPayload.md) |  | 

### Return type

[**BlocklistListAddResponse**](BlocklistListAddResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

