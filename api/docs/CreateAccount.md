# CreateAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **NullableString** | Account ID. If set to &#x60;null&#x60;, a unique ID will be generated automatically. If you provide an existing account ID, the settings for that account will be updated instead | 
**Copy** | Pointer to **NullableBool** | Copy submitted messages to Sent folder. Set to &#x60;null&#x60; to unset and use provider specific default. | [optional] 
**Email** | Pointer to **string** | Default email address of the account | [optional] 
**Imap** | Pointer to [**ImapConfiguration**](ImapConfiguration.md) |  | [optional] 
**ImapIndexer** | Pointer to [**NullableIMAPIndexer**](IMAPIndexer.md) |  | [optional] 
**Locale** | Pointer to **string** | Optional locale | [optional] 
**Logs** | Pointer to **bool** | Store recent logs | [optional] [default to false]
**Name** | **string** | Display name for the account | 
**NotifyFrom** | Pointer to **NullableTime** | Only send webhooks for messages received after this date. Defaults to account creation time. IMAP only. | [optional] 
**Oauth2** | Pointer to [**OAuth2**](OAuth2.md) |  | [optional] 
**Path** | Pointer to **[]string** | Mailbox folders to monitor for changes (IMAP only). Use \&quot;*\&quot; to monitor all folders (default). While you can still access unmonitored folders via API, you won&#39;t receive webhooks for changes in those folders. | [optional] 
**Proxy** | Pointer to **string** | Proxy server URL for outbound connections | [optional] 
**Smtp** | Pointer to [**SmtpConfiguration**](SmtpConfiguration.md) |  | [optional] 
**SmtpEhloName** | Pointer to **NullableString** | Hostname to use in SMTP EHLO/HELO commands (defaults to system hostname) | [optional] 
**Subconnections** | Pointer to **[]string** | Additional mailbox paths to monitor with dedicated IMAP connections for faster change detection. Use sparingly as connection limits are strict. | [optional] 
**Tz** | Pointer to **string** | Optional timezone | [optional] 
**Webhooks** | Pointer to **string** | Account-specific webhook URL | [optional] 
**WebhooksCustomHeaders** | Pointer to [**[]WebhooksCustomHeader**](WebhooksCustomHeader.md) | Additional HTTP headers to include with every webhook request for authentication or tracking | [optional] 

## Methods

### NewCreateAccount

`func NewCreateAccount(account NullableString, name string, ) *CreateAccount`

NewCreateAccount instantiates a new CreateAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountWithDefaults

`func NewCreateAccountWithDefaults() *CreateAccount`

NewCreateAccountWithDefaults instantiates a new CreateAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CreateAccount) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateAccount) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateAccount) SetAccount(v string)`

SetAccount sets Account field to given value.


### SetAccountNil

`func (o *CreateAccount) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *CreateAccount) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetCopy

`func (o *CreateAccount) GetCopy() bool`

GetCopy returns the Copy field if non-nil, zero value otherwise.

### GetCopyOk

`func (o *CreateAccount) GetCopyOk() (*bool, bool)`

GetCopyOk returns a tuple with the Copy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopy

`func (o *CreateAccount) SetCopy(v bool)`

SetCopy sets Copy field to given value.

### HasCopy

`func (o *CreateAccount) HasCopy() bool`

HasCopy returns a boolean if a field has been set.

### SetCopyNil

`func (o *CreateAccount) SetCopyNil(b bool)`

 SetCopyNil sets the value for Copy to be an explicit nil

### UnsetCopy
`func (o *CreateAccount) UnsetCopy()`

UnsetCopy ensures that no value is present for Copy, not even an explicit nil
### GetEmail

`func (o *CreateAccount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateAccount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateAccount) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateAccount) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetImap

`func (o *CreateAccount) GetImap() ImapConfiguration`

GetImap returns the Imap field if non-nil, zero value otherwise.

### GetImapOk

`func (o *CreateAccount) GetImapOk() (*ImapConfiguration, bool)`

GetImapOk returns a tuple with the Imap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImap

`func (o *CreateAccount) SetImap(v ImapConfiguration)`

SetImap sets Imap field to given value.

### HasImap

`func (o *CreateAccount) HasImap() bool`

HasImap returns a boolean if a field has been set.

### GetImapIndexer

`func (o *CreateAccount) GetImapIndexer() IMAPIndexer`

GetImapIndexer returns the ImapIndexer field if non-nil, zero value otherwise.

### GetImapIndexerOk

`func (o *CreateAccount) GetImapIndexerOk() (*IMAPIndexer, bool)`

GetImapIndexerOk returns a tuple with the ImapIndexer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapIndexer

`func (o *CreateAccount) SetImapIndexer(v IMAPIndexer)`

SetImapIndexer sets ImapIndexer field to given value.

### HasImapIndexer

`func (o *CreateAccount) HasImapIndexer() bool`

HasImapIndexer returns a boolean if a field has been set.

### SetImapIndexerNil

`func (o *CreateAccount) SetImapIndexerNil(b bool)`

 SetImapIndexerNil sets the value for ImapIndexer to be an explicit nil

### UnsetImapIndexer
`func (o *CreateAccount) UnsetImapIndexer()`

UnsetImapIndexer ensures that no value is present for ImapIndexer, not even an explicit nil
### GetLocale

`func (o *CreateAccount) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CreateAccount) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CreateAccount) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CreateAccount) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetLogs

`func (o *CreateAccount) GetLogs() bool`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *CreateAccount) GetLogsOk() (*bool, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *CreateAccount) SetLogs(v bool)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *CreateAccount) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetName

`func (o *CreateAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAccount) SetName(v string)`

SetName sets Name field to given value.


### GetNotifyFrom

`func (o *CreateAccount) GetNotifyFrom() time.Time`

GetNotifyFrom returns the NotifyFrom field if non-nil, zero value otherwise.

### GetNotifyFromOk

`func (o *CreateAccount) GetNotifyFromOk() (*time.Time, bool)`

GetNotifyFromOk returns a tuple with the NotifyFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFrom

`func (o *CreateAccount) SetNotifyFrom(v time.Time)`

SetNotifyFrom sets NotifyFrom field to given value.

### HasNotifyFrom

`func (o *CreateAccount) HasNotifyFrom() bool`

HasNotifyFrom returns a boolean if a field has been set.

### SetNotifyFromNil

`func (o *CreateAccount) SetNotifyFromNil(b bool)`

 SetNotifyFromNil sets the value for NotifyFrom to be an explicit nil

### UnsetNotifyFrom
`func (o *CreateAccount) UnsetNotifyFrom()`

UnsetNotifyFrom ensures that no value is present for NotifyFrom, not even an explicit nil
### GetOauth2

`func (o *CreateAccount) GetOauth2() OAuth2`

GetOauth2 returns the Oauth2 field if non-nil, zero value otherwise.

### GetOauth2Ok

`func (o *CreateAccount) GetOauth2Ok() (*OAuth2, bool)`

GetOauth2Ok returns a tuple with the Oauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2

`func (o *CreateAccount) SetOauth2(v OAuth2)`

SetOauth2 sets Oauth2 field to given value.

### HasOauth2

`func (o *CreateAccount) HasOauth2() bool`

HasOauth2 returns a boolean if a field has been set.

### GetPath

`func (o *CreateAccount) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateAccount) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreateAccount) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CreateAccount) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *CreateAccount) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *CreateAccount) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetProxy

`func (o *CreateAccount) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *CreateAccount) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *CreateAccount) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *CreateAccount) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetSmtp

`func (o *CreateAccount) GetSmtp() SmtpConfiguration`

GetSmtp returns the Smtp field if non-nil, zero value otherwise.

### GetSmtpOk

`func (o *CreateAccount) GetSmtpOk() (*SmtpConfiguration, bool)`

GetSmtpOk returns a tuple with the Smtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtp

`func (o *CreateAccount) SetSmtp(v SmtpConfiguration)`

SetSmtp sets Smtp field to given value.

### HasSmtp

`func (o *CreateAccount) HasSmtp() bool`

HasSmtp returns a boolean if a field has been set.

### GetSmtpEhloName

`func (o *CreateAccount) GetSmtpEhloName() string`

GetSmtpEhloName returns the SmtpEhloName field if non-nil, zero value otherwise.

### GetSmtpEhloNameOk

`func (o *CreateAccount) GetSmtpEhloNameOk() (*string, bool)`

GetSmtpEhloNameOk returns a tuple with the SmtpEhloName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpEhloName

`func (o *CreateAccount) SetSmtpEhloName(v string)`

SetSmtpEhloName sets SmtpEhloName field to given value.

### HasSmtpEhloName

`func (o *CreateAccount) HasSmtpEhloName() bool`

HasSmtpEhloName returns a boolean if a field has been set.

### SetSmtpEhloNameNil

`func (o *CreateAccount) SetSmtpEhloNameNil(b bool)`

 SetSmtpEhloNameNil sets the value for SmtpEhloName to be an explicit nil

### UnsetSmtpEhloName
`func (o *CreateAccount) UnsetSmtpEhloName()`

UnsetSmtpEhloName ensures that no value is present for SmtpEhloName, not even an explicit nil
### GetSubconnections

`func (o *CreateAccount) GetSubconnections() []string`

GetSubconnections returns the Subconnections field if non-nil, zero value otherwise.

### GetSubconnectionsOk

`func (o *CreateAccount) GetSubconnectionsOk() (*[]string, bool)`

GetSubconnectionsOk returns a tuple with the Subconnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubconnections

`func (o *CreateAccount) SetSubconnections(v []string)`

SetSubconnections sets Subconnections field to given value.

### HasSubconnections

`func (o *CreateAccount) HasSubconnections() bool`

HasSubconnections returns a boolean if a field has been set.

### GetTz

`func (o *CreateAccount) GetTz() string`

GetTz returns the Tz field if non-nil, zero value otherwise.

### GetTzOk

`func (o *CreateAccount) GetTzOk() (*string, bool)`

GetTzOk returns a tuple with the Tz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTz

`func (o *CreateAccount) SetTz(v string)`

SetTz sets Tz field to given value.

### HasTz

`func (o *CreateAccount) HasTz() bool`

HasTz returns a boolean if a field has been set.

### GetWebhooks

`func (o *CreateAccount) GetWebhooks() string`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *CreateAccount) GetWebhooksOk() (*string, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *CreateAccount) SetWebhooks(v string)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *CreateAccount) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetWebhooksCustomHeaders

`func (o *CreateAccount) GetWebhooksCustomHeaders() []WebhooksCustomHeader`

GetWebhooksCustomHeaders returns the WebhooksCustomHeaders field if non-nil, zero value otherwise.

### GetWebhooksCustomHeadersOk

`func (o *CreateAccount) GetWebhooksCustomHeadersOk() (*[]WebhooksCustomHeader, bool)`

GetWebhooksCustomHeadersOk returns a tuple with the WebhooksCustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooksCustomHeaders

`func (o *CreateAccount) SetWebhooksCustomHeaders(v []WebhooksCustomHeader)`

SetWebhooksCustomHeaders sets WebhooksCustomHeaders field to given value.

### HasWebhooksCustomHeaders

`func (o *CreateAccount) HasWebhooksCustomHeaders() bool`

HasWebhooksCustomHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


