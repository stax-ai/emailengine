# QueueJobs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int32** | Jobs that are currently being processed | [optional] 
**Delayed** | Pointer to **int32** | Jobs that are processed in the future | [optional] 
**Paused** | Pointer to **int32** | Jobs that would be processed once queue processing is resumed | [optional] 
**Waiting** | Pointer to **int32** | Jobs that should be processed, but are waiting until there are any free handlers | [optional] 

## Methods

### NewQueueJobs

`func NewQueueJobs() *QueueJobs`

NewQueueJobs instantiates a new QueueJobs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueJobsWithDefaults

`func NewQueueJobsWithDefaults() *QueueJobs`

NewQueueJobsWithDefaults instantiates a new QueueJobs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *QueueJobs) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *QueueJobs) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *QueueJobs) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *QueueJobs) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDelayed

`func (o *QueueJobs) GetDelayed() int32`

GetDelayed returns the Delayed field if non-nil, zero value otherwise.

### GetDelayedOk

`func (o *QueueJobs) GetDelayedOk() (*int32, bool)`

GetDelayedOk returns a tuple with the Delayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayed

`func (o *QueueJobs) SetDelayed(v int32)`

SetDelayed sets Delayed field to given value.

### HasDelayed

`func (o *QueueJobs) HasDelayed() bool`

HasDelayed returns a boolean if a field has been set.

### GetPaused

`func (o *QueueJobs) GetPaused() int32`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *QueueJobs) GetPausedOk() (*int32, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *QueueJobs) SetPaused(v int32)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *QueueJobs) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetWaiting

`func (o *QueueJobs) GetWaiting() int32`

GetWaiting returns the Waiting field if non-nil, zero value otherwise.

### GetWaitingOk

`func (o *QueueJobs) GetWaitingOk() (*int32, bool)`

GetWaitingOk returns a tuple with the Waiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiting

`func (o *QueueJobs) SetWaiting(v int32)`

SetWaiting sets Waiting field to given value.

### HasWaiting

`func (o *QueueJobs) HasWaiting() bool`

HasWaiting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


