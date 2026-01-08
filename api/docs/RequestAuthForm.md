# RequestAuthForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **NullableString** | Account ID. If set to &#x60;null&#x60;, a unique ID will be generated automatically. If you provide an existing account ID, the settings for that account will be updated instead | [optional] 
**Delegated** | Pointer to **bool** | If true, configures this account as a shared mailbox. Currently supported by MS365 OAuth2 accounts | [optional] [default to false]
**Email** | Pointer to **string** | Specifies the default email address for this account. Users can change it if needed. | [optional] 
**Name** | Pointer to **string** | Display name for the account | [optional] 
**NotifyFrom** | Pointer to **NullableTime** | Only send webhooks for messages received after this date. Defaults to account creation time. IMAP only. | [optional] 
**Path** | Pointer to **[]string** | Mailbox folders to monitor for changes (IMAP only). Use \&quot;*\&quot; to monitor all folders (default). While you can still access unmonitored folders via API, you won&#39;t receive webhooks for changes in those folders. | [optional] 
**RedirectUrl** | **string** | After the authentication process is completed, the user is redirected to this URL | 
**Subconnections** | Pointer to **[]string** | Additional mailbox paths to monitor with dedicated IMAP connections for faster change detection. Use sparingly as connection limits are strict. | [optional] 
**Type** | Pointer to [**DefaultAccountType**](DefaultAccountType.md) |  | [optional] [default to DEFAULTACCOUNTTYPE_FALSE]

## Methods

### NewRequestAuthForm

`func NewRequestAuthForm(redirectUrl string, ) *RequestAuthForm`

NewRequestAuthForm instantiates a new RequestAuthForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestAuthFormWithDefaults

`func NewRequestAuthFormWithDefaults() *RequestAuthForm`

NewRequestAuthFormWithDefaults instantiates a new RequestAuthForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *RequestAuthForm) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *RequestAuthForm) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *RequestAuthForm) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *RequestAuthForm) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *RequestAuthForm) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *RequestAuthForm) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetDelegated

`func (o *RequestAuthForm) GetDelegated() bool`

GetDelegated returns the Delegated field if non-nil, zero value otherwise.

### GetDelegatedOk

`func (o *RequestAuthForm) GetDelegatedOk() (*bool, bool)`

GetDelegatedOk returns a tuple with the Delegated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegated

`func (o *RequestAuthForm) SetDelegated(v bool)`

SetDelegated sets Delegated field to given value.

### HasDelegated

`func (o *RequestAuthForm) HasDelegated() bool`

HasDelegated returns a boolean if a field has been set.

### GetEmail

`func (o *RequestAuthForm) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RequestAuthForm) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RequestAuthForm) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RequestAuthForm) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *RequestAuthForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestAuthForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestAuthForm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestAuthForm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifyFrom

`func (o *RequestAuthForm) GetNotifyFrom() time.Time`

GetNotifyFrom returns the NotifyFrom field if non-nil, zero value otherwise.

### GetNotifyFromOk

`func (o *RequestAuthForm) GetNotifyFromOk() (*time.Time, bool)`

GetNotifyFromOk returns a tuple with the NotifyFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFrom

`func (o *RequestAuthForm) SetNotifyFrom(v time.Time)`

SetNotifyFrom sets NotifyFrom field to given value.

### HasNotifyFrom

`func (o *RequestAuthForm) HasNotifyFrom() bool`

HasNotifyFrom returns a boolean if a field has been set.

### SetNotifyFromNil

`func (o *RequestAuthForm) SetNotifyFromNil(b bool)`

 SetNotifyFromNil sets the value for NotifyFrom to be an explicit nil

### UnsetNotifyFrom
`func (o *RequestAuthForm) UnsetNotifyFrom()`

UnsetNotifyFrom ensures that no value is present for NotifyFrom, not even an explicit nil
### GetPath

`func (o *RequestAuthForm) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RequestAuthForm) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RequestAuthForm) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RequestAuthForm) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *RequestAuthForm) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *RequestAuthForm) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetRedirectUrl

`func (o *RequestAuthForm) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *RequestAuthForm) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *RequestAuthForm) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.


### GetSubconnections

`func (o *RequestAuthForm) GetSubconnections() []string`

GetSubconnections returns the Subconnections field if non-nil, zero value otherwise.

### GetSubconnectionsOk

`func (o *RequestAuthForm) GetSubconnectionsOk() (*[]string, bool)`

GetSubconnectionsOk returns a tuple with the Subconnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubconnections

`func (o *RequestAuthForm) SetSubconnections(v []string)`

SetSubconnections sets Subconnections field to given value.

### HasSubconnections

`func (o *RequestAuthForm) HasSubconnections() bool`

HasSubconnections returns a boolean if a field has been set.

### GetType

`func (o *RequestAuthForm) GetType() DefaultAccountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestAuthForm) GetTypeOk() (*DefaultAccountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestAuthForm) SetType(v DefaultAccountType)`

SetType sets Type field to given value.

### HasType

`func (o *RequestAuthForm) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


