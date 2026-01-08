# MailboxResponseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to **int32** | Message count from STATUS command | [optional] 
**Unseen** | Pointer to **int32** | Unread message count from STATUS command | [optional] 

## Methods

### NewMailboxResponseStatus

`func NewMailboxResponseStatus() *MailboxResponseStatus`

NewMailboxResponseStatus instantiates a new MailboxResponseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxResponseStatusWithDefaults

`func NewMailboxResponseStatusWithDefaults() *MailboxResponseStatus`

NewMailboxResponseStatusWithDefaults instantiates a new MailboxResponseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *MailboxResponseStatus) GetMessages() int32`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MailboxResponseStatus) GetMessagesOk() (*int32, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *MailboxResponseStatus) SetMessages(v int32)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *MailboxResponseStatus) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetUnseen

`func (o *MailboxResponseStatus) GetUnseen() int32`

GetUnseen returns the Unseen field if non-nil, zero value otherwise.

### GetUnseenOk

`func (o *MailboxResponseStatus) GetUnseenOk() (*int32, bool)`

GetUnseenOk returns a tuple with the Unseen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnseen

`func (o *MailboxResponseStatus) SetUnseen(v int32)`

SetUnseen sets Unseen field to given value.

### HasUnseen

`func (o *MailboxResponseStatus) HasUnseen() bool`

HasUnseen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


