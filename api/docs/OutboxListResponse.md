# OutboxListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]OutboxEntry**](OutboxEntry.md) |  | [optional] 
**Page** | Pointer to **int32** | Current page (0-based index) | [optional] 
**Pages** | Pointer to **int32** | Total page count | [optional] 
**Total** | Pointer to **int32** | How many matching entries | [optional] 

## Methods

### NewOutboxListResponse

`func NewOutboxListResponse() *OutboxListResponse`

NewOutboxListResponse instantiates a new OutboxListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboxListResponseWithDefaults

`func NewOutboxListResponseWithDefaults() *OutboxListResponse`

NewOutboxListResponseWithDefaults instantiates a new OutboxListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *OutboxListResponse) GetMessages() []OutboxEntry`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *OutboxListResponse) GetMessagesOk() (*[]OutboxEntry, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *OutboxListResponse) SetMessages(v []OutboxEntry)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *OutboxListResponse) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetPage

`func (o *OutboxListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OutboxListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OutboxListResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *OutboxListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *OutboxListResponse) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *OutboxListResponse) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *OutboxListResponse) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *OutboxListResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetTotal

`func (o *OutboxListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OutboxListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OutboxListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *OutboxListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


