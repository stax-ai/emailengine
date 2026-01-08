# Imap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Error code. Only present if success&#x3D;false | [optional] 
**Error** | Pointer to **string** | Error messages for IMAP verification. Only present if success&#x3D;false | [optional] 
**Success** | Pointer to **bool** | Was IMAP account verified | [optional] 

## Methods

### NewImap

`func NewImap() *Imap`

NewImap instantiates a new Imap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImapWithDefaults

`func NewImapWithDefaults() *Imap`

NewImapWithDefaults instantiates a new Imap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Imap) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Imap) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Imap) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Imap) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetError

`func (o *Imap) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Imap) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Imap) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Imap) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSuccess

`func (o *Imap) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Imap) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Imap) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *Imap) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


