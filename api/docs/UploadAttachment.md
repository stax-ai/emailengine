# UploadAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | Pointer to **string** | Content-ID value for embedded images | [optional] 
**Content** | **string** | Base64 formatted attachment file | 
**ContentDisposition** | Pointer to [**ContentDisposition**](ContentDisposition.md) |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Encoding** | Pointer to [**Encoding**](Encoding.md) |  | [optional] [default to ENCODING_BASE64]
**Filename** | Pointer to **string** |  | [optional] 
**Reference** | Pointer to [**NullableReference**](Reference.md) |  | [optional] 

## Methods

### NewUploadAttachment

`func NewUploadAttachment(content string, ) *UploadAttachment`

NewUploadAttachment instantiates a new UploadAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadAttachmentWithDefaults

`func NewUploadAttachmentWithDefaults() *UploadAttachment`

NewUploadAttachmentWithDefaults instantiates a new UploadAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *UploadAttachment) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *UploadAttachment) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *UploadAttachment) SetCid(v string)`

SetCid sets Cid field to given value.

### HasCid

`func (o *UploadAttachment) HasCid() bool`

HasCid returns a boolean if a field has been set.

### GetContent

`func (o *UploadAttachment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UploadAttachment) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UploadAttachment) SetContent(v string)`

SetContent sets Content field to given value.


### GetContentDisposition

`func (o *UploadAttachment) GetContentDisposition() ContentDisposition`

GetContentDisposition returns the ContentDisposition field if non-nil, zero value otherwise.

### GetContentDispositionOk

`func (o *UploadAttachment) GetContentDispositionOk() (*ContentDisposition, bool)`

GetContentDispositionOk returns a tuple with the ContentDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentDisposition

`func (o *UploadAttachment) SetContentDisposition(v ContentDisposition)`

SetContentDisposition sets ContentDisposition field to given value.

### HasContentDisposition

`func (o *UploadAttachment) HasContentDisposition() bool`

HasContentDisposition returns a boolean if a field has been set.

### GetContentType

`func (o *UploadAttachment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *UploadAttachment) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *UploadAttachment) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *UploadAttachment) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetEncoding

`func (o *UploadAttachment) GetEncoding() Encoding`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *UploadAttachment) GetEncodingOk() (*Encoding, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *UploadAttachment) SetEncoding(v Encoding)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *UploadAttachment) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetFilename

`func (o *UploadAttachment) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *UploadAttachment) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *UploadAttachment) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *UploadAttachment) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetReference

`func (o *UploadAttachment) GetReference() Reference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *UploadAttachment) GetReferenceOk() (*Reference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *UploadAttachment) SetReference(v Reference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *UploadAttachment) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *UploadAttachment) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *UploadAttachment) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


