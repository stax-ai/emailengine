# \SMTPGatewayAPI

All URIs are relative to *https://emailengine.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1GatewayGateway**](SMTPGatewayAPI.md#DeleteV1GatewayGateway) | **Delete** /v1/gateway/{gateway} | Remove SMTP gateway
[**GetV1GatewayGateway**](SMTPGatewayAPI.md#GetV1GatewayGateway) | **Get** /v1/gateway/{gateway} | Get gateway info
[**GetV1Gateways**](SMTPGatewayAPI.md#GetV1Gateways) | **Get** /v1/gateways | List gateways
[**PostV1Gateway**](SMTPGatewayAPI.md#PostV1Gateway) | **Post** /v1/gateway | Register new gateway
[**PutV1GatewayEditGateway**](SMTPGatewayAPI.md#PutV1GatewayEditGateway) | **Put** /v1/gateway/edit/{gateway} | Update gateway info



## DeleteV1GatewayGateway

> Model13 DeleteV1GatewayGateway(ctx, gateway).XEeTimeout(xEeTimeout).Execute()

Remove SMTP gateway



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
	gateway := "gateway_example" // string | Gateway ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMTPGatewayAPI.DeleteV1GatewayGateway(context.Background(), gateway).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMTPGatewayAPI.DeleteV1GatewayGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1GatewayGateway`: Model13
	fmt.Fprintf(os.Stdout, "Response from `SMTPGatewayAPI.DeleteV1GatewayGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gateway** | **string** | Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1GatewayGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

[**Model13**](Model13.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1GatewayGateway

> GatewayResponse GetV1GatewayGateway(ctx, gateway).XEeTimeout(xEeTimeout).Execute()

Get gateway info



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
	gateway := "gateway_example" // string | Gateway ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMTPGatewayAPI.GetV1GatewayGateway(context.Background(), gateway).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMTPGatewayAPI.GetV1GatewayGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1GatewayGateway`: GatewayResponse
	fmt.Fprintf(os.Stdout, "Response from `SMTPGatewayAPI.GetV1GatewayGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gateway** | **string** | Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1GatewayGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

[**GatewayResponse**](GatewayResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1Gateways

> GatewaysFilterResponse GetV1Gateways(ctx).XEeTimeout(xEeTimeout).Page(page).PageSize(pageSize).Execute()

List gateways



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
	resp, r, err := apiClient.SMTPGatewayAPI.GetV1Gateways(context.Background()).XEeTimeout(xEeTimeout).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMTPGatewayAPI.GetV1Gateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1Gateways`: GatewaysFilterResponse
	fmt.Fprintf(os.Stdout, "Response from `SMTPGatewayAPI.GetV1Gateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1GatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **page** | **int32** | Page number (zero indexed, so use 0 for first page) | [default to 0]
 **pageSize** | **int32** | How many entries per page | [default to 20]

### Return type

[**GatewaysFilterResponse**](GatewaysFilterResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1Gateway

> CreateGatewayResponse PostV1Gateway(ctx).XEeTimeout(xEeTimeout).CreateGateway(createGateway).Execute()

Register new gateway



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
	createGateway := *openapiclient.NewCreateGateway("Gateway_example", "Host_example", "Name_example", int32(123)) // CreateGateway |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMTPGatewayAPI.PostV1Gateway(context.Background()).XEeTimeout(xEeTimeout).CreateGateway(createGateway).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMTPGatewayAPI.PostV1Gateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1Gateway`: CreateGatewayResponse
	fmt.Fprintf(os.Stdout, "Response from `SMTPGatewayAPI.PostV1Gateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1GatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **createGateway** | [**CreateGateway**](CreateGateway.md) |  | 

### Return type

[**CreateGatewayResponse**](CreateGatewayResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1GatewayEditGateway

> UpdateGatewayResponse PutV1GatewayEditGateway(ctx, gateway).XEeTimeout(xEeTimeout).UpdateGateway(updateGateway).Execute()

Update gateway info



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
	gateway := "gateway_example" // string | Gateway ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)
	updateGateway := *openapiclient.NewUpdateGateway() // UpdateGateway |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMTPGatewayAPI.PutV1GatewayEditGateway(context.Background(), gateway).XEeTimeout(xEeTimeout).UpdateGateway(updateGateway).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMTPGatewayAPI.PutV1GatewayEditGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1GatewayEditGateway`: UpdateGatewayResponse
	fmt.Fprintf(os.Stdout, "Response from `SMTPGatewayAPI.PutV1GatewayEditGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gateway** | **string** | Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1GatewayEditGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **updateGateway** | [**UpdateGateway**](UpdateGateway.md) |  | 

### Return type

[**UpdateGatewayResponse**](UpdateGatewayResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

