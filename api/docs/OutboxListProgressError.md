# OutboxListProgressError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Error code | [optional] 
**Message** | Pointer to **string** | Error description | [optional] 
**StatusCode** | Pointer to **string** | SMTP response code | [optional] 

## Methods

### NewOutboxListProgressError

`func NewOutboxListProgressError() *OutboxListProgressError`

NewOutboxListProgressError instantiates a new OutboxListProgressError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboxListProgressErrorWithDefaults

`func NewOutboxListProgressErrorWithDefaults() *OutboxListProgressError`

NewOutboxListProgressErrorWithDefaults instantiates a new OutboxListProgressError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *OutboxListProgressError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OutboxListProgressError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OutboxListProgressError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *OutboxListProgressError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *OutboxListProgressError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OutboxListProgressError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OutboxListProgressError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OutboxListProgressError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatusCode

`func (o *OutboxListProgressError) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *OutboxListProgressError) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *OutboxListProgressError) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *OutboxListProgressError) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


