# Model11

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | An error message if referenced message processing failed | [optional] 
**Message** | **string** | Referenced message ID | 
**Success** | Pointer to **bool** | Was the referenced message processed successfully | [optional] 

## Methods

### NewModel11

`func NewModel11(message string, ) *Model11`

NewModel11 instantiates a new Model11 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel11WithDefaults

`func NewModel11WithDefaults() *Model11`

NewModel11WithDefaults instantiates a new Model11 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *Model11) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Model11) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Model11) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Model11) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *Model11) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Model11) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Model11) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSuccess

`func (o *Model11) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Model11) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Model11) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *Model11) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


