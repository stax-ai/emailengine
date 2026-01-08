# SMTPInfoStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Error type identifier | [optional] 
**Command** | Pointer to **string** | SMTP command that failed | [optional] 
**Created** | Pointer to **NullableTime** | When was the status for SMTP connection last updated | [optional] 
**Description** | Pointer to **string** | Error information | [optional] 
**Response** | Pointer to **string** | SMTP response message for delivery attempt | [optional] 
**ResponseCode** | Pointer to **int32** | Error status code | [optional] 
**Status** | Pointer to [**SMTPStatusStatus**](SMTPStatusStatus.md) |  | [optional] 

## Methods

### NewSMTPInfoStatus

`func NewSMTPInfoStatus() *SMTPInfoStatus`

NewSMTPInfoStatus instantiates a new SMTPInfoStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMTPInfoStatusWithDefaults

`func NewSMTPInfoStatusWithDefaults() *SMTPInfoStatus`

NewSMTPInfoStatusWithDefaults instantiates a new SMTPInfoStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SMTPInfoStatus) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SMTPInfoStatus) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SMTPInfoStatus) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SMTPInfoStatus) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCommand

`func (o *SMTPInfoStatus) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *SMTPInfoStatus) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *SMTPInfoStatus) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *SMTPInfoStatus) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetCreated

`func (o *SMTPInfoStatus) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SMTPInfoStatus) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SMTPInfoStatus) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SMTPInfoStatus) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *SMTPInfoStatus) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *SMTPInfoStatus) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetDescription

`func (o *SMTPInfoStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SMTPInfoStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SMTPInfoStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SMTPInfoStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResponse

`func (o *SMTPInfoStatus) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *SMTPInfoStatus) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *SMTPInfoStatus) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *SMTPInfoStatus) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetResponseCode

`func (o *SMTPInfoStatus) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *SMTPInfoStatus) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *SMTPInfoStatus) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *SMTPInfoStatus) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### GetStatus

`func (o *SMTPInfoStatus) GetStatus() SMTPStatusStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SMTPInfoStatus) GetStatusOk() (*SMTPStatusStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SMTPInfoStatus) SetStatus(v SMTPStatusStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SMTPInfoStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


