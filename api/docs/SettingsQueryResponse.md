# SettingsQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthServer** | Pointer to **string** | External authentication service URL for retrieving mailbox credentials dynamically | [optional] 
**DeliveryAttempts** | Pointer to **int32** | Maximum number of delivery attempts before marking a message as permanently failed | [optional] 
**EnableApiProxy** | Pointer to **bool** | Trust X-Forwarded-* headers when behind a reverse proxy | [optional] 
**GenerateEmailSummary** | Pointer to **bool** | Generate AI-powered summaries for incoming emails using OpenAI | [optional] 
**IgnoreMailCertErrors** | Pointer to **bool** | Allow connections to mail servers with self-signed or invalid TLS certificates (not recommended for production) | [optional] 
**ImapClientName** | Pointer to **string** | Client name advertised via IMAP ID extension | [optional] 
**ImapClientSupportUrl** | Pointer to **string** | Support URL advertised via IMAP ID extension | [optional] 
**ImapClientVendor** | Pointer to **string** | Vendor name advertised via IMAP ID extension | [optional] 
**ImapClientVersion** | Pointer to **string** | Client version advertised via IMAP ID extension | [optional] 
**ImapIndexer** | Pointer to [**ImapIndexer**](ImapIndexer.md) |  | [optional] 
**ImapProxyServerEnabled** | Pointer to **bool** | Enable the IMAP proxy server | [optional] 
**ImapProxyServerHost** | Pointer to **string** | IP address to bind the IMAP proxy to (empty &#x3D; all interfaces) | [optional] 
**ImapProxyServerPassword** | Pointer to **NullableString** | Password for IMAP proxy authentication (null &#x3D; disable authentication) | [optional] 
**ImapProxyServerPort** | Pointer to **int32** | Port number for the IMAP proxy server | [optional] 
**ImapProxyServerProxy** | Pointer to **bool** | Enable PROXY protocol support for the IMAP proxy | [optional] 
**ImapProxyServerTLSEnabled** | Pointer to **bool** | Enable TLS support for the IMAP proxy | [optional] 
**ImapStrategy** | Pointer to [**ImapStrategy**](ImapStrategy.md) |  | [optional] 
**InboxNewOnly** | Pointer to **bool** | Trigger \&quot;messageNew\&quot; webhooks only for messages arriving in the Inbox folder | [optional] 
**LocalAddresses** | Pointer to **[]string** | List of local IP addresses to use for outbound connections (requires appropriate network configuration) | [optional] 
**Locale** | Pointer to [**Locale**](Locale.md) |  | [optional] 
**Logs** | Pointer to [**LogSettings**](LogSettings.md) |  | [optional] 
**NotificationBaseUrl** | Pointer to **NullableString** | Public callback URL for external OAuth providers. Falls back to serviceUrl if not set | [optional] 
**NotifyAttachmentSize** | Pointer to **int32** | Maximum size (in bytes) per attachment to include in webhook payloads | [optional] 
**NotifyAttachments** | Pointer to **bool** | Include attachment data in webhook payloads | [optional] 
**NotifyCalendarEvents** | Pointer to **bool** | Include parsed calendar event data in webhook payloads | [optional] 
**NotifyHeaders** | Pointer to **[]string** | Email headers to include in webhook payloads for additional context | [optional] 
**NotifyText** | Pointer to **bool** | Include plain text message content in webhook payloads | [optional] 
**NotifyTextSize** | Pointer to **int32** | Maximum size (in bytes) of text content to include in webhook payloads | [optional] 
**NotifyWebSafeHtml** | Pointer to **bool** | Sanitize HTML content to remove potentially dangerous elements before including in webhooks | [optional] 
**OpenAiAPIKey** | Pointer to **string** | Your OpenAI API key for AI features | [optional] 
**OpenAiAPIUrl** | Pointer to **string** | Base URL for OpenAI API (change for OpenAI-compatible services) | [optional] 
**OpenAiGenerateEmbeddings** | Pointer to **bool** | Generate vector embeddings for semantic search and similarity matching | [optional] 
**OpenAiMaxTokens** | Pointer to **float32** | Maximum tokens limit for OpenAI API requests (defaults: GPT-5: 18000, GPT-4: 6500, GPT-3.5: 3000) | [optional] 
**OpenAiModel** | Pointer to **string** | OpenAI model to use for text generation | [optional] 
**OpenAiPreProcessingFn** | Pointer to **string** | JavaScript function to filter emails before AI processing (return true to process, false to skip) | [optional] 
**OpenAiPrompt** | Pointer to **string** | Custom system prompt to guide AI behavior when processing emails | [optional] 
**OpenAiTemperature** | Pointer to **float32** | Controls randomness in AI responses (0 &#x3D; deterministic, 2 &#x3D; very creative) | [optional] 
**OpenAiTopP** | Pointer to **float32** | Nucleus sampling parameter for AI text generation (0-1, lower &#x3D; more focused) | [optional] 
**PageBrandName** | Pointer to **NullableString** | Brand name displayed in page titles | [optional] 
**ProxyEnabled** | Pointer to **bool** | Route outbound connections through a proxy server | [optional] 
**ProxyUrl** | Pointer to **string** | Proxy server URL for outbound connections | [optional] 
**QueueKeep** | Pointer to **int32** | Number of completed and failed queue entries to retain for debugging | [optional] 
**ResolveGmailCategories** | Pointer to **bool** | Automatically detect and categorize Gmail tabs (Primary, Social, Promotions, etc.) for IMAP connections | [optional] 
**ScriptEnv** | Pointer to **string** | JSON object containing environment variables available to pre-processing scripts | [optional] 
**ServiceSecret** | Pointer to **string** | HMAC secret for signing API requests and validating webhooks | [optional] 
**ServiceUrl** | Pointer to **string** | Base URL of this EmailEngine instance (used for generating public URLs, path component is ignored) | [optional] 
**SmtpEhloName** | Pointer to **NullableString** | Hostname to use in SMTP EHLO/HELO commands (defaults to system hostname) | [optional] 
**SmtpServerAuthEnabled** | Pointer to **bool** | Require SMTP authentication for incoming connections | [optional] 
**SmtpServerEnabled** | Pointer to **bool** | Enable the built-in SMTP server for receiving emails | [optional] 
**SmtpServerHost** | Pointer to **string** | IP address to bind the SMTP server to (empty &#x3D; all interfaces) | [optional] 
**SmtpServerPassword** | Pointer to **NullableString** | Password for SMTP authentication (null &#x3D; disable authentication) | [optional] 
**SmtpServerPort** | Pointer to **int32** | Port number for the built-in SMTP server | [optional] 
**SmtpServerProxy** | Pointer to **bool** | Enable PROXY protocol support for the SMTP server | [optional] 
**SmtpServerTLSEnabled** | Pointer to **bool** | Enable TLS/STARTTLS support for the SMTP server | [optional] 
**SmtpStrategy** | Pointer to [**SmtpStrategy**](SmtpStrategy.md) |  | [optional] 
**TemplateHeader** | Pointer to **string** | Custom HTML to inject at the top of hosted pages (e.g., for branding) | [optional] 
**TemplateHtmlHead** | Pointer to **string** | Custom HTML to inject into the &lt;head&gt; section of hosted pages (e.g., for analytics) | [optional] 
**Timezone** | Pointer to **string** | Default timezone for date/time display (IANA timezone identifier) | [optional] 
**TrackClicks** | Pointer to **bool** | Rewrite links in outgoing HTML emails to track click-through rates | [optional] 
**TrackOpens** | Pointer to **bool** | Insert a tracking pixel in outgoing HTML emails to detect when messages are opened | [optional] 
**WebhookEvents** | Pointer to **[]string** | List of event types that will trigger webhook notifications | [optional] 
**Webhooks** | Pointer to **string** | Target URL that will receive webhook notifications via POST requests | [optional] 
**WebhooksCustomHeaders** | Pointer to [**[]WebhooksCustomHeader**](WebhooksCustomHeader.md) | Additional HTTP headers to include with every webhook request for authentication or tracking | [optional] 
**WebhooksEnabled** | Pointer to **bool** | Enable or disable webhook delivery for all accounts | [optional] 

## Methods

### NewSettingsQueryResponse

`func NewSettingsQueryResponse() *SettingsQueryResponse`

NewSettingsQueryResponse instantiates a new SettingsQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsQueryResponseWithDefaults

`func NewSettingsQueryResponseWithDefaults() *SettingsQueryResponse`

NewSettingsQueryResponseWithDefaults instantiates a new SettingsQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthServer

`func (o *SettingsQueryResponse) GetAuthServer() string`

GetAuthServer returns the AuthServer field if non-nil, zero value otherwise.

### GetAuthServerOk

`func (o *SettingsQueryResponse) GetAuthServerOk() (*string, bool)`

GetAuthServerOk returns a tuple with the AuthServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServer

`func (o *SettingsQueryResponse) SetAuthServer(v string)`

SetAuthServer sets AuthServer field to given value.

### HasAuthServer

`func (o *SettingsQueryResponse) HasAuthServer() bool`

HasAuthServer returns a boolean if a field has been set.

### GetDeliveryAttempts

`func (o *SettingsQueryResponse) GetDeliveryAttempts() int32`

GetDeliveryAttempts returns the DeliveryAttempts field if non-nil, zero value otherwise.

### GetDeliveryAttemptsOk

`func (o *SettingsQueryResponse) GetDeliveryAttemptsOk() (*int32, bool)`

GetDeliveryAttemptsOk returns a tuple with the DeliveryAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAttempts

`func (o *SettingsQueryResponse) SetDeliveryAttempts(v int32)`

SetDeliveryAttempts sets DeliveryAttempts field to given value.

### HasDeliveryAttempts

`func (o *SettingsQueryResponse) HasDeliveryAttempts() bool`

HasDeliveryAttempts returns a boolean if a field has been set.

### GetEnableApiProxy

`func (o *SettingsQueryResponse) GetEnableApiProxy() bool`

GetEnableApiProxy returns the EnableApiProxy field if non-nil, zero value otherwise.

### GetEnableApiProxyOk

`func (o *SettingsQueryResponse) GetEnableApiProxyOk() (*bool, bool)`

GetEnableApiProxyOk returns a tuple with the EnableApiProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableApiProxy

`func (o *SettingsQueryResponse) SetEnableApiProxy(v bool)`

SetEnableApiProxy sets EnableApiProxy field to given value.

### HasEnableApiProxy

`func (o *SettingsQueryResponse) HasEnableApiProxy() bool`

HasEnableApiProxy returns a boolean if a field has been set.

### GetGenerateEmailSummary

`func (o *SettingsQueryResponse) GetGenerateEmailSummary() bool`

GetGenerateEmailSummary returns the GenerateEmailSummary field if non-nil, zero value otherwise.

### GetGenerateEmailSummaryOk

`func (o *SettingsQueryResponse) GetGenerateEmailSummaryOk() (*bool, bool)`

GetGenerateEmailSummaryOk returns a tuple with the GenerateEmailSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateEmailSummary

`func (o *SettingsQueryResponse) SetGenerateEmailSummary(v bool)`

SetGenerateEmailSummary sets GenerateEmailSummary field to given value.

### HasGenerateEmailSummary

`func (o *SettingsQueryResponse) HasGenerateEmailSummary() bool`

HasGenerateEmailSummary returns a boolean if a field has been set.

### GetIgnoreMailCertErrors

`func (o *SettingsQueryResponse) GetIgnoreMailCertErrors() bool`

GetIgnoreMailCertErrors returns the IgnoreMailCertErrors field if non-nil, zero value otherwise.

### GetIgnoreMailCertErrorsOk

`func (o *SettingsQueryResponse) GetIgnoreMailCertErrorsOk() (*bool, bool)`

GetIgnoreMailCertErrorsOk returns a tuple with the IgnoreMailCertErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMailCertErrors

`func (o *SettingsQueryResponse) SetIgnoreMailCertErrors(v bool)`

SetIgnoreMailCertErrors sets IgnoreMailCertErrors field to given value.

### HasIgnoreMailCertErrors

`func (o *SettingsQueryResponse) HasIgnoreMailCertErrors() bool`

HasIgnoreMailCertErrors returns a boolean if a field has been set.

### GetImapClientName

`func (o *SettingsQueryResponse) GetImapClientName() string`

GetImapClientName returns the ImapClientName field if non-nil, zero value otherwise.

### GetImapClientNameOk

`func (o *SettingsQueryResponse) GetImapClientNameOk() (*string, bool)`

GetImapClientNameOk returns a tuple with the ImapClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapClientName

`func (o *SettingsQueryResponse) SetImapClientName(v string)`

SetImapClientName sets ImapClientName field to given value.

### HasImapClientName

`func (o *SettingsQueryResponse) HasImapClientName() bool`

HasImapClientName returns a boolean if a field has been set.

### GetImapClientSupportUrl

`func (o *SettingsQueryResponse) GetImapClientSupportUrl() string`

GetImapClientSupportUrl returns the ImapClientSupportUrl field if non-nil, zero value otherwise.

### GetImapClientSupportUrlOk

`func (o *SettingsQueryResponse) GetImapClientSupportUrlOk() (*string, bool)`

GetImapClientSupportUrlOk returns a tuple with the ImapClientSupportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapClientSupportUrl

`func (o *SettingsQueryResponse) SetImapClientSupportUrl(v string)`

SetImapClientSupportUrl sets ImapClientSupportUrl field to given value.

### HasImapClientSupportUrl

`func (o *SettingsQueryResponse) HasImapClientSupportUrl() bool`

HasImapClientSupportUrl returns a boolean if a field has been set.

### GetImapClientVendor

`func (o *SettingsQueryResponse) GetImapClientVendor() string`

GetImapClientVendor returns the ImapClientVendor field if non-nil, zero value otherwise.

### GetImapClientVendorOk

`func (o *SettingsQueryResponse) GetImapClientVendorOk() (*string, bool)`

GetImapClientVendorOk returns a tuple with the ImapClientVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapClientVendor

`func (o *SettingsQueryResponse) SetImapClientVendor(v string)`

SetImapClientVendor sets ImapClientVendor field to given value.

### HasImapClientVendor

`func (o *SettingsQueryResponse) HasImapClientVendor() bool`

HasImapClientVendor returns a boolean if a field has been set.

### GetImapClientVersion

`func (o *SettingsQueryResponse) GetImapClientVersion() string`

GetImapClientVersion returns the ImapClientVersion field if non-nil, zero value otherwise.

### GetImapClientVersionOk

`func (o *SettingsQueryResponse) GetImapClientVersionOk() (*string, bool)`

GetImapClientVersionOk returns a tuple with the ImapClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapClientVersion

`func (o *SettingsQueryResponse) SetImapClientVersion(v string)`

SetImapClientVersion sets ImapClientVersion field to given value.

### HasImapClientVersion

`func (o *SettingsQueryResponse) HasImapClientVersion() bool`

HasImapClientVersion returns a boolean if a field has been set.

### GetImapIndexer

`func (o *SettingsQueryResponse) GetImapIndexer() ImapIndexer`

GetImapIndexer returns the ImapIndexer field if non-nil, zero value otherwise.

### GetImapIndexerOk

`func (o *SettingsQueryResponse) GetImapIndexerOk() (*ImapIndexer, bool)`

GetImapIndexerOk returns a tuple with the ImapIndexer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapIndexer

`func (o *SettingsQueryResponse) SetImapIndexer(v ImapIndexer)`

SetImapIndexer sets ImapIndexer field to given value.

### HasImapIndexer

`func (o *SettingsQueryResponse) HasImapIndexer() bool`

HasImapIndexer returns a boolean if a field has been set.

### GetImapProxyServerEnabled

`func (o *SettingsQueryResponse) GetImapProxyServerEnabled() bool`

GetImapProxyServerEnabled returns the ImapProxyServerEnabled field if non-nil, zero value otherwise.

### GetImapProxyServerEnabledOk

`func (o *SettingsQueryResponse) GetImapProxyServerEnabledOk() (*bool, bool)`

GetImapProxyServerEnabledOk returns a tuple with the ImapProxyServerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapProxyServerEnabled

`func (o *SettingsQueryResponse) SetImapProxyServerEnabled(v bool)`

SetImapProxyServerEnabled sets ImapProxyServerEnabled field to given value.

### HasImapProxyServerEnabled

`func (o *SettingsQueryResponse) HasImapProxyServerEnabled() bool`

HasImapProxyServerEnabled returns a boolean if a field has been set.

### GetImapProxyServerHost

`func (o *SettingsQueryResponse) GetImapProxyServerHost() string`

GetImapProxyServerHost returns the ImapProxyServerHost field if non-nil, zero value otherwise.

### GetImapProxyServerHostOk

`func (o *SettingsQueryResponse) GetImapProxyServerHostOk() (*string, bool)`

GetImapProxyServerHostOk returns a tuple with the ImapProxyServerHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapProxyServerHost

`func (o *SettingsQueryResponse) SetImapProxyServerHost(v string)`

SetImapProxyServerHost sets ImapProxyServerHost field to given value.

### HasImapProxyServerHost

`func (o *SettingsQueryResponse) HasImapProxyServerHost() bool`

HasImapProxyServerHost returns a boolean if a field has been set.

### GetImapProxyServerPassword

`func (o *SettingsQueryResponse) GetImapProxyServerPassword() string`

GetImapProxyServerPassword returns the ImapProxyServerPassword field if non-nil, zero value otherwise.

### GetImapProxyServerPasswordOk

`func (o *SettingsQueryResponse) GetImapProxyServerPasswordOk() (*string, bool)`

GetImapProxyServerPasswordOk returns a tuple with the ImapProxyServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapProxyServerPassword

`func (o *SettingsQueryResponse) SetImapProxyServerPassword(v string)`

SetImapProxyServerPassword sets ImapProxyServerPassword field to given value.

### HasImapProxyServerPassword

`func (o *SettingsQueryResponse) HasImapProxyServerPassword() bool`

HasImapProxyServerPassword returns a boolean if a field has been set.

### SetImapProxyServerPasswordNil

`func (o *SettingsQueryResponse) SetImapProxyServerPasswordNil(b bool)`

 SetImapProxyServerPasswordNil sets the value for ImapProxyServerPassword to be an explicit nil

### UnsetImapProxyServerPassword
`func (o *SettingsQueryResponse) UnsetImapProxyServerPassword()`

UnsetImapProxyServerPassword ensures that no value is present for ImapProxyServerPassword, not even an explicit nil
### GetImapProxyServerPort

`func (o *SettingsQueryResponse) GetImapProxyServerPort() int32`

GetImapProxyServerPort returns the ImapProxyServerPort field if non-nil, zero value otherwise.

### GetImapProxyServerPortOk

`func (o *SettingsQueryResponse) GetImapProxyServerPortOk() (*int32, bool)`

GetImapProxyServerPortOk returns a tuple with the ImapProxyServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapProxyServerPort

`func (o *SettingsQueryResponse) SetImapProxyServerPort(v int32)`

SetImapProxyServerPort sets ImapProxyServerPort field to given value.

### HasImapProxyServerPort

`func (o *SettingsQueryResponse) HasImapProxyServerPort() bool`

HasImapProxyServerPort returns a boolean if a field has been set.

### GetImapProxyServerProxy

`func (o *SettingsQueryResponse) GetImapProxyServerProxy() bool`

GetImapProxyServerProxy returns the ImapProxyServerProxy field if non-nil, zero value otherwise.

### GetImapProxyServerProxyOk

`func (o *SettingsQueryResponse) GetImapProxyServerProxyOk() (*bool, bool)`

GetImapProxyServerProxyOk returns a tuple with the ImapProxyServerProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapProxyServerProxy

`func (o *SettingsQueryResponse) SetImapProxyServerProxy(v bool)`

SetImapProxyServerProxy sets ImapProxyServerProxy field to given value.

### HasImapProxyServerProxy

`func (o *SettingsQueryResponse) HasImapProxyServerProxy() bool`

HasImapProxyServerProxy returns a boolean if a field has been set.

### GetImapProxyServerTLSEnabled

`func (o *SettingsQueryResponse) GetImapProxyServerTLSEnabled() bool`

GetImapProxyServerTLSEnabled returns the ImapProxyServerTLSEnabled field if non-nil, zero value otherwise.

### GetImapProxyServerTLSEnabledOk

`func (o *SettingsQueryResponse) GetImapProxyServerTLSEnabledOk() (*bool, bool)`

GetImapProxyServerTLSEnabledOk returns a tuple with the ImapProxyServerTLSEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapProxyServerTLSEnabled

`func (o *SettingsQueryResponse) SetImapProxyServerTLSEnabled(v bool)`

SetImapProxyServerTLSEnabled sets ImapProxyServerTLSEnabled field to given value.

### HasImapProxyServerTLSEnabled

`func (o *SettingsQueryResponse) HasImapProxyServerTLSEnabled() bool`

HasImapProxyServerTLSEnabled returns a boolean if a field has been set.

### GetImapStrategy

`func (o *SettingsQueryResponse) GetImapStrategy() ImapStrategy`

GetImapStrategy returns the ImapStrategy field if non-nil, zero value otherwise.

### GetImapStrategyOk

`func (o *SettingsQueryResponse) GetImapStrategyOk() (*ImapStrategy, bool)`

GetImapStrategyOk returns a tuple with the ImapStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapStrategy

`func (o *SettingsQueryResponse) SetImapStrategy(v ImapStrategy)`

SetImapStrategy sets ImapStrategy field to given value.

### HasImapStrategy

`func (o *SettingsQueryResponse) HasImapStrategy() bool`

HasImapStrategy returns a boolean if a field has been set.

### GetInboxNewOnly

`func (o *SettingsQueryResponse) GetInboxNewOnly() bool`

GetInboxNewOnly returns the InboxNewOnly field if non-nil, zero value otherwise.

### GetInboxNewOnlyOk

`func (o *SettingsQueryResponse) GetInboxNewOnlyOk() (*bool, bool)`

GetInboxNewOnlyOk returns a tuple with the InboxNewOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboxNewOnly

`func (o *SettingsQueryResponse) SetInboxNewOnly(v bool)`

SetInboxNewOnly sets InboxNewOnly field to given value.

### HasInboxNewOnly

`func (o *SettingsQueryResponse) HasInboxNewOnly() bool`

HasInboxNewOnly returns a boolean if a field has been set.

### GetLocalAddresses

`func (o *SettingsQueryResponse) GetLocalAddresses() []string`

GetLocalAddresses returns the LocalAddresses field if non-nil, zero value otherwise.

### GetLocalAddressesOk

`func (o *SettingsQueryResponse) GetLocalAddressesOk() (*[]string, bool)`

GetLocalAddressesOk returns a tuple with the LocalAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddresses

`func (o *SettingsQueryResponse) SetLocalAddresses(v []string)`

SetLocalAddresses sets LocalAddresses field to given value.

### HasLocalAddresses

`func (o *SettingsQueryResponse) HasLocalAddresses() bool`

HasLocalAddresses returns a boolean if a field has been set.

### GetLocale

`func (o *SettingsQueryResponse) GetLocale() Locale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *SettingsQueryResponse) GetLocaleOk() (*Locale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *SettingsQueryResponse) SetLocale(v Locale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *SettingsQueryResponse) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetLogs

`func (o *SettingsQueryResponse) GetLogs() LogSettings`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *SettingsQueryResponse) GetLogsOk() (*LogSettings, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *SettingsQueryResponse) SetLogs(v LogSettings)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *SettingsQueryResponse) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetNotificationBaseUrl

`func (o *SettingsQueryResponse) GetNotificationBaseUrl() string`

GetNotificationBaseUrl returns the NotificationBaseUrl field if non-nil, zero value otherwise.

### GetNotificationBaseUrlOk

`func (o *SettingsQueryResponse) GetNotificationBaseUrlOk() (*string, bool)`

GetNotificationBaseUrlOk returns a tuple with the NotificationBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationBaseUrl

`func (o *SettingsQueryResponse) SetNotificationBaseUrl(v string)`

SetNotificationBaseUrl sets NotificationBaseUrl field to given value.

### HasNotificationBaseUrl

`func (o *SettingsQueryResponse) HasNotificationBaseUrl() bool`

HasNotificationBaseUrl returns a boolean if a field has been set.

### SetNotificationBaseUrlNil

`func (o *SettingsQueryResponse) SetNotificationBaseUrlNil(b bool)`

 SetNotificationBaseUrlNil sets the value for NotificationBaseUrl to be an explicit nil

### UnsetNotificationBaseUrl
`func (o *SettingsQueryResponse) UnsetNotificationBaseUrl()`

UnsetNotificationBaseUrl ensures that no value is present for NotificationBaseUrl, not even an explicit nil
### GetNotifyAttachmentSize

`func (o *SettingsQueryResponse) GetNotifyAttachmentSize() int32`

GetNotifyAttachmentSize returns the NotifyAttachmentSize field if non-nil, zero value otherwise.

### GetNotifyAttachmentSizeOk

`func (o *SettingsQueryResponse) GetNotifyAttachmentSizeOk() (*int32, bool)`

GetNotifyAttachmentSizeOk returns a tuple with the NotifyAttachmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyAttachmentSize

`func (o *SettingsQueryResponse) SetNotifyAttachmentSize(v int32)`

SetNotifyAttachmentSize sets NotifyAttachmentSize field to given value.

### HasNotifyAttachmentSize

`func (o *SettingsQueryResponse) HasNotifyAttachmentSize() bool`

HasNotifyAttachmentSize returns a boolean if a field has been set.

### GetNotifyAttachments

`func (o *SettingsQueryResponse) GetNotifyAttachments() bool`

GetNotifyAttachments returns the NotifyAttachments field if non-nil, zero value otherwise.

### GetNotifyAttachmentsOk

`func (o *SettingsQueryResponse) GetNotifyAttachmentsOk() (*bool, bool)`

GetNotifyAttachmentsOk returns a tuple with the NotifyAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyAttachments

`func (o *SettingsQueryResponse) SetNotifyAttachments(v bool)`

SetNotifyAttachments sets NotifyAttachments field to given value.

### HasNotifyAttachments

`func (o *SettingsQueryResponse) HasNotifyAttachments() bool`

HasNotifyAttachments returns a boolean if a field has been set.

### GetNotifyCalendarEvents

`func (o *SettingsQueryResponse) GetNotifyCalendarEvents() bool`

GetNotifyCalendarEvents returns the NotifyCalendarEvents field if non-nil, zero value otherwise.

### GetNotifyCalendarEventsOk

`func (o *SettingsQueryResponse) GetNotifyCalendarEventsOk() (*bool, bool)`

GetNotifyCalendarEventsOk returns a tuple with the NotifyCalendarEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCalendarEvents

`func (o *SettingsQueryResponse) SetNotifyCalendarEvents(v bool)`

SetNotifyCalendarEvents sets NotifyCalendarEvents field to given value.

### HasNotifyCalendarEvents

`func (o *SettingsQueryResponse) HasNotifyCalendarEvents() bool`

HasNotifyCalendarEvents returns a boolean if a field has been set.

### GetNotifyHeaders

`func (o *SettingsQueryResponse) GetNotifyHeaders() []string`

GetNotifyHeaders returns the NotifyHeaders field if non-nil, zero value otherwise.

### GetNotifyHeadersOk

`func (o *SettingsQueryResponse) GetNotifyHeadersOk() (*[]string, bool)`

GetNotifyHeadersOk returns a tuple with the NotifyHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyHeaders

`func (o *SettingsQueryResponse) SetNotifyHeaders(v []string)`

SetNotifyHeaders sets NotifyHeaders field to given value.

### HasNotifyHeaders

`func (o *SettingsQueryResponse) HasNotifyHeaders() bool`

HasNotifyHeaders returns a boolean if a field has been set.

### GetNotifyText

`func (o *SettingsQueryResponse) GetNotifyText() bool`

GetNotifyText returns the NotifyText field if non-nil, zero value otherwise.

### GetNotifyTextOk

`func (o *SettingsQueryResponse) GetNotifyTextOk() (*bool, bool)`

GetNotifyTextOk returns a tuple with the NotifyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyText

`func (o *SettingsQueryResponse) SetNotifyText(v bool)`

SetNotifyText sets NotifyText field to given value.

### HasNotifyText

`func (o *SettingsQueryResponse) HasNotifyText() bool`

HasNotifyText returns a boolean if a field has been set.

### GetNotifyTextSize

`func (o *SettingsQueryResponse) GetNotifyTextSize() int32`

GetNotifyTextSize returns the NotifyTextSize field if non-nil, zero value otherwise.

### GetNotifyTextSizeOk

`func (o *SettingsQueryResponse) GetNotifyTextSizeOk() (*int32, bool)`

GetNotifyTextSizeOk returns a tuple with the NotifyTextSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTextSize

`func (o *SettingsQueryResponse) SetNotifyTextSize(v int32)`

SetNotifyTextSize sets NotifyTextSize field to given value.

### HasNotifyTextSize

`func (o *SettingsQueryResponse) HasNotifyTextSize() bool`

HasNotifyTextSize returns a boolean if a field has been set.

### GetNotifyWebSafeHtml

`func (o *SettingsQueryResponse) GetNotifyWebSafeHtml() bool`

GetNotifyWebSafeHtml returns the NotifyWebSafeHtml field if non-nil, zero value otherwise.

### GetNotifyWebSafeHtmlOk

`func (o *SettingsQueryResponse) GetNotifyWebSafeHtmlOk() (*bool, bool)`

GetNotifyWebSafeHtmlOk returns a tuple with the NotifyWebSafeHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyWebSafeHtml

`func (o *SettingsQueryResponse) SetNotifyWebSafeHtml(v bool)`

SetNotifyWebSafeHtml sets NotifyWebSafeHtml field to given value.

### HasNotifyWebSafeHtml

`func (o *SettingsQueryResponse) HasNotifyWebSafeHtml() bool`

HasNotifyWebSafeHtml returns a boolean if a field has been set.

### GetOpenAiAPIKey

`func (o *SettingsQueryResponse) GetOpenAiAPIKey() string`

GetOpenAiAPIKey returns the OpenAiAPIKey field if non-nil, zero value otherwise.

### GetOpenAiAPIKeyOk

`func (o *SettingsQueryResponse) GetOpenAiAPIKeyOk() (*string, bool)`

GetOpenAiAPIKeyOk returns a tuple with the OpenAiAPIKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAiAPIKey

`func (o *SettingsQueryResponse) SetOpenAiAPIKey(v string)`

SetOpenAiAPIKey sets OpenAiAPIKey field to given value.

### HasOpenAiAPIKey

`func (o *SettingsQueryResponse) HasOpenAiAPIKey() bool`

HasOpenAiAPIKey returns a boolean if a field has been set.

### GetOpenAiAPIUrl

`func (o *SettingsQueryResponse) GetOpenAiAPIUrl() string`

GetOpenAiAPIUrl returns the OpenAiAPIUrl field if non-nil, zero value otherwise.

### GetOpenAiAPIUrlOk

`func (o *SettingsQueryResponse) GetOpenAiAPIUrlOk() (*string, bool)`

GetOpenAiAPIUrlOk returns a tuple with the OpenAiAPIUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAiAPIUrl

`func (o *SettingsQueryResponse) SetOpenAiAPIUrl(v string)`

SetOpenAiAPIUrl sets OpenAiAPIUrl field to given value.

### HasOpenAiAPIUrl

`func (o *SettingsQueryResponse) HasOpenAiAPIUrl() bool`

HasOpenAiAPIUrl returns a boolean if a field has been set.

### GetOpenAiGenerateEmbeddings

`func (o *SettingsQueryResponse) GetOpenAiGenerateEmbeddings() bool`

GetOpenAiGenerateEmbeddings returns the OpenAiGenerateEmbeddings field if non-nil, zero value otherwise.

### GetOpenAiGenerateEmbeddingsOk

`func (o *SettingsQueryResponse) GetOpenAiGenerateEmbeddingsOk() (*bool, bool)`

GetOpenAiGenerateEmbeddingsOk returns a tuple with the OpenAiGenerateEmbeddings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAiGenerateEmbeddings

`func (o *SettingsQueryResponse) SetOpenAiGenerateEmbeddings(v bool)`

SetOpenAiGenerateEmbeddings sets OpenAiGenerateEmbeddings field to given value.

### HasOpenAiGenerateEmbeddings

`func (o *SettingsQueryResponse) HasOpenAiGenerateEmbeddings() bool`

HasOpenAiGenerateEmbeddings returns a boolean if a field has been set.

### GetOpenAiMaxTokens

`func (o *SettingsQueryResponse) GetOpenAiMaxTokens() float32`

GetOpenAiMaxTokens returns the OpenAiMaxTokens field if non-nil, zero value otherwise.

### GetOpenAiMaxTokensOk

`func (o *SettingsQueryResponse) GetOpenAiMaxTokensOk() (*float32, bool)`

GetOpenAiMaxTokensOk returns a tuple with the OpenAiMaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAiMaxTokens

`func (o *SettingsQueryResponse) SetOpenAiMaxTokens(v float32)`

SetOpenAiMaxTokens sets OpenAiMaxTokens field to given value.

### HasOpenAiMaxTokens

`func (o *SettingsQueryResponse) HasOpenAiMaxTokens() bool`

HasOpenAiMaxTokens returns a boolean if a field has been set.

### GetOpenAiModel

`func (o *SettingsQueryResponse) GetOpenAiModel() string`

GetOpenAiModel returns the OpenAiModel field if non-nil, zero value otherwise.

### GetOpenAiModelOk

`func (o *SettingsQueryResponse) GetOpenAiModelOk() (*string, bool)`

GetOpenAiModelOk returns a tuple with the OpenAiModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAiModel

`func (o *SettingsQueryResponse) SetOpenAiModel(v string)`

SetOpenAiModel sets OpenAiModel field to given value.

### HasOpenAiModel

`func (o *SettingsQueryResponse) HasOpenAiModel() bool`

HasOpenAiModel returns a boolean if a field has been set.

### GetOpenAiPreProcessingFn

`func (o *SettingsQueryResponse) GetOpenAiPreProcessingFn() string`

GetOpenAiPreProcessingFn returns the OpenAiPreProcessingFn field if non-nil, zero value otherwise.

### GetOpenAiPreProcessingFnOk

`func (o *SettingsQueryResponse) GetOpenAiPreProcessingFnOk() (*string, bool)`

GetOpenAiPreProcessingFnOk returns a tuple with the OpenAiPreProcessingFn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAiPreProcessingFn

`func (o *SettingsQueryResponse) SetOpenAiPreProcessingFn(v string)`

SetOpenAiPreProcessingFn sets OpenAiPreProcessingFn field to given value.

### HasOpenAiPreProcessingFn

`func (o *SettingsQueryResponse) HasOpenAiPreProcessingFn() bool`

HasOpenAiPreProcessingFn returns a boolean if a field has been set.

### GetOpenAiPrompt

`func (o *SettingsQueryResponse) GetOpenAiPrompt() string`

GetOpenAiPrompt returns the OpenAiPrompt field if non-nil, zero value otherwise.

### GetOpenAiPromptOk

`func (o *SettingsQueryResponse) GetOpenAiPromptOk() (*string, bool)`

GetOpenAiPromptOk returns a tuple with the OpenAiPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAiPrompt

`func (o *SettingsQueryResponse) SetOpenAiPrompt(v string)`

SetOpenAiPrompt sets OpenAiPrompt field to given value.

### HasOpenAiPrompt

`func (o *SettingsQueryResponse) HasOpenAiPrompt() bool`

HasOpenAiPrompt returns a boolean if a field has been set.

### GetOpenAiTemperature

`func (o *SettingsQueryResponse) GetOpenAiTemperature() float32`

GetOpenAiTemperature returns the OpenAiTemperature field if non-nil, zero value otherwise.

### GetOpenAiTemperatureOk

`func (o *SettingsQueryResponse) GetOpenAiTemperatureOk() (*float32, bool)`

GetOpenAiTemperatureOk returns a tuple with the OpenAiTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAiTemperature

`func (o *SettingsQueryResponse) SetOpenAiTemperature(v float32)`

SetOpenAiTemperature sets OpenAiTemperature field to given value.

### HasOpenAiTemperature

`func (o *SettingsQueryResponse) HasOpenAiTemperature() bool`

HasOpenAiTemperature returns a boolean if a field has been set.

### GetOpenAiTopP

`func (o *SettingsQueryResponse) GetOpenAiTopP() float32`

GetOpenAiTopP returns the OpenAiTopP field if non-nil, zero value otherwise.

### GetOpenAiTopPOk

`func (o *SettingsQueryResponse) GetOpenAiTopPOk() (*float32, bool)`

GetOpenAiTopPOk returns a tuple with the OpenAiTopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAiTopP

`func (o *SettingsQueryResponse) SetOpenAiTopP(v float32)`

SetOpenAiTopP sets OpenAiTopP field to given value.

### HasOpenAiTopP

`func (o *SettingsQueryResponse) HasOpenAiTopP() bool`

HasOpenAiTopP returns a boolean if a field has been set.

### GetPageBrandName

`func (o *SettingsQueryResponse) GetPageBrandName() string`

GetPageBrandName returns the PageBrandName field if non-nil, zero value otherwise.

### GetPageBrandNameOk

`func (o *SettingsQueryResponse) GetPageBrandNameOk() (*string, bool)`

GetPageBrandNameOk returns a tuple with the PageBrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageBrandName

`func (o *SettingsQueryResponse) SetPageBrandName(v string)`

SetPageBrandName sets PageBrandName field to given value.

### HasPageBrandName

`func (o *SettingsQueryResponse) HasPageBrandName() bool`

HasPageBrandName returns a boolean if a field has been set.

### SetPageBrandNameNil

`func (o *SettingsQueryResponse) SetPageBrandNameNil(b bool)`

 SetPageBrandNameNil sets the value for PageBrandName to be an explicit nil

### UnsetPageBrandName
`func (o *SettingsQueryResponse) UnsetPageBrandName()`

UnsetPageBrandName ensures that no value is present for PageBrandName, not even an explicit nil
### GetProxyEnabled

`func (o *SettingsQueryResponse) GetProxyEnabled() bool`

GetProxyEnabled returns the ProxyEnabled field if non-nil, zero value otherwise.

### GetProxyEnabledOk

`func (o *SettingsQueryResponse) GetProxyEnabledOk() (*bool, bool)`

GetProxyEnabledOk returns a tuple with the ProxyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyEnabled

`func (o *SettingsQueryResponse) SetProxyEnabled(v bool)`

SetProxyEnabled sets ProxyEnabled field to given value.

### HasProxyEnabled

`func (o *SettingsQueryResponse) HasProxyEnabled() bool`

HasProxyEnabled returns a boolean if a field has been set.

### GetProxyUrl

`func (o *SettingsQueryResponse) GetProxyUrl() string`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *SettingsQueryResponse) GetProxyUrlOk() (*string, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *SettingsQueryResponse) SetProxyUrl(v string)`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *SettingsQueryResponse) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.

### GetQueueKeep

`func (o *SettingsQueryResponse) GetQueueKeep() int32`

GetQueueKeep returns the QueueKeep field if non-nil, zero value otherwise.

### GetQueueKeepOk

`func (o *SettingsQueryResponse) GetQueueKeepOk() (*int32, bool)`

GetQueueKeepOk returns a tuple with the QueueKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueKeep

`func (o *SettingsQueryResponse) SetQueueKeep(v int32)`

SetQueueKeep sets QueueKeep field to given value.

### HasQueueKeep

`func (o *SettingsQueryResponse) HasQueueKeep() bool`

HasQueueKeep returns a boolean if a field has been set.

### GetResolveGmailCategories

`func (o *SettingsQueryResponse) GetResolveGmailCategories() bool`

GetResolveGmailCategories returns the ResolveGmailCategories field if non-nil, zero value otherwise.

### GetResolveGmailCategoriesOk

`func (o *SettingsQueryResponse) GetResolveGmailCategoriesOk() (*bool, bool)`

GetResolveGmailCategoriesOk returns a tuple with the ResolveGmailCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveGmailCategories

`func (o *SettingsQueryResponse) SetResolveGmailCategories(v bool)`

SetResolveGmailCategories sets ResolveGmailCategories field to given value.

### HasResolveGmailCategories

`func (o *SettingsQueryResponse) HasResolveGmailCategories() bool`

HasResolveGmailCategories returns a boolean if a field has been set.

### GetScriptEnv

`func (o *SettingsQueryResponse) GetScriptEnv() string`

GetScriptEnv returns the ScriptEnv field if non-nil, zero value otherwise.

### GetScriptEnvOk

`func (o *SettingsQueryResponse) GetScriptEnvOk() (*string, bool)`

GetScriptEnvOk returns a tuple with the ScriptEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptEnv

`func (o *SettingsQueryResponse) SetScriptEnv(v string)`

SetScriptEnv sets ScriptEnv field to given value.

### HasScriptEnv

`func (o *SettingsQueryResponse) HasScriptEnv() bool`

HasScriptEnv returns a boolean if a field has been set.

### GetServiceSecret

`func (o *SettingsQueryResponse) GetServiceSecret() string`

GetServiceSecret returns the ServiceSecret field if non-nil, zero value otherwise.

### GetServiceSecretOk

`func (o *SettingsQueryResponse) GetServiceSecretOk() (*string, bool)`

GetServiceSecretOk returns a tuple with the ServiceSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSecret

`func (o *SettingsQueryResponse) SetServiceSecret(v string)`

SetServiceSecret sets ServiceSecret field to given value.

### HasServiceSecret

`func (o *SettingsQueryResponse) HasServiceSecret() bool`

HasServiceSecret returns a boolean if a field has been set.

### GetServiceUrl

`func (o *SettingsQueryResponse) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *SettingsQueryResponse) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *SettingsQueryResponse) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *SettingsQueryResponse) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetSmtpEhloName

`func (o *SettingsQueryResponse) GetSmtpEhloName() string`

GetSmtpEhloName returns the SmtpEhloName field if non-nil, zero value otherwise.

### GetSmtpEhloNameOk

`func (o *SettingsQueryResponse) GetSmtpEhloNameOk() (*string, bool)`

GetSmtpEhloNameOk returns a tuple with the SmtpEhloName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpEhloName

`func (o *SettingsQueryResponse) SetSmtpEhloName(v string)`

SetSmtpEhloName sets SmtpEhloName field to given value.

### HasSmtpEhloName

`func (o *SettingsQueryResponse) HasSmtpEhloName() bool`

HasSmtpEhloName returns a boolean if a field has been set.

### SetSmtpEhloNameNil

`func (o *SettingsQueryResponse) SetSmtpEhloNameNil(b bool)`

 SetSmtpEhloNameNil sets the value for SmtpEhloName to be an explicit nil

### UnsetSmtpEhloName
`func (o *SettingsQueryResponse) UnsetSmtpEhloName()`

UnsetSmtpEhloName ensures that no value is present for SmtpEhloName, not even an explicit nil
### GetSmtpServerAuthEnabled

`func (o *SettingsQueryResponse) GetSmtpServerAuthEnabled() bool`

GetSmtpServerAuthEnabled returns the SmtpServerAuthEnabled field if non-nil, zero value otherwise.

### GetSmtpServerAuthEnabledOk

`func (o *SettingsQueryResponse) GetSmtpServerAuthEnabledOk() (*bool, bool)`

GetSmtpServerAuthEnabledOk returns a tuple with the SmtpServerAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServerAuthEnabled

`func (o *SettingsQueryResponse) SetSmtpServerAuthEnabled(v bool)`

SetSmtpServerAuthEnabled sets SmtpServerAuthEnabled field to given value.

### HasSmtpServerAuthEnabled

`func (o *SettingsQueryResponse) HasSmtpServerAuthEnabled() bool`

HasSmtpServerAuthEnabled returns a boolean if a field has been set.

### GetSmtpServerEnabled

`func (o *SettingsQueryResponse) GetSmtpServerEnabled() bool`

GetSmtpServerEnabled returns the SmtpServerEnabled field if non-nil, zero value otherwise.

### GetSmtpServerEnabledOk

`func (o *SettingsQueryResponse) GetSmtpServerEnabledOk() (*bool, bool)`

GetSmtpServerEnabledOk returns a tuple with the SmtpServerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServerEnabled

`func (o *SettingsQueryResponse) SetSmtpServerEnabled(v bool)`

SetSmtpServerEnabled sets SmtpServerEnabled field to given value.

### HasSmtpServerEnabled

`func (o *SettingsQueryResponse) HasSmtpServerEnabled() bool`

HasSmtpServerEnabled returns a boolean if a field has been set.

### GetSmtpServerHost

`func (o *SettingsQueryResponse) GetSmtpServerHost() string`

GetSmtpServerHost returns the SmtpServerHost field if non-nil, zero value otherwise.

### GetSmtpServerHostOk

`func (o *SettingsQueryResponse) GetSmtpServerHostOk() (*string, bool)`

GetSmtpServerHostOk returns a tuple with the SmtpServerHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServerHost

`func (o *SettingsQueryResponse) SetSmtpServerHost(v string)`

SetSmtpServerHost sets SmtpServerHost field to given value.

### HasSmtpServerHost

`func (o *SettingsQueryResponse) HasSmtpServerHost() bool`

HasSmtpServerHost returns a boolean if a field has been set.

### GetSmtpServerPassword

`func (o *SettingsQueryResponse) GetSmtpServerPassword() string`

GetSmtpServerPassword returns the SmtpServerPassword field if non-nil, zero value otherwise.

### GetSmtpServerPasswordOk

`func (o *SettingsQueryResponse) GetSmtpServerPasswordOk() (*string, bool)`

GetSmtpServerPasswordOk returns a tuple with the SmtpServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServerPassword

`func (o *SettingsQueryResponse) SetSmtpServerPassword(v string)`

SetSmtpServerPassword sets SmtpServerPassword field to given value.

### HasSmtpServerPassword

`func (o *SettingsQueryResponse) HasSmtpServerPassword() bool`

HasSmtpServerPassword returns a boolean if a field has been set.

### SetSmtpServerPasswordNil

`func (o *SettingsQueryResponse) SetSmtpServerPasswordNil(b bool)`

 SetSmtpServerPasswordNil sets the value for SmtpServerPassword to be an explicit nil

### UnsetSmtpServerPassword
`func (o *SettingsQueryResponse) UnsetSmtpServerPassword()`

UnsetSmtpServerPassword ensures that no value is present for SmtpServerPassword, not even an explicit nil
### GetSmtpServerPort

`func (o *SettingsQueryResponse) GetSmtpServerPort() int32`

GetSmtpServerPort returns the SmtpServerPort field if non-nil, zero value otherwise.

### GetSmtpServerPortOk

`func (o *SettingsQueryResponse) GetSmtpServerPortOk() (*int32, bool)`

GetSmtpServerPortOk returns a tuple with the SmtpServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServerPort

`func (o *SettingsQueryResponse) SetSmtpServerPort(v int32)`

SetSmtpServerPort sets SmtpServerPort field to given value.

### HasSmtpServerPort

`func (o *SettingsQueryResponse) HasSmtpServerPort() bool`

HasSmtpServerPort returns a boolean if a field has been set.

### GetSmtpServerProxy

`func (o *SettingsQueryResponse) GetSmtpServerProxy() bool`

GetSmtpServerProxy returns the SmtpServerProxy field if non-nil, zero value otherwise.

### GetSmtpServerProxyOk

`func (o *SettingsQueryResponse) GetSmtpServerProxyOk() (*bool, bool)`

GetSmtpServerProxyOk returns a tuple with the SmtpServerProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServerProxy

`func (o *SettingsQueryResponse) SetSmtpServerProxy(v bool)`

SetSmtpServerProxy sets SmtpServerProxy field to given value.

### HasSmtpServerProxy

`func (o *SettingsQueryResponse) HasSmtpServerProxy() bool`

HasSmtpServerProxy returns a boolean if a field has been set.

### GetSmtpServerTLSEnabled

`func (o *SettingsQueryResponse) GetSmtpServerTLSEnabled() bool`

GetSmtpServerTLSEnabled returns the SmtpServerTLSEnabled field if non-nil, zero value otherwise.

### GetSmtpServerTLSEnabledOk

`func (o *SettingsQueryResponse) GetSmtpServerTLSEnabledOk() (*bool, bool)`

GetSmtpServerTLSEnabledOk returns a tuple with the SmtpServerTLSEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServerTLSEnabled

`func (o *SettingsQueryResponse) SetSmtpServerTLSEnabled(v bool)`

SetSmtpServerTLSEnabled sets SmtpServerTLSEnabled field to given value.

### HasSmtpServerTLSEnabled

`func (o *SettingsQueryResponse) HasSmtpServerTLSEnabled() bool`

HasSmtpServerTLSEnabled returns a boolean if a field has been set.

### GetSmtpStrategy

`func (o *SettingsQueryResponse) GetSmtpStrategy() SmtpStrategy`

GetSmtpStrategy returns the SmtpStrategy field if non-nil, zero value otherwise.

### GetSmtpStrategyOk

`func (o *SettingsQueryResponse) GetSmtpStrategyOk() (*SmtpStrategy, bool)`

GetSmtpStrategyOk returns a tuple with the SmtpStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpStrategy

`func (o *SettingsQueryResponse) SetSmtpStrategy(v SmtpStrategy)`

SetSmtpStrategy sets SmtpStrategy field to given value.

### HasSmtpStrategy

`func (o *SettingsQueryResponse) HasSmtpStrategy() bool`

HasSmtpStrategy returns a boolean if a field has been set.

### GetTemplateHeader

`func (o *SettingsQueryResponse) GetTemplateHeader() string`

GetTemplateHeader returns the TemplateHeader field if non-nil, zero value otherwise.

### GetTemplateHeaderOk

`func (o *SettingsQueryResponse) GetTemplateHeaderOk() (*string, bool)`

GetTemplateHeaderOk returns a tuple with the TemplateHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateHeader

`func (o *SettingsQueryResponse) SetTemplateHeader(v string)`

SetTemplateHeader sets TemplateHeader field to given value.

### HasTemplateHeader

`func (o *SettingsQueryResponse) HasTemplateHeader() bool`

HasTemplateHeader returns a boolean if a field has been set.

### GetTemplateHtmlHead

`func (o *SettingsQueryResponse) GetTemplateHtmlHead() string`

GetTemplateHtmlHead returns the TemplateHtmlHead field if non-nil, zero value otherwise.

### GetTemplateHtmlHeadOk

`func (o *SettingsQueryResponse) GetTemplateHtmlHeadOk() (*string, bool)`

GetTemplateHtmlHeadOk returns a tuple with the TemplateHtmlHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateHtmlHead

`func (o *SettingsQueryResponse) SetTemplateHtmlHead(v string)`

SetTemplateHtmlHead sets TemplateHtmlHead field to given value.

### HasTemplateHtmlHead

`func (o *SettingsQueryResponse) HasTemplateHtmlHead() bool`

HasTemplateHtmlHead returns a boolean if a field has been set.

### GetTimezone

`func (o *SettingsQueryResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SettingsQueryResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SettingsQueryResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *SettingsQueryResponse) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTrackClicks

`func (o *SettingsQueryResponse) GetTrackClicks() bool`

GetTrackClicks returns the TrackClicks field if non-nil, zero value otherwise.

### GetTrackClicksOk

`func (o *SettingsQueryResponse) GetTrackClicksOk() (*bool, bool)`

GetTrackClicksOk returns a tuple with the TrackClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackClicks

`func (o *SettingsQueryResponse) SetTrackClicks(v bool)`

SetTrackClicks sets TrackClicks field to given value.

### HasTrackClicks

`func (o *SettingsQueryResponse) HasTrackClicks() bool`

HasTrackClicks returns a boolean if a field has been set.

### GetTrackOpens

`func (o *SettingsQueryResponse) GetTrackOpens() bool`

GetTrackOpens returns the TrackOpens field if non-nil, zero value otherwise.

### GetTrackOpensOk

`func (o *SettingsQueryResponse) GetTrackOpensOk() (*bool, bool)`

GetTrackOpensOk returns a tuple with the TrackOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackOpens

`func (o *SettingsQueryResponse) SetTrackOpens(v bool)`

SetTrackOpens sets TrackOpens field to given value.

### HasTrackOpens

`func (o *SettingsQueryResponse) HasTrackOpens() bool`

HasTrackOpens returns a boolean if a field has been set.

### GetWebhookEvents

`func (o *SettingsQueryResponse) GetWebhookEvents() []string`

GetWebhookEvents returns the WebhookEvents field if non-nil, zero value otherwise.

### GetWebhookEventsOk

`func (o *SettingsQueryResponse) GetWebhookEventsOk() (*[]string, bool)`

GetWebhookEventsOk returns a tuple with the WebhookEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEvents

`func (o *SettingsQueryResponse) SetWebhookEvents(v []string)`

SetWebhookEvents sets WebhookEvents field to given value.

### HasWebhookEvents

`func (o *SettingsQueryResponse) HasWebhookEvents() bool`

HasWebhookEvents returns a boolean if a field has been set.

### GetWebhooks

`func (o *SettingsQueryResponse) GetWebhooks() string`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *SettingsQueryResponse) GetWebhooksOk() (*string, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *SettingsQueryResponse) SetWebhooks(v string)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *SettingsQueryResponse) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetWebhooksCustomHeaders

`func (o *SettingsQueryResponse) GetWebhooksCustomHeaders() []WebhooksCustomHeader`

GetWebhooksCustomHeaders returns the WebhooksCustomHeaders field if non-nil, zero value otherwise.

### GetWebhooksCustomHeadersOk

`func (o *SettingsQueryResponse) GetWebhooksCustomHeadersOk() (*[]WebhooksCustomHeader, bool)`

GetWebhooksCustomHeadersOk returns a tuple with the WebhooksCustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooksCustomHeaders

`func (o *SettingsQueryResponse) SetWebhooksCustomHeaders(v []WebhooksCustomHeader)`

SetWebhooksCustomHeaders sets WebhooksCustomHeaders field to given value.

### HasWebhooksCustomHeaders

`func (o *SettingsQueryResponse) HasWebhooksCustomHeaders() bool`

HasWebhooksCustomHeaders returns a boolean if a field has been set.

### GetWebhooksEnabled

`func (o *SettingsQueryResponse) GetWebhooksEnabled() bool`

GetWebhooksEnabled returns the WebhooksEnabled field if non-nil, zero value otherwise.

### GetWebhooksEnabledOk

`func (o *SettingsQueryResponse) GetWebhooksEnabledOk() (*bool, bool)`

GetWebhooksEnabledOk returns a tuple with the WebhooksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooksEnabled

`func (o *SettingsQueryResponse) SetWebhooksEnabled(v bool)`

SetWebhooksEnabled sets WebhooksEnabled field to given value.

### HasWebhooksEnabled

`func (o *SettingsQueryResponse) HasWebhooksEnabled() bool`

HasWebhooksEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


