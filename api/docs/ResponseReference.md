# ResponseReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | An error message if referenced message processing failed | [optional] 
**Message** | **string** | Referenced message ID | 
**Success** | Pointer to **bool** | Was the referenced message processed | [optional] 

## Methods

### NewResponseReference

`func NewResponseReference(message string, ) *ResponseReference`

NewResponseReference instantiates a new ResponseReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseReferenceWithDefaults

`func NewResponseReferenceWithDefaults() *ResponseReference`

NewResponseReferenceWithDefaults instantiates a new ResponseReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ResponseReference) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResponseReference) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResponseReference) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ResponseReference) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseReference) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseReference) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseReference) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSuccess

`func (o *ResponseReference) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResponseReference) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResponseReference) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ResponseReference) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


