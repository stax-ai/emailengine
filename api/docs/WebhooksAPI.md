# \WebhooksAPI

All URIs are relative to *https://emailengine.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1Webhookroutes**](WebhooksAPI.md#GetV1Webhookroutes) | **Get** /v1/webhookRoutes | List webhook routes
[**GetV1WebhookroutesWebhookrouteWebhookroute**](WebhooksAPI.md#GetV1WebhookroutesWebhookrouteWebhookroute) | **Get** /v1/webhookRoutes/webhookRoute/{webhookRoute} | Get webhook route information



## GetV1Webhookroutes

> WebhookRoutesListResponse GetV1Webhookroutes(ctx).XEeTimeout(xEeTimeout).Page(page).PageSize(pageSize).Execute()

List webhook routes



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
	resp, r, err := apiClient.WebhooksAPI.GetV1Webhookroutes(context.Background()).XEeTimeout(xEeTimeout).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetV1Webhookroutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1Webhookroutes`: WebhookRoutesListResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetV1Webhookroutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1WebhookroutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **page** | **int32** | Page number (zero indexed, so use 0 for first page) | [default to 0]
 **pageSize** | **int32** | How many entries per page | [default to 20]

### Return type

[**WebhookRoutesListResponse**](WebhookRoutesListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1WebhookroutesWebhookrouteWebhookroute

> WebhookRouteResponse GetV1WebhookroutesWebhookrouteWebhookroute(ctx, webhookRoute).XEeTimeout(xEeTimeout).Execute()

Get webhook route information



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
	webhookRoute := "webhookRoute_example" // string | Webhook ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.GetV1WebhookroutesWebhookrouteWebhookroute(context.Background(), webhookRoute).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetV1WebhookroutesWebhookrouteWebhookroute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1WebhookroutesWebhookrouteWebhookroute`: WebhookRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetV1WebhookroutesWebhookrouteWebhookroute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookRoute** | **string** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1WebhookroutesWebhookrouteWebhookrouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

[**WebhookRouteResponse**](WebhookRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

