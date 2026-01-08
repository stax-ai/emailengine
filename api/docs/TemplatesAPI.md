# \TemplatesAPI

All URIs are relative to *https://emailengine.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1TemplatesAccountAccount**](TemplatesAPI.md#DeleteV1TemplatesAccountAccount) | **Delete** /v1/templates/account/{account} | Flush templates
[**DeleteV1TemplatesTemplateTemplate**](TemplatesAPI.md#DeleteV1TemplatesTemplateTemplate) | **Delete** /v1/templates/template/{template} | Remove a template
[**GetV1Templates**](TemplatesAPI.md#GetV1Templates) | **Get** /v1/templates | List account templates
[**GetV1TemplatesTemplateTemplate**](TemplatesAPI.md#GetV1TemplatesTemplateTemplate) | **Get** /v1/templates/template/{template} | Get template information
[**PostV1TemplatesTemplate**](TemplatesAPI.md#PostV1TemplatesTemplate) | **Post** /v1/templates/template | Create template
[**PutV1TemplatesTemplateTemplate**](TemplatesAPI.md#PutV1TemplatesTemplateTemplate) | **Put** /v1/templates/template/{template} | Update a template



## DeleteV1TemplatesAccountAccount

> DeleteTemplateRequestResponse DeleteV1TemplatesAccountAccount(ctx, account).XEeTimeout(xEeTimeout).Force(force).Execute()

Flush templates



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
	force := true // bool | Must be true in order to flush templates (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.DeleteV1TemplatesAccountAccount(context.Background(), account).XEeTimeout(xEeTimeout).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.DeleteV1TemplatesAccountAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1TemplatesAccountAccount`: DeleteTemplateRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.DeleteV1TemplatesAccountAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1TemplatesAccountAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **force** | **bool** | Must be true in order to flush templates | [default to false]

### Return type

[**DeleteTemplateRequestResponse**](DeleteTemplateRequestResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteV1TemplatesTemplateTemplate

> Model14 DeleteV1TemplatesTemplateTemplate(ctx, template).XEeTimeout(xEeTimeout).Execute()

Remove a template



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
	template := "template_example" // string | Template ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.DeleteV1TemplatesTemplateTemplate(context.Background(), template).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.DeleteV1TemplatesTemplateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1TemplatesTemplateTemplate`: Model14
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.DeleteV1TemplatesTemplateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**template** | **string** | Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1TemplatesTemplateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

[**Model14**](Model14.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1Templates

> AccountTemplatesResponse GetV1Templates(ctx).XEeTimeout(xEeTimeout).Account(account).Page(page).PageSize(pageSize).Execute()

List account templates



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
	account := "account_example" // string | Account ID to list the templates for (optional)
	page := int32(56) // int32 | Page number (zero indexed, so use 0 for first page) (optional) (default to 0)
	pageSize := int32(56) // int32 | How many entries per page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.GetV1Templates(context.Background()).XEeTimeout(xEeTimeout).Account(account).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetV1Templates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1Templates`: AccountTemplatesResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetV1Templates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1TemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **account** | **string** | Account ID to list the templates for | 
 **page** | **int32** | Page number (zero indexed, so use 0 for first page) | [default to 0]
 **pageSize** | **int32** | How many entries per page | [default to 20]

### Return type

[**AccountTemplatesResponse**](AccountTemplatesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1TemplatesTemplateTemplate

> AccountTemplateResponse GetV1TemplatesTemplateTemplate(ctx, template).XEeTimeout(xEeTimeout).Execute()

Get template information



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
	template := "template_example" // string | Template ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.GetV1TemplatesTemplateTemplate(context.Background(), template).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetV1TemplatesTemplateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1TemplatesTemplateTemplate`: AccountTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetV1TemplatesTemplateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**template** | **string** | Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1TemplatesTemplateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

[**AccountTemplateResponse**](AccountTemplateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1TemplatesTemplate

> CreateTemplateResponse PostV1TemplatesTemplate(ctx).XEeTimeout(xEeTimeout).CreateTemplate(createTemplate).Execute()

Create template



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
	createTemplate := *openapiclient.NewCreateTemplate("Account_example", *openapiclient.NewCreateTemplateContent(), "Name_example") // CreateTemplate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.PostV1TemplatesTemplate(context.Background()).XEeTimeout(xEeTimeout).CreateTemplate(createTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.PostV1TemplatesTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TemplatesTemplate`: CreateTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.PostV1TemplatesTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TemplatesTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **createTemplate** | [**CreateTemplate**](CreateTemplate.md) |  | 

### Return type

[**CreateTemplateResponse**](CreateTemplateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1TemplatesTemplateTemplate

> UpdateTemplateResponse PutV1TemplatesTemplateTemplate(ctx, template).XEeTimeout(xEeTimeout).UpdateTemplate(updateTemplate).Execute()

Update a template



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
	template := "template_example" // string | Template ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)
	updateTemplate := *openapiclient.NewUpdateTemplate() // UpdateTemplate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.PutV1TemplatesTemplateTemplate(context.Background(), template).XEeTimeout(xEeTimeout).UpdateTemplate(updateTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.PutV1TemplatesTemplateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1TemplatesTemplateTemplate`: UpdateTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.PutV1TemplatesTemplateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**template** | **string** | Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1TemplatesTemplateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **updateTemplate** | [**UpdateTemplate**](UpdateTemplate.md) |  | 

### Return type

[**UpdateTemplateResponse**](UpdateTemplateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

