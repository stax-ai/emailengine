# MessageMove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Destination mailbox folder path | 
**Source** | Pointer to **string** | Source mailbox folder path (Gmail API only). Needed to remove the label from the message. | [optional] 

## Methods

### NewMessageMove

`func NewMessageMove(path string, ) *MessageMove`

NewMessageMove instantiates a new MessageMove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageMoveWithDefaults

`func NewMessageMoveWithDefaults() *MessageMove`

NewMessageMoveWithDefaults instantiates a new MessageMove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *MessageMove) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MessageMove) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MessageMove) SetPath(v string)`

SetPath sets Path field to given value.


### GetSource

`func (o *MessageMove) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MessageMove) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MessageMove) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MessageMove) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


