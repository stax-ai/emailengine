# SettingsPutQueueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paused** | Pointer to **bool** | Is the queue paused or not | [optional] 
**Queue** | [**Queue**](Queue.md) |  | 

## Methods

### NewSettingsPutQueueResponse

`func NewSettingsPutQueueResponse(queue Queue, ) *SettingsPutQueueResponse`

NewSettingsPutQueueResponse instantiates a new SettingsPutQueueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsPutQueueResponseWithDefaults

`func NewSettingsPutQueueResponseWithDefaults() *SettingsPutQueueResponse`

NewSettingsPutQueueResponseWithDefaults instantiates a new SettingsPutQueueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaused

`func (o *SettingsPutQueueResponse) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *SettingsPutQueueResponse) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *SettingsPutQueueResponse) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *SettingsPutQueueResponse) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetQueue

`func (o *SettingsPutQueueResponse) GetQueue() Queue`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *SettingsPutQueueResponse) GetQueueOk() (*Queue, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *SettingsPutQueueResponse) SetQueue(v Queue)`

SetQueue sets Queue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


