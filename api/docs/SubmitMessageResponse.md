# SubmitMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailMerge** | Pointer to [**[]BulkResponseEntry**](BulkResponseEntry.md) | Bulk message responses | [optional] 
**MessageId** | Pointer to **string** | Message-ID header value. Not present for bulk messages. | [optional] 
**Preview** | Pointer to **string** | Base64 encoded RFC822 email if a preview was requested | [optional] 
**QueueId** | Pointer to **string** | Queue identifier for scheduled email. Not present for bulk messages. | [optional] 
**Reference** | Pointer to [**Model11**](Model11.md) |  | [optional] 
**Response** | Pointer to **string** |  | [optional] 
**SendAt** | Pointer to **string** | Scheduled send time | [optional] 

## Methods

### NewSubmitMessageResponse

`func NewSubmitMessageResponse() *SubmitMessageResponse`

NewSubmitMessageResponse instantiates a new SubmitMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitMessageResponseWithDefaults

`func NewSubmitMessageResponseWithDefaults() *SubmitMessageResponse`

NewSubmitMessageResponseWithDefaults instantiates a new SubmitMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailMerge

`func (o *SubmitMessageResponse) GetMailMerge() []BulkResponseEntry`

GetMailMerge returns the MailMerge field if non-nil, zero value otherwise.

### GetMailMergeOk

`func (o *SubmitMessageResponse) GetMailMergeOk() (*[]BulkResponseEntry, bool)`

GetMailMergeOk returns a tuple with the MailMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailMerge

`func (o *SubmitMessageResponse) SetMailMerge(v []BulkResponseEntry)`

SetMailMerge sets MailMerge field to given value.

### HasMailMerge

`func (o *SubmitMessageResponse) HasMailMerge() bool`

HasMailMerge returns a boolean if a field has been set.

### GetMessageId

`func (o *SubmitMessageResponse) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *SubmitMessageResponse) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *SubmitMessageResponse) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *SubmitMessageResponse) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetPreview

`func (o *SubmitMessageResponse) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *SubmitMessageResponse) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *SubmitMessageResponse) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *SubmitMessageResponse) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetQueueId

`func (o *SubmitMessageResponse) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *SubmitMessageResponse) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *SubmitMessageResponse) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *SubmitMessageResponse) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetReference

`func (o *SubmitMessageResponse) GetReference() Model11`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SubmitMessageResponse) GetReferenceOk() (*Model11, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SubmitMessageResponse) SetReference(v Model11)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SubmitMessageResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetResponse

`func (o *SubmitMessageResponse) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *SubmitMessageResponse) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *SubmitMessageResponse) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *SubmitMessageResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSendAt

`func (o *SubmitMessageResponse) GetSendAt() string`

GetSendAt returns the SendAt field if non-nil, zero value otherwise.

### GetSendAtOk

`func (o *SubmitMessageResponse) GetSendAtOk() (*string, bool)`

GetSendAtOk returns a tuple with the SendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAt

`func (o *SubmitMessageResponse) SetSendAt(v string)`

SetSendAt sets SendAt field to given value.

### HasSendAt

`func (o *SubmitMessageResponse) HasSendAt() bool`

HasSendAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


