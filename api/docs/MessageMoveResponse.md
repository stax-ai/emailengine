# MessageMoveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the moved message. Only included if the server provides it. | [optional] 
**Path** | **string** | Destination mailbox folder path | 
**Uid** | Pointer to **int32** | UID of the moved message, applies only to IMAP accounts. Only included if the server provides it. | [optional] 

## Methods

### NewMessageMoveResponse

`func NewMessageMoveResponse(path string, ) *MessageMoveResponse`

NewMessageMoveResponse instantiates a new MessageMoveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageMoveResponseWithDefaults

`func NewMessageMoveResponseWithDefaults() *MessageMoveResponse`

NewMessageMoveResponseWithDefaults instantiates a new MessageMoveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageMoveResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageMoveResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageMoveResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MessageMoveResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *MessageMoveResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MessageMoveResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MessageMoveResponse) SetPath(v string)`

SetPath sets Path field to given value.


### GetUid

`func (o *MessageMoveResponse) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MessageMoveResponse) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MessageMoveResponse) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MessageMoveResponse) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


