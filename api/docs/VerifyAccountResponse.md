# VerifyAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Imap** | Pointer to [**Imap**](Imap.md) |  | [optional] 
**Mailboxes** | Pointer to [**[]MailboxShortResponseItem**](MailboxShortResponseItem.md) |  | [optional] 
**Smtp** | Pointer to [**Smtp**](Smtp.md) |  | [optional] 

## Methods

### NewVerifyAccountResponse

`func NewVerifyAccountResponse() *VerifyAccountResponse`

NewVerifyAccountResponse instantiates a new VerifyAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyAccountResponseWithDefaults

`func NewVerifyAccountResponseWithDefaults() *VerifyAccountResponse`

NewVerifyAccountResponseWithDefaults instantiates a new VerifyAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImap

`func (o *VerifyAccountResponse) GetImap() Imap`

GetImap returns the Imap field if non-nil, zero value otherwise.

### GetImapOk

`func (o *VerifyAccountResponse) GetImapOk() (*Imap, bool)`

GetImapOk returns a tuple with the Imap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImap

`func (o *VerifyAccountResponse) SetImap(v Imap)`

SetImap sets Imap field to given value.

### HasImap

`func (o *VerifyAccountResponse) HasImap() bool`

HasImap returns a boolean if a field has been set.

### GetMailboxes

`func (o *VerifyAccountResponse) GetMailboxes() []MailboxShortResponseItem`

GetMailboxes returns the Mailboxes field if non-nil, zero value otherwise.

### GetMailboxesOk

`func (o *VerifyAccountResponse) GetMailboxesOk() (*[]MailboxShortResponseItem, bool)`

GetMailboxesOk returns a tuple with the Mailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxes

`func (o *VerifyAccountResponse) SetMailboxes(v []MailboxShortResponseItem)`

SetMailboxes sets Mailboxes field to given value.

### HasMailboxes

`func (o *VerifyAccountResponse) HasMailboxes() bool`

HasMailboxes returns a boolean if a field has been set.

### GetSmtp

`func (o *VerifyAccountResponse) GetSmtp() Smtp`

GetSmtp returns the Smtp field if non-nil, zero value otherwise.

### GetSmtpOk

`func (o *VerifyAccountResponse) GetSmtpOk() (*Smtp, bool)`

GetSmtpOk returns a tuple with the Smtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtp

`func (o *VerifyAccountResponse) SetSmtp(v Smtp)`

SetSmtp sets Smtp field to given value.

### HasSmtp

`func (o *VerifyAccountResponse) HasSmtp() bool`

HasSmtp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


