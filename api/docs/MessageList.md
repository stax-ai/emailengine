# MessageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]MessageListEntry**](MessageListEntry.md) |  | [optional] 
**NextPageCursor** | Pointer to **NullableString** | Cursor for fetching the next page (null when no more pages) | [optional] 
**Page** | Pointer to **int32** | Current page number (zero-based) | [optional] 
**Pages** | Pointer to **int32** | Total number of pages available (exact for IMAP, approximate for Gmail API) | [optional] 
**PrevPageCursor** | Pointer to **NullableString** | Cursor for fetching the previous page | [optional] 
**Total** | Pointer to **int32** | Total number of messages matching the query (exact for IMAP, approximate for Gmail API) | [optional] 

## Methods

### NewMessageList

`func NewMessageList() *MessageList`

NewMessageList instantiates a new MessageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageListWithDefaults

`func NewMessageListWithDefaults() *MessageList`

NewMessageListWithDefaults instantiates a new MessageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *MessageList) GetMessages() []MessageListEntry`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MessageList) GetMessagesOk() (*[]MessageListEntry, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *MessageList) SetMessages(v []MessageListEntry)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *MessageList) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetNextPageCursor

`func (o *MessageList) GetNextPageCursor() string`

GetNextPageCursor returns the NextPageCursor field if non-nil, zero value otherwise.

### GetNextPageCursorOk

`func (o *MessageList) GetNextPageCursorOk() (*string, bool)`

GetNextPageCursorOk returns a tuple with the NextPageCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageCursor

`func (o *MessageList) SetNextPageCursor(v string)`

SetNextPageCursor sets NextPageCursor field to given value.

### HasNextPageCursor

`func (o *MessageList) HasNextPageCursor() bool`

HasNextPageCursor returns a boolean if a field has been set.

### SetNextPageCursorNil

`func (o *MessageList) SetNextPageCursorNil(b bool)`

 SetNextPageCursorNil sets the value for NextPageCursor to be an explicit nil

### UnsetNextPageCursor
`func (o *MessageList) UnsetNextPageCursor()`

UnsetNextPageCursor ensures that no value is present for NextPageCursor, not even an explicit nil
### GetPage

`func (o *MessageList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *MessageList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *MessageList) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *MessageList) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *MessageList) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *MessageList) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *MessageList) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *MessageList) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetPrevPageCursor

`func (o *MessageList) GetPrevPageCursor() string`

GetPrevPageCursor returns the PrevPageCursor field if non-nil, zero value otherwise.

### GetPrevPageCursorOk

`func (o *MessageList) GetPrevPageCursorOk() (*string, bool)`

GetPrevPageCursorOk returns a tuple with the PrevPageCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPageCursor

`func (o *MessageList) SetPrevPageCursor(v string)`

SetPrevPageCursor sets PrevPageCursor field to given value.

### HasPrevPageCursor

`func (o *MessageList) HasPrevPageCursor() bool`

HasPrevPageCursor returns a boolean if a field has been set.

### SetPrevPageCursorNil

`func (o *MessageList) SetPrevPageCursorNil(b bool)`

 SetPrevPageCursorNil sets the value for PrevPageCursor to be an explicit nil

### UnsetPrevPageCursor
`func (o *MessageList) UnsetPrevPageCursor()`

UnsetPrevPageCursor ensures that no value is present for PrevPageCursor, not even an explicit nil
### GetTotal

`func (o *MessageList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MessageList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MessageList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MessageList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


