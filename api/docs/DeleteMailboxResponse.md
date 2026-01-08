# DeleteMailboxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **bool** | Was the mailbox deleted | [optional] 
**Path** | **string** | Full path to mailbox | 

## Methods

### NewDeleteMailboxResponse

`func NewDeleteMailboxResponse(path string, ) *DeleteMailboxResponse`

NewDeleteMailboxResponse instantiates a new DeleteMailboxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteMailboxResponseWithDefaults

`func NewDeleteMailboxResponseWithDefaults() *DeleteMailboxResponse`

NewDeleteMailboxResponseWithDefaults instantiates a new DeleteMailboxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *DeleteMailboxResponse) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DeleteMailboxResponse) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DeleteMailboxResponse) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DeleteMailboxResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetPath

`func (o *DeleteMailboxResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DeleteMailboxResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DeleteMailboxResponse) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


