# BulkResponseEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** | Message ID | [optional] 
**QueueId** | Pointer to **string** | Queue identifier for scheduled email. Not present for bulk messages. | [optional] 
**Reference** | Pointer to [**Model12**](Model12.md) |  | [optional] 
**SendAt** | Pointer to **time.Time** | Send message at specified time. Overrides message level &#x60;sendAt&#x60; value. | [optional] 
**Skipped** | Pointer to [**Skipped**](Skipped.md) |  | [optional] 
**Success** | Pointer to **bool** | Was the referenced message processed successfully | [optional] 
**To** | Pointer to [**ToAddressSingle**](ToAddressSingle.md) |  | [optional] 

## Methods

### NewBulkResponseEntry

`func NewBulkResponseEntry() *BulkResponseEntry`

NewBulkResponseEntry instantiates a new BulkResponseEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkResponseEntryWithDefaults

`func NewBulkResponseEntryWithDefaults() *BulkResponseEntry`

NewBulkResponseEntryWithDefaults instantiates a new BulkResponseEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *BulkResponseEntry) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *BulkResponseEntry) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *BulkResponseEntry) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *BulkResponseEntry) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetQueueId

`func (o *BulkResponseEntry) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *BulkResponseEntry) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *BulkResponseEntry) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *BulkResponseEntry) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetReference

`func (o *BulkResponseEntry) GetReference() Model12`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *BulkResponseEntry) GetReferenceOk() (*Model12, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *BulkResponseEntry) SetReference(v Model12)`

SetReference sets Reference field to given value.

### HasReference

`func (o *BulkResponseEntry) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSendAt

`func (o *BulkResponseEntry) GetSendAt() time.Time`

GetSendAt returns the SendAt field if non-nil, zero value otherwise.

### GetSendAtOk

`func (o *BulkResponseEntry) GetSendAtOk() (*time.Time, bool)`

GetSendAtOk returns a tuple with the SendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAt

`func (o *BulkResponseEntry) SetSendAt(v time.Time)`

SetSendAt sets SendAt field to given value.

### HasSendAt

`func (o *BulkResponseEntry) HasSendAt() bool`

HasSendAt returns a boolean if a field has been set.

### GetSkipped

`func (o *BulkResponseEntry) GetSkipped() Skipped`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *BulkResponseEntry) GetSkippedOk() (*Skipped, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *BulkResponseEntry) SetSkipped(v Skipped)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *BulkResponseEntry) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetSuccess

`func (o *BulkResponseEntry) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *BulkResponseEntry) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *BulkResponseEntry) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *BulkResponseEntry) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTo

`func (o *BulkResponseEntry) GetTo() ToAddressSingle`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BulkResponseEntry) GetToOk() (*ToAddressSingle, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BulkResponseEntry) SetTo(v ToAddressSingle)`

SetTo sets To field to given value.

### HasTo

`func (o *BulkResponseEntry) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


