# SettingsQueueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**QueueJobs**](QueueJobs.md) |  | [optional] 
**Paused** | Pointer to **bool** | Is the queue paused or not | [optional] 
**Queue** | [**Queue**](Queue.md) |  | 

## Methods

### NewSettingsQueueResponse

`func NewSettingsQueueResponse(queue Queue, ) *SettingsQueueResponse`

NewSettingsQueueResponse instantiates a new SettingsQueueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsQueueResponseWithDefaults

`func NewSettingsQueueResponseWithDefaults() *SettingsQueueResponse`

NewSettingsQueueResponseWithDefaults instantiates a new SettingsQueueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *SettingsQueueResponse) GetJobs() QueueJobs`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *SettingsQueueResponse) GetJobsOk() (*QueueJobs, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *SettingsQueueResponse) SetJobs(v QueueJobs)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *SettingsQueueResponse) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetPaused

`func (o *SettingsQueueResponse) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *SettingsQueueResponse) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *SettingsQueueResponse) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *SettingsQueueResponse) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetQueue

`func (o *SettingsQueueResponse) GetQueue() Queue`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *SettingsQueueResponse) GetQueueOk() (*Queue, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *SettingsQueueResponse) SetQueue(v Queue)`

SetQueue sets Queue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


