# OutboxEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Unique identifier for the email account | 
**Attempts** | Pointer to **int32** | Maximum delivery attempts before marking as failed | [optional] 
**AttemptsMade** | Pointer to **int32** | Number of delivery attempts made | [optional] 
**Created** | Pointer to **time.Time** | When this message was added to the queue | [optional] 
**Envelope** | Pointer to [**Envelope**](Envelope.md) |  | [optional] 
**MessageId** | Pointer to **string** | Message-ID header value | [optional] 
**NextAttempt** | Pointer to **time.Time** | Next delivery attempt time | [optional] 
**Progress** | Pointer to [**OutboxEntryProgress**](OutboxEntryProgress.md) |  | [optional] 
**QueueId** | Pointer to **string** | Unique queue entry identifier | [optional] 
**Scheduled** | Pointer to **time.Time** | Scheduled delivery time | [optional] 
**Source** | Pointer to [**Source**](Source.md) |  | [optional] 
**Subject** | Pointer to **string** | Email subject line | [optional] 

## Methods

### NewOutboxEntry

`func NewOutboxEntry(account string, ) *OutboxEntry`

NewOutboxEntry instantiates a new OutboxEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboxEntryWithDefaults

`func NewOutboxEntryWithDefaults() *OutboxEntry`

NewOutboxEntryWithDefaults instantiates a new OutboxEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *OutboxEntry) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *OutboxEntry) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *OutboxEntry) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetAttempts

`func (o *OutboxEntry) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *OutboxEntry) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *OutboxEntry) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *OutboxEntry) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetAttemptsMade

`func (o *OutboxEntry) GetAttemptsMade() int32`

GetAttemptsMade returns the AttemptsMade field if non-nil, zero value otherwise.

### GetAttemptsMadeOk

`func (o *OutboxEntry) GetAttemptsMadeOk() (*int32, bool)`

GetAttemptsMadeOk returns a tuple with the AttemptsMade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptsMade

`func (o *OutboxEntry) SetAttemptsMade(v int32)`

SetAttemptsMade sets AttemptsMade field to given value.

### HasAttemptsMade

`func (o *OutboxEntry) HasAttemptsMade() bool`

HasAttemptsMade returns a boolean if a field has been set.

### GetCreated

`func (o *OutboxEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OutboxEntry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OutboxEntry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OutboxEntry) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEnvelope

`func (o *OutboxEntry) GetEnvelope() Envelope`

GetEnvelope returns the Envelope field if non-nil, zero value otherwise.

### GetEnvelopeOk

`func (o *OutboxEntry) GetEnvelopeOk() (*Envelope, bool)`

GetEnvelopeOk returns a tuple with the Envelope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvelope

`func (o *OutboxEntry) SetEnvelope(v Envelope)`

SetEnvelope sets Envelope field to given value.

### HasEnvelope

`func (o *OutboxEntry) HasEnvelope() bool`

HasEnvelope returns a boolean if a field has been set.

### GetMessageId

`func (o *OutboxEntry) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *OutboxEntry) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *OutboxEntry) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *OutboxEntry) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetNextAttempt

`func (o *OutboxEntry) GetNextAttempt() time.Time`

GetNextAttempt returns the NextAttempt field if non-nil, zero value otherwise.

### GetNextAttemptOk

`func (o *OutboxEntry) GetNextAttemptOk() (*time.Time, bool)`

GetNextAttemptOk returns a tuple with the NextAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAttempt

`func (o *OutboxEntry) SetNextAttempt(v time.Time)`

SetNextAttempt sets NextAttempt field to given value.

### HasNextAttempt

`func (o *OutboxEntry) HasNextAttempt() bool`

HasNextAttempt returns a boolean if a field has been set.

### GetProgress

`func (o *OutboxEntry) GetProgress() OutboxEntryProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *OutboxEntry) GetProgressOk() (*OutboxEntryProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *OutboxEntry) SetProgress(v OutboxEntryProgress)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *OutboxEntry) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetQueueId

`func (o *OutboxEntry) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *OutboxEntry) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *OutboxEntry) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *OutboxEntry) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetScheduled

`func (o *OutboxEntry) GetScheduled() time.Time`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *OutboxEntry) GetScheduledOk() (*time.Time, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *OutboxEntry) SetScheduled(v time.Time)`

SetScheduled sets Scheduled field to given value.

### HasScheduled

`func (o *OutboxEntry) HasScheduled() bool`

HasScheduled returns a boolean if a field has been set.

### GetSource

`func (o *OutboxEntry) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OutboxEntry) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OutboxEntry) SetSource(v Source)`

SetSource sets Source field to given value.

### HasSource

`func (o *OutboxEntry) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSubject

`func (o *OutboxEntry) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *OutboxEntry) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *OutboxEntry) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *OutboxEntry) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


