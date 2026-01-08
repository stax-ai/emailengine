# \AccountAPI

All URIs are relative to *https://emailengine.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1AccountAccount**](AccountAPI.md#DeleteV1AccountAccount) | **Delete** /v1/account/{account} | Remove account
[**GetV1AccountAccount**](AccountAPI.md#GetV1AccountAccount) | **Get** /v1/account/{account} | Get account info
[**GetV1AccountAccountOauthtoken**](AccountAPI.md#GetV1AccountAccountOauthtoken) | **Get** /v1/account/{account}/oauth-token | Get OAuth2 access token
[**GetV1AccountAccountServersignatures**](AccountAPI.md#GetV1AccountAccountServersignatures) | **Get** /v1/account/{account}/server-signatures | List Account Signatures
[**GetV1Accounts**](AccountAPI.md#GetV1Accounts) | **Get** /v1/accounts | List accounts
[**GetV1Changes**](AccountAPI.md#GetV1Changes) | **Get** /v1/changes | Stream state changes
[**PostV1Account**](AccountAPI.md#PostV1Account) | **Post** /v1/account | Register new account
[**PostV1AuthenticationForm**](AccountAPI.md#PostV1AuthenticationForm) | **Post** /v1/authentication/form | Generate authentication link
[**PostV1Verifyaccount**](AccountAPI.md#PostV1Verifyaccount) | **Post** /v1/verifyAccount | Verify IMAP and SMTP settings
[**PutV1AccountAccount**](AccountAPI.md#PutV1AccountAccount) | **Put** /v1/account/{account} | Update account info
[**PutV1AccountAccountFlush**](AccountAPI.md#PutV1AccountAccountFlush) | **Put** /v1/account/{account}/flush | Request account flush
[**PutV1AccountAccountReconnect**](AccountAPI.md#PutV1AccountAccountReconnect) | **Put** /v1/account/{account}/reconnect | Request reconnect
[**PutV1AccountAccountSync**](AccountAPI.md#PutV1AccountAccountSync) | **Put** /v1/account/{account}/sync | Request syncing



## DeleteV1AccountAccount

> DeleteRequestResponse DeleteV1AccountAccount(ctx, account).XEeTimeout(xEeTimeout).Execute()

Remove account



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.DeleteV1AccountAccount(context.Background(), account).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.DeleteV1AccountAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1AccountAccount`: DeleteRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.DeleteV1AccountAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1AccountAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

[**DeleteRequestResponse**](DeleteRequestResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1AccountAccount

> AccountResponse GetV1AccountAccount(ctx, account).XEeTimeout(xEeTimeout).Quota(quota).Execute()

Get account info



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
	quota := true // bool | If true, then include quota information in the response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetV1AccountAccount(context.Background(), account).XEeTimeout(xEeTimeout).Quota(quota).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetV1AccountAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1AccountAccount`: AccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetV1AccountAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AccountAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **quota** | **bool** | If true, then include quota information in the response | [default to false]

### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1AccountAccountOauthtoken

> Model2 GetV1AccountAccountOauthtoken(ctx, account).XEeTimeout(xEeTimeout).Execute()

Get OAuth2 access token



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetV1AccountAccountOauthtoken(context.Background(), account).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetV1AccountAccountOauthtoken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1AccountAccountOauthtoken`: Model2
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetV1AccountAccountOauthtoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AccountAccountOauthtokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

[**Model2**](Model2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1AccountAccountServersignatures

> AccountTokenResponse GetV1AccountAccountServersignatures(ctx, account).XEeTimeout(xEeTimeout).Execute()

List Account Signatures



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetV1AccountAccountServersignatures(context.Background(), account).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetV1AccountAccountServersignatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1AccountAccountServersignatures`: AccountTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetV1AccountAccountServersignatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AccountAccountServersignaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

[**AccountTokenResponse**](AccountTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1Accounts

> AccountsFilterResponse GetV1Accounts(ctx).XEeTimeout(xEeTimeout).Page(page).PageSize(pageSize).State(state).Query(query).Execute()

List accounts



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
	state := "state_example" // string | Filter accounts by state (optional)
	query := "query_example" // string | Filter accounts by string match (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetV1Accounts(context.Background()).XEeTimeout(xEeTimeout).Page(page).PageSize(pageSize).State(state).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetV1Accounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1Accounts`: AccountsFilterResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetV1Accounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **page** | **int32** | Page number (zero indexed, so use 0 for first page) | [default to 0]
 **pageSize** | **int32** | How many entries per page | [default to 20]
 **state** | **string** | Filter accounts by state | 
 **query** | **string** | Filter accounts by string match | 

### Return type

[**AccountsFilterResponse**](AccountsFilterResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1Changes

> string GetV1Changes(ctx).XEeTimeout(xEeTimeout).Execute()

Stream state changes



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetV1Changes(context.Background()).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetV1Changes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1Changes`: string
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetV1Changes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1Account

> CreateAccountResponse PostV1Account(ctx).XEeTimeout(xEeTimeout).CreateAccount(createAccount).Execute()

Register new account



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
	createAccount := *openapiclient.NewCreateAccount("Account_example", "Name_example") // CreateAccount |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PostV1Account(context.Background()).XEeTimeout(xEeTimeout).CreateAccount(createAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PostV1Account``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1Account`: CreateAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PostV1Account`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1AccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **createAccount** | [**CreateAccount**](CreateAccount.md) |  | 

### Return type

[**CreateAccountResponse**](CreateAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1AuthenticationForm

> RequestAuthFormResponse PostV1AuthenticationForm(ctx).XEeTimeout(xEeTimeout).RequestAuthForm(requestAuthForm).Execute()

Generate authentication link



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
	requestAuthForm := *openapiclient.NewRequestAuthForm("RedirectUrl_example") // RequestAuthForm |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PostV1AuthenticationForm(context.Background()).XEeTimeout(xEeTimeout).RequestAuthForm(requestAuthForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PostV1AuthenticationForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1AuthenticationForm`: RequestAuthFormResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PostV1AuthenticationForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1AuthenticationFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **requestAuthForm** | [**RequestAuthForm**](RequestAuthForm.md) |  | 

### Return type

[**RequestAuthFormResponse**](RequestAuthFormResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1Verifyaccount

> VerifyAccountResponse PostV1Verifyaccount(ctx).XEeTimeout(xEeTimeout).VerifyAccount(verifyAccount).Execute()

Verify IMAP and SMTP settings



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
	verifyAccount := *openapiclient.NewVerifyAccount() // VerifyAccount |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PostV1Verifyaccount(context.Background()).XEeTimeout(xEeTimeout).VerifyAccount(verifyAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PostV1Verifyaccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1Verifyaccount`: VerifyAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PostV1Verifyaccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1VerifyaccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **verifyAccount** | [**VerifyAccount**](VerifyAccount.md) |  | 

### Return type

[**VerifyAccountResponse**](VerifyAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1AccountAccount

> Model18 PutV1AccountAccount(ctx, account).XEeTimeout(xEeTimeout).UpdateAccount(updateAccount).Execute()

Update account info



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
	updateAccount := *openapiclient.NewUpdateAccount() // UpdateAccount |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PutV1AccountAccount(context.Background(), account).XEeTimeout(xEeTimeout).UpdateAccount(updateAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PutV1AccountAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1AccountAccount`: Model18
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PutV1AccountAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1AccountAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **updateAccount** | [**UpdateAccount**](UpdateAccount.md) |  | 

### Return type

[**Model18**](Model18.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1AccountAccountFlush

> RequestFlushResponse PutV1AccountAccountFlush(ctx, account).XEeTimeout(xEeTimeout).RequestFlush(requestFlush).Execute()

Request account flush



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
	requestFlush := *openapiclient.NewRequestFlush() // RequestFlush |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PutV1AccountAccountFlush(context.Background(), account).XEeTimeout(xEeTimeout).RequestFlush(requestFlush).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PutV1AccountAccountFlush``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1AccountAccountFlush`: RequestFlushResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PutV1AccountAccountFlush`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1AccountAccountFlushRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **requestFlush** | [**RequestFlush**](RequestFlush.md) |  | 

### Return type

[**RequestFlushResponse**](RequestFlushResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1AccountAccountReconnect

> RequestReconnectResponse PutV1AccountAccountReconnect(ctx, account).XEeTimeout(xEeTimeout).RequestReconnect(requestReconnect).Execute()

Request reconnect



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
	requestReconnect := *openapiclient.NewRequestReconnect() // RequestReconnect |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PutV1AccountAccountReconnect(context.Background(), account).XEeTimeout(xEeTimeout).RequestReconnect(requestReconnect).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PutV1AccountAccountReconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1AccountAccountReconnect`: RequestReconnectResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PutV1AccountAccountReconnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1AccountAccountReconnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **requestReconnect** | [**RequestReconnect**](RequestReconnect.md) |  | 

### Return type

[**RequestReconnectResponse**](RequestReconnectResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1AccountAccountSync

> RequestSyncResponse PutV1AccountAccountSync(ctx, account).XEeTimeout(xEeTimeout).RequestSync(requestSync).Execute()

Request syncing



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
	requestSync := *openapiclient.NewRequestSync() // RequestSync |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PutV1AccountAccountSync(context.Background(), account).XEeTimeout(xEeTimeout).RequestSync(requestSync).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PutV1AccountAccountSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1AccountAccountSync`: RequestSyncResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PutV1AccountAccountSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1AccountAccountSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **requestSync** | [**RequestSync**](RequestSync.md) |  | 

### Return type

[**RequestSyncResponse**](RequestSyncResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

