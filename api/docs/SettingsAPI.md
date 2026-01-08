# \SettingsAPI

All URIs are relative to *https://emailengine.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1Autoconfig**](SettingsAPI.md#GetV1Autoconfig) | **Get** /v1/autoconfig | Discover Email settings
[**GetV1Settings**](SettingsAPI.md#GetV1Settings) | **Get** /v1/settings | List specific settings
[**GetV1SettingsQueueQueue**](SettingsAPI.md#GetV1SettingsQueueQueue) | **Get** /v1/settings/queue/{queue} | Show queue information
[**PostV1Settings**](SettingsAPI.md#PostV1Settings) | **Post** /v1/settings | Set setting values
[**PutV1SettingsQueueQueue**](SettingsAPI.md#PutV1SettingsQueueQueue) | **Put** /v1/settings/queue/{queue} | Set queue settings



## GetV1Autoconfig

> DiscoveredEmailSettings GetV1Autoconfig(ctx).Email(email).XEeTimeout(xEeTimeout).Execute()

Discover Email settings



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
	email := "email_example" // string | Email address to discover email settings for
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.GetV1Autoconfig(context.Background()).Email(email).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetV1Autoconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1Autoconfig`: DiscoveredEmailSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetV1Autoconfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AutoconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Email address to discover email settings for | 
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

[**DiscoveredEmailSettings**](DiscoveredEmailSettings.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1Settings

> SettingsQueryResponse GetV1Settings(ctx).XEeTimeout(xEeTimeout).EventTypes(eventTypes).WebhooksEnabled(webhooksEnabled).Webhooks(webhooks).WebhookEvents(webhookEvents).WebhooksCustomHeaders(webhooksCustomHeaders).NotifyHeaders(notifyHeaders).ServiceUrl(serviceUrl).NotificationBaseUrl(notificationBaseUrl).TrackSentMessages(trackSentMessages).TrackClicks(trackClicks).TrackOpens(trackOpens).ImapIndexer(imapIndexer).ResolveGmailCategories(resolveGmailCategories).IgnoreMailCertErrors(ignoreMailCertErrors).GenerateEmailSummary(generateEmailSummary).GenerateRiskAssessment(generateRiskAssessment).OpenAiAPIKey(openAiAPIKey).OpenAiModel(openAiModel).OpenAiAPIUrl(openAiAPIUrl).DocumentStoreChatModel(documentStoreChatModel).OpenAiTemperature(openAiTemperature).OpenAiTopP(openAiTopP).OpenAiMaxTokens(openAiMaxTokens).OpenAiPrompt(openAiPrompt).OpenAiGenerateEmbeddings(openAiGenerateEmbeddings).InboxNewOnly(inboxNewOnly).ServiceSecret(serviceSecret).AuthServer(authServer).ProxyEnabled(proxyEnabled).ProxyUrl(proxyUrl).SmtpEhloName(smtpEhloName).NotifyText(notifyText).NotifyWebSafeHtml(notifyWebSafeHtml).NotifyTextSize(notifyTextSize).NotifyAttachments(notifyAttachments).NotifyAttachmentSize(notifyAttachmentSize).NotifyCalendarEvents(notifyCalendarEvents).GmailEnabled(gmailEnabled).GmailClientId(gmailClientId).GmailClientSecret(gmailClientSecret).GmailRedirectUrl(gmailRedirectUrl).GmailExtraScopes(gmailExtraScopes).OutlookEnabled(outlookEnabled).OutlookClientId(outlookClientId).OutlookClientSecret(outlookClientSecret).OutlookRedirectUrl(outlookRedirectUrl).OutlookAuthority(outlookAuthority).OutlookExtraScopes(outlookExtraScopes).MailRuEnabled(mailRuEnabled).MailRuClientId(mailRuClientId).MailRuClientSecret(mailRuClientSecret).MailRuRedirectUrl(mailRuRedirectUrl).MailRuExtraScopes(mailRuExtraScopes).ServiceClient(serviceClient).ServiceKey(serviceKey).ServiceExtraScopes(serviceExtraScopes).DocumentStoreEnabled(documentStoreEnabled).DocumentStoreUrl(documentStoreUrl).DocumentStoreIndex(documentStoreIndex).DocumentStoreAuthEnabled(documentStoreAuthEnabled).DocumentStoreUsername(documentStoreUsername).DocumentStorePassword(documentStorePassword).DocumentStoreGenerateEmbeddings(documentStoreGenerateEmbeddings).DocumentStorePreProcessingEnabled(documentStorePreProcessingEnabled).Logs(logs).ImapStrategy(imapStrategy).SmtpStrategy(smtpStrategy).LocalAddresses(localAddresses).SmtpServerEnabled(smtpServerEnabled).SmtpServerPort(smtpServerPort).SmtpServerHost(smtpServerHost).SmtpServerProxy(smtpServerProxy).SmtpServerAuthEnabled(smtpServerAuthEnabled).SmtpServerPassword(smtpServerPassword).SmtpServerTLSEnabled(smtpServerTLSEnabled).ImapProxyServerEnabled(imapProxyServerEnabled).ImapProxyServerPort(imapProxyServerPort).ImapProxyServerHost(imapProxyServerHost).ImapProxyServerProxy(imapProxyServerProxy).ImapProxyServerPassword(imapProxyServerPassword).ImapProxyServerTLSEnabled(imapProxyServerTLSEnabled).QueueKeep(queueKeep).DeliveryAttempts(deliveryAttempts).TemplateHeader(templateHeader).TemplateHtmlHead(templateHtmlHead).ScriptEnv(scriptEnv).EnableApiProxy(enableApiProxy).Locale(locale).Timezone(timezone).PageBrandName(pageBrandName).OpenAiPreProcessingFn(openAiPreProcessingFn).ImapClientName(imapClientName).ImapClientVersion(imapClientVersion).ImapClientVendor(imapClientVendor).ImapClientSupportUrl(imapClientSupportUrl).Execute()

List specific settings



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
	eventTypes := true // bool |  (optional) (default to false)
	webhooksEnabled := true // bool |  (optional) (default to false)
	webhooks := true // bool |  (optional) (default to false)
	webhookEvents := true // bool |  (optional) (default to false)
	webhooksCustomHeaders := true // bool |  (optional) (default to false)
	notifyHeaders := true // bool |  (optional) (default to false)
	serviceUrl := true // bool |  (optional) (default to false)
	notificationBaseUrl := true // bool |  (optional) (default to false)
	trackSentMessages := true // bool |  (optional) (default to false)
	trackClicks := true // bool |  (optional) (default to false)
	trackOpens := true // bool |  (optional) (default to false)
	imapIndexer := true // bool |  (optional) (default to false)
	resolveGmailCategories := true // bool |  (optional) (default to false)
	ignoreMailCertErrors := true // bool |  (optional) (default to false)
	generateEmailSummary := true // bool |  (optional) (default to false)
	generateRiskAssessment := true // bool |  (optional) (default to false)
	openAiAPIKey := true // bool |  (optional) (default to false)
	openAiModel := true // bool |  (optional) (default to false)
	openAiAPIUrl := true // bool |  (optional) (default to false)
	documentStoreChatModel := true // bool |  (optional) (default to false)
	openAiTemperature := true // bool |  (optional) (default to false)
	openAiTopP := true // bool |  (optional) (default to false)
	openAiMaxTokens := true // bool |  (optional) (default to false)
	openAiPrompt := true // bool |  (optional) (default to false)
	openAiGenerateEmbeddings := true // bool |  (optional) (default to false)
	inboxNewOnly := true // bool |  (optional) (default to false)
	serviceSecret := true // bool |  (optional) (default to false)
	authServer := true // bool |  (optional) (default to false)
	proxyEnabled := true // bool |  (optional) (default to false)
	proxyUrl := true // bool |  (optional) (default to false)
	smtpEhloName := true // bool |  (optional) (default to false)
	notifyText := true // bool |  (optional) (default to false)
	notifyWebSafeHtml := true // bool |  (optional) (default to false)
	notifyTextSize := true // bool |  (optional) (default to false)
	notifyAttachments := true // bool |  (optional) (default to false)
	notifyAttachmentSize := true // bool |  (optional) (default to false)
	notifyCalendarEvents := true // bool |  (optional) (default to false)
	gmailEnabled := true // bool |  (optional) (default to false)
	gmailClientId := true // bool |  (optional) (default to false)
	gmailClientSecret := true // bool |  (optional) (default to false)
	gmailRedirectUrl := true // bool |  (optional) (default to false)
	gmailExtraScopes := true // bool |  (optional) (default to false)
	outlookEnabled := true // bool |  (optional) (default to false)
	outlookClientId := true // bool |  (optional) (default to false)
	outlookClientSecret := true // bool |  (optional) (default to false)
	outlookRedirectUrl := true // bool |  (optional) (default to false)
	outlookAuthority := true // bool |  (optional) (default to false)
	outlookExtraScopes := true // bool |  (optional) (default to false)
	mailRuEnabled := true // bool |  (optional) (default to false)
	mailRuClientId := true // bool |  (optional) (default to false)
	mailRuClientSecret := true // bool |  (optional) (default to false)
	mailRuRedirectUrl := true // bool |  (optional) (default to false)
	mailRuExtraScopes := true // bool |  (optional) (default to false)
	serviceClient := true // bool |  (optional) (default to false)
	serviceKey := true // bool |  (optional) (default to false)
	serviceExtraScopes := true // bool |  (optional) (default to false)
	documentStoreEnabled := true // bool |  (optional) (default to false)
	documentStoreUrl := true // bool |  (optional) (default to false)
	documentStoreIndex := true // bool |  (optional) (default to false)
	documentStoreAuthEnabled := true // bool |  (optional) (default to false)
	documentStoreUsername := true // bool |  (optional) (default to false)
	documentStorePassword := true // bool |  (optional) (default to false)
	documentStoreGenerateEmbeddings := true // bool |  (optional) (default to false)
	documentStorePreProcessingEnabled := true // bool |  (optional) (default to false)
	logs := true // bool |  (optional) (default to false)
	imapStrategy := true // bool |  (optional) (default to false)
	smtpStrategy := true // bool |  (optional) (default to false)
	localAddresses := true // bool |  (optional) (default to false)
	smtpServerEnabled := true // bool |  (optional) (default to false)
	smtpServerPort := true // bool |  (optional) (default to false)
	smtpServerHost := true // bool |  (optional) (default to false)
	smtpServerProxy := true // bool |  (optional) (default to false)
	smtpServerAuthEnabled := true // bool |  (optional) (default to false)
	smtpServerPassword := true // bool |  (optional) (default to false)
	smtpServerTLSEnabled := true // bool |  (optional) (default to false)
	imapProxyServerEnabled := true // bool |  (optional) (default to false)
	imapProxyServerPort := true // bool |  (optional) (default to false)
	imapProxyServerHost := true // bool |  (optional) (default to false)
	imapProxyServerProxy := true // bool |  (optional) (default to false)
	imapProxyServerPassword := true // bool |  (optional) (default to false)
	imapProxyServerTLSEnabled := true // bool |  (optional) (default to false)
	queueKeep := true // bool |  (optional) (default to false)
	deliveryAttempts := true // bool |  (optional) (default to false)
	templateHeader := true // bool |  (optional) (default to false)
	templateHtmlHead := true // bool |  (optional) (default to false)
	scriptEnv := true // bool |  (optional) (default to false)
	enableApiProxy := true // bool |  (optional) (default to false)
	locale := true // bool |  (optional) (default to false)
	timezone := true // bool |  (optional) (default to false)
	pageBrandName := true // bool |  (optional) (default to false)
	openAiPreProcessingFn := true // bool |  (optional) (default to false)
	imapClientName := true // bool |  (optional) (default to false)
	imapClientVersion := true // bool |  (optional) (default to false)
	imapClientVendor := true // bool |  (optional) (default to false)
	imapClientSupportUrl := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.GetV1Settings(context.Background()).XEeTimeout(xEeTimeout).EventTypes(eventTypes).WebhooksEnabled(webhooksEnabled).Webhooks(webhooks).WebhookEvents(webhookEvents).WebhooksCustomHeaders(webhooksCustomHeaders).NotifyHeaders(notifyHeaders).ServiceUrl(serviceUrl).NotificationBaseUrl(notificationBaseUrl).TrackSentMessages(trackSentMessages).TrackClicks(trackClicks).TrackOpens(trackOpens).ImapIndexer(imapIndexer).ResolveGmailCategories(resolveGmailCategories).IgnoreMailCertErrors(ignoreMailCertErrors).GenerateEmailSummary(generateEmailSummary).GenerateRiskAssessment(generateRiskAssessment).OpenAiAPIKey(openAiAPIKey).OpenAiModel(openAiModel).OpenAiAPIUrl(openAiAPIUrl).DocumentStoreChatModel(documentStoreChatModel).OpenAiTemperature(openAiTemperature).OpenAiTopP(openAiTopP).OpenAiMaxTokens(openAiMaxTokens).OpenAiPrompt(openAiPrompt).OpenAiGenerateEmbeddings(openAiGenerateEmbeddings).InboxNewOnly(inboxNewOnly).ServiceSecret(serviceSecret).AuthServer(authServer).ProxyEnabled(proxyEnabled).ProxyUrl(proxyUrl).SmtpEhloName(smtpEhloName).NotifyText(notifyText).NotifyWebSafeHtml(notifyWebSafeHtml).NotifyTextSize(notifyTextSize).NotifyAttachments(notifyAttachments).NotifyAttachmentSize(notifyAttachmentSize).NotifyCalendarEvents(notifyCalendarEvents).GmailEnabled(gmailEnabled).GmailClientId(gmailClientId).GmailClientSecret(gmailClientSecret).GmailRedirectUrl(gmailRedirectUrl).GmailExtraScopes(gmailExtraScopes).OutlookEnabled(outlookEnabled).OutlookClientId(outlookClientId).OutlookClientSecret(outlookClientSecret).OutlookRedirectUrl(outlookRedirectUrl).OutlookAuthority(outlookAuthority).OutlookExtraScopes(outlookExtraScopes).MailRuEnabled(mailRuEnabled).MailRuClientId(mailRuClientId).MailRuClientSecret(mailRuClientSecret).MailRuRedirectUrl(mailRuRedirectUrl).MailRuExtraScopes(mailRuExtraScopes).ServiceClient(serviceClient).ServiceKey(serviceKey).ServiceExtraScopes(serviceExtraScopes).DocumentStoreEnabled(documentStoreEnabled).DocumentStoreUrl(documentStoreUrl).DocumentStoreIndex(documentStoreIndex).DocumentStoreAuthEnabled(documentStoreAuthEnabled).DocumentStoreUsername(documentStoreUsername).DocumentStorePassword(documentStorePassword).DocumentStoreGenerateEmbeddings(documentStoreGenerateEmbeddings).DocumentStorePreProcessingEnabled(documentStorePreProcessingEnabled).Logs(logs).ImapStrategy(imapStrategy).SmtpStrategy(smtpStrategy).LocalAddresses(localAddresses).SmtpServerEnabled(smtpServerEnabled).SmtpServerPort(smtpServerPort).SmtpServerHost(smtpServerHost).SmtpServerProxy(smtpServerProxy).SmtpServerAuthEnabled(smtpServerAuthEnabled).SmtpServerPassword(smtpServerPassword).SmtpServerTLSEnabled(smtpServerTLSEnabled).ImapProxyServerEnabled(imapProxyServerEnabled).ImapProxyServerPort(imapProxyServerPort).ImapProxyServerHost(imapProxyServerHost).ImapProxyServerProxy(imapProxyServerProxy).ImapProxyServerPassword(imapProxyServerPassword).ImapProxyServerTLSEnabled(imapProxyServerTLSEnabled).QueueKeep(queueKeep).DeliveryAttempts(deliveryAttempts).TemplateHeader(templateHeader).TemplateHtmlHead(templateHtmlHead).ScriptEnv(scriptEnv).EnableApiProxy(enableApiProxy).Locale(locale).Timezone(timezone).PageBrandName(pageBrandName).OpenAiPreProcessingFn(openAiPreProcessingFn).ImapClientName(imapClientName).ImapClientVersion(imapClientVersion).ImapClientVendor(imapClientVendor).ImapClientSupportUrl(imapClientSupportUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetV1Settings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1Settings`: SettingsQueryResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetV1Settings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1SettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **eventTypes** | **bool** |  | [default to false]
 **webhooksEnabled** | **bool** |  | [default to false]
 **webhooks** | **bool** |  | [default to false]
 **webhookEvents** | **bool** |  | [default to false]
 **webhooksCustomHeaders** | **bool** |  | [default to false]
 **notifyHeaders** | **bool** |  | [default to false]
 **serviceUrl** | **bool** |  | [default to false]
 **notificationBaseUrl** | **bool** |  | [default to false]
 **trackSentMessages** | **bool** |  | [default to false]
 **trackClicks** | **bool** |  | [default to false]
 **trackOpens** | **bool** |  | [default to false]
 **imapIndexer** | **bool** |  | [default to false]
 **resolveGmailCategories** | **bool** |  | [default to false]
 **ignoreMailCertErrors** | **bool** |  | [default to false]
 **generateEmailSummary** | **bool** |  | [default to false]
 **generateRiskAssessment** | **bool** |  | [default to false]
 **openAiAPIKey** | **bool** |  | [default to false]
 **openAiModel** | **bool** |  | [default to false]
 **openAiAPIUrl** | **bool** |  | [default to false]
 **documentStoreChatModel** | **bool** |  | [default to false]
 **openAiTemperature** | **bool** |  | [default to false]
 **openAiTopP** | **bool** |  | [default to false]
 **openAiMaxTokens** | **bool** |  | [default to false]
 **openAiPrompt** | **bool** |  | [default to false]
 **openAiGenerateEmbeddings** | **bool** |  | [default to false]
 **inboxNewOnly** | **bool** |  | [default to false]
 **serviceSecret** | **bool** |  | [default to false]
 **authServer** | **bool** |  | [default to false]
 **proxyEnabled** | **bool** |  | [default to false]
 **proxyUrl** | **bool** |  | [default to false]
 **smtpEhloName** | **bool** |  | [default to false]
 **notifyText** | **bool** |  | [default to false]
 **notifyWebSafeHtml** | **bool** |  | [default to false]
 **notifyTextSize** | **bool** |  | [default to false]
 **notifyAttachments** | **bool** |  | [default to false]
 **notifyAttachmentSize** | **bool** |  | [default to false]
 **notifyCalendarEvents** | **bool** |  | [default to false]
 **gmailEnabled** | **bool** |  | [default to false]
 **gmailClientId** | **bool** |  | [default to false]
 **gmailClientSecret** | **bool** |  | [default to false]
 **gmailRedirectUrl** | **bool** |  | [default to false]
 **gmailExtraScopes** | **bool** |  | [default to false]
 **outlookEnabled** | **bool** |  | [default to false]
 **outlookClientId** | **bool** |  | [default to false]
 **outlookClientSecret** | **bool** |  | [default to false]
 **outlookRedirectUrl** | **bool** |  | [default to false]
 **outlookAuthority** | **bool** |  | [default to false]
 **outlookExtraScopes** | **bool** |  | [default to false]
 **mailRuEnabled** | **bool** |  | [default to false]
 **mailRuClientId** | **bool** |  | [default to false]
 **mailRuClientSecret** | **bool** |  | [default to false]
 **mailRuRedirectUrl** | **bool** |  | [default to false]
 **mailRuExtraScopes** | **bool** |  | [default to false]
 **serviceClient** | **bool** |  | [default to false]
 **serviceKey** | **bool** |  | [default to false]
 **serviceExtraScopes** | **bool** |  | [default to false]
 **documentStoreEnabled** | **bool** |  | [default to false]
 **documentStoreUrl** | **bool** |  | [default to false]
 **documentStoreIndex** | **bool** |  | [default to false]
 **documentStoreAuthEnabled** | **bool** |  | [default to false]
 **documentStoreUsername** | **bool** |  | [default to false]
 **documentStorePassword** | **bool** |  | [default to false]
 **documentStoreGenerateEmbeddings** | **bool** |  | [default to false]
 **documentStorePreProcessingEnabled** | **bool** |  | [default to false]
 **logs** | **bool** |  | [default to false]
 **imapStrategy** | **bool** |  | [default to false]
 **smtpStrategy** | **bool** |  | [default to false]
 **localAddresses** | **bool** |  | [default to false]
 **smtpServerEnabled** | **bool** |  | [default to false]
 **smtpServerPort** | **bool** |  | [default to false]
 **smtpServerHost** | **bool** |  | [default to false]
 **smtpServerProxy** | **bool** |  | [default to false]
 **smtpServerAuthEnabled** | **bool** |  | [default to false]
 **smtpServerPassword** | **bool** |  | [default to false]
 **smtpServerTLSEnabled** | **bool** |  | [default to false]
 **imapProxyServerEnabled** | **bool** |  | [default to false]
 **imapProxyServerPort** | **bool** |  | [default to false]
 **imapProxyServerHost** | **bool** |  | [default to false]
 **imapProxyServerProxy** | **bool** |  | [default to false]
 **imapProxyServerPassword** | **bool** |  | [default to false]
 **imapProxyServerTLSEnabled** | **bool** |  | [default to false]
 **queueKeep** | **bool** |  | [default to false]
 **deliveryAttempts** | **bool** |  | [default to false]
 **templateHeader** | **bool** |  | [default to false]
 **templateHtmlHead** | **bool** |  | [default to false]
 **scriptEnv** | **bool** |  | [default to false]
 **enableApiProxy** | **bool** |  | [default to false]
 **locale** | **bool** |  | [default to false]
 **timezone** | **bool** |  | [default to false]
 **pageBrandName** | **bool** |  | [default to false]
 **openAiPreProcessingFn** | **bool** |  | [default to false]
 **imapClientName** | **bool** |  | [default to false]
 **imapClientVersion** | **bool** |  | [default to false]
 **imapClientVendor** | **bool** |  | [default to false]
 **imapClientSupportUrl** | **bool** |  | [default to false]

### Return type

[**SettingsQueryResponse**](SettingsQueryResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV1SettingsQueueQueue

> SettingsQueueResponse GetV1SettingsQueueQueue(ctx, queue).XEeTimeout(xEeTimeout).Execute()

Show queue information



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
	queue := "queue_example" // string | Queue ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.GetV1SettingsQueueQueue(context.Background(), queue).XEeTimeout(xEeTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetV1SettingsQueueQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1SettingsQueueQueue`: SettingsQueueResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetV1SettingsQueueQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queue** | **string** | Queue ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1SettingsQueueQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 

### Return type

[**SettingsQueueResponse**](SettingsQueueResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1Settings

> SettingsUpdatedResponse PostV1Settings(ctx).XEeTimeout(xEeTimeout).Settings(settings).Execute()

Set setting values



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
	settings := *openapiclient.NewSettings() // Settings |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.PostV1Settings(context.Background()).XEeTimeout(xEeTimeout).Settings(settings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PostV1Settings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1Settings`: SettingsUpdatedResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.PostV1Settings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1SettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **settings** | [**Settings**](Settings.md) |  | 

### Return type

[**SettingsUpdatedResponse**](SettingsUpdatedResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1SettingsQueueQueue

> SettingsPutQueueResponse PutV1SettingsQueueQueue(ctx, queue).XEeTimeout(xEeTimeout).SettingsPutQueuePayload(settingsPutQueuePayload).Execute()

Set queue settings



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
	queue := "queue_example" // string | Queue ID
	xEeTimeout := int32(56) // int32 | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) (optional)
	settingsPutQueuePayload := *openapiclient.NewSettingsPutQueuePayload() // SettingsPutQueuePayload |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.PutV1SettingsQueueQueue(context.Background(), queue).XEeTimeout(xEeTimeout).SettingsPutQueuePayload(settingsPutQueuePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PutV1SettingsQueueQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1SettingsQueueQueue`: SettingsPutQueueResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.PutV1SettingsQueueQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queue** | **string** | Queue ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1SettingsQueueQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEeTimeout** | **int32** | Request timeout in milliseconds (overrides EENGINE_TIMEOUT environment variable) | 
 **settingsPutQueuePayload** | [**SettingsPutQueuePayload**](SettingsPutQueuePayload.md) |  | 

### Return type

[**SettingsPutQueueResponse**](SettingsPutQueueResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

