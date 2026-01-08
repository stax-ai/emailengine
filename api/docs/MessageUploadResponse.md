# MessageUploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the message. NB! This and other fields might not be present if server did not provide enough information | [optional] 
**MessageId** | Pointer to **string** | Message ID | [optional] 
**Path** | Pointer to **string** | Folder this message was uploaded to | [optional] 
**Reference** | Pointer to [**ResponseReference**](ResponseReference.md) |  | [optional] 
**Seq** | Pointer to **int32** | Sequence number of uploaded message | [optional] 
**Uid** | Pointer to **int32** | UID of uploaded message | [optional] 
**UidValidity** | Pointer to **string** | UIDVALIDITY of the target folder. Numeric value cast as string. | [optional] 

## Methods

### NewMessageUploadResponse

`func NewMessageUploadResponse() *MessageUploadResponse`

NewMessageUploadResponse instantiates a new MessageUploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageUploadResponseWithDefaults

`func NewMessageUploadResponseWithDefaults() *MessageUploadResponse`

NewMessageUploadResponseWithDefaults instantiates a new MessageUploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageUploadResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageUploadResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageUploadResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MessageUploadResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessageId

`func (o *MessageUploadResponse) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageUploadResponse) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageUploadResponse) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *MessageUploadResponse) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetPath

`func (o *MessageUploadResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MessageUploadResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MessageUploadResponse) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MessageUploadResponse) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetReference

`func (o *MessageUploadResponse) GetReference() ResponseReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *MessageUploadResponse) GetReferenceOk() (*ResponseReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *MessageUploadResponse) SetReference(v ResponseReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *MessageUploadResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSeq

`func (o *MessageUploadResponse) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *MessageUploadResponse) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *MessageUploadResponse) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *MessageUploadResponse) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetUid

`func (o *MessageUploadResponse) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MessageUploadResponse) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MessageUploadResponse) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MessageUploadResponse) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUidValidity

`func (o *MessageUploadResponse) GetUidValidity() string`

GetUidValidity returns the UidValidity field if non-nil, zero value otherwise.

### GetUidValidityOk

`func (o *MessageUploadResponse) GetUidValidityOk() (*string, bool)`

GetUidValidityOk returns a tuple with the UidValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidValidity

`func (o *MessageUploadResponse) SetUidValidity(v string)`

SetUidValidity sets UidValidity field to given value.

### HasUidValidity

`func (o *MessageUploadResponse) HasUidValidity() bool`

HasUidValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


