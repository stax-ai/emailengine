# \MessageAPI

All URIs are relative to *https://emailengine.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1AccountAccountMessageMessage**](MessageAPI.md#DeleteV1AccountAccountMessageMessage) | **Delete** /v1/account/{account}/message/{message} | Delete message
[**GetV1AccountAccountAttachmentAttachment**](MessageAPI.md#GetV1AccountAccountAttachmentAttachment) | **Get** /v1/account/{account}/attachment/{attachment} | Download attachment
[**GetV1AccountAccountMessageMessage**](MessageAPI.md#GetV1AccountAccountMessageMessage) | **Get** /v1/account/{account}/message/{message} | Get message information
[**GetV1AccountAccountMessageMessageSource**](MessageAPI.md#GetV1AccountAccountMessageMessageSource) | **Get** /v1/account/{account}/message/{message}/source | Download raw message
[**GetV1AccountAccountMessages**](MessageAPI.md#GetV1AccountAccountMessages) | **Get** /v1/account/{account}/messages | List messages in a folder
[**GetV1AccountAccountTextText**](MessageAPI.md#GetV1AccountAccountTextText) | **Get** /v1/account/{account}/text/{text} | Retrieve message text
[**PostV1AccountAccountMessage**](MessageAPI.md#PostV1AccountAccountMessage) | **Post** /v1/account/{account}/message | Upload message
[**PostV1AccountAccountSearch**](MessageAPI.md#PostV1AccountAccountSearch) | **Post** /v1/account/{account}/search | Search for messages
[**PutV1AccountAccountMessageMessage**](MessageAPI.md#PutV1AccountAccountMessageMessage) | **Put** /v1/account/{account}/message/{message} | Update message
[**PutV1AccountAccountMessageMessageMove**](MessageAPI.md#PutV1AccountAccountMessageMessageMove) | **Put** /v1/account/{account}/message/{message}/move | Move a message to a specified folder



## DeleteV1AccountAccountMessageMessage

> MessageDeleteResponse DeleteV1AccountAccountMessageMessage(ctx, account, message).XEeTimeout(xEeTimeout).Force(force).Execute()

Delete message



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
	message := "message_example" // string | Message ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)
	force := true // bool | Delete message even if not in Trash. Not supported for Gmail API accounts. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.DeleteV1AccountAccountMessageMessage(context.Background(), account, message).XEeTimeout(xEeTimeout).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.DeleteV1AccountAccountMessageMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1AccountAccountMessageMessage`: MessageDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.DeleteV1AccountAccountMessageMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 
**message** | **string** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1AccountAccountMessageMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **force** | **bool** | Delete message even if not in Trash. Not supported for Gmail API accounts. | [default to false]

### Return type

[**MessageDeleteResponse**](MessageDeleteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1AccountAccountAttachmentAttachment

> string GetV1AccountAccountAttachmentAttachment(ctx, account, attachment).XEeTimeout(xEeTimeout).Execute()

Download attachment



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
	attachment := "attachment_example" // string | Attachment ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.GetV1AccountAccountAttachmentAttachment(context.Background(), account, attachment).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.GetV1AccountAccountAttachmentAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1AccountAccountAttachmentAttachment`: string
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.GetV1AccountAccountAttachmentAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 
**attachment** | **string** | Attachment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AccountAccountAttachmentAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1AccountAccountMessageMessage

> MessageDetails GetV1AccountAccountMessageMessage(ctx, account, message).XEeTimeout(xEeTimeout).MaxBytes(maxBytes).TextType(textType).WebSafeHtml(webSafeHtml).EmbedAttachedImages(embedAttachedImages).PreProcessHtml(preProcessHtml).MarkAsSeen(markAsSeen).Execute()

Get message information



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
	message := "message_example" // string | Message ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)
	maxBytes := int32(56) // int32 | Max length of text content (optional)
	textType := "textType_example" // string | Which text content to return, use * for all. By default text content is not returned. (optional)
	webSafeHtml := true // bool | Shorthand option to fetch and preprocess HTML and inline images. Overrides `textType`, `preProcessHtml`, and `embedAttachedImages` options. (optional) (default to false)
	embedAttachedImages := true // bool | If true, then fetches attached images and embeds these in the HTML as data URIs (optional) (default to false)
	preProcessHtml := true // bool | If true, then pre-processes HTML for compatibility (optional) (default to false)
	markAsSeen := true // bool | If true, then marks unseen email as seen while returning the message (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.GetV1AccountAccountMessageMessage(context.Background(), account, message).XEeTimeout(xEeTimeout).MaxBytes(maxBytes).TextType(textType).WebSafeHtml(webSafeHtml).EmbedAttachedImages(embedAttachedImages).PreProcessHtml(preProcessHtml).MarkAsSeen(markAsSeen).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.GetV1AccountAccountMessageMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1AccountAccountMessageMessage`: MessageDetails
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.GetV1AccountAccountMessageMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 
**message** | **string** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AccountAccountMessageMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **maxBytes** | **int32** | Max length of text content | 
 **textType** | **string** | Which text content to return, use * for all. By default text content is not returned. | 
 **webSafeHtml** | **bool** | Shorthand option to fetch and preprocess HTML and inline images. Overrides &#x60;textType&#x60;, &#x60;preProcessHtml&#x60;, and &#x60;embedAttachedImages&#x60; options. | [default to false]
 **embedAttachedImages** | **bool** | If true, then fetches attached images and embeds these in the HTML as data URIs | [default to false]
 **preProcessHtml** | **bool** | If true, then pre-processes HTML for compatibility | [default to false]
 **markAsSeen** | **bool** | If true, then marks unseen email as seen while returning the message | [default to false]

### Return type

[**MessageDetails**](MessageDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1AccountAccountMessageMessageSource

> string GetV1AccountAccountMessageMessageSource(ctx, account, message).XEeTimeout(xEeTimeout).Execute()

Download raw message



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
	message := "message_example" // string | Message ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.GetV1AccountAccountMessageMessageSource(context.Background(), account, message).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.GetV1AccountAccountMessageMessageSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1AccountAccountMessageMessageSource`: string
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.GetV1AccountAccountMessageMessageSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 
**message** | **string** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AccountAccountMessageMessageSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: message/rfc822

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1AccountAccountMessages

> MessageList GetV1AccountAccountMessages(ctx, account).Path(path).XEeTimeout(xEeTimeout).Cursor(cursor).Page(page).PageSize(pageSize).Execute()

List messages in a folder



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
	cursor := "cursor_example" // string | Paging cursor from `nextPageCursor` or `prevPageCursor` value (optional)
	page := int32(56) // int32 | Page number (zero-indexed, so use 0 for the first page). Only supported for IMAP accounts. Deprecated; use the paging cursor instead. If the page cursor value is provided, then the page number value is ignored. (optional) (default to 0)
	pageSize := int32(56) // int32 | How many entries per page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.GetV1AccountAccountMessages(context.Background(), account).Path(path).XEeTimeout(xEeTimeout).Cursor(cursor).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.GetV1AccountAccountMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1AccountAccountMessages`: MessageList
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.GetV1AccountAccountMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AccountAccountMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | Mailbox folder path. Can use special use labels like \&quot;\\Sent\&quot;. Special value \&quot;\\All\&quot; is available for Gmail IMAP, Gmail API, MS Graph API accounts. | 
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **cursor** | **string** | Paging cursor from &#x60;nextPageCursor&#x60; or &#x60;prevPageCursor&#x60; value | 
 **page** | **int32** | Page number (zero-indexed, so use 0 for the first page). Only supported for IMAP accounts. Deprecated; use the paging cursor instead. If the page cursor value is provided, then the page number value is ignored. | [default to 0]
 **pageSize** | **int32** | How many entries per page | [default to 20]

### Return type

[**MessageList**](MessageList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1AccountAccountTextText

> TextResponse GetV1AccountAccountTextText(ctx, account, text).XEeTimeout(xEeTimeout).MaxBytes(maxBytes).TextType(textType).Execute()

Retrieve message text



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
	text := "text_example" // string | Message text ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)
	maxBytes := int32(56) // int32 | Max length of text content (optional)
	textType := "textType_example" // string | Which text content to return, use * for all. By default all contents are returned. (optional) (default to "*")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.GetV1AccountAccountTextText(context.Background(), account, text).XEeTimeout(xEeTimeout).MaxBytes(maxBytes).TextType(textType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.GetV1AccountAccountTextText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1AccountAccountTextText`: TextResponse
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.GetV1AccountAccountTextText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 
**text** | **string** | Message text ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AccountAccountTextTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **maxBytes** | **int32** | Max length of text content | 
 **textType** | **string** | Which text content to return, use * for all. By default all contents are returned. | [default to &quot;*&quot;]

### Return type

[**TextResponse**](TextResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1AccountAccountMessage

> MessageUploadResponse PostV1AccountAccountMessage(ctx, account).XEeTimeout(xEeTimeout).MessageUpload(messageUpload).Execute()

Upload message



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
	messageUpload := *openapiclient.NewMessageUpload("Path_example") // MessageUpload |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.PostV1AccountAccountMessage(context.Background(), account).XEeTimeout(xEeTimeout).MessageUpload(messageUpload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.PostV1AccountAccountMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1AccountAccountMessage`: MessageUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.PostV1AccountAccountMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1AccountAccountMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **messageUpload** | [**MessageUpload**](MessageUpload.md) |  | 

### Return type

[**MessageUploadResponse**](MessageUploadResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1AccountAccountSearch

> MessageList PostV1AccountAccountSearch(ctx, account).XEeTimeout(xEeTimeout).Path(path).Cursor(cursor).Page(page).PageSize(pageSize).UseOutlookSearch(useOutlookSearch).Model9(model9).Execute()

Search for messages



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
	path := "path_example" // string | Mailbox folder path. Can use special use labels like \"\\Sent\". Special value \"\\All\" is available for Gmail IMAP, Gmail API, MS Graph API accounts. (optional)
	cursor := "cursor_example" // string | Paging cursor from `nextPageCursor` or `prevPageCursor` value (optional)
	page := int32(56) // int32 | Page number (zero-indexed, so use 0 for the first page). Only supported for IMAP accounts. Deprecated; use the paging cursor instead. If the page cursor value is provided, then the page number value is ignored. (optional) (default to 0)
	pageSize := int32(56) // int32 | How many entries per page (optional) (default to 20)
	useOutlookSearch := true // bool | MS Graph only. If enabled, uses the $search parameter for MS Graph search queries instead of $filter. This allows searching the \"to\", \"cc\", \"bcc\", \"larger\", \"smaller\", \"body\", \"before\", \"sentBefore\", \"since\", and the \"sentSince\" fields. Note that $search returns up to 1,000 results, does not indicate the total number of matching results or pages, and returns results sorted by relevance rather than date. (optional)
	model9 := *openapiclient.NewModel9(*openapiclient.NewSearchQuery()) // Model9 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.PostV1AccountAccountSearch(context.Background(), account).XEeTimeout(xEeTimeout).Path(path).Cursor(cursor).Page(page).PageSize(pageSize).UseOutlookSearch(useOutlookSearch).Model9(model9).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.PostV1AccountAccountSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1AccountAccountSearch`: MessageList
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.PostV1AccountAccountSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1AccountAccountSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **path** | **string** | Mailbox folder path. Can use special use labels like \&quot;\\Sent\&quot;. Special value \&quot;\\All\&quot; is available for Gmail IMAP, Gmail API, MS Graph API accounts. | 
 **cursor** | **string** | Paging cursor from &#x60;nextPageCursor&#x60; or &#x60;prevPageCursor&#x60; value | 
 **page** | **int32** | Page number (zero-indexed, so use 0 for the first page). Only supported for IMAP accounts. Deprecated; use the paging cursor instead. If the page cursor value is provided, then the page number value is ignored. | [default to 0]
 **pageSize** | **int32** | How many entries per page | [default to 20]
 **useOutlookSearch** | **bool** | MS Graph only. If enabled, uses the $search parameter for MS Graph search queries instead of $filter. This allows searching the \&quot;to\&quot;, \&quot;cc\&quot;, \&quot;bcc\&quot;, \&quot;larger\&quot;, \&quot;smaller\&quot;, \&quot;body\&quot;, \&quot;before\&quot;, \&quot;sentBefore\&quot;, \&quot;since\&quot;, and the \&quot;sentSince\&quot; fields. Note that $search returns up to 1,000 results, does not indicate the total number of matching results or pages, and returns results sorted by relevance rather than date. | 
 **model9** | [**Model9**](Model9.md) |  | 

### Return type

[**MessageList**](MessageList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1AccountAccountMessageMessage

> Model29 PutV1AccountAccountMessageMessage(ctx, account, message).XEeTimeout(xEeTimeout).MessageUpdate(messageUpdate).Execute()

Update message



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
	message := "message_example" // string | Message ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)
	messageUpdate := *openapiclient.NewMessageUpdate() // MessageUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.PutV1AccountAccountMessageMessage(context.Background(), account, message).XEeTimeout(xEeTimeout).MessageUpdate(messageUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.PutV1AccountAccountMessageMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1AccountAccountMessageMessage`: Model29
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.PutV1AccountAccountMessageMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 
**message** | **string** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1AccountAccountMessageMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **messageUpdate** | [**MessageUpdate**](MessageUpdate.md) |  | 

### Return type

[**Model29**](Model29.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1AccountAccountMessageMessageMove

> MessageMoveResponse PutV1AccountAccountMessageMessageMove(ctx, account, message).XEeTimeout(xEeTimeout).MessageMove(messageMove).Execute()

Move a message to a specified folder



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
	message := "message_example" // string | Message ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)
	messageMove := *openapiclient.NewMessageMove("Path_example") // MessageMove |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.PutV1AccountAccountMessageMessageMove(context.Background(), account, message).XEeTimeout(xEeTimeout).MessageMove(messageMove).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.PutV1AccountAccountMessageMessageMove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1AccountAccountMessageMessageMove`: MessageMoveResponse
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.PutV1AccountAccountMessageMessageMove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | Unique identifier for the email account | 
**message** | **string** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1AccountAccountMessageMessageMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **messageMove** | [**MessageMove**](MessageMove.md) |  | 

### Return type

[**MessageMoveResponse**](MessageMoveResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

