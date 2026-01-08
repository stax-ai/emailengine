# MessagesMoveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailIds** | Pointer to **[]string** | An optional list of emailId values, if the server supports unique email IDs | [optional] 
**IdMap** | Pointer to **[][]string** | An optional map of source and target ID values, if the server provided this info | [optional] 
**Path** | **string** | Target mailbox folder path | 

## Methods

### NewMessagesMoveResponse

`func NewMessagesMoveResponse(path string, ) *MessagesMoveResponse`

NewMessagesMoveResponse instantiates a new MessagesMoveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesMoveResponseWithDefaults

`func NewMessagesMoveResponseWithDefaults() *MessagesMoveResponse`

NewMessagesMoveResponseWithDefaults instantiates a new MessagesMoveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailIds

`func (o *MessagesMoveResponse) GetEmailIds() []string`

GetEmailIds returns the EmailIds field if non-nil, zero value otherwise.

### GetEmailIdsOk

`func (o *MessagesMoveResponse) GetEmailIdsOk() (*[]string, bool)`

GetEmailIdsOk returns a tuple with the EmailIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailIds

`func (o *MessagesMoveResponse) SetEmailIds(v []string)`

SetEmailIds sets EmailIds field to given value.

### HasEmailIds

`func (o *MessagesMoveResponse) HasEmailIds() bool`

HasEmailIds returns a boolean if a field has been set.

### GetIdMap

`func (o *MessagesMoveResponse) GetIdMap() [][]string`

GetIdMap returns the IdMap field if non-nil, zero value otherwise.

### GetIdMapOk

`func (o *MessagesMoveResponse) GetIdMapOk() (*[][]string, bool)`

GetIdMapOk returns a tuple with the IdMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdMap

`func (o *MessagesMoveResponse) SetIdMap(v [][]string)`

SetIdMap sets IdMap field to given value.

### HasIdMap

`func (o *MessagesMoveResponse) HasIdMap() bool`

HasIdMap returns a boolean if a field has been set.

### GetPath

`func (o *MessagesMoveResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MessagesMoveResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MessagesMoveResponse) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


