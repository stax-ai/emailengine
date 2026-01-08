# UpdateAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Copy** | Pointer to **NullableBool** | Copy submitted messages to Sent folder. Set to &#x60;null&#x60; to unset and use provider specific default. | [optional] 
**Email** | Pointer to **string** | Default email address of the account | [optional] 
**Imap** | Pointer to [**IMAPUpdate**](IMAPUpdate.md) |  | [optional] 
**Locale** | Pointer to **string** | Optional locale | [optional] 
**Logs** | Pointer to **bool** | Store recent logs | [optional] 
**Name** | Pointer to **string** | Display name for the account | [optional] 
**NotifyFrom** | Pointer to **NullableTime** | Only send webhooks for messages received after this date. Defaults to account creation time. IMAP only. | [optional] 
**Oauth2** | Pointer to [**OAuth2Update**](OAuth2Update.md) |  | [optional] 
**Path** | Pointer to **[]string** | Mailbox folders to monitor for changes (IMAP only). Use \&quot;*\&quot; to monitor all folders (default). While you can still access unmonitored folders via API, you won&#39;t receive webhooks for changes in those folders. | [optional] 
**Proxy** | Pointer to **string** | Proxy server URL for outbound connections | [optional] 
**Smtp** | Pointer to [**SMTPUpdate**](SMTPUpdate.md) |  | [optional] 
**SmtpEhloName** | Pointer to **NullableString** | Hostname to use in SMTP EHLO/HELO commands (defaults to system hostname) | [optional] 
**Subconnections** | Pointer to **[]string** | Additional mailbox paths to monitor with dedicated IMAP connections for faster change detection. Use sparingly as connection limits are strict. | [optional] 
**Tz** | Pointer to **string** | Optional timezone | [optional] 
**Webhooks** | Pointer to **string** | Account-specific webhook URL | [optional] 
**WebhooksCustomHeaders** | Pointer to [**[]WebhooksCustomHeader**](WebhooksCustomHeader.md) | Additional HTTP headers to include with every webhook request for authentication or tracking | [optional] 

## Methods

### NewUpdateAccount

`func NewUpdateAccount() *UpdateAccount`

NewUpdateAccount instantiates a new UpdateAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountWithDefaults

`func NewUpdateAccountWithDefaults() *UpdateAccount`

NewUpdateAccountWithDefaults instantiates a new UpdateAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopy

`func (o *UpdateAccount) GetCopy() bool`

GetCopy returns the Copy field if non-nil, zero value otherwise.

### GetCopyOk

`func (o *UpdateAccount) GetCopyOk() (*bool, bool)`

GetCopyOk returns a tuple with the Copy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopy

`func (o *UpdateAccount) SetCopy(v bool)`

SetCopy sets Copy field to given value.

### HasCopy

`func (o *UpdateAccount) HasCopy() bool`

HasCopy returns a boolean if a field has been set.

### SetCopyNil

`func (o *UpdateAccount) SetCopyNil(b bool)`

 SetCopyNil sets the value for Copy to be an explicit nil

### UnsetCopy
`func (o *UpdateAccount) UnsetCopy()`

UnsetCopy ensures that no value is present for Copy, not even an explicit nil
### GetEmail

`func (o *UpdateAccount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateAccount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateAccount) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateAccount) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetImap

`func (o *UpdateAccount) GetImap() IMAPUpdate`

GetImap returns the Imap field if non-nil, zero value otherwise.

### GetImapOk

`func (o *UpdateAccount) GetImapOk() (*IMAPUpdate, bool)`

GetImapOk returns a tuple with the Imap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImap

`func (o *UpdateAccount) SetImap(v IMAPUpdate)`

SetImap sets Imap field to given value.

### HasImap

`func (o *UpdateAccount) HasImap() bool`

HasImap returns a boolean if a field has been set.

### GetLocale

`func (o *UpdateAccount) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UpdateAccount) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UpdateAccount) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UpdateAccount) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetLogs

`func (o *UpdateAccount) GetLogs() bool`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *UpdateAccount) GetLogsOk() (*bool, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *UpdateAccount) SetLogs(v bool)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *UpdateAccount) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetName

`func (o *UpdateAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifyFrom

`func (o *UpdateAccount) GetNotifyFrom() time.Time`

GetNotifyFrom returns the NotifyFrom field if non-nil, zero value otherwise.

### GetNotifyFromOk

`func (o *UpdateAccount) GetNotifyFromOk() (*time.Time, bool)`

GetNotifyFromOk returns a tuple with the NotifyFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFrom

`func (o *UpdateAccount) SetNotifyFrom(v time.Time)`

SetNotifyFrom sets NotifyFrom field to given value.

### HasNotifyFrom

`func (o *UpdateAccount) HasNotifyFrom() bool`

HasNotifyFrom returns a boolean if a field has been set.

### SetNotifyFromNil

`func (o *UpdateAccount) SetNotifyFromNil(b bool)`

 SetNotifyFromNil sets the value for NotifyFrom to be an explicit nil

### UnsetNotifyFrom
`func (o *UpdateAccount) UnsetNotifyFrom()`

UnsetNotifyFrom ensures that no value is present for NotifyFrom, not even an explicit nil
### GetOauth2

`func (o *UpdateAccount) GetOauth2() OAuth2Update`

GetOauth2 returns the Oauth2 field if non-nil, zero value otherwise.

### GetOauth2Ok

`func (o *UpdateAccount) GetOauth2Ok() (*OAuth2Update, bool)`

GetOauth2Ok returns a tuple with the Oauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2

`func (o *UpdateAccount) SetOauth2(v OAuth2Update)`

SetOauth2 sets Oauth2 field to given value.

### HasOauth2

`func (o *UpdateAccount) HasOauth2() bool`

HasOauth2 returns a boolean if a field has been set.

### GetPath

`func (o *UpdateAccount) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *UpdateAccount) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *UpdateAccount) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *UpdateAccount) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *UpdateAccount) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *UpdateAccount) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetProxy

`func (o *UpdateAccount) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *UpdateAccount) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *UpdateAccount) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *UpdateAccount) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetSmtp

`func (o *UpdateAccount) GetSmtp() SMTPUpdate`

GetSmtp returns the Smtp field if non-nil, zero value otherwise.

### GetSmtpOk

`func (o *UpdateAccount) GetSmtpOk() (*SMTPUpdate, bool)`

GetSmtpOk returns a tuple with the Smtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtp

`func (o *UpdateAccount) SetSmtp(v SMTPUpdate)`

SetSmtp sets Smtp field to given value.

### HasSmtp

`func (o *UpdateAccount) HasSmtp() bool`

HasSmtp returns a boolean if a field has been set.

### GetSmtpEhloName

`func (o *UpdateAccount) GetSmtpEhloName() string`

GetSmtpEhloName returns the SmtpEhloName field if non-nil, zero value otherwise.

### GetSmtpEhloNameOk

`func (o *UpdateAccount) GetSmtpEhloNameOk() (*string, bool)`

GetSmtpEhloNameOk returns a tuple with the SmtpEhloName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpEhloName

`func (o *UpdateAccount) SetSmtpEhloName(v string)`

SetSmtpEhloName sets SmtpEhloName field to given value.

### HasSmtpEhloName

`func (o *UpdateAccount) HasSmtpEhloName() bool`

HasSmtpEhloName returns a boolean if a field has been set.

### SetSmtpEhloNameNil

`func (o *UpdateAccount) SetSmtpEhloNameNil(b bool)`

 SetSmtpEhloNameNil sets the value for SmtpEhloName to be an explicit nil

### UnsetSmtpEhloName
`func (o *UpdateAccount) UnsetSmtpEhloName()`

UnsetSmtpEhloName ensures that no value is present for SmtpEhloName, not even an explicit nil
### GetSubconnections

`func (o *UpdateAccount) GetSubconnections() []string`

GetSubconnections returns the Subconnections field if non-nil, zero value otherwise.

### GetSubconnectionsOk

`func (o *UpdateAccount) GetSubconnectionsOk() (*[]string, bool)`

GetSubconnectionsOk returns a tuple with the Subconnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubconnections

`func (o *UpdateAccount) SetSubconnections(v []string)`

SetSubconnections sets Subconnections field to given value.

### HasSubconnections

`func (o *UpdateAccount) HasSubconnections() bool`

HasSubconnections returns a boolean if a field has been set.

### GetTz

`func (o *UpdateAccount) GetTz() string`

GetTz returns the Tz field if non-nil, zero value otherwise.

### GetTzOk

`func (o *UpdateAccount) GetTzOk() (*string, bool)`

GetTzOk returns a tuple with the Tz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTz

`func (o *UpdateAccount) SetTz(v string)`

SetTz sets Tz field to given value.

### HasTz

`func (o *UpdateAccount) HasTz() bool`

HasTz returns a boolean if a field has been set.

### GetWebhooks

`func (o *UpdateAccount) GetWebhooks() string`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *UpdateAccount) GetWebhooksOk() (*string, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *UpdateAccount) SetWebhooks(v string)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *UpdateAccount) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetWebhooksCustomHeaders

`func (o *UpdateAccount) GetWebhooksCustomHeaders() []WebhooksCustomHeader`

GetWebhooksCustomHeaders returns the WebhooksCustomHeaders field if non-nil, zero value otherwise.

### GetWebhooksCustomHeadersOk

`func (o *UpdateAccount) GetWebhooksCustomHeadersOk() (*[]WebhooksCustomHeader, bool)`

GetWebhooksCustomHeadersOk returns a tuple with the WebhooksCustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooksCustomHeaders

`func (o *UpdateAccount) SetWebhooksCustomHeaders(v []WebhooksCustomHeader)`

SetWebhooksCustomHeaders sets WebhooksCustomHeaders field to given value.

### HasWebhooksCustomHeaders

`func (o *UpdateAccount) HasWebhooksCustomHeaders() bool`

HasWebhooksCustomHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


