# MessagesDeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **bool** | Was the delete action executed | [optional] 
**Moved** | Pointer to [**MessagesMovedToTrash**](MessagesMovedToTrash.md) |  | [optional] 

## Methods

### NewMessagesDeleteResponse

`func NewMessagesDeleteResponse() *MessagesDeleteResponse`

NewMessagesDeleteResponse instantiates a new MessagesDeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesDeleteResponseWithDefaults

`func NewMessagesDeleteResponseWithDefaults() *MessagesDeleteResponse`

NewMessagesDeleteResponseWithDefaults instantiates a new MessagesDeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *MessagesDeleteResponse) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *MessagesDeleteResponse) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *MessagesDeleteResponse) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *MessagesDeleteResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetMoved

`func (o *MessagesDeleteResponse) GetMoved() MessagesMovedToTrash`

GetMoved returns the Moved field if non-nil, zero value otherwise.

### GetMovedOk

`func (o *MessagesDeleteResponse) GetMovedOk() (*MessagesMovedToTrash, bool)`

GetMovedOk returns a tuple with the Moved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoved

`func (o *MessagesDeleteResponse) SetMoved(v MessagesMovedToTrash)`

SetMoved sets Moved field to given value.

### HasMoved

`func (o *MessagesDeleteResponse) HasMoved() bool`

HasMoved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


