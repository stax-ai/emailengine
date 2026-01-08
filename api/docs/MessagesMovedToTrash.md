# MessagesMovedToTrash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | **string** | Trash folder path | 
**EmailIds** | Pointer to **[]string** | An optional list of emailId values, if the server supports unique email IDs | [optional] 
**IdMap** | Pointer to **[][]string** | An optional map of source and target ID values, if the server provided this info | [optional] 

## Methods

### NewMessagesMovedToTrash

`func NewMessagesMovedToTrash(destination string, ) *MessagesMovedToTrash`

NewMessagesMovedToTrash instantiates a new MessagesMovedToTrash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesMovedToTrashWithDefaults

`func NewMessagesMovedToTrashWithDefaults() *MessagesMovedToTrash`

NewMessagesMovedToTrashWithDefaults instantiates a new MessagesMovedToTrash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *MessagesMovedToTrash) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *MessagesMovedToTrash) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *MessagesMovedToTrash) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetEmailIds

`func (o *MessagesMovedToTrash) GetEmailIds() []string`

GetEmailIds returns the EmailIds field if non-nil, zero value otherwise.

### GetEmailIdsOk

`func (o *MessagesMovedToTrash) GetEmailIdsOk() (*[]string, bool)`

GetEmailIdsOk returns a tuple with the EmailIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailIds

`func (o *MessagesMovedToTrash) SetEmailIds(v []string)`

SetEmailIds sets EmailIds field to given value.

### HasEmailIds

`func (o *MessagesMovedToTrash) HasEmailIds() bool`

HasEmailIds returns a boolean if a field has been set.

### GetIdMap

`func (o *MessagesMovedToTrash) GetIdMap() [][]string`

GetIdMap returns the IdMap field if non-nil, zero value otherwise.

### GetIdMapOk

`func (o *MessagesMovedToTrash) GetIdMapOk() (*[][]string, bool)`

GetIdMapOk returns a tuple with the IdMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdMap

`func (o *MessagesMovedToTrash) SetIdMap(v [][]string)`

SetIdMap sets IdMap field to given value.

### HasIdMap

`func (o *MessagesMovedToTrash) HasIdMap() bool`

HasIdMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


