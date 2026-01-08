# \DeliveryTestAPI

All URIs are relative to *https://emailengine.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1DeliverytestCheckDeliverytest**](DeliveryTestAPI.md#GetV1DeliverytestCheckDeliverytest) | **Get** /v1/delivery-test/check/{deliveryTest} | Check test status
[**PostV1DeliverytestAccountAccount**](DeliveryTestAPI.md#PostV1DeliverytestAccountAccount) | **Post** /v1/delivery-test/account/{account} | Create delivery test



## GetV1DeliverytestCheckDeliverytest

> DeliveryCheckResponse GetV1DeliverytestCheckDeliverytest(ctx, deliveryTest).XEeTimeout(xEeTimeout).Execute()

Check test status



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
	deliveryTest := "deliveryTest_example" // string | Test ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeliveryTestAPI.GetV1DeliverytestCheckDeliverytest(context.Background(), deliveryTest).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliveryTestAPI.GetV1DeliverytestCheckDeliverytest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1DeliverytestCheckDeliverytest`: DeliveryCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `DeliveryTestAPI.GetV1DeliverytestCheckDeliverytest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryTest** | **string** | Test ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1DeliverytestCheckDeliverytestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

[**DeliveryCheckResponse**](DeliveryCheckResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1DeliverytestAccountAccount

> DeliveryStartResponse PostV1DeliverytestAccountAccount(ctx, account).XEeTimeout(xEeTimeout).DeliveryStartRequest(deliveryStartRequest).Execute()

Create delivery test



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
	deliveryStartRequest := *openapiclient.NewDeliveryStartRequest() // DeliveryStartRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeliveryTestAPI.PostV1DeliverytestAccountAccount(context.Background(), account).XEeTimeout(xEeTimeout).DeliveryStartRequest(deliveryStartRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliveryTestAPI.PostV1DeliverytestAccountAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1DeliverytestAccountAccount`: DeliveryStartResponse
	fmt.Fprintf(os.Stdout, "Response from `DeliveryTestAPI.PostV1DeliverytestAccountAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1DeliverytestAccountAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **deliveryStartRequest** | [**DeliveryStartRequest**](DeliveryStartRequest.md) |  | 

### Return type

[**DeliveryStartResponse**](DeliveryStartResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

