# AccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Unique identifier for the email account | 
**App** | Pointer to **string** | OAuth2 application ID | [optional] 
**BaseScopes** | Pointer to [**BaseScopes**](BaseScopes.md) |  | [optional] 
**Copy** | Pointer to **bool** | Copy submitted messages to Sent folder | [optional] 
**Counters** | Pointer to [**AccountCounters**](AccountCounters.md) |  | [optional] 
**Email** | Pointer to **string** | Default email address of the account | [optional] 
**Imap** | Pointer to [**IMAPResponse**](IMAPResponse.md) |  | [optional] 
**ImapIndexer** | Pointer to [**NullableIMAPIndexer**](IMAPIndexer.md) |  | [optional] 
**LastError** | Pointer to [**NullableAccountErrorEntry**](AccountErrorEntry.md) |  | [optional] 
**Locale** | Pointer to **string** | Optional locale | [optional] 
**Logs** | Pointer to **bool** | Store recent logs | [optional] 
**Name** | Pointer to **string** | Display name for the account | [optional] 
**NotifyFrom** | Pointer to **NullableTime** | Only send webhooks for messages received after this date. Defaults to account creation time. IMAP only. | [optional] 
**Oauth2** | Pointer to [**Oauth2Response**](Oauth2Response.md) |  | [optional] 
**Path** | Pointer to **[]string** | Mailbox folders to monitor for changes (IMAP only). Use \&quot;*\&quot; to monitor all folders (default). While you can still access unmonitored folders via API, you won&#39;t receive webhooks for changes in those folders. | [optional] 
**Proxy** | Pointer to **string** | Proxy server URL for outbound connections | [optional] 
**Quota** | Pointer to [**AccountQuota**](AccountQuota.md) |  | [optional] 
**Smtp** | Pointer to [**SMTPResponse**](SMTPResponse.md) |  | [optional] 
**SmtpEhloName** | Pointer to **NullableString** | Hostname to use in SMTP EHLO/HELO commands (defaults to system hostname) | [optional] 
**SmtpStatus** | Pointer to [**SMTPInfoStatus**](SMTPInfoStatus.md) |  | [optional] 
**State** | Pointer to [**AccountInfoState**](AccountInfoState.md) |  | [optional] 
**Subconnections** | Pointer to **[]string** | Additional mailbox paths to monitor with dedicated IMAP connections for faster change detection. Use sparingly as connection limits are strict. | [optional] 
**SyncTime** | Pointer to **time.Time** | Last sync time | [optional] 
**Type** | [**ModelType**](ModelType.md) |  | 
**Tz** | Pointer to **string** | Optional timezone | [optional] 
**Webhooks** | Pointer to **string** | Account-specific webhook URL | [optional] 
**WebhooksCustomHeaders** | Pointer to [**[]WebhooksCustomHeader**](WebhooksCustomHeader.md) | Additional HTTP headers to include with every webhook request for authentication or tracking | [optional] 

## Methods

### NewAccountResponse

`func NewAccountResponse(account string, type_ ModelType, ) *AccountResponse`

NewAccountResponse instantiates a new AccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResponseWithDefaults

`func NewAccountResponseWithDefaults() *AccountResponse`

NewAccountResponseWithDefaults instantiates a new AccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AccountResponse) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AccountResponse) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AccountResponse) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetApp

`func (o *AccountResponse) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AccountResponse) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AccountResponse) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *AccountResponse) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetBaseScopes

`func (o *AccountResponse) GetBaseScopes() BaseScopes`

GetBaseScopes returns the BaseScopes field if non-nil, zero value otherwise.

### GetBaseScopesOk

`func (o *AccountResponse) GetBaseScopesOk() (*BaseScopes, bool)`

GetBaseScopesOk returns a tuple with the BaseScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScopes

`func (o *AccountResponse) SetBaseScopes(v BaseScopes)`

SetBaseScopes sets BaseScopes field to given value.

### HasBaseScopes

`func (o *AccountResponse) HasBaseScopes() bool`

HasBaseScopes returns a boolean if a field has been set.

### GetCopy

`func (o *AccountResponse) GetCopy() bool`

GetCopy returns the Copy field if non-nil, zero value otherwise.

### GetCopyOk

`func (o *AccountResponse) GetCopyOk() (*bool, bool)`

GetCopyOk returns a tuple with the Copy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopy

`func (o *AccountResponse) SetCopy(v bool)`

SetCopy sets Copy field to given value.

### HasCopy

`func (o *AccountResponse) HasCopy() bool`

HasCopy returns a boolean if a field has been set.

### GetCounters

`func (o *AccountResponse) GetCounters() AccountCounters`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *AccountResponse) GetCountersOk() (*AccountCounters, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *AccountResponse) SetCounters(v AccountCounters)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *AccountResponse) HasCounters() bool`

HasCounters returns a boolean if a field has been set.

### GetEmail

`func (o *AccountResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetImap

`func (o *AccountResponse) GetImap() IMAPResponse`

GetImap returns the Imap field if non-nil, zero value otherwise.

### GetImapOk

`func (o *AccountResponse) GetImapOk() (*IMAPResponse, bool)`

GetImapOk returns a tuple with the Imap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImap

`func (o *AccountResponse) SetImap(v IMAPResponse)`

SetImap sets Imap field to given value.

### HasImap

`func (o *AccountResponse) HasImap() bool`

HasImap returns a boolean if a field has been set.

### GetImapIndexer

`func (o *AccountResponse) GetImapIndexer() IMAPIndexer`

GetImapIndexer returns the ImapIndexer field if non-nil, zero value otherwise.

### GetImapIndexerOk

`func (o *AccountResponse) GetImapIndexerOk() (*IMAPIndexer, bool)`

GetImapIndexerOk returns a tuple with the ImapIndexer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapIndexer

`func (o *AccountResponse) SetImapIndexer(v IMAPIndexer)`

SetImapIndexer sets ImapIndexer field to given value.

### HasImapIndexer

`func (o *AccountResponse) HasImapIndexer() bool`

HasImapIndexer returns a boolean if a field has been set.

### SetImapIndexerNil

`func (o *AccountResponse) SetImapIndexerNil(b bool)`

 SetImapIndexerNil sets the value for ImapIndexer to be an explicit nil

### UnsetImapIndexer
`func (o *AccountResponse) UnsetImapIndexer()`

UnsetImapIndexer ensures that no value is present for ImapIndexer, not even an explicit nil
### GetLastError

`func (o *AccountResponse) GetLastError() AccountErrorEntry`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *AccountResponse) GetLastErrorOk() (*AccountErrorEntry, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *AccountResponse) SetLastError(v AccountErrorEntry)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *AccountResponse) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *AccountResponse) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *AccountResponse) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetLocale

`func (o *AccountResponse) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *AccountResponse) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *AccountResponse) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *AccountResponse) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetLogs

`func (o *AccountResponse) GetLogs() bool`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *AccountResponse) GetLogsOk() (*bool, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *AccountResponse) SetLogs(v bool)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *AccountResponse) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetName

`func (o *AccountResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifyFrom

`func (o *AccountResponse) GetNotifyFrom() time.Time`

GetNotifyFrom returns the NotifyFrom field if non-nil, zero value otherwise.

### GetNotifyFromOk

`func (o *AccountResponse) GetNotifyFromOk() (*time.Time, bool)`

GetNotifyFromOk returns a tuple with the NotifyFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFrom

`func (o *AccountResponse) SetNotifyFrom(v time.Time)`

SetNotifyFrom sets NotifyFrom field to given value.

### HasNotifyFrom

`func (o *AccountResponse) HasNotifyFrom() bool`

HasNotifyFrom returns a boolean if a field has been set.

### SetNotifyFromNil

`func (o *AccountResponse) SetNotifyFromNil(b bool)`

 SetNotifyFromNil sets the value for NotifyFrom to be an explicit nil

### UnsetNotifyFrom
`func (o *AccountResponse) UnsetNotifyFrom()`

UnsetNotifyFrom ensures that no value is present for NotifyFrom, not even an explicit nil
### GetOauth2

`func (o *AccountResponse) GetOauth2() Oauth2Response`

GetOauth2 returns the Oauth2 field if non-nil, zero value otherwise.

### GetOauth2Ok

`func (o *AccountResponse) GetOauth2Ok() (*Oauth2Response, bool)`

GetOauth2Ok returns a tuple with the Oauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2

`func (o *AccountResponse) SetOauth2(v Oauth2Response)`

SetOauth2 sets Oauth2 field to given value.

### HasOauth2

`func (o *AccountResponse) HasOauth2() bool`

HasOauth2 returns a boolean if a field has been set.

### GetPath

`func (o *AccountResponse) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AccountResponse) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AccountResponse) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AccountResponse) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *AccountResponse) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *AccountResponse) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetProxy

`func (o *AccountResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *AccountResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *AccountResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *AccountResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetQuota

`func (o *AccountResponse) GetQuota() AccountQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *AccountResponse) GetQuotaOk() (*AccountQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *AccountResponse) SetQuota(v AccountQuota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *AccountResponse) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetSmtp

`func (o *AccountResponse) GetSmtp() SMTPResponse`

GetSmtp returns the Smtp field if non-nil, zero value otherwise.

### GetSmtpOk

`func (o *AccountResponse) GetSmtpOk() (*SMTPResponse, bool)`

GetSmtpOk returns a tuple with the Smtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtp

`func (o *AccountResponse) SetSmtp(v SMTPResponse)`

SetSmtp sets Smtp field to given value.

### HasSmtp

`func (o *AccountResponse) HasSmtp() bool`

HasSmtp returns a boolean if a field has been set.

### GetSmtpEhloName

`func (o *AccountResponse) GetSmtpEhloName() string`

GetSmtpEhloName returns the SmtpEhloName field if non-nil, zero value otherwise.

### GetSmtpEhloNameOk

`func (o *AccountResponse) GetSmtpEhloNameOk() (*string, bool)`

GetSmtpEhloNameOk returns a tuple with the SmtpEhloName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpEhloName

`func (o *AccountResponse) SetSmtpEhloName(v string)`

SetSmtpEhloName sets SmtpEhloName field to given value.

### HasSmtpEhloName

`func (o *AccountResponse) HasSmtpEhloName() bool`

HasSmtpEhloName returns a boolean if a field has been set.

### SetSmtpEhloNameNil

`func (o *AccountResponse) SetSmtpEhloNameNil(b bool)`

 SetSmtpEhloNameNil sets the value for SmtpEhloName to be an explicit nil

### UnsetSmtpEhloName
`func (o *AccountResponse) UnsetSmtpEhloName()`

UnsetSmtpEhloName ensures that no value is present for SmtpEhloName, not even an explicit nil
### GetSmtpStatus

`func (o *AccountResponse) GetSmtpStatus() SMTPInfoStatus`

GetSmtpStatus returns the SmtpStatus field if non-nil, zero value otherwise.

### GetSmtpStatusOk

`func (o *AccountResponse) GetSmtpStatusOk() (*SMTPInfoStatus, bool)`

GetSmtpStatusOk returns a tuple with the SmtpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpStatus

`func (o *AccountResponse) SetSmtpStatus(v SMTPInfoStatus)`

SetSmtpStatus sets SmtpStatus field to given value.

### HasSmtpStatus

`func (o *AccountResponse) HasSmtpStatus() bool`

HasSmtpStatus returns a boolean if a field has been set.

### GetState

`func (o *AccountResponse) GetState() AccountInfoState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AccountResponse) GetStateOk() (*AccountInfoState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AccountResponse) SetState(v AccountInfoState)`

SetState sets State field to given value.

### HasState

`func (o *AccountResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubconnections

`func (o *AccountResponse) GetSubconnections() []string`

GetSubconnections returns the Subconnections field if non-nil, zero value otherwise.

### GetSubconnectionsOk

`func (o *AccountResponse) GetSubconnectionsOk() (*[]string, bool)`

GetSubconnectionsOk returns a tuple with the Subconnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubconnections

`func (o *AccountResponse) SetSubconnections(v []string)`

SetSubconnections sets Subconnections field to given value.

### HasSubconnections

`func (o *AccountResponse) HasSubconnections() bool`

HasSubconnections returns a boolean if a field has been set.

### GetSyncTime

`func (o *AccountResponse) GetSyncTime() time.Time`

GetSyncTime returns the SyncTime field if non-nil, zero value otherwise.

### GetSyncTimeOk

`func (o *AccountResponse) GetSyncTimeOk() (*time.Time, bool)`

GetSyncTimeOk returns a tuple with the SyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncTime

`func (o *AccountResponse) SetSyncTime(v time.Time)`

SetSyncTime sets SyncTime field to given value.

### HasSyncTime

`func (o *AccountResponse) HasSyncTime() bool`

HasSyncTime returns a boolean if a field has been set.

### GetType

`func (o *AccountResponse) GetType() ModelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountResponse) GetTypeOk() (*ModelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountResponse) SetType(v ModelType)`

SetType sets Type field to given value.


### GetTz

`func (o *AccountResponse) GetTz() string`

GetTz returns the Tz field if non-nil, zero value otherwise.

### GetTzOk

`func (o *AccountResponse) GetTzOk() (*string, bool)`

GetTzOk returns a tuple with the Tz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTz

`func (o *AccountResponse) SetTz(v string)`

SetTz sets Tz field to given value.

### HasTz

`func (o *AccountResponse) HasTz() bool`

HasTz returns a boolean if a field has been set.

### GetWebhooks

`func (o *AccountResponse) GetWebhooks() string`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *AccountResponse) GetWebhooksOk() (*string, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *AccountResponse) SetWebhooks(v string)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *AccountResponse) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetWebhooksCustomHeaders

`func (o *AccountResponse) GetWebhooksCustomHeaders() []WebhooksCustomHeader`

GetWebhooksCustomHeaders returns the WebhooksCustomHeaders field if non-nil, zero value otherwise.

### GetWebhooksCustomHeadersOk

`func (o *AccountResponse) GetWebhooksCustomHeadersOk() (*[]WebhooksCustomHeader, bool)`

GetWebhooksCustomHeadersOk returns a tuple with the WebhooksCustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooksCustomHeaders

`func (o *AccountResponse) SetWebhooksCustomHeaders(v []WebhooksCustomHeader)`

SetWebhooksCustomHeaders sets WebhooksCustomHeaders field to given value.

### HasWebhooksCustomHeaders

`func (o *AccountResponse) HasWebhooksCustomHeaders() bool`

HasWebhooksCustomHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


