# ModifyMailboxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPath** | Pointer to **string** | Full path to mailbox if renamed | [optional] 
**Path** | **string** | Mailbox folder path | 
**Renamed** | Pointer to **bool** | Was the mailbox renamed | [optional] 
**Subscribed** | Pointer to **bool** | Subscription status after modification | [optional] 

## Methods

### NewModifyMailboxResponse

`func NewModifyMailboxResponse(path string, ) *ModifyMailboxResponse`

NewModifyMailboxResponse instantiates a new ModifyMailboxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyMailboxResponseWithDefaults

`func NewModifyMailboxResponseWithDefaults() *ModifyMailboxResponse`

NewModifyMailboxResponseWithDefaults instantiates a new ModifyMailboxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPath

`func (o *ModifyMailboxResponse) GetNewPath() string`

GetNewPath returns the NewPath field if non-nil, zero value otherwise.

### GetNewPathOk

`func (o *ModifyMailboxResponse) GetNewPathOk() (*string, bool)`

GetNewPathOk returns a tuple with the NewPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPath

`func (o *ModifyMailboxResponse) SetNewPath(v string)`

SetNewPath sets NewPath field to given value.

### HasNewPath

`func (o *ModifyMailboxResponse) HasNewPath() bool`

HasNewPath returns a boolean if a field has been set.

### GetPath

`func (o *ModifyMailboxResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ModifyMailboxResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ModifyMailboxResponse) SetPath(v string)`

SetPath sets Path field to given value.


### GetRenamed

`func (o *ModifyMailboxResponse) GetRenamed() bool`

GetRenamed returns the Renamed field if non-nil, zero value otherwise.

### GetRenamedOk

`func (o *ModifyMailboxResponse) GetRenamedOk() (*bool, bool)`

GetRenamedOk returns a tuple with the Renamed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenamed

`func (o *ModifyMailboxResponse) SetRenamed(v bool)`

SetRenamed sets Renamed field to given value.

### HasRenamed

`func (o *ModifyMailboxResponse) HasRenamed() bool`

HasRenamed returns a boolean if a field has been set.

### GetSubscribed

`func (o *ModifyMailboxResponse) GetSubscribed() bool`

GetSubscribed returns the Subscribed field if non-nil, zero value otherwise.

### GetSubscribedOk

`func (o *ModifyMailboxResponse) GetSubscribedOk() (*bool, bool)`

GetSubscribedOk returns a tuple with the Subscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribed

`func (o *ModifyMailboxResponse) SetSubscribed(v bool)`

SetSubscribed sets Subscribed field to given value.

### HasSubscribed

`func (o *ModifyMailboxResponse) HasSubscribed() bool`

HasSubscribed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


