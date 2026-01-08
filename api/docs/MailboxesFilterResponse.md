# MailboxesFilterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mailboxes** | Pointer to [**[]MailboxResponseItem**](MailboxResponseItem.md) |  | [optional] 

## Methods

### NewMailboxesFilterResponse

`func NewMailboxesFilterResponse() *MailboxesFilterResponse`

NewMailboxesFilterResponse instantiates a new MailboxesFilterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxesFilterResponseWithDefaults

`func NewMailboxesFilterResponseWithDefaults() *MailboxesFilterResponse`

NewMailboxesFilterResponseWithDefaults instantiates a new MailboxesFilterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailboxes

`func (o *MailboxesFilterResponse) GetMailboxes() []MailboxResponseItem`

GetMailboxes returns the Mailboxes field if non-nil, zero value otherwise.

### GetMailboxesOk

`func (o *MailboxesFilterResponse) GetMailboxesOk() (*[]MailboxResponseItem, bool)`

GetMailboxesOk returns a tuple with the Mailboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxes

`func (o *MailboxesFilterResponse) SetMailboxes(v []MailboxResponseItem)`

SetMailboxes sets Mailboxes field to given value.

### HasMailboxes

`func (o *MailboxesFilterResponse) HasMailboxes() bool`

HasMailboxes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


