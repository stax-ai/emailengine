# VerifyAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Imap** | Pointer to [**ImapConfiguration**](ImapConfiguration.md) |  | [optional] 
**Mailboxes** | Pointer to **bool** | Include mailbox listing in response | [optional] [default to false]
**Proxy** | Pointer to **string** | Proxy server URL for outbound connections | [optional] 
**Smtp** | Pointer to [**SmtpConfiguration**](SmtpConfiguration.md) |  | [optional] 
**SmtpEhloName** | Pointer to **NullableString** | Hostname to use in SMTP EHLO/HELO commands (defaults to system hostname) | [optional] 

## Methods

### NewVerifyAccount

`func NewVerifyAccount() *VerifyAccount`

NewVerifyAccount instantiates a new VerifyAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyAccountWithDefaults

`func NewVerifyAccountWithDefaults() *VerifyAccount`

NewVerifyAccountWithDefaults instantiates a new VerifyAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImap

`func (o *VerifyAccount) GetImap() ImapConfiguration`

GetImap returns the Imap field if non-nil, zero value otherwise.

### GetImapOk

`func (o *VerifyAccount) GetImapOk() (*ImapConfiguration, bool)`

GetImapOk returns a tuple with the Imap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImap

`func (o *VerifyAccount) SetImap(v ImapConfiguration)`

SetImap sets Imap field to given value.

### HasImap

`func (o *VerifyAccount) HasImap() bool`

HasImap returns a boolean if a field has been set.

### GetMailboxes

`func (o *VerifyAccount) GetMailboxes() bool`

GetMailboxes returns the Mailboxes field if non-nil, zero value otherwise.

### GetMailboxesOk

`func (o *VerifyAccount) GetMailboxesOk() (*bool, bool)`

GetMailboxesOk returns a tuple with the Mailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxes

`func (o *VerifyAccount) SetMailboxes(v bool)`

SetMailboxes sets Mailboxes field to given value.

### HasMailboxes

`func (o *VerifyAccount) HasMailboxes() bool`

HasMailboxes returns a boolean if a field has been set.

### GetProxy

`func (o *VerifyAccount) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *VerifyAccount) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *VerifyAccount) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *VerifyAccount) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetSmtp

`func (o *VerifyAccount) GetSmtp() SmtpConfiguration`

GetSmtp returns the Smtp field if non-nil, zero value otherwise.

### GetSmtpOk

`func (o *VerifyAccount) GetSmtpOk() (*SmtpConfiguration, bool)`

GetSmtpOk returns a tuple with the Smtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtp

`func (o *VerifyAccount) SetSmtp(v SmtpConfiguration)`

SetSmtp sets Smtp field to given value.

### HasSmtp

`func (o *VerifyAccount) HasSmtp() bool`

HasSmtp returns a boolean if a field has been set.

### GetSmtpEhloName

`func (o *VerifyAccount) GetSmtpEhloName() string`

GetSmtpEhloName returns the SmtpEhloName field if non-nil, zero value otherwise.

### GetSmtpEhloNameOk

`func (o *VerifyAccount) GetSmtpEhloNameOk() (*string, bool)`

GetSmtpEhloNameOk returns a tuple with the SmtpEhloName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpEhloName

`func (o *VerifyAccount) SetSmtpEhloName(v string)`

SetSmtpEhloName sets SmtpEhloName field to given value.

### HasSmtpEhloName

`func (o *VerifyAccount) HasSmtpEhloName() bool`

HasSmtpEhloName returns a boolean if a field has been set.

### SetSmtpEhloNameNil

`func (o *VerifyAccount) SetSmtpEhloNameNil(b bool)`

 SetSmtpEhloNameNil sets the value for SmtpEhloName to be an explicit nil

### UnsetSmtpEhloName
`func (o *VerifyAccount) UnsetSmtpEhloName()`

UnsetSmtpEhloName ensures that no value is present for SmtpEhloName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


