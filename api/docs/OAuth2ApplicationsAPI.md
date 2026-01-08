# \OAuth2ApplicationsAPI

All URIs are relative to *https://emailengine.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1Oauth2App**](OAuth2ApplicationsAPI.md#DeleteV1Oauth2App) | **Delete** /v1/oauth2/{app} | Remove OAuth2 application
[**GetV1Oauth2**](OAuth2ApplicationsAPI.md#GetV1Oauth2) | **Get** /v1/oauth2 | List OAuth2 applications
[**GetV1Oauth2App**](OAuth2ApplicationsAPI.md#GetV1Oauth2App) | **Get** /v1/oauth2/{app} | Get application info
[**PostV1Oauth2**](OAuth2ApplicationsAPI.md#PostV1Oauth2) | **Post** /v1/oauth2 | Register OAuth2 application
[**PutV1Oauth2App**](OAuth2ApplicationsAPI.md#PutV1Oauth2App) | **Put** /v1/oauth2/{app} | Update OAuth2 application



## DeleteV1Oauth2App

> DeleteAppRequestResponse DeleteV1Oauth2App(ctx, app).XEeTimeout(xEeTimeout).Execute()

Remove OAuth2 application



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
	app := "app_example" // string | OAuth2 application ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth2ApplicationsAPI.DeleteV1Oauth2App(context.Background(), app).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ApplicationsAPI.DeleteV1Oauth2App``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1Oauth2App`: DeleteAppRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth2ApplicationsAPI.DeleteV1Oauth2App`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | OAuth2 application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1Oauth2AppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

[**DeleteAppRequestResponse**](DeleteAppRequestResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1Oauth2

> OAuth2FilterResponse GetV1Oauth2(ctx).XEeTimeout(xEeTimeout).Page(page).PageSize(pageSize).Execute()

List OAuth2 applications



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
	resp, r, err := apiClient.OAuth2ApplicationsAPI.GetV1Oauth2(context.Background()).XEeTimeout(xEeTimeout).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ApplicationsAPI.GetV1Oauth2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1Oauth2`: OAuth2FilterResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth2ApplicationsAPI.GetV1Oauth2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1Oauth2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **page** | **int32** | Page number (zero indexed, so use 0 for first page) | [default to 0]
 **pageSize** | **int32** | How many entries per page | [default to 20]

### Return type

[**OAuth2FilterResponse**](OAuth2FilterResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1Oauth2App

> ApplicationResponse GetV1Oauth2App(ctx, app).XEeTimeout(xEeTimeout).Execute()

Get application info



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
	app := "app_example" // string | OAuth2 application ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth2ApplicationsAPI.GetV1Oauth2App(context.Background(), app).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ApplicationsAPI.GetV1Oauth2App``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1Oauth2App`: ApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth2ApplicationsAPI.GetV1Oauth2App`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | OAuth2 application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1Oauth2AppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

[**ApplicationResponse**](ApplicationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1Oauth2

> CreateAppResponse PostV1Oauth2(ctx).XEeTimeout(xEeTimeout).CreateOAuth2App(createOAuth2App).Execute()

Register OAuth2 application



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
	createOAuth2App := *openapiclient.NewCreateOAuth2App("Authority_example", "Name_example", openapiclient.provider("gmail"), "ServiceClient_example", "ServiceClientEmail_example", "ServiceKey_example") // CreateOAuth2App |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth2ApplicationsAPI.PostV1Oauth2(context.Background()).XEeTimeout(xEeTimeout).CreateOAuth2App(createOAuth2App).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ApplicationsAPI.PostV1Oauth2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1Oauth2`: CreateAppResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth2ApplicationsAPI.PostV1Oauth2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1Oauth2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **createOAuth2App** | [**CreateOAuth2App**](CreateOAuth2App.md) |  | 

### Return type

[**CreateAppResponse**](CreateAppResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1Oauth2App

> UpdateOAuthAppResponse PutV1Oauth2App(ctx, app).XEeTimeout(xEeTimeout).UpdateOAuthApp(updateOAuthApp).Execute()

Update OAuth2 application



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
	app := "app_example" // string | OAuth2 application ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)
	updateOAuthApp := *openapiclient.NewUpdateOAuthApp() // UpdateOAuthApp |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth2ApplicationsAPI.PutV1Oauth2App(context.Background(), app).XEeTimeout(xEeTimeout).UpdateOAuthApp(updateOAuthApp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ApplicationsAPI.PutV1Oauth2App``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1Oauth2App`: UpdateOAuthAppResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth2ApplicationsAPI.PutV1Oauth2App`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | OAuth2 application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1Oauth2AppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **updateOAuthApp** | [**UpdateOAuthApp**](UpdateOAuthApp.md) |  | 

### Return type

[**UpdateOAuthAppResponse**](UpdateOAuthAppResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

