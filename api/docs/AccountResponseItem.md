# AccountResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Unique identifier for the email account | 
**App** | Pointer to **string** | OAuth2 application ID | [optional] 
**Counters** | Pointer to [**AccountCounters**](AccountCounters.md) |  | [optional] 
**Email** | Pointer to **string** | Default email address of the account | [optional] 
**LastError** | Pointer to [**NullableAccountErrorEntry**](AccountErrorEntry.md) |  | [optional] 
**Name** | Pointer to **string** | Display name for the account | [optional] 
**Proxy** | Pointer to **string** | Proxy server URL for outbound connections | [optional] 
**SmtpEhloName** | Pointer to **NullableString** | Hostname to use in SMTP EHLO/HELO commands (defaults to system hostname) | [optional] 
**State** | [**State**](State.md) |  | 
**SyncTime** | Pointer to **time.Time** | Last sync time | [optional] 
**Type** | [**ModelType**](ModelType.md) |  | 
**Webhooks** | Pointer to **string** | Account-specific webhook URL | [optional] 

## Methods

### NewAccountResponseItem

`func NewAccountResponseItem(account string, state State, type_ ModelType, ) *AccountResponseItem`

NewAccountResponseItem instantiates a new AccountResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResponseItemWithDefaults

`func NewAccountResponseItemWithDefaults() *AccountResponseItem`

NewAccountResponseItemWithDefaults instantiates a new AccountResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AccountResponseItem) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AccountResponseItem) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AccountResponseItem) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetApp

`func (o *AccountResponseItem) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AccountResponseItem) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AccountResponseItem) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *AccountResponseItem) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetCounters

`func (o *AccountResponseItem) GetCounters() AccountCounters`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *AccountResponseItem) GetCountersOk() (*AccountCounters, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *AccountResponseItem) SetCounters(v AccountCounters)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *AccountResponseItem) HasCounters() bool`

HasCounters returns a boolean if a field has been set.

### GetEmail

`func (o *AccountResponseItem) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountResponseItem) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountResponseItem) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountResponseItem) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLastError

`func (o *AccountResponseItem) GetLastError() AccountErrorEntry`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *AccountResponseItem) GetLastErrorOk() (*AccountErrorEntry, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *AccountResponseItem) SetLastError(v AccountErrorEntry)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *AccountResponseItem) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *AccountResponseItem) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *AccountResponseItem) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetName

`func (o *AccountResponseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountResponseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountResponseItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountResponseItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProxy

`func (o *AccountResponseItem) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *AccountResponseItem) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *AccountResponseItem) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *AccountResponseItem) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetSmtpEhloName

`func (o *AccountResponseItem) GetSmtpEhloName() string`

GetSmtpEhloName returns the SmtpEhloName field if non-nil, zero value otherwise.

### GetSmtpEhloNameOk

`func (o *AccountResponseItem) GetSmtpEhloNameOk() (*string, bool)`

GetSmtpEhloNameOk returns a tuple with the SmtpEhloName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpEhloName

`func (o *AccountResponseItem) SetSmtpEhloName(v string)`

SetSmtpEhloName sets SmtpEhloName field to given value.

### HasSmtpEhloName

`func (o *AccountResponseItem) HasSmtpEhloName() bool`

HasSmtpEhloName returns a boolean if a field has been set.

### SetSmtpEhloNameNil

`func (o *AccountResponseItem) SetSmtpEhloNameNil(b bool)`

 SetSmtpEhloNameNil sets the value for SmtpEhloName to be an explicit nil

### UnsetSmtpEhloName
`func (o *AccountResponseItem) UnsetSmtpEhloName()`

UnsetSmtpEhloName ensures that no value is present for SmtpEhloName, not even an explicit nil
### GetState

`func (o *AccountResponseItem) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AccountResponseItem) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AccountResponseItem) SetState(v State)`

SetState sets State field to given value.


### GetSyncTime

`func (o *AccountResponseItem) GetSyncTime() time.Time`

GetSyncTime returns the SyncTime field if non-nil, zero value otherwise.

### GetSyncTimeOk

`func (o *AccountResponseItem) GetSyncTimeOk() (*time.Time, bool)`

GetSyncTimeOk returns a tuple with the SyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncTime

`func (o *AccountResponseItem) SetSyncTime(v time.Time)`

SetSyncTime sets SyncTime field to given value.

### HasSyncTime

`func (o *AccountResponseItem) HasSyncTime() bool`

HasSyncTime returns a boolean if a field has been set.

### GetType

`func (o *AccountResponseItem) GetType() ModelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountResponseItem) GetTypeOk() (*ModelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountResponseItem) SetType(v ModelType)`

SetType sets Type field to given value.


### GetWebhooks

`func (o *AccountResponseItem) GetWebhooks() string`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *AccountResponseItem) GetWebhooksOk() (*string, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *AccountResponseItem) SetWebhooks(v string)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *AccountResponseItem) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


