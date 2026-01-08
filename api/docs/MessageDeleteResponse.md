# MessageDeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **bool** | Was the delete action executed | [optional] 
**Moved** | Pointer to [**Moved**](Moved.md) |  | [optional] 

## Methods

### NewMessageDeleteResponse

`func NewMessageDeleteResponse() *MessageDeleteResponse`

NewMessageDeleteResponse instantiates a new MessageDeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDeleteResponseWithDefaults

`func NewMessageDeleteResponseWithDefaults() *MessageDeleteResponse`

NewMessageDeleteResponseWithDefaults instantiates a new MessageDeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *MessageDeleteResponse) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *MessageDeleteResponse) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *MessageDeleteResponse) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *MessageDeleteResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetMoved

`func (o *MessageDeleteResponse) GetMoved() Moved`

GetMoved returns the Moved field if non-nil, zero value otherwise.

### GetMovedOk

`func (o *MessageDeleteResponse) GetMovedOk() (*Moved, bool)`

GetMovedOk returns a tuple with the Moved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoved

`func (o *MessageDeleteResponse) SetMoved(v Moved)`

SetMoved sets Moved field to given value.

### HasMoved

`func (o *MessageDeleteResponse) HasMoved() bool`

HasMoved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


