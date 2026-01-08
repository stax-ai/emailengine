# Smtp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Error code. Only present if success&#x3D;false | [optional] 
**Error** | Pointer to **string** | Error messages for SMTP verification. Only present if success&#x3D;false | [optional] 
**Success** | Pointer to **bool** | Was SMTP account verified | [optional] 

## Methods

### NewSmtp

`func NewSmtp() *Smtp`

NewSmtp instantiates a new Smtp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpWithDefaults

`func NewSmtpWithDefaults() *Smtp`

NewSmtpWithDefaults instantiates a new Smtp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Smtp) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Smtp) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Smtp) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Smtp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetError

`func (o *Smtp) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Smtp) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Smtp) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Smtp) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSuccess

`func (o *Smtp) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Smtp) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Smtp) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *Smtp) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


