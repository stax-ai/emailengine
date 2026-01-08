# RequestFlush

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flush** | Pointer to **bool** | Only flush the account if true | [optional] [default to false]
**ImapIndexer** | Pointer to [**NullableIMAPIndexer**](IMAPIndexer.md) |  | [optional] 
**NotifyFrom** | Pointer to **NullableTime** | Only send webhooks for messages received after this date. Defaults to account creation time. IMAP only. | [optional] 

## Methods

### NewRequestFlush

`func NewRequestFlush() *RequestFlush`

NewRequestFlush instantiates a new RequestFlush object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestFlushWithDefaults

`func NewRequestFlushWithDefaults() *RequestFlush`

NewRequestFlushWithDefaults instantiates a new RequestFlush object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlush

`func (o *RequestFlush) GetFlush() bool`

GetFlush returns the Flush field if non-nil, zero value otherwise.

### GetFlushOk

`func (o *RequestFlush) GetFlushOk() (*bool, bool)`

GetFlushOk returns a tuple with the Flush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlush

`func (o *RequestFlush) SetFlush(v bool)`

SetFlush sets Flush field to given value.

### HasFlush

`func (o *RequestFlush) HasFlush() bool`

HasFlush returns a boolean if a field has been set.

### GetImapIndexer

`func (o *RequestFlush) GetImapIndexer() IMAPIndexer`

GetImapIndexer returns the ImapIndexer field if non-nil, zero value otherwise.

### GetImapIndexerOk

`func (o *RequestFlush) GetImapIndexerOk() (*IMAPIndexer, bool)`

GetImapIndexerOk returns a tuple with the ImapIndexer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapIndexer

`func (o *RequestFlush) SetImapIndexer(v IMAPIndexer)`

SetImapIndexer sets ImapIndexer field to given value.

### HasImapIndexer

`func (o *RequestFlush) HasImapIndexer() bool`

HasImapIndexer returns a boolean if a field has been set.

### SetImapIndexerNil

`func (o *RequestFlush) SetImapIndexerNil(b bool)`

 SetImapIndexerNil sets the value for ImapIndexer to be an explicit nil

### UnsetImapIndexer
`func (o *RequestFlush) UnsetImapIndexer()`

UnsetImapIndexer ensures that no value is present for ImapIndexer, not even an explicit nil
### GetNotifyFrom

`func (o *RequestFlush) GetNotifyFrom() time.Time`

GetNotifyFrom returns the NotifyFrom field if non-nil, zero value otherwise.

### GetNotifyFromOk

`func (o *RequestFlush) GetNotifyFromOk() (*time.Time, bool)`

GetNotifyFromOk returns a tuple with the NotifyFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFrom

`func (o *RequestFlush) SetNotifyFrom(v time.Time)`

SetNotifyFrom sets NotifyFrom field to given value.

### HasNotifyFrom

`func (o *RequestFlush) HasNotifyFrom() bool`

HasNotifyFrom returns a boolean if a field has been set.

### SetNotifyFromNil

`func (o *RequestFlush) SetNotifyFromNil(b bool)`

 SetNotifyFromNil sets the value for NotifyFrom to be an explicit nil

### UnsetNotifyFrom
`func (o *RequestFlush) UnsetNotifyFrom()`

UnsetNotifyFrom ensures that no value is present for NotifyFrom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


